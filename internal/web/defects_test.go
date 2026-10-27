package web

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// defectRegistry arma el proyecto 1 con la historia 1, el sprint 1 cerrado y el sprint 2
// abierto, que es lo que piden los escenarios de defectos.
func defectRegistry(t *testing.T) (*domain.ProjectRegistry, *domain.Project) {
	t.Helper()
	registry := domain.NewProjectRegistry()
	project := addProject(t, registry, "Sistema de metricas")
	addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
	addSprint(t, registry, 1)
	if err := project.CloseSprint(nil, time.Date(2026, 10, 30, 0, 0, 0, 0, time.UTC)); err != nil {
		t.Fatalf("no se pudo cerrar el sprint 1: %v", err)
	}
	if _, err := project.AddSprint(domain.SprintData{
		Goal:  "Cerrar la interfaz",
		Start: time.Date(2026, 11, 2, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2026, 11, 13, 0, 0, 0, 0, time.UTC),
	}); err != nil {
		t.Fatalf("no se pudo preparar el sprint 2: %v", err)
	}
	return registry, project
}

// addDefect carga en el proyecto el defecto "Suma mal las horas", de severidad Mayor, sobre la
// historia 1 y detectado en ese sprint.
func addDefect(t *testing.T, project *domain.Project, detectedIn int) *domain.Defect {
	t.Helper()
	defect, err := project.AddDefect(domain.DefectData{Title: "Suma mal las horas", Description: "No suma el ultimo",
		Severity: domain.SeverityMajor, StoryID: 1, DetectedIn: detectedIn})
	if err != nil {
		t.Fatalf("no se pudo preparar el defecto: %v", err)
	}
	return defect
}

// Escenario: Cargar un defecto, pasarlo a en curso y resolverlo
// Cubre: RN-1, RN-2, RN-16
func TestCargarUnDefectoPasarloAEnCursoYResolverlo(t *testing.T) {
	registry, project := defectRegistry(t)
	saver := &fakeSaver{}
	navegador := newBrowser(t, NewServer(registry, saver, testToday))

	// Lo que tiene valores fijos o sale de lo cargado se elige de una lista (RN-2, RN-16).
	html := navegador.get("/proyectos/1/defectos/nuevo").Body.String()
	for _, lista := range []struct {
		nombre string
		quiero []string
	}{
		{"severidad", []string{"Crítica", "Mayor", "Menor"}},
		{"historia", []string{"1"}},
		{"sprint", []string{"1", "2"}},
	} {
		if got := options(html, lista.nombre); !slices.Equal(got, lista.quiero) {
			t.Errorf("se esperaba elegir %s entre %v, salio %v", lista.nombre, lista.quiero, got)
		}
	}

	res := navegador.post("/proyectos/1/defectos/nuevo", url.Values{
		"titulo":      {"Suma mal las horas"},
		"descripcion": {"El total por historia no suma el ultimo registro"},
		"severidad":   {"Mayor"},
		"historia":    {"1"},
		"sprint":      {"1"},
	})
	if res.Code != http.StatusSeeOther || res.Header().Get("Location") != "/proyectos/1/defectos" {
		t.Fatalf("se esperaba volver a la lista de defectos con el estado %d, salio %d hacia %q:\n%s", http.StatusSeeOther, res.Code, res.Header().Get("Location"), res.Body.String())
	}
	html = joinTags(between(navegador.get("/proyectos/1/defectos").Body.String(), "<main", "</main>"))
	if !strings.Contains(html, "Se cargó el defecto 1") {
		t.Errorf("se esperaba el mensaje %q, y no esta en:\n%s", "Se cargó el defecto 1", html)
	}
	fila := between(html, "<tbody>", "</tbody>")
	for _, texto := range []string{"<td>1</td>", "Suma mal las horas", "<td>Mayor</td>", "<td>Abierto</td>"} {
		if !strings.Contains(fila, texto) {
			t.Errorf("se esperaba en la lista de defectos %q, y no esta en:\n%s", texto, fila)
		}
	}

	// Pasarlo a En curso se elige de la lista de estados, en la pagina del defecto.
	html = navegador.get("/proyectos/1/defectos/1").Body.String()
	if got := options(html, "estado"); !slices.Equal(got, []string{"Abierto", "En curso"}) {
		t.Errorf("se esperaba elegir el estado entre Abierto y En curso, salio %v", got)
	}
	res = navegador.post("/proyectos/1/defectos/1/estado", url.Values{"estado": {"En curso"}})
	if res.Code != http.StatusSeeOther || res.Header().Get("Location") != "/proyectos/1/defectos/1" {
		t.Fatalf("se esperaba volver al defecto 1 con el estado %d, salio %d hacia %q:\n%s", http.StatusSeeOther, res.Code, res.Header().Get("Location"), res.Body.String())
	}
	html = joinTags(between(navegador.get("/proyectos/1/defectos/1").Body.String(), "<main", "</main>"))
	if !strings.Contains(html, "El defecto 1 pasó a En curso") || !strings.Contains(html, "<dt>Estado</dt><dd>En curso</dd>") {
		t.Errorf("se esperaba el defecto 1 En curso, con su mensaje, y no esta en:\n%s", html)
	}
	if n := len(project.DefectsByState(domain.DefectInProgress)); n != 1 || saver.saves != 2 {
		t.Errorf("se esperaba un defecto En curso y dos guardados, hay %d y %d", n, saver.saves)
	}

	// Resolverlo: el sprint se elige de la lista y se pide confirmar, porque no se deshace.
	html = navegador.get("/proyectos/1/defectos/1").Body.String()
	if got := options(html, "sprint"); !slices.Equal(got, []string{"1", "2"}) {
		t.Errorf("se esperaba elegir el sprint de resolucion entre 1 y 2, salio %v", got)
	}
	html = navegador.get("/proyectos/1/defectos/1/resolver?sprint=2").Body.String()
	if !strings.Contains(html, "¿Resolver el defecto 1 en el sprint 2?") {
		t.Fatalf("se esperaba la pregunta para confirmar, y no esta en:\n%s", html)
	}
	res = navegador.post("/proyectos/1/defectos/1/resolver", url.Values{"sprint": {"2"}})
	if res.Code != http.StatusSeeOther || res.Header().Get("Location") != "/proyectos/1/defectos/1" {
		t.Fatalf("se esperaba volver al defecto 1 con el estado %d, salio %d hacia %q:\n%s", http.StatusSeeOther, res.Code, res.Header().Get("Location"), res.Body.String())
	}
	html = joinTags(between(navegador.get("/proyectos/1/defectos/1").Body.String(), "<main", "</main>"))
	for _, texto := range []string{"Se resolvió el defecto 1 en el sprint 2", "<dt>Estado</dt><dd>Resuelto</dd>", "<dt>Resuelto en</dt><dd>Sprint 2</dd>"} {
		if !strings.Contains(html, texto) {
			t.Errorf("se esperaba en el defecto resuelto %q, y no esta en:\n%s", texto, html)
		}
	}
	// Un resuelto ya no ofrece cambiar su estado ni resolverlo.
	for _, accion := range []string{`action="/proyectos/1/defectos/1/estado"`, `action="/proyectos/1/defectos/1/resolver"`} {
		if strings.Contains(html, accion) {
			t.Errorf("un defecto resuelto no deberia ofrecer %s:\n%s", accion, html)
		}
	}
}

