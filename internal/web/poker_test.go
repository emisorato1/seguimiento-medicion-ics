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
	votos := joinTags(between(html, `<table id="votos">`, "</table>"))
	for _, fila := range []string{"<td>Ana Perez</td><td>5</td>", "<td>Juan Gomez</td><td>8</td>"} {
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

// Escenario: Diferencias, nueva ronda y acuerdo
// Cubre: RN-1, RN-8
func TestDiferenciasNuevaRondaYAcuerdo(t *testing.T) {
	// Dado la ronda 1 de la historia 1 revelada con 3 de Ana Perez y 8 de Juan Gomez.
	registry, project := pokerRegistry(t)
	round, err := project.OpenVotingRound(1)
	if err != nil {
		t.Fatalf("no se pudo abrir la ronda: %v", err)
	}
	for nombre, puntos := range map[string]int{"Ana Perez": 3, "Juan Gomez": 8} {
		if err := round.Vote(nombre, puntos); err != nil {
			t.Fatalf("no se pudo votar: %v", err)
		}
	}
	if err := round.Reveal(); err != nil {
		t.Fatalf("no se pudo revelar: %v", err)
	}
	navegador := newBrowser(t, newTestServer(t, registry))

	// Las diferencias, tal como las informa el nucleo.
	html := navegador.get("/proyectos/1/historias/1/rondas/1").Body.String()
	diferencias := between(html, `<section id="diferencias"`, "</section>")
	if !strings.Contains(diferencias, "Hay diferencias: el más bajo es 3, de Ana Perez, y el más alto es 8, de Juan Gomez.") {
		t.Errorf("se esperaban las diferencias con el mas bajo y el mas alto, y no estan en:\n%s", html)
	}

	// Una ronda nueva desde la ronda revelada; los dos votan 5 y se revela.
	res := navegador.post("/proyectos/1/planning-poker", url.Values{"historia": {"1"}})
	if res.Code != http.StatusSeeOther || res.Header().Get("Location") != "/proyectos/1/historias/1/rondas/2" {
		t.Fatalf("se esperaba ir a la ronda 2 con el estado %d, salio %d hacia %q", http.StatusSeeOther, res.Code, res.Header().Get("Location"))
	}
	if !strings.Contains(html, `action="/proyectos/1/planning-poker"`) || !strings.Contains(html, `name="historia" value="1"`) {
		t.Errorf("se esperaba en la ronda revelada el boton para abrir una ronda nueva:\n%s", html)
	}
	for _, nombre := range []string{"Ana Perez", "Juan Gomez"} {
		navegador.post("/proyectos/1/historias/1/rondas/2/votar", url.Values{"integrante": {nombre}, "puntos": {"5"}})
	}
	navegador.post("/proyectos/1/historias/1/rondas/2/revelar", nil)
	html = navegador.get("/proyectos/1/historias/1/rondas/2").Body.String()
	if diferencias := between(html, `<section id="diferencias"`, "</section>"); !strings.Contains(diferencias, "No hay diferencias: todos votaron 5.") {
		t.Errorf("se esperaba que no hay diferencias, y no esta en:\n%s", html)
	}

	// Acordar se elige de la escala, y deja la historia estimada.
	if got := options(html, "puntos"); !slices.Contains(got, "5") || !slices.Contains(got, "8") {
		t.Errorf("se esperaba acordar eligiendo de la escala, salio %v", got)
	}
	res = navegador.post("/proyectos/1/historias/1/rondas/2/acordar", url.Values{"puntos": {"5"}})
	if res.Code != http.StatusSeeOther || res.Header().Get("Location") != "/proyectos/1/historias/1/rondas/2" {
		t.Fatalf("se esperaba volver a la ronda 2 con el estado %d, salio %d hacia %q:\n%s", http.StatusSeeOther, res.Code, res.Header().Get("Location"), res.Body.String())
	}
	html = joinTags(between(navegador.get("/proyectos/1/historias/1/rondas/2").Body.String(), "<main", "</main>"))
	if !strings.Contains(html, "Se acordó la historia 1 en 5 story points") || !strings.Contains(html, "<dt>Estimación acordada</dt><dd>5</dd>") {
		t.Errorf("se esperaba el mensaje y la estimacion acordada, y no estan en:\n%s", html)
	}
	if strings.Contains(html, `/rondas/2/acordar"`) {
		t.Errorf("una ronda acordada no deberia ofrecer acordar otra vez:\n%s", html)
	}
	if backlog := joinTags(navegador.get("/proyectos/1/backlog").Body.String()); !strings.Contains(backlog, "Registrar esfuerzo</a></td><td>Alta</td><td>Pendiente</td><td>5</td>") {
		t.Errorf("se esperaba la historia 1 con 5 story points en el backlog:\n%s", backlog)
	}

	// La ronda 1 sigue en la lista de rondas, sin acordar, y ya no ofrece acordar.
	lista := joinTags(between(navegador.get("/proyectos/1/historias/1/rondas").Body.String(), "<tbody>", "</tbody>"))
	filas := strings.Split(lista, "<tr>")[1:]
	if len(filas) != 2 || !strings.Contains(filas[0], "Ronda 1") || !strings.Contains(filas[0], "sin acordar") || !strings.Contains(filas[1], "<td>5</td>") {
		t.Errorf("se esperaban la ronda 1 sin acordar y la 2 acordada en 5, y la lista es:\n%s", lista)
	}
	if html := navegador.get("/proyectos/1/historias/1/rondas/1").Body.String(); strings.Contains(html, `/rondas/1/acordar"`) {
		t.Errorf("una ronda que no es la ultima no deberia ofrecer acordar:\n%s", html)
	}
}

// Defecto #220: con tres o mas integrantes en el voto mas bajo o en el mas alto, la frase de
// las diferencias los separa con comas y une el ultimo con «y»; con dos, solo con «y».
// Cubre: #78 RN-8
func TestLasDiferenciasSeparanConComasTresOMasNombres(t *testing.T) {
	casos := []struct {
		nombre string
		votos  map[string]int
		frase  string
	}{
		{
			nombre: "tres en el mas bajo",
			votos:  map[string]int{"Emiliano Sorato": 5, "Agustín Salinas": 5, "Maximiliano Eula": 5, "Juliana Bustos": 8},
			frase:  "Hay diferencias: el más bajo es 5, de Emiliano Sorato, Agustín Salinas y Maximiliano Eula, y el más alto es 8, de Juliana Bustos.",
		},
		{
			nombre: "tres en el mas alto",
			votos:  map[string]int{"Emiliano Sorato": 3, "Agustín Salinas": 8, "Maximiliano Eula": 8, "Juliana Bustos": 8},
			frase:  "Hay diferencias: el más bajo es 3, de Emiliano Sorato, y el más alto es 8, de Agustín Salinas, Maximiliano Eula y Juliana Bustos.",
		},
		{
			nombre: "dos en cada uno",
			votos:  map[string]int{"Emiliano Sorato": 3, "Agustín Salinas": 3, "Maximiliano Eula": 8, "Juliana Bustos": 8},
			frase:  "Hay diferencias: el más bajo es 3, de Emiliano Sorato y Agustín Salinas, y el más alto es 8, de Maximiliano Eula y Juliana Bustos.",
		},
	}
	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			// Dado la ronda 1 de la historia 1, con cuatro integrantes, revelada.
			registry := domain.NewProjectRegistry()
			project := addProject(t, registry, "Sistema de metricas")
			for _, nombre := range []string{"Emiliano Sorato", "Agustín Salinas", "Maximiliano Eula", "Juliana Bustos"} {
				addMember(t, project, nombre, domain.RoleProductBuilder)
			}
			addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
			round, err := project.OpenVotingRound(1)
			if err != nil {
				t.Fatalf("no se pudo abrir la ronda: %v", err)
			}
			for nombre, puntos := range caso.votos {
				if err := round.Vote(nombre, puntos); err != nil {
					t.Fatalf("no se pudo votar: %v", err)
				}
			}
			if err := round.Reveal(); err != nil {
				t.Fatalf("no se pudo revelar: %v", err)
			}

			html := get(t, newTestServer(t, registry), "/proyectos/1/historias/1/rondas/1").Body.String()

			diferencias := between(html, `<section id="diferencias"`, "</section>")
			if !strings.Contains(diferencias, caso.frase) {
				t.Errorf("se esperaba la frase %q, y las diferencias son:\n%s", caso.frase, diferencias)
			}
		})
	}
}

