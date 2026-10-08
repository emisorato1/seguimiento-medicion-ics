package main

import (
	"bufio"
	"fmt"
	"io"
)

// prompt es el indicador con el que la consola pide una operación.
const prompt = "> "

// runConsole lee operaciones de in, una por línea, y escribe los resultados en out.
// Antes de cada línea muestra el indicador. Termina cuando el usuario ejecuta la
// operación de salir (RN-6).
func runConsole(in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprint(out, prompt)
		if !scanner.Scan() {
			return scanner.Err()
		}
		if scanner.Text() == "salir" {
			return nil
		}
	}
}
