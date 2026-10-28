package store

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// Escenario: Exportar e importar da los mismos datos
// Cubre: #177 RN-7, RN-8, RN-11
func TestExportarEImportarDaLoMismo(t *testing.T) {
	dir := t.TempDir()
	registry := unRegistroCompleto(t)
	origen := abrirBase(t, filepath.Join(dir, "origen.db"))
	sinError(t, origen.Save(registry))
	archivo := filepath.Join(dir, "exportado.json")

	sinError(t, Export(origen, archivo))
	destino := abrirBase(t, filepath.Join(dir, "destino.db"))
	sinError(t, Import(destino, archivo))

	cargado, err := destino.Load()
	sinError(t, err)
	if got, want := comoTexto(t, cargado), comoTexto(t, registry); !bytes.Equal(got, want) {
		t.Errorf("la base nueva no tiene los mismos datos.\nOrigen:\n%s\nDestino:\n%s", want, got)
	}
	// El archivo exportado tiene el formato de siempre: lo lee el archivo de la #77.
	leido, err := NewFile(archivo).Load()
	sinError(t, err)
	if got, want := comoTexto(t, leido), comoTexto(t, registry); !bytes.Equal(got, want) {
		t.Errorf("el archivo exportado no tiene el formato de la #77:\n%s", got)
	}
}

// Esquema del escenario: Importar algo que no se puede no cambia nada
// Cubre: #177 RN-9, RN-10
func TestImportarAlgoQueNoSePuedeNoCambiaNada(t *testing.T) {
	invalido := `{"next_id": 2, "projects": [{"id": 1, "name": "Sistema", "start": "2026-10-01", "next_story_id": 2,
		"members": [], "efforts": [], "stories": [{"id": 1, "title": "Una", "description": "Algo", "priority": "Alta",
		"state": "Pendiente", "story_points": 4, "acceptance_criteria": ["Se puede"]}]}]}`
	casos := []struct {
		nombre  string
		base    *domain.ProjectRegistry
		archivo func(t *testing.T, dir string) string
		err     error
	}{
		{"un archivo que no existe", domain.NewProjectRegistry(), func(t *testing.T, dir string) string {
			return filepath.Join(dir, "no-esta.json")
		}, domain.ErrSavedDataInvalid},
		{"un archivo con una historia de 4 story points", domain.NewProjectRegistry(), func(t *testing.T, dir string) string {
			path := filepath.Join(dir, "invalido.json")
			sinError(t, os.WriteFile(path, []byte(invalido), 0o644))
			return path
		}, domain.ErrStoryPointsInvalid},
		{"una base con datos", unRegistroCompleto(t), func(t *testing.T, dir string) string {
			path := filepath.Join(dir, "valido.json")
			sinError(t, NewFile(path).Save(unProyecto(t, func(p *domain.Project) {})))
			return path
		}, ErrDatabaseNotEmpty},
	}
	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			dir := t.TempDir()
			db := abrirBase(t, filepath.Join(dir, "datos-ics.db"))
			sinError(t, db.Save(caso.base))
			path := caso.archivo(t, dir)

			err := Import(db, path)

			if !errors.Is(err, caso.err) || !strings.Contains(err.Error(), path) {
				t.Errorf("se esperaba %v con la ruta del archivo, salio: %v", caso.err, err)
			}
			cargado, err := db.Load()
			sinError(t, err)
			if got, want := comoTexto(t, cargado), comoTexto(t, caso.base); !bytes.Equal(got, want) {
				t.Errorf("la base tenia que quedar como estaba, y tiene:\n%s", got)
			}
		})
	}
}

