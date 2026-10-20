# Plantilla de especificación SDD

Copiar a `docs/specs/sprint-N/NN-nombre-corto.md` (N el sprint donde se comprometió la
historia, NN el número de la issue) y completar **antes** de escribir código.

```
Requisitos → [Puerta 1] → Diseño → [Puerta 2] → Tareas → [Puerta 3] → Implementación
```

Se revisa **en las puertas, no durante la implementación**. Cada puerta la aprueba alguien
distinto del que escribió la fase, y queda registrado quién y cuándo.

## Las dos reglas que más se incumplen

**1. En Requisitos no va nada técnico.** Ni base de datos, ni estructuras, ni nombres de
funciones. Solo comportamiento. La prueba: si un requisito solo lo entiende un
programador, tiene decisiones técnicas disfrazadas y hay que reescribirlo.

**2. En Diseño van decisiones, no código.** Nada de structs, firmas ni pseudocódigo: si el
diseño se vuelve código en prosa, el agente lo traduce mecánicamente y se pierde todo el
valor. La heurística: *si dos personas con experiencia tomarían la misma decisión sin
discutirla, no va en el diseño; si podrían elegir distinto y las consecuencias cambian, sí va.*

Si la spec termina siendo más larga que el código que genera, está sobreespecificada.

---

# Spec: [título de la funcionalidad]

- **Historia:** #[N]
- **Sprint:** [N]
- **Responsable:** [quién]
- **Estado:** Requisitos | Diseño | Tareas | Implementación | Terminada
- **Depende de:** [#N, o "nada"]
- **La necesitan:** [#N, #M, o "nadie todavía"]

## Fase 1 — Requisitos

Qué tiene que hacer el sistema, desde el punto de vista de quien lo usa.

### Objetivo

[Una o dos frases: qué problema resuelve y para quién.]

### Qué queda afuera

[Lo que alguien podría suponer que entra y no entra. Evita discusiones en la puerta.]

### Entradas

| Dato | Obligatorio | Descripción |
|---|---|---|

### Salidas

[Qué devuelve o qué cambia en el sistema.]

### Reglas de negocio

Se escriben con la notación **EARS**, que es la que recomienda el apunte de la cátedra
porque elimina ambigüedad sin dejar de ser legible. Cinco patrones:

| Patrón | Forma | Para qué |
|---|---|---|
| Ubicuo | El sistema debe… | propiedad permanente |
| Dirigido por evento | Cuando [evento], el sistema debe… | se dispara con algo |
| Dirigido por estado | Mientras [condición], el sistema debe… | vale mientras dure algo |
| No deseado | Si [situación], el sistema debe… | errores y excepciones |
| Opcional | Donde [configuración], el sistema debe… | depende de una opción |

1. [RN-1] …
2. [RN-2] …

### Casos límite

[Los bordes: cero, vacío, el máximo, el primero, el último, el duplicado.]

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|

### Criterios de aceptación

Tienen que ser **verificables**: mirando el resultado se tiene que poder decir si se cumplen
o no. "Es rápido" no sirve; "responde en menos de 2 segundos" sí.

- [ ]

### Escenarios BDD

```gherkin
Escenario: [caso normal]
  Dado [contexto]
  Cuando [acción]
  Entonces [resultado esperado]
```

Cubrir casos normales, alternativos, límite y de error. Cada escenario termina en al menos
un test.

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: [quién] — Fecha: [dd/mm] — Comentarios:

## Fase 2 — Diseño

Solo se escribe **después** de que la Puerta 1 esté aprobada.

### Enfoque

[Cómo se va a resolver, en prosa. Qué patrón o convención del proyecto se sigue. Si es la
primera funcionalidad que hace algo de este tipo, qué convención se está estableciendo.]

### Archivos afectados

| Archivo | Se crea o se modifica | Para qué |
|---|---|---|

### Decisiones

Solo las que podrían haberse resuelto de otra forma razonable.

| Decisión | Alternativa descartada | Por qué |
|---|---|---|

### Dependencias y restricciones

[Qué tiene que estar listo antes. Qué **no** hay que tocar para no romper otra cosa.]

> **Puerta 2** — ¿el diseño es viable y coherente con el resto del proyecto?
> Aprobó: [quién] — Fecha: [dd/mm] — Comentarios:

## Fase 3 — Tareas

Cada tarea tiene que poder resolverse en un ciclo de TDD y terminar en commits.

| # | Tarea | Regla o criterio que cubre | Estado |
|---|---|---|---|

> **Puerta 3** — ¿las tareas cubren todos los criterios de aceptación?
> Aprobó: [quién] — Fecha: [dd/mm] — Comentarios:

## Fase 4 — Implementación

Ciclo RED → GREEN → REFACTOR. **El test en un commit y la implementación en otro.**

| Tarea | Commit del test | Commit de la implementación |
|---|---|---|

### Desvíos respecto de la spec

[Si hubo que cambiar algo de lo especificado, anotarlo acá y actualizar la fase que
corresponda, en el mismo pull request.]

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #[N] |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | [archivos `_test.go`] |
| Código Go | [archivos] |
