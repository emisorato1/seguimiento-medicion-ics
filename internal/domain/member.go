package domain

import "strings"

// Role es el rol de un integrante dentro de su proyecto. Es un tipo propio y no un texto
// libre: la RN-4 solo acepta Agile Enabler y Product Builder.
type Role int

const (
	// RoleAgileEnabler y RoleProductBuilder son los dos unicos roles validos.
	RoleAgileEnabler Role = iota
	RoleProductBuilder
)

// String devuelve el nombre del rol tal como lo nombra la Fase 1.
func (r Role) String() string {
	switch r {
	case RoleProductBuilder:
		return "Product Builder"
	default:
		return "Agile Enabler"
	}
}

// ParseRole convierte el texto que escribe el usuario en un rol, ignorando mayusculas y
// los espacios de los extremos. Cualquier otro texto se rechaza con el error de la RN-4.
// La conversion vive en el dominio y no en la consola: si la consola rechazara un texto
// invalido por su cuenta, la RN-4 quedaria duplicada fuera del dominio (ajuste de la Fase 3).
func ParseRole(text string) (Role, error) {
	switch strings.ToLower(strings.TrimSpace(text)) {
	case "agile enabler":
		return RoleAgileEnabler, nil
	case "product builder":
		return RoleProductBuilder, nil
	default:
		return RoleAgileEnabler, ErrMemberRoleInvalid
	}
}
