package domain

import (
	"errors"
	"testing"
)

// unSprintAbierto devuelve un proyecto con el sprint 1 abierto, para los tests que no prueban
// la creacion del sprint.
func unSprintAbierto(t *testing.T) (*Project, *Sprint) {
	t.Helper()
	project := unProyecto(t)
	sprint, err := project.AddSprint(SprintData{Goal: "Tener el backlog estimado", Start: unDia(2026, 11, 1), End: unDia(2026, 11, 14)})
	if err != nil {
		t.Fatalf("no se pudo crear el sprint de prueba: %v", err)
	}
	return project, sprint
}

// Escenario: Un sprint recien abierto tiene el Sprint Backlog vacio
// Cubre: RN-9
func TestUnSprintRecienAbiertoTieneElSprintBacklogVacio(t *testing.T) {
	project, _ := unSprintAbierto(t)

	sprint, ok := project.OpenSprint()

	if !ok {
		t.Fatalf("se esperaba el sprint 1 abierto, salio que no hay")
	}
	stories := sprint.Stories()
	if stories == nil {
		t.Fatalf("se esperaba una lista vacia, no nil")
	}
	if len(stories) != 0 {
		t.Fatalf("se esperaba el Sprint Backlog vacio, salio %v", stories)
	}
}

// unaHistoriaEstimada agrega al proyecto una historia Pendiente estimada en esos puntos.
func unaHistoriaEstimada(t *testing.T, project *Project, points int) *Story {
	t.Helper()
	story, err := project.AddStory(unosDatos())
	if err != nil {
		t.Fatalf("no se pudo crear la historia de prueba: %v", err)
	}
	if err := project.EstimateStory(story.ID(), points); err != nil {
		t.Fatalf("no se pudo estimar la historia de prueba: %v", err)
	}
	return story
}

// Escenario: Asignar una historia estimada al sprint abierto
// Cubre: RN-1, RN-4
func TestAsignarUnaHistoriaEstimadaAlSprintAbierto(t *testing.T) {
	project, sprint := unSprintAbierto(t)
	story := unaHistoriaEstimada(t, project, 5)

	err := project.AssignStory(story.ID())

	if err != nil {
		t.Fatalf("se esperaba asignar la historia sin error, salio: %v", err)
	}
	if stories := sprint.Stories(); len(stories) != 1 || stories[0] != story.ID() {
		t.Errorf("se esperaba el Sprint Backlog del sprint 1 con la historia %d, salio %v", story.ID(), stories)
	}
	if story.State() != StoryInProgress {
		t.Errorf("se esperaba la historia En curso, salio %v", story.State())
	}
}

// Escenario: El Sprint Backlog respeta el orden en que se asignaron las historias
// Cubre: RN-9, RN-15
func TestElSprintBacklogRespetaElOrdenEnQueSeAsignaronLasHistorias(t *testing.T) {
	project, sprint := unSprintAbierto(t)
	first := unaHistoriaEstimada(t, project, 3)
	second := unaHistoriaEstimada(t, project, 5)
	third := unaHistoriaEstimada(t, project, 8)

	for _, story := range []*Story{third, first, second} {
		if err := project.AssignStory(story.ID()); err != nil {
			t.Fatalf("no se pudo asignar la historia %d: %v", story.ID(), err)
		}
	}

	want := []int{third.ID(), first.ID(), second.ID()}
	got := sprint.Stories()
	if len(got) != len(want) || got[0] != want[0] || got[1] != want[1] || got[2] != want[2] {
		t.Fatalf("se esperaba el Sprint Backlog %v, salio %v", want, got)
	}
}

// Escenario: Rechazar una historia sin estimar
// Cubre: RN-3, RN-14
func TestRechazarUnaHistoriaSinEstimar(t *testing.T) {
	project, sprint := unSprintAbierto(t)
	story, err := project.AddStory(unosDatos())
	if err != nil {
		t.Fatalf("no se pudo crear la historia de prueba: %v", err)
	}

	err = project.AssignStory(story.ID())

	if !errors.Is(err, ErrStoryNotEstimated) {
		t.Fatalf("se esperaba ErrStoryNotEstimated, salio: %v", err)
	}
	if stories := sprint.Stories(); len(stories) != 0 {
		t.Errorf("se esperaba el Sprint Backlog vacio, salio %v", stories)
	}
	if story.State() != StoryPending {
		t.Errorf("se esperaba la historia todavia Pendiente, salio %v", story.State())
	}
}

// Escenario: Rechazar una historia que ya esta en el sprint
// Cubre: RN-5, RN-14
func TestRechazarUnaHistoriaQueYaEstaEnElSprint(t *testing.T) {
	project, sprint := unSprintAbierto(t)
	story := unaHistoriaEstimada(t, project, 5)
	if err := project.AssignStory(story.ID()); err != nil {
		t.Fatalf("no se pudo asignar la historia de prueba: %v", err)
	}

	err := project.AssignStory(story.ID())

	if !errors.Is(err, ErrStoryAlreadyInSprint) {
		t.Fatalf("se esperaba ErrStoryAlreadyInSprint, salio: %v", err)
	}
	if stories := sprint.Stories(); len(stories) != 1 {
		t.Errorf("se esperaba el Sprint Backlog con una sola historia, salio %v", stories)
	}
}

