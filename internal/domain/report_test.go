package domain

import (
	"slices"
	"testing"
)

// unProyectoLlamado crea un proyecto con ese nombre, para los tests que miran el titulo.
func unProyectoLlamado(t *testing.T, name string) *Project {
	t.Helper()
	project, err := NewProjectRegistry().Create(ProjectData{Name: name, Start: unDia(2026, 9, 1)})
	if err != nil {
		t.Fatalf("no se pudo crear el proyecto de prueba %q: %v", name, err)
	}
	return project
}

// encabezados devuelve los encabezados de las secciones del reporte, en orden.
func encabezados(r Report) []string {
	headings := []string{}
	for _, section := range r.Sections {
		headings = append(headings, section.Heading)
	}
	return headings
}

// Escenario: Generar el reporte de un sprint cerrado
// Cubre: RN-1, RN-3, RN-4
func TestGenerarElReporteDeUnSprintCerrado(t *testing.T) {
	project := unProyectoLlamado(t, "Seguimiento ICS")
	sprint := unSprintCerrado(t, project, "Tener el MVP", unDia(2026, 10, 20), unDia(2026, 10, 25), unDia(2026, 10, 25), nil)
	today := unDia(2026, 10, 27)

	report := project.buildSprintReport(sprint, Metrics{}, today)

	if report.Title != "Reporte del sprint 1 — Seguimiento ICS" {
		t.Errorf("se esperaba el titulo %q, salio %q", "Reporte del sprint 1 — Seguimiento ICS", report.Title)
	}
	if report.Scope != ScopeSprint {
		t.Errorf("se esperaba un reporte de sprint, salio el alcance %v", report.Scope)
	}
	if !report.GeneratedOn.Equal(today) {
		t.Errorf("se esperaba la fecha de generacion %v, salio %v", today, report.GeneratedOn)
	}
	want := []string{"Resumen", "Historias", "Estimación y esfuerzo", "Defectos"}
	if got := encabezados(report); !slices.Equal(got, want) {
		t.Errorf("se esperaban las secciones %q en ese orden, salieron %q", want, got)
	}
}
