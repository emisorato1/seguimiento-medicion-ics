package domain

import "testing"

// Sin escenario: consulta que necesita la #77 para su página de inicio (Fase 2, Dependencias)
// Cubre: dependencia de la #77, RN-7 y RN-23 (la lista vacía no es un error)
func TestListarLosProyectosDeUnRegistroVacio(t *testing.T) {
	registry := NewProjectRegistry()

	projects := registry.Projects()

	if len(projects) != 0 {
		t.Fatalf("se esperaba una lista vacia, salieron %d proyectos", len(projects))
	}
}

// Sin escenario: consulta que necesita la #77 para su página de inicio (Fase 2, Dependencias)
// Cubre: dependencia de la #77
//
// Son doce proyectos y no dos o tres: con pocos, el orden en que Go recorre un mapa puede
// coincidir por casualidad con el de los números y el test pasaría sin ordenar.
func TestListarTodosLosProyectosEnOrdenDeNumero(t *testing.T) {
	registry := NewProjectRegistry()
	for i := 1; i <= 12; i++ {
		if _, err := registry.Create(ProjectData{Name: "Proyecto", Start: unDia(2026, 10, 1)}); err != nil {
			t.Fatalf("no se pudo crear el proyecto de prueba %d: %v", i, err)
		}
	}

	projects := registry.Projects()

	if len(projects) != 12 {
		t.Fatalf("se esperaban 12 proyectos, salieron %d", len(projects))
	}
	for i, project := range projects {
		if project.ID() != i+1 {
			t.Fatalf("se esperaba el proyecto %d en la posicion %d, salio el %d", i+1, i, project.ID())
		}
		saved, _ := registry.Get(i + 1)
		if project != saved {
			t.Fatalf("se esperaba recibir el proyecto %d guardado, no una copia", i+1)
		}
	}
}
