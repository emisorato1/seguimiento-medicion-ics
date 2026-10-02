# Spec: Registrar integrantes de un proyecto

- **Historia:** #4
- **Sprint:** 1
- **Responsable:** Juliana Bustos
- **Estado:** Puerta 1 aprobada — **sigue la Fase 2 (Diseño)**
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de
  la consigna y de los criterios de aceptación de la issue. El responsable lo
  revisa, lo corrige si hace falta y lo presenta a la Puerta 1. Corregido por la
  responsable y presentado a la Puerta 1 por pull request
- **Depende de:** #3 (tiene que existir un proyecto al que agregarle integrantes)
- **La necesitan:** #14 (registro de esfuerzo por integrante), #13 (Planning Poker)

## Fase 1 — Requisitos

### Objetivo

Que se pueda cargar quiénes forman parte de un proyecto, con el rol de cada uno, para
después poder asignarles historias, registrar su esfuerzo y saber quiénes votan en una
estimación.

### Qué queda afuera

- Usuarios y contraseñas: un integrante es un dato del proyecto, no alguien que entra al
  sistema.
- Modificar el nombre y apellido o el rol de un integrante ya cargado: para corregirlo se lo
  quita y se lo vuelve a agregar.
- Asignarle historias a un integrante.
- El esfuerzo que carga cada integrante, que es la historia #14.

### Entradas

| Dato | Obligatorio | Descripción |
|---|---|---|
| Proyecto | sí | a qué proyecto se agrega o de cuál se quita |
| Nombre y apellido | sí | un solo dato de texto que identifica a la persona dentro del proyecto |
| Rol | sí, al agregar | Agile Enabler o Product Builder |

### Salidas

El integrante queda asociado al proyecto, con su nombre y apellido y su rol. Se puede pedir
la lista de integrantes de un proyecto, con el rol de cada uno.

### Reglas de negocio

1. **[RN-1]** Cuando se agrega un integrante a un proyecto, el sistema debe registrar su
   nombre y apellido y su rol.
2. **[RN-2]** Si se intenta agregar un integrante sin nombre y apellido, o con uno compuesto
   solo por espacios, el sistema debe rechazar la operación e informar que el nombre y
   apellido es obligatorio.
3. **[RN-3]** Cuando se agrega un integrante, el sistema debe guardar su nombre y apellido sin
   los espacios del principio y del final.
4. **[RN-4]** Si se intenta agregar un integrante con un rol distinto de **Agile Enabler** y
   **Product Builder**, el sistema debe rechazar la operación e informar cuáles son los roles
   válidos.
5. **[RN-5]** Si se intenta agregar un integrante cuyo nombre y apellido ya está cargado en ese
   proyecto, el sistema debe rechazar la operación e informar que ya está cargado. La
   comparación no distingue mayúsculas de minúsculas y se hace sobre el nombre y apellido
   sin los espacios del principio y del final.
6. **[RN-6]** El sistema debe permitir que el mismo nombre y apellido esté cargado en
   proyectos distintos, sin considerarlo duplicado.
7. **[RN-7]** Si se intenta agregar un Agile Enabler a un proyecto que ya tiene uno, el
   sistema debe rechazar la operación e informar que el proyecto ya tiene un Agile Enabler.
8. **[RN-8]** Cuando se pide la lista de integrantes de un proyecto, el sistema debe devolver
   todos sus integrantes con su rol. Si el proyecto no tiene integrantes, debe devolver una
   lista vacía, y eso no es un error.
9. **[RN-9]** Cuando se quita un integrante de un proyecto, el sistema debe dejar de incluirlo
   en la lista de integrantes. El integrante se indica por su nombre y apellido, con la misma
   comparación que la RN-5.
10. **[RN-10]** Si se intenta agregar, quitar o listar integrantes de un proyecto que no
    existe, el sistema debe informar que no se encontró el proyecto.
