package domain

import (
	"fmt"
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
			{Heading: "Resumen"},
			{Heading: "Historias"},
			{Heading: "Estimación y esfuerzo"},
			{Heading: "Defectos"},
		},
	}
}
