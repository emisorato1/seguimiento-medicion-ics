package domain

import "testing"

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
