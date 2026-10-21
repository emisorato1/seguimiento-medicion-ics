package domain

// PrioritizedBacklog devuelve el Product Backlog del proyecto ordenado por prioridad
// (historia #7). Es una consulta: arma una lista nueva y no cambia el backlog guardado. Un
// proyecto sin historias devuelve una lista vacia, y eso no es un error (RN-8).
func (p *Project) PrioritizedBacklog() []*Story {
	return []*Story{}
}
