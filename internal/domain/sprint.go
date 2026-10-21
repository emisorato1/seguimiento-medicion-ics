package domain

import "time"

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

// Number devuelve el numero del sprint dentro de su proyecto, que no cambia nunca.
func (s *Sprint) Number() int {
	return s.number
}

// Sprints devuelve todos los sprints del proyecto, en orden de numero. La lista es nueva,
// pero los sprints son los guardados, no copias. Sin sprints devuelve una lista vacia.
func (p *Project) Sprints() []*Sprint {
	sprints := make([]*Sprint, len(p.sprints))
	copy(sprints, p.sprints)
	return sprints
}

// OpenSprint devuelve el sprint abierto del proyecto, y si hay alguno. Que no haya ninguno no
// es un error (RN-10): se informa aparte, para que quien consulta no tenga que distinguir
// "no hay" de "fallo".
func (p *Project) OpenSprint() (*Sprint, bool) {
	return nil, false
}
