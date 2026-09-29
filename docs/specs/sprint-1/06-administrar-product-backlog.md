# Spec: Administrar el Product Backlog

- **Historia:** #6
- **Sprint:** 1
- **Responsable:** Emiliano Sorato
- **Estado:** Requisitos — **borrador, pendiente de Puerta 1**
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de
  la consigna y de los criterios de aceptación de la issue. El responsable lo
  revisa, lo corrige si hace falta y lo presenta a la Puerta 1
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
| Story points | no | los carga la historia #12 |

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
5. **[RN-5]** Si se intenta crear o modificar una historia sin título, sin descripción o sin
   al menos un criterio de aceptación, el sistema debe rechazar la operación e informar qué
   falta.
6. **[RN-6]** Una historia recién creada debe quedar **sin estimar**: los story points son
   opcionales en esta historia.
7. **[RN-7]** Mientras una historia esté en estado Pendiente o En curso, el sistema debe
   permitir modificar su título, descripción, prioridad y criterios de aceptación.
8. **[RN-8]** Si se intenta modificar una historia que ya está Terminada, el sistema debe
   rechazar la operación e informar que una historia terminada no se modifica.
9. **[RN-9]** Si se intenta agregar una historia a un proyecto que no existe, el sistema debe
   informar que no se encontró el proyecto.
10. **[RN-10]** El sistema debe permitir que dos historias del mismo proyecto tengan el mismo
    título: lo que las distingue es el identificador.
11. **[RN-11]** Mientras un proyecto tenga historias cargadas, el sistema debe poder devolver
    el backlog completo de ese proyecto.

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

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| Título vacío o solo espacios | rechaza e informa que el título es obligatorio |
| Descripción vacía | rechaza e informa que la descripción es obligatoria |
| Sin criterios de aceptación | rechaza e informa que hace falta al menos uno |
| Prioridad distinta de las tres permitidas | rechaza e informa cuáles son válidas |
| Modificar una historia Terminada | rechaza e informa que no se puede modificar |
| Proyecto inexistente | informa que no se encontró el proyecto |
| Historia inexistente al modificar | informa que no se encontró la historia |

### Criterios de aceptación

- [ ] Se puede crear una historia con título, descripción, prioridad y criterios de
      aceptación, y queda con identificador propio y estado Pendiente.
- [ ] El identificador no se repite dentro del mismo proyecto.
- [ ] No se puede crear una historia sin título, sin descripción o sin criterios de aceptación.
- [ ] No se puede crear una historia con una prioridad que no sea Alta, Media o Baja.
- [ ] Una historia recién creada queda sin estimar.
- [ ] Se puede modificar una historia mientras no esté Terminada.
- [ ] No se puede modificar una historia Terminada.
- [ ] Se puede pedir el backlog completo de un proyecto.
- [ ] El backlog de un proyecto sin historias se devuelve vacío.

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
| Historia de usuario | issue #6 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | pendiente |
| Código Go | pendiente |
