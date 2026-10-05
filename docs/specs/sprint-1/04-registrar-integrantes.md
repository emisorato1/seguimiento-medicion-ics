# Spec: Registrar integrantes de un proyecto

- **Historia:** #4
- **Sprint:** 1
- **Responsable:** Juliana Bustos
- **Estado:** Diseño — **pendiente de Puerta 2**
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

Las decisiones transversales del proyecto —un solo paquete `internal/domain` con un archivo
por entidad, datos en memoria, errores como valores del dominio, mensajes en español,
validaciones en el dominio— están en [AGENTS.md](../../../AGENTS.md) y no se repiten acá.
Esta fase registra solo lo propio de esta historia.

### Enfoque

Los integrantes **no son una colección aparte: viven dentro del proyecto**, igual que el
backlog de la #6. Un integrante no existe sin proyecto, y guardarlo adentro hace que la
RN-6 salga sola: el duplicado se busca solo entre los integrantes de ese proyecto, así que
el mismo nombre en otro proyecto nunca choca.

Agregar y quitar son operaciones del proyecto, y siguen el mismo orden que la #3 y la #6:

1. **Normalizar** el nombre y apellido: sacar los espacios de los extremos (RN-3).
2. **Validar** sobre lo ya normalizado, en este orden: nombre y apellido obligatorio (RN-2),
   rol válido (RN-4), duplicado (RN-5) y Agile Enabler único (RN-7). Para quitar, que el
   integrante exista (RN-11).
3. **Aplicar** el cambio recién si todo validó.

Como nada se toca antes de validar todo, una operación rechazada deja los integrantes como
estaban sin tener que deshacer nada (RN-12).

Encontrar el proyecto es del registro de la #3. Si el proyecto no existe, esta historia usa
el mismo error de "no se encontró el proyecto" que declara la #3 (RN-10), no uno propio.

### Archivos afectados

| Archivo | Se crea o se modifica | Para qué |
|---|---|---|
| `internal/domain/member.go` | se crea | el integrante, su rol, la comparación de nombres y las operaciones de agregar, quitar y listar integrantes del proyecto |
| `internal/domain/member_test.go` | se crea | los tests de los escenarios BDD de la Fase 1 |
| `internal/domain/project.go` | se modifica | el proyecto pasa a contener su lista de integrantes; nada más |
| `internal/domain/errors.go` | se modifica | los errores de esta historia |

### Decisiones

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| Los integrantes viven **dentro del proyecto** | una colección de integrantes aparte, indexada por proyecto | un integrante no existe sin proyecto, y adentro la RN-6 (mismo nombre en otro proyecto) no necesita ningún código especial. Es el mismo criterio que la #6 usa para el backlog |
| Las operaciones sobre integrantes van en `member.go`, aunque sean del proyecto; `project.go` solo gana el dato | escribirlas en `project.go` | la #6 también modifica `project.go`. Dejarle a ese archivo un solo cambio chico reduce los conflictos entre las dos historias, con el mismo criterio con que la #3 separó el registro |
| El nombre y apellido es **un solo dato de texto** | dos datos, nombre por un lado y apellido por el otro | lo decidió la Puerta 1. Con dos datos aparecen preguntas que la Fase 1 no tiene: si se puede cargar solo el nombre, o cómo se compara un duplicado con un apellido compuesto |
| El integrante se identifica por su **nombre y apellido**, sin un número propio | asignarle un identificador entero, como a las historias | la Fase 1 lo identifica por su nombre (RN-9), que es único dentro del proyecto, y modificar un integrante quedó afuera, así que el nombre no cambia. Si la #13 o la #14 necesitan referirse a un integrante de otra forma, lo agregan ellas |
| Se guarda el nombre **como se escribió** (sin los espacios de los extremos); las mayúsculas solo se ignoran **al comparar** | guardarlo todo en minúsculas | la lista tiene que mostrar "Juliana Bustos", no "juliana bustos". Normalizar para comparar y no para guardar cumple la RN-5 sin perder cómo lo escribió el usuario |
| La comparación ignora mayúsculas también en **letras con tilde y la ñ**, pero **no ignora las tildes** ni junta los espacios del medio | ignorar también las tildes ("Agustin" igual a "Agustín") o juntar espacios repetidos | la Puerta 1 decidió ignorar solo las mayúsculas. "ÍÑIGO" e "íñigo" son el mismo nombre; "Agustin" y "Agustín" pueden ser dos personas, y decidir que no lo son sería agregar una regla que nadie aprobó |
| El rol es un **tipo propio con dos valores**, y el dominio rechaza cualquier otro valor aunque la consola solo ofrezca esos dos | un texto libre, o confiar en que la consola manda siempre un rol válido | AGENTS.md pone las validaciones en el dominio, y la RN-4 tiene escenario propio. Es el mismo trato que la #6 le da al estado de una historia |
| Si una operación tiene varios errores, se informa **el primero** del orden del Enfoque | informar todos los errores juntos | cada escenario de la Fase 1 espera un solo mensaje. Fijar el orden hace que el resultado sea predecible: un nombre vacío con un rol inválido informa el nombre |
| La lista de integrantes sale **en el orden en que se agregaron**, y es una copia que no permite cambiar los integrantes del proyecto | orden alfabético, o entregar la lista guardada | el orden de alta es predecible para los tests, y ordenar para mostrar es de la consola. Entregar la lista guardada permitiría agregar o quitar integrantes sin pasar por las validaciones |

### Dependencias y restricciones

- **Depende de la #3.** Necesita `project.go`, el registro de proyectos y `errors.go`, que
  crea la #3. No se puede empezar a implementar hasta que la #3 esté en `dev`.
- **Depende de lo que se decida en la Puerta 2 de la #3** (PR #41): si al buscar un
  proyecto se recibe **el proyecto guardado** o **una copia**. Este diseño supone la opción 3
  que se está discutiendo ahí: se recibe el proyecto guardado, con sus datos ocultos fuera del
  paquete, así que agregar o quitar un integrante se hace directamente sobre el proyecto. Si
  la #3 elige la copia, esta historia tiene que guardar cada cambio a través del registro, y
  esta fase se ajusta antes de la Puerta 2.
- **Choca con la #6 en `project.go`.** Las dos historias le agregan un dato al proyecto. El
  cambio de esta historia en ese archivo se limita a la lista de integrantes; quien mergee
  segundo trae `dev` y resuelve el conflicto.
- **No se toca `cmd/`.** Pedir los datos por teclado, traducir lo que escribe el usuario a uno
  de los dos roles y mostrar la lista es de la consola (#24).
- **La #13 y la #14 se van a apoyar en esta historia** para saber quiénes votan y quién carga
  esfuerzo. Esta historia no les deja nada preparado: ver la decisión sobre el identificador.

> **Puerta 2** — ¿el diseño es viable y coherente con el resto del proyecto?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios:

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