11. **[RN-11]** Si se intenta quitar un integrante que no pertenece al proyecto, el sistema
    debe informar que no se encontró el integrante.
12. **[RN-12]** Si una operación se rechaza, el sistema debe dejar los integrantes del
    proyecto como estaban.

### Casos límite

- Proyecto sin ningún integrante: la lista se devuelve vacía, no es un error.
- Agregar el primer integrante.
- Quitar el último integrante: el proyecto queda sin integrantes y sigue siendo válido.
- El mismo nombre y apellido en proyectos distintos: válido.
- Nombre con espacios al principio y al final: se guarda sin esos espacios, y la comparación
  de duplicados se hace sobre el nombre ya limpio.
- El mismo nombre escrito con otras mayúsculas ("juliana bustos" y "Juliana Bustos") en el
  mismo proyecto: es duplicado.
- Quitar al Agile Enabler: el proyecto queda sin Agile Enabler y se puede agregar otro.
- Un proyecto sin Agile Enabler, o con varios Product Builders: válido.

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| Nombre y apellido vacío o solo espacios | rechaza e informa que el nombre y apellido es obligatorio |
| Rol distinto de los dos permitidos | rechaza e informa cuáles son los roles válidos |
| Integrante ya cargado en ese proyecto | rechaza e informa que ya está cargado |
| Segundo Agile Enabler en el mismo proyecto | rechaza e informa que el proyecto ya tiene uno |
| Proyecto inexistente | informa que no se encontró el proyecto |
| Integrante inexistente al quitarlo | informa que no se encontró el integrante |

### Criterios de aceptación

- [ ] Se puede agregar un integrante indicando su nombre y apellido y su rol.
- [ ] No se puede agregar un integrante sin nombre y apellido.
- [ ] No se puede agregar un integrante con un rol distinto de Agile Enabler o Product Builder.
- [ ] Se pueden listar los integrantes de un proyecto, con su rol.
- [ ] Listar los integrantes de un proyecto sin integrantes devuelve una lista vacía.
- [ ] No se puede cargar dos veces el mismo integrante en un proyecto, aunque cambien las
      mayúsculas o los espacios de los extremos.
- [ ] El mismo integrante se puede cargar en dos proyectos distintos.
- [ ] No se puede cargar un segundo Agile Enabler en el mismo proyecto.
- [ ] Se puede quitar un integrante del proyecto.
- [ ] Quitar un integrante que no está en el proyecto informa el error.
- [ ] Operar sobre un proyecto inexistente informa que no se encontró.
- [ ] Una operación rechazada deja los integrantes del proyecto como estaban.

### Escenarios BDD

