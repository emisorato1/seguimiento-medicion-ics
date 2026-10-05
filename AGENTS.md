# Guía del proyecto para agentes de IA

Convenciones técnicas del repositorio, para agentes de IA (Claude Code, Copilot, Cursor) y
para los integrantes del equipo.

**Las reglas de trabajo — el flujo, los sprints, las historias y el tablero — están en
[CLAUDE.md](CLAUDE.md).** Leelo primero. Este archivo es solo el cómo técnico.

## Qué es este proyecto

Aplicación de estimación, seguimiento y medición de proyectos de software. El núcleo y todas
las reglas de negocio van en **Go**. Ver el [README](README.md).

## Comandos

```bash
scripts/revisiones.sh    # PRs esperando mi revision
scripts/coherencia.sh    # specs, indice y puertas dicen lo mismo (corre en CI)
scripts/estado.sh        # regenerar ESTADO.md desde el tablero
scripts/tablero.sh ver   # ver el estado de las historias
go test ./...            # correr todos los tests
go test ./internal/...   # solo los tests del dominio
go test -run TestX ./... # correr un test puntual
go build ./...           # compilar
go vet ./...             # analisis estatico
gofmt -l .               # listar archivos mal formateados
```

## Testing

- Framework: el `testing` de la biblioteca estándar. **No agregar dependencias de testing**
  sin acuerdo del equipo.
- Los tests viven al lado del código, en `paquete_test.go`.
- Se escribe **el test primero**: RED → GREEN → REFACTOR.
- El test va en un commit y la implementación en otro. El historial de commits es la
  evidencia del TDD que se evalúa, así que no aplastar la historia ni subir el test junto
  con la implementación ya terminada.
- Todo lo que sea regla de negocio, métrica, estimación o validación tiene que tener test.
- Cómo se nombran y se escriben, en la sección "Cómo se escriben los tests" de este archivo.

## Cómo se escriben los tests

La trazabilidad que se evalúa va **escenario BDD → test**, así que el test tiene que poder
encontrarse desde el escenario sin buscar.

**Un test por escenario**, con el nombre del escenario en PascalCase, sin tildes ni ñ para
que se escriba siempre igual, y un comentario arriba que diga a qué escenario corresponde y
qué reglas cubre:

```go
// Escenario: Rechazar un proyecto sin nombre
// Cubre: RN-2
func TestRechazarUnProyectoSinNombre(t *testing.T) {
	registry := NewProjectRegistry()

	_, err := registry.Create("", date(2026, 10, 1), nil)

	if !errors.Is(err, ErrProjectNameRequired) {
		t.Fatalf("se esperaba ErrProjectNameRequired, salio: %v", err)
	}
}
```

Las reglas:

- **Un "Esquema del escenario" se escribe como un solo test con tabla de casos**, un caso por
  fila de los Ejemplos. El nombre de cada caso es el valor de la fila.
- **Los errores se verifican con `errors.Is` contra el error del dominio, nunca comparando el
  texto del mensaje.** El mensaje es para el usuario y va a cambiar; el error no.
- **Cada test arranca de cero.** Nada de estado compartido entre tests: si un escenario dice
  "Dado un proyecto registrado", el test lo crea.
- **El mensaje de fallo dice qué se esperaba y qué salió.** Un `t.Fatal("fallo")` no sirve
  para nada cuando el CI se pone en rojo.
- Los tests van en `archivo_test.go`, al lado del archivo que prueban.
- Sin librerías de testing: solo `testing` de la biblioteca estándar.

Si un escenario de la Fase 1 no tiene test, la historia no está terminada.

## Estructura del proyecto

```
cmd/            punto de entrada de la aplicación
internal/       nucleo de la solución y reglas de negocio (Go)
docs/specs/     especificaciones SDD, una carpeta por sprint
docs/actas/     actas de plannings, reviews y retrospectivas
scripts/        utilidades del tablero y del estado
```

## Decisiones de arquitectura

