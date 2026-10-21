package domain

import (
	"strings"
	"time"
)

// SprintState es la etapa en la que esta un sprint (historia #8). Es un tipo propio para que
// no se pueda guardar un estado que no existe.
type SprintState int

const (
	// SprintOpen: es el sprint en curso del proyecto. Un sprint nace abierto (RN-3).
	SprintOpen SprintState = iota
	// SprintClosed: ya termino. El pase a cerrado lo hace la historia #10.
	SprintClosed
)

// Sprint es un sprint de un proyecto, con los datos que fija el contrato del dominio. Sus
// datos estan ocultos fuera del paquete: solo se cambian con operaciones que validan. Esta
// historia usa el numero, el goal, las fechas y el estado; las historias asignadas las llena
// la #9, y lo del cierre la #10.
type Sprint struct {
	number           int
	goal             string
	start            time.Time
	end              time.Time
	state            SprintState
	stories          []int
	completedStories []int
	pointsAtClose    map[int]int
	closedOn         time.Time
}

// SprintData son los datos con los que se crea un sprint. El numero no es un dato: lo asigna
// el sistema (RN-1).
type SprintData struct {
	Goal  string
	Start time.Time
	End   time.Time
}

// prepareSprint normaliza y valida los datos de un sprint, sin tocar ningun sprint ni el
// proyecto: si algo falla, no se agrega nada (RN-12).
func prepareSprint(data SprintData) (SprintData, error) {
	// Primero se normaliza y despues se valida sobre lo normalizado: asi un goal de solo
	// espacios cuenta como vacio (RN-4, RN-7). Las fechas son dias, sin hora ni zona, igual
	// que las del proyecto (defecto #59).
	data.Goal = strings.TrimSpace(data.Goal)
	data.Start = toDay(data.Start)
	data.End = toDay(data.End)
	if data.Goal == "" {
		return SprintData{}, ErrSprintGoalRequired
	}
	if data.Start.IsZero() {
		return SprintData{}, ErrSprintStartRequired
	}
	if data.End.IsZero() {
		return SprintData{}, ErrSprintEndRequired
	}
	if data.End.Before(data.Start) {
		return SprintData{}, ErrSprintDatesInconsistent
	}
	return data, nil
}

// Number devuelve el numero del sprint dentro de su proyecto, que no cambia nunca.
func (s *Sprint) Number() int {
	return s.number
}

// Goal devuelve el Sprint Goal: la meta del sprint, en una frase.
func (s *Sprint) Goal() string {
	return s.goal
}

// Start devuelve el dia de inicio del sprint.
func (s *Sprint) Start() time.Time {
	return s.start
}

// End devuelve el dia de fin planificado del sprint.
func (s *Sprint) End() time.Time {
	return s.end
}

// State devuelve en que etapa esta el sprint.
func (s *Sprint) State() SprintState {
	return s.state
}

// Stories devuelve una copia de los identificadores de las historias asignadas al sprint:
// cambiar lo que se recibe no cambia el sprint.
func (s *Sprint) Stories() []int {
	stories := make([]int, len(s.stories))
	copy(stories, s.stories)
	return stories
}

// AddSprint crea un sprint en el proyecto con los datos indicados. Recibe el numero que
// sigue al del ultimo sprint creado en el proyecto, nace abierto y sin historias (RN-1,
// RN-3).
func (p *Project) AddSprint(data SprintData) (*Sprint, error) {
	data, err := prepareSprint(data)
	if err != nil {
		return nil, err
	}

	sprint := &Sprint{
		number: p.lastSprintNumber + 1,
		goal:   data.Goal,
		start:  data.Start,
		end:    data.End,
		state:  SprintOpen,
	}
	p.sprints = append(p.sprints, sprint)
	p.lastSprintNumber = sprint.number
	return sprint, nil
}

// Sprints devuelve todos los sprints del proyecto, en orden de numero. La lista es nueva,
// pero los sprints son los guardados, no copias. Sin sprints devuelve una lista vacia.
func (p *Project) Sprints() []*Sprint {
	sprints := make([]*Sprint, len(p.sprints))
	copy(sprints, p.sprints)
	return sprints
}

// Sprint devuelve el sprint del proyecto con ese numero, abierto o cerrado, el guardado y no
// una copia. Si el proyecto no tiene ningun sprint con ese numero, informa que no se
// encontro (RN-15).
func (p *Project) Sprint(number int) (*Sprint, error) {
	for _, sprint := range p.sprints {
		if sprint.number == number {
			return sprint, nil
		}
	}
	return nil, ErrSprintNotFound
}

// OpenSprint devuelve el sprint abierto del proyecto, y si hay alguno. Que no haya ninguno no
// es un error (RN-10): se informa aparte, para que quien consulta no tenga que distinguir
// "no hay" de "fallo".
func (p *Project) OpenSprint() (*Sprint, bool) {
	for _, sprint := range p.sprints {
		if sprint.state == SprintOpen {
			return sprint, true
		}
	}
	return nil, false
}
