package domain

import (
	"reflect"
	"testing"
)

// Los tests del dominio del dashboard (historia #17) arman las métricas a mano, contra el
// contrato (D8): así cada test elige sus números sin depender del cálculo de la #16.

// dashboardSprint crea el sprint siguiente del proyecto con el alta de la #8. Cerrarlo es de
// la #10: se hace desde dentro del paquete, como en la #5.
func dashboardSprint(t *testing.T, project *Project, closed bool) *Sprint {
	t.Helper()
	sprint, err := project.AddSprint(SprintData{Goal: "Tener el dashboard", Start: date(2026, 10, 26), End: date(2026, 10, 30)})
	if err != nil {
		t.Fatalf("no se pudo crear el sprint de prueba: %v", err)
	}
	if closed {
		sprint.state = SprintClosed
	}
	return sprint
}

// Escenario: Un proyecto sin sprints avisa que no hay datos
// Cubre: RN-7
func TestUnProyectoSinSprintsAvisaQueNoHayDatos(t *testing.T) {
	project := unProyecto(t)

	dashboard := buildDashboard(dashboardSource{sprints: project.Sprints()})

	if dashboard.HasSprints {
		t.Fatalf("se esperaba un dashboard sin sprints, salio HasSprints en true")
	}
	if len(dashboard.Overview) != 0 {
		t.Fatalf("se esperaba un dashboard sin estado general, salio: %v", dashboard.Overview)
	}
	if dashboard.CurrentSprint != nil {
		t.Fatalf("se esperaba un dashboard sin sprint en curso, salio: %+v", dashboard.CurrentSprint)
	}
	if len(dashboard.Charts) != 0 {
		t.Fatalf("se esperaba un dashboard sin graficos, salio: %v", dashboard.Charts)
	}
}

// Escenario: El estado general muestra las metricas tal como se calcularon
// Cubre: RN-2, RN-4
func TestElEstadoGeneralMuestraLasMetricasTalComoSeCalcularon(t *testing.T) {
	project := unProyecto(t)
	dashboardSprint(t, project, true)
	metrics := Metrics{
		PlannedPoints: 40, CompletedPoints: 30,
		StoriesPlanned: 10, StoriesCompleted: 7, CompletionRate: 70,
		EstimatedHours: 40, ActualHours: 45, HoursDeviation: 12.5,
		DefectsDetected: 5, DefectsResolved: 3,
		Velocity: 12,
	}

	dashboard := buildDashboard(dashboardSource{sprints: project.Sprints(), projectMetrics: metrics})

	// Este test mira qué número va en cada lugar; cómo se escribe cada uno es de la tarea 3.
	want := []Figure{
		{Label: "Story points planificados", Text: formatCount(40)},
		{Label: "Story points completados", Text: formatCount(30)},
		{Label: "Historias planificadas", Text: formatCount(10)},
		{Label: "Historias completadas", Text: formatCount(7)},
		{Label: "Porcentaje de historias completas", Text: formatPercent(70)},
		{Label: "Horas estimadas", Text: formatHours(40)},
		{Label: "Horas reales", Text: formatHours(45)},
		{Label: "Desviación de horas", Text: formatPercent(12.5)},
		{Label: "Defectos detectados", Text: formatCount(5)},
		{Label: "Defectos resueltos", Text: formatCount(3)},
		{Label: "Velocidad", Text: formatCount(12)},
	}
	if !reflect.DeepEqual(dashboard.Overview, want) {
		t.Fatalf("se esperaba el estado general %v, salio %v", want, dashboard.Overview)
	}
}