// existe dice si hay un archivo en esa ruta.
func existe(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Escenario: Lo guardado con la versión anterior se importa solo
// Escenario: Lo de la versión anterior no vuelve aunque se borre todo
// Escenario: Si lo de la versión anterior no se puede importar, no se toca
// Cubre: #177 RN-14, RN-15, RN-9, D10
func TestElArchivoViejoSeImportaUnaSolaVez(t *testing.T) {
	viejo := unProyecto(t, func(p *domain.Project) {})
	preparar := func(t *testing.T, contenido *domain.ProjectRegistry) (string, *Database) {
		t.Helper()
		dir := t.TempDir()
		if contenido != nil {
			sinError(t, NewFile(filepath.Join(dir, LegacyFile)).Save(contenido))
		}
		return dir, abrirBase(t, filepath.Join(dir, "datos-ics.db"))
	}

	t.Run("con la base vacia lo importa, lo renombra y no lo vuelve a importar", func(t *testing.T) {
		dir, db := preparar(t, viejo)

		imported, err := ImportLegacy(db, dir)

		sinError(t, err)
		if !imported || existe(filepath.Join(dir, LegacyFile)) || !existe(filepath.Join(dir, LegacyFile+".importado")) {
			t.Errorf("se esperaba importarlo y renombrarlo, salio importado=%v", imported)
		}
		cargado, err := db.Load()
		sinError(t, err)
		if got, want := comoTexto(t, cargado), comoTexto(t, viejo); !bytes.Equal(got, want) {
			t.Errorf("la base tenia que tener los datos del archivo viejo, tiene:\n%s", got)
		}
		if imported, err := ImportLegacy(db, dir); imported || err != nil {
			t.Errorf("la segunda vez no tenia que importar nada, salio importado=%v, %v", imported, err)
		}

		// Aunque despues se borre todo, no se vuelve a importar.
		sinError(t, db.Save(domain.NewProjectRegistry()))
		if imported, err := ImportLegacy(db, dir); imported || err != nil {
			t.Errorf("con la base vaciada despues no tenia que volver a importar, salio importado=%v, %v", imported, err)
		}
	})

	t.Run("sin archivo viejo no hace nada", func(t *testing.T) {
		dir, db := preparar(t, nil)
		if imported, err := ImportLegacy(db, dir); imported || err != nil {
			t.Errorf("sin archivo no tenia que importar nada, salio importado=%v, %v", imported, err)
		}
	})

	t.Run("con la base con datos no lo toca", func(t *testing.T) {
		dir, db := preparar(t, viejo)
		sinError(t, db.Save(unRegistroCompleto(t)))
		if imported, err := ImportLegacy(db, dir); imported || err != nil || !existe(filepath.Join(dir, LegacyFile)) {
			t.Errorf("con datos en la base no tenia que importar ni renombrar, salio importado=%v, %v", imported, err)
		}
	})

	t.Run("un archivo viejo con una historia de 4 story points no se toca y la base queda vacia", func(t *testing.T) {
		dir, db := preparar(t, nil)
		path := filepath.Join(dir, LegacyFile)
		invalido := []byte(`{"next_id": 2, "projects": [{"id": 1, "name": "Sistema de metricas", "start": "2026-10-01", "next_story_id": 2, "members": [],
			"stories": [{"id": 1, "title": "Crear proyectos", "description": "Cargar un proyecto", "priority": "Alta", "state": "Pendiente", "story_points": 4, "acceptance_criteria": ["Se crea con su nombre"]}],
			"efforts": []}]}`)
		sinError(t, os.WriteFile(path, invalido, 0o644))

		imported, err := ImportLegacy(db, dir)

		if imported || !errors.Is(err, domain.ErrSavedDataInvalid) || !strings.Contains(err.Error(), path) {
			t.Errorf("se esperaba ErrSavedDataInvalid con la ruta del archivo, salio importado=%v, %v", imported, err)
		}
		if err == nil || !strings.Contains(err.Error(), domain.ErrStoryPointsInvalid.Error()) {
			t.Errorf("el aviso tenia que decir que dato esta mal (los story points), salio: %v", err)
		}
		if quedo, _ := os.ReadFile(path); !bytes.Equal(quedo, invalido) || existe(path+".importado") {
			t.Error("el archivo viejo tenia que quedar como estaba, sin renombrar")
		}
		if empty, _ := db.IsEmpty(); !empty {
			t.Error("la base tenia que seguir vacia")
		}
	})

	t.Run("si no se puede renombrar, la base vuelve a quedar vacia", func(t *testing.T) {
		dir, db := preparar(t, viejo)
		sinError(t, os.WriteFile(filepath.Join(dir, LegacyFile+".importado"), []byte("de antes"), 0o644))

		imported, err := ImportLegacy(db, dir)

		if imported || err == nil || !strings.Contains(err.Error(), LegacyFile+".importado") {
			t.Errorf("se esperaba un error que nombre el archivo que ya existe, salio importado=%v, %v", imported, err)
		}
		if empty, _ := db.IsEmpty(); !empty || !existe(filepath.Join(dir, LegacyFile)) {
			t.Error("la base tenia que quedar vacia y el archivo viejo, en su lugar")
		}
	})
}

// Esquema del escenario: Lo que no se puede pedir al exportar o importar (exportar sobre la base)
// Cubre: #177 RN-16
func TestExportarAlMismoArchivoDeLaBaseSeRechaza(t *testing.T) {
	path := filepath.Join(t.TempDir(), "datos-ics.db")
	db := abrirBase(t, path)
	registry := unRegistroCompleto(t)
	sinError(t, db.Save(registry))
	antes, err := os.ReadFile(path)
	sinError(t, err)

	err = Export(db, path)

	if !errors.Is(err, ErrExportOverDatabase) || !strings.Contains(err.Error(), path) {
		t.Errorf("se esperaba ErrExportOverDatabase con la ruta, salio: %v", err)
	}
	if despues, _ := os.ReadFile(path); !bytes.Equal(despues, antes) {
		t.Error("la base tenia que quedar como estaba")
	}
	// Tambien con otra forma de escribir la misma ruta.
	otra := filepath.Join(filepath.Dir(path), ".", "datos-ics.db")
	if err := Export(db, otra); !errors.Is(err, ErrExportOverDatabase) {
		t.Errorf("con la ruta escrita de otra forma se esperaba ErrExportOverDatabase, salio: %v", err)
	}
}
