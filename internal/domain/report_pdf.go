package domain

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
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

// Medidas de la pagina, en puntos (1/72 de pulgada): A4 vertical con margenes de 2 cm. Un
// renglon entra si no pasa el ancho del A4 menos los dos margenes (Fase 2 de la spec).
const (
	pdfPageWidth  = 595.28
	pdfPageHeight = 841.89
	pdfMargin     = 56.69
	pdfLineWidth  = 481.9
	// pdfRowsPerPage: todos los renglones miden 18 pt, y en el alto util entran 39 mas el pie.
	pdfRowsPerPage = 39
)

// fontSize es el tamano de cada letra: el cuerpo en 10 pt, los encabezados en 12 pt y el
// titulo en 16 pt, los dos ultimos en negrita (RN-5).
func (s pdfStyle) fontSize() float64 {
	switch s {
	case styleTitle:
		return 16
	case styleHeading:
		return 12
	}
	return 10
}

// Los anchos de Helvetica y Helvetica-Bold, en milesimas del tamano de la letra, copiados de
// los AFM estandar de Adobe (Helvetica.afm y Helvetica-Bold.afm). Se escriben agrupando las
// letras que miden lo mismo; las letras con tilde, la ñ y la ü miden como su letra sin marca,
// salvo la i con tilde, que mide como la t. Un caracter que no esta en la tabla se mide con
// unknownWidth, el mas ancho, para que ningun renglon pase el margen.
var helveticaWidths = widthTable(map[int]string{
	191:  "'",
	222:  "ijl",
	260:  "|",
	278:  " !,./:;I[\\]ftíìîïÍÌÎÏ·",
	333:  "()-`r¡",
	334:  "{}",
	355:  "\"",
	365:  "º",
	370:  "ª",
	389:  "*",
	400:  "°",
	469:  "^",
	500:  "Jcksvxyzç",
	556:  "#$0123456789?L_abdeghnopquáàâäãåéèêëñóòôöõúùûü–€«»",
	584:  "+<=>~",
	611:  "FTZ¿",
	667:  "&ABEKPSVXYÁÀÂÄÃÉÈÊË",
	722:  "CDHNRUwÑÚÙÛÜÇ",
	778:  "GOQÓÒÔÖÕ",
	833:  "Mm",
	889:  "%",
	944:  "W",
	1000: "—…",
	1015: "@",
})

var helveticaBoldWidths = widthTable(map[int]string{
	238:  "'",
	278:  " ,./I\\ijlíìîïÍÌÎÏ·",
	280:  "|",
	333:  "!()-:;[]`ft¡",
	365:  "º",
	370:  "ª",
	389:  "*{}r",
	400:  "°",
	474:  "\"",
	500:  "z",
	556:  "#$0123456789Jaceksvxy_áàâäãåéèêëç–€«»",
	584:  "+<=>^~",
	611:  "?FLTZbdghnopquñóòôöõúùûü¿",
	667:  "EPSVXYÉÈÊË",
	722:  "&ABCDHKNRUÁÀÂÄÃÑÚÙÛÜÇ",
	778:  "GOQwÓÒÔÖÕ",
	833:  "M",
	889:  "%m",
	944:  "W",
	975:  "@",
	1000: "—…",
})

const unknownWidth = 1015

// widthTable pasa la tabla agrupada por ancho a una tabla por caracter.
func widthTable(groups map[int]string) map[rune]int {
	table := map[rune]int{}
	for width, chars := range groups {
		for _, r := range chars {
			table[r] = width
		}
	}
	return table
}

// textWidth mide un texto en puntos, con la letra y el tamano de su estilo.
func textWidth(text string, style pdfStyle) float64 {
	table := helveticaWidths
	if style != styleBody {
		table = helveticaBoldWidths
	}
	total := 0
	for _, r := range text {
		width, ok := table[r]
		if !ok {
			width = unknownWidth
		}
		total += width
	}
	return float64(total) * style.fontSize() / 1000
}

// wrapText corta un texto en renglones que entran en el ancho de la pagina, entre palabras
// (RN-6). Un texto que entra queda como esta, con todos sus espacios. Una palabra que sola no
// entra se corta letra por letra donde llega al margen.
func wrapText(text string, style pdfStyle) []string {
	if textWidth(text, style) <= pdfLineWidth {
		return []string{text}
	}
	var rows []string
	current := ""
	for i, word := range strings.Split(text, " ") {
		if i == 0 {
			current = word
		} else if candidate := current + " " + word; textWidth(candidate, style) <= pdfLineWidth {
			current = candidate
		} else {
			rows = append(rows, current)
			current = word
		}
		for textWidth(current, style) > pdfLineWidth {
			head, rest := splitAtMargin(current, style)
			rows = append(rows, head)
			current = rest
		}
	}
	return append(rows, current)
}

// splitAtMargin separa de un texto que no entra la parte que si entra en el renglon, sin
// partir ninguna letra: se recorre por runas y no por bytes.
func splitAtMargin(text string, style pdfStyle) (string, string) {
	end := 0
	for i, r := range text {
		if textWidth(text[:i+utf8.RuneLen(r)], style) > pdfLineWidth {
			break
		}
		end = i + utf8.RuneLen(r)
	}
	return text[:end], text[end:]
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
	var rows []pdfRow
	add := func(text string, style pdfStyle) {
		for _, row := range wrapText(text, style) {
			rows = append(rows, pdfRow{text: row, style: style})
		}
	}
	add(doc.title, styleTitle)
	add("Generado el "+doc.generatedOn.Format("02/01/2006"), styleBody)
	for _, section := range doc.sections {
		add(section.heading, styleHeading)
		for _, line := range section.lines {
			add(line, styleBody)
		}
	}
	pages := []pdfPage{{}}
	for i, row := range rows {
		last := &pages[len(pages)-1]
		if len(last.rows) > 0 && len(last.rows)+1+headingRows(rows, i) > pdfRowsPerPage {
			pages = append(pages, pdfPage{})
			last = &pages[len(pages)-1]
		}
		last.rows = append(last.rows, row)
	}
	for i := range pages {
		pages[i].footer = fmt.Sprintf("Página %d de %d", i+1, len(pages))
	}
	return pages
}

// headingRows cuenta los renglones del encabezado que empieza en el renglon i (uno largo
// ocupa varios), o da 0 si ahi no empieza un encabezado. Un encabezado necesita lugar para
// el y para un renglon mas: si no lo tiene, pasa a la pagina siguiente y nunca queda como
// ultimo renglon (RN-7).
func headingRows(rows []pdfRow, i int) int {
	if rows[i].style != styleHeading || (i > 0 && rows[i-1].style == styleHeading) {
		return 0
	}
	count := 0
	for i+count < len(rows) && rows[i+count].style == styleHeading {
		count++
	}
	return count
}
