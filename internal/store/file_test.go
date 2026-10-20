package store

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// unRegistro arma un registro con un proyecto, un integrante y una historia estimada.
func unRegistro(t *testing.T) *domain.ProjectRegistry {
	t.Helper()
	registry := domain.NewProjectRegistry()
	project, err := registry.Create(domain.ProjectData{
		Name: "Sistema de metricas", Start: time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		t.Fatalf("no se pudo crear el proyecto de prueba: %v", err)
	}
	if _, err := project.AddMember("Ana Perez", domain.RoleAgileEnabler); err != nil {
		t.Fatalf("no se pudo agregar el integrante de prueba: %v", err)
	}
	story, err := project.AddStory(domain.StoryData{
		Title: "Registrar esfuerzo", Description: "Cargar horas", Priority: domain.PriorityHigh,
		AcceptanceCriteria: []string{"Valida las horas"},
	})
	if err != nil {
		t.Fatalf("no se pudo crear la historia de prueba: %v", err)
	}
	if err := project.EstimateStory(story.ID(), 3); err != nil {
		t.Fatalf("no se pudo estimar la historia de prueba: %v", err)
	}
	return registry
}

// comoTexto devuelve el registro convertido, para comparar dos registros enteros: la
// comparacion dato por dato ya la hace el test de la conversion en el dominio.
func comoTexto(t *testing.T, registry *domain.ProjectRegistry) []byte {
	t.Helper()
	data, err := registry.ToJSON()
	if err != nil {
		t.Fatalf("no se pudo convertir el registro: %v", err)
	}
	return data
}

// Escenario: Las listas vacias dicen que hacer (fila "el sistema sin datos guardados")
// Cubre: #77 RN-23
func TestSinDatosGuardadosSeArrancaSinProyectos(t *testing.T) {
	file := NewFile(filepath.Join(t.TempDir(), "datos-ics.json"))

	registry, err := file.Load()

	if err != nil {
		t.Fatalf("se esperaba arrancar sin error cuando no hay datos guardados, salio: %v", err)
	}
	if len(registry.Projects()) != 0 {
		t.Fatalf("se esperaba un registro sin proyectos, salieron %d", len(registry.Projects()))
	}
}

// Escenario: Los datos siguen al cerrar y volver a abrir (el archivo)
// Cubre: #77 RN-21
func TestLoGuardadoSeVuelveALeerIgual(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "datos-ics.json")
	registry := unRegistro(t)

	if err := NewFile(path).Save(registry); err != nil {
		t.Fatalf("se esperaba guardar sin error, salio: %v", err)
	}
	loaded, err := NewFile(path).Load()
	if err != nil {
		t.Fatalf("se esperaba volver a leer sin error, salio: %v", err)
	}
	if !bytes.Equal(comoTexto(t, loaded), comoTexto(t, registry)) {
		t.Fatalf("se esperaba leer lo mismo que se guardo:\n%s\nsalio:\n%s", comoTexto(t, registry), comoTexto(t, loaded))
	}

	// guardar otra vez reemplaza el archivo, no le agrega nada
	if _, err := registry.Create(domain.ProjectData{Name: "Otro", Start: time.Date(2026, 10, 5, 0, 0, 0, 0, time.UTC)}); err != nil {
		t.Fatalf("no se pudo crear el segundo proyecto de prueba: %v", err)
	}
	if err := NewFile(path).Save(registry); err != nil {
		t.Fatalf("se esperaba volver a guardar sin error, salio: %v", err)
	}
	loaded, err = NewFile(path).Load()
	if err != nil {
		t.Fatalf("se esperaba volver a leer sin error, salio: %v", err)
	}
	if len(loaded.Projects()) != 2 {
		t.Fatalf("se esperaban 2 proyectos despues de guardar de nuevo, salieron %d", len(loaded.Projects()))
	}

	// en la carpeta queda solo el archivo de datos: ningun temporal
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("no se pudo leer la carpeta de prueba: %v", err)
	}
	if len(entries) != 1 || entries[0].Name() != "datos-ics.json" {
		names := []string{}
		for _, entry := range entries {
			names = append(names, entry.Name())
		}
		t.Fatalf("se esperaba que en la carpeta quedara solo datos-ics.json, quedaron %q", names)
	}
}

