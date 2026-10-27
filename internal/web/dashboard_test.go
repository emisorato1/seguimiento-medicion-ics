package web

import (
	"fmt"
	"testing"
)

// Escenario: El largo de cada barra es proporcional al valor mas grande del grafico
// Cubre: RN-11, RN-12, RN-13
//
// El largo es un porcentaje del ancho del gráfico con un decimal; escribirlo con punto en el
// SVG es de la tarea 13 (D3).
func TestElLargoDeCadaBarraEsProporcionalAlValorMasGrandeDelGrafico(t *testing.T) {
	cases := []struct {
		max, value float64
		want       float64
	}{
		{40, 40, 100},
		{40, 10, 25},
		{8, 5, 62.5},
		{50, 12.5, 25},
		{1000, 1, 1},
		{40, 0, 0}, // sin barra
		// Dos casos que el escenario no muestra. Con un decimal: 8 de 13 es 61,538… y va
		// 61,5, como en el gráfico de velocidad. Y un gráfico todo en cero no divide por
		// cero: sin barra (RN-14, que la tarea 13 prueba con el gráfico entero).
		{13, 8, 61.5},
		{0, 0, 0},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v de %v", c.value, c.max), func(t *testing.T) {
			if got := barWidth(c.value, c.max); got != c.want {
				t.Fatalf("se esperaba que %v de %v midiera %v %%, salio %v %%", c.value, c.max, c.want, got)
			}
		})
	}
}
