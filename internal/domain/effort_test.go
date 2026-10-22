package domain

import (
	"errors"
	"math"
	"strconv"
	"testing"
	"time"
)

// mismasHorasReales compara dos cantidades de horas con una tolerancia chica: comparar
// decimales con == falla por el redondeo de la computadora. Se llama distinto que el de la
// #16 para que las dos historias no choquen al juntarse en dev.
func mismasHorasReales(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}

// hoy es el día que los tests pasan como "hoy": el dominio no lee el reloj (AGENTS.md).
var hoy = unDia(2026, 10, 21)

// unProyectoConJuliana arma un proyecto con la integrante "Juliana Bustos" y la historia 1.
func unProyectoConJuliana(t *testing.T) *Project {
	t.Helper()
	project := unProyecto(t)
	if _, err := project.AddMember("Juliana Bustos", RoleProductBuilder); err != nil {
		t.Fatalf("no se pudo agregar la integrante de prueba: %v", err)
	}
	if _, err := project.AddStory(unosDatos()); err != nil {
		t.Fatalf("no se pudo crear la historia de prueba: %v", err)
	}
	return project
}

// unEsfuerzo devuelve los datos de un registro válido: 2 horas de Juliana Bustos en la
// historia 1, hoy. Cada test cambia solo lo que prueba.
func unEsfuerzo() EffortData {
	return EffortData{StoryID: 1, MemberName: "Juliana Bustos", Day: hoy, Hours: 2, Activity: "Escribir los tests"}
}

// Escenario: Registrar horas trabajadas en una historia
// Cubre: RN-1, RN-15
func TestRegistrarHorasTrabajadasEnUnaHistoria(t *testing.T) {
	project := unProyectoConJuliana(t)
	data := unEsfuerzo()
	data.Hours = 3
	data.Activity = "Escribir los tests de la estimacion"

	effort, err := project.LogEffort(data, hoy)

	if err != nil {
		t.Fatalf("se esperaba registrar el esfuerzo sin error, salio: %v", err)
	}
	if effort.StoryID() != 1 || effort.MemberName() != "Juliana Bustos" || !effort.Day().Equal(hoy) ||
		!mismasHorasReales(effort.Hours(), 3) || effort.Activity() != "Escribir los tests de la estimacion" {
		t.Fatalf("se esperaba el registro de la historia 1, Juliana Bustos, %v, 3 horas y la actividad, salio %d, %q, %v, %v, %q",
			hoy, effort.StoryID(), effort.MemberName(), effort.Day(), effort.Hours(), effort.Activity())
	}
	efforts := project.Efforts()
	if len(efforts) != 1 || efforts[0] != effort {
		t.Fatalf("se esperaba que el proyecto tuviera ese unico registro, tiene %d", len(efforts))
	}
}

// Escenario: Registrar media hora
// Cubre: RN-2
func TestRegistrarMediaHora(t *testing.T) {
	project := unProyectoConJuliana(t)
	data := unEsfuerzo()
	data.Hours = 0.5

	effort, err := project.LogEffort(data, hoy)

	if err != nil {
		t.Fatalf("se esperaba registrar media hora sin error, salio: %v", err)
	}
	if !mismasHorasReales(effort.Hours(), 0.5) {
		t.Fatalf("se esperaba el registro con media hora, salio con %v", effort.Hours())
	}
}

// Esquema del escenario: Rechazar horas que no son mayores que cero
// Cubre: RN-3, RN-18
func TestRechazarHorasQueNoSonMayoresQueCero(t *testing.T) {
	cases := []struct {
		name  string
		hours float64
	}{
		{"0", 0},
		{"-2", -2},
		{"infinito", math.Inf(1)},
		{"no es un numero", math.NaN()},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			project := unProyectoConJuliana(t)
			data := unEsfuerzo()
			data.Hours = tc.hours

			_, err := project.LogEffort(data, hoy)

			if !errors.Is(err, ErrEffortHoursInvalid) {
				t.Fatalf("se esperaba ErrEffortHoursInvalid con %v horas, salio: %v", tc.hours, err)
			}
			if len(project.Efforts()) != 0 {
				t.Fatalf("se esperaba que el proyecto siguiera sin registros, tiene %d", len(project.Efforts()))
			}
		})
	}
}

