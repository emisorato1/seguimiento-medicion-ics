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

// Escenario: Los defectos del sprint son los detectados o resueltos en el
// Cubre: RN-12
func TestLosDefectosDelSprintSonLosDetectadosOResueltosEnEl(t *testing.T) {
	project := unProyectoConDefectos(t, 3)
	for _, detectedIn := range []int{1, 2, 3} {
		data := unDefecto()
		data.DetectedIn = detectedIn
		if _, err := project.AddDefect(data); err != nil {
			t.Fatalf("no se pudo registrar el defecto detectado en el sprint %d: %v", detectedIn, err)
		}
	}
	if err := project.ResolveDefect(1, 2); err != nil {
		t.Fatalf("no se pudo resolver el defecto 1 en el sprint 2: %v", err)
	}
	sprint, err := project.Sprint(2)
	if err != nil {
		t.Fatalf("no se encontro el sprint 2: %v", err)
	}
	metrics, err := project.SprintMetrics(2)
	if err != nil {
		t.Fatalf("no se pudieron calcular las metricas del sprint 2: %v", err)
	}

	report := project.buildSprintReport(sprint, metrics, unDia(2026, 10, 27))

	want := []string{
		"Defectos detectados: 1",
		"Defectos resueltos: 1",
		"#1 — La velocidad sale en cero — Mayor — Resuelto en el sprint 2",
		"#2 — La velocidad sale en cero — Mayor — Abierto",
	}
	if got := seccion(t, report, "Defectos"); !slices.Equal(got, want) {
		t.Errorf("se esperaba la seccion Defectos\n%q\nsalio\n%q", want, got)
	}
}

// Esquema del escenario: Una seccion sin datos lo dice
// Cubre: RN-5
func TestUnaSeccionSinDatosLoDice(t *testing.T) {
	cases := []struct {
		que     string
		seccion string
		want    []string
	}{
		{"historias", "Historias", []string{"Historias planificadas: 0", "Historias completadas: 0 (0,0 %)", "No hay historias en este sprint."}},
		{"defectos", "Defectos", []string{"Defectos detectados: 0", "Defectos resueltos: 0", "No hay defectos en este sprint."}},
	}
	for _, c := range cases {
		t.Run(c.que, func(t *testing.T) {
			project := unProyecto(t)
			sprint := unSprintCerrado(t, project, "Tener el MVP", unDia(2026, 10, 20), unDia(2026, 10, 25), unDia(2026, 10, 25), nil)
			metrics, err := project.SprintMetrics(1)
			if err != nil {
				t.Fatalf("no se pudieron calcular las metricas del sprint 1: %v", err)
			}

			report := project.buildSprintReport(sprint, metrics, unDia(2026, 10, 27))

			if got := seccion(t, report, c.seccion); !slices.Equal(got, c.want) {
				t.Errorf("se esperaba la seccion %s\n%q\nsalio\n%q", c.seccion, c.want, got)
			}
		})
	}
}

