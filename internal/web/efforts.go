package web

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// errInvalidHours es el error de forma de unas horas que no son un numero (D11). Que sean
// mayores que cero lo decide el dominio.
var errInvalidHours = errors.New("las horas tienen que ser un número, por ejemplo 2,5")

// effortFields son los campos del formulario de esfuerzo.
var effortFields = []string{"integrante", "historia", "dia", "horas", "actividad"}

// formatHours muestra unas horas con coma decimal y sin ceros de mas, como 2,5 (RN-5 de la
// #77).
func formatHours(hours float64) string {
	return strings.Replace(strconv.FormatFloat(hours, 'f', -1, 64), ".", ",", 1)
}

// decimalHours es la forma de unas horas: un numero, con decimales opcionales despues de una
// coma o de un punto. El signo menos se deja pasar para que el dominio diga que tienen que ser
// mayores que cero.
var decimalHours = regexp.MustCompile(`^-?[0-9]+([.,][0-9]+)?$`)

// hours lee unas horas escritas con coma o con punto (D11). Si no son un numero, marca el
// campo y devuelve cero, sin llamar al dominio. NaN e Inf tampoco son un numero que alguien
// escriba como horas, aunque la conversion de Go los acepte.
func (f *form) hours(name string) float64 {
	text := strings.TrimSpace(f.Values[name])
	// Solo numeros decimales, con coma o con punto: ParseFloat acepta tambien exponentes
	// (1e3) y hexadecimales (0x1p4), que nadie escribe como horas (defecto #192).
	if !decimalHours.MatchString(text) {
		f.fail(name, errInvalidHours.Error())
		return 0
	}
	value, err := strconv.ParseFloat(strings.Replace(text, ",", ".", 1), 64)
	if err != nil || math.IsNaN(value) || math.IsInf(value, 0) {
		f.fail(name, errInvalidHours.Error())
		return 0
	}
	return value
}

// memberOptions son los integrantes del proyecto para elegir, por su nombre (D10).
func memberOptions(project *domain.Project) []string {
	var names []string
	for _, member := range project.Members() {
		names = append(names, member.Name())
	}
	return names
}

// total es una fila de los totales de horas: de que es y cuantas horas suma.
type total struct {
	Name  string
	Hours float64
}

// effortPage es lo que muestra la pagina de esfuerzo: los registros con el titulo de su
// historia, y los totales por historia y por integrante.
type effortPage struct {
	ProjectID int
	Efforts   []effortRow
	ByStory   []total
	ByMember  []total
}

// effortRow es un registro de esfuerzo con el titulo de su historia, para leerlo en la tabla.
type effortRow struct {
	*domain.Effort
	StoryTitle string
}

// listEfforts muestra los registros de esfuerzo del proyecto y sus totales, que suma el
// nucleo (#14). Sin registros muestra que no hay y ofrece registrar los primeros.
func (s *Server) listEfforts(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		s.notFound(w, r, nil, err)
		return
	}
	data := effortPage{ProjectID: project.ID()}
	for _, effort := range project.Efforts() {
		row := effortRow{Effort: effort}
		if story, err := project.Story(effort.StoryID()); err == nil {
			row.StoryTitle = story.Title()
		}
		data.Efforts = append(data.Efforts, row)
	}
	for _, story := range project.Backlog() {
		if hours, err := project.StoryHours(story.ID()); err == nil && hours > 0 {
			data.ByStory = append(data.ByStory, total{Name: fmt.Sprintf("%d — %s", story.ID(), story.Title()), Hours: hours})
		}
	}
	for _, member := range project.Members() {
		if hours, err := project.MemberHours(member.Name()); err == nil && hours > 0 {
			data.ByMember = append(data.ByMember, total{Name: member.Name(), Hours: hours})
		}
	}
	s.render(w, r, http.StatusOK, "esfuerzo.html", page{
		Title:   "Esfuerzo de " + project.Name(),
		Project: project,
		Section: "esfuerzo",
		Data:    data,
	})
}

// effortForm es lo que muestra el formulario de esfuerzo: las listas para elegir y lo
// escrito.
type effortForm struct {
	ProjectID int
	Members   []string
	Stories   []option
	Form      *form
}

