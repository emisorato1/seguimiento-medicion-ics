package domain

import (
	"strconv"
	"testing"
)

// Escenario: Consultar una historia sin estimar
// Cubre: RN-4
//
// Desvío: este test pasa sin código nuevo, porque IsEstimated ya estaba en dev desde el
// PR #56 de la historia #6. Queda anotado en la Fase 4.
func TestConsultarUnaHistoriaSinEstimar(t *testing.T) {
	story := &Story{state: StoryPending}

	if story.IsEstimated() {
		t.Fatalf("se esperaba que la historia figurara sin estimar, salio estimada con %d story points", story.StoryPoints())
	}
}

// Escenario: Estimar con cada valor de la escala
// Cubre: RN-1, RN-5
func TestEstimarConCadaValorDeLaEscala(t *testing.T) {
	for _, points := range []int{1, 2, 3, 5, 8, 13, 21} {
		t.Run(strconv.Itoa(points), func(t *testing.T) {
			story := &Story{state: StoryPending}

			err := story.Estimate(points)

			if err != nil {
				t.Fatalf("se esperaba estimar con %d sin error, salio: %v", points, err)
			}
			if story.StoryPoints() != points || !story.IsEstimated() {
				t.Fatalf("se esperaba la historia estimada en %d, salio %d", points, story.StoryPoints())
			}
		})
	}
}
