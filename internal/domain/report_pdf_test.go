package domain

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
)

// unReporteChico arma a mano el reporte del escenario, como un Report de la #18: el sprint 1,
// generado el 25/10/2026, con dos secciones de dos lineas. Los escenarios se arman con el
// Report y ya no con un tipo propio del PDF, para que la trazabilidad pase por el reporte real
// (pedido de Maximiliano Eula en el #184, Fase 4 de la spec).
func unReporteChico() Report {
	return Report{
		Title:       "Reporte del sprint 1 — Seguimiento ICS",
		GeneratedOn: unDia(2026, 10, 25),
		Scope:       ScopeSprint,
		Sections: []Section{
			{Heading: "Resumen", Lines: []string{"Proyecto: Seguimiento ICS", "Sprint Goal: Tener el MVP"}},
			{Heading: "Defectos", Lines: []string{"Defectos detectados: 0", "No hay defectos en este sprint."}},
		},
	}
}

// exportar escribe el reporte como PDF con r.WritePDF en un destino en memoria y devuelve el
// archivo. Si la exportacion falla, corta el test.
func exportar(t *testing.T, report Report) []byte {
	t.Helper()
	var buf bytes.Buffer
	if err := report.WritePDF(&buf); err != nil {
		t.Fatalf("no se pudo exportar el reporte: %v", err)
	}
	return buf.Bytes()
}

