package web

import (
	"net/http"
	"net/url"
	"slices"
	"strings"
	"testing"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// Escenario: Registrar esfuerzo y ver los totales
// Cubre: RN-1, RN-11, RN-12
func TestRegistrarEsfuerzoYVerLosTotales(t *testing.T) {
	registry := domain.NewProjectRegistry()
	project := addProject(t, registry, "Sistema de metricas")
	addMember(t, project, "Ana Perez", domain.RoleProductBuilder)
	addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
	saver := &fakeSaver{}
	navegador := newBrowser(t, NewServer(registry, saver, testToday))

	// El dia propuesto es el de hoy, y el integrante y la historia se eligen de la lista.
	html := navegador.get("/proyectos/1/esfuerzo/nuevo").Body.String()
	if !strings.Contains(html, `name="dia" value="2026-10-27"`) {
		t.Errorf("se esperaba el dia propuesto 27/10/2026, y no esta en:\n%s", html)
	}
	if got := options(html, "integrante"); !slices.Equal(got, []string{"Ana Perez"}) {
		t.Errorf("se esperaba elegir a Ana Perez, salio %v", got)
	}
	if got := options(html, "historia"); !slices.Equal(got, []string{"1"}) {
		t.Errorf("se esperaba elegir la historia 1, salio %v", got)
	}

	res := navegador.post("/proyectos/1/esfuerzo/nuevo", url.Values{
		"integrante": {"Ana Perez"},
		"historia":   {"1"},
		"dia":        {"2026-10-27"},
		"horas":      {"2,5"},
		"actividad":  {"Escribir tests"},
	})

	if res.Code != http.StatusSeeOther || res.Header().Get("Location") != "/proyectos/1/esfuerzo" {
		t.Fatalf("se esperaba volver al esfuerzo con el estado %d, salio %d hacia %q:\n%s", http.StatusSeeOther, res.Code, res.Header().Get("Location"), res.Body.String())
	}
	html = between(navegador.get("/proyectos/1/esfuerzo").Body.String(), "<main", "</main>")
	if !strings.Contains(html, "Se registraron 2,5 horas de Ana Perez en la historia 1") {
		t.Errorf("se esperaba el mensaje del registro, y no esta en:\n%s", html)
	}
	registros := between(html, `<table id="registros">`, "</table>")
	for _, texto := range []string{"27/10/2026", "Ana Perez", "Registrar esfuerzo", "Escribir tests", "<td>2,5</td>"} {
		if !strings.Contains(registros, texto) {
			t.Errorf("se esperaba en la tabla de esfuerzo %q, y no esta en:\n%s", texto, registros)
		}
	}
	for _, total := range []string{`<table id="por-historia">`, `<table id="por-integrante">`} {
		tabla := between(html, total, "</table>")
		if !strings.Contains(tabla, "<td>2,5</td>") {
			t.Errorf("se esperaba el total de 2,5 horas en %s, y no esta en:\n%s", total, tabla)
		}
	}
	if saver.saves != 1 {
		t.Errorf("se esperaba guardar una vez, se guardo %d", saver.saves)
	}
}

// Escenario: Cargar las horas estimadas de una historia
// Cubre: RN-1
func TestCargarLasHorasEstimadasDeUnaHistoria(t *testing.T) {
	registry := domain.NewProjectRegistry()
	project := addProject(t, registry, "Sistema de metricas")
	addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
	navegador := newBrowser(t, newTestServer(t, registry))

	html := between(navegador.get("/proyectos/1/historias/1").Body.String(), "<main", "</main>")
	if !strings.Contains(html, "sin horas estimadas") || !strings.Contains(html, `href="/proyectos/1/historias/1/horas-estimadas"`) {
		t.Errorf("se esperaba la historia sin horas estimadas y el acceso para cargarlas, y no estan en:\n%s", html)
	}

	res := navegador.post("/proyectos/1/historias/1/horas-estimadas", url.Values{"horas": {"6,5"}})

	if res.Code != http.StatusSeeOther || res.Header().Get("Location") != "/proyectos/1/historias/1" {
		t.Fatalf("se esperaba volver a la historia con el estado %d, salio %d hacia %q:\n%s", http.StatusSeeOther, res.Code, res.Header().Get("Location"), res.Body.String())
	}
	html = between(navegador.get("/proyectos/1/historias/1").Body.String(), "<main", "</main>")
	if !strings.Contains(html, "Se cargaron las horas estimadas de la historia 1") {
		t.Errorf("se esperaba el mensaje, y no esta en:\n%s", html)
	}
	if !strings.Contains(html, "<dt>Horas estimadas</dt>\n  <dd>6,5</dd>") {
		t.Errorf("se esperaba la historia con 6,5 horas estimadas, y no esta en:\n%s", html)
	}
	if story, _ := project.Story(1); story.EstimatedHours() != 6.5 {
		t.Errorf("se esperaban 6,5 horas estimadas en el nucleo, hay %v", story.EstimatedHours())
	}
}
