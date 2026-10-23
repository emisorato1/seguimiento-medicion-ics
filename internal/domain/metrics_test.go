package domain

import (
	"errors"
	"fmt"
	"testing"
)

// Escenario: Pedir algo que no existe
// Cubre: RN-20
//
// La fila de la historia se adelantó a la tarea 16 para que la carga de horas estimadas
// quedara cubierta entera. Las métricas del proyecto 99 se piden como las pediría la web: el
// proyecto se busca en el registro, y si no existe no se llega a pedir nada.
func TestPedirAlgoQueNoExiste(t *testing.T) {
	cases := []struct {
		name string
		run  func(project *Project) error
		want error
	}{
		{"pedir las metricas del sprint 99",
			func(project *Project) error { _, err := project.SprintMetrics(99); return err }, ErrSprintNotFound},
		{"cargar horas estimadas en la historia 99",
			func(project *Project) error { return project.EstimateHours(99, 4) }, ErrStoryNotFound},
		{"pedir las metricas del proyecto 99",
			func(project *Project) error { _, err := NewProjectRegistry().Get(99); return err }, ErrProjectNotFound},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			project := unProyecto(t)

			err := tc.run(project)

			if !errors.Is(err, tc.want) {
				t.Fatalf("se esperaba %v, salio: %v", tc.want, err)
			}
		})
	}
}

// unSprint crea el sprint siguiente del proyecto, abierto y sin historias.
func unSprint(t *testing.T, project *Project) *Sprint {
	t.Helper()
	sprint, err := project.AddSprint(SprintData{Goal: "Tener el MVP", Start: unDia(2026, 10, 26), End: unDia(2026, 10, 30)})
	if err != nil {
		t.Fatalf("no se pudo crear el sprint de prueba: %v", err)
	}
	return sprint
}

// Esquema del escenario: Sin datos cargados las metricas valen cero
// Cubre: RN-1, RN-14, RN-15
func TestSinDatosCargadosLasMetricasValenCero(t *testing.T) {
	cases := []struct {
		name string
		ask  func(t *testing.T, project *Project) (Metrics, error)
	}{
		{"un proyecto con el sprint 1 recien abierto", func(t *testing.T, project *Project) (Metrics, error) {
			unSprint(t, project)
			return project.SprintMetrics(1)
		}},
		{"un proyecto sin sprints", func(t *testing.T, project *Project) (Metrics, error) {
			return project.ProjectMetrics(), nil
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			project := unProyecto(t)

			got, err := tc.ask(t, project)

			if err != nil {
				t.Fatalf("se esperaba que no fuera un error, salio: %v", err)
			}
			if got != (Metrics{}) {
				t.Fatalf("se esperaban las once metricas en cero, salio %+v", got)
			}
		})
	}
}

// historiasEstimadas carga en el backlog una historia por cada valor, estimada con esos
// story points; un 0 la deja sin estimar.
func historiasEstimadas(t *testing.T, project *Project, points ...int) {
	t.Helper()
	for _, p := range points {
		story, err := project.AddStory(unosDatos())
		if err != nil {
			t.Fatalf("no se pudo crear la historia de prueba: %v", err)
		}
		if p == 0 {
			continue
		}
		if err := project.EstimateStory(story.ID(), p); err != nil {
			t.Fatalf("no se pudo estimar la historia de prueba en %d: %v", p, err)
		}
	}
}

// conIntegrante agrega al proyecto la integrante "Juliana Bustos", para poder cargarle horas
// con cargar (de la #14).
func conIntegrante(t *testing.T, project *Project) {
	t.Helper()
	if _, err := project.AddMember("Juliana Bustos", RoleProductBuilder); err != nil {
		t.Fatalf("no se pudo agregar la integrante de prueba: %v", err)
	}
}

// conHorasEstimadas carga las horas estimadas de la historia o corta el test.
func conHorasEstimadas(t *testing.T, project *Project, id int, hours float64) {
	t.Helper()
	if err := project.EstimateHours(id, hours); err != nil {
		t.Fatalf("no se pudieron cargar las horas estimadas de la historia %d: %v", id, err)
	}
}