// Escenario: Rechazar horas de un dia que todavia no paso
// Cubre: RN-4, RN-18
func TestRechazarHorasDeUnDiaQueTodaviaNoPaso(t *testing.T) {
	project := unProyectoConJuliana(t)
	data := unEsfuerzo()
	data.Day = hoy.AddDate(0, 0, 1)

	_, err := project.LogEffort(data, hoy)

	if !errors.Is(err, ErrEffortDayInFuture) {
		t.Fatalf("se esperaba ErrEffortDayInFuture con fecha de mañana, salio: %v", err)
	}
	if len(project.Efforts()) != 0 {
		t.Fatalf("se esperaba que el proyecto siguiera sin registros, tiene %d", len(project.Efforts()))
	}
}

// Escenario: Registrar horas de un dia anterior al inicio del proyecto
// Cubre: RN-17
func TestRegistrarHorasDeUnDiaAnteriorAlInicioDelProyecto(t *testing.T) {
	project := unProyectoConJuliana(t) // empezó el 01/10/2026
	data := unEsfuerzo()
	data.Day = unDia(2026, 9, 15)

	_, err := project.LogEffort(data, hoy)

	if err != nil {
		t.Fatalf("se esperaba registrar un día anterior al inicio del proyecto sin error, salio: %v", err)
	}
	if len(project.Efforts()) != 1 {
		t.Fatalf("se esperaba que quedara guardado el registro, hay %d", len(project.Efforts()))
	}
}

// Escenario: Rechazar un registro sin fecha
// Cubre: RN-5, RN-18
func TestRechazarUnRegistroSinFecha(t *testing.T) {
	project := unProyectoConJuliana(t)
	data := unEsfuerzo()
	data.Day = time.Time{}

	_, err := project.LogEffort(data, hoy)

	if !errors.Is(err, ErrEffortDayRequired) {
		t.Fatalf("se esperaba ErrEffortDayRequired sin fecha, salio: %v", err)
	}
	if len(project.Efforts()) != 0 {
		t.Fatalf("se esperaba que el proyecto siguiera sin registros, tiene %d", len(project.Efforts()))
	}
}

// Esquema del escenario: Rechazar un registro sin actividad realizada
// Cubre: RN-6, RN-18
func TestRechazarUnRegistroSinActividadRealizada(t *testing.T) {
	for _, activity := range []string{"", "   "} {
		t.Run(strconv.Quote(activity), func(t *testing.T) {
			project := unProyectoConJuliana(t)
			data := unEsfuerzo()
			data.Activity = activity

			_, err := project.LogEffort(data, hoy)

			if !errors.Is(err, ErrEffortActivityRequired) {
				t.Fatalf("se esperaba ErrEffortActivityRequired con la actividad %q, salio: %v", activity, err)
			}
			if len(project.Efforts()) != 0 {
				t.Fatalf("se esperaba que el proyecto siguiera sin registros, tiene %d", len(project.Efforts()))
			}
		})
	}
}

// Escenario: Guardar la actividad sin los espacios de los extremos
// Cubre: RN-6
func TestGuardarLaActividadSinLosEspaciosDeLosExtremos(t *testing.T) {
	project := unProyectoConJuliana(t)
	data := unEsfuerzo()
	data.Activity = "  Revisar el pull request  "

	effort, err := project.LogEffort(data, hoy)

	if err != nil {
		t.Fatalf("se esperaba registrar el esfuerzo sin error, salio: %v", err)
	}
	if effort.Activity() != "Revisar el pull request" {
		t.Fatalf("se esperaba la actividad %q, salio %q", "Revisar el pull request", effort.Activity())
	}
}

