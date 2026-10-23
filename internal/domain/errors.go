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
	ErrProjectDatesInconsistent = errors.New("las fechas son incoherentes: la fecha de finalización es anterior a la de inicio")
	// ErrProjectNotFound: consultar o modificar un proyecto que no existe (historia #3, RN-6).
	ErrProjectNotFound = errors.New("no se encontró el proyecto")
	// ErrStoryPriorityInvalid: prioridad que no es Alta, Media ni Baja (historia #6, RN-4).
	ErrStoryPriorityInvalid = errors.New("la prioridad debe ser Alta, Media o Baja")
	// ErrStoryTitleRequired: crear o modificar una historia sin titulo (historia #6, RN-5).
	ErrStoryTitleRequired = errors.New("el título de la historia es obligatorio")
	// ErrStoryDescriptionRequired: crear o modificar una historia sin descripcion (historia #6, RN-5).
	ErrStoryDescriptionRequired = errors.New("la descripción de la historia es obligatoria")
	// ErrStoryCriteriaRequired: historia sin ningun criterio de aceptacion (historia #6, RN-5).
	ErrStoryCriteriaRequired = errors.New("la historia necesita al menos un criterio de aceptación")
	// ErrStoryDone: modificar una historia ya terminada (historia #6, RN-10).
	ErrStoryDone = errors.New("una historia terminada no se modifica")
	// ErrStoryNotFound: modificar una historia que no existe en el proyecto (historia #6, RN-13).
	ErrStoryNotFound = errors.New("no se encontró la historia")
	// ErrMemberRoleInvalid: rol que no es Agile Enabler ni Product Builder (historia #4, RN-4).
	ErrMemberRoleInvalid = errors.New("el rol debe ser Agile Enabler o Product Builder")
	// ErrMemberNameRequired: agregar un integrante sin nombre y apellido (historia #4, RN-2).
	ErrMemberNameRequired = errors.New("el nombre y apellido del integrante es obligatorio")
	// ErrMemberDuplicate: agregar un integrante que ya esta cargado en el proyecto (historia #4, RN-5).
	ErrMemberDuplicate = errors.New("el integrante ya está cargado en el proyecto")
	// ErrMemberAgileEnablerTaken: agregar un segundo Agile Enabler al mismo proyecto (historia #4, RN-7).
	ErrMemberAgileEnablerTaken = errors.New("el proyecto ya tiene un Agile Enabler")
	// ErrMemberNotFound: quitar un integrante que no pertenece al proyecto (historia #4, RN-11).
	ErrMemberNotFound = errors.New("no se encontró el integrante en el proyecto")
	// ErrDefectTitleRequired y ErrDefectDescriptionRequired: registrar un defecto sin título o
	// sin descripción (historia #15, RN-4). Los errores de la #15 van acá y no al final, para
	// no chocar con los que agregan otras historias al mismo tiempo.
	ErrDefectTitleRequired       = errors.New("falta el título del defecto")
	ErrDefectDescriptionRequired = errors.New("falta la descripción del defecto")
	// ErrDefectSeverityInvalid: severidad que no es Crítica, Mayor ni Menor, o ninguna
	// (historia #15, RN-6).
	ErrDefectSeverityInvalid = errors.New("la severidad debe ser Crítica, Mayor o Menor")
	// ErrDefectResolveNeedsSprint: pasar un defecto a Resuelto cambiándole el estado, sin
	// decir en qué sprint se resolvió (historia #15, RN-10).
	ErrDefectResolveNeedsSprint = errors.New("para resolver un defecto hay que indicar en qué sprint se resolvió")
	// ErrDefectStateInvalid: cambiar un defecto a un estado que no es Abierto ni En curso
	// (historia #15, RN-9).
	ErrDefectStateInvalid = errors.New("el estado de un defecto solo cambia entre Abierto y En curso")
	// ErrDefectResolvedBeforeDetected: resolver un defecto en un sprint de número menor al de
	// detección (historia #15, RN-13).
	ErrDefectResolvedBeforeDetected = errors.New("un defecto no se resuelve en un sprint anterior al de detección")
	// ErrDefectAlreadyResolved: cambiarle el estado o volver a resolver un defecto Resuelto
	// (historia #15, RN-14).
	ErrDefectAlreadyResolved = errors.New("el defecto ya está resuelto")
	// ErrDefectNotFound: cambiarle el estado o resolver un defecto que no existe en el
	// proyecto (historia #15, RN-19).
	ErrDefectNotFound = errors.New("no se encontró el defecto")
	// ErrSavedDataInvalid: los datos guardados no se pueden leer o tienen algo que el dominio
	// no acepta (historia #77, RN-24). Va antes de los de la #12 y no al final, para no chocar
	// con los errores que agregan otras historias al mismo tiempo.
	ErrSavedDataInvalid = errors.New("los datos guardados no se pueden leer")
	// ErrEffortHoursInvalid: registrar horas en cero, negativas, infinitas o que no son un
	// número (historia #14, RN-3). Van aca y no al final, para no chocar con los errores que
	// agregan otras historias al mismo tiempo.
	ErrEffortHoursInvalid = errors.New("las horas tienen que ser mayores que cero")
	// ErrEffortDayInFuture: registrar horas de un día posterior al de hoy (historia #14, RN-4).
	ErrEffortDayInFuture = errors.New("no se pueden cargar horas de un día que todavía no pasó")
	// ErrEffortDayRequired: registrar horas sin indicar el día (historia #14, RN-5).
	ErrEffortDayRequired = errors.New("la fecha es obligatoria")
	// ErrEffortActivityRequired: registrar horas sin decir qué se hizo (historia #14, RN-6).
	ErrEffortActivityRequired = errors.New("la actividad realizada es obligatoria")
	// ErrStoryPointsInvalid: valor fuera de la escala, cero o negativo (historia #12, RN-2, RN-3).
	ErrStoryPointsInvalid = errors.New("los story points tienen que ser uno de estos valores: " + scaleText())
	// ErrStoryEstimateDone: estimar, reestimar o quitar la estimacion de una historia
	// Terminada (historia #12, RN-8).
	ErrStoryEstimateDone = errors.New("la estimación de una historia terminada no se modifica")
	// ErrEstimatedHoursInvalid: cargar horas estimadas en cero o negativas (historia #16, RN-18).
	ErrEstimatedHoursInvalid = errors.New("las horas estimadas tienen que ser mayores que cero")
	// ErrSprintNotFound: consultar un sprint con un numero que no existe en el proyecto
	// (historia #8, RN-15). Lo usan tambien la #11 y la #15.
	ErrSprintNotFound = errors.New("no se encontró el sprint")
	// ErrSprintGoalRequired: crear un sprint sin Sprint Goal, o con uno de solo espacios
	// (historia #8, RN-4).
	ErrSprintGoalRequired = errors.New("el Sprint Goal es obligatorio")
	// ErrSprintStartRequired: crear un sprint sin fecha de inicio (historia #8, RN-5).
	ErrSprintStartRequired = errors.New("la fecha de inicio del sprint es obligatoria")
	// ErrSprintEndRequired: crear un sprint sin fecha de fin (historia #8, RN-5).
	ErrSprintEndRequired = errors.New("la fecha de fin del sprint es obligatoria")
	// ErrSprintDatesInconsistent: fecha de fin del sprint anterior a la de inicio (historia #8,
	// RN-6).
	ErrSprintDatesInconsistent = errors.New("las fechas del sprint son incoherentes: la fecha de fin es anterior a la de inicio")
	// ErrSprintAlreadyOpen: crear un sprint mientras el proyecto tiene otro abierto (historia
	// #8, RN-8). Se devuelve envuelto con el numero del sprint abierto, asi el aviso dice cual
	// es; errors.Is lo sigue reconociendo.
	ErrSprintAlreadyOpen = errors.New("ya hay un sprint abierto en el proyecto")
	// ErrStoryNotEstimated: asignar a un sprint una historia sin estimar (historia #9, RN-3).
	ErrStoryNotEstimated = errors.New("solo se asignan a un sprint las historias estimadas")
	// ErrStoryAlreadyInSprint: asignar una historia que ya esta en el Sprint Backlog (historia
	// #9, RN-5).
	ErrStoryAlreadyInSprint = errors.New("la historia ya está asignada al sprint")
	// ErrNoOpenSprint: asignar o quitar una historia cuando el proyecto no tiene sprint abierto
	// (historia #9, RN-2). Lo reusa la #10 al cerrar.
	ErrNoOpenSprint = errors.New("el proyecto no tiene un sprint abierto")
	// ErrStoryNotInSprint: quitar del sprint una historia que no esta en su Sprint Backlog
	// (historia #9, RN-8).
	ErrStoryNotInSprint = errors.New("la historia no está en el sprint")
)
