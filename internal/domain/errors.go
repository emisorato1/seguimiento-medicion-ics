package domain

import "errors"

// Errores del dominio, compartidos por todas las historias: cada una agrega los suyos aca.
// Los nombres van en ingles y los mensajes, que los lee el usuario, en español. Los tests
// los comparan con errors.Is, nunca por el texto.
var (
	// ErrProjectNameRequired: crear o modificar un proyecto sin nombre (historia #3, RN-2).
	ErrProjectNameRequired = errors.New("el nombre del proyecto es obligatorio")
	// ErrProjectStartRequired: crear o modificar un proyecto sin fecha de inicio (historia #3, RN-3).
	ErrProjectStartRequired = errors.New("la fecha de inicio del proyecto es obligatoria")
	// ErrProjectDatesInconsistent: fecha de finalizacion anterior a la de inicio (historia #3, RN-4).
	ErrProjectDatesInconsistent = errors.New("las fechas son incoherentes: la fecha de finalizacion es anterior a la de inicio")
)
