package domain

import (
	"errors"
	"testing"
)

// Escenario: Pedir algo que no existe
// Cubre: RN-20
//
// Por ahora solo tiene la fila de la historia, adelantada de la tarea 16 para que la carga
// de horas estimadas quede cubierta entera. Las filas del sprint y del proyecto se agregan
// en la tarea 16, cuando existan las métricas.
func TestPedirAlgoQueNoExiste(t *testing.T) {
	cases := []struct {
		name string
		run  func(project *Project) error
		want error
	}{
		{"cargar horas estimadas en la historia 99",
			func(project *Project) error { return project.EstimateHours(99, 4) }, ErrStoryNotFound},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			project := unProyecto(t)

			err := tc.run(project)

			if !errors.Is(err, tc.want) {
				t.Fatalf("se esperaba %v, salio: %v", tc.want, err)
			}
		})
	}
}
