package main

import (
	"io"
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
