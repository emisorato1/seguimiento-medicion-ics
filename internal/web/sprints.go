package web

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

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
// abierto y, para el abierto, el formulario para asignar. De uno cerrado, cuales historias se
// completaron y los story points que tenian al cerrarlo, que guarda la #10.
type sprintPage struct {
	ProjectID int
	Sprint    *domain.Sprint
	Open      bool
	Stories   []*domain.Story
	Completed map[int]bool
	Points    map[int]int
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
	data.Stories = sprintStories(project, sprint)
	data.Completed = map[int]bool{}
	for _, id := range sprint.CompletedStories() {
		data.Completed[id] = true
	}
	data.Points = sprint.PointsAtClose()
	s.render(w, r, status, "sprint.html", page{
		Title:   fmt.Sprintf("Sprint %d", sprint.Number()),
		Project: project,
		Section: "sprints",
		Data:    data,
	})
}

// sprintStories son las historias asignadas al sprint, en el orden en que se asignaron.
func sprintStories(project *domain.Project, sprint *domain.Sprint) []*domain.Story {
	var stories []*domain.Story
	for _, id := range sprint.Stories() {
		if story, err := project.Story(id); err == nil {
			stories = append(stories, story)
		}
	}
	return stories
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

// closeForm es lo que muestra el formulario para cerrar el sprint abierto: la fecha de cierre
// y una casilla por cada historia asignada, para marcar las completadas (RN-10, D8).
type closeForm struct {
	ProjectID int
	Sprint    *domain.Sprint
	Stories   []*domain.Story
	Checked   map[int]bool
	Form      *form
}

// renderCloseForm muestra el formulario para cerrar el sprint, con lo elegido y sus errores.
func (s *Server) renderCloseForm(w http.ResponseWriter, r *http.Request, status int, project *domain.Project, sprint *domain.Sprint, checked []int, f *form) {
	data := closeForm{
		ProjectID: project.ID(),
		Sprint:    sprint,
		Stories:   sprintStories(project, sprint),
		Checked:   map[int]bool{},
		Form:      f,
	}
	for _, id := range checked {
		data.Checked[id] = true
	}
	s.render(w, r, status, "sprint-cerrar.html", page{
		Title:   fmt.Sprintf("Cerrar el sprint %d", sprint.Number()),
		Project: project,
		Section: "sprints",
		Data:    data,
	})
}

// closeSprintForm muestra el formulario para cerrar el sprint abierto, con la fecha de hoy
// propuesta y ninguna historia marcada (RN-10, RN-11).
func (s *Server) closeSprintForm(w http.ResponseWriter, r *http.Request) {
	project, sprint, err := s.openSprint(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	f := newForm()
	f.Values["cierre"] = s.today().Format(dateLayout)
	s.renderCloseForm(w, r, http.StatusOK, project, sprint, nil, f)
}

// readCompleted lee los numeros de las historias marcadas como completadas. Algo que no es un
// numero no viene de las casillas de la pagina: se marca arriba del formulario (RN-2).
func readCompleted(f *form, values []string) []int {
	var ids []int
	for _, value := range values {
		id, err := strconv.Atoi(value)
		if err != nil {
			f.General = errInvalidStory.Error()
			return nil
		}
		ids = append(ids, id)
	}
	return ids
}

// closeConfirmation es la pagina que pide confirmar antes de cerrar el sprint, que no se
// deshace: dice cuales quedan completadas y cuales vuelven al Product Backlog (RN-9, RN-10).
func closeConfirmation(project *domain.Project, sprint *domain.Sprint, closedOn string, completed []int) confirmation {
	done := map[int]bool{}
	for _, id := range completed {
		done[id] = true
	}
	var completedNames, pendingNames []string
	for _, story := range sprintStories(project, sprint) {
		name := fmt.Sprintf("%d — %s", story.ID(), story.Title())
		if done[story.ID()] {
			completedNames = append(completedNames, name)
		} else {
			pendingNames = append(pendingNames, name)
		}
	}
	question := fmt.Sprintf("¿Cerrar el sprint %d?", sprint.Number())
	if day, err := time.Parse(dateLayout, closedOn); err == nil {
		question = fmt.Sprintf("¿Cerrar el sprint %d el %s?", sprint.Number(), formatDate(day))
	}
	var ids []string
	for _, id := range completed {
		ids = append(ids, strconv.Itoa(id))
	}
	return confirmation{
		Question: question,
		Detail:   fmt.Sprintf("Quedan completadas: %s. Vuelven al Product Backlog: %s.", listOrNone(completedNames), listOrNone(pendingNames)),
		Action:   fmt.Sprintf("/proyectos/%d/sprints/cerrar", project.ID()),
		Fields:   map[string]string{"cierre": closedOn, "completadas": strings.Join(ids, ",")},
		Confirm:  "Sí, cerrar",
		Cancel:   fmt.Sprintf("/proyectos/%d/sprints/%d", project.ID(), sprint.Number()),
	}
}

// listOrNone junta los nombres separados por coma, o dice "ninguna" si no hay.
func listOrNone(names []string) string {
	if len(names) == 0 {
		return "ninguna"
	}
	return strings.Join(names, ", ")
}

// confirmClose pide confirmar antes de cerrar el sprint con la fecha y las historias
// elegidas. Una fecha que no es fecha vuelve al formulario sin preguntar nada (RN-14 de la
// #77); que este, la controla el dominio al cerrar.
func (s *Server) confirmClose(w http.ResponseWriter, r *http.Request) {
	project, sprint, err := s.openSprint(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	f := readForm(r, "cierre")
	f.date("cierre")
	completed := readCompleted(f, r.Form["completada"])
	if f.hasErrors() || f.General != "" {
		s.renderCloseForm(w, r, http.StatusUnprocessableEntity, project, sprint, completed, f)
		return
	}
	s.render(w, r, http.StatusOK, "confirmar.html", page{
		Title:   fmt.Sprintf("Cerrar el sprint %d", sprint.Number()),
		Project: project,
		Section: "sprints",
		Data:    closeConfirmation(project, sprint, f.Values["cierre"], completed),
	})
}

// closeSprint cierra el sprint abierto, despues de que se confirmo, y vuelve a su pagina con
// 303. Si el dominio lo rechaza, vuelve al formulario con lo elegido y el mensaje (RN-2).
func (s *Server) closeSprint(w http.ResponseWriter, r *http.Request) {
	project, sprint, err := s.openSprint(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	f := readForm(r, "cierre", "completadas")
	closedOn := f.date("cierre")
	var values []string
	if f.Values["completadas"] != "" {
		values = strings.Split(f.Values["completadas"], ",")
	}
	completed := readCompleted(f, values)
	if f.hasErrors() || f.General != "" {
		s.renderCloseForm(w, r, http.StatusUnprocessableEntity, project, sprint, completed, f)
		return
	}
	if err := project.CloseSprint(completed, closedOn); err != nil {
		f.reject(err)
		s.renderCloseForm(w, r, http.StatusUnprocessableEntity, project, sprint, completed, f)
		return
	}
	s.done(w, r, fmt.Sprintf("Se cerró el sprint %d", sprint.Number()), fmt.Sprintf("/proyectos/%d/sprints/%d", project.ID(), sprint.Number()))
}
