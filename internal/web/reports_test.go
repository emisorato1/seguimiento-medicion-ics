package web

import (
	"bytes"
	"html"
	"html/template"
	"io"
	"net/http"
	"regexp"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// el31DeOctubre es la fecha de hoy del escenario del reporte.
func el31DeOctubre() time.Time {
	return time.Date(2026, 10, 31, 0, 0, 0, 0, time.UTC)
}

// reportRegistry arma el proyecto 1 del escenario del reporte: el sprint 1 cerrado el
// 30/10/2026 con la historia 1 completada y la 2 no, y un defecto detectado en el.
func reportRegistry(t *testing.T) (*domain.ProjectRegistry, *domain.Project) {
	t.Helper()
	registry, project := closeRegistry(t)
	addDefect(t, project, 1)
	if err := project.CloseSprint([]int{1}, el30DeOctubre()); err != nil {
		t.Fatalf("no se pudo cerrar el sprint: %v", err)
	}
	return registry, project
}

// Expresiones para leer el reporte en la pagina: los encabezados de las secciones y cada
// linea de una seccion.
var (
	reportHeading = regexp.MustCompile(`<h2[^>]*>([^<]*)</h2>`)
	reportLine    = regexp.MustCompile(`<li>([^<]*)</li>`)
)

// Esquema del escenario: Ver el reporte
// Cubre: RN-1, RN-12, RN-14
//
// Las lineas que se esperan las arma el propio dominio con la misma fecha de hoy, sin
// escribirlas a mano: la pagina tiene que mostrarlas tal cual, sin cambiarles el formato
// (D6). Se leen sin el escape del HTML, que el navegador deshace.
func TestVerElReporte(t *testing.T) {
	casos := []struct {
		ambito string
		pagina string
		titulo string
		sprint int // 0 es el reporte del proyecto
	}{
		{ambito: "el proyecto", pagina: "/proyectos/1/reporte", titulo: "Reporte del proyecto Sistema de metricas"},
		{ambito: "el sprint 1", pagina: "/proyectos/1/reporte?sprint=1", titulo: "Reporte del sprint 1 — Sistema de metricas", sprint: 1},
	}
	for _, caso := range casos {
		t.Run(caso.ambito, func(t *testing.T) {
			registry, project := reportRegistry(t)
			saver := &fakeSaver{}
			server := NewServer(registry, saver, el31DeOctubre)

			res := get(t, server, caso.pagina)

			if res.Code != http.StatusOK {
				t.Fatalf("se esperaba el estado %d, salio %d:\n%s", http.StatusOK, res.Code, res.Body.String())
			}
			contenido := joinTags(between(res.Body.String(), "<main", "</main>"))
			if !strings.Contains(html.UnescapeString(contenido), "<h1>"+caso.titulo+"</h1>") {
				t.Errorf("se esperaba el titulo %q, y no esta en:\n%s", caso.titulo, contenido)
			}
			if !strings.Contains(contenido, "Generado el 31/10/2026") {
				t.Errorf("se esperaba \"Generado el 31/10/2026\", con la fecha de hoy, y no esta en:\n%s", contenido)
			}

			report := project.ProjectReport(el31DeOctubre())
			if caso.sprint != 0 {
				var err error
				if report, err = project.SprintReport(caso.sprint, el31DeOctubre()); err != nil {
					t.Fatalf("no se pudo armar el reporte esperado: %v", err)
				}
			}
			// Las cuatro secciones, en el orden del dominio.
			var encabezados []string
			for _, m := range reportHeading.FindAllStringSubmatch(contenido, -1) {
				encabezados = append(encabezados, html.UnescapeString(m[1]))
			}
			if want := []string{"Resumen", "Historias", "Estimación y esfuerzo", "Defectos"}; !slices.Equal(encabezados, want) {
				t.Errorf("se esperaban las secciones %q en ese orden, salieron %q", want, encabezados)
			}
			// Cada seccion con sus lineas, en orden y con el mismo texto que arma el dominio.
			for _, section := range report.Sections {
				seccion := between(contenido, ">"+section.Heading+"</h2>", "</section>")
				var lineas []string
				for _, m := range reportLine.FindAllStringSubmatch(seccion, -1) {
					lineas = append(lineas, html.UnescapeString(m[1]))
				}
				if !slices.Equal(lineas, section.Lines) {
					t.Errorf("en la seccion %s se esperaban las lineas del dominio:\n %q\nsalieron:\n %q", section.Heading, section.Lines, lineas)
				}
			}

			// El reporte se elige de una lista con los sprints del proyecto; sin sprint es el
			// del proyecto.
			if sprints := options(contenido, "sprint"); !slices.Equal(sprints, []string{"1"}) {
				t.Errorf("se esperaba elegir el sprint 1 de la lista, salio %q", sprints)
			}
			if elegido := strings.Contains(contenido, `<option value="1" selected>`); elegido != (caso.sprint == 1) {
				t.Errorf("se esperaba el sprint 1 elegido solo en su reporte, y la lista es:\n%s", between(contenido, `<select id="sprint"`, "</select>"))
			}
			// Ver el reporte no cambia ni guarda nada.
			if saver.saves != 0 {
				t.Errorf("ver el reporte no deberia guardar, y guardo %d veces", saver.saves)
			}
		})
	}
}

// Sin escenario: un sprint que no es un numero, escrito a mano en la direccion, se marca
// junto al campo y no se arma ningun reporte, como una fecha que no es fecha (Condiciones de
// error de la Fase 1). Tampoco se descarga ningun PDF (tarea 21).
// Cubre: RN-2
func TestUnSprintQueNoEsUnNumeroSeMarcaJuntoAlCampo(t *testing.T) {
	for _, pagina := range []string{"/proyectos/1/reporte?sprint=uno", "/proyectos/1/reporte/pdf?sprint=uno"} {
		t.Run(pagina, func(t *testing.T) {
			registry, _ := reportRegistry(t)
			server := newTestServer(t, registry)

			res := get(t, server, pagina)

			if res.Code != http.StatusUnprocessableEntity {
				t.Fatalf("se esperaba el estado %d, salio %d", http.StatusUnprocessableEntity, res.Code)
			}
			contenido := between(res.Body.String(), "<main", "</main>")
			if !strings.Contains(contenido, `id="sprint-error">`+errInvalidSprint.Error()) {
				t.Errorf("se esperaba el error junto al sprint, y no esta en:\n%s", contenido)
			}
			if strings.Contains(contenido, "Generado el") || strings.Contains(contenido, "<h2") {
				t.Errorf("con el sprint mal escrito no se deberia mostrar ningun reporte:\n%s", contenido)
			}
			if archivo := res.Header().Get("Content-Disposition"); archivo != "" {
				t.Errorf("con el sprint mal escrito no se deberia entregar ningun archivo, y se entrega %q", archivo)
			}
		})
	}
}

// Escenario: Descargar el reporte en PDF
// Cubre: RN-15
//
// El archivo tiene que ser el PDF de ese mismo reporte: se compara byte por byte con el que
// escribe la #19 para el reporte que arma el dominio con la misma fecha de hoy, asi tiene el
// mismo titulo y las mismas lineas que la pagina. La fila del proyecto no esta en el
// escenario: es el otro nombre de archivo de la RN-15.
func TestDescargarElReporteEnPDF(t *testing.T) {
	casos := []struct {
		reporte string
		pagina  string
		pdf     string
		archivo string
		sprint  int // 0 es el reporte del proyecto
	}{
		{reporte: "el del sprint 1", pagina: "/proyectos/1/reporte?sprint=1", pdf: "/proyectos/1/reporte/pdf?sprint=1", archivo: "reporte-proyecto-1-sprint-1.pdf", sprint: 1},
		{reporte: "el del proyecto", pagina: "/proyectos/1/reporte", pdf: "/proyectos/1/reporte/pdf", archivo: "reporte-proyecto-1.pdf"},
	}
	for _, caso := range casos {
		t.Run(caso.reporte, func(t *testing.T) {
			registry, project := reportRegistry(t)
			saver := &fakeSaver{}
			server := NewServer(registry, saver, el31DeOctubre)

			// La pagina del reporte ofrece descargar ese mismo reporte.
			contenido := between(get(t, server, caso.pagina).Body.String(), "<main", "</main>")
			if !strings.Contains(contenido, `href="`+caso.pdf+`"`) {
				t.Fatalf("se esperaba en la pagina el acceso para descargar el PDF (%s), y no esta en:\n%s", caso.pdf, contenido)
			}

			res := get(t, server, caso.pdf)

			if res.Code != http.StatusOK {
				t.Fatalf("se esperaba el estado %d, salio %d:\n%s", http.StatusOK, res.Code, res.Body.String())
			}
			if tipo := res.Header().Get("Content-Type"); tipo != "application/pdf" {
				t.Errorf("se esperaba un archivo PDF, y el tipo es %q", tipo)
			}
			if archivo, want := res.Header().Get("Content-Disposition"), `attachment; filename="`+caso.archivo+`"`; archivo != want {
				t.Errorf("se esperaba un archivo para guardar llamado %q (%s), salio %q", caso.archivo, want, archivo)
			}
			report := project.ProjectReport(el31DeOctubre())
			if caso.sprint != 0 {
				var err error
				if report, err = project.SprintReport(caso.sprint, el31DeOctubre()); err != nil {
					t.Fatalf("no se pudo armar el reporte esperado: %v", err)
				}
			}
			var want bytes.Buffer
			if err := report.WritePDF(&want); err != nil {
				t.Fatalf("no se pudo exportar el reporte esperado: %v", err)
			}
			if !bytes.Equal(res.Body.Bytes(), want.Bytes()) {
				t.Errorf("se esperaba el PDF del mismo reporte que muestra la pagina (%d bytes), y llegaron %d bytes distintos", want.Len(), res.Body.Len())
			}
			if saver.saves != 0 {
				t.Errorf("descargar el reporte no deberia guardar, y guardo %d veces", saver.saves)
			}
		})
	}
}

// Sin escenario: un PDF que no se puede generar se avisa con el mensaje del nucleo, en la
// pagina del mismo reporte, y no se entrega ningun archivo, ni siquiera lo que se llego a
// escribir. Escrito en memoria, el PDF de verdad no falla: el test reemplaza la exportacion
// por una que escribe un pedazo y falla como lo informa la #19.
// Cubre: RN-19
func TestUnPDFQueNoSePuedeGenerarSeAvisaSinEntregarArchivo(t *testing.T) {
	registry, _ := reportRegistry(t)
	server := newTestServer(t, registry)
	server.exportPDF = func(_ domain.Report, w io.Writer) error {
		io.WriteString(w, "%PDF-1.4 cortado")
		return domain.ErrPDFNotGenerated
	}

	res := get(t, server, "/proyectos/1/reporte/pdf?sprint=1")

	if res.Code != http.StatusInternalServerError {
		t.Fatalf("se esperaba el estado %d, salio %d", http.StatusInternalServerError, res.Code)
	}
	if archivo := res.Header().Get("Content-Disposition"); archivo != "" || strings.Contains(res.Body.String(), "%PDF") {
		t.Errorf("no se deberia entregar ningun archivo, y se entrega %q", archivo)
	}
	if tipo := res.Header().Get("Content-Type"); !strings.HasPrefix(tipo, "text/html") {
		t.Errorf("se esperaba una pagina con el aviso, y el tipo es %q", tipo)
	}
	contenido := between(res.Body.String(), "<main", "</main>")
	if mensaje := template.HTMLEscapeString(domain.ErrPDFNotGenerated.Error()); !strings.Contains(contenido, mensaje) {
		t.Errorf("se esperaba el aviso con el mensaje del nucleo %q, y no esta en:\n%s", mensaje, contenido)
	}
	if !strings.Contains(html.UnescapeString(contenido), "<h1>Reporte del sprint 1 — Sistema de metricas</h1>") || !strings.Contains(contenido, `href="/proyectos/1/reporte/pdf?sprint=1"`) {
		t.Errorf("se esperaba el aviso en la pagina del mismo reporte, con el acceso para volver a descargarlo:\n%s", contenido)
	}
}
