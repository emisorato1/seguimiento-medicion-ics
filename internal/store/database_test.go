package store

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// abrirBase abre la base de esa ruta y la cierra al terminar el test: en Windows un archivo
// abierto no se puede borrar, y t.TempDir lo borra al final.
func abrirBase(t *testing.T, path string) *Database {
	t.Helper()
	db, err := OpenDatabase(path)
	if err != nil {
		t.Fatalf("no se pudo abrir la base %s: %v", path, err)
	}
	t.Cleanup(func() { db.Close() })
	return db
}

// versionDelEsquema devuelve el numero de la ultima migracion aplicada a la base.
func versionDelEsquema(t *testing.T, db *Database) int {
	t.Helper()
	var version int
	if err := db.db.QueryRow("PRAGMA user_version").Scan(&version); err != nil {
		t.Fatalf("no se pudo leer la version del esquema: %v", err)
	}
	return version
}

// Escenario: Abrir por primera vez
// Cubre: #177 RN-5, D2, D6
func TestAbrirUnaBaseNuevaCreaElEsquema(t *testing.T) {
	path := filepath.Join(t.TempDir(), "datos-ics.db")

	db := abrirBase(t, path)

	if v := versionDelEsquema(t, db); v != 1 {
		t.Errorf("se esperaba el esquema en la version 1, esta en %d", v)
	}
	rows, err := db.db.Query("SELECT name FROM sqlite_master WHERE type = 'table' ORDER BY name")
	if err != nil {
		t.Fatalf("no se pudieron leer las tablas: %v", err)
	}
	var tablas []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			t.Fatalf("no se pudo leer una tabla: %v", err)
		}
		tablas = append(tablas, name)
	}
	rows.Close()
	quiero := []string{"acceptance_criteria", "defects", "efforts", "members", "projects", "registry",
		"sprint_completed_stories", "sprint_points_at_close", "sprint_stories", "sprints", "stories",
		"votes", "voting_round_participants", "voting_rounds"}
	if !slices.Equal(tablas, quiero) {
		t.Errorf("se esperaban las tablas %v, salieron %v", quiero, tablas)
	}

	// Volver a abrirla no vuelve a crear el esquema ni cambia su version.
	db.Close()
	otra := abrirBase(t, path)
	if v := versionDelEsquema(t, otra); v != 1 {
		t.Errorf("al volver a abrirla se esperaba la version 1, esta en %d", v)
	}
}

// guardarYReabrir guarda el registro en la base de esa ruta, la cierra, la vuelve a abrir y
// carga lo guardado, como al cerrar el sistema y volver a abrirlo.
func guardarYReabrir(t *testing.T, path string, registry *domain.ProjectRegistry) *domain.ProjectRegistry {
	t.Helper()
	db, err := OpenDatabase(path)
	if err != nil {
		t.Fatalf("no se pudo abrir la base: %v", err)
	}
	if err := db.Save(registry); err != nil {
		db.Close()
		t.Fatalf("no se pudo guardar: %v", err)
	}
	db.Close()
	cargado, err := abrirBase(t, path).Load()
	if err != nil {
		t.Fatalf("no se pudo cargar lo guardado: %v", err)
	}
	return cargado
}

