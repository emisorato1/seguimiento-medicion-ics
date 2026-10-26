package web

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// errInvalidSprint es el error de forma del sprint elegido: no se eligio ninguno de la lista,
// o llego algo que no es un numero (RN-2). Que exista lo decide el dominio.
var errInvalidSprint = errors.New("hay que elegir el sprint de la lista")

// defectFields son los campos del formulario de defecto.
var defectFields = []string{"titulo", "descripcion", "severidad", "historia", "sprint"}

// severityOptions son las severidades para elegir, con el nombre que les da el dominio
// (D10). El dominio convierte el texto elegido con ParseSeverity.
func severityOptions() []string {
	return []string{domain.SeverityCritical.String(), domain.SeverityMajor.String(), domain.SeverityMinor.String()}
}

// stateOptions son los estados a los que se puede pasar un defecto cambiandole el estado:
// a Resuelto se llega resolviendolo, con su sprint (Entradas). El dominio convierte el
// texto elegido con ParseDefectState.
func stateOptions() []string {
	return []string{domain.DefectOpen.String(), domain.DefectInProgress.String()}
}

// sprintOptions son los sprints del proyecto para elegir, todos, abiertos o cerrados (D10):
// se elige por el numero y se lee con el Sprint Goal.
func sprintOptions(project *domain.Project) []option {
	var options []option
	for _, sprint := range project.Sprints() {
		options = append(options, option{Value: strconv.Itoa(sprint.Number()), Label: fmt.Sprintf("%d — %s", sprint.Number(), sprint.Goal())})
	}
	return options
}

// listDefects muestra los defectos del proyecto, en el orden del nucleo. Sin defectos
// muestra que no hay y ofrece cargar el primero (RN-7 de la #77).
func (s *Server) listDefects(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		s.notFound(w, r, nil, err)
		return
	}
	data := struct {
		ProjectID int
		Defects   []*domain.Defect
	}{
		ProjectID: project.ID(),
		Defects:   project.Defects(),
	}
	s.render(w, r, http.StatusOK, "defectos.html", page{
		Title:   "Defectos de " + project.Name(),
		Project: project,
		Section: "defectos",
		Data:    data,
	})
}

// defectForm es lo que muestra el formulario de defecto: lo escrito y las listas para elegir.
type defectForm struct {
	ProjectID  int
	Severities []string
	Stories    []option
	Sprints    []option
	Form       *form
}

// renderDefectForm muestra el formulario para cargar un defecto, con lo escrito y sus errores.
// Sin sprints dice que falta uno y ofrece crearlo (RN-17).
func (s *Server) renderDefectForm(w http.ResponseWriter, r *http.Request, status int, project *domain.Project, f *form) {
	s.render(w, r, status, "defecto-form.html", page{
		Title:   "Nuevo defecto",
		Project: project,
		Section: "defectos",
		Data: defectForm{
			ProjectID:  project.ID(),
			Severities: severityOptions(),
			Stories:    storyOptions(project),
			Sprints:    sprintOptions(project),
			Form:       f,
		},
	})
}

// newDefect muestra el formulario para cargar un defecto.
func (s *Server) newDefect(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		s.notFound(w, r, nil, err)
		return
	}
	s.renderDefectForm(w, r, http.StatusOK, project, newForm())
}

// createDefect carga el defecto con los datos del formulario y vuelve a la lista con 303. Si
// el dominio lo rechaza, vuelve a mostrar el formulario con lo escrito y el mensaje (RN-2).
func (s *Server) createDefect(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		s.notFound(w, r, nil, err)
		return
	}
	f := readForm(r, defectFields...)
	storyID := f.number("historia", errInvalidStory)
	sprint := f.number("sprint", errInvalidSprint)
	if f.hasErrors() {
		s.renderDefectForm(w, r, http.StatusUnprocessableEntity, project, f)
		return
	}
	severity, _ := domain.ParseSeverity(f.Values["severidad"])

	defect, err := project.AddDefect(domain.DefectData{
		Title:       f.Values["titulo"],
		Description: f.Values["descripcion"],
		Severity:    severity,
		StoryID:     storyID,
		DetectedIn:  sprint,
	})
	if err != nil {
		f.reject(err)
		s.renderDefectForm(w, r, http.StatusUnprocessableEntity, project, f)
		return
	}

	s.done(w, r, fmt.Sprintf("Se cargó el defecto %d", defect.ID()), fmt.Sprintf("/proyectos/%d/defectos", project.ID()))
}

