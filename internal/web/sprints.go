package web

import (
	"net/http"
)

// listSprints muestra los sprints del proyecto.
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
