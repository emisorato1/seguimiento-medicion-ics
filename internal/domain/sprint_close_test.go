package domain

import "testing"

// unSprintParaCerrar arma un proyecto con el sprint 1 abierto y una historia asignada por
// cada valor de puntos, estimada con ese valor. Asignar es de la #9.
func unSprintParaCerrar(t *testing.T, points ...int) (*Project, *Sprint) {
	t.Helper()
	project := unProyecto(t)
	historiasEstimadas(t, project, points...)
	sprint := unSprint(t, project)
	for i := range points {
		if err := project.AssignStory(i + 1); err != nil {
			t.Fatalf("no se pudo asignar la historia de prueba %d: %v", i+1, err)
		}
	}
	return project, sprint
}

// Escenario: Despues de cerrar se puede crear el sprint siguiente
// Cubre: RN-1, RN-14
func TestDespuesDeCerrarSePuedeCrearElSprintSiguiente(t *testing.T) {
	project, sprint := unSprintParaCerrar(t)

	err := project.CloseSprint(nil, unDia(2026, 10, 30))

	if err != nil {
		t.Fatalf("se esperaba cerrar el sprint sin error, salio: %v", err)
	}
	if sprint.State() != SprintClosed || !sprint.ClosedOn().Equal(unDia(2026, 10, 30)) {
		t.Fatalf("se esperaba el sprint Cerrado el 30/10/2026, quedo %v con fecha %v", sprint.State(), sprint.ClosedOn())
	}
	if _, ok := project.OpenSprint(); ok {
		t.Fatalf("se esperaba que el proyecto quedara sin sprint abierto")
	}
	next, err := project.AddSprint(SprintData{Goal: "Seguir midiendo", Start: unDia(2026, 10, 31), End: unDia(2026, 11, 1)})
	if err != nil || next.Number() != 2 {
		t.Fatalf("se esperaba poder crear el sprint 2, salio %v (error: %v)", next, err)
	}
}