// unaHistoriaTerminada deja la historia Terminada desde dentro del paquete, porque terminar
// es de la #10, que todavia no existe. Es lo mismo que hace la #12.
func unaHistoriaTerminada(story *Story) {
	story.state = StoryDone
}

// Escenario: Rechazar una historia terminada
// Cubre: RN-6, RN-14
func TestRechazarUnaHistoriaTerminada(t *testing.T) {
	project := unProyecto(t)
	first, err := project.AddSprint(SprintData{Goal: "Tener el backlog estimado", Start: unDia(2026, 11, 1), End: unDia(2026, 11, 14)})
	if err != nil {
		t.Fatalf("no se pudo crear el sprint 1 de prueba: %v", err)
	}
	cerrarSprint(first)
	sprint, err := project.AddSprint(SprintData{Goal: "Medir la velocidad", Start: unDia(2026, 11, 15), End: unDia(2026, 11, 28)})
	if err != nil {
		t.Fatalf("no se pudo crear el sprint 2 de prueba: %v", err)
	}
	story := unaHistoriaEstimada(t, project, 5)
	unaHistoriaTerminada(story)

	err = project.AssignStory(story.ID())

	if !errors.Is(err, ErrStoryDone) {
		t.Fatalf("se esperaba ErrStoryDone, salio: %v", err)
	}
	if stories := sprint.Stories(); len(stories) != 0 {
		t.Errorf("se esperaba el Sprint Backlog del sprint 2 vacio, salio %v", stories)
	}
}

// Escenario: Rechazar quitar una historia terminada
// Cubre: RN-6, RN-14
func TestRechazarQuitarUnaHistoriaTerminada(t *testing.T) {
	project, first := unSprintAbierto(t)
	story := unaHistoriaEstimada(t, project, 5)
	if err := project.AssignStory(story.ID()); err != nil {
		t.Fatalf("no se pudo asignar la historia de prueba: %v", err)
	}
	// El sprint 1 se cierra con la historia completada: es lo que hara la #10.
	unaHistoriaTerminada(story)
	cerrarSprint(first)
	if _, err := project.AddSprint(SprintData{Goal: "Medir la velocidad", Start: unDia(2026, 11, 15), End: unDia(2026, 11, 28)}); err != nil {
		t.Fatalf("no se pudo crear el sprint 2 de prueba: %v", err)
	}

	err := project.UnassignStory(story.ID())

	if !errors.Is(err, ErrStoryDone) {
		t.Fatalf("se esperaba ErrStoryDone, salio: %v", err)
	}
	if story.State() != StoryDone {
		t.Errorf("se esperaba la historia todavia Terminada, salio %v", story.State())
	}
}

// Esquema del escenario: Operar sin un sprint abierto
// Cubre: RN-2
func TestOperarSinUnSprintAbierto(t *testing.T) {
	cases := []struct {
		name string
		run  func(project *Project, id int) error
	}{
		{"asignar", func(project *Project, id int) error { return project.AssignStory(id) }},
		{"quitar", func(project *Project, id int) error { return project.UnassignStory(id) }},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			project := unProyecto(t)
			story := unaHistoriaEstimada(t, project, 5)

			err := tc.run(project, story.ID())

			if !errors.Is(err, ErrNoOpenSprint) {
				t.Fatalf("se esperaba ErrNoOpenSprint, salio: %v", err)
			}
		})
	}
}

// Escenario: Quitar una historia del sprint abierto
// Cubre: RN-7
func TestQuitarUnaHistoriaDelSprintAbierto(t *testing.T) {
	project, sprint := unSprintAbierto(t)
	story := unaHistoriaEstimada(t, project, 5)
	if err := project.AssignStory(story.ID()); err != nil {
		t.Fatalf("no se pudo asignar la historia de prueba: %v", err)
	}

	err := project.UnassignStory(story.ID())

	if err != nil {
		t.Fatalf("se esperaba quitar la historia sin error, salio: %v", err)
	}
	if stories := sprint.Stories(); len(stories) != 0 {
		t.Errorf("se esperaba el Sprint Backlog del sprint 1 vacio, salio %v", stories)
	}
	if story.State() != StoryPending {
		t.Errorf("se esperaba la historia Pendiente, salio %v", story.State())
	}
}

// Escenario: Quitar una historia y volver a asignarla
// Cubre: RN-7, RN-9
func TestQuitarUnaHistoriaYVolverAAsignarla(t *testing.T) {
	project, sprint := unSprintAbierto(t)
	first := unaHistoriaEstimada(t, project, 3)
	second := unaHistoriaEstimada(t, project, 5)
	for _, story := range []*Story{first, second} {
		if err := project.AssignStory(story.ID()); err != nil {
			t.Fatalf("no se pudo asignar la historia %d: %v", story.ID(), err)
		}
	}

	if err := project.UnassignStory(first.ID()); err != nil {
		t.Fatalf("se esperaba quitar la historia %d sin error, salio: %v", first.ID(), err)
	}
	if err := project.AssignStory(first.ID()); err != nil {
		t.Fatalf("se esperaba volver a asignar la historia %d sin error, salio: %v", first.ID(), err)
	}

	got := sprint.Stories()
	if len(got) != 2 || got[0] != second.ID() || got[1] != first.ID() {
		t.Fatalf("se esperaba el Sprint Backlog [%d %d], salio %v", second.ID(), first.ID(), got)
	}
}
