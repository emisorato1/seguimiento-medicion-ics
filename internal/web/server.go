// Package web es la interfaz web del sistema. Es una capa delgada, igual que la consola:
// lee el formulario, llama al nucleo y muestra lo que responde. No valida reglas de negocio
// ni guarda datos propios.
package web

import (
	"bytes"
	"embed"
	"encoding/base64"
	"errors"
	"html/template"
	"io/fs"
	"net/http"
	"path"
	"sync"
	"time"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// files tiene las plantillas y el CSS adentro del binario (D2): la web funciona desde
// cualquier carpeta y sin internet.
//
//go:embed templates static
var files embed.FS

// errCrossOrigin es el aviso de que un pedido que cambia algo vino de otro sitio.
var errCrossOrigin = errors.New("el pedido vino de otra página y no se aceptó: los cambios se hacen desde esta web")

// errPageNotFound es para una direccion que no es de ninguna pagina. Lo que no existe en
// el dominio, como un proyecto o una historia, se informa con el error del dominio.
var errPageNotFound = errors.New("no se encontró la página")

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
	// handler es el mux detras de la proteccion contra pedidos de otro sitio.
	handler http.Handler
	pages   map[string]*template.Template
	// lock hace que los pedidos usen el registro de a uno (RN-20, D12): net/http atiende
	// cada pedido en paralelo y el dominio no esta hecho para eso.
	lock sync.RWMutex
}

// NewServer arma la web sobre el registro de proyectos que recibe, guardandolo con saver.
func NewServer(registry *domain.ProjectRegistry, saver Saver) *Server {
	s := &Server{
		registry: registry,
		saver:    saver,
		mux:      http.NewServeMux(),
		pages:    parsePages(),
	}

	s.mux.HandleFunc("/", s.pageNotFound)
	s.mux.HandleFunc("GET /{$}", s.home)
	s.mux.Handle("GET /static/", http.FileServerFS(files))
	s.mux.HandleFunc("GET /proyectos", s.listProjects)
	s.mux.HandleFunc("GET /proyectos/nuevo", s.newProject)
	s.mux.HandleFunc("POST /proyectos/nuevo", s.createProject)
	s.mux.HandleFunc("GET /proyectos/{id}", s.showProject)
	s.mux.HandleFunc("GET /proyectos/{id}/editar", s.editProject)
	s.mux.HandleFunc("POST /proyectos/{id}/editar", s.updateProject)
	s.mux.HandleFunc("GET /proyectos/{id}/integrantes", s.listMembers)
	s.mux.HandleFunc("GET /proyectos/{id}/integrantes/nuevo", s.newMember)
	s.mux.HandleFunc("POST /proyectos/{id}/integrantes/nuevo", s.createMember)
	s.mux.HandleFunc("GET /proyectos/{id}/integrantes/quitar", s.confirmRemoveMember)
	s.mux.HandleFunc("POST /proyectos/{id}/integrantes/quitar", s.removeMember)
	s.mux.HandleFunc("GET /proyectos/{id}/backlog", s.listBacklog)
	s.mux.HandleFunc("GET /proyectos/{id}/historias/nueva", s.newStory)
	s.mux.HandleFunc("POST /proyectos/{id}/historias/nueva", s.createStory)
	s.mux.HandleFunc("GET /proyectos/{id}/historias/{n}", s.showStory)
	s.mux.HandleFunc("GET /proyectos/{id}/historias/{n}/editar", s.editStory)
	s.mux.HandleFunc("POST /proyectos/{id}/historias/{n}/editar", s.updateStory)
	s.mux.HandleFunc("GET /proyectos/{id}/historias/{n}/estimar", s.estimateForm)
	s.mux.HandleFunc("POST /proyectos/{id}/historias/{n}/estimar", s.estimateStory)
	s.mux.HandleFunc("GET /proyectos/{id}/historias/{n}/quitar-estimacion", s.confirmRemoveEstimate)
	s.mux.HandleFunc("POST /proyectos/{id}/historias/{n}/quitar-estimacion", s.removeEstimate)
	s.mux.HandleFunc("GET /proyectos/{id}/sprints", s.listSprints)

	// Sin usuarios ni contraseñas (D20), otra pagina abierta en el mismo navegador podria
	// enviar un formulario a la web. La proteccion de la biblioteca estandar rechaza todo
	// pedido que cambia algo y viene de otro sitio (revision del PR #95).
	protection := http.NewCrossOriginProtection()
	protection.SetDenyHandler(http.HandlerFunc(s.crossOrigin))
	s.handler = protection.Handler(s.mux)
	return s
}

