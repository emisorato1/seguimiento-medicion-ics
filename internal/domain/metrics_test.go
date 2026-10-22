package domain

import (
	"errors"
	"testing"
)

// Escenario: Pedir algo que no existe
// Cubre: RN-20
//
// La fila de la historia se adelantó a la tarea 16 para que la carga de horas estimadas
// quedara cubierta entera. Las métricas del proyecto 99 se piden como las pediría la web: el
// proyecto se busca en el registro, y si no existe no se llega a pedir nada.
func TestPedirAlgoQueNoExiste(t *testing.T) {
	cases := []struct {
		name string
		run  func(project *Project) error
		want error
	}{
		{"pedir las metricas del sprint 99",
			func(project *Project) error { _, err := project.SprintMetrics(99); return err }, ErrSprintNotFound},
		{"cargar horas estimadas en la historia 99",
			func(project *Project) error { return project.EstimateHours(99, 4) }, ErrStoryNotFound},
		{"pedir las metricas del proyecto 99",
			func(project *Project) error { _, err := NewProjectRegistry().Get(99); return err }, ErrProjectNotFound},
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

// historiasEstimadas carga en el backlog una historia por cada valor, estimada con esos
// story points; un 0 la deja sin estimar.
func historiasEstimadas(t *testing.T, project *Project, points ...int) {
	t.Helper()
	for _, p := range points {
		story, err := project.AddStory(unosDatos())
		if err != nil {
			t.Fatalf("no se pudo crear la historia de prueba: %v", err)
		}
		if p == 0 {
			continue
		}
		if err := project.EstimateStory(story.ID(), p); err != nil {
			t.Fatalf("no se pudo estimar la historia de prueba en %d: %v", p, err)
		}
	}
}

// Escenario: Metricas de un sprint cerrado
// Cubre: RN-3, RN-16
//
// El escenario se completa en varias tareas (Fase 3): por ahora verifica los planificados, y
// las tareas 6, 9, 10, 11 y 12 le suman su número. Asignar historias es de la #9 y cerrar el
// sprint de la #10, que todavía no están en dev: el test arma el sprint dentro del paquete.
func TestMetricasDeUnSprintCerrado(t *testing.T) {
	project := unProyecto(t)
	historiasEstimadas(t, project, 3, 5, 2, 3)
	sprint := unSprint(t, project)
	sprint.stories = []int{1, 2, 3, 4}
	sprint.state = SprintClosed

	got, err := project.SprintMetrics(1)

	if err != nil {
		t.Fatalf("se esperaba que no fuera un error, salio: %v", err)
	}
	if got.PlannedPoints != 13 || got.StoriesPlanned != 4 {
		t.Fatalf("se esperaban 13 story points y 4 historias planificadas, salieron %d y %d", got.PlannedPoints, got.StoriesPlanned)
	}
}
