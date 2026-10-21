package domain

import "testing"

// Escenario: Un proyecto sin sprints no tiene sprint abierto
// Cubre: RN-10
func TestUnProyectoSinSprintsNoTieneSprintAbierto(t *testing.T) {
	project := unProyecto(t)

	sprint, ok := project.OpenSprint()

	if ok {
		t.Fatalf("se esperaba que el proyecto no tuviera sprint abierto, salio el sprint %d", sprint.Number())
	}
	if sprint != nil {
		t.Fatalf("se esperaba no recibir ningun sprint, salio uno")
	}
	if sprints := project.Sprints(); len(sprints) != 0 {
		t.Fatalf("se esperaba el proyecto sin sprints, salio con %d", len(sprints))
	}
}
