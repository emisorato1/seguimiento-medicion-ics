package domain

import (
	"math"
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
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

// Sin escenario: la tabla de anchos de la Fase 2
// Cubre: RN-4
//
// Los anchos son los de los AFM de Adobe, en milesimas del tamano de la letra. Se prueban
// algunos valores conocidos, y que todos los caracteres ASCII imprimibles esten en las dos
// tablas, para que ninguno se mida con el ancho de un caracter desconocido.
func TestAnchosDeHelvetica(t *testing.T) {
	cases := []struct {
		text  string
		style pdfStyle
		want  float64
	}{
		{text: "i", style: styleBody, want: 2.22},
		{text: "W", style: styleBody, want: 9.44},
		{text: " ", style: styleBody, want: 2.78},
		{text: "W", style: styleTitle, want: 15.104},
		{text: "i", style: styleHeading, want: 3.336},
		{text: "ñ", style: styleBody, want: 5.56},
	}
	for _, c := range cases {
		if got := textWidth(c.text, c.style); math.Abs(got-c.want) > 1e-9 {
			t.Errorf("ancho de %q con la letra %d: se esperaba %v, salio %v", c.text, c.style, c.want, got)
		}
	}
	for r := rune(32); r <= 126; r++ {
		if _, ok := helveticaWidths[r]; !ok {
			t.Errorf("falta %q en la tabla de Helvetica", r)
		}
		if _, ok := helveticaBoldWidths[r]; !ok {
			t.Errorf("falta %q en la tabla de Helvetica-Bold", r)
		}
	}
}

// Escenario: Una linea larga sigue en el renglon siguiente
// Cubre: RN-4, RN-6
func TestUnaLineaLargaSigueEnElRenglonSiguiente(t *testing.T) {
	line := strings.TrimSpace(strings.Repeat("palabra ", 120))

	rows := wrapText(line, styleBody)

	if len(rows) < 2 {
		t.Fatalf("se esperaba que la linea siguiera en otro renglon, salio %d renglon", len(rows))
	}
	for _, row := range rows {
		if textWidth(row, styleBody) > pdfLineWidth {
			t.Errorf("el renglon %q pasa el margen: mide %v pt", row, textWidth(row, styleBody))
		}
		if strings.HasPrefix(row, " ") || strings.HasSuffix(row, " ") {
			t.Errorf("el renglon %q empieza o termina con un espacio: no se corto entre palabras", row)
		}
	}
	if strings.Join(rows, " ") != line {
		t.Errorf("al unir los renglones no vuelve la linea original")
	}

	// Casos del foco de la revision: los espacios de una linea corta no se tocan.
	for _, short := range []string{"a  b", "   "} {
		if got := wrapText(short, styleBody); !reflect.DeepEqual(got, []string{short}) {
			t.Errorf("se esperaba %q en un solo renglon, salio %q", short, got)
		}
	}
}

// Escenario: Un titulo largo sigue en varios renglones con su letra
// Cubre: RN-4, RN-5, RN-6
func TestUnTituloLargoSigueEnVariosRenglonesConSuLetra(t *testing.T) {
	doc := pdfDoc{title: strings.TrimSpace(strings.Repeat("REPORTE DEL PROYECTO ", 4)), generatedOn: unDia(2026, 10, 25)}

	rows := layoutReport(doc)[0].rows

	titleRows := 0
	for _, row := range rows {
		if row.style == styleTitle {
			titleRows++
			if textWidth(row.text, styleTitle) > pdfLineWidth {
				t.Errorf("el renglon del titulo %q pasa el margen", row.text)
			}
		}
	}
	if titleRows < 2 {
		t.Errorf("se esperaba el titulo en varios renglones con su letra, salieron %d", titleRows)
	}
	if rows[titleRows].text != "Generado el 25/10/2026" {
		t.Errorf("se esperaba \"Generado el\" despues del titulo, salio %q", rows[titleRows].text)
	}
}

// Escenario: Una palabra mas larga que el renglon se corta en el margen
// Cubre: RN-6
func TestUnaPalabraMasLargaQueElRenglonSeCortaEnElMargen(t *testing.T) {
	cases := []struct {
		name string
		word string
	}{
		{name: "200 letras sin espacios", word: strings.Repeat("x", 200)},
		// Caso del foco de la revision: el corte va entre letras y no parte una ñ en dos.
		{name: "200 eñes sin espacios", word: strings.Repeat("ñ", 200)},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			rows := wrapText(c.word, styleBody)

			if len(rows) < 2 {
				t.Fatalf("se esperaba la palabra en varios renglones, salio %d", len(rows))
			}
			for _, row := range rows {
				if textWidth(row, styleBody) > pdfLineWidth {
					t.Errorf("un renglon pasa el margen: mide %v pt", textWidth(row, styleBody))
				}
				if !utf8.ValidString(row) {
					t.Errorf("el renglon %q tiene una letra partida", row)
				}
			}
			if strings.Join(rows, "") != c.word {
				t.Errorf("al unir los renglones no vuelve la palabra original")
			}
		})
	}
}