// Escenario: Metricas de un sprint cerrado
// Cubre: RN-3, RN-7, RN-8, RN-12, RN-16
//
// El escenario se completa en varias tareas (Fase 3): por ahora verifica los planificados,
// las horas, la desviación y los defectos, y las tareas 6 y 11 le suman su número. Asignar historias es de la #9 y
// cerrar el sprint de la #10, que todavía no están en dev: el test arma el sprint dentro del
// paquete.
func TestMetricasDeUnSprintCerrado(t *testing.T) {
	project := unProyecto(t)
	conIntegrante(t, project)
	historiasEstimadas(t, project, 3, 5, 2, 3)
	for i, hours := range []float64{4, 6, 2, 4} {
		conHorasEstimadas(t, project, i+1, hours)
	}
	for i, hours := range []float64{5, 8, 3, 4} {
		cargar(t, project, i+1, "Juliana Bustos", hours)
	}
	sprint := unSprint(t, project)
	sprint.stories = []int{1, 2, 3, 4}
	sprint.state = SprintClosed
	for i := 0; i < 2; i++ {
		if _, err := project.AddDefect(unDefecto()); err != nil {
			t.Fatalf("no se pudo registrar el defecto de prueba: %v", err)
		}
	}
	if err := project.ResolveDefect(1, 1); err != nil {
		t.Fatalf("no se pudo resolver el defecto de prueba: %v", err)
	}

	got, err := project.SprintMetrics(1)

	if err != nil {
		t.Fatalf("se esperaba que no fuera un error, salio: %v", err)
	}
	if got.PlannedPoints != 13 || got.StoriesPlanned != 4 {
		t.Fatalf("se esperaban 13 story points y 4 historias planificadas, salieron %d y %d", got.PlannedPoints, got.StoriesPlanned)
	}
	if !mismasHoras(got.EstimatedHours, 16) || !mismasHoras(got.ActualHours, 20) {
		t.Fatalf("se esperaban 16 horas estimadas y 20 reales, salieron %v y %v", got.EstimatedHours, got.ActualHours)
	}
	if !mismasHoras(got.HoursDeviation, 25) {
		t.Fatalf("se esperaba una desviacion de 25, salio %v", got.HoursDeviation)
	}
	if got.DefectsDetected != 2 || got.DefectsResolved != 1 {
		t.Fatalf("se esperaban 2 defectos detectados y 1 resuelto, salieron %d y %d", got.DefectsDetected, got.DefectsResolved)
	}
}

// Escenario: Las horas de una historia que paso por dos sprints cuentan en el ultimo
// Cubre: RN-7
//
// La historia 2 no se completó en el sprint 1 y volvió a comprometerse en el 2. Sus horas
// cuentan una sola vez, en el último sprint en que estuvo, aunque se hayan cargado antes.
func TestLasHorasDeUnaHistoriaQuePasoPorDosSprintsCuentanEnElUltimo(t *testing.T) {
	project := unProyecto(t)
	conIntegrante(t, project)
	historiasEstimadas(t, project, 3, 5)
	conHorasEstimadas(t, project, 2, 6)
	cargar(t, project, 2, "Juliana Bustos", 3)
	cargar(t, project, 2, "Juliana Bustos", 4)
	first := unSprint(t, project)
	first.stories = []int{2}
	first.state = SprintClosed
	second := unSprint(t, project)
	second.stories = []int{2}

	cases := []struct {
		sprint            int
		estimated, actual float64
	}{
		{1, 0, 0},
		{2, 6, 7},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("sprint %d", tc.sprint), func(t *testing.T) {
			got, err := project.SprintMetrics(tc.sprint)

			if err != nil {
				t.Fatalf("se esperaba que no fuera un error, salio: %v", err)
			}
			if !mismasHoras(got.EstimatedHours, tc.estimated) || !mismasHoras(got.ActualHours, tc.actual) {
				t.Fatalf("se esperaban %v horas estimadas y %v reales, salieron %v y %v", tc.estimated, tc.actual, got.EstimatedHours, got.ActualHours)
			}
		})
	}
}

