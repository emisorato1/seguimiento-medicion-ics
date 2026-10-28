package web

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// Escenario: Abrir el dashboard desde la barra de navegacion
// Cubre: RN-18
func TestAbrirElDashboardDesdeLaBarraDeNavegacion(t *testing.T) {
	registry := domain.NewProjectRegistry()
	addProject(t, registry, "Sistema de metricas")
	server := newTestServer(t, registry)
	enlace := `<a href="/proyectos/1/dashboard"`

	for _, desde := range []string{"/proyectos/1", "/proyectos/1/backlog", "/proyectos/1/sprints", "/proyectos/1/dashboard"} {
		t.Run(desde, func(t *testing.T) {
			res := get(t, server, desde)
			if res.Code != http.StatusOK {
				t.Fatalf("se esperaba %s con el estado %d, salio %d", desde, http.StatusOK, res.Code)
			}
			barra := between(res.Body.String(), `<nav aria-label="Navegación principal">`, "</nav>")
			if !strings.Contains(barra, enlace) || !strings.Contains(barra, ">Dashboard</a>") {
				t.Fatalf("se esperaba en la barra el acceso al dashboard, y no esta en:\n%s", barra)
			}
			// Va junto a los datos del proyecto: después de Datos y antes de Integrantes.
			datos := strings.Index(barra, `<a href="/proyectos/1"`)
			dashboard := strings.Index(barra, enlace)
			integrantes := strings.Index(barra, `<a href="/proyectos/1/integrantes"`)
			if !(datos < dashboard && dashboard < integrantes) {
				t.Errorf("se esperaba el dashboard entre Datos e Integrantes, y no esta asi en:\n%s", barra)
			}
		})
	}

	// En su propia página, el dashboard figura como la sección actual.
	barra := between(get(t, server, "/proyectos/1/dashboard").Body.String(), `<nav aria-label="Navegación principal">`, "</nav>")
	if !strings.Contains(barra, enlace+` aria-current="page">`) {
		t.Errorf("se esperaba el dashboard marcado como la seccion actual, y no esta en:\n%s", barra)
	}
}

// Escenario: Encabezado de un proyecto sin fecha de finalizacion
// Cubre: RN-3, RN-15
func TestEncabezadoDeUnProyectoSinFechaDeFinalizacion(t *testing.T) {
	registry := domain.NewProjectRegistry()
	if _, err := registry.Create(domain.ProjectData{Name: "Sistema de metricas", Start: time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC)}); err != nil {
		t.Fatalf("no se pudo preparar el proyecto: %v", err)
	}
	addSprint(t, registry, 1)
	server := newTestServer(t, registry)

	res := get(t, server, "/proyectos/1/dashboard")

	if res.Code != http.StatusOK {
		t.Fatalf("se esperaba el estado %d, salio %d", http.StatusOK, res.Code)
	}
	body := res.Body.String()
	if titulo := between(body, "<title>", "</title>"); !strings.Contains(titulo, "Dashboard de Sistema de metricas") {
		t.Errorf("se esperaba el titulo %q, salio %q", "Dashboard de Sistema de metricas", titulo)
	}
	contenido := between(body, "<main", "</main>")
	for _, want := range []string{"<h1>Sistema de metricas</h1>", "01/10/2026", "sin fecha de fin"} {
		if !strings.Contains(contenido, want) {
			t.Errorf("se esperaba %q en el encabezado, y no esta en:\n%s", want, contenido)
		}
	}
}

