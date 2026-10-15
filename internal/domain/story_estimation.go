package domain

import (
	"strconv"
	"strings"
)

// storyPointScale es la escala válida de story points (RN-1). Es la única fuente: la usan
// la validación y el mensaje de error, así el mensaje nunca dice algo distinto (Fase 2, D3).
var storyPointScale = []int{1, 2, 3, 5, 8, 13, 21}

// Estimate asigna story points a la historia. Sirve tanto para la primera estimación
// como para reestimar (RN-5, RN-6). Una historia Terminada no se estima (RN-8). Controla
// el estado antes que el valor (Fase 2, D5) y valida antes de cambiar: si algo falla, la
// historia queda como estaba (RN-2, RN-3).
func (s *Story) Estimate(points int) error {
	if s.state == StoryDone {
		return ErrStoryEstimateDone
	}
	if !inScale(points) {
		return ErrStoryPointsInvalid
	}
	s.storyPoints = points
	return nil
}

// inScale indica si el valor está en la escala de story points.
func inScale(points int) bool {
	for _, valid := range storyPointScale {
		if points == valid {
			return true
		}
	}
	return false
}

// scaleText escribe la escala como la lee el usuario: "1, 2, 3, 5, 8, 13 y 21".
func scaleText() string {
	texts := make([]string, len(storyPointScale))
	for i, points := range storyPointScale {
		texts[i] = strconv.Itoa(points)
	}
	last := len(texts) - 1
	return strings.Join(texts[:last], ", ") + " y " + texts[last]
}

// RemoveEstimate quita la estimación de la historia, que vuelve a quedar sin estimar
// (RN-7). Es una operación propia y no "estimar con 0", porque el 0 no es un valor válido
// para estimar (RN-3, Fase 2). A una historia Terminada no se le quita (RN-8).
func (s *Story) RemoveEstimate() error {
	if s.state == StoryDone {
		return ErrStoryEstimateDone
	}
	s.storyPoints = 0
	return nil
}
