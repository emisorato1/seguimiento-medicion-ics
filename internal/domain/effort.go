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
// el dominio no lee el reloj (AGENTS.md). Valida todo antes de guardar: si algo falla, el
// proyecto queda como estaba (RN-18).
func (p *Project) LogEffort(data EffortData, today time.Time) (*Effort, error) {
	effort, err := p.newEffort(data, today, false)
	if err != nil {
		return nil, err
	}
	p.efforts = append(p.efforts, effort)
	return effort, nil
}

// newEffort arma un registro validando sus datos en el orden de la Fase 2: integrante,
// historia, día, horas y actividad. No lo guarda. La carga de los datos guardados lo usa con
// loading en verdadero, y ahí cambian dos controles (Fase 2): no se mira el día futuro,
// porque al cargar no hay un "hoy" con qué comparar y un día que no era futuro al guardarse
// no puede volverse futuro después; y se acepta el registro de alguien que ya no es
// integrante, porque quitarlo del proyecto no borra sus horas (decisión de Emiliano del 21/10).
func (p *Project) newEffort(data EffortData, today time.Time, loading bool) (*Effort, error) {
	// el integrante se busca como en la #4, ignorando mayúsculas y los espacios de los
	// extremos, y se guarda su nombre tal como está cargado (RN-7, RN-8). Un nombre vacío no
	// se acepta nunca, ni al cargar
	name := strings.TrimSpace(data.MemberName)
	member := p.findMember(name)
	if member != nil {
		name = member.name
	} else if !loading || name == "" {
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
	if !loading && toDay(data.Day).After(toDay(today)) {
		return nil, ErrEffortDayInFuture
	}
	// "mayor que cero" escrito asi rechaza tambien NaN, que no es mayor ni menor que nada;
	// el infinito si es mayor que cero, y se rechaza aparte porque no son horas (RN-3)
	if !(data.Hours > 0) || math.IsInf(data.Hours, 0) {
		return nil, ErrEffortHoursInvalid
	}
	// un registro es de un solo dia: no pasa de 24 horas. Al cargar no se controla, como el dia
	// futuro, para no rechazar lo que se guardo antes de esta regla (RN-20, defecto #192)
	if !loading && data.Hours > 24 {
		return nil, ErrEffortHoursOverDay
	}
	// la actividad se controla ya sin los espacios de los extremos, así un texto de solo
	// espacios cuenta como vacío, y se guarda así (RN-6)
	activity := strings.TrimSpace(data.Activity)
	if activity == "" {
		return nil, ErrEffortActivityRequired
	}
	return &Effort{
		storyID:    data.StoryID,
		memberName: name,
		activity:   activity,
		hours:      data.Hours,
		day:        toDay(data.Day),
	}, nil
}

// Efforts devuelve los registros de esfuerzo del proyecto, en el orden en que se cargaron
// (RN-15). Sin registros devuelve una lista vacía, que no es un error. La lista es nueva,
// pero los registros son los guardados.
func (p *Project) Efforts() []*Effort {
	efforts := make([]*Effort, len(p.efforts))
	copy(efforts, p.efforts)
	return efforts
}

// StoryHours devuelve las horas reales de una historia: la suma de todos sus registros, de
// cualquier integrante y día. Sin registros son cero, que no es un error; una historia que
// no es del proyecto informa que no se encontró, para que no pase por "no se trabajó" (RN-13).
func (p *Project) StoryHours(id int) (float64, error) {
	if _, err := p.Story(id); err != nil {
		return 0, err
	}
	total := 0.0
	for _, effort := range p.efforts {
		if effort.storyID == id {
			total += effort.hours
		}
	}
	return total, nil
}

// MemberHours devuelve las horas de un integrante: la suma de todos los registros que cargó
// en el proyecto. El nombre se busca como al registrar, ignorando mayúsculas y los espacios
// de los extremos. Sin registros son cero, que no es un error; alguien que no es integrante
// informa que no se encontró (RN-14).
func (p *Project) MemberHours(name string) (float64, error) {
	member := p.findMember(strings.TrimSpace(name))
	if member == nil {
		return 0, ErrMemberNotFound
	}
	total := 0.0
	for _, effort := range p.efforts {
		if effort.memberName == member.name {
			total += effort.hours
		}
	}
	return total, nil
}
