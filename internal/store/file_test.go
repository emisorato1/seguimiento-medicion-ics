package store

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strings"
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

// Los errores se fabrican con el lugar de las cosas y no con permisos de Unix: Windows ignora
// los permisos de una carpeta y no impide leer un archivo de permiso 0o000, así que esos casos
// no fallaban ahí (defecto #108). Que haya un archivo donde se espera una carpeta, o una
// carpeta donde se espera un archivo, hace fallar la operación en cualquier sistema.

// unaCarpetaEnLugarDelArchivo pone en la ruta del archivo de datos una carpeta con algo
// adentro, y devuelve lo que tiene, para controlar después que quedó igual.
func unaCarpetaEnLugarDelArchivo(t *testing.T, path string) []byte {
	t.Helper()
	contenido := []byte(`{"projects": [ esto no es json`)
	if err := os.Mkdir(path, 0o700); err != nil {
		t.Fatalf("no se pudo crear la carpeta de prueba: %v", err)
	}
	if err := os.WriteFile(filepath.Join(path, "algo"), contenido, 0o600); err != nil {
		t.Fatalf("no se pudo llenar la carpeta de prueba: %v", err)
	}
	return contenido
}

// unArchivoEnLugarDeLaCarpeta pone un archivo común donde tendría que estar la carpeta del
// archivo de datos, y devuelve la ruta del archivo de datos: escribir ahí adentro no se puede.
func unArchivoEnLugarDeLaCarpeta(t *testing.T, dir string) string {
	t.Helper()
	carpeta := filepath.Join(dir, "carpeta")
	if err := os.WriteFile(carpeta, []byte("no soy una carpeta"), 0o600); err != nil {
		t.Fatalf("no se pudo preparar el archivo de prueba: %v", err)
	}
	return filepath.Join(carpeta, "datos-ics.json")
}

// sinTemporales falla si en la carpeta quedó algún archivo temporal del guardado.
func sinTemporales(t *testing.T, dir string) {
	t.Helper()
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("no se pudo leer la carpeta de prueba: %v", err)
	}
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".datos-") {
			t.Fatalf("se esperaba que no quedara ningun temporal, quedo %q", entry.Name())
		}
	}
}

// Escenario: Los datos guardados estan danados (el archivo)
// Cubre: #77 RN-24
func TestUnArchivoDanadoQuedaComoEstaba(t *testing.T) {
	cases := []struct {
		name    string
		prepare func(t *testing.T, path string) (string, []byte)
	}{
		{"texto que no es JSON", func(t *testing.T, path string) (string, []byte) {
			original := []byte(`{"projects": [ esto no es json`)
			if err := os.WriteFile(path, original, 0o600); err != nil {
				t.Fatalf("no se pudo preparar el archivo de prueba: %v", err)
			}
			return path, original
		}},
		{"en el lugar del archivo hay una carpeta", func(t *testing.T, path string) (string, []byte) {
			return filepath.Join(path, "algo"), unaCarpetaEnLugarDelArchivo(t, path)
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "datos-ics.json")
			left, original := tc.prepare(t, path)

			registry, err := NewFile(path).Load()

			if !errors.Is(err, domain.ErrSavedDataInvalid) {
				t.Fatalf("se esperaba ErrSavedDataInvalid, salio: %v", err)
			}
			if registry != nil {
				t.Fatalf("se esperaba que no se cargara ningun registro")
			}
			after, err := os.ReadFile(left)
			if err != nil {
				t.Fatalf("no se pudo releer lo que habia: %v", err)
			}
			if !bytes.Equal(after, original) {
				t.Fatalf("se esperaba que lo que habia quedara igual, quedo:\n%s", after)
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
		{"la carpeta del archivo es un archivo", unArchivoEnLugarDeLaCarpeta},
		{"en el lugar del archivo hay una carpeta", func(t *testing.T, dir string) string {
			path := filepath.Join(dir, "datos-ics.json")
			unaCarpetaEnLugarDelArchivo(t, path)
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
			sinTemporales(t, dir)
		})
	}
}

// Sin escenario: hallazgo de front en el PR #92, los mensajes van en castellano (AGENTS.md)
// Cubre: #77 RN-24
//
// El error del sistema operativo viene en inglés: el mensaje dice que no se pudo abrir el
// archivo y cuál, sin ese detalle.
func TestElMensajeDeUnArchivoQueNoSePuedeLeerEstaEnCastellano(t *testing.T) {
	path := filepath.Join(t.TempDir(), "datos-ics.json")
	unaCarpetaEnLugarDelArchivo(t, path)
	want := "los datos guardados no se pueden leer: no se pudo abrir el archivo " + path

	_, err := NewFile(path).Load()

	if err == nil || err.Error() != want {
		t.Fatalf("se esperaba el mensaje %q, salio: %v", want, err)
	}
}

// Sin escenario: hallazgo de front en el PR #92, los mensajes van en castellano (AGENTS.md)
// Cubre: #77 RN-25
//
// Lo mismo al guardar: el aviso que muestra la web dice qué archivo no se pudo escribir, sin
// el detalle en inglés del sistema operativo.
func TestElMensajeDeUnCambioQueNoSePudoGuardarEstaEnCastellano(t *testing.T) {
	path := unArchivoEnLugarDeLaCarpeta(t, t.TempDir())
	want := "el cambio no quedó guardado: no se pudo escribir el archivo " + path

	err := NewFile(path).Save(unRegistro(t))

	if err == nil || err.Error() != want {
		t.Fatalf("se esperaba el mensaje %q, salio: %v", want, err)
	}
}
