package domain

import (
	"errors"
	"fmt"
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

// Escenario: Pedir el detalle de un sprint por su numero
// Cubre: RN-5
func TestPedirElDetalleDeUnSprintPorSuNumero(t *testing.T) {
	project := unProyecto(t)
	unSprintCerrado(t, project, "Tener el MVP", unDia(2026, 10, 12), unDia(2026, 10, 19), unDia(2026, 10, 19), nil)
	unSprintCerrado(t, project, "Tener los sprints", unDia(2026, 10, 20), unDia(2026, 10, 25), unDia(2026, 10, 24), nil)

	sprint, err := project.Sprint(2)

	if err != nil {
		t.Fatalf("se esperaba encontrar el sprint 2, salio: %v", err)
	}
	if sprint.Number() != 2 || sprint.Goal() != "Tener los sprints" {
		t.Fatalf("se esperaba el sprint 2 con su goal, salio el %d con %q", sprint.Number(), sprint.Goal())
	}
	if !sprint.Start().Equal(unDia(2026, 10, 20)) || !sprint.End().Equal(unDia(2026, 10, 25)) || !sprint.ClosedOn().Equal(unDia(2026, 10, 24)) {
		t.Errorf("se esperaban las fechas 20/10/2026 y 25/10/2026 y el cierre el 24/10/2026, salio %v, %v y %v", sprint.Start(), sprint.End(), sprint.ClosedOn())
	}
}

// Escenario: Ver que historias se completaron en un sprint cerrado
// Cubre: RN-4
func TestVerQueHistoriasSeCompletaronEnUnSprintCerrado(t *testing.T) {
	project := unProyecto(t)
	unSprintCerrado(t, project, "Tener el MVP", unDia(2026, 10, 12), unDia(2026, 10, 19), unDia(2026, 10, 19), []int{1}, 3, 5, 2)

	sprint, err := project.Sprint(1)

	if err != nil {
		t.Fatalf("se esperaba encontrar el sprint 1, salio: %v", err)
	}
	if !slices.Equal(sprint.Stories(), []int{1, 2, 3}) {
		t.Errorf("se esperaban asignadas las historias [1 2 3], salieron %v", sprint.Stories())
	}
	if !slices.Equal(sprint.CompletedStories(), []int{1}) {
		t.Errorf("se esperaba completada la historia [1], salieron %v", sprint.CompletedStories())
	}
}

// Escenario: Un sprint cerrado sin historias se consulta igual
// Cubre: RN-4
func TestUnSprintCerradoSinHistoriasSeConsultaIgual(t *testing.T) {
	project := unProyecto(t)
	unSprintCerrado(t, project, "Tener el MVP", unDia(2026, 10, 12), unDia(2026, 10, 19), unDia(2026, 10, 19), nil)

	sprint, err := project.Sprint(1)

	if err != nil {
		t.Fatalf("se esperaba encontrar el sprint 1, salio: %v", err)
	}
	if sprint.Goal() != "Tener el MVP" || !sprint.Start().Equal(unDia(2026, 10, 12)) || !sprint.End().Equal(unDia(2026, 10, 19)) {
		t.Errorf("se esperaba el sprint 1 con su goal y sus fechas, salio %q del %v al %v", sprint.Goal(), sprint.Start(), sprint.End())
	}
	if stories := sprint.Stories(); stories == nil || len(stories) != 0 {
		t.Errorf("se esperaba la lista de asignadas vacia y no nil, salio %#v", stories)
	}
	if completed := sprint.CompletedStories(); completed == nil || len(completed) != 0 {
		t.Errorf("se esperaba la lista de completadas vacia y no nil, salio %#v", completed)
	}
}

// Esquema del escenario: Pedir un sprint que no existe
// Cubre: RN-6
func TestPedirUnSprintQueNoExiste(t *testing.T) {
	for _, number := range []int{0, 2, -1} {
		t.Run(fmt.Sprintf("el sprint %d", number), func(t *testing.T) {
			project := unProyecto(t)
			unSprintCerrado(t, project, "Tener el MVP", unDia(2026, 10, 12), unDia(2026, 10, 19), unDia(2026, 10, 19), nil)

			sprint, err := project.Sprint(number)

			if !errors.Is(err, ErrSprintNotFound) {
				t.Fatalf("se esperaba ErrSprintNotFound para el sprint %d, salio: %v", number, err)
			}
			if sprint != nil {
				t.Fatalf("se esperaba no recibir ningun sprint, salio el %d", sprint.Number())
			}
		})
	}
}

// Escenario: Pedir el detalle del sprint abierto
// Cubre: RN-7
func TestPedirElDetalleDelSprintAbierto(t *testing.T) {
	project := unProyecto(t)
	unSprint(t, project)

	sprint, err := project.Sprint(1)

	if err != nil {
		t.Fatalf("se esperaba encontrar el sprint 1, salio: %v", err)
	}
	if sprint.State() != SprintOpen {
		t.Errorf("se esperaba el sprint 1 en estado Abierto, salio %v", sprint.State())
	}
	if !sprint.ClosedOn().IsZero() {
		t.Errorf("se esperaba la fecha de cierre vacia, salio %v", sprint.ClosedOn())
	}
}

// Escenario: Consultar un sprint no lo modifica
// Cubre: RN-8
func TestConsultarUnSprintNoLoModifica(t *testing.T) {
	project := unProyecto(t)
	unSprintCerrado(t, project, "Tener el MVP", unDia(2026, 10, 12), unDia(2026, 10, 19), unDia(2026, 10, 19), []int{1}, 3, 5)

	sprint, err := project.Sprint(1)
	if err != nil {
		t.Fatalf("se esperaba encontrar el sprint 1, salio: %v", err)
	}
	stories := sprint.Stories()
	stories[0] = 99
	completed := sprint.CompletedStories()
	completed[0] = 99

	again, _ := project.Sprint(1)
	if !slices.Equal(again.Stories(), []int{1, 2}) {
		t.Errorf("se esperaba que el sprint 1 siguiera con sus historias [1 2] asignadas, tiene %v", again.Stories())
	}
	if !slices.Equal(again.CompletedStories(), []int{1}) {
		t.Errorf("se esperaba que el sprint 1 siguiera registrando completada la historia [1], tiene %v", again.CompletedStories())
	}
}

// Escenario: Cada proyecto lista solamente sus sprints
// Cubre: RN-9
func TestCadaProyectoListaSolamenteSusSprints(t *testing.T) {
	registry := NewProjectRegistry()
	a, err := registry.Create(ProjectData{Name: "Proyecto A", Start: unDia(2026, 10, 1)})
	if err != nil {
		t.Fatalf("no se pudo crear el proyecto A: %v", err)
	}
	b, err := registry.Create(ProjectData{Name: "Proyecto B", Start: unDia(2026, 10, 1)})
	if err != nil {
		t.Fatalf("no se pudo crear el proyecto B: %v", err)
	}
	unSprintCerrado(t, a, "Meta de A", unDia(2026, 10, 12), unDia(2026, 10, 19), unDia(2026, 10, 19), nil)
	unSprintCerrado(t, b, "Meta de B", unDia(2026, 10, 12), unDia(2026, 10, 19), unDia(2026, 10, 19), nil)

	sprints := a.ClosedSprints()

	if len(sprints) != 1 {
		t.Fatalf("se esperaba un solo sprint, salieron %d", len(sprints))
	}
	if sprints[0].Number() != 1 || sprints[0].Goal() != "Meta de A" {
		t.Errorf("se esperaba el sprint 1 del proyecto A, salio el %d con %q", sprints[0].Number(), sprints[0].Goal())
	}
}

// Escenario: Cerrar un sprint nuevo no cambia los anteriores
// Cubre: RN-11
func TestCerrarUnSprintNuevoNoCambiaLosAnteriores(t *testing.T) {
	project := unProyecto(t)
	first := unSprintCerrado(t, project, "Tener el MVP", unDia(2026, 10, 12), unDia(2026, 10, 19), unDia(2026, 10, 19), []int{1}, 3)
	if _, err := project.AddSprint(SprintData{Goal: "Tener los sprints", Start: unDia(2026, 10, 20), End: unDia(2026, 10, 25)}); err != nil {
		t.Fatalf("no se pudo crear el sprint 2 de prueba: %v", err)
	}
	story := unaHistoriaEstimada(t, project, 5)
	if err := project.AssignStory(story.ID()); err != nil {
		t.Fatalf("no se pudo asignar la historia de prueba %d: %v", story.ID(), err)
	}

	if err := project.CloseSprint(nil, unDia(2026, 10, 25)); err != nil {
		t.Fatalf("se esperaba cerrar el sprint 2 sin error, salio: %v", err)
	}

	if !slices.Equal(first.CompletedStories(), []int{1}) {
		t.Errorf("se esperaba que el sprint 1 siguiera registrando completada la historia [1], tiene %v", first.CompletedStories())
	}
	if !slices.Equal(first.Stories(), []int{1}) || !first.ClosedOn().Equal(unDia(2026, 10, 19)) {
		t.Errorf("se esperaba el sprint 1 con la historia [1] y cerrado el 19/10/2026, tiene %v y cierre %v", first.Stories(), first.ClosedOn())
	}
}
