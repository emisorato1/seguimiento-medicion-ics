package domain

import (
	"errors"
	"strings"
	"testing"
)

// Tarea 2 de la Fase 3 de la historia #4. No tiene escenario BDD propio: los escenarios
// empiezan en la tarea 4, cuando ya se puede agregar un integrante a un proyecto.
// Cubre: RN-1
func TestUnIntegranteExponeSuNombreYSuRol(t *testing.T) {
	member := &Member{
		name: "Juliana Bustos",
		role: RoleProductBuilder,
	}

	if member.Name() != "Juliana Bustos" {
		t.Fatalf("se esperaba el nombre y apellido %q, salio: %q", "Juliana Bustos", member.Name())
	}
	if member.Role() != RoleProductBuilder {
		t.Fatalf("se esperaba el rol %v, salio: %v", RoleProductBuilder, member.Role())
	}
}

// Escenario: Listar un proyecto sin integrantes
// Cubre: RN-8
func TestListarUnProyectoSinIntegrantes(t *testing.T) {
	registry := NewProjectRegistry()
	project, err := registry.Create(ProjectData{Name: "Sistema de metricas", Start: date(2026, 10, 1)})
	if err != nil {
		t.Fatalf("se esperaba registrar el proyecto de prueba sin error, salio: %v", err)
	}

	members := project.Members()

	if len(members) != 0 {
		t.Fatalf("se esperaba una lista vacia de integrantes, salieron %d", len(members))
	}
}

// Escenario: Agregar el primer integrante
// Cubre: RN-1, RN-8
func TestAgregarElPrimerIntegrante(t *testing.T) {
	registry := NewProjectRegistry()
	project, err := registry.Create(ProjectData{Name: "Sistema de metricas", Start: date(2026, 10, 1)})
	if err != nil {
		t.Fatalf("se esperaba registrar el proyecto de prueba sin error, salio: %v", err)
	}

	member, err := project.AddMember("Juliana Bustos", RoleProductBuilder)

	if err != nil {
		t.Fatalf("se esperaba que el integrante se agregara, salio: %v", err)
	}
	if member.Name() != "Juliana Bustos" || member.Role() != RoleProductBuilder {
		t.Fatalf("se esperaba a %q con rol %v, salio: %q con rol %v",
			"Juliana Bustos", RoleProductBuilder, member.Name(), member.Role())
	}
	members := project.Members()
	if len(members) != 1 {
		t.Fatalf("se esperaba que el proyecto tuviera 1 integrante, salieron %d", len(members))
	}
	if members[0].Name() != "Juliana Bustos" || members[0].Role() != RoleProductBuilder {
		t.Fatalf("se esperaba que la lista incluyera a %q con rol %v, salio: %q con rol %v",
			"Juliana Bustos", RoleProductBuilder, members[0].Name(), members[0].Role())
	}
}

// Escenario: Guardar el nombre sin espacios en los extremos
// Cubre: RN-3
func TestGuardarElNombreDelIntegranteSinEspaciosEnLosExtremos(t *testing.T) {
	registry := NewProjectRegistry()
	project, err := registry.Create(ProjectData{Name: "Sistema de metricas", Start: date(2026, 10, 1)})
	if err != nil {
		t.Fatalf("se esperaba registrar el proyecto de prueba sin error, salio: %v", err)
	}

	member, err := project.AddMember("  Juliana Bustos  ", RoleProductBuilder)

	if err != nil {
		t.Fatalf("se esperaba que el integrante se agregara, salio: %v", err)
	}
	if member.Name() != "Juliana Bustos" {
		t.Fatalf("se esperaba el nombre sin espacios en los extremos %q, salio: %q", "Juliana Bustos", member.Name())
	}
	if got := project.Members()[0].Name(); got != "Juliana Bustos" {
		t.Fatalf("se esperaba que la lista incluyera a %q, salio: %q", "Juliana Bustos", got)
	}
}

