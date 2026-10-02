# Spec: Administrar el Product Backlog

- **Historia:** #6
- **Sprint:** 1
- **Responsable:** Emiliano Sorato
- **Estado:** Requisitos — **corregido, pendiente de Puerta 1**
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de
  la consigna y de los criterios de aceptación de la issue. El responsable lo
  revisa, lo corrige si hace falta y lo presenta a la Puerta 1. Corregido por el
  responsable y presentado a la Puerta 1 por pull request
- **Depende de:** #3 (el backlog pertenece a un proyecto)
- **La necesitan:** #7 (priorizar), #9 (asignar a sprint), #12 (estimar), #16 (métricas)

## Fase 1 — Requisitos

### Objetivo

Que el equipo pueda cargar y mantener las historias de usuario de un proyecto en un solo
lugar, con todo lo necesario para después estimarlas, planificarlas y medirlas.

### Qué queda afuera

- El orden y la priorización del backlog, que es la historia #7.
- Asignar historias a un sprint, que es la #9.
- Cargar los story points, que es la #12. Acá la historia puede quedar sin estimar.
- Borrar historias: en esta historia solo se agregan y se modifican.

### Entradas

| Dato | Obligatorio | Descripción |
|---|---|---|
| Proyecto | sí | a qué proyecto pertenece la historia |
| Título | sí | nombre corto de la historia |
| Descripción | sí | qué se quiere y para qué |
| Prioridad | sí | Alta, Media o Baja |
| Criterios de aceptación | sí | al menos uno |

Los **story points** no son una entrada de esta historia: la historia nace sin estimar
(RN-7) y el valor lo carga la historia #12. Siguen siendo uno de los datos que se registran
por la RN-2.

### Salidas

Una historia registrada en el backlog del proyecto, con un identificador propio y un estado.
Se puede pedir el backlog completo de un proyecto.

### Reglas de negocio

1. **[RN-1]** El sistema debe asignar a cada historia un identificador único dentro de su
   proyecto, que no se repite y no cambia.
2. **[RN-2]** El sistema debe registrar para cada historia: identificador, título,
   descripción, prioridad, estado, story points y criterios de aceptación.
3. **[RN-3]** Cuando se crea una historia, el sistema debe dejarla en estado **Pendiente**.
4. **[RN-4]** El sistema debe aceptar únicamente las prioridades **Alta**, **Media** y **Baja**.
5. **[RN-5]** Si se intenta crear o modificar una historia sin título, sin descripción, sin
   prioridad o sin al menos un criterio de aceptación, el sistema debe rechazar la operación
   e informar qué falta. Un texto compuesto solo por espacios cuenta como vacío.
6. **[RN-6]** Cuando se registran el título, la descripción y los criterios de aceptación, el
   sistema debe guardarlos sin los espacios del principio y del final, y debe descartar los
   criterios de aceptación que queden vacíos.
7. **[RN-7]** Una historia recién creada debe quedar **sin estimar**: los story points son
   opcionales en esta historia.
8. **[RN-8]** Mientras una historia esté en estado Pendiente o En curso, el sistema debe
   permitir modificar su título, descripción, prioridad y criterios de aceptación.
9. **[RN-9]** Cuando se modifica una historia, el sistema debe cambiar solo los datos que se
   indicaron en esa modificación y dejar el resto como estaba. Los cuatro datos modificables
   son obligatorios, así que ninguno se puede dejar vacío: para cambiar los criterios de
   aceptación se envía la lista completa que reemplaza a la anterior.
10. **[RN-10]** Si se intenta modificar una historia que ya está Terminada, el sistema debe
    rechazar la operación e informar que una historia terminada no se modifica.
11. **[RN-11]** Si una operación se rechaza, el sistema debe dejar la historia y el backlog
    como estaban antes.
12. **[RN-12]** Si se intenta agregar, modificar o listar historias de un proyecto que no
    existe, el sistema debe informar que no se encontró el proyecto.
13. **[RN-13]** Si se intenta modificar una historia que no existe en ese proyecto, el sistema
    debe informar que no se encontró la historia y no crear ninguna.
14. **[RN-14]** El sistema debe permitir que dos historias del mismo proyecto tengan el mismo
    título: lo que las distingue es el identificador.
15. **[RN-15]** Cuando se pide el backlog de un proyecto, el sistema debe devolver todas sus
    historias. Si el proyecto no tiene ninguna, debe devolver una lista vacía, y eso no es un
    error.
