package web

import (
	"math"
	"net/http"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// La página del dashboard (historia #17). El dominio entrega los números ya escritos; la web
// solo decide cuánto mide cada barra, que depende de la pantalla (D2).

// dashboardPage es lo que dibuja la plantilla: el proyecto, para el encabezado, y lo que
// arma el dominio.
type dashboardPage struct {
	*domain.Project
	Dashboard domain.Dashboard
}

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
	s.render(w, r, http.StatusOK, "dashboard.html", page{
		Title:   "Dashboard de " + project.Name(),
		Project: project,
		Section: "dashboard",
		Data:    dashboardPage{Project: project, Dashboard: project.Dashboard()},
	})
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
