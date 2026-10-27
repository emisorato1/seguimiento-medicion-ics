package web

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
	"github.com/emisorato1/seguimiento-medicion-ics/internal/store"
)

// fakeSaver es el guardado falso de los tests (D21): en vez de escribir un archivo, cuenta
// cuantas veces se guardo el registro, y con err simula que guardar falla.
type fakeSaver struct {
	saves int
	err   error
}

func (f *fakeSaver) Save(registry *domain.ProjectRegistry) error {
	f.saves++
	return f.err
}

// testToday es la fecha de hoy de los tests (D1 de la #78): fija, asi los escenarios con
// "hoy es 27/10/2026" se prueban igual cualquier dia.
func testToday() time.Time {
	return time.Date(2026, 10, 27, 0, 0, 0, 0, time.UTC)
}

// newTestServer arma una web nueva sobre el registro que recibe, con un guardado falso.
// Cada test arma el suyo y arranca de cero.
func newTestServer(t *testing.T, registry *domain.ProjectRegistry) *Server {
	t.Helper()
	return NewServer(registry, &fakeSaver{}, testToday)
}

// get pide una pagina a la web, como lo haria el navegador.
func get(t *testing.T, server http.Handler, path string) *httptest.ResponseRecorder {
	t.Helper()
	res := httptest.NewRecorder()
	server.ServeHTTP(res, httptest.NewRequest(http.MethodGet, path, nil))
	return res
}

// addProject registra un proyecto con inicio el 01/11/2026, sin descripcion ni fin, para
// preparar un escenario.
func addProject(t *testing.T, registry *domain.ProjectRegistry, name string) *domain.Project {
	t.Helper()
	project, err := registry.Create(domain.ProjectData{
		Name:  name,
		Start: time.Date(2026, 11, 1, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		t.Fatalf("no se pudo preparar el proyecto %q: %v", name, err)
	}
	return project
}

// addMember agrega un integrante a un proyecto, para preparar un escenario.
func addMember(t *testing.T, project *domain.Project, name string, role domain.Role) {
	t.Helper()
	if _, err := project.AddMember(name, role); err != nil {
		t.Fatalf("no se pudo preparar el integrante %q: %v", name, err)
	}
}

// addStory carga una historia en el backlog de un proyecto, para preparar un escenario.
func addStory(t *testing.T, project *domain.Project, title string, priority domain.Priority) *domain.Story {
	t.Helper()
	story, err := project.AddStory(domain.StoryData{
		Title:              title,
		Description:        "Como integrante quiero " + strings.ToLower(title),
		Priority:           priority,
		AcceptanceCriteria: []string{"Se puede " + strings.ToLower(title)},
	})
	if err != nil {
		t.Fatalf("no se pudo preparar la historia %q: %v", title, err)
	}
	return story
}

// doneStoryRegistry arma, con la conversion de los datos guardados, un registro con el
// proyecto 1 y su historia 1 Terminada: hasta que existan los sprints es la unica forma de
// tener una historia Terminada.
func doneStoryRegistry(t *testing.T) *domain.ProjectRegistry {
	t.Helper()
	registry, err := domain.RegistryFromJSON([]byte(`{
		"next_id": 2,
		"projects": [{
			"id": 1, "name": "Sistema de metricas", "start": "2026-11-01", "next_story_id": 2,
			"members": [],
			"stories": [{
				"id": 1, "title": "Registrar esfuerzo", "description": "Como integrante quiero registrar esfuerzo",
				"priority": "Media", "state": "Terminada", "story_points": 3,
				"acceptance_criteria": ["Se puede registrar esfuerzo"]
			}]
		}]
	}`))
	if err != nil {
		t.Fatalf("no se pudo preparar la historia Terminada: %v", err)
	}
	return registry
}

// joinTags quita los espacios y los saltos de linea entre una etiqueta y la siguiente, para
// buscar una fila sin depender de la sangria ni del fin de linea: en Windows las plantillas
// pueden bajar con CRLF (pedido de Agustin en la revision del #149).
func joinTags(html string) string {
	return betweenTags.ReplaceAllString(html, "><")
}

var betweenTags = regexp.MustCompile(`>\s+<`)

// between devuelve el pedazo del HTML que va de start a end, para mirar solo una parte de
// la pagina, como la barra. Si no lo encuentra devuelve vacio.
func between(html, start, end string) string {
	i := strings.Index(html, start)
	if i < 0 {
		return ""
	}
	j := strings.Index(html[i:], end)
	if j < 0 {
		return ""
	}
	return html[i : i+j]
}

// options devuelve los valores que se pueden elegir en la lista con ese nombre. La opcion
// vacia del principio no es un valor: es el pedido de que se elija uno.
func options(html, name string) []string {
	list := between(html, `<select id="`+name+`"`, "</select>")
	var values []string
	for _, part := range strings.Split(list, `<option value="`)[1:] {
		if value := part[:strings.Index(part, `"`)]; value != "" {
			values = append(values, value)
		}
	}
	return values
}

// scaleText pasa la escala de story points a texto, que es como la muestra la lista.
func scaleText(scale []int) []string {
	var values []string
	for _, points := range scale {
		values = append(values, strconv.Itoa(points))
	}
	return values
}

// snapshot describe en texto todo lo que tiene el registro, para comparar antes y despues
// de una operacion y saber si cambio algun dato.
func snapshot(registry *domain.ProjectRegistry) string {
	var b strings.Builder
	for _, p := range registry.Projects() {
		fmt.Fprintf(&b, "proyecto %d %q %q %v %v\n", p.ID(), p.Name(), p.Description(), p.Start(), p.End())
		for _, m := range p.Members() {
			fmt.Fprintf(&b, "  integrante %q %v\n", m.Name(), m.Role())
		}
		for _, h := range p.Backlog() {
			fmt.Fprintf(&b, "  historia %d %q %q %v %v %d %q\n", h.ID(), h.Title(), h.Description(),
				h.Priority(), h.State(), h.StoryPoints(), h.AcceptanceCriteria())
		}
	}
	// Lo que se guarda incluye lo que suman las historias de la #78 (sprints, esfuerzo,
	// defectos y rondas), asi un cambio en cualquiera de ellos tambien se nota.
	guardado, err := registry.ToJSON()
	if err != nil {
		return "no se pudo describir el registro: " + err.Error()
	}
	b.Write(guardado)
	return b.String()
}

// browser es un navegador de prueba: guarda las cookies que manda la web y las devuelve en
// cada pedido siguiente, como el navegador de verdad. Hace falta para el mensaje de una sola
// vez (D8), que viaja en una cookie.
type browser struct {
	t       *testing.T
	server  http.Handler
	cookies map[string]*http.Cookie
}

func newBrowser(t *testing.T, server http.Handler) *browser {
	return &browser{t: t, server: server, cookies: make(map[string]*http.Cookie)}
}

func (b *browser) get(path string) *httptest.ResponseRecorder {
	b.t.Helper()
	return b.do(httptest.NewRequest(http.MethodGet, path, nil))
}

func (b *browser) post(path string, form url.Values) *httptest.ResponseRecorder {
	b.t.Helper()
	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return b.do(req)
}

func (b *browser) do(req *http.Request) *httptest.ResponseRecorder {
	for _, cookie := range b.cookies {
		req.AddCookie(cookie)
	}
	res := httptest.NewRecorder()
	b.server.ServeHTTP(res, req)
	for _, cookie := range res.Result().Cookies() {
		if cookie.MaxAge < 0 {
			delete(b.cookies, cookie.Name)
		} else {
			b.cookies[cookie.Name] = cookie
		}
	}
	return res
}

// Escenario: Abrir el sistema con un solo comando y sin internet
// Cubre: RN-3, RN-4. Esta es la parte de la web; el comando lo prueba cmd/web.
func TestAbrirElSistemaConUnSoloComandoYSinInternet(t *testing.T) {
	registry := domain.NewProjectRegistry()
	project := addProject(t, registry, "Sistema de metricas")
	addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
	server := newTestServer(t, registry)

	t.Run("la raiz lleva a la lista de proyectos", func(t *testing.T) {
		res := get(t, server, "/")

		if res.Code != http.StatusSeeOther {
			t.Fatalf("se esperaba el estado %d, salio %d", http.StatusSeeOther, res.Code)
		}
		if destino := res.Header().Get("Location"); destino != "/proyectos" {
			t.Fatalf("se esperaba ir a /proyectos, salio %q", destino)
		}
	})

	t.Run("el CSS sale del binario", func(t *testing.T) {
		res := get(t, server, "/static/app.css")

		if res.Code != http.StatusOK {
			t.Fatalf("se esperaba el estado %d, salio %d", http.StatusOK, res.Code)
		}
		if tipo := res.Header().Get("Content-Type"); !strings.HasPrefix(tipo, "text/css") {
			t.Fatalf("se esperaba un CSS, salio %q", tipo)
		}
		if res.Body.Len() == 0 {
			t.Fatal("se esperaba el CSS con contenido, salio vacio")
		}
	})

	// Toda pagina de la web: cada pantalla nueva suma aca su direccion.
	paginas := []string{
		"/proyectos",
		"/proyectos/nuevo",
		"/proyectos/1",
		"/proyectos/1/editar",
		"/proyectos/1/integrantes",
		"/proyectos/1/integrantes/nuevo",
		"/proyectos/1/backlog",
		"/proyectos/1/historias/nueva",
		"/proyectos/1/historias/1",
		"/proyectos/1/historias/1/editar",
		"/proyectos/1/historias/1/estimar",
	}

	for _, pagina := range paginas {
		t.Run(pagina, func(t *testing.T) {
			res := get(t, server, pagina)

			if res.Code != http.StatusOK {
				t.Fatalf("se esperaba el estado %d, salio %d", http.StatusOK, res.Code)
			}
			if tipo := res.Header().Get("Content-Type"); !strings.HasPrefix(tipo, "text/html") {
				t.Fatalf("se esperaba una pagina HTML, salio %q", tipo)
			}

			html := res.Body.String()
			debeTener := map[string]string{
				"en castellano":         `<html lang="es">`,
				"adaptada al ancho":     `<meta name="viewport" content="width=device-width, initial-scale=1">`,
				"con los estilos":       `<link rel="stylesheet" href="/static/app.css">`,
				"con la barra":          `<nav aria-label="Navegación principal">`,
				"con acceso a la lista": `href="/proyectos"`,
			}
			for que, texto := range debeTener {
				if !strings.Contains(html, texto) {
					t.Errorf("se esperaba la pagina %s (%s), y no esta en:\n%s", que, texto, html)
				}
			}

			// Ninguna pagina pide nada de afuera: sin internet tiene que verse completa.
			for _, afuera := range []string{"http://", "https://", `="//`} {
				if strings.Contains(html, afuera) {
					t.Errorf("la pagina no deberia pedir nada de afuera, y tiene %q", afuera)
				}
			}
		})
	}
}

// Esquema del escenario: Las listas vacias dicen que hacer
// Cubre: RN-7, RN-23
func TestLasListasVaciasDicenQueHacer(t *testing.T) {
	casos := []struct {
		nombre   string
		preparar func(t *testing.T, registry *domain.ProjectRegistry)
		pagina   string
		mensaje  string
		acceso   string
	}{
		{
			nombre:   "el sistema sin datos guardados",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {},
			pagina:   "/proyectos",
			mensaje:  "Todavía no hay proyectos",
			acceso:   `href="/proyectos/nuevo"`,
		},
		{
			nombre: "un proyecto sin integrantes",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				addProject(t, registry, "Sistema de metricas")
			},
			pagina:  "/proyectos/1/integrantes",
			mensaje: "Este proyecto todavía no tiene integrantes",
			acceso:  `href="/proyectos/1/integrantes/nuevo"`,
		},
		{
			nombre: "un proyecto sin historias",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				addProject(t, registry, "Sistema de metricas")
			},
			pagina:  "/proyectos/1/backlog",
			mensaje: "El backlog está vacío",
			acceso:  `href="/proyectos/1/historias/nueva"`,
		},
		{
			nombre: "un proyecto sin sprints (#78)",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				addProject(t, registry, "Sistema de metricas")
			},
			pagina:  "/proyectos/1/sprints",
			mensaje: "Todavía no hay sprints",
			acceso:  `href="/proyectos/1/sprints/nuevo"`,
		},
		{
			nombre: "un proyecto sin esfuerzo (#78)",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				addProject(t, registry, "Sistema de metricas")
			},
			pagina:  "/proyectos/1/esfuerzo",
			mensaje: "Todavía no hay horas registradas",
			acceso:  `href="/proyectos/1/esfuerzo/nuevo"`,
		},
		{
			nombre: "un proyecto sin integrantes al registrar esfuerzo (#78)",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				project := addProject(t, registry, "Sistema de metricas")
				addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
			},
			pagina:  "/proyectos/1/esfuerzo/nuevo",
			mensaje: "Este proyecto todavía no tiene integrantes",
			acceso:  `href="/proyectos/1/integrantes/nuevo"`,
		},
		{
			nombre: "un proyecto sin defectos (#78)",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				addProject(t, registry, "Sistema de metricas")
			},
			pagina:  "/proyectos/1/defectos",
			mensaje: "Todavía no hay defectos",
			acceso:  `href="/proyectos/1/defectos/nuevo"`,
		},
		{
			nombre: "un proyecto sin sprints al cargar un defecto (#78, RN-17)",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				project := addProject(t, registry, "Sistema de metricas")
				addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
			},
			pagina:  "/proyectos/1/defectos/nuevo",
			mensaje: "Este proyecto todavía no tiene sprints",
			acceso:  `href="/proyectos/1/sprints/nuevo"`,
		},
	}

	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			registry := domain.NewProjectRegistry()
			caso.preparar(t, registry)
			server := newTestServer(t, registry)

			res := get(t, server, caso.pagina)

			if res.Code != http.StatusOK {
				t.Fatalf("se esperaba el estado %d, salio %d", http.StatusOK, res.Code)
			}
			html := res.Body.String()
			if !strings.Contains(html, caso.mensaje) {
				t.Errorf("se esperaba el mensaje %q, y no esta en:\n%s", caso.mensaje, html)
			}
			if !strings.Contains(html, caso.acceso) {
				t.Errorf("se esperaba el acceso %s, y no esta en:\n%s", caso.acceso, html)
			}
		})
	}
}