// Escenario: Un proyecto sin sprints avisa que no hay datos
// Cubre: RN-7
// El test del dominio tiene el mismo nombre, en su paquete (Fase 3).
func TestUnProyectoSinSprintsAvisaQueNoHayDatos(t *testing.T) {
	registry := domain.NewProjectRegistry()
	addProject(t, registry, "Sistema de metricas")
	server := newTestServer(t, registry)

	res := get(t, server, "/proyectos/1/dashboard")

	// No es un error: la página se muestra con su encabezado.
	if res.Code != http.StatusOK {
		t.Fatalf("se esperaba el estado %d, salio %d", http.StatusOK, res.Code)
	}
	contenido := between(res.Body.String(), "<main", "</main>")
	if !strings.Contains(contenido, "<h1>Sistema de metricas</h1>") {
		t.Errorf("se esperaba el encabezado del proyecto, y no esta en:\n%s", contenido)
	}
	aviso := between(contenido, `<div class="vacio">`, "</div>")
	if !strings.Contains(aviso, "Todavía no hay datos para mostrar.") || !strings.Contains(aviso, `href="/proyectos/1/sprints/nuevo"`) {
		t.Errorf("se esperaba el aviso de que no hay datos, con el acceso a crear el primer sprint, y no esta en:\n%s", contenido)
	}
	if strings.Contains(contenido, "<svg") || strings.Contains(contenido, "<figure") {
		t.Errorf("no se esperaba ningun grafico, y hay en:\n%s", contenido)
	}
}

// Desde la tarea 12, los tests de la página arman el proyecto con las operaciones de verdad
// y no con un dashboard armado a mano (desvío de la D10, en la Fase 4).

// dashboardProject crea "Sistema de metricas" con una historia por cada valor de points,
// estimada con esos puntos, y el sprint 1 cerrado con la historia 1 comprometida y
// completada.
func dashboardProject(t *testing.T, registry *domain.ProjectRegistry, points ...int) *domain.Project {
	t.Helper()
	project := addProject(t, registry, "Sistema de metricas")
	for i, p := range points {
		story := addStory(t, project, fmt.Sprintf("Historia %d", i+1), domain.PriorityHigh)
		if err := project.EstimateStory(story.ID(), p); err != nil {
			t.Fatalf("no se pudo estimar la historia %d: %v", story.ID(), err)
		}
	}
	addSprint(t, registry, project.ID())
	if err := project.AssignStory(1); err != nil {
		t.Fatalf("no se pudo asignar la historia 1: %v", err)
	}
	if err := project.CloseSprint([]int{1}, time.Date(2026, 10, 30, 0, 0, 0, 0, time.UTC)); err != nil {
		t.Fatalf("no se pudo cerrar el sprint 1: %v", err)
	}
	return project
}

// Escenario: El estado general muestra las metricas tal como se calcularon
// Cubre: RN-4, RN-15
func TestElEstadoGeneralMuestraLasMetricasTalComoSeCalcularon(t *testing.T) {
	registry := domain.NewProjectRegistry()
	project := dashboardProject(t, registry, 3, 5)
	server := newTestServer(t, registry)

	res := get(t, server, "/proyectos/1/dashboard")

	estado := between(res.Body.String(), `<section id="estado-general"`, "</section>")
	if !strings.Contains(estado, "Estado general") {
		t.Fatalf("se esperaba la parte del estado general, y no esta en:\n%s", res.Body.String())
	}
	// Cada número tal como lo entrega el dominio, con su rótulo y en su orden.
	last := 0
	for _, figure := range project.Dashboard().Overview {
		fila := "<dt>" + figure.Label + "</dt><dd>" + figure.Text + "</dd>"
		i := strings.Index(estado, fila)
		if i < last {
			t.Fatalf("se esperaba %q en el estado general, despues del anterior, y no esta asi en:\n%s", fila, estado)
		}
		last = i
	}
	// Y los que da la #16 con estos datos: 8 planificados, 3 completados en el sprint 1.
	for _, fila := range []string{"<dt>Story points planificados</dt><dd>8</dd>", "<dt>Story points completados</dt><dd>3</dd>", "<dt>Velocidad</dt><dd>3</dd>"} {
		if !strings.Contains(estado, fila) {
			t.Errorf("se esperaba %q en el estado general, y no esta en:\n%s", fila, estado)
		}
	}
}

