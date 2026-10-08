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
// Normaliza y valida todo antes de registrar: si algo falla, no se registra nada.
func (r *ProjectRegistry) Create(data ProjectData) (*Project, error) {
	data, err := prepare(data)
	if err != nil {
		return nil, err
	}

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

// Get devuelve el proyecto guardado con ese identificador, no una copia.
func (r *ProjectRegistry) Get(id int) (*Project, error) {
	project, ok := r.projects[id]
	if !ok {
		return nil, ErrProjectNotFound
	}
	return project, nil
}

// Update modifica el proyecto con ese identificador segun el pedido de cambios.
// Arma el resultado aparte, lo normaliza y lo valida completo, y recien ahi lo aplica:
// si algo falla, el proyecto queda como estaba (RN-5).
func (r *ProjectRegistry) Update(id int, changes ProjectChanges) error {
	project, err := r.Get(id)
	if err != nil {
		return err
	}

	result := ProjectData{
		Name:        project.name,
		Description: project.description,
		Start:       project.start,
		End:         project.end,
	}
	if changes.Name != nil {
		result.Name = *changes.Name
	}
	if changes.Start != nil {
		result.Start = *changes.Start
	}

	result, err = prepare(result)
	if err != nil {
		return err
	}

	project.name = result.Name
	project.start = result.Start
	return nil
}
