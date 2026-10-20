package web

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// fakeSaver es el guardado falso de los tests (D21): en vez de escribir un archivo, cuenta
// cuantas veces se guardo el registro.
type fakeSaver struct {
	saves int
}

func (f *fakeSaver) Save(registry *domain.ProjectRegistry) error {
	f.saves++
	return nil
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
	addProject(t, registry, "Sistema de metricas")
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
		figura     string
	}{
		{
			nombre: "agregar un integrante",
			preparar: func(t *testing.T, registry *domain.ProjectRegistry) {
				addProject(t, registry, "Sistema de metricas")
			},
			formulario: "/proyectos/1/integrantes/nuevo",
			datos:      url.Values{"nombre": {"Ana Perez"}, "rol": {"Product Builder"}},
			mensaje:    "Se agregó a Ana Perez",
			figura:     "<td>Product Builder</td>",
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
			figura:     "<h1>Metricas ICS</h1>",
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
			figura:     "sin descripción",
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
			if !strings.Contains(html, caso.figura) {
				t.Errorf("se esperaba que figure %s, y no esta en:\n%s", caso.figura, html)
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
	}

	for _, caso := range casos {
		t.Run(caso.formulario, func(t *testing.T) {
			registry := domain.NewProjectRegistry()
			addProject(t, registry, "Sistema de metricas")
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