// Esquema del escenario: Calcular la desviacion de horas
// Cubre: RN-12, RN-14
//
// Cada fila arma un sprint con una sola historia que tiene esas horas. Cero horas estimadas
// es no cargarlas, porque la carga rechaza el cero (RN-18), y cero reales es no registrar
// ninguna.
func TestCalcularLaDesviacionDeHoras(t *testing.T) {
	cases := []struct {
		estimated, actual, deviation float64
	}{
		{16, 20, 25},
		{10, 8, -20},
		{8, 0, -100},
		{0, 5, 0},
		{0, 0, 0},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v estimadas y %v reales", tc.estimated, tc.actual), func(t *testing.T) {
			project := unProyecto(t)
			conIntegrante(t, project)
			historiasEstimadas(t, project, 3)
			if tc.estimated > 0 {
				conHorasEstimadas(t, project, 1, tc.estimated)
			}
			if tc.actual > 0 {
				cargar(t, project, 1, "Juliana Bustos", tc.actual)
			}
			unSprint(t, project).stories = []int{1}

			got, err := project.SprintMetrics(1)

			if err != nil {
				t.Fatalf("se esperaba que no fuera un error, salio: %v", err)
			}
			if !mismasHoras(got.HoursDeviation, tc.deviation) {
				t.Fatalf("se esperaba una desviacion de %v, salio %v", tc.deviation, got.HoursDeviation)
			}
		})
	}
}

// Escenario: Las metricas de un sprint abierto son parciales
// Cubre: RN-6
//
// Un sprint abierto todavía no tiene registro de cierre, así que sus completados y su
// velocidad valen cero, y eso no es un error. Lo planificado y las horas sí se calculan.
func TestLasMetricasDeUnSprintAbiertoSonParciales(t *testing.T) {
	project := unProyecto(t)
	conIntegrante(t, project)
	historiasEstimadas(t, project, 3)
	cargar(t, project, 1, "Juliana Bustos", 2)
	unSprint(t, project).stories = []int{1}

	got, err := project.SprintMetrics(1)

	if err != nil {
		t.Fatalf("se esperaba que no fuera un error, salio: %v", err)
	}
	if got.PlannedPoints != 3 || got.CompletedPoints != 0 || got.Velocity != 0 {
		t.Fatalf("se esperaban 3 story points planificados, 0 completados y velocidad 0, salieron %d, %d y %d",
			got.PlannedPoints, got.CompletedPoints, got.Velocity)
	}
	if !mismasHoras(got.ActualHours, 2) {
		t.Fatalf("se esperaban 2 horas reales, salieron %v", got.ActualHours)
	}
}

// Escenario: Defectos detectados y resueltos de cada sprint
// Cubre: RN-8
//
// Un defecto cuenta como detectado en su sprint de detección y como resuelto en el de
// resolución, aunque sean sprints distintos.
func TestDefectosDetectadosYResueltosDeCadaSprint(t *testing.T) {
	project := unProyectoConDefectos(t, 2)
	for _, detectedIn := range []int{1, 2, 2} {
		data := unDefecto()
		data.DetectedIn = detectedIn
		if _, err := project.AddDefect(data); err != nil {
			t.Fatalf("no se pudo registrar el defecto de prueba: %v", err)
		}
	}
	for _, id := range []int{1, 2} {
		if err := project.ResolveDefect(id, 2); err != nil {
			t.Fatalf("no se pudo resolver el defecto de prueba %d: %v", id, err)
		}
	}

	cases := []struct {
		sprint, detected, resolved int
	}{
		{1, 1, 0},
		{2, 2, 2},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("sprint %d", tc.sprint), func(t *testing.T) {
			got, err := project.SprintMetrics(tc.sprint)

			if err != nil {
				t.Fatalf("se esperaba que no fuera un error, salio: %v", err)
			}
			if got.DefectsDetected != tc.detected || got.DefectsResolved != tc.resolved {
				t.Fatalf("se esperaban %d defectos detectados y %d resueltos, salieron %d y %d",
					tc.detected, tc.resolved, got.DefectsDetected, got.DefectsResolved)
			}
		})
	}
}

