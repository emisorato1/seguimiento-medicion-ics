# Spec: Priorizar el Product Backlog

- **Historia:** #7
- **Sprint:** 2
- **Responsable:** a definir en la planning
- **Estado:** Requisitos — **borrador, pendiente de Puerta 1**
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de la consigna y
  de los criterios de aceptación de la issue. El responsable lo revisa, lo corrige si hace
  falta y lo presenta a la Puerta 1
- **Depende de:** #6 (el backlog y la prioridad de cada historia viven ahí), y por lo tanto
  también de #3
- **La necesitan:** #9 (en la planning se eligen primero las historias más prioritarias) y
  la interfaz del Sprint 2, que es donde se muestra el backlog ordenado

## Fase 1 — Requisitos

### Objetivo

Que el equipo pueda ver el Product Backlog de un proyecto ordenado por prioridad, de lo más
importante a lo menos importante, para saber de un vistazo qué conviene comprometer en el
próximo sprint.

### Qué queda afuera

- **Definir cuáles son las prioridades y validarlas.** Alta, Media y Baja ya están definidas
  en la historia #6 (su RN-4), y acá no se agrega ni se quita ninguna.
- **Asignar y cambiar la prioridad de una historia.** También es de la #6: una historia nace
  con prioridad y se le cambia modificándola. Esta historia solo se ocupa del efecto que ese
  cambio tiene en el orden.
- **Ordenar el backlog a mano**, arrastrando historias o dándoles una posición propia
  independiente de la prioridad.
- **Ordenar por otra cosa** que no sea la prioridad: por story points, por estado o por
  fecha.
- **Filtrar el backlog**, por prioridad, por estado o por lo que sea: el listado devuelve
  todas las historias del proyecto.
- Elegir qué historias entran a un sprint, que es la #9.
- Cómo se muestra el listado en pantalla: eso es de la interfaz.

### Entradas

| Dato | Obligatorio | Descripción |
|---|---|---|
| Proyecto | sí | de qué proyecto se quiere el backlog ordenado |

No hay ninguna otra entrada. El criterio de orden **no se elige**: siempre es el mismo
(RN-2 y RN-3). La prioridad de cada historia ya se cargó en la #6.

### Salidas

El listado de todas las historias del proyecto, ordenado de la prioridad más importante a la
menos importante, y cada historia con los mismos datos que tiene en el backlog. Pedir el
listado es una consulta: no cambia nada.

### Reglas de negocio

1. **[RN-1]** El sistema debe considerar **Alta** la prioridad más importante, **Media** la
   intermedia y **Baja** la menos importante. Son las tres prioridades que ya define la
   historia #6, y esta historia no agrega ni quita ninguna.
2. **[RN-2]** Cuando se pide el Product Backlog ordenado de un proyecto, el sistema debe
   devolver sus historias de la más importante a la menos importante: primero las de
   prioridad Alta, después las de prioridad Media y por último las de prioridad Baja.
3. **[RN-3]** Cuando dos o más historias tienen la misma prioridad, el sistema debe
   ordenarlas entre sí por su identificador, de menor a mayor, de manera que un mismo backlog
   se devuelva siempre en el mismo orden.
4. **[RN-4]** El sistema debe incluir en el listado ordenado **todas** las historias del
   proyecto, cualquiera sea su estado y estén estimadas o no, sin omitir ni repetir ninguna.
5. **[RN-5]** El sistema debe devolver cada historia del listado con los mismos datos que
   tiene en el backlog: identificador, título, descripción, prioridad, estado, story points y
   criterios de aceptación.
6. **[RN-6]** Cuando cambia la prioridad de una historia, el sistema debe ubicarla en el
   lugar que le corresponde a su nueva prioridad la próxima vez que se pida el listado
   ordenado.
7. **[RN-7]** El sistema debe devolver el listado ordenado sin modificar nada: ni los datos
   de ninguna historia, ni el orden en que el backlog devuelve sus historias cuando se lo
   pide sin ordenar.
8. **[RN-8]** Mientras un proyecto no tenga ninguna historia, el sistema debe devolver un
   listado ordenado vacío, y eso no es un error.
9. **[RN-9]** Si se pide el backlog ordenado de un proyecto que no existe, el sistema debe
   informar que no se encontró el proyecto, igual que cualquier otra operación sobre un
   proyecto inexistente.

### Casos límite

- **Proyecto sin historias:** el listado se devuelve vacío, no es un error.
- **Proyecto con una sola historia:** el listado la devuelve sola, sea cual sea su prioridad.
- **Todas las historias con la misma prioridad:** no hay nada que ordenar por prioridad, así
  que quedan ordenadas por identificador de menor a mayor.