// Esquema del escenario: Confirmar o cancelar antes de lo que no se deshace (resolver)
// Cubre: RN-9
func TestConfirmarOCancelarAntesDeResolverUnDefecto(t *testing.T) {
	for _, confirmo := range []bool{true, false} {
		nombre := map[bool]string{true: "confirmo", false: "cancelo"}[confirmo]
		t.Run(nombre, func(t *testing.T) {
			// Dado el sprint 1 abierto y el defecto 1 Abierto.
			registry := domain.NewProjectRegistry()
			project := addProject(t, registry, "Sistema de metricas")
			addStory(t, project, "Registrar esfuerzo", domain.PriorityHigh)
			addSprint(t, registry, 1)
			if _, err := project.AddDefect(domain.DefectData{Title: "Suma mal las horas", Description: "No suma el ultimo",
				Severity: domain.SeverityMajor, StoryID: 1, DetectedIn: 1}); err != nil {
				t.Fatalf("no se pudo preparar el defecto: %v", err)
			}
			saver := &fakeSaver{}
			navegador := newBrowser(t, NewServer(registry, saver, testToday))
			antes := snapshot(registry)

			// Pedir resolverlo no resuelve nada: muestra la confirmacion.
			res := navegador.get("/proyectos/1/defectos/1/resolver?sprint=1")
			if res.Code != http.StatusOK {
				t.Fatalf("se esperaba la confirmacion con el estado %d, salio %d", http.StatusOK, res.Code)
			}
			confirmacion := between(res.Body.String(), "<main", "</main>")
			for _, texto := range []string{`action="/proyectos/1/defectos/1/resolver"`, `href="/proyectos/1/defectos/1"`, `name="sprint" value="1"`} {
				if !strings.Contains(confirmacion, texto) {
					t.Errorf("se esperaba en la confirmacion %s, y no esta en:\n%s", texto, confirmacion)
				}
			}
			if despues := snapshot(registry); despues != antes {
				t.Fatalf("pedir que se resuelva no deberia cambiar nada, y el registro paso de:\n%s\na:\n%s", antes, despues)
			}

			if confirmo {
				navegador.post("/proyectos/1/defectos/1/resolver", url.Values{"sprint": {"1"}})
			} else {
				navegador.get("/proyectos/1/defectos/1")
			}

			quiero := map[bool]string{true: "<td>Resuelto</td>", false: "<td>Abierto</td>"}[confirmo]
			if html := navegador.get("/proyectos/1/defectos").Body.String(); !strings.Contains(html, quiero) {
				t.Errorf("se esperaba el defecto 1 con %s, y no esta en:\n%s", quiero, html)
			}
			if !confirmo && (snapshot(registry) != antes || saver.saves != 0) {
				t.Errorf("cancelar no deberia cambiar nada")
			}
		})
	}
}

