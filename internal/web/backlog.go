package web

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// priorityOptions son las prioridades que se ofrecen para elegir, con el nombre que les da
// el dominio (RN-11, D10). El dominio convierte el texto elegido con ParsePriority.
func priorityOptions() []string {
	return []string{domain.PriorityHigh.String(), domain.PriorityMedium.String(), domain.PriorityLow.String()}
}

// storyFields son los campos del formulario de historia.
var storyFields = []string{"titulo", "descripcion", "prioridad", "criterios"}

// splitLines separa los criterios de aceptacion, que se escriben uno por renglon. Los
// renglones en blanco los ignora el dominio, no la web.
func splitLines(text string) []string {
	return strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
}

// findStory busca el proyecto y la historia de los numeros que vienen en la direccion. Si
// el proyecto existe lo devuelve aunque la historia no, para que la barra lo siga mostrando.
func (s *Server) findStory(r *http.Request) (*domain.Project, *domain.Story, error) {
	project, err := s.findProject(r)
	if err != nil {
		return nil, nil, err
	}
	n, _ := strconv.Atoi(r.PathValue("n"))
	story, err := project.Story(n)
	if err != nil {
		return project, nil, err
	}
	return project, story, nil
}

// listBacklog muestra el Product Backlog del proyecto, en el orden del nucleo (punto 6 de
// la puerta). Sin historias muestra que esta vacio y ofrece cargar la primera (RN-7).
func (s *Server) listBacklog(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		s.notFound(w, r, nil, err)
		return
	}
	s.render(w, r, http.StatusOK, "backlog.html", page{
		Title:   "Backlog de " + project.Name(),
		Project: project,
		Section: "backlog",
		Data:    project,
	})
}

// storyForm es lo que muestra el formulario de historia, que sirve para cargar y para
// modificar: a donde se envia, el titulo, el boton, las prioridades y lo escrito.
type storyForm struct {
	Action     string
	Heading    string
	Submit     string
	Cancel     string
	Priorities []string
	Form       *form
}

// newStoryForm es el formulario para cargar una historia en el backlog del proyecto.
func newStoryForm(project *domain.Project, f *form) storyForm {
	return storyForm{
		Action:     fmt.Sprintf("/proyectos/%d/historias/nueva", project.ID()),
		Heading:    "Cargar historia",
		Submit:     "Cargar historia",
		Cancel:     fmt.Sprintf("/proyectos/%d/backlog", project.ID()),
		Priorities: priorityOptions(),
		Form:       f,
	}
}

// editStoryForm es el formulario para modificar una historia.
func editStoryForm(project *domain.Project, story *domain.Story, f *form) storyForm {
	storyPage := fmt.Sprintf("/proyectos/%d/historias/%d", project.ID(), story.ID())
	return storyForm{
		Action:     storyPage + "/editar",
		Heading:    fmt.Sprintf("Modificar la historia %d", story.ID()),
		Submit:     "Guardar cambios",
		Cancel:     storyPage,
		Priorities: priorityOptions(),
		Form:       f,
	}
}

// renderStoryForm muestra el formulario de historia con lo escrito y sus errores.
func (s *Server) renderStoryForm(w http.ResponseWriter, r *http.Request, status int, project *domain.Project, data storyForm) {
	s.render(w, r, status, "historia-form.html", page{
		Title:   data.Heading,
		Project: project,
		Section: "backlog",
		Data:    data,
	})
}

// newStory muestra el formulario para cargar una historia.
func (s *Server) newStory(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		s.notFound(w, r, nil, err)
		return
	}
	s.renderStoryForm(w, r, http.StatusOK, project, newStoryForm(project, newForm()))
}

// createStory carga la historia en el backlog. La prioridad se convierte en el dominio y,
// si no es valida, la rechaza AddStory en su propio orden (punto 7 de la puerta).
func (s *Server) createStory(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		s.notFound(w, r, nil, err)
		return
	}

	f := readForm(r, storyFields...)
	priority, _ := domain.ParsePriority(f.Values["prioridad"])
	story, err := project.AddStory(domain.StoryData{
		Title:              f.Values["titulo"],
		Description:        f.Values["descripcion"],
		Priority:           priority,
		AcceptanceCriteria: splitLines(f.Values["criterios"]),
	})
	if err != nil {
		f.reject(err)
		s.renderStoryForm(w, r, http.StatusUnprocessableEntity, project, newStoryForm(project, f))
		return
	}

	s.saver.Save(s.registry)
	setFlash(w, fmt.Sprintf("Se cargó la historia %d", story.ID()))
	http.Redirect(w, r, fmt.Sprintf("/proyectos/%d/backlog", project.ID()), http.StatusSeeOther)
}

