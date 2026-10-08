package main

import (
	"bufio"
	"io"
)

// runConsole lee operaciones de in, una por línea, y escribe los resultados en out.
// Termina cuando el usuario ejecuta la operación de salir (RN-6).
func runConsole(in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		if scanner.Text() == "salir" {
			return nil
		}
	}
	return scanner.Err()
}
