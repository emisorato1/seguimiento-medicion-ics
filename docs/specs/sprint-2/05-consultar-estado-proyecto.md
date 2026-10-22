# Spec: Consultar el estado de un proyecto

- **Historia:** #5
- **Sprint:** 2
- **Responsable:** Juliana Bustos
- **Estado:** Puerta 3 aprobada — **lista para implementar**
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de la consigna y
  de los criterios de aceptación de la issue. El responsable lo revisa, lo corrige si hace
  falta y lo presenta a las puertas. Revisado por la responsable el 20/10 y presentado a las
  tres puertas por pull request.
- **Depende de:** #3 (el proyecto y sus datos), #4 (los integrantes), #6 (el backlog y el
  estado de cada historia) y #8 (el sprint abierto); los cambios de estado de las historias
  los hacen la #9 y la #10
- **La necesitan:** la interfaz, que es donde se muestra; ninguna otra historia la consume

## Fase 1 — Requisitos

### Objetivo

Que cualquier integrante pueda ver de un vistazo en qué anda un proyecto —sus datos, quiénes
trabajan en él, cuánto del backlog está terminado y qué sprint está en curso— sin tener que
revisar sprint por sprint ni historia por historia.

### Qué queda afuera

- **Las métricas**: story points, velocidad, horas, desviación, porcentaje de historias
  completas y defectos. Son la #16, y el dashboard (#17) las muestra. Ver el punto 2 de
  "Para decidir en la puerta".
- **El reporte del proyecto**, que junta todo para la review: es la #18.
- **Los sprints ya cerrados** y su detalle: son la #11.
- **El listado de historias**: el backlog completo es la #6 y ordenado por prioridad, la #7.
  Acá solo se cuentan.
- **Modificar cualquier cosa**: esta historia es solo una consulta.
- **Cómo se muestra en pantalla**: eso es de la interfaz.

### Entradas

| Dato | Obligatorio | Descripción |
|---|---|---|
| Proyecto | sí | de qué proyecto se quiere ver el estado |

### Salidas

El estado del proyecto en el momento de la consulta: sus datos (identificador, nombre,
descripción y fechas), sus integrantes con su rol, el resumen del backlog (historias totales,
completadas y pendientes) y el sprint en curso, o el aviso de que no hay ninguno. Pedir el
estado no cambia nada.

### Reglas de negocio

1. **[RN-1]** Cuando se pide el estado de un proyecto, el sistema debe devolver su
   identificador, su nombre, su descripción, su fecha de inicio y su fecha de finalización.
2. **[RN-2]** Si el proyecto no tiene descripción o no tiene fecha de finalización, el sistema
   debe informar que ese dato no está cargado, y eso no es un error.
3. **[RN-3]** Cuando se pide el estado, el sistema debe devolver los integrantes del proyecto,
   cada uno con su nombre y apellido y su rol, en el orden en que se cargaron.
4. **[RN-4]** Mientras el proyecto no tenga integrantes, el sistema debe devolver la lista de
   integrantes vacía, y eso no es un error.
5. **[RN-5]** Cuando se pide el estado, el sistema debe devolver el resumen del backlog: la
   cantidad total de historias, cuántas están completadas y cuántas están pendientes.
6. **[RN-6]** El sistema debe contar como completada cada historia que está Terminada, y como
   pendiente cada historia que está Pendiente o En curso.
7. **[RN-7]** El sistema debe devolver un resumen en el que el total sea siempre igual a las
   completadas más las pendientes.
8. **[RN-8]** Mientras el proyecto tenga un sprint abierto, el sistema debe incluir en el
   estado su número, su Sprint Goal, su fecha de inicio, su fecha de fin y cuántas historias
   tiene asignadas.
9. **[RN-9]** Mientras el proyecto no tenga ningún sprint abierto, el sistema debe informar
   que no hay sprint en curso, y eso no es un error.
10. **[RN-10]** El sistema debe devolver el estado sin modificar nada: ni el proyecto, ni sus
    integrantes, ni sus historias, ni sus sprints.
11. **[RN-11]** Si se pide el estado de un proyecto que no existe, el sistema debe informar
    que no se encontró el proyecto y no devolver ningún estado.
12. **[RN-12]** El sistema debe armar el estado con lo que el proyecto tiene en el momento de
    la consulta: dos consultas seguidas sin cambios en el medio dan lo mismo, y una consulta
    hecha después de un cambio lo refleja.
