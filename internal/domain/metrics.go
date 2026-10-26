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
	// completedPoints tiene los story points de cada historia completada, uno por historia,
	// igual que plannedPoints.
	completedPoints []int
	// estimatedHours y actualHours son las horas estimadas y las registradas de las historias
	// cuyas horas cuentan en este ámbito.
	estimatedHours float64
	actualHours    float64
	// defectsDetected y defectsResolved son los defectos que cuentan en este ámbito.
	defectsDetected int
	defectsResolved int
}

// metrics convierte el ámbito en las once métricas. Es el único cálculo, compartido por el
// sprint y el proyecto.
func (sc metricsScope) metrics() Metrics {
	m := Metrics{
		StoriesPlanned:   len(sc.plannedPoints),
		StoriesCompleted: len(sc.completedPoints),
		EstimatedHours:   sc.estimatedHours,
		ActualHours:      sc.actualHours,
		DefectsDetected:  sc.defectsDetected,
		DefectsResolved:  sc.defectsResolved,
	}
	// sin horas estimadas no hay contra qué comparar: la desviación queda en cero, y no es un
	// error (RN-14)
	if sc.estimatedHours > 0 {
		m.HoursDeviation = (sc.actualHours - sc.estimatedHours) / sc.estimatedHours * 100
	}
	for _, points := range sc.plannedPoints {
		m.PlannedPoints += points
	}
	for _, points := range sc.completedPoints {
		m.CompletedPoints += points
	}
	// sin historias planificadas el porcentaje queda en cero, y no es un error (RN-13, RN-14)
	if m.StoriesPlanned > 0 {
		m.CompletionRate = float64(m.StoriesCompleted) / float64(m.StoriesPlanned) * 100
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

	// los planificados son todas las historias comprometidas en el sprint (RN-3); sus horas
	// cuentan solo si este es el último sprint en que estuvo comprometida, así una historia
	// que pasó por dos sprints no las suma dos veces (RN-7)
	points := p.pointsByStory()
	estimated := p.estimatedHoursByStory()
	actual := p.actualHoursByStory()
	var scope metricsScope
	for _, id := range sprint.stories {
		// un sprint cerrado usa los puntos que tenía cada historia al cerrar, así reestimar
		// después una que volvió al backlog no cambia un sprint viejo (RN-3)
		planned := points[id]
		if sprint.state == SprintClosed {
			planned = sprint.pointsAtClose[id]
		}
		scope.plannedPoints = append(scope.plannedPoints, planned)
		if p.lastSprintOf(id) == sprint.number {
			scope.estimatedHours += estimated[id]
			scope.actualHours += actual[id]
		}
	}
	// las completadas son las que quedaron registradas al cerrar, con los puntos que tenían
	// en ese momento; un sprint abierto no tiene ninguna todavía (RN-4, RN-6)
	for _, id := range sprint.completedStories {
		scope.completedPoints = append(scope.completedPoints, sprint.pointsAtClose[id])
	}
	// un defecto cuenta como detectado en su sprint de detección y como resuelto en el de
	// resolución; uno sin resolver tiene 0, que no es el número de ningún sprint (RN-8)
	for _, defect := range p.defects {
		if defect.detectedIn == sprint.number {
			scope.defectsDetected++
		}
		if defect.resolvedIn == sprint.number {
			scope.defectsResolved++
		}
	}
	m := scope.metrics()
	// la velocidad de un sprint son sus puntos completados (RN-5)
	m.Velocity = m.CompletedPoints
	return m, nil
}

// lastSprintOf devuelve el número del último sprint, el de número más alto, en que la
// historia estuvo comprometida (RN-7). Se averigua al calcular y no se guarda (Fase 2).
func (p *Project) lastSprintOf(id int) int {
	last := 0
	for _, sprint := range p.sprints {
		for _, storyID := range sprint.stories {
			if storyID == id && sprint.number > last {
				last = sprint.number
			}
		}
	}
	return last
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

// estimatedHoursByStory devuelve las horas estimadas actuales de cada historia del backlog,
// por su número.
func (p *Project) estimatedHoursByStory() map[int]float64 {
	hours := map[int]float64{}
	for _, story := range p.backlog {
		hours[story.id] = story.estimatedHours
	}
	return hours
}

// actualHoursByStory devuelve las horas registradas sobre cada historia, de cualquier
// integrante y día (RN-7), por su número.
func (p *Project) actualHoursByStory() map[int]float64 {
	hours := map[int]float64{}
	for _, effort := range p.efforts {
		hours[effort.storyID] += effort.hours
	}
	return hours
}

// ProjectMetrics calcula las métricas del proyecto entero, con los datos de este momento y
// sin cambiar nada (RN-1, RN-2).
func (p *Project) ProjectMetrics() Metrics {
	// los planificados son todas las historias del Product Backlog, cada una una vez, hayan
	// entrado a un sprint o no (RN-9). Las horas son las estimadas de todas las historias y
	// todas las registradas, y los defectos, todos, con los Resueltos como resueltos (RN-11)
	var scope metricsScope
	for _, story := range p.backlog {
		scope.plannedPoints = append(scope.plannedPoints, story.storyPoints)
		scope.estimatedHours += story.estimatedHours
	}
	for _, effort := range p.efforts {
		scope.actualHours += effort.hours
	}
	// las completadas son las registradas como completadas en algún sprint cerrado, con los
	// puntos que tenían al cerrarlo; una historia Terminada no vuelve a un sprint, así que
	// cuenta una sola vez (RN-9)
	for _, sprint := range p.sprints {
		for _, id := range sprint.completedStories {
			scope.completedPoints = append(scope.completedPoints, sprint.pointsAtClose[id])
		}
	}
	for _, defect := range p.defects {
		scope.defectsDetected++
		if defect.state == DefectResolved {
			scope.defectsResolved++
		}
	}
	m := scope.metrics()
	// la velocidad del proyecto es el promedio de las velocidades de sus sprints cerrados,
	// redondeado al entero más cercano con las mitades hacia arriba; sin ninguno cerrado, es
	// cero (RN-10). Con enteros: (2*suma + n) / (2*n) redondea así sin pasar por decimales
	closed, total := 0, 0
	for _, sprint := range p.sprints {
		if sprint.state == SprintClosed {
			closed++
			for _, id := range sprint.completedStories {
				total += sprint.pointsAtClose[id]
			}
		}
	}
	if closed > 0 {
		m.Velocity = (2*total + closed) / (2 * closed)
	}
	return m
}
