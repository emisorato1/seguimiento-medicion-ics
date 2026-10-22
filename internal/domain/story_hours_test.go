package domain

import (
	"errors"
	"fmt"
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

// Escenario: Rechazar horas estimadas que no son mayores que cero
// Cubre: RN-18
func TestRechazarHorasEstimadasQueNoSonMayoresQueCero(t *testing.T) {
	cases := []struct {
		name  string
		hours float64
	}{
		{"0", 0},
		{"-2", -2},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			project := unProyecto(t)
			story, err := project.AddStory(unosDatos())
			if err != nil {
				t.Fatalf("no se pudo crear la historia de prueba: %v", err)
			}
			if err := project.EstimateHours(story.ID(), 4); err != nil {
				t.Fatalf("no se pudieron cargar las 4 horas estimadas de prueba: %v", err)
			}

			err = project.EstimateHours(story.ID(), tc.hours)

			if !errors.Is(err, ErrEstimatedHoursInvalid) {
				t.Fatalf("se esperaba ErrEstimatedHoursInvalid al cargar %v horas, salio: %v", tc.hours, err)
			}
			if !mismasHoras(story.EstimatedHours(), 4) {
				t.Fatalf("se esperaba que la historia siguiera con 4 horas estimadas, salio %v", story.EstimatedHours())
			}
		})
	}
}

// Escenario: No cambiar las horas estimadas de una historia terminada
// Cubre: RN-19
func TestNoCambiarLasHorasEstimadasDeUnaHistoriaTerminada(t *testing.T) {
	project := unProyecto(t)
	story, err := project.AddStory(unosDatos())
	if err != nil {
		t.Fatalf("no se pudo crear la historia de prueba: %v", err)
	}
	// Terminar una historia es de la #10, que todavia no esta en dev: el test la prepara
	// desde dentro del paquete, como la #12 en su tarea 9.
	story.state = StoryDone
	story.estimatedHours = 4

	err = project.EstimateHours(story.ID(), 6)

	if !errors.Is(err, ErrStoryEstimateDone) {
		t.Fatalf("se esperaba ErrStoryEstimateDone, salio: %v", err)
	}
	if !mismasHoras(story.EstimatedHours(), 4) {
		t.Fatalf("se esperaba que la historia siguiera con 4 horas estimadas, salio %v", story.EstimatedHours())
	}
}

// Sin escenario: caso límite de la RN-18, hallazgo de la revisión del PR #81
// Cubre: RN-18
//
// NaN ("no es un número") no es mayor que cero, así que la regla lo rechaza. Puede llegar
// desde la consola, porque strconv.ParseFloat acepta el texto "NaN".
func TestRechazarHorasEstimadasQueNoSonUnNumero(t *testing.T) {
	project := unProyecto(t)
	story, err := project.AddStory(unosDatos())
	if err != nil {
		t.Fatalf("no se pudo crear la historia de prueba: %v", err)
	}
	if err := project.EstimateHours(story.ID(), 4); err != nil {
		t.Fatalf("no se pudieron cargar las 4 horas estimadas de prueba: %v", err)
	}

	err = project.EstimateHours(story.ID(), math.NaN())

	if !errors.Is(err, ErrEstimatedHoursInvalid) {
		t.Fatalf("se esperaba ErrEstimatedHoursInvalid al cargar NaN horas, salio: %v", err)
	}
	if !mismasHoras(story.EstimatedHours(), 4) {
		t.Fatalf("se esperaba que la historia siguiera con 4 horas estimadas, salio %v", story.EstimatedHours())
	}
}

// Sin escenario: caso límite de la RN-18, pedido de cambios de Agustín en el PR #81 (20/10)
// Cubre: RN-18
//
// El infinito positivo es mayor que cero, pero no son horas: strconv.ParseFloat acepta el
// texto "Inf", y con una historia en horas infinitas la web ya no puede guardar el archivo,
// porque JSON no tiene cómo escribirlas. El negativo ya lo rechazaba la regla; va igual para
// que los dos queden cubiertos.
func TestRechazarHorasEstimadasInfinitas(t *testing.T) {
	for _, hours := range []float64{math.Inf(1), math.Inf(-1)} {
		t.Run(fmt.Sprint(hours), func(t *testing.T) {
			project := unProyecto(t)
			story, err := project.AddStory(unosDatos())
			if err != nil {
				t.Fatalf("no se pudo crear la historia de prueba: %v", err)
			}
			if err := project.EstimateHours(story.ID(), 4); err != nil {
				t.Fatalf("no se pudieron cargar las 4 horas estimadas de prueba: %v", err)
			}

			err = project.EstimateHours(story.ID(), hours)

			if !errors.Is(err, ErrEstimatedHoursInvalid) {
				t.Fatalf("se esperaba ErrEstimatedHoursInvalid al cargar %v horas, salio: %v", hours, err)
			}
			if !mismasHoras(story.EstimatedHours(), 4) {
				t.Fatalf("se esperaba que la historia siguiera con 4 horas estimadas, salio %v", story.EstimatedHours())
			}
		})
	}
}
