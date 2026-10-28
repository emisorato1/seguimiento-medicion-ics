package store

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// Importar y exportar usan el archivo de la #77, con su formato de siempre (historia #177,
// D8): el que guardaba la version anterior y el de los datos de demo.

// ErrDatabaseNotEmpty: se pidio importar un archivo a una base que ya tiene datos (historia
// #177, RN-10). Importar reemplazaria todo lo que hay.
var ErrDatabaseNotEmpty = errors.New("la base ya tiene datos y no se importa sobre datos existentes")

// ErrExportOverDatabase: se pidio exportar al mismo archivo de la base (historia #177, RN-16).
// Escribir el archivo ahi reemplazaria la base por el texto.
var ErrExportOverDatabase = errors.New("no se exporta al mismo archivo de la base: elegí otro nombre para el archivo")

// Export escribe todos los datos de la base en el archivo, con el formato de la #77 (RN-7).
// Escribe con el archivo de la #77, que no deja un archivo a medias. Si el archivo es la base
// misma, no escribe nada y devuelve ErrExportOverDatabase (RN-16).
func Export(d *Database, path string) error {
	if sameFile(path, d.path) {
		return fmt.Errorf("%w (%s)", ErrExportOverDatabase, path)
	}
	registry, err := d.Load()
	if err != nil {
		return err
	}
	return NewFile(path).Save(registry)
}

// Import trae los datos del archivo a una base sin datos. Los controla todos con las reglas
// del dominio y solo si las cumplen los guarda (RN-8). Si la base ya tiene datos devuelve
// ErrDatabaseNotEmpty (RN-10). Si el archivo no existe, no se puede leer o tiene un dato que
// no cumple las reglas, devuelve un error que cumple errors.Is con domain.ErrSavedDataInvalid
// (RN-9). En los dos casos el error nombra el archivo y la base no cambia.
func Import(d *Database, path string) error {
	empty, err := d.IsEmpty()
	if err != nil {
		return err
	}
	if !empty {
		return fmt.Errorf("%w (se pidió importar %s)", ErrDatabaseNotEmpty, path)
	}
	if _, err := os.Stat(path); errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("%w: no se pudo abrir el archivo %s", domain.ErrSavedDataInvalid, path)
	}
	registry, err := NewFile(path).Load()
	if err != nil {
		return fmt.Errorf("%w (en el archivo %s)", err, path)
	}
	return d.Save(registry)
}

// LegacyFile es el archivo donde la version anterior, la de la #77, guardaba los datos.
const LegacyFile = "datos-ics.json"

// ImportLegacy trae solo el archivo de la version anterior, si esta en dir y la base no tiene
// datos (RN-14), y lo renombra a LegacyFile + ".importado" para no traerlo otra vez (RN-15).
// Devuelve si lo importo. Sin archivo, o con la base con datos, no hace nada. Si el archivo no
// cumple las reglas, no lo toca y la base queda vacia (RN-9). El orden es importar y guardar,
// y despues renombrar; si el renombre falla, vacia la base otra vez, asi se importa una sola
// vez cuando se resuelva (D10).
func ImportLegacy(d *Database, dir string) (bool, error) {
	path := filepath.Join(dir, LegacyFile)
	if _, err := os.Stat(path); err != nil {
		return false, nil
	}
	empty, err := d.IsEmpty()
	if err != nil || !empty {
		return false, err
	}
	// En Linux y macOS renombrar pisa el destino sin avisar: se mira antes, para no perder el
	// archivo de una importacion anterior.
	target := path + ".importado"
	if _, err := os.Stat(target); err == nil {
		return false, fmt.Errorf("ya existe %s, así que no se importa %s para no importarlo dos veces: mové o borrá uno de los dos y volvé a abrir el sistema", target, path)
	}
	if err := Import(d, path); err != nil {
		return false, err
	}
	if err := os.Rename(path, target); err != nil {
		if emptyErr := d.Save(domain.NewProjectRegistry()); emptyErr != nil {
			return false, emptyErr
		}
		return false, fmt.Errorf("no se pudo renombrar %s a %s, así que la base quedó vacía para importarlo una sola vez: revisá que se pueda escribir en la carpeta y volvé a abrir el sistema", path, target)
	}
	return true, nil
}

// sameFile dice si las dos rutas llevan al mismo archivo, aunque esten escritas de otra forma.
func sameFile(a, b string) bool {
	infoA, errA := os.Stat(a)
	infoB, errB := os.Stat(b)
	if errA == nil && errB == nil {
		return os.SameFile(infoA, infoB)
	}
	absA, errA := filepath.Abs(a)
	absB, errB := filepath.Abs(b)
	return errA == nil && errB == nil && absA == absB
}
