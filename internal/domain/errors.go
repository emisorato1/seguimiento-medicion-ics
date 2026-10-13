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
	// ErrProjectNotFound: consultar o modificar un proyecto que no existe (historia #3, RN-6).
	ErrProjectNotFound = errors.New("no se encontro el proyecto")
	// ErrStoryPriorityInvalid: prioridad que no es Alta, Media ni Baja (historia #6, RN-4).
	ErrStoryPriorityInvalid = errors.New("la prioridad debe ser Alta, Media o Baja")
	// ErrMemberRoleInvalid: rol que no es Agile Enabler ni Product Builder (historia #4, RN-4).
	ErrMemberRoleInvalid = errors.New("el rol debe ser Agile Enabler o Product Builder")
	// ErrMemberNameRequired: agregar un integrante sin nombre y apellido (historia #4, RN-2).
	ErrMemberNameRequired = errors.New("el nombre y apellido del integrante es obligatorio")
	// ErrMemberDuplicate: agregar un integrante que ya esta cargado en el proyecto (historia #4, RN-5).
	ErrMemberDuplicate = errors.New("el integrante ya esta cargado en el proyecto")
	// ErrMemberAgileEnablerTaken: agregar un segundo Agile Enabler al mismo proyecto (historia #4, RN-7).
	ErrMemberAgileEnablerTaken = errors.New("el proyecto ya tiene un Agile Enabler")
	// ErrMemberNotFound: quitar un integrante que no pertenece al proyecto (historia #4, RN-11).
	ErrMemberNotFound = errors.New("no se encontro el integrante en el proyecto")
)
