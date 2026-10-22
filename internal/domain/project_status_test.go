package domain

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"testing"
)

// statusProject registra un proyecto con historias en los estados que se piden, para los
// escenarios del estado. Los estados En curso y Terminada son de la #9 y la #10: se fijan
// desde dentro del paquete, como en la #12.
func statusProject(t *testing.T, registry *ProjectRegistry, pending, inProgress, done int) *Project {
	t.Helper()
	project, err := registry.Create(ProjectData{Name: "Sistema de metricas", Start: date(2026, 10, 1)})
	if err != nil {
		t.Fatalf("no se pudo preparar el proyecto: %v", err)
	}
	states := []StoryState{}
	for i := 0; i < pending; i++ {
		states = append(states, StoryPending)
	}
	for i := 0; i < inProgress; i++ {
		states = append(states, StoryInProgress)
	}
	for i := 0; i < done; i++ {
		states = append(states, StoryDone)
	}
	for i, state := range states {
		story, err := project.AddStory(StoryData{
			Title:              fmt.Sprintf("Historia %d", i+1),
			Description:        "Como integrante quiero algo",
			Priority:           PriorityMedium,
			AcceptanceCriteria: []string{"Se puede hacer"},
		})
		if err != nil {
			t.Fatalf("no se pudo preparar la historia %d: %v", i+1, err)
		}
		story.state = state
	}
	return project
}

// Escenario: Consultar el estado de un proyecto recien creado
// Cubre: RN-1, RN-4, RN-9
func TestConsultarElEstadoDeUnProyectoRecienCreado(t *testing.T) {
	registry := NewProjectRegistry()
	project, err := registry.Create(ProjectData{
		Name:        "Sistema de metricas",
		Description: "Seguimiento del proyecto integrador",
		Start:       date(2026, 10, 1),
		End:         date(2026, 11, 30),
	})
	if err != nil {
		t.Fatalf("no se pudo preparar el proyecto: %v", err)
	}

	status := project.Status()

	if status.ID != 1 || status.Name != "Sistema de metricas" {
		t.Errorf("se esperaba el proyecto 1 \"Sistema de metricas\", salio %d %q", status.ID, status.Name)
	}
	if !status.HasDescription || status.Description != "Seguimiento del proyecto integrador" {
		t.Errorf("se esperaba la descripcion cargada, salio %v %q", status.HasDescription, status.Description)
	}
	if !status.Start.Equal(date(2026, 10, 1)) || !status.HasEnd || !status.End.Equal(date(2026, 11, 30)) {
		t.Errorf("se esperaban las fechas 01/10/2026 y 30/11/2026, salio %v y %v (fin cargado: %v)", status.Start, status.End, status.HasEnd)
	}
	if len(status.Members) != 0 {
		t.Errorf("se esperaba la lista de integrantes vacia, salio %v", status.Members)
	}
	if status.Backlog != (BacklogSummary{}) {
		t.Errorf("se esperaba el backlog en cero, salio %+v", status.Backlog)
	}
	if status.HasCurrentSprint {
		t.Errorf("no se esperaba sprint en curso, salio %+v", status.CurrentSprint)
	}
}

// Escenario: Un proyecto sin descripcion ni fecha de finalizacion
// Cubre: RN-2
func TestUnProyectoSinDescripcionNiFechaDeFinalizacion(t *testing.T) {
	registry := NewProjectRegistry()
	project, err := registry.Create(ProjectData{Name: "Sistema de metricas", Start: date(2026, 10, 1)})
	if err != nil {
		t.Fatalf("no se pudo preparar el proyecto: %v", err)
	}

	status := project.Status()

	if status.HasDescription {
		t.Errorf("se esperaba la descripcion como no cargada, salio %q", status.Description)
	}
	if status.HasEnd {
		t.Errorf("se esperaba la fecha de finalizacion como no cargada, salio %v", status.End)
	}
}

// Escenario: Consultar el estado de un proyecto que no existe
// Cubre: RN-11. No agrega codigo: lo resuelve el registro de la #3 al buscar el proyecto, que
// es de donde se pide el estado.
func TestConsultarElEstadoDeUnProyectoQueNoExiste(t *testing.T) {
	registry := NewProjectRegistry()

	project, err := registry.Get(99)

	if !errors.Is(err, ErrProjectNotFound) {
		t.Fatalf("se esperaba ErrProjectNotFound, salio: %v", err)
	}
	if project != nil {
		t.Errorf("no se esperaba ningun proyecto del que pedir el estado, salio %v", project)
	}
}

