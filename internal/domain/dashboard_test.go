package domain

import "testing"

// Los tests del dominio del dashboard (historia #17) arman las métricas a mano, contra el
// contrato (D8): así cada test elige sus números sin depender del cálculo de la #16.

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
