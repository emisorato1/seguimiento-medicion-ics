// Package domain es el nucleo de la aplicacion: las entidades y sus reglas de negocio.
package domain

import "time"

// ProjectData son los datos con los que se crea un proyecto.
// Una fecha en cero significa que no se indico.
type ProjectData struct {
	Name        string
	Description string
	Start       time.Time
	End         time.Time
}

// ProjectChanges es el pedido de modificacion de un proyecto. Un dato en nil no se toca;
// uno con valor cambia a ese valor. Un dato que no se indica nunca se borra (RN-8).
type ProjectChanges struct {
	Name *string
}

// Project es un proyecto de software registrado. Sus datos estan ocultos fuera del
// paquete: solo se cambian con operaciones que validan.
type Project struct {
	id          int
	name        string
	description string
	start       time.Time
	end         time.Time
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