16. **[RN-16]** Los identificadores de las historias deben ser únicos dentro de cada proyecto:
    dos proyectos distintos pueden tener cada uno una historia con el mismo identificador.

### Estados de una historia

| Estado | Qué significa |
|---|---|
| Pendiente | está en el backlog, sin comprometer |
| En curso | fue asignada a un sprint abierto |
| Terminada | se completó y el sprint se cerró |

En esta historia solo se crea en **Pendiente**. Los cambios de estado los hacen las historias
#9 (asignar a un sprint) y #10 (cerrar el sprint).

### Casos límite

- Proyecto sin historias: el backlog se devuelve vacío, no es un error.
- Historia con un solo criterio de aceptación: es válida.
- Historia sin story points: es válida y queda marcada como sin estimar.
- Título o descripción con espacios al principio y al final: se guardan sin esos espacios.
- Primera historia de un proyecto: recibe el primer identificador.
- Criterios de aceptación donde alguno viene vacío: se descartan los vacíos; si no queda
  ninguno, la operación se rechaza.
- Modificar una historia sin cambiar nada: es válido y no da error.
- Dos proyectos distintos, cada uno con su primera historia: los identificadores pueden
  coincidir entre proyectos, porque son únicos dentro de cada proyecto.

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| Título vacío o solo espacios | rechaza e informa que el título es obligatorio |
| Descripción vacía | rechaza e informa que la descripción es obligatoria |
| Sin criterios de aceptación | rechaza e informa que hace falta al menos uno |
| Prioridad distinta de las tres permitidas | rechaza e informa cuáles son válidas |
| Todos los criterios de aceptación vienen vacíos | rechaza e informa que hace falta al menos uno |
| Modificar una historia Terminada | rechaza e informa que no se puede modificar |
| Proyecto inexistente | informa que no se encontró el proyecto |
| Historia inexistente al modificar | informa que no se encontró la historia y no crea ninguna |

### Criterios de aceptación

- [ ] Se puede crear una historia con título, descripción, prioridad y criterios de
      aceptación, y queda con identificador propio y estado Pendiente.
- [ ] El identificador no se repite dentro del mismo proyecto.
- [ ] No se puede crear una historia sin título, sin descripción, sin prioridad o sin
      criterios de aceptación.
- [ ] No se puede crear una historia con una prioridad que no sea Alta, Media o Baja.
- [ ] Una historia recién creada queda sin estimar.
- [ ] Se puede modificar una historia mientras no esté Terminada.
- [ ] No se puede modificar una historia Terminada.
- [ ] Se puede pedir el backlog completo de un proyecto.
- [ ] El backlog de un proyecto sin historias se devuelve vacío.
- [ ] Un criterio de aceptación vacío se descarta, y si no queda ninguno la operación se
      rechaza.
- [ ] Una operación rechazada deja la historia y el backlog como estaban.
- [ ] Modificar una historia inexistente informa el error y no crea ninguna.
- [ ] Operar sobre un proyecto inexistente informa que no se encontró.

### Escenarios BDD