// Escenario: Rechazar el registro de alguien que no es integrante del proyecto
// Cubre: RN-7, RN-18
func TestRechazarElRegistroDeAlguienQueNoEsIntegranteDelProyecto(t *testing.T) {
	project := unProyectoConJuliana(t)
	data := unEsfuerzo()
	data.MemberName = "Pedro Gomez"

	_, err := project.LogEffort(data, hoy)

	if !errors.Is(err, ErrMemberNotFound) {
		t.Fatalf("se esperaba ErrMemberNotFound para Pedro Gomez, salio: %v", err)
	}
	if len(project.Efforts()) != 0 {
		t.Fatalf("se esperaba que el proyecto siguiera sin registros, tiene %d", len(project.Efforts()))
	}
}

// Escenario: Encontrar al integrante aunque el nombre venga escrito distinto
// Cubre: RN-7, RN-8
func TestEncontrarAlIntegranteAunqueElNombreVengaEscritoDistinto(t *testing.T) {
	project := unProyectoConJuliana(t)
	data := unEsfuerzo()
	data.MemberName = "  juliana bustos  "

	effort, err := project.LogEffort(data, hoy)

	if err != nil {
		t.Fatalf("se esperaba encontrar a la integrante, salio: %v", err)
	}
	if effort.MemberName() != "Juliana Bustos" {
		t.Fatalf("se esperaba el registro a nombre de %q, salio %q", "Juliana Bustos", effort.MemberName())
	}
}

// Escenario: Rechazar el registro sobre una historia que no es del proyecto
// Cubre: RN-9, RN-18
func TestRechazarElRegistroSobreUnaHistoriaQueNoEsDelProyecto(t *testing.T) {
	project := unProyectoConJuliana(t)
	data := unEsfuerzo()
	data.StoryID = 99

	_, err := project.LogEffort(data, hoy)

	if !errors.Is(err, ErrStoryNotFound) {
		t.Fatalf("se esperaba ErrStoryNotFound para la historia 99, salio: %v", err)
	}
	if len(project.Efforts()) != 0 {
		t.Fatalf("se esperaba que el proyecto siguiera sin registros, tiene %d", len(project.Efforts()))
	}
}

// Escenario: Registrar horas sobre una historia terminada
// Cubre: RN-10
func TestRegistrarHorasSobreUnaHistoriaTerminada(t *testing.T) {
	project := unProyectoConJuliana(t)
	story, _ := project.Story(1)
	// Terminar una historia es de la #10, que todavia no esta en dev: el test la prepara
	// desde dentro del paquete, como la #12 en su tarea 9.
	story.state = StoryDone

	_, err := project.LogEffort(unEsfuerzo(), hoy)

	if err != nil {
		t.Fatalf("se esperaba registrar horas sobre una historia Terminada sin error, salio: %v", err)
	}
	if len(project.Efforts()) != 1 {
		t.Fatalf("se esperaba que quedara guardado el registro, hay %d", len(project.Efforts()))
	}
}

// Escenario: Dos cargas del mismo integrante, la misma historia y el mismo dia
// Cubre: RN-11, RN-12, RN-15
//
// Las dos cargas llegan con horas del día distintas (9 y 18): las dos tienen que quedar con
// el mismo día y sin la hora (RN-12, que no tiene escenario propio).
func TestDosCargasDelMismoIntegranteLaMismaHistoriaYElMismoDia(t *testing.T) {
	project := unProyectoConJuliana(t)
	first := unEsfuerzo()
	first.Hours = 2
	first.Activity = "Escribir los tests"
	first.Day = hoy.Add(9 * time.Hour)
	second := unEsfuerzo()
	second.Hours = 1
	second.Activity = "Corregir la revision"
	second.Day = hoy.Add(18 * time.Hour)

	if _, err := project.LogEffort(first, hoy); err != nil {
		t.Fatalf("se esperaba registrar la primera carga sin error, salio: %v", err)
	}
	if _, err := project.LogEffort(second, hoy); err != nil {
		t.Fatalf("se esperaba registrar la segunda carga sin error, salio: %v", err)
	}

	efforts := project.Efforts()
	if len(efforts) != 2 {
		t.Fatalf("se esperaban 2 registros de esfuerzo, hay %d", len(efforts))
	}
	if efforts[0].Activity() != "Escribir los tests" || efforts[1].Activity() != "Corregir la revision" {
		t.Fatalf("se esperaban los registros en el orden de carga, salieron %q y %q", efforts[0].Activity(), efforts[1].Activity())
	}
	for _, effort := range efforts {
		if !effort.Day().Equal(hoy) {
			t.Fatalf("se esperaba el dia %v sin la hora, salio %v", hoy, effort.Day())
		}
	}
	hours, err := project.StoryHours(1)
	if err != nil || !mismasHorasReales(hours, 3) {
		t.Fatalf("se esperaba que la historia 1 tuviera 3 horas, salio %v (error: %v)", hours, err)
	}
}