// sinPermisos saltea el test si corre como administrador, que puede leer y escribir aunque
// el archivo o la carpeta no lo permitan.
func sinPermisos(t *testing.T) {
	t.Helper()
	if os.Geteuid() == 0 {
		t.Skip("como administrador los permisos no frenan nada")
	}
}

// Escenario: Los datos guardados estan danados (el archivo)
// Cubre: #77 RN-24
func TestUnArchivoDanadoQuedaComoEstaba(t *testing.T) {
	cases := []struct {
		name string
		mode os.FileMode
	}{
		{"texto que no es JSON", 0o600},
		{"archivo que no se puede leer", 0o000},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.mode == 0o000 {
				sinPermisos(t)
			}
			path := filepath.Join(t.TempDir(), "datos-ics.json")
			original := []byte(`{"projects": [ esto no es json`)
			if err := os.WriteFile(path, original, tc.mode); err != nil {
				t.Fatalf("no se pudo preparar el archivo de prueba: %v", err)
			}

			registry, err := NewFile(path).Load()

			if !errors.Is(err, domain.ErrSavedDataInvalid) {
				t.Fatalf("se esperaba ErrSavedDataInvalid, salio: %v", err)
			}
			if registry != nil {
				t.Fatalf("se esperaba que no se cargara ningun registro")
			}
			os.Chmod(path, 0o600)
			after, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("no se pudo releer el archivo de prueba: %v", err)
			}
			if !bytes.Equal(after, original) {
				t.Fatalf("se esperaba que el archivo quedara igual, quedo:\n%s", after)
			}
		})
	}
}

// Escenario: Un cambio que no se pudo guardar se avisa (el archivo)
// Cubre: #77 RN-25
func TestUnCambioQueNoSePudoGuardarSeAvisa(t *testing.T) {
	cases := []struct {
		name    string
		prepare func(t *testing.T, dir string) string
	}{
		{"carpeta donde no se puede escribir", func(t *testing.T, dir string) string {
			sinPermisos(t)
			if err := os.Chmod(dir, 0o500); err != nil {
				t.Fatalf("no se pudo quitar el permiso de escritura: %v", err)
			}
			t.Cleanup(func() { os.Chmod(dir, 0o700) })
			return filepath.Join(dir, "datos-ics.json")
		}},
		{"en el lugar del archivo hay una carpeta", func(t *testing.T, dir string) string {
			path := filepath.Join(dir, "datos-ics.json")
			if err := os.Mkdir(path, 0o700); err != nil {
				t.Fatalf("no se pudo crear la carpeta de prueba: %v", err)
			}
			if err := os.WriteFile(filepath.Join(path, "algo"), []byte("x"), 0o600); err != nil {
				t.Fatalf("no se pudo llenar la carpeta de prueba: %v", err)
			}
			return path
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			dir := t.TempDir()
			path := tc.prepare(t, dir)

			err := NewFile(path).Save(unRegistro(t))

			if !errors.Is(err, ErrNotSaved) {
				t.Fatalf("se esperaba ErrNotSaved, salio: %v", err)
			}
			entries, err := os.ReadDir(dir)
			if err != nil {
				t.Fatalf("no se pudo leer la carpeta de prueba: %v", err)
			}
			for _, entry := range entries {
				if entry.Name() != "datos-ics.json" {
					t.Fatalf("se esperaba que no quedara ningun temporal, quedo %q", entry.Name())
				}
			}
		})
	}
}

// Sin escenario: hallazgo de front en el PR #92, los mensajes van en castellano (AGENTS.md)
// Cubre: #77 RN-24
//
// El error del sistema operativo viene en inglés («permission denied»): el mensaje dice que no
// se pudo abrir el archivo y cuál, sin ese detalle.
func TestElMensajeDeUnArchivoQueNoSePuedeLeerEstaEnCastellano(t *testing.T) {
	sinPermisos(t)
	path := filepath.Join(t.TempDir(), "datos-ics.json")
	if err := os.WriteFile(path, []byte(`{}`), 0o000); err != nil {
		t.Fatalf("no se pudo preparar el archivo de prueba: %v", err)
	}
	want := "los datos guardados no se pueden leer: no se pudo abrir el archivo " + path

	_, err := NewFile(path).Load()

	if err == nil || err.Error() != want {
		t.Fatalf("se esperaba el mensaje %q, salio: %v", want, err)
	}
}