13. **[RN-13]** El sistema debe tratar cada proyecto por separado: el estado de un proyecto no
    cuenta integrantes, historias ni sprints de otro.

### Casos límite

- **Proyecto recién creado:** sin integrantes, con el backlog en cero (0 totales, 0
  completadas, 0 pendientes) y sin sprint en curso. Nada de eso es un error.
- **Proyecto sin descripción y sin fecha de finalización:** se informan como no cargadas.
- **Todas las historias terminadas:** pendientes en 0.
- **Ninguna historia terminada:** completadas en 0.
- **Historias En curso:** cuentan como pendientes, no como completadas.
- **Proyecto con sprints, todos cerrados:** no hay sprint en curso.
- **Sprint abierto sin historias asignadas:** se muestra, con 0 historias asignadas.
- **Historia que no se terminó en un sprint y volvió al backlog:** cuenta una sola vez, como
  pendiente.

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| El proyecto no existe | informa que no se encontró el proyecto y no devuelve ningún estado |
| El proyecto no tiene integrantes, historias o sprint abierto | **no es un error**: lo informa vacío o en cero |

### Criterios de aceptación

- [ ] El estado muestra el identificador, el nombre, la descripción y las dos fechas del
      proyecto.
- [ ] Un proyecto sin descripción o sin fecha de finalización las muestra como no cargadas,
      sin error.
- [ ] El estado muestra cada integrante con su nombre y su rol, en el orden en que se cargó.
- [ ] Un proyecto sin integrantes muestra la lista vacía, sin error.
- [ ] El estado muestra la cantidad de historias totales, completadas y pendientes, y el
      total es igual a completadas más pendientes.
- [ ] Solo las historias Terminadas cuentan como completadas; las En curso cuentan como
      pendientes.
- [ ] Si hay un sprint abierto, el estado muestra su número, su goal, sus fechas y cuántas
      historias tiene.
- [ ] Si no hay sprint abierto, el estado lo informa y eso no es un error.
- [ ] Una consulta hecha después de un cambio en el proyecto lo refleja.
- [ ] Consultar el estado no modifica nada del proyecto.
- [ ] El estado de un proyecto no incluye nada de otro.
- [ ] Pedir el estado de un proyecto que no existe informa que no se encontró.

### Escenarios BDD

```gherkin
Escenario: Consultar el estado de un proyecto recien creado
  Dado un proyecto "Sistema de metricas" con inicio 01/10/2026 y fin 30/11/2026, sin integrantes, historias ni sprints
  Cuando pido el estado del proyecto
  Entonces veo su identificador, su nombre, su descripcion y sus dos fechas
  Y la lista de integrantes esta vacia
  Y el backlog tiene 0 historias totales, 0 completadas y 0 pendientes
  Y se informa que no hay sprint en curso
  Y no se informa ningun error

Escenario: Un proyecto sin descripcion ni fecha de finalizacion
  Dado un proyecto registrado solo con nombre y fecha de inicio
  Cuando pido el estado del proyecto
  Entonces la descripcion y la fecha de finalizacion figuran como no cargadas

Escenario: El estado muestra los integrantes con su rol
  Dado un proyecto con "Emiliano Sorato" como Agile Enabler y "Juliana Bustos" como Product Builder, cargados en ese orden
  Cuando pido el estado del proyecto
  Entonces veo 2 integrantes, en ese orden
  Y cada uno con su rol

Esquema del escenario: Resumen del backlog segun el estado de las historias
  Dado un proyecto con <pendientes sin empezar> historias "Pendiente", <en curso> "En curso" y <terminadas> "Terminada"
  Cuando pido el estado del proyecto
  Entonces el backlog tiene <total> historias totales, <completadas> completadas y <pendientes> pendientes

  Ejemplos:
    | pendientes sin empezar | en curso | terminadas | total | completadas | pendientes |
    | 3                      | 0        | 0          | 3     | 0           | 3          |
    | 1                      | 2        | 0          | 3     | 0           | 3          |
    | 1                      | 1        | 2          | 4     | 2           | 2          |
    | 0                      | 0        | 3          | 3     | 3           | 0          |

Escenario: El estado muestra el sprint en curso
  Dado un proyecto con el sprint 1 cerrado y el sprint 2 abierto con el goal "Tener el dashboard" y 3 historias asignadas
  Cuando pido el estado del proyecto
  Entonces el sprint en curso es el 2, con el goal "Tener el dashboard" y sus fechas
  Y tiene 3 historias asignadas

Escenario: Un proyecto con todos sus sprints cerrados no tiene sprint en curso
  Dado un proyecto cuyos sprints 1 y 2 estan cerrados
  Cuando pido el estado del proyecto
  Entonces se informa que no hay sprint en curso
  Y no se informa ningun error

Escenario: El estado refleja los cambios del proyecto
  Dado un proyecto con 2 historias "Pendiente"
  Y pido el estado del proyecto
  Cuando agrego una historia al backlog
  Y vuelvo a pedir el estado del proyecto
  Entonces el backlog tiene 3 historias totales y 3 pendientes

Escenario: Consultar el estado no cambia nada
  Dado un proyecto con 2 integrantes, 3 historias y el sprint 1 abierto
  Cuando pido el estado del proyecto dos veces seguidas
  Entonces las dos veces recibo lo mismo
  Y el proyecto sigue con sus 2 integrantes, sus 3 historias y el sprint 1 abierto

Escenario: Cada proyecto muestra solo lo suyo
  Dado un proyecto A con 2 integrantes y 3 historias, y un proyecto B sin integrantes ni historias
  Cuando pido el estado del proyecto B
  Entonces la lista de integrantes esta vacia
  Y el backlog tiene 0 historias totales

Escenario: Consultar el estado de un proyecto que no existe
  Dado que no existe ningun proyecto con identificador 99
  Cuando pido el estado del proyecto 99
  Entonces se informa que no se encontro el proyecto
  Y no recibo ningun estado
```

