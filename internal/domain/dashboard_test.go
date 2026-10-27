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

// Escenario: Dashboard de un proyecto con un sprint en curso
// Cubre: RN-8, RN-10
// Mira los tres primeros gráficos; el de velocidad, cuarto, es de la tarea 6.
func TestDashboardDeUnProyectoConUnSprintEnCurso(t *testing.T) {
	project := unProyecto(t)
	dashboardSprint(t, project, true)
	open := dashboardSprint(t, project, false)
	metrics := Metrics{
		PlannedPoints: 40, CompletedPoints: 30,
		EstimatedHours: 40, ActualHours: 45.25,
		DefectsDetected: 5, DefectsResolved: 3,
	}
	// Las del sprint abierto son otras, para ver que los gráficos usan las del proyecto.
	sprintMetrics := map[int]Metrics{
		open.Number(): {PlannedPoints: 13, EstimatedHours: 10, ActualHours: 6.5, DefectsDetected: 2},
	}

	dashboard := buildDashboard(dashboardSource{sprints: project.Sprints(), projectMetrics: metrics, sprintMetrics: sprintMetrics})

	// Value lleva el número tal como llegó, para medir la barra; Text, ya redondeado.
	want := []Chart{
		{Title: "Story points", Bars: []Bar{
			{Label: "Planificados", Value: 40, Text: "40"},
			{Label: "Completados", Value: 30, Text: "30"},
		}},
		{Title: "Horas", Bars: []Bar{
			{Label: "Estimadas", Value: 40, Text: "40,0 h"},
			{Label: "Reales", Value: 45.25, Text: "45,3 h"},
		}},
		{Title: "Defectos", Bars: []Bar{
			{Label: "Detectados", Value: 5, Text: "5"},
			{Label: "Resueltos", Value: 3, Text: "3"},
		}},
	}
	if len(dashboard.Charts) < len(want) {
		t.Fatalf("se esperaban al menos %d graficos, salieron %d: %+v", len(want), len(dashboard.Charts), dashboard.Charts)
	}
	if got := dashboard.Charts[:len(want)]; !reflect.DeepEqual(got, want) {
		t.Fatalf("se esperaban los graficos %+v, salieron %+v", want, got)
	}
}

// Escenario: El grafico de velocidad tiene una barra por sprint cerrado
// Cubre: RN-8, RN-9
// Cuánto mide cada barra (61,5 %, 100 % y 76,9 %) es de la web, en la tarea 13.
func TestElGraficoDeVelocidadTieneUnaBarraPorSprintCerrado(t *testing.T) {
	project := unProyecto(t)
	sprintMetrics := map[int]Metrics{}
	for _, velocity := range []int{8, 13, 10} {
		sprint := dashboardSprint(t, project, true)
		sprintMetrics[sprint.Number()] = Metrics{Velocity: velocity}
	}

	dashboard := buildDashboard(dashboardSource{sprints: project.Sprints(), sprintMetrics: sprintMetrics})

	want := Chart{Title: "Velocidad", Bars: []Bar{
		{Label: "Sprint 1", Value: 8, Text: "8"},
		{Label: "Sprint 2", Value: 13, Text: "13"},
		{Label: "Sprint 3", Value: 10, Text: "10"},
	}}
	if len(dashboard.Charts) != 4 {
		t.Fatalf("se esperaban 4 graficos, salieron %d: %+v", len(dashboard.Charts), dashboard.Charts)
	}
	if got := dashboard.Charts[3]; !reflect.DeepEqual(got, want) {
		t.Fatalf("se esperaba el grafico de velocidad %+v, salio %+v", want, got)
	}
}

// Escenario: Todavia no hay sprints cerrados para el grafico de velocidad
// Cubre: RN-4, RN-9
func TestTodaviaNoHaySprintsCerradosParaElGraficoDeVelocidad(t *testing.T) {
	project := unProyecto(t)
	open := dashboardSprint(t, project, false)
	// El sprint abierto trae velocidad, para ver que no aparece como barra.
	sprintMetrics := map[int]Metrics{open.Number(): {Velocity: 5}}

	dashboard := buildDashboard(dashboardSource{sprints: project.Sprints(), sprintMetrics: sprintMetrics})

	if len(dashboard.Charts) != 4 {
		t.Fatalf("se esperaban 4 graficos, salieron %d: %+v", len(dashboard.Charts), dashboard.Charts)
	}
	velocity := dashboard.Charts[3]
	if velocity.Title != "Velocidad" || len(velocity.Bars) != 0 || velocity.Empty != "Todavía no hay sprints cerrados" {
		t.Fatalf("se esperaba el grafico de velocidad sin barras y con el aviso, salio %+v", velocity)
	}
	if got, _ := figureText(dashboard.Overview, "Velocidad"); got != "sin sprints cerrados" {
		t.Fatalf("se esperaba la velocidad del estado general como %q, salio %q", "sin sprints cerrados", got)
	}

	// Un caso límite que el escenario no muestra: con sprints cerrados, una velocidad promedio
	// de cero se escribe "0" y no el aviso. El aviso depende de que no haya sprints cerrados,
	// no de que el valor sea cero (RN-4).
	t.Run("con sprints cerrados y velocidad cero", func(t *testing.T) {
		project := unProyecto(t)
		dashboardSprint(t, project, true)

		dashboard := buildDashboard(dashboardSource{sprints: project.Sprints()})

		if got, _ := figureText(dashboard.Overview, "Velocidad"); got != "0" {
			t.Fatalf("se esperaba la velocidad del estado general como %q, salio %q", "0", got)
		}
		if len(dashboard.Charts) != 4 || dashboard.Charts[3].Empty != "" {
			t.Fatalf("se esperaba el grafico de velocidad sin el aviso, salio %+v", dashboard.Charts)
		}
	})
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
