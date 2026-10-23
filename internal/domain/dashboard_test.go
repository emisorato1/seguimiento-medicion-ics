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
