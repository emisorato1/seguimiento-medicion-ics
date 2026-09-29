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
   rechazar la operación e informar cuáles son los valores válidos.
4. **[RN-4]** Una historia sin story points debe considerarse **no estimada**, y eso no es un
   error. Que una historia recién creada quede sin estimar lo define la historia #6 (RN-6);
   acá solo se garantiza que al consultarla figure como "sin estimar" y no con un cero.
5. **[RN-5]** Mientras la historia no esté Terminada, el sistema debe permitir estimarla, y
   volver a estimarla con otro valor o con el mismo.
6. **[RN-6]** Cuando se reestima una historia, el sistema debe reemplazar el valor anterior.
7. **[RN-7]** Mientras la historia no esté Terminada, el sistema debe permitir quitarle la
   estimación, dejándola de nuevo como no estimada.
8. **[RN-8]** Si la historia está Terminada, el sistema debe rechazar cualquier cambio en su
   estimación —estimarla, reestimarla o quitarle la estimación— e informar que la
   estimación de una historia terminada no se modifica.
9. **[RN-9]** Si se intenta estimar una historia que no existe, o quitarle la estimación, el
   sistema debe informar que no se encontró la historia.

Los estados de una historia (Pendiente, En curso, Terminada) son los que define la
historia #6. "Terminada" es la historia completada en un sprint cerrado, que es la condición
que pone la issue: _"se puede reestimar mientras no esté en un sprint cerrado"_.

### Casos límite

- Estimar con **1**, el valor más chico de la escala: válido.
- Estimar con **21**, el valor más grande: válido, aunque conviene partir la historia.
- Estimar con **4** o con **7**: inválido, no están en la escala.
- Reestimar con el mismo valor que ya tenía: válido, no es error.
- Quitarle la estimación a una historia que ya estaba sin estimar: válido, sigue sin estimar.
- Estimar o reestimar una historia **En curso**: válido, porque no está Terminada.
- Consultar los story points de una historia no estimada: devuelve "sin estimar", no un cero.

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| Valor fuera de la escala de Fibonacci | rechaza e informa los valores válidos |
| Valor cero o negativo | rechaza e informa los valores válidos |
| Estimar, reestimar o quitar la estimación de una historia Terminada | rechaza e informa que no se modifica |
| Historia inexistente | informa que no se encontró la historia |

### Criterios de aceptación

- [ ] Se puede asignar story points a una historia con cualquiera de los valores 1, 2, 3, 5,
      8, 13 o 21.
- [ ] No se puede asignar un valor fuera de esa escala, por ejemplo 4 o 7.
- [ ] No se puede asignar cero ni un valor negativo.
- [ ] Una historia sin estimar figura como no estimada, y eso no es un error.
- [ ] Se puede reestimar una historia mientras no esté Terminada, y queda el último valor.
- [ ] No se puede cambiar la estimación de una historia Terminada.
- [ ] Se le puede quitar la estimación a una historia no terminada.
- [ ] Estimar una historia que no existe informa que no se encontró.

### Escenarios BDD

```gherkin
Esquema del escenario: Estimar con cada valor de la escala
  Dado una historia en estado "Pendiente" sin estimar
  Cuando le asigno <puntos> story points
  Entonces la historia queda estimada en <puntos>

  Ejemplos:
    | puntos |
    | 1      |
    | 2      |
    | 3      |
    | 5      |
    | 8      |
    | 13     |
    | 21     |

Esquema del escenario: Rechazar valores fuera de la escala
  Dado una historia en estado "Pendiente" sin estimar
  Cuando intento asignarle <puntos> story points
  Entonces la operacion se rechaza
  Y se informa que los valores validos son 1, 2, 3, 5, 8, 13 y 21
  Y la historia sigue sin estimar

  Ejemplos:
    | puntos | por que                |
    | 4      | no esta en Fibonacci   |
    | 7      | no esta en Fibonacci   |
    | 0      | cero                   |
    | -3     | negativo               |

Escenario: Consultar una historia sin estimar
  Dado una historia en estado "Pendiente" sin estimar
  Cuando consulto sus story points
  Entonces figura como "sin estimar"

Escenario: Reestimar una historia
  Dado una historia en estado "Pendiente" estimada en 3
  Cuando la reestimo en 8
  Entonces la historia queda estimada en 8

Escenario: Reestimar con el mismo valor
  Dado una historia en estado "Pendiente" estimada en 5
  Cuando la reestimo en 5
  Entonces la historia queda estimada en 5

Escenario: Reestimar una historia en curso
  Dado una historia en estado "En curso" estimada en 3
  Cuando la reestimo en 5
  Entonces la historia queda estimada en 5

Escenario: No reestimar una historia terminada
  Dado una historia en estado "Terminada" estimada en 5
  Cuando intento reestimarla en 8
  Entonces la operacion se rechaza
  Y se informa que la estimacion de una historia terminada no se modifica
  Y la historia sigue estimada en 5

Escenario: Quitar la estimacion
  Dado una historia en estado "Pendiente" estimada en 5
  Cuando le quito la estimacion
  Entonces la historia figura como no estimada

Escenario: No quitar la estimacion de una historia terminada
  Dado una historia en estado "Terminada" estimada en 5
  Cuando intento quitarle la estimacion
  Entonces la operacion se rechaza
  Y la historia sigue estimada en 5

Escenario: Estimar una historia que no existe
  Dado que no existe ninguna historia con identificador 99
  Cuando intento asignarle 5 story points a la historia 99
  Entonces se informa que no se encontro la historia
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