// Escenario: Abrir por primera vez
// Cubre: #177 RN-1, RN-3, RN-5
func TestGuardarYCargarUnRegistro(t *testing.T) {
	// Una base nueva, sin guardar nada, se carga como un registro sin proyectos.
	nueva, err := abrirBase(t, filepath.Join(t.TempDir(), "datos-ics.db")).Load()
	if err != nil {
		t.Fatalf("no se pudo cargar una base nueva: %v", err)
	}
	if got, want := comoTexto(t, nueva), comoTexto(t, domain.NewProjectRegistry()); !bytes.Equal(got, want) {
		t.Errorf("una base nueva tenia que cargarse vacia, salio:\n%s", got)
	}

	conUnProyecto := domain.NewProjectRegistry()
	if _, err := conUnProyecto.Create(domain.ProjectData{
		Name: "Sistema de métricas", Description: "Seguimiento del integrador",
		Start: time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC), End: time.Date(2026, 11, 2, 0, 0, 0, 0, time.UTC),
	}); err != nil {
		t.Fatalf("no se pudo crear el proyecto: %v", err)
	}
	casos := []struct {
		nombre   string
		registry *domain.ProjectRegistry
	}{
		{"un registro sin proyectos", domain.NewProjectRegistry()},
		{"un registro con un proyecto", conUnProyecto},
	}
	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			cargado := guardarYReabrir(t, filepath.Join(t.TempDir(), "datos-ics.db"), caso.registry)

			if got, want := comoTexto(t, cargado), comoTexto(t, caso.registry); !bytes.Equal(got, want) {
				t.Errorf("lo cargado no es lo que se guardo.\nGuardado:\n%s\nCargado:\n%s", want, got)
			}
		})
	}
}

// dia arma una fecha sin hora, como las que guarda el sistema.
func dia(y int, m time.Month, d int) time.Time {
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

// sinError corta el test si una operacion del dominio para preparar los datos fallo.
func sinError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("no se pudo preparar el registro de prueba: %v", err)
	}
}

// unRegistroCompleto arma, con las operaciones del dominio, un registro que usa todas las
// entidades y todos sus datos: dos proyectos; integrantes en un orden que no es el
// alfabetico; historias estimadas, sin estimar y con horas estimadas; un sprint cerrado con
// lo completado y uno abierto; esfuerzo; un defecto resuelto y uno abierto; una ronda
// revelada y acordada y otra abierta con un voto.
func unRegistroCompleto(t *testing.T) *domain.ProjectRegistry {
	t.Helper()
	registry := domain.NewProjectRegistry()
	p, err := registry.Create(domain.ProjectData{Name: "Sistema de métricas", Description: "El integrador", Start: dia(2026, 10, 1)})
	sinError(t, err)
	_, err = p.AddMember("Zoe Ríos", domain.RoleAgileEnabler)
	sinError(t, err)
	_, err = p.AddMember("Ana Pérez", domain.RoleProductBuilder)
	sinError(t, err)
	for _, h := range []domain.StoryData{
		{Title: "Registrar esfuerzo", Description: "Cargar horas", Priority: domain.PriorityHigh, AcceptanceCriteria: []string{"Se cargan las horas", "Se ve el total"}},
		{Title: "Cargar defectos", Description: "Registrar errores", Priority: domain.PriorityMedium, AcceptanceCriteria: []string{"Se carga un defecto"}},
		{Title: "Ver el tablero", Description: "Mirar cómo venimos", Priority: domain.PriorityLow, AcceptanceCriteria: []string{"Se ven los gráficos"}},
	} {
		_, err = p.AddStory(h)
		sinError(t, err)
	}
	sinError(t, p.EstimateStory(1, 3))
	sinError(t, p.EstimateHours(1, 6.5))

	_, err = p.AddSprint(domain.SprintData{Goal: "Tener el MVP", Start: dia(2026, 10, 12), End: dia(2026, 10, 23)})
	sinError(t, err)
	sinError(t, p.AssignStory(1))
	sinError(t, p.CloseSprint([]int{1}, dia(2026, 10, 23)))
	_, err = p.AddSprint(domain.SprintData{Goal: "Cerrar la interfaz", Start: dia(2026, 10, 26), End: dia(2026, 11, 6)})
	sinError(t, err)

	hoy := dia(2026, 10, 31)
	_, err = p.LogEffort(domain.EffortData{StoryID: 1, MemberName: "Ana Pérez", Activity: "Escribir tests", Hours: 2.5, Day: dia(2026, 10, 20)}, hoy)
	sinError(t, err)
	_, err = p.LogEffort(domain.EffortData{StoryID: 1, MemberName: "Zoe Ríos", Activity: "Revisar el PR", Hours: 1, Day: dia(2026, 10, 21)}, hoy)
	sinError(t, err)

	_, err = p.AddDefect(domain.DefectData{Title: "Suma mal", Description: "No suma el último", Severity: domain.SeverityMajor, StoryID: 1, DetectedIn: 1})
	sinError(t, err)
	sinError(t, p.ResolveDefect(1, 2))
	_, err = p.AddDefect(domain.DefectData{Title: "Se pierde la severidad", Description: "Al recargar", Severity: domain.SeverityMinor, StoryID: 2, DetectedIn: 2})
	sinError(t, err)

	ronda, err := p.OpenVotingRound(2)
	sinError(t, err)
	sinError(t, ronda.Vote("Zoe Ríos", 3))
	sinError(t, ronda.Vote("Ana Pérez", 8))
	sinError(t, ronda.Reveal())
	sinError(t, ronda.Agree(5))
	otra, err := p.OpenVotingRound(2)
	sinError(t, err)
	sinError(t, otra.Vote("Zoe Ríos", 5))
	sinError(t, p.AssignStory(2))

	_, err = registry.Create(domain.ProjectData{Name: "Otro proyecto", Start: dia(2026, 9, 1), End: dia(2026, 12, 1)})
	sinError(t, err)
	return registry
}

