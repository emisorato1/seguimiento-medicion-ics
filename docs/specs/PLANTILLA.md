# Plantilla de especificación SDD

Copiar este archivo a `docs/specs/NN-nombre-de-la-funcionalidad.md` (el NN es el número de
la issue) y completarlo **antes** de escribir código.

El flujo tiene cuatro fases con una puerta de aprobación entre cada una:

```
Requisitos → [Puerta 1] → Diseño → [Puerta 2] → Tareas → [Puerta 3] → Implementación
```

La regla es **revisar en las puertas, no durante la implementación**. Cada puerta la aprueba
un integrante distinto del que escribió la fase, y se registra abajo quién aprobó y cuándo.
La aprobación puede ser un comentario en el pull request de la spec: no hace falta reunión.

**No todas las funcionalidades necesitan las cuatro fases con la misma profundidad.** Una
historia chica puede tener requisitos y tareas con un diseño de tres líneas. Inflar la spec
de algo simple es un anti-patrón, no una virtud.

---

# Spec: [título de la funcionalidad]

- **Historia:** #[número de la issue]
- **Autor de la spec:** [quién]
- **Estado:** Requisitos | Diseño | Tareas | Implementación | Terminada

## Fase 1 — Requisitos

Qué se va a construir, desde el punto de vista del usuario. Acá **no** se habla de
tecnología ni de cómo está estructurado el código.

### Objetivo

[Una o dos frases: qué problema resuelve y para quién.]

### Entradas

[Qué datos recibe, con su tipo y de dónde vienen.]

### Salidas

[Qué devuelve o qué cambia en el sistema.]

### Reglas de negocio

[Numeradas, una por línea, de forma que cada una se pueda testear por separado.]

1.
2.

### Restricciones

[Límites técnicos o de alcance: qué queda explícitamente afuera.]

### Casos límite

[Valores en el borde: cero, vacío, el máximo, el primero, el último.]

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| | |

### Criterios de aceptación

[Los mismos que están en la issue. Si acá aparecen criterios nuevos, actualizar la issue.]

- [ ]

### Escenarios BDD

```gherkin
Escenario: [nombre del caso normal]
  Dado [contexto]
  Cuando [acción]
  Entonces [resultado esperado]

Escenario: [nombre de un caso de error]
  Dado [contexto]
  Cuando [acción]
  Entonces [resultado esperado]
```

Cubrir casos normales, alternativos, límite y de error.

> **Puerta 1** — ¿los requisitos son los correctos y están completos?
> Aprobó: [quién] — Fecha: [dd/mm] — Comentarios: [si hubo cambios, cuáles]

## Fase 2 — Diseño

Cómo se va a construir. Acá sí se habla de código.

### Tipos y funciones

[Structs, interfaces y firmas de las funciones principales en Go.]

### Decisiones de diseño

[Qué se decidió y por qué. Si se descartó una alternativa, dejar dicho cuál y el motivo:
esto es lo que se pregunta en la defensa.]

### Qué se toca

[Archivos y paquetes que se crean o se modifican.]

> **Puerta 2** — ¿el diseño es viable y coherente con el resto?
> Aprobó: [quién] — Fecha: [dd/mm] — Comentarios:

## Fase 3 — Tareas

Descomposición en unidades chicas. Cada tarea tiene que poder resolverse en un ciclo de TDD
(un test que falla, el código que lo hace pasar, refactor) y terminar en un commit.

| # | Tarea | Test que la cubre | Estado |
|---|---|---|---|
| 1 | | | |
| 2 | | | |

> **Puerta 3** — ¿las tareas cubren todos los criterios de aceptación?
> Aprobó: [quién] — Fecha: [dd/mm] — Comentarios:

## Fase 4 — Implementación

Se ejecutan las tareas con el ciclo RED → GREEN → REFACTOR. El historial de commits es la
evidencia del TDD, así que **el test va en un commit y la implementación en otro**.

### Commits

[Se completa a medida que se avanza.]

| Tarea | Commit del test | Commit de la implementación |
|---|---|---|
| | | |

### Desvíos respecto de la spec

[Si durante la implementación hubo que cambiar algo de lo especificado, anotarlo acá y
actualizar la fase correspondiente. Una spec que no se actualiza deja de servir.]

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #[N] |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | [archivos `_test.go`] |
| Código Go | [archivos] |
