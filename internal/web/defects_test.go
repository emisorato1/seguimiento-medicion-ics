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

// defectRegistry arma el proyecto 1 con la historia 1, el sprint 1 cerrado y el sprint 2
// abierto, que es lo que piden los escenarios de defectos.
func defectRegistry(t *testing.T) (*domain.ProjectRegistry, *domain.Project) {
	t.Helper()
	registry := domain.NewProjectRegistry()
	project := addProject(t, registry, "Sistema de metricas")
	addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
	addSprint(t, registry, 1)
	if err := project.CloseSprint(nil, time.Date(2026, 10, 30, 0, 0, 0, 0, time.UTC)); err != nil {
		t.Fatalf("no se pudo cerrar el sprint 1: %v", err)
	}
	if _, err := project.AddSprint(domain.SprintData{
		Goal:  "Cerrar la interfaz",
		Start: time.Date(2026, 11, 2, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2026, 11, 13, 0, 0, 0, 0, time.UTC),
	}); err != nil {
		t.Fatalf("no se pudo preparar el sprint 2: %v", err)
	}
	return registry, project
}

// Escenario: Cargar un defecto, pasarlo a en curso y resolverlo
// Cubre: RN-1, RN-2, RN-16
func TestCargarUnDefectoPasarloAEnCursoYResolverlo(t *testing.T) {
	registry, project := defectRegistry(t)
	saver := &fakeSaver{}
	navegador := newBrowser(t, NewServer(registry, saver, testToday))

	// Lo que tiene valores fijos o sale de lo cargado se elige de una lista (RN-2, RN-16).
	html := navegador.get("/proyectos/1/defectos/nuevo").Body.String()
	for _, lista := range []struct {
		nombre string
		quiero []string
	}{
		{"severidad", []string{"Crítica", "Mayor", "Menor"}},
		{"historia", []string{"1"}},
		{"sprint", []string{"1", "2"}},
	} {
		if got := options(html, lista.nombre); !slices.Equal(got, lista.quiero) {
			t.Errorf("se esperaba elegir %s entre %v, salio %v", lista.nombre, lista.quiero, got)
		}
	}

	res := navegador.post("/proyectos/1/defectos/nuevo", url.Values{
		"titulo":      {"Suma mal las horas"},
		"descripcion": {"El total por historia no suma el ultimo registro"},
		"severidad":   {"Mayor"},
		"historia":    {"1"},
		"sprint":      {"1"},
	})
	if res.Code != http.StatusSeeOther || res.Header().Get("Location") != "/proyectos/1/defectos" {
		t.Fatalf("se esperaba volver a la lista de defectos con el estado %d, salio %d hacia %q:\n%s", http.StatusSeeOther, res.Code, res.Header().Get("Location"), res.Body.String())
	}
	html = joinTags(between(navegador.get("/proyectos/1/defectos").Body.String(), "<main", "</main>"))
	if !strings.Contains(html, "Se cargó el defecto 1") {
		t.Errorf("se esperaba el mensaje %q, y no esta en:\n%s", "Se cargó el defecto 1", html)
	}
	fila := between(html, "<tbody>", "</tbody>")
	for _, texto := range []string{"<td>1</td>", "Suma mal las horas", "<td>Mayor</td>", "<td>Abierto</td>"} {
		if !strings.Contains(fila, texto) {
			t.Errorf("se esperaba en la lista de defectos %q, y no esta en:\n%s", texto, fila)
		}
	}

	// Pasarlo a En curso se elige de la lista de estados, en la pagina del defecto.
	html = navegador.get("/proyectos/1/defectos/1").Body.String()
	if got := options(html, "estado"); !slices.Equal(got, []string{"Abierto", "En curso"}) {
		t.Errorf("se esperaba elegir el estado entre Abierto y En curso, salio %v", got)
	}
	res = navegador.post("/proyectos/1/defectos/1/estado", url.Values{"estado": {"En curso"}})
	if res.Code != http.StatusSeeOther || res.Header().Get("Location") != "/proyectos/1/defectos/1" {
		t.Fatalf("se esperaba volver al defecto 1 con el estado %d, salio %d hacia %q:\n%s", http.StatusSeeOther, res.Code, res.Header().Get("Location"), res.Body.String())
	}
	html = joinTags(between(navegador.get("/proyectos/1/defectos/1").Body.String(), "<main", "</main>"))
	if !strings.Contains(html, "El defecto 1 pasó a En curso") || !strings.Contains(html, "<dt>Estado</dt><dd>En curso</dd>") {
		t.Errorf("se esperaba el defecto 1 En curso, con su mensaje, y no esta en:\n%s", html)
	}
	if n := len(project.DefectsByState(domain.DefectInProgress)); n != 1 || saver.saves != 2 {
		t.Errorf("se esperaba un defecto En curso y dos guardados, hay %d y %d", n, saver.saves)
	}
}
