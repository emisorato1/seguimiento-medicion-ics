package domain

import (
	"math"
	"testing"
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
