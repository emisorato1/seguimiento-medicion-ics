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

// Escenario: Agregar y listar integrantes
// Cubre: RN-1
func TestAgregarYListarIntegrantes(t *testing.T) {
	registry := domain.NewProjectRegistry()
	var out strings.Builder
	in := "crear-proyecto\nSistema de metricas\n\n01/10/2026\n\n" +
		"agregar-integrante\n1\nJuliana Bustos\nproduct builder\n" +
		"listar-integrantes\n1\n" +
		"salir\n"

	err := runConsole(strings.NewReader(in), &out, registry)

	if err != nil {
		t.Fatalf("se esperaba que la consola terminara sin error, salio: %v", err)
	}
	project, _ := registry.Get(1)
	members := project.Members()
	if len(members) != 1 || members[0].Name() != "Juliana Bustos" || members[0].Role() != domain.RoleProductBuilder {
		t.Fatalf("se esperaba a Juliana Bustos como Product Builder en el proyecto, salio: %v", members)
	}
	if !strings.Contains(out.String(), "Juliana Bustos") {
		t.Fatalf("se esperaba ver a Juliana Bustos en la lista de integrantes, salio: %q", out.String())
	}
}

// Escenario: Listar un backlog vacio
// Cubre: RN-1
func TestListarUnBacklogVacio(t *testing.T) {
	registry := domain.NewProjectRegistry()
	var out strings.Builder
	in := "crear-proyecto\nSistema de metricas\n\n01/10/2026\n\n" +
		"listar-backlog\n1\n" +
		"salir\n"

	err := runConsole(strings.NewReader(in), &out, registry)

	if err != nil {
		t.Fatalf("se esperaba que la consola terminara sin error, salio: %v", err)
	}
	if !strings.Contains(out.String(), "No hay historias cargadas") {
		t.Fatalf("se esperaba ver que el backlog no tiene historias, salio: %q", out.String())
	}
}

// Escenario: Cargar una historia y listar el backlog
// Cubre: RN-1, RN-3
func TestCargarUnaHistoriaYListarElBacklog(t *testing.T) {
	registry := domain.NewProjectRegistry()
	var out strings.Builder
	in := "crear-proyecto\nSistema de metricas\n\n01/10/2026\n\n" +
		"cargar-historia\n1\nRegistrar esfuerzo\nCargar las horas trabajadas\nalta\n" +
		"Se cargan las horas de una historia\nNo se aceptan horas negativas\n\n" +
		"listar-backlog\n1\n" +
		"salir\n"

	err := runConsole(strings.NewReader(in), &out, registry)

	if err != nil {
		t.Fatalf("se esperaba que la consola terminara sin error, salio: %v", err)
	}
	project, _ := registry.Get(1)
	story, err := project.Story(1)
	if err != nil {
		t.Fatalf("se esperaba la historia 1 en el backlog, salio: %v", err)
	}
	if story.Title() != "Registrar esfuerzo" || story.Priority() != domain.PriorityHigh || len(story.AcceptanceCriteria()) != 2 {
		t.Fatalf("se esperaba la historia con su titulo, prioridad Alta y dos criterios, salio: %q, %v, %v",
			story.Title(), story.Priority(), story.AcceptanceCriteria())
	}
	if !strings.Contains(out.String(), "#1 Registrar esfuerzo") {
		t.Fatalf("se esperaba ver la historia \"Registrar esfuerzo\" con su identificador, salio: %q", out.String())
	}
	if !strings.Contains(out.String(), "sin estimar") {
		t.Fatalf("se esperaba ver la historia como sin estimar, salio: %q", out.String())
	}
}

// unProyectoConUnaHistoria es la entrada que crea el proyecto 1 y le carga la historia 1, el
// "Dado un proyecto con una historia cargada" de los escenarios de estimar.
const unProyectoConUnaHistoria = "crear-proyecto\nSistema de metricas\n\n01/10/2026\n\n" +
	"cargar-historia\n1\nRegistrar esfuerzo\nCargar las horas trabajadas\nalta\n" +
	"Se cargan las horas de una historia\n\n"

// Escenario: Estimar una historia desde la consola
// Cubre: RN-1
func TestEstimarUnaHistoriaDesdeLaConsola(t *testing.T) {
	registry := domain.NewProjectRegistry()
	var out strings.Builder
	in := unProyectoConUnaHistoria + "estimar\n1\n1\n5\nsalir\n"

	err := runConsole(strings.NewReader(in), &out, registry)

	if err != nil {
		t.Fatalf("se esperaba que la consola terminara sin error, salio: %v", err)
	}
	project, _ := registry.Get(1)
	story, _ := project.Story(1)
	if story.StoryPoints() != 5 {
		t.Fatalf("se esperaba la historia 1 estimada en 5, salio %d", story.StoryPoints())
	}
	if !strings.Contains(out.String(), "estimada en 5") {
		t.Fatalf("se esperaba ver que la historia quedo estimada en 5, salio: %q", out.String())
	}
}

// Escenario: El nucleo rechaza un valor fuera de la escala
// Cubre: RN-2, RN-7
//
// Desvío: este test pasa sin código nuevo, porque la tarea 9 ya muestra el mensaje del
// núcleo cuando rechaza la estimación. Queda anotado en la Fase 4.
func TestElNucleoRechazaUnValorFueraDeLaEscala(t *testing.T) {
	registry := domain.NewProjectRegistry()
	var out strings.Builder
	in := unProyectoConUnaHistoria + "estimar\n1\n1\n4\nlistar-backlog\n1\nsalir\n"

	err := runConsole(strings.NewReader(in), &out, registry)

	if err != nil {
		t.Fatalf("se esperaba que la consola terminara sin error, salio: %v", err)
	}
	if !strings.Contains(out.String(), domain.ErrStoryPointsInvalid.Error()) {
		t.Fatalf("se esperaba ver el mensaje del nucleo con los valores validos, salio: %q", out.String())
	}
	if !strings.Contains(out.String(), "#1 Registrar esfuerzo") {
		t.Fatalf("se esperaba que el programa siguiera funcionando y listara el backlog, salio: %q", out.String())
	}
	project, _ := registry.Get(1)
	if story, _ := project.Story(1); story.IsEstimated() {
		t.Fatalf("se esperaba que la historia siguiera sin estimar, salio estimada en %d", story.StoryPoints())
	}
}

// Escenario: Operacion incompleta
// Cubre: RN-5
func TestOperacionIncompleta(t *testing.T) {
	registry := domain.NewProjectRegistry()
	var out strings.Builder
	in := unProyectoConUnaHistoria + "estimar\n1\n\n5\nsalir\n"

	err := runConsole(strings.NewReader(in), &out, registry)

	if err != nil {
		t.Fatalf("se esperaba que la consola terminara sin error, salio: %v", err)
	}
	if !strings.Contains(out.String(), "Falta indicar la historia") {
		t.Fatalf("se esperaba ver que falta indicar la historia, salio: %q", out.String())
	}
	project, _ := registry.Get(1)
	if story, _ := project.Story(1); story.IsEstimated() {
		t.Fatalf("se esperaba que la operacion no se ejecutara, la historia salio estimada en %d", story.StoryPoints())
	}
}
