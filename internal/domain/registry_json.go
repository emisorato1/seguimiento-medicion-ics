package domain

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	Efforts     []effortJSON `json:"efforts"`
	// Los sprints son de la #8. Van con omitempty: un archivo de antes, sin sprints, se lee
	// como un proyecto sin sprints y con el contador en cero.
	LastSprintNumber int          `json:"last_sprint_number,omitempty"`
	Sprints          []sprintJSON `json:"sprints,omitempty"`
}

type effortJSON struct {
	StoryID    int     `json:"story_id"`
	MemberName string  `json:"member_name"`
	Activity   string  `json:"activity"`
	Hours      float64 `json:"hours"`
	Day        string  `json:"day"`
}

// sprintJSON es un sprint guardado (historia #8). El estado va con su nombre, como el de las
// historias. Las historias asignadas y los datos del cierre los suman la #9 y la #10.
type sprintJSON struct {
	Number int    `json:"number"`
	Goal   string `json:"goal"`
	Start  string `json:"start"`
	End    string `json:"end"`
	State  string `json:"state"`
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
	EstimatedHours     float64  `json:"estimated_hours,omitempty"`
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
			Efforts:     []effortJSON{},
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
				EstimatedHours:     s.estimatedHours,
				AcceptanceCriteria: s.acceptanceCriteria,
			})
		}
		for _, e := range p.efforts {
			project.Efforts = append(project.Efforts, effortJSON{
				StoryID:    e.storyID,
				MemberName: e.memberName,
				Activity:   e.activity,
				Hours:      e.hours,
				Day:        e.day.Format(dayLayout),
			})
		}
		project.LastSprintNumber = p.lastSprintNumber
		for _, s := range p.sprints {
			project.Sprints = append(project.Sprints, sprintJSON{
				Number: s.number,
				Goal:   s.goal,
				Start:  s.start.Format(dayLayout),
				End:    s.end.Format(dayLayout),
				State:  s.state.String(),
			})
		}
		out.Projects = append(out.Projects, project)
	}
	return json.MarshalIndent(out, "", "  ")
}

// RegistryFromJSON reconstruye un registro a partir del texto guardado, con los mismos
// numeros y contadores que tenia. Cada entidad pasa por las mismas validaciones que al
// crearla, y un solo dato invalido rechaza la carga entera (D15): asi el resto del sistema
// puede seguir dando por ciertas sus reglas aunque el archivo se haya editado a mano. El
// error cumple errors.Is con ErrSavedDataInvalid y, si el motivo es una regla del dominio,
// tambien con el error de esa regla.
func RegistryFromJSON(data []byte) (*ProjectRegistry, error) {
	registry, err := registryFromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrSavedDataInvalid, err)
	}
	return registry, nil
}

// registryFromJSON hace la reconstruccion; RegistryFromJSON solo le agrega al error que son
// los datos guardados los que no se pueden leer.
func registryFromJSON(data []byte) (*ProjectRegistry, error) {
	// Un dato que el sistema no conoce se rechaza en vez de ignorarse: si un programa viejo
	// abriera un archivo con datos de una historia posterior, los perderia en silencio al
	// volver a guardar.
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	var in registryJSON
	if err := decoder.Decode(&in); err != nil {
		// el detalle de encoding/json viene en ingles y el mensaje es para el usuario
		// (AGENTS.md): no se arrastra
		return nil, errors.New("el archivo está dañado, incompleto o tiene datos que el sistema no conoce")
	}
	// Decode lee solo el primer valor: si despues viene algo mas que espacios, el archivo
	// no es lo que el sistema escribio y tampoco se acepta.
	if _, err := decoder.Token(); err != io.EOF {
		return nil, fmt.Errorf("tienen texto de más después de los datos")
	}

	ids := make([]int, 0, len(in.Projects))
	for _, pj := range in.Projects {
		ids = append(ids, pj.ID)
	}
	if err := checkNumbers(ids, in.NextID); err != nil {
		return nil, fmt.Errorf("proyectos: %w", err)
	}

	registry := NewProjectRegistry()
	for _, pj := range in.Projects {
		project, err := projectFromJSON(pj)
		if err != nil {
			return nil, fmt.Errorf("proyecto %d: %w", pj.ID, err)
		}
		registry.projects[project.id] = project
	}
	registry.nextID = in.NextID
	return registry, nil
}

