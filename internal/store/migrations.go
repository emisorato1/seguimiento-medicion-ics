package store

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"sort"
	"strconv"
	"strings"
)

// migrations son los cambios del esquema, numerados y embebidos en el binario (historia
// #177, D6): "0001_inicial.sql", "0002_...". Un dato nuevo de la conversion suma su
// migracion; nunca se edita una que ya se uso.
//
//go:embed migrations/*.sql
var migrations embed.FS

// migration es un archivo de migracion con su numero.
type migration struct {
	number int
	name   string
}

// knownMigrations lista las migraciones embebidas, en orden de numero.
func knownMigrations() ([]migration, error) {
	names, err := fs.Glob(migrations, "migrations/*.sql")
	if err != nil {
		return nil, err
	}
	var list []migration
	for _, name := range names {
		base := strings.TrimPrefix(name, "migrations/")
		number, err := strconv.Atoi(strings.SplitN(base, "_", 2)[0])
		if err != nil {
			return nil, fmt.Errorf("la migracion %s no empieza con su numero", base)
		}
		list = append(list, migration{number: number, name: name})
	}
	sort.Slice(list, func(i, j int) bool { return list[i].number < list[j].number })
	return list, nil
}

// errNewerSchema es una base con una migracion que este sistema no conoce (RN-13).
var errNewerSchema = fmt.Errorf("es de una versión más nueva del sistema")

// migrate aplica, en una transaccion, las migraciones que la base todavia no tiene. La base
// guarda en PRAGMA user_version el numero de la ultima que se le aplico (D6). Si tiene un
// numero mayor que el de la ultima migracion conocida, no la toca y devuelve errNewerSchema.
func migrate(db *sql.DB) error {
	var current int
	if err := db.QueryRow("PRAGMA user_version").Scan(&current); err != nil {
		return err
	}
	list, err := knownMigrations()
	if err != nil {
		return err
	}
	if len(list) > 0 && current > list[len(list)-1].number {
		return errNewerSchema
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	for _, m := range list {
		if m.number <= current {
			continue
		}
		text, err := migrations.ReadFile(m.name)
		if err != nil {
			return err
		}
		if _, err := tx.Exec(string(text)); err != nil {
			return fmt.Errorf("migracion %s: %w", m.name, err)
		}
		// PRAGMA no admite parametros: el numero sale del nombre del archivo, no del usuario.
		if _, err := tx.Exec(fmt.Sprintf("PRAGMA user_version = %d", m.number)); err != nil {
			return err
		}
	}
	return tx.Commit()
}
