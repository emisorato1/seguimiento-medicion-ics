# Spec: Priorizar el Product Backlog

- **Historia:** #7
- **Sprint:** 2
- **Responsable:** Agustín Salinas
- **Estado:** Implementada — **en review**
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de la consigna y
  de los criterios de aceptación de la issue. El responsable lo revisa, lo corrige si hace
  falta y lo presenta a las puertas. Revisado por el responsable el 20/10: decidió los cinco
  puntos de la puerta y cambió el segundo, para que las historias Terminadas vayan al final
  del listado (RN-10).
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
   prioridad Alta, después las de prioridad Media y por último las de prioridad Baja. Las
   historias Terminadas van aparte, al final (RN-10).
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
10. **[RN-10]** Cuando se pide el listado ordenado, el sistema debe ubicar las historias
    Terminadas después de todas las demás, ordenadas entre sí con el mismo criterio: por
    prioridad y, a igual prioridad, por identificador. Así lo que queda por hacer aparece
    primero, que es para lo que se usa el listado al planificar.

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
- **Historias en distintos estados**, incluso Terminadas: aparecen todas. Pendiente y En
  curso no cambian el orden; las Terminadas van al final.
- **Una historia Terminada de prioridad Alta:** queda después de cualquier historia no
  terminada, aunque esa sea Baja.
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
- [ ] Las historias Terminadas aparecen al final del listado, después de todas las demás, y
      ordenadas entre sí por prioridad.

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

Escenario: Las historias terminadas van al final
  Dado un proyecto con la historia 1 "Alta" en estado "Terminada", la 2 "Baja" en estado
    "Pendiente", la 3 "Media" en estado "En curso" y la 4 "Baja" en estado "Terminada"
  Cuando pido el backlog ordenado
  Entonces recibo las historias en el orden 3, 2, 1, 4
