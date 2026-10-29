package store

import (
	"bytes"
	"encoding/json"
	"errors"
)

// Los tipos que siguen son la forma del archivo de la #77, la que entrega la conversion del
// dominio (internal/domain/registry_json.go), con las mismas etiquetas. Son el puente entre
// el dominio y las tablas (historia #177, D3): al guardar se desarman en filas y al cargar se
// rearman y se le pasan al dominio, que valida. Nunca se guardan como texto en la base.

type fileRegistry struct {
	NextID   int           `json:"next_id"`
	Projects []fileProject `json:"projects"`
}

type fileProject struct {
	ID               int               `json:"id"`
	Name             string            `json:"name"`
	Description      string            `json:"description,omitempty"`
	Start            string            `json:"start"`
	End              string            `json:"end,omitempty"`
	NextStoryID      int               `json:"next_story_id"`
	Members          []fileMember      `json:"members"`
	Stories          []fileStory       `json:"stories"`
	Efforts          []fileEffort      `json:"efforts"`
	LastSprintNumber int               `json:"last_sprint_number,omitempty"`
	Sprints          []fileSprint      `json:"sprints,omitempty"`
	LastDefectID     int               `json:"last_defect_id,omitempty"`
	Defects          []fileDefect      `json:"defects,omitempty"`
	VotingRounds     []fileVotingRound `json:"voting_rounds,omitempty"`
}

type fileMember struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type fileStory struct {
	ID                 int      `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Priority           string   `json:"priority"`
	State              string   `json:"state"`
	StoryPoints        int      `json:"story_points"`
	EstimatedHours     float64  `json:"estimated_hours,omitempty"`
	AcceptanceCriteria []string `json:"acceptance_criteria"`
}

type fileEffort struct {
	StoryID    int     `json:"story_id"`
	MemberName string  `json:"member_name"`
	Activity   string  `json:"activity"`
	Hours      float64 `json:"hours"`
	Day        string  `json:"day"`
}

type fileSprint struct {
	Number           int         `json:"number"`
	Goal             string      `json:"goal"`
	Start            string      `json:"start"`
	End              string      `json:"end"`
	State            string      `json:"state"`
	Stories          []int       `json:"stories,omitempty"`
	ClosedOn         string      `json:"closed_on,omitempty"`
	CompletedStories []int       `json:"completed_stories,omitempty"`
	PointsAtClose    map[int]int `json:"points_at_close,omitempty"`
}

type fileDefect struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Severity    string `json:"severity"`
	State       string `json:"state"`
	StoryID     int    `json:"story_id"`
	DetectedIn  int    `json:"detected_in"`
	ResolvedIn  int    `json:"resolved_in,omitempty"`
}

type fileVotingRound struct {
	StoryID      int        `json:"story_id"`
	Number       int        `json:"number"`
	Participants []string   `json:"participants"`
	Votes        []fileVote `json:"votes,omitempty"`
	Revealed     bool       `json:"revealed,omitempty"`
	Agreed       int        `json:"agreed,omitempty"`
}

type fileVote struct {
	MemberName string `json:"member_name"`
	Points     int    `json:"points"`
}

// errUnknownData es un dato de la conversion que la base no tiene donde guardar (D5).
var errUnknownData = errors.New("la conversión trae un dato que la base no sabe guardar")

// decodeRegistry lee lo que entrega la conversion del dominio. Un dato que estos tipos no
// conocen es un error y no se ignora (D5): si una historia suma un dato a la conversion y no
// a la base, el guardado falla en vez de perderlo en silencio.
func decodeRegistry(data []byte) (fileRegistry, error) {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	var registry fileRegistry
	if err := decoder.Decode(&registry); err != nil {
		return fileRegistry{}, errUnknownData
	}
	return registry, nil
}