Estas valen para **todo el proyecto**. Cada spec las da por hechas y no las vuelve a
discutir en su Fase 2: ahí solo van las decisiones propias de esa funcionalidad. Si alguna
de estas hay que cambiarla, se cambia acá y se avisa al equipo.

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| El núcleo vive en un solo paquete `internal/domain`, con **un archivo por entidad** (`project.go`, `member.go`, `story.go`, …), y se puede **separar en más archivos por funcionalidad** cuando varias historias tocan la misma entidad (`story_estimation.go`, `project_registry.go`) | un paquete por entidad | con paquetes separados aparecen ciclos de importación en cuanto una historia referencia a un proyecto. La regla existe para que varias personas no editen el mismo archivo: cuando separar por funcionalidad sirve mejor a ese objetivo, se separa |
| Durante el Sprint 1 los datos viven **en memoria**: al cerrar el programa se pierden | archivo o base de datos | ningún requerimiento mínimo pide persistencia todavía, y meterla ahora obligaría a decidir el formato antes de saber qué se guarda. Se revisa en el Sprint 2 |
| Los errores se devuelven como **valores de error del dominio**, declarados en el paquete y comparables con `errors.Is` | devolver strings, o `panic` | permite que los tests verifiquen *qué* error salió y no el texto del mensaje, que va a cambiar |
| Los **mensajes** de error son para el usuario y van en español; los **nombres** de los errores en inglés, como el resto del código | todo en inglés | el mensaje lo lee quien usa el sistema, y varias reglas de negocio dicen literalmente qué se le informa |
| Los identificadores los asigna el dominio: enteros que arrancan en 1 y son **únicos dentro de cada proyecto** | UUID, o únicos globales | es lo que esperan las specs (#6 RN-16) y lo que espera el usuario: la primera historia de un proyecto es la 1 |
| El núcleo **no sabe nada de la interfaz**: no imprime, no lee del teclado, no formatea | que el dominio imprima directamente | es lo que permite testearlo, y lo que hace que la consola de la #24 sea reemplazable en el Sprint 2 |
| **Sin dependencias externas**: solo la biblioteca estándar de Go | usar un framework o una librería de tests | el proyecto es chico y la cátedra evalúa el código propio. Agregar una dependencia se acuerda con el equipo |
| Las validaciones viven **en el dominio**, nunca en la interfaz | validar en la consola para dar mejor respuesta | si la validación está en la interfaz, cambiar de interfaz la pierde, y los tests no la cubren |
| Quien guarda una entidad **entrega la entidad guardada, no una copia**, y sus datos están **ocultos fuera del paquete**: solo se cambian con operaciones que validan | entregar una copia y obligar a guardar cada cambio | la consola vive en `cmd/`, así que no puede tocar datos ocultos de `internal/domain` y no tiene forma de saltear las validaciones. Dentro del paquete la copia no protege nada, porque es el mismo paquete: paga el costo sin comprar la protección. Además, con copias cada historia tendría que acordarse de guardar, y un olvido pierde el cambio en silencio, sin error y sin que ningún escenario BDD lo muestre |
| Toda operación **valida todo antes de cambiar cualquier dato** | validar y aplicar a medida que se avanza | es la condición que hace que una operación rechazada deje la entidad como estaba, sin tener que deshacer. Lo pide la RN-5 de la #3, la RN-12 de la #4 y la RN-11 de la #6 |

## Estilo de código

- Formatear siempre con `gofmt`. Sin excepciones.
- Nombres de tipos, funciones y variables **en inglés**, que es la convención de Go
  (`Project`, `Story`, `CalculateVelocity`). Los comentarios y la documentación, en español.
- Cuando el nombre traduce un término del dominio, usar siempre la misma traducción:
  proyecto → `Project`, historia → `Story`, sprint → `Sprint`, defecto → `Defect`,
  esfuerzo → `Effort`, métrica → `Metric`, story points → `StoryPoints`.
- Código simple y explícito. Es un trabajo práctico de la facultad, no una librería: se
  prefiere la versión que se entiende leyéndola una vez.
- Los errores se devuelven, no se hace `panic`.

```go
// Story es un elemento del Product Backlog.
type Story struct {
	ID          int
	Title       string
	StoryPoints int
}

// IsEstimated indica si la historia ya tiene una estimación acordada.
func (s Story) IsEstimated() bool {
	return s.StoryPoints > 0
}
```

## Flujo de git

- Ramas: `main` (producción) ← PR ← `dev` (integración) ← PR ← rama de historia.
- Una rama por historia, sacada de `dev`: `historia-<numero>-<descripcion-corta>`
  (ej. `historia-3-crear-proyecto`). Para defectos: `defecto-<numero>-<descripcion>`.
- Nadie commitea directo a `main` ni a `dev`.
- Mensajes de commit en minúscula, con prefijo: `feat:`, `fix:`, `test:`, `docs:`, `refactor:`.
- El PR se completa con la plantilla, enlaza la issue y la spec, y lo revisa otro integrante.
- No se mergea con el CI en rojo.
- Se mergea con **merge commit**: squash y rebase están desactivados para no perder la
  evidencia del TDD. Tampoco se hace `push --force`.
- La rama se borra sola al mergear el PR; localmente se limpia con `git fetch --prune`.

Las reglas completas del flujo de ramas están en [CLAUDE.md](CLAUDE.md).

## Boundaries

**Nunca:**

- Commitear secretos, tokens o credenciales.
- Escribir código de una funcionalidad que todavía no tiene su spec aprobada en
  `docs/specs/`.
- Subir la implementación en el mismo commit que su test.
- Reescribir el historial de `main` (`push --force`, rebase de commits ya publicados).
- Agregar dependencias externas sin acuerdo del equipo: el núcleo va en Go estándar.

**Preguntar antes de:**

- Crear un paquete o una capa nueva que no esté en la estructura de arriba.
- Cambiar una decisión que ya está escrita en una spec aprobada.
- Tocar la configuración del repositorio o del tablero.

**Siempre:**

- Leer la spec de la funcionalidad antes de tocar su código.
- Correr `go test ./...` y `gofmt -l .` antes de abrir un pull request.
- Si la implementación se desvía de la spec, actualizar la spec en el mismo PR.

## Cómo trabajamos

El flujo completo está en [CLAUDE.md](CLAUDE.md): SDD de cuatro fases con puertas de
aprobación, BDD, TDD y las reglas del tablero.
