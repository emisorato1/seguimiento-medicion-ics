package domain

// ClosedSprints devuelve los sprints cerrados del proyecto, del mas viejo al mas nuevo
// (historia #11, RN-1). El orden sale del que tienen guardado: la #8 agrega cada sprint al
// final con el numero siguiente, asi que no hay que ordenar. La lista es nueva, pero los
// sprints son los guardados, no copias, igual que Sprints. Si no tiene ninguno devuelve
// una lista vacia, que no es un error (RN-3).
func (p *Project) ClosedSprints() []*Sprint {
	closed := []*Sprint{}
	for _, sprint := range p.sprints {
		if sprint.state == SprintClosed {
			closed = append(closed, sprint)
		}
	}
	return closed
}
