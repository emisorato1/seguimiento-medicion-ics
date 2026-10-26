package store

import (
	"bytes"
	"path/filepath"
	"slices"
	"testing"
	"time"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// abrirBase abre la base de esa ruta y la cierra al terminar el test: en Windows un archivo
// abierto no se puede borrar, y t.TempDir lo borra al final.
func abrirBase(t *testing.T, path string) *Database {
	t.Helper()
	db, err := OpenDatabase(path)
	if err != nil {
		t.Fatalf("no se pudo abrir la base %s: %v", path, err)
	}
	t.Cleanup(func() { db.Close() })
	return db
}

// versionDelEsquema devuelve el numero de la ultima migracion aplicada a la base.
func versionDelEsquema(t *testing.T, db *Database) int {
	t.Helper()
	var version int
	if err := db.db.QueryRow("PRAGMA user_version").Scan(&version); err != nil {
		t.Fatalf("no se pudo leer la version del esquema: %v", err)
	}
	return version
}

// Escenario: Abrir por primera vez
// Cubre: #177 RN-5, D2, D6
func TestAbrirUnaBaseNuevaCreaElEsquema(t *testing.T) {
	path := filepath.Join(t.TempDir(), "datos-ics.db")

	db := abrirBase(t, path)

	if v := versionDelEsquema(t, db); v != 1 {
		t.Errorf("se esperaba el esquema en la version 1, esta en %d", v)
	}
	rows, err := db.db.Query("SELECT name FROM sqlite_master WHERE type = 'table' ORDER BY name")
	if err != nil {
		t.Fatalf("no se pudieron leer las tablas: %v", err)
	}
	var tablas []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			t.Fatalf("no se pudo leer una tabla: %v", err)
		}
		tablas = append(tablas, name)
	}
	rows.Close()
	quiero := []string{"acceptance_criteria", "defects", "efforts", "members", "projects", "registry",
		"sprint_completed_stories", "sprint_points_at_close", "sprint_stories", "sprints", "stories",
		"votes", "voting_round_participants", "voting_rounds"}
	if !slices.Equal(tablas, quiero) {
		t.Errorf("se esperaban las tablas %v, salieron %v", quiero, tablas)
	}

	// Volver a abrirla no vuelve a crear el esquema ni cambia su version.
	db.Close()
	otra := abrirBase(t, path)
	if v := versionDelEsquema(t, otra); v != 1 {
		t.Errorf("al volver a abrirla se esperaba la version 1, esta en %d", v)
	}
}

// guardarYReabrir guarda el registro en la base de esa ruta, la cierra, la vuelve a abrir y
// carga lo guardado, como al cerrar el sistema y volver a abrirlo.
func guardarYReabrir(t *testing.T, path string, registry *domain.ProjectRegistry) *domain.ProjectRegistry {
	t.Helper()
	db, err := OpenDatabase(path)
	if err != nil {
		t.Fatalf("no se pudo abrir la base: %v", err)
	}
	if err := db.Save(registry); err != nil {
		db.Close()
		t.Fatalf("no se pudo guardar: %v", err)
	}
	db.Close()
	cargado, err := abrirBase(t, path).Load()
	if err != nil {
		t.Fatalf("no se pudo cargar lo guardado: %v", err)
	}
	return cargado
}

// Escenario: Abrir por primera vez
// Cubre: #177 RN-1, RN-3, RN-5
func TestGuardarYCargarUnRegistro(t *testing.T) {
	// Una base nueva, sin guardar nada, se carga como un registro sin proyectos.
	nueva, err := abrirBase(t, filepath.Join(t.TempDir(), "datos-ics.db")).Load()
	if err != nil {
		t.Fatalf("no se pudo cargar una base nueva: %v", err)
	}
	if got, want := comoTexto(t, nueva), comoTexto(t, domain.NewProjectRegistry()); !bytes.Equal(got, want) {
		t.Errorf("una base nueva tenia que cargarse vacia, salio:\n%s", got)
	}

	conUnProyecto := domain.NewProjectRegistry()
	if _, err := conUnProyecto.Create(domain.ProjectData{
		Name: "Sistema de métricas", Description: "Seguimiento del integrador",
		Start: time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC), End: time.Date(2026, 11, 2, 0, 0, 0, 0, time.UTC),
	}); err != nil {
		t.Fatalf("no se pudo crear el proyecto: %v", err)
	}
	casos := []struct {
		nombre   string
		registry *domain.ProjectRegistry
	}{
		{"un registro sin proyectos", domain.NewProjectRegistry()},
		{"un registro con un proyecto", conUnProyecto},
	}
	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			cargado := guardarYReabrir(t, filepath.Join(t.TempDir(), "datos-ics.db"), caso.registry)

			if got, want := comoTexto(t, cargado), comoTexto(t, caso.registry); !bytes.Equal(got, want) {
				t.Errorf("lo cargado no es lo que se guardo.\nGuardado:\n%s\nCargado:\n%s", want, got)
			}
		})
	}
}