// Esquema del escenario: Mostrar el rechazo sin perder lo escrito
// Cubre: RN-12, RN-14, RN-15
func TestMostrarElRechazoSinPerderLoEscrito(t *testing.T) {
	casos := []struct {
		nombre     string
		preparar   func(t *testing.T, registry *domain.ProjectRegistry)
		inicial    func(t *testing.T) *domain.ProjectRegistry
		formulario string
		datos      url.Values
		campo      string // vacio si el mensaje va arriba del formulario
		err        error
		// fueraDeLista son los campos de una lista para elegir que recibieron un valor que
		// la lista no ofrece: no hay opcion que lo muestre, asi que no se puede conservar.
		fueraDeLista []string
	}{
		{
			nombre:     "proyecto con el nombre vacio",
			preparar:   func(t *testing.T, registry *domain.ProjectRegistry) {},
			formulario: "/proyectos/nuevo",
			datos:      url.Values{"nombre": {""}, "descripcion": {"Tablero & métricas"}, "inicio": {"2026-11-01"}},
			campo:      "nombre",
			err:        domain.ErrProjectNameRequired,
		},
		{
			nombre:     "proyecto con la fecha de fin antes que la de inicio",
			preparar:   func(t *testing.T, registry *domain.ProjectRegistry) {},
			formulario: "/proyectos/nuevo",
			datos:      url.Values{"nombre": {"Sistema de metricas"}, "inicio": {"2026-11-01"}, "fin": {"2026-10-01"}},
			campo:      "fin",
			err:        domain.ErrProjectDatesInconsistent,
		},
		{
			nombre:     "proyecto con una fecha de inicio que no es fecha",
			preparar:   func(t *testing.T, registry *domain.ProjectRegistry) {},
			formulario: "/proyectos/nuevo",
			datos:      url.Values{"nombre": {"Sistema de metricas"}, "inicio": {"31/02/2026"}},
			campo:      "inicio",
			err:        errInvalidDate,
		},
		{
			// Si la fecha mal escrita llegara al dominio como "sin fecha", el proyecto se
			// crearia sin fecha de fin: por eso este caso prueba que no se lo llama.
			nombre:     "proyecto con una fecha de fin que no es fecha",
			preparar:   func(t *testing.T, registry *domain.ProjectRegistry) {},
			formulario: "/proyectos/nuevo",
			datos:      url.Values{"nombre": {"Sistema de metricas"}, "inicio": {"2026-11-01"}, "fin": {"el lunes"}},
			campo:      "fin",
			err:        errInvalidDate,
		},
		{
			nombre: "integrante que ya esta cargado",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				project := addProject(t, registry, "Sistema de metricas")
				addMember(t, project, "Ana Perez", domain.RoleProductBuilder)
			},
			formulario: "/proyectos/1/integrantes/nuevo",
			datos:      url.Values{"nombre": {"ana perez"}, "rol": {"Product Builder"}},
			campo:      "nombre",
			err:        domain.ErrMemberDuplicate,
		},
		{
			nombre: "segundo Agile Enabler",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				project := addProject(t, registry, "Sistema de metricas")
				addMember(t, project, "Ana Perez", domain.RoleAgileEnabler)
			},
			formulario: "/proyectos/1/integrantes/nuevo",
			datos:      url.Values{"nombre": {"Luis Gomez"}, "rol": {"Agile Enabler"}},
			campo:      "rol",
			err:        domain.ErrMemberAgileEnablerTaken,
		},
		{
			nombre: "historia con el titulo vacio",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				addProject(t, registry, "Sistema de metricas")
			},
			formulario: "/proyectos/1/historias/nueva",
			datos: url.Values{
				"titulo":      {""},
				"descripcion": {"Como integrante quiero registrar mis horas"},
				"prioridad":   {"Media"},
				"criterios":   {"Se cargan las horas"},
			},
			campo: "titulo",
			err:   domain.ErrStoryTitleRequired,
		},
		{
			nombre: "historia con los criterios en blanco",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				addProject(t, registry, "Sistema de metricas")
			},
			formulario: "/proyectos/1/historias/nueva",
			datos: url.Values{
				"titulo":      {"Registrar esfuerzo"},
				"descripcion": {"Como integrante quiero registrar mis horas"},
				"prioridad":   {"Media"},
				"criterios":   {"  \n   \n"},
			},
			campo: "criterios",
			err:   domain.ErrStoryCriteriaRequired,
		},
		{
			// La lista no ofrece el 4: se envia igual, y lo rechaza el dominio (punto 5).
			nombre: "estimacion fuera de la escala",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				project := addProject(t, registry, "Sistema de metricas")
				addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
			},
			formulario:   "/proyectos/1/historias/1/estimar",
			datos:        url.Values{"puntos": {"4"}},
			campo:        "puntos",
			err:          domain.ErrStoryPointsInvalid,
			fueraDeLista: []string{"puntos"},
		},
		{
			// Sin elegir de la lista, o con letras: es la forma del dato, como una fecha
			// que no es fecha, y no se llama al dominio.
			nombre: "estimacion sin elegir",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				project := addProject(t, registry, "Sistema de metricas")
				addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
			},
			formulario: "/proyectos/1/historias/1/estimar",
			datos:      url.Values{"puntos": {""}},
			campo:      "puntos",
			err:        errInvalidPoints,
		},
		{
			nombre: "estimacion con letras",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				project := addProject(t, registry, "Sistema de metricas")
				addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
			},
			formulario:   "/proyectos/1/historias/1/estimar",
			datos:        url.Values{"puntos": {"cinco"}},
			campo:        "puntos",
			err:          errInvalidPoints,
			fueraDeLista: []string{"puntos"},
		},
		{
			// Las tres que siguen no estan en los Ejemplos: los mismos rechazos al modificar
			// un proyecto y al quitar una estimacion.
			nombre: "modificar un proyecto con el nombre vacio",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				addProject(t, registry, "Sistema de metricas")
			},
			formulario: "/proyectos/1/editar",
			datos:      url.Values{"nombre": {""}, "inicio": {"2026-11-01"}},
			campo:      "nombre",
			err:        domain.ErrProjectNameRequired,
		},
		{
			nombre: "modificar un proyecto con una fecha que no es fecha",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				addProject(t, registry, "Sistema de metricas")
			},
			formulario: "/proyectos/1/editar",
			datos:      url.Values{"nombre": {"Metricas ICS"}, "inicio": {"2026-11-01"}, "fin": {"pronto"}},
			campo:      "fin",
			err:        errInvalidDate,
		},
		{
			nombre:     "quitar la estimacion de una historia Terminada",
			inicial:    doneStoryRegistry,
			formulario: "/proyectos/1/historias/1/quitar-estimacion",
			datos:      url.Values{},
			campo:      "",
			err:        domain.ErrStoryEstimateDone,
		},
		{
			nombre:     "modificar una historia Terminada",
			inicial:    doneStoryRegistry,
			formulario: "/proyectos/1/historias/1/editar",
			datos: url.Values{
				"titulo":      {"Otro titulo"},
				"descripcion": {"Como integrante quiero registrar esfuerzo"},
				"prioridad":   {"Media"},
				"criterios":   {"Se puede registrar esfuerzo"},
			},
			campo: "",
			err:   domain.ErrStoryDone,
		},
		{
			nombre:     "estimar una historia Terminada",
			inicial:    doneStoryRegistry,
			formulario: "/proyectos/1/historias/1/estimar",
			datos:      url.Values{"puntos": {"5"}},
			campo:      "",
			err:        domain.ErrStoryEstimateDone,
		},
		{
			nombre: "sprint con el goal vacio (#78)",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				addProject(t, registry, "Sistema de metricas")
			},
			formulario: "/proyectos/1/sprints/nuevo",
			datos:      url.Values{"goal": {"  "}, "inicio": {"2026-10-26"}, "fin": {"2026-10-30"}},
			campo:      "goal",
			err:        domain.ErrSprintGoalRequired,
		},
		{
			nombre: "segundo sprint abierto (#78)",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				addProject(t, registry, "Sistema de metricas")
				addSprint(t, registry, 1)
			},
			formulario: "/proyectos/1/sprints/nuevo",
			datos:      url.Values{"goal": {"Cerrar el MVP"}, "inicio": {"2026-11-02"}, "fin": {"2026-11-06"}},
			campo:      "",
			err:        domain.ErrSprintAlreadyOpen,
		},
		{
			nombre: "asignar una historia sin estimar (#78)",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				project := addProject(t, registry, "Sistema de metricas")
				addStory(t, project, "Sin estimar", domain.PriorityLow)
				addSprint(t, registry, 1)
			},
			formulario: "/proyectos/1/sprints/asignar",
			datos:      url.Values{"historia": {"1"}},
			campo:      "historia",
			err:        domain.ErrStoryNotEstimated,
		},
		{
			nombre: "asignar sin elegir la historia (#78)",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				project := addProject(t, registry, "Sistema de metricas")
				addStory(t, project, "Sin estimar", domain.PriorityLow)
				addSprint(t, registry, 1)
			},
			formulario: "/proyectos/1/sprints/asignar",
			datos:      url.Values{"historia": {""}},
			campo:      "historia",
			err:        errInvalidStory,
		},
		{
			nombre: "esfuerzo de un dia que todavia no paso (#78)",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				project := addProject(t, registry, "Sistema de metricas")
				addMember(t, project, "Ana Perez", domain.RoleProductBuilder)
				addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
			},
			formulario: "/proyectos/1/esfuerzo/nuevo",
			datos:      url.Values{"integrante": {"Ana Perez"}, "historia": {"1"}, "dia": {"2026-10-28"}, "horas": {"2"}, "actividad": {"Escribir tests"}},
			campo:      "dia",
			err:        domain.ErrEffortDayInFuture,
		},
		{
			nombre: "esfuerzo de cero horas (#78)",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				project := addProject(t, registry, "Sistema de metricas")
				addMember(t, project, "Ana Perez", domain.RoleProductBuilder)
				addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
			},
			formulario: "/proyectos/1/esfuerzo/nuevo",
			datos:      url.Values{"integrante": {"Ana Perez"}, "historia": {"1"}, "dia": {"2026-10-27"}, "horas": {"0"}, "actividad": {"Escribir tests"}},
			campo:      "horas",
			err:        domain.ErrEffortHoursInvalid,
		},
		{
			nombre: "esfuerzo con horas que no son numero (#78)",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				project := addProject(t, registry, "Sistema de metricas")
				addMember(t, project, "Ana Perez", domain.RoleProductBuilder)
				addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
			},
			formulario: "/proyectos/1/esfuerzo/nuevo",
			datos:      url.Values{"integrante": {"Ana Perez"}, "historia": {"1"}, "dia": {"2026-10-27"}, "horas": {"dos"}, "actividad": {"Escribir tests"}},
			campo:      "horas",
			err:        errInvalidHours,
		},
		{
			nombre: "esfuerzo con NaN horas (#78)",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				project := addProject(t, registry, "Sistema de metricas")
				addMember(t, project, "Ana Perez", domain.RoleProductBuilder)
				addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
			},
			formulario: "/proyectos/1/esfuerzo/nuevo",
			datos:      url.Values{"integrante": {"Ana Perez"}, "historia": {"1"}, "dia": {"2026-10-27"}, "horas": {"NaN"}, "actividad": {"Escribir tests"}},
			campo:      "horas",
			err:        errInvalidHours,
		},
		{
			nombre: "esfuerzo con horas infinitas (#78)",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				project := addProject(t, registry, "Sistema de metricas")
				addMember(t, project, "Ana Perez", domain.RoleProductBuilder)
				addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
			},
			formulario: "/proyectos/1/esfuerzo/nuevo",
			datos:      url.Values{"integrante": {"Ana Perez"}, "historia": {"1"}, "dia": {"2026-10-27"}, "horas": {"Inf"}, "actividad": {"Escribir tests"}},
			campo:      "horas",
			err:        errInvalidHours,
		},
		{
			nombre: "horas estimadas en cero (#78)",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				project := addProject(t, registry, "Sistema de metricas")
				addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
			},
			formulario: "/proyectos/1/historias/1/horas-estimadas",
			datos:      url.Values{"horas": {"0"}},
			campo:      "horas",
			err:        domain.ErrEstimatedHoursInvalid,
		},
		{
			// No esta en los Ejemplos: la historia que no existe la rechaza el nucleo (RN-2).
			nombre: "abrir una ronda de una historia que no existe (#78)",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				project := addProject(t, registry, "Sistema de metricas")
				addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
			},
			formulario:   "/proyectos/1/planning-poker",
			datos:        url.Values{"historia": {"99"}},
			campo:        "historia",
			err:          domain.ErrStoryNotFound,
			fueraDeLista: []string{"historia"},
		},
		{
			// No esta en los Ejemplos: la fecha de cierre la controla el nucleo (#10).
			nombre:     "cerrar el sprint sin fecha de cierre (#78)",
			inicial:    func(t *testing.T) *domain.ProjectRegistry { registry, _ := closeRegistry(t); return registry },
			formulario: "/proyectos/1/sprints/cerrar",
			datos:      url.Values{"cierre": {""}, "completadas": {"1"}},
			campo:      "cierre",
			err:        domain.ErrSprintCloseDateRequired,
		},
		{
			// No esta en los Ejemplos: una fecha que no es fecha no llega al nucleo (RN-14 de la #77).
			nombre:     "cerrar el sprint con una fecha que no es fecha (#78)",
			inicial:    func(t *testing.T) *domain.ProjectRegistry { registry, _ := closeRegistry(t); return registry },
			formulario: "/proyectos/1/sprints/cerrar",
			datos:      url.Values{"cierre": {"el viernes"}, "completadas": {"1"}},
			campo:      "cierre",
			err:        errInvalidDate,
		},
		{
			// No esta en los Ejemplos: algo que no viene de las casillas va arriba.
			nombre:     "cerrar el sprint con una historia que no es un numero (#78)",
			inicial:    func(t *testing.T) *domain.ProjectRegistry { registry, _ := closeRegistry(t); return registry },
			formulario: "/proyectos/1/sprints/cerrar",
			datos:      url.Values{"cierre": {"2026-10-30"}, "completadas": {"uno"}},
			err:        errInvalidStory,
			// las completadas viajan en un campo oculto de la confirmacion, no en el formulario
			fueraDeLista: []string{"completadas"},
		},
		{
			// No esta en los Ejemplos: acordar sin elegir es un dato sin forma (RN-2).
			nombre: "acordar sin elegir los story points (#78)",
			inicial: func(t *testing.T) *domain.ProjectRegistry {
				registry, project := pokerRegistry(t)
				revealedRound(t, project)
				return registry
			},
			formulario: "/proyectos/1/historias/1/rondas/1/acordar",
			datos:      url.Values{"puntos": {""}},
			campo:      "puntos",
			err:        errInvalidPoints,
		},
		{
			// No esta en los Ejemplos: acordar antes de revelar lo rechaza el nucleo (#73).
			nombre: "acordar una ronda sin revelar (#78)",
			inicial: func(t *testing.T) *domain.ProjectRegistry {
				registry, project := pokerRegistry(t)
				if _, err := project.OpenVotingRound(1); err != nil {
					t.Fatalf("no se pudo abrir la ronda: %v", err)
				}
				return registry
			},
			formulario: "/proyectos/1/historias/1/rondas/1/acordar",
			datos:      url.Values{"puntos": {"5"}},
			err:        domain.ErrRoundNotRevealed,
		},
		{
			nombre:     "defecto con el titulo vacio (#78)",
			inicial:    func(t *testing.T) *domain.ProjectRegistry { registry, _ := defectRegistry(t); return registry },
			formulario: "/proyectos/1/defectos/nuevo",
			datos:      url.Values{"titulo": {"  "}, "descripcion": {"No suma el ultimo"}, "severidad": {"Mayor"}, "historia": {"1"}, "sprint": {"1"}},
			campo:      "titulo",
			err:        domain.ErrDefectTitleRequired,
		},
		{
			// No esta en los Ejemplos: la severidad sin elegir la rechaza el nucleo (RN-2).
			nombre:     "defecto sin elegir la severidad (#78)",
			inicial:    func(t *testing.T) *domain.ProjectRegistry { registry, _ := defectRegistry(t); return registry },
			formulario: "/proyectos/1/defectos/nuevo",
			datos:      url.Values{"titulo": {"Suma mal las horas"}, "descripcion": {"No suma el ultimo"}, "severidad": {""}, "historia": {"1"}, "sprint": {"1"}},
			campo:      "severidad",
			err:        domain.ErrDefectSeverityInvalid,
		},
		{
			// No esta en los Ejemplos: el sprint sin elegir es un dato sin forma (RN-2).
			nombre:     "defecto sin elegir el sprint (#78)",
			inicial:    func(t *testing.T) *domain.ProjectRegistry { registry, _ := defectRegistry(t); return registry },
			formulario: "/proyectos/1/defectos/nuevo",
			datos:      url.Values{"titulo": {"Suma mal las horas"}, "descripcion": {"No suma el ultimo"}, "severidad": {"Mayor"}, "historia": {"1"}, "sprint": {""}},
			campo:      "sprint",
			err:        errInvalidSprint,
		},
		{
			// No esta en los Ejemplos: el estado sin elegir lo rechaza el nucleo (RN-2).
			nombre: "cambiar el estado sin elegirlo (#78)",
			inicial: func(t *testing.T) *domain.ProjectRegistry {
				registry, project := defectRegistry(t)
				addDefect(t, project, 1)
				return registry
			},
			formulario: "/proyectos/1/defectos/1/estado",
			datos:      url.Values{"estado": {""}},
			campo:      "estado",
			err:        domain.ErrDefectStateUnknown,
		},
		{
			nombre: "resolver en un sprint anterior al de deteccion (#78)",
			inicial: func(t *testing.T) *domain.ProjectRegistry {
				registry, project := defectRegistry(t)
				addDefect(t, project, 2)
				return registry
			},
			formulario: "/proyectos/1/defectos/1/resolver",
			datos:      url.Values{"sprint": {"1"}},
			campo:      "sprint",
			err:        domain.ErrDefectResolvedBeforeDetected,
		},
	}

	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			registry := domain.NewProjectRegistry()
			if caso.inicial != nil {
				registry = caso.inicial(t)
			}
			if caso.preparar != nil {
				caso.preparar(t, registry)
			}
			saver := &fakeSaver{}
			server := NewServer(registry, saver, testToday)
			antes := snapshot(registry)

			res := newBrowser(t, server).post(caso.formulario, caso.datos)

			if res.Code != http.StatusUnprocessableEntity {
				t.Fatalf("se esperaba el estado %d, salio %d", http.StatusUnprocessableEntity, res.Code)
			}
			html := res.Body.String()

			// El formulario vuelve con todo lo que se habia escrito.
			for campo, valores := range caso.datos {
				if slices.Contains(caso.fueraDeLista, campo) {
					continue
				}
				if valor := template.HTMLEscapeString(valores[0]); !strings.Contains(html, valor) {
					t.Errorf("se esperaba que el campo %s conservara %q, y no esta en:\n%s", campo, valor, html)
				}
			}

			// El mensaje del dominio, junto a su campo o, si no es de ningun campo, arriba.
			donde := "arriba del formulario"
			// Sin el cierre del elemento: el dominio puede sumarle un detalle al mensaje, como
			// el numero del sprint que ya esta abierto (#8).
			junto := fmt.Sprintf(`id="error-formulario">%s`, template.HTMLEscapeString(caso.err.Error()))
			if caso.campo != "" {
				donde = "junto al campo " + caso.campo
				junto = fmt.Sprintf(`id="%s-error">%s`, caso.campo, template.HTMLEscapeString(caso.err.Error()))
			}
			if !strings.Contains(html, junto) {
				t.Errorf("se esperaba %s el mensaje %q, y no esta en:\n%s", donde, caso.err, html)
			}

			// No cambio ningun dato.
			if despues := snapshot(registry); despues != antes {
				t.Errorf("no se esperaba ningun cambio, y el registro paso de:\n%s\na:\n%s", antes, despues)
			}
			if saver.saves != 0 {
				t.Errorf("no se esperaba que se guardara nada, se guardo %d veces", saver.saves)
			}
		})
	}
}

