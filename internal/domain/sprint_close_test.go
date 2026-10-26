package domain

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

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

// estadoDe devuelve el estado de la historia con ese identificador o corta el test.
func estadoDe(t *testing.T, project *Project, id int) StoryState {
	t.Helper()
	story, err := project.Story(id)
	if err != nil {
		t.Fatalf("no se encontro la historia %d: %v", id, err)
	}
	return story.State()
}

// Escenario: Cerrar el sprint registrando que se completo y que no
// Cubre: RN-1, RN-2, RN-3, RN-4
func TestCerrarElSprintRegistrandoQueSeCompletoYQueNo(t *testing.T) {
	project, sprint := unSprintParaCerrar(t, 3, 5, 2)

	err := project.CloseSprint([]int{1, 2}, unDia(2026, 10, 25))

	if err != nil {
		t.Fatalf("se esperaba cerrar el sprint sin error, salio: %v", err)
	}
	if sprint.State() != SprintClosed || !sprint.ClosedOn().Equal(unDia(2026, 10, 25)) {
		t.Fatalf("se esperaba el sprint Cerrado el 25/10/2026, quedo %v con fecha %v", sprint.State(), sprint.ClosedOn())
	}
	if got := fmt.Sprint(sprint.CompletedStories()); got != "[1 2]" {
		t.Fatalf("se esperaban completadas las historias [1 2], salieron %s", got)
	}
	for id, want := range map[int]StoryState{1: StoryDone, 2: StoryDone, 3: StoryPending} {
		if got := estadoDe(t, project, id); got != want {
			t.Fatalf("se esperaba la historia %d en %v, quedo en %v", id, want, got)
		}
	}
}

// Escenario: El sprint cerrado conserva lo que se le habia comprometido
// Cubre: RN-5
func TestElSprintCerradoConservaLoQueSeLeHabiaComprometido(t *testing.T) {
	project, sprint := unSprintParaCerrar(t, 3, 5, 2)

	if err := project.CloseSprint([]int{1}, unDia(2026, 10, 25)); err != nil {
		t.Fatalf("se esperaba cerrar el sprint sin error, salio: %v", err)
	}

	if got := fmt.Sprint(sprint.Stories()); got != "[1 2 3]" {
		t.Fatalf("se esperaba que el sprint siguiera con las historias [1 2 3], salieron %s", got)
	}
}

// Esquema del escenario: Cerrar con distinta cantidad de historias completadas
// Cubre: RN-3, RN-4, RN-9
func TestCerrarConDistintaCantidadDeHistoriasCompletadas(t *testing.T) {
	cases := []struct {
		completed              []int
		terminadas, pendientes int
	}{
		{nil, 0, 3},
		{[]int{2}, 1, 2},
		{[]int{1, 2, 3}, 3, 0},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d completadas", len(tc.completed)), func(t *testing.T) {
			project, sprint := unSprintParaCerrar(t, 3, 5, 2)

			if err := project.CloseSprint(tc.completed, unDia(2026, 10, 25)); err != nil {
				t.Fatalf("se esperaba cerrar el sprint sin error, salio: %v", err)
			}

			terminadas, pendientes := 0, 0
			for _, id := range []int{1, 2, 3} {
				switch estadoDe(t, project, id) {
				case StoryDone:
					terminadas++
				case StoryPending:
					pendientes++
				}
			}
			if terminadas != tc.terminadas || pendientes != tc.pendientes || sprint.State() != SprintClosed {
				t.Fatalf("se esperaban %d Terminadas, %d Pendientes y el sprint Cerrado, salieron %d, %d y %v",
					tc.terminadas, tc.pendientes, terminadas, pendientes, sprint.State())
			}
		})
	}
}

// Escenario: Cerrar un sprint sin ninguna historia asignada
// Cubre: RN-9
func TestCerrarUnSprintSinNingunaHistoriaAsignada(t *testing.T) {
	project, sprint := unSprintParaCerrar(t)

	if err := project.CloseSprint(nil, unDia(2026, 10, 25)); err != nil {
		t.Fatalf("se esperaba cerrar el sprint sin error, salio: %v", err)
	}

	if sprint.State() != SprintClosed || len(sprint.CompletedStories()) != 0 {
		t.Fatalf("se esperaba el sprint Cerrado y sin completadas, quedo %v con %v", sprint.State(), sprint.CompletedStories())
	}
}

