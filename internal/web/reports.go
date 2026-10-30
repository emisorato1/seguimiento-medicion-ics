package web

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// La pagina del reporte (historia #78, tareas 20 y 21). El reporte lo arma el nucleo con sus
// lineas ya escritas (#18): la web lo muestra tal cual, sin reformatear nada (D6).

// reportPage es lo que muestra la pagina del reporte: el reporte, o nada si el sprint pedido
// no tiene la forma de un numero; de que sprint es (0 para el del proyecto), que tambien arma
// el acceso al PDF; la lista para elegir el sprint, el formulario con lo elegido y su error, y
// el aviso de que el PDF no se pudo generar (RN-19).
type reportPage struct {
	ProjectID int
	Report    *domain.Report
	Sprint    int
	Sprints   []option
	Form      *form
	Error     string
}

// readReport arma el reporte que pide la direccion, con la fecha de hoy (RN-12): sin sprint,
// el del proyecto, y con sprint, el de ese sprint. Un sprint que no es un numero se marca
// junto al campo y no se arma ningun reporte (RN-2); uno que no existe lo informa el nucleo,
// con su error.
func (s *Server) readReport(r *http.Request, project *domain.Project) (reportPage, error) {
	data := reportPage{ProjectID: project.ID(), Sprints: sprintOptions(project), Form: readForm(r, "sprint")}
	if strings.TrimSpace(data.Form.Values["sprint"]) == "" {
		report := project.ProjectReport(s.today())
		data.Report = &report
		return data, nil
	}
	n := data.Form.number("sprint", errInvalidSprint)
	if data.Form.hasErrors() {
		return data, nil
	}
	report, err := project.SprintReport(n, s.today())
	if err != nil {
		return data, err
	}
	data.Report, data.Sprint = &report, n
	return data, nil
}

// renderReport muestra la pagina del reporte. El titulo de la pestaña es el del reporte, o
// "Reporte" si no se pudo armar.
func (s *Server) renderReport(w http.ResponseWriter, r *http.Request, status int, project *domain.Project, data reportPage) {
	title := "Reporte"
	if data.Report != nil {
		title = data.Report.Title
	}
	s.render(w, r, status, "reporte.html", page{Title: title, Project: project, Section: "reporte", Data: data})
}

// showReport muestra el reporte del proyecto o el de un sprint, con su titulo, su fecha de
// generacion y sus secciones en el orden y con las lineas del nucleo (RN-14). Solo lee, asi
// que va con el candado de lectura (#77, D12). Un proyecto o un sprint que no existe se
// informa con el mensaje del nucleo y el acceso para volver (RN-18).
func (s *Server) showReport(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		s.notFound(w, r, nil, err)
		return
	}
	data, err := s.readReport(r, project)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	status := http.StatusOK
	if data.Form.hasErrors() {
		status = http.StatusUnprocessableEntity
	}
	s.renderReport(w, r, status, project, data)
}

// fileName es el nombre del archivo del PDF: reporte-proyecto-N.pdf o
// reporte-proyecto-N-sprint-M.pdf (RN-15).
func (p reportPage) fileName() string {
	if p.Sprint == 0 {
		return fmt.Sprintf("reporte-proyecto-%d.pdf", p.ProjectID)
	}
	return fmt.Sprintf("reporte-proyecto-%d-sprint-%d.pdf", p.ProjectID, p.Sprint)
}

// downloadReport entrega el PDF del mismo reporte que muestra la pagina, como un archivo para
// guardar con su nombre (RN-15). Lo arma entero en memoria con la exportacion de la #19 antes
// de enviar nada (D5): si no se pudo generar, vuelve a la pagina del reporte con el mensaje
// del nucleo y no entrega ningun archivo, ni siquiera lo que se llego a escribir (RN-19). Lo
// que no existe o no tiene la forma de un numero se informa como en la pagina.
func (s *Server) downloadReport(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		s.notFound(w, r, nil, err)
		return
	}
	data, err := s.readReport(r, project)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	if data.Report == nil {
		s.renderReport(w, r, http.StatusUnprocessableEntity, project, data)
		return
	}
	var pdf bytes.Buffer
	if err := s.exportPDF(*data.Report, &pdf); err != nil {
		data.Error = err.Error()
		s.renderReport(w, r, http.StatusInternalServerError, project, data)
		return
	}
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", `attachment; filename="`+data.fileName()+`"`)
	w.Header().Set("Content-Length", strconv.Itoa(pdf.Len()))
	pdf.WriteTo(w)
}
