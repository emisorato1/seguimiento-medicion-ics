package store

import (
	"bytes"
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
