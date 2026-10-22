package domain

import (
	"errors"
	"fmt"
	"testing"
)

// conSprints crea en el proyecto los sprints del 1 al n, con el alta de la #8.
func conSprints(t *testing.T, project *Project, n int) {
	t.Helper()
	for i := 1; i <= n; i++ {
		if _, err := project.AddSprint(SprintData{Goal: "Sprint de prueba", Start: unDia(2026, 10, 1), End: unDia(2026, 10, 5)}); err != nil {
			t.Fatalf("no se pudo crear el sprint de prueba %d: %v", i, err)
		}
	}
}

// unProyectoConDefectos arma un proyecto con la historia 1 y los sprints del 1 al n.
func unProyectoConDefectos(t *testing.T, sprints int) *Project {
	t.Helper()
	project := unProyecto(t)
	if _, err := project.AddStory(unosDatos()); err != nil {
		t.Fatalf("no se pudo crear la historia de prueba: %v", err)
	}
	conSprints(t, project, sprints)
	return project
}

// unDefecto devuelve los datos de un defecto válido: Mayor, sobre la historia 1 y detectado
// en el sprint 1. Cada test cambia solo lo que prueba.
func unDefecto() DefectData {
	return DefectData{
		Title:       "La velocidad sale en cero",
		Description: "El resumen del sprint muestra 0 puntos",
		Severity:    SeverityMajor,
		StoryID:     1,
		DetectedIn:  1,
	}
}

// Escenario: Registrar un defecto
// Cubre: RN-1, RN-2, RN-3
func TestRegistrarUnDefecto(t *testing.T) {
	project := unProyectoConDefectos(t, 1)

	defect, err := project.AddDefect(unDefecto())

	if err != nil {
		t.Fatalf("se esperaba registrar el defecto sin error, salio: %v", err)
	}
	if defect.ID() != 1 || defect.State() != DefectOpen || defect.ResolvedIn() != 0 {
		t.Fatalf("se esperaba el defecto 1, Abierto y sin sprint de resolucion, salio %d, %v, %d", defect.ID(), defect.State(), defect.ResolvedIn())
	}
	if defect.Title() != "La velocidad sale en cero" || defect.Description() != "El resumen del sprint muestra 0 puntos" ||
		defect.Severity() != SeverityMajor || defect.StoryID() != 1 || defect.DetectedIn() != 1 {
		t.Fatalf("se esperaban los datos registrados, salio %q, %q, %v, historia %d, sprint %d",
			defect.Title(), defect.Description(), defect.Severity(), defect.StoryID(), defect.DetectedIn())
	}
	if defects := project.Defects(); len(defects) != 1 || defects[0] != defect {
		t.Fatalf("se esperaba que el proyecto tuviera ese unico defecto, tiene %d", len(defects))
	}
}

// Escenario: El defecto nuevo recibe el identificador siguiente
// Cubre: RN-1, RN-17
func TestElDefectoNuevoRecibeElIdentificadorSiguiente(t *testing.T) {
	project := unProyectoConDefectos(t, 1)
	if _, err := project.AddDefect(unDefecto()); err != nil {
		t.Fatalf("no se pudo registrar el primer defecto de prueba: %v", err)
	}

	defect, err := project.AddDefect(unDefecto())

	if err != nil || defect.ID() != 2 {
		t.Fatalf("se esperaba el defecto 2 con los mismos datos, salio %v (error: %v)", defect, err)
	}
	if len(project.Defects()) != 2 {
		t.Fatalf("se esperaba que el proyecto quedara con 2 defectos, tiene %d", len(project.Defects()))
	}
}

