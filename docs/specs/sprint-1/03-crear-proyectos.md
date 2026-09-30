# Spec: Crear y modificar proyectos

- **Historia:** #3
- **Sprint:** 1
- **Responsable:** Agustín Salinas
- **Estado:** Requisitos — **corregido, pendiente de Puerta 1**
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de
  la consigna y de los criterios de aceptación de la issue. El responsable lo
  revisa, lo corrige si hace falta y lo presenta a la Puerta 1. Corregido por el
  responsable y presentado a la Puerta 1 por pull request
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
   validaciones que al crearlo. Si alguna falla, el proyecto debe quedar como estaba antes.
6. **[RN-6]** Si se intenta modificar un proyecto que no existe, el sistema debe informar que
   no se encontró el proyecto y no modificar ni crear nada.
7. **[RN-7]** El sistema debe permitir que dos proyectos distintos tengan el mismo nombre:
   lo que los distingue es el identificador.
8. **[RN-8]** Cuando se modifica un proyecto, el sistema debe cambiar solo los datos que se
   indicaron en esa modificación y dejar el resto como estaba.
9. **[RN-9]** Cuando se modifica un proyecto, el sistema debe permitir dejar sin descripción
   o sin fecha de finalización un proyecto que las tenía, porque son datos opcionales. El
   nombre y la fecha de inicio no se pueden quitar (RN-2 y RN-3).
10. **[RN-10]** Cuando se registra el nombre de un proyecto, el sistema debe guardarlo sin los
    espacios del principio y del final.

### Casos límite

- Proyecto sin descripción y sin fecha de finalización: es válido, son opcionales.
- Fecha de finalización **igual** a la de inicio: es válida, un proyecto puede durar un día.
- Nombre con espacios al principio y al final: se guarda sin esos espacios.
- Nombre de un solo carácter: es válido.
- Modificar un proyecto sin cambiar nada: es válido y no da error.
- Modificar solo la fecha de inicio de un proyecto que ya tiene fecha de finalización: si la
  nueva fecha de inicio queda después de la de finalización, se rechaza por fechas
  incoherentes, aunque la fecha de finalización no se haya tocado.

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
- [ ] No se puede crear ni modificar un proyecto sin fecha de inicio.
- [ ] No se puede crear ni modificar un proyecto con fecha de fin anterior a la de inicio.
- [ ] Se pueden modificar los datos de un proyecto ya creado y los cambios quedan guardados.
- [ ] Una modificación rechazada deja el proyecto como estaba.
- [ ] Se le puede quitar la descripción o la fecha de finalización a un proyecto.
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

Esquema del escenario: Rechazar un proyecto sin nombre
  Cuando intento crear un proyecto con nombre <nombre> e inicio 01/10/2026
  Entonces la operacion se rechaza
  Y se informa que el nombre es obligatorio

  Ejemplos:
    | nombre      |
    | ""          |
    | "   "       |

Escenario: Rechazar un proyecto sin fecha de inicio
  Cuando intento crear un proyecto con nombre "Sistema de metricas" y sin fecha de inicio
  Entonces la operacion se rechaza
  Y se informa que la fecha de inicio es obligatoria

Escenario: Guardar el nombre sin espacios en los extremos
  Cuando creo un proyecto con nombre "  Sistema de metricas  " e inicio 01/10/2026
  Entonces el proyecto queda registrado con nombre "Sistema de metricas"

Escenario: Aceptar un nombre de un solo caracter
  Cuando creo un proyecto con nombre "X" e inicio 01/10/2026
  Entonces el proyecto queda registrado con nombre "X"

Escenario: Dos proyectos con el mismo nombre
  Dado un proyecto registrado con nombre "Sistema de metricas"
  Cuando creo otro proyecto con nombre "Sistema de metricas" e inicio 01/10/2026
  Entonces el proyecto queda registrado
  Y su identificador es distinto del identificador del primero

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

Escenario: Modificar un proyecto sin cambiar nada
  Dado un proyecto registrado con nombre "Sistema de metricas" e inicio 01/10/2026
  Cuando lo modifico con los mismos datos que ya tiene
  Entonces la operacion no da error
  Y el proyecto conserva sus datos

Escenario: Rechazar una modificacion que deja el proyecto sin nombre
  Dado un proyecto registrado con nombre "Sistema de metricas"
  Cuando intento modificar su nombre a "   "
  Entonces la operacion se rechaza
  Y se informa que el nombre es obligatorio
  Y el proyecto sigue llamandose "Sistema de metricas"

Escenario: Rechazar una fecha de inicio posterior a la de finalizacion al modificar
  Dado un proyecto registrado con inicio 01/10/2026 y fin 30/11/2026
  Cuando intento modificar su fecha de inicio a 01/12/2026
  Entonces la operacion se rechaza
  Y se informa que las fechas son incoherentes
  Y el proyecto conserva inicio 01/10/2026 y fin 30/11/2026

Escenario: Quitar la fecha de finalizacion
  Dado un proyecto registrado con inicio 01/10/2026 y fin 30/11/2026
  Cuando le quito la fecha de finalizacion
  Entonces el proyecto queda sin fecha de finalizacion
  Y conserva su nombre y su fecha de inicio

Escenario: Modificar un proyecto que no existe
  Dado que no existe ningun proyecto con identificador 99
  Cuando intento modificar el proyecto 99
  Entonces se informa que no se encontro el proyecto
  Y no se crea ningun proyecto
```

**Para decidir en la puerta:** RN-9 permite quitarle a un proyecto la descripción o la fecha
de finalización. La issue solo pide que se puedan "modificar los datos"; se eligió permitirlo
porque son datos opcionales y un proyecto puede no tener fecha de fin estimada todavía.

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