// Escenario: Exportar el reporte de un sprint a PDF
// Cubre: RN-2, RN-3, RN-8
func TestExportarElReporteDeUnSprintAPDF(t *testing.T) {
	pages := layoutReport(unReporteChico())

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
	report := Report{Title: "Reporte", GeneratedOn: unDia(2026, 10, 25), Sections: []Section{
		{Heading: "Resumen", Lines: []string{"a", "", "b"}},
	}}

	rows := layoutReport(report)[0].rows

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
	pages := layoutReport(Report{Title: "Reporte", GeneratedOn: unDia(2026, 10, 25)})

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
		// Un caracter que no esta en la tabla se mide con el ancho mas grande.
		{text: "😀", style: styleBody, want: 10.15},
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
	report := Report{Title: strings.TrimSpace(strings.Repeat("REPORTE DEL PROYECTO ", 4)), GeneratedOn: unDia(2026, 10, 25)}

	rows := layoutReport(report)[0].rows

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

// unReporteConLineas arma un reporte con una seccion por cada cantidad de lineas pedida. Las
// lineas son cortas, asi que cada una ocupa un renglon.
func unReporteConLineas(counts ...int) Report {
	report := Report{Title: "Reporte", GeneratedOn: unDia(2026, 10, 25)}
	for s, count := range counts {
		section := Section{Heading: fmt.Sprintf("Seccion %d", s+1)}
		for i := 0; i < count; i++ {
			section.Lines = append(section.Lines, fmt.Sprintf("Linea %d", i+1))
		}
		report.Sections = append(report.Sections, section)
	}
	return report
}

// Escenario: Un reporte largo ocupa varias paginas numeradas
// Cubre: RN-7, RN-8
func TestUnReporteLargoOcupaVariasPaginasNumeradas(t *testing.T) {
	cases := []struct {
		paginas int
		lineas  int
		pie     string
	}{
		{paginas: 1, lineas: 10, pie: "Página 1 de 1"},
		// Titulo, "Generado el", encabezado y 100 lineas: 103 renglones, en paginas de 39.
		{paginas: 3, lineas: 100, pie: "Página 3 de 3"},
	}
	for _, c := range cases {
		t.Run(c.pie, func(t *testing.T) {
			pages := layoutReport(unReporteConLineas(c.lineas))

			if len(pages) != c.paginas {
				t.Fatalf("se esperaban %d paginas, salieron %d", c.paginas, len(pages))
			}
			total := 0
			for i, page := range pages {
				if len(page.rows) > 39 {
					t.Errorf("la pagina %d tiene %d renglones, mas de %d", i+1, len(page.rows), 39)
				}
				if want := fmt.Sprintf("Página %d de %d", i+1, c.paginas); page.footer != want {
					t.Errorf("se esperaba el pie %q, salio %q", want, page.footer)
				}
				total += len(page.rows)
			}
			if total != 3+c.lineas {
				t.Errorf("se esperaban %d renglones en total, salieron %d", 3+c.lineas, total)
			}
			if pages[len(pages)-1].footer != c.pie {
				t.Errorf("se esperaba que la ultima dijera %q, dice %q", c.pie, pages[len(pages)-1].footer)
			}
		})
	}
}

// Escenario: Un encabezado nunca queda solo al pie de una pagina
// Cubre: RN-7
func TestUnEncabezadoNuncaQuedaSoloAlPieDeUnaPagina(t *testing.T) {
	// Titulo, "Generado el", encabezado y 35 lineas: 38 renglones. El encabezado de la seccion
	// 2 caeria en el renglon 39, el ultimo de la pagina.
	pages := layoutReport(unReporteConLineas(35, 3))

	if len(pages) != 2 {
		t.Fatalf("se esperaban 2 paginas, salieron %d", len(pages))
	}
	first := pages[0].rows
	if last := first[len(first)-1]; last.text != "Linea 35" {
		t.Errorf("se esperaba que la pagina 1 terminara en \"Linea 35\", termina en %q", last.text)
	}
	second := pages[1].rows
	if second[0] != (pdfRow{text: "Seccion 2", style: styleHeading}) || second[1].text != "Linea 1" {
		t.Errorf("se esperaba que la pagina 2 empezara con el encabezado y su primera linea, empieza con %v", second[:2])
	}
}

// Sin escenario: caso del foco de la revision del plan
// Cubre: RN-7
//
// Un encabezado sin lineas tampoco queda como ultimo renglon, y pasarlo a la pagina siguiente
// no deja ninguna pagina vacia.
func TestUnEncabezadoSinLineasAlFinalNoAgregaPaginaVacia(t *testing.T) {
	pages := layoutReport(unReporteConLineas(35, 0))

	if len(pages) != 2 {
		t.Fatalf("se esperaban 2 paginas, salieron %d", len(pages))
	}
	if len(pages[1].rows) != 1 || pages[1].rows[0].text != "Seccion 2" {
		t.Errorf("se esperaba la pagina 2 con solo el encabezado de la seccion 2, salio %v", pages[1].rows)
	}
	for i, page := range pages {
		if len(page.rows) == 0 {
			t.Errorf("la pagina %d quedo vacia", i+1)
		}
	}
}

// parseXref lee la tabla de posiciones del PDF: devuelve donde empieza la tabla y la posicion
// de cada objeto en uso, en orden (el objeto 1 primero). Corta el test si la tabla no se puede
// leer.
func parseXref(t *testing.T, pdf []byte) (int, []int) {
	t.Helper()
	tail := string(pdf[bytes.LastIndex(pdf, []byte("startxref")):])
	fields := strings.Fields(tail)
	if len(fields) < 3 || fields[2] != "%%EOF" {
		t.Fatalf("el final del archivo no es startxref, su posicion y %%%%EOF: %q", tail)
	}
	start, err := strconv.Atoi(fields[1])
	if err != nil || start >= len(pdf) || !bytes.HasPrefix(pdf[start:], []byte("xref\n")) {
		t.Fatalf("startxref no apunta al comienzo de la tabla xref (%q)", fields[1])
	}
	lines := strings.Split(string(pdf[start:]), "\n")
	var first, count int
	fmt.Sscanf(lines[1], "%d %d", &first, &count)
	var offsets []int
	for _, entry := range lines[2 : 2+count] {
		if len(entry) != 19 {
			t.Fatalf("cada entrada de la tabla mide 20 bytes con su fin de linea, %q no", entry)
		}
		if strings.HasSuffix(entry, "n ") {
			offset, _ := strconv.Atoi(entry[:10])
			offsets = append(offsets, offset)
		}
	}
	return start, offsets
}

// Escenario: El archivo se abre en un lector de PDF
// Cubre: RN-1, RN-4
//
// Lo que un lector controla antes de mostrar nada (Fase 3 de la spec). Que de verdad abra se
// verifica a mano, y se anota en la Fase 4.
func TestElArchivoSeAbreEnUnLectorDePDF(t *testing.T) {
	pdf := exportar(t, unReporteConLineas(100))

	if !bytes.HasPrefix(pdf, []byte("%PDF-1.4\n")) {
		t.Errorf("el archivo no empieza con la cabecera de PDF 1.4: %q", pdf[:min(len(pdf), 12)])
	}
	if !bytes.HasSuffix(pdf, []byte("%%EOF\n")) {
		t.Errorf("el archivo no termina con la marca de fin")
	}
	start, offsets := parseXref(t, pdf)
	for i, offset := range offsets {
		if want := fmt.Sprintf("%d 0 obj\n", i+1); !bytes.HasPrefix(pdf[offset:], []byte(want)) {
			t.Errorf("la tabla dice que el objeto %d empieza en %d, y ahi no empieza", i+1, offset)
		}
	}
	trailer := string(pdf[start:])
	if !strings.Contains(trailer, "/Root 1 0 R") || !strings.Contains(trailer, fmt.Sprintf("/Size %d", len(offsets)+1)) {
		t.Errorf("el trailer no nombra el catalogo o no cuenta los objetos: %q", trailer)
	}
	text := string(pdf)
	if !strings.Contains(text, "/Type /Pages") || !strings.Contains(text, "/Count 3") {
		t.Errorf("el arbol de paginas no cuenta las 3 paginas")
	}
	if got := strings.Count(text, "/MediaBox [0 0 595.28 841.89]"); got != 3 {
		t.Errorf("se esperaban 3 paginas A4, salieron %d", got)
	}
	for _, font := range []string{"/BaseFont /Helvetica ", "/BaseFont /Helvetica-Bold "} {
		if !strings.Contains(text, font+"/Encoding /WinAnsiEncoding") {
			t.Errorf("falta la letra %q con la codificacion WinAnsi", font)
		}
	}
}

// textoPDF es un texto escrito en el contenido de una pagina, con su letra y su posicion.
type textoPDF struct {
	font, size, text string
	x, y             float64
}

var textoEnPDF = regexp.MustCompile(`BT /(F\d) (\d+) Tf ([\d.]+) ([\d.]+) Td \(((?:[^\\()]|\\.)*)\) Tj ET`)

// winAnsiDeTest pasa a texto los bytes de 0x80 a 0x9F que se usan en los tests; el resto de la
// codificacion WinAnsi coincide con Latin-1.
var winAnsiDeTest = map[byte]rune{0x80: '€', 0x85: '…', 0x91: '‘', 0x92: '’', 0x93: '“', 0x94: '”', 0x95: '•', 0x96: '–', 0x97: '—'}

// extraerTextos devuelve, en orden, los textos que el PDF escribe en sus paginas, como los
// leeria un lector: sin las barras del escape y con los bytes WinAnsi pasados a letras.
func extraerTextos(pdf []byte) []textoPDF {
	var textos []textoPDF
	for _, m := range textoEnPDF.FindAllSubmatch(pdf, -1) {
		var text []rune
		raw := m[5]
		for i := 0; i < len(raw); i++ {
			b := raw[i]
			if b == '\\' {
				i++
				b = raw[i]
			}
			if r, ok := winAnsiDeTest[b]; ok {
				text = append(text, r)
			} else {
				text = append(text, rune(b))
			}
		}
		x, _ := strconv.ParseFloat(string(m[3]), 64)
		y, _ := strconv.ParseFloat(string(m[4]), 64)
		textos = append(textos, textoPDF{font: string(m[1]), size: string(m[2]), text: string(text), x: x, y: y})
	}
	return textos
}

// Escenario: Exportar el reporte de un sprint a PDF
// Cubre: RN-2, RN-5, RN-9
//
// Amplia TestExportarElReporteDeUnSprintAPDF, que es del mismo escenario: aca se mira el
// archivo escrito y no el maquetado. La raya del titulo es de la tarea 9.
func TestExportarElReporteDeUnSprintAPDFTextoYLetras(t *testing.T) {
	report := unReporteChico()
	report.Title = "Reporte del sprint 1 - Seguimiento ICS"

	textos := extraerTextos(exportar(t, report))

	want := []textoPDF{
		{font: "F2", size: "16", text: "Reporte del sprint 1 - Seguimiento ICS"},
		{font: "F1", size: "10", text: "Generado el 25/10/2026"},
		{font: "F2", size: "12", text: "Resumen"},
		{font: "F1", size: "10", text: "Proyecto: Seguimiento ICS"},
		{font: "F1", size: "10", text: "Sprint Goal: Tener el MVP"},
		{font: "F2", size: "12", text: "Defectos"},
		{font: "F1", size: "10", text: "Defectos detectados: 0"},
		{font: "F1", size: "10", text: "No hay defectos en este sprint."},
		{font: "F1", size: "10", text: "Página 1 de 1"},
	}
	if len(textos) != len(want) {
		t.Fatalf("se esperaban %d textos en el archivo, salieron %d: %v", len(want), len(textos), textos)
	}
	for i, w := range want {
		got := textos[i]
		if got.font != w.font || got.size != w.size || got.text != w.text {
			t.Errorf("texto %d: se esperaba %s %s %q, salio %s %s %q", i+1, w.font, w.size, w.text, got.font, got.size, got.text)
		}
		if got.x != pdfMargin || got.y < pdfMargin || got.y > pdfPageHeight-pdfMargin {
			t.Errorf("el texto %q esta fuera de los margenes: (%v, %v)", got.text, got.x, got.y)
		}
	}
}

// Escenario: Los parentesis y las barras no rompen el archivo
// Cubre: RN-11
func TestLosParentesisYLasBarrasNoRompenElArchivo(t *testing.T) {
	line := `Desviacion (sobre lo estimado): 25,0 % \ revisar`
	report := Report{Title: "Reporte", GeneratedOn: unDia(2026, 10, 25), Sections: []Section{{Heading: "Resumen", Lines: []string{line}}}}

	pdf := exportar(t, report)

	if !bytes.Contains(pdf, []byte(`(Desviacion \(sobre lo estimado\): 25,0 % \\ revisar) Tj`)) {
		t.Errorf("la linea no se escribio con los parentesis y la barra escapados")
	}
	start, offsets := parseXref(t, pdf)
	if start == 0 || len(offsets) != 6 {
		t.Errorf("la tabla de posiciones no quedo bien: empieza en %d con %d objetos", start, len(offsets))
	}
	textos := extraerTextos(pdf)
	if len(textos) < 4 || textos[3].text != line {
		t.Errorf("se esperaba leer la linea tal cual, salio %v", textos)
	}
}

// Escenario: Letras del castellano y caracteres que no se pueden representar
// Cubre: RN-10
func TestLetrasDelCastellanoYCaracteresQueNoSePuedenRepresentar(t *testing.T) {
	cases := []struct {
		contenido string
		text      string
		want      []byte
	}{
		{contenido: "tildes, enie, dieresis y signos de apertura", text: "áéíóú ñ ü ¿¡", want: []byte{0xE1, 0xE9, 0xED, 0xF3, 0xFA, ' ', 0xF1, ' ', 0xFC, ' ', 0xBF, 0xA1}},
		// La raya no esta en Latin-1: en WinAnsi es el byte 0x97.
		{contenido: "la raya larga del titulo del reporte de un sprint", text: "Reporte del sprint 1 — X", want: []byte("Reporte del sprint 1 \x97 X")},
		{contenido: "un emoji", text: "ok 😀 ok", want: []byte("ok ? ok")},
		{contenido: "un caracter de control de Latin-1", text: "a\u0085b", want: []byte("a?b")},
	}
	for _, c := range cases {
		t.Run(c.contenido, func(t *testing.T) {
			if got := encodeWinAnsi(c.text); !bytes.Equal(got, c.want) {
				t.Errorf("se esperaba % X, salio % X", c.want, got)
			}
		})
	}

	report := unReporteChico()
	textos := extraerTextos(exportar(t, report))
	if textos[0].text != report.Title || textos[len(textos)-1].text != "Página 1 de 1" {
		t.Errorf("se esperaba leer el titulo con su raya y el pie con su tilde, salio %q y %q", textos[0].text, textos[len(textos)-1].text)
	}
}

// writerQueFalla es un destino que devuelve un error al escribir, con un mensaje en ingles
// como los de los archivos y las conexiones.
type writerQueFalla struct{}

func (writerQueFalla) Write([]byte) (int, error) { return 0, errors.New("broken pipe") }

// writerCorto es un destino que escribe la mitad de lo que recibe y no avisa.
type writerCorto struct{}

func (writerCorto) Write(p []byte) (int, error) { return len(p) / 2, nil }

// Escenario: Un PDF que no se puede entregar se informa
// Cubre: RN-12
func TestUnPDFQueNoSePuedeEntregarSeInforma(t *testing.T) {
	cases := []struct {
		name string
		w    io.Writer
	}{
		{name: "el destino devuelve un error", w: writerQueFalla{}},
		// Caso del foco de la revision: escribir de menos sin error tambien es una falla.
		{name: "el destino escribe la mitad", w: writerCorto{}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := unReporteChico().WritePDF(c.w)

			if !errors.Is(err, ErrPDFNotGenerated) {
				t.Fatalf("se esperaba que se informara que no se pudo generar el PDF, salio: %v", err)
			}
			if strings.Contains(err.Error(), "broken pipe") {
				t.Errorf("el mensaje es para el usuario y no lleva el error del destino: %q", err)
			}
		})
	}
}

