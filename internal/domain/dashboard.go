package domain

import "time"

// Dashboard son los datos del dashboard de un proyecto (historia #17). El dominio los arma
// con los números ya escritos y la web los dibuja: cuánto mide cada barra depende de la
// pantalla y lo decide la web (D2). Es un dato plano y sus campos son públicos, como
// Metrics: no tiene reglas que proteger, y la web lo arma a mano en sus tests (D10).
type Dashboard struct {
	// HasSprints dice si el proyecto tiene algún sprint. Sin sprints, la página muestra solo
	// el encabezado y el aviso de que no hay datos (RN-7).
	HasSprints bool
	// Overview es el estado general: cada métrica del proyecto con su valor escrito (RN-4).
	Overview []Figure
	// CurrentSprint es el sprint abierto, o nil si no hay ninguno (RN-5, RN-6).
	CurrentSprint *CurrentSprint
	// Charts son los cuatro gráficos, en el orden de la RN-8.
	Charts []Chart
}

// CurrentSprint es el sprint en curso del dashboard: sus datos y lo que lleva registrado,
// sin completados ni velocidad, que se registran recién al cerrarlo (RN-5). Las fechas las
// escribe la web, como día/mes/año (RN-15).
type CurrentSprint struct {
	Number  int
	Goal    string
	Start   time.Time
	End     time.Time
	Figures []Figure
}

// Chart es un gráfico de barras del dashboard. Empty es el aviso que se muestra cuando no
// tiene barras, como «Todavía no hay sprints cerrados» (RN-9).
type Chart struct {
	Title string
	Bars  []Bar
	Empty string
}

// Bar es una barra de un gráfico: Value es el número con el que la web mide la barra, y Text
// el mismo valor ya escrito con la regla de formato (RN-15).
type Bar struct {
	Label string
	Value float64
	Text  string
}

// Figure es un número del estado general o del sprint en curso, ya escrito (RN-15).
type Figure struct {
	Label string
	Text  string
}

// dashboardSource es lo que necesita el armado del dashboard: los sprints del proyecto, las
// métricas del proyecto entero y las de cada sprint, por su número. Las métricas llegan
// calculadas, nunca se calculan acá (D8).
type dashboardSource struct {
	sprints        []*Sprint
	projectMetrics Metrics
	sprintMetrics  map[int]Metrics
}

// buildDashboard arma los datos del dashboard a partir de lo que trae source, sin cambiar
// nada. Un proyecto sin sprints no tiene estado general, sprint en curso ni gráficos, y eso
// no es un error (RN-7).
func buildDashboard(source dashboardSource) Dashboard {
	if len(source.sprints) == 0 {
		return Dashboard{}
	}
	return Dashboard{
		HasSprints:    true,
		Overview:      overview(source.projectMetrics),
		CurrentSprint: currentSprint(source),
		Charts:        charts(source.projectMetrics),
	}
}

// charts arma los gráficos del proyecto entero en el orden de la RN-8: story points, horas
// y defectos, con las mismas métricas del estado general. Cada barra lleva el valor tal como
// llegó, para que la web la mida, y ya escrito (RN-10, RN-15).
func charts(m Metrics) []Chart {
	return []Chart{
		{Title: "Story points", Bars: []Bar{
			{Label: "Planificados", Value: float64(m.PlannedPoints), Text: formatCount(m.PlannedPoints)},
			{Label: "Completados", Value: float64(m.CompletedPoints), Text: formatCount(m.CompletedPoints)},
		}},
		{Title: "Horas", Bars: []Bar{
			{Label: "Estimadas", Value: m.EstimatedHours, Text: formatHours(m.EstimatedHours)},
			{Label: "Reales", Value: m.ActualHours, Text: formatHours(m.ActualHours)},
		}},
		{Title: "Defectos", Bars: []Bar{
			{Label: "Detectados", Value: float64(m.DefectsDetected), Text: formatCount(m.DefectsDetected)},
			{Label: "Resueltos", Value: float64(m.DefectsResolved), Text: formatCount(m.DefectsResolved)},
		}},
	}
}

// currentSprint arma el sprint en curso con el sprint abierto y lo que lleva registrado. No
// trae los completados ni la velocidad, que se registran recién al cerrarlo (RN-5). Sin
// sprint abierto devuelve nil, y la web informa que no hay sprint en curso (RN-6).
func currentSprint(source dashboardSource) *CurrentSprint {
	for _, sprint := range source.sprints {
		if sprint.State() != SprintOpen {
			continue
		}
		m := source.sprintMetrics[sprint.Number()]
		return &CurrentSprint{
			Number: sprint.Number(),
			Goal:   sprint.Goal(),
			Start:  sprint.Start(),
			End:    sprint.End(),
			Figures: []Figure{
				{Label: "Story points planificados", Text: formatCount(m.PlannedPoints)},
				{Label: "Horas reales", Text: formatHours(m.ActualHours)},
				{Label: "Defectos detectados", Text: formatCount(m.DefectsDetected)},
			},
		}
	}
	return nil
}

// overview arma el estado general con las métricas del proyecto entero, en el orden de la
// RN-4, tal como llegaron: no vuelve a calcular ninguna (RN-2).
func overview(m Metrics) []Figure {
	return []Figure{
		{Label: "Story points planificados", Text: formatCount(m.PlannedPoints)},
		{Label: "Story points completados", Text: formatCount(m.CompletedPoints)},
		{Label: "Historias planificadas", Text: formatCount(m.StoriesPlanned)},
		{Label: "Historias completadas", Text: formatCount(m.StoriesCompleted)},
		{Label: "Porcentaje de historias completas", Text: formatPercent(m.CompletionRate)},
		{Label: "Horas estimadas", Text: formatHours(m.EstimatedHours)},
		{Label: "Horas reales", Text: formatHours(m.ActualHours)},
		{Label: "Desviación de horas", Text: formatPercent(m.HoursDeviation)},
		{Label: "Defectos detectados", Text: formatCount(m.DefectsDetected)},
		{Label: "Defectos resueltos", Text: formatCount(m.DefectsResolved)},
		{Label: "Velocidad", Text: formatCount(m.Velocity)},
	}
}
