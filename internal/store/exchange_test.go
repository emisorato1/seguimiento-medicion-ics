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