// Esquema del escenario: Rechazar un defecto sin titulo o sin descripcion
// Cubre: RN-4, RN-18
func TestRechazarUnDefectoSinTituloOSinDescripcion(t *testing.T) {
	cases := []struct {
		name   string
		change func(d *DefectData)
		want   error
	}{
		{"el titulo", func(d *DefectData) { d.Title = "   " }, ErrDefectTitleRequired},
		{"la descripcion", func(d *DefectData) { d.Description = "" }, ErrDefectDescriptionRequired},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			project := unProyectoConDefectos(t, 1)
			data := unDefecto()
			tc.change(&data)

			_, err := project.AddDefect(data)

			if !errors.Is(err, tc.want) {
				t.Fatalf("se esperaba %v, salio: %v", tc.want, err)
			}
			if len(project.Defects()) != 0 {
				t.Fatalf("se esperaba que el proyecto siguiera sin defectos, tiene %d", len(project.Defects()))
			}
		})
	}
}

// Escenario: Guardar el titulo y la descripcion sin los espacios de los extremos
// Cubre: RN-5
func TestGuardarElTituloYLaDescripcionSinLosEspaciosDeLosExtremos(t *testing.T) {
	project := unProyectoConDefectos(t, 1)
	data := unDefecto()
	data.Title = "  La velocidad sale en cero  "
	data.Description = "  El resumen muestra 0  "

	defect, err := project.AddDefect(data)

	if err != nil {
		t.Fatalf("se esperaba registrar el defecto sin error, salio: %v", err)
	}
	if defect.Title() != "La velocidad sale en cero" || defect.Description() != "El resumen muestra 0" {
		t.Fatalf("se esperaban el titulo y la descripcion sin espacios de mas, salio %q y %q", defect.Title(), defect.Description())
	}
}

// Esquema del escenario: Registrar un defecto con cada severidad
// Cubre: RN-6
func TestRegistrarUnDefectoConCadaSeveridad(t *testing.T) {
	cases := []struct {
		written string
		want    string
	}{
		{"Crítica", "Crítica"},
		{"critica", "Crítica"},
		{"  MAYOR  ", "Mayor"},
		{"Menor", "Menor"},
	}
	for _, tc := range cases {
		t.Run(tc.written, func(t *testing.T) {
			project := unProyectoConDefectos(t, 1)
			severity, err := ParseSeverity(tc.written)
			if err != nil {
				t.Fatalf("se esperaba aceptar la severidad %q, salio: %v", tc.written, err)
			}
			data := unDefecto()
			data.Severity = severity

			defect, err := project.AddDefect(data)

			if err != nil {
				t.Fatalf("se esperaba registrar el defecto sin error, salio: %v", err)
			}
			if defect.Severity().String() != tc.want {
				t.Fatalf("se esperaba la severidad %q, salio %q", tc.want, defect.Severity().String())
			}
		})
	}
}

// Escenario: Rechazar una severidad que no existe
// Cubre: RN-6, RN-18
//
// Además del texto del escenario, cubre lo que llega a AddDefect sin pasar por la conversión:
// una severidad sin indicar (el valor cero) y una armada a mano fuera de rango (RN-6).
func TestRechazarUnaSeveridadQueNoExiste(t *testing.T) {
	if _, err := ParseSeverity("Urgente"); !errors.Is(err, ErrDefectSeverityInvalid) {
		t.Fatalf("se esperaba ErrDefectSeverityInvalid para %q, salio: %v", "Urgente", err)
	}
	for _, severity := range []Severity{0, Severity(9)} {
		project := unProyectoConDefectos(t, 1)
		data := unDefecto()
		data.Severity = severity

		_, err := project.AddDefect(data)

		if !errors.Is(err, ErrDefectSeverityInvalid) {
			t.Fatalf("se esperaba ErrDefectSeverityInvalid con la severidad %d, salio: %v", severity, err)
		}
		if len(project.Defects()) != 0 {
			t.Fatalf("se esperaba que el proyecto siguiera sin defectos, tiene %d", len(project.Defects()))
		}
	}
}

// Sin escenario: tarea 4 de la #15, una severidad que no existe no se muestra como Crítica
// Cubre: RN-6
func TestUnaSeveridadQueNoExisteSeMuestraSinSeveridad(t *testing.T) {
	for _, severity := range []Severity{0, Severity(9)} {
		if severity.String() != "sin severidad" {
			t.Fatalf("se esperaba %q para la severidad %d, salio %q", "sin severidad", severity, severity.String())
		}
	}
}

