package web

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// Escenario: Crear un proyecto desde la web
// Cubre: RN-1, RN-8, RN-9
func TestCrearUnProyectoDesdeLaWeb(t *testing.T) {
	registry := domain.NewProjectRegistry()
	saver := &fakeSaver{}
	navegador := newBrowser(t, NewServer(registry, saver))
	mensaje := "Se creó el proyecto Sistema de metricas"

	res := navegador.post("/proyectos/nuevo", url.Values{
		"nombre": {"Sistema de metricas"},
		"inicio": {"2026-11-01"},
	})

	if res.Code != http.StatusSeeOther {
		t.Fatalf("se esperaba el estado %d, salio %d:\n%s", http.StatusSeeOther, res.Code, res.Body.String())
	}
	destino := res.Header().Get("Location")
	if destino != "/proyectos/1" {
		t.Fatalf("se esperaba ir a /proyectos/1, salio %q", destino)
	}

	// El navegador sigue la redireccion: ve el mensaje.
	res = navegador.get(destino)
	if res.Code != http.StatusOK {
		t.Fatalf("se esperaba ver el proyecto con el estado %d, salio %d", http.StatusOK, res.Code)
	}
	if html := res.Body.String(); !strings.Contains(html, mensaje) {
		t.Errorf("se esperaba el mensaje %q, y no esta en:\n%s", mensaje, html)
	}

	// Se recarga la pagina: el mensaje ya no aparece.
	res = navegador.get(destino)
	if res.Code != http.StatusOK {
		t.Fatalf("se esperaba ver el proyecto con el estado %d, salio %d", http.StatusOK, res.Code)
	}
	if html := res.Body.String(); strings.Contains(html, mensaje) {
		t.Errorf("al recargar no se esperaba el mensaje %q, y aparece en:\n%s", mensaje, html)
	}

	if cantidad := len(registry.Projects()); cantidad != 1 {
		t.Fatalf("se esperaba un solo proyecto despues de recargar, salieron %d", cantidad)
	}
	if saver.saves != 1 {
		t.Errorf("se esperaba que el registro se guardara una vez, se guardo %d", saver.saves)
	}

	lista := navegador.get("/proyectos").Body.String()
	for _, texto := range []string{"<td>1</td>", `<a href="/proyectos/1">Sistema de metricas</a>`} {
		if !strings.Contains(lista, texto) {
			t.Errorf("se esperaba el proyecto en la lista con el numero 1 (%s), y no esta en:\n%s", texto, lista)
		}
	}
}