// crossOrigin responde 403 a un pedido de otro sitio, con el aviso en castellano.
func (s *Server) crossOrigin(w http.ResponseWriter, r *http.Request) {
	s.render(w, r, http.StatusForbidden, "rechazado.html", page{Title: "Pedido rechazado", Data: errCrossOrigin.Error()})
}

// ServeHTTP atiende cada pedido del navegador. Toma el candado alrededor de todo el pedido,
// asi una operacion y su guardado quedan juntos (D12). Los pedidos que solo leen —las paginas
// que se piden con GET— pueden ir a la vez entre ellos; los que cambian algo van solos.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet || r.Method == http.MethodHead {
		s.lock.RLock()
		defer s.lock.RUnlock()
	} else {
		s.lock.Lock()
		defer s.lock.Unlock()
	}
	s.handler.ServeHTTP(w, r)
}

// home lleva a la lista de proyectos, que es la pagina de inicio.
func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/proyectos", http.StatusSeeOther)
}

// templateFuncs son las funciones que pueden usar las plantillas.
var templateFuncs = template.FuncMap{
	"fecha":   formatDate,
	"choices": newChoices,
}

// choices es una lista para elegir: el texto de la opcion vacia que pide elegir, las
// opciones y la que esta elegida.
type choices struct {
	Prompt   string
	Options  []string
	Selected string
}

// newChoices arma una lista para elegir desde una plantilla.
func newChoices(prompt string, options []string, selected string) choices {
	return choices{Prompt: prompt, Options: options, Selected: selected}
}

// formatDate muestra una fecha como dd/mm/aaaa (RN-5).
func formatDate(day time.Time) string {
	return day.Format("02/01/2006")
}

// confirmation es la pagina que pide confirmar antes de quitar algo (D9): la pregunta, a
// donde se envia lo que se confirma, con que datos, a donde vuelve si se cancela y, si el
// dominio rechazo lo confirmado, su motivo.
type confirmation struct {
	Question string
	Detail   string
	Action   string
	Fields   map[string]string
	Confirm  string
	Cancel   string
	Error    string
}

// notSaved es la pagina que avisa que una operacion se hizo pero no quedo guardada.
type notSaved struct {
	Operation string
	Reason    string
	Continue  string
}

// done termina una operacion aceptada: guarda el registro (D13), deja el mensaje que la
// confirma (RN-8) y redirige con 303 a la pagina del resultado (D5). Si el guardado falla,
// el cambio queda en memoria y lo incluye el siguiente guardado que funcione; se avisa con
// estado 500 (RN-25, D19).
func (s *Server) done(w http.ResponseWriter, r *http.Request, message, location string) {
	if err := s.saver.Save(s.registry); err != nil {
		s.render(w, r, http.StatusInternalServerError, "no-guardado.html", page{
			Title: "No quedó guardado",
			Data:  notSaved{Operation: message, Reason: err.Error(), Continue: location},
		})
		return
	}
	setFlash(w, message)
	http.Redirect(w, r, location, http.StatusSeeOther)
}

// notFound responde 404 con el mensaje de lo que no se encontro y el acceso para volver a
// la lista de proyectos (RN-16). Si se estaba dentro de un proyecto, la barra lo sigue
// mostrando.
func (s *Server) notFound(w http.ResponseWriter, r *http.Request, project *domain.Project, err error) {
	s.render(w, r, http.StatusNotFound, "no-encontrado.html", page{
		Title:   "No se encontró",
		Project: project,
		Data:    err.Error(),
	})
}

// pageNotFound atiende toda direccion que no es de ninguna pagina.
func (s *Server) pageNotFound(w http.ResponseWriter, r *http.Request) {
	s.notFound(w, r, nil, errPageNotFound)
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
