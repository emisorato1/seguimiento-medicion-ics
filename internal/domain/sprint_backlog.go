package domain

// AssignStory agrega la historia con ese identificador al Sprint Backlog del sprint abierto
// del proyecto y la deja En curso (historia #9, RN-1, RN-4). Se le pide al proyecto, porque
// hacen falta las dos puntas: el sprint abierto y la historia del Product Backlog.
func (p *Project) AssignStory(id int) error {
	sprint, ok := p.OpenSprint()
	if !ok {
		return ErrNoOpenSprint
	}
	story, err := p.Story(id)
	if err != nil {
		return err
	}
	// Los controles van antes de cambiar nada, asi un rechazo deja el Sprint Backlog y la
	// historia como estaban (RN-14), y de lo mas general a lo mas particular: no tiene
	// sentido mirar la estimacion de una historia que no se puede mover.
	if story.state == StoryDone {
		return ErrStoryDone
	}
	if !story.IsEstimated() {
		return ErrStoryNotEstimated
	}
	if sprint.hasStory(story.id) {
		return ErrStoryAlreadyInSprint
	}
	sprint.stories = append(sprint.stories, story.id)
	story.state = StoryInProgress
	return nil
}

// UnassignStory saca la historia con ese identificador del Sprint Backlog del sprint abierto
// y la devuelve al Product Backlog como Pendiente (historia #9, RN-7). Una historia Terminada
// no se mueve (RN-6).
func (p *Project) UnassignStory(id int) error {
	sprint, ok := p.OpenSprint()
	if !ok {
		return ErrNoOpenSprint
	}
	story, err := p.Story(id)
	if err != nil {
		return err
	}
	if story.state == StoryDone {
		return ErrStoryDone
	}
	if !sprint.hasStory(story.id) {
		return ErrStoryNotInSprint
	}
	// Se arma la lista nueva sin la historia, en el mismo orden, y recien despues se aplica
	// junto con el cambio de estado.
	remaining := make([]int, 0, len(sprint.stories))
	for _, storyID := range sprint.stories {
		if storyID != story.id {
			remaining = append(remaining, storyID)
		}
	}
	sprint.stories = remaining
	story.state = StoryPending
	return nil
}

// hasStory indica si la historia con ese identificador esta en el Sprint Backlog.
func (s *Sprint) hasStory(id int) bool {
	for _, storyID := range s.stories {
		if storyID == id {
			return true
		}
	}
	return false
}
