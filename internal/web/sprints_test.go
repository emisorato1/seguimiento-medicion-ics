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