// Esquema del escenario: Cada operacion se confirma con un mensaje
// Cubre: RN-1, RN-8
func TestCadaOperacionSeConfirmaConUnMensaje(t *testing.T) {
	casos := []struct {
		nombre     string
		preparar   func(t *testing.T, registry *domain.ProjectRegistry)
		formulario string
		precargado []string
		datos      url.Values
		mensaje    string
		figura     []string
		verificar  func(t *testing.T, registry *domain.ProjectRegistry)
	}{
		{
			nombre: "agregar un integrante",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				addProject(t, registry, "Sistema de metricas")
			},
			formulario: "/proyectos/1/integrantes/nuevo",
			datos:      url.Values{"nombre": {"Ana Perez"}, "rol": {"Product Builder"}},
			mensaje:    "Se agregó a Ana Perez",
			figura:     []string{"<td>Product Builder</td>"},
		},
		{
			nombre: "cargar una historia",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				addProject(t, registry, "Sistema de metricas")
			},
			formulario: "/proyectos/1/historias/nueva",
			datos: url.Values{
				"titulo":      {"Registrar esfuerzo"},
				"descripcion": {"Como integrante quiero registrar mis horas"},
				"prioridad":   {"Alta"},
				"criterios":   {"Se cargan las horas\r\n\r\nSe ve el total"},
			},
			mensaje: "Se cargó la historia 1",
			figura:  []string{"<td>Pendiente</td>", "sin estimar"},
			verificar: func(t *testing.T, registry *domain.ProjectRegistry) {
				project, _ := registry.Get(1)
				story, err := project.Story(1)
				if err != nil {
					t.Fatalf("se esperaba la historia 1 cargada: %v", err)
				}
				// Un criterio por renglon; el renglon en blanco lo ignora el dominio.
				criterios := []string{"Se cargan las horas", "Se ve el total"}
				if !slices.Equal(story.AcceptanceCriteria(), criterios) {
					t.Errorf("se esperaban los criterios %q, salieron %q", criterios, story.AcceptanceCriteria())
				}
			},
		},
		{
			nombre: "cambiar la prioridad de una historia",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				project := addProject(t, registry, "Sistema de metricas")
				addStory(t, project, "Registrar esfuerzo", domain.PriorityMedium)
			},
			formulario: "/proyectos/1/historias/1/editar",
			precargado: []string{`value="Registrar esfuerzo"`, `<option value="Media" selected>`, ">Se puede registrar esfuerzo</textarea>"},
			datos: url.Values{
				"titulo":      {"Registrar esfuerzo"},
				"descripcion": {"Como integrante quiero registrar esfuerzo"},
				"prioridad":   {"Alta"},
				"criterios":   {"Se puede registrar esfuerzo"},
			},
			mensaje: "Se modificó la historia 1",
			figura:  []string{"<dd>Alta</dd>"},
		},
		{
			nombre: "estimar una historia",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				project := addProject(t, registry, "Sistema de metricas")
				addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
			},
			formulario: "/proyectos/1/historias/1/estimar",
			datos:      url.Values{"puntos": {"5"}},
			mensaje:    "Se estimó la historia 1 en 5 story points",
			figura:     []string{"<td>5</td>"},
		},
		{
			nombre: "cambiar el nombre de un proyecto",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				addProject(t, registry, "Sistema de metricas")
			},
			formulario: "/proyectos/1/editar",
			precargado: []string{`value="Sistema de metricas"`, `value="2026-11-01"`},
			datos:      url.Values{"nombre": {"Metricas ICS"}, "inicio": {"2026-11-01"}},
			mensaje:    "Se modificó el proyecto Metricas ICS",
			figura:     []string{"<h1>Metricas ICS</h1>"},
		},
		{
			nombre: "borrar la descripcion de un proyecto",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				_, err := registry.Create(domain.ProjectData{
					Name:        "Sistema de metricas",
					Description: "Tablero del curso",
					Start:       time.Date(2026, 11, 1, 0, 0, 0, 0, time.UTC),
				})
				if err != nil {
					t.Fatalf("no se pudo preparar el proyecto: %v", err)
				}
			},
			formulario: "/proyectos/1/editar",
			precargado: []string{">Tablero del curso</textarea>"},
			datos:      url.Values{"nombre": {"Sistema de metricas"}, "descripcion": {""}, "inicio": {"2026-11-01"}},
			mensaje:    "Se modificó el proyecto Sistema de metricas",
			figura:     []string{"sin descripción"},
		},
	}

	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			registry := domain.NewProjectRegistry()
			caso.preparar(t, registry)
			navegador := newBrowser(t, newTestServer(t, registry))

			// El formulario se abre con lo que ya tiene cargado.
			res := navegador.get(caso.formulario)
			if res.Code != http.StatusOK {
				t.Fatalf("se esperaba abrir el formulario con el estado %d, salio %d", http.StatusOK, res.Code)
			}
			for _, texto := range caso.precargado {
				if html := res.Body.String(); !strings.Contains(html, texto) {
					t.Errorf("se esperaba el formulario precargado con %s, y no esta en:\n%s", texto, html)
				}
			}

			res = navegador.post(caso.formulario, caso.datos)
			if res.Code != http.StatusSeeOther {
				t.Fatalf("se esperaba el estado %d, salio %d:\n%s", http.StatusSeeOther, res.Code, res.Body.String())
			}

			html := between(navegador.get(res.Header().Get("Location")).Body.String(), "<main", "</main>")
			if !strings.Contains(html, caso.mensaje) {
				t.Errorf("se esperaba el mensaje %q, y no esta en:\n%s", caso.mensaje, html)
			}
			for _, texto := range caso.figura {
				if !strings.Contains(html, texto) {
					t.Errorf("se esperaba que figure %s, y no esta en:\n%s", texto, html)
				}
			}
			if caso.verificar != nil {
				caso.verificar(t, registry)
			}
		})
	}
}

