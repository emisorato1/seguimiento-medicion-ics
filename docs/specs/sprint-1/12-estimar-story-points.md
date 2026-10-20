# Spec: Estimar historias con Story Points

- **Historia:** #12
- **Sprint:** 1
- **Responsable:** Maximiliano Eula
- **Estado:** Puerta 3 aprobada — **lista para implementar**
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de
  la consigna y de los criterios de aceptación de la issue. El responsable lo
  revisa, lo corrige si hace falta y lo presenta a la Puerta 1. Corregido por el
  responsable y presentado en el PR #25
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
   error. Que una historia recién creada quede sin estimar lo define la historia #6 (RN-7);
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
> Aprobó: Emiliano Sorato — Fecha: 2026-09-29 — Comentarios: aprobado en el PR #25, sin
> observaciones. El PR planteaba dos decisiones para la puerta: quitar la estimación a una
> historia sin estimar no es error, y una historia En curso se puede reestimar.

## Fase 2 — Diseño

Las decisiones transversales del proyecto —un solo paquete `internal/domain`, datos en
memoria, errores como valores del dominio comparables y mensajes en español— están en
[AGENTS.md](../../../AGENTS.md) y no se repiten acá. El diseño se apoya además en la Fase 2
de la historia #6: la historia, su estado como tipo propio y el backlog dentro del proyecto.
Esta fase registra solo lo propio de la estimación.

### Enfoque

