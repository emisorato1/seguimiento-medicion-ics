package store

import (
	"database/sql"
	"fmt"

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