// Esquema del escenario: Pedir algo que no existe
// Cubre: RN-16
func TestPedirAlgoQueNoExiste(t *testing.T) {
	casos := []struct {
		pedido string
		err    error
		volver string // el acceso al proyecto, si el pedido es de algo dentro de uno
	}{
		{pedido: "/proyectos/99", err: domain.ErrProjectNotFound},
		{pedido: "/proyectos/abc", err: domain.ErrProjectNotFound},
		{pedido: "/proyectos/1/historias/99", err: domain.ErrStoryNotFound},
		// No esta en los Ejemplos: una direccion que no es de ninguna pagina tambien se
		// informa en castellano, con el mismo acceso a la lista.
		{pedido: "/proyectos/1/cualquiera", err: errPageNotFound},
		// Tampoco estan en los Ejemplos: la misma regla en cada pantalla del proyecto.
		{pedido: "/proyectos/99/editar", err: domain.ErrProjectNotFound},
		{pedido: "/proyectos/99/integrantes", err: domain.ErrProjectNotFound},
		{pedido: "/proyectos/99/integrantes/nuevo", err: domain.ErrProjectNotFound},
		{pedido: "/proyectos/99/backlog", err: domain.ErrProjectNotFound},
		{pedido: "/proyectos/99/historias/nueva", err: domain.ErrProjectNotFound},
		{pedido: "/proyectos/1/historias/99/editar", err: domain.ErrStoryNotFound},
		{pedido: "/proyectos/1/historias/99/estimar", err: domain.ErrStoryNotFound},
		{pedido: "/proyectos/1/historias/99/quitar-estimacion", err: domain.ErrStoryNotFound},
		{pedido: "/proyectos/1/integrantes/quitar?nombre=Nadie", err: domain.ErrMemberNotFound},
		// La #78 (RN-18) pide ademas el acceso para volver al proyecto.
		{pedido: "/proyectos/1/sprints/99", err: domain.ErrSprintNotFound, volver: "/proyectos/1"},
		{pedido: "/proyectos/99/sprints", err: domain.ErrProjectNotFound},
		{pedido: "/proyectos/99/sprints/nuevo", err: domain.ErrProjectNotFound},
		{pedido: "/proyectos/99/sprints/1", err: domain.ErrProjectNotFound},
		{pedido: "/proyectos/1/sprints/quitar?historia=1", err: domain.ErrNoOpenSprint, volver: "/proyectos/1"},
		{pedido: "/proyectos/99/esfuerzo", err: domain.ErrProjectNotFound},
		{pedido: "/proyectos/99/esfuerzo/nuevo", err: domain.ErrProjectNotFound},
		{pedido: "/proyectos/1/historias/99/horas-estimadas", err: domain.ErrStoryNotFound, volver: "/proyectos/1"},
		{pedido: "/proyectos/99/planning-poker", err: domain.ErrProjectNotFound},
		{pedido: "/proyectos/1/historias/99/rondas/1", err: domain.ErrStoryNotFound, volver: "/proyectos/1"},
		{pedido: "/proyectos/1/historias/1/rondas/1", err: errRoundNotFound, volver: "/proyectos/1"},
		{pedido: "/proyectos/1/historias/99/rondas", err: domain.ErrStoryNotFound, volver: "/proyectos/1"},
		{pedido: "/proyectos/1/sprints/cerrar", err: domain.ErrNoOpenSprint, volver: "/proyectos/1"},
		// La confirmacion del cierre sin sprint abierto ya no es una pagina de "No se encontro":
		// es una pagina vieja, y avisa en el formulario de cerrar (defecto #190).
		{pedido: "/proyectos/99/historias/1/rondas", err: domain.ErrProjectNotFound},
		{pedido: "/proyectos/1/defectos/99", err: domain.ErrDefectNotFound, volver: "/proyectos/1"},
		{pedido: "/proyectos/1/defectos/99/resolver?sprint=1", err: domain.ErrDefectNotFound, volver: "/proyectos/1"},
		{pedido: "/proyectos/99/defectos", err: domain.ErrProjectNotFound},
		{pedido: "/proyectos/99/defectos/nuevo", err: domain.ErrProjectNotFound},
	}

	for _, caso := range casos {
		t.Run(caso.pedido, func(t *testing.T) {
			registry := domain.NewProjectRegistry()
			project := addProject(t, registry, "Sistema de metricas")
			// La historia 1 existe, sin rondas, para pedir una ronda que no tiene (#78).
			addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
			server := newTestServer(t, registry)

			res := get(t, server, caso.pedido)

			if res.Code != http.StatusNotFound {
				t.Fatalf("se esperaba el estado %d, salio %d", http.StatusNotFound, res.Code)
			}
			contenido := between(res.Body.String(), "<main", "</main>")
			if mensaje := template.HTMLEscapeString(caso.err.Error()); !strings.Contains(contenido, mensaje) {
				t.Errorf("se esperaba el mensaje %q, y no esta en:\n%s", mensaje, contenido)
			}
			if !strings.Contains(contenido, `href="/proyectos"`) {
				t.Errorf("se esperaba el acceso para volver a la lista de proyectos, y no esta en:\n%s", contenido)
			}
			if caso.volver != "" && !strings.Contains(contenido, `href="`+caso.volver+`"`) {
				t.Errorf("se esperaba el acceso para volver al proyecto (%s), y no esta en:\n%s", caso.volver, contenido)
			}
		})
	}
}

