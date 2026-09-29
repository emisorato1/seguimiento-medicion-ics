# Spec: Registrar integrantes de un proyecto

- **Historia:** #4
- **Sprint:** 1
- **Responsable:** Juliana Bustos
- **Estado:** Requisitos — **borrador, pendiente de Puerta 1**
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de
  la consigna y de los criterios de aceptación de la issue. El responsable lo
  revisa, lo corrige si hace falta y lo presenta a la Puerta 1
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
- Que un integrante pertenezca a varios proyectos a la vez.
- El esfuerzo que carga cada integrante, que es la historia #14.

### Entradas

| Dato | Obligatorio | Descripción |
|---|---|---|
| Proyecto | sí | a qué proyecto se agrega |
| Nombre y apellido | sí | cómo se identifica a la persona |
| Rol | sí | Agile Enabler o Product Builder |

### Salidas

El integrante queda asociado al proyecto. Se puede pedir la lista de integrantes de un
proyecto.

### Reglas de negocio

1. **[RN-1]** Cuando se agrega un integrante a un proyecto, el sistema debe registrar su
   nombre y apellido y su rol.
2. **[RN-2]** El sistema debe aceptar únicamente los roles **Agile Enabler** y
   **Product Builder**.
3. **[RN-3]** Si se intenta agregar un integrante con un nombre que ya existe en ese
   proyecto, el sistema debe rechazar la operación e informar que ya está cargado.
4. **[RN-4]** El sistema debe permitir que la misma persona exista en proyectos distintos
   sin que eso se considere duplicado.
5. **[RN-5]** Mientras un proyecto tenga integrantes cargados, el sistema debe poder
   devolver la lista completa.
6. **[RN-6]** El sistema debe permitir quitar un integrante de un proyecto.
7. **[RN-7]** Si se intenta agregar o quitar un integrante de un proyecto que no existe, el
   sistema debe informar que no se encontró el proyecto.
8. **[RN-8]** Si se intenta quitar un integrante que no pertenece al proyecto, el sistema debe
   informar que no se encontró el integrante.
9. **[RN-9]** El sistema debe permitir que un proyecto tenga un solo Agile Enabler.

### Casos límite

- Proyecto sin ningún integrante: la lista se devuelve vacía, no es un error.
- Agregar el primer integrante.
- Quitar el último integrante: el proyecto queda sin integrantes y sigue siendo válido.
- Dos personas con el mismo nombre en proyectos distintos: válido.
- Nombre con espacios al principio y al final: se guarda sin esos espacios, y la comparación
  de duplicados se hace sobre el nombre ya limpio.

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| Nombre vacío o solo espacios | rechaza e informa que el nombre es obligatorio |
| Rol distinto de los dos permitidos | rechaza e informa cuáles son los roles válidos |
| Integrante ya cargado en ese proyecto | rechaza e informa que ya está cargado |
| Segundo Agile Enabler en el mismo proyecto | rechaza e informa que ya hay uno |
| Proyecto inexistente | informa que no se encontró el proyecto |
| Integrante inexistente al quitarlo | informa que no se encontró el integrante |

### Criterios de aceptación

- [ ] Se puede agregar un integrante indicando nombre, apellido y rol.
- [ ] Se puede listar los integrantes de un proyecto.
- [ ] No se puede cargar dos veces la misma persona en un mismo proyecto.
- [ ] La misma persona se puede cargar en dos proyectos distintos.
- [ ] No se puede cargar un segundo Agile Enabler en el mismo proyecto.
- [ ] Se puede quitar un integrante del proyecto.
- [ ] Listar los integrantes de un proyecto sin integrantes devuelve una lista vacía.

### Escenarios BDD

```gherkin
Escenario: Agregar el primer integrante
  Dado un proyecto registrado sin integrantes
  Cuando agrego a "Juliana Bustos" con rol "Product Builder"
  Entonces el proyecto tiene 1 integrante
  Y la lista incluye a "Juliana Bustos"

Escenario: Rechazar un integrante duplicado
  Dado un proyecto con "Juliana Bustos" ya cargada
  Cuando intento agregar a "Juliana Bustos" otra vez
  Entonces la operacion se rechaza
  Y se informa que ya esta cargada

Escenario: La misma persona en dos proyectos
  Dado un proyecto A con "Juliana Bustos" cargada
  Y un proyecto B sin integrantes
  Cuando agrego a "Juliana Bustos" al proyecto B
  Entonces la operacion se acepta

Escenario: Rechazar un segundo Agile Enabler
  Dado un proyecto con "Emiliano Sorato" como "Agile Enabler"
  Cuando intento agregar a "Agustin Salinas" como "Agile Enabler"
  Entonces la operacion se rechaza
  Y se informa que el proyecto ya tiene un Agile Enabler

Escenario: Rechazar un rol invalido
  Dado un proyecto registrado
  Cuando intento agregar a "Maximiliano Eula" con rol "Product Owner"
  Entonces la operacion se rechaza
  Y se informan los roles validos

Escenario: Quitar un integrante
  Dado un proyecto con "Juliana Bustos" cargada
  Cuando quito a "Juliana Bustos"
  Entonces el proyecto queda sin integrantes

Escenario: Listar un proyecto sin integrantes
  Dado un proyecto registrado sin integrantes
  Cuando pido la lista de integrantes
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
| Historia de usuario | issue #4 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | pendiente |
| Código Go | pendiente |
