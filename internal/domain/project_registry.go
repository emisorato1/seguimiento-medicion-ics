package domain

import "strings"

// ProjectRegistry guarda los proyectos mientras corre el programa y es el unico que
// les asigna identificador.
type ProjectRegistry struct {
	projects map[int]*Project
	nextID   int
}

// NewProjectRegistry crea un registro vacio. El primer proyecto recibe el identificador 1.
func NewProjectRegistry() *ProjectRegistry {
	return &ProjectRegistry{
		projects: make(map[int]*Project),
		nextID:   1,
	}
}

// Create registra un proyecto nuevo con los datos indicados y le asigna su identificador.
// Normaliza y valida todo antes de registrar: si algo falla, no se registra nada.
func (r *ProjectRegistry) Create(data ProjectData) (*Project, error) {
	name := strings.TrimSpace(data.Name)
	if name == "" {
		return nil, ErrProjectNameRequired
	}
	if data.Start.IsZero() {
		return nil, ErrProjectStartRequired
	}
	if !data.End.IsZero() && data.End.Before(data.Start) {
		return nil, ErrProjectDatesInconsistent
	}

	project := &Project{
		id:          r.nextID,
		name:        name,
		description: data.Description,
		start:       data.Start,
		end:         data.End,
	}
	r.projects[project.id] = project
	r.nextID++
	return project, nil
}

// Get devuelve el proyecto guardado con ese identificador, no una copia.
func (r *ProjectRegistry) Get(id int) (*Project, error) {
	project, ok := r.projects[id]
	if !ok {
		return nil, ErrProjectNotFound
	}
	return project, nil
}

// Update modifica el proyecto con ese identificador segun el pedido de cambios.
func (r *ProjectRegistry) Update(id int, changes ProjectChanges) error {
	project, err := r.Get(id)
	if err != nil {
		return err
	}
	if changes.Name != nil {
		project.name = *changes.Name
	}
	return nil
}
