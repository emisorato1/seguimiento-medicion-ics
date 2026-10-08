package domain

import (
	"errors"
	"testing"
	"time"
)

// date arma una fecha sin hora, que es como el dominio maneja las fechas.
func date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// Escenario: Crear un proyecto con todos los datos
// Cubre: RN-1
func TestCrearUnProyectoConTodosLosDatos(t *testing.T) {
	registry := NewProjectRegistry()

	project, err := registry.Create(ProjectData{
		Name:        "Sistema de metricas",
		Description: "TP de ICS",
		Start:       date(2026, 10, 1),
		End:         date(2026, 11, 30),
	})

	if err != nil {
		t.Fatalf("se esperaba crear el proyecto sin error, salio: %v", err)
	}
	if project.ID() != 1 {
		t.Fatalf("se esperaba el identificador 1, salio: %d", project.ID())
	}

	found, err := registry.Get(project.ID())
	if err != nil {
		t.Fatalf("se esperaba encontrar el proyecto %d, salio: %v", project.ID(), err)
	}
	if found.Name() != "Sistema de metricas" {
		t.Errorf("se esperaba el nombre %q, salio: %q", "Sistema de metricas", found.Name())
	}
	if found.Description() != "TP de ICS" {
		t.Errorf("se esperaba la descripcion %q, salio: %q", "TP de ICS", found.Description())
	}
	if !found.Start().Equal(date(2026, 10, 1)) {
		t.Errorf("se esperaba el inicio 01/10/2026, salio: %v", found.Start())
	}
	if !found.End().Equal(date(2026, 11, 30)) {
		t.Errorf("se esperaba el fin 30/11/2026, salio: %v", found.End())
	}
}

// Escenario: Crear un proyecto solo con los datos obligatorios
// Cubre: datos opcionales de la Fase 1 (descripcion y fecha de finalizacion)
func TestCrearUnProyectoSoloConLosDatosObligatorios(t *testing.T) {
	registry := NewProjectRegistry()

	project, err := registry.Create(ProjectData{
		Name:  "Sistema de metricas",
		Start: date(2026, 10, 1),
	})

	if err != nil {
		t.Fatalf("se esperaba crear el proyecto sin error, salio: %v", err)
	}
	if project.Description() != "" {
		t.Errorf("se esperaba el proyecto sin descripcion, salio: %q", project.Description())
	}
	if project.HasEnd() {
		t.Errorf("se esperaba el proyecto sin fecha de finalizacion, salio: %v", project.End())
	}
}

// Escenario: Dos proyectos con el mismo nombre
// Cubre: RN-1, RN-7
func TestDosProyectosConElMismoNombre(t *testing.T) {
	registry := NewProjectRegistry()
	first, err := registry.Create(ProjectData{Name: "Sistema de metricas", Start: date(2026, 10, 1)})
	if err != nil {
		t.Fatalf("se esperaba crear el primer proyecto sin error, salio: %v", err)
	}

	second, err := registry.Create(ProjectData{Name: "Sistema de metricas", Start: date(2026, 10, 1)})

	if err != nil {
		t.Fatalf("se esperaba crear el segundo proyecto con el mismo nombre sin error, salio: %v", err)
	}
	if second.ID() == first.ID() {
		t.Fatalf("se esperaban identificadores distintos, los dos salieron: %d", first.ID())
	}
	if _, err := registry.Get(first.ID()); err != nil {
		t.Errorf("se esperaba seguir encontrando el primer proyecto, salio: %v", err)
	}
}

// Escenario: Guardar el nombre sin espacios en los extremos
// Cubre: RN-10
func TestGuardarElNombreSinEspaciosEnLosExtremos(t *testing.T) {
	registry := NewProjectRegistry()

	project, err := registry.Create(ProjectData{Name: "  Sistema de metricas  ", Start: date(2026, 10, 1)})

	if err != nil {
		t.Fatalf("se esperaba crear el proyecto sin error, salio: %v", err)
	}
	if project.Name() != "Sistema de metricas" {
		t.Errorf("se esperaba el nombre %q, salio: %q", "Sistema de metricas", project.Name())
	}
}

// Escenario: Aceptar un nombre de un solo caracter
// Cubre: RN-10
func TestAceptarUnNombreDeUnSoloCaracter(t *testing.T) {
	registry := NewProjectRegistry()

	project, err := registry.Create(ProjectData{Name: "X", Start: date(2026, 10, 1)})

	if err != nil {
		t.Fatalf("se esperaba crear el proyecto sin error, salio: %v", err)
	}
	if project.Name() != "X" {
		t.Errorf("se esperaba el nombre %q, salio: %q", "X", project.Name())
	}
}

// Escenario: Rechazar un proyecto sin nombre
// Cubre: RN-2
func TestRechazarUnProyectoSinNombre(t *testing.T) {
	cases := []struct {
		name        string
		projectName string
	}{
		{name: "vacio", projectName: ""},
		{name: "solo espacios", projectName: "   "},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			registry := NewProjectRegistry()

			_, err := registry.Create(ProjectData{Name: c.projectName, Start: date(2026, 10, 1)})

			if !errors.Is(err, ErrProjectNameRequired) {
				t.Fatalf("se esperaba ErrProjectNameRequired, salio: %v", err)
			}

			// La creacion rechazada no registro nada: el proximo proyecto sigue siendo el 1.
			next, err := registry.Create(ProjectData{Name: "Sistema de metricas", Start: date(2026, 10, 1)})
			if err != nil {
				t.Fatalf("se esperaba crear el proyecto valido sin error, salio: %v", err)
			}
			if next.ID() != 1 {
				t.Errorf("se esperaba que el rechazo no registrara nada y el proximo fuera el 1, salio: %d", next.ID())
			}
		})
	}
}

// Escenario: Rechazar un proyecto sin fecha de inicio
// Cubre: RN-3
func TestRechazarUnProyectoSinFechaDeInicio(t *testing.T) {
	registry := NewProjectRegistry()

	_, err := registry.Create(ProjectData{Name: "Sistema de metricas"})

	if !errors.Is(err, ErrProjectStartRequired) {
		t.Fatalf("se esperaba ErrProjectStartRequired, salio: %v", err)
	}
}
