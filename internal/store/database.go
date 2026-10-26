package store

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"

	// El driver de SQLite en Go puro, sin CGO: compila igual en Linux, macOS y Windows
	// (historia #177, D1). Se registra con el nombre "sqlite".
	_ "modernc.org/sqlite"
)

// Database es la base SQLite donde la web guarda el registro de proyectos (historia #177).
// Leer y escribir la base es de este paquete; convertir el registro y validarlo al
// reconstruirlo sigue siendo del dominio (D3).
type Database struct {
	db   *sql.DB
	path string
}

// OpenDatabase abre la base de esa ruta, o la crea si no existe (RN-5), y le aplica las
// migraciones que le falten (D6). Usa una sola conexion: todos los cambios ya pasan de a uno
// por el candado de la web (D7).
func OpenDatabase(path string) (*Database, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("no se pudo abrir la base %s", path)
	}
	db.SetMaxOpenConns(1)
	if err := migrate(db); err != nil {
		db.Close()
		return nil, fmt.Errorf("no se pudo abrir la base %s: %w", path, err)
	}
	return &Database{db: db, path: path}, nil
}

// Close cierra la base. En Windows, un archivo abierto no se puede borrar ni renombrar.
func (d *Database) Close() error {
	return d.db.Close()
}

// tables son las tablas de la base, en el orden en que se vacian al guardar.
var tables = []string{"registry", "projects", "members", "stories", "acceptance_criteria", "efforts",
	"sprints", "sprint_stories", "sprint_completed_stories", "sprint_points_at_close", "defects",
	"voting_rounds", "voting_round_participants", "votes"}

// Save guarda el registro entero dentro de una transaccion: vacia todas las tablas y vuelve a
// insertar todo (D4). Si algo falla, la transaccion se deshace y la base queda como estaba
// (RN-2), y el error cumple errors.Is con ErrNotSaved.
func (d *Database) Save(registry *domain.ProjectRegistry) error {
	data, err := registry.ToJSON()
	if err != nil {
		return d.notSaved()
	}
	rows, err := decodeRegistry(data)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrNotSaved, err)
	}
	tx, err := d.db.Begin()
	if err != nil {
		return d.notSaved()
	}
	defer tx.Rollback()
	for _, table := range tables {
		if _, err := tx.Exec("DELETE FROM " + table); err != nil {
			return d.notSaved()
		}
	}
	if err := insertRegistry(tx, rows); err != nil {
		return d.notSaved()
	}
	if err := tx.Commit(); err != nil {
		return d.notSaved()
	}
	return nil
}

// notSaved arma el error de un guardado que fallo, en castellano y con la ruta de la base.
func (d *Database) notSaved() error {
	return fmt.Errorf("%w: no se pudo escribir la base %s", ErrNotSaved, d.path)
}

// insertRegistry escribe las filas del registro en la transaccion. Cada lista guarda su
// orden en "position", porque el dominio lo conserva (integrantes, backlog, esfuerzo,
// sprints, defectos y rondas).
func insertRegistry(tx *sql.Tx, registry fileRegistry) error {
	if _, err := tx.Exec("INSERT INTO registry (id, next_id) VALUES (1, ?)", registry.NextID); err != nil {
		return err
	}
	for i, p := range registry.Projects {
		if err := insertProject(tx, i, p); err != nil {
			return err
		}
	}
	return nil
}

