package domain

// Metrics son las once métricas de un sprint o de un proyecto entero (historia #16). Es un
// resultado calculado, no algo que se guarda: se pide cada vez que hace falta (RN-2).
//
// Es un dato plano y sus campos son públicos, a diferencia de las entidades: no tiene reglas
// que proteger, y la #17 (dashboard), la #18 (reporte) y la #19 (PDF) lo arman a mano en sus
// tests desde fuera del paquete. Se declara antes que su cálculo para que esas historias no
// esperen a que existan el sprint, el esfuerzo y los defectos (contrato, decidido el 21/10).
//
// Las horas, la desviación y el porcentaje van con sus decimales, sin redondear (RN-16):
// redondear es de quien muestra.
type Metrics struct {
	// PlannedPoints son los story points de las historias comprometidas.
	PlannedPoints int
	// CompletedPoints son los story points de las historias completadas.
	CompletedPoints int
	// Velocity es la velocidad: en un sprint, sus puntos completados; en el proyecto, el
	// promedio de sus sprints cerrados redondeado a entero (RN-5, RN-10).
	Velocity int
	// EstimatedHours son las horas estimadas de las historias.
	EstimatedHours float64
	// ActualHours son las horas registradas como esfuerzo sobre esas historias.
	ActualHours float64
	// HoursDeviation es cuánto se apartaron las horas reales de las estimadas, en porcentaje:
	// positiva si se trabajó más de lo estimado (RN-12).
	HoursDeviation float64
	// StoriesPlanned es cuántas historias se comprometieron.
	StoriesPlanned int
	// StoriesCompleted es cuántas se completaron.
	StoriesCompleted int
	// CompletionRate es el porcentaje de historias completas (RN-13).
	CompletionRate float64
	// DefectsDetected es cuántos defectos se encontraron.
	DefectsDetected int
	// DefectsResolved es cuántos se resolvieron.
	DefectsResolved int
}

// metricsScope es el ámbito de un cálculo: qué cuenta en él (Fase 2, "Elegir el ámbito").
// Las dos consultas, la del sprint y la del proyecto, lo arman distinto y después comparten
// un único cálculo, así las fórmulas están escritas una sola vez.
type metricsScope struct {
	// plannedPoints tiene los story points de cada historia planificada, uno por historia:
	// la cantidad de historias es su largo. Una historia sin estimar aporta 0 (RN-16).
	plannedPoints []int
}

// metrics convierte el ámbito en las once métricas. Es el único cálculo, compartido por el
// sprint y el proyecto.
func (sc metricsScope) metrics() Metrics {
	m := Metrics{StoriesPlanned: len(sc.plannedPoints)}
	for _, points := range sc.plannedPoints {
		m.PlannedPoints += points
	}
	return m
}

// SprintMetrics calcula las métricas del sprint con ese número, con los datos de este
// momento y sin cambiar nada (RN-1, RN-2). Si el proyecto no tiene ese sprint, informa que
// no se encontró (RN-20).
func (p *Project) SprintMetrics(number int) (Metrics, error) {
	sprint, err := p.Sprint(number)
	if err != nil {
		return Metrics{}, err
	}

	// los planificados son todas las historias comprometidas en el sprint (RN-3)
	points := p.pointsByStory()
	var scope metricsScope
	for _, id := range sprint.stories {
		scope.plannedPoints = append(scope.plannedPoints, points[id])
	}
	return scope.metrics(), nil
}

// pointsByStory devuelve los story points actuales de cada historia del backlog, por su
// número.
func (p *Project) pointsByStory() map[int]int {
	points := map[int]int{}
	for _, story := range p.backlog {
		points[story.id] = story.storyPoints
	}
	return points
}

// ProjectMetrics calcula las métricas del proyecto entero, con los datos de este momento y
// sin cambiar nada (RN-1, RN-2).
func (p *Project) ProjectMetrics() Metrics {
	return Metrics{}
}
