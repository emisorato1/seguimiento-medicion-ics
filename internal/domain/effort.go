package domain

import "time"

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
	effort := &Effort{
		storyID:    data.StoryID,
		memberName: data.MemberName,
		activity:   data.Activity,
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
