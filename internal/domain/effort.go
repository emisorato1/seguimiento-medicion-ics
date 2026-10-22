package domain

import (
	"math"
	"strings"
	"time"
)

// Effort es un registro de esfuerzo: las horas que un integrante trabajó de verdad en una
// historia, un día, y qué hizo (historia #14). Vive dentro del proyecto, como los
// integrantes y el backlog: así los registros de un proyecto nunca suman en otro (RN-16).
// Sus datos están ocultos fuera del paquete: solo se cargan con LogEffort, que valida.
type Effort struct {
	storyID    int
	memberName string
	activity   string
	hours      float64
	day        time.Time
}

// StoryID devuelve el número de la historia sobre la que se trabajó.
func (e *Effort) StoryID() int {
	return e.storyID
}

// MemberName devuelve el nombre y apellido de quien trabajó, tal como está cargado en el
// proyecto.
func (e *Effort) MemberName() string {
	return e.memberName
}

// Activity devuelve qué se hizo en esas horas.
func (e *Effort) Activity() string {
	return e.activity
}

// Hours devuelve las horas trabajadas, con sus fracciones.
func (e *Effort) Hours() float64 {
	return e.hours
}

// Day devuelve el día en que se trabajó, sin hora.
func (e *Effort) Day() time.Time {
	return e.day
}

// EffortData son los datos con los que se registra esfuerzo.
type EffortData struct {
	StoryID    int
	MemberName string
	Activity   string
	Hours      float64
	Day        time.Time
}

// LogEffort registra esfuerzo en el proyecto (RN-1). El día de hoy lo recibe de quien llama:
// el dominio no lee el reloj (AGENTS.md).
func (p *Project) LogEffort(data EffortData, today time.Time) (*Effort, error) {
	// el integrante se busca como en la #4, ignorando mayúsculas y los espacios de los
	// extremos, y se guarda su nombre tal como está cargado (RN-7, RN-8)
	member := p.findMember(strings.TrimSpace(data.MemberName))
	if member == nil {
		return nil, ErrMemberNotFound
	}
	// la historia tiene que ser del proyecto; su estado no importa, porque las horas se
	// cargan después de trabajar (RN-9, RN-10)
	if _, err := p.Story(data.StoryID); err != nil {
		return nil, err
	}
	if data.Day.IsZero() {
		return nil, ErrEffortDayRequired
	}
	// el dia se compara solo con hoy, sin la hora, y no con las fechas del proyecto ni de los
	// sprints (RN-4, RN-17)
	if toDay(data.Day).After(toDay(today)) {
		return nil, ErrEffortDayInFuture
	}
	// "mayor que cero" escrito asi rechaza tambien NaN, que no es mayor ni menor que nada;
	// el infinito si es mayor que cero, y se rechaza aparte porque no son horas (RN-3)
	if !(data.Hours > 0) || math.IsInf(data.Hours, 0) {
		return nil, ErrEffortHoursInvalid
	}
	// la actividad se controla ya sin los espacios de los extremos, así un texto de solo
	// espacios cuenta como vacío, y se guarda así (RN-6)
	activity := strings.TrimSpace(data.Activity)
	if activity == "" {
		return nil, ErrEffortActivityRequired
	}
	effort := &Effort{
		storyID:    data.StoryID,
		memberName: member.name,
		activity:   activity,
		hours:      data.Hours,
		day:        data.Day,
	}
	p.efforts = append(p.efforts, effort)
	return effort, nil
}

// Efforts devuelve los registros de esfuerzo del proyecto, en el orden en que se cargaron
// (RN-15). Sin registros devuelve una lista vacía, que no es un error. La lista es nueva,
// pero los registros son los guardados.
func (p *Project) Efforts() []*Effort {
	efforts := make([]*Effort, len(p.efforts))
	copy(efforts, p.efforts)
	return efforts
}
