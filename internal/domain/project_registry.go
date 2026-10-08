package domain

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
func (r *ProjectRegistry) Create(data ProjectData) (*Project, error) {
	project := &Project{
		id:          r.nextID,
		name:        data.Name,
		description: data.Description,
		start:       data.Start,
		end:         data.End,
	}
	r.projects[project.id] = project
	r.nextID++
	return project, nil
}

// Get devuelve el proyecto guardado con ese identificador.
func (r *ProjectRegistry) Get(id int) (*Project, error) {
	return r.projects[id], nil
}