**Para decidir en la puerta:**

1. **RN-6 — las historias En curso cuentan como pendientes.** La issue pide "totales,
   completadas y pendientes", y con tres estados alguno tiene que quedar adentro de otro. Se
   eligió que "pendiente" quiera decir "todavía no terminada", que es lo que le importa a
   quien pregunta cómo viene el proyecto. Si el equipo quiere ver las En curso aparte, es un
   número más en el resumen y cambia la RN-6, pero ya no se cumpliría la RN-7 tal como está.
   **Decidido el 21/10: las En curso cuentan como pendientes.**
2. **No se muestra el porcentaje de historias completas.** Sale fácil de los dos números, pero
   es una de las métricas de la #16, y el equipo decidió que las métricas se calculen en un
   solo lugar. Si se lo agrega acá, hay dos cálculos que pueden dar distinto.
   **Decidido el 21/10: no se muestra; el porcentaje queda en la #16.**
3. **Las completadas se cuentan por el estado actual de cada historia**, no por lo que quedó
   registrado al cerrar cada sprint. Para el proyecto entero da lo mismo: cada historia está
   una sola vez en el backlog, así que no se cuenta dos veces aunque haya pasado por dos
   sprints. Lo registrado al cerrar importa para medir sprint por sprint, que es la #16.
   **Decidido el 21/10: se acepta como está.**