// Escenario: Rechazar un defecto sobre una historia que no es del proyecto
// Cubre: RN-7, RN-18
func TestRechazarUnDefectoSobreUnaHistoriaQueNoEsDelProyecto(t *testing.T) {
	project := unProyectoConDefectos(t, 1)
	data := unDefecto()
	data.StoryID = 99

	_, err := project.AddDefect(data)

	if !errors.Is(err, ErrStoryNotFound) {
		t.Fatalf("se esperaba ErrStoryNotFound para la historia 99, salio: %v", err)
	}
	if len(project.Defects()) != 0 {
		t.Fatalf("se esperaba que el proyecto siguiera sin defectos, tiene %d", len(project.Defects()))
	}
}

// Escenario: Registrar un defecto sobre una historia terminada
// Cubre: RN-7
func TestRegistrarUnDefectoSobreUnaHistoriaTerminada(t *testing.T) {
	project := unProyectoConDefectos(t, 1)
	story, _ := project.Story(1)
	// Terminar una historia es de la #10, que todavia no esta en dev: el test la prepara
	// desde dentro del paquete, como la #12 en su tarea 9.
	story.state = StoryDone

	if _, err := project.AddDefect(unDefecto()); err != nil {
		t.Fatalf("se esperaba registrar un defecto sobre una historia Terminada, salio: %v", err)
	}
}

// Escenario: Rechazar un defecto con un sprint de deteccion que no existe
// Cubre: RN-8, RN-18
func TestRechazarUnDefectoConUnSprintDeDeteccionQueNoExiste(t *testing.T) {
	project := unProyectoConDefectos(t, 1)
	data := unDefecto()
	data.DetectedIn = 4

	_, err := project.AddDefect(data)

	if !errors.Is(err, ErrSprintNotFound) {
		t.Fatalf("se esperaba ErrSprintNotFound para el sprint 4, salio: %v", err)
	}
	if len(project.Defects()) != 0 {
		t.Fatalf("se esperaba que el proyecto siguiera sin defectos, tiene %d", len(project.Defects()))
	}
}

// Escenario: Registrar un defecto detectado en un sprint ya cerrado
// Cubre: RN-8
func TestRegistrarUnDefectoDetectadoEnUnSprintYaCerrado(t *testing.T) {
	project := unProyectoConDefectos(t, 1)
	sprint, _ := project.Sprint(1)
	// Cerrar un sprint es de la #10, que todavia no esta en dev: el test lo crea con el alta de
	// la #8 y le fija el estado desde dentro del paquete (tarea 6 de la Fase 3).
	sprint.state = SprintClosed

	if _, err := project.AddDefect(unDefecto()); err != nil {
		t.Fatalf("se esperaba registrar un defecto detectado en un sprint cerrado, salio: %v", err)
	}
}

// unDefectoRegistrado registra el defecto 1 en un proyecto con la historia 1 y los sprints del
// 1 al n, y lo devuelve.
func unDefectoRegistrado(t *testing.T, sprints int) (*Project, *Defect) {
	t.Helper()
	project := unProyectoConDefectos(t, sprints)
	defect, err := project.AddDefect(unDefecto())
	if err != nil {
		t.Fatalf("no se pudo registrar el defecto de prueba: %v", err)
	}
	return project, defect
}

// Esquema del escenario: Cambiar el estado de un defecto no resuelto
// Cubre: RN-9
func TestCambiarElEstadoDeUnDefectoNoResuelto(t *testing.T) {
	cases := []struct {
		name     string
		from, to DefectState
	}{
		{"de Abierto a En curso", DefectOpen, DefectInProgress},
		{"de En curso a Abierto", DefectInProgress, DefectOpen},
		{"de Abierto a Abierto", DefectOpen, DefectOpen},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			project, defect := unDefectoRegistrado(t, 1)
			defect.state = tc.from

			err := project.ChangeDefectState(defect.ID(), tc.to)

			if err != nil {
				t.Fatalf("se esperaba cambiar el estado sin error, salio: %v", err)
			}
			if defect.State() != tc.to {
				t.Fatalf("se esperaba el estado %d, salio %d", tc.to, defect.State())
			}
		})
	}
}

