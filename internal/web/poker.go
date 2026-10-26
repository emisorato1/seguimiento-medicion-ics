package web

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// errRoundNotFound es para el numero de una ronda que la historia no tiene. El nucleo da la
// ronda actual de una historia, y la web la busca por su numero en la direccion.
var errRoundNotFound = errors.New("no se encontró la ronda")

// roundSummary es una historia del backlog con su ronda actual, si tiene, para la pagina de
// Planning Poker.
type roundSummary struct {
	Story *domain.Story
	Round *domain.VotingRound
}

// pokerPage es lo que muestra la pagina de Planning Poker del proyecto: cada historia con su
// ronda actual y el formulario para abrir una.
type pokerPage struct {
	ProjectID int
	Stories   []roundSummary
	Choices   []option
	Form      *form
}

// renderPoker muestra la pagina de Planning Poker del proyecto.
func (s *Server) renderPoker(w http.ResponseWriter, r *http.Request, status int, project *domain.Project, f *form) {
	data := pokerPage{ProjectID: project.ID(), Choices: storyOptions(project), Form: f}
	for _, story := range project.Backlog() {
		summary := roundSummary{Story: story}
		if round, ok, err := project.CurrentVotingRound(story.ID()); err == nil && ok {
			summary.Round = round
		}
		data.Stories = append(data.Stories, summary)
	}
	s.render(w, r, status, "planning-poker.html", page{
		Title:   "Planning Poker de " + project.Name(),
		Project: project,
		Section: "planning-poker",
		Data:    data,
	})
}

// showPoker muestra la pagina de Planning Poker del proyecto.
func (s *Server) showPoker(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		s.notFound(w, r, nil, err)
		return
	}
	s.renderPoker(w, r, http.StatusOK, project, newForm())
}

// roundPath es la direccion de una ronda.
func roundPath(projectID int, round *domain.VotingRound) string {
	return fmt.Sprintf("/proyectos/%d/historias/%d/rondas/%d", projectID, round.StoryID(), round.Number())
}

// openRound abre una ronda sobre la historia elegida (#13) y va a ella con 303. Si el
// dominio la rechaza, vuelve a mostrar la pagina con el mensaje.
func (s *Server) openRound(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		s.notFound(w, r, nil, err)
		return
	}
	f := readForm(r, "historia")
	id := f.number("historia", errInvalidStory)
	if f.hasErrors() {
		s.renderPoker(w, r, http.StatusUnprocessableEntity, project, f)
		return
	}
	round, err := project.OpenVotingRound(id)
	if err != nil {
		f.reject(err)
		s.renderPoker(w, r, http.StatusUnprocessableEntity, project, f)
		return
	}
	s.done(w, r, fmt.Sprintf("Se abrió la ronda %d de la historia %d", round.Number(), round.StoryID()), roundPath(project.ID(), round))
}

// findRound busca el proyecto, la historia y la ronda de la direccion entre las rondas de la
// historia (VotingRounds, #73), asi tambien se ven las anteriores. Devuelve ademas si es la
// ultima, la unica sobre la que se puede acordar.
func (s *Server) findRound(r *http.Request) (*domain.Project, *domain.Story, *domain.VotingRound, bool, error) {
	project, story, err := s.findStory(r)
	if err != nil {
		return project, nil, nil, false, err
	}
	n, _ := strconv.Atoi(r.PathValue("r"))
	rounds, err := project.VotingRounds(story.ID())
	if err != nil {
		return project, story, nil, false, err
	}
	for i, round := range rounds {
		if round.Number() == n {
			return project, story, round, i == len(rounds)-1, nil
		}
	}
	return project, story, nil, false, errRoundNotFound
}

// differencesText arma la frase de las diferencias de una ronda revelada tal como las
// informa el nucleo: el voto mas bajo y el mas alto con quienes los dieron (RN-8).
func differencesText(d domain.Differences) string {
	if !d.HasDifferences {
		return fmt.Sprintf("No hay diferencias: todos votaron %d.", d.Lowest.Points)
	}
	return fmt.Sprintf("Hay diferencias: el más bajo es %d, de %s, y el más alto es %d, de %s.",
		d.Lowest.Points, strings.Join(d.Lowest.Members, " y "), d.Highest.Points, strings.Join(d.Highest.Members, " y "))
}

// roundPage es lo que muestra la pagina de una ronda. Antes de revelar, solo lo que el
// nucleo deja leer: quien voto y quien falta (RN-4, D2).
type roundPage struct {
	ProjectID int
	Story     *domain.Story
	Round     *domain.VotingRound
	Voted     []string
	Pending   []string
	// Participants es cuantos participan: los que votaron mas los que faltan.
	Participants int
	// Votes son los votos con su integrante, que el nucleo entrega recien al revelar, y
	// Differences la frase que los compara.
	Votes       []domain.Vote
	Differences string
	// Last dice si es la ultima ronda de la historia: solo sobre esa se acuerda (#73).
	Last   bool
	Points []string
	Form   *form
}

