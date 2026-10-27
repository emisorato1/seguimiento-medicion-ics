package web

import "math"

// La página del dashboard (historia #17). El dominio entrega los números ya escritos; la web
// solo decide cuánto mide cada barra, que depende de la pantalla (D2).

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
