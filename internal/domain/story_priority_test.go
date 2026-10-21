package domain

import (
	"errors"
	"testing"
)

// conHistorias devuelve un proyecto con una historia por cada prioridad indicada, en ese
// orden: la primera recibe el identificador 1, la segunda el 2, y asi.
func conHistorias(t *testing.T, priorities ...Priority) (*Project, []*Story) {
	t.Helper()
	project := unProyecto(t)
	stories := make([]*Story, 0, len(priorities))
	for _, priority := range priorities {
		data := unosDatos()
		data.Priority = priority
		story, err := project.AddStory(data)
		if err != nil {
			t.Fatalf("no se pudo crear la historia de prueba: %v", err)
		}
		stories = append(stories, story)
	}
	return project, stories
}

// identificadores devuelve los identificadores de las historias, en el orden en que vienen.
func identificadores(stories []*Story) []int {
	ids := make([]int, len(stories))
	for i, story := range stories {
		ids[i] = story.ID()
	}
	return ids
}

// mismoOrden indica si los identificadores son exactamente esos, en ese orden.
func mismoOrden(got []int, want ...int) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range want {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}

// Escenario: Backlog ordenado de un proyecto sin historias
// Cubre: RN-8
func TestBacklogOrdenadoDeUnProyectoSinHistorias(t *testing.T) {
	project := unProyecto(t)

	stories := project.PrioritizedBacklog()

	if stories == nil {
		t.Fatalf("se esperaba una lista vacia, no nil")
	}
	if len(stories) != 0 {
		t.Fatalf("se esperaba el listado vacio, salio con %d historias", len(stories))
	}
}

// Escenario: Pedir el backlog ordenado de un proyecto que no existe
// Cubre: RN-9
//
// Pasa sin codigo de esta historia: lo resuelve el registro de la #3, que es el unico camino
// para llegar a un proyecto. Va en la tarea 1, como dice la Fase 3.
func TestPedirElBacklogOrdenadoDeUnProyectoQueNoExiste(t *testing.T) {
	registry := NewProjectRegistry()

	project, err := registry.Get(99)

	if !errors.Is(err, ErrProjectNotFound) {
		t.Fatalf("se esperaba ErrProjectNotFound, salio: %v", err)
	}
	if project != nil {
		t.Fatalf("se esperaba no recibir ningun proyecto, salio el %d", project.ID())
	}
}

// Escenario: El listado incluye las historias en cualquier estado
// Cubre: RN-4
func TestElListadoIncluyeLasHistoriasEnCualquierEstado(t *testing.T) {
	project, stories := conHistorias(t, PriorityHigh, PriorityMedium, PriorityLow)
	// Asignar a un sprint (#9) y cerrarlo (#10) todavia no existen: el test fija el estado
	// desde dentro del paquete, como la #12 en su tarea 9.
	stories[1].state = StoryInProgress
	stories[2].state = StoryDone

	listed := project.PrioritizedBacklog()

	if len(listed) != 3 {
		t.Fatalf("se esperaban las 3 historias, salieron %d: %v", len(listed), identificadores(listed))
	}
	seen := map[int]int{}
	for _, story := range listed {
		seen[story.ID()]++
	}
	for _, story := range stories {
		if seen[story.ID()] != 1 {
			t.Errorf("se esperaba la historia %d (%s) una sola vez, aparecio %d veces",
				story.ID(), story.State(), seen[story.ID()])
		}
	}
}

// Escenario: Las historias del listado traen todos sus datos
// Cubre: RN-5
func TestLasHistoriasDelListadoTraenTodosSusDatos(t *testing.T) {
	project, stories := conHistorias(t, PriorityLow, PriorityHigh, PriorityMedium)
	if err := project.EstimateStory(stories[0].ID(), 5); err != nil {
		t.Fatalf("no se pudo estimar la historia de prueba: %v", err)
	}

	listed := project.PrioritizedBacklog()

	if len(listed) != 3 {
		t.Fatalf("se esperaban las 3 historias, salieron %d", len(listed))
	}
	for _, story := range listed {
		saved, err := project.Story(story.ID())
		if err != nil {
			t.Fatalf("no se encontro en el backlog la historia %d del listado: %v", story.ID(), err)
		}
		// La historia del listado es la misma que guarda el backlog, no una copia
		// incompleta: trae el identificador, el titulo, la descripcion, la prioridad, el
		// estado, los story points y los criterios que tiene guardados.
		if story != saved {
			t.Fatalf("se esperaba en el listado la historia %d guardada en el backlog, salio otra", story.ID())
		}
	}
	if listed[0].Title() == "" || listed[0].Description() == "" || len(listed[0].AcceptanceCriteria()) == 0 {
		t.Errorf("se esperaba la historia %d con titulo, descripcion y criterios, salio: %q, %q, %v",
			listed[0].ID(), listed[0].Title(), listed[0].Description(), listed[0].AcceptanceCriteria())
	}
}

// Escenario: Backlog ordenado de un proyecto con una sola historia
// Cubre: RN-4
func TestBacklogOrdenadoDeUnProyectoConUnaSolaHistoria(t *testing.T) {
	project, _ := conHistorias(t, PriorityLow)

	listed := project.PrioritizedBacklog()

	if !mismoOrden(identificadores(listed), 1) {
		t.Fatalf("se esperaba solo la historia 1, salio: %v", identificadores(listed))
	}
}
