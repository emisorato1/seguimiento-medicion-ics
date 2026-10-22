package domain

import (
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
