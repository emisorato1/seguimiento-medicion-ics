package domain

import "strings"

// Severity es qué tan grave es un defecto (historia #15). Es un tipo propio para que no se
// pueda guardar una severidad que no existe. El valor cero no es ninguna: significa "sin
// severidad", así un defecto armado sin indicarla no queda como Crítico sin que nadie se
// entere (Fase 2), igual que el rol de la #4.
type Severity int

const (
	// SeverityCritical, SeverityMajor y SeverityMinor son las tres severidades válidas.
	// Arrancan en 1 para que el valor cero quede libre como "sin severidad".
	SeverityCritical Severity = iota + 1
	SeverityMajor
	SeverityMinor
)

// String devuelve el nombre de la severidad tal como la nombra la Fase 1, o "sin severidad"
// si no es ninguna de las tres.
func (s Severity) String() string {
	switch s {
	case SeverityCritical:
		return "Crítica"
	case SeverityMajor:
		return "Mayor"
	case SeverityMinor:
		return "Menor"
	default:
		return "sin severidad"
	}
}

// isValid indica si la severidad es una de las tres válidas. Un valor armado a mano, como
// Severity(9), no lo es (RN-6).
func (s Severity) isValid() bool {
	return s == SeverityCritical || s == SeverityMajor || s == SeverityMinor
}

// ParseSeverity convierte el texto que escribe el usuario en una severidad, sin distinguir
// mayúsculas ni los espacios de los extremos, y acepta "critica" sin tilde, que es como se
// escribe en un teclado sin tildes (Fase 2). Cualquier otro texto se rechaza con el error de
// la RN-6. Vive en el dominio y no en la interfaz, para que la regla no quede duplicada.
func ParseSeverity(text string) (Severity, error) {
	switch strings.ToLower(strings.TrimSpace(text)) {
	case "crítica", "critica":
		return SeverityCritical, nil
	case "mayor":
		return SeverityMajor, nil
	case "menor":
		return SeverityMinor, nil
	default:
		return 0, ErrDefectSeverityInvalid
	}
}

// DefectState es en qué anda un defecto. A diferencia de la severidad, no tiene un valor
// "ninguno": el defecto siempre nace Abierto (RN-3).
type DefectState int

const (
	// DefectOpen: se detectó y todavía nadie lo está arreglando.
	DefectOpen DefectState = iota
	// DefectInProgress: alguien lo está arreglando.
	DefectInProgress
	// DefectResolved: se arregló, y quedó registrado en qué sprint. Solo se llega con
	// ResolveDefect.
	DefectResolved
)

// Defect es un defecto encontrado en el producto. Vive dentro del proyecto, como las
// historias: así su identificador es único dentro del proyecto sin llevar un índice aparte
// (RN-1, RN-16). Sus datos están ocultos fuera del paquete: solo se cambian con operaciones
// que validan.
type Defect struct {
	id          int
	title       string
	description string
	severity    Severity
	state       DefectState
	storyID     int
	detectedIn  int
	resolvedIn  int
}

// ID devuelve el identificador del defecto dentro de su proyecto, que no cambia nunca.
func (d *Defect) ID() int {
	return d.id
}

// Title devuelve cómo se lo nombra, en una frase.
func (d *Defect) Title() string {
	return d.title
}

// Description devuelve qué pasa.
func (d *Defect) Description() string {
	return d.description
}

// Severity devuelve qué tan grave es.
func (d *Defect) Severity() Severity {
	return d.severity
}

// State devuelve en qué anda.
func (d *Defect) State() DefectState {
	return d.state
}

// StoryID devuelve el número de la historia sobre la que apareció.
func (d *Defect) StoryID() int {
	return d.storyID
}

// DetectedIn devuelve el número del sprint en que se lo encontró.
func (d *Defect) DetectedIn() int {
	return d.detectedIn
}

// ResolvedIn devuelve el número del sprint en que se resolvió, o 0 si no está resuelto.
func (d *Defect) ResolvedIn() int {
	return d.resolvedIn
}

// DefectData son los datos con los que se registra un defecto. El identificador no está: lo
// asigna el sistema (RN-1).
type DefectData struct {
	Title       string
	Description string
	Severity    Severity
	StoryID     int
	DetectedIn  int
}

// AddDefect registra un defecto en el proyecto con el identificador que sigue al del último
// registrado, en estado Abierto y sin sprint de resolución (RN-1, RN-3).
func (p *Project) AddDefect(data DefectData) (*Defect, error) {
	// primero se busca lo que se nombra (Fase 2): la historia tiene que ser del proyecto, en
	// cualquier estado (RN-7)
	if _, err := p.Story(data.StoryID); err != nil {
		return nil, err
	}
	// el sprint de detección tiene que existir en el proyecto, abierto o cerrado (RN-8)
	if _, err := p.Sprint(data.DetectedIn); err != nil {
		return nil, err
	}
	// el título y la descripción se controlan ya sin los espacios de los extremos, así un
	// texto de solo espacios cuenta como vacío, y se guardan así (RN-4, RN-5)
	title := strings.TrimSpace(data.Title)
	if title == "" {
		return nil, ErrDefectTitleRequired
	}
	description := strings.TrimSpace(data.Description)
	if description == "" {
		return nil, ErrDefectDescriptionRequired
	}
	if !data.Severity.isValid() {
		return nil, ErrDefectSeverityInvalid
	}
	defect := &Defect{
		id:          p.lastDefectID + 1,
		title:       title,
		description: description,
		severity:    data.Severity,
		state:       DefectOpen,
		storyID:     data.StoryID,
		detectedIn:  data.DetectedIn,
	}
	p.defects = append(p.defects, defect)
	p.lastDefectID = defect.id
	return defect, nil
}

// Defects devuelve todos los defectos del proyecto, en el orden en que se registraron
// (RN-15). Sin defectos devuelve una lista vacía, que no es un error. La lista es nueva, pero
// los defectos son los guardados.
func (p *Project) Defects() []*Defect {
	defects := make([]*Defect, len(p.defects))
	copy(defects, p.defects)
	return defects
}

// ChangeDefectState pasa el defecto con ese identificador al estado indicado: entre Abierto y
// En curso, en los dos sentidos e incluso al que ya tenía (RN-9).
func (p *Project) ChangeDefectState(id int, state DefectState) error {
	// a Resuelto solo se llega resolviendo, porque hace falta el sprint (RN-10)
	if state == DefectResolved {
		return ErrDefectResolveNeedsSprint
	}
	if state != DefectOpen && state != DefectInProgress {
		return ErrDefectStateInvalid
	}
	for _, defect := range p.defects {
		if defect.id == id {
			defect.state = state
		}
	}
	return nil
}

// ResolveDefect pasa el defecto con ese identificador a Resuelto y registra el sprint en que
// se resolvió, esté Abierto o En curso (RN-11).
func (p *Project) ResolveDefect(id, sprintNumber int) error {
	for _, defect := range p.defects {
		if defect.id == id {
			defect.state = DefectResolved
			defect.resolvedIn = sprintNumber
		}
	}
	return nil
}