// unReporteDeTreintaLineas arma un reporte con las cuatro secciones de la #18 y 30 lineas en
// total, inventadas: el PDF no mira que dicen. Lleva metricas que no estan en ninguna linea,
// para que se note si el PDF las escribe por su cuenta (RN-3).
func unReporteDeTreintaLineas(title string, scope ReportScope) Report {
	report := Report{
		Title:       title,
		GeneratedOn: unDia(2026, 10, 25),
		Scope:       scope,
		Metrics:     Metrics{PlannedPoints: 13, CompletedPoints: 8, Velocity: 8, EstimatedHours: 40, ActualHours: 52},
	}
	for s, heading := range []string{"Resumen", "Historias", "Estimación y esfuerzo", "Defectos"} {
		section := Section{Heading: heading}
		for i := 0; i < []int{6, 12, 6, 6}[s]; i++ {
			section.Lines = append(section.Lines, fmt.Sprintf("%s, linea %d", heading, i+1))
		}
		report.Sections = append(report.Sections, section)
	}
	return report
}

// textosDelReporte son los textos que tiene que tener el PDF de un reporte de una pagina, en
// orden: el titulo, "Generado el" con la fecha, cada encabezado con sus lineas y el pie. Es
// lo unico que el PDF agrega al reporte (RN-3).
func textosDelReporte(report Report) []string {
	want := []string{report.Title, "Generado el " + report.GeneratedOn.Format("02/01/2006")}
	for _, section := range report.Sections {
		want = append(want, section.Heading)
		want = append(want, section.Lines...)
	}
	return append(want, "Página 1 de 1")
}

