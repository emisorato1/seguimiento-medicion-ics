package web

import (
	"fmt"
	"html/template"
	"net/http"
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
