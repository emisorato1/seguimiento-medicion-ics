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
// Cubre: RN-3, RN-4, RN-5, RN-7, RN-8, RN-12, RN-13, RN-16
//
// Desde la tarea 7 el sprint se cierra con el cierre de la #10, que registra las completadas y
// los puntos de cada historia.
func TestMetricasDeUnSprintCerrado(t *testing.T) {
	project, _ := unSprintParaCerrar(t, 3, 5, 2, 3)
	conIntegrante(t, project)
	for i, hours := range []float64{4, 6, 2, 4} {
		conHorasEstimadas(t, project, i+1, hours)
	}
	for i, hours := range []float64{5, 8, 3, 4} {
		cargar(t, project, i+1, "Juliana Bustos", hours)
	}
	for i := 0; i < 2; i++ {
		if _, err := project.AddDefect(unDefecto()); err != nil {
			t.Fatalf("no se pudo registrar el defecto de prueba: %v", err)
		}
	}
	if err := project.CloseSprint([]int{1, 2, 3}, unDia(2026, 10, 30)); err != nil {
		t.Fatalf("no se pudo cerrar el sprint de prueba: %v", err)
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
	if got.CompletedPoints != 10 || got.Velocity != 10 || got.StoriesCompleted != 3 {
		t.Fatalf("se esperaban 10 puntos completados, velocidad 10 y 3 historias completadas, salieron %d, %d y %d",
			got.CompletedPoints, got.Velocity, got.StoriesCompleted)
	}
	if !mismasHoras(got.CompletionRate, 75) {
		t.Fatalf("se esperaba un porcentaje de historias completas de 75, salio %v", got.CompletionRate)
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
// Cubre: RN-9, RN-10, RN-11, RN-12, RN-13
//
// El sprint 1 se cierra con el cierre de la #10, que registra las completadas.
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
	unSprint(t, project)
	for _, id := range []int{1, 2, 3} {
		if err := project.AssignStory(id); err != nil {
			t.Fatalf("no se pudo asignar la historia %d al sprint 1: %v", id, err)
		}
	}
	for i := 0; i < 2; i++ {
		if _, err := project.AddDefect(unDefecto()); err != nil {
			t.Fatalf("no se pudo registrar el defecto de prueba: %v", err)
		}
	}
	if err := project.CloseSprint([]int{1, 2}, unDia(2026, 10, 30)); err != nil {
		t.Fatalf("no se pudo cerrar el sprint 1 de prueba: %v", err)
	}
	unSprint(t, project)
	if err := project.AssignStory(3); err != nil {
		t.Fatalf("no se pudo asignar la historia 3 al sprint 2: %v", err)
	}
	data := unDefecto()
	data.DetectedIn = 2
	if _, err := project.AddDefect(data); err != nil {
		t.Fatalf("no se pudo registrar el defecto de prueba: %v", err)
	}
	if err := project.ResolveDefect(1, 2); err != nil {
		t.Fatalf("no se pudo resolver el defecto de prueba: %v", err)
	}

	got := project.ProjectMetrics()

	if got.PlannedPoints != 21 || got.StoriesPlanned != 5 {
		t.Fatalf("se esperaban 21 story points y 5 historias planificadas, salieron %d y %d", got.PlannedPoints, got.StoriesPlanned)
	}
	if got.CompletedPoints != 8 || got.StoriesCompleted != 2 || !mismasHoras(got.CompletionRate, 40) {
		t.Fatalf("se esperaban 8 puntos completados, 2 historias completadas y un porcentaje de 40, salieron %d, %d y %v",
			got.CompletedPoints, got.StoriesCompleted, got.CompletionRate)
	}
	if got.Velocity != 8 {
		t.Fatalf("se esperaba velocidad 8, salio %d", got.Velocity)
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

// Escenario: Una historia que volvio al backlog cuenta como completada solo donde se completo
// Cubre: RN-4, RN-5
//
// Los dos sprints se cierran con el cierre de la #10: la historia 2 no se completó en el 1,
// volvió al backlog, se comprometió en el 2 y se completó ahí.
func TestUnaHistoriaQueVolvioAlBacklogCuentaComoCompletadaSoloDondeSeCompleto(t *testing.T) {
	project, _ := unSprintParaCerrar(t, 3, 5)
	if err := project.CloseSprint([]int{1}, unDia(2026, 10, 30)); err != nil {
		t.Fatalf("no se pudo cerrar el sprint 1 de prueba: %v", err)
	}
	unSprint(t, project)
	if err := project.AssignStory(2); err != nil {
		t.Fatalf("no se pudo asignar la historia 2 al sprint 2: %v", err)
	}
	if err := project.CloseSprint([]int{2}, unDia(2026, 10, 31)); err != nil {
		t.Fatalf("no se pudo cerrar el sprint 2 de prueba: %v", err)
	}

	cases := []struct {
		sprint, planned, completed int
	}{
		{1, 8, 3},
		{2, 5, 5},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("sprint %d", tc.sprint), func(t *testing.T) {
			got, err := project.SprintMetrics(tc.sprint)

			if err != nil {
				t.Fatalf("se esperaba que no fuera un error, salio: %v", err)
			}
			if got.PlannedPoints != tc.planned || got.CompletedPoints != tc.completed || got.Velocity != tc.completed {
				t.Fatalf("se esperaban %d puntos planificados, %d completados y velocidad %d, salieron %d, %d y %d",
					tc.planned, tc.completed, tc.completed, got.PlannedPoints, got.CompletedPoints, got.Velocity)
			}
		})
	}
}

// Escenario: Los puntos de un sprint cerrado no cambian si se reestima una historia
// Cubre: RN-3
func TestLosPuntosDeUnSprintCerradoNoCambianSiSeReestimaUnaHistoria(t *testing.T) {
	project, _ := unSprintParaCerrar(t, 3, 5)
	if err := project.CloseSprint([]int{1}, unDia(2026, 10, 30)); err != nil {
		t.Fatalf("no se pudo cerrar el sprint de prueba: %v", err)
	}

	if err := project.EstimateStory(2, 8); err != nil {
		t.Fatalf("no se pudo reestimar la historia 2: %v", err)
	}
	got, err := project.SprintMetrics(1)

	if err != nil {
		t.Fatalf("se esperaba que no fuera un error, salio: %v", err)
	}
	if got.PlannedPoints != 8 || got.CompletedPoints != 3 {
		t.Fatalf("se esperaban 8 story points planificados y 3 completados, salieron %d y %d", got.PlannedPoints, got.CompletedPoints)
	}
}

// Esquema del escenario: Calcular el porcentaje de historias completas
// Cubre: RN-13, RN-14, RN-16
func TestCalcularElPorcentajeDeHistoriasCompletas(t *testing.T) {
	cases := []struct {
		planned, completed int
		rate               float64
	}{
		{4, 3, 75},
		{8, 1, 12.5},
		{3, 3, 100},
		{0, 0, 0},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d planificadas y %d completadas", tc.planned, tc.completed), func(t *testing.T) {
			points := make([]int, tc.planned)
			completed := []int{}
			for i := range points {
				points[i] = 3
				if i < tc.completed {
					completed = append(completed, i+1)
				}
			}
			project, _ := unSprintParaCerrar(t, points...)
			if err := project.CloseSprint(completed, unDia(2026, 10, 30)); err != nil {
				t.Fatalf("no se pudo cerrar el sprint de prueba: %v", err)
			}

			got, err := project.SprintMetrics(1)

			if err != nil {
				t.Fatalf("se esperaba que no fuera un error, salio: %v", err)
			}
			if !mismasHoras(got.CompletionRate, tc.rate) {
				t.Fatalf("se esperaba un porcentaje de %v, salio %v", tc.rate, got.CompletionRate)
			}
		})
	}
}

// Esquema del escenario: Velocidad del proyecto
// Cubre: RN-10
//
// Cada sprint cerrado completa historias que suman su velocidad, con puntos de la escala: el 6
// son dos historias de 3. La fila "ninguna" tiene un sprint abierto, que no cuenta.
func TestVelocidadDelProyecto(t *testing.T) {
	cases := []struct {
		name     string
		sprints  [][]int
		velocity int
	}{
		{"8, 5 y 6", [][]int{{8}, {5}, {3, 3}}, 6},
		{"10 y 13", [][]int{{5, 5}, {13}}, 12},
		{"ninguna", nil, 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			project := unProyecto(t)
			next := 1
			for _, points := range tc.sprints {
				historiasEstimadas(t, project, points...)
				unSprint(t, project)
				completed := []int{}
				for range points {
					if err := project.AssignStory(next); err != nil {
						t.Fatalf("no se pudo asignar la historia %d: %v", next, err)
					}
					completed = append(completed, next)
					next++
				}
				if err := project.CloseSprint(completed, unDia(2026, 10, 30)); err != nil {
					t.Fatalf("no se pudo cerrar el sprint de prueba: %v", err)
				}
			}
			if tc.sprints == nil {
				historiasEstimadas(t, project, 5)
				unSprint(t, project)
				if err := project.AssignStory(1); err != nil {
					t.Fatalf("no se pudo asignar la historia al sprint abierto: %v", err)
				}
			}

			got := project.ProjectMetrics()

			if got.Velocity != tc.velocity {
				t.Fatalf("se esperaba velocidad %d, salio %d", tc.velocity, got.Velocity)
			}
		})
	}
}