// Escenario: Los datos siguen al cerrar y volver a abrir
// Cubre: #177 RN-3
func TestUnRegistroCompletoVuelveIgual(t *testing.T) {
	registry := unRegistroCompleto(t)

	cargado := guardarYReabrir(t, filepath.Join(t.TempDir(), "datos-ics.db"), registry)

	// Se compara lo que da la conversion del dominio: si un dato, su orden o un contador no
	// volvio igual, el texto es distinto.
	if got, want := comoTexto(t, cargado), comoTexto(t, registry); !bytes.Equal(got, want) {
		t.Errorf("lo cargado no es lo que se guardo.\nGuardado:\n%s\nCargado:\n%s", want, got)
	}

	// La numeracion sigue donde habia quedado.
	p, err := cargado.Get(1)
	sinError(t, err)
	historia, err := p.AddStory(domain.StoryData{Title: "Otra", Description: "Una más", Priority: domain.PriorityLow, AcceptanceCriteria: []string{"Algo"}})
	sinError(t, err)
	defecto, err := p.AddDefect(domain.DefectData{Title: "Otro", Description: "Uno más", Severity: domain.SeverityMinor, StoryID: 1, DetectedIn: 2})
	sinError(t, err)
	proyecto, err := cargado.Create(domain.ProjectData{Name: "Tercero", Start: dia(2026, 10, 1)})
	sinError(t, err)
	if historia.ID() != 4 || defecto.ID() != 3 || proyecto.ID() != 3 {
		t.Errorf("se esperaban la historia 4, el defecto 3 y el proyecto 3, salieron %d, %d y %d", historia.ID(), defecto.ID(), proyecto.ID())
	}
}

// unProyecto arma un registro con un proyecto y le aplica cambios con las operaciones del
// dominio.
func unProyecto(t *testing.T, cambios func(p *domain.Project)) *domain.ProjectRegistry {
	t.Helper()
	registry := domain.NewProjectRegistry()
	p, err := registry.Create(domain.ProjectData{Name: "Sistema de métricas", Start: dia(2026, 10, 1)})
	sinError(t, err)
	cambios(p)
	return registry
}