// Esquema del escenario: Las opciones para elegir son las del sistema
// Cubre: RN-11
func TestLasOpcionesParaElegirSonLasDelSistema(t *testing.T) {
	casos := []struct {
		formulario string
		pagina     string
		lista      string
		opciones   []string
	}{
		{
			formulario: "integrante",
			pagina:     "/proyectos/1/integrantes/nuevo",
			lista:      "rol",
			opciones:   []string{domain.RoleAgileEnabler.String(), domain.RoleProductBuilder.String()},
		},
		{
			formulario: "historia",
			pagina:     "/proyectos/1/historias/nueva",
			lista:      "prioridad",
			opciones:   []string{domain.PriorityHigh.String(), domain.PriorityMedium.String(), domain.PriorityLow.String()},
		},
		{
			formulario: "estimacion",
			pagina:     "/proyectos/1/historias/1/estimar",
			lista:      "puntos",
			opciones:   scaleText(domain.StoryPointScale()),
		},
	}

	for _, caso := range casos {
		t.Run(caso.formulario, func(t *testing.T) {
			registry := domain.NewProjectRegistry()
			project := addProject(t, registry, "Sistema de metricas")
			addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
			server := newTestServer(t, registry)

			res := get(t, server, caso.pagina)

			if res.Code != http.StatusOK {
				t.Fatalf("se esperaba el estado %d, salio %d", http.StatusOK, res.Code)
			}
			if salio := options(res.Body.String(), caso.lista); !slices.Equal(salio, caso.opciones) {
				t.Errorf("se esperaban las opciones de %s %v, salieron %v", caso.lista, caso.opciones, salio)
			}
		})
	}
}

