package domain

import (
	"bytes"
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

// buildPDF escribe las paginas maquetadas como un archivo PDF 1.4: la cabecera, los objetos
// numerados, la tabla con la posicion de cada objeto y la marca de fin. Los objetos son el
// catalogo (1), el arbol de paginas (2), Helvetica (3), Helvetica-Bold (4) y, por cada
// pagina, la pagina y su contenido. Las letras son de las 14 estandar, sin incrustar, con la
// codificacion WinAnsi (Fase 2 de la spec).
func buildPDF(pages []pdfPage) []byte {
	var objects []string
	kids := make([]string, len(pages))
	for i := range pages {
		kids[i] = fmt.Sprintf("%d 0 R", 5+2*i)
	}
	objects = append(objects,
		"<< /Type /Catalog /Pages 2 0 R >>",
		fmt.Sprintf("<< /Type /Pages /Kids [%s] /Count %d >>", strings.Join(kids, " "), len(pages)),
		"<< /Type /Font /Subtype /Type1 /BaseFont /Helvetica /Encoding /WinAnsiEncoding >>",
		"<< /Type /Font /Subtype /Type1 /BaseFont /Helvetica-Bold /Encoding /WinAnsiEncoding >>",
	)
	for i := range pages {
		content := pageContent(pages[i])
		objects = append(objects,
			fmt.Sprintf("<< /Type /Page /Parent 2 0 R /MediaBox [0 0 %.2f %.2f] /Resources << /Font << /F1 3 0 R /F2 4 0 R >> >> /Contents %d 0 R >>",
				pdfPageWidth, pdfPageHeight, 6+2*i),
			fmt.Sprintf("<< /Length %d >>\nstream\n%s\nendstream", len(content), content),
		)
	}

	var out bytes.Buffer
	out.WriteString("%PDF-1.4\n")
	offsets := make([]int, len(objects))
	for i, object := range objects {
		offsets[i] = out.Len()
		fmt.Fprintf(&out, "%d 0 obj\n%s\nendobj\n", i+1, object)
	}
	xref := out.Len()
	fmt.Fprintf(&out, "xref\n0 %d\n0000000000 65535 f \n", len(objects)+1)
	for _, offset := range offsets {
		fmt.Fprintf(&out, "%010d 00000 n \n", offset)
	}
	fmt.Fprintf(&out, "trailer\n<< /Size %d /Root 1 0 R >>\nstartxref\n%d\n%%%%EOF\n", len(objects)+1, xref)
	return out.Bytes()
}

// pdfTop es la altura del margen de arriba: el renglon i se escribe en pdfTop - 18*(i+1), y
// el pie, en el lugar del renglon 40.
const (
	pdfTop       = pdfPageHeight - pdfMargin
	pdfRowHeight = 18
)

// pageContent escribe el contenido de una pagina: cada renglon con su letra y su tamano, y el
// pie. Un renglon en blanco no se escribe, pero ocupa su lugar.
func pageContent(page pdfPage) string {
	var out strings.Builder
	write := func(text string, style pdfStyle, row int) {
		font := "F2"
		if style == styleBody {
			font = "F1"
		}
		y := pdfTop - pdfRowHeight*float64(row+1)
		fmt.Fprintf(&out, "BT /%s %d Tf %.2f %.2f Td (%s) Tj ET\n", font, int(style.fontSize()), pdfMargin, y, escapePDFString(encodeWinAnsi(text)))
	}
	for i, row := range page.rows {
		if row.text != "" {
			write(row.text, row.style, i)
		}
	}
	write(page.footer, styleBody, pdfRowsPerPage)
	return out.String()
}

// winAnsiExtra son los caracteres de la codificacion WinAnsi que no estan en Latin-1, con su
// byte: van del 0x80 al 0x9F, donde Latin-1 tiene caracteres de control. La raya del titulo
// del reporte de un sprint es uno de ellos (tabla del apendice D de la referencia de PDF).
var winAnsiExtra = map[rune]byte{
	'€': 0x80, '‚': 0x82, 'ƒ': 0x83, '„': 0x84, '…': 0x85, '†': 0x86, '‡': 0x87, 'ˆ': 0x88,
	'‰': 0x89, 'Š': 0x8A, '‹': 0x8B, 'Œ': 0x8C, 'Ž': 0x8E, '‘': 0x91, '’': 0x92, '“': 0x93,
	'”': 0x94, '•': 0x95, '–': 0x96, '—': 0x97, '˜': 0x98, '™': 0x99, 'š': 0x9A, '›': 0x9B,
	'œ': 0x9C, 'ž': 0x9E, 'Ÿ': 0x9F,
}

// encodeWinAnsi pasa el texto a la codificacion de las letras, un byte por caracter. Los
// caracteres imprimibles de ASCII y de Latin-1 van con su mismo numero, los de winAnsiExtra con
// su byte, y cualquier otro, como un emoji o un caracter de control, sale como "?" (RN-10).
func encodeWinAnsi(text string) []byte {
	var out []byte
	for _, r := range text {
		if b, ok := winAnsiExtra[r]; ok {
			out = append(out, b)
		} else if (r >= 0x20 && r <= 0x7E) || (r >= 0xA0 && r <= 0xFF) {
			out = append(out, byte(r))
		} else {
			out = append(out, '?')
		}
	}
	return out
}

// escapePDFString pone una barra delante de los parentesis y de las barras del texto: en un
// PDF el texto va entre parentesis, y sin la barra un parentesis del reporte cerraria el
// texto antes de tiempo y el archivo no abriria (RN-11).
func escapePDFString(text []byte) []byte {
	var out []byte
	for _, b := range text {
		if b == '(' || b == ')' || b == '\\' {
			out = append(out, '\\')
		}
		out = append(out, b)
	}
	return out
}
