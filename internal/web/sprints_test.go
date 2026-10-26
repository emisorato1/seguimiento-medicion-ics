package web

import (
	"net/http"
	"net/url"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// Escenario: La barra de un proyecto lleva a todas sus secciones
// Cubre: RN-3
//
// Cada pantalla de la #78 suma su fila cuando llega a dev, asi la barra no lleva nunca a una
// pagina que todavia no existe (Fase 4, Desvios).
func TestLaBarraDeUnProyectoLlevaATodasSusSecciones(t *testing.T) {
	secciones := []struct {
		nombre string
		pagina string
	}{
		{nombre: "Sprints", pagina: "/proyectos/1/sprints"},
		{nombre: "Esfuerzo", pagina: "/proyectos/1/esfuerzo"},
		{nombre: "Defectos", pagina: "/proyectos/1/defectos"},
		{nombre: "Planning Poker", pagina: "/proyectos/1/planning-poker"},
	}

	registry := domain.NewProjectRegistry()
	addProject(t, registry, "Sistema de metricas")
	server := newTestServer(t, registry)

	for _, seccion := range secciones {
		t.Run(seccion.nombre, func(t *testing.T) {
			// Desde cualquier pagina del proyecto, la barra lleva a la seccion.
			for _, desde := range []string{"/proyectos/1", "/proyectos/1/backlog", seccion.pagina} {
				res := get(t, server, desde)
				if res.Code != http.StatusOK {
					t.Fatalf("se esperaba %s con el estado %d, salio %d", desde, http.StatusOK, res.Code)
				}
				barra := between(res.Body.String(), `<nav aria-label="Navegación principal">`, "</nav>")
				enlace := `<a href="` + seccion.pagina + `"`
				if !strings.Contains(barra, "Sistema de metricas") || !strings.Contains(barra, enlace) || !strings.Contains(barra, ">"+seccion.nombre+"</a>") {
					t.Errorf("en %s se esperaba en la barra el proyecto y el acceso a %s, y no estan en:\n%s", desde, seccion.nombre, barra)
				}
			}

			// En su propia pagina, la seccion figura como la actual.
			barra := between(get(t, server, seccion.pagina).Body.String(), `<nav aria-label="Navegación principal">`, "</nav>")
			if !strings.Contains(barra, `<a href="`+seccion.pagina+`" aria-current="page">`) {
				t.Errorf("se esperaba %s marcada como la seccion actual, y no esta en:\n%s", seccion.nombre, barra)
			}
		})
	}
}

// Escenario: Crear un sprint desde la web
// Cubre: RN-1, RN-2
func TestCrearUnSprintDesdeLaWeb(t *testing.T) {
	registry := domain.NewProjectRegistry()
	project := addProject(t, registry, "Sistema de metricas")
	saver := &fakeSaver{}
	navegador := newBrowser(t, NewServer(registry, saver, testToday))

	res := navegador.post("/proyectos/1/sprints/nuevo", url.Values{
		"goal":   {"Tener el MVP"},
		"inicio": {"2026-10-26"},
		"fin":    {"2026-10-30"},
	})

	if res.Code != http.StatusSeeOther {
		t.Fatalf("se esperaba el estado %d, salio %d:\n%s", http.StatusSeeOther, res.Code, res.Body.String())
	}
	lista := res.Header().Get("Location")
	if lista != "/proyectos/1/sprints" {
		t.Fatalf("se esperaba volver a la lista de sprints, salio %q", lista)
	}
	html := between(navegador.get(lista).Body.String(), "<main", "</main>")
	if !strings.Contains(html, "Se creó el sprint 1") {
		t.Errorf("se esperaba el mensaje %q, y no esta en:\n%s", "Se creó el sprint 1", html)
	}
	fila := between(html, "<tbody>", "</tbody>")
	for _, texto := range []string{"<td>1</td>", "Tener el MVP", "26/10/2026", "30/10/2026", "<td>Abierto</td>"} {
		if !strings.Contains(fila, texto) {
			t.Errorf("se esperaba en la lista de sprints %q, y no esta en:\n%s", texto, fila)
		}
	}

	// Recargar la lista no crea otro sprint, y el mensaje ya no aparece.
	html = between(navegador.get(lista).Body.String(), "<main", "</main>")
	if strings.Contains(html, "Se creó el sprint 1") {
		t.Errorf("el mensaje no tenia que repetirse al recargar:\n%s", html)
	}
	if n := len(project.Sprints()); n != 1 || saver.saves != 1 {
		t.Errorf("se esperaba un solo sprint guardado una vez, hay %d sprints y %d guardados", n, saver.saves)
	}
}

// addSprint crea en el proyecto un sprint abierto del 26/10/2026 al 30/10/2026.
func addSprint(t *testing.T, registry *domain.ProjectRegistry, projectID int) *domain.Sprint {
	t.Helper()
	project, err := registry.Get(projectID)
	if err != nil {
		t.Fatalf("no se encontro el proyecto %d: %v", projectID, err)
	}
	sprint, err := project.AddSprint(domain.SprintData{
		Goal:  "Tener el MVP",
		Start: time.Date(2026, 10, 26, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2026, 10, 30, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		t.Fatalf("no se pudo preparar el sprint: %v", err)
	}
	return sprint
}

// Escenario: Asignar y quitar historias del sprint abierto
// Cubre: RN-1, RN-16
func TestAsignarYQuitarHistoriasDelSprintAbierto(t *testing.T) {
	registry := domain.NewProjectRegistry()
	project := addProject(t, registry, "Sistema de metricas")
	for _, h := range []struct {
		titulo string
		puntos int
	}{{"Registrar esfuerzo", 3}, {"Cargar defectos", 5}} {
		story := addStory(t, project, h.titulo, domain.PriorityHigh)
		if err := project.EstimateStory(story.ID(), h.puntos); err != nil {
			t.Fatalf("no se pudo estimar la historia: %v", err)
		}
	}
	addStory(t, project, "Sin estimar", domain.PriorityLow)
	addSprint(t, registry, 1)
	navegador := newBrowser(t, newTestServer(t, registry))

	// Para asignar se elige de la lista con todo el backlog, aunque el nucleo despues
	// rechace la historia sin estimar (RN-16).
	html := navegador.get("/proyectos/1/sprints/1").Body.String()
	if got := options(html, "historia"); !slices.Equal(got, []string{"1", "2", "3"}) {
		t.Errorf("se esperaban para asignar las historias 1, 2 y 3, salieron %v", got)
	}

	for _, historia := range []string{"1", "2"} {
		res := navegador.post("/proyectos/1/sprints/asignar", url.Values{"historia": {historia}})
		if res.Code != http.StatusSeeOther || res.Header().Get("Location") != "/proyectos/1/sprints/1" {
			t.Fatalf("se esperaba volver al sprint 1 con el estado %d, salio %d hacia %q", http.StatusSeeOther, res.Code, res.Header().Get("Location"))
		}
	}
	html = between(navegador.get("/proyectos/1/sprints/1").Body.String(), "<main", "</main>")
	if !strings.Contains(html, "Se asignó la historia 2 al sprint 1") {
		t.Errorf("se esperaba el mensaje de la asignacion, y no esta en:\n%s", html)
	}
	backlog := between(html, "<tbody>", "</tbody>")
	for _, texto := range []string{"Registrar esfuerzo", "Cargar defectos", "<td>3</td>", "<td>5</td>"} {
		if !strings.Contains(backlog, texto) {
			t.Errorf("se esperaba en el Sprint Backlog %q, y no esta en:\n%s", texto, backlog)
		}
	}
	if n := strings.Count(backlog, "<td>En curso</td>"); n != 2 {
		t.Errorf("se esperaban las dos historias En curso, hay %d:\n%s", n, backlog)
	}

	// Quitar pide confirmacion, y al confirmar la historia vuelve a Pendiente.
	html = navegador.get("/proyectos/1/sprints/quitar?historia=1").Body.String()
	if !strings.Contains(html, "¿Quitar la historia 1 del sprint 1?") {
		t.Fatalf("se esperaba la pregunta para confirmar, y no esta en:\n%s", html)
	}
	res := navegador.post("/proyectos/1/sprints/quitar", url.Values{"historia": {"1"}})
	if res.Code != http.StatusSeeOther {
		t.Fatalf("se esperaba el estado %d, salio %d:\n%s", http.StatusSeeOther, res.Code, res.Body.String())
	}
	html = between(navegador.get(res.Header().Get("Location")).Body.String(), "<main", "</main>")
	if !strings.Contains(html, "Se quitó la historia 1 del sprint 1") {
		t.Errorf("se esperaba el mensaje de quitar, y no esta en:\n%s", html)
	}
	backlog = between(html, "<tbody>", "</tbody>")
	if strings.Contains(backlog, "Registrar esfuerzo") || !strings.Contains(backlog, "Cargar defectos") {
		t.Errorf("se esperaba que solo quede la historia 2 en el Sprint Backlog:\n%s", backlog)
	}
	if story, _ := project.Story(1); story.State() != domain.StoryPending {
		t.Errorf("se esperaba la historia 1 Pendiente, esta %v", story.State())
	}
}

// Sin escenario: quitar del sprint una historia que no esta asignada vuelve a pedir confirmar
// con el mensaje del nucleo, sin cambiar nada (RN-2; RN-10 y RN-13 de la #77).
// Cubre: RN-2
func TestQuitarDelSprintUnaHistoriaQueNoEstaAsignada(t *testing.T) {
	registry := domain.NewProjectRegistry()
	project := addProject(t, registry, "Sistema de metricas")
	addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
	addSprint(t, registry, 1)
	saver := &fakeSaver{}
	navegador := newBrowser(t, NewServer(registry, saver, testToday))
	antes := snapshot(registry)

	res := navegador.post("/proyectos/1/sprints/quitar", url.Values{"historia": {"1"}})

	if res.Code != http.StatusUnprocessableEntity {
		t.Fatalf("se esperaba el estado %d, salio %d", http.StatusUnprocessableEntity, res.Code)
	}
	html := res.Body.String()
	if !strings.Contains(html, `id="error-formulario">`) || !strings.Contains(html, "¿Quitar la historia 1 del sprint 1?") {
		t.Errorf("se esperaba la confirmacion otra vez con el mensaje del nucleo arriba:\n%s", html)
	}
	if snapshot(registry) != antes || saver.saves != 0 {
		t.Errorf("no se esperaba ningun cambio")
	}
}

// closeRegistry arma el sprint 1 abierto con las historias 1 y 2 asignadas, estimadas en 3 y 5.
func closeRegistry(t *testing.T) (*domain.ProjectRegistry, *domain.Project) {
	t.Helper()
	registry := domain.NewProjectRegistry()
	project := addProject(t, registry, "Sistema de metricas")
	for _, h := range []struct {
		titulo string
		puntos int
	}{{"Registrar esfuerzo", 3}, {"Cargar defectos", 5}} {
		story := addStory(t, project, h.titulo, domain.PriorityHigh)
		if err := project.EstimateStory(story.ID(), h.puntos); err != nil {
			t.Fatalf("no se pudo estimar la historia: %v", err)
		}
	}
	addSprint(t, registry, 1)
	for _, id := range []int{1, 2} {
		if err := project.AssignStory(id); err != nil {
			t.Fatalf("no se pudo asignar la historia %d: %v", id, err)
		}
	}
	return registry, project
}

// el30DeOctubre es la fecha de hoy del escenario del cierre.
func el30DeOctubre() time.Time {
	return time.Date(2026, 10, 30, 0, 0, 0, 0, time.UTC)
}

// Escenario: Cerrar el sprint marcando las completadas
// Cubre: RN-1, RN-9, RN-10, RN-11
func TestCerrarElSprintMarcandoLasCompletadas(t *testing.T) {
	registry, project := closeRegistry(t)
	navegador := newBrowser(t, NewServer(registry, &fakeSaver{}, el30DeOctubre))

	// El sprint abierto ofrece cerrarlo.
	if html := navegador.get("/proyectos/1/sprints/1").Body.String(); !strings.Contains(html, `href="/proyectos/1/sprints/cerrar"`) {
		t.Fatalf("se esperaba en el sprint abierto el acceso para cerrarlo:\n%s", html)
	}

	// La fecha de cierre propuesta es hoy y ninguna historia esta marcada.
	html := navegador.get("/proyectos/1/sprints/cerrar").Body.String()
	if !strings.Contains(html, `id="cierre" name="cierre" value="2026-10-30"`) {
		t.Errorf("se esperaba la fecha de cierre propuesta 2026-10-30, y no esta en:\n%s", html)
	}
	for _, id := range []string{"1", "2"} {
		if !strings.Contains(html, `type="checkbox" id="completada-`+id+`" name="completada" value="`+id+`"`) {
			t.Errorf("se esperaba la casilla de la historia %s, y no esta en:\n%s", id, html)
		}
	}
	if strings.Contains(html, "checked") {
		t.Errorf("ninguna historia tenia que estar marcada de entrada:\n%s", html)
	}

	// Marcar la 1 y enviar pide confirmar, con lo que queda completado y lo que vuelve.
	html = navegador.get("/proyectos/1/sprints/cerrar/confirmar?cierre=2026-10-30&completada=1").Body.String()
	for _, texto := range []string{"¿Cerrar el sprint 1 el 30/10/2026?", "Quedan completadas: 1 — Registrar esfuerzo.",
		"Vuelven al Product Backlog: 2 — Cargar defectos.", `name="completadas" value="1"`} {
		if !strings.Contains(html, texto) {
			t.Errorf("se esperaba en la confirmacion %q, y no esta en:\n%s", texto, html)
		}
	}
	if _, ok := project.OpenSprint(); !ok {
		t.Fatal("pedir que se cierre no tenia que cerrar el sprint")
	}

	// Confirmar lo cierra.
	res := navegador.post("/proyectos/1/sprints/cerrar", url.Values{"cierre": {"2026-10-30"}, "completadas": {"1"}})
	if res.Code != http.StatusSeeOther || res.Header().Get("Location") != "/proyectos/1/sprints/1" {
		t.Fatalf("se esperaba volver al sprint 1 con el estado %d, salio %d hacia %q:\n%s", http.StatusSeeOther, res.Code, res.Header().Get("Location"), res.Body.String())
	}
	html = joinTags(between(navegador.get("/proyectos/1/sprints/1").Body.String(), "<main", "</main>"))
	if !strings.Contains(html, "Se cerró el sprint 1") || !strings.Contains(html, "<dt>Estado</dt><dd>Cerrado</dd>") {
		t.Errorf("se esperaba el mensaje y el sprint Cerrado, y no estan en:\n%s", html)
	}
	if _, ok := project.OpenSprint(); ok {
		t.Error("se esperaba que no quede sprint abierto")
	}
	if !slices.Equal(project.Sprints()[0].CompletedStories(), []int{1}) {
		t.Errorf("se esperaba completada solo la historia 1, quedo %v", project.Sprints()[0].CompletedStories())
	}
}

// Esquema del escenario: Confirmar o cancelar antes de lo que no se deshace (cerrar)
// Cubre: RN-9
func TestConfirmarOCancelarAntesDeCerrarElSprint(t *testing.T) {
	for _, confirmo := range []bool{true, false} {
		nombre := map[bool]string{true: "confirmo", false: "cancelo"}[confirmo]
		t.Run(nombre, func(t *testing.T) {
			registry, _ := closeRegistry(t)
			saver := &fakeSaver{}
			navegador := newBrowser(t, NewServer(registry, saver, el30DeOctubre))
			antes := snapshot(registry)

			confirmacion := between(navegador.get("/proyectos/1/sprints/cerrar/confirmar?cierre=2026-10-30").Body.String(), "<main", "</main>")
			for _, texto := range []string{`action="/proyectos/1/sprints/cerrar"`, `href="/proyectos/1/sprints/1"`, "Quedan completadas: ninguna."} {
				if !strings.Contains(confirmacion, texto) {
					t.Errorf("se esperaba en la confirmacion %s, y no esta en:\n%s", texto, confirmacion)
				}
			}
			if confirmo {
				navegador.post("/proyectos/1/sprints/cerrar", url.Values{"cierre": {"2026-10-30"}, "completadas": {""}})
			} else {
				navegador.get("/proyectos/1/sprints/1")
			}

			quiero := map[bool]string{true: "<td>Cerrado</td>", false: "<td>Abierto</td>"}[confirmo]
			if html := navegador.get("/proyectos/1/sprints").Body.String(); !strings.Contains(html, quiero) {
				t.Errorf("se esperaba el sprint 1 con %s, y no esta en:\n%s", quiero, html)
			}
			if !confirmo && (snapshot(registry) != antes || saver.saves != 0) {
				t.Error("cancelar no deberia cambiar nada")
			}
		})
	}
}
