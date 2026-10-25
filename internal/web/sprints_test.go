package web

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// Escenario: La barra de un proyecto lleva a todas sus secciones
// Cubre: RN-3
//
// Cada pantalla de la #78 suma su fila cuando llega a dev, asi la barra no lleva nunca a una
// pagina que todavia no existe (Fase 4, Desvios).
func TestLaBarraDeUnProyectoLlevaATodasSusSecciones(t *testing.T) {
	secciones := []struct {
		nombre string
		pagina string
	}{
		{nombre: "Sprints", pagina: "/proyectos/1/sprints"},
	}

	registry := domain.NewProjectRegistry()
	addProject(t, registry, "Sistema de metricas")
	server := newTestServer(t, registry)

	for _, seccion := range secciones {
		t.Run(seccion.nombre, func(t *testing.T) {
			// Desde cualquier pagina del proyecto, la barra lleva a la seccion.
			for _, desde := range []string{"/proyectos/1", "/proyectos/1/backlog", seccion.pagina} {
				res := get(t, server, desde)
				if res.Code != http.StatusOK {
					t.Fatalf("se esperaba %s con el estado %d, salio %d", desde, http.StatusOK, res.Code)
				}
				barra := between(res.Body.String(), `<nav aria-label="Navegación principal">`, "</nav>")
				enlace := `<a href="` + seccion.pagina + `"`
				if !strings.Contains(barra, "Sistema de metricas") || !strings.Contains(barra, enlace) || !strings.Contains(barra, ">"+seccion.nombre+"</a>") {
					t.Errorf("en %s se esperaba en la barra el proyecto y el acceso a %s, y no estan en:\n%s", desde, seccion.nombre, barra)
				}
			}

			// En su propia pagina, la seccion figura como la actual.
			barra := between(get(t, server, seccion.pagina).Body.String(), `<nav aria-label="Navegación principal">`, "</nav>")
			if !strings.Contains(barra, `<a href="`+seccion.pagina+`" aria-current="page">`) {
				t.Errorf("se esperaba %s marcada como la seccion actual, y no esta en:\n%s", seccion.nombre, barra)
			}
		})
	}
}

// Escenario: Crear un sprint desde la web
// Cubre: RN-1, RN-2
func TestCrearUnSprintDesdeLaWeb(t *testing.T) {
	registry := domain.NewProjectRegistry()
	project := addProject(t, registry, "Sistema de metricas")
	saver := &fakeSaver{}
	navegador := newBrowser(t, NewServer(registry, saver))

	res := navegador.post("/proyectos/1/sprints/nuevo", url.Values{
		"goal":   {"Tener el MVP"},
		"inicio": {"2026-10-26"},
		"fin":    {"2026-10-30"},
	})

	if res.Code != http.StatusSeeOther {
		t.Fatalf("se esperaba el estado %d, salio %d:\n%s", http.StatusSeeOther, res.Code, res.Body.String())
	}
	lista := res.Header().Get("Location")
	if lista != "/proyectos/1/sprints" {
		t.Fatalf("se esperaba volver a la lista de sprints, salio %q", lista)
	}
	html := between(navegador.get(lista).Body.String(), "<main", "</main>")
	if !strings.Contains(html, "Se creó el sprint 1") {
		t.Errorf("se esperaba el mensaje %q, y no esta en:\n%s", "Se creó el sprint 1", html)
	}
	fila := between(html, "<tbody>", "</tbody>")
	for _, texto := range []string{"<td>1</td>", "Tener el MVP", "26/10/2026", "30/10/2026", "<td>Abierto</td>"} {
		if !strings.Contains(fila, texto) {
			t.Errorf("se esperaba en la lista de sprints %q, y no esta en:\n%s", texto, fila)
		}
	}

	// Recargar la lista no crea otro sprint, y el mensaje ya no aparece.
	html = between(navegador.get(lista).Body.String(), "<main", "</main>")
	if strings.Contains(html, "Se creó el sprint 1") {
		t.Errorf("el mensaje no tenia que repetirse al recargar:\n%s", html)
	}
	if n := len(project.Sprints()); n != 1 || saver.saves != 1 {
		t.Errorf("se esperaba un solo sprint guardado una vez, hay %d sprints y %d guardados", n, saver.saves)
	}
}