// insertProject escribe un proyecto con todo lo que tiene.
func insertProject(tx *sql.Tx, position int, p fileProject) error {
	exec := func(query string, args ...any) error {
		_, err := tx.Exec(query, args...)
		return err
	}
	if err := exec(`INSERT INTO projects (id, position, name, description, start_day, end_day,
		next_story_id, last_sprint_number, last_defect_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		p.ID, position, p.Name, p.Description, p.Start, p.End, p.NextStoryID, p.LastSprintNumber, p.LastDefectID); err != nil {
		return err
	}
	for i, m := range p.Members {
		if err := exec("INSERT INTO members (project_id, position, name, role) VALUES (?, ?, ?, ?)",
			p.ID, i, m.Name, m.Role); err != nil {
			return err
		}
	}
	for i, st := range p.Stories {
		if err := exec(`INSERT INTO stories (project_id, id, position, title, description, priority, state,
			story_points, estimated_hours) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			p.ID, st.ID, i, st.Title, st.Description, st.Priority, st.State, st.StoryPoints, st.EstimatedHours); err != nil {
			return err
		}
		for j, text := range st.AcceptanceCriteria {
			if err := exec("INSERT INTO acceptance_criteria (project_id, story_id, position, text) VALUES (?, ?, ?, ?)",
				p.ID, st.ID, j, text); err != nil {
				return err
			}
		}
	}
	for i, e := range p.Efforts {
		if err := exec(`INSERT INTO efforts (project_id, position, story_id, member_name, activity, hours, day)
			VALUES (?, ?, ?, ?, ?, ?, ?)`, p.ID, i, e.StoryID, e.MemberName, e.Activity, e.Hours, e.Day); err != nil {
			return err
		}
	}
	for i, sp := range p.Sprints {
		if err := exec(`INSERT INTO sprints (project_id, number, position, goal, start_day, end_day, state, closed_on)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, p.ID, sp.Number, i, sp.Goal, sp.Start, sp.End, sp.State, sp.ClosedOn); err != nil {
			return err
		}
		for j, id := range sp.Stories {
			if err := exec("INSERT INTO sprint_stories (project_id, sprint_number, position, story_id) VALUES (?, ?, ?, ?)",
				p.ID, sp.Number, j, id); err != nil {
				return err
			}
		}
		for j, id := range sp.CompletedStories {
			if err := exec("INSERT INTO sprint_completed_stories (project_id, sprint_number, position, story_id) VALUES (?, ?, ?, ?)",
				p.ID, sp.Number, j, id); err != nil {
				return err
			}
		}
		for id, points := range sp.PointsAtClose {
			if err := exec("INSERT INTO sprint_points_at_close (project_id, sprint_number, story_id, points) VALUES (?, ?, ?, ?)",
				p.ID, sp.Number, id, points); err != nil {
				return err
			}
		}
	}
	for i, d := range p.Defects {
		if err := exec(`INSERT INTO defects (project_id, id, position, title, description, severity, state,
			story_id, detected_in, resolved_in) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			p.ID, d.ID, i, d.Title, d.Description, d.Severity, d.State, d.StoryID, d.DetectedIn, d.ResolvedIn); err != nil {
			return err
		}
	}
	for i, r := range p.VotingRounds {
		if err := exec(`INSERT INTO voting_rounds (project_id, position, story_id, number, revealed, agreed)
			VALUES (?, ?, ?, ?, ?, ?)`, p.ID, i, r.StoryID, r.Number, r.Revealed, r.Agreed); err != nil {
			return err
		}
		for j, name := range r.Participants {
			if err := exec("INSERT INTO voting_round_participants (project_id, round_position, position, name) VALUES (?, ?, ?, ?)",
				p.ID, i, j, name); err != nil {
				return err
			}
		}
		for j, v := range r.Votes {
			if err := exec("INSERT INTO votes (project_id, round_position, position, member_name, points) VALUES (?, ?, ?, ?, ?)",
				p.ID, i, j, v.MemberName, v.Points); err != nil {
				return err
			}
		}
	}
	return nil
}

// Load lee la base y reconstruye el registro con la conversion del dominio, que controla cada
// dato con las mismas reglas que al cargarlo desde la web (RN-4, D3). Una base sin nada
// guardado da un registro sin proyectos (RN-5). Si algo no se puede leer o no cumple las
// reglas, el error cumple errors.Is con domain.ErrSavedDataInvalid y la base no se toca.
func (d *Database) Load() (*domain.ProjectRegistry, error) {
	registry, err := readRegistry(d.db)
	if err != nil {
		return nil, fmt.Errorf("%w: no se pudo leer la base %s", domain.ErrSavedDataInvalid, d.path)
	}
	data, err := json.Marshal(registry)
	if err != nil {
		return nil, fmt.Errorf("%w: no se pudo leer la base %s", domain.ErrSavedDataInvalid, d.path)
	}
	return domain.RegistryFromJSON(data)
}