Estimar no crea nada nuevo: **agrega comportamiento a la historia que define la #6**. El dato
de story points ya existe en la historia (RN-2 de la #6), y esta historia define cómo se
carga, se cambia y se quita.

Hay dos operaciones:

- **Estimar**, que sirve tanto para la primera estimación como para reestimar (RN-5, RN-6).
- **Quitar la estimación** (RN-7), que es un pedido explícito y separado.

Las dos siguen el mismo orden, de lo más general a lo más particular:

1. **Encontrar la historia** dentro de su proyecto. Si no existe, se informa que no se
   encontró (RN-9).
2. **Controlar el estado.** Si está Terminada, se rechaza cualquier cambio en la estimación
   (RN-8).
3. **Validar el valor**, solo al estimar: tiene que estar en la escala (RN-1 a RN-3).
4. **Aplicar** el cambio recién si todo lo anterior pasó.

Como en la #6, aplicar al final es lo que garantiza que una operación rechazada deje la
historia como estaba.

**"Sin estimar" se representa con el valor 0.** El 0 nunca es una estimación válida (RN-3),
así que no hay ambigüedad. Para cumplir la RN-4, el sistema ofrece una forma explícita de
consultar si la historia está estimada, y quien muestra el dato informa "sin estimar" en
lugar de un cero.

### Archivos afectados

| Archivo | Se crea o se modifica | Para qué |
|---|---|---|
| `internal/domain/story_estimation.go` | se crea | la escala válida, las operaciones de estimar y quitar la estimación, y la consulta de si la historia está estimada |
| `internal/domain/story_estimation_test.go` | se crea | los tests de los escenarios BDD de la Fase 1, nombrados como cada escenario |
| `internal/domain/errors.go` | se modifica | los errores de esta historia: valor fuera de la escala e historia terminada |

### Decisiones

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| "Sin estimar" se representa con el valor 0 | un dato aparte que indique si está estimada | el 0 nunca es válido (RN-3), y un solo dato no puede quedar en un estado contradictorio |
| Quitar la estimación es una operación propia | estimar con 0 para quitar | estimar con 0 está prohibido por la RN-3 |
| La escala válida es una lista fija, usada por la validación y por el mensaje de error | calcular la sucesión de Fibonacci | la escala es una lista cerrada, y con una sola fuente el mensaje de la RN-2 nunca miente |
| Un solo error para cualquier valor inválido | un error para fuera de la escala y otro para cero o negativo | la Fase 1 da la misma respuesta en los dos casos (RN-2, RN-3) |
| Orden de los controles: existencia → estado → valor | validar el valor primero | no tiene sentido validar el valor de una historia que no existe o no se puede tocar |
| La estimación va en `story_estimation.go` y sus tests en `story_estimation_test.go`; el dato sigue en `story.go` | ponerlo todo en `story.go` | evita que varias historias toquen el mismo archivo y deja la trazabilidad directa; **se aparta de la letra de AGENTS.md: para decidir en la puerta** |

### Dependencias y restricciones

- **Depende de la #6** (y por ella, de la #3): necesita la historia, su estado y el backlog
  dentro del proyecto. El código no se empieza hasta que la #6 esté en `dev`.
- **No se modifica `story.go`.** El dato de story points lo define la #6, con 0 como valor
  inicial (su RN-7): la historia nace sin estimar sin trabajo extra. Esta historia solo le
  agrega comportamiento.
- **El estado de la historia solo se lee**, no se cambia: las transiciones son de la #9 y
  la #10.
- **`errors.go` es compartido** (Fase 2 de la #6): los errores de esta historia se agregan ahí.
- **La usan otras historias:** la #13 (Planning Poker) registra la estimación acordada usando
  la operación de estimar de esta historia, sin duplicarla; y la #24 la expone en la consola.

> **Puerta 2** — ¿el diseño es viable y coherente con el resto del proyecto?
> Aprobó: Emiliano Sorato — Fecha: 2026-10-05 — Comentarios: aprobada en el PR #39. Se acepto separar story_estimation.go: la regla de AGENTS.md de un archivo por entidad existia para evitar que varias historias editen el mismo archivo, y la separacion sirve mejor a ese objetivo. Se corrigio AGENTS.md en consecuencia.

## Fase 3 — Tareas

Cada tarea es un ciclo de TDD completo: el test que falla en un commit, la implementación que
lo hace pasar en otro. Están en el orden en que hay que hacerlas, y cada una deja el código
compilando y los tests en verde.

Los tests siguen la convención de [AGENTS.md](../../../AGENTS.md): un test por escenario, con
el nombre del escenario. Cada tarea de rechazo verifica, además del error, que la historia
quedó como estaba: es la condición de "validar todo antes de cambiar cualquier dato".

| # | Tarea | Escenario BDD | Reglas | Estado |
|---|---|---|---|---|
| 1 | Una historia recién creada figura como no estimada al consultarla | Consultar una historia sin estimar | RN-4 | pendiente |
| 2 | Estimar con cada valor de la escala: la historia queda estimada en ese valor | Estimar con cada valor de la escala | RN-1, RN-5 | pendiente |
| 3 | Rechazar los valores fuera de la escala con un único error que informa los valores válidos; la historia sigue sin estimar | Rechazar valores fuera de la escala | RN-1, RN-2, RN-3 | pendiente |
| 4 | Reestimar reemplaza el valor anterior | Reestimar una historia | RN-5, RN-6 | pendiente |
| 5 | Reestimar con el mismo valor no es error | Reestimar con el mismo valor | RN-5 | pendiente |
| 6 | Una historia En curso se puede reestimar | Reestimar una historia en curso | RN-5 | pendiente |
| 7 | Quitar la estimación deja la historia sin estimar | Quitar la estimacion | RN-7 | pendiente |
| 8 | Quitarle la estimación a una historia ya sin estimar no es error | — (caso límite) | RN-7 | pendiente |
| 9 | Una historia Terminada no se estima ni se reestima, y conserva su valor | No reestimar una historia terminada | RN-8 | pendiente |
| 10 | A una historia Terminada no se le quita la estimación, y conserva su valor | No quitar la estimacion de una historia terminada | RN-8 | pendiente |
| 11 | Estimar una historia que no existe informa que no se encontró | Estimar una historia que no existe | RN-9 | pendiente |
| 12 | Quitarle la estimación a una historia que no existe informa que no se encontró | — (RN-9) | RN-9 | pendiente |
| 13 | Con más de un problema se informa el primero, en el orden existencia → estado → valor | — (decisión de la Fase 2) | RN-8, RN-9 | pendiente |

**Tareas sin escenario.** La 8, la 12 y la 13 cubren un caso límite, la segunda mitad de la
RN-9 y el orden de los controles decidido en la Fase 2, que no tienen escenario propio en la
Fase 1. Sus tests se nombran por lo que verifican (`TestQuitarLaEstimacionDeUnaHistoriaYaSinEstimar`,
`TestQuitarLaEstimacionDeUnaHistoriaInexistente`, `TestConMasDeUnProblemaSeInformaElPrimero`),
y el comentario de arriba dice qué regla o decisión cubren en lugar de qué escenario.

**Cómo se verifica la RN-2.** El error se verifica con `errors.Is`, como pide la convención.
Además, el test de la tarea 3 verifica que el mensaje **incluya cada valor de la escala**, sin
comparar el texto completo: la redacción puede cambiar, pero la RN-2 exige que los valores
válidos estén.

### Por qué este orden

- **La 1 va primero** porque es la consulta de si la historia está estimada, que usan todas
  las demás para verificar el resultado.
- **La 2 es el camino feliz**, y la 3 agrega la validación. Al revés, el primer test de
  validación no tendría contra qué comparar.
- **De la 4 a la 8** son reestimar y quitar sobre historias que se pueden modificar.
- **De la 9 a la 12** son los rechazos, que se agregan uno por uno.
- **La 13 va al final** porque solo tiene sentido cuando ya existen los tres errores que hay
  que ordenar.

### Cobertura

Las nueve reglas de negocio quedan cubiertas: RN-1 (2, 3), RN-2 (3), RN-3 (3), RN-4 (1),
RN-5 (2, 4, 5, 6), RN-6 (4), RN-7 (7, 8), RN-8 (9, 10, 13), RN-9 (11, 12, 13).

| Criterio de aceptación | Tareas |
|---|---|
| Se puede asignar story points con cualquiera de los valores 1, 2, 3, 5, 8, 13 o 21 | 2 |
| No se puede asignar un valor fuera de esa escala, por ejemplo 4 o 7 | 3 |
| No se puede asignar cero ni un valor negativo | 3 |
| Una historia sin estimar figura como no estimada, y eso no es un error | 1 |
| Se puede reestimar una historia mientras no esté Terminada, y queda el último valor | 4, 5, 6 |
| No se puede cambiar la estimación de una historia Terminada | 9, 10 |
| Se le puede quitar la estimación a una historia no terminada | 7 |
| Estimar una historia que no existe informa que no se encontró | 11 |

Los diez escenarios BDD de la Fase 1 tienen su tarea, en la columna "Escenario BDD".

### Dependencias

- **Todas las tareas necesitan las de la #6** en `dev`: la historia, su estado y el backlog
  dentro del proyecto.
- **Las tareas 6, 9 y 10 necesitan historias En curso o Terminada**, y los cambios de estado
  son de la #9 y la #10, que todavía no existen. Los tests las preparan desde dentro del
  paquete, que puede fijar el estado directamente.

> **Puerta 3** — ¿las tareas cubren todos los criterios de aceptación?
> Aprobó: Emiliano Sorato — Fecha: 2026-10-08 — Comentarios: aprobada en el PR #47. Las tareas 9 y 10 quedan separadas porque son dos caminos distintos de la RN-8 y, juntas, uno podria quedar sin cubrir sin que se note.

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
