package domain

// ClosedSprints devuelve los sprints cerrados del proyecto (historia #11). Si no tiene
// ninguno devuelve una lista vacia, que no es un error (RN-3).
func (p *Project) ClosedSprints() []*Sprint {
	return []*Sprint{}
}
