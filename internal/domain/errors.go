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
	// ErrStoryTitleRequired: crear o modificar una historia sin titulo (historia #6, RN-5).
	ErrStoryTitleRequired = errors.New("el titulo de la historia es obligatorio")
	// ErrStoryDescriptionRequired: crear o modificar una historia sin descripcion (historia #6, RN-5).
	ErrStoryDescriptionRequired = errors.New("la descripcion de la historia es obligatoria")
	// ErrStoryCriteriaRequired: historia sin ningun criterio de aceptacion (historia #6, RN-5).
	ErrStoryCriteriaRequired = errors.New("la historia necesita al menos un criterio de aceptacion")
	// ErrStoryDone: modificar una historia ya terminada (historia #6, RN-10).
	ErrStoryDone = errors.New("una historia terminada no se modifica")
	// ErrStoryNotFound: modificar una historia que no existe en el proyecto (historia #6, RN-13).
	ErrStoryNotFound = errors.New("no se encontro la historia")
)
