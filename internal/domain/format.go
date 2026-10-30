package domain

import (
	"math"
	"strconv"
	"strings"
	"time"
)

// La regla de formato de los números (RN-15 de la #17, RN-14 de la #18): cómo se escribe cada
// número para mostrarlo, igual en el dashboard y en el reporte. La escribe la primera de las
// dos historias que llega a `dev`, y la otra la reutiliza (contrato del dominio).
//
// Los enteros van sin decimales. Las horas y los porcentajes, con un decimal, coma decimal y
// su unidad separada por un espacio: "12,5 h", "66,7 %". Un negativo lleva su signo, y un
// valor que se redondea a cero se escribe sin signo.

// formatCount escribe un entero: story points, historias, defectos y velocidad.
func formatCount(n int) string {
	return strconv.Itoa(n)
}

// formatHours escribe una cantidad de horas, por ejemplo "12,5 h".
func formatHours(hours float64) string {
	return formatDecimal(hours) + " h"
}

// formatPercent escribe un porcentaje: el de historias completas y la desviación de horas,
// por ejemplo "66,7 %" o "-12,5 %".
func formatPercent(percent float64) string {
	return formatDecimal(percent) + " %"
}

// formatDecimal escribe un número con un decimal y coma decimal. Redondea el empate hacia
// afuera del cero, como se redondea a mano (12,25 da 12,3), y no al par como la biblioteca
// estándar. La coma se arma cambiando el punto, porque Go no tiene formato regional y el
// resultado no puede depender de la máquina.
func formatDecimal(x float64) string {
	rounded := math.Round(x*10) / 10
	if rounded == 0 {
		// Un negativo que se redondea a cero queda en -0, que se escribiría "-0,0".
		rounded = 0
	}
	return strings.Replace(strconv.FormatFloat(rounded, 'f', 1, 64), ".", ",", 1)
}

// formatDate escribe un día como día/mes/año, con el día y el mes en dos dígitos, por
// ejemplo "01/10/2026". Es el mismo formato que usa la web, así una fecha se lee igual en la
// pantalla, en el reporte y en el PDF (RN-14 de la #18). La suma la #18: en la #17 las fechas
// las escribe la web.
func formatDate(day time.Time) string {
	return day.Format("02/01/2006")
}
