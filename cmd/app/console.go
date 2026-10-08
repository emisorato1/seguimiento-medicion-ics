package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// prompt es el indicador con el que la consola pide una operación.
const prompt = "> "

// operationNames devuelve las operaciones disponibles, en el orden en que se muestran.
func operationNames() []string {
	return []string{"salir"}
}

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
		line := scanner.Text()
		switch {
		case line == "":
			// una línea vacía no hace nada y vuelve a pedir una operación
		case line == "salir":
			return nil
		default:
			fmt.Fprintf(out, "La operación %q no existe. Operaciones disponibles: %s\n",
				line, strings.Join(operationNames(), ", "))
		}
	}
}
