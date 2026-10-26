package domain

import (
	"testing"
)

// Esquema del escenario: Un proyecto sin sprints cerrados devuelve un listado vacio
// Cubre: RN-3
func TestUnProyectoSinSprintsCerradosDevuelveUnListadoVacio(t *testing.T) {
	cases := []struct {
		situacion string
		prepare   func(t *testing.T, project *Project)
	}{
		{"sin ningun sprint", func(t *testing.T, project *Project) {}},
		{"con un sprint abierto", func(t *testing.T, project *Project) { unSprint(t, project) }},
	}
	for _, c := range cases {
		t.Run(c.situacion, func(t *testing.T) {
			project := unProyecto(t)
			c.prepare(t, project)

			sprints := project.ClosedSprints()

			if sprints == nil {
				t.Fatalf("se esperaba una lista vacia, no nil")
			}
			if len(sprints) != 0 {
				t.Fatalf("se esperaba el listado vacio, salieron %d sprints", len(sprints))
			}
		})
	}
}
