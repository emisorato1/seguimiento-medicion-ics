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

// todasLasLineas devuelve las lineas de todas las secciones del reporte, en orden.
func todasLasLineas(r Report) []string {
	lines := []string{}
	for _, section := range r.Sections {
		lines = append(lines, section.Lines...)
	}
	return lines
}

// Esquema del escenario: Formato de los numeros
// Cubre: RN-14
func TestFormatoDeLosNumeros(t *testing.T) {
	cases := []struct {
		metrica string
		metrics Metrics
		linea   string
	}{
		{"horas reales", Metrics{ActualHours: 7.5}, "Horas reales: 7,5 h"},
		{"desviacion 25", Metrics{HoursDeviation: 25}, "Desviación: 25,0 %"},
		{"desviacion -20", Metrics{HoursDeviation: -20}, "Desviación: -20,0 %"},
		{"desviacion 0", Metrics{HoursDeviation: 0}, "Desviación: 0,0 %"},
		{"porcentaje con 2 completadas", Metrics{StoriesPlanned: 3, StoriesCompleted: 2, CompletionRate: 200.0 / 3}, "Historias completadas: 2 (66,7 %)"},
	}
	for _, c := range cases {
		t.Run(c.metrica, func(t *testing.T) {
			project := unProyecto(t)
			sprint := unSprintCerrado(t, project, "Tener el MVP", unDia(2026, 10, 20), unDia(2026, 10, 25), unDia(2026, 10, 25), nil)

			report := project.buildSprintReport(sprint, c.metrics, unDia(2026, 10, 27))

			if lines := todasLasLineas(report); !slices.Contains(lines, c.linea) {
				t.Errorf("se esperaba la linea %q, y no esta en:\n%q", c.linea, lines)
			}
		})
	}
}

// Esquema del escenario: El resumen dice en que estado esta el sprint
// Cubre: RN-6, RN-7, RN-14
func TestElResumenDiceEnQueEstadoEstaElSprint(t *testing.T) {
	cases := []struct {
		estado string
		closed bool
		linea  string
	}{
		{"cerrado el 25/10/2026", true, "Estado: cerrado el 25/10/2026"},
		{"abierto", false, "Estado: abierto. Los números son parciales."},
	}
	for _, c := range cases {
		t.Run(c.estado, func(t *testing.T) {
			project := unProyectoLlamado(t, "Seguimiento ICS")
			sprint, err := project.AddSprint(SprintData{Goal: "Tener el MVP", Start: unDia(2026, 10, 1), End: unDia(2026, 10, 25)})
			if err != nil {
				t.Fatalf("no se pudo crear el sprint de prueba: %v", err)
			}
			if c.closed {
				if err := project.CloseSprint(nil, unDia(2026, 10, 25)); err != nil {
					t.Fatalf("no se pudo cerrar el sprint de prueba: %v", err)
				}
			}

			report := project.buildSprintReport(sprint, Metrics{}, unDia(2026, 10, 27))

			want := []string{
				"Proyecto: Seguimiento ICS",
				"Sprint Goal: Tener el MVP",
				"Inicio: 01/10/2026",
				"Fin: 25/10/2026",
				c.linea,
			}
			if got := seccion(t, report, "Resumen"); !slices.Equal(got, want) {
				t.Errorf("se esperaba el resumen\n%q\nsalio\n%q", want, got)
			}
		})
	}
}

// Escenario: Las historias del sprint figuran segun el cierre de ese sprint
// Cubre: RN-9, RN-10
func TestLasHistoriasDelSprintFiguranSegunElCierreDeEseSprint(t *testing.T) {
	project := unProyecto(t)
	first := unSprintCerrado(t, project, "Tener el MVP", unDia(2026, 10, 12), unDia(2026, 10, 19), unDia(2026, 10, 19), []int{1}, 3, 5)
	if _, err := project.AddSprint(SprintData{Goal: "Terminar lo pendiente", Start: unDia(2026, 10, 20), End: unDia(2026, 10, 25)}); err != nil {
		t.Fatalf("no se pudo crear el sprint 2 de prueba: %v", err)
	}
	if err := project.AssignStory(2); err != nil {
		t.Fatalf("no se pudo asignar la historia 2 al sprint 2: %v", err)
	}
	if err := project.CloseSprint([]int{2}, unDia(2026, 10, 25)); err != nil {
		t.Fatalf("no se pudo cerrar el sprint 2: %v", err)
	}
	if estadoDe(t, project, 2) != StoryDone {
		t.Fatalf("se esperaba la historia 2 Terminada despues del sprint 2")
	}
	metrics, err := project.SprintMetrics(1)
	if err != nil {
		t.Fatalf("no se pudieron calcular las metricas del sprint 1: %v", err)
	}

	report := project.buildSprintReport(first, metrics, unDia(2026, 10, 27))

	want := []string{
		"Historias planificadas: 2",
		"Historias completadas: 1 (50,0 %)",
		"#1 — Registrar esfuerzo — 3 SP — completada",
		"#2 — Registrar esfuerzo — 5 SP — no completada",
	}
	if got := seccion(t, report, "Historias"); !slices.Equal(got, want) {
		t.Errorf("se esperaba la seccion Historias\n%q\nsalio\n%q", want, got)
	}
}

// Escenario: Una historia sin estimar figura como sin estimar
// Cubre: RN-9
//
// La historia 3 entra al sprint estimada (la #9 no asigna historias sin estimar) y se le quita
// la estimacion mientras esta En curso, que la #12 permite; al cerrar, la #10 le registra 0
// puntos. La historia 1, estimada, sigue mostrando sus puntos.
func TestUnaHistoriaSinEstimarFiguraComoSinEstimar(t *testing.T) {
	project := unProyecto(t)
	sprint := unSprint(t, project)
	for _, points := range []int{3, 5, 8} {
		story := unaHistoriaEstimada(t, project, points)
		if err := project.AssignStory(story.ID()); err != nil {
			t.Fatalf("no se pudo asignar la historia de prueba %d: %v", story.ID(), err)
		}
	}
	if err := project.RemoveStoryEstimate(3); err != nil {
		t.Fatalf("no se pudo quitar la estimacion de la historia 3: %v", err)
	}
	if err := project.CloseSprint(nil, unDia(2026, 10, 30)); err != nil {
		t.Fatalf("no se pudo cerrar el sprint de prueba: %v", err)
	}

	report := project.buildSprintReport(sprint, Metrics{}, unDia(2026, 10, 31))

	lines := seccion(t, report, "Historias")
	for _, want := range []string{
		"#1 — Registrar esfuerzo — 3 SP — no completada",
		"#3 — Registrar esfuerzo — sin estimar — no completada",
	} {
		if !slices.Contains(lines, want) {
			t.Errorf("se esperaba la linea %q, y no esta en:\n%q", want, lines)
		}
	}
}

// Escenario: Las historias de un sprint abierto figuran en curso
// Cubre: RN-10
func TestLasHistoriasDeUnSprintAbiertoFiguranEnCurso(t *testing.T) {
	project := unProyecto(t)
	sprint := unSprint(t, project)
	story := unaHistoriaEstimada(t, project, 3)
	if err := project.AssignStory(story.ID()); err != nil {
		t.Fatalf("no se pudo asignar la historia de prueba: %v", err)
	}

	report := project.buildSprintReport(sprint, Metrics{}, unDia(2026, 10, 27))

	want := "#1 — Registrar esfuerzo — 3 SP — en curso"
	if lines := seccion(t, report, "Historias"); !slices.Contains(lines, want) {
		t.Errorf("se esperaba la linea %q, y no esta en:\n%q", want, lines)
	}
}