// Esquema del escenario: Filtrar los defectos
// Cubre: RN-13, D7
func TestFiltrarLosDefectos(t *testing.T) {
	// Dado el defecto 1 Abierto de severidad Critica y el defecto 2 En curso de severidad Menor.
	registry, project := defectRegistry(t)
	for _, d := range []domain.DefectData{
		{Title: "Suma mal las horas", Description: "No suma el ultimo", Severity: domain.SeverityCritical, StoryID: 1, DetectedIn: 1},
		{Title: "No guarda el voto", Description: "Se pierde al recargar", Severity: domain.SeverityMinor, StoryID: 1, DetectedIn: 1},
	} {
		if _, err := project.AddDefect(d); err != nil {
			t.Fatalf("no se pudo preparar el defecto: %v", err)
		}
	}
	if err := project.ChangeDefectState(2, domain.DefectInProgress); err != nil {
		t.Fatalf("no se pudo pasar el defecto 2 a En curso: %v", err)
	}
	server := newTestServer(t, registry)

	// Los filtros se eligen de una lista, con los nombres del sistema.
	html := get(t, server, "/proyectos/1/defectos").Body.String()
	if got := options(html, "estado"); !slices.Equal(got, []string{"Abierto", "En curso", "Resuelto"}) {
		t.Errorf("se esperaba filtrar por los estados Abierto, En curso y Resuelto, salio %v", got)
	}
	if got := options(html, "severidad"); !slices.Equal(got, []string{"Crítica", "Mayor", "Menor"}) {
		t.Errorf("se esperaba filtrar por las severidades Crítica, Mayor y Menor, salio %v", got)
	}

	casos := []struct {
		filtro  string
		pagina  string
		veo     []string
		noVeo   []string
		mensaje string // el rechazo arriba de la lista, si el filtro no es de la lista
	}{
		{"el estado Abierto", "/proyectos/1/defectos?estado=Abierto", []string{"Suma mal las horas"}, []string{"No guarda el voto"}, ""},
		{"la severidad Menor", "/proyectos/1/defectos?severidad=Menor", []string{"No guarda el voto"}, []string{"Suma mal las horas"}, ""},
		{"nada", "/proyectos/1/defectos", []string{"Suma mal las horas", "No guarda el voto"}, nil, ""},
		// No estan en los Ejemplos: un filtro que no es de la lista, escrito a mano en la
		// direccion, lo rechaza el nucleo, y se ven todos con su mensaje arriba.
		{"un estado que no existe", "/proyectos/1/defectos?estado=Cerrado", []string{"Suma mal las horas", "No guarda el voto"}, nil, domain.ErrDefectStateUnknown.Error()},
		{"una severidad que no existe", "/proyectos/1/defectos?severidad=Leve", []string{"Suma mal las horas", "No guarda el voto"}, nil, domain.ErrDefectSeverityInvalid.Error()},
	}
	for _, caso := range casos {
		t.Run(caso.filtro, func(t *testing.T) {
			res := get(t, server, caso.pagina)
			if res.Code != http.StatusOK {
				t.Fatalf("se esperaba el estado %d, salio %d", http.StatusOK, res.Code)
			}
			if caso.mensaje != "" && !strings.Contains(res.Body.String(), `id="error-formulario">`+caso.mensaje) {
				t.Errorf("se esperaba arriba el mensaje %q, y no esta en:\n%s", caso.mensaje, res.Body.String())
			}
			filas := between(res.Body.String(), "<tbody>", "</tbody>")
			for _, titulo := range caso.veo {
				if !strings.Contains(filas, titulo) {
					t.Errorf("se esperaba ver %q, y no esta en:\n%s", titulo, filas)
				}
			}
			for _, titulo := range caso.noVeo {
				if strings.Contains(filas, titulo) {
					t.Errorf("no se esperaba ver %q, y esta en:\n%s", titulo, filas)
				}
			}
		})
	}
}

