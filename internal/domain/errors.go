package domain

import "errors"

// Errores del dominio, compartidos por todas las historias: cada una agrega los suyos aca.
// Los nombres van en ingles y los mensajes, que los lee el usuario, en español. Los tests
// los comparan con errors.Is, nunca por el texto.
var (
	// ErrProjectNameRequired: crear o modificar un proyecto sin nombre (historia #3, RN-2).
	ErrProjectNameRequired = errors.New("el nombre del proyecto es obligatorio")
)
