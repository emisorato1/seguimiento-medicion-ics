package domain

import (
	"errors"
	"testing"
)

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