// Escenario: Porcentajes, horas y enteros se muestran con su formato
// Cubre: RN-15
func TestPorcentajesHorasYEnterosSeMuestranConSuFormato(t *testing.T) {
	cases := []struct {
		name  string
		label string
		set   func(m *Metrics)
		want  string
	}{
		{"porcentaje con decimales", "Porcentaje de historias completas", func(m *Metrics) { m.CompletionRate = 66.6666 }, "66,7 %"},
		{"porcentaje entero", "Porcentaje de historias completas", func(m *Metrics) { m.CompletionRate = 100 }, "100,0 %"},
		{"horas con decimales", "Horas reales", func(m *Metrics) { m.ActualHours = 12.5 }, "12,5 h"},
		{"horas enteras", "Horas estimadas", func(m *Metrics) { m.EstimatedHours = 40 }, "40,0 h"},
		{"desviacion negativa", "Desviación de horas", func(m *Metrics) { m.HoursDeviation = -12.5 }, "-12,5 %"},
		{"desviacion cero", "Desviación de horas", func(m *Metrics) { m.HoursDeviation = 0 }, "0,0 %"},
		{"desviacion positiva", "Desviación de horas", func(m *Metrics) { m.HoursDeviation = 25 }, "25,0 %"},
		{"velocidad", "Velocidad", func(m *Metrics) { m.Velocity = 12 }, "12"},
		{"story points completados", "Story points completados", func(m *Metrics) { m.CompletedPoints = 30 }, "30"},
		// Dos bordes que la RN-15 implica y el escenario no muestra: una desviación negativa
		// que se redondea a cero se escribe como cero, sin signo; y el empate se redondea
		// hacia arriba, como se redondea a mano (un cuarto de hora).
		{"desviacion negativa que redondea a cero", "Desviación de horas", func(m *Metrics) { m.HoursDeviation = -0.04 }, "0,0 %"},
		{"horas que empatan al redondear", "Horas reales", func(m *Metrics) { m.ActualHours = 12.25 }, "12,3 h"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			project := unProyecto(t)
			dashboardSprint(t, project, true)
			var metrics Metrics
			c.set(&metrics)

			dashboard := buildDashboard(dashboardSource{sprints: project.Sprints(), projectMetrics: metrics})

			got, ok := figureText(dashboard.Overview, c.label)
			if !ok {
				t.Fatalf("el estado general no tiene %q: %v", c.label, dashboard.Overview)
			}
			if got != c.want {
				t.Fatalf("se esperaba %q como %q, salio %q", c.label, c.want, got)
			}
		})
	}
}

// Escenario: El sprint en curso muestra su meta y lo que lleva registrado
// Cubre: RN-5
func TestElSprintEnCursoMuestraSuMetaYLoQueLlevaRegistrado(t *testing.T) {
	project := unProyecto(t)
	dashboardSprint(t, project, true)
	open := dashboardSprint(t, project, false)
	// Las métricas de un sprint abierto traen los completados y la velocidad en cero (#16):
	// se cargan igual con otro valor para ver que el sprint en curso no los muestra.
	sprintMetrics := map[int]Metrics{
		open.Number(): {PlannedPoints: 13, CompletedPoints: 5, ActualHours: 6.5, DefectsDetected: 2, Velocity: 5},
	}

	dashboard := buildDashboard(dashboardSource{sprints: project.Sprints(), sprintMetrics: sprintMetrics})

	want := &CurrentSprint{
		Number: 2,
		Goal:   "Tener el dashboard",
		Start:  date(2026, 10, 26),
		End:    date(2026, 10, 30),
		Figures: []Figure{
			{Label: "Story points planificados", Text: "13"},
			{Label: "Horas reales", Text: "6,5 h"},
			{Label: "Defectos detectados", Text: "2"},
		},
	}
	if !reflect.DeepEqual(dashboard.CurrentSprint, want) {
		t.Fatalf("se esperaba el sprint en curso %+v, salio %+v", want, dashboard.CurrentSprint)
	}
}

// Escenario: Un proyecto con todos sus sprints cerrados no tiene sprint en curso
// Cubre: RN-6
// El escenario se llama igual que uno de la #5 (project_status_test.go): el sufijo evita que
// los dos tests choquen en el mismo paquete.
func TestUnProyectoConTodosSusSprintsCerradosNoTieneSprintEnCursoEnElDashboard(t *testing.T) {
	project := unProyecto(t)
	dashboardSprint(t, project, true)
	dashboardSprint(t, project, true)

	dashboard := buildDashboard(dashboardSource{sprints: project.Sprints(), projectMetrics: Metrics{PlannedPoints: 21}})

	if dashboard.CurrentSprint != nil {
		t.Fatalf("se esperaba un dashboard sin sprint en curso, salio: %+v", dashboard.CurrentSprint)
	}
	if got, ok := figureText(dashboard.Overview, "Story points planificados"); !ok || got != "21" {
		t.Fatalf("se esperaba el estado general igual, con 21 story points planificados, salio: %v", dashboard.Overview)
	}
}

// figureText devuelve el valor escrito del número con ese rótulo, y si está.
func figureText(figures []Figure, label string) (string, bool) {
	for _, figure := range figures {
		if figure.Label == label {
			return figure.Text, true
		}
	}
	return "", false
}
