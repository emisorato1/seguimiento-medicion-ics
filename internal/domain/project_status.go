package domain

import "time"

// ProjectStatus es el estado de un proyecto en el momento de la consulta (historia #5): una
// foto con copias de los datos, no las entidades guardadas. Cambiarla no cambia el proyecto
// (RN-10), y no cambia si el proyecto cambia despues: para ver un cambio se vuelve a pedir
// (RN-12).
type ProjectStatus struct {
	ID          int
	Name        string
	Description string
	// HasDescription y HasEnd dicen si esos datos opcionales estan cargados (RN-2).
	HasDescription bool
	Start          time.Time
	End            time.Time
	HasEnd         bool
	// Members son los integrantes, cada uno con su nombre y su rol, en el orden en que se
	// cargaron (RN-3). Sin integrantes es una lista vacia (RN-4).
	Members []MemberStatus
	Backlog BacklogSummary
	// CurrentSprint es el sprint abierto, si HasCurrentSprint lo dice. Que no haya ninguno
	// no es un error (RN-9).
	CurrentSprint    SprintSummary
	HasCurrentSprint bool
}

// MemberStatus es un integrante en el estado del proyecto: su nombre y apellido y su rol.
type MemberStatus struct {
	Name string
	Role Role
}

// BacklogSummary es el resumen del backlog: cuantas historias tiene, cuantas estan
// completadas y cuantas pendientes (RN-5).
type BacklogSummary struct {
	Total     int
	Completed int
	Pending   int
}

// SprintSummary es el sprint en curso en el estado del proyecto: su numero, su Sprint Goal,
// sus fechas y cuantas historias tiene asignadas (RN-8).
type SprintSummary struct {
	Number          int
	Goal            string
	Start           time.Time
	End             time.Time
	AssignedStories int
}

// Status arma el estado del proyecto con lo que tiene en este momento, sin cambiar nada
// (RN-10). El proyecto que no existe lo informa el registro al buscarlo (RN-11).
func (p *Project) Status() ProjectStatus {
	members := make([]MemberStatus, 0, len(p.members))
	for _, member := range p.members {
		members = append(members, MemberStatus{Name: member.name, Role: member.role})
	}

	// El resumen se cuenta en cada consulta, mirando el estado de cada historia: solo las
	// Terminadas son completadas, y las Pendientes y En curso son pendientes (RN-6). Asi el
	// total es siempre la suma de las dos (RN-7).
	var backlog BacklogSummary
	for _, story := range p.backlog {
		backlog.Total++
		if story.state == StoryDone {
			backlog.Completed++
		} else {
			backlog.Pending++
		}
	}

	return ProjectStatus{
		ID:             p.id,
		Name:           p.name,
		Description:    p.description,
		HasDescription: p.description != "",
		Start:          p.start,
		End:            p.end,
		HasEnd:         p.HasEnd(),
		Members:        members,
		Backlog:        backlog,
	}
}
