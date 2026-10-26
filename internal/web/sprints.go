package web

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// errInvalidStory es el error de forma de la historia elegida: no se eligio ninguna de la
// lista, o llego algo que no es un numero (RN-2). Que se pueda asignar lo decide el dominio.
var errInvalidStory = errors.New("hay que elegir la historia de la lista")

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

// storyOptions son las historias del backlog para elegir, todas, sin filtrar por reglas
// (RN-16, D10): se elige por el numero y se lee con el titulo.
func storyOptions(project *domain.Project) []option {
	var options []option
	for _, story := range project.Backlog() {
		options = append(options, option{Value: strconv.Itoa(story.ID()), Label: fmt.Sprintf("%d — %s", story.ID(), story.Title())})
	}
	return options
}

// sprintPage es lo que muestra la pagina de un sprint: el sprint, sus historias, si es el
// abierto y, para el abierto, el formulario para asignar.
type sprintPage struct {
	ProjectID int
	Sprint    *domain.Sprint
	Open      bool
	Stories   []*domain.Story
	Choices   []option
	Form      *form
}

// renderSprint muestra la pagina de un sprint con su Sprint Backlog.
func (s *Server) renderSprint(w http.ResponseWriter, r *http.Request, status int, project *domain.Project, sprint *domain.Sprint, f *form) {
	data := sprintPage{
		ProjectID: project.ID(),
		Sprint:    sprint,
		Open:      sprint.State() == domain.SprintOpen,
		Choices:   storyOptions(project),
		Form:      f,
	}
	for _, id := range sprint.Stories() {
		if story, err := project.Story(id); err == nil {
			data.Stories = append(data.Stories, story)
		}
	}
	s.render(w, r, status, "sprint.html", page{
		Title:   fmt.Sprintf("Sprint %d", sprint.Number()),
		Project: project,
		Section: "sprints",
		Data:    data,
	})
}

// showSprint muestra un sprint, abierto o cerrado.
func (s *Server) showSprint(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		s.notFound(w, r, nil, err)
		return
	}
	n, _ := strconv.Atoi(r.PathValue("n"))
	sprint, err := project.Sprint(n)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	s.renderSprint(w, r, http.StatusOK, project, sprint, newForm())
}

// openSprint busca el proyecto y su sprint abierto. Sin sprint abierto devuelve el error del
// dominio para asignar o quitar.
func (s *Server) openSprint(r *http.Request) (*domain.Project, *domain.Sprint, error) {
	project, err := s.findProject(r)
	if err != nil {
		return nil, nil, err
	}
	sprint, ok := project.OpenSprint()
	if !ok {
		return project, nil, domain.ErrNoOpenSprint
	}
	return project, sprint, nil
}

// assignStory asigna al sprint abierto la historia elegida y vuelve al sprint con 303. Si el
// dominio la rechaza, vuelve a mostrar el sprint con el mensaje (RN-2, RN-16).
func (s *Server) assignStory(w http.ResponseWriter, r *http.Request) {
	project, sprint, err := s.openSprint(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	f := readForm(r, "historia")
	id := f.number("historia", errInvalidStory)
	if f.hasErrors() {
		s.renderSprint(w, r, http.StatusUnprocessableEntity, project, sprint, f)
		return
	}
	if err := project.AssignStory(id); err != nil {
		f.reject(err)
		s.renderSprint(w, r, http.StatusUnprocessableEntity, project, sprint, f)
		return
	}
	s.done(w, r, fmt.Sprintf("Se asignó la historia %d al sprint %d", id, sprint.Number()), fmt.Sprintf("/proyectos/%d/sprints/%d", project.ID(), sprint.Number()))
}

// unassignConfirmation es la pagina que pide confirmar antes de quitar una historia del
// sprint abierto (RN-10 de la #77).
func unassignConfirmation(project *domain.Project, sprint *domain.Sprint, story *domain.Story) confirmation {
	return confirmation{
		Question: fmt.Sprintf("¿Quitar la historia %d del sprint %d?", story.ID(), sprint.Number()),
		Detail:   fmt.Sprintf("«%s» vuelve al Product Backlog como Pendiente. Se puede volver a asignar después.", story.Title()),
		Action:   fmt.Sprintf("/proyectos/%d/sprints/quitar", project.ID()),
		Fields:   map[string]string{"historia": strconv.Itoa(story.ID())},
		Confirm:  "Sí, quitar",
		Cancel:   fmt.Sprintf("/proyectos/%d/sprints/%d", project.ID(), sprint.Number()),
	}
}

// findSprintStory busca el proyecto, su sprint abierto y la historia elegida.
func (s *Server) findSprintStory(r *http.Request) (*domain.Project, *domain.Sprint, *domain.Story, error) {
	project, sprint, err := s.openSprint(r)
	if err != nil {
		return project, nil, nil, err
	}
	id, _ := strconv.Atoi(r.FormValue("historia"))
	story, err := project.Story(id)
	if err != nil {
		return project, sprint, nil, err
	}
	return project, sprint, story, nil
}

// confirmUnassign pide confirmar antes de quitar una historia del sprint abierto.
func (s *Server) confirmUnassign(w http.ResponseWriter, r *http.Request) {
	project, sprint, story, err := s.findSprintStory(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	s.render(w, r, http.StatusOK, "confirmar.html", page{
		Title:   "Quitar historia del sprint",
		Project: project,
		Section: "sprints",
		Data:    unassignConfirmation(project, sprint, story),
	})
}

// unassignStory quita la historia del sprint abierto, despues de que se confirmo. Si el
// dominio lo rechaza, vuelve a pedir confirmar con su mensaje.
func (s *Server) unassignStory(w http.ResponseWriter, r *http.Request) {
	project, sprint, story, err := s.findSprintStory(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	if err := project.UnassignStory(story.ID()); err != nil {
		data := unassignConfirmation(project, sprint, story)
		data.Error = err.Error()
		s.render(w, r, http.StatusUnprocessableEntity, "confirmar.html", page{
			Title:   "Quitar historia del sprint",
			Project: project,
			Section: "sprints",
			Data:    data,
		})
		return
	}
	s.done(w, r, fmt.Sprintf("Se quitó la historia %d del sprint %d", story.ID(), sprint.Number()), fmt.Sprintf("/proyectos/%d/sprints/%d", project.ID(), sprint.Number()))
}
