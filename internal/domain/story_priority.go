package domain

// PrioritizedBacklog devuelve el Product Backlog del proyecto ordenado por prioridad
// (historia #7). Es una consulta: arma una lista nueva y no cambia el backlog guardado. Un
// proyecto sin historias devuelve una lista vacia, y eso no es un error (RN-8). La lista es
// nueva, pero las historias son las guardadas, con todos sus datos (RN-4, RN-5).
func (p *Project) PrioritizedBacklog() []*Story {
	return p.Backlog()
}