// Escenario: El sprint en curso muestra su meta y lo que lleva registrado
// Cubre: RN-5, RN-15
func TestElSprintEnCursoMuestraSuMetaYLoQueLlevaRegistrado(t *testing.T) {
	registry := domain.NewProjectRegistry()
	project := dashboardProject(t, registry, 3, 13)
	if _, err := project.AddSprint(domain.SprintData{
		Goal:  "Tener el dashboard",
		Start: time.Date(2026, 10, 26, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2026, 10, 30, 0, 0, 0, 0, time.UTC),
	}); err != nil {
		t.Fatalf("no se pudo crear el sprint 2: %v", err)
	}
	if err := project.AssignStory(2); err != nil {
		t.Fatalf("no se pudo asignar la historia 2: %v", err)
	}
	server := newTestServer(t, registry)

	res := get(t, server, "/proyectos/1/dashboard")

	enCurso := between(res.Body.String(), `<section id="sprint-en-curso"`, "</section>")
	for _, want := range []string{"Sprint 2", "Tener el dashboard", "26/10/2026", "30/10/2026",
		"<dt>Story points planificados</dt><dd>13</dd>", "<dt>Horas reales</dt><dd>0,0 h</dd>", "<dt>Defectos detectados</dt><dd>0</dd>"} {
		if !strings.Contains(enCurso, want) {
			t.Errorf("se esperaba %q en el sprint en curso, y no esta en:\n%s", want, enCurso)
		}
	}
	for _, unwanted := range []string{"completados", "Velocidad"} {
		if strings.Contains(enCurso, unwanted) {
			t.Errorf("no se esperaba %q en el sprint en curso, y esta en:\n%s", unwanted, enCurso)
		}
	}
}

// Escenario: Un proyecto con todos sus sprints cerrados no tiene sprint en curso
// Cubre: RN-6
func TestUnProyectoConTodosSusSprintsCerradosNoTieneSprintEnCurso(t *testing.T) {
	registry := domain.NewProjectRegistry()
	project := dashboardProject(t, registry, 3, 5)
	addSprint(t, registry, project.ID())
	if err := project.CloseSprint(nil, time.Date(2026, 10, 30, 0, 0, 0, 0, time.UTC)); err != nil {
		t.Fatalf("no se pudo cerrar el sprint 2: %v", err)
	}
	server := newTestServer(t, registry)

	res := get(t, server, "/proyectos/1/dashboard")

	if res.Code != http.StatusOK {
		t.Fatalf("se esperaba el estado %d, salio %d", http.StatusOK, res.Code)
	}
	body := res.Body.String()
	if enCurso := between(body, `<section id="sprint-en-curso"`, "</section>"); !strings.Contains(enCurso, "No hay ningún sprint en curso") {
		t.Errorf("se esperaba el aviso de que no hay sprint en curso, y no esta en:\n%s", body)
	}
	if estado := between(body, `<section id="estado-general"`, "</section>"); !strings.Contains(estado, "<dt>Story points planificados</dt><dd>8</dd>") {
		t.Errorf("se esperaba igual el estado general, y no esta en:\n%s", body)
	}
}

// Escenario: Pedir el dashboard de un proyecto que no existe
// Cubre: RN-17
func TestPedirElDashboardDeUnProyectoQueNoExiste(t *testing.T) {
	server := newTestServer(t, domain.NewProjectRegistry())

	res := get(t, server, "/proyectos/99/dashboard")

	if res.Code != http.StatusNotFound {
		t.Fatalf("se esperaba el estado %d, salio %d", http.StatusNotFound, res.Code)
	}
	contenido := between(res.Body.String(), "<main", "</main>")
	if mensaje := template.HTMLEscapeString(domain.ErrProjectNotFound.Error()); !strings.Contains(contenido, mensaje) {
		t.Errorf("se esperaba el mensaje %q, y no esta en:\n%s", mensaje, contenido)
	}
	if !strings.Contains(contenido, `href="/proyectos"`) {
		t.Errorf("se esperaba el acceso para volver a la lista de proyectos, y no esta en:\n%s", contenido)
	}
}

