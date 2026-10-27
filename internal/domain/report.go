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
			{Heading: "Estimación y esfuerzo", Lines: effortLines(metrics)},
			{Heading: "Defectos"},
		},
	}
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
	lines := []string{
		"Historias planificadas: " + formatCount(metrics.StoriesPlanned),
		fmt.Sprintf("Historias completadas: %s (%s)", formatCount(metrics.StoriesCompleted), formatPercent(metrics.CompletionRate)),
	}
	for _, id := range sprint.stories {
		story, _ := p.Story(id)
		points, situation := story.storyPoints, "en curso"
		if sprint.state == SprintClosed {
			points, situation = sprint.pointsAtClose[id], "no completada"
			if slices.Contains(sprint.completedStories, id) {
				situation = "completada"
			}
		}
		lines = append(lines, fmt.Sprintf("#%d — %s — %s — %s", id, story.title, pointsText(points), situation))
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

// effortLines escribe la seccion Estimacion y esfuerzo con los numeros de las metricas, sin
// calcular ninguno (RN-2, RN-11), y con la regla de formato de format.go (RN-14).
func effortLines(m Metrics) []string {
	return []string{
		"Story points planificados: " + formatCount(m.PlannedPoints),
		"Story points completados: " + formatCount(m.CompletedPoints),
		"Velocidad: " + formatCount(m.Velocity),
		"Horas estimadas: " + formatHours(m.EstimatedHours),
		"Horas reales: " + formatHours(m.ActualHours),
		"Desviación: " + formatPercent(m.HoursDeviation),
	}
}
