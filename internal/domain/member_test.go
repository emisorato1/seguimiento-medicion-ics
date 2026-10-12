package domain

import (
	"errors"
	"strings"
	"testing"
)

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
