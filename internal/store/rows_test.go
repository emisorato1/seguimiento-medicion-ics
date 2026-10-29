package store

import (
	"errors"
	"testing"
)

// Sin escenario: un dato que una historia suma a la conversion del dominio y no a la base
// hace fallar el guardado en vez de perderse en silencio.
// Cubre: #177 D5, RN-2
func TestUnDatoQueLaBaseNoConoceNoSeGuarda(t *testing.T) {
	conocido := `{"next_id": 2, "projects": [{"id": 1, "name": "Sistema", "start": "2026-10-01",
		"next_story_id": 1, "members": [], "stories": [], "efforts": []}]}`
	desconocido := `{"next_id": 2, "projects": [{"id": 1, "name": "Sistema", "start": "2026-10-01",
		"next_story_id": 1, "members": [], "stories": [], "efforts": [], "dato_nuevo": 1}]}`

	if _, err := decodeRegistry([]byte(conocido)); err != nil {
		t.Fatalf("los datos que la base conoce tenian que leerse, salio: %v", err)
	}
	if _, err := decodeRegistry([]byte(desconocido)); !errors.Is(err, errUnknownData) {
		t.Errorf("un dato que la base no conoce tenia que dar errUnknownData, salio: %v", err)
	}
}
