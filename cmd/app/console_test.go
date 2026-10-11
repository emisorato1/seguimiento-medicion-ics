package main

import (
	"errors"
	"io"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// Escenario: Salir del programa
// Cubre: RN-6
func TestSalirDelPrograma(t *testing.T) {
	reader, writer := io.Pipe()
	defer writer.Close()
	done := make(chan error, 1)

	go func() {
		done <- runConsole(reader, io.Discard, domain.NewProjectRegistry())
	}()
	go writer.Write([]byte("salir\n"))

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("se esperaba que la consola terminara sin error, salio: %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("se esperaba que la consola terminara al recibir \"salir\", pero siguio esperando entrada")
	}
}

// Escenario: Entrada vacia
// Cubre: caso límite de la Fase 1
func TestEntradaVacia(t *testing.T) {
	var out strings.Builder

	err := runConsole(strings.NewReader("\nsalir\n"), &out, domain.NewProjectRegistry())

	if err != nil {
		t.Fatalf("se esperaba que la consola terminara sin error, salio: %v", err)
	}
	want := prompt + prompt
	if out.String() != want {
		t.Fatalf("se esperaba que la consola volviera a pedir una operacion sin mostrar nada mas (%q), salio: %q", want, out.String())
	}
}

// Escenario: Operacion inexistente
// Cubre: RN-4
func TestOperacionInexistente(t *testing.T) {
	var out strings.Builder

	err := runConsole(strings.NewReader("volar\nsalir\n"), &out, domain.NewProjectRegistry())

	if err != nil {
		t.Fatalf("se esperaba que la consola terminara sin error, salio: %v", err)
	}
	if !strings.Contains(out.String(), "volar") {
		t.Fatalf("se esperaba que la consola informara que la operacion \"volar\" no existe, salio: %q", out.String())
	}
	for _, name := range operationNames() {
		if !strings.Contains(out.String(), name) {
			t.Fatalf("se esperaba que la consola listara la operacion %q, salio: %q", name, out.String())
		}
	}
}

// Escenario: Crear un proyecto desde la consola
// Cubre: RN-1, RN-3
func TestCrearUnProyectoDesdeLaConsola(t *testing.T) {
	registry := domain.NewProjectRegistry()
	var out strings.Builder
	in := "crear-proyecto\nSistema de metricas\n\n01/10/2026\n\nsalir\n"

	err := runConsole(strings.NewReader(in), &out, registry)

	if err != nil {
		t.Fatalf("se esperaba que la consola terminara sin error, salio: %v", err)
	}
	project, err := registry.Get(1)
	if err != nil {
		t.Fatalf("se esperaba el proyecto 1 en el registro, salio: %v", err)
	}
	if project.Name() != "Sistema de metricas" {
		t.Fatalf("se esperaba el nombre %q, salio: %q", "Sistema de metricas", project.Name())
	}
	start := time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC)
	if !project.Start().Equal(start) {
		t.Fatalf("se esperaba el inicio %v, salio: %v", start, project.Start())
	}
	if !strings.Contains(out.String(), strconv.Itoa(project.ID())) {
		t.Fatalf("se esperaba ver el identificador %d en la salida, salio: %q", project.ID(), out.String())
	}
}

// Escenario: Mostrar el motivo de un rechazo
// Cubre: RN-2, RN-7
//
// La consola solo se puede observar por lo que escribe, y la RN-2 pide mostrar el motivo
// tal cual lo informa el nucleo. Por eso el test busca el mensaje del error del dominio:
// no lo escribe a mano, asi que si cambia la redaccion el test sigue valiendo.
func TestMostrarElMotivoDeUnRechazo(t *testing.T) {
	registry := domain.NewProjectRegistry()
	var out strings.Builder
	in := "crear-proyecto\n\n\n01/10/2026\n\nvolar\nsalir\n"

	err := runConsole(strings.NewReader(in), &out, registry)

	if err != nil {
		t.Fatalf("se esperaba que la consola terminara sin error, salio: %v", err)
	}
	if !strings.Contains(out.String(), domain.ErrProjectNameRequired.Error()) {
		t.Fatalf("se esperaba ver el mensaje del nucleo %q, salio: %q",
			domain.ErrProjectNameRequired.Error(), out.String())
	}
	if !strings.Contains(out.String(), "volar") {
		t.Fatalf("se esperaba que la consola siguiera funcionando despues del rechazo, salio: %q", out.String())
	}
	if _, err := registry.Get(1); !errors.Is(err, domain.ErrProjectNotFound) {
		t.Fatalf("se esperaba que no se registrara ningun proyecto, salio: %v", err)
	}
}

// Sin escenario: tarea 14, una fecha mal escrita
// Cubre: RN-5
func TestUnaFechaMalEscritaSeInformaYLaOperacionNoSeEjecuta(t *testing.T) {
	cases := []struct {
		field string
		input string
	}{
		{"fecha de inicio", "crear-proyecto\nSistema de metricas\n\nprimero de octubre\nsalir\n"},
		{"fecha de finalización", "crear-proyecto\nSistema de metricas\n\n01/10/2026\n31-12-2026\nsalir\n"},
	}
	for _, tc := range cases {
		t.Run(tc.field, func(t *testing.T) {
			registry := domain.NewProjectRegistry()
			var out strings.Builder

			err := runConsole(strings.NewReader(tc.input), &out, registry)

			if err != nil {
				t.Fatalf("se esperaba que la consola terminara sin error, salio: %v", err)
			}
			if !strings.Contains(out.String(), tc.field) {
				t.Fatalf("se esperaba que la consola indicara que la %s esta mal escrita, salio: %q", tc.field, out.String())
			}
			if _, err := registry.Get(1); !errors.Is(err, domain.ErrProjectNotFound) {
				t.Fatalf("se esperaba que no se creara ningun proyecto, salio: %v", err)
			}
		})
	}
}