// Esquema del escenario: Confirmar o cancelar antes de quitar
// Cubre: RN-8, RN-10
func TestConfirmarOCancelarAntesDeQuitar(t *testing.T) {
	casos := []struct {
		nombre    string
		pedido    string
		confirmar string
		datos     url.Values
		cancelar  string
		confirmo  bool
		mensaje   string
		pagina    string
		figura    string
		queda     bool
	}{
		{
			nombre:    "quitar a Ana Perez y confirmar",
			pedido:    "/proyectos/1/integrantes/quitar?nombre=Ana+Perez",
			confirmar: "/proyectos/1/integrantes/quitar",
			datos:     url.Values{"nombre": {"Ana Perez"}},
			cancelar:  "/proyectos/1/integrantes",
			confirmo:  true,
			mensaje:   "Se quitó a Ana Perez",
			pagina:    "/proyectos/1/integrantes",
			figura:    "<td>Ana Perez</td>",
			queda:     false,
		},
		{
			nombre:    "quitar a Ana Perez y cancelar",
			pedido:    "/proyectos/1/integrantes/quitar?nombre=Ana+Perez",
			confirmar: "/proyectos/1/integrantes/quitar",
			datos:     url.Values{"nombre": {"Ana Perez"}},
			cancelar:  "/proyectos/1/integrantes",
			confirmo:  false,
			pagina:    "/proyectos/1/integrantes",
			figura:    "<td>Ana Perez</td>",
			queda:     true,
		},
		{
			nombre:    "quitar la estimacion de la 1 y confirmar",
			pedido:    "/proyectos/1/historias/1/quitar-estimacion",
			confirmar: "/proyectos/1/historias/1/quitar-estimacion",
			datos:     url.Values{},
			cancelar:  "/proyectos/1/historias/1",
			confirmo:  true,
			mensaje:   "Se quitó la estimación de la historia 1",
			pagina:    "/proyectos/1/backlog",
			figura:    "sin estimar",
			queda:     true,
		},
		{
			nombre:    "quitar la estimacion de la 1 y cancelar",
			pedido:    "/proyectos/1/historias/1/quitar-estimacion",
			confirmar: "/proyectos/1/historias/1/quitar-estimacion",
			datos:     url.Values{},
			cancelar:  "/proyectos/1/historias/1",
			confirmo:  false,
			pagina:    "/proyectos/1/backlog",
			figura:    "<td>5</td>",
			queda:     true,
		},
	}

	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			// Dado el integrante "Ana Perez" y la historia 1 estimada en 5.
			registry := domain.NewProjectRegistry()
			project := addProject(t, registry, "Sistema de metricas")
			addMember(t, project, "Ana Perez", domain.RoleProductBuilder)
			addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
			if err := project.EstimateStory(1, 5); err != nil {
				t.Fatalf("no se pudo preparar la estimacion: %v", err)
			}
			saver := &fakeSaver{}
			navegador := newBrowser(t, NewServer(registry, saver, testToday))
			antes := snapshot(registry)

			// Pedir que se quite no quita nada: muestra la confirmacion.
			res := navegador.get(caso.pedido)
			if res.Code != http.StatusOK {
				t.Fatalf("se esperaba la confirmacion con el estado %d, salio %d", http.StatusOK, res.Code)
			}
			confirmacion := between(res.Body.String(), "<main", "</main>")
			for _, texto := range []string{`action="` + caso.confirmar + `"`, `href="` + caso.cancelar + `"`} {
				if !strings.Contains(confirmacion, texto) {
					t.Errorf("se esperaba en la confirmacion %s, y no esta en:\n%s", texto, confirmacion)
				}
			}
			if despues := snapshot(registry); despues != antes {
				t.Fatalf("pedir que se quite no deberia cambiar nada, y el registro paso de:\n%s\na:\n%s", antes, despues)
			}

			if caso.confirmo {
				res = navegador.post(caso.confirmar, caso.datos)
				if res.Code != http.StatusSeeOther {
					t.Fatalf("se esperaba el estado %d, salio %d:\n%s", http.StatusSeeOther, res.Code, res.Body.String())
				}
			} else {
				navegador.get(caso.cancelar)
				if despues := snapshot(registry); despues != antes || saver.saves != 0 {
					t.Errorf("cancelar no deberia cambiar nada, y el registro paso de:\n%s\na:\n%s", antes, despues)
				}
			}

			html := between(navegador.get(caso.pagina).Body.String(), "<main", "</main>")
			if caso.mensaje != "" && !strings.Contains(html, caso.mensaje) {
				t.Errorf("se esperaba el mensaje %q, y no esta en:\n%s", caso.mensaje, html)
			}
			if strings.Contains(html, caso.figura) != caso.queda {
				t.Errorf("se esperaba que %s figure: %v, y la pagina es:\n%s", caso.figura, caso.queda, html)
			}
		})
	}
}

// Escenario: Un cambio que no se pudo guardar se avisa
// Cubre: RN-25
func TestUnCambioQueNoSePudoGuardarSeAvisa(t *testing.T) {
	registry := domain.NewProjectRegistry()
	saver := &fakeSaver{err: errors.New("el cambio no quedó guardado: permiso denegado")}
	navegador := newBrowser(t, NewServer(registry, saver, testToday))

	res := navegador.post("/proyectos/nuevo", url.Values{
		"nombre": {"Sistema de metricas"},
		"inicio": {"2026-11-01"},
	})

	if res.Code != http.StatusInternalServerError {
		t.Fatalf("se esperaba el estado %d, salio %d", http.StatusInternalServerError, res.Code)
	}
	contenido := between(res.Body.String(), "<main", "</main>")
	if aviso := template.HTMLEscapeString(saver.err.Error()); !strings.Contains(contenido, aviso) {
		t.Errorf("se esperaba el aviso %q, y no esta en:\n%s", aviso, contenido)
	}
	// El cambio queda en memoria y lo incluye el siguiente guardado que funcione (D19).
	if cantidad := len(registry.Projects()); cantidad != 1 {
		t.Errorf("se esperaba el proyecto en memoria, hay %d proyectos", cantidad)
	}
}

// Escenario: Dos pestanas a la vez
// Cubre: RN-20
func TestDosPestanasALaVez(t *testing.T) {
	registry := domain.NewProjectRegistry()
	saver := &fakeSaver{}
	server := NewServer(registry, saver, testToday)
	const pestanas = 20

	// Veinte pestañas crean un proyecto cada una, todas al mismo tiempo. Con go test -race,
	// cualquier dato que dos pedidos toquen a la vez sin el candado hace fallar el test.
	var wg sync.WaitGroup
	codigos := make([]int, pestanas)
	for i := 0; i < pestanas; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			form := url.Values{"nombre": {fmt.Sprintf("Proyecto %d", i)}, "inicio": {"2026-11-01"}}
			req := httptest.NewRequest(http.MethodPost, "/proyectos/nuevo", strings.NewReader(form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			res := httptest.NewRecorder()
			server.ServeHTTP(res, req)
			codigos[i] = res.Code
		}(i)
	}
	wg.Wait()

	for i, codigo := range codigos {
		if codigo != http.StatusSeeOther {
			t.Errorf("la pestaña %d esperaba el estado %d, salio %d", i, http.StatusSeeOther, codigo)
		}
	}
	numeros := make(map[int]bool)
	for _, project := range registry.Projects() {
		numeros[project.ID()] = true
	}
	if len(numeros) != pestanas {
		t.Errorf("se esperaban %d proyectos con numeros distintos, quedaron %d", pestanas, len(numeros))
	}
	if saver.saves != pestanas {
		t.Errorf("se esperaba un guardado por operacion, %d, y hubo %d", pestanas, saver.saves)
	}
}

// Expresiones para revisar el HTML sin librerias: los campos de un formulario, un atributo
// de una etiqueta, y los errores que se muestran junto a un campo.
var (
	fieldTag   = regexp.MustCompile(`<(input|select|textarea)\b[^>]*>`)
	errorTag   = regexp.MustCompile(`<span class="error" id="([^"]+)">`)
	tableStart = regexp.MustCompile(`<div class="tabla">\s*<table`)
)

// attribute devuelve el valor de un atributo de una etiqueta, el primero si esta repetido,
// que es el que usa el navegador, y cuantas veces aparece.
func attribute(tag, name string) (string, int) {
	found := regexp.MustCompile(`\s`+name+`="([^"]*)"`).FindAllStringSubmatch(tag, -1)
	if len(found) == 0 {
		return "", 0
	}
	return found[0][1], len(found)
}