```

**Decidido por el responsable antes de la puerta (20/10).** Los puntos 1, 3, 4 y 5 quedan como
estaban propuestos; el 2 cambió (ver abajo). Si quien aprueba la puerta no está de acuerdo con
alguno, es el momento de decirlo:

1. **RN-3 — el desempate dentro de la misma prioridad es por identificador.** Las otras
   opciones eran dejar el orden en que se cargaron las historias (que hoy es el mismo, porque
   el identificador crece con cada alta) o permitir un orden manual dentro de cada prioridad.
   Se eligió el identificador porque es el único dato estable de la historia y porque hace
   que el listado sea siempre el mismo; el orden manual es bastante más trabajo y la issue no
   lo pide. **Si el equipo quiere orden manual, es otra historia.**
2. **RN-4 y RN-10 — el listado incluye las historias Terminadas, pero al final.** El borrador
   las mezclaba con las demás, ordenadas solo por prioridad: así una historia Alta ya
   terminada quedaba arriba de todo, tapando lo que hay que planificar, que es justo para lo
   que la issue pide el listado. Dejar que la interfaz las filtre tampoco servía: decidir qué
   importa al planificar es una regla de negocio, y `AGENTS.md` dice que esas viven en el
   dominio. Se descartó también sacarlas del listado, porque dejaría de ser el backlog
   completo (RN-4). Ubicarlas al final cumple las dos cosas.
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
> Aprobó: Maximiliano Eula — Fecha: 2026-10-20 — Comentarios: aprobada sin cambios en el PR #105 (https://github.com/emisorato1/seguimiento-medicion-ics/pull/105#issuecomment-6303588269). Verificó a mano los órdenes esperados de los 15 escenarios.

## Fase 2 — Diseño

Las decisiones transversales del proyecto están en [AGENTS.md](../../../AGENTS.md) y no se
repiten acá. El diseño se apoya en la Fase 2 de la #6: el backlog vive dentro del proyecto y
la prioridad es un tipo propio con sus tres valores. Esta fase registra solo lo propio del
orden.

### Enfoque

Ordenar es **una consulta**: arma una lista nueva con las historias del backlog, la ordena y
la devuelve. La lista que guarda el proyecto no se toca, así que el backlog sin ordenar de la
#6 sigue devolviendo lo mismo (RN-7 y punto 3 de "Para decidir en la puerta").

El orden usa tres criterios, en este orden: si la historia está Terminada, que la manda al
final (RN-10); la importancia de la prioridad (RN-1, RN-2); y, a igual prioridad, el
identificador de menor a mayor (RN-3). **No se guarda nada**: el orden se
arma cada vez que se pide, y por eso una prioridad que se cambia con la modificación de la #6
aparece en su nuevo lugar en el listado siguiente, sin trabajo extra (RN-6).

La lista es nueva, pero las historias son las mismas que guarda el backlog, igual que en el
listado de la #6: agregar o quitar elementos de lo que se recibe no cambia el backlog, y los
datos de cada historia solo cambian con sus operaciones (RN-5).

### Archivos afectados

| Archivo | Se crea o se modifica | Para qué |
|---|---|---|
| `internal/domain/story_priority.go` | se crea | el listado del backlog ordenado por prioridad |
| `internal/domain/story_priority_test.go` | se crea | los tests de los escenarios BDD de la Fase 1, nombrados como cada escenario |

No se modifica `errors.go`: el único error posible es el proyecto inexistente, que ya informa
el registro de la #3.

### Decisiones

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| La importancia sale del orden en que la #6 declaró las prioridades: Alta, Media, Baja | una lista propia de esta historia que diga qué prioridad va primero | con dos fuentes podrían decir cosas distintas. El costo es que ese orden pasa a importar: si alguien reordena las prioridades, el test de la tarea 3 falla y lo avisa |
| El desempate por identificador se hace explícito | confiar en que el backlog está guardado en orden de identificador y ordenar respetando el orden original | hoy coinciden porque el identificador crece con cada alta, pero el día que haya borrado o reordenamiento dejan de coincidir y el listado cambiaría sin que ningún escenario lo note. Lo protege un test propio de la tarea 4 |
| El listado es una lista nueva; el backlog guardado no se reordena | ordenar el backlog guardado, y que la #6 lo devuelva ordenado | rompe la RN-7 y cambia lo que la #6 ya tiene aprobado y testeado |
| El orden se arma en cada pedido | guardar el backlog ordenado y reordenarlo al cambiar una prioridad | obligaría a la modificación de la #6 a acordarse de reordenar, y un olvido no da error; armándolo cada vez, la RN-6 sale sola. Son decenas de historias: no hace falta ahorrarse el ordenamiento |
| Las Terminadas van al final dentro del mismo listado | mezclarlas con las demás, o dejar que la interfaz las filtre | mezcladas tapan lo que hay que planificar, y filtrarlas en la interfaz pondría una regla de negocio fuera del dominio, repetida en la consola y en la web |
| El listado va en `story_priority.go` | agregarlo a `story.go` | mismo criterio que la estimación de la #12: no pisar el archivo de la #6 y dejar la trazabilidad directa |

### Dependencias y restricciones

- **Solo depende de lo que ya está en `dev`**: el backlog, la prioridad y la modificación de
  una historia (#6), y el registro de proyectos (#3). Se puede empezar apenas pase la Puerta 3.
- **No se modifica `story.go`** ni ninguna operación de la #6.
- **No es una entidad nueva del contrato del dominio**: es una operación más del proyecto
  sobre el backlog que ya existe, como dice el contrato ("lo nuevo sigue el mismo camino").
- **La usan la #9**, para elegir en la planning, y la interfaz, para mostrar el backlog.

> **Puerta 2** — ¿el diseño es viable y coherente con el resto del proyecto?
> Aprobó: Maximiliano Eula — Fecha: 2026-10-20 — Comentarios: aprobada sin cambios en el PR #105 (https://github.com/emisorato1/seguimiento-medicion-ics/pull/105#issuecomment-6303588269). Coincide con el orden de las constantes de prioridad de `story.go`.

## Fase 3 — Tareas

Cada tarea es un ciclo de TDD completo: el test que falla en un commit, la implementación que
lo hace pasar en otro. Van en el orden en que hay que hacerlas, y cada una deja el código
compilando y los tests en verde.

Los tests siguen la convención de [AGENTS.md](../../../AGENTS.md): un test por escenario, con
el nombre del escenario.

**Son cinco tareas para quince escenarios, a propósito.** En la #12 seis tareas pasaron en
verde sin código nuevo porque separaban en ciclos lo que salía del mismo código, y su Fase 4
lo anotó como un error de la Fase 3. Acá cada tarea agrupa los escenarios que dependen de la
misma implementación, y cada una tiene al menos un test que falla antes de escribirla.

| # | Tarea | Escenario BDD | Reglas | Estado |
|---|---|---|---|---|
| 1 | El listado ordenado de un proyecto sin historias es una lista vacía, sin error; el de un proyecto inexistente informa que no se encontró | Backlog ordenado de un proyecto sin historias; Pedir el backlog ordenado de un proyecto que no existe | RN-8, RN-9 | hecha |
| 2 | El listado trae todas las historias del backlog con todos sus datos, cualquiera sea su estado y estén estimadas o no, sin omitir ni repetir ninguna | El listado incluye las historias en cualquier estado; Las historias del listado traen todos sus datos; Backlog ordenado de un proyecto con una sola historia | RN-4, RN-5 | hecha |
| 3 | El listado va de Alta a Baja, sin huecos si falta una prioridad; se arma en cada pedido, así que un cambio de prioridad se ve en el siguiente; y no toca el backlog guardado | Listar el backlog ordenado por prioridad; No hay historias de prioridad alta; El orden se actualiza al cambiar una prioridad; Bajar la prioridad de la historia que estaba primera; Cambiar una prioridad por la que ya tenia; Pedir el listado ordenado no cambia el backlog | RN-1, RN-2, RN-6, RN-7 | hecha |
| 4 | A igual prioridad, el listado desempata por identificador de menor a mayor aunque el backlog esté guardado en otro orden, y dos listados seguidos dan lo mismo | Desempatar por identificador dentro de la misma prioridad; Todas las historias tienen la misma prioridad; Dos listados seguidos devuelven lo mismo | RN-3 | hecha |
| 5 | Las historias Terminadas van al final del listado, ordenadas entre sí por prioridad e identificador | Las historias terminadas van al final | RN-10 | hecha |

**Dónde falla cada una antes de implementarse.**

- **La 1**, porque el listado no existe. El test del proyecto inexistente pasa solo, porque lo
  resuelve el registro de la #3, y por eso va acá y no en una tarea propia.
- **La 2**, porque la tarea 1 devuelve siempre la lista vacía.
- **La 3**, porque la tarea 2 devuelve las historias en el orden en que se cargaron: fallan el
  primer escenario, el de "no hay prioridad alta", el esquema y el de bajar la prioridad.
  "Cambiar una prioridad por la que ya tenia" y "Pedir el listado ordenado no cambia el
  backlog" pasarían solos: van acá porque protegen *cómo* se ordena, sobre una lista nueva y
  en cada pedido. Para que pase alcanza con ordenar por prioridad respetando el orden en que
  vienen las historias; el desempate explícito es de la 4.
- **La 4 es la delicada**: con los escenarios de la Fase 1 solos no falla, porque las
  historias que se cargan con la operación de la #6 quedan guardadas en orden de
  identificador. Por eso lleva además un test que arma, desde dentro del paquete, un backlog
  guardado fuera de orden (`TestDesempatarPorIdentificadorAunqueElBacklogEsteGuardadoEnOtroOrden`,
  encabezado con `// Tarea 4:` y la decisión de la Fase 2 que protege). Ese es el que falla.
