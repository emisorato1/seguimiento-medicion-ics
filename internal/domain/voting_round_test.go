package domain

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
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

// Esquema del escenario: Rechazar un voto fuera de la escala
// Cubre: RN-9, RN-18
func TestRechazarUnVotoFueraDeLaEscala(t *testing.T) {
	for _, points := range []int{4, 7, 0, -3} {
		t.Run(strconv.Itoa(points), func(t *testing.T) {
			_, round := openRound(t)

			err := round.Vote("Juliana Bustos", points)

			if !errors.Is(err, ErrStoryPointsInvalid) {
				t.Fatalf("se esperaba ErrStoryPointsInvalid al votar %d, salio: %v", points, err)
			}
			if !slices.Contains(round.Pending(), "Juliana Bustos") {
				t.Fatalf("se esperaba que Juliana Bustos siguiera entre los que faltan, salio: %v", round.Pending())
			}
		})
	}
}

// Escenario: Rechazar un segundo voto del mismo integrante
// Cubre: RN-10, RN-12, RN-18
//
// Los valores 13 y 21 no aparecen en ningún otro dato de la ronda (historia 1, ronda 1,
// nombres sin números): así el test puede comprobar que el aviso no lleva ninguno de los dos.
func TestRechazarUnSegundoVotoDelMismoIntegrante(t *testing.T) {
	cases := []struct {
		name   string
		second int
	}{
		{"con un valor de la escala", 21},
		// El control de "ya votó" va antes que el del valor: un segundo voto fuera de la
		// escala también informa que ya votó (orden de la Fase 2).
		{"con un valor fuera de la escala", 4},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, round := openRound(t)
			if err := round.Vote("Juliana Bustos", 13); err != nil {
				t.Fatalf("no se pudo registrar el primer voto: %v", err)
			}

			err := round.Vote("Juliana Bustos", tc.second)

			if !errors.Is(err, ErrAlreadyVoted) {
				t.Fatalf("se esperaba ErrAlreadyVoted, salio: %v", err)
			}
			if strings.Contains(err.Error(), "13") || strings.Contains(err.Error(), strconv.Itoa(tc.second)) {
				t.Fatalf("se esperaba un aviso sin ningun valor votado, salio: %q", err.Error())
			}
			for _, name := range []string{"Agustin Salinas", "Maximiliano Eula"} {
				if err := round.Vote(name, 5); err != nil {
					t.Fatalf("no se pudo votar: %v", err)
				}
			}
			if err := round.Reveal(); err != nil {
				t.Fatalf("no se pudo revelar la ronda: %v", err)
			}
			votes, _ := round.Votes()
			if votes[1] != (Vote{"Juliana Bustos", 13}) {
				t.Fatalf("se esperaba que quedara el primer voto de Juliana Bustos, 13, salio: %v", votes[1])
			}
		})
	}
}

// Esquema del escenario: Rechazar el voto de quien no participa de la ronda
// Cubre: RN-2, RN-11, RN-18
func TestRechazarElVotoDeQuienNoParticipaDeLaRonda(t *testing.T) {
	cases := []struct {
		name   string
		voter  string
		points int
	}{
		{"Pedro Gomez no es integrante del proyecto", "Pedro Gomez", 5},
		{"Emiliano Sorato se agrego al proyecto despues de abrirla", "Emiliano Sorato", 5},
		// El control de "no participa" va antes que el del valor (orden de la Fase 2).
		{"quien no participa, con un valor fuera de la escala", "Pedro Gomez", 4},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			project, round := openRound(t)
			if _, err := project.AddMember("Emiliano Sorato", RoleAgileEnabler); err != nil {
				t.Fatalf("no se pudo agregar al integrante nuevo: %v", err)
			}

			err := round.Vote(tc.voter, tc.points)

			if !errors.Is(err, ErrNotParticipant) {
				t.Fatalf("se esperaba ErrNotParticipant, salio: %v", err)
			}
			if len(round.Voted()) != 0 {
				t.Fatalf("se esperaba que no votara nadie, salio: %v", round.Voted())
			}
		})
	}
}

// Escenario: Reconocer al integrante aunque el nombre venga escrito distinto
// Cubre: RN-11
func TestReconocerAlIntegranteAunqueElNombreVengaEscritoDistinto(t *testing.T) {
	_, round := openRound(t)

	err := round.Vote("  juliana bustos  ", 5)

	if err != nil {
		t.Fatalf("se esperaba aceptar el voto, salio: %v", err)
	}
	if want := []string{"Juliana Bustos"}; !slices.Equal(round.Voted(), want) {
		t.Fatalf("se esperaba que votara %v, escrito como esta cargado, salio: %v", want, round.Voted())
	}
}

// Escenario: Quien sale del proyecto despues de abrir la ronda sigue participando
// Cubre: RN-2
func TestQuienSaleDelProyectoDespuesDeAbrirLaRondaSigueParticipando(t *testing.T) {
	project, round := openRound(t)
	if err := project.RemoveMember("Juliana Bustos"); err != nil {
		t.Fatalf("no se pudo quitar a Juliana Bustos del proyecto: %v", err)
	}
	for _, name := range []string{"Agustin Salinas", "Maximiliano Eula"} {
		if err := round.Vote(name, 5); err != nil {
			t.Fatalf("no se pudo votar: %v", err)
		}
	}

	err := round.Reveal()

	if want := []string{"Juliana Bustos"}; !slices.Equal(round.Pending(), want) {
		t.Fatalf("se esperaba que siguiera faltando %v, salio: %v", want, round.Pending())
	}
	if !errors.Is(err, ErrVotesMissing) {
		t.Fatalf("se esperaba que la ronda no se pudiera revelar sin su voto, salio: %v", err)
	}
}

