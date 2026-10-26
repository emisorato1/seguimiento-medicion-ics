package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strings"
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
	// las horas estimadas son de la #16, que las suma a la conversion
	if err := first.EstimateHours(estimated.ID(), 6.5); err != nil {
		t.Fatalf("no se pudieron cargar las horas estimadas de prueba: %v", err)
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
	done.estimatedHours = 4
	done.state = StoryDone

	// Los sprints son de la #8: el 1 cerrado y el 2 abierto. Cerrar es de la #10, asi que el
	// test fija el estado desde dentro del paquete.
	closed, err := first.AddSprint(SprintData{Goal: "Tener el backlog estimado", Start: unDia(2026, 10, 1), End: unDia(2026, 10, 14)})
	if err != nil {
		t.Fatalf("no se pudo crear el sprint 1 de prueba: %v", err)
	}
	closed.state = SprintClosed
	// el sprint 1 tuvo la historia Terminada en su Sprint Backlog (#9), y la completo al
	// cerrar: los datos del cierre son de la #10, que los suma a la conversion
	closed.stories = []int{done.ID()}
	closed.closedOn = unDia(2026, 10, 14)
	closed.completedStories = []int{done.ID()}
	closed.pointsAtClose = map[int]int{done.ID(): 3}
	if _, err := first.AddSprint(SprintData{Goal: "Medir la velocidad", Start: unDia(2026, 10, 15), End: unDia(2026, 10, 28)}); err != nil {
		t.Fatalf("no se pudo crear el sprint 2 de prueba: %v", err)
	}
	// el Sprint Backlog del sprint abierto es de la #9, que lo suma a la conversion
	if err := first.AssignStory(estimated.ID()); err != nil {
		t.Fatalf("no se pudo asignar la historia de prueba al sprint 2: %v", err)
	}

	// el esfuerzo es de la #14, que lo suma a la conversion: dos registros, uno de media hora
	for _, data := range []EffortData{
		{StoryID: estimated.ID(), MemberName: "Ana Perez", Activity: "Escribir los tests", Hours: 2, Day: unDia(2026, 10, 20)},
		{StoryID: done.ID(), MemberName: "Juan Gomez", Activity: "Revisar el PR", Hours: 0.5, Day: unDia(2026, 10, 21)},
	} {
		if _, err := first.LogEffort(data, unDia(2026, 10, 21)); err != nil {
			t.Fatalf("no se pudo registrar el esfuerzo de prueba: %v", err)
		}
	}

	// los defectos son de la #15, que los suma a la conversion: uno Abierto, uno En curso y
	// uno Resuelto, con las tres severidades, detectados en los dos sprints
	for _, data := range []DefectData{
		{Title: "La velocidad sale en cero", Description: "El resumen muestra 0 puntos", Severity: SeverityCritical, StoryID: estimated.ID(), DetectedIn: 1},
		{Title: "Se pierde la media hora", Description: "El total redondea las horas", Severity: SeverityMinor, StoryID: 2, DetectedIn: 2},
		{Title: "La fecha sale al reves", Description: "Muestra el mes antes que el dia", Severity: SeverityMajor, StoryID: done.ID(), DetectedIn: 1},
	} {
		if _, err := first.AddDefect(data); err != nil {
			t.Fatalf("no se pudo registrar el defecto de prueba: %v", err)
		}
	}
	if err := first.ChangeDefectState(2, DefectInProgress); err != nil {
		t.Fatalf("no se pudo pasar a En curso el defecto de prueba: %v", err)
	}
	if err := first.ResolveDefect(3, 2); err != nil {
		t.Fatalf("no se pudo resolver el defecto de prueba: %v", err)
	}

	// las rondas de Planning Poker son de la #13, que las suma a la conversion (tarea 14): una
	// revelada, sobre la historia estimada, y una sin revelar con un voto oculto, sobre la 2
	revealed, err := first.OpenVotingRound(estimated.ID())
	if err != nil {
		t.Fatalf("no se pudo abrir la ronda de prueba: %v", err)
	}
	for name, points := range map[string]int{"Ana Perez": 5, "Juan Gomez": 8} {
		if err := revealed.Vote(name, points); err != nil {
			t.Fatalf("no se pudo votar en la ronda de prueba: %v", err)
		}
	}
	if err := revealed.Reveal(); err != nil {
		t.Fatalf("no se pudo revelar la ronda de prueba: %v", err)
	}
	hidden, err := first.OpenVotingRound(2)
	if err != nil {
		t.Fatalf("no se pudo abrir la ronda de prueba: %v", err)
	}
	if err := hidden.Vote("Ana Perez", 3); err != nil {
		t.Fatalf("no se pudo votar en la ronda de prueba: %v", err)
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
			if !mismasHoras(gs.EstimatedHours(), ws.EstimatedHours()) {
				t.Fatalf("proyecto %d: se esperaba la historia %d con %v horas estimadas, salio con %v",
					w.ID(), ws.ID(), ws.EstimatedHours(), gs.EstimatedHours())
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
			wantStories, gotStories := ws.Stories(), gs.Stories()
			if len(gotStories) != len(wantStories) {
				t.Fatalf("sprint %d: se esperaba el Sprint Backlog %v, salio %v", ws.Number(), wantStories, gotStories)
			}
			for k := range wantStories {
				if gotStories[k] != wantStories[k] {
					t.Fatalf("sprint %d: se esperaba el Sprint Backlog %v, salio %v", ws.Number(), wantStories, gotStories)
				}
			}
		}
		for j, ws := range wantSprints {
			gs := gotSprints[j]
			// los datos del cierre de la #10
			if !gs.ClosedOn().Equal(ws.ClosedOn()) || fmt.Sprint(gs.CompletedStories()) != fmt.Sprint(ws.CompletedStories()) ||
				fmt.Sprint(gs.PointsAtClose()) != fmt.Sprint(ws.PointsAtClose()) {
				t.Fatalf("sprint %d: se esperaba el cierre del %v con %v y los puntos %v, salio %v con %v y %v",
					ws.Number(), ws.ClosedOn(), ws.CompletedStories(), ws.PointsAtClose(),
					gs.ClosedOn(), gs.CompletedStories(), gs.PointsAtClose())
			}
		}
		if g.lastSprintNumber != w.lastSprintNumber {
			t.Fatalf("proyecto %d: se esperaba el contador de sprints en %d, salio en %d", w.ID(), w.lastSprintNumber, g.lastSprintNumber)
		}

		wantEfforts, gotEfforts := w.Efforts(), g.Efforts()
		if len(gotEfforts) != len(wantEfforts) {
			t.Fatalf("proyecto %d: se esperaban %d registros de esfuerzo, salieron %d", w.ID(), len(wantEfforts), len(gotEfforts))
		}
		for j, we := range wantEfforts {
			ge := gotEfforts[j]
			if ge.StoryID() != we.StoryID() || ge.MemberName() != we.MemberName() || ge.Activity() != we.Activity() ||
				!mismasHorasReales(ge.Hours(), we.Hours()) || !ge.Day().Equal(we.Day()) {
				t.Fatalf("proyecto %d: se esperaba el esfuerzo %d de %q (%q, %v h, %v), salio %d de %q (%q, %v h, %v)",
					w.ID(), we.StoryID(), we.MemberName(), we.Activity(), we.Hours(), we.Day(),
					ge.StoryID(), ge.MemberName(), ge.Activity(), ge.Hours(), ge.Day())
			}
		}

		// Los defectos de la #15, con su numero, su estado y sus sprints, y el contador.
		wantDefects, gotDefects := w.Defects(), g.Defects()
		if len(gotDefects) != len(wantDefects) {
			t.Fatalf("proyecto %d: se esperaban %d defectos, salieron %d", w.ID(), len(wantDefects), len(gotDefects))
		}
		for j, wd := range wantDefects {
			gd := gotDefects[j]
			if gd.ID() != wd.ID() || gd.Title() != wd.Title() || gd.Description() != wd.Description() ||
				gd.Severity() != wd.Severity() || gd.State() != wd.State() || gd.StoryID() != wd.StoryID() ||
				gd.DetectedIn() != wd.DetectedIn() || gd.ResolvedIn() != wd.ResolvedIn() {
				t.Fatalf("proyecto %d: se esperaba el defecto %+v, salio %+v", w.ID(), *wd, *gd)
			}
		}
		if g.lastDefectID != w.lastDefectID {
			t.Fatalf("proyecto %d: se esperaba el contador de defectos en %d, salio en %d", w.ID(), w.lastDefectID, g.lastDefectID)
		}

		// Las rondas de la #13, con sus participantes y sus votos, tambien los de una ronda sin
		// revelar. El test esta dentro del paquete y compara los datos ocultos directamente.
		if len(g.votingRounds) != len(w.votingRounds) {
			t.Fatalf("proyecto %d: se esperaban %d rondas, salieron %d", w.ID(), len(w.votingRounds), len(g.votingRounds))
		}
		for j, wr := range w.votingRounds {
			gr := g.votingRounds[j]
			if gr.storyID != wr.storyID || gr.number != wr.number || gr.revealed != wr.revealed ||
				!slices.Equal(gr.participants, wr.participants) || !slices.Equal(gr.votes, wr.votes) {
				t.Fatalf("proyecto %d: se esperaba la ronda %d de la historia %d (revelada %v, %v, %v), salio la %d de la %d (revelada %v, %v, %v)",
					w.ID(), wr.number, wr.storyID, wr.revealed, wr.participants, wr.votes,
					gr.number, gr.storyID, gr.revealed, gr.participants, gr.votes)
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
				Efforts: []effortJSON{
					{StoryID: 1, MemberName: "Ana Perez", Activity: "Escribir los tests", Hours: 2, Day: "2026-10-20"},
				},
				LastSprintNumber: 2,
				Sprints: []sprintJSON{
					{Number: 1, Goal: "Tener el backlog estimado", Start: "2026-10-01", End: "2026-10-14", State: "Cerrado", Stories: []int{2},
						ClosedOn: "2026-10-14", CompletedStories: []int{2}, PointsAtClose: map[int]int{2: 3}},
					{Number: 2, Goal: "Medir la velocidad", Start: "2026-10-15", End: "2026-10-28", State: "Abierto"},
				},
				LastDefectID: 2,
				Defects: []defectJSON{
					{ID: 1, Title: "La velocidad sale en cero", Description: "El resumen muestra 0 puntos", Severity: "Crítica",
						State: "Abierto", StoryID: 1, DetectedIn: 1},
					{ID: 2, Title: "La fecha sale al reves", Description: "Muestra el mes antes que el dia", Severity: "Mayor",
						State: "Resuelto", StoryID: 2, DetectedIn: 1, ResolvedIn: 2},
				},
				VotingRounds: []votingRoundJSON{
					{StoryID: 1, Number: 1, Participants: []string{"Ana Perez", "Juan Gomez"},
						Votes: []voteJSON{{MemberName: "Ana Perez", Points: 5}}},
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
		{name: "dos sprints abiertos", change: func(d *registryJSON) {
			// desde la #10 el sprint 1 trae los datos del cierre: se le sacan, para que el
			// motivo siga siendo los dos abiertos
			d.Projects[0].Sprints[0].State = "Abierto"
			d.Projects[0].Sprints[0].ClosedOn = ""
			d.Projects[0].Sprints[0].CompletedStories = nil
			d.Projects[0].Sprints[0].PointsAtClose = nil
		}, want: ErrSprintAlreadyOpen},
		{name: "numero de sprint repetido", change: func(d *registryJSON) { d.Projects[0].Sprints[1].Number = 1 }, want: ErrSavedDataInvalid},
		{name: "contador de sprints que no llega a los numeros usados", change: func(d *registryJSON) { d.Projects[0].LastSprintNumber = 1 }, want: ErrSavedDataInvalid},
		{name: "estado de sprint inexistente", change: func(d *registryJSON) { d.Projects[0].Sprints[0].State = "Planificado" }, want: ErrSavedDataInvalid},
		{name: "fecha de inicio de sprint que no es fecha", change: func(d *registryJSON) { d.Projects[0].Sprints[0].Start = "01/10/2026" }, want: ErrSavedDataInvalid},
		{name: "fecha de fin de sprint que no es fecha", change: func(d *registryJSON) { d.Projects[0].Sprints[0].End = "mañana" }, want: ErrSavedDataInvalid},
		// El Sprint Backlog de la #9 (tarea 14): solo historias que existen en el proyecto.
		{name: "sprint con una historia que no existe", change: func(d *registryJSON) { d.Projects[0].Sprints[0].Stories = []int{99} }, want: ErrStoryNotFound},
		{name: "historia repetida en un Sprint Backlog", change: func(d *registryJSON) { d.Projects[0].Sprints[0].Stories = []int{2, 2} }, want: ErrStoryAlreadyInSprint},
		// Los datos del cierre de la #10 (tarea 15): tienen que ser coherentes con el estado y
		// con el Sprint Backlog.
		{name: "sprint cerrado sin fecha de cierre", change: func(d *registryJSON) { d.Projects[0].Sprints[0].ClosedOn = "" }, want: ErrSprintCloseDateRequired},
		{name: "fecha de cierre que no es fecha", change: func(d *registryJSON) { d.Projects[0].Sprints[0].ClosedOn = "14/10/2026" }, want: ErrSavedDataInvalid},
		{name: "sprint abierto con fecha de cierre", change: func(d *registryJSON) { d.Projects[0].Sprints[1].ClosedOn = "2026-10-28" }, want: ErrSavedDataInvalid},
		{name: "sprint abierto con puntos al cerrar", change: func(d *registryJSON) { d.Projects[0].Sprints[1].PointsAtClose = map[int]int{1: 5} }, want: ErrSavedDataInvalid},
		{name: "historia completada repetida", change: func(d *registryJSON) { d.Projects[0].Sprints[0].CompletedStories = []int{2, 2} }, want: ErrSavedDataInvalid},
		{name: "historia completada que no estaba en el sprint", change: func(d *registryJSON) { d.Projects[0].Sprints[0].CompletedStories = []int{1} }, want: ErrStoryNotInSprint},
		{name: "faltan los puntos al cerrar de una historia del sprint", change: func(d *registryJSON) { d.Projects[0].Sprints[0].PointsAtClose = nil }, want: ErrSavedDataInvalid},
		{name: "puntos al cerrar de una historia que no estaba en el sprint", change: func(d *registryJSON) { d.Projects[0].Sprints[0].PointsAtClose = map[int]int{2: 3, 1: 5} }, want: ErrSavedDataInvalid},
		{name: "puntos al cerrar fuera de la escala", change: func(d *registryJSON) { d.Projects[0].Sprints[0].PointsAtClose = map[int]int{2: 4} }, want: ErrStoryPointsInvalid},
		{name: "horas estimadas negativas", change: func(d *registryJSON) { d.Projects[0].Stories[0].EstimatedHours = -2 }, want: ErrEstimatedHoursInvalid},
		{name: "esfuerzo con horas en cero", change: func(d *registryJSON) { d.Projects[0].Efforts[0].Hours = 0 }, want: ErrEffortHoursInvalid},
		{name: "esfuerzo sin actividad", change: func(d *registryJSON) { d.Projects[0].Efforts[0].Activity = "  " }, want: ErrEffortActivityRequired},
		{name: "esfuerzo sin integrante", change: func(d *registryJSON) { d.Projects[0].Efforts[0].MemberName = "  " }, want: ErrMemberNotFound},
		{name: "esfuerzo sobre una historia que no existe", change: func(d *registryJSON) { d.Projects[0].Efforts[0].StoryID = 99 }, want: ErrStoryNotFound},
		{name: "esfuerzo con una fecha que no es fecha", change: func(d *registryJSON) { d.Projects[0].Efforts[0].Day = "20/10/2026" }, want: ErrSavedDataInvalid},
		// Los defectos de la #15 (tarea 16): las mismas reglas que al registrarlos, y lo que
		// solo un archivo puede traer mal.
		{name: "defecto sin titulo", change: func(d *registryJSON) { d.Projects[0].Defects[0].Title = "  " }, want: ErrDefectTitleRequired},
		{name: "defecto sin descripcion", change: func(d *registryJSON) { d.Projects[0].Defects[0].Description = "" }, want: ErrDefectDescriptionRequired},
		{name: "severidad de defecto inexistente", change: func(d *registryJSON) { d.Projects[0].Defects[0].Severity = "Urgente" }, want: ErrDefectSeverityInvalid},
		{name: "estado de defecto inexistente", change: func(d *registryJSON) { d.Projects[0].Defects[0].State = "Cerrado" }, want: ErrSavedDataInvalid},
		{name: "defecto sobre una historia que no existe", change: func(d *registryJSON) { d.Projects[0].Defects[0].StoryID = 99 }, want: ErrStoryNotFound},
		{name: "defecto detectado en un sprint que no existe", change: func(d *registryJSON) { d.Projects[0].Defects[0].DetectedIn = 9 }, want: ErrSprintNotFound},
		{name: "defecto Resuelto sin sprint de resolucion", change: func(d *registryJSON) { d.Projects[0].Defects[1].ResolvedIn = 0 }, want: ErrDefectResolveNeedsSprint},
		{name: "defecto resuelto en un sprint que no existe", change: func(d *registryJSON) { d.Projects[0].Defects[1].ResolvedIn = 9 }, want: ErrSprintNotFound},
		{name: "defecto resuelto antes de detectarse", change: func(d *registryJSON) {
			d.Projects[0].Defects[1].DetectedIn = 2
			d.Projects[0].Defects[1].ResolvedIn = 1
		}, want: ErrDefectResolvedBeforeDetected},
		{name: "defecto sin resolver con sprint de resolucion", change: func(d *registryJSON) { d.Projects[0].Defects[0].ResolvedIn = 2 }, want: ErrSavedDataInvalid},
		{name: "numero de defecto repetido", change: func(d *registryJSON) { d.Projects[0].Defects[1].ID = 1 }, want: ErrSavedDataInvalid},
		{name: "contador de defectos que no llega a los numeros usados", change: func(d *registryJSON) { d.Projects[0].LastDefectID = 1 }, want: ErrSavedDataInvalid},
		// D16 de la #77: un dato que el sistema no conoce rechaza la carga, también dentro de
		// un defecto.
		{name: "defecto con un dato que el sistema no conoce", raw: strings.Replace(string(base),
			`"title":"La velocidad sale en cero"`, `"title":"La velocidad sale en cero","responsable":"Ana Perez"`, 1), want: ErrSavedDataInvalid},
		// Las rondas de la #13 (tarea 14): las mismas reglas que al votar, y lo que solo un
		// archivo puede traer mal.
		{name: "ronda sobre una historia que no existe", change: func(d *registryJSON) { d.Projects[0].VotingRounds[0].StoryID = 99 }, want: ErrStoryNotFound},
		{name: "ronda con numero cero", change: func(d *registryJSON) { d.Projects[0].VotingRounds[0].Number = 0 }, want: ErrSavedDataInvalid},
		{name: "voto fuera de la escala", change: func(d *registryJSON) { d.Projects[0].VotingRounds[0].Votes[0].Points = 4 }, want: ErrStoryPointsInvalid},
		{name: "voto de quien no participa", change: func(d *registryJSON) { d.Projects[0].VotingRounds[0].Votes[0].MemberName = "Pedro Gomez" }, want: ErrNotParticipant},
		{name: "dos votos de la misma persona", change: func(d *registryJSON) {
			d.Projects[0].VotingRounds[0].Votes = append(d.Projects[0].VotingRounds[0].Votes, voteJSON{MemberName: "Ana Perez", Points: 3})
		}, want: ErrAlreadyVoted},
		{name: "dos rondas sin revelar sobre la misma historia", change: func(d *registryJSON) {
			d.Projects[0].VotingRounds = append(d.Projects[0].VotingRounds, votingRoundJSON{StoryID: 1, Number: 2, Participants: []string{"Ana Perez"}})
		}, want: ErrVotingRoundInProgress},
		{name: "ronda revelada a la que le faltan votos", change: func(d *registryJSON) { d.Projects[0].VotingRounds[0].Revealed = true }, want: ErrVotesMissing},
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

// Sin escenario: tarea 14 de la #14, la carga no controla el día futuro
// Cubre: Fase 2 de la #14, la decisión sobre los datos guardados
//
// Al cargar no hay "hoy" con qué comparar, y no hace falta: un día que no era futuro cuando
// se guardó no puede volverse futuro después. Un registro con un día lejano se carga igual.
func TestUnEsfuerzoDeUnDiaPosteriorAHoySeCargaIgual(t *testing.T) {
	saved := unosDatosGuardados()
	saved.Projects[0].Efforts[0].Day = "2099-01-01"
	data, err := json.Marshal(saved)
	if err != nil {
		t.Fatalf("no se pudieron armar los datos de prueba: %v", err)
	}

	registry, err := RegistryFromJSON(data)

	if err != nil {
		t.Fatalf("se esperaba cargar los datos sin error, salio: %v", err)
	}
	project, _ := registry.Get(1)
	if len(project.Efforts()) != 1 || !project.Efforts()[0].Day().Equal(unDia(2099, 1, 1)) {
		t.Fatalf("se esperaba el registro del 01/01/2099, salieron %d registros", len(project.Efforts()))
	}
}

// Sin escenario: hallazgo de review-merge en el PR #100, decisión (a) de Emiliano del 21/10
// Cubre: Fase 2 de la #14, la decisión sobre los datos guardados
//
// Quitar a un integrante (#4) no borra las horas que cargó. Si la carga exigiera que siga
// siendo integrante, el archivo que guarda la propia web no se volvería a abrir. Sus horas
// se cargan con su nombre, siguen sumando en la historia, y consultar las suyas informa que
// no se encontró el integrante, igual que antes de cerrar (RN-14).
func TestElEsfuerzoDeUnIntegranteQuitadoSeCargaIgual(t *testing.T) {
	registry := unRegistroCargado(t)
	project, _ := registry.Get(1)
	if err := project.RemoveMember("Juan Gomez"); err != nil {
		t.Fatalf("no se pudo quitar el integrante de prueba: %v", err)
	}
	data, err := registry.ToJSON()
	if err != nil {
		t.Fatalf("no se pudieron convertir los datos: %v", err)
	}

	loaded, err := RegistryFromJSON(data)

	if err != nil {
		t.Fatalf("se esperaba cargar los datos sin error, salio: %v", err)
	}
	mismoRegistro(t, registry, loaded)
	got, _ := loaded.Get(1)
	if hours, err := got.StoryHours(3); err != nil || hours != 0.5 {
		t.Fatalf("se esperaba que la historia 3 siguiera sumando 0,5 horas, salio %v (%v)", hours, err)
	}
	if _, err := got.MemberHours("Juan Gomez"); !errors.Is(err, ErrMemberNotFound) {
		t.Fatalf("se esperaba ErrMemberNotFound al consultar las horas de quien ya no es integrante, salio: %v", err)
	}
}
