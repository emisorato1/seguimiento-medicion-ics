package web

import (
	"html/template"
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

	// Con todos los votos, al revelar se ve cada voto con su integrante (tarea 18, RN-8).
	navegador.post(ronda+"/votar", url.Values{"integrante": {"Juan Gomez"}, "puntos": {"8"}})
	res = navegador.post(ronda+"/revelar", nil)
	if res.Code != http.StatusSeeOther || res.Header().Get("Location") != ronda {
		t.Fatalf("se esperaba volver a la ronda con el estado %d, salio %d:\n%s", http.StatusSeeOther, res.Code, res.Body.String())
	}
	html = between(navegador.get(ronda).Body.String(), "<main", "</main>")
	votos := between(html, `<table id="votos">`, "</table>")
	for _, fila := range []string{"<td>Ana Perez</td>\n        <td>5</td>", "<td>Juan Gomez</td>\n        <td>8</td>"} {
		if !strings.Contains(votos, fila) {
			t.Errorf("se esperaba en los votos revelados %q, y no esta en:\n%s", fila, votos)
		}
	}
}

// Escenario: No se puede revelar con votos faltantes
// Cubre: RN-2
func TestNoSePuedeRevelarConVotosFaltantes(t *testing.T) {
	registry, project := pokerRegistry(t)
	round, err := project.OpenVotingRound(1)
	if err != nil {
		t.Fatalf("no se pudo abrir la ronda: %v", err)
	}
	if err := round.Vote("Ana Perez", 5); err != nil {
		t.Fatalf("no se pudo preparar el voto: %v", err)
	}
	saver := &fakeSaver{}
	server := NewServer(registry, saver, testToday)

	res := newBrowser(t, server).post("/proyectos/1/historias/1/rondas/1/revelar", nil)

	if res.Code != http.StatusUnprocessableEntity {
		t.Fatalf("se esperaba el estado %d, salio %d", http.StatusUnprocessableEntity, res.Code)
	}
	html := res.Body.String()
	if !strings.Contains(html, `id="error-formulario">`+template.HTMLEscapeString(domain.ErrVotesMissing.Error())) {
		t.Errorf("se esperaba arriba el mensaje %q, y no esta en:\n%s", domain.ErrVotesMissing, html)
	}
	if round.Revealed() || saver.saves != 0 {
		t.Errorf("la ronda tenia que seguir sin revelar y sin guardarse")
	}
}

// elements devuelve el contenido de cada elemento de la pagina con esas etiquetas, como cada
// fila de una tabla o cada renglon de una lista, para mirar que va junto a que.
func elements(html string, tags ...string) []string {
	var found []string
	for _, tag := range tags {
		rest := html
		for {
			i := strings.Index(rest, "<"+tag)
			if i < 0 {
				break
			}
			j := strings.Index(rest[i:], "</"+tag+">")
			if j < 0 {
				break
			}
			found = append(found, rest[i:i+j])
			rest = rest[i+j:]
		}
	}
	return found
}

// Escenario: Nadie ve un voto antes de revelar
// Cubre: RN-4
//
// No se busca el 13 suelto, porque el formulario de votar ofrece la escala entera y el
// numero puede estar en una fecha o un identificador (revision de Maximiliano Eula, puerta
// de la #78): se mira que la tabla de votos no exista, que ningun elemento que nombra a Ana
// Perez lleve el 13 y que el mensaje de su voto no lleve ningun numero.
func TestNadieVeUnVotoAntesDeRevelar(t *testing.T) {
	registry, project := pokerRegistry(t)
	addSprint(t, registry, 1)
	if _, err := project.OpenVotingRound(1); err != nil {
		t.Fatalf("no se pudo abrir la ronda: %v", err)
	}
	navegador := newBrowser(t, newTestServer(t, registry))
	ronda := "/proyectos/1/historias/1/rondas/1"

	res := navegador.post(ronda+"/votar", url.Values{"integrante": {"Ana Perez"}, "puntos": {"13"}})
	if res.Code != http.StatusSeeOther {
		t.Fatalf("se esperaba aceptar el voto con el estado %d, salio %d:\n%s", http.StatusSeeOther, res.Code, res.Body.String())
	}

	paginas := []string{ronda, ronda, "/proyectos/1", "/proyectos/1/backlog", "/proyectos/1/historias/1",
		"/proyectos/1/integrantes", "/proyectos/1/sprints", "/proyectos/1/sprints/1", "/proyectos/1/esfuerzo",
		"/proyectos/1/planning-poker", "/proyectos"}
	for i, pagina := range paginas {
		html := navegador.get(pagina).Body.String()
		// La primera vez, la ronda trae el mensaje del voto, que no lleva ningun numero.
		if i == 0 {
			mensaje := between(html, `role="status">`, "</p>")
			if !strings.Contains(mensaje, "Ana Perez") || strings.ContainsAny(mensaje, "0123456789") {
				t.Errorf("se esperaba el mensaje del voto de Ana Perez sin ningun numero, salio %q", mensaje)
			}
		}
		if strings.Contains(html, `id="votos"`) {
			t.Errorf("en %s no tenia que estar la tabla de votos antes de revelar", pagina)
		}
		for _, elemento := range elements(html, "tr", "li", "p", "dd", "td", "option", "span", "h1", "h2") {
			if strings.Contains(elemento, "Ana Perez") && strings.Contains(elemento, "13") {
				t.Errorf("en %s el 13 aparece junto a Ana Perez: %s", pagina, elemento)
			}
		}
	}
}

// Escenario: Un voto rechazado no muestra el valor elegido
// Cubre: RN-6
func TestUnVotoRechazadoNoMuestraElValorElegido(t *testing.T) {
	registry, project := pokerRegistry(t)
	round, err := project.OpenVotingRound(1)
	if err != nil {
		t.Fatalf("no se pudo abrir la ronda: %v", err)
	}
	if err := round.Vote("Ana Perez", 5); err != nil {
		t.Fatalf("no se pudo preparar el voto: %v", err)
	}
	server := newTestServer(t, registry)

	res := newBrowser(t, server).post("/proyectos/1/historias/1/rondas/1/votar", url.Values{"integrante": {"Ana Perez"}, "puntos": {"21"}})

	if res.Code != http.StatusUnprocessableEntity {
		t.Fatalf("se esperaba el estado %d, salio %d", http.StatusUnprocessableEntity, res.Code)
	}
	html := res.Body.String()
	if !strings.Contains(html, domain.ErrAlreadyVoted.Error()) {
		t.Errorf("se esperaba el mensaje %q, y no esta en:\n%s", domain.ErrAlreadyVoted, html)
	}
	// El 21 esta en la escala que se ofrece, pero no puede volver elegido.
	puntos := between(html, `<select id="puntos"`, "</select>")
	if strings.Contains(puntos, "selected") {
		t.Errorf("no se esperaba ningun valor elegido en los story points:\n%s", puntos)
	}
	for _, elemento := range elements(html, "tr", "li", "p", "div", "span", "option") {
		if strings.Contains(elemento, "Ana Perez") && strings.Contains(elemento, "21") && !strings.Contains(elemento, "<select") {
			t.Errorf("el 21 aparece junto a Ana Perez: %s", elemento)
		}
	}
}
