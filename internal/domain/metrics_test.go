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

// unSprint crea el sprint siguiente del proyecto, abierto y sin historias.
func unSprint(t *testing.T, project *Project) *Sprint {
	t.Helper()
	sprint, err := project.AddSprint(SprintData{Goal: "Tener el MVP", Start: unDia(2026, 10, 26), End: unDia(2026, 10, 30)})
	if err != nil {
		t.Fatalf("no se pudo crear el sprint de prueba: %v", err)
	}
	return sprint
}

// Esquema del escenario: Sin datos cargados las metricas valen cero
// Cubre: RN-1, RN-14, RN-15
func TestSinDatosCargadosLasMetricasValenCero(t *testing.T) {
	cases := []struct {
		name string
		ask  func(t *testing.T, project *Project) (Metrics, error)
	}{
		{"un proyecto con el sprint 1 recien abierto", func(t *testing.T, project *Project) (Metrics, error) {
			unSprint(t, project)
			return project.SprintMetrics(1)
		}},
		{"un proyecto sin sprints", func(t *testing.T, project *Project) (Metrics, error) {
			return project.ProjectMetrics(), nil
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			project := unProyecto(t)

			got, err := tc.ask(t, project)

			if err != nil {
				t.Fatalf("se esperaba que no fuera un error, salio: %v", err)
			}
			if got != (Metrics{}) {
				t.Fatalf("se esperaban las once metricas en cero, salio %+v", got)
			}
		})
	}
}
