package domain

import (
	"errors"
	"strings"
	"testing"
	"time"
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

// unDia arma una fecha sin hora, que es como las compara el dominio.
func unDia(anio int, mes time.Month, dia int) time.Time {
	return time.Date(anio, mes, dia, 0, 0, 0, 0, time.UTC)
}

// unProyecto devuelve un proyecto registrado y vacio, para los tests que no prueban la
// creacion del proyecto en si.
func unProyecto(t *testing.T) *Project {
	t.Helper()
	registry := NewProjectRegistry()
	project, err := registry.Create(ProjectData{Name: "Sistema de metricas", Start: unDia(2026, 10, 1)})
	if err != nil {
		t.Fatalf("no se pudo crear el proyecto de prueba: %v", err)
	}
	return project
}

// unosDatos devuelve datos validos de historia, para los tests que no prueban la validacion.
func unosDatos() StoryData {
	return StoryData{
		Title:              "Registrar esfuerzo",
		Description:        "Poder cargar las horas trabajadas",
		Priority:           PriorityHigh,
		AcceptanceCriteria: []string{"El sistema valida las horas"},
	}
}

// Escenario: Backlog de un proyecto sin historias
// Cubre: RN-15
func TestBacklogDeUnProyectoSinHistorias(t *testing.T) {
	registry := NewProjectRegistry()
	project, err := registry.Create(ProjectData{Name: "Sistema de metricas", Start: unDia(2026, 10, 1)})
	if err != nil {
		t.Fatalf("no se pudo crear el proyecto de prueba: %v", err)
	}

	backlog := project.Backlog()

	if len(backlog) != 0 {
		t.Fatalf("se esperaba un backlog vacio, salieron %d historias", len(backlog))
	}
}

// Escenario: Crear una historia en el backlog
// Cubre: RN-1, RN-3, RN-7
func TestCrearUnaHistoriaEnElBacklog(t *testing.T) {
	project := unProyecto(t)

	story, err := project.AddStory(StoryData{
		Title:              "Registrar esfuerzo",
		Description:        "Poder cargar las horas trabajadas",
		Priority:           PriorityHigh,
		AcceptanceCriteria: []string{"El sistema valida las horas"},
	})

	if err != nil {
		t.Fatalf("se esperaba que la historia se creara, salio: %v", err)
	}
	if story.ID() != 1 {
		t.Fatalf("se esperaba el identificador 1 para la primera historia, salio: %d", story.ID())
	}
	if story.State() != StoryPending {
		t.Fatalf("se esperaba el estado Pendiente, salio: %v", story.State())
	}
	if story.IsEstimated() {
		t.Fatalf("se esperaba que la historia naciera sin estimar")
	}
	if len(project.Backlog()) != 1 {
		t.Fatalf("se esperaba 1 historia en el backlog, salieron: %d", len(project.Backlog()))
	}
}

// Escenario: Los identificadores son unicos dentro de cada proyecto
// Cubre: RN-1, RN-16
func TestLosIdentificadoresSonUnicosDentroDeCadaProyecto(t *testing.T) {
	registry := NewProjectRegistry()
	proyectoA, _ := registry.Create(ProjectData{Name: "Proyecto A", Start: unDia(2026, 10, 1)})
	proyectoB, _ := registry.Create(ProjectData{Name: "Proyecto B", Start: unDia(2026, 10, 1)})

	primeraDeA, _ := proyectoA.AddStory(unosDatos())
	primeraDeB, _ := proyectoB.AddStory(unosDatos())

	if primeraDeA.ID() != 1 || primeraDeB.ID() != 1 {
		t.Fatalf("se esperaba que la primera historia de cada proyecto tuviera el identificador 1, salieron: %d y %d",
			primeraDeA.ID(), primeraDeB.ID())
	}
}

// Escenario: Los identificadores no se repiten
// Cubre: RN-1
func TestLosIdentificadoresDeLasHistoriasNoSeRepiten(t *testing.T) {
	project := unProyecto(t)

	primera, _ := project.AddStory(unosDatos())
	segunda, _ := project.AddStory(unosDatos())

	if primera.ID() == segunda.ID() {
		t.Fatalf("se esperaban identificadores distintos, las dos salieron con %d", primera.ID())
	}
}

// Escenario: Guardar el titulo sin espacios en los extremos
// Escenario: Descartar los criterios de aceptacion vacios
// Cubre: RN-6
func TestNormalizarLosDatosDeUnaHistoria(t *testing.T) {
	project := unProyecto(t)

	story, err := project.AddStory(StoryData{
		Title:              "  Registrar esfuerzo  ",
		Description:        "  Poder cargar las horas  ",
		Priority:           PriorityHigh,
		AcceptanceCriteria: []string{"  El sistema valida las horas  ", "", "   "},
	})

	if err != nil {
		t.Fatalf("se esperaba que la historia se creara, salio: %v", err)
	}
	if story.Title() != "Registrar esfuerzo" {
		t.Fatalf("se esperaba el titulo sin espacios en los extremos, salio: %q", story.Title())
	}
	if story.Description() != "Poder cargar las horas" {
		t.Fatalf("se esperaba la descripcion sin espacios en los extremos, salio: %q", story.Description())
	}
	criterios := story.AcceptanceCriteria()
	if len(criterios) != 1 || criterios[0] != "El sistema valida las horas" {
		t.Fatalf("se esperaba 1 criterio sin espacios y sin los vacios, salieron: %q", criterios)
	}
}

// Esquema del escenario: Rechazar una historia sin los datos obligatorios
// Escenario: Un texto con solo espacios cuenta como vacio
// Escenario: Rechazar si todos los criterios quedan vacios
// Cubre: RN-5, RN-11
func TestRechazarUnaHistoriaSinLosDatosObligatorios(t *testing.T) {
	casos := []struct {
		dato     string
		datos    StoryData
		esperado error
	}{
		{"titulo", StoryData{Title: "", Description: "d", AcceptanceCriteria: []string{"c"}}, ErrStoryTitleRequired},
		{"titulo con solo espacios", StoryData{Title: "   ", Description: "d", AcceptanceCriteria: []string{"c"}}, ErrStoryTitleRequired},
		{"descripcion", StoryData{Title: "t", Description: "", AcceptanceCriteria: []string{"c"}}, ErrStoryDescriptionRequired},
		{"descripcion con solo espacios", StoryData{Title: "t", Description: "  ", AcceptanceCriteria: []string{"c"}}, ErrStoryDescriptionRequired},
		{"criterios de aceptacion", StoryData{Title: "t", Description: "d", AcceptanceCriteria: nil}, ErrStoryCriteriaRequired},
		{"criterios todos vacios", StoryData{Title: "t", Description: "d", AcceptanceCriteria: []string{"", "   "}}, ErrStoryCriteriaRequired},
	}

	for _, caso := range casos {
		t.Run(caso.dato, func(t *testing.T) {
			project := unProyecto(t)

			_, err := project.AddStory(caso.datos)

			if !errors.Is(err, caso.esperado) {
				t.Fatalf("se esperaba %v al faltar %s, salio: %v", caso.esperado, caso.dato, err)
			}
			if len(project.Backlog()) != 0 {
				t.Fatalf("se esperaba que el backlog siguiera vacio tras el rechazo, salieron %d historias",
					len(project.Backlog()))
			}
		})
	}
}

// Escenario: Dos historias con el mismo titulo
// Cubre: RN-14
func TestDosHistoriasDelMismoProyectoPuedenTenerElMismoTitulo(t *testing.T) {
	project := unProyecto(t)

	primera, _ := project.AddStory(unosDatos())
	segunda, err := project.AddStory(unosDatos())

	if err != nil {
		t.Fatalf("se esperaba que dos historias pudieran tener el mismo titulo, salio: %v", err)
	}
	if primera.Title() != segunda.Title() {
		t.Fatalf("los titulos deberian ser iguales en este caso")
	}
	if primera.ID() == segunda.ID() {
		t.Fatalf("se esperaban identificadores distintos, los dos salieron %d", primera.ID())
	}
}

// Escenario: Modificar una historia pendiente
// Escenario: Modificar una historia sin cambiar nada
// Escenario: Reemplazar los criterios de aceptacion
// Cubre: RN-8, RN-9
func TestModificarUnaHistoria(t *testing.T) {
	project := unProyecto(t)
	story, _ := project.AddStory(unosDatos())
	nuevaPrioridad := PriorityLow

	err := project.UpdateStory(story.ID(), StoryChanges{Priority: &nuevaPrioridad})

	if err != nil {
		t.Fatalf("se esperaba que la historia se modificara, salio: %v", err)
	}
	if story.Priority() != PriorityLow {
		t.Fatalf("se esperaba la prioridad Baja, salio: %v", story.Priority())
	}
	if story.Title() != "Registrar esfuerzo" {
		t.Fatalf("se esperaba que el titulo quedara como estaba, salio: %q", story.Title())
	}
}

// Escenario: Modificar una historia en curso
// Cubre: RN-8
func TestModificarUnaHistoriaEnCurso(t *testing.T) {
	project := unProyecto(t)
	story, _ := project.AddStory(unosDatos())
	story.state = StoryInProgress
	nuevoTitulo := "Registrar esfuerzo por integrante"

	err := project.UpdateStory(story.ID(), StoryChanges{Title: &nuevoTitulo})

	if err != nil {
		t.Fatalf("se esperaba poder modificar una historia En curso, salio: %v", err)
	}
	if story.Title() != nuevoTitulo {
		t.Fatalf("se esperaba el titulo nuevo, salio: %q", story.Title())
	}
}

// Escenario: Modificar una historia sin cambiar nada
// Cubre: RN-9
func TestModificarUnaHistoriaSinCambiarNada(t *testing.T) {
	project := unProyecto(t)
	story, _ := project.AddStory(unosDatos())

	err := project.UpdateStory(story.ID(), StoryChanges{})

	if err != nil {
		t.Fatalf("se esperaba que modificar sin cambios no diera error, salio: %v", err)
	}
	if story.Title() != "Registrar esfuerzo" || story.Priority() != PriorityHigh {
		t.Fatalf("se esperaba que la historia conservara sus datos")
	}
}

// Escenario: Reemplazar los criterios de aceptacion
// Cubre: RN-9
func TestReemplazarLosCriteriosDeAceptacion(t *testing.T) {
	project := unProyecto(t)
	story, _ := project.AddStory(unosDatos())
	nuevos := []string{"Primero", "Segundo", "Tercero"}

	err := project.UpdateStory(story.ID(), StoryChanges{AcceptanceCriteria: &nuevos})

	if err != nil {
		t.Fatalf("se esperaba que los criterios se reemplazaran, salio: %v", err)
	}
	if len(story.AcceptanceCriteria()) != 3 {
		t.Fatalf("se esperaban 3 criterios, salieron: %d", len(story.AcceptanceCriteria()))
	}
}

// Escenario: Una modificacion rechazada no cambia nada
// Cubre: RN-5, RN-11
func TestUnaModificacionRechazadaNoCambiaNada(t *testing.T) {
	project := unProyecto(t)
	story, _ := project.AddStory(unosDatos())
	enBlanco := "   "

	err := project.UpdateStory(story.ID(), StoryChanges{Title: &enBlanco})

	if !errors.Is(err, ErrStoryTitleRequired) {
		t.Fatalf("se esperaba ErrStoryTitleRequired, salio: %v", err)
	}
	if story.Title() != "Registrar esfuerzo" {
		t.Fatalf("se esperaba que la historia conservara su titulo, salio: %q", story.Title())
	}
}

// Escenario: No modificar una historia terminada
// Cubre: RN-10
func TestNoModificarUnaHistoriaTerminada(t *testing.T) {
	project := unProyecto(t)
	story, _ := project.AddStory(unosDatos())
	// El test arma el estado directamente: las transiciones son de la #9 y la #10 y
	// todavia no existen. Es posible porque el test vive en el mismo paquete.
	story.state = StoryDone
	nuevoTitulo := "Otro titulo"

	err := project.UpdateStory(story.ID(), StoryChanges{Title: &nuevoTitulo})

	if !errors.Is(err, ErrStoryDone) {
		t.Fatalf("se esperaba ErrStoryDone, salio: %v", err)
	}
	if story.Title() != "Registrar esfuerzo" {
		t.Fatalf("se esperaba que la historia terminada conservara su titulo, salio: %q", story.Title())
	}
}
