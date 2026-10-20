// Package web es la interfaz web del sistema. Es una capa delgada, igual que la consola:
// lee el formulario, llama al nucleo y muestra lo que responde. No valida reglas de negocio
// ni guarda datos propios.
package web

import (
	"bytes"
	"embed"
	"encoding/base64"
	"html/template"
	"io/fs"
	"net/http"
	"path"
	"time"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// files tiene las plantillas y el CSS adentro del binario (D2): la web funciona desde
// cualquier carpeta y sin internet.
//
//go:embed templates static
var files embed.FS

// Saver guarda el registro completo despues de cada operacion aceptada (D13). En el
// programa es el archivo de internal/store; en los tests, uno falso en memoria (D21).
type Saver interface {
	Save(registry *domain.ProjectRegistry) error
}

// Server es la web: las rutas, las plantillas, el registro de proyectos que muestra y
// donde lo guarda.
type Server struct {
	registry *domain.ProjectRegistry
	saver    Saver
	mux      *http.ServeMux
	pages    map[string]*template.Template
}

// NewServer arma la web sobre el registro de proyectos que recibe, guardandolo con saver.
func NewServer(registry *domain.ProjectRegistry, saver Saver) *Server {
	s := &Server{
		registry: registry,
		saver:    saver,
		mux:      http.NewServeMux(),
		pages:    parsePages(),
	}

	s.mux.HandleFunc("GET /{$}", s.home)
	s.mux.Handle("GET /static/", http.FileServerFS(files))
	s.mux.HandleFunc("GET /proyectos", s.listProjects)
	s.mux.HandleFunc("GET /proyectos/nuevo", s.newProject)
	s.mux.HandleFunc("POST /proyectos/nuevo", s.createProject)
	s.mux.HandleFunc("GET /proyectos/{id}", s.showProject)
	return s
}

// ServeHTTP atiende cada pedido del navegador.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

// home lleva a la lista de proyectos, que es la pagina de inicio.
func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/proyectos", http.StatusSeeOther)
}

// templateFuncs son las funciones que pueden usar las plantillas.
var templateFuncs = template.FuncMap{
	"fecha": formatDate,
}

// formatDate muestra una fecha como dd/mm/aaaa (RN-5).
func formatDate(day time.Time) string {
	return day.Format("02/01/2006")
}

// parsePages arma una plantilla por pagina: cada una es el layout comun mas el archivo de
// la pagina, que define el bloque "content". Una pantalla nueva solo suma su archivo.
func parsePages() map[string]*template.Template {
	layout := template.Must(template.New("layout").Funcs(templateFuncs).ParseFS(files, "templates/layout.html"))

	names, err := fs.Glob(files, "templates/*.html")
	if err != nil {
		panic(err)
	}

	pages := make(map[string]*template.Template)
	for _, name := range names {
		base := path.Base(name)
		if base == "layout.html" {
			continue
		}
		page := template.Must(layout.Clone())
		pages[base] = template.Must(page.ParseFS(files, name))
	}
	return pages
}

// page es lo que recibe el layout: el titulo de la pestaña, el mensaje de confirmacion si
// hay uno, el proyecto en el que se esta (para la barra, RN-4) con la seccion abierta, y
// los datos de la pagina.
type page struct {
	Title   string
	Flash   string
	Project *domain.Project
	Section string
	Data    any
}

// flashCookie es la cookie que lleva el mensaje de confirmacion hasta la pagina siguiente.
const flashCookie = "aviso"

// setFlash deja el mensaje que confirma una operacion para la pagina a la que se redirige
// (RN-8). Viaja en una cookie y no en la direccion, asi recargar o guardar el enlace no lo
// vuelve a mostrar (D8). Va en base64 porque una cookie no admite tildes ni espacios.
func setFlash(w http.ResponseWriter, message string) {
	http.SetCookie(w, &http.Cookie{
		Name:     flashCookie,
		Value:    base64.URLEncoding.EncodeToString([]byte(message)),
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
}

// takeFlash lee el mensaje de confirmacion y borra la cookie, asi se muestra una sola vez
// (RN-9). Sin cookie, o con una que no se puede leer, no hay mensaje.
func takeFlash(w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie(flashCookie)
	if err != nil {
		return ""
	}
	http.SetCookie(w, &http.Cookie{Name: flashCookie, Path: "/", MaxAge: -1})
	message, err := base64.URLEncoding.DecodeString(cookie.Value)
	if err != nil {
		return ""
	}
	return string(message)
}

// render muestra una pagina dentro del layout, con el mensaje de confirmacion pendiente si
// lo hay. Arma todo antes de escribir, asi una plantilla que falla no deja una pagina
// cortada a la mitad.
func (s *Server) render(w http.ResponseWriter, r *http.Request, status int, name string, p page) {
	p.Flash = takeFlash(w, r)

	var buf bytes.Buffer
	if err := s.pages[name].ExecuteTemplate(&buf, "layout", p); err != nil {
		http.Error(w, "No se pudo mostrar la página.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	buf.WriteTo(w)
}
