package web

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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

// projectForm es lo que muestra el formulario de proyecto, que sirve para crear y para
// modificar: a donde se envia, el titulo, el boton y lo escrito.
type projectForm struct {
	Action  string
	Heading string
	Submit  string
	Cancel  string
	Form    *form
}

// renderProjectForm muestra el formulario de proyecto. Si se esta modificando, project es
// el proyecto, para la barra; si se esta creando, es nil.
func (s *Server) renderProjectForm(w http.ResponseWriter, r *http.Request, status int, project *domain.Project, data projectForm) {
	p := page{Title: data.Heading, Data: data}
	if project != nil {
		p.Project = project
		p.Section = "datos"
	}
	s.render(w, r, status, "proyecto-form.html", p)
}

// newProjectForm es el formulario para crear un proyecto.
func newProjectForm(f *form) projectForm {
	return projectForm{
		Action:  "/proyectos/nuevo",
		Heading: "Nuevo proyecto",
		Submit:  "Crear proyecto",
		Cancel:  "/proyectos",
		Form:    f,
	}
}

// editProjectForm es el formulario para modificar un proyecto.
func editProjectForm(project *domain.Project, f *form) projectForm {
	return projectForm{
		Action:  fmt.Sprintf("/proyectos/%d/editar", project.ID()),
		Heading: "Modificar proyecto",
		Submit:  "Guardar cambios",
		Cancel:  fmt.Sprintf("/proyectos/%d", project.ID()),
		Form:    f,
	}
}

// newProject muestra el formulario para crear un proyecto.
func (s *Server) newProject(w http.ResponseWriter, r *http.Request) {
	s.renderProjectForm(w, r, http.StatusOK, nil, newProjectForm(newForm()))
}

// createProject crea el proyecto con los datos del formulario y redirige con 303 a su
// pagina, asi recargar no lo crea de nuevo (RN-9, D5). Si el dominio lo rechaza, vuelve a
// mostrar el formulario con lo escrito y el mensaje junto a su campo (RN-12, D6).
func (s *Server) createProject(w http.ResponseWriter, r *http.Request) {
	f := readForm(r, projectFields...)
	start := f.date("inicio")
	end := f.date("fin")
	if f.hasErrors() {
		s.renderProjectForm(w, r, http.StatusUnprocessableEntity, nil, newProjectForm(f))
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
		s.renderProjectForm(w, r, http.StatusUnprocessableEntity, nil, newProjectForm(f))
		return
	}

	s.saver.Save(s.registry)
	setFlash(w, "Se creó el proyecto "+project.Name())
	http.Redirect(w, r, fmt.Sprintf("/proyectos/%d", project.ID()), http.StatusSeeOther)
}

// findProject busca el proyecto del numero que viene en la direccion.
func (s *Server) findProject(r *http.Request) (*domain.Project, error) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	return s.registry.Get(id)
}

// showProject muestra la pagina de un proyecto.
func (s *Server) showProject(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
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

// editProject muestra el formulario para modificar un proyecto, precargado con sus datos.
func (s *Server) editProject(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	f := newForm()
	f.Values["nombre"] = project.Name()
	f.Values["descripcion"] = project.Description()
	f.Values["inicio"] = project.Start().Format(dateLayout)
	if project.HasEnd() {
		f.Values["fin"] = project.End().Format(dateLayout)
	}
	s.renderProjectForm(w, r, http.StatusOK, project, editProjectForm(project, f))
}

// updateProject modifica el proyecto con lo que quedo en el formulario. El formulario trae
// todos los datos: los obligatorios se mandan tal cual, y un dato opcional que quedo vacio
// se quita (caso limite de la #3).
func (s *Server) updateProject(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	f := readForm(r, projectFields...)
	start := f.date("inicio")
	end := f.date("fin")
	if f.hasErrors() {
		s.renderProjectForm(w, r, http.StatusUnprocessableEntity, project, editProjectForm(project, f))
		return
	}

	name := f.Values["nombre"]
	changes := domain.ProjectChanges{
		Name:        &name,
		Start:       &start,
		Description: domain.SetTo(f.Values["descripcion"]),
		End:         domain.SetTo(end),
	}
	if f.Values["descripcion"] == "" {
		changes.Description = domain.Remove[string]()
	}
	if end.IsZero() {
		changes.End = domain.Remove[time.Time]()
	}

	if err := s.registry.Update(project.ID(), changes); err != nil {
		f.reject(err)
		s.renderProjectForm(w, r, http.StatusUnprocessableEntity, project, editProjectForm(project, f))
		return
	}

	s.saver.Save(s.registry)
	setFlash(w, "Se modificó el proyecto "+project.Name())
	http.Redirect(w, r, fmt.Sprintf("/proyectos/%d", project.ID()), http.StatusSeeOther)
}
