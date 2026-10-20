package domain

// EstimateHours carga las horas estimadas de la historia con ese identificador dentro del
// proyecto. Admite fracciones, y volver a cargarlas reemplaza el valor anterior (historia
// #16, RN-17). Primero busca la historia: si no existe, informa que no se encontró.
func (p *Project) EstimateHours(id int, hours float64) error {
	story, err := p.Story(id)
	if err != nil {
		return err
	}
	story.estimatedHours = hours
	return nil
}
