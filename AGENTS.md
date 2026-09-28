# Guía del proyecto para agentes de IA

Este archivo le dice a un agente de IA (Claude Code, Copilot, Cursor) cómo trabajar en este
repositorio. También sirve como referencia rápida para los integrantes del equipo.

La cátedra permite usar IA, pero **todo lo generado tiene que quedar entendido y validado
por el equipo**: respondemos por el código igual que si lo hubiéramos escrito a mano. Nada
de pegar sin leer.

## Qué es este proyecto

Aplicación de estimación, seguimiento y medición de proyectos de software. El núcleo y todas
las reglas de negocio van en **Go**. Ver el [README](README.md).

## Comandos

```bash
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
```

## Estilo de código

- Formatear siempre con `gofmt`. Sin excepciones.
- Nombres de tipos, funciones y variables **en español**, como el dominio del problema
  (`Proyecto`, `Historia`, `calcularVelocidad`). Los comentarios también en español.
- Código simple y explícito. Es un trabajo práctico de la facultad, no una librería: se
  prefiere la versión que se entiende leyéndola una vez.
- Los errores se devuelven, no se hace `panic`.

```go
// Historia es un elemento del Product Backlog.
type Historia struct {
	ID          int
	Titulo      string
	StoryPoints int
}

func (h Historia) EstaEstimada() bool {
	return h.StoryPoints > 0
}
```

## Flujo de git

- Rama principal: `main`.
- Una rama por historia: `historia-<numero>-<descripcion-corta>` (ej. `historia-3-crear-proyecto`).
- Mensajes de commit en minúscula, con prefijo: `feat:`, `fix:`, `test:`, `docs:`, `refactor:`.
- Todo entra a `main` por **pull request**, revisado por otro integrante. El PR tiene que
  enlazar la issue y la spec.
- No aplastar commits al mergear: se pierde la evidencia del TDD.

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

El flujo es SDD de cuatro fases con puertas de aprobación:

```
Requisitos → [P1] → Diseño → [P2] → Tareas → [P3] → Implementación
```

La plantilla está en [docs/specs/PLANTILLA.md](docs/specs/PLANTILLA.md). Se revisa en las
puertas, no durante la implementación.
