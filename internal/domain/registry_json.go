package domain

import (
	"encoding/json"
	"time"
)

// Conversion del registro completo a JSON y de vuelta (historia #77, D13 a D16). Vive en el
// dominio porque los datos de las entidades estan ocultos fuera del paquete: afuera no se
// podrian leer ni reconstruir sin exportarlos (D14). El dominio no toca archivos: convierte a
// texto y desde texto, y leer o escribir el archivo es de internal/store.
//
// Toda historia que agrega un dato a una entidad lo suma tambien aca, con su caso en el test
// de ida y vuelta: si no, se pierde al cerrar sin que ningun test lo note.

// dayLayout es como se guardan las fechas: solo el dia, sin hora ni zona (D16, defecto #59).
const dayLayout = "2006-01-02"

// Los tipos que siguen son la forma del archivo. No se exportan: solo los usa la conversion.
// El rol, la prioridad y el estado van con su nombre y no con su numero, asi el archivo se
// lee y no se rompe si se reordenan las constantes (D16).

type registryJSON struct {
	NextID   int           `json:"next_id"`
	Projects []projectJSON `json:"projects"`
}

type projectJSON struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description,omitempty"`
	Start       string       `json:"start"`
	End         string       `json:"end,omitempty"`
	NextStoryID int          `json:"next_story_id"`
	Members     []memberJSON `json:"members"`
	Stories     []storyJSON  `json:"stories"`
}

type memberJSON struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type storyJSON struct {
	ID                 int      `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Priority           string   `json:"priority"`
	State              string   `json:"state"`
	StoryPoints        int      `json:"story_points"`
	AcceptanceCriteria []string `json:"acceptance_criteria"`
}

// ToJSON convierte el registro completo, con sus proyectos, integrantes, historias y
// contadores, en el texto que se guarda.
func (r *ProjectRegistry) ToJSON() ([]byte, error) {
	out := registryJSON{NextID: r.nextID, Projects: []projectJSON{}}
	for _, p := range r.Projects() {
		project := projectJSON{
			ID:          p.id,
			Name:        p.name,
			Description: p.description,
			Start:       p.start.Format(dayLayout),
			NextStoryID: p.nextStoryID,
			Members:     []memberJSON{},
			Stories:     []storyJSON{},
		}
		if p.HasEnd() {
			project.End = p.end.Format(dayLayout)
		}
		for _, m := range p.members {
			project.Members = append(project.Members, memberJSON{Name: m.name, Role: m.role.String()})
		}
		for _, s := range p.backlog {
			project.Stories = append(project.Stories, storyJSON{
				ID:                 s.id,
				Title:              s.title,
				Description:        s.description,
				Priority:           s.priority.String(),
				State:              s.state.String(),
				StoryPoints:        s.storyPoints,
				AcceptanceCriteria: s.acceptanceCriteria,
			})
		}
		out.Projects = append(out.Projects, project)
	}
	return json.MarshalIndent(out, "", "  ")
}

// RegistryFromJSON reconstruye un registro a partir del texto guardado, con los mismos
// numeros y contadores que tenia.
func RegistryFromJSON(data []byte) (*ProjectRegistry, error) {
	var in registryJSON
	if err := json.Unmarshal(data, &in); err != nil {
		return nil, err
	}

	registry := NewProjectRegistry()
	for _, pj := range in.Projects {
		project, err := projectFromJSON(pj)
		if err != nil {
			return nil, err
		}
		registry.projects[project.id] = project
	}
	registry.nextID = in.NextID
	return registry, nil
}

// projectFromJSON reconstruye un proyecto con sus integrantes y su backlog.
func projectFromJSON(pj projectJSON) (*Project, error) {
	start, err := parseDay(pj.Start)
	if err != nil {
		return nil, err
	}
	end, err := parseDay(pj.End)
	if err != nil {
		return nil, err
	}
	project := &Project{
		id:          pj.ID,
		name:        pj.Name,
		description: pj.Description,
		start:       start,
		end:         end,
		nextStoryID: pj.NextStoryID,
	}

	for _, mj := range pj.Members {
		role, err := ParseRole(mj.Role)
		if err != nil {
			return nil, err
		}
		project.members = append(project.members, &Member{name: mj.Name, role: role})
	}

	for _, sj := range pj.Stories {
		priority, err := ParsePriority(sj.Priority)
		if err != nil {
			return nil, err
		}
		state, err := parseStoryState(sj.State)
		if err != nil {
			return nil, err
		}
		project.backlog = append(project.backlog, &Story{
			id:                 sj.ID,
			title:              sj.Title,
			description:        sj.Description,
			priority:           priority,
			state:              state,
			storyPoints:        sj.StoryPoints,
			acceptanceCriteria: sj.AcceptanceCriteria,
		})
	}
	return project, nil
}

// parseDay lee una fecha guardada como dia. Un texto vacio es una fecha sin cargar.
func parseDay(text string) (time.Time, error) {
	if text == "" {
		return time.Time{}, nil
	}
	return time.Parse(dayLayout, text)
}

// parseStoryState lee el estado de una historia por su nombre.
func parseStoryState(text string) (StoryState, error) {
	for _, state := range []StoryState{StoryPending, StoryInProgress, StoryDone} {
		if text == state.String() {
			return state, nil
		}
	}
	return StoryPending, ErrSavedDataInvalid
}
