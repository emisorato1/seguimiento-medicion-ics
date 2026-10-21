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

// Escenario: Crear el primer sprint de un proyecto
// Cubre: RN-1, RN-2, RN-3
func TestCrearElPrimerSprintDeUnProyecto(t *testing.T) {
	project := unProyecto(t)

	sprint, err := project.AddSprint(SprintData{
		Goal:  "Tener el backlog estimado",
		Start: unDia(2026, 11, 1),
		End:   unDia(2026, 11, 14),
	})

	if err != nil {
		t.Fatalf("se esperaba crear el sprint sin error, salio: %v", err)
	}
	if sprint.Number() != 1 {
		t.Errorf("se esperaba el sprint numero 1, salio el %d", sprint.Number())
	}
	if sprint.State() != SprintOpen {
		t.Errorf("se esperaba el sprint en estado Abierto, salio %v", sprint.State())
	}
	if stories := sprint.Stories(); len(stories) != 0 {
		t.Errorf("se esperaba el sprint sin historias asignadas, salio con %v", stories)
	}
	if sprints := project.Sprints(); len(sprints) != 1 || sprints[0] != sprint {
		t.Errorf("se esperaba el proyecto con ese unico sprint, salio con %d", len(sprints))
	}
}
