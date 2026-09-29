# Spec: Estimar historias con Story Points

- **Historia:** #12
- **Sprint:** 1
- **Responsable:** Maximiliano Eula
- **Estado:** Requisitos — **borrador, pendiente de Puerta 1**
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de
  la consigna y de los criterios de aceptación de la issue. El responsable lo
  revisa, lo corrige si hace falta y lo presenta a la Puerta 1
- **Depende de:** #6 (tiene que existir la historia a estimar)
- **La necesitan:** #9 (solo entran al sprint las historias estimadas), #13 (Planning Poker
  registra la estimación acordada), #16 (métricas de story points y velocidad)

## Fase 1 — Requisitos

### Objetivo

Que se pueda asignarle a una historia del backlog un tamaño en story points, para poder
después planificar cuánto entra en un sprint y medir la velocidad del equipo.

### Qué queda afuera

- Cómo se llega al número entre varias personas: eso es Planning Poker, la historia #13.
  Acá se guarda un valor ya acordado.
- Estimar en horas: las horas son la historia #14 y son otra cosa.

### Entradas

| Dato | Obligatorio | Descripción |
|---|---|---|
| Historia | sí | qué historia se está estimando |
| Story points | sí | el valor acordado |

### Salidas

La historia queda con sus story points registrados y pasa a estar estimada.

### Reglas de negocio

1. **[RN-1]** El sistema debe aceptar como story points únicamente los valores de la
   sucesión de Fibonacci **1, 2, 3, 5, 8, 13 y 21**.
2. **[RN-2]** Si se intenta estimar con un valor que no está en esa lista, el sistema debe
   rechazar la operación e informar cuáles son los valores válidos.
3. **[RN-3]** Si se intenta estimar con cero o con un valor negativo, el sistema debe
   rechazar la operación.
4. **[RN-4]** Una historia sin story points debe considerarse **no estimada**, y eso no es un
   error: es el estado normal de una historia recién creada.
5. **[RN-5]** Mientras la historia no esté Terminada, el sistema debe permitir volver a
   estimarla con otro valor.
6. **[RN-6]** Si se intenta reestimar una historia que ya está Terminada, el sistema debe
   rechazar la operación e informar que una historia terminada no se reestima.
7. **[RN-7]** Cuando se reestima una historia, el sistema debe reemplazar el valor anterior.
8. **[RN-8]** El sistema debe permitir quitarle la estimación a una historia, dejándola de
   nuevo como no estimada, siempre que no esté Terminada.
9. **[RN-9]** Si se intenta estimar una historia que no existe, el sistema debe informar que
   no se encontró la historia.

### Casos límite

- Estimar con **1**, el valor más chico de la escala: válido.
- Estimar con **21**, el valor más grande: válido, aunque conviene partir la historia.
- Estimar con **4** o con **7**: inválido, no están en la escala.
- Reestimar con el mismo valor que ya tenía: válido, no es error.
- Consultar los story points de una historia no estimada: devuelve "sin estimar", no un cero.

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| Valor fuera de la escala de Fibonacci | rechaza e informa los valores válidos |
| Valor cero o negativo | rechaza e informa los valores válidos |
| Reestimar una historia Terminada | rechaza e informa que no se reestima |
| Historia inexistente | informa que no se encontró la historia |

### Criterios de aceptación

- [ ] Se puede asignar story points a una historia con cualquiera de los valores 1, 2, 3, 5,
      8, 13 o 21.
- [ ] No se puede asignar un valor fuera de esa escala, por ejemplo 4 o 7.
- [ ] No se puede asignar cero ni un valor negativo.
- [ ] Una historia recién creada figura como no estimada, y eso no es un error.
- [ ] Se puede reestimar una historia mientras no esté Terminada, y queda el último valor.
- [ ] No se puede reestimar una historia Terminada.
- [ ] Se le puede quitar la estimación a una historia no terminada.

### Escenarios BDD

```gherkin
Escenario: Estimar una historia
  Dado una historia en estado "Pendiente" sin estimar
  Cuando le asigno 5 story points
  Entonces la historia queda estimada en 5

Escenario: Aceptar los extremos de la escala
  Dado una historia en estado "Pendiente" sin estimar
  Cuando le asigno 1 story point
  Entonces la historia queda estimada en 1

Escenario: Rechazar un valor fuera de Fibonacci
  Dado una historia en estado "Pendiente"
  Cuando intento asignarle 4 story points
  Entonces la operacion se rechaza
  Y se informa que los valores validos son 1, 2, 3, 5, 8, 13 y 21

Escenario: Rechazar cero
  Dado una historia en estado "Pendiente"
  Cuando intento asignarle 0 story points
  Entonces la operacion se rechaza

Escenario: Una historia nueva esta sin estimar
  Dado un proyecto registrado
  Cuando creo una historia nueva
  Entonces la historia figura como no estimada

Escenario: Reestimar una historia
  Dado una historia estimada en 3
  Cuando la reestimo en 8
  Entonces la historia queda estimada en 8

Escenario: No reestimar una historia terminada
  Dado una historia en estado "Terminada" estimada en 5
  Cuando intento reestimarla en 8
  Entonces la operacion se rechaza
  Y se informa que una historia terminada no se reestima

Escenario: Quitar la estimacion
  Dado una historia en estado "Pendiente" estimada en 5
  Cuando le quito la estimacion
  Entonces la historia figura como no estimada
```

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios:

## Fase 2 — Diseño

_Se escribe después de que la Puerta 1 esté aprobada._

## Fase 3 — Tareas

_Se escribe después de que la Puerta 2 esté aprobada._

## Fase 4 — Implementación

_Se completa a medida que se avanza._

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #12 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | pendiente |
| Código Go | pendiente |
