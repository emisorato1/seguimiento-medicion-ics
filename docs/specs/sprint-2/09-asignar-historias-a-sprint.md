# Spec: Asignar historias a un sprint

- **Historia:** #9
- **Sprint:** 2
- **Responsable:** Agustín Salinas
- **Estado:** Borrador completo — pendiente de las tres puertas
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de la consigna y
  de los criterios de aceptación de la issue. El responsable lo revisa, lo corrige si hace
  falta y lo presenta a las puertas.
- **Depende de:** #8 (tiene que existir el sprint abierto), #6 (las historias y su estado) y
  #12 (solo entran historias estimadas); por la #6 y la #8, también de la #3
- **La necesitan:** #10 (cierra el sprint sobre las historias que esta historia le asignó),
  #11 (muestra las historias de cada sprint cerrado), #16 (mide sobre ellas) y la interfaz
  del Sprint 2

## Fase 1 — Requisitos

### Objetivo

Que en la planificación el equipo pueda armar el Sprint Backlog: elegir del Product Backlog
las historias que se compromete a hacer en el sprint abierto, y poder sacar una mientras el
sprint siga abierto si el equipo se dio cuenta de que no entra.

### Qué queda afuera

- **Crear el sprint**, que es la #8. Acá ya existe y está abierto.
- **Cerrar el sprint** y decidir qué se completó y qué no, que es la #10.
- **Consultar los sprints cerrados** y sus historias, que es la #11.
- **Los story points planificados, los completados y la velocidad**, que son la #16. Esta
  historia deja registrado qué historias se comprometieron; los números los saca la #16.
- **Estimar una historia**, que es la #12. Acá solo se exige que ya esté estimada.
- **Ordenar el backlog por prioridad** para elegir cuál asignar, que es la #7.
- **Controlar cuánto entra en el sprint**: no hay límite de story points ni de cantidad de
  historias. Cuánto se compromete lo decide el equipo en la planning, no el sistema.
- **Asignarle la historia a una persona**: el Sprint Backlog es del equipo.

### Entradas

| Dato | Obligatorio | Descripción |
|---|---|---|
| Proyecto | sí | en qué proyecto se está trabajando |
| Historia | sí | cuál es la historia que se asigna o se quita, por su identificador |

