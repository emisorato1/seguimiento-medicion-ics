package domain

import (
	"fmt"
	"slices"
	"time"
)

// ReportScope dice de que es un reporte: del proyecto entero o de uno de sus sprints
// (historia #18, RN-3).
type ReportScope int

const (
	// ScopeProject: el reporte del proyecto entero.
	ScopeProject ReportScope = iota
	// ScopeSprint: el reporte de un sprint.
	ScopeSprint
)

// Report es un reporte ya armado (historia #18). Son datos planos, como Metrics y el
// Dashboard: se arma de una vez con los datos de ese momento y no se guarda (Fase 2). Las
// metricas son las del mismo alcance, tal como las da la #16 (RN-2), y las secciones traen
// el texto ya escrito, que la exportacion a PDF de la #19 copia sin tocar.
type Report struct {
	Title       string
	GeneratedOn time.Time
	Scope       ReportScope
	Metrics     Metrics
	Sections    []Section
}

// Section es una seccion del reporte: su encabezado y sus lineas de texto, en orden.
type Section struct {
	Heading string
	Lines   []string
}

// SprintReport genera el reporte del sprint con ese numero, con los datos de este momento,
// sin guardarlo y sin cambiar nada (RN-15). Le pide las metricas del sprint a la #16 y se
// las pasa al armado (RN-2). Si el proyecto no tiene ese sprint, informa que no se encontro
// (RN-16). El dia de generacion lo indica quien lo pide.
func (p *Project) SprintReport(number int, today time.Time) (Report, error) {
	metrics, err := p.SprintMetrics(number)
	if err != nil {
		return Report{}, err
	}
	// si las metricas salieron, el sprint existe
	sprint, _ := p.Sprint(number)
	return p.buildSprintReport(sprint, metrics, today), nil
}

// ProjectReport genera el reporte del proyecto entero con las metricas del proyecto de la
// #16, igual que SprintReport (RN-2, RN-15). No informa errores: se le pide a un proyecto
// que ya se encontro, y el que no existe lo informa el registro (#3).
func (p *Project) ProjectReport(today time.Time) Report {
	return p.buildProjectReport(p.ProjectMetrics(), today)
}

// buildSprintReport arma el reporte del sprint con las metricas que recibe, sin pedirlas
// ni calcular ninguna (RN-2), y con el dia que le pasa quien lo pide: el dominio no lee el
// reloj. Tiene siempre las cuatro secciones, en el orden de la RN-4.
func (p *Project) buildSprintReport(sprint *Sprint, metrics Metrics, today time.Time) Report {
	return Report{
		Title:       fmt.Sprintf("Reporte del sprint %d — %s", sprint.number, p.name),
		GeneratedOn: today,
		Scope:       ScopeSprint,
		Metrics:     metrics,
		Sections: []Section{
			{Heading: "Resumen", Lines: p.sprintSummaryLines(sprint)},
			{Heading: "Historias", Lines: p.sprintStoryLines(sprint, metrics)},
			{Heading: "Estimación y esfuerzo", Lines: effortLines(metrics, "Velocidad: "+formatCount(metrics.Velocity))},
			{Heading: "Defectos", Lines: p.defectLines(metrics, "No hay defectos en este sprint.", func(d *Defect) bool {
				// uno detectado en un sprint y resuelto en otro aparece en los dos (RN-12)
				return d.detectedIn == sprint.number || d.resolvedIn == sprint.number
			})},
		},
	}
}

// buildProjectReport arma el reporte del proyecto entero con las metricas del proyecto que
// recibe, igual que el de un sprint: las mismas cuatro secciones y las mismas piezas, con lo
// que cambia el alcance (RN-1, RN-8, RN-10, RN-11, RN-12).
func (p *Project) buildProjectReport(metrics Metrics, today time.Time) Report {
	return Report{
		Title:       "Reporte del proyecto " + p.name,
		GeneratedOn: today,
		Scope:       ScopeProject,
		Metrics:     metrics,
		Sections: []Section{
			{Heading: "Resumen", Lines: p.projectSummaryLines()},
			{Heading: "Historias", Lines: p.projectStoryLines(metrics)},
			{Heading: "Estimación y esfuerzo", Lines: effortLines(metrics, p.averageVelocityLine(metrics))},
			{Heading: "Defectos", Lines: p.defectLines(metrics, "No hay defectos en este proyecto.", func(*Defect) bool { return true })},
		},
	}
}

// projectSummaryLines escribe el Resumen del reporte del proyecto: el nombre, las fechas y
// cuantos sprints tiene y cuantos estan cerrados (RN-8). La fecha de fin es optativa (#3):
// sin ella dice "sin fecha de fin".
func (p *Project) projectSummaryLines() []string {
	end := "sin fecha de fin"
	if p.HasEnd() {
		end = formatDate(p.end)
	}
	sprints := "sprints"
	if len(p.sprints) == 1 {
		sprints = "sprint"
	}
	return []string{
		"Proyecto: " + p.name,
		"Inicio: " + formatDate(p.start),
		"Fin: " + end,
		fmt.Sprintf("%d %s, %d cerrados", len(p.sprints), sprints, len(p.ClosedSprints())),
	}
}

// averageVelocityLine escribe la velocidad promedio del proyecto. Mientras no hay ningun
// sprint cerrado dice "sin sprints cerrados", como el dashboard, y no 0: un cero haria
// parecer que el equipo no hizo nada (RN-11).
func (p *Project) averageVelocityLine(metrics Metrics) string {
	if len(p.ClosedSprints()) == 0 {
		return "Velocidad promedio: sin sprints cerrados"
	}
	return "Velocidad promedio: " + formatCount(metrics.Velocity)
}

