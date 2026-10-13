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

// Escenario: Rechazar un rol invalido
// Cubre: RN-4, RN-12
//
// El texto "Product Owner" ya lo rechaza ParseRole (tarea 1). Este test cubre lo que llega
// a AddMember sin pasar por la conversion: un rol sin inicializar y uno fuera de rango.
func TestRechazarUnRolInvalido(t *testing.T) {
	casos := []struct {
		nombre string
		rol    Role
	}{
		{"sin rol", Role(0)},
		{"fuera de rango", Role(99)},
	}

	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			registry := NewProjectRegistry()
			project, err := registry.Create(ProjectData{Name: "Sistema de metricas", Start: date(2026, 10, 1)})
			if err != nil {
				t.Fatalf("se esperaba registrar el proyecto de prueba sin error, salio: %v", err)
			}

			_, err = project.AddMember("Maximiliano Eula", caso.rol)

			if !errors.Is(err, ErrMemberRoleInvalid) {
				t.Fatalf("se esperaba ErrMemberRoleInvalid con un rol %s, salio: %v", caso.nombre, err)
			}
			if len(project.Members()) != 0 {
				t.Fatalf("se esperaba que el proyecto siguiera sin integrantes, salieron %d", len(project.Members()))
			}
		})
	}
}

// Tarea 7: el valor cero del rol no es ninguno de los dos roles validos, asi un rol sin
// inicializar no se confunde con Agile Enabler. Se muestra como "sin rol".
// Cubre: RN-4
func TestUnRolSinInicializarNoEsNingunoDeLosValidos(t *testing.T) {
	var rol Role

	if rol == RoleAgileEnabler || rol == RoleProductBuilder {
		t.Fatalf("se esperaba que el rol sin inicializar no fuera ninguno de los validos, salio: %v", rol)
	}
	if rol.String() != "sin rol" {
		t.Fatalf("se esperaba que el rol sin inicializar se mostrara como %q, salio: %q", "sin rol", rol.String())
	}
}

// Esquema del escenario: Rechazar un integrante duplicado
// Cubre: RN-5, RN-12
//
// Las tres primeras filas son las de los Ejemplos de la Fase 1. Las dos ultimas son la
// decision de la Fase 2: las mayusculas se ignoran tambien en letras con tilde y la ñ.
func TestRechazarUnIntegranteDuplicado(t *testing.T) {
	casos := []struct {
		cargado string
		nuevo   string
	}{
		{"Juliana Bustos", "Juliana Bustos"},
		{"Juliana Bustos", "juliana bustos"},
		{"Juliana Bustos", "  JULIANA BUSTOS  "},
		{"Íñigo Peña", "íñigo peña"},
		{"Agustín Salinas", "AGUSTÍN SALINAS"},
	}

	for _, caso := range casos {
		t.Run(caso.nuevo, func(t *testing.T) {
			registry := NewProjectRegistry()
			project, err := registry.Create(ProjectData{Name: "Sistema de metricas", Start: date(2026, 10, 1)})
			if err != nil {
				t.Fatalf("se esperaba registrar el proyecto de prueba sin error, salio: %v", err)
			}
			if _, err := project.AddMember(caso.cargado, RoleProductBuilder); err != nil {
				t.Fatalf("se esperaba cargar a %q sin error, salio: %v", caso.cargado, err)
			}

			_, err = project.AddMember(caso.nuevo, RoleProductBuilder)

			if !errors.Is(err, ErrMemberDuplicate) {
				t.Fatalf("se esperaba ErrMemberDuplicate al agregar %q con %q cargado, salio: %v", caso.nuevo, caso.cargado, err)
			}
			members := project.Members()
			if len(members) != 1 || members[0].Name() != caso.cargado {
				t.Fatalf("se esperaba que el proyecto siguiera teniendo solo a %q, salieron %d integrantes", caso.cargado, len(members))
			}
		})
	}
}

// Tarea 8: no son duplicados los nombres que difieren en una tilde o en los espacios del
// medio. La Fase 2 decide ignorar solo las mayusculas: "Agustin" y "Agustín" pueden ser dos
// personas distintas.
// Cubre: RN-5
func TestNoSonDuplicadosLosNombresQueDifierenEnTildesOEspaciosDelMedio(t *testing.T) {
	casos := []struct {
		cargado string
		nuevo   string
	}{
		{"Agustín Salinas", "Agustin Salinas"},
		{"Juliana Bustos", "Juliana  Bustos"},
	}

	for _, caso := range casos {
		t.Run(caso.nuevo, func(t *testing.T) {
			registry := NewProjectRegistry()
			project, err := registry.Create(ProjectData{Name: "Sistema de metricas", Start: date(2026, 10, 1)})
			if err != nil {
				t.Fatalf("se esperaba registrar el proyecto de prueba sin error, salio: %v", err)
			}
			if _, err := project.AddMember(caso.cargado, RoleProductBuilder); err != nil {
				t.Fatalf("se esperaba cargar a %q sin error, salio: %v", caso.cargado, err)
			}

			_, err = project.AddMember(caso.nuevo, RoleProductBuilder)

			if err != nil {
				t.Fatalf("se esperaba que %q no fuera duplicado de %q, salio: %v", caso.nuevo, caso.cargado, err)
			}
			if len(project.Members()) != 2 {
				t.Fatalf("se esperaban 2 integrantes, salieron %d", len(project.Members()))
			}
		})
	}
}

// Escenario: La misma persona en dos proyectos
// Cubre: RN-6
func TestLaMismaPersonaEnDosProyectos(t *testing.T) {
	registry := NewProjectRegistry()
	proyectoA, err := registry.Create(ProjectData{Name: "Proyecto A", Start: date(2026, 10, 1)})
	if err != nil {
		t.Fatalf("se esperaba registrar el proyecto A sin error, salio: %v", err)
	}
	proyectoB, err := registry.Create(ProjectData{Name: "Proyecto B", Start: date(2026, 10, 1)})
	if err != nil {
		t.Fatalf("se esperaba registrar el proyecto B sin error, salio: %v", err)
	}
	if _, err := proyectoA.AddMember("Juliana Bustos", RoleProductBuilder); err != nil {
		t.Fatalf("se esperaba cargar a Juliana Bustos en el proyecto A sin error, salio: %v", err)
	}

	_, err = proyectoB.AddMember("Juliana Bustos", RoleProductBuilder)

	if err != nil {
		t.Fatalf("se esperaba aceptar a Juliana Bustos en el proyecto B, salio: %v", err)
	}
	for nombre, project := range map[string]*Project{"A": proyectoA, "B": proyectoB} {
		members := project.Members()
		if len(members) != 1 || members[0].Name() != "Juliana Bustos" {
			t.Fatalf("se esperaba que el proyecto %s incluyera solo a Juliana Bustos, salieron %d integrantes", nombre, len(members))
		}
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