// Escenario: Usar la pagina en un celular y con el teclado
// Cubre: RN-18, RN-19. El contraste, el ancho de 360 y el foco se verifican a mano.
func TestUsarLaPaginaEnUnCelularYConElTeclado(t *testing.T) {
	registry := domain.NewProjectRegistry()
	project := addProject(t, registry, "Sistema de metricas")
	addMember(t, project, "Ana Perez", domain.RoleProductBuilder)
	addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
	// Lo que usan los formularios de la #78: un sprint, un registro de esfuerzo y una ronda.
	addSprint(t, registry, 1)
	if _, err := project.LogEffort(domain.EffortData{StoryID: 1, MemberName: "Ana Perez", Activity: "Escribir tests",
		Hours: 2, Day: testToday()}, testToday()); err != nil {
		t.Fatalf("no se pudo preparar el esfuerzo: %v", err)
	}
	if _, err := project.OpenVotingRound(1); err != nil {
		t.Fatalf("no se pudo preparar la ronda: %v", err)
	}
	addDefect(t, project, 1)
	// Una ronda revelada en otra historia, para el formulario del acuerdo.
	addStory(t, project, "Cargar defectos", domain.PriorityLow)
	round, err := project.OpenVotingRound(2)
	if err != nil {
		t.Fatalf("no se pudo abrir la ronda: %v", err)
	}
	if err := round.Vote("Ana Perez", 5); err != nil {
		t.Fatalf("no se pudo votar: %v", err)
	}
	if err := round.Reveal(); err != nil {
		t.Fatalf("no se pudo revelar: %v", err)
	}
	navegador := newBrowser(t, newTestServer(t, registry))
	defecto := url.Values{"titulo": {"Uno"}, "descripcion": {"Otro"}, "severidad": {"Mayor"}, "historia": {"1"}, "sprint": {"1"}}

	historia := url.Values{"titulo": {"Una"}, "descripcion": {"Otra"}, "prioridad": {"Alta"}, "criterios": {"Un criterio"}}
	esfuerzo := url.Values{"integrante": {"Ana Perez"}, "historia": {"1"}, "dia": {"2026-10-27"}, "horas": {"2"}, "actividad": {"Escribir tests"}}
	sin := func(datos url.Values, campo string) url.Values {
		copia := url.Values{}
		for k, v := range datos {
			copia[k] = v
		}
		copia.Set(campo, "")
		return copia
	}

	// Cada formulario vacio, y cada uno rechazado con el error en cada uno de sus campos.
	paginas := map[string]*httptest.ResponseRecorder{
		"proyecto nuevo":              navegador.get("/proyectos/nuevo"),
		"proyecto a modificar":        navegador.get("/proyectos/1/editar"),
		"integrante nuevo":            navegador.get("/proyectos/1/integrantes/nuevo"),
		"historia nueva":              navegador.get("/proyectos/1/historias/nueva"),
		"historia a modificar":        navegador.get("/proyectos/1/historias/1/editar"),
		"estimacion":                  navegador.get("/proyectos/1/historias/1/estimar"),
		"proyecto sin nombre":         navegador.post("/proyectos/nuevo", url.Values{"inicio": {"2026-11-01"}}),
		"proyecto sin fechas":         navegador.post("/proyectos/nuevo", url.Values{"nombre": {"Uno"}, "inicio": {"x"}, "fin": {"y"}}),
		"integrante sin nombre":       navegador.post("/proyectos/1/integrantes/nuevo", url.Values{"rol": {"Product Builder"}}),
		"integrante sin rol":          navegador.post("/proyectos/1/integrantes/nuevo", url.Values{"nombre": {"Luis Gomez"}}),
		"historia sin titulo":         navegador.post("/proyectos/1/historias/nueva", sin(historia, "titulo")),
		"historia sin descripcion":    navegador.post("/proyectos/1/historias/nueva", sin(historia, "descripcion")),
		"historia sin prioridad":      navegador.post("/proyectos/1/historias/nueva", sin(historia, "prioridad")),
		"historia sin criterios":      navegador.post("/proyectos/1/historias/nueva", sin(historia, "criterios")),
		"estimacion sin story points": navegador.post("/proyectos/1/historias/1/estimar", url.Values{}),
		// Los formularios de la #78 (tarea 24).
		"sprint nuevo":                 navegador.get("/proyectos/1/sprints/nuevo"),
		"sprint sin goal ni fechas":    navegador.post("/proyectos/1/sprints/nuevo", url.Values{"inicio": {"x"}, "fin": {"y"}}),
		"asignar al sprint":            navegador.get("/proyectos/1/sprints/1"),
		"asignar sin historia":         navegador.post("/proyectos/1/sprints/asignar", url.Values{}),
		"esfuerzo nuevo":               navegador.get("/proyectos/1/esfuerzo/nuevo"),
		"esfuerzo sin integrante":      navegador.post("/proyectos/1/esfuerzo/nuevo", sin(esfuerzo, "integrante")),
		"esfuerzo sin historia":        navegador.post("/proyectos/1/esfuerzo/nuevo", sin(esfuerzo, "historia")),
		"esfuerzo sin dia":             navegador.post("/proyectos/1/esfuerzo/nuevo", sin(esfuerzo, "dia")),
		"esfuerzo sin horas":           navegador.post("/proyectos/1/esfuerzo/nuevo", sin(esfuerzo, "horas")),
		"esfuerzo sin actividad":       navegador.post("/proyectos/1/esfuerzo/nuevo", sin(esfuerzo, "actividad")),
		"horas estimadas":              navegador.get("/proyectos/1/historias/1/horas-estimadas"),
		"horas estimadas que no son":   navegador.post("/proyectos/1/historias/1/horas-estimadas", url.Values{"horas": {"x"}}),
		"abrir una ronda":              navegador.get("/proyectos/1/planning-poker"),
		"abrir una ronda sin historia": navegador.post("/proyectos/1/planning-poker", url.Values{}),
		"votar":                        navegador.get("/proyectos/1/historias/1/rondas/1"),
		"votar sin elegir":             navegador.post("/proyectos/1/historias/1/rondas/1/votar", url.Values{}),
		"votar sin elegir quien vota":  navegador.post("/proyectos/1/historias/1/rondas/1/votar", url.Values{"puntos": {"5"}}),
		"defecto nuevo":                navegador.get("/proyectos/1/defectos/nuevo"),
		"defecto sin titulo":           navegador.post("/proyectos/1/defectos/nuevo", sin(defecto, "titulo")),
		"defecto sin descripcion":      navegador.post("/proyectos/1/defectos/nuevo", sin(defecto, "descripcion")),
		"defecto sin severidad":        navegador.post("/proyectos/1/defectos/nuevo", sin(defecto, "severidad")),
		"defecto sin historia":         navegador.post("/proyectos/1/defectos/nuevo", sin(defecto, "historia")),
		"defecto sin sprint":           navegador.post("/proyectos/1/defectos/nuevo", sin(defecto, "sprint")),
		"filtrar los defectos":         navegador.get("/proyectos/1/defectos"),
		"cambiar el estado y resolver": navegador.get("/proyectos/1/defectos/1"),
		"estado sin elegir":            navegador.post("/proyectos/1/defectos/1/estado", url.Values{}),
		"resolver sin sprint":          navegador.post("/proyectos/1/defectos/1/resolver", url.Values{}),
		"cerrar el sprint":             navegador.get("/proyectos/1/sprints/cerrar"),
		"cerrar sin fecha":             navegador.post("/proyectos/1/sprints/cerrar", url.Values{}),
		"acordar":                      navegador.get("/proyectos/1/historias/2/rondas/1"),
		"acordar sin elegir":           navegador.post("/proyectos/1/historias/2/rondas/1/acordar", url.Values{}),
	}

	for nombre, res := range paginas {
		t.Run(nombre, func(t *testing.T) {
			html := res.Body.String()

			// Todo campo tiene una etiqueta visible asociada.
			campos := 0
			for _, tag := range fieldTag.FindAllString(html, -1) {
				if tipo, _ := attribute(tag, "type"); tipo == "hidden" {
					continue
				}
				campos++
				id, _ := attribute(tag, "id")
				if id == "" || !strings.Contains(html, `<label for="`+id+`">`) {
					t.Errorf("el campo %s no tiene su etiqueta asociada", tag)
				}
			}
			if campos == 0 {
				t.Fatalf("se esperaba un formulario con campos, y la pagina es:\n%s", html)
			}

			// Con errores, el foco arranca en el primer campo con error, asi el lector de
			// pantalla lo anuncia con su mensaje; sin errores, la pagina no mueve el foco.
			// Hallazgo de la verificacion con el teclado.
			conFoco := 0
			for _, tag := range fieldTag.FindAllString(html, -1) {
				if strings.Contains(tag, " autofocus") {
					conFoco++
					if invalido, _ := attribute(tag, "aria-invalid"); invalido != "true" {
						t.Errorf("el foco arranca en %s, que no tiene error", tag)
					}
				}
			}
			tieneErrores := errorTag.MatchString(html)
			if tieneErrores && conFoco != 1 {
				t.Errorf("con errores se esperaba el foco en un solo campo, el primero con error, y hay %d", conFoco)
			}
			if !tieneErrores && conFoco != 0 {
				t.Errorf("sin errores la pagina no deberia mover el foco, y lo mueve en %d campos", conFoco)
			}

			// Todo error junto a un campo esta enlazado con el campo, para los lectores de
			// pantalla, y el campo esta marcado como invalido.
			for _, match := range errorTag.FindAllStringSubmatch(html, -1) {
				errorID := match[1]
				enlazado := false
				for _, tag := range fieldTag.FindAllString(html, -1) {
					describe, veces := attribute(tag, "aria-describedby")
					if veces > 1 {
						t.Errorf("el campo %s repite aria-describedby: el navegador solo usa el primero", tag)
					}
					if slices.Contains(strings.Fields(describe), errorID) {
						enlazado = true
						if invalido, _ := attribute(tag, "aria-invalid"); invalido != "true" {
							t.Errorf("el campo %s tiene error pero no esta marcado como invalido", tag)
						}
					}
				}
				if !enlazado {
					t.Errorf("el error %s no esta enlazado con ningun campo", errorID)
				}
			}
		})
	}

	// Toda tabla esta dentro de su recuadro, que se desplaza de costado si no entra.
	for _, pagina := range []string{"/proyectos", "/proyectos/1/integrantes", "/proyectos/1/backlog",
		"/proyectos/1/sprints", "/proyectos/1/esfuerzo", "/proyectos/1/planning-poker", "/proyectos/1/defectos", "/proyectos/1/historias/2/rondas"} {
		html := navegador.get(pagina).Body.String()
		tablas := strings.Count(html, "<table")
		if tablas == 0 || len(tableStart.FindAllString(html, -1)) != tablas {
			t.Errorf("en %s se esperaba cada tabla dentro de su recuadro, y la pagina es:\n%s", pagina, html)
		}
	}
}