// Esquema del escenario: Una ronda revelada ya no acepta cambios
// Cubre: RN-16, RN-18
func TestUnaRondaReveladaYaNoAceptaCambios(t *testing.T) {
	cases := []struct {
		name   string
		action func(round *VotingRound) error
	}{
		// "Ronda revelada" va antes que cualquier otro control: Juliana ya votó, y el aviso
		// igual es que la ronda ya fue revelada (orden de la Fase 2).
		{"votar 3 otra vez como Juliana Bustos", func(round *VotingRound) error { return round.Vote("Juliana Bustos", 3) }},
		{"revelarla de nuevo", func(round *VotingRound) error { return round.Reveal() }},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, round := openRound(t)
			for name, points := range map[string]int{"Agustin Salinas": 5, "Juliana Bustos": 8, "Maximiliano Eula": 5} {
				if err := round.Vote(name, points); err != nil {
					t.Fatalf("no se pudo votar: %v", err)
				}
			}
			if err := round.Reveal(); err != nil {
				t.Fatalf("no se pudo revelar la ronda: %v", err)
			}
			before, _ := round.Votes()

			err := tc.action(round)

			if !errors.Is(err, ErrRoundRevealed) {
				t.Fatalf("se esperaba ErrRoundRevealed, salio: %v", err)
			}
			if after, _ := round.Votes(); !slices.Equal(after, before) {
				t.Fatalf("se esperaba que los votos revelados siguieran siendo %v, salio: %v", before, after)
			}
		})
	}
}

// Escenario: Mostrar una ronda sin revelar no deja ver los votos
// Cubre: RN-17
//
// Se votan 13 y 21, que no aparecen en ningún otro dato de la ronda. Se prueban todas las
// formas de imprimir de Go, también la de depuración (%#v), que no usa el texto de la ronda
// y mostraría los datos ocultos. Y la ronda como valor, no solo como puntero.
func TestMostrarUnaRondaSinRevelarNoDejaVerLosVotos(t *testing.T) {
	_, round := openRound(t)
	if err := round.Vote("Agustin Salinas", 13); err != nil {
		t.Fatalf("no se pudo votar: %v", err)
	}
	if err := round.Vote("Juliana Bustos", 21); err != nil {
		t.Fatalf("no se pudo votar: %v", err)
	}

	for _, format := range []string{"%v", "%+v", "%#v", "%s"} {
		for _, printed := range []string{fmt.Sprintf(format, round), fmt.Sprintf(format, *round)} {
			if strings.Contains(printed, "13") || strings.Contains(printed, "21") {
				t.Fatalf("se esperaba que %s no mostrara ningun valor votado, salio: %q", format, printed)
			}
			for _, want := range []string{"ronda 1", "historia 1", "2 de 3"} {
				if !strings.Contains(printed, want) {
					t.Fatalf("se esperaba que %s mostrara %q, salio: %q", format, want, printed)
				}
			}
		}
	}
}

// Esquema del escenario: Consultar la ronda actual de una historia
// Cubre: RN-19
func TestConsultarLaRondaActualDeUnaHistoria(t *testing.T) {
	cases := []struct {
		name     string
		open     bool
		reveal   bool
		found    bool
		revealed bool
	}{
		{"ninguna ronda", false, false, false, false},
		{"la ronda 1 sin revelar", true, false, true, false},
		{"la ronda 1 revelada", true, true, true, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			project, story := projectWithThreeMembers(t)
			if tc.open {
				round, err := project.OpenVotingRound(story.ID())
				if err != nil {
					t.Fatalf("no se pudo abrir la ronda de prueba: %v", err)
				}
				if tc.reveal {
					for _, name := range round.Pending() {
						if err := round.Vote(name, 5); err != nil {
							t.Fatalf("no se pudo votar: %v", err)
						}
					}
					if err := round.Reveal(); err != nil {
						t.Fatalf("no se pudo revelar la ronda: %v", err)
					}
				}
			}

			round, found, err := project.CurrentVotingRound(story.ID())

			if err != nil {
				t.Fatalf("se esperaba consultar sin error, salio: %v", err)
			}
			if found != tc.found {
				t.Fatalf("se esperaba encontrada: %v, salio: %v", tc.found, found)
			}
			if tc.found && (round.Number() != 1 || round.Revealed() != tc.revealed) {
				t.Fatalf("se esperaba la ronda 1 con revelada: %v, salio la ronda %d con revelada: %v",
					tc.revealed, round.Number(), round.Revealed())
			}
		})
	}
}

// Escenario: Consultar la ronda actual de una historia que no existe
// Cubre: RN-19
func TestConsultarLaRondaActualDeUnaHistoriaQueNoExiste(t *testing.T) {
	project, _ := projectWithThreeMembers(t)

	_, found, err := project.CurrentVotingRound(99)

	if !errors.Is(err, ErrStoryNotFound) {
		t.Fatalf("se esperaba ErrStoryNotFound, salio: %v", err)
	}
	if found {
		t.Fatal("se esperaba que no se encontrara ninguna ronda")
	}
}
