package main

import (
	"io"
	"strings"
	"testing"
	"time"
)

// Escenario: Salir del programa
// Cubre: RN-6
func TestSalirDelPrograma(t *testing.T) {
	reader, writer := io.Pipe()
	defer writer.Close()
	done := make(chan error, 1)

	go func() {
		done <- runConsole(reader, io.Discard)
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

	err := runConsole(strings.NewReader("\nsalir\n"), &out)

	if err != nil {
		t.Fatalf("se esperaba que la consola terminara sin error, salio: %v", err)
	}
	want := prompt + prompt
	if out.String() != want {
		t.Fatalf("se esperaba que la consola volviera a pedir una operacion sin mostrar nada mas (%q), salio: %q", want, out.String())
	}
}