// Escenario: Los datos siguen al cerrar y volver a abrir
// Cubre: RN-21, RN-22
func TestLosDatosSiguenAlCerrarYVolverAAbrir(t *testing.T) {
	path := filepath.Join(t.TempDir(), "datos-ics.json")
	// abrir arma la web desde el archivo, como hace el comando al arrancar.
	abrir := func() *browser {
		file := store.NewFile(path)
		registry, err := file.Load()
		if err != nil {
			t.Fatalf("no se pudieron leer los datos guardados: %v", err)
		}
		return newBrowser(t, NewServer(registry, file, testToday))
	}
	operar := func(navegador *browser, path string, datos url.Values) *httptest.ResponseRecorder {
		t.Helper()
		res := navegador.post(path, datos)
		if res.Code != http.StatusSeeOther {
			t.Fatalf("%s: se esperaba el estado %d, salio %d:\n%s", path, http.StatusSeeOther, res.Code, res.Body.String())
		}
		return res
	}

	// Dado que cargue el proyecto 1 con el integrante "Ana Perez" y la historia 1 estimada en 3.
	primera := abrir()
	operar(primera, "/proyectos/nuevo", url.Values{"nombre": {"Sistema de metricas"}, "inicio": {"2026-11-01"}})
	operar(primera, "/proyectos/1/integrantes/nuevo", url.Values{"nombre": {"Ana Perez"}, "rol": {"Product Builder"}})
	operar(primera, "/proyectos/1/historias/nueva", url.Values{
		"titulo": {"Registrar esfuerzo"}, "descripcion": {"Como integrante quiero registrar esfuerzo"},
		"prioridad": {"Alta"}, "criterios": {"Se puede registrar esfuerzo"},
	})
	operar(primera, "/proyectos/1/historias/1/estimar", url.Values{"puntos": {"3"}})

	// Cuando cierro el sistema y lo vuelvo a abrir.
	segunda := abrir()

	esperado := map[string]string{
		"/proyectos/1":             "<h1>Sistema de metricas</h1>",
		"/proyectos/1/integrantes": "<td>Ana Perez</td>",
		"/proyectos/1/backlog":     "<td>3</td>",
	}
	for pagina, texto := range esperado {
		if html := segunda.get(pagina).Body.String(); !strings.Contains(html, texto) {
			t.Errorf("al volver a abrir se esperaba en %s %s, y no esta en:\n%s", pagina, texto, html)
		}
	}

	res := operar(segunda, "/proyectos/nuevo", url.Values{"nombre": {"Metricas ICS"}, "inicio": {"2026-11-02"}})
	if destino := res.Header().Get("Location"); destino != "/proyectos/2" {
		t.Errorf("el siguiente proyecto deberia recibir el numero 2, y la web fue a %q", destino)
	}
}

// Escenario: Los datos nuevos siguen al cerrar y volver a abrir
// Cubre: RN-2 de la #78 (RN-21 a RN-25 de la #77)
//
// Por ahora con sprints, esfuerzo y rondas: los defectos se suman con la tarea 12, y que el
// sprint siguiente reciba el numero 2 necesita cerrar el sprint (#10).
func TestLosDatosNuevosSiguenAlCerrarYVolverAAbrir(t *testing.T) {
	path := filepath.Join(t.TempDir(), "datos-ics.json")
	abrir := func() *browser {
		file := store.NewFile(path)
		registry, err := file.Load()
		if err != nil {
			t.Fatalf("no se pudieron leer los datos guardados: %v", err)
		}
		return newBrowser(t, NewServer(registry, file, testToday))
	}
	operar := func(navegador *browser, path string, datos url.Values) {
		t.Helper()
		if res := navegador.post(path, datos); res.Code != http.StatusSeeOther {
			t.Fatalf("%s: se esperaba el estado %d, salio %d:\n%s", path, http.StatusSeeOther, res.Code, res.Body.String())
		}
	}

	// Dado que cargue el sprint 1, un registro de esfuerzo y una ronda revelada.
	primera := abrir()
	operar(primera, "/proyectos/nuevo", url.Values{"nombre": {"Sistema de metricas"}, "inicio": {"2026-10-01"}})
	operar(primera, "/proyectos/1/integrantes/nuevo", url.Values{"nombre": {"Ana Perez"}, "rol": {"Product Builder"}})
	operar(primera, "/proyectos/1/historias/nueva", url.Values{
		"titulo": {"Registrar esfuerzo"}, "descripcion": {"Como integrante quiero registrar esfuerzo"},
		"prioridad": {"Alta"}, "criterios": {"Se puede registrar esfuerzo"},
	})
	operar(primera, "/proyectos/1/historias/1/estimar", url.Values{"puntos": {"3"}})
	operar(primera, "/proyectos/1/sprints/nuevo", url.Values{"goal": {"Tener el MVP"}, "inicio": {"2026-10-26"}, "fin": {"2026-10-30"}})
	operar(primera, "/proyectos/1/sprints/asignar", url.Values{"historia": {"1"}})
	operar(primera, "/proyectos/1/esfuerzo/nuevo", url.Values{
		"integrante": {"Ana Perez"}, "historia": {"1"}, "dia": {"2026-10-27"}, "horas": {"2,5"}, "actividad": {"Escribir tests"},
	})
	operar(primera, "/proyectos/1/planning-poker", url.Values{"historia": {"1"}})
	operar(primera, "/proyectos/1/historias/1/rondas/1/votar", url.Values{"integrante": {"Ana Perez"}, "puntos": {"5"}})
	operar(primera, "/proyectos/1/historias/1/rondas/1/revelar", nil)
	operar(primera, "/proyectos/1/defectos/nuevo", url.Values{
		"titulo": {"Suma mal las horas"}, "descripcion": {"No suma el ultimo registro"},
		"severidad": {"Mayor"}, "historia": {"1"}, "sprint": {"1"},
	})
	operar(primera, "/proyectos/1/defectos/1/estado", url.Values{"estado": {"En curso"}})
	operar(primera, "/proyectos/1/historias/1/rondas/1/acordar", url.Values{"puntos": {"5"}})
	operar(primera, "/proyectos/1/sprints/cerrar", url.Values{"cierre": {"2026-10-27"}, "completadas": {"1"}})
	operar(primera, "/proyectos/1/sprints/nuevo", url.Values{"goal": {"Cerrar la interfaz"}, "inicio": {"2026-11-02"}, "fin": {"2026-11-13"}})

	// Cuando cierro el sistema y lo vuelvo a abrir.
	segunda := abrir()

	esperado := map[string]string{
		"/proyectos/1/sprints":              "Tener el MVP</a></td><td>26/10/2026</td><td>30/10/2026</td><td>Cerrado</td><td>27/10/2026</td>",
		"/proyectos/1/sprints/1":            "Registrar esfuerzo",
		"/proyectos/1/esfuerzo":             "Escribir tests",
		"/proyectos/1/historias/1/rondas/1": "<td>Ana Perez</td><td>5</td>",
		"/proyectos/1/defectos":             "<td>Mayor</td><td>En curso</td>",
		"/proyectos/1/sprints/2":            "Cerrar la interfaz",
		"/proyectos/1/historias/1/rondas":   "<td>Revelada</td><td>5</td>",
	}
	for pagina, texto := range esperado {
		if html := joinTags(segunda.get(pagina).Body.String()); !strings.Contains(html, texto) {
			t.Errorf("al volver a abrir se esperaba en %s %q, y no esta en:\n%s", pagina, texto, html)
		}
	}
}

// Sin escenario: hallazgo de la revisión de Maximiliano Eula (pedidos de otro sitio).
// Cubre: #77, D20 (sin usuarios ni contraseñas)
// Sin usuarios ni contraseñas, otra pagina abierta en el mismo navegador podria enviar un
// formulario a la web y cambiar los datos. Un pedido que cambia algo y viene de otro sitio
// se rechaza.
func TestUnPedidoDeOtroSitioSeRechaza(t *testing.T) {
	enviar := func(server http.Handler, encabezados map[string]string) *httptest.ResponseRecorder {
		form := url.Values{"nombre": {"Proyecto ajeno"}, "inicio": {"2026-11-01"}}
		req := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:8080/proyectos/nuevo", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		for nombre, valor := range encabezados {
			req.Header.Set(nombre, valor)
		}
		res := httptest.NewRecorder()
		server.ServeHTTP(res, req)
		return res
	}

	casos := map[string]map[string]string{
		"con el Origin de otro sitio":           {"Origin": "http://otro-sitio.example"},
		"con Sec-Fetch-Site de otro sitio":      {"Sec-Fetch-Site": "cross-site"},
		"con Sec-Fetch-Site de otro subdominio": {"Sec-Fetch-Site": "same-site"},
	}
	for nombre, encabezados := range casos {
		t.Run(nombre, func(t *testing.T) {
			registry := domain.NewProjectRegistry()
			saver := &fakeSaver{}

			res := enviar(NewServer(registry, saver, testToday), encabezados)

			if res.Code != http.StatusForbidden {
				t.Fatalf("se esperaba el estado %d, salio %d", http.StatusForbidden, res.Code)
			}
			if mensaje := template.HTMLEscapeString(errCrossOrigin.Error()); !strings.Contains(res.Body.String(), mensaje) {
				t.Errorf("se esperaba el aviso %q, y la respuesta es:\n%s", mensaje, res.Body.String())
			}
			if len(registry.Projects()) != 0 || saver.saves != 0 {
				t.Errorf("no se esperaba ningun cambio, y hay %d proyectos y %d guardados", len(registry.Projects()), saver.saves)
			}
		})
	}

	t.Run("el mismo pedido desde la propia web se acepta", func(t *testing.T) {
		registry := domain.NewProjectRegistry()

		res := enviar(NewServer(registry, &fakeSaver{}, testToday), map[string]string{
			"Origin":         "http://127.0.0.1:8080",
			"Sec-Fetch-Site": "same-origin",
		})

		if res.Code != http.StatusSeeOther || len(registry.Projects()) != 1 {
			t.Errorf("se esperaba crear el proyecto con el estado %d, salio %d y %d proyectos", http.StatusSeeOther, res.Code, len(registry.Projects()))
		}
	})
}
