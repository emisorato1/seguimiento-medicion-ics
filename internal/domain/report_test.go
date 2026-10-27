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

// seccion devuelve las lineas de la seccion con ese encabezado, o corta el test si el
// reporte no la tiene.
func seccion(t *testing.T, r Report, heading string) []string {
	t.Helper()
	for _, section := range r.Sections {
		if section.Heading == heading {
			return section.Lines
		}
	}
	t.Fatalf("el reporte no tiene la seccion %q; tiene %q", heading, encabezados(r))
	return nil
}

// Escenario: Los numeros del reporte son los de las metricas
// Cubre: RN-2, RN-11
func TestLosNumerosDelReporteSonLosDeLasMetricas(t *testing.T) {
	project := unProyecto(t)
	sprint := unSprintCerrado(t, project, "Tener el MVP", unDia(2026, 10, 20), unDia(2026, 10, 25), unDia(2026, 10, 25), nil)
	metrics := Metrics{PlannedPoints: 13, CompletedPoints: 10, Velocity: 10, EstimatedHours: 16, ActualHours: 20, HoursDeviation: 25}

	report := project.buildSprintReport(sprint, metrics, unDia(2026, 10, 27))

	if report.Metrics != metrics {
		t.Errorf("se esperaban las metricas %+v, salieron %+v", metrics, report.Metrics)
	}
	want := []string{
		"Story points planificados: 13",
		"Story points completados: 10",
		"Velocidad: 10",
		"Horas estimadas: 16,0 h",
		"Horas reales: 20,0 h",
		"Desviación: 25,0 %",
	}
	if got := seccion(t, report, "Estimación y esfuerzo"); !slices.Equal(got, want) {
		t.Errorf("se esperaban las lineas\n%q\nsalieron\n%q", want, got)
	}
}