// Escenario: Rechazar el pase a Resuelto por cambio de estado
// Cubre: RN-10, RN-18
func TestRechazarElPaseAResueltoPorCambioDeEstado(t *testing.T) {
	project, defect := unDefectoRegistrado(t, 1)

	err := project.ChangeDefectState(defect.ID(), DefectResolved)

	if !errors.Is(err, ErrDefectResolveNeedsSprint) {
		t.Fatalf("se esperaba ErrDefectResolveNeedsSprint, salio: %v", err)
	}
	if defect.State() != DefectOpen {
		t.Fatalf("se esperaba que el defecto siguiera Abierto, quedo en %d", defect.State())
	}
}

// Sin escenario: tarea 8 de la #15, un estado armado a mano que no existe
// Cubre: RN-9, RN-18
//
// La RN-9 solo deja pasar entre Abierto y En curso. Un estado que no es ninguno de los tres
// —un valor armado a mano, como DefectState(9)— también se rechaza, para que nunca quede
// guardado un estado que no existe.
func TestRechazarUnEstadoQueNoExiste(t *testing.T) {
	project, defect := unDefectoRegistrado(t, 1)

	err := project.ChangeDefectState(defect.ID(), DefectState(9))

	if !errors.Is(err, ErrDefectStateInvalid) {
		t.Fatalf("se esperaba ErrDefectStateInvalid, salio: %v", err)
	}
	if defect.State() != DefectOpen {
		t.Fatalf("se esperaba que el defecto siguiera Abierto, quedo en %d", defect.State())
	}
}

// Esquema del escenario: Resolver un defecto
// Cubre: RN-11
func TestResolverUnDefecto(t *testing.T) {
	for _, from := range []DefectState{DefectOpen, DefectInProgress} {
		project, defect := unDefectoRegistrado(t, 2)
		defect.state = from

		err := project.ResolveDefect(defect.ID(), 2)

		if err != nil {
			t.Fatalf("se esperaba resolver el defecto desde el estado %d sin error, salio: %v", from, err)
		}
		if defect.State() != DefectResolved || defect.ResolvedIn() != 2 {
			t.Fatalf("se esperaba el defecto Resuelto en el sprint 2, salio en %d con el sprint %d", defect.State(), defect.ResolvedIn())
		}
	}
}

// Escenario: Rechazar la resolucion con un sprint que no existe
// Cubre: RN-12, RN-18
func TestRechazarLaResolucionConUnSprintQueNoExiste(t *testing.T) {
	project, defect := unDefectoRegistrado(t, 1)

	err := project.ResolveDefect(defect.ID(), 4)

	if !errors.Is(err, ErrSprintNotFound) {
		t.Fatalf("se esperaba ErrSprintNotFound para el sprint 4, salio: %v", err)
	}
	if defect.State() != DefectOpen || defect.ResolvedIn() != 0 {
		t.Fatalf("se esperaba que el defecto siguiera Abierto y sin sprint de resolucion, quedo en %d con el sprint %d", defect.State(), defect.ResolvedIn())
	}
}

// Escenario: Rechazar la resolucion en un sprint anterior al de deteccion
// Cubre: RN-13, RN-18
func TestRechazarLaResolucionEnUnSprintAnteriorAlDeDeteccion(t *testing.T) {
	project := unProyectoConDefectos(t, 2)
	data := unDefecto()
	data.DetectedIn = 2
	defect, err := project.AddDefect(data)
	if err != nil {
		t.Fatalf("no se pudo registrar el defecto de prueba: %v", err)
	}

	err = project.ResolveDefect(defect.ID(), 1)

	if !errors.Is(err, ErrDefectResolvedBeforeDetected) {
		t.Fatalf("se esperaba ErrDefectResolvedBeforeDetected, salio: %v", err)
	}
	if defect.State() != DefectOpen || defect.ResolvedIn() != 0 {
		t.Fatalf("se esperaba que el defecto siguiera sin resolver, quedo en %d con el sprint %d", defect.State(), defect.ResolvedIn())
	}
}

