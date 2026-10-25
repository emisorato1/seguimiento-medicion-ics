package web

import (
	"net/http"
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