**A qué sprint se asigna no es una entrada**: siempre es el sprint abierto del proyecto
(RN-1). Como solo puede haber uno abierto a la vez (#8, RN-8), no hay nada que elegir.

### Salidas

La historia queda en el Sprint Backlog del sprint abierto y pasa a estar En curso. Al
quitarla, sale del Sprint Backlog y vuelve al Product Backlog como Pendiente. En cualquier
momento se puede consultar el Sprint Backlog del sprint abierto.

### Reglas de negocio

1. **[RN-1]** Cuando se asigna una historia, el sistema debe agregarla al Sprint Backlog del
   sprint abierto del proyecto.
2. **[RN-2]** Si el proyecto no tiene ningún sprint abierto, el sistema debe rechazar la
   asignación y la desasignación e informar que el proyecto no tiene un sprint abierto.
3. **[RN-3]** Si la historia no está estimada, el sistema debe rechazar la asignación e
   informar que solo se asignan historias estimadas.
4. **[RN-4]** Cuando se asigna una historia, el sistema debe dejarla en estado **En curso**.
5. **[RN-5]** Si la historia ya está en el Sprint Backlog, el sistema debe rechazar la
   asignación e informar que la historia ya está asignada al sprint.
6. **[RN-6]** Si la historia está Terminada, el sistema debe rechazar la asignación e
   informar que una historia terminada no se modifica.
7. **[RN-7]** Mientras el sprint esté abierto, el sistema debe permitir quitarle una
   historia: sale del Sprint Backlog y vuelve al Product Backlog en estado **Pendiente**.
8. **[RN-8]** Si se intenta quitar una historia que no está en el Sprint Backlog, el sistema
   debe rechazar la operación e informar que la historia no está en el sprint.
9. **[RN-9]** El sistema debe permitir consultar el Sprint Backlog del sprint abierto, en el
   orden en que se asignaron las historias. Un sprint sin historias devuelve un Sprint
   Backlog vacío, y eso no es un error.
10. **[RN-10]** El sistema debe asignar únicamente historias del mismo proyecto que el
    sprint, y mantener el Sprint Backlog de cada sprint por separado.
11. **[RN-11]** Mientras un sprint esté cerrado, el sistema no debe permitir agregarle ni
    quitarle historias.
12. **[RN-12]** El sistema debe permitir asignar al sprint abierto una historia que ya estuvo
    en un sprint cerrado y no se completó, porque volvió al Product Backlog.
13. **[RN-13]** Si se intenta asignar o quitar una historia que no existe en el proyecto, el
    sistema debe informar que no se encontró la historia.
14. **[RN-14]** Si una asignación o una desasignación se rechaza, el sistema debe dejar el
    Sprint Backlog y la historia como estaban.
15. **[RN-15]** El sistema debe aceptar cualquier cantidad de historias en el Sprint Backlog,
    sin límite de cantidad ni de story points.
16. **[RN-16]** Si se intenta asignar o quitar una historia en un proyecto que no existe, el
    sistema debe informar que no se encontró el proyecto y no cambiar nada.

Los estados de una historia —Pendiente, En curso y Terminada— son los que define la
historia #6. Esta historia usa dos de las tres transiciones: Pendiente a En curso al
asignar, y En curso a Pendiente al quitar. La que falta, a Terminada, es de la #10.

### Casos límite

- Sprint abierto recién creado: el Sprint Backlog está vacío y eso no es un error.
- Asignar la primera historia, y quitar la única que había: el Sprint Backlog vuelve a
  quedar vacío.
- Historia estimada en 1, el valor más chico de la escala: se asigna igual. No se mira
  cuánto vale la estimación, solo que exista.
- Historia a la que se le quitó la estimación (#12, RN-7): queda sin estimar y no se puede
  asignar.
- Historia que estuvo en un sprint cerrado sin completarse: se puede volver a asignar.
- Historia Terminada en un sprint cerrado: no se puede volver a asignar nunca.
- Quitar una historia y volver a asignarla en el mismo sprint: es válido, y queda última en
  el orden del Sprint Backlog.
- Dos proyectos, cada uno con su sprint abierto: asignar en uno no cambia el otro.
- Sprint abierto con todas las historias del Product Backlog asignadas: es válido, el
  Product Backlog no queda "vacío", sus historias quedan En curso.

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| El proyecto no tiene sprint abierto | rechaza e informa que no hay un sprint abierto |
| La historia no está estimada | rechaza e informa que solo se asignan historias estimadas |
| La historia ya está en el Sprint Backlog | rechaza e informa que ya está asignada |
| La historia está Terminada | rechaza e informa que una historia terminada no se modifica |
| Se quiere quitar una historia que no está en el sprint | rechaza e informa que no está en el sprint |
| La historia no existe en el proyecto | informa que no se encontró la historia |
| El proyecto no existe | informa que no se encontró el proyecto y no cambia nada |

### Criterios de aceptación

- [ ] Se puede asignar una historia estimada al sprint abierto, y queda en el Sprint Backlog.
- [ ] Al asignarla, la historia queda En curso.
- [ ] No se puede asignar una historia sin estimar.
- [ ] No se puede asignar dos veces la misma historia al mismo sprint.
- [ ] No se puede asignar una historia Terminada.
- [ ] No se puede asignar ni quitar nada si el proyecto no tiene un sprint abierto.
- [ ] Se puede quitar una historia del sprint mientras el sprint esté abierto, y vuelve al
      Product Backlog en estado Pendiente.
- [ ] No se puede quitar una historia que no está en el sprint.
- [ ] Se puede consultar el Sprint Backlog del sprint abierto, en el orden en que se
      asignaron las historias, y si está vacío eso no es un error.
- [ ] Una historia que quedó sin completar en un sprint cerrado se puede asignar al sprint
      abierto siguiente.
- [ ] Una operación rechazada no cambia el Sprint Backlog ni el estado de la historia.
- [ ] Asignar o quitar en un proyecto o sobre una historia que no existe informa que no se
      encontró.

### Escenarios BDD

```gherkin
Escenario: Asignar una historia estimada al sprint abierto
  Dado un proyecto con el sprint 1 abierto y una historia estimada en 5, en estado "Pendiente"
  Cuando asigno la historia al sprint
  Entonces la historia queda en el Sprint Backlog del sprint 1
  Y la historia queda en estado "En curso"

Escenario: El Sprint Backlog respeta el orden en que se asignaron las historias
  Dado un proyecto con el sprint 1 abierto y tres historias estimadas
  Cuando las asigno en el orden 3, 1, 2
  Entonces el Sprint Backlog del sprint 1 tiene las historias 3, 1 y 2 en ese orden

Escenario: Un sprint recien abierto tiene el Sprint Backlog vacio
  Dado un proyecto con el sprint 1 abierto y sin historias asignadas
  Cuando consulto el Sprint Backlog del sprint
  Entonces el Sprint Backlog esta vacio
  Y eso no es un error

Escenario: Rechazar una historia sin estimar
  Dado un proyecto con el sprint 1 abierto y una historia sin estimar
  Cuando intento asignar la historia al sprint
  Entonces la operacion se rechaza
  Y se informa que solo se asignan historias estimadas
  Y el Sprint Backlog sigue vacio
  Y la historia sigue en estado "Pendiente"

Escenario: Rechazar una historia que ya esta en el sprint
  Dado un proyecto con el sprint 1 abierto y la historia 1 ya asignada
  Cuando intento asignar otra vez la historia 1
  Entonces la operacion se rechaza
  Y se informa que la historia ya esta asignada al sprint
  Y el Sprint Backlog sigue teniendo una sola historia

Escenario: Rechazar una historia terminada
  Dado un proyecto con el sprint 2 abierto y una historia en estado "Terminada"
  Cuando intento asignar la historia al sprint 2
  Entonces la operacion se rechaza
  Y se informa que una historia terminada no se modifica
  Y el Sprint Backlog sigue vacio

Esquema del escenario: Operar sin un sprint abierto
  Dado un proyecto sin ningun sprint abierto y una historia estimada
  Cuando intento <operacion> la historia
  Entonces la operacion se rechaza
  Y se informa que el proyecto no tiene un sprint abierto

  Ejemplos:
    | operacion |
    | "asignar" |
    | "quitar"  |

Escenario: Quitar una historia del sprint abierto
  Dado un proyecto con el sprint 1 abierto y la historia 1 asignada y en estado "En curso"
  Cuando quito la historia 1 del sprint
  Entonces el Sprint Backlog del sprint 1 queda vacio
  Y la historia 1 queda en estado "Pendiente"

Escenario: Quitar una historia y volver a asignarla
  Dado un proyecto con el sprint 1 abierto y las historias 1 y 2 asignadas en ese orden
  Cuando quito la historia 1 y la asigno de nuevo
  Entonces el Sprint Backlog del sprint 1 tiene las historias 2 y 1 en ese orden

Escenario: Rechazar quitar una historia que no esta en el sprint
  Dado un proyecto con el sprint 1 abierto y una historia estimada sin asignar
  Cuando intento quitar la historia del sprint
  Entonces la operacion se rechaza
  Y se informa que la historia no esta en el sprint
  Y la historia sigue en estado "Pendiente"

Escenario: Una historia que quedo sin completar se puede asignar al sprint siguiente
  Dado un proyecto cuyo sprint 1 se cerro con la historia 1 sin completar
  Y el sprint 2 abierto
  Cuando asigno la historia 1 al sprint 2
  Entonces la historia 1 queda en el Sprint Backlog del sprint 2
  Y la historia 1 queda en estado "En curso"
  Y el sprint 1 sigue teniendo a la historia 1 entre las que se le asignaron

Escenario: Cada proyecto arma su propio Sprint Backlog
  Dado un proyecto A y un proyecto B, cada uno con su sprint abierto y sus historias
  Cuando asigno una historia en el proyecto A
  Entonces el Sprint Backlog del sprint abierto del proyecto B sigue vacio

Esquema del escenario: Operar sobre algo que no existe
  Dado un proyecto con el sprint 1 abierto
  Cuando intento asignar la historia <caso>
  Entonces se informa que no se encontro <que>
  Y el Sprint Backlog sigue vacio

  Ejemplos:
    | caso                                | que           |
    | "99, que no existe en el proyecto"  | "la historia" |
    | "1 en el proyecto 99, que no existe"| "el proyecto" |
```

**Para decidir en la puerta:**

1. **La asignación es siempre al sprint abierto; no se elige a qué sprint.** La issue pide
   que una historia no esté en dos sprints abiertos a la vez, y con la regla de la #8 —un
   solo sprint abierto por proyecto— eso no puede pasar: no hay un segundo sprint abierto al
   que asignarla. Por eso no hay una validación propia para esa condición, y lo que se
   verifica es lo que sí puede pasar: que una historia esté en un sprint cerrado y en el
   abierto. Si alguna vez se permiten dos sprints abiertos, hay que agregar esa validación.
2. **Asignar dos veces la misma historia es un error (RN-5), no una operación que no hace
   nada.** Es el mismo criterio que usó la #4 con los integrantes repetidos. Si el equipo
   prefiere que no moleste, se saca el error y la operación pasa a no hacer nada.
3. **Quitar una historia la deja Pendiente (RN-7), sin recordar en qué estado estaba antes.**
   Como la única forma de entrar al sprint es estando Pendiente, hoy es lo mismo; si más
   adelante aparece otro estado previo, la regla hay que revisarla.
4. **La historia que se quita del sprint abierto desaparece del Sprint Backlog como si nunca
   hubiera estado.** No queda registro de que estuvo y se sacó. Para el cierre y para las
   métricas lo que vale es lo que quedó comprometido al cerrar, no los cambios de opinión
   durante la planning.
5. **Lo que esta historia le deja a la #10 y a la #16:** la lista de historias asignadas al
   sprint es el registro de lo que se comprometió y **no se vacía al cerrar**. Una historia
   que vuelve al Product Backlog sigue figurando en el sprint donde se la había
   comprometido; si no, la #16 no podría decir cuántos puntos se habían planificado.

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios: [pendiente]

## Fase 2 — Diseño

Las decisiones transversales del proyecto están en [AGENTS.md](../../../AGENTS.md) y no se
repiten acá. Los datos del sprint y los nombres de sus operaciones están en
[el contrato del dominio](../../diseno-dominio.md). Esta fase se apoya además en la Fase 2 de
la #8 (el sprint dentro del proyecto y la consulta del sprint abierto) y en la de la #6 (la
historia, su estado y el backlog).

### Enfoque

Esta historia **no crea ninguna entidad**: le agrega comportamiento al sprint que define la
#8 y a la historia que define la #6. El Sprint Backlog no es una cosa aparte: es la lista de
historias que el sprint ya tiene entre sus datos, y que hasta ahora nacía vacía.

Las dos operaciones —asignar y quitar— se le piden al proyecto, no al sprint, porque hacen
falta las dos puntas: el sprint abierto y la historia del Product Backlog, y las dos cuelgan
del proyecto.

Las dos siguen el mismo orden de controles, de lo más general a lo más particular:

1. **Hay un sprint abierto** en el proyecto (RN-2). Sin eso no hay nada que hacer.
2. **La historia existe** en el proyecto (RN-13).
3. **La historia se puede mover**: no está Terminada (RN-6).
4. **La historia está estimada** (RN-3), solo al asignar.
5. **Está o no está en el Sprint Backlog**, según corresponda (RN-5, RN-8).

Recién después se aplica el cambio: agregar o sacar el identificador de la lista del sprint y
cambiar el estado de la historia. Como los dos cambios van juntos y al final, una operación
rechazada deja todo como estaba (RN-14) sin tener que deshacer nada.

### Archivos afectados

| Archivo | Se crea o se modifica | Para qué |
|---|---|---|
| `internal/domain/sprint_backlog.go` | se crea | asignar y quitar una historia del sprint abierto, y consultar el Sprint Backlog |
| `internal/domain/sprint_backlog_test.go` | se crea | los tests de los escenarios BDD de la Fase 1, nombrados como cada escenario |
| `internal/domain/errors.go` | se modifica | los cuatro errores nuevos de esta historia |

### Decisiones

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| Lo de esta historia va en `sprint_backlog.go` y no en `sprint.go` | agregarlo al archivo que crea la #8 | cuatro historias tocan el sprint en dos sprints seguidos; separar por funcionalidad es lo que ya se aprobó en la Puerta 2 de la #12 para que dos personas no editen el mismo archivo |
| Las dos operaciones se le piden al proyecto | pedírselas al sprint | el sprint no conoce el Product Backlog y no puede buscar la historia ni cambiarle el estado; además el sprint destino siempre es el abierto, y eso lo sabe el proyecto |
| El sprint guarda identificadores de historia, y quien quiera el detalle se lo pide al Product Backlog | guardar las historias enteras dentro del sprint | lo fija el contrato del dominio, y guardar la historia en dos lugares deja dos fuentes de verdad que se pueden contradecir |
| La historia **no** guarda a qué sprint pertenece | agregarle a la historia el sprint actual | sería el mismo dato dos veces, y la #16 necesita el registro por sprint (qué se comprometió en cada uno), no un dato por historia que solo sirve para el último |
| Orden de los controles: sprint abierto → existe → terminada → estimada → ya asignada | validar primero la estimación, que es la regla más propia de esta historia | no tiene sentido mirar la estimación de una historia que no existe o que no se puede mover, y así el mensaje que sale es siempre el del primer problema real |
| Quitar es una operación propia | una sola operación que recibe el sprint destino, o vacío para sacarla | son dos reglas distintas con dos errores distintos (RN-5 y RN-8); juntarlas obliga a adivinar cuál se quiso hacer |
| Se reutiliza el error de historia terminada que ya declaró la #6 | un error propio "no se puede asignar una historia terminada" | es el mismo caso y el mismo mensaje; con dos errores para lo mismo, quien consume tiene que acordarse de mirar los dos |
| El error de "no hay sprint abierto" lo declara esta historia y lo reusa la #10 | que cada historia declare el suyo | es la misma situación para las dos, y es la primera que lo necesita |

### Dependencias y restricciones

- **Depende de la #8** (el sprint, su lista de historias y la consulta del sprint abierto),
  de la **#6** (la historia, su estado y el backlog) y de la **#12** (que la historia esté
  estimada). Como manda el plan del 20/10, **no se espera a que estén en `dev`**: los tests
  arman el sprint y las historias dentro del paquete, contra el contrato.
