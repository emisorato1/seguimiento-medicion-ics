package web

import (
	"net/http"
	"net/url"
	"slices"
	"strings"
	"testing"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// pokerRegistry arma el proyecto 1 con los integrantes Ana Perez y Juan Gomez y la historia
// 1, que es con lo que arrancan los escenarios de Planning Poker.
func pokerRegistry(t *testing.T) (*domain.ProjectRegistry, *domain.Project) {
	t.Helper()
	registry := domain.NewProjectRegistry()
	project := addProject(t, registry, "Sistema de metricas")
	addMember(t, project, "Ana Perez", domain.RoleProductBuilder)
	addMember(t, project, "Juan Gomez", domain.RoleProductBuilder)
	addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
	return registry, project
}

// Escenario: Una ronda de Planning Poker sin ver los votos hasta revelar
// Cubre: RN-1, RN-5, RN-7
func TestUnaRondaDePlanningPokerSinVerLosVotosHastaRevelar(t *testing.T) {
	registry, _ := pokerRegistry(t)
	navegador := newBrowser(t, newTestServer(t, registry))

	// La ronda se abre eligiendo la historia de la lista.
	html := navegador.get("/proyectos/1/planning-poker").Body.String()
	if got := options(html, "historia"); !slices.Equal(got, []string{"1"}) {
		t.Errorf("se esperaba abrir la ronda eligiendo la historia 1, salio %v", got)
	}
	res := navegador.post("/proyectos/1/planning-poker", url.Values{"historia": {"1"}})
	ronda := "/proyectos/1/historias/1/rondas/1"
	if res.Code != http.StatusSeeOther || res.Header().Get("Location") != ronda {
		t.Fatalf("se esperaba ir a la ronda con el estado %d, salio %d hacia %q:\n%s", http.StatusSeeOther, res.Code, res.Header().Get("Location"), res.Body.String())
	}
	html = navegador.get(ronda).Body.String()
	if !strings.Contains(html, "Se abrió la ronda 1 de la historia 1") {
		t.Errorf("se esperaba el mensaje de la ronda abierta, y no esta en:\n%s", html)
	}
	if got := options(html, "integrante"); !slices.Equal(got, []string{"Ana Perez", "Juan Gomez"}) {
		t.Errorf("se esperaba que voten Ana Perez y Juan Gomez, salio %v", got)
	}

	res = navegador.post(ronda+"/votar", url.Values{"integrante": {"Ana Perez"}, "puntos": {"5"}})
	if res.Code != http.StatusSeeOther || res.Header().Get("Location") != ronda {
		t.Fatalf("se esperaba volver a la ronda con el estado %d, salio %d:\n%s", http.StatusSeeOther, res.Code, res.Body.String())
	}
	html = between(navegador.get(ronda).Body.String(), "<main", "</main>")
	if !strings.Contains(html, "Se registró el voto de Ana Perez") {
		t.Errorf("se esperaba el mensaje del voto, y no esta en:\n%s", html)
	}
	votaron := between(html, `<ul id="votaron">`, "</ul>")
	faltan := between(html, `<ul id="faltan">`, "</ul>")
	if !strings.Contains(votaron, "Ana Perez") || !strings.Contains(faltan, "Juan Gomez") || strings.Contains(faltan, "Ana Perez") {
		t.Errorf("se esperaba que voto Ana Perez y falta Juan Gomez, salio:\nvotaron: %s\nfaltan: %s", votaron, faltan)
	}
	if !strings.Contains(html, "Votaron 1 de 2") {
		t.Errorf("se esperaba ver que votaron 1 de 2, y no esta en:\n%s", html)
	}
	if got := options(html, "integrante"); !slices.Equal(got, []string{"Juan Gomez"}) {
		t.Errorf("se esperaba que solo falte votar Juan Gomez, salio %v", got)
	}
}
