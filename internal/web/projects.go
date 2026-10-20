package web

import "net/http"

// listProjects muestra la lista de proyectos, que es la pagina de inicio.
func (s *Server) listProjects(w http.ResponseWriter, r *http.Request) {
	s.render(w, http.StatusOK, "proyectos.html", page{Title: "Proyectos"})
}