- **No hay historias de alguna de las tres prioridades** (por ejemplo, ninguna Alta): el
  listado arranca directamente por el grupo siguiente, sin huecos ni marcas de que ese grupo
  está vacío.
- **Dos historias con el mismo título y la misma prioridad:** las distingue el identificador,
  que además decide cuál va primero.
- **Historias en distintos estados**, incluso Terminadas: aparecen todas, y el estado no
  cambia el orden.
- **Cambiar la prioridad de la historia que estaba primera a Baja:** pasa al final del
  listado.
- **Cambiar una prioridad por la misma que ya tenía:** el listado queda exactamente igual.
- **Pedir el listado dos veces seguidas sin cambiar nada:** devuelve exactamente lo mismo.

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| Se pide el backlog ordenado de un proyecto que no existe | informa que no se encontró el proyecto y no devuelve ningún listado |
| El proyecto existe pero no tiene historias | **no es un error**: devuelve un listado vacío |
| Se pide el listado ordenado dos veces seguidas | **no es un error**: devuelve lo mismo las dos veces |

Una historia con una prioridad que no sea una de las tres no se contempla, porque la #6 no
deja guardarla (su RN-4 y su RN-5). Queda anotado para la puerta, más abajo.

### Criterios de aceptación

- [ ] Pedir el backlog ordenado de un proyecto devuelve primero las historias de prioridad
      Alta, después las Media y por último las Baja.
- [ ] Dentro de una misma prioridad, las historias quedan ordenadas por identificador de
      menor a mayor.
- [ ] El listado ordenado devuelve la misma cantidad de historias que tiene el backlog, sin
      omitir ni repetir ninguna.
- [ ] Las historias aparecen en el listado cualquiera sea su estado y estén estimadas o no.
- [ ] Cada historia del listado se devuelve con su identificador, título, descripción,
      prioridad, estado, story points y criterios de aceptación.
- [ ] Al cambiar la prioridad de una historia, el listado siguiente la ubica en el lugar que
      corresponde a su nueva prioridad.
- [ ] Cambiar una prioridad por la misma que ya tenía deja el listado igual.
- [ ] Pedir el listado ordenado no cambia ningún dato de las historias ni el orden en que el
      backlog las devuelve sin ordenar.
- [ ] Pedir el listado ordenado dos veces seguidas, sin cambiar nada en el medio, devuelve
      exactamente el mismo resultado.
- [ ] El backlog ordenado de un proyecto sin historias se devuelve vacío, y eso no es un
      error.
- [ ] Un proyecto con una sola historia devuelve un listado con esa única historia.
- [ ] Un backlog donde todas las historias tienen la misma prioridad se devuelve ordenado por
      identificador.
- [ ] Pedir el backlog ordenado de un proyecto que no existe informa que no se encontró.

### Escenarios BDD

