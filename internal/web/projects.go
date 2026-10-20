package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// listProjects muestra la lista de proyectos, que es la pagina de inicio. Sin proyectos
// muestra que no hay y ofrece crear el primero (RN-7, RN-23).
func (s *Server) listProjects(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Projects []*domain.Project
	}{
		Projects: s.registry.Projects(),
	}
	s.render(w, r, http.StatusOK, "proyectos.html", page{Title: "Proyectos", Data: data})
}

// projectFields son los campos del formulario de proyecto.
var projectFields = []string{"nombre", "descripcion", "inicio", "fin"}

// newProject muestra el formulario para crear un proyecto.
func (s *Server) newProject(w http.ResponseWriter, r *http.Request) {
	s.renderProjectForm(w, r, http.StatusOK, newForm())
}

// renderProjectForm muestra el formulario de proyecto con lo escrito y sus errores.
func (s *Server) renderProjectForm(w http.ResponseWriter, r *http.Request, status int, f *form) {
	data := struct{ Form *form }{Form: f}
	s.render(w, r, status, "proyecto-nuevo.html", page{Title: "Nuevo proyecto", Data: data})
}

// createProject crea el proyecto con los datos del formulario y redirige con 303 a su
// pagina, asi recargar no lo crea de nuevo (RN-9, D5). Si el dominio lo rechaza, vuelve a
// mostrar el formulario con lo escrito y el mensaje junto a su campo (RN-12, D6).
func (s *Server) createProject(w http.ResponseWriter, r *http.Request) {
	f := readForm(r, projectFields...)
	start := f.date("inicio")
	end := f.date("fin")
	if f.hasErrors() {
		s.renderProjectForm(w, r, http.StatusUnprocessableEntity, f)
		return
	}

	project, err := s.registry.Create(domain.ProjectData{
		Name:        f.Values["nombre"],
		Description: f.Values["descripcion"],
		Start:       start,
		End:         end,
	})
	if err != nil {
		f.reject(err)
		s.renderProjectForm(w, r, http.StatusUnprocessableEntity, f)
		return
	}

	s.saver.Save(s.registry)
	setFlash(w, "Se creó el proyecto "+project.Name())
	http.Redirect(w, r, fmt.Sprintf("/proyectos/%d", project.ID()), http.StatusSeeOther)
}

// showProject muestra la pagina de un proyecto.
func (s *Server) showProject(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	project, err := s.registry.Get(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	s.render(w, r, http.StatusOK, "proyecto.html", page{
		Title:   project.Name(),
		Project: project,
		Section: "datos",
		Data:    project,
	})
}