```gherkin
Escenario: Crear una historia en el backlog
  Dado un proyecto registrado sin historias
  Cuando creo una historia con titulo "Registrar esfuerzo", prioridad "Alta" y un criterio de aceptacion
  Entonces la historia queda en el backlog con estado "Pendiente"
  Y queda sin estimar
  Y recibe un identificador

Escenario: Los identificadores no se repiten
  Dado un proyecto con una historia ya creada
  Cuando creo una segunda historia
  Entonces las dos historias tienen identificadores distintos

Escenario: Rechazar una historia sin criterios de aceptacion
  Dado un proyecto registrado
  Cuando intento crear una historia sin ningun criterio de aceptacion
  Entonces la operacion se rechaza
  Y se informa que hace falta al menos un criterio

Escenario: Rechazar una prioridad invalida
  Dado un proyecto registrado
  Cuando intento crear una historia con prioridad "Urgente"
  Entonces la operacion se rechaza
  Y se informan las prioridades validas

Escenario: Modificar una historia pendiente
  Dado una historia en estado "Pendiente" con prioridad "Baja"
  Cuando cambio su prioridad a "Alta"
  Entonces la historia queda con prioridad "Alta"

Escenario: No modificar una historia terminada
  Dado una historia en estado "Terminada"
  Cuando intento cambiar su titulo
  Entonces la operacion se rechaza
  Y se informa que una historia terminada no se modifica

Escenario: Backlog de un proyecto sin historias
  Dado un proyecto registrado sin historias
  Cuando pido el backlog
  Entonces recibo una lista vacia

Esquema del escenario: Rechazar una historia sin los datos obligatorios
  Dado un proyecto registrado
  Cuando intento crear una historia sin <dato>
  Entonces la operacion se rechaza
  Y se informa que falta <dato>
  Y el backlog sigue vacio

  Ejemplos:
    | dato                      |
    | "titulo"                  |
    | "descripcion"             |
    | "prioridad"               |
    | "criterios de aceptacion" |

Escenario: Un texto con solo espacios cuenta como vacio
  Dado un proyecto registrado
  Cuando intento crear una historia con titulo "   "
  Entonces la operacion se rechaza
  Y se informa que el titulo es obligatorio

Escenario: Descartar los criterios de aceptacion vacios
  Dado un proyecto registrado
  Cuando creo una historia con los criterios "El sistema valida el nombre", "" y "   "
  Entonces la historia queda registrada con 1 criterio de aceptacion

Escenario: Rechazar si todos los criterios quedan vacios
  Dado un proyecto registrado
  Cuando intento crear una historia con los criterios "" y "   "
  Entonces la operacion se rechaza
  Y se informa que hace falta al menos un criterio

Escenario: Guardar el titulo sin espacios en los extremos
  Dado un proyecto registrado
  Cuando creo una historia con titulo "  Registrar esfuerzo  "
  Entonces la historia queda registrada con titulo "Registrar esfuerzo"

Escenario: Una modificacion rechazada no cambia nada
  Dado una historia en estado "Pendiente" con titulo "Registrar esfuerzo"
  Cuando intento modificar su titulo a "   "
  Entonces la operacion se rechaza
  Y la historia sigue teniendo titulo "Registrar esfuerzo"

Escenario: Modificar una historia que no existe
  Dado un proyecto registrado con una historia
  Cuando intento modificar la historia 99
  Entonces se informa que no se encontro la historia
  Y el backlog sigue teniendo 1 historia

Escenario: Reemplazar los criterios de aceptacion
  Dado una historia en estado "Pendiente" con 2 criterios de aceptacion
  Cuando la modifico enviando una lista con 3 criterios
  Entonces la historia queda con esos 3 criterios

Escenario: Los identificadores son unicos dentro de cada proyecto
  Dado un proyecto A y un proyecto B, los dos sin historias
  Cuando creo una historia en cada uno
  Entonces la historia del proyecto A tiene identificador 1
  Y la historia del proyecto B tiene identificador 1

Escenario: Pedir el backlog de un proyecto con historias
  Dado un proyecto con 3 historias cargadas
  Cuando pido el backlog
  Entonces recibo las 3 historias
  Y cada una con su identificador, titulo, descripcion, prioridad, estado, story points y
  criterios de aceptacion

Esquema del escenario: Operar sobre un proyecto que no existe
  Dado que no existe ningun proyecto con identificador 99
  Cuando intento <operacion> en el proyecto 99
  Entonces se informa que no se encontro el proyecto

  Ejemplos:
    | operacion            |
    | "crear una historia" |
    | "modificar una historia" |
    | "pedir el backlog"   |
```

**Para decidir en la puerta:**

1. **RN-9** — los criterios de aceptación se reemplazan enviando la lista completa, no se
   agregan ni se quitan de a uno. Es la opción más simple y evita el problema de distinguir
   "no mandé el dato" de "mandé el dato vacío", que es el mismo tema que quedó abierto en la
   spec de la #3.
2. **RN-16** — los identificadores son únicos *por proyecto*, no globales. Así el primer
   ítem de cada proyecto es el 1, que es lo que espera el usuario.
3. Los criterios de aceptación vacíos **se descartan en silencio** en vez de rechazar toda la
   operación (RN-6), salvo que no quede ninguno.

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: Maximiliano Eula — Fecha: 2026-10-02 — Comentarios: aprobada en el PR #28, con
> los ajustes de la revisión aplicados y las tres decisiones que planteaba el PR.

## Fase 2 — Diseño

_Se escribe después de que la Puerta 1 esté aprobada._

## Fase 3 — Tareas

_Se escribe después de que la Puerta 2 esté aprobada._

## Fase 4 — Implementación

_Se completa a medida que se avanza._

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #6 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | pendiente |
| Código Go | pendiente |