```gherkin
Escenario: Agregar el primer integrante
  Dado un proyecto registrado sin integrantes
  Cuando agrego a "Juliana Bustos" con rol "Product Builder"
  Entonces el proyecto tiene 1 integrante
  Y la lista incluye a "Juliana Bustos" con rol "Product Builder"

Esquema del escenario: Rechazar un integrante sin nombre y apellido
  Dado un proyecto registrado sin integrantes
  Cuando intento agregar a <nombre> con rol "Product Builder"
  Entonces la operacion se rechaza
  Y se informa que el nombre y apellido es obligatorio
  Y el proyecto sigue sin integrantes

  Ejemplos:
    | nombre |
    | ""     |
    | "   "  |

Escenario: Guardar el nombre sin espacios en los extremos
  Dado un proyecto registrado sin integrantes
  Cuando agrego a "  Juliana Bustos  " con rol "Product Builder"
  Entonces la lista incluye a "Juliana Bustos"

Escenario: Rechazar un rol invalido
  Dado un proyecto registrado sin integrantes
  Cuando intento agregar a "Maximiliano Eula" con rol "Product Owner"
  Entonces la operacion se rechaza
  Y se informa que los roles validos son "Agile Enabler" y "Product Builder"
  Y el proyecto sigue sin integrantes

Esquema del escenario: Rechazar un integrante duplicado
  Dado un proyecto con "Juliana Bustos" cargada como "Product Builder"
  Cuando intento agregar a <nombre> con rol "Product Builder"
  Entonces la operacion se rechaza
  Y se informa que ya esta cargado
  Y el proyecto sigue teniendo 1 integrante

  Ejemplos:
    | nombre               |
    | "Juliana Bustos"     |
    | "juliana bustos"     |
    | "  JULIANA BUSTOS  " |

Escenario: La misma persona en dos proyectos
  Dado un proyecto A con "Juliana Bustos" cargada
  Y un proyecto B sin integrantes
  Cuando agrego a "Juliana Bustos" al proyecto B con rol "Product Builder"
  Entonces la operacion se acepta
  Y los dos proyectos incluyen a "Juliana Bustos"

Escenario: Rechazar un segundo Agile Enabler
  Dado un proyecto con "Emiliano Sorato" como "Agile Enabler"
  Cuando intento agregar a "Agustin Salinas" como "Agile Enabler"
  Entonces la operacion se rechaza
  Y se informa que el proyecto ya tiene un Agile Enabler
  Y el proyecto sigue teniendo 1 integrante

Escenario: Agregar un Agile Enabler despues de quitar al anterior
  Dado un proyecto con "Emiliano Sorato" como "Agile Enabler"
  Cuando quito a "Emiliano Sorato"
  Y agrego a "Agustin Salinas" como "Agile Enabler"
  Entonces la lista incluye a "Agustin Salinas" con rol "Agile Enabler"

Escenario: Listar los integrantes de un proyecto
  Dado un proyecto con "Emiliano Sorato" como "Agile Enabler"
  Y "Juliana Bustos" como "Product Builder"
  Cuando pido la lista de integrantes
  Entonces recibo 2 integrantes, cada uno con su rol

Escenario: Listar un proyecto sin integrantes
  Dado un proyecto registrado sin integrantes
  Cuando pido la lista de integrantes
  Entonces recibo una lista vacia

Escenario: Quitar el ultimo integrante
  Dado un proyecto con "Juliana Bustos" cargada
  Cuando quito a "juliana bustos"
  Entonces el proyecto queda sin integrantes

Escenario: Quitar un integrante que no pertenece al proyecto
  Dado un proyecto con "Juliana Bustos" cargada
  Cuando intento quitar a "Maximiliano Eula"
  Entonces se informa que no se encontro el integrante
  Y el proyecto sigue teniendo 1 integrante

Escenario: Operar sobre un proyecto que no existe
  Dado que no existe ningun proyecto con identificador 99
  Cuando intento agregar a "Juliana Bustos" al proyecto 99 con rol "Product Builder"
  Entonces se informa que no se encontro el proyecto
```

**Decisiones tomadas por la responsable, para validar en la puerta:**

1. Los duplicados se comparan **sin distinguir mayúsculas de minúsculas** (RN-5).
2. Los únicos roles son **Agile Enabler** y **Product Builder** (RN-4).
3. **Nombre y apellido es un solo dato** de texto, no dos campos (Entradas).
4. Para quitar un integrante se lo indica por su nombre y apellido, que es único dentro del
   proyecto (RN-9).
5. Modificar un integrante ya cargado queda afuera de esta historia (Qué queda afuera).

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: Emiliano Sorato — Fecha: 2026-10-02 — Comentarios: aprobada en el PR #27, con las
> cinco decisiones que planteaba el PR. Observación menor registrada: "nombre y apellido es
> un solo dato de texto" es una decisión de modelado y su lugar natural es la Fase 2.

## Fase 2 — Diseño

_Se escribe después de que la Puerta 1 esté aprobada._

## Fase 3 — Tareas

_Se escribe después de que la Puerta 2 esté aprobada._

## Fase 4 — Implementación

_Se completa a medida que se avanza._

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #4 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | pendiente |
| Código Go | pendiente |