// readRegistry lee las filas de la base en el orden en que se guardaron. Sin fila en
// registry, como en una base nueva, el proximo proyecto es el 1, como en un registro nuevo.
// La base tiene una sola conexion (D7): cada tabla se lee entera y se cierra antes de leer la
// siguiente, y despues se reparte por proyecto.
func readRegistry(db *sql.DB) (fileRegistry, error) {
	registry := fileRegistry{NextID: 1, Projects: []fileProject{}}
	err := db.QueryRow("SELECT next_id FROM registry WHERE id = 1").Scan(&registry.NextID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fileRegistry{}, err
	}

	// index lleva de un proyecto a su lugar en la lista, para repartir lo que se lee.
	index := map[int]int{}
	err = query(db, `SELECT id, name, description, start_day, end_day, next_story_id,
		last_sprint_number, last_defect_id FROM projects ORDER BY position`, func(rows *sql.Rows) error {
		p := fileProject{Members: []fileMember{}, Stories: []fileStory{}, Efforts: []fileEffort{}}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Start, &p.End, &p.NextStoryID,
			&p.LastSprintNumber, &p.LastDefectID); err != nil {
			return err
		}
		index[p.ID] = len(registry.Projects)
		registry.Projects = append(registry.Projects, p)
		return nil
	})
	if err != nil {
		return fileRegistry{}, err
	}
	project := func(id int) (*fileProject, error) {
		i, ok := index[id]
		if !ok {
			return nil, fmt.Errorf("una fila apunta al proyecto %d, que no esta", id)
		}
		return &registry.Projects[i], nil
	}

	steps := []struct {
		query string
		read  func(rows *sql.Rows) error
	}{
		{"SELECT project_id, name, role FROM members ORDER BY project_id, position", func(rows *sql.Rows) error {
			var id int
			var m fileMember
			if err := rows.Scan(&id, &m.Name, &m.Role); err != nil {
				return err
			}
			p, err := project(id)
			if err != nil {
				return err
			}
			p.Members = append(p.Members, m)
			return nil
		}},
		{`SELECT project_id, id, title, description, priority, state, story_points, estimated_hours
			FROM stories ORDER BY project_id, position`, func(rows *sql.Rows) error {
			var id int
			st := fileStory{AcceptanceCriteria: []string{}}
			if err := rows.Scan(&id, &st.ID, &st.Title, &st.Description, &st.Priority, &st.State,
				&st.StoryPoints, &st.EstimatedHours); err != nil {
				return err
			}
			p, err := project(id)
			if err != nil {
				return err
			}
			p.Stories = append(p.Stories, st)
			return nil
		}},
		{"SELECT project_id, story_id, text FROM acceptance_criteria ORDER BY project_id, story_id, position", func(rows *sql.Rows) error {
			var id, storyID int
			var text string
			if err := rows.Scan(&id, &storyID, &text); err != nil {
				return err
			}
			p, err := project(id)
			if err != nil {
				return err
			}
			for i := range p.Stories {
				if p.Stories[i].ID == storyID {
					p.Stories[i].AcceptanceCriteria = append(p.Stories[i].AcceptanceCriteria, text)
					return nil
				}
			}
			return fmt.Errorf("un criterio apunta a la historia %d, que no esta", storyID)
		}},
		{"SELECT project_id, story_id, member_name, activity, hours, day FROM efforts ORDER BY project_id, position", func(rows *sql.Rows) error {
			var id int
			var e fileEffort
			if err := rows.Scan(&id, &e.StoryID, &e.MemberName, &e.Activity, &e.Hours, &e.Day); err != nil {
				return err
			}
			p, err := project(id)
			if err != nil {
				return err
			}
			p.Efforts = append(p.Efforts, e)
			return nil
		}},
		{"SELECT project_id, number, goal, start_day, end_day, state, closed_on FROM sprints ORDER BY project_id, position", func(rows *sql.Rows) error {
			var id int
			var sp fileSprint
			if err := rows.Scan(&id, &sp.Number, &sp.Goal, &sp.Start, &sp.End, &sp.State, &sp.ClosedOn); err != nil {
				return err
			}
			p, err := project(id)
			if err != nil {
				return err
			}
			p.Sprints = append(p.Sprints, sp)
			return nil
		}},
		{"SELECT project_id, sprint_number, story_id FROM sprint_stories ORDER BY project_id, sprint_number, position", func(rows *sql.Rows) error {
			return readSprintList(rows, project, func(sp *fileSprint, storyID int) { sp.Stories = append(sp.Stories, storyID) })
		}},
		{"SELECT project_id, sprint_number, story_id FROM sprint_completed_stories ORDER BY project_id, sprint_number, position", func(rows *sql.Rows) error {
			return readSprintList(rows, project, func(sp *fileSprint, storyID int) {
				sp.CompletedStories = append(sp.CompletedStories, storyID)
			})
		}},
		{"SELECT project_id, sprint_number, story_id, points FROM sprint_points_at_close ORDER BY project_id, sprint_number, story_id", func(rows *sql.Rows) error {
			var id, number, storyID, points int
			if err := rows.Scan(&id, &number, &storyID, &points); err != nil {
				return err
			}
			sp, err := findSprint(project, id, number)
			if err != nil {
				return err
			}
			if sp.PointsAtClose == nil {
				sp.PointsAtClose = map[int]int{}
			}
			sp.PointsAtClose[storyID] = points
			return nil
		}},
		{`SELECT project_id, id, title, description, severity, state, story_id, detected_in, resolved_in
			FROM defects ORDER BY project_id, position`, func(rows *sql.Rows) error {
			var id int
			var d fileDefect
			if err := rows.Scan(&id, &d.ID, &d.Title, &d.Description, &d.Severity, &d.State, &d.StoryID,
				&d.DetectedIn, &d.ResolvedIn); err != nil {
				return err
			}
			p, err := project(id)
			if err != nil {
				return err
			}
			p.Defects = append(p.Defects, d)
			return nil
		}},
		{"SELECT project_id, story_id, number, revealed, agreed FROM voting_rounds ORDER BY project_id, position", func(rows *sql.Rows) error {
			var id int
			r := fileVotingRound{Participants: []string{}}
			if err := rows.Scan(&id, &r.StoryID, &r.Number, &r.Revealed, &r.Agreed); err != nil {
				return err
			}
			p, err := project(id)
			if err != nil {
				return err
			}
			p.VotingRounds = append(p.VotingRounds, r)
			return nil
		}},
		{"SELECT project_id, round_position, name FROM voting_round_participants ORDER BY project_id, round_position, position", func(rows *sql.Rows) error {
			var id, round int
			var name string
			if err := rows.Scan(&id, &round, &name); err != nil {
				return err
			}
			r, err := findRound(project, id, round)
			if err != nil {
				return err
			}
			r.Participants = append(r.Participants, name)
			return nil
		}},
		{"SELECT project_id, round_position, member_name, points FROM votes ORDER BY project_id, round_position, position", func(rows *sql.Rows) error {
			var id, round int
			var v fileVote
			if err := rows.Scan(&id, &round, &v.MemberName, &v.Points); err != nil {
				return err
			}
			r, err := findRound(project, id, round)
			if err != nil {
				return err
			}
			r.Votes = append(r.Votes, v)
			return nil
		}},
	}
	for _, step := range steps {
		if err := query(db, step.query, step.read); err != nil {
			return fileRegistry{}, err
		}
	}
	return registry, nil
}