// projectFromJSON reconstruye un proyecto con sus integrantes y su backlog, validando cada
// uno como al crearlo.
func projectFromJSON(pj projectJSON) (*Project, error) {
	start, err := parseDay(pj.Start)
	if err != nil {
		return nil, err
	}
	end, err := parseDay(pj.End)
	if err != nil {
		return nil, err
	}
	data, err := prepare(ProjectData{Name: pj.Name, Description: pj.Description, Start: start, End: end})
	if err != nil {
		return nil, err
	}
	project := &Project{
		id:          pj.ID,
		name:        data.Name,
		description: data.Description,
		start:       data.Start,
		end:         data.End,
	}

	// Los integrantes entran con la misma operacion que al cargarlos, que controla el nombre,
	// el duplicado y que haya un solo Agile Enabler.
	for _, mj := range pj.Members {
		role, err := ParseRole(mj.Role)
		if err != nil {
			return nil, fmt.Errorf("integrante %q: %w", mj.Name, err)
		}
		if _, err := project.AddMember(mj.Name, role); err != nil {
			return nil, fmt.Errorf("integrante %q: %w", mj.Name, err)
		}
	}

	// Las historias no pueden entrar con AddStory, que les daria un numero nuevo: se validan
	// con lo mismo que usa AddStory y se conserva su numero.
	ids := make([]int, 0, len(pj.Stories))
	for _, sj := range pj.Stories {
		ids = append(ids, sj.ID)
	}
	if err := checkNumbers(ids, pj.NextStoryID); err != nil {
		return nil, fmt.Errorf("historias: %w", err)
	}
	for _, sj := range pj.Stories {
		story, err := storyFromJSON(sj)
		if err != nil {
			return nil, fmt.Errorf("historia %d: %w", sj.ID, err)
		}
		project.backlog = append(project.backlog, story)
	}
	project.nextStoryID = pj.NextStoryID

	sprints, err := sprintsFromJSON(pj.Sprints, pj.LastSprintNumber)
	if err != nil {
		return nil, fmt.Errorf("sprints: %w", err)
	}
	project.sprints = sprints
	project.lastSprintNumber = pj.LastSprintNumber

	// El esfuerzo va al final, porque apunta a integrantes e historias que ya tienen que estar
	// cargados. Pasa por las mismas reglas que al registrarlo, menos la del día futuro y la de
	// seguir siendo integrante (Fase 2 de la #14).
	for i, ej := range pj.Efforts {
		day, err := parseDay(ej.Day)
		if err != nil {
			return nil, fmt.Errorf("esfuerzo %d: %w", i+1, err)
		}
		effort, err := project.newEffort(EffortData{
			StoryID:    ej.StoryID,
			MemberName: ej.MemberName,
			Activity:   ej.Activity,
			Hours:      ej.Hours,
			Day:        day,
		}, time.Time{}, true)
		if err != nil {
			return nil, fmt.Errorf("esfuerzo %d: %w", i+1, err)
		}
		project.efforts = append(project.efforts, effort)
	}
	return project, nil
}

// sprintsFromJSON reconstruye los sprints de un proyecto (historia #8, tarea 15). No puede
// usar AddSprint, que los crearia abiertos y con numeros nuevos: cada uno se valida con lo
// mismo que usa AddSprint (goal y fechas), y despues se controla lo que solo un archivo puede
// traer mal: numeros repetidos o que el contador no alcanza, y mas de un sprint abierto.
func sprintsFromJSON(in []sprintJSON, lastNumber int) ([]*Sprint, error) {
	numbers := make([]int, 0, len(in))
	for _, sj := range in {
		numbers = append(numbers, sj.Number)
	}
	// el contador guarda el ultimo numero asignado; checkNumbers recibe el proximo
	if err := checkNumbers(numbers, lastNumber+1); err != nil {
		return nil, err
	}
	sprints := make([]*Sprint, 0, len(in))
	var open *Sprint
	for _, sj := range in {
		sprint, err := sprintFromJSON(sj)
		if err != nil {
			return nil, fmt.Errorf("sprint %d: %w", sj.Number, err)
		}
		if sprint.state == SprintOpen {
			if open != nil {
				return nil, fmt.Errorf("%w: el sprint %d y el sprint %d", ErrSprintAlreadyOpen, open.number, sprint.number)
			}
			open = sprint
		}
		sprints = append(sprints, sprint)
	}
	return sprints, nil
}

