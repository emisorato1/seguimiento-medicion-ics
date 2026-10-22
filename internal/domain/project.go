// Package domain es el nucleo de la aplicacion: las entidades y sus reglas de negocio.
package domain

import (
	"strings"
	"time"
)

// ProjectData son los datos con los que se crea un proyecto.
// Una fecha en cero significa que no se indico.
type ProjectData struct {
	Name        string
	Description string
	Start       time.Time
	End         time.Time
}

// ProjectChanges es el pedido de modificacion de un proyecto. Un dato que no se indica
// nunca se borra (RN-8). El nombre y la fecha de inicio son obligatorios: en nil no se
// tocan y con valor cambian, pero no existe la opcion de quitarlos. La descripcion y la
// fecha de finalizacion son opcionales y se pueden quitar con un pedido explicito (RN-9).
type ProjectChanges struct {
	Name        *string
	Start       *time.Time
	Description OptionalChange[string]
	End         OptionalChange[time.Time]
}

// OptionalChange es lo que se pide sobre un dato opcional al modificar un proyecto: su
// valor cero no toca el dato, SetTo lo cambia y Remove lo quita. Como es un solo valor,
// un pedido no puede decir dos cosas a la vez sobre el mismo dato.
type OptionalChange[T any] struct {
	set    bool
	value  T
	remove bool
}

// SetTo pide cambiar un dato opcional por un valor nuevo.
func SetTo[T any](value T) OptionalChange[T] {
	return OptionalChange[T]{set: true, value: value}
}

// Remove pide quitar un dato opcional.
func Remove[T any]() OptionalChange[T] {
	return OptionalChange[T]{remove: true}
}

// apply devuelve el dato despues de aplicar el pedido: el mismo, el valor nuevo o vacio.
func (c OptionalChange[T]) apply(current T) T {
	switch {
	case c.set:
		return c.value
	case c.remove:
		var empty T
		return empty
	default:
		return current
	}
}

// prepare normaliza y valida los datos de un proyecto, en ese orden, sin tocar ningun
// proyecto. Crear y modificar usan esta misma funcion, asi aplican las mismas reglas (RN-5).
func prepare(data ProjectData) (ProjectData, error) {
	data.Name = strings.TrimSpace(data.Name)
	data.Start = toDay(data.Start)
	data.End = toDay(data.End)
	if data.Name == "" {
		return ProjectData{}, ErrProjectNameRequired
	}
	if data.Start.IsZero() {
		return ProjectData{}, ErrProjectStartRequired
	}
	if !data.End.IsZero() && data.End.Before(data.Start) {
		return ProjectData{}, ErrProjectDatesInconsistent
	}
	return data, nil
}

// toDay deja solo el dia de una fecha, sin la hora ni la zona horaria: las fechas de un
// proyecto son dias y se comparan como tales (defecto #59). El dia se toma tal como lo
// ve quien lo indico y se arma siempre en UTC, asi dos fechas del mismo dia son iguales
// aunque vengan en zonas distintas. Una fecha en cero sigue en cero.
func toDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// Project es un proyecto de software registrado. Sus datos estan ocultos fuera del
// paquete: solo se cambian con operaciones que validan.
type Project struct {
	id          int
	name        string
	description string
	start       time.Time
	end         time.Time

	// backlog y nextStoryID son el Product Backlog del proyecto (historia #6). Las
	// historias viven dentro del proyecto: no existen sin el, y asi el identificador
	// unico por proyecto (RN-16) sale solo, sin llevar un indice aparte.
	backlog     []*Story
	nextStoryID int
	// members son los integrantes del proyecto (historia #4). Viven dentro del proyecto: no
	// existen sin el, y asi el mismo nombre en otro proyecto nunca cuenta como duplicado
	// (RN-6). Las operaciones sobre integrantes estan en member.go.
	members []*Member
	// sprints y lastSprintNumber son los sprints del proyecto (historia #8). Viven dentro
	// del proyecto, asi el numero unico por proyecto (RN-1, RN-9) sale solo. El contador
	// guarda el ultimo numero asignado y nunca retrocede: en cero, el proximo sprint es el 1,
	// tambien en un proyecto cargado de un archivo que todavia no guardaba sprints. Las
	// operaciones sobre sprints estan en sprint.go.
	sprints          []*Sprint
	lastSprintNumber int
	// efforts son los registros de esfuerzo del proyecto, en el orden en que se cargaron
	// (historia #14). Viven dentro del proyecto, así las horas de un proyecto no suman en
	// otro (RN-16). Las operaciones están en effort.go.
	efforts []*Effort
}

// ID devuelve el identificador del proyecto, que no cambia nunca.
func (p *Project) ID() int {
	return p.id
}

// Name devuelve el nombre del proyecto.
func (p *Project) Name() string {
	return p.name
}

// Description devuelve la descripcion del proyecto, o vacio si no tiene.
func (p *Project) Description() string {
	return p.description
}

// Start devuelve la fecha de inicio del proyecto.
func (p *Project) Start() time.Time {
	return p.start
}

// End devuelve la fecha de finalizacion del proyecto, o una fecha en cero si no tiene.
func (p *Project) End() time.Time {
	return p.end
}

// HasEnd indica si el proyecto tiene fecha de finalizacion, que es opcional.
func (p *Project) HasEnd() bool {
	return !p.end.IsZero()
}
