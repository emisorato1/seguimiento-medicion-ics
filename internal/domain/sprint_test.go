package domain

import (
	"errors"
	"testing"
)

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

// Escenario: Consultar el sprint abierto de un proyecto
// Cubre: RN-10
func TestConsultarElSprintAbiertoDeUnProyecto(t *testing.T) {
	project := unProyecto(t)
	created, err := project.AddSprint(SprintData{
		Goal:  "Tener el backlog estimado",
		Start: unDia(2026, 11, 1),
		End:   unDia(2026, 11, 14),
	})
	if err != nil {
		t.Fatalf("no se pudo crear el sprint de prueba: %v", err)
	}

	sprint, ok := project.OpenSprint()

	if !ok {
		t.Fatalf("se esperaba que el proyecto tuviera sprint abierto, salio que no")
	}
	if sprint != created || sprint.Number() != 1 {
		t.Fatalf("se esperaba el sprint 1 recien creado, salio otro")
	}
	if sprint.Goal() != "Tener el backlog estimado" {
		t.Errorf("se esperaba el goal %q, salio %q", "Tener el backlog estimado", sprint.Goal())
	}
	if !sprint.Start().Equal(unDia(2026, 11, 1)) || !sprint.End().Equal(unDia(2026, 11, 14)) {
		t.Errorf("se esperaban las fechas 01/11/2026 y 14/11/2026, salio %v y %v", sprint.Start(), sprint.End())
	}
	if sprint.State() != SprintOpen {
		t.Errorf("se esperaba el sprint en estado Abierto, salio %v", sprint.State())
	}
}

// Escenario: Consultar un sprint por su numero
// Cubre: RN-15
func TestConsultarUnSprintPorSuNumero(t *testing.T) {
	project := unProyecto(t)
	first, err := project.AddSprint(SprintData{Goal: "Tener el backlog estimado", Start: unDia(2026, 11, 1), End: unDia(2026, 11, 14)})
	if err != nil {
		t.Fatalf("no se pudo crear el sprint 1 de prueba: %v", err)
	}
	// Cerrar un sprint es de la #10, que todavia no existe: el test fija el estado desde
	// dentro del paquete, como la #12 con las historias Terminadas.
	first.state = SprintClosed
	if _, err := project.AddSprint(SprintData{Goal: "Medir la velocidad", Start: unDia(2026, 11, 15), End: unDia(2026, 11, 28)}); err != nil {
		t.Fatalf("no se pudo crear el sprint 2 de prueba: %v", err)
	}

	sprint, err := project.Sprint(1)

	if err != nil {
		t.Fatalf("se esperaba encontrar el sprint 1, salio: %v", err)
	}
	if sprint.Number() != 1 || sprint.Goal() != "Tener el backlog estimado" {
		t.Fatalf("se esperaba el sprint 1 con su goal, salio el %d con %q", sprint.Number(), sprint.Goal())
	}
	if !sprint.Start().Equal(unDia(2026, 11, 1)) || !sprint.End().Equal(unDia(2026, 11, 14)) {
		t.Errorf("se esperaban las fechas 01/11/2026 y 14/11/2026, salio %v y %v", sprint.Start(), sprint.End())
	}
	if sprint.State() != SprintClosed {
		t.Errorf("se esperaba el sprint 1 en estado Cerrado, salio %v", sprint.State())
	}
}

// Escenario: Consultar un sprint que no existe
// Cubre: RN-15
func TestConsultarUnSprintQueNoExiste(t *testing.T) {
	project := unProyecto(t)
	if _, err := project.AddSprint(SprintData{Goal: "Tener el backlog estimado", Start: unDia(2026, 11, 1), End: unDia(2026, 11, 14)}); err != nil {
		t.Fatalf("no se pudo crear el sprint de prueba: %v", err)
	}

	sprint, err := project.Sprint(5)

	if !errors.Is(err, ErrSprintNotFound) {
		t.Fatalf("se esperaba ErrSprintNotFound, salio: %v", err)
	}
	if sprint != nil {
		t.Fatalf("se esperaba no recibir ningun sprint, salio el %d", sprint.Number())
	}
}