// Escenario: El largo de cada barra es proporcional al valor mas grande del grafico
// Cubre: RN-11, RN-12, RN-13
//
// El largo es un porcentaje del ancho del gráfico con un decimal; escribirlo con punto en el
// SVG es de la tarea 13 (D3).
func TestElLargoDeCadaBarraEsProporcionalAlValorMasGrandeDelGrafico(t *testing.T) {
	cases := []struct {
		max, value float64
		want       float64
	}{
		{40, 40, 100},
		{40, 10, 25},
		{8, 5, 62.5},
		{50, 12.5, 25},
		{1000, 1, 1},
		{40, 0, 0}, // sin barra
		// Dos casos que el escenario no muestra. Con un decimal: 8 de 13 es 61,538… y va
		// 61,5, como en el gráfico de velocidad. Y un gráfico todo en cero no divide por
		// cero: sin barra (RN-14, que la tarea 13 prueba con el gráfico entero).
		{13, 8, 61.5},
		{0, 0, 0},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v de %v", c.value, c.max), func(t *testing.T) {
			if got := barWidth(c.value, c.max); got != c.want {
				t.Fatalf("se esperaba que %v de %v midiera %v %%, salio %v %%", c.value, c.max, c.want, got)
			}
		})
	}
}

// barrasDelGrafico devuelve el gráfico con ese título tal como sale en la página, y el ancho
// de cada barra que dibuja, en orden. Un valor sin barra no suma ancho.
func barrasDelGrafico(t *testing.T, html, title string) (string, []string) {
	t.Helper()
	chart := between(html, "<figcaption>"+title+"</figcaption>", "</figure>")
	if chart == "" {
		t.Fatalf("se esperaba el grafico %q, y no esta en:\n%s", title, html)
	}
	widths := []string{}
	for _, m := range regexp.MustCompile(`<rect[^>]* width="([^"]*)"`).FindAllStringSubmatch(chart, -1) {
		widths = append(widths, m[1])
	}
	return chart, widths
}

// sprintCerradoConPuntos arma un proyecto con un sprint cerrado que comprometió historias con
// esos puntos y completó las de los números que se le pasan.
func sprintCerradoConPuntos(t *testing.T, registry *domain.ProjectRegistry, points []int, completed []int) *domain.Project {
	t.Helper()
	project := addProject(t, registry, "Sistema de metricas")
	addSprint(t, registry, project.ID())
	for i, p := range points {
		story := addStory(t, project, fmt.Sprintf("Historia %d", i+1), domain.PriorityHigh)
		if err := project.EstimateStory(story.ID(), p); err != nil {
			t.Fatalf("no se pudo estimar la historia %d: %v", story.ID(), err)
		}
		if err := project.AssignStory(story.ID()); err != nil {
			t.Fatalf("no se pudo asignar la historia %d: %v", story.ID(), err)
		}
	}
	if err := project.CloseSprint(completed, time.Date(2026, 10, 30, 0, 0, 0, 0, time.UTC)); err != nil {
		t.Fatalf("no se pudo cerrar el sprint 1: %v", err)
	}
	return project
}

