package domain

import (
	"errors"
	"testing"
)

// conSprints crea en el proyecto los sprints del 1 al n, con el alta de la #8.
func conSprints(t *testing.T, project *Project, n int) {
	t.Helper()
	for i := 1; i <= n; i++ {
		if _, err := project.AddSprint(SprintData{Goal: "Sprint de prueba", Start: unDia(2026, 10, 1), End: unDia(2026, 10, 5)}); err != nil {
			t.Fatalf("no se pudo crear el sprint de prueba %d: %v", i, err)
		}
	}
}

// unProyectoConDefectos arma un proyecto con la historia 1 y los sprints del 1 al n.
func unProyectoConDefectos(t *testing.T, sprints int) *Project {
	t.Helper()
	project := unProyecto(t)
	if _, err := project.AddStory(unosDatos()); err != nil {
		t.Fatalf("no se pudo crear la historia de prueba: %v", err)
	}
	conSprints(t, project, sprints)
	return project
}

// unDefecto devuelve los datos de un defecto válido: Mayor, sobre la historia 1 y detectado
// en el sprint 1. Cada test cambia solo lo que prueba.
func unDefecto() DefectData {
	return DefectData{
		Title:       "La velocidad sale en cero",
		Description: "El resumen del sprint muestra 0 puntos",
		Severity:    SeverityMajor,
		StoryID:     1,
		DetectedIn:  1,
	}
}

// Escenario: Registrar un defecto
// Cubre: RN-1, RN-2, RN-3
func TestRegistrarUnDefecto(t *testing.T) {
	project := unProyectoConDefectos(t, 1)

	defect, err := project.AddDefect(unDefecto())

	if err != nil {
		t.Fatalf("se esperaba registrar el defecto sin error, salio: %v", err)
	}
	if defect.ID() != 1 || defect.State() != DefectOpen || defect.ResolvedIn() != 0 {
		t.Fatalf("se esperaba el defecto 1, Abierto y sin sprint de resolucion, salio %d, %v, %d", defect.ID(), defect.State(), defect.ResolvedIn())
	}
	if defect.Title() != "La velocidad sale en cero" || defect.Description() != "El resumen del sprint muestra 0 puntos" ||
		defect.Severity() != SeverityMajor || defect.StoryID() != 1 || defect.DetectedIn() != 1 {
		t.Fatalf("se esperaban los datos registrados, salio %q, %q, %v, historia %d, sprint %d",
			defect.Title(), defect.Description(), defect.Severity(), defect.StoryID(), defect.DetectedIn())
	}
	if defects := project.Defects(); len(defects) != 1 || defects[0] != defect {
		t.Fatalf("se esperaba que el proyecto tuviera ese unico defecto, tiene %d", len(defects))
	}
}

// Escenario: El defecto nuevo recibe el identificador siguiente
// Cubre: RN-1, RN-17
func TestElDefectoNuevoRecibeElIdentificadorSiguiente(t *testing.T) {
	project := unProyectoConDefectos(t, 1)
	if _, err := project.AddDefect(unDefecto()); err != nil {
		t.Fatalf("no se pudo registrar el primer defecto de prueba: %v", err)
	}

	defect, err := project.AddDefect(unDefecto())

	if err != nil || defect.ID() != 2 {
		t.Fatalf("se esperaba el defecto 2 con los mismos datos, salio %v (error: %v)", defect, err)
	}
	if len(project.Defects()) != 2 {
		t.Fatalf("se esperaba que el proyecto quedara con 2 defectos, tiene %d", len(project.Defects()))
	}
}

// Esquema del escenario: Rechazar un defecto sin titulo o sin descripcion
// Cubre: RN-4, RN-18
func TestRechazarUnDefectoSinTituloOSinDescripcion(t *testing.T) {
	cases := []struct {
		name   string
		change func(d *DefectData)
		want   error
	}{
		{"el titulo", func(d *DefectData) { d.Title = "   " }, ErrDefectTitleRequired},
		{"la descripcion", func(d *DefectData) { d.Description = "" }, ErrDefectDescriptionRequired},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			project := unProyectoConDefectos(t, 1)
			data := unDefecto()
			tc.change(&data)

			_, err := project.AddDefect(data)

			if !errors.Is(err, tc.want) {
				t.Fatalf("se esperaba %v, salio: %v", tc.want, err)
			}
			if len(project.Defects()) != 0 {
				t.Fatalf("se esperaba que el proyecto siguiera sin defectos, tiene %d", len(project.Defects()))
			}
		})
	}
}

// Escenario: Guardar el titulo y la descripcion sin los espacios de los extremos
// Cubre: RN-5
func TestGuardarElTituloYLaDescripcionSinLosEspaciosDeLosExtremos(t *testing.T) {
	project := unProyectoConDefectos(t, 1)
	data := unDefecto()
	data.Title = "  La velocidad sale en cero  "
	data.Description = "  El resumen muestra 0  "

	defect, err := project.AddDefect(data)

	if err != nil {
		t.Fatalf("se esperaba registrar el defecto sin error, salio: %v", err)
	}
	if defect.Title() != "La velocidad sale en cero" || defect.Description() != "El resumen muestra 0" {
		t.Fatalf("se esperaban el titulo y la descripcion sin espacios de mas, salio %q y %q", defect.Title(), defect.Description())
	}
}

// Esquema del escenario: Registrar un defecto con cada severidad
// Cubre: RN-6
func TestRegistrarUnDefectoConCadaSeveridad(t *testing.T) {
	cases := []struct {
		written string
		want    string
	}{
		{"Crítica", "Crítica"},
		{"critica", "Crítica"},
		{"  MAYOR  ", "Mayor"},
		{"Menor", "Menor"},
	}
	for _, tc := range cases {
		t.Run(tc.written, func(t *testing.T) {
			project := unProyectoConDefectos(t, 1)
			severity, err := ParseSeverity(tc.written)
			if err != nil {
				t.Fatalf("se esperaba aceptar la severidad %q, salio: %v", tc.written, err)
			}
			data := unDefecto()
			data.Severity = severity

			defect, err := project.AddDefect(data)

			if err != nil {
				t.Fatalf("se esperaba registrar el defecto sin error, salio: %v", err)
			}
			if defect.Severity().String() != tc.want {
				t.Fatalf("se esperaba la severidad %q, salio %q", tc.want, defect.Severity().String())
			}
		})
	}
}

// Escenario: Rechazar una severidad que no existe
// Cubre: RN-6, RN-18
//
// Además del texto del escenario, cubre lo que llega a AddDefect sin pasar por la conversión:
// una severidad sin indicar (el valor cero) y una armada a mano fuera de rango (RN-6).
func TestRechazarUnaSeveridadQueNoExiste(t *testing.T) {
	if _, err := ParseSeverity("Urgente"); !errors.Is(err, ErrDefectSeverityInvalid) {
		t.Fatalf("se esperaba ErrDefectSeverityInvalid para %q, salio: %v", "Urgente", err)
	}
	for _, severity := range []Severity{0, Severity(9)} {
		project := unProyectoConDefectos(t, 1)
		data := unDefecto()
		data.Severity = severity

		_, err := project.AddDefect(data)

		if !errors.Is(err, ErrDefectSeverityInvalid) {
			t.Fatalf("se esperaba ErrDefectSeverityInvalid con la severidad %d, salio: %v", severity, err)
		}
		if len(project.Defects()) != 0 {
			t.Fatalf("se esperaba que el proyecto siguiera sin defectos, tiene %d", len(project.Defects()))
		}
	}
}
