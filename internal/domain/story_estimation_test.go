package domain

import "testing"

// Escenario: Consultar una historia sin estimar
// Cubre: RN-4
//
// Desvío: este test pasa sin código nuevo, porque IsEstimated ya estaba en dev desde el
// PR #56 de la historia #6. Queda anotado en la Fase 4.
func TestConsultarUnaHistoriaSinEstimar(t *testing.T) {
	story := &Story{state: StoryPending}

	if story.IsEstimated() {
		t.Fatalf("se esperaba que la historia figurara sin estimar, salio estimada con %d story points", story.StoryPoints())
	}
}
