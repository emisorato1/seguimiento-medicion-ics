package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// prompt es el indicador con el que la consola pide una operación.
const prompt = "> "

// dateLayout es el formato en que se escriben las fechas: dd/mm/aaaa.
const dateLayout = "02/01/2006"

// exitOperation es la operación que termina la consola (RN-6). No está en la tabla
// porque no ejecuta nada: corta el ciclo.
const exitOperation = "salir"

// operation es una operación de la consola: su nombre y lo que ejecuta.
type operation struct {
	name string
	run  func(scanner *bufio.Scanner, out io.Writer, registry *domain.ProjectRegistry)
}

// operations son las operaciones de la consola, en el orden en que se muestran. Es la
// única lista: la usan el ciclo para ejecutarlas y el mensaje de operación inexistente.
var operations = []operation{
	{"crear-proyecto", createProject},
	{"agregar-integrante", addMember},
	{"listar-integrantes", listMembers},
	{"listar-backlog", listBacklog},
}

// operationNames devuelve las operaciones disponibles, en el orden en que se muestran.
func operationNames() []string {
	names := make([]string, 0, len(operations)+1)
	for _, op := range operations {
		names = append(names, op.name)
	}
	return append(names, exitOperation)
}

// findOperation busca una operación de la tabla por su nombre.
func findOperation(name string) (operation, bool) {
	for _, op := range operations {
		if op.name == name {
			return op, true
		}
	}
	return operation{}, false
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
		case line == exitOperation:
			return nil
		default:
			if op, ok := findOperation(line); ok {
				op.run(scanner, out, registry)
			} else {
				fmt.Fprintf(out, "La operación %q no existe. Operaciones disponibles: %s\n",
					line, strings.Join(operationNames(), ", "))
			}
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

// askProject pide el identificador de un proyecto y lo busca en el registro. Si no existe,
// muestra el mensaje del núcleo y devuelve nil (RN-2). Un identificador escrito con letras
// todavía no se informa como error de forma: es la tarea 12.
func askProject(scanner *bufio.Scanner, out io.Writer, registry *domain.ProjectRegistry) *domain.Project {
	id, _ := strconv.Atoi(ask(scanner, out, "Proyecto (identificador): "))
	project, err := registry.Get(id)
	if err != nil {
		fmt.Fprintln(out, err)
		return nil
	}
	return project
}

// addMember pide un integrante y lo agrega al proyecto (RN-1). El rol se pasa como texto
// y lo convierte el dominio: la consola no valida roles (RN-7).
func addMember(scanner *bufio.Scanner, out io.Writer, registry *domain.ProjectRegistry) {
	project := askProject(scanner, out, registry)
	if project == nil {
		return
	}
	name := ask(scanner, out, "Nombre y apellido: ")
	role, err := domain.ParseRole(ask(scanner, out, "Rol (Agile Enabler o Product Builder): "))
	if err != nil {
		fmt.Fprintln(out, err)
		return
	}
	if _, err := project.AddMember(name, role); err != nil {
		fmt.Fprintln(out, err)
		return
	}
	fmt.Fprintf(out, "Integrante agregado al proyecto %d\n", project.ID())
}

// listMembers muestra los integrantes de un proyecto con su rol (RN-1).
func listMembers(scanner *bufio.Scanner, out io.Writer, registry *domain.ProjectRegistry) {
	project := askProject(scanner, out, registry)
	if project == nil {
		return
	}
	for _, member := range project.Members() {
		fmt.Fprintf(out, "- %s (%s)\n", member.Name(), member.Role())
	}
}

// listBacklog muestra las historias del backlog de un proyecto, o avisa que no tiene.
func listBacklog(scanner *bufio.Scanner, out io.Writer, registry *domain.ProjectRegistry) {
	project := askProject(scanner, out, registry)
	if project == nil {
		return
	}
	if len(project.Backlog()) == 0 {
		fmt.Fprintln(out, "No hay historias cargadas en el backlog.")
	}
}
