package domain

// StoryState es el estado de una historia dentro de su proyecto. Es un tipo propio y no un
// texto libre, para que no se pueda guardar un estado que no existe. Las transiciones entre
// estados son de las historias #9 y #10; esta historia solo crea en Pendiente.
type StoryState int

const (
	// StoryPending: esta en el backlog, sin comprometer.
	StoryPending StoryState = iota
	// StoryInProgress: fue asignada a un sprint abierto.
	StoryInProgress
	// StoryDone: se completo y el sprint se cerro.
	StoryDone
)

// String devuelve el nombre del estado tal como lo nombra la Fase 1.
func (s StoryState) String() string {
	switch s {
	case StoryInProgress:
		return "En curso"
	case StoryDone:
		return "Terminada"
	default:
		return "Pendiente"
	}
}

// Priority es la prioridad de una historia. Igual que el estado, es un tipo propio: la
// RN-4 solo acepta Alta, Media y Baja.
type Priority int

const (
	// PriorityHigh, PriorityMedium y PriorityLow son las tres unicas prioridades validas.
	PriorityHigh Priority = iota
	PriorityMedium
	PriorityLow
)

// String devuelve el nombre de la prioridad tal como la nombra la Fase 1.
func (p Priority) String() string {
	switch p {
	case PriorityMedium:
		return "Media"
	case PriorityLow:
		return "Baja"
	default:
		return "Alta"
	}
}

// Story es una historia del Product Backlog de un proyecto. Sus datos estan ocultos fuera
// del paquete: solo se cambian con operaciones que validan.
type Story struct {
	id                 int
	title              string
	description        string
	priority           Priority
	state              StoryState
	storyPoints        int
	acceptanceCriteria []string
}

// ID devuelve el identificador de la historia, unico dentro de su proyecto (RN-16).
func (s *Story) ID() int {
	return s.id
}

// Title devuelve el titulo de la historia.
func (s *Story) Title() string {
	return s.title
}

// Description devuelve la descripcion de la historia.
func (s *Story) Description() string {
	return s.description
}

// Priority devuelve la prioridad de la historia.
func (s *Story) Priority() Priority {
	return s.priority
}

// State devuelve el estado de la historia.
func (s *Story) State() StoryState {
	return s.state
}

// StoryPoints devuelve la estimacion de la historia. El valor 0 significa que no esta
// estimada: quien lo muestra informa "sin estimar" (historia #12).
func (s *Story) StoryPoints() int {
	return s.storyPoints
}

// IsEstimated indica si la historia tiene una estimacion cargada (historia #12).
func (s *Story) IsEstimated() bool {
	return s.storyPoints > 0
}

// AcceptanceCriteria devuelve una copia de los criterios de aceptacion: cambiar lo que se
// recibe no cambia los de la historia.
func (s *Story) AcceptanceCriteria() []string {
	criterios := make([]string, len(s.acceptanceCriteria))
	copy(criterios, s.acceptanceCriteria)
	return criterios
}
