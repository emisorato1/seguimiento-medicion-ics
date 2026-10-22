package domain

import (
	"errors"
	"slices"
	"testing"
)

// projectWithThreeMembers devuelve un proyecto con "Agustin Salinas", "Juliana Bustos" y
// "Maximiliano Eula", cargados en ese orden, y la historia 1 Pendiente y sin estimar: el
// "Dado" de la mayoria de los escenarios de la #13.
func projectWithThreeMembers(t *testing.T) (*Project, *Story) {
	t.Helper()
	project := unProyecto(t)
	for _, name := range []string{"Agustin Salinas", "Juliana Bustos", "Maximiliano Eula"} {
		if _, err := project.AddMember(name, RoleProductBuilder); err != nil {
			t.Fatalf("no se pudo agregar el integrante de prueba %q: %v", name, err)
		}
	}
	story, err := project.AddStory(unosDatos())
	if err != nil {
		t.Fatalf("no se pudo crear la historia de prueba: %v", err)
	}
	return project, story
}

// Esquema del escenario: Abrir una ronda sobre una historia que no esta terminada
// Cubre: RN-1, RN-2, RN-3, RN-13
func TestAbrirUnaRondaSobreUnaHistoriaQueNoEstaTerminada(t *testing.T) {
	cases := []struct {
		name   string
		state  StoryState
		points int
	}{
		{"Pendiente y sin estimar", StoryPending, 0},
		{"Pendiente y estimada en 5", StoryPending, 5},
		{"En curso y estimada en 3", StoryInProgress, 3},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			project, story := projectWithThreeMembers(t)
			story.state = tc.state
			story.storyPoints = tc.points

			round, err := project.OpenVotingRound(story.ID())

			if err != nil {
				t.Fatalf("se esperaba abrir la ronda sin error, salio: %v", err)
			}
			if round.Number() != 1 || round.StoryID() != story.ID() || round.Revealed() {
				t.Fatalf("se esperaba la ronda 1 de la historia %d sin revelar, salio la ronda %d de la historia %d, revelada: %v",
					story.ID(), round.Number(), round.StoryID(), round.Revealed())
			}
			if len(round.Voted()) != 0 {
				t.Fatalf("se esperaba que no votara nadie, salio: %v", round.Voted())
			}
			want := []string{"Agustin Salinas", "Juliana Bustos", "Maximiliano Eula"}
			if !slices.Equal(round.Pending(), want) {
				t.Fatalf("se esperaba que faltaran votar %v, salio: %v", want, round.Pending())
			}
			if story.State() != tc.state || story.StoryPoints() != tc.points {
				t.Fatalf("se esperaba que la historia siguiera %v con %d puntos, salio %v con %d",
					tc.state, tc.points, story.State(), story.StoryPoints())
			}
		})
	}
}

// Esquema del escenario: Rechazar una ronda sobre una historia que no se puede votar
// Cubre: RN-4, RN-5
func TestRechazarUnaRondaSobreUnaHistoriaQueNoSePuedeVotar(t *testing.T) {
	cases := []struct {
		name    string
		storyID int
		done    bool
		want    error
	}{
		{"la historia 1 en estado Terminada", 1, true, ErrStoryEstimateDone},
		{"ninguna historia con identificador 99", 99, false, ErrStoryNotFound},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			project, story := projectWithThreeMembers(t)
			if tc.done {
				story.state = StoryDone
			}

			round, err := project.OpenVotingRound(tc.storyID)

			if !errors.Is(err, tc.want) {
				t.Fatalf("se esperaba %v, salio: %v", tc.want, err)
			}
			if round != nil {
				t.Fatalf("se esperaba que no se abriera ninguna ronda, salio la ronda %d", round.Number())
			}
		})
	}
}

// Escenario: Rechazar una ronda en un proyecto sin integrantes
// Cubre: RN-6
func TestRechazarUnaRondaEnUnProyectoSinIntegrantes(t *testing.T) {
	project := unProyecto(t)
	story, err := project.AddStory(unosDatos())
	if err != nil {
		t.Fatalf("no se pudo crear la historia de prueba: %v", err)
	}

	round, err := project.OpenVotingRound(story.ID())

	if !errors.Is(err, ErrVotingRoundNoMembers) {
		t.Fatalf("se esperaba ErrVotingRoundNoMembers, salio: %v", err)
	}
	if round != nil {
		t.Fatalf("se esperaba que no se abriera ninguna ronda, salio la ronda %d", round.Number())
	}
}

// Escenario: Una historia no puede tener dos rondas en curso
// Cubre: RN-7
func TestUnaHistoriaNoPuedeTenerDosRondasEnCurso(t *testing.T) {
	project, story := projectWithThreeMembers(t)
	other, err := project.AddStory(unosDatos())
	if err != nil {
		t.Fatalf("no se pudo crear la segunda historia de prueba: %v", err)
	}
	if _, err := project.OpenVotingRound(story.ID()); err != nil {
		t.Fatalf("no se pudo abrir la primera ronda: %v", err)
	}

	round, err := project.OpenVotingRound(story.ID())

	if !errors.Is(err, ErrVotingRoundInProgress) {
		t.Fatalf("se esperaba ErrVotingRoundInProgress, salio: %v", err)
	}
	if round != nil {
		t.Fatalf("se esperaba que no se abriera otra ronda, salio la ronda %d", round.Number())
	}
	if _, err := project.OpenVotingRound(other.ID()); err != nil {
		t.Fatalf("se esperaba abrir la ronda de la historia %d sin error, salio: %v", other.ID(), err)
	}
}

