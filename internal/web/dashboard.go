package web

import (
	"math"
	"net/http"
	"strconv"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// La página del dashboard (historia #17). El dominio entrega los números ya escritos; la web
// solo decide cuánto mide cada barra, que depende de la pantalla (D2).

// dashboardPage es lo que dibuja la plantilla: el proyecto, para el encabezado, lo que arma
// el dominio y sus gráficos ya medidos.
type dashboardPage struct {
	*domain.Project
	Dashboard domain.Dashboard
	Charts    []chartView
}

// chartView es un gráfico listo para dibujar en SVG: su título, el aviso si no tiene datos
// (el de velocidad sin sprints cerrados, RN-9), el alto del dibujo y sus barras.
type chartView struct {
	Title  string
	Empty  string
	Height int
	Bars   []barView
}

// barView es una barra lista para dibujar: lo que entrega el dominio, su ancho, su color y
// dónde va dentro del SVG. Width va con punto decimal, como pide el SVG (D3), y queda vacío
// cuando el valor es cero y no lleva barra (RN-12).
type barView struct {
	domain.Bar
	Width string
	Class string
	TextY int
	RectY int
}

// Cada barra ocupa una fila del SVG: arriba el rótulo y el valor, debajo la barra (D5).
const (
	chartRow  = 44
	chartText = 16
	chartBar  = 22
)

// showDashboard muestra el dashboard de un proyecto, con su encabezado: el nombre, el inicio
// y el fin, o «sin fecha de fin» (RN-3). Sin sprints, debajo va solo el aviso de que no hay
// datos, que no es un error (RN-7). Solo lee, así que va con el candado de lectura (#77,
// D12). Un proyecto que no existe se informa como en el resto de la web (RN-17).
func (s *Server) showDashboard(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		s.notFound(w, r, nil, err)
		return
	}
	dashboard := project.Dashboard()
	s.render(w, r, http.StatusOK, "dashboard.html", page{
		Title:   "Dashboard de " + project.Name(),
		Project: project,
		Section: "dashboard",
		Data:    dashboardPage{Project: project, Dashboard: dashboard, Charts: chartViews(dashboard.Charts)},
	})
}

// chartViews mide las barras de cada gráfico contra su propio valor más grande (D9) y les da
// su color (D6). Los gráficos llegan en el orden de la RN-8: en los tres primeros, que
// comparan, la primera barra es lo planificado y la segunda lo real; el último es el de
// velocidad, con todas sus barras del mismo color.
func chartViews(charts []domain.Chart) []chartView {
	views := make([]chartView, 0, len(charts))
	for i, chart := range charts {
		max := 0.0
		for _, bar := range chart.Bars {
			max = math.Max(max, bar.Value)
		}
		view := chartView{Title: chart.Title, Empty: chart.Empty, Height: len(chart.Bars) * chartRow}
		for j, bar := range chart.Bars {
			class := "barra-planificada"
			switch {
			case i == len(charts)-1:
				class = "barra-velocidad"
			case j == 1:
				class = "barra-real"
			}
			width := ""
			if w := barWidth(bar.Value, max); w > 0 {
				width = strconv.FormatFloat(w, 'f', -1, 64) + "%"
			}
			view.Bars = append(view.Bars, barView{Bar: bar, Width: width, Class: class,
				TextY: j*chartRow + chartText, RectY: j*chartRow + chartBar})
		}
		views = append(views, view)
	}
	return views
}

// barWidth devuelve el largo de la barra de un valor, como porcentaje del ancho del gráfico
// con un decimal: el valor más grande del gráfico mide 100 y los demás su parte (RN-11). Un
// cero no tiene barra (RN-12), y un valor mayor que cero mide al menos 1, para que se vea
// (RN-13). Un gráfico todo en cero no tiene barras (RN-14).
func barWidth(value, max float64) float64 {
	if value <= 0 || max <= 0 {
		return 0
	}
	width := math.Round(value/max*1000) / 10
	return math.Max(width, 1)
}
