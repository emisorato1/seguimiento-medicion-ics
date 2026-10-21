package domain

import "sort"

// PrioritizedBacklog devuelve el Product Backlog del proyecto ordenado por prioridad
// (historia #7). Es una consulta: arma una lista nueva y no cambia el backlog guardado
// (RN-7). Un proyecto sin historias devuelve una lista vacia, y eso no es un error (RN-8).
// La lista es nueva, pero las historias son las guardadas, con todos sus datos (RN-4,
// RN-5). El orden se arma en cada pedido, asi una prioridad que cambio aparece en su nuevo
// lugar sin trabajo extra (RN-6).
func (p *Project) PrioritizedBacklog() []*Story {
	stories := p.Backlog()
	// La importancia sale del orden en que la #6 declara las prioridades: Alta, Media, Baja
	// (RN-1, RN-2). SliceStable deja en el orden en que vienen las de igual prioridad.
	sort.SliceStable(stories, func(i, j int) bool {
		return stories[i].priority < stories[j].priority
	})
	return stories
}
