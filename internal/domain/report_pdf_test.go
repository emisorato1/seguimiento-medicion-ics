package domain

import (
	"reflect"
	"testing"
)

// unDocChico arma el reporte del escenario: el sprint 1, generado el 25/10/2026, con dos
// secciones de dos lineas. Hasta que la #18 este en dev, el PDF trabaja sobre pdfDoc, que
// tiene los mismos tres datos que lee del reporte: titulo, fecha y secciones.
func unDocChico() pdfDoc {
	return pdfDoc{
		title:       "Reporte del sprint 1 — Seguimiento ICS",
		generatedOn: unDia(2026, 10, 25),
		sections: []pdfSection{
			{heading: "Resumen", lines: []string{"Proyecto: Seguimiento ICS", "Sprint Goal: Tener el MVP"}},
			{heading: "Defectos", lines: []string{"Defectos detectados: 0", "No hay defectos en este sprint."}},
		},
	}
}

// Escenario: Exportar el reporte de un sprint a PDF
// Cubre: RN-2, RN-3, RN-8
func TestExportarElReporteDeUnSprintAPDF(t *testing.T) {
	pages := layoutReport(unDocChico())

	want := []pdfRow{
		{text: "Reporte del sprint 1 — Seguimiento ICS", style: styleTitle},
		{text: "Generado el 25/10/2026", style: styleBody},
		{text: "Resumen", style: styleHeading},
		{text: "Proyecto: Seguimiento ICS", style: styleBody},
		{text: "Sprint Goal: Tener el MVP", style: styleBody},
		{text: "Defectos", style: styleHeading},
		{text: "Defectos detectados: 0", style: styleBody},
		{text: "No hay defectos en este sprint.", style: styleBody},
	}
	if len(pages) != 1 {
		t.Fatalf("se esperaba una sola pagina, salieron %d", len(pages))
	}
	if !reflect.DeepEqual(pages[0].rows, want) {
		t.Errorf("renglones de la pagina:\n se esperaba %v\n salio       %v", want, pages[0].rows)
	}
	if pages[0].footer != "Página 1 de 1" {
		t.Errorf("se esperaba el pie %q, salio %q", "Página 1 de 1", pages[0].footer)
	}
}

// Escenario: Una linea vacia queda como renglon en blanco
// Cubre: RN-2
func TestUnaLineaVaciaQuedaComoRenglonEnBlanco(t *testing.T) {
	doc := pdfDoc{title: "Reporte", generatedOn: unDia(2026, 10, 25), sections: []pdfSection{
		{heading: "Resumen", lines: []string{"a", "", "b"}},
	}}

	rows := layoutReport(doc)[0].rows

	want := []pdfRow{{text: "a", style: styleBody}, {text: "", style: styleBody}, {text: "b", style: styleBody}}
	if !reflect.DeepEqual(rows[3:], want) {
		t.Errorf("se esperaba %v despues del encabezado, salio %v", want, rows[3:])
	}
}

// Sin escenario: caso del foco de la revision del plan
// Cubre: RN-3
//
// Un reporte sin secciones no es un error: da una pagina con el titulo, "Generado el" y el pie.
func TestUnReporteSinSeccionesTieneUnaPagina(t *testing.T) {
	pages := layoutReport(pdfDoc{title: "Reporte", generatedOn: unDia(2026, 10, 25)})

	if len(pages) != 1 || len(pages[0].rows) != 2 || pages[0].footer != "Página 1 de 1" {
		t.Errorf("se esperaba una pagina con dos renglones y su pie, salio %v", pages)
	}
}