// findDefect busca el proyecto y el defecto con el numero de la direccion. Que no exista lo
// informa el dominio con su error (RN-18).
func (s *Server) findDefect(r *http.Request) (*domain.Project, *domain.Defect, error) {
	project, err := s.findProject(r)
	if err != nil {
		return nil, nil, err
	}
	n, _ := strconv.Atoi(r.PathValue("n"))
	for _, defect := range project.Defects() {
		if defect.ID() == n {
			return project, defect, nil
		}
	}
	return project, nil, domain.ErrDefectNotFound
}

// defectPage es lo que muestra la pagina de un defecto: sus datos y, si no esta resuelto,
// los formularios para cambiarle el estado y para resolverlo.
type defectPage struct {
	ProjectID int
	Defect    *domain.Defect
	Resolved  bool
	States    []string
	Sprints   []option
	Form      *form
}

// renderDefect muestra la pagina de un defecto.
func (s *Server) renderDefect(w http.ResponseWriter, r *http.Request, status int, project *domain.Project, defect *domain.Defect, f *form) {
	s.render(w, r, status, "defecto.html", page{
		Title:   fmt.Sprintf("Defecto %d", defect.ID()),
		Project: project,
		Section: "defectos",
		Data: defectPage{
			ProjectID: project.ID(),
			Defect:    defect,
			Resolved:  defect.State() == domain.DefectResolved,
			States:    stateOptions(),
			Sprints:   sprintOptions(project),
			Form:      f,
		},
	})
}

// showDefect muestra un defecto.
func (s *Server) showDefect(w http.ResponseWriter, r *http.Request) {
	project, defect, err := s.findDefect(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	s.renderDefect(w, r, http.StatusOK, project, defect, newForm())
}

// changeDefectState pasa el defecto al estado elegido y vuelve a su pagina con 303. Si el
// dominio lo rechaza, vuelve a mostrar la pagina con el mensaje (RN-2).
func (s *Server) changeDefectState(w http.ResponseWriter, r *http.Request) {
	project, defect, err := s.findDefect(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	f := readForm(r, "estado")
	state, err := domain.ParseDefectState(f.Values["estado"])
	if err == nil {
		err = project.ChangeDefectState(defect.ID(), state)
	}
	if err != nil {
		f.reject(err)
		s.renderDefect(w, r, http.StatusUnprocessableEntity, project, defect, f)
		return
	}
	s.done(w, r, fmt.Sprintf("El defecto %d pasó a %s", defect.ID(), state), fmt.Sprintf("/proyectos/%d/defectos/%d", project.ID(), defect.ID()))
}

// resolveConfirmation es la pagina que pide confirmar antes de resolver un defecto, que no
// se deshace (RN-9, D9).
func resolveConfirmation(project *domain.Project, defect *domain.Defect, sprint string) confirmation {
	return confirmation{
		Question: fmt.Sprintf("¿Resolver el defecto %d en el sprint %s?", defect.ID(), sprint),
		Detail:   fmt.Sprintf("«%s» queda Resuelto y ya no cambia de estado.", defect.Title()),
		Action:   fmt.Sprintf("/proyectos/%d/defectos/%d/resolver", project.ID(), defect.ID()),
		Fields:   map[string]string{"sprint": sprint},
		Confirm:  "Sí, resolver",
		Cancel:   fmt.Sprintf("/proyectos/%d/defectos/%d", project.ID(), defect.ID()),
	}
}

// confirmResolve pide confirmar antes de resolver el defecto en el sprint elegido.
func (s *Server) confirmResolve(w http.ResponseWriter, r *http.Request) {
	project, defect, err := s.findDefect(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	s.render(w, r, http.StatusOK, "confirmar.html", page{
		Title:   "Resolver defecto",
		Project: project,
		Section: "defectos",
		Data:    resolveConfirmation(project, defect, r.FormValue("sprint")),
	})
}

// resolveDefect resuelve el defecto en el sprint elegido, despues de que se confirmo, y
// vuelve a su pagina con 303. Si el dominio lo rechaza, vuelve a pedir confirmar con su
// mensaje.
func (s *Server) resolveDefect(w http.ResponseWriter, r *http.Request) {
	project, defect, err := s.findDefect(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	sprint, _ := strconv.Atoi(r.FormValue("sprint"))
	if err := project.ResolveDefect(defect.ID(), sprint); err != nil {
		data := resolveConfirmation(project, defect, r.FormValue("sprint"))
		data.Error = err.Error()
		s.render(w, r, http.StatusUnprocessableEntity, "confirmar.html", page{
			Title:   "Resolver defecto",
			Project: project,
			Section: "defectos",
			Data:    data,
		})
		return
	}
	s.done(w, r, fmt.Sprintf("Se resolvió el defecto %d en el sprint %d", defect.ID(), sprint), fmt.Sprintf("/proyectos/%d/defectos/%d", project.ID(), defect.ID()))
}