// Sin escenario: pedir resolver sin elegir el sprint no lleva a confirmar nada: vuelve a la
// pagina del defecto con el error junto al sprint (RN-2).
// Cubre: RN-2, RN-9
func TestResolverSinElegirElSprint(t *testing.T) {
	registry, project := defectRegistry(t)
	addDefect(t, project, 1)
	server := newTestServer(t, registry)

	res := get(t, server, "/proyectos/1/defectos/1/resolver?sprint=")

	if res.Code != http.StatusUnprocessableEntity {
		t.Fatalf("se esperaba el estado %d, salio %d", http.StatusUnprocessableEntity, res.Code)
	}
	html := res.Body.String()
	if !strings.Contains(html, `id="sprint-error">`+errInvalidSprint.Error()) || strings.Contains(html, "¿Resolver") {
		t.Errorf("se esperaba el error junto al sprint y ninguna pregunta para confirmar:\n%s", html)
	}
}

// Defecto #190 (QA-3): con una pagina vieja de un defecto que ya se resolvio en otra pestaña,
// cambiarle el estado o resolverlo otra vez muestra el mensaje del nucleo, aunque la pagina ya
// no tenga esos formularios.
// Cubre: #78 RN-2; #15 RN-14
func TestUnDefectoYaResueltoDesdeUnaPaginaViejaMuestraElMensaje(t *testing.T) {
	registry, project := defectRegistry(t)
	addDefect(t, project, 1)
	if err := project.ResolveDefect(1, 2); err != nil {
		t.Fatalf("no se pudo resolver el defecto: %v", err)
	}
	navegador := newBrowser(t, newTestServer(t, registry))

	for pedido, res := range map[string]*httptest.ResponseRecorder{
		"cambiar el estado": navegador.post("/proyectos/1/defectos/1/estado", url.Values{"estado": {"En curso"}}),
		"resolver":          navegador.post("/proyectos/1/defectos/1/resolver", url.Values{"sprint": {"2"}}),
	} {
		if res.Code != http.StatusUnprocessableEntity || !strings.Contains(res.Body.String(), domain.ErrDefectAlreadyResolved.Error()) {
			t.Errorf("%s: se esperaba el estado %d y el mensaje %q, salio %d:\n%s", pedido, http.StatusUnprocessableEntity, domain.ErrDefectAlreadyResolved, res.Code, res.Body.String())
		}
	}
}

// Defecto #190 (QA-10): pedir resolver en un sprint que no existe avisa antes de preguntar,
// junto al sprint, en la pagina del defecto.
// Cubre: #78 RN-2, RN-9; #15 RN-12
func TestResolverEnUnSprintQueNoExisteAvisaAntesDePreguntar(t *testing.T) {
	registry, project := defectRegistry(t)
	addDefect(t, project, 1)
	server := newTestServer(t, registry)

	res := get(t, server, "/proyectos/1/defectos/1/resolver?sprint=99")

	html := res.Body.String()
	if res.Code != http.StatusUnprocessableEntity || strings.Contains(html, "¿Resolver") {
		t.Errorf("se esperaba el estado %d y ninguna pregunta para confirmar, salio %d:\n%s", http.StatusUnprocessableEntity, res.Code, html)
	}
	if !strings.Contains(html, `id="sprint-error">`+domain.ErrSprintNotFound.Error()) {
		t.Errorf("se esperaba el mensaje %q junto al sprint:\n%s", domain.ErrSprintNotFound, html)
	}
}

// Defecto #193: un filtro escrito en minusculas en la direccion filtra bien y ademas queda
// elegido en su lista, escrito como lo muestra la lista.
// Cubre: #78 RN-13, D7
func TestUnFiltroEnMinusculasQuedaElegido(t *testing.T) {
	registry, project := defectRegistry(t)
	addDefect(t, project, 1)
	server := newTestServer(t, registry)

	for direccion, elegido := range map[string]string{
		"/proyectos/1/defectos?estado=abierto":  `<option value="Abierto" selected>`,
		"/proyectos/1/defectos?estado=EN+CURSO": `<option value="En curso" selected>`,
		"/proyectos/1/defectos?severidad=mayor": `<option value="Mayor" selected>`,
	} {
		html := get(t, server, direccion).Body.String()
		if !strings.Contains(html, elegido) {
			t.Errorf("%s: se esperaba %s en la lista, y no esta en:\n%s", direccion, elegido, between(html, "<main", "</main>"))
		}
	}
}