- **El dato de las historias del sprint lo declara la #8**; esta historia solo lo llena y lo
  vacía. No se cambia su nombre ni su forma.
- **No se toca el cierre.** El pase a Cerrado, la fecha de cierre y el registro de lo
  completado son de la #10.
- **La lista de historias asignadas no se vacía al cerrar el sprint.** La #10 y la #16
  dependen de eso.
- **`errors.go` es compartido**: los errores nuevos se agregan al final para achicar el
  conflicto con las otras historias del Sprint 2.

> **Puerta 2** — ¿el diseño es viable y coherente con el resto del proyecto?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios: [pendiente]

## Fase 3 — Tareas

Cada tarea es un ciclo de TDD completo: el test que falla en un commit, la implementación que
lo hace pasar en otro. Van en el orden en que hay que hacerlas, y cada una deja el código
compilando y los tests en verde.

Los tests siguen la convención de [AGENTS.md](../../../AGENTS.md): un test por escenario, con
el nombre del escenario. Cada tarea de rechazo verifica, además del error, que el Sprint
Backlog y el estado de la historia quedaron como estaban (RN-14).

| # | Tarea | Escenario BDD | Reglas | Estado |
|---|---|---|---|---|
| 1 | El Sprint Backlog de un sprint recién abierto está vacío, y eso no es un error | Un sprint recien abierto tiene el Sprint Backlog vacio | RN-9 | pendiente |
| 2 | Asignar una historia estimada la deja en el Sprint Backlog y En curso | Asignar una historia estimada al sprint abierto | RN-1, RN-4 | pendiente |
| 3 | El Sprint Backlog conserva el orden en que se asignaron las historias | El Sprint Backlog respeta el orden en que se asignaron las historias | RN-9, RN-15 | pendiente |
| 4 | Rechazar una historia sin estimar; el Sprint Backlog y la historia quedan como estaban | Rechazar una historia sin estimar | RN-3, RN-14 | pendiente |
| 5 | Rechazar una historia ya asignada; el Sprint Backlog no la repite | Rechazar una historia que ya esta en el sprint | RN-5, RN-14 | pendiente |
| 6 | Rechazar una historia Terminada, con el error que ya existe para ese caso | Rechazar una historia terminada | RN-6, RN-14 | pendiente |
| 7 | Asignar o quitar sin sprint abierto se rechaza e informa que no hay sprint abierto | Operar sin un sprint abierto | RN-2 | pendiente |
| 8 | Quitar una historia la saca del Sprint Backlog y la deja Pendiente | Quitar una historia del sprint abierto | RN-7 | pendiente |
| 9 | Quitar una historia y volver a asignarla la deja última en el orden | Quitar una historia y volver a asignarla | RN-7, RN-9 | pendiente |
| 10 | Rechazar quitar una historia que no está en el Sprint Backlog | Rechazar quitar una historia que no esta en el sprint | RN-8, RN-14 | pendiente |
| 11 | Una historia que quedó sin completar en un sprint cerrado se asigna al sprint siguiente, y el sprint cerrado la sigue teniendo entre las suyas | Una historia que quedo sin completar se puede asignar al sprint siguiente | RN-11, RN-12 | pendiente |
| 12 | Asignar en un proyecto no cambia el Sprint Backlog de otro | Cada proyecto arma su propio Sprint Backlog | RN-10 | pendiente |
| 13 | Asignar o quitar una historia o un proyecto que no existen informa que no se encontró | Operar sobre algo que no existe | RN-13, RN-16 | pendiente |