// Esquema del escenario: Los datos difíciles vuelven iguales
// Cubre: #177 RN-3, RN-6
func TestLosDatosDificilesVuelvenIguales(t *testing.T) {
	hoy := dia(2026, 10, 31)
	conHistoria := func(p *domain.Project, criterios ...string) {
		_, err := p.AddStory(domain.StoryData{Title: "Registrar esfuerzo", Description: "Cargar horas",
			Priority: domain.PriorityHigh, AcceptanceCriteria: criterios})
		sinError(t, err)
	}
	casos := []struct {
		nombre string
		armar  func(p *domain.Project)
		// despues mira lo que no se ve en la conversion, si hace falta
		despues func(t *testing.T, cargado *domain.ProjectRegistry)
	}{
		{nombre: "un criterio con signos y en dos renglones", armar: func(p *domain.Project) {
			conHistoria(p, "Se ve <b>el total</b> y \"las horas\"\nen dos renglones", "Ñandú, acción & 'comillas'")
		}},
		{nombre: "horas con decimales", armar: func(p *domain.Project) {
			conHistoria(p, "Se cargan")
			_, err := p.AddMember("Ana Pérez", domain.RoleProductBuilder)
			sinError(t, err)
			for _, horas := range []float64{2.5, 0.1} {
				_, err := p.LogEffort(domain.EffortData{StoryID: 1, MemberName: "Ana Pérez", Activity: "Tests", Hours: horas, Day: dia(2026, 10, 27)}, hoy)
				sinError(t, err)
			}
		}},
		{nombre: "una ronda sin revelar con un voto", armar: func(p *domain.Project) {
			conHistoria(p, "Se vota")
			for _, nombre := range []string{"Ana Pérez", "Luis Gómez"} {
				_, err := p.AddMember(nombre, domain.RoleProductBuilder)
				sinError(t, err)
			}
			ronda, err := p.OpenVotingRound(1)
			sinError(t, err)
			sinError(t, ronda.Vote("Ana Pérez", 13))
		}, despues: func(t *testing.T, cargado *domain.ProjectRegistry) {
			p, err := cargado.Get(1)
			sinError(t, err)
			ronda, _, err := p.CurrentVotingRound(1)
			sinError(t, err)
			if _, err := ronda.Votes(); !errors.Is(err, domain.ErrVotesHidden) {
				t.Errorf("los votos de una ronda sin revelar tenian que seguir ocultos, salio: %v", err)
			}
		}},
		{nombre: "un integrante quitado con horas cargadas", armar: func(p *domain.Project) {
			conHistoria(p, "Se cargan")
			_, err := p.AddMember("Ana Pérez", domain.RoleProductBuilder)
			sinError(t, err)
			_, err = p.LogEffort(domain.EffortData{StoryID: 1, MemberName: "Ana Pérez", Activity: "Tests", Hours: 2, Day: dia(2026, 10, 27)}, hoy)
			sinError(t, err)
			sinError(t, p.RemoveMember("Ana Pérez"))
		}},
		{nombre: "un proyecto al que se le quito el unico integrante", armar: func(p *domain.Project) {
			_, err := p.AddMember("Ana Pérez", domain.RoleProductBuilder)
			sinError(t, err)
			sinError(t, p.RemoveMember("Ana Pérez"))
		}},
		{nombre: "un proyecto sin nada", armar: func(p *domain.Project) {}},
	}
	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			registry := unProyecto(t, caso.armar)

			cargado := guardarYReabrir(t, filepath.Join(t.TempDir(), "datos-ics.db"), registry)

			if got, want := comoTexto(t, cargado), comoTexto(t, registry); !bytes.Equal(got, want) {
				t.Errorf("lo cargado no es lo que se guardo.\nGuardado:\n%s\nCargado:\n%s", want, got)
			}
			if caso.despues != nil {
				caso.despues(t, cargado)
			}
		})
	}
}

