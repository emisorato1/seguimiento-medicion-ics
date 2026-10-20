package domain

// EstimateHours carga las horas estimadas de la historia con ese identificador dentro del
// proyecto. Admite fracciones, y volver a cargarlas reemplaza el valor anterior (historia
// #16, RN-17). Sigue el mismo orden que la estimación en story points de la #12: busca la
// historia, controla que no esté Terminada y valida el valor, y recién ahí lo aplica. Si
// algo falla, la historia queda como estaba (RN-18, RN-19).
func (p *Project) EstimateHours(id int, hours float64) error {
	story, err := p.Story(id)
	if err != nil {
		return err
	}
	if story.state == StoryDone {
		return ErrStoryEstimateDone
	}
	if hours <= 0 {
		return ErrEstimatedHoursInvalid
	}
	story.estimatedHours = hours
	return nil
}
