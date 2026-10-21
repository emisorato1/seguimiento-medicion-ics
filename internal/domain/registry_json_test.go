package domain

import (
	"encoding/json"
	"errors"
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

	// Los sprints son de la #8: el 1 cerrado y el 2 abierto. Cerrar es de la #10, asi que el
	// test fija el estado desde dentro del paquete.
	closed, err := first.AddSprint(SprintData{Goal: "Tener el backlog estimado", Start: unDia(2026, 10, 1), End: unDia(2026, 10, 14)})
	if err != nil {
		t.Fatalf("no se pudo crear el sprint 1 de prueba: %v", err)
	}
	closed.state = SprintClosed
	if _, err := first.AddSprint(SprintData{Goal: "Medir la velocidad", Start: unDia(2026, 10, 15), End: unDia(2026, 10, 28)}); err != nil {
		t.Fatalf("no se pudo crear el sprint 2 de prueba: %v", err)
	}
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

		// Los sprints de la #8, con su numero, goal, fechas y estado, y el contador.
		wantSprints, gotSprints := w.Sprints(), g.Sprints()
		if len(gotSprints) != len(wantSprints) {
			t.Fatalf("proyecto %d: se esperaban %d sprints, salieron %d", w.ID(), len(wantSprints), len(gotSprints))
		}
		for j, ws := range wantSprints {
			gs := gotSprints[j]
			if gs.Number() != ws.Number() || gs.Goal() != ws.Goal() || gs.State() != ws.State() ||
				!gs.Start().Equal(ws.Start()) || !gs.End().Equal(ws.End()) {
				t.Fatalf("proyecto %d: se esperaba el sprint %d %q (%v, %v a %v), salio %d %q (%v, %v a %v)",
					w.ID(), ws.Number(), ws.Goal(), ws.State(), ws.Start(), ws.End(),
					gs.Number(), gs.Goal(), gs.State(), gs.Start(), gs.End())
			}
		}
		if g.lastSprintNumber != w.lastSprintNumber {
			t.Fatalf("proyecto %d: se esperaba el contador de sprints en %d, salio en %d", w.ID(), w.lastSprintNumber, g.lastSprintNumber)
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

// unosDatosGuardados devuelve datos guardados válidos, en la forma del archivo, para que
// cada caso de rechazo cambie uno solo.
func unosDatosGuardados() registryJSON {
	return registryJSON{
		NextID: 3,
		Projects: []projectJSON{
			{ID: 1, Name: "Sistema de metricas", Start: "2026-10-01", End: "2026-11-02", NextStoryID: 3,
				Members: []memberJSON{{Name: "Ana Perez", Role: "Agile Enabler"}, {Name: "Juan Gomez", Role: "Product Builder"}},
				Stories: []storyJSON{
					{ID: 1, Title: "Registrar esfuerzo", Description: "Cargar horas", Priority: "Alta", State: "Pendiente",
						StoryPoints: 5, AcceptanceCriteria: []string{"Valida las horas"}},
					{ID: 2, Title: "Crear proyectos", Description: "Alta de proyectos", Priority: "Baja", State: "Terminada",
						StoryPoints: 3, AcceptanceCriteria: []string{"Pide el nombre"}},
				},
				LastSprintNumber: 2,
				Sprints: []sprintJSON{
					{Number: 1, Goal: "Tener el backlog estimado", Start: "2026-10-01", End: "2026-10-14", State: "Cerrado"},
					{Number: 2, Goal: "Medir la velocidad", Start: "2026-10-15", End: "2026-10-28", State: "Abierto"},
				}},
			{ID: 2, Name: "Otro proyecto", Start: "2026-10-05", NextStoryID: 1},
		},
	}
}

// Escenario: Los datos guardados estan danados (la reconstrucción del registro)
// Cubre: #77 RN-24
//
// Los cinco primeros casos son los de la tarea 3. Los demás cubren cada control que hace la
// reconstrucción, para que ninguno quede sin probar.
func TestLosDatosGuardadosEstanDanados(t *testing.T) {
	base, err := json.Marshal(unosDatosGuardados())
	if err != nil {
		t.Fatalf("no se pudieron armar los datos de prueba: %v", err)
	}
	if _, err := RegistryFromJSON(base); err != nil {
		t.Fatalf("los datos de prueba sin cambios tendrian que cargarse, salio: %v", err)
	}

	cases := []struct {
		name   string
		raw    string
		change func(d *registryJSON)
		want   error
	}{
		{name: "texto que no es JSON", raw: `{"projects": [`, want: ErrSavedDataInvalid},
		{name: "proyecto sin nombre", change: func(d *registryJSON) { d.Projects[0].Name = "" }, want: ErrProjectNameRequired},
		{name: "4 story points", change: func(d *registryJSON) { d.Projects[0].Stories[0].StoryPoints = 4 }, want: ErrStoryPointsInvalid},
		{name: "rol inexistente", change: func(d *registryJSON) { d.Projects[0].Members[1].Role = "Scrum Master" }, want: ErrMemberRoleInvalid},
		{name: "dos Agile Enabler", change: func(d *registryJSON) { d.Projects[0].Members[1].Role = "Agile Enabler" }, want: ErrMemberAgileEnablerTaken},
		{name: "un dato que el sistema no conoce", raw: `{"next_id": 1, "projects": [], "sprints": []}`, want: ErrSavedDataInvalid},
		{name: "texto sobrante despues de los datos", raw: `{"next_id": 1, "projects": []} basura`, want: ErrSavedDataInvalid},
		{name: "fecha de inicio que no es fecha", change: func(d *registryJSON) { d.Projects[0].Start = "01/10/2026" }, want: ErrSavedDataInvalid},
		{name: "fecha de fin que no es fecha", change: func(d *registryJSON) { d.Projects[0].End = "mañana" }, want: ErrSavedDataInvalid},
		{name: "numero de proyecto repetido", change: func(d *registryJSON) { d.Projects[1].ID = 1 }, want: ErrSavedDataInvalid},
		{name: "contador de proyectos que no supera los numeros usados", change: func(d *registryJSON) { d.NextID = 2 }, want: ErrSavedDataInvalid},
		{name: "contador de historias en cero", change: func(d *registryJSON) { d.Projects[0].NextStoryID = 0 }, want: ErrSavedDataInvalid},
		{name: "historia sin titulo", change: func(d *registryJSON) { d.Projects[0].Stories[0].Title = "" }, want: ErrStoryTitleRequired},
		{name: "prioridad inexistente", change: func(d *registryJSON) { d.Projects[0].Stories[0].Priority = "Urgente" }, want: ErrStoryPriorityInvalid},
		{name: "estado inexistente", change: func(d *registryJSON) { d.Projects[0].Stories[1].State = "Cerrada" }, want: ErrSavedDataInvalid},
		// Los sprints de la #8 (tarea 15): lo mismo que se controla al crear, mas lo que solo
		// puede venir mal de un archivo.
		{name: "sprint sin Sprint Goal", change: func(d *registryJSON) { d.Projects[0].Sprints[0].Goal = "   " }, want: ErrSprintGoalRequired},
		{name: "sprint sin fecha de inicio", change: func(d *registryJSON) { d.Projects[0].Sprints[0].Start = "" }, want: ErrSprintStartRequired},
		{name: "sprint con las fechas al reves", change: func(d *registryJSON) { d.Projects[0].Sprints[0].End = "2026-09-30" }, want: ErrSprintDatesInconsistent},
		{name: "dos sprints abiertos", change: func(d *registryJSON) { d.Projects[0].Sprints[0].State = "Abierto" }, want: ErrSprintAlreadyOpen},
		{name: "numero de sprint repetido", change: func(d *registryJSON) { d.Projects[0].Sprints[1].Number = 1 }, want: ErrSavedDataInvalid},
		{name: "contador de sprints que no llega a los numeros usados", change: func(d *registryJSON) { d.Projects[0].LastSprintNumber = 1 }, want: ErrSavedDataInvalid},
		{name: "estado de sprint inexistente", change: func(d *registryJSON) { d.Projects[0].Sprints[0].State = "Planificado" }, want: ErrSavedDataInvalid},
		{name: "fecha de inicio de sprint que no es fecha", change: func(d *registryJSON) { d.Projects[0].Sprints[0].Start = "01/10/2026" }, want: ErrSavedDataInvalid},
		{name: "fecha de fin de sprint que no es fecha", change: func(d *registryJSON) { d.Projects[0].Sprints[0].End = "mañana" }, want: ErrSavedDataInvalid},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			data := []byte(tc.raw)
			if tc.change != nil {
				saved := unosDatosGuardados()
				tc.change(&saved)
				changed, marshalErr := json.Marshal(saved)
				if marshalErr != nil {
					t.Fatalf("no se pudieron armar los datos de prueba: %v", marshalErr)
				}
				data = changed
			}

			registry, err := RegistryFromJSON(data)

			if !errors.Is(err, ErrSavedDataInvalid) {
				t.Fatalf("se esperaba ErrSavedDataInvalid, salio: %v", err)
			}
			if !errors.Is(err, tc.want) {
				t.Fatalf("se esperaba que el motivo fuera %v, salio: %v", tc.want, err)
			}
			if registry != nil {
				t.Fatalf("se esperaba que no se cargara ningun registro, se cargo uno con %d proyectos", len(registry.Projects()))
			}
		})
	}
}

// Sin escenario: hallazgo de front en el PR #92, los mensajes van en castellano (AGENTS.md)
// Cubre: #77 RN-24
//
// El detalle de encoding/json viene en inglés («unexpected EOF», «unknown field»): no tiene
// que llegar al mensaje que lee el usuario al arrancar.
func TestElMensajeDeDatosDanadosEstaEnCastellano(t *testing.T) {
	want := "los datos guardados no se pueden leer: el archivo está dañado, incompleto o tiene datos que el sistema no conoce"
	cases := []struct {
		name string
		raw  string
	}{
		{"archivo cortado", `{"next_id": 1, "projects": [`},
		{"texto que no es JSON", `hola`},
		{"un dato que el sistema no conoce", `{"next_id": 1, "projects": [], "sprints": []}`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := RegistryFromJSON([]byte(tc.raw))

			if err == nil || err.Error() != want {
				t.Fatalf("se esperaba el mensaje %q, salio: %v", want, err)
			}
		})
	}
}