```gherkin
Escenario: Listar el backlog ordenado por prioridad
  Dado un proyecto con las historias 1 "Baja", 2 "Alta" y 3 "Media"
  Cuando pido el backlog ordenado
  Entonces recibo las historias en el orden 2, 3, 1

Escenario: Desempatar por identificador dentro de la misma prioridad
  Dado un proyecto con las historias 1 "Media", 2 "Alta" y 3 "Alta"
  Cuando pido el backlog ordenado
  Entonces recibo las historias en el orden 2, 3, 1

Esquema del escenario: El orden se actualiza al cambiar una prioridad
  Dado un proyecto con las historias 1 "Alta", 2 "Media" y 3 "Baja"
  Cuando cambio la prioridad de la historia 3 a <prioridad>
  Y pido el backlog ordenado
  Entonces recibo las historias en el orden <orden>

  Ejemplos:
    | prioridad | orden   |
    | "Alta"    | 1, 3, 2 |
    | "Media"   | 1, 2, 3 |
    | "Baja"    | 1, 2, 3 |

Escenario: Bajar la prioridad de la historia que estaba primera
  Dado un proyecto con las historias 1 "Alta", 2 "Media" y 3 "Media"
  Cuando cambio la prioridad de la historia 1 a "Baja"
  Y pido el backlog ordenado
  Entonces recibo las historias en el orden 2, 3, 1

Escenario: Backlog ordenado de un proyecto sin historias
  Dado un proyecto registrado sin historias
  Cuando pido el backlog ordenado
  Entonces recibo una lista vacia
  Y no se informa ningun error

Escenario: Backlog ordenado de un proyecto con una sola historia
  Dado un proyecto con una unica historia de prioridad "Baja"
  Cuando pido el backlog ordenado
  Entonces recibo esa unica historia

Escenario: Todas las historias tienen la misma prioridad
  Dado un proyecto con las historias 1, 2 y 3, todas de prioridad "Media"
  Cuando pido el backlog ordenado
  Entonces recibo las historias en el orden 1, 2, 3

Escenario: No hay historias de prioridad alta
  Dado un proyecto con las historias 1 "Baja" y 2 "Media"
  Cuando pido el backlog ordenado
  Entonces recibo las historias en el orden 2, 1
  Y el listado tiene 2 historias

Escenario: El listado incluye las historias en cualquier estado
  Dado un proyecto con una historia "Pendiente", una "En curso" y una "Terminada"
  Cuando pido el backlog ordenado
  Entonces el listado tiene las 3 historias

Escenario: Las historias del listado traen todos sus datos
  Dado un proyecto con 3 historias cargadas
  Cuando pido el backlog ordenado
  Entonces cada historia del listado trae su identificador, titulo, descripcion, prioridad,
  estado, story points y criterios de aceptacion

Escenario: Pedir el listado ordenado no cambia el backlog
  Dado un proyecto con las historias 1 "Baja", 2 "Alta" y 3 "Media"
  Cuando pido el backlog ordenado
  Y despues pido el backlog sin ordenar
  Entonces el backlog sin ordenar sigue devolviendo las historias en el orden 1, 2, 3
  Y ninguna historia cambio su prioridad

Escenario: Dos listados seguidos devuelven lo mismo
  Dado un proyecto con 3 historias de distinta prioridad
  Cuando pido el backlog ordenado dos veces seguidas sin cambiar nada
  Entonces las dos veces recibo las historias en el mismo orden

Escenario: Cambiar una prioridad por la que ya tenia
  Dado un proyecto con las historias 1 "Alta", 2 "Media" y 3 "Baja"
  Cuando cambio la prioridad de la historia 2 a "Media"
  Y pido el backlog ordenado
  Entonces recibo las historias en el orden 1, 2, 3

Escenario: Pedir el backlog ordenado de un proyecto que no existe
  Dado que no existe ningun proyecto con identificador 99
  Cuando pido el backlog ordenado del proyecto 99
  Entonces se informa que no se encontro el proyecto
  Y no recibo ningun listado
```

**Para decidir en la puerta:**

1. **RN-3 — el desempate dentro de la misma prioridad es por identificador.** Las otras
   opciones eran dejar el orden en que se cargaron las historias (que hoy es el mismo, porque
   el identificador crece con cada alta) o permitir un orden manual dentro de cada prioridad.
   Se eligió el identificador porque es el único dato estable de la historia y porque hace
   que el listado sea siempre el mismo; el orden manual es bastante más trabajo y la issue no
   lo pide. **Si el equipo quiere orden manual, es otra historia.**
2. **RN-4 — el listado incluye también las historias Terminadas.** Se podría argumentar que,
   como el listado se usa para planificar el próximo sprint, las Terminadas molestan. Se
   decidió incluirlas porque el backlog es el del proyecto entero y filtrar sería una
   decisión de quien lo muestra, no del orden. Vale confirmarlo en la puerta.
3. **El listado ordenado se pide aparte: el backlog de la #6 se sigue devolviendo como está.**
   La alternativa era que el backlog pasara a devolverse siempre ordenado por prioridad y no
   hubiera dos formas de pedirlo. Se dejaron las dos para no cambiar el comportamiento que la
   #6 ya tiene aprobado y testeado, pero si la puerta prefiere una sola forma, conviene
   decidirlo ahora y no después de implementarlo.
4. **Asignar la prioridad ya está hecho.** El primer criterio de aceptación de la issue
   ("se puede asignar prioridad alta, media o baja a cada historia") se cumple con la #6: acá
   no se agrega ninguna forma nueva de asignarla ni de cambiarla. Si la puerta quiere que la
   #7 tenga su propia operación de cambiar prioridad, hay que decidirlo ahí, porque cambia el
   alcance y la estimación.
5. **No se contempla una historia sin una de las tres prioridades**, porque la #6 no deja
   guardarla. Si alguna vez se pudiera, habría que decidir dónde va en el listado.

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios:

## Fase 2 — Diseño

_Se escribe después de que la Puerta 1 esté aprobada._

## Fase 3 — Tareas

_Se escribe después de que la Puerta 2 esté aprobada._

## Fase 4 — Implementación

_Se escribe después de que la Puerta 3 esté aprobada._

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #7 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | se completa en la Fase 3 |
| Código Go | se completa en la Fase 2 |
