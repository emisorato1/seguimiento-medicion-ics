package domain

import "math"

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
	// "no es mayor que cero" y no "menor o igual que cero": las dos dicen lo mismo salvo con
	// NaN, que no es mayor ni menor que nada y con la segunda pasaría. El infinito sí es
	// mayor que cero, y se rechaza aparte porque no son horas y JSON no lo puede guardar
	// (pedido de cambios de Agustín en el PR #81, como hace la #14 con las horas reales).
	if !(hours > 0) || math.IsInf(hours, 0) {
		return ErrEstimatedHoursInvalid
	}
	story.estimatedHours = hours
	return nil
}
