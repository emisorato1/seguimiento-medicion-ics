package web

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
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

// newTestServer arma una web nueva sobre el registro que recibe, con un guardado falso.
// Cada test arma el suyo y arranca de cero.
func newTestServer(t *testing.T, registry *domain.ProjectRegistry) *Server {
	t.Helper()
	return NewServer(registry, &fakeSaver{})
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
		formulario string
		datos      url.Values
		campo      string
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
	}

	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			registry := domain.NewProjectRegistry()
			caso.preparar(t, registry)
			saver := &fakeSaver{}
			server := NewServer(registry, saver)
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

			// El mensaje del dominio, junto a su campo.
			junto := fmt.Sprintf(`id="%s-error">%s<`, caso.campo, template.HTMLEscapeString(caso.err.Error()))
			if !strings.Contains(html, junto) {
				t.Errorf("se esperaba junto al campo %s el mensaje %q, y no esta en:\n%s", caso.campo, caso.err, html)
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
	}{
		{pedido: "/proyectos/99", err: domain.ErrProjectNotFound},
		{pedido: "/proyectos/abc", err: domain.ErrProjectNotFound},
		{pedido: "/proyectos/1/historias/99", err: domain.ErrStoryNotFound},
		// No esta en los Ejemplos: una direccion que no es de ninguna pagina tambien se
		// informa en castellano, con el mismo acceso a la lista.
		{pedido: "/proyectos/1/cualquiera", err: errPageNotFound},
	}

	for _, caso := range casos {
		t.Run(caso.pedido, func(t *testing.T) {
			registry := domain.NewProjectRegistry()
			addProject(t, registry, "Sistema de metricas")
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
			navegador := newBrowser(t, NewServer(registry, saver))
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
	navegador := newBrowser(t, NewServer(registry, saver))

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
	server := NewServer(registry, saver)
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
	navegador := newBrowser(t, newTestServer(t, registry))

	historia := url.Values{"titulo": {"Una"}, "descripcion": {"Otra"}, "prioridad": {"Alta"}, "criterios": {"Un criterio"}}
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
	for _, pagina := range []string{"/proyectos", "/proyectos/1/integrantes", "/proyectos/1/backlog"} {
		html := navegador.get(pagina).Body.String()
		tablas := strings.Count(html, "<table")
		if tablas == 0 || len(tableStart.FindAllString(html, -1)) != tablas {
			t.Errorf("en %s se esperaba cada tabla dentro de su recuadro, y la pagina es:\n%s", pagina, html)
		}
	}
}