// Escenario: Resolver un defecto en el mismo sprint en que se detecto
// Cubre: RN-13
func TestResolverUnDefectoEnElMismoSprintEnQueSeDetecto(t *testing.T) {
	project := unProyectoConDefectos(t, 2)
	data := unDefecto()
	data.DetectedIn = 2
	defect, err := project.AddDefect(data)
	if err != nil {
		t.Fatalf("no se pudo registrar el defecto de prueba: %v", err)
	}

	err = project.ResolveDefect(defect.ID(), 2)

	if err != nil || defect.State() != DefectResolved || defect.ResolvedIn() != 2 {
		t.Fatalf("se esperaba el defecto Resuelto en el sprint 2, salio en %d con el sprint %d (error: %v)", defect.State(), defect.ResolvedIn(), err)
	}
}

// Esquema del escenario: Un defecto resuelto no se vuelve a tocar
// Cubre: RN-14, RN-18
func TestUnDefectoResueltoNoSeVuelveATocar(t *testing.T) {
	cases := []struct {
		name string
		run  func(project *Project, id int) error
	}{
		{"cambiarle el estado a Abierto", func(project *Project, id int) error { return project.ChangeDefectState(id, DefectOpen) }},
		{"resolver en el sprint 3", func(project *Project, id int) error { return project.ResolveDefect(id, 3) }},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			project, defect := unDefectoRegistrado(t, 3)
			if err := project.ResolveDefect(defect.ID(), 2); err != nil {
				t.Fatalf("no se pudo resolver el defecto de prueba: %v", err)
			}

			err := tc.run(project, defect.ID())

			if !errors.Is(err, ErrDefectAlreadyResolved) {
				t.Fatalf("se esperaba ErrDefectAlreadyResolved, salio: %v", err)
			}
			if defect.State() != DefectResolved || defect.ResolvedIn() != 2 {
				t.Fatalf("se esperaba que siguiera Resuelto en el sprint 2, quedo en %d con el sprint %d", defect.State(), defect.ResolvedIn())
			}
		})
	}
}

// numerosDeDefectos devuelve los identificadores de una lista de defectos, en su orden.
func numerosDeDefectos(defects []*Defect) []int {
	ids := []int{}
	for _, defect := range defects {
		ids = append(ids, defect.ID())
	}
	return ids
}

// Esquema del escenario: Listar los defectos del proyecto
// Cubre: RN-15
func TestListarLosDefectosDelProyecto(t *testing.T) {
	project := unProyectoConDefectos(t, 1)
	for _, severity := range []Severity{SeverityCritical, SeverityMinor, SeverityMinor} {
		data := unDefecto()
		data.Severity = severity
		if _, err := project.AddDefect(data); err != nil {
			t.Fatalf("no se pudo registrar el defecto de prueba: %v", err)
		}
	}
	if err := project.ResolveDefect(3, 1); err != nil {
		t.Fatalf("no se pudo resolver el defecto de prueba: %v", err)
	}

	cases := []struct {
		name string
		got  []*Defect
		want []int
	}{
		{"todos los defectos", project.Defects(), []int{1, 2, 3}},
		{"los defectos en Abierto", project.DefectsByState(DefectOpen), []int{1, 2}},
		{"los defectos de severidad Menor", project.DefectsBySeverity(SeverityMinor), []int{2, 3}},
		{"los defectos de severidad Mayor", project.DefectsBySeverity(SeverityMajor), []int{}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := numerosDeDefectos(tc.got)
			if fmt.Sprint(got) != fmt.Sprint(tc.want) {
				t.Fatalf("se esperaban los defectos %v, salieron %v", tc.want, got)
			}
		})
	}
}
