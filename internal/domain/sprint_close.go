package domain

import (
	"fmt"
	"time"
)

// CloseSprint cierra el sprint abierto del proyecto con la fecha de cierre indicada (historia
// #10, RN-1). El sprint que se cierra es siempre el abierto: por eso se le pide al proyecto, y
// cerrado el sprint, el proyecto queda sin sprint abierto y puede crear el siguiente (RN-14).
func (p *Project) CloseSprint(completed []int, closedOn time.Time) error {
	sprint, _ := p.OpenSprint()
	// la fecha de cierre es obligatoria, y no se compara con ninguna otra (RN-10, RN-11)
	if closedOn.IsZero() {
		return ErrSprintCloseDateRequired
	}
	// las completadas quedan registradas aparte de las asignadas (RN-2); las asignadas quedan
	// Terminadas si se completaron y vuelven al Product Backlog como Pendientes si no (RN-3,
	// RN-4)
	// se revisa la lista entera antes de tocar una sola historia, así un cierre rechazado deja
	// todo como estaba (RN-7, RN-13); una historia indicada dos veces se registra una sola
	// vez, en el orden en que apareció primero (RN-8)
	done := map[int]bool{}
	unique := []int{}
	for _, id := range completed {
		if !sprint.hasStory(id) {
			return fmt.Errorf("%w: la historia %d", ErrStoryNotInSprint, id)
		}
		if !done[id] {
			unique = append(unique, id)
		}
		done[id] = true
	}
	for _, id := range sprint.stories {
		story, _ := p.Story(id)
		if done[id] {
			story.state = StoryDone
		} else {
			story.state = StoryPending
		}
	}
	sprint.completedStories = unique
	sprint.state = SprintClosed
	sprint.closedOn = toDay(closedOn)
	return nil
}

// CompletedStories devuelve una copia de los identificadores de las historias que se
// registraron como completadas al cerrar el sprint, vacía mientras está abierto. Cambiar lo
// que se recibe no cambia el sprint, como en Stories().
func (s *Sprint) CompletedStories() []int {
	completed := make([]int, len(s.completedStories))
	copy(completed, s.completedStories)
	return completed
}

// ClosedOn devuelve el día en que el sprint se cerró de verdad, o la fecha cero si sigue
// abierto.
func (s *Sprint) ClosedOn() time.Time {
	return s.closedOn
}