// Esquema del escenario: Rechazar un integrante sin nombre y apellido
// Cubre: RN-2, RN-12
func TestRechazarUnIntegranteSinNombreYApellido(t *testing.T) {
	casos := []string{"", "   "}

	for _, nombre := range casos {
		t.Run(nombre, func(t *testing.T) {
			registry := NewProjectRegistry()
			project, err := registry.Create(ProjectData{Name: "Sistema de metricas", Start: date(2026, 10, 1)})
			if err != nil {
				t.Fatalf("se esperaba registrar el proyecto de prueba sin error, salio: %v", err)
			}

			_, err = project.AddMember(nombre, RoleProductBuilder)

			if !errors.Is(err, ErrMemberNameRequired) {
				t.Fatalf("se esperaba ErrMemberNameRequired para %q, salio: %v", nombre, err)
			}
			if len(project.Members()) != 0 {
				t.Fatalf("se esperaba que el proyecto siguiera sin integrantes, salieron %d", len(project.Members()))
			}
		})
	}
}

// Tarea 1: el dominio convierte el texto que escribe el usuario en un rol, ignorando
// mayusculas y los espacios de los extremos.
// Cubre: RN-4
func TestConvertirElTextoEnUnRol(t *testing.T) {
	casos := []struct {
		texto    string
		esperado Role
	}{
		{"Agile Enabler", RoleAgileEnabler},
		{"agile enabler", RoleAgileEnabler},
		{"  AGILE ENABLER  ", RoleAgileEnabler},
		{"Product Builder", RoleProductBuilder},
		{"product builder", RoleProductBuilder},
		{"  Product BUILDER ", RoleProductBuilder},
	}

	for _, caso := range casos {
		t.Run(caso.texto, func(t *testing.T) {
			rol, err := ParseRole(caso.texto)
			if err != nil {
				t.Fatalf("se esperaba que %q fuera un rol valido, salio: %v", caso.texto, err)
			}
			if rol != caso.esperado {
				t.Fatalf("se esperaba el rol %v para %q, salio: %v", caso.esperado, caso.texto, rol)
			}
		})
	}
}

// Tarea 1: cualquier texto que no sea uno de los dos roles se rechaza con el error de la
// RN-4. Solo se ignoran los espacios de los extremos, no los del medio.
// Cubre: RN-4
func TestRechazarUnTextoQueNoEsUnRol(t *testing.T) {
	casos := []string{"Product Owner", "Scrum Master", "", "   ", "AgileEnabler", "Product  Builder"}

	for _, texto := range casos {
		t.Run(texto, func(t *testing.T) {
			_, err := ParseRole(texto)

			if !errors.Is(err, ErrMemberRoleInvalid) {
				t.Fatalf("se esperaba ErrMemberRoleInvalid para %q, salio: %v", texto, err)
			}
		})
	}
}

// Tarea 1: el mensaje del error nombra los dos roles validos, asi la consola solo tiene que
// mostrarlo (RN-4: "informar cuales son los roles validos").
// Cubre: RN-4
func TestElErrorDeRolNombraLosValidos(t *testing.T) {
	_, err := ParseRole("Product Owner")
	if err == nil {
		t.Fatal("se esperaba un error para \"Product Owner\", salio: nil")
	}

	for _, esperado := range []string{"Agile Enabler", "Product Builder"} {
		if !strings.Contains(err.Error(), esperado) {
			t.Fatalf("se esperaba que el mensaje nombrara el rol %q, salio: %q", esperado, err.Error())
		}
	}
}

// Tarea 1: los dos roles se muestran con el nombre que usa la Fase 1, y es el mismo texto
// que ParseRole acepta: lo que se muestra se puede volver a escribir.
// Cubre: RN-4
func TestLosDosRolesSeMuestranYSeVuelvenAConvertir(t *testing.T) {
	casos := map[Role]string{
		RoleAgileEnabler:   "Agile Enabler",
		RoleProductBuilder: "Product Builder",
	}

	for rol, texto := range casos {
		if rol.String() != texto {
			t.Fatalf("se esperaba que el rol se mostrara como %q, salio: %q", texto, rol.String())
		}
		vuelta, err := ParseRole(rol.String())
		if err != nil || vuelta != rol {
			t.Fatalf("se esperaba que %q volviera a convertirse en el mismo rol, salio: %v (%v)", texto, vuelta, err)
		}
	}
}
