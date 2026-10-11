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
	start, ok := askDate(scanner, out, "Fecha de inicio (dd/mm/aaaa): ", "fecha de inicio")
	if !ok {
		return
	}
	end, ok := askDate(scanner, out, "Fecha de finalización (dd/mm/aaaa, opcional): ", "fecha de finalización")
	if !ok {
		return
	}

	project, err := registry.Create(domain.ProjectData{
		Name: name, Description: description, Start: start, End: end,
	})
	if err != nil {
		fmt.Fprintln(out, err)
		return
	}
	fmt.Fprintf(out, "Proyecto creado con identificador %d\n", project.ID())
}

// askDate pide una fecha. Si no tiene la forma dd/mm/aaaa lo informa nombrando el dato y
// devuelve false: la operación no se ejecuta (RN-5).
func askDate(scanner *bufio.Scanner, out io.Writer, label, field string) (time.Time, bool) {
	date, err := parseDate(ask(scanner, out, label))
	if err != nil {
		fmt.Fprintf(out, "La %s no tiene la forma dd/mm/aaaa: la operación no se ejecuta.\n", field)
		return time.Time{}, false
	}
	return date, true
}

// parseDate convierte una fecha escrita dd/mm/aaaa en un día, sin hora. Una línea vacía
// es "sin fecha" y no es un error: si el dato es obligatorio, lo rechaza el núcleo (RN-7).
// Cualquier otro texto que no tenga esa forma es un error de forma (RN-5).
func parseDate(text string) (time.Time, error) {
	if text == "" {
		return time.Time{}, nil
	}
	return time.Parse(dateLayout, text)
}