// cargar registra esfuerzo en el proyecto o corta el test: es para preparar los datos de las
// consultas, no para probar el registro.
func cargar(t *testing.T, project *Project, storyID int, name string, hours float64) {
	t.Helper()
	data := unEsfuerzo()
	data.StoryID = storyID
	data.MemberName = name
	data.Hours = hours
	if _, err := project.LogEffort(data, hoy); err != nil {
		t.Fatalf("no se pudo cargar el esfuerzo de prueba: %v", err)
	}
}

// Escenario: Consultar el total de horas de una historia
// Cubre: RN-13
func TestConsultarElTotalDeHorasDeUnaHistoria(t *testing.T) {
	project := unProyectoConJuliana(t)
	cargar(t, project, 1, "Juliana Bustos", 2)
	cargar(t, project, 1, "Juliana Bustos", 1.5)

	hours, err := project.StoryHours(1)

	if err != nil || !mismasHorasReales(hours, 3.5) {
		t.Fatalf("se esperaban 3 horas y media, salio %v (error: %v)", hours, err)
	}
}

// Esquema del escenario: Sin registros, el total es cero
// Cubre: RN-13, RN-14
func TestSinRegistrosElTotalEsCero(t *testing.T) {
	cases := []struct {
		name string
		ask  func(project *Project) (float64, error)
	}{
		{"la historia 1", func(project *Project) (float64, error) { return project.StoryHours(1) }},
		{"Juliana Bustos", func(project *Project) (float64, error) { return project.MemberHours("Juliana Bustos") }},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			project := unProyectoConJuliana(t)

			hours, err := tc.ask(project)

			if err != nil {
				t.Fatalf("se esperaba que no fuera un error, salio: %v", err)
			}
			if hours != 0 {
				t.Fatalf("se esperaban cero horas, salio %v", hours)
			}
		})
	}
}

// Esquema del escenario: Consultar las horas de algo que no existe
// Cubre: RN-13, RN-14
func TestConsultarLasHorasDeAlgoQueNoExiste(t *testing.T) {
	cases := []struct {
		name string
		ask  func(project *Project) (float64, error)
		want error
	}{
		{"la historia 99", func(project *Project) (float64, error) { return project.StoryHours(99) }, ErrStoryNotFound},
		{"Pedro Gomez", func(project *Project) (float64, error) { return project.MemberHours("Pedro Gomez") }, ErrMemberNotFound},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			project := unProyectoConJuliana(t)

			_, err := tc.ask(project)

			if !errors.Is(err, tc.want) {
				t.Fatalf("se esperaba %v, salio: %v", tc.want, err)
			}
		})
	}
}

// Escenario: Consultar el total de horas de un integrante
// Cubre: RN-14
//
// Se consulta con otras mayúsculas y espacios de más: el nombre se busca igual que al
// registrar (RN-14).
func TestConsultarElTotalDeHorasDeUnIntegrante(t *testing.T) {
	project := unProyectoConJuliana(t)
	if _, err := project.AddStory(unosDatos()); err != nil {
		t.Fatalf("no se pudo crear la segunda historia de prueba: %v", err)
	}
	cargar(t, project, 1, "Juliana Bustos", 2)
	cargar(t, project, 2, "Juliana Bustos", 3)

	hours, err := project.MemberHours("  JULIANA bustos ")

	if err != nil || !mismasHorasReales(hours, 5) {
		t.Fatalf("se esperaban 5 horas, salio %v (error: %v)", hours, err)
	}
}
