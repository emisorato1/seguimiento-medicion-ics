package domain

import (
	"errors"
	"strings"
	"testing"
)

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

// Tarea 2 de la Fase 3: el dominio convierte el texto que escribe el usuario en una
// prioridad, ignorando mayusculas y espacios de los extremos. Asi la consola (#24) no
// duplica la RN-4: solo pasa el texto.
// Cubre: RN-4
func TestConvertirElTextoEnUnaPrioridad(t *testing.T) {
	casos := []struct {
		texto    string
		esperada Priority
	}{
		{"Alta", PriorityHigh},
		{"Media", PriorityMedium},
		{"Baja", PriorityLow},
		{"alta", PriorityHigh},
		{"MEDIA", PriorityMedium},
		{"  Baja  ", PriorityLow},
	}

	for _, caso := range casos {
		t.Run(caso.texto, func(t *testing.T) {
			prioridad, err := ParsePriority(caso.texto)
			if err != nil {
				t.Fatalf("se esperaba que %q fuera una prioridad valida, salio: %v", caso.texto, err)
			}
			if prioridad != caso.esperada {
				t.Fatalf("se esperaba la prioridad %v para %q, salio: %v", caso.esperada, caso.texto, prioridad)
			}
		})
	}
}

// Escenario: Rechazar una prioridad invalida
// Cubre: RN-4
func TestRechazarUnaPrioridadInvalida(t *testing.T) {
	casos := []string{"Urgente", "", "   ", "altisima"}

	for _, texto := range casos {
		t.Run(texto, func(t *testing.T) {
			_, err := ParsePriority(texto)

			if !errors.Is(err, ErrStoryPriorityInvalid) {
				t.Fatalf("se esperaba ErrStoryPriorityInvalid para %q, salio: %v", texto, err)
			}
		})
	}
}

// Tarea 2: el mensaje del error nombra las tres prioridades validas, asi la consola solo
// tiene que mostrarlo (RN-2 de la #24).
// Cubre: RN-4
func TestElErrorDePrioridadNombraLasValidas(t *testing.T) {
	_, err := ParsePriority("Urgente")

	for _, esperada := range []string{"Alta", "Media", "Baja"} {
		if !strings.Contains(err.Error(), esperada) {
			t.Fatalf("se esperaba que el mensaje nombrara la prioridad %q, salio: %q", esperada, err.Error())
		}
	}
}

// Tarea 2: las tres prioridades se muestran con el nombre que usa la Fase 1, y es el mismo
// texto que ParsePriority acepta: lo que se muestra se puede volver a escribir.
// Cubre: RN-4
func TestLasTresPrioridadesSeMuestranYSeVuelvenAConvertir(t *testing.T) {
	casos := map[Priority]string{
		PriorityHigh:   "Alta",
		PriorityMedium: "Media",
		PriorityLow:    "Baja",
	}

	for prioridad, texto := range casos {
		if prioridad.String() != texto {
			t.Fatalf("se esperaba que la prioridad se mostrara como %q, salio: %q", texto, prioridad.String())
		}
		vuelta, err := ParsePriority(prioridad.String())
		if err != nil || vuelta != prioridad {
			t.Fatalf("se esperaba que %q volviera a convertirse en la misma prioridad, salio: %v (%v)",
				prioridad.String(), vuelta, err)
		}
	}
}

// Tarea 1: los story points son uno de los datos que se registran por la RN-2. El valor lo
// carga la historia #12; aca solo se guarda y se lee.
// Cubre: RN-2
func TestUnaHistoriaDevuelveSusStoryPoints(t *testing.T) {
	sinEstimar := &Story{storyPoints: 0}
	estimada := &Story{storyPoints: 5}

	if sinEstimar.StoryPoints() != 0 || sinEstimar.IsEstimated() {
		t.Fatalf("se esperaba que una historia con 0 story points figurara sin estimar, salio: %d",
			sinEstimar.StoryPoints())
	}
	if estimada.StoryPoints() != 5 || !estimada.IsEstimated() {
		t.Fatalf("se esperaban 5 story points y la historia estimada, salio: %d", estimada.StoryPoints())
	}
}