// Esquema del escenario: Rechazar un sprint sin Sprint Goal
// Cubre: RN-4, RN-12
func TestRechazarUnSprintSinSprintGoal(t *testing.T) {
	cases := []struct {
		name string
		goal string
	}{
		{"vacio", ""},
		{"solo espacios", "   "},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			project := unProyecto(t)

			_, err := project.AddSprint(SprintData{Goal: tc.goal, Start: unDia(2026, 11, 1), End: unDia(2026, 11, 14)})

			if !errors.Is(err, ErrSprintGoalRequired) {
				t.Fatalf("se esperaba ErrSprintGoalRequired, salio: %v", err)
			}
			if sprints := project.Sprints(); len(sprints) != 0 {
				t.Fatalf("se esperaba el proyecto sin sprints, salio con %d", len(sprints))
			}
		})
	}
}

// Esquema del escenario: Rechazar un sprint sin las fechas obligatorias
// Cubre: RN-5, RN-12
func TestRechazarUnSprintSinLasFechasObligatorias(t *testing.T) {
	cases := []struct {
		name string
		data SprintData
		want error
	}{
		{"fecha de inicio", SprintData{Goal: "Tener el backlog estimado", End: unDia(2026, 11, 14)}, ErrSprintStartRequired},
		{"fecha de fin", SprintData{Goal: "Tener el backlog estimado", Start: unDia(2026, 11, 1)}, ErrSprintEndRequired},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			project := unProyecto(t)

			_, err := project.AddSprint(tc.data)

			if !errors.Is(err, tc.want) {
				t.Fatalf("se esperaba %v, salio: %v", tc.want, err)
			}
			if sprints := project.Sprints(); len(sprints) != 0 {
				t.Fatalf("se esperaba el proyecto sin sprints, salio con %d", len(sprints))
			}
		})
	}
}

// Escenario: Rechazar un sprint con fechas incoherentes
// Cubre: RN-6, RN-12
func TestRechazarUnSprintConFechasIncoherentes(t *testing.T) {
	project := unProyecto(t)

	_, err := project.AddSprint(SprintData{Goal: "Tener el backlog estimado", Start: unDia(2026, 11, 14), End: unDia(2026, 11, 1)})

	if !errors.Is(err, ErrSprintDatesInconsistent) {
		t.Fatalf("se esperaba ErrSprintDatesInconsistent, salio: %v", err)
	}
	if sprints := project.Sprints(); len(sprints) != 0 {
		t.Fatalf("se esperaba el proyecto sin sprints, salio con %d", len(sprints))
	}
}

// Escenario: Un sprint de un solo dia es valido
// Cubre: RN-6
func TestUnSprintDeUnSoloDiaEsValido(t *testing.T) {
	project := unProyecto(t)

	sprint, err := project.AddSprint(SprintData{Goal: "Tener el backlog estimado", Start: unDia(2026, 11, 1), End: unDia(2026, 11, 1)})

	if err != nil {
		t.Fatalf("se esperaba crear el sprint de un dia sin error, salio: %v", err)
	}
	if sprint.Number() != 1 {
		t.Errorf("se esperaba el sprint 1 registrado, salio el %d", sprint.Number())
	}
}

// Escenario: Un sprint con fechas pasadas es valido
// Cubre: RN-14
func TestUnSprintConFechasPasadasEsValido(t *testing.T) {
	project := unProyecto(t)

	sprint, err := project.AddSprint(SprintData{Goal: "Tener el backlog estimado", Start: unDia(2026, 3, 1), End: unDia(2026, 3, 14)})

	if err != nil {
		t.Fatalf("se esperaba crear el sprint con fechas pasadas sin error, salio: %v", err)
	}
	if !sprint.Start().Equal(unDia(2026, 3, 1)) || !sprint.End().Equal(unDia(2026, 3, 14)) {
		t.Errorf("se esperaban las fechas tal como se indicaron, salio %v y %v", sprint.Start(), sprint.End())
	}
}

// Escenario: Guardar el Sprint Goal sin espacios en los extremos
// Cubre: RN-7
func TestGuardarElSprintGoalSinEspaciosEnLosExtremos(t *testing.T) {
	project := unProyecto(t)

	sprint, err := project.AddSprint(SprintData{Goal: "  Tener el backlog estimado  ", Start: unDia(2026, 11, 1), End: unDia(2026, 11, 14)})

	if err != nil {
		t.Fatalf("se esperaba crear el sprint sin error, salio: %v", err)
	}
	if sprint.Goal() != "Tener el backlog estimado" {
		t.Errorf("se esperaba el goal %q, salio %q", "Tener el backlog estimado", sprint.Goal())
	}
}
