package domain

import (
	"testing"
)

// unRegistroCargado arma un registro con dos proyectos: el primero con descripción, fecha de
// fin, dos integrantes y tres historias (estimada, sin estimar y Terminada); el segundo sin
// los datos opcionales y vacío.
func unRegistroCargado(t *testing.T) *ProjectRegistry {
	t.Helper()
	registry := NewProjectRegistry()
	first, err := registry.Create(ProjectData{
		Name: "Sistema de metricas", Description: "Para la materia",
		Start: unDia(2026, 10, 1), End: unDia(2026, 11, 2),
	})
	if err != nil {
		t.Fatalf("no se pudo crear el primer proyecto de prueba: %v", err)
	}
	if _, err := registry.Create(ProjectData{Name: "Otro proyecto", Start: unDia(2026, 10, 5)}); err != nil {
		t.Fatalf("no se pudo crear el segundo proyecto de prueba: %v", err)
	}
	if _, err := first.AddMember("Ana Perez", RoleAgileEnabler); err != nil {
		t.Fatalf("no se pudo agregar el integrante de prueba: %v", err)
	}
	if _, err := first.AddMember("Juan Gomez", RoleProductBuilder); err != nil {
		t.Fatalf("no se pudo agregar el integrante de prueba: %v", err)
	}

	estimated, err := first.AddStory(StoryData{
		Title: "Registrar esfuerzo", Description: "Cargar horas", Priority: PriorityHigh,
		AcceptanceCriteria: []string{"Valida las horas", "Suma por historia"},
	})
	if err != nil {
		t.Fatalf("no se pudo crear la historia de prueba: %v", err)
	}
	if err := first.EstimateStory(estimated.ID(), 5); err != nil {
		t.Fatalf("no se pudo estimar la historia de prueba: %v", err)
	}
	if _, err := first.AddStory(StoryData{
		Title: "Gestionar defectos", Description: "Cargar defectos", Priority: PriorityMedium,
		AcceptanceCriteria: []string{"Tiene severidad"},
	}); err != nil {
		t.Fatalf("no se pudo crear la historia de prueba: %v", err)
	}
	done, err := first.AddStory(StoryData{
		Title: "Crear proyectos", Description: "Alta de proyectos", Priority: PriorityLow,
		AcceptanceCriteria: []string{"Pide el nombre"},
	})
	if err != nil {
		t.Fatalf("no se pudo crear la historia de prueba: %v", err)
	}
	// Terminar una historia es de la #10, que todavia no esta en dev: el test la prepara
	// desde dentro del paquete.
	done.storyPoints = 3
	done.state = StoryDone
	return registry
}

// mismoRegistro compara dos registros dato por dato: proyectos, integrantes e historias.
func mismoRegistro(t *testing.T, want, got *ProjectRegistry) {
	t.Helper()
	wantProjects, gotProjects := want.Projects(), got.Projects()
	if len(gotProjects) != len(wantProjects) {
		t.Fatalf("se esperaban %d proyectos, salieron %d", len(wantProjects), len(gotProjects))
	}
	for i, w := range wantProjects {
		g := gotProjects[i]
		if g.ID() != w.ID() || g.Name() != w.Name() || g.Description() != w.Description() ||
			!g.Start().Equal(w.Start()) || !g.End().Equal(w.End()) || g.HasEnd() != w.HasEnd() {
			t.Fatalf("se esperaba el proyecto %d %q (%q, %v a %v), salio %d %q (%q, %v a %v)",
				w.ID(), w.Name(), w.Description(), w.Start(), w.End(),
				g.ID(), g.Name(), g.Description(), g.Start(), g.End())
		}

		wantMembers, gotMembers := w.Members(), g.Members()
		if len(gotMembers) != len(wantMembers) {
			t.Fatalf("proyecto %d: se esperaban %d integrantes, salieron %d", w.ID(), len(wantMembers), len(gotMembers))
		}
		for j, wm := range wantMembers {
			if gotMembers[j].Name() != wm.Name() || gotMembers[j].Role() != wm.Role() {
				t.Fatalf("proyecto %d: se esperaba el integrante %q (%v), salio %q (%v)",
					w.ID(), wm.Name(), wm.Role(), gotMembers[j].Name(), gotMembers[j].Role())
			}
		}

		wantStories, gotStories := w.Backlog(), g.Backlog()
		if len(gotStories) != len(wantStories) {
			t.Fatalf("proyecto %d: se esperaban %d historias, salieron %d", w.ID(), len(wantStories), len(gotStories))
		}
		for j, ws := range wantStories {
			gs := gotStories[j]
			if gs.ID() != ws.ID() || gs.Title() != ws.Title() || gs.Description() != ws.Description() ||
				gs.Priority() != ws.Priority() || gs.State() != ws.State() || gs.StoryPoints() != ws.StoryPoints() {
				t.Fatalf("proyecto %d: se esperaba la historia %d %q (%v, %v, %d SP), salio %d %q (%v, %v, %d SP)",
					w.ID(), ws.ID(), ws.Title(), ws.Priority(), ws.State(), ws.StoryPoints(),
					gs.ID(), gs.Title(), gs.Priority(), gs.State(), gs.StoryPoints())
			}
			wantCriteria, gotCriteria := ws.AcceptanceCriteria(), gs.AcceptanceCriteria()
			if len(gotCriteria) != len(wantCriteria) {
				t.Fatalf("historia %d: se esperaban los criterios %q, salieron %q", ws.ID(), wantCriteria, gotCriteria)
			}
			for k := range wantCriteria {
				if gotCriteria[k] != wantCriteria[k] {
					t.Fatalf("historia %d: se esperaban los criterios %q, salieron %q", ws.ID(), wantCriteria, gotCriteria)
				}
			}
		}
	}
}

// Escenario: Los datos siguen al cerrar y volver a abrir (la conversión del registro)
// Cubre: #77 RN-21, RN-22
func TestLosDatosSiguenAlCerrarYVolverAAbrir(t *testing.T) {
	registry := unRegistroCargado(t)

	data, err := registry.ToJSON()
	if err != nil {
		t.Fatalf("se esperaba convertir el registro sin error, salio: %v", err)
	}
	loaded, err := RegistryFromJSON(data)
	if err != nil {
		t.Fatalf("se esperaba reconstruir el registro sin error, salio: %v\n%s", err, data)
	}

	mismoRegistro(t, registry, loaded)

	// los contadores siguen donde habian quedado (RN-22)
	next, err := loaded.Create(ProjectData{Name: "Tercero", Start: unDia(2026, 10, 9)})
	if err != nil {
		t.Fatalf("no se pudo crear un proyecto en el registro reconstruido: %v", err)
	}
	if next.ID() != 3 {
		t.Fatalf("se esperaba que el proyecto nuevo recibiera el numero 3, recibio el %d", next.ID())
	}
	first, _ := loaded.Get(1)
	story, err := first.AddStory(unosDatos())
	if err != nil {
		t.Fatalf("no se pudo cargar una historia en el proyecto reconstruido: %v", err)
	}
	if story.ID() != 4 {
		t.Fatalf("se esperaba que la historia nueva recibiera el numero 4, recibio el %d", story.ID())
	}
}
