package web

import (
	"net/http"
	"strconv"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

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

// showStory muestra una historia del backlog.
func (s *Server) showStory(w http.ResponseWriter, r *http.Request) {
	project, story, err := s.findStory(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	s.render(w, r, http.StatusOK, "historia.html", page{
		Title:   story.Title(),
		Project: project,
		Section: "backlog",
		Data:    story,
	})
}