// revealedRound abre la ronda 1 de la historia 1, con 3 de Ana Perez y 8 de Juan Gomez, y la
// revela.
func revealedRound(t *testing.T, project *domain.Project) *domain.VotingRound {
	t.Helper()
	round, err := project.OpenVotingRound(1)
	if err != nil {
		t.Fatalf("no se pudo abrir la ronda: %v", err)
	}
	for nombre, puntos := range map[string]int{"Ana Perez": 3, "Juan Gomez": 8} {
		if err := round.Vote(nombre, puntos); err != nil {
			t.Fatalf("no se pudo votar: %v", err)
		}
	}
	if err := round.Reveal(); err != nil {
		t.Fatalf("no se pudo revelar: %v", err)
	}
	return round
}

// Sin escenario propio: la verificacion de la Fase 3 de una ronda votada desde dos pestañas
// (pedida por Emiliano Sorato el 26/10). Ana Perez y Juan Gomez votan cada uno desde su pestaña;
// ninguna ve un voto hasta revelar, una pestaña con la pagina vieja no puede votar otra vez
// por quien ya voto, y al revelar las dos ven la diferencia.
// Cubre: #78 RN-4, RN-5, RN-7, RN-8; RN-20 de la #77
func TestUnaRondaVotadaDesdeDosPestanas(t *testing.T) {
	registry, project := pokerRegistry(t)
	if _, err := project.OpenVotingRound(1); err != nil {
		t.Fatalf("no se pudo abrir la ronda: %v", err)
	}
	server := newTestServer(t, registry)
	pestanaDeAna, pestanaDeJuan := newBrowser(t, server), newBrowser(t, server)
	ronda := "/proyectos/1/historias/1/rondas/1"

	// Las dos abren la ronda y ofrecen votar a los dos.
	for _, pestana := range []*browser{pestanaDeAna, pestanaDeJuan} {
		if got := options(pestana.get(ronda).Body.String(), "integrante"); !slices.Equal(got, []string{"Ana Perez", "Juan Gomez"}) {
			t.Fatalf("se esperaba ofrecer votar a Ana Perez y Juan Gomez, salio %v", got)
		}
	}

	// Cada uno vota desde su pestaña.
	for _, voto := range []struct {
		pestana *browser
		nombre  string
		puntos  string
	}{{pestanaDeAna, "Ana Perez", "3"}, {pestanaDeJuan, "Juan Gomez", "8"}} {
		if res := voto.pestana.post(ronda+"/votar", url.Values{"integrante": {voto.nombre}, "puntos": {voto.puntos}}); res.Code != http.StatusSeeOther {
			t.Fatalf("se esperaba aceptar el voto de %s con el estado %d, salio %d:\n%s", voto.nombre, http.StatusSeeOther, res.Code, res.Body.String())
		}
	}

	// Cada pestaña ve el mensaje de su voto, no el de la otra, y ninguna ve un voto.
	for _, caso := range []struct {
		pestana *browser
		suyo    string
		otro    string
	}{{pestanaDeAna, "Ana Perez", "Juan Gomez"}, {pestanaDeJuan, "Juan Gomez", "Ana Perez"}} {
		html := between(caso.pestana.get(ronda).Body.String(), "<main", "</main>")
		mensaje := between(html, `role="status">`, "</p>")
		if !strings.Contains(mensaje, "Se registró el voto de "+caso.suyo) || strings.Contains(mensaje, caso.otro) {
			t.Errorf("en la pestaña de %s se esperaba solo el mensaje de su voto, salio %q", caso.suyo, mensaje)
		}
		if !strings.Contains(html, "Votaron 2 de 2") || strings.Contains(html, `id="votos"`) {
			t.Errorf("en la pestaña de %s se esperaba que votaron los dos, sin ningun voto a la vista:\n%s", caso.suyo, html)
		}
		for _, elemento := range elements(html, "tr", "li", "p", "dd", "td", "span", "h1", "h2") {
			if (strings.Contains(elemento, "Ana Perez") && strings.Contains(elemento, "3")) || (strings.Contains(elemento, "Juan Gomez") && strings.Contains(elemento, "8")) {
				t.Errorf("en la pestaña de %s un voto aparece junto a su integrante: %s", caso.suyo, elemento)
			}
		}
	}

	// La pestaña de Juan, con la pagina vieja, intenta votar otra vez por Ana Perez: el
	// nucleo lo rechaza con su mensaje, sin mostrar el valor, y el voto de Ana no cambia
	// (lo muestra la tabla al revelar).
	res := pestanaDeJuan.post(ronda+"/votar", url.Values{"integrante": {"Ana Perez"}, "puntos": {"21"}})
	rechazo := between(res.Body.String(), "<main", "</main>")
	if res.Code != http.StatusUnprocessableEntity || !strings.Contains(rechazo, `id="error-formulario">`+domain.ErrAlreadyVoted.Error()) || strings.Contains(rechazo, "21") {
		t.Errorf("se esperaba rechazar el segundo voto de Ana Perez con el mensaje del nucleo y sin el valor, salio %d:\n%s", res.Code, rechazo)
	}

	// Ana revela desde su pestaña, y la de Juan ve los votos y la diferencia.
	if res := pestanaDeAna.post(ronda+"/revelar", nil); res.Code != http.StatusSeeOther {
		t.Fatalf("se esperaba revelar con el estado %d, salio %d:\n%s", http.StatusSeeOther, res.Code, res.Body.String())
	}
	html := pestanaDeJuan.get(ronda).Body.String()
	votos := joinTags(between(html, `<table id="votos">`, "</table>"))
	for _, fila := range []string{"<td>Ana Perez</td><td>3</td>", "<td>Juan Gomez</td><td>8</td>"} {
		if !strings.Contains(votos, fila) {
			t.Errorf("en la pestaña de Juan se esperaba el voto %q, y no esta en:\n%s", fila, votos)
		}
	}
	if diferencias := between(html, `<section id="diferencias"`, "</section>"); !strings.Contains(diferencias, "Hay diferencias: el más bajo es 3, de Ana Perez, y el más alto es 8, de Juan Gomez.") {
		t.Errorf("en la pestaña de Juan se esperaba la diferencia, y no esta en:\n%s", html)
	}
}