// renderEffortForm muestra el formulario para registrar esfuerzo.
func (s *Server) renderEffortForm(w http.ResponseWriter, r *http.Request, status int, project *domain.Project, f *form) {
	s.render(w, r, status, "esfuerzo-form.html", page{
		Title:   "Registrar esfuerzo",
		Project: project,
		Section: "esfuerzo",
		Data: effortForm{
			ProjectID: project.ID(),
			Members:   memberOptions(project),
			Stories:   storyOptions(project),
			Form:      f,
		},
	})
}

// newEffort muestra el formulario para registrar esfuerzo, con el dia de hoy propuesto
// (RN-11).
func (s *Server) newEffort(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		s.notFound(w, r, nil, err)
		return
	}
	f := newForm()
	f.Values["dia"] = s.today().Format(dateLayout)
	s.renderEffortForm(w, r, http.StatusOK, project, f)
}

// createEffort registra el esfuerzo con la fecha de hoy para que el dominio rechace un dia
// que todavia no paso (RN-12), y vuelve al esfuerzo con 303.
func (s *Server) createEffort(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		s.notFound(w, r, nil, err)
		return
	}
	f := readForm(r, effortFields...)
	storyID := f.number("historia", errInvalidStory)
	day := f.date("dia")
	hours := f.hours("horas")
	if f.hasErrors() {
		s.renderEffortForm(w, r, http.StatusUnprocessableEntity, project, f)
		return
	}

	effort, err := project.LogEffort(domain.EffortData{
		StoryID:    storyID,
		MemberName: f.Values["integrante"],
		Activity:   f.Values["actividad"],
		Hours:      hours,
		Day:        day,
	}, s.today())
	if err != nil {
		f.reject(err)
		s.renderEffortForm(w, r, http.StatusUnprocessableEntity, project, f)
		return
	}

	message := fmt.Sprintf("Se registraron %s horas de %s en la historia %d", formatHours(effort.Hours()), effort.MemberName(), effort.StoryID())
	s.done(w, r, message, fmt.Sprintf("/proyectos/%d/esfuerzo", project.ID()))
}

// renderEstimatedHoursForm muestra el formulario para cargar las horas estimadas de una
// historia.
func (s *Server) renderEstimatedHoursForm(w http.ResponseWriter, r *http.Request, status int, project *domain.Project, story *domain.Story, f *form) {
	s.render(w, r, status, "horas-estimadas.html", page{
		Title:   fmt.Sprintf("Horas estimadas de la historia %d", story.ID()),
		Project: project,
		Section: "backlog",
		Data: struct {
			ProjectID int
			Story     *domain.Story
			Form      *form
		}{ProjectID: project.ID(), Story: story, Form: f},
	})
}

// estimatedHoursForm muestra el formulario de horas estimadas, con las que ya tenga.
func (s *Server) estimatedHoursForm(w http.ResponseWriter, r *http.Request) {
	project, story, err := s.findStory(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	f := newForm()
	if story.EstimatedHours() > 0 {
		f.Values["horas"] = formatHours(story.EstimatedHours())
	}
	s.renderEstimatedHoursForm(w, r, http.StatusOK, project, story, f)
}

// estimateHours carga las horas estimadas de la historia (#16) y vuelve a ella con 303.
func (s *Server) estimateHours(w http.ResponseWriter, r *http.Request) {
	project, story, err := s.findStory(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	f := readForm(r, "horas")
	hours := f.hours("horas")
	if f.hasErrors() {
		s.renderEstimatedHoursForm(w, r, http.StatusUnprocessableEntity, project, story, f)
		return
	}
	if err := project.EstimateHours(story.ID(), hours); err != nil {
		f.reject(err)
		s.renderEstimatedHoursForm(w, r, http.StatusUnprocessableEntity, project, story, f)
		return
	}
	s.done(w, r, fmt.Sprintf("Se cargaron las horas estimadas de la historia %d", story.ID()), fmt.Sprintf("/proyectos/%d/historias/%d", project.ID(), story.ID()))
}
