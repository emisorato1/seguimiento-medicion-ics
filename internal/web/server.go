// Package web es la interfaz web del sistema. Es una capa delgada, igual que la consola:
// lee el formulario, llama al nucleo y muestra lo que responde. No valida reglas de negocio
// ni guarda datos propios.
package web

import (
	"bytes"
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"path"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// files tiene las plantillas y el CSS adentro del binario (D2): la web funciona desde
// cualquier carpeta y sin internet.
//
//go:embed templates static
var files embed.FS

// Server es la web: las rutas, las plantillas y el registro de proyectos que muestra.
type Server struct {
	registry *domain.ProjectRegistry
	mux      *http.ServeMux
	pages    map[string]*template.Template
}

// NewServer arma la web sobre el registro de proyectos que recibe.
func NewServer(registry *domain.ProjectRegistry) *Server {
	s := &Server{
		registry: registry,
		mux:      http.NewServeMux(),
		pages:    parsePages(),
	}

	s.mux.HandleFunc("GET /{$}", s.home)
	s.mux.Handle("GET /static/", http.FileServerFS(files))
	s.mux.HandleFunc("GET /proyectos", s.listProjects)
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

// parsePages arma una plantilla por pagina: cada una es el layout comun mas el archivo de
// la pagina, que define el bloque "content". Una pantalla nueva solo suma su archivo.
func parsePages() map[string]*template.Template {
	layout := template.Must(template.ParseFS(files, "templates/layout.html"))

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

// page es lo que recibe el layout: el titulo de la pestaña y los datos de la pagina.
type page struct {
	Title string
	Data  any
}

// render muestra una pagina dentro del layout. Arma todo antes de escribir, asi una
// plantilla que falla no deja una pagina cortada a la mitad.
func (s *Server) render(w http.ResponseWriter, status int, name string, p page) {
	var buf bytes.Buffer
	if err := s.pages[name].ExecuteTemplate(&buf, "layout", p); err != nil {
		http.Error(w, "No se pudo mostrar la página.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	buf.WriteTo(w)
}
