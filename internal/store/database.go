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

// insertRegistry escribe las filas del registro en la transaccion.
func insertRegistry(tx *sql.Tx, registry fileRegistry) error {
	if _, err := tx.Exec("INSERT INTO registry (id, next_id) VALUES (1, ?)", registry.NextID); err != nil {
		return err
	}
	for i, p := range registry.Projects {
		if _, err := tx.Exec(`INSERT INTO projects (id, position, name, description, start_day, end_day,
			next_story_id, last_sprint_number, last_defect_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			p.ID, i, p.Name, p.Description, p.Start, p.End, p.NextStoryID, p.LastSprintNumber, p.LastDefectID); err != nil {
			return err
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
func readRegistry(db *sql.DB) (fileRegistry, error) {
	registry := fileRegistry{NextID: 1, Projects: []fileProject{}}
	err := db.QueryRow("SELECT next_id FROM registry WHERE id = 1").Scan(&registry.NextID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fileRegistry{}, err
	}
	rows, err := db.Query(`SELECT id, name, description, start_day, end_day, next_story_id,
		last_sprint_number, last_defect_id FROM projects ORDER BY position`)
	if err != nil {
		return fileRegistry{}, err
	}
	defer rows.Close()
	for rows.Next() {
		p := fileProject{Members: []fileMember{}, Stories: []fileStory{}, Efforts: []fileEffort{}}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Start, &p.End, &p.NextStoryID,
			&p.LastSprintNumber, &p.LastDefectID); err != nil {
			return fileRegistry{}, err
		}
		registry.Projects = append(registry.Projects, p)
	}
	return registry, rows.Err()
}
