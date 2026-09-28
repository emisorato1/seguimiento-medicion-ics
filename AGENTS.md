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

## Estructura del proyecto

```
cmd/            punto de entrada de la aplicación
internal/       nucleo de la solución y reglas de negocio (Go)
docs/specs/     especificaciones SDD, una por funcionalidad
docs/actas/     actas de plannings, reviews y retrospectivas
scripts/        utilidades del tablero y del estado
```

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
- No aplastar commits al mergear ni hacer `push --force`: se pierde la evidencia del TDD.

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