### Por qué este orden

- **La 1 va primero** porque la consulta del Sprint Backlog es lo que usan todas las demás
  para verificar el resultado, y en un sprint vacío se puede escribir antes de que exista la
  asignación.
- **La 2 es el camino feliz.** De la 4 a la 7 se agregan los rechazos de a uno, que recién
  ahí tienen contra qué fallar.
- **La 3 va temprano** porque el orden del Sprint Backlog es lo que después verifica la 9.
- **La 8 abre la segunda operación**, y la 9 y la 10 la completan.
- **La 11, la 12 y la 13 van al final** porque miran los bordes —el sprint anterior, otro
  proyecto, algo que no existe— y solo tienen sentido con lo de adentro ya andando.

### Cobertura

Las dieciséis reglas quedan cubiertas: RN-1 (2), RN-2 (7), RN-3 (4), RN-4 (2), RN-5 (5),
RN-6 (6), RN-7 (8, 9), RN-8 (10), RN-9 (1, 3, 9), RN-10 (12), RN-11 (11), RN-12 (11),
RN-13 (13), RN-14 (4, 5, 6, 10), RN-15 (3), RN-16 (13).

| Criterio de aceptación | Tareas |
|---|---|
| Se puede asignar una historia estimada al sprint abierto | 2 |
| Al asignarla, la historia queda En curso | 2 |
| No se puede asignar una historia sin estimar | 4 |
| No se puede asignar dos veces la misma historia | 5 |
| No se puede asignar una historia Terminada | 6 |
| No se puede asignar ni quitar sin un sprint abierto | 7 |
| Se puede quitar una historia mientras el sprint esté abierto, y vuelve a Pendiente | 8, 9 |
| No se puede quitar una historia que no está en el sprint | 10 |
| Se puede consultar el Sprint Backlog, en orden, y vacío no es un error | 1, 3 |
| Una historia sin completar se puede asignar al sprint siguiente | 11 |
| Una operación rechazada no cambia nada | 4, 5, 6, 10 |
| Asignar o quitar sobre algo que no existe informa que no se encontró | 13 |

Los trece escenarios de la Fase 1 tienen su tarea en la columna "Escenario BDD".

### Dependencias

- **Todas las tareas necesitan la #8 y la #6**: el sprint abierto y la historia con su
  estado. Mientras no estén en `dev`, los tests las arman dentro del paquete.
- **Las tareas 6 y 11 necesitan historias Terminadas y un sprint cerrado**, que son de la
  #10. Los tests los preparan desde dentro del paquete, igual que hizo la #12.
- **La tarea 4 necesita una historia sin estimar**, que es el estado en el que nace toda
  historia (#6, RN-7): no hace falta nada de la #12 para prepararla.

> **Puerta 3** — ¿las tareas cubren todos los criterios de aceptación?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios: [pendiente]

## Fase 4 — Implementación

_Se completa durante la implementación._

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #9 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | `internal/domain/sprint_backlog_test.go` |
| Código Go | `internal/domain/sprint_backlog.go` y sus errores en `internal/domain/errors.go` |
