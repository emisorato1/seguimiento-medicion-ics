package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// prompt es el indicador con el que la consola pide una operación.
const prompt = "> "

// dateLayout es el formato en que se escriben las fechas: dd/mm/aaaa.
const dateLayout = "02/01/2006"

// operationNames devuelve las operaciones disponibles, en el orden en que se muestran.
func operationNames() []string {
	return []string{"crear-proyecto", "salir"}
}

// runConsole lee operaciones de in, una por línea, y escribe los resultados en out.
// Trabaja sobre el registro de proyectos que recibe: no guarda datos propios.
// Antes de cada línea muestra el indicador. Termina cuando el usuario ejecuta la
// operación de salir (RN-6).
func runConsole(in io.Reader, out io.Writer, registry *domain.ProjectRegistry) error {
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
		case line == "crear-proyecto":
			createProject(scanner, out, registry)
		default:
			fmt.Fprintf(out, "La operación %q no existe. Operaciones disponibles: %s\n",
				line, strings.Join(operationNames(), ", "))
		}
	}
}

// ask muestra la etiqueta de un dato y devuelve la línea que escribe el usuario.
func ask(scanner *bufio.Scanner, out io.Writer, label string) string {
	fmt.Fprint(out, label)
	scanner.Scan()
	return scanner.Text()
}

// createProject pide los datos de un proyecto, lo crea en el registro y muestra su
// identificador (RN-1, RN-3). No valida nada: las reglas las aplica el núcleo (RN-7).
func createProject(scanner *bufio.Scanner, out io.Writer, registry *domain.ProjectRegistry) {
	name := ask(scanner, out, "Nombre: ")
	description := ask(scanner, out, "Descripción (opcional): ")
	start := parseDate(ask(scanner, out, "Fecha de inicio (dd/mm/aaaa): "))
	end := parseDate(ask(scanner, out, "Fecha de finalización (dd/mm/aaaa, opcional): "))

	project, err := registry.Create(domain.ProjectData{
		Name: name, Description: description, Start: start, End: end,
	})
	// Mostrar el motivo cuando el núcleo rechaza el proyecto es la tarea 5.
	if err == nil {
		fmt.Fprintf(out, "Proyecto creado con identificador %d\n", project.ID())
	}
}

// parseDate convierte una fecha escrita dd/mm/aaaa en un día, sin hora. Una línea vacía
// es "sin fecha". Una fecha mal escrita todavía no se informa: es la tarea 14.
func parseDate(text string) time.Time {
	date, _ := time.Parse(dateLayout, text)
	return date
}
