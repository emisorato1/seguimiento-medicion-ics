package domain

// Estimate asigna story points a la historia. Sirve tanto para la primera estimación
// como para reestimar (RN-5, RN-6).
func (s *Story) Estimate(points int) error {
	s.storyPoints = points
	return nil
}
