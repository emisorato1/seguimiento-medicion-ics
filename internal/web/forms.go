package web

import (
	"errors"
	"net/http"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// form es un formulario tal como lo escribio el usuario: lo que tiene cada campo, el error
// que va junto a cada uno y el que no corresponde a ninguno (RN-12, RN-13).
type form struct {
	Values  map[string]string
	Errors  map[string]string
	General string
}

// field es un campo listo para mostrar: su nombre, lo escrito y su error si lo tiene.
type field struct {
	Name  string
	Value string
	Error string
}

// newForm arma un formulario vacio, para la primera vez que se muestra.
func newForm() *form {
	return &form{Values: make(map[string]string), Errors: make(map[string]string)}
}

// readForm arma el formulario con lo que mando el navegador en esos campos.
func readForm(r *http.Request, names ...string) *form {
	f := newForm()
	for _, name := range names {
		f.Values[name] = r.FormValue(name)
	}
	return f
}

// Field devuelve un campo del formulario para la plantilla.
func (f *form) Field(name string) field {
	return field{Name: name, Value: f.Values[name], Error: f.Errors[name]}
}

// fieldErrors dice junto a que campo va cada error del dominio (D6). Se compara con
// errors.Is, nunca con el texto. Un error que no esta aca no se pierde: va arriba.
var fieldErrors = []struct {
	err   error
	field string
}{
	{domain.ErrProjectNameRequired, "nombre"},
	{domain.ErrProjectStartRequired, "inicio"},
	{domain.ErrProjectDatesInconsistent, "fin"},
}

// reject anota el rechazo del dominio junto al campo que lo causo, con el mismo mensaje
// del dominio (RN-12). Si no corresponde a ningun campo, va arriba del formulario (RN-13).
func (f *form) reject(err error) {
	for _, fe := range fieldErrors {
		if errors.Is(err, fe.err) {
			f.Errors[fe.field] = err.Error()
			return
		}
	}
	f.General = err.Error()
}
