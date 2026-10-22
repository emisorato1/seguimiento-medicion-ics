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