// Escenario: Generar el reporte del proyecto entero
// Cubre: RN-1, RN-3, RN-8, RN-10, RN-11, RN-12
func TestGenerarElReporteDelProyectoEntero(t *testing.T) {
	project, err := NewProjectRegistry().Create(ProjectData{Name: "Seguimiento ICS", Start: unDia(2026, 9, 1), End: unDia(2026, 11, 2)})
	if err != nil {
		t.Fatalf("no se pudo crear el proyecto de prueba: %v", err)
	}
	unSprintCerrado(t, project, "Tener el MVP", unDia(2026, 10, 5), unDia(2026, 10, 11), unDia(2026, 10, 11), []int{1}, 3)
	unSprintCerrado(t, project, "Tener los sprints", unDia(2026, 10, 12), unDia(2026, 10, 18), unDia(2026, 10, 18), []int{2}, 5, 2)
	unSprint(t, project)
	unaHistoriaEstimada(t, project, 8)
	data := unDefecto()
	if _, err := project.AddDefect(data); err != nil {
		t.Fatalf("no se pudo registrar el defecto 1: %v", err)
	}
	data.DetectedIn = 2
	if _, err := project.AddDefect(data); err != nil {
		t.Fatalf("no se pudo registrar el defecto 2: %v", err)
	}
	if err := project.ResolveDefect(1, 2); err != nil {
		t.Fatalf("no se pudo resolver el defecto 1: %v", err)
	}
	metrics := project.ProjectMetrics()
	today := unDia(2026, 10, 27)

	report := project.buildProjectReport(metrics, today)

	if report.Title != "Reporte del proyecto Seguimiento ICS" || report.Scope != ScopeProject || !report.GeneratedOn.Equal(today) {
		t.Errorf("se esperaba el reporte del proyecto Seguimiento ICS del %v, salio %q, alcance %v, del %v", today, report.Title, report.Scope, report.GeneratedOn)
	}
	if report.Metrics != metrics {
		t.Errorf("se esperaban las metricas del proyecto %+v, salieron %+v", metrics, report.Metrics)
	}
	want := map[string][]string{
		"Resumen": {
			"Proyecto: Seguimiento ICS",
			"Inicio: 01/09/2026",
			"Fin: 02/11/2026",
			"3 sprints, 2 cerrados",
		},
		"Historias": {
			"Historias planificadas: " + formatCount(metrics.StoriesPlanned),
			"Historias completadas: " + formatCount(metrics.StoriesCompleted) + " (" + formatPercent(metrics.CompletionRate) + ")",
			"#1 — Registrar esfuerzo — 3 SP — Terminada",
			"#2 — Registrar esfuerzo — 5 SP — Terminada",
			"#3 — Registrar esfuerzo — 2 SP — Pendiente",
			"#4 — Registrar esfuerzo — 8 SP — Pendiente",
		},
		"Defectos": {
			"Defectos detectados: 2",
			"Defectos resueltos: 1",
			"#1 — La velocidad sale en cero — Mayor — Resuelto en el sprint 2",
			"#2 — La velocidad sale en cero — Mayor — Abierto",
		},
	}
	for _, heading := range []string{"Resumen", "Historias", "Defectos"} {
		if got := seccion(t, report, heading); !slices.Equal(got, want[heading]) {
			t.Errorf("se esperaba la seccion %s\n%q\nsalio\n%q", heading, want[heading], got)
		}
	}
	velocity := "Velocidad promedio: " + formatCount(metrics.Velocity)
	if lines := seccion(t, report, "Estimación y esfuerzo"); !slices.Contains(lines, velocity) {
		t.Errorf("se esperaba la linea %q, y no esta en:\n%q", velocity, lines)
	}
}

// Escenario: Un proyecto sin fecha de fin
// Cubre: RN-8
func TestUnProyectoSinFechaDeFin(t *testing.T) {
	project := unProyectoLlamado(t, "Seguimiento ICS")

	report := project.buildProjectReport(project.ProjectMetrics(), unDia(2026, 10, 27))

	if lines := seccion(t, report, "Resumen"); !slices.Contains(lines, "Fin: sin fecha de fin") {
		t.Errorf("se esperaba la linea %q, y no esta en:\n%q", "Fin: sin fecha de fin", lines)
	}
}

// Escenario: Un proyecto sin sprints cerrados no muestra una velocidad de cero
// Cubre: RN-8, RN-11
func TestUnProyectoSinSprintsCerradosNoMuestraUnaVelocidadDeCero(t *testing.T) {
	project := unProyectoLlamado(t, "Seguimiento ICS")
	unSprint(t, project)

	report := project.buildProjectReport(project.ProjectMetrics(), unDia(2026, 10, 27))

	if lines := seccion(t, report, "Resumen"); !slices.Contains(lines, "1 sprint, 0 cerrados") {
		t.Errorf("se esperaba la linea %q, y no esta en:\n%q", "1 sprint, 0 cerrados", lines)
	}
	velocity := "Velocidad promedio: sin sprints cerrados"
	if lines := seccion(t, report, "Estimación y esfuerzo"); !slices.Contains(lines, velocity) {
		t.Errorf("se esperaba la linea %q, y no esta en:\n%q", velocity, lines)
	}
}
