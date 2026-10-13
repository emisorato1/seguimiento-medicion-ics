package domain

import "strings"

// Role es el rol de un integrante dentro de su proyecto. Es un tipo propio y no un texto
// libre: la RN-4 solo acepta Agile Enabler y Product Builder.
//
// El valor cero no es ningun rol: significa "sin rol". Asi un integrante armado sin indicar
// el rol no queda como Agile Enabler sin que nadie se entere, y AddMember lo rechaza.
type Role int

const (
	// RoleAgileEnabler y RoleProductBuilder son los dos unicos roles validos. Arrancan en 1
	// para que el valor cero quede libre como "sin rol".
	RoleAgileEnabler Role = iota + 1
	RoleProductBuilder
)

// String devuelve el nombre del rol tal como lo nombra la Fase 1, o "sin rol" si no es
// ninguno de los dos.
func (r Role) String() string {
	switch r {
	case RoleAgileEnabler:
		return "Agile Enabler"
	case RoleProductBuilder:
		return "Product Builder"
	default:
		return "sin rol"
	}
}

// valid indica si el rol es uno de los dos que acepta la RN-4.
func (r Role) valid() bool {
	return r == RoleAgileEnabler || r == RoleProductBuilder
}

// Member es un integrante de un proyecto. Se identifica por su nombre y apellido, que es
// un solo dato y es unico dentro del proyecto: no tiene un identificador numerico propio
// (Fase 2). Sus datos estan ocultos fuera del paquete: solo se cambian con operaciones que
// validan.
type Member struct {
	name string
	role Role
}

// Name devuelve el nombre y apellido del integrante, tal como se escribio y sin los espacios
// de los extremos.
func (m *Member) Name() string {
	return m.name
}

// Role devuelve el rol del integrante dentro del proyecto.
func (m *Member) Role() Role {
	return m.role
}

// Members devuelve los integrantes del proyecto. Si no tiene ninguno devuelve una lista
// vacia, que no es un error (RN-8). Esta en member.go y no en project.go para que la #4 y
// la #6 choquen lo menos posible en ese archivo (Fase 2).
func (p *Project) Members() []*Member {
	members := make([]*Member, len(p.members))
	copy(members, p.members)
	return members
}

// AddMember agrega un integrante al proyecto con su nombre y apellido y su rol (RN-1), y
// queda en la lista de integrantes (RN-8). El rol llega ya convertido: el texto que escribe
// el usuario se convierte antes con ParseRole. Primero normaliza el nombre, sacando los
// espacios de los extremos (RN-3); las validaciones corren despues, sobre lo ya normalizado,
// y todas antes de tocar la lista: si alguna falla, el proyecto queda como estaba (RN-12).
func (p *Project) AddMember(name string, role Role) (*Member, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, ErrMemberNameRequired
	}
	if !role.valid() {
		return nil, ErrMemberRoleInvalid
	}

	member := &Member{name: name, role: role}
	p.members = append(p.members, member)
	return member, nil
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
		return 0, ErrMemberRoleInvalid
	}
}