// sprintFromJSON reconstruye un sprint con su numero y su estado.
func sprintFromJSON(sj sprintJSON) (*Sprint, error) {
	start, err := parseDay(sj.Start)
	if err != nil {
		return nil, err
	}
	end, err := parseDay(sj.End)
	if err != nil {
		return nil, err
	}
	data, err := prepareSprint(SprintData{Goal: sj.Goal, Start: start, End: end})
	if err != nil {
		return nil, err
	}
	state, err := parseSprintState(sj.State)
	if err != nil {
		return nil, err
	}
	return &Sprint{number: sj.Number, goal: data.Goal, start: data.Start, end: data.End, state: state}, nil
}

// storyFromJSON reconstruye una historia con su numero, su estado y su estimacion.
func storyFromJSON(sj storyJSON) (*Story, error) {
	priority, err := ParsePriority(sj.Priority)
	if err != nil {
		return nil, err
	}
	state, err := parseStoryState(sj.State)
	if err != nil {
		return nil, err
	}
	data, err := prepareStory(StoryData{
		Title:              sj.Title,
		Description:        sj.Description,
		Priority:           priority,
		AcceptanceCriteria: sj.AcceptanceCriteria,
	})
	if err != nil {
		return nil, err
	}
	// cero es "sin estimar" (historia #12); cualquier otro valor tiene que estar en la escala
	if sj.StoryPoints != 0 && !inScale(sj.StoryPoints) {
		return nil, ErrStoryPointsInvalid
	}
	// cero es "sin horas estimadas" (historia #16, RN-17); un valor cargado tiene que ser mayor
	// que cero (RN-18). NaN no puede llegar: JSON no tiene como escribirlo
	if sj.EstimatedHours < 0 {
		return nil, ErrEstimatedHoursInvalid
	}
	return &Story{
		id:                 sj.ID,
		title:              data.Title,
		description:        data.Description,
		priority:           data.Priority,
		state:              state,
		storyPoints:        sj.StoryPoints,
		estimatedHours:     sj.EstimatedHours,
		acceptanceCriteria: data.AcceptanceCriteria,
	}, nil
}

// checkNumbers controla los numeros de un grupo (los proyectos del registro o las historias
// de un proyecto) contra su contador: el contador es el numero que va a recibir el proximo,
// asi que tiene que ser al menos 1 y mayor que todos los usados, y ningun numero se repite.
// Es lo que garantiza que al volver a abrir nunca se reparta un numero ya usado (RN-22).
func checkNumbers(ids []int, next int) error {
	if next < 1 {
		return fmt.Errorf("el próximo número a asignar es %d y tiene que ser al menos 1", next)
	}
	seen := map[int]bool{}
	for _, id := range ids {
		if id < 1 || id >= next {
			return fmt.Errorf("el número %d no es válido con el próximo a asignar en %d", id, next)
		}
		if seen[id] {
			return fmt.Errorf("el número %d está repetido", id)
		}
		seen[id] = true
	}
	return nil
}

// parseDay lee una fecha guardada como dia. Un texto vacio es una fecha sin cargar.
func parseDay(text string) (time.Time, error) {
	if text == "" {
		return time.Time{}, nil
	}
	day, err := time.Parse(dayLayout, text)
	if err != nil {
		return time.Time{}, fmt.Errorf("la fecha %q no es un día con la forma aaaa-mm-dd", text)
	}
	return day, nil
}

// parseStoryState lee el estado de una historia por su nombre.
func parseStoryState(text string) (StoryState, error) {
	for _, state := range []StoryState{StoryPending, StoryInProgress, StoryDone} {
		if text == state.String() {
			return state, nil
		}
	}
	return StoryPending, fmt.Errorf("el estado %q no existe", text)
}
