package domain

import "strconv"

// La regla de formato de los números (RN-15 de la #17, RN-14 de la #18): cómo se escribe cada
// número para mostrarlo, igual en el dashboard y en el reporte. La escribe la primera de las
// dos historias que llega a `dev`, y la otra la reutiliza (contrato del dominio).
//
// Por ahora escribe cada número tal como llega; los decimales, la coma y la unidad son de la
// tarea 3 de la #17.

// formatCount escribe un entero: story points, historias, defectos y velocidad.
func formatCount(n int) string {
	return strconv.Itoa(n)
}

// formatHours escribe una cantidad de horas.
func formatHours(hours float64) string {
	return strconv.FormatFloat(hours, 'f', -1, 64)
}

// formatPercent escribe un porcentaje: el de historias completas y la desviación de horas.
func formatPercent(percent float64) string {
	return strconv.FormatFloat(percent, 'f', -1, 64)
}
