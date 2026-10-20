package domain

import (
	"math"
	"testing"
)

// mismasHoras compara dos cantidades con decimales con una tolerancia chica, como pide la
// Fase 3 de la #16: comparar decimales con == falla por el redondeo de la computadora.
func mismasHoras(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}

// Escenario: Cargar y reemplazar las horas estimadas de una historia
// Cubre: RN-17
func TestCargarYReemplazarLasHorasEstimadasDeUnaHistoria(t *testing.T) {
	project := unProyecto(t)
	story, err := project.AddStory(unosDatos())
	if err != nil {
		t.Fatalf("no se pudo crear la historia de prueba: %v", err)
	}

	if story.EstimatedHours() != 0 {
		t.Fatalf("se esperaba que la historia naciera con 0 horas estimadas, salio %v", story.EstimatedHours())
	}

	if err := project.EstimateHours(story.ID(), 6.5); err != nil {
		t.Fatalf("se esperaba cargar 6,5 horas estimadas sin error, salio: %v", err)
	}
	if !mismasHoras(story.EstimatedHours(), 6.5) {
		t.Fatalf("se esperaba la historia con 6,5 horas estimadas, salio %v", story.EstimatedHours())
	}

	if err := project.EstimateHours(story.ID(), 8); err != nil {
		t.Fatalf("se esperaba reemplazar por 8 horas estimadas sin error, salio: %v", err)
	}
	if !mismasHoras(story.EstimatedHours(), 8) {
		t.Fatalf("se esperaba la historia con 8 horas estimadas, salio %v", story.EstimatedHours())
	}
}