// Escenario: El estado muestra los integrantes con su rol
// Cubre: RN-3
func TestElEstadoMuestraLosIntegrantesConSuRol(t *testing.T) {
	registry := NewProjectRegistry()
	project, err := registry.Create(ProjectData{Name: "Sistema de metricas", Start: date(2026, 10, 1)})
	if err != nil {
		t.Fatalf("no se pudo preparar el proyecto: %v", err)
	}
	for _, m := range []MemberStatus{{"Emiliano Sorato", RoleAgileEnabler}, {"Juliana Bustos", RoleProductBuilder}} {
		if _, err := project.AddMember(m.Name, m.Role); err != nil {
			t.Fatalf("no se pudo preparar el integrante %q: %v", m.Name, err)
		}
	}

	status := project.Status()

	want := []MemberStatus{{"Emiliano Sorato", RoleAgileEnabler}, {"Juliana Bustos", RoleProductBuilder}}
	if !slices.Equal(status.Members, want) {
		t.Errorf("se esperaban los integrantes %v, en ese orden, salio %v", want, status.Members)
	}
}

// Esquema del escenario: Resumen del backlog segun el estado de las historias
// Cubre: RN-5, RN-6, RN-7
func TestResumenDelBacklogSegunElEstadoDeLasHistorias(t *testing.T) {
	casos := []struct {
		pendientesSinEmpezar, enCurso, terminadas int
		want                                      BacklogSummary
	}{
		{3, 0, 0, BacklogSummary{Total: 3, Completed: 0, Pending: 3}},
		{1, 2, 0, BacklogSummary{Total: 3, Completed: 0, Pending: 3}},
		{1, 1, 2, BacklogSummary{Total: 4, Completed: 2, Pending: 2}},
		{0, 0, 3, BacklogSummary{Total: 3, Completed: 3, Pending: 0}},
	}

	for _, caso := range casos {
		nombre := fmt.Sprintf("%d pendientes, %d en curso, %d terminadas", caso.pendientesSinEmpezar, caso.enCurso, caso.terminadas)
		t.Run(nombre, func(t *testing.T) {
			project := statusProject(t, NewProjectRegistry(), caso.pendientesSinEmpezar, caso.enCurso, caso.terminadas)

			got := project.Status().Backlog

			if got != caso.want {
				t.Errorf("se esperaba el resumen %+v, salio %+v", caso.want, got)
			}
			if got.Total != got.Completed+got.Pending {
				t.Errorf("el total tiene que ser completadas mas pendientes, salio %+v", got)
			}
		})
	}
}

// Escenario: El estado refleja los cambios del proyecto
// Cubre: RN-12
func TestElEstadoReflejaLosCambiosDelProyecto(t *testing.T) {
	project := statusProject(t, NewProjectRegistry(), 2, 0, 0)
	antes := project.Status()

	if _, err := project.AddStory(StoryData{
		Title: "Una mas", Description: "Como integrante quiero otra", Priority: PriorityLow, AcceptanceCriteria: []string{"Se carga"},
	}); err != nil {
		t.Fatalf("no se pudo agregar la historia: %v", err)
	}
	despues := project.Status()

	if antes.Backlog.Total != 2 {
		t.Errorf("antes del cambio se esperaban 2 historias, salieron %d", antes.Backlog.Total)
	}
	if despues.Backlog != (BacklogSummary{Total: 3, Completed: 0, Pending: 3}) {
		t.Errorf("despues del cambio se esperaban 3 historias pendientes, salio %+v", despues.Backlog)
	}
}