// openRound abre una ronda sobre la historia 1 de un proyecto con "Agustin Salinas",
// "Juliana Bustos" y "Maximiliano Eula": el "Dado una ronda sin revelar" de los escenarios.
func openRound(t *testing.T) (*Project, *VotingRound) {
	t.Helper()
	project, story := projectWithThreeMembers(t)
	round, err := project.OpenVotingRound(story.ID())
	if err != nil {
		t.Fatalf("no se pudo abrir la ronda de prueba: %v", err)
	}
	return project, round
}

// Escenario: Votar registra quien voto sin dejar ver los votos, ni siquiera el propio
// Cubre: RN-8, RN-12, RN-13
func TestVotarRegistraQuienVotoSinDejarVerLosVotosNiSiquieraElPropio(t *testing.T) {
	_, round := openRound(t)
	if err := round.Vote("Agustin Salinas", 3); err != nil {
		t.Fatalf("se esperaba aceptar el voto de Agustin Salinas, salio: %v", err)
	}
	if err := round.Vote("Juliana Bustos", 8); err != nil {
		t.Fatalf("se esperaba aceptar el voto de Juliana Bustos, salio: %v", err)
	}

	votes, err := round.Votes()

	if !errors.Is(err, ErrVotesHidden) {
		t.Fatalf("se esperaba ErrVotesHidden, salio: %v", err)
	}
	if votes != nil {
		t.Fatalf("se esperaba no recibir ningun voto, salio: %v", votes)
	}
	if want := []string{"Agustin Salinas", "Juliana Bustos"}; !slices.Equal(round.Voted(), want) {
		t.Fatalf("se esperaba que votaran %v, salio: %v", want, round.Voted())
	}
	if want := []string{"Maximiliano Eula"}; !slices.Equal(round.Pending(), want) {
		t.Fatalf("se esperaba que faltara %v, salio: %v", want, round.Pending())
	}
}

// Esquema del escenario: Revelar cuando votaron todos
// Cubre: RN-14
func TestRevelarCuandoVotaronTodos(t *testing.T) {
	cases := []struct {
		name    string
		members []string
		votes   []Vote
	}{
		{"tres integrantes", []string{"Agustin Salinas", "Juliana Bustos", "Maximiliano Eula"},
			[]Vote{{"Agustin Salinas", 3}, {"Juliana Bustos", 8}, {"Maximiliano Eula", 5}}},
		{"un unico integrante", []string{"Maximiliano Eula"},
			[]Vote{{"Maximiliano Eula", 13}}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			project := unProyecto(t)
			for _, name := range tc.members {
				if _, err := project.AddMember(name, RoleProductBuilder); err != nil {
					t.Fatalf("no se pudo agregar el integrante de prueba %q: %v", name, err)
				}
			}
			story, err := project.AddStory(unosDatos())
			if err != nil {
				t.Fatalf("no se pudo crear la historia de prueba: %v", err)
			}
			round, err := project.OpenVotingRound(story.ID())
			if err != nil {
				t.Fatalf("no se pudo abrir la ronda de prueba: %v", err)
			}
			// Votan en el orden inverso al de los integrantes: los votos revelados tienen que
			// salir en el orden de los integrantes, no en el que se votó.
			for i := len(tc.votes) - 1; i >= 0; i-- {
				if err := round.Vote(tc.votes[i].MemberName, tc.votes[i].Points); err != nil {
					t.Fatalf("no se pudo votar: %v", err)
				}
			}

			err = round.Reveal()

			if err != nil {
				t.Fatalf("se esperaba revelar sin error, salio: %v", err)
			}
			if !round.Revealed() {
				t.Fatal("se esperaba la ronda revelada")
			}
			votes, err := round.Votes()
			if err != nil {
				t.Fatalf("se esperaba ver los votos de la ronda revelada, salio: %v", err)
			}
			if !slices.Equal(votes, tc.votes) {
				t.Fatalf("se esperaban los votos %v, salio: %v", tc.votes, votes)
			}
		})
	}
}

// Esquema del escenario: Rechazar revelar mientras falten votos
// Cubre: RN-15, RN-18
func TestRechazarRevelarMientrasFaltenVotos(t *testing.T) {
	cases := []struct {
		name  string
		voted []string
	}{
		{"nadie", nil},
		{"Agustin Salinas y Juliana Bustos", []string{"Agustin Salinas", "Juliana Bustos"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, round := openRound(t)
			for _, name := range tc.voted {
				if err := round.Vote(name, 5); err != nil {
					t.Fatalf("no se pudo votar: %v", err)
				}
			}

			err := round.Reveal()

			if !errors.Is(err, ErrVotesMissing) {
				t.Fatalf("se esperaba ErrVotesMissing, salio: %v", err)
			}
			if round.Revealed() {
				t.Fatal("se esperaba que la ronda siguiera sin revelar")
			}
			if _, err := round.Votes(); !errors.Is(err, ErrVotesHidden) {
				t.Fatalf("se esperaba que los votos siguieran ocultos, salio: %v", err)
			}
		})
	}
}