// projectStoryLines escribe la seccion Historias del reporte del proyecto: los totales y una
// linea por historia del backlog, con su estado actual (RN-10).
func (p *Project) projectStoryLines(metrics Metrics) []string {
	lines := storyTotals(metrics)
	for _, story := range p.backlog {
		lines = append(lines, storyLine(story, story.storyPoints, story.state.String()))
	}
	return withEmptyNotice(lines, len(p.backlog), "No hay historias en este proyecto.")
}

// sprintSummaryLines escribe el Resumen del reporte de un sprint: el proyecto, el Sprint Goal,
// las fechas y el estado (RN-6). Mientras el sprint esta abierto avisa que los numeros son
// parciales (RN-7).
func (p *Project) sprintSummaryLines(sprint *Sprint) []string {
	state := "Estado: abierto. Los números son parciales."
	if sprint.state == SprintClosed {
		state = "Estado: cerrado el " + formatDate(sprint.closedOn)
	}
	return []string{
		"Proyecto: " + p.name,
		"Sprint Goal: " + sprint.goal,
		"Inicio: " + formatDate(sprint.start),
		"Fin: " + formatDate(sprint.end),
		state,
	}
}

// sprintStoryLines escribe la seccion Historias del reporte de un sprint: los totales de las
// metricas y una linea por historia comprometida (RN-9). Que se completo sale de lo que
// registro el cierre de ese sprint, y no del estado actual de la historia, que pudo
// terminarse despues en otro sprint (RN-10, punto 2 de la puerta). Los puntos tambien son
// los que tenia al cerrar (#10). Mientras el sprint esta abierto no hay cierre: cada
// historia va con sus puntos actuales y "en curso".
func (p *Project) sprintStoryLines(sprint *Sprint, metrics Metrics) []string {
	lines := storyTotals(metrics)
	for _, id := range sprint.stories {
		story, _ := p.Story(id)
		points, situation := story.storyPoints, "en curso"
		if sprint.state == SprintClosed {
			points, situation = sprint.pointsAtClose[id], "no completada"
			if slices.Contains(sprint.completedStories, id) {
				situation = "completada"
			}
		}
		lines = append(lines, storyLine(story, points, situation))
	}
	return withEmptyNotice(lines, len(sprint.stories), "No hay historias en este sprint.")
}

// storyTotals son las dos primeras lineas de la seccion Historias: las planificadas y las
// completadas con su porcentaje, tal como vienen en las metricas (RN-2, RN-9).
func storyTotals(metrics Metrics) []string {
	return []string{
		"Historias planificadas: " + formatCount(metrics.StoriesPlanned),
		fmt.Sprintf("Historias completadas: %s (%s)", formatCount(metrics.StoriesCompleted), formatPercent(metrics.CompletionRate)),
	}
}

// storyLine escribe la linea de una historia: su numero, su titulo, sus puntos y su situacion
// (RN-9).
func storyLine(story *Story, points int, situation string) string {
	return fmt.Sprintf("#%d — %s — %s — %s", story.id, story.title, pointsText(points), situation)
}

// withEmptyNotice agrega el aviso de que no hay nada que listar cuando la seccion quedo solo
// con sus totales (RN-5).
func withEmptyNotice(lines []string, listed int, notice string) []string {
	if listed == 0 {
		lines = append(lines, notice)
	}
	return lines
}

// pointsText escribe los story points de una historia, o "sin estimar" si no tiene: la
// escala no tiene el 0, asi que 0 es siempre una historia sin estimar (RN-9).
func pointsText(points int) string {
	if points == 0 {
		return "sin estimar"
	}
	return formatCount(points) + " SP"
}

// defectLines escribe la seccion Defectos: los totales de las metricas y una linea por cada
// defecto que entra en el reporte, con su severidad, su estado y, si esta resuelto, el sprint
// en que se resolvio (RN-12). Que defectos entran lo decide cada reporte.
func (p *Project) defectLines(metrics Metrics, empty string, include func(*Defect) bool) []string {
	lines := []string{
		"Defectos detectados: " + formatCount(metrics.DefectsDetected),
		"Defectos resueltos: " + formatCount(metrics.DefectsResolved),
	}
	listed := 0
	for _, defect := range p.defects {
		if !include(defect) {
			continue
		}
		state := defect.state.String()
		if defect.state == DefectResolved {
			state = fmt.Sprintf("%s en el sprint %d", state, defect.resolvedIn)
		}
		lines = append(lines, fmt.Sprintf("#%d — %s — %s — %s", defect.id, defect.title, defect.severity, state))
		listed++
	}
	return withEmptyNotice(lines, listed, empty)
}

// effortLines escribe la seccion Estimacion y esfuerzo con los numeros de las metricas, sin
// calcular ninguno (RN-2, RN-11), y con la regla de formato de format.go (RN-14). La linea
// de la velocidad la escribe cada reporte: "Velocidad" en un sprint y "Velocidad promedio"
// en el proyecto, que sin sprints cerrados no tiene un numero que mostrar.
func effortLines(m Metrics, velocity string) []string {
	return []string{
		"Story points planificados: " + formatCount(m.PlannedPoints),
		"Story points completados: " + formatCount(m.CompletedPoints),
		velocity,
		"Horas estimadas: " + formatHours(m.EstimatedHours),
		"Horas reales: " + formatHours(m.ActualHours),
		"Desviación: " + formatPercent(m.HoursDeviation),
	}
}
