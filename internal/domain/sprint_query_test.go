package domain

import (
	"slices"
	"testing"
	"time"
)

// unSprintCerrado crea en el proyecto el sprint siguiente con ese goal y esas fechas, le
// asigna una historia nueva por cada valor de puntos y lo cierra en closedOn con las
// historias completadas indicadas. Crear es de la #8, asignar de la #9 y cerrar de la #10.
func unSprintCerrado(t *testing.T, project *Project, goal string, start, end, closedOn time.Time, completed []int, points ...int) *Sprint {
	t.Helper()
	sprint, err := project.AddSprint(SprintData{Goal: goal, Start: start, End: end})
	if err != nil {
		t.Fatalf("no se pudo crear el sprint de prueba %q: %v", goal, err)
	}
	for _, p := range points {
		story := unaHistoriaEstimada(t, project, p)
		if err := project.AssignStory(story.ID()); err != nil {
			t.Fatalf("no se pudo asignar la historia de prueba %d: %v", story.ID(), err)
		}
	}
	if err := project.CloseSprint(completed, closedOn); err != nil {
		t.Fatalf("no se pudo cerrar el sprint de prueba %q: %v", goal, err)
	}
	return sprint
}

// Esquema del escenario: Un proyecto sin sprints cerrados devuelve un listado vacio
// Cubre: RN-3
func TestUnProyectoSinSprintsCerradosDevuelveUnListadoVacio(t *testing.T) {
	cases := []struct {
		situacion string
		prepare   func(t *testing.T, project *Project)
	}{
		{"sin ningun sprint", func(t *testing.T, project *Project) {}},
		{"con un sprint abierto", func(t *testing.T, project *Project) { unSprint(t, project) }},
	}
	for _, c := range cases {
		t.Run(c.situacion, func(t *testing.T) {
			project := unProyecto(t)
			c.prepare(t, project)

			sprints := project.ClosedSprints()

			if sprints == nil {
				t.Fatalf("se esperaba una lista vacia, no nil")
			}
			if len(sprints) != 0 {
				t.Fatalf("se esperaba el listado vacio, salieron %d sprints", len(sprints))
			}
		})
	}
}

// Escenario: Listar los sprints cerrados de un proyecto
// Cubre: RN-1, RN-4
func TestListarLosSprintsCerradosDeUnProyecto(t *testing.T) {
	project := unProyecto(t)
	unSprintCerrado(t, project, "Tener el MVP", unDia(2026, 10, 12), unDia(2026, 10, 19), unDia(2026, 10, 19), nil, 3, 5)
	unSprintCerrado(t, project, "Tener los sprints", unDia(2026, 10, 20), unDia(2026, 10, 25), unDia(2026, 10, 24), nil, 2)

	sprints := project.ClosedSprints()

	if len(sprints) != 2 {
		t.Fatalf("se esperaban los sprints 1 y 2, salieron %d sprints", len(sprints))
	}
	want := []struct {
		number          int
		goal            string
		start, end, cls time.Time
		stories         []int
	}{
		{1, "Tener el MVP", unDia(2026, 10, 12), unDia(2026, 10, 19), unDia(2026, 10, 19), []int{1, 2}},
		{2, "Tener los sprints", unDia(2026, 10, 20), unDia(2026, 10, 25), unDia(2026, 10, 24), []int{3}},
	}
	for i, w := range want {
		s := sprints[i]
		if s.Number() != w.number || s.Goal() != w.goal {
			t.Errorf("en el lugar %d se esperaba el sprint %d %q, salio el %d %q", i+1, w.number, w.goal, s.Number(), s.Goal())
		}
		if !s.Start().Equal(w.start) || !s.End().Equal(w.end) || !s.ClosedOn().Equal(w.cls) {
			t.Errorf("el sprint %d tiene las fechas %v, %v y cierre %v; se esperaban %v, %v y %v", w.number, s.Start(), s.End(), s.ClosedOn(), w.start, w.end, w.cls)
		}
		if !slices.Equal(s.Stories(), w.stories) {
			t.Errorf("el sprint %d tiene las historias %v, se esperaban %v", w.number, s.Stories(), w.stories)
		}
	}
}

// Escenario: El sprint abierto no aparece en el listado
// Cubre: RN-2
func TestElSprintAbiertoNoApareceEnElListado(t *testing.T) {
	project := unProyecto(t)
	unSprintCerrado(t, project, "Tener el MVP", unDia(2026, 10, 12), unDia(2026, 10, 19), unDia(2026, 10, 19), nil)
	unSprint(t, project)

	sprints := project.ClosedSprints()

	if len(sprints) != 1 || sprints[0].Number() != 1 {
		numbers := []int{}
		for _, s := range sprints {
			numbers = append(numbers, s.Number())
		}
		t.Fatalf("se esperaba solamente el sprint 1, salieron %v", numbers)
	}
}
