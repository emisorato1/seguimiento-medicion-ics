# Spec: Crear y modificar proyectos

- **Historia:** #3
- **Sprint:** 1
- **Responsable:** Agustín Salinas
- **Estado:** Requisitos — **borrador, pendiente de Puerta 1**
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de
  la consigna y de los criterios de aceptación de la issue. El responsable lo
  revisa, lo corrige si hace falta y lo presenta a la Puerta 1
- **Depende de:** nada. Es la base del Sprint 1.
- **La necesitan:** #4 (integrantes), #6 (Product Backlog), #8 (sprints)

## Fase 1 — Requisitos

### Objetivo

Que un integrante del equipo pueda registrar en el sistema el proyecto de software sobre el
que va a trabajar, y corregir sus datos cuando cambien. Es la entidad raíz: todo lo demás
—integrantes, historias, sprints, métricas— cuelga de un proyecto.

### Qué queda afuera

- Borrar proyectos.
- Archivar o cerrar un proyecto.
- Permisos: en esta historia cualquiera puede crear y modificar cualquier proyecto.
- Los integrantes del proyecto, que son la historia #4.

### Entradas

| Dato | Obligatorio | Descripción |
|---|---|---|
| Nombre | sí | cómo se llama el proyecto |
| Descripción | no | texto libre que explica de qué se trata |
| Fecha de inicio | sí | cuándo arranca |
| Fecha de finalización | no | cuándo se estima terminarlo |

### Salidas

Un proyecto registrado, con un identificador propio que el sistema asigna y que no cambia.
Al consultarlo se devuelven sus datos y ese identificador.

### Reglas de negocio

1. **[RN-1]** El sistema debe asignar a cada proyecto un identificador único que no se
   repite y que no cambia durante toda la vida del proyecto.
2. **[RN-2]** Si se intenta crear o modificar un proyecto sin nombre, o con un nombre
   compuesto solo por espacios, el sistema debe rechazar la operación e informar que el
   nombre es obligatorio.
3. **[RN-3]** Si se intenta crear o modificar un proyecto sin fecha de inicio, el sistema debe
   rechazar la operación e informar que la fecha de inicio es obligatoria.
4. **[RN-4]** Si se indica una fecha de finalización anterior a la fecha de inicio, el sistema
   debe rechazar la operación e informar que las fechas son incoherentes.
5. **[RN-5]** Cuando se modifica un proyecto existente, el sistema debe aplicar las mismas
   validaciones que al crearlo.
6. **[RN-6]** Si se intenta modificar un proyecto que no existe, el sistema debe informar que
   no se encontró el proyecto y no modificar nada.
7. **[RN-7]** El sistema debe permitir que dos proyectos distintos tengan el mismo nombre:
   lo que los distingue es el identificador.
8. **[RN-8]** Cuando se modifica un proyecto, el sistema debe conservar los datos que no se
   enviaron en la modificación.

### Casos límite

- Proyecto sin descripción y sin fecha de finalización: es válido, son opcionales.
- Fecha de finalización **igual** a la de inicio: es válida, un proyecto puede durar un día.
- Nombre con espacios al principio y al final: se guarda sin esos espacios.
- Nombre de un solo carácter: es válido.
- Modificar un proyecto sin cambiar nada: es válido y no da error.

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| Nombre vacío o solo espacios | rechaza e informa que el nombre es obligatorio |
| Falta la fecha de inicio | rechaza e informa que la fecha de inicio es obligatoria |
| Fecha de fin anterior a la de inicio | rechaza e informa que las fechas son incoherentes |
| Identificador inexistente al modificar | informa que no se encontró el proyecto |

### Criterios de aceptación

- [ ] Se puede crear un proyecto con nombre, descripción, fecha de inicio y fecha de fin, y
      queda registrado con un identificador propio.
- [ ] Se puede crear un proyecto con solo nombre y fecha de inicio.
- [ ] No se puede crear ni modificar un proyecto sin nombre.
- [ ] No se puede crear ni modificar un proyecto con fecha de fin anterior a la de inicio.
- [ ] Se pueden modificar los datos de un proyecto ya creado y los cambios quedan guardados.
- [ ] Dos proyectos creados con el mismo nombre reciben identificadores distintos.
- [ ] Modificar un proyecto inexistente informa el error y no crea nada.

### Escenarios BDD

```gherkin
Escenario: Crear un proyecto con todos los datos
  Dado que no existe ningun proyecto registrado
  Cuando creo un proyecto con nombre "Sistema de metricas", inicio 01/10/2026 y fin 30/11/2026
  Entonces el proyecto queda registrado
  Y recibo su identificador

Escenario: Crear un proyecto solo con los datos obligatorios
  Dado que no existe ningun proyecto registrado
  Cuando creo un proyecto con nombre "Sistema de metricas" e inicio 01/10/2026
  Entonces el proyecto queda registrado sin descripcion y sin fecha de finalizacion

Escenario: Rechazar un proyecto sin nombre
  Cuando intento crear un proyecto con nombre vacio e inicio 01/10/2026
  Entonces la operacion se rechaza
  Y se informa que el nombre es obligatorio

Escenario: Rechazar fechas incoherentes
  Cuando intento crear un proyecto con inicio 01/10/2026 y fin 30/09/2026
  Entonces la operacion se rechaza
  Y se informa que las fechas son incoherentes

Escenario: Aceptar un proyecto que empieza y termina el mismo dia
  Cuando creo un proyecto con inicio 01/10/2026 y fin 01/10/2026
  Entonces el proyecto queda registrado

Escenario: Modificar el nombre de un proyecto existente
  Dado un proyecto registrado con nombre "Sistema de metricas"
  Cuando modifico su nombre a "Seguimiento y medicion"
  Entonces el proyecto pasa a llamarse "Seguimiento y medicion"
  Y conserva sus fechas

Escenario: Modificar un proyecto que no existe
  Dado que no existe ningun proyecto con identificador 99
  Cuando intento modificar el proyecto 99
  Entonces se informa que no se encontro el proyecto
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
| Historia de usuario | issue #3 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | pendiente |
| Código Go | pendiente |