// Escenario: Metricas del proyecto entero
// Cubre: RN-9, RN-11, RN-12
//
// El escenario se completa en varias tareas (Fase 3): por ahora verifica los planificados,
// las horas, la desviación y los defectos. Los completados, el porcentaje y la velocidad salen
// del registro del cierre de la #10 y los suman las tareas 6, 11 y 14. Asignar historias es de
// la #9 y cerrar el sprint de la #10: el test arma los sprints dentro del paquete.
func TestMetricasDelProyectoEntero(t *testing.T) {
	project := unProyecto(t)
	conIntegrante(t, project)
	historiasEstimadas(t, project, 3, 5, 2, 3, 8)
	for i, hours := range []float64{4, 6, 2, 8, 10} {
		conHorasEstimadas(t, project, i+1, hours)
	}
	for i, hours := range []float64{5, 8, 3, 7, 10} {
		cargar(t, project, i+1, "Juliana Bustos", hours)
	}
	first := unSprint(t, project)
	first.stories = []int{1, 2, 3}
	first.state = SprintClosed
	unSprint(t, project).stories = []int{3}
	for _, detectedIn := range []int{1, 1, 2} {
		data := unDefecto()
		data.DetectedIn = detectedIn
		if _, err := project.AddDefect(data); err != nil {
			t.Fatalf("no se pudo registrar el defecto de prueba: %v", err)
		}
	}
	if err := project.ResolveDefect(1, 2); err != nil {
		t.Fatalf("no se pudo resolver el defecto de prueba: %v", err)
	}

	got := project.ProjectMetrics()

	if got.PlannedPoints != 21 || got.StoriesPlanned != 5 {
		t.Fatalf("se esperaban 21 story points y 5 historias planificadas, salieron %d y %d", got.PlannedPoints, got.StoriesPlanned)
	}
	if !mismasHoras(got.EstimatedHours, 30) || !mismasHoras(got.ActualHours, 33) || !mismasHoras(got.HoursDeviation, 10) {
		t.Fatalf("se esperaban 30 horas estimadas, 33 reales y una desviacion de 10, salieron %v, %v y %v",
			got.EstimatedHours, got.ActualHours, got.HoursDeviation)
	}
	if got.DefectsDetected != 3 || got.DefectsResolved != 1 {
		t.Fatalf("se esperaban 3 defectos detectados y 1 resuelto, salieron %d y %d", got.DefectsDetected, got.DefectsResolved)
	}
}

// Escenario: Las metricas reflejan los datos cargados despues
// Cubre: RN-2
//
// Las métricas se calculan cada vez que se piden: la segunda vez ven las horas cargadas en el
// medio. Y pedirlas no cambia nada: el registro entero, convertido a JSON, queda igual.
func TestLasMetricasReflejanLosDatosCargadosDespues(t *testing.T) {
	registry := NewProjectRegistry()
	project, err := registry.Create(ProjectData{Name: "Sistema de metricas", Start: unDia(2026, 10, 1)})
	if err != nil {
		t.Fatalf("no se pudo crear el proyecto de prueba: %v", err)
	}
	conIntegrante(t, project)
	historiasEstimadas(t, project, 3)
	cargar(t, project, 1, "Juliana Bustos", 5)
	sprint := unSprint(t, project)
	sprint.stories = []int{1}
	sprint.state = SprintClosed
	before, err := registry.ToJSON()
	if err != nil {
		t.Fatalf("no se pudieron convertir los datos: %v", err)
	}

	first, err := project.SprintMetrics(1)
	if err != nil {
		t.Fatalf("se esperaba que no fuera un error, salio: %v", err)
	}
	project.ProjectMetrics()
	after, err := registry.ToJSON()
	if err != nil {
		t.Fatalf("no se pudieron convertir los datos: %v", err)
	}
	cargar(t, project, 1, "Juliana Bustos", 2)
	second, err := project.SprintMetrics(1)
	if err != nil {
		t.Fatalf("se esperaba que no fuera un error, salio: %v", err)
	}

	if !mismasHoras(first.ActualHours, 5) || !mismasHoras(second.ActualHours, 7) {
		t.Fatalf("se esperaban 5 horas reales la primera vez y 7 la segunda, salieron %v y %v", first.ActualHours, second.ActualHours)
	}
	if string(after) != string(before) {
		t.Fatalf("se esperaba que pedir las metricas no cambiara ningun dato, y el registro cambio")
	}
}
