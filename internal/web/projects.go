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
	s.render(w, http.StatusOK, "proyectos.html", page{Title: "Proyectos", Data: data})
}

// newProject muestra el formulario para crear un proyecto.
func (s *Server) newProject(w http.ResponseWriter, r *http.Request) {
	s.render(w, http.StatusOK, "proyecto-nuevo.html", page{Title: "Nuevo proyecto"})
}

// createProject crea el proyecto con los datos del formulario y redirige con 303 a su
// pagina, asi recargar no lo crea de nuevo (RN-9, D5).
func (s *Server) createProject(w http.ResponseWriter, r *http.Request) {
	start, _ := time.Parse("2006-01-02", r.FormValue("inicio"))
	end, _ := time.Parse("2006-01-02", r.FormValue("fin"))

	project, err := s.registry.Create(domain.ProjectData{
		Name:        r.FormValue("nombre"),
		Description: r.FormValue("descripcion"),
		Start:       start,
		End:         end,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	s.saver.Save(s.registry)
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
	s.render(w, http.StatusOK, "proyecto.html", page{Title: project.Name(), Data: project})
}
