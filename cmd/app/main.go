// Command app es el punto de entrada de la aplicacion de seguimiento y medicion
// de proyectos de software.
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

const version = "0.1.0"

// main arranca la consola conectada al teclado y la pantalla, con un registro de proyectos
// nuevo (tarea 15 de la historia 24). Los datos viven en memoria mientras el programa corre.
func main() {
	fmt.Printf("Seguimiento y Medición ICS %s\n", version)
	fmt.Printf("Escribí el nombre de una operación. Operaciones disponibles: %s\n",
		strings.Join(operationNames(), ", "))
	if err := runConsole(os.Stdin, os.Stdout, domain.NewProjectRegistry()); err != nil {
		fmt.Fprintln(os.Stderr, "Error al leer la entrada:", err)
		os.Exit(1)
	}
}
