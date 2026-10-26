package store

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// Importar y exportar usan el archivo de la #77, con su formato de siempre (historia #177,
// D8): el que guardaba la version anterior y el de los datos de demo.

// Export escribe todos los datos de la base en el archivo, con el formato de la #77 (RN-7).
// Escribe con el archivo de la #77, que no deja un archivo a medias.
func Export(d *Database, path string) error {
	registry, err := d.Load()
	if err != nil {
		return err
	}
	return NewFile(path).Save(registry)
}

// Import trae los datos del archivo a la base. Los controla todos con las reglas del dominio
// y solo si las cumplen los guarda (RN-8). Si el archivo no existe, no se puede leer o tiene
// un dato que no cumple las reglas, devuelve un error que cumple errors.Is con
// domain.ErrSavedDataInvalid y no cambia nada (RN-9).
func Import(d *Database, path string) error {
	if _, err := os.Stat(path); errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("%w: no se pudo abrir el archivo %s", domain.ErrSavedDataInvalid, path)
	}
	registry, err := NewFile(path).Load()
	if err != nil {
		return err
	}
	return d.Save(registry)
}
