package web

import (
	"net/http"

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