// Esquema del escenario: Las barras de un grafico se miden contra su valor mas grande
// Cubre: RN-10, RN-11, RN-12, RN-14, D3
//
// Con datos reales lo completado nunca supera a lo comprometido, así que el caso "mas de lo
// planificado" se prueba con el gráfico de horas: 20 h estimadas y 25 h reales (Fase 4).
func TestLasBarrasDeUnGraficoSeMidenContraSuValorMasGrande(t *testing.T) {
	cases := []struct {
		caso      string
		points    []int
		completed []int
		want      []string // el ancho de cada barra dibujada; un valor en cero no tiene
		texts     []string
	}{
		{"normal", []int{21, 8, 1, 8, 2}, []int{1, 2, 3}, []string{"100%", "75%"}, []string{"40", "30"}},
		{"valores iguales", []int{21, 13, 5, 1}, []int{1, 2, 3, 4}, []string{"100%", "100%"}, []string{"40", "40"}},
		{"un valor en cero", []int{3}, nil, []string{"100%"}, []string{"3", "0"}},
		{"todo el grafico en cero", nil, nil, []string{}, []string{"0", "0"}},
	}
	for _, c := range cases {
		t.Run(c.caso, func(t *testing.T) {
			registry := domain.NewProjectRegistry()
			sprintCerradoConPuntos(t, registry, c.points, c.completed)
			server := newTestServer(t, registry)

			res := get(t, server, "/proyectos/1/dashboard")

			if res.Code != http.StatusOK {
				t.Fatalf("se esperaba el estado %d, salio %d", http.StatusOK, res.Code)
			}
			chart, widths := barrasDelGrafico(t, res.Body.String(), "Story points")
			if strings.Join(widths, " ") != strings.Join(c.want, " ") {
				t.Errorf("se esperaban las barras %q, salieron %q en:\n%s", c.want, widths, chart)
			}
			// Cada barra muestra su rótulo y su valor, aunque el valor sea cero.
			for i, label := range []string{"Planificados", "Completados"} {
				for _, text := range []string{">" + label + "</text>", ">" + c.texts[i] + "</text>"} {
					if !strings.Contains(chart, text) {
						t.Errorf("se esperaba %q en el grafico, y no esta en:\n%s", text, chart)
					}
				}
			}
		})
	}

	t.Run("mas de lo planificado", func(t *testing.T) {
		registry := domain.NewProjectRegistry()
		project := addProject(t, registry, "Sistema de metricas")
		addMember(t, project, "Ana Perez", domain.RoleProductBuilder)
		addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
		addSprint(t, registry, project.ID())
		if err := project.EstimateHours(1, 20); err != nil {
			t.Fatalf("no se pudieron estimar las horas: %v", err)
		}
		for _, day := range []int{26, 27} {
			if _, err := project.LogEffort(domain.EffortData{StoryID: 1, MemberName: "Ana Perez", Activity: "Escribir tests",
				Hours: 12.5, Day: time.Date(2026, 10, day, 0, 0, 0, 0, time.UTC)}, testToday()); err != nil {
				t.Fatalf("no se pudo registrar el esfuerzo: %v", err)
			}
		}
		server := newTestServer(t, registry)

		chart, widths := barrasDelGrafico(t, get(t, server, "/proyectos/1/dashboard").Body.String(), "Horas")

		if want := []string{"80%", "100%"}; strings.Join(widths, " ") != strings.Join(want, " ") {
			t.Errorf("se esperaban las barras %q, salieron %q en:\n%s", want, widths, chart)
		}
		for _, text := range []string{">Estimadas</text>", ">20,0 h</text>", ">Reales</text>", ">25,0 h</text>"} {
			if !strings.Contains(chart, text) {
				t.Errorf("se esperaba %q en el grafico, y no esta en:\n%s", text, chart)
			}
		}
	})
}

