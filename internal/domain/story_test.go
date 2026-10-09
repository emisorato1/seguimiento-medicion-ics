package domain

import "testing"

// Tarea 1 de la Fase 3 de la historia #6. No tiene escenario BDD propio: los escenarios
// empiezan en la tarea 4, cuando ya se puede crear una historia en un backlog.
// Cubre: RN-2
func TestUnaHistoriaExponeTodosSusDatos(t *testing.T) {
	criterios := []string{"El sistema valida el titulo", "El sistema valida la prioridad"}
	story := &Story{
		id:                 1,
		title:              "Registrar esfuerzo",
		description:        "Poder cargar las horas trabajadas en una historia",
		priority:           PriorityHigh,
		state:              StoryPending,
		storyPoints:        0,
		acceptanceCriteria: criterios,
	}

	if story.ID() != 1 {
		t.Fatalf("se esperaba el identificador 1, salio: %d", story.ID())
	}
	if story.Title() != "Registrar esfuerzo" {
		t.Fatalf("se esperaba el titulo %q, salio: %q", "Registrar esfuerzo", story.Title())
	}
	if story.Description() != "Poder cargar las horas trabajadas en una historia" {
		t.Fatalf("se esperaba la descripcion cargada, salio: %q", story.Description())
	}
	if story.Priority() != PriorityHigh {
		t.Fatalf("se esperaba la prioridad Alta, salio: %v", story.Priority())
	}
	if story.State() != StoryPending {
		t.Fatalf("se esperaba el estado Pendiente, salio: %v", story.State())
	}
	if story.IsEstimated() {
		t.Fatalf("se esperaba que una historia con 0 story points figurara sin estimar")
	}
	if len(story.AcceptanceCriteria()) != 2 {
		t.Fatalf("se esperaban 2 criterios de aceptacion, salieron: %d", len(story.AcceptanceCriteria()))
	}
}

// Tarea 1: los criterios que se devuelven son una copia, asi nadie cambia los de la historia
// sin pasar por sus validaciones (decision de encapsulamiento de AGENTS.md).
// Cubre: RN-2
func TestCambiarLosCriteriosRecibidosNoCambiaLosDeLaHistoria(t *testing.T) {
	story := &Story{acceptanceCriteria: []string{"El sistema valida el titulo"}}

	recibidos := story.AcceptanceCriteria()
	recibidos[0] = "otra cosa"

	if story.AcceptanceCriteria()[0] != "El sistema valida el titulo" {
		t.Fatalf("se esperaba que la historia conservara su criterio original, salio: %q",
			story.AcceptanceCriteria()[0])
	}
}

// Tarea 1: el estado es un tipo propio con los tres valores de la Fase 1, para que no se
// pueda guardar un estado que no existe.
// Cubre: RN-2
func TestLosTresEstadosDeUnaHistoriaSeDistinguen(t *testing.T) {
	casos := map[StoryState]string{
		StoryPending:    "Pendiente",
		StoryInProgress: "En curso",
		StoryDone:       "Terminada",
	}

	vistos := make(map[string]bool)
	for estado, texto := range casos {
		if estado.String() != texto {
			t.Fatalf("se esperaba que el estado se mostrara como %q, salio: %q", texto, estado.String())
		}
		if vistos[texto] {
			t.Fatalf("dos estados distintos se muestran igual: %q", texto)
		}
		vistos[texto] = true
	}
}