// Escenario: Consultar el estado no cambia nada
// Cubre: RN-10
func TestConsultarElEstadoNoCambiaNada(t *testing.T) {
	project := statusProject(t, NewProjectRegistry(), 3, 0, 0)
	for _, name := range []string{"Emiliano Sorato", "Juliana Bustos"} {
		if _, err := project.AddMember(name, RoleProductBuilder); err != nil {
			t.Fatalf("no se pudo preparar el integrante %q: %v", name, err)
		}
	}
	if _, err := project.AddSprint(SprintData{Goal: "Tener el dashboard", Start: date(2026, 10, 1), End: date(2026, 10, 15)}); err != nil {
		t.Fatalf("no se pudo preparar el sprint: %v", err)
	}

	primera := project.Status()
	segunda := project.Status()

	if !reflect.DeepEqual(primera, segunda) {
		t.Errorf("dos consultas seguidas tenian que dar lo mismo, salieron:\n%+v\n%+v", primera, segunda)
	}
	// Cambiar lo recibido no cambia el proyecto: es una foto con copias.
	primera.Members[0].Name = "Otra persona"
	if len(project.Members()) != 2 || project.Members()[0].Name() != "Emiliano Sorato" {
		t.Errorf("el proyecto tenia que seguir con sus 2 integrantes, quedo con %d", len(project.Members()))
	}
	if len(project.Backlog()) != 3 {
		t.Errorf("el proyecto tenia que seguir con sus 3 historias, quedo con %d", len(project.Backlog()))
	}
	if sprint, ok := project.OpenSprint(); !ok || sprint.Number() != 1 {
		t.Errorf("el sprint 1 tenia que seguir abierto")
	}
}

// Escenario: Cada proyecto muestra solo lo suyo
// Cubre: RN-13
func TestCadaProyectoMuestraSoloLoSuyo(t *testing.T) {
	registry := NewProjectRegistry()
	a := statusProject(t, registry, 3, 0, 0)
	for _, name := range []string{"Emiliano Sorato", "Juliana Bustos"} {
		if _, err := a.AddMember(name, RoleProductBuilder); err != nil {
			t.Fatalf("no se pudo preparar el integrante %q: %v", name, err)
		}
	}
	b, err := registry.Create(ProjectData{Name: "Otro proyecto", Start: date(2026, 10, 1)})
	if err != nil {
		t.Fatalf("no se pudo preparar el proyecto B: %v", err)
	}

	status := b.Status()

	if len(status.Members) != 0 || status.Backlog.Total != 0 {
		t.Errorf("el proyecto B no tenia que contar nada del A, salio %d integrantes y %d historias", len(status.Members), status.Backlog.Total)
	}
}

// statusSprint crea un sprint en el proyecto con el alta de la #8. Cerrarlo es de la #10 y
// asignarle historias de la #9: se hace desde dentro del paquete, como en la #8.
func statusSprint(t *testing.T, project *Project, goal string, closed bool, stories ...int) *Sprint {
	t.Helper()
	sprint, err := project.AddSprint(SprintData{Goal: goal, Start: date(2026, 10, 20), End: date(2026, 10, 25)})
	if err != nil {
		t.Fatalf("no se pudo preparar el sprint %q: %v", goal, err)
	}
	sprint.stories = stories
	if closed {
		sprint.state = SprintClosed
	}
	return sprint
}

// Escenario: El estado muestra el sprint en curso
// Cubre: RN-8
func TestElEstadoMuestraElSprintEnCurso(t *testing.T) {
	project := statusProject(t, NewProjectRegistry(), 3, 0, 0)
	statusSprint(t, project, "Tener el MVP", true)
	statusSprint(t, project, "Tener el dashboard", false, 1, 2, 3)

	status := project.Status()

	want := SprintSummary{Number: 2, Goal: "Tener el dashboard", Start: date(2026, 10, 20), End: date(2026, 10, 25), AssignedStories: 3}
	if !status.HasCurrentSprint || status.CurrentSprint != want {
		t.Errorf("se esperaba el sprint en curso %+v, salio %+v (hay sprint en curso: %v)", want, status.CurrentSprint, status.HasCurrentSprint)
	}
}

// Escenario: Un proyecto con todos sus sprints cerrados no tiene sprint en curso
// Cubre: RN-9
func TestUnProyectoConTodosSusSprintsCerradosNoTieneSprintEnCurso(t *testing.T) {
	project := statusProject(t, NewProjectRegistry(), 1, 0, 0)
	statusSprint(t, project, "Tener el MVP", true)
	statusSprint(t, project, "Tener el dashboard", true)

	status := project.Status()

	if status.HasCurrentSprint {
		t.Errorf("no se esperaba sprint en curso, salio %+v", status.CurrentSprint)
	}
}