// Esquema del escenario: El PDF no agrega ni quita nada del reporte
// Cubre: RN-3, RN-13
func TestElPDFNoAgregaNiQuitaNadaDelReporte(t *testing.T) {
	cases := []struct {
		reporte string
		title   string
		scope   ReportScope
	}{
		{reporte: "el reporte del proyecto", title: "Reporte del proyecto Seguimiento ICS", scope: ScopeProject},
		{reporte: "el reporte de un sprint", title: "Reporte del sprint 1 — Seguimiento ICS", scope: ScopeSprint},
	}
	for _, c := range cases {
		t.Run(c.reporte, func(t *testing.T) {
			report := unReporteDeTreintaLineas(c.title, c.scope)

			pdf := exportar(t, report)

			want := textosDelReporte(report)
			var got []string
			for _, texto := range extraerTextos(pdf) {
				got = append(got, texto.text)
			}
			if len(want) != 2+4+30+1 || !reflect.DeepEqual(got, want) {
				t.Errorf("se esperaba exactamente el reporte, \"Generado el\" y el pie:\n %q\nsalio:\n %q", want, got)
			}
		})
	}
}

// Escenario: Exportar dos veces el mismo reporte da el mismo archivo
// Cubre: RN-3
func TestExportarDosVecesElMismoReporteDaElMismoArchivo(t *testing.T) {
	first := exportar(t, unReporteChico())
	second := exportar(t, unReporteChico())

	if !bytes.Equal(first, second) {
		t.Errorf("los dos archivos no son iguales byte por byte")
	}
	for _, from := range []string{"/CreationDate", "/ModDate", "/ID"} {
		if bytes.Contains(first, []byte(from)) {
			t.Errorf("el archivo tiene %s, que cambia en cada exportacion", from)
		}
	}
}

