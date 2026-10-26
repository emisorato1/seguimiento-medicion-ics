package store

import (
	"bytes"
	"path/filepath"
	"testing"
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
