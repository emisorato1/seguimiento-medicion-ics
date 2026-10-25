package web

import (
	"fmt"
	"net/http"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// sprintFields son los campos del formulario de sprint.
var sprintFields = []string{"goal", "inicio", "fin"}

// listSprints muestra los sprints del proyecto, en el orden del nucleo. Sin sprints muestra
// que no hay y ofrece crear el primero (RN-7 de la #77).
func (s *Server) listSprints(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		s.notFound(w, r, nil, err)
		return
	}
	s.render(w, r, http.StatusOK, "sprints.html", page{
		Title:   "Sprints de " + project.Name(),
		Project: project,
		Section: "sprints",
		Data:    project,
	})
}

// renderSprintForm muestra el formulario para crear un sprint, con lo escrito y sus errores.
func (s *Server) renderSprintForm(w http.ResponseWriter, r *http.Request, status int, project *domain.Project, f *form) {
	s.render(w, r, status, "sprint-form.html", page{
		Title:   "Nuevo sprint",
		Project: project,
		Section: "sprints",
		Data:    f,
	})
}

// newSprint muestra el formulario para crear un sprint.
func (s *Server) newSprint(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		s.notFound(w, r, nil, err)
		return
	}
	s.renderSprintForm(w, r, http.StatusOK, project, newForm())
}

// createSprint crea el sprint con los datos del formulario y vuelve a la lista de sprints con
// 303, asi recargar no lo crea de nuevo. Si el dominio lo rechaza, vuelve a mostrar el
// formulario con lo escrito y el mensaje (RN-2).
func (s *Server) createSprint(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		s.notFound(w, r, nil, err)
		return
	}
	f := readForm(r, sprintFields...)
	start := f.date("inicio")
	end := f.date("fin")
	if f.hasErrors() {
		s.renderSprintForm(w, r, http.StatusUnprocessableEntity, project, f)
		return
	}

	sprint, err := project.AddSprint(domain.SprintData{Goal: f.Values["goal"], Start: start, End: end})
	if err != nil {
		f.reject(err)
		s.renderSprintForm(w, r, http.StatusUnprocessableEntity, project, f)
		return
	}

	s.done(w, r, fmt.Sprintf("Se creó el sprint %d", sprint.Number()), fmt.Sprintf("/proyectos/%d/sprints", project.ID()))
}
