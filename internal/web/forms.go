package web

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// errInvalidDate es el error de forma de una fecha. No es una regla de negocio sino la forma
// del dato (RN-14), igual que la RN-5 de la consola.
var errInvalidDate = errors.New("la fecha no es válida")

// errInvalidPoints es el error de forma de los story points: no se eligio ninguno de la
// lista, o llego algo que no es un numero (RN-14). Que el numero este en la escala lo
// decide el dominio.
var errInvalidPoints = errors.New("hay que elegir los story points de la lista")

// dateLayout es como manda la fecha el calendario del navegador (D11).
const dateLayout = "2006-01-02"

// form es un formulario tal como lo escribio el usuario: lo que tiene cada campo, el error
// que va junto a cada uno y el que no corresponde a ninguno (RN-12, RN-13).
type form struct {
	Values  map[string]string
	Errors  map[string]string
	General string
	// focus es el primer campo que quedo con error: ahi arranca el foco al volver a
	// mostrar el formulario, para que el lector de pantalla anuncie el error.
	focus string
}

// field es un campo listo para mostrar: su nombre, lo escrito, su error si lo tiene y si
// el foco arranca en el.
type field struct {
	Name  string
	Value string
	Error string
	Focus bool
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
	return field{Name: name, Value: f.Values[name], Error: f.Errors[name], Focus: name == f.focus}
}

// fail anota el error junto al campo. El primero que falla es el que recibe el foco.
func (f *form) fail(name, message string) {
	f.Errors[name] = message
	if f.focus == "" {
		f.focus = name
	}
}

// date lee una fecha del formulario. Vacia es una fecha sin indicar, y que sea obligatoria
// lo decide el dominio. Un texto que no es una fecha se marca junto al campo (RN-14): quien
// llama no tiene que llamar al dominio si el formulario quedo con errores.
func (f *form) date(name string) time.Time {
	value := strings.TrimSpace(f.Values[name])
	if value == "" {
		return time.Time{}
	}
	day, err := time.Parse(dateLayout, value)
	if err != nil {
		f.fail(name, errInvalidDate.Error())
		return time.Time{}
	}
	return day
}

// number lee un numero entero del formulario. Vacio o con letras se marca junto al campo
// con el error que recibe (RN-14), y quien llama no llama al dominio.
func (f *form) number(name string, shapeErr error) int {
	value, err := strconv.Atoi(strings.TrimSpace(f.Values[name]))
	if err != nil {
		f.fail(name, shapeErr.Error())
		return 0
	}
	return value
}

// hasErrors indica si algun campo quedo con un error de forma.
func (f *form) hasErrors() bool {
	return len(f.Errors) > 0
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
	{domain.ErrMemberNameRequired, "nombre"},
	{domain.ErrMemberRoleInvalid, "rol"},
	{domain.ErrMemberDuplicate, "nombre"},
	{domain.ErrMemberAgileEnablerTaken, "rol"},
	{domain.ErrStoryTitleRequired, "titulo"},
	{domain.ErrStoryDescriptionRequired, "descripcion"},
	{domain.ErrStoryPriorityInvalid, "prioridad"},
	{domain.ErrStoryCriteriaRequired, "criterios"},
	{domain.ErrStoryPointsInvalid, "puntos"},
	{domain.ErrSprintGoalRequired, "goal"},
	{domain.ErrSprintStartRequired, "inicio"},
	{domain.ErrSprintEndRequired, "fin"},
	{domain.ErrSprintDatesInconsistent, "fin"},
	{domain.ErrStoryNotEstimated, "historia"},
	{domain.ErrStoryAlreadyInSprint, "historia"},
}

// reject anota el rechazo del dominio junto al campo que lo causo, con el mismo mensaje
// del dominio (RN-12). Si no corresponde a ningun campo, va arriba del formulario (RN-13).
func (f *form) reject(err error) {
	for _, fe := range fieldErrors {
		if errors.Is(err, fe.err) {
			f.fail(fe.field, err.Error())
			return
		}
	}
	f.General = err.Error()
}
