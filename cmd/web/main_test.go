package main

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// Escenario: Abrir el sistema con un solo comando y sin internet
// Cubre: RN-2. Esta es la parte del comando; que la pagina no pida nada de afuera lo prueba
// internal/web.
func TestAbrirElSistemaConUnSoloComandoYSinInternet(t *testing.T) {
	t.Run("sin opciones escucha solo en esta computadora, en el 8080", func(t *testing.T) {
		opts, err := parseOptions(nil)
		if err != nil {
			t.Fatalf("no se esperaba error, salio: %v", err)
		}
		if opts.addr() != "127.0.0.1:8080" {
			t.Errorf("se esperaba escuchar en 127.0.0.1:8080, salio %q", opts.addr())
		}
		if opts.dataPath != "datos-ics.json" {
			t.Errorf("se esperaban los datos en datos-ics.json, salio %q", opts.dataPath)
		}
	})

	t.Run("respeta -puerto, -datos y -red", func(t *testing.T) {
		opts, err := parseOptions([]string{"-puerto", "9090", "-datos", "otros.json", "-red"})
		if err != nil {
			t.Fatalf("no se esperaba error, salio: %v", err)
		}
		if opts.addr() != ":9090" {
			t.Errorf("con -red se esperaba escuchar en toda la red en el 9090, salio %q", opts.addr())
		}
		if opts.dataPath != "otros.json" {
			t.Errorf("se esperaban los datos en otros.json, salio %q", opts.dataPath)
		}
	})

	t.Run("una opcion que no existe es un error", func(t *testing.T) {
		if _, err := parseOptions([]string{"-color"}); err == nil {
			t.Fatal("se esperaba un error por la opcion -color, y no salio ninguno")
		}
	})

	t.Run("muestra la direccion para abrirlo en el navegador", func(t *testing.T) {
		opts, _ := parseOptions([]string{"-puerto", "9090"})
		var out bytes.Buffer

		showAddress(&out, opts)

		if !strings.Contains(out.String(), "http://127.0.0.1:9090") {
			t.Errorf("se esperaba la direccion http://127.0.0.1:9090, salio:\n%s", out.String())
		}
	})

	t.Run("en esa direccion la lista de proyectos se ve completa, con sus estilos", func(t *testing.T) {
		opts, _ := parseOptions([]string{"-datos", filepath.Join(t.TempDir(), "datos-ics.json")})

		handler, err := open(opts)
		if err != nil {
			t.Fatalf("no se esperaba error al abrir sin datos guardados, salio: %v", err)
		}

		for _, pagina := range []string{"/proyectos", "/static/app.css"} {
			res := httptest.NewRecorder()
			handler.ServeHTTP(res, httptest.NewRequest(http.MethodGet, pagina, nil))
			if res.Code != http.StatusOK {
				t.Errorf("se esperaba %s con el estado %d, salio %d", pagina, http.StatusOK, res.Code)
			}
		}
	})
}

// Escenario: Los datos guardados estan danados
// Cubre: RN-24
func TestLosDatosGuardadosEstanDanados(t *testing.T) {
	path := filepath.Join(t.TempDir(), "datos-ics.json")
	danado := []byte(`{"next_id": 2, "projects": [{"id": 1, "name": ""`)
	if err := os.WriteFile(path, danado, 0o644); err != nil {
		t.Fatalf("no se pudo preparar el archivo danado: %v", err)
	}
	opts, _ := parseOptions([]string{"-datos", path})

	_, err := open(opts)

	if !errors.Is(err, domain.ErrSavedDataInvalid) {
		t.Fatalf("se esperaba ErrSavedDataInvalid, salio: %v", err)
	}
	quedo, _ := os.ReadFile(path)
	if !bytes.Equal(quedo, danado) {
		t.Errorf("los datos guardados tenian que quedar como estaban, y cambiaron a:\n%s", quedo)
	}
}
