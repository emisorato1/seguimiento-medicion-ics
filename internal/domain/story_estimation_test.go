package domain

import (
	"errors"
	"strconv"
	"strings"
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

// Escenario: Rechazar valores fuera de la escala
// Cubre: RN-1, RN-2, RN-3
func TestRechazarValoresFueraDeLaEscala(t *testing.T) {
	for _, points := range []int{4, 7, 0, -3} {
		t.Run(strconv.Itoa(points), func(t *testing.T) {
			story := &Story{state: StoryPending}

			err := story.Estimate(points)

			if !errors.Is(err, ErrStoryPointsInvalid) {
				t.Fatalf("se esperaba ErrStoryPointsInvalid al estimar con %d, salio: %v", points, err)
			}
			informed := numbersIn(err.Error())
			for _, valid := range []int{1, 2, 3, 5, 8, 13, 21} {
				if !informed[valid] {
					t.Fatalf("se esperaba que el mensaje informara el valor valido %d, salio: %q", valid, err.Error())
				}
			}
			if story.IsEstimated() {
				t.Fatalf("se esperaba que la historia siguiera sin estimar, salio estimada en %d", story.StoryPoints())
			}
		})
	}
}

// numbersIn devuelve los números enteros que aparecen en un texto.
func numbersIn(text string) map[int]bool {
	numbers := map[int]bool{}
	for _, field := range strings.FieldsFunc(text, func(r rune) bool { return r < '0' || r > '9' }) {
		if n, err := strconv.Atoi(field); err == nil {
			numbers[n] = true
		}
	}
	return numbers
}

// Escenario: Reestimar una historia
// Cubre: RN-5, RN-6
func TestReestimarUnaHistoria(t *testing.T) {
	story := &Story{state: StoryPending, storyPoints: 3}

	err := story.Estimate(8)

	if err != nil {
		t.Fatalf("se esperaba reestimar sin error, salio: %v", err)
	}
	if story.StoryPoints() != 8 {
		t.Fatalf("se esperaba la historia estimada en 8, salio %d", story.StoryPoints())
	}
}

// Escenario: Reestimar con el mismo valor
// Cubre: RN-5
func TestReestimarConElMismoValor(t *testing.T) {
	story := &Story{state: StoryPending, storyPoints: 5}

	err := story.Estimate(5)

	if err != nil {
		t.Fatalf("se esperaba reestimar con el mismo valor sin error, salio: %v", err)
	}
	if story.StoryPoints() != 5 {
		t.Fatalf("se esperaba la historia estimada en 5, salio %d", story.StoryPoints())
	}
}

// Escenario: Reestimar una historia en curso
// Cubre: RN-5
func TestReestimarUnaHistoriaEnCurso(t *testing.T) {
	story := &Story{state: StoryInProgress, storyPoints: 3}

	err := story.Estimate(5)

	if err != nil {
		t.Fatalf("se esperaba reestimar una historia En curso sin error, salio: %v", err)
	}
	if story.StoryPoints() != 5 {
		t.Fatalf("se esperaba la historia estimada en 5, salio %d", story.StoryPoints())
	}
}

// Escenario: Quitar la estimacion
// Cubre: RN-7
func TestQuitarLaEstimacion(t *testing.T) {
	story := &Story{state: StoryPending, storyPoints: 5}

	err := story.RemoveEstimate()

	if err != nil {
		t.Fatalf("se esperaba quitar la estimacion sin error, salio: %v", err)
	}
	if story.IsEstimated() {
		t.Fatalf("se esperaba que la historia figurara sin estimar, salio estimada en %d", story.StoryPoints())
	}
}

// Sin escenario: caso límite de la RN-7
// Cubre: RN-7
func TestQuitarLaEstimacionDeUnaHistoriaYaSinEstimar(t *testing.T) {
	story := &Story{state: StoryPending}

	err := story.RemoveEstimate()

	if err != nil {
		t.Fatalf("se esperaba que quitar la estimacion de una historia sin estimar no diera error, salio: %v", err)
	}
	if story.IsEstimated() {
		t.Fatalf("se esperaba que la historia siguiera sin estimar, salio estimada en %d", story.StoryPoints())
	}
}

// Escenario: No reestimar una historia terminada
// Cubre: RN-8
func TestNoReestimarUnaHistoriaTerminada(t *testing.T) {
	cases := []struct {
		name    string
		initial int
	}{
		{"reestimar una terminada estimada en 5", 5},
		{"estimar una terminada sin estimar", 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			story := &Story{state: StoryDone, storyPoints: tc.initial}

			err := story.Estimate(8)

			if !errors.Is(err, ErrStoryEstimateDone) {
				t.Fatalf("se esperaba ErrStoryEstimateDone, salio: %v", err)
			}
			if story.StoryPoints() != tc.initial {
				t.Fatalf("se esperaba que la historia conservara %d story points, salio %d", tc.initial, story.StoryPoints())
			}
		})
	}
}

// Escenario: No quitar la estimacion de una historia terminada
// Cubre: RN-8
func TestNoQuitarLaEstimacionDeUnaHistoriaTerminada(t *testing.T) {
	story := &Story{state: StoryDone, storyPoints: 5}

	err := story.RemoveEstimate()

	if !errors.Is(err, ErrStoryEstimateDone) {
		t.Fatalf("se esperaba ErrStoryEstimateDone, salio: %v", err)
	}
	if story.StoryPoints() != 5 {
		t.Fatalf("se esperaba que la historia siguiera estimada en 5, salio %d", story.StoryPoints())
	}
}