- **La 5**, porque hasta la 4 el estado no cuenta para el orden: la historia 1, Alta y
  Terminada, sale primera.

### Por qué este orden

- **La 1 va primero** porque crea la operación, con el caso que no necesita ninguna historia.
- **La 2 antes que la 3**: primero que estén todas, después en qué orden. Al revés, un test de
  orden en rojo podría estar fallando por una historia que falta, y no se sabría cuál de las
  dos cosas anda mal.
- **La 3 antes que la 4**: el desempate solo tiene sentido cuando ya hay un orden que
  desempatar.
- **La 5 al final**: suma un criterio por encima de los otros dos, y se ve con claridad recién
  cuando la prioridad y el desempate ya funcionan.

### Cobertura

Las diez reglas quedan cubiertas: RN-1 (3), RN-2 (3), RN-3 (4), RN-4 (2), RN-5 (2),
RN-6 (3), RN-7 (3), RN-8 (1), RN-9 (1), RN-10 (5).

| Criterio de aceptación | Tareas |
|---|---|
| El backlog ordenado devuelve primero Alta, después Media y por último Baja | 3 |
| Dentro de una misma prioridad, por identificador de menor a mayor | 4 |
| La misma cantidad de historias que el backlog, sin omitir ni repetir | 2 |
| Las historias aparecen cualquiera sea su estado y estén estimadas o no | 2 |
| Cada historia trae todos sus datos | 2 |
| Al cambiar la prioridad, el listado siguiente la ubica en su nuevo lugar | 3 |
| Cambiar una prioridad por la misma deja el listado igual | 3 |
| Pedir el listado no cambia los datos ni el orden del backlog sin ordenar | 3 |
| Dos listados seguidos devuelven exactamente lo mismo | 4 |
| Un proyecto sin historias devuelve el listado vacío, sin error | 1 |
| Un proyecto con una sola historia devuelve esa historia | 2 |
| Con todas las historias de la misma prioridad, el orden es por identificador | 4 |
| Un proyecto inexistente informa que no se encontró | 1 |
| Las Terminadas aparecen al final, ordenadas entre sí por prioridad | 5 |

