package web

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

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

// post envia un formulario a la web, como lo haria el navegador.
func post(t *testing.T, server http.Handler, path string, form url.Values) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res := httptest.NewRecorder()
	server.ServeHTTP(res, req)
	return res
}

// Escenario: Abrir el sistema con un solo comando y sin internet
// Cubre: RN-3, RN-4. Esta es la parte de la web; el comando lo prueba cmd/web.
func TestAbrirElSistemaConUnSoloComandoYSinInternet(t *testing.T) {
	server := newTestServer(t, domain.NewProjectRegistry())

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
	paginas := []string{"/proyectos", "/proyectos/nuevo"}

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
