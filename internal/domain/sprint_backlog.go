package domain

// AssignStory agrega la historia con ese identificador al Sprint Backlog del sprint abierto
// del proyecto y la deja En curso (historia #9, RN-1, RN-4). Se le pide al proyecto, porque
// hacen falta las dos puntas: el sprint abierto y la historia del Product Backlog.
func (p *Project) AssignStory(id int) error {
	sprint, _ := p.OpenSprint()
	story, err := p.Story(id)
	if err != nil {
		return err
	}
	sprint.stories = append(sprint.stories, story.id)
	story.state = StoryInProgress
	return nil
}