// Sin escenario: los reportes que genera la #18, exportados de punta a punta
// Cubre: RN-2, RN-3, RN-13
//
// Los escenarios arman el reporte a mano. Este exporta los dos que genera el dominio para un
// proyecto con un sprint cerrado, el del proyecto y el del sprint, y lee en el archivo
// exactamente su titulo, "Generado el", sus encabezados, sus lineas y el pie.
func TestLosReportesQueGeneraLaHistoria18SeExportanTalCual(t *testing.T) {
	project := unProyectoLlamado(t, "Seguimiento ICS")
	unSprintCerrado(t, project, "Tener el MVP", unDia(2026, 10, 20), unDia(2026, 10, 25), unDia(2026, 10, 25), []int{1}, 3, 5)
	today := unDia(2026, 10, 27)
	sprintReport, err := project.SprintReport(1, today)
	if err != nil {
		t.Fatalf("no se pudo generar el reporte del sprint 1: %v", err)
	}
	cases := []struct {
		reporte string
		report  Report
	}{
		{reporte: "el reporte del proyecto", report: project.ProjectReport(today)},
		{reporte: "el reporte del sprint 1", report: sprintReport},
	}
	for _, c := range cases {
		t.Run(c.reporte, func(t *testing.T) {
			pdf := exportar(t, c.report)

			want := textosDelReporte(c.report)
			var got []string
			for _, texto := range extraerTextos(pdf) {
				got = append(got, texto.text)
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("se esperaba en el archivo exactamente el reporte de la #18:\n %q\nsalio:\n %q", want, got)
			}
		})
	}
}