// renderRound muestra la pagina de una ronda.
func (s *Server) renderRound(w http.ResponseWriter, r *http.Request, status int, project *domain.Project, story *domain.Story, round *domain.VotingRound, last bool, f *form) {
	// Antes de revelar, Votes y Differences responden con un error y no hay nada que mostrar
	// (D2).
	votes, _ := round.Votes()
	var differences string
	if d, err := round.Differences(); err == nil {
		differences = differencesText(d)
	}
	s.render(w, r, status, "ronda.html", page{
		Title:   fmt.Sprintf("Ronda %d de la historia %d", round.Number(), story.ID()),
		Project: project,
		Section: "planning-poker",
		Data: roundPage{
			ProjectID:    project.ID(),
			Story:        story,
			Round:        round,
			Voted:        round.Voted(),
			Pending:      round.Pending(),
			Participants: len(round.Voted()) + len(round.Pending()),
			Votes:        votes,
			Differences:  differences,
			Last:         last,
			Points:       pointOptions(),
			Form:         f,
		},
	})
}

// showRound muestra una ronda de una historia.
func (s *Server) showRound(w http.ResponseWriter, r *http.Request) {
	project, story, round, last, err := s.findRound(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	s.renderRound(w, r, http.StatusOK, project, story, round, last, newForm())
}

// vote registra el voto elegido y vuelve a la ronda con 303. El mensaje nombra a quien voto
// y no lleva el valor (RN-5).
func (s *Server) vote(w http.ResponseWriter, r *http.Request) {
	project, story, round, last, err := s.findRound(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	f := readForm(r, "integrante", "puntos")
	points := f.number("puntos", errInvalidPoints)
	if f.hasErrors() {
		f.Values["puntos"] = ""
		s.renderRound(w, r, http.StatusUnprocessableEntity, project, story, round, last, f)
		return
	}
	if err := round.Vote(f.Values["integrante"], points); err != nil {
		// El formulario vuelve sin el valor elegido: mostrarlo romperia el secreto delante
		// de quien este mirando (RN-6, D3). Es la unica excepcion a volver con lo escrito.
		f.Values["puntos"] = ""
		f.reject(err)
		s.renderRound(w, r, http.StatusUnprocessableEntity, project, story, round, last, f)
		return
	}
	s.done(w, r, "Se registró el voto de "+f.Values["integrante"], roundPath(project.ID(), round))
}

// reveal revela la ronda y vuelve a ella con 303. Si faltan votos, el rechazo no es de
// ningun campo y va arriba (RN-2).
func (s *Server) reveal(w http.ResponseWriter, r *http.Request) {
	project, story, round, last, err := s.findRound(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	if err := round.Reveal(); err != nil {
		f := newForm()
		f.reject(err)
		s.renderRound(w, r, http.StatusUnprocessableEntity, project, story, round, last, f)
		return
	}
	s.done(w, r, fmt.Sprintf("Se reveló la ronda %d de la historia %d", round.Number(), story.ID()), roundPath(project.ID(), round))
}

// agree registra la estimacion acordada de la ronda y vuelve a ella con 303 (#73). Si el
// dominio la rechaza, vuelve a mostrar la ronda con el mensaje: junto a los story points si
// es de la escala, arriba si es de la ronda.
func (s *Server) agree(w http.ResponseWriter, r *http.Request) {
	project, story, round, last, err := s.findRound(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	f := readForm(r, "puntos")
	points := f.number("puntos", errInvalidPoints)
	if f.hasErrors() {
		s.renderRound(w, r, http.StatusUnprocessableEntity, project, story, round, last, f)
		return
	}
	if err := round.Agree(points); err != nil {
		f.reject(err)
		s.renderRound(w, r, http.StatusUnprocessableEntity, project, story, round, last, f)
		return
	}
	s.done(w, r, fmt.Sprintf("Se acordó la historia %d en %d story points", story.ID(), points), roundPath(project.ID(), round))
}

// listRounds muestra las rondas de una historia, en orden, con su acuerdo si lo tienen (#73).
func (s *Server) listRounds(w http.ResponseWriter, r *http.Request) {
	project, story, err := s.findStory(r)
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	rounds, err := project.VotingRounds(story.ID())
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}
	data := struct {
		ProjectID int
		Story     *domain.Story
		Rounds    []*domain.VotingRound
	}{ProjectID: project.ID(), Story: story, Rounds: rounds}
	s.render(w, r, http.StatusOK, "rondas.html", page{
		Title:   fmt.Sprintf("Rondas de la historia %d", story.ID()),
		Project: project,
		Section: "planning-poker",
		Data:    data,
	})
}