// Escenario: La misma historia indicada dos veces cuenta una sola vez
// Cubre: RN-2, RN-8
func TestLaMismaHistoriaIndicadaDosVecesCuentaUnaSolaVez(t *testing.T) {
	project, sprint := unSprintParaCerrar(t, 3)

	if err := project.CloseSprint([]int{1, 1}, unDia(2026, 10, 25)); err != nil {
		t.Fatalf("se esperaba cerrar el sprint sin error, salio: %v", err)
	}

	if got := fmt.Sprint(sprint.CompletedStories()); got != "[1]" {
		t.Fatalf("se esperaba una sola historia completada, salieron %s", got)
	}
	if got := estadoDe(t, project, 1); got != StoryDone {
		t.Fatalf("se esperaba la historia 1 Terminada, quedo en %v", got)
	}
}

// Escenario: Rechazar el cierre si una historia completada no esta en el sprint
// Cubre: RN-7, RN-13
func TestRechazarElCierreSiUnaHistoriaCompletadaNoEstaEnElSprint(t *testing.T) {
	project, sprint := unSprintParaCerrar(t, 3)
	historiasEstimadas(t, project, 5)

	err := project.CloseSprint([]int{1, 2}, unDia(2026, 10, 25))

	if !errors.Is(err, ErrStoryNotInSprint) {
		t.Fatalf("se esperaba ErrStoryNotInSprint, salio: %v", err)
	}
	if sprint.State() != SprintOpen || len(sprint.CompletedStories()) != 0 {
		t.Fatalf("se esperaba el sprint Abierto y sin completadas, quedo %v con %v", sprint.State(), sprint.CompletedStories())
	}
	if got := estadoDe(t, project, 1); got != StoryInProgress {
		t.Fatalf("se esperaba la historia 1 En curso, quedo en %v", got)
	}
}

// Escenario: Rechazar el cierre sin fecha de cierre
// Cubre: RN-10, RN-13
func TestRechazarElCierreSinFechaDeCierre(t *testing.T) {
	project, sprint := unSprintParaCerrar(t, 3)

	err := project.CloseSprint([]int{1}, time.Time{})

	if !errors.Is(err, ErrSprintCloseDateRequired) {
		t.Fatalf("se esperaba ErrSprintCloseDateRequired, salio: %v", err)
	}
	if sprint.State() != SprintOpen || !sprint.ClosedOn().IsZero() {
		t.Fatalf("se esperaba el sprint Abierto y sin fecha de cierre, quedo %v con %v", sprint.State(), sprint.ClosedOn())
	}
	if got := estadoDe(t, project, 1); got != StoryInProgress {
		t.Fatalf("se esperaba la historia 1 En curso, quedo en %v", got)
	}
}

// Escenario: Rechazar el cierre si no hay sprint abierto
// Cubre: RN-6
func TestRechazarElCierreSiNoHaySprintAbierto(t *testing.T) {
	project := unProyecto(t)

	err := project.CloseSprint(nil, unDia(2026, 10, 25))

	if !errors.Is(err, ErrNoOpenSprint) {
		t.Fatalf("se esperaba ErrNoOpenSprint, salio: %v", err)
	}
}

// Escenario: No se puede cerrar dos veces el mismo sprint
// Cubre: RN-12, RN-13
func TestNoSePuedeCerrarDosVecesElMismoSprint(t *testing.T) {
	project, sprint := unSprintParaCerrar(t, 3, 5)
	if err := project.CloseSprint([]int{1}, unDia(2026, 10, 25)); err != nil {
		t.Fatalf("no se pudo cerrar el sprint de prueba: %v", err)
	}

	err := project.CloseSprint([]int{2}, unDia(2026, 10, 28))

	if !errors.Is(err, ErrNoOpenSprint) {
		t.Fatalf("se esperaba ErrNoOpenSprint, salio: %v", err)
	}
	if !sprint.ClosedOn().Equal(unDia(2026, 10, 25)) || fmt.Sprint(sprint.CompletedStories()) != "[1]" {
		t.Fatalf("se esperaba que el sprint conservara el cierre del 25/10 con la historia 1, quedo %v con %v",
			sprint.ClosedOn(), sprint.CompletedStories())
	}
}