// showStory muestra una historia del backlog con todos sus datos.
func (s *Server) showStory(w http.ResponseWriter, r *http.Request) {
	project, story, err := s.findStory(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	data := struct {
		ProjectID int
		Story     *domain.Story
	}{
		ProjectID: project.ID(),
		Story:     story,
	}
	s.render(w, r, http.StatusOK, "historia.html", page{
		Title:   story.Title(),
		Project: project,
		Section: "backlog",
		Data:    data,
	})
}

// editStory muestra el formulario para modificar una historia, precargado con sus datos.
func (s *Server) editStory(w http.ResponseWriter, r *http.Request) {
	project, story, err := s.findStory(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}

	f := newForm()
	f.Values["titulo"] = story.Title()
	f.Values["descripcion"] = story.Description()
	f.Values["prioridad"] = story.Priority().String()
	f.Values["criterios"] = strings.Join(story.AcceptanceCriteria(), "\n")
	s.renderStoryForm(w, r, http.StatusOK, project, editStoryForm(project, story, f))
}

// updateStory modifica la historia con lo que quedo en el formulario, que trae los cuatro
// datos. Si el dominio la rechaza —tambien si esta Terminada—, vuelve el formulario con lo
// escrito y el mensaje junto a su campo o arriba (RN-12, RN-13).
func (s *Server) updateStory(w http.ResponseWriter, r *http.Request) {
	project, story, err := s.findStory(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}

	f := readForm(r, storyFields...)
	title := f.Values["titulo"]
	description := f.Values["descripcion"]
	priority, _ := domain.ParsePriority(f.Values["prioridad"])
	criteria := splitLines(f.Values["criterios"])
	changes := domain.StoryChanges{
		Title:              &title,
		Description:        &description,
		Priority:           &priority,
		AcceptanceCriteria: &criteria,
	}

	if err := project.UpdateStory(story.ID(), changes); err != nil {
		f.reject(err)
		s.renderStoryForm(w, r, http.StatusUnprocessableEntity, project, editStoryForm(project, story, f))
		return
	}

	s.saver.Save(s.registry)
	setFlash(w, fmt.Sprintf("Se modificó la historia %d", story.ID()))
	http.Redirect(w, r, fmt.Sprintf("/proyectos/%d/historias/%d", project.ID(), story.ID()), http.StatusSeeOther)
}

// pointOptions es la escala de story points que se ofrece para elegir, la misma con la que
// valida el dominio (RN-11, D10).
func pointOptions() []string {
	var options []string
	for _, points := range domain.StoryPointScale() {
		options = append(options, strconv.Itoa(points))
	}
	return options
}

// renderEstimateForm muestra el formulario para estimar una historia.
func (s *Server) renderEstimateForm(w http.ResponseWriter, r *http.Request, status int, project *domain.Project, story *domain.Story, f *form) {
	data := struct {
		ProjectID int
		Story     *domain.Story
		Points    []string
		Form      *form
	}{
		ProjectID: project.ID(),
		Story:     story,
		Points:    pointOptions(),
		Form:      f,
	}
	s.render(w, r, status, "estimar.html", page{
		Title:   fmt.Sprintf("Estimar la historia %d", story.ID()),
		Project: project,
		Section: "backlog",
		Data:    data,
	})
}

// estimateForm muestra el formulario para estimar, con la estimacion actual elegida.
func (s *Server) estimateForm(w http.ResponseWriter, r *http.Request) {
	project, story, err := s.findStory(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	f := newForm()
	if story.IsEstimated() {
		f.Values["puntos"] = strconv.Itoa(story.StoryPoints())
	}
	s.renderEstimateForm(w, r, http.StatusOK, project, story, f)
}

// estimateStory estima la historia con los story points elegidos. Un valor fuera de la
// escala lo rechaza el dominio (RN-15); uno que no es un numero, la web (RN-14).
func (s *Server) estimateStory(w http.ResponseWriter, r *http.Request) {
	project, story, err := s.findStory(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}

	f := readForm(r, "puntos")
	points := f.number("puntos", errInvalidPoints)
	if f.hasErrors() {
		s.renderEstimateForm(w, r, http.StatusUnprocessableEntity, project, story, f)
		return
	}
	if err := project.EstimateStory(story.ID(), points); err != nil {
		f.reject(err)
		s.renderEstimateForm(w, r, http.StatusUnprocessableEntity, project, story, f)
		return
	}

	s.saver.Save(s.registry)
	setFlash(w, fmt.Sprintf("Se estimó la historia %d en %d story points", story.ID(), points))
	http.Redirect(w, r, fmt.Sprintf("/proyectos/%d/backlog", project.ID()), http.StatusSeeOther)
}

// removeEstimateConfirmation es la pagina que pide confirmar antes de quitar la estimacion.
func removeEstimateConfirmation(project *domain.Project, story *domain.Story) confirmation {
	storyPage := fmt.Sprintf("/proyectos/%d/historias/%d", project.ID(), story.ID())
	return confirmation{
		Question: fmt.Sprintf("¿Quitar la estimación de la historia %d?", story.ID()),
		Detail:   fmt.Sprintf("Está estimada en %d story points y va a quedar sin estimar.", story.StoryPoints()),
		Action:   storyPage + "/quitar-estimacion",
		Confirm:  "Sí, quitar la estimación",
		Cancel:   storyPage,
	}
}

// confirmRemoveEstimate pide confirmar antes de quitar la estimacion (RN-10, D9).
func (s *Server) confirmRemoveEstimate(w http.ResponseWriter, r *http.Request) {
	project, story, err := s.findStory(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	s.render(w, r, http.StatusOK, "confirmar.html", page{
		Title:   "Quitar estimación",
		Project: project,
		Section: "backlog",
		Data:    removeEstimateConfirmation(project, story),
	})
}

// removeEstimate quita la estimacion, despues de que se confirmo. Si el dominio lo rechaza,
// como en una historia Terminada, el motivo se muestra arriba de la confirmacion (RN-13).
func (s *Server) removeEstimate(w http.ResponseWriter, r *http.Request) {
	project, story, err := s.findStory(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	if err := project.RemoveStoryEstimate(story.ID()); err != nil {
		data := removeEstimateConfirmation(project, story)
		data.Error = err.Error()
		s.render(w, r, http.StatusUnprocessableEntity, "confirmar.html", page{
			Title:   "Quitar estimación",
			Project: project,
			Section: "backlog",
			Data:    data,
		})
		return
	}

	s.saver.Save(s.registry)
	setFlash(w, fmt.Sprintf("Se quitó la estimación de la historia %d", story.ID()))
	http.Redirect(w, r, fmt.Sprintf("/proyectos/%d/backlog", project.ID()), http.StatusSeeOther)
}
