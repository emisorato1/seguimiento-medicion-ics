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

// ptr devuelve un puntero al valor, para indicar en un pedido de cambios que un dato cambia.
func ptr[T any](value T) *T {
	return &value
}

// fullProject registra el proyecto de los escenarios de modificacion: "Sistema de metricas",
// descripcion "TP de ICS", inicio 01/10/2026 y fin 30/11/2026.
func fullProject(t *testing.T, registry *ProjectRegistry) *Project {
	t.Helper()
	project, err := registry.Create(ProjectData{
		Name:        "Sistema de metricas",
		Description: "TP de ICS",
		Start:       date(2026, 10, 1),
		End:         date(2026, 11, 30),
	})
	if err != nil {
		t.Fatalf("se esperaba registrar el proyecto de prueba sin error, salio: %v", err)
	}
	return project
}

// assertProject verifica los cuatro datos de un proyecto.
func assertProject(t *testing.T, p *Project, name, description string, start, end time.Time) {
	t.Helper()
	if p.Name() != name {
		t.Errorf("se esperaba el nombre %q, salio: %q", name, p.Name())
	}
	if p.Description() != description {
		t.Errorf("se esperaba la descripcion %q, salio: %q", description, p.Description())
	}
	if !p.Start().Equal(start) {
		t.Errorf("se esperaba el inicio %v, salio: %v", start, p.Start())
	}
	if !p.End().Equal(end) {
		t.Errorf("se esperaba el fin %v, salio: %v", end, p.End())
	}
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

// Escenario: Rechazar fechas incoherentes
// Cubre: RN-4
func TestRechazarFechasIncoherentes(t *testing.T) {
	registry := NewProjectRegistry()

	_, err := registry.Create(ProjectData{
		Name:  "Sistema de metricas",
		Start: date(2026, 10, 1),
		End:   date(2026, 9, 30),
	})

	if !errors.Is(err, ErrProjectDatesInconsistent) {
		t.Fatalf("se esperaba ErrProjectDatesInconsistent, salio: %v", err)
	}
}

// Escenario: Aceptar un proyecto que empieza y termina el mismo dia
// Cubre: RN-4
func TestAceptarUnProyectoQueEmpiezaYTerminaElMismoDia(t *testing.T) {
	registry := NewProjectRegistry()

	project, err := registry.Create(ProjectData{
		Name:  "Sistema de metricas",
		Start: date(2026, 10, 1),
		End:   date(2026, 10, 1),
	})

	if err != nil {
		t.Fatalf("se esperaba crear el proyecto sin error, salio: %v", err)
	}
	if !project.End().Equal(date(2026, 10, 1)) {
		t.Errorf("se esperaba el fin 01/10/2026, salio: %v", project.End())
	}
}

// Escenario: Consultar un proyecto que no existe
// Cubre: RN-6
func TestConsultarUnProyectoQueNoExiste(t *testing.T) {
	registry := NewProjectRegistry()

	_, err := registry.Get(99)

	if !errors.Is(err, ErrProjectNotFound) {
		t.Fatalf("se esperaba ErrProjectNotFound, salio: %v", err)
	}
}

// Escenario: Modificar un proyecto que no existe
// Cubre: RN-6
func TestModificarUnProyectoQueNoExiste(t *testing.T) {
	registry := NewProjectRegistry()

	err := registry.Update(99, ProjectChanges{})

	if !errors.Is(err, ErrProjectNotFound) {
		t.Fatalf("se esperaba ErrProjectNotFound, salio: %v", err)
	}
	if _, err := registry.Get(99); !errors.Is(err, ErrProjectNotFound) {
		t.Errorf("se esperaba que no se creara el proyecto 99, salio: %v", err)
	}
	// Tampoco se creo ningun otro: el proximo proyecto sigue siendo el 1.
	next, err := registry.Create(ProjectData{Name: "Sistema de metricas", Start: date(2026, 10, 1)})
	if err != nil {
		t.Fatalf("se esperaba crear el proyecto valido sin error, salio: %v", err)
	}
	if next.ID() != 1 {
		t.Errorf("se esperaba que no se hubiera creado ningun proyecto y el proximo fuera el 1, salio: %d", next.ID())
	}
}

// Escenario: Modificar solo el nombre conserva los demas datos
// Cubre: RN-8
func TestModificarSoloElNombreConservaLosDemasDatos(t *testing.T) {
	registry := NewProjectRegistry()
	project := fullProject(t, registry)

	err := registry.Update(project.ID(), ProjectChanges{Name: ptr("Seguimiento y medicion")})

	if err != nil {
		t.Fatalf("se esperaba modificar el proyecto sin error, salio: %v", err)
	}
	assertProject(t, project, "Seguimiento y medicion", "TP de ICS", date(2026, 10, 1), date(2026, 11, 30))
}

// Escenario: Modificar un proyecto sin cambiar nada
// Cubre: RN-8
func TestModificarUnProyectoSinCambiarNada(t *testing.T) {
	registry := NewProjectRegistry()
	project := fullProject(t, registry)

	err := registry.Update(project.ID(), ProjectChanges{Name: ptr("Sistema de metricas")})

	if err != nil {
		t.Fatalf("se esperaba que modificar sin cambios no diera error, salio: %v", err)
	}
	assertProject(t, project, "Sistema de metricas", "TP de ICS", date(2026, 10, 1), date(2026, 11, 30))
}

// Escenario: Guardar sin espacios en los extremos el nombre modificado
// Cubre: RN-5, RN-10
func TestGuardarSinEspaciosEnLosExtremosElNombreModificado(t *testing.T) {
	registry := NewProjectRegistry()
	project := fullProject(t, registry)

	err := registry.Update(project.ID(), ProjectChanges{Name: ptr("  Seguimiento y medicion  ")})

	if err != nil {
		t.Fatalf("se esperaba modificar el proyecto sin error, salio: %v", err)
	}
	if project.Name() != "Seguimiento y medicion" {
		t.Errorf("se esperaba el nombre %q, salio: %q", "Seguimiento y medicion", project.Name())
	}
}

// Escenario: Rechazar una modificacion que deja el proyecto sin nombre
// Cubre: RN-2, RN-5
func TestRechazarUnaModificacionQueDejaElProyectoSinNombre(t *testing.T) {
	registry := NewProjectRegistry()
	project := fullProject(t, registry)

	err := registry.Update(project.ID(), ProjectChanges{Name: ptr("   ")})

	if !errors.Is(err, ErrProjectNameRequired) {
		t.Fatalf("se esperaba ErrProjectNameRequired, salio: %v", err)
	}
	assertProject(t, project, "Sistema de metricas", "TP de ICS", date(2026, 10, 1), date(2026, 11, 30))
}

// Escenario: Rechazar una fecha de inicio posterior a la de finalizacion al modificar
// Cubre: RN-4, RN-5
func TestRechazarUnaFechaDeInicioPosteriorALaDeFinalizacionAlModificar(t *testing.T) {
	registry := NewProjectRegistry()
	project := fullProject(t, registry)

	err := registry.Update(project.ID(), ProjectChanges{Start: ptr(date(2026, 12, 1))})

	if !errors.Is(err, ErrProjectDatesInconsistent) {
		t.Fatalf("se esperaba ErrProjectDatesInconsistent, salio: %v", err)
	}
	assertProject(t, project, "Sistema de metricas", "TP de ICS", date(2026, 10, 1), date(2026, 11, 30))
}

// Escenario: Quitar la fecha de finalizacion
// Cubre: RN-9
func TestQuitarLaFechaDeFinalizacion(t *testing.T) {
	registry := NewProjectRegistry()
	project := fullProject(t, registry)

	err := registry.Update(project.ID(), ProjectChanges{End: Remove[time.Time]()})

	if err != nil {
		t.Fatalf("se esperaba quitar la fecha de finalizacion sin error, salio: %v", err)
	}
	if project.HasEnd() {
		t.Errorf("se esperaba el proyecto sin fecha de finalizacion, salio: %v", project.End())
	}
	assertProject(t, project, "Sistema de metricas", "TP de ICS", date(2026, 10, 1), time.Time{})
}

// Escenario: Quitar la descripcion
// Cubre: RN-9
func TestQuitarLaDescripcion(t *testing.T) {
	registry := NewProjectRegistry()
	project := fullProject(t, registry)

	err := registry.Update(project.ID(), ProjectChanges{Description: Remove[string]()})

	if err != nil {
		t.Fatalf("se esperaba quitar la descripcion sin error, salio: %v", err)
	}
	assertProject(t, project, "Sistema de metricas", "", date(2026, 10, 1), date(2026, 11, 30))
}
