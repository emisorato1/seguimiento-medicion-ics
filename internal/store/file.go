// Package store guarda el registro de proyectos en un archivo, para que los datos sigan al
// cerrar y volver a abrir el sistema (historia #77, D13, D17 y D18). Solo lee y escribe el
// archivo: convertir el registro a texto y validarlo al reconstruirlo es del dominio (D14).
package store

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// ErrNotSaved: el registro no se pudo guardar en el archivo (historia #77, RN-25). El cambio
// sigue en memoria y el proximo guardado que funcione lo incluye (D19).
var ErrNotSaved = errors.New("el cambio no quedó guardado")

// File es el archivo donde se guarda el registro.
type File struct {
	path string
}

// NewFile prepara el archivo de datos en esa ruta. No lo abre ni lo crea: eso pasa al leer
// y al guardar.
func NewFile(path string) *File {
	return &File{path: path}
}

// Load lee el registro guardado. Si el archivo no existe, como la primera vez, devuelve un
// registro vacio, y eso no es un error (RN-23). Si existe pero no se puede leer, o el
// dominio no acepta lo que tiene, devuelve un error que cumple errors.Is con
// domain.ErrSavedDataInvalid y no toca el archivo (RN-24, D18).
func (f *File) Load() (*domain.ProjectRegistry, error) {
	data, err := os.ReadFile(f.path)
	if errors.Is(err, fs.ErrNotExist) {
		return domain.NewProjectRegistry(), nil
	}
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrSavedDataInvalid, err)
	}
	return domain.RegistryFromJSON(data)
}

// Save reescribe el archivo entero con el registro (D13). Escribe primero en un temporal de
// la misma carpeta y despues lo renombra (D17): renombrar reemplaza el archivo de una sola
// vez, asi un corte a mitad de la escritura deja el archivo anterior entero y no uno roto.
// El temporal va en la misma carpeta porque renombrar entre discos distintos no es de una
// sola vez. Si algo falla devuelve un error que cumple errors.Is con ErrNotSaved, deja el
// archivo anterior como estaba y borra el temporal.
func (f *File) Save(registry *domain.ProjectRegistry) error {
	data, err := registry.ToJSON()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrNotSaved, err)
	}
	tmp, err := os.CreateTemp(filepath.Dir(f.path), ".datos-*.tmp")
	if err != nil {
		return fmt.Errorf("%w: %v", ErrNotSaved, err)
	}
	_, err = tmp.Write(data)
	if err == nil {
		// Sync obliga a que el contenido llegue al disco antes de renombrar: sin eso, un
		// corte de luz justo despues del renombre puede dejar el archivo vacio.
		err = tmp.Sync()
	}
	if closeErr := tmp.Close(); err == nil {
		err = closeErr
	}
	if err == nil {
		err = os.Rename(tmp.Name(), f.path)
	}
	if err != nil {
		os.Remove(tmp.Name())
		return fmt.Errorf("%w: %v", ErrNotSaved, err)
	}
	return nil
}