4. **Del sprint en curso se muestra cuántas historias tiene, no cuáles.** El detalle de un
   sprint ya existe (#9 para el abierto, #11 para los cerrados).
   **Decidido el 21/10: se acepta como está.**
5. **Los integrantes salen en el orden en que se cargaron**, que es como los devuelve la #4.
   No se ordenan por rol ni por nombre.
   **Decidido el 21/10: se acepta como está.**
6. **La tarea 4 espera al `Sprint` de la #8 en `dev`.** El borrador decía que su test "arma
   el sprint dentro del paquete" mientras la #8 no esté, pero en Go un test no puede armar un
   sprint si el tipo todavía no existe: no compila. Las tareas 1 a 3 no lo necesitan, porque
   el estado guarda una copia propia de los datos del sprint y no el sprint en sí. La tarea 4
   se programa cuando estén en `dev` el `Sprint` y la consulta del sprint abierto de la #8,
   que se le pidió adelantar a Agustín.
   **Decidido el 21/10: las tareas 1 a 3 arrancan ya, y la 4 espera el PR temprano de la
   #8.** Cómo arma el test los sprints que necesita está en las Dependencias de la Fase 3.

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: Maximiliano Eula — Fecha: 2026-10-20 — Comentarios: aprobada en el PR #102 (https://github.com/emisorato1/seguimiento-medicion-ics/pull/102#issuecomment-6303513354), con dos aclaraciones. Los ajustes se aplicaron el 21/10 en ae10336, y Maxi verificó el 21/10 que quedan cumplidas las condiciones (https://github.com/emisorato1/seguimiento-medicion-ics/pull/102#issuecomment-6314721364).

## Fase 2 — Diseño

Las decisiones transversales del proyecto están en [AGENTS.md](../../../AGENTS.md) y no se
repiten acá. Los datos del sprint están en [el contrato del dominio](../../diseno-dominio.md).
Esta fase registra solo lo propio de la consulta.

### Enfoque

Esta historia **no cambia nada**: junta en una sola respuesta lo que el proyecto ya sabe. No
hay que normalizar, ni validar, ni aplicar. Son cuatro lecturas, y todas existen o están en
el contrato:

1. **Los datos del proyecto**, de la #3.
2. **Los integrantes**, de la #4.
3. **El resumen del backlog**, que se cuenta recorriendo las historias de la #6 y mirando el
   estado de cada una (RN-5 a RN-7).
4. **El sprint en curso**, con la consulta del sprint abierto que define la #8.

El estado es **una foto**: un resultado plano que se arma en el momento, como las métricas de
la #16, con copias de los datos y no con las entidades guardadas. No es una entidad
compartida, porque solo la consume la interfaz, y por eso no está en el contrato.

El proyecto que no existe lo informa el registro de la #3 al buscarlo, con el error que ya
existe. Esta historia no agrega errores.

### Archivos afectados

| Archivo | Se crea o se modifica | Para qué |
|---|---|---|
| `internal/domain/project_status.go` | se crea | la consulta que arma el estado: datos, integrantes, resumen del backlog y sprint en curso |
| `internal/domain/project_status_test.go` | se crea | los tests de los escenarios BDD de la Fase 1, nombrados como cada escenario |

### Decisiones

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| El estado es una foto con copias de los datos | devolver el proyecto, sus integrantes y su sprint guardados, más los números | mezclar datos vivos con conteos congelados hace que el estado se contradiga solo si el proyecto cambia después; y una copia no deja modificar nada desde afuera (RN-10) |
| El resumen se cuenta en cada consulta | llevar los contadores guardados en el proyecto y actualizarlos con cada cambio | un número guardado se desactualiza en silencio el día que una historia nueva (la #9, la #10) cambie un estado y no se acuerde de sumarlo; contando, la RN-12 sale sola |
| El resumen trae números, no la lista de historias | incluir las historias completas | el listado ya existe en la #6 y la #7; repetirlo acá agranda la consulta sin agregar nada |
| El sprint en curso se pide con la consulta del sprint abierto de la #8 | recorrer los sprints buscando el que está abierto | cuál es el sprint abierto es una regla de la #8; repetirla acá la duplica |
| Una sola consulta trae las cuatro partes | una consulta por parte | la issue pide una foto del proyecto; con cuatro sueltas, la interfaz tendría que armarla, y entre una y otra el proyecto podría cambiar |
| La consulta va en `project_status.go` | agregarla a `project.go` | en este sprint la #8 y la #9 también tocan el proyecto; un archivo propio evita que se pisen, como ya se hizo con la estimación de la #12 |

### Dependencias y restricciones

- **La #3, la #4 y la #6 ya están en `dev`**: los datos, los integrantes y el backlog se leen
  con lo que existe.
- **El sprint abierto es de la #8**, que todavía no está en `dev`. Las tareas 1 a 3 no lo
  necesitan; la tarea 4 sí, y espera a que el `Sprint` y la consulta del sprint abierto estén
  en `dev` (punto 6 de la puerta).
- **Los estados En curso y Terminada los ponen la #9 y la #10.** Mientras no existan, los
  tests fijan el estado de la historia desde dentro del paquete, como ya hizo la #12.
- **No se modifica ningún archivo existente.** Esta historia solo lee.
- **No se calcula ninguna métrica**: ni porcentajes, ni puntos, ni horas. Son de la #16.

> **Puerta 2** — ¿el diseño es viable y coherente con el resto del proyecto?
> Aprobó: Maximiliano Eula — Fecha: 2026-10-20 — Comentarios: aprobada en el PR #102 (https://github.com/emisorato1/seguimiento-medicion-ics/pull/102#issuecomment-6303513354), con dos aclaraciones. Los ajustes se aplicaron el 21/10 en ae10336, y Maxi verificó el 21/10 que quedan cumplidas las condiciones (https://github.com/emisorato1/seguimiento-medicion-ics/pull/102#issuecomment-6314721364).

## Fase 3 — Tareas

Cada tarea es un ciclo de TDD completo: el test que falla en un commit, la implementación que
lo hace pasar en otro. Van en el orden en que hay que hacerlas, y cada una deja el código
compilando y los tests en verde.

Los tests siguen la convención de [AGENTS.md](../../../AGENTS.md): un test por escenario, con
el nombre del escenario.

**Hay pocas tareas, y algunas cubren varios escenarios, a propósito.** En la #12 seis tareas
pasaron en verde sin código nuevo porque separaban en ciclos lo que salía del mismo código, y
su Fase 4 lo anotó como un error de la Fase 3. Acá cada tarea agrupa los escenarios que
dependen de la misma implementación, y cada una tiene al menos un test que falla antes de
escribirla.

| # | Tarea | Escenario BDD | Reglas | Estado |
|---|---|---|---|---|
| 1 | El estado de un proyecto recién creado trae sus datos, la lista de integrantes vacía, el backlog en cero y ningún sprint en curso; la descripción y la fecha de fin vacías figuran como no cargadas; un proyecto inexistente informa que no se encontró | Consultar el estado de un proyecto recien creado; Un proyecto sin descripcion ni fecha de finalizacion; Consultar el estado de un proyecto que no existe | RN-1, RN-2, RN-4, RN-9, RN-11 | hecha |
| 2 | El estado trae los integrantes con su nombre y su rol, en el orden en que se cargaron | El estado muestra los integrantes con su rol | RN-3 | hecha |
| 3 | El resumen cuenta totales, completadas y pendientes en cada consulta, las En curso cuentan como pendientes, solo cuenta lo del proyecto y consultar no cambia nada | Resumen del backlog segun el estado de las historias; El estado refleja los cambios del proyecto; Cada proyecto muestra solo lo suyo; Consultar el estado no cambia nada | RN-5, RN-6, RN-7, RN-10, RN-12, RN-13 | hecha |
| 4 | El estado trae el sprint abierto con su número, su goal, sus fechas y cuántas historias tiene; con todos los sprints cerrados informa que no hay sprint en curso | El estado muestra el sprint en curso; Un proyecto con todos sus sprints cerrados no tiene sprint en curso | RN-8, RN-9 | hecha |

**Dónde falla cada una antes de implementarse.** La 1 falla porque la consulta no existe. La
2, porque la tarea 1 devuelve los integrantes vacíos. La 3, porque el resumen está en cero:
los tests de "no cambia nada" y de "cada proyecto" pasarían solos, y por eso van con el
conteo y no en una tarea propia. La 4, porque el estado todavía no trae ningún sprint.

### Por qué este orden

- **La 1 va primero** porque crea la consulta con todo vacío: es contra lo que se escriben
  las demás. El proyecto inexistente va acá porque no agrega código, ya lo resuelve el
  registro de la #3.
- **La 2 y la 3 son lo que ya está en `dev`**: integrantes y backlog. Se pueden programar
  apenas pase la Puerta 3.
- **La 4 va al final** porque es la única que depende de una historia que todavía no está en
  `dev` (la #8). Si la #8 llega antes, se puede adelantar.

### Cobertura

Las trece reglas quedan cubiertas: RN-1 (1), RN-2 (1), RN-3 (2), RN-4 (1), RN-5 (3),
RN-6 (3), RN-7 (3), RN-8 (4), RN-9 (1, 4), RN-10 (3), RN-11 (1), RN-12 (3), RN-13 (3).

| Criterio de aceptación | Tareas |
|---|---|
| El estado muestra el identificador, el nombre, la descripción y las fechas del proyecto | 1 |
| Sin descripción o sin fecha de fin, figuran como no cargadas | 1 |
| Cada integrante con su nombre y su rol, en orden | 2 |
| Un proyecto sin integrantes muestra la lista vacía | 1 |
| Totales, completadas y pendientes, con el total igual a la suma | 3 |
| Solo las Terminadas cuentan como completadas | 3 |
| Con un sprint abierto se ve su número, su goal, sus fechas y sus historias | 4 |
| Sin sprint abierto se informa y no es un error | 1, 4 |
| Una consulta posterior a un cambio lo refleja | 3 |
| Consultar no modifica nada | 3 |
| El estado de un proyecto no incluye nada de otro | 3 |
| Un proyecto inexistente informa que no se encontró | 1 |

Los diez escenarios de la Fase 1 tienen su tarea en la columna "Escenario BDD".

### Dependencias

- **Las tareas 1 a 3 solo necesitan lo que ya está en `dev`** (#3, #4 y #6).
- **La tarea 3 necesita historias En curso y Terminadas**, y esos cambios de estado son de la
  #9 y la #10: el test los fija desde dentro del paquete.
- **La tarea 4 necesita el `Sprint` y la consulta del sprint abierto de la #8** en `dev`. Sin
  el tipo, su test no compila. Se le pidió a Agustín adelantarlos (punto 6 de la puerta);
  mientras tanto se programan las tareas 1 a 3. Maxi confirmó en el PR #104 que
  `p.OpenSprint()` es una consulta del sprint abierto y no la acción de abrir uno.
- **La tarea 4 necesita además un sprint Cerrado y un sprint con 3 historias asignadas.**
  Cerrar es de la #10 y asignar es de la #9, que todavía no existen. Con el `Sprint` en
  `dev`, el test crea los sprints con el alta de la #8 y, desde dentro del paquete, le fija
  el estado Cerrado a uno y le carga los 3 identificadores de historia en su lista de
  historias asignadas al otro. Es lo mismo que se hace con las historias En curso y
  Terminadas, y lo mismo que hace la #8 en sus tareas 10 a 12.

> **Puerta 3** — ¿las tareas cubren todos los criterios de aceptación?
> Aprobó: Maximiliano Eula — Fecha: 2026-10-20 — Comentarios: aprobada en el PR #102 (https://github.com/emisorato1/seguimiento-medicion-ics/pull/102#issuecomment-6303513354), con dos aclaraciones. Los ajustes se aplicaron el 21/10 en ae10336, y Maxi verificó el 21/10 que quedan cumplidas las condiciones (https://github.com/emisorato1/seguimiento-medicion-ics/pull/102#issuecomment-6314721364).

## Fase 4 — Implementación

**La implementó la sesión de Emiliano Sorato, con autorización de Juliana Bustos, el 21/10**,
para adelantar sus historias. La responsable sigue siendo Juliana.

Ciclo RED → GREEN → REFACTOR. El test en un commit y la implementación en otro, con la salida
del test fallando en el mensaje del commit del test, después de `RED:`.

| Tarea | Commit del test | Commit de la implementación |
|---|---|---|
| 1 | `7de777d` | `718ae62` |
| 2 | `d7aa701` | `674ff06` |
| 3 | `a385451` | `431b247` |
| 4 | `27cd558` | `96b9633` |

Las cuatro fallaron antes de su implementación. La 1, por compilación: la consulta no existía.
La 2, la 3 y la 4, por comportamiento: la tarea 1 devolvía los integrantes vacíos, el resumen
en cero y ningún sprint. Cuatro tests pasaron solos, como preveía la Fase 3, y por eso van en
el mismo commit que uno que falla:

- **El proyecto que no existe** (tarea 1): lo resuelve el registro de la #3 al buscarlo.
- **Consultar no cambia nada y cada proyecto muestra solo lo suyo** (tarea 3): la consulta
  nunca escribe y cada proyecto cuenta lo suyo.
- **Todos los sprints cerrados** (tarea 4): antes de la tarea, el estado nunca traía sprint.

Nombres que la Fase 2 no fija: la consulta es `p.Status()`, y devuelve un `ProjectStatus`
con `MemberStatus` (nombre y rol), `BacklogSummary` (`Total`, `Completed`, `Pending`) y
`SprintSummary` (número, goal, fechas y `AssignedStories`). Los datos opcionales y el sprint
en curso llevan su indicador aparte —`HasDescription`, `HasEnd`, `HasCurrentSprint`—, para que
"no está cargado" y "no hay sprint en curso" se informen y no haya que deducirlos de un valor
vacío (RN-2, RN-9). El sprint en curso sale de `p.OpenSprint()` de la #8, como pide la Fase 2.

Los tests fijan el estado En curso y Terminada de las historias, y el sprint Cerrado con sus
historias asignadas, desde dentro del paquete, como dicen las Dependencias de la Fase 3. No
se modificó ningún archivo que ya existía. `internal/domain` sigue en 100% de cobertura.

### Desvíos respecto de la spec

Ninguno.

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #5 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | `internal/domain/project_status_test.go` |
| Código Go | `internal/domain/project_status.go` |
