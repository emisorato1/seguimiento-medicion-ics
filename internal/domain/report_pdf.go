package domain

import (
	"fmt"
	"time"
)

// La exportacion del reporte a PDF (historia #19) tiene tres pasos, cada uno con sus tests:
// maquetar el reporte en paginas de renglones, escribir esas paginas como un PDF y entregarlo
// a quien lo pide. Este archivo empieza por el primero, que no sabe nada de PDF.

// pdfStyle es la letra de un renglon: el cuerpo, un encabezado de seccion o el titulo.
type pdfStyle int

const (
	styleBody pdfStyle = iota
	styleHeading
	styleTitle
)

// pdfSection y pdfDoc son lo que el PDF lee del reporte: el titulo, el dia en que se genero y
// las secciones con su encabezado y sus lineas. Se usan en lugar de Report mientras el codigo
// de la #18 no este en dev (decidido el 26/10, ver la Fase 4 de la spec).
type pdfSection struct {
	heading string
	lines   []string
}

type pdfDoc struct {
	title       string
	generatedOn time.Time
	sections    []pdfSection
}

// pdfRow es un renglon de una pagina. Un texto vacio es un renglon en blanco.
type pdfRow struct {
	text  string
	style pdfStyle
}

// pdfPage es una pagina maquetada: sus renglones y el pie con su numero.
type pdfPage struct {
	rows   []pdfRow
	footer string
}

// layoutReport convierte el reporte en paginas: el titulo, "Generado el" con la fecha, y cada
// encabezado con sus lineas, en el orden del reporte (RN-2, RN-3).
func layoutReport(doc pdfDoc) []pdfPage {
	rows := []pdfRow{
		{text: doc.title, style: styleTitle},
		{text: "Generado el " + doc.generatedOn.Format("02/01/2006"), style: styleBody},
	}
	for _, section := range doc.sections {
		rows = append(rows, pdfRow{text: section.heading, style: styleHeading})
		for _, line := range section.lines {
			rows = append(rows, pdfRow{text: line, style: styleBody})
		}
	}
	pages := []pdfPage{{rows: rows}}
	for i := range pages {
		pages[i].footer = fmt.Sprintf("Página %d de %d", i+1, len(pages))
	}
	return pages
}