// query corre la consulta, le pasa cada fila a read y cierra las filas antes de volver.
func query(db *sql.DB, q string, read func(rows *sql.Rows) error) error {
	rows, err := db.Query(q)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		if err := read(rows); err != nil {
			return err
		}
	}
	return rows.Err()
}

// readSprintList lee una fila de una lista de historias de un sprint y se la pasa a add.
func readSprintList(rows *sql.Rows, project func(int) (*fileProject, error), add func(*fileSprint, int)) error {
	var id, number, storyID int
	if err := rows.Scan(&id, &number, &storyID); err != nil {
		return err
	}
	sp, err := findSprint(project, id, number)
	if err != nil {
		return err
	}
	add(sp, storyID)
	return nil
}

// findSprint busca un sprint ya leido, por su proyecto y su numero.
func findSprint(project func(int) (*fileProject, error), id, number int) (*fileSprint, error) {
	p, err := project(id)
	if err != nil {
		return nil, err
	}
	for i := range p.Sprints {
		if p.Sprints[i].Number == number {
			return &p.Sprints[i], nil
		}
	}
	return nil, fmt.Errorf("una fila apunta al sprint %d, que no esta", number)
}

// findRound busca una ronda ya leida, por su proyecto y su lugar en la lista.
func findRound(project func(int) (*fileProject, error), id, position int) (*fileVotingRound, error) {
	p, err := project(id)
	if err != nil {
		return nil, err
	}
	if position < 0 || position >= len(p.VotingRounds) {
		return nil, fmt.Errorf("una fila apunta a la ronda %d, que no esta", position)
	}
	return &p.VotingRounds[position], nil
}