// Escenario: Un cambio que no se pudo guardar no deja nada a medias
// Cubre: #177 RN-2, D4, D9
func TestUnGuardadoQueFallaNoDejaNadaAMedias(t *testing.T) {
	path := filepath.Join(t.TempDir(), "datos-ics.db")
	antes := unProyecto(t, func(p *domain.Project) {})
	despues := unRegistroCompleto(t)
	db := abrirBase(t, path)
	sinError(t, db.Save(antes))

	// La base rechaza los votos, que se insertan al final: el guardado se corta despues de
	// haber vaciado las tablas y escrito casi todo.
	if _, err := db.db.Exec("CREATE TRIGGER falla BEFORE INSERT ON votes BEGIN SELECT RAISE(ABORT, 'falla'); END"); err != nil {
		t.Fatalf("no se pudo preparar la falla: %v", err)
	}
	err := db.Save(despues)

	if !errors.Is(err, ErrNotSaved) {
		t.Fatalf("se esperaba ErrNotSaved, salio: %v", err)
	}
	if !strings.Contains(err.Error(), path) {
		t.Errorf("el mensaje tenia que nombrar la base, salio: %v", err)
	}
	cargado, err := db.Load()
	sinError(t, err)
	if got, want := comoTexto(t, cargado), comoTexto(t, antes); !bytes.Equal(got, want) {
		t.Errorf("la base tenia que quedar como antes del guardado que fallo, y tiene:\n%s", got)
	}

	// El siguiente guardado que funciona guarda todo.
	if _, err := db.db.Exec("DROP TRIGGER falla"); err != nil {
		t.Fatalf("no se pudo sacar la falla: %v", err)
	}
	sinError(t, db.Save(despues))
	cargado, err = db.Load()
	sinError(t, err)
	if got, want := comoTexto(t, cargado), comoTexto(t, despues); !bytes.Equal(got, want) {
		t.Errorf("el guardado siguiente tenia que guardar todo, y la base tiene:\n%s", got)
	}
}

// Escenarios: Un dato guardado que no cumple las reglas no se carga / Una base de una
// versión más nueva no se abre
// Cubre: #177 RN-4, RN-13
func TestLoQueNoSePuedeAbrirSeAvisaYNoSeToca(t *testing.T) {
	casos := []struct {
		nombre  string
		armar   func(t *testing.T, path string)
		mensaje string
	}{
		{nombre: "un archivo que no es una base", armar: func(t *testing.T, path string) {
			if err := os.WriteFile(path, []byte("no soy una base"), 0o644); err != nil {
				t.Fatalf("no se pudo preparar el archivo: %v", err)
			}
		}, mensaje: "no se pudo abrir la base"},
		{nombre: "una historia con 4 story points", armar: func(t *testing.T, path string) {
			db, err := OpenDatabase(path)
			sinError(t, err)
			sinError(t, db.Save(unRegistroCompleto(t)))
			_, err = db.db.Exec("UPDATE stories SET story_points = 4 WHERE id = 1")
			sinError(t, err)
			db.Close()
		}, mensaje: domain.ErrStoryPointsInvalid.Error()},
		{nombre: "una base de una version mas nueva", armar: func(t *testing.T, path string) {
			db, err := OpenDatabase(path)
			sinError(t, err)
			_, err = db.db.Exec("PRAGMA user_version = 99")
			sinError(t, err)
			db.Close()
		}, mensaje: "versión más nueva"},
	}
	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "datos-ics.db")
			caso.armar(t, path)
			antes, err := os.ReadFile(path)
			sinError(t, err)

			db, err := OpenDatabase(path)
			if err == nil {
				_, err = db.Load()
				db.Close()
			}

			if !errors.Is(err, domain.ErrSavedDataInvalid) {
				t.Fatalf("se esperaba ErrSavedDataInvalid, salio: %v", err)
			}
			if !strings.Contains(err.Error(), caso.mensaje) || !strings.Contains(err.Error(), path) {
				t.Errorf("el mensaje tenia que decir %q y nombrar la base, salio: %v", caso.mensaje, err)
			}
			if despues, _ := os.ReadFile(path); !bytes.Equal(despues, antes) {
				t.Error("la base tenia que quedar como estaba")
			}
		})
	}
}
