package domain

import (
	"errors"
	"slices"
	"testing"
)

// revealedRound abre una ronda sobre la historia 1 de un proyecto con los integrantes de
// los votos, en ese orden, la vota y la revela: el "Dado una ronda revelada donde ..." de
// los escenarios de la #73.
func revealedRound(t *testing.T, votes []Vote) (*Project, *VotingRound) {
	t.Helper()
	project := unProyecto(t)
	for _, vote := range votes {
		if _, err := project.AddMember(vote.MemberName, RoleProductBuilder); err != nil {
			t.Fatalf("no se pudo agregar el integrante de prueba %q: %v", vote.MemberName, err)
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
	for _, vote := range votes {
		if err := round.Vote(vote.MemberName, vote.Points); err != nil {
			t.Fatalf("no se pudo votar: %v", err)
		}
	}
	if err := round.Reveal(); err != nil {
		t.Fatalf("no se pudo revelar la ronda de prueba: %v", err)
	}
	return project, round
}

// sameExtreme indica si un extremo de las diferencias tiene ese valor y esos integrantes,
// en ese orden.
func sameExtreme(got VoteExtreme, points int, members ...string) bool {
	return got.Points == points && slices.Equal(got.Members, members)
}

// Esquema del escenario: Informar las diferencias de una ronda revelada
// Cubre: RN-1, RN-2
func TestInformarLasDiferenciasDeUnaRondaRevelada(t *testing.T) {
	cases := []struct {
		name    string
		votes   []Vote
		lowest  []string
		highest []string
	}{
		{"3, 8 y 5", []Vote{{"Agustin Salinas", 3}, {"Juliana Bustos", 8}, {"Maximiliano Eula", 5}},
			[]string{"Agustin Salinas"}, []string{"Juliana Bustos"}},
		{"3, 8 y 8: empate en el mas alto", []Vote{{"Agustin Salinas", 3}, {"Juliana Bustos", 8}, {"Maximiliano Eula", 8}},
			[]string{"Agustin Salinas"}, []string{"Juliana Bustos", "Maximiliano Eula"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, round := revealedRound(t, tc.votes)

			differences, err := round.Differences()

			if err != nil {
				t.Fatalf("se esperaban las diferencias sin error, salio: %v", err)
			}
			if !differences.HasDifferences {
				t.Fatal("se esperaba que hubiera diferencias")
			}
			if !sameExtreme(differences.Lowest, 3, tc.lowest...) {
				t.Fatalf("se esperaba el mas bajo en 3, de %v, salio: %+v", tc.lowest, differences.Lowest)
			}
			if !sameExtreme(differences.Highest, 8, tc.highest...) {
				t.Fatalf("se esperaba el mas alto en 8, de %v, salio: %+v", tc.highest, differences.Highest)
			}
		})
	}
}

// Esquema del escenario: Sin diferencias cuando todos votaron lo mismo
// Cubre: RN-3
func TestSinDiferenciasCuandoTodosVotaronLoMismo(t *testing.T) {
	cases := []struct {
		name   string
		votes  []Vote
		points int
	}{
		{"los tres integrantes votaron 5", []Vote{{"Agustin Salinas", 5}, {"Juliana Bustos", 5}, {"Maximiliano Eula", 5}}, 5},
		{"Maximiliano Eula, el unico integrante, voto 13", []Vote{{"Maximiliano Eula", 13}}, 13},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, round := revealedRound(t, tc.votes)

			differences, err := round.Differences()

			if err != nil {
				t.Fatalf("se esperaban las diferencias sin error, salio: %v", err)
			}
			if differences.HasDifferences {
				t.Fatal("se esperaba que no hubiera diferencias")
			}
			members := make([]string, 0, len(tc.votes))
			for _, vote := range tc.votes {
				members = append(members, vote.MemberName)
			}
			if !sameExtreme(differences.Lowest, tc.points, members...) || !sameExtreme(differences.Highest, tc.points, members...) {
				t.Fatalf("se esperaba el mas bajo y el mas alto en %d, de %v, salio: %+v y %+v",
					tc.points, members, differences.Lowest, differences.Highest)
			}
		})
	}
}

// Escenario: No se informan diferencias de una ronda sin revelar
// Cubre: RN-4
func TestNoSeInformanDiferenciasDeUnaRondaSinRevelar(t *testing.T) {
	_, round := openRound(t)
	for _, name := range round.Pending() {
		if err := round.Vote(name, 5); err != nil {
			t.Fatalf("no se pudo votar: %v", err)
		}
	}

	differences, err := round.Differences()

	if !errors.Is(err, ErrVotesHidden) {
		t.Fatalf("se esperaba ErrVotesHidden, salio: %v", err)
	}
	if differences.HasDifferences || differences.Lowest.Members != nil || differences.Highest.Members != nil {
		t.Fatalf("se esperaba no informar nada de los votos, salio: %+v", differences)
	}
}