Los quince escenarios de la Fase 1 tienen su tarea en la columna "Escenario BDD".

### Dependencias

- **Ninguna tarea espera a otra historia**: todo lo que necesitan ya está en `dev`.
- **La tarea 2 necesita historias En curso y Terminadas**, y esos cambios de estado son de la
  #9 y la #10: el test fija el estado desde dentro del paquete, como hizo la #12.
- **La tarea 3 cambia prioridades con la modificación de una historia de la #6**, sin agregar
  ninguna forma nueva de hacerlo (punto 4 de "Para decidir en la puerta").

> **Puerta 3** — ¿las tareas cubren todos los criterios de aceptación?
> Aprobó: Maximiliano Eula — Fecha: 2026-10-20 — Comentarios: aprobada sin cambios en el PR #105 (https://github.com/emisorato1/seguimiento-medicion-ics/pull/105#issuecomment-6303588269). Destacó que cada tarea explica dónde falla antes de implementarse.

## Fase 4 — Implementación

Ciclo RED → GREEN → REFACTOR. **El test en un commit y la implementación en otro.**

| Tarea | Commit del test | Commit de la implementación |
|---|---|---|
| 1 | `41f2471` | `d194fa8` |
| 2 | `ab5ea38` | `5c75d88` |
| 3 | `eb1c628` | `f11be53` |
| 4 | `50a9d3f` | `deef1b0` |
| 5 | `cc4e1f9` | `ba0c908` |

Cada commit de test dice en su mensaje qué falló, y todos fallaron antes de su
implementación. Los tests que pasaron solos son los que la Fase 3 anticipó, y cada uno va
en la tarea de la implementación que lo sostiene:

- **Tarea 1:** "Pedir el backlog ordenado de un proyecto que no existe" pasa sin código de
  esta historia, porque lo resuelve el registro de la #3. El RED de la tarea es el otro
  escenario: el listado no existía.
- **Tarea 3:** "Cambiar una prioridad por la que ya tenia", "Pedir el listado ordenado no
  cambia el backlog" y las filas "Media" y "Baja" del esquema pasan con la tarea 2. Fallan los
  otros cuatro.
- **Tarea 4:** los tres escenarios pasan con la tarea 3, porque la #6 guarda las historias en
  orden de identificador. El RED es el test que arma un backlog guardado fuera de orden
  (`TestDesempatarPorIdentificadorAunqueElBacklogEsteGuardadoEnOtroOrden`).

### Desvíos respecto de la spec

- **El listado se llama `PrioritizedBacklog()`** y cuelga del proyecto, como las demás
  operaciones del backlog. La spec no fijaba el nombre y el contrato del dominio tampoco lo
  nombra, porque no es una entidad nueva.
- **La tarea 3 se implementó con un ordenamiento estable solo por prioridad**, y la tarea 4
  lo reemplazó por uno con el desempate explícito. Es lo que pedía la Fase 3: cada tarea
  agrega un criterio y su test falla antes.
- No hubo otros desvíos.

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #7 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | `internal/domain/story_priority_test.go`: un test por escenario, 15 en total, más el de la tarea 4 que arma un backlog guardado fuera de orden |
| Código Go | `PrioritizedBacklog()` en `internal/domain/story_priority.go` |
