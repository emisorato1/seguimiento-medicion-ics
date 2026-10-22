package domain

import (
	"errors"
	"slices"
	"testing"
)

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
