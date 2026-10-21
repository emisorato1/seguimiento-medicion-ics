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

// Escenario: Listar el backlog ordenado por prioridad
// Cubre: RN-1, RN-2
func TestListarElBacklogOrdenadoPorPrioridad(t *testing.T) {
	project, _ := conHistorias(t, PriorityLow, PriorityHigh, PriorityMedium)

	listed := project.PrioritizedBacklog()

	if !mismoOrden(identificadores(listed), 2, 3, 1) {
		t.Fatalf("se esperaba el orden [2 3 1], salio: %v", identificadores(listed))
	}
}

// Escenario: No hay historias de prioridad alta
// Cubre: RN-2
func TestNoHayHistoriasDePrioridadAlta(t *testing.T) {
	project, _ := conHistorias(t, PriorityLow, PriorityMedium)

	listed := project.PrioritizedBacklog()

	if !mismoOrden(identificadores(listed), 2, 1) {
		t.Fatalf("se esperaba el orden [2 1], sin huecos, salio: %v", identificadores(listed))
	}
}

// Esquema del escenario: El orden se actualiza al cambiar una prioridad
// Cubre: RN-6
func TestElOrdenSeActualizaAlCambiarUnaPrioridad(t *testing.T) {
	cases := []struct {
		name     string
		priority Priority
		want     []int
	}{
		{"Alta", PriorityHigh, []int{1, 3, 2}},
		{"Media", PriorityMedium, []int{1, 2, 3}},
		{"Baja", PriorityLow, []int{1, 2, 3}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			project, _ := conHistorias(t, PriorityHigh, PriorityMedium, PriorityLow)
			if err := project.UpdateStory(3, StoryChanges{Priority: ptr(tc.priority)}); err != nil {
				t.Fatalf("no se pudo cambiar la prioridad de la historia 3: %v", err)
			}

			listed := project.PrioritizedBacklog()

			if !mismoOrden(identificadores(listed), tc.want...) {
				t.Fatalf("se esperaba el orden %v, salio: %v", tc.want, identificadores(listed))
			}
		})
	}
}

// Escenario: Bajar la prioridad de la historia que estaba primera
// Cubre: RN-6
func TestBajarLaPrioridadDeLaHistoriaQueEstabaPrimera(t *testing.T) {
	project, _ := conHistorias(t, PriorityHigh, PriorityMedium, PriorityMedium)
	if err := project.UpdateStory(1, StoryChanges{Priority: ptr(PriorityLow)}); err != nil {
		t.Fatalf("no se pudo bajar la prioridad de la historia 1: %v", err)
	}

	listed := project.PrioritizedBacklog()

	if !mismoOrden(identificadores(listed), 2, 3, 1) {
		t.Fatalf("se esperaba el orden [2 3 1], salio: %v", identificadores(listed))
	}
}

// Escenario: Cambiar una prioridad por la que ya tenia
// Cubre: RN-6
func TestCambiarUnaPrioridadPorLaQueYaTenia(t *testing.T) {
	project, _ := conHistorias(t, PriorityHigh, PriorityMedium, PriorityLow)
	if err := project.UpdateStory(2, StoryChanges{Priority: ptr(PriorityMedium)}); err != nil {
		t.Fatalf("no se pudo cambiar la prioridad de la historia 2: %v", err)
	}

	listed := project.PrioritizedBacklog()

	if !mismoOrden(identificadores(listed), 1, 2, 3) {
		t.Fatalf("se esperaba el orden [1 2 3], salio: %v", identificadores(listed))
	}
}

// Escenario: Pedir el listado ordenado no cambia el backlog
// Cubre: RN-7
func TestPedirElListadoOrdenadoNoCambiaElBacklog(t *testing.T) {
	project, stories := conHistorias(t, PriorityLow, PriorityHigh, PriorityMedium)

	project.PrioritizedBacklog()

	if !mismoOrden(identificadores(project.Backlog()), 1, 2, 3) {
		t.Fatalf("se esperaba el backlog sin ordenar en [1 2 3], salio: %v", identificadores(project.Backlog()))
	}
	want := []Priority{PriorityLow, PriorityHigh, PriorityMedium}
	for i, story := range stories {
		if story.Priority() != want[i] {
			t.Errorf("se esperaba la historia %d con prioridad %s, salio %s", story.ID(), want[i], story.Priority())
		}
	}
}

// Escenario: Desempatar por identificador dentro de la misma prioridad
// Cubre: RN-3
func TestDesempatarPorIdentificadorDentroDeLaMismaPrioridad(t *testing.T) {
	project, _ := conHistorias(t, PriorityMedium, PriorityHigh, PriorityHigh)

	listed := project.PrioritizedBacklog()

	if !mismoOrden(identificadores(listed), 2, 3, 1) {
		t.Fatalf("se esperaba el orden [2 3 1], salio: %v", identificadores(listed))
	}
}

// Escenario: Todas las historias tienen la misma prioridad
// Cubre: RN-3
func TestTodasLasHistoriasTienenLaMismaPrioridad(t *testing.T) {
	project, _ := conHistorias(t, PriorityMedium, PriorityMedium, PriorityMedium)

	listed := project.PrioritizedBacklog()

	if !mismoOrden(identificadores(listed), 1, 2, 3) {
		t.Fatalf("se esperaba el orden [1 2 3], salio: %v", identificadores(listed))
	}
}

// Escenario: Dos listados seguidos devuelven lo mismo
// Cubre: RN-3
func TestDosListadosSeguidosDevuelvenLoMismo(t *testing.T) {
	project, _ := conHistorias(t, PriorityLow, PriorityHigh, PriorityMedium)

	first := identificadores(project.PrioritizedBacklog())
	second := identificadores(project.PrioritizedBacklog())

	if !mismoOrden(second, first...) {
		t.Fatalf("se esperaba el mismo orden las dos veces, salio %v y despues %v", first, second)
	}
}

// Tarea 4: desempatar por identificador aunque el backlog este guardado en otro orden
// Cubre: RN-3 y la decision de la Fase 2 "El desempate por identificador se hace explicito"
//
// Con los escenarios de la Fase 1 solos el desempate no falla, porque las historias que se
// cargan con la operacion de la #6 quedan guardadas en orden de identificador. Este test
// arma desde dentro del paquete un backlog guardado al reves, que es lo que pasaria si algun
// dia hay borrado o reordenamiento, y comprueba que el listado no depende de eso.
func TestDesempatarPorIdentificadorAunqueElBacklogEsteGuardadoEnOtroOrden(t *testing.T) {
	project, stories := conHistorias(t, PriorityHigh, PriorityHigh, PriorityHigh)
	project.backlog = []*Story{stories[2], stories[0], stories[1]}

	listed := project.PrioritizedBacklog()

	if !mismoOrden(identificadores(listed), 1, 2, 3) {
		t.Fatalf("se esperaba el orden [1 2 3] aunque el backlog este guardado como [3 1 2], salio: %v",
			identificadores(listed))
	}
}
