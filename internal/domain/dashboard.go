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

// dashboardSource es lo que necesita el armado del dashboard: los sprints del proyecto. Las
// métricas llegan calculadas, nunca se calculan acá (D8).
type dashboardSource struct {
	sprints []*Sprint
}

// buildDashboard arma los datos del dashboard a partir de lo que trae source, sin cambiar
// nada. Un proyecto sin sprints no tiene estado general, sprint en curso ni gráficos, y eso
// no es un error (RN-7).
func buildDashboard(source dashboardSource) Dashboard {
	return Dashboard{HasSprints: len(source.sprints) > 0}
}