// Escenario: El grafico de velocidad tiene una barra por sprint cerrado
// Cubre: RN-9, RN-11, D3
//
// El ancho va con punto en el atributo del SVG (61.5%), y el valor que se lee con la regla de
// formato.
func TestElGraficoDeVelocidadTieneUnaBarraPorSprintCerrado(t *testing.T) {
	registry := domain.NewProjectRegistry()
	project := addProject(t, registry, "Sistema de metricas")
	// Velocidades 8, 13 y 10: el 10 no está en la escala, así que el sprint 3 completa 8 y 2.
	for i, points := range [][]int{{8}, {13}, {8, 2}} {
		start := time.Date(2026, 10, 1+7*i, 0, 0, 0, 0, time.UTC)
		if _, err := project.AddSprint(domain.SprintData{Goal: "Avanzar", Start: start, End: start.AddDate(0, 0, 6)}); err != nil {
			t.Fatalf("no se pudo crear el sprint %d: %v", i+1, err)
		}
		completed := []int{}
		for _, p := range points {
			story := addStory(t, project, fmt.Sprintf("Historia de %d puntos", p), domain.PriorityHigh)
			if err := project.EstimateStory(story.ID(), p); err != nil {
				t.Fatalf("no se pudo estimar la historia %d: %v", story.ID(), err)
			}
			if err := project.AssignStory(story.ID()); err != nil {
				t.Fatalf("no se pudo asignar la historia %d: %v", story.ID(), err)
			}
			completed = append(completed, story.ID())
		}
		if err := project.CloseSprint(completed, start.AddDate(0, 0, 6)); err != nil {
			t.Fatalf("no se pudo cerrar el sprint %d: %v", i+1, err)
		}
	}
	server := newTestServer(t, registry)

	chart, widths := barrasDelGrafico(t, get(t, server, "/proyectos/1/dashboard").Body.String(), "Velocidad")

	if want := []string{"61.5%", "100%", "76.9%"}; strings.Join(widths, " ") != strings.Join(want, " ") {
		t.Errorf("se esperaban las barras %q, salieron %q en:\n%s", want, widths, chart)
	}
	last := 0
	for _, text := range []string{">Sprint 1</text>", ">8</text>", ">Sprint 2</text>", ">13</text>", ">Sprint 3</text>", ">10</text>"} {
		i := strings.Index(chart, text)
		if i < last {
			t.Fatalf("se esperaba %q en el grafico, despues del anterior, y no esta asi en:\n%s", text, chart)
		}
		last = i
	}
}

// Escenario: Todavia no hay sprints cerrados para el grafico de velocidad
// Cubre: RN-9
func TestTodaviaNoHaySprintsCerradosParaElGraficoDeVelocidad(t *testing.T) {
	registry := domain.NewProjectRegistry()
	addProject(t, registry, "Sistema de metricas")
	addSprint(t, registry, 1)
	server := newTestServer(t, registry)

	chart, widths := barrasDelGrafico(t, get(t, server, "/proyectos/1/dashboard").Body.String(), "Velocidad")

	if !strings.Contains(chart, "Todavía no hay sprints cerrados") {
		t.Errorf("se esperaba el aviso de que no hay sprints cerrados, y no esta en:\n%s", chart)
	}
	if len(widths) != 0 || strings.Contains(chart, "Sprint 1") {
		t.Errorf("se esperaba el grafico sin barras y sin el sprint 1, y salio:\n%s", chart)
	}
}

// Escenario: Dashboard de un proyecto con un sprint en curso
// Cubre: RN-1, RN-8
func TestDashboardDeUnProyectoConUnSprintEnCurso(t *testing.T) {
	registry := domain.NewProjectRegistry()
	project := dashboardProject(t, registry, 3, 5)
	if _, err := project.AddSprint(domain.SprintData{
		Goal:  "Tener el dashboard",
		Start: time.Date(2026, 10, 31, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2026, 11, 6, 0, 0, 0, 0, time.UTC),
	}); err != nil {
		t.Fatalf("no se pudo crear el sprint 2: %v", err)
	}
	server := newTestServer(t, registry)

	res := get(t, server, "/proyectos/1/dashboard")

	if res.Code != http.StatusOK {
		t.Fatalf("se esperaba el estado %d, salio %d", http.StatusOK, res.Code)
	}
	html := res.Body.String()
	// Las cuatro partes, en este orden: el encabezado, el estado general, el sprint en curso y
	// los gráficos. Y los gráficos, en el orden de la RN-8.
	for _, orden := range [][]string{
		{"<h1>Sistema de metricas</h1>", `<section id="estado-general"`, `<section id="sprint-en-curso"`, `<section id="graficos"`},
		{"<figcaption>Story points</figcaption>", "<figcaption>Horas</figcaption>", "<figcaption>Defectos</figcaption>", "<figcaption>Velocidad</figcaption>"},
	} {
		last := -1
		for _, parte := range orden {
			i := strings.Index(html, parte)
			if i <= last {
				t.Fatalf("se esperaba %q despues de la parte anterior (en este orden: %q), y no esta asi en:\n%s", parte, orden, html)
			}
			last = i
		}
	}
}
