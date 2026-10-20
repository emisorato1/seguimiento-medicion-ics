# Spec: Gestionar defectos

- **Historia:** #15
- **Sprint:** 2
- **Responsable:** Juliana Bustos
- **Estado:** Borrador completo — pendiente de las tres puertas
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de la consigna y de
  los criterios de aceptación de la issue. El responsable lo revisa, lo corrige si hace falta
  y lo presenta a las puertas. Revisado por la responsable el 20/10 y presentado a las tres
  puertas por pull request.
- **Depende de:** #3 (el proyecto), #6 (la historia relacionada) y #8 (los sprints de
  detección y de resolución)
- **La necesitan:** #16 (defectos detectados y resueltos), #17 (dashboard), #18 (reporte)

## Fase 1 — Requisitos

### Objetivo

Que el equipo pueda anotar los defectos que va encontrando, decir qué tan graves son, seguir
en qué estado está cada uno y dejar registrado en qué sprint se detectó y en cuál se
resolvió, para llevar el control de la calidad del producto.

### Qué queda afuera

- **Contar los defectos detectados y resueltos** para las métricas y el informe: eso es la
  #16. Acá se registran y se listan, no se cuentan ni se calculan porcentajes.
- **Reabrir un defecto ya resuelto** (ver el punto 3 de "Para decidir en la puerta").
- **Modificar o borrar un defecto** ya registrado: ni su título, ni su descripción, ni su
  severidad, ni su historia, ni su sprint de detección.
- **Asignarle el defecto a un integrante**, darle una **prioridad** distinta de la severidad
  o clasificarlo **por tipo** (funcional, de interfaz, de datos).
- **Registrar la fecha exacta** de detección o de resolución: lo que se registra es el sprint.
- Cómo se muestran los defectos en pantalla: eso es de la interfaz.

### Entradas

Para registrar un defecto:

| Dato | Obligatorio | Descripción |
|---|---|---|
| Proyecto | sí | en qué proyecto se registra |
| Título | sí | cómo se lo nombra, en una frase |
| Descripción | sí | qué pasa |
| Severidad | sí | Crítica, Mayor o Menor |
| Historia relacionada | sí | sobre qué historia apareció |
| Sprint de detección | sí | en qué sprint se lo encontró |

Para cambiar el estado: el proyecto, el defecto y el estado nuevo. Para resolverlo: el
proyecto, el defecto y el sprint de resolución.

El **identificador** no es una entrada: lo asigna el sistema al registrar el defecto (RN-1).

### Salidas

Un defecto registrado dentro del proyecto, con su identificador, en estado Abierto y sin
sprint de resolución. Se pueden consultar todos los defectos del proyecto, los de un estado y
los de una severidad.

### Severidades y estados

| Severidad | Qué significa |
|---|---|
| Crítica | impide usar el sistema o hace perder datos |
| Mayor | una funcionalidad no hace lo que dice, pero se puede seguir trabajando |
| Menor | molesta o se ve mal, pero no impide trabajar |

| Estado | Qué significa |
|---|---|
| Abierto | se detectó y todavía nadie lo está arreglando |
| En curso | alguien lo está arreglando |
| Resuelto | se arregló, y quedó registrado en qué sprint |

### Reglas de negocio

1. **[RN-1]** Cuando se registra un defecto, el sistema debe asignarle el identificador que
   sigue al del último defecto registrado en ese proyecto, empezando por 1. El identificador
   no se repite dentro del proyecto y no cambia.
2. **[RN-2]** El sistema debe registrar de cada defecto: identificador, título, descripción,
   severidad, estado, historia relacionada, sprint de detección y sprint de resolución.
3. **[RN-3]** Cuando se registra un defecto, el sistema debe dejarlo en estado **Abierto** y
   sin sprint de resolución.
4. **[RN-4]** Si se registra un defecto sin título o sin descripción, el sistema debe
   rechazar la operación e informar cuál de los dos falta. Un texto compuesto solo por
   espacios cuenta como vacío.
5. **[RN-5]** Cuando se registran el título y la descripción, el sistema debe guardarlos sin
   los espacios del principio y del final.
6. **[RN-6]** El sistema debe aceptar como severidad únicamente **Crítica**, **Mayor** o
   **Menor**. Si se indica otra, o no se indica ninguna, debe rechazar la operación e
   informar cuáles son las válidas.
7. **[RN-7]** Si la historia relacionada no pertenece al proyecto, el sistema debe rechazar
   la operación e informar que no se encontró la historia. La historia se acepta en cualquier
   estado, incluida una Terminada.
8. **[RN-8]** Si el sprint de detección no existe en el proyecto, el sistema debe rechazar la
   operación e informar que no se encontró el sprint. El sprint de detección se acepta esté
   abierto o cerrado.
9. **[RN-9]** Mientras un defecto no esté Resuelto, el sistema debe permitir pasarlo entre
   Abierto y En curso, en los dos sentidos, incluso al estado que ya tenía.
10. **[RN-10]** Si se intenta pasar un defecto a Resuelto cambiándole el estado, el sistema
    debe rechazar la operación e informar que para resolverlo hay que indicar en qué sprint
    se resolvió.
11. **[RN-11]** Cuando se resuelve un defecto, el sistema debe pasarlo a **Resuelto** y
    registrar el sprint de resolución, esté el defecto Abierto o En curso.
12. **[RN-12]** Si el sprint de resolución no existe en el proyecto, el sistema debe rechazar
    la operación e informar que no se encontró el sprint.
13. **[RN-13]** Si el sprint de resolución es anterior al de detección, el sistema debe
    rechazar la operación e informar que un defecto no se resuelve en un sprint anterior al
    de detección. Resolverlo en el mismo sprint en que se detectó es válido.
14. **[RN-14]** Si el defecto ya está Resuelto, el sistema debe rechazar cualquier cambio de
    estado y cualquier resolución nueva, e informar que el defecto ya está resuelto.
15. **[RN-15]** El sistema debe permitir consultar todos los defectos del proyecto, los de un
    estado y los de una severidad, siempre en el orden en que se registraron. Una consulta
    sin resultados devuelve una lista vacía, que no es un error.
16. **[RN-16]** El sistema debe tratar cada proyecto por separado: los identificadores de
    defecto son únicos dentro de cada proyecto, y un defecto nunca aparece en las consultas
    de otro proyecto.
17. **[RN-17]** El sistema debe permitir que dos defectos del mismo proyecto tengan el mismo
    título, la misma descripción, la misma severidad y la misma historia: lo que los
    distingue es el identificador.
18. **[RN-18]** Si una operación se rechaza, el sistema debe dejar el proyecto y sus defectos
    como estaban antes.
19. **[RN-19]** Si se opera sobre un defecto o sobre un proyecto que no existe, el sistema
    debe informar que no se encontró y no debe cambiar nada.

### Casos límite

- Primer defecto de un proyecto: recibe el identificador 1. Dos proyectos, cada uno con su
  primer defecto: los dos llevan el 1.
- Proyecto sin defectos, o consulta por una severidad que ninguno tiene: lista vacía, y no es
  un error.
- Resolver en el **mismo** sprint de detección: válido, es el límite exacto de la RN-13.
  Resolver en uno anterior: rechazado, aunque ese sprint exista.
- Resolver un defecto Abierto sin que haya pasado nunca por En curso: válido.
- Cambiar el estado de un defecto al estado que ya tenía: válido, no es un error.
- Defecto detectado en un sprint ya cerrado, o sobre una historia Terminada: válido.
- Dos defectos con exactamente los mismos datos: quedan los dos, con identificadores
  distintos.

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| Título o descripción vacíos o de solo espacios | rechaza e informa cuál de los dos falta |
| Severidad que no es Crítica, Mayor ni Menor, o ninguna | rechaza e informa cuáles son las válidas |
| La historia no pertenece al proyecto | rechaza e informa que no se encontró la historia |
| El sprint de detección o el de resolución no existen | rechaza e informa que no se encontró el sprint |
| Se intenta pasar a Resuelto cambiando el estado | rechaza e informa que hay que indicar el sprint de resolución |
| Sprint de resolución anterior al de detección | rechaza e informa que no se resuelve en un sprint anterior al de detección |
| El defecto ya está Resuelto | rechaza e informa que el defecto ya está resuelto |
| Defecto o proyecto inexistente | informa que no se encontró y no cambia nada |

### Criterios de aceptación

- [ ] Se registra un defecto con título, descripción, severidad, historia relacionada y
      sprint de detección, y queda con su identificador, en Abierto y sin sprint de
      resolución. El primero del proyecto recibe el 1 y cada uno nuevo el siguiente.
- [ ] No se puede registrar un defecto sin título o sin descripción; un texto de solo
      espacios cuenta como vacío y lo que se guarda va sin los espacios de los extremos.
- [ ] La severidad solo puede ser Crítica, Mayor o Menor.
- [ ] No se puede registrar un defecto sobre una historia que no es del proyecto, y sí sobre
      una historia Terminada.
- [ ] No se puede registrar un defecto con un sprint de detección que no existe, y sí con uno
      ya cerrado.
- [ ] Un defecto no resuelto se puede pasar de Abierto a En curso y de En curso a Abierto.
- [ ] No se puede pasar un defecto a Resuelto cambiándole el estado: hay que indicar el
      sprint de resolución.
- [ ] Al resolver un defecto queda en Resuelto con su sprint de resolución, viniera de
      Abierto o de En curso.
- [ ] No se puede resolver un defecto en un sprint anterior al de detección ni con un sprint
      que no existe en el proyecto, y sí en el mismo sprint en que se detectó.
- [ ] Un defecto Resuelto no cambia de estado ni se vuelve a resolver.
- [ ] Se pueden listar todos los defectos, los de un estado y los de una severidad, en el
      orden en que se registraron, y una consulta sin resultados devuelve una lista vacía sin
      ser un error.
- [ ] Los defectos de un proyecto no aparecen en las consultas de otro, y cada proyecto
      numera desde 1.
- [ ] Una operación rechazada deja el proyecto y sus defectos como estaban, y operar sobre un
      defecto o un proyecto que no existe informa que no se encontró.

### Escenarios BDD

```gherkin
Escenario: Registrar un defecto
  Dado un proyecto con la historia 1 y el sprint 1
  Cuando registro un defecto con titulo "La velocidad sale en cero", descripcion "El resumen del sprint muestra 0 puntos", severidad "Mayor", sobre la historia 1 y detectado en el sprint 1
  Entonces el defecto queda registrado con identificador 1
  Y queda en estado "Abierto"
  Y queda sin sprint de resolucion

Escenario: El defecto nuevo recibe el identificador siguiente
  Dado un proyecto con el defecto 1 registrado
  Cuando registro otro defecto con los mismos datos
  Entonces el defecto queda registrado con identificador 2
  Y el proyecto queda con 2 defectos

Esquema del escenario: Rechazar un defecto sin titulo o sin descripcion
  Dado un proyecto con la historia 1 y el sprint 1
  Cuando intento registrar un defecto con <dato> vacio
  Entonces la operacion se rechaza
  Y se informa que falta <dato>
  Y el proyecto sigue sin defectos

  Ejemplos:
    | dato             |
    | "el titulo"      |
    | "la descripcion" |

Escenario: Guardar el titulo y la descripcion sin los espacios de los extremos
  Dado un proyecto con la historia 1 y el sprint 1
  Cuando registro un defecto con titulo "  La velocidad sale en cero  "
  Entonces el defecto queda registrado con titulo "La velocidad sale en cero"

Esquema del escenario: Registrar un defecto con cada severidad
  Dado un proyecto con la historia 1 y el sprint 1
  Cuando registro un defecto con severidad <severidad>
  Entonces el defecto queda registrado con severidad <severidad>

  Ejemplos:
    | severidad |
    | "Critica" |
    | "Mayor"   |
    | "Menor"   |

Escenario: Rechazar una severidad que no existe
  Dado un proyecto con la historia 1 y el sprint 1
  Cuando intento registrar un defecto con severidad "Urgente"
  Entonces la operacion se rechaza
  Y se informa que la severidad debe ser Critica, Mayor o Menor
  Y el proyecto sigue sin defectos

Escenario: Rechazar un defecto sobre una historia que no es del proyecto
  Dado un proyecto con el sprint 1 y sin ninguna historia con identificador 99
  Cuando intento registrar un defecto sobre la historia 99
  Entonces la operacion se rechaza
  Y se informa que no se encontro la historia

Escenario: Registrar un defecto sobre una historia terminada
  Dado un proyecto con el sprint 1 y la historia 1 en estado "Terminada"
  Cuando registro un defecto sobre la historia 1
  Entonces el defecto queda registrado

Escenario: Rechazar un defecto con un sprint de deteccion que no existe
  Dado un proyecto con la historia 1 y un solo sprint, el numero 1
  Cuando intento registrar un defecto detectado en el sprint 4
  Entonces la operacion se rechaza
  Y se informa que no se encontro el sprint
  Y el proyecto sigue sin defectos

Escenario: Registrar un defecto detectado en un sprint ya cerrado
  Dado un proyecto con la historia 1 y el sprint 1 en estado "Cerrado"
  Cuando registro un defecto detectado en el sprint 1
  Entonces el defecto queda registrado

Esquema del escenario: Cambiar el estado de un defecto no resuelto
  Dado un proyecto con el defecto 1 en estado <desde>
  Cuando le cambio el estado a <hasta>
  Entonces el defecto 1 queda en estado <hasta>

  Ejemplos:
    | desde      | hasta      |
    | "Abierto"  | "En curso" |
    | "En curso" | "Abierto"  |
    | "Abierto"  | "Abierto"  |

Escenario: Rechazar el pase a Resuelto por cambio de estado
  Dado un proyecto con el defecto 1 en estado "Abierto"
  Cuando intento cambiarle el estado a "Resuelto"
  Entonces la operacion se rechaza
  Y se informa que para resolverlo hay que indicar el sprint de resolucion
  Y el defecto 1 sigue en estado "Abierto"

Esquema del escenario: Resolver un defecto
  Dado un proyecto con el sprint 2 y el defecto 1 detectado en el sprint 1, en estado <desde>
  Cuando resuelvo el defecto 1 en el sprint 2
  Entonces el defecto 1 queda en estado "Resuelto"
  Y queda con el sprint de resolucion 2

  Ejemplos:
    | desde      |
    | "Abierto"  |
    | "En curso" |

Escenario: Resolver un defecto en el mismo sprint en que se detecto
  Dado un proyecto con el defecto 1 detectado en el sprint 2
  Cuando resuelvo el defecto 1 en el sprint 2
  Entonces el defecto 1 queda en estado "Resuelto"
  Y queda con el sprint de resolucion 2

Escenario: Rechazar la resolucion en un sprint anterior al de deteccion
  Dado un proyecto con los sprints 1 y 2, y el defecto 1 detectado en el sprint 2
  Cuando intento resolver el defecto 1 en el sprint 1
  Entonces la operacion se rechaza
  Y se informa que un defecto no se resuelve en un sprint anterior al de deteccion
  Y el defecto 1 sigue sin sprint de resolucion

Escenario: Rechazar la resolucion con un sprint que no existe
  Dado un proyecto con un solo sprint, el numero 1, y el defecto 1 detectado en el sprint 1
  Cuando intento resolver el defecto 1 en el sprint 4
  Entonces la operacion se rechaza
  Y se informa que no se encontro el sprint
  Y el defecto 1 sigue en estado "Abierto"

Esquema del escenario: Un defecto resuelto no se vuelve a tocar
  Dado un proyecto con el defecto 1 en estado "Resuelto" en el sprint 2
  Cuando intento <operacion> el defecto 1
  Entonces la operacion se rechaza
  Y se informa que el defecto ya esta resuelto
  Y el defecto 1 sigue en estado "Resuelto" con el sprint de resolucion 2

  Ejemplos:
    | operacion                       |
    | "cambiarle el estado a Abierto" |
    | "resolver en el sprint 3"       |

Esquema del escenario: Listar los defectos del proyecto
  Dado un proyecto con el defecto 1 "Critica" y "Abierto", el defecto 2 "Menor" y "Abierto" y el defecto 3 "Menor" y "Resuelto"
  Cuando pido <consulta>
  Entonces recibo <resultado>, en el orden en que se registraron

  Ejemplos:
    | consulta                          | resultado               |
    | "todos los defectos"              | "los defectos 1, 2 y 3" |
    | "los defectos en Abierto"         | "los defectos 1 y 2"    |
    | "los defectos de severidad Menor" | "los defectos 2 y 3"    |
    | "los defectos de severidad Mayor" | "una lista vacia"       |

Escenario: Los defectos de un proyecto no aparecen en otro
  Dado un proyecto A con un defecto y un proyecto B sin defectos
  Cuando registro un defecto en el proyecto B
  Entonces el defecto del proyecto B queda registrado con identificador 1
  Y el proyecto B tiene 1 solo defecto

Esquema del escenario: Operar sobre algo que no existe
  Dado un proyecto sin ningun defecto con identificador 99
  Cuando intento <operacion>
  Entonces se informa que no se encontro
  Y no cambia nada

  Ejemplos:
    | operacion                                     |
    | "cambiarle el estado al defecto 99"           |
    | "resolver el defecto 99"                      |
    | "registrar un defecto en el proyecto 99"      |
```

**Para decidir en la puerta:**

1. **El cambio de estado y las consultas por estado y por severidad.** La issue pide seguir
   el estado del defecto y consultarlos por estado y por severidad. **Resuelto en el
   contrato:** [`docs/diseno-dominio.md`](../../diseno-dominio.md) ya nombra las seis
   operaciones (registrar, listar, cambiar el estado, resolver y las dos consultas
   filtradas). No hace falta tocar el contrato antes de implementar.
2. **El título es obligatorio aunque la issue no lo nombre.** La issue pide "descripción,
   severidad, estado, historia y sprint de detección"; el contrato le agrega un título, y acá
   se lo tomó como obligatorio igual que la descripción. Si el equipo lo quiere opcional,
   cambia la RN-4.
3. **Un defecto resuelto no se reabre (RN-14).** Si el arreglo no funcionó, se registra un
   defecto nuevo. La alternativa —poder volver un defecto a Abierto— obligaría a decidir qué
   pasa con el sprint de resolución ya registrado y, sobre todo, cómo lo cuenta la #16: un
   defecto reabierto, ¿sigue contando como resuelto en el sprint anterior? Mientras no se
   pueda reabrir, esa pregunta no existe.
4. **El estado Resuelto solo se alcanza resolviendo (RN-10).** Hay una operación para cambiar
   el estado y otra para resolver, porque resolver necesita un dato más. Si el equipo quiere
   una sola operación, hay que decidir qué hace cuando el destino es Resuelto y no viene el
   sprint.
5. **Los dos sprints se validan contra los sprints del proyecto (RN-8, RN-12).** La
   alternativa era aceptar cualquier número. La consecuencia de validarlos es que un defecto
   no se puede cargar antes de que su sprint exista: con el sprint creándose en la #8, en la
   práctica siempre existe.
6. **No hay responsable, ni prioridad, ni fecha, y contar los defectos es de la #16.** El
   contrato no tiene esos tres datos y la issue no los pide; cuándo pasó se deduce del
   sprint. Acá los defectos se listan, y la #16 cuenta sobre esas listas los detectados y los
   resueltos de cada sprint.
7. **Esta historia necesita que el `Sprint` de la #8 exista en `dev` antes de su tarea 1.**
   El borrador decía que los tests "arman los sprints a mano" mientras la #8 no esté, pero en
   Go un test no puede armar un sprint si el tipo todavía no existe en el paquete: no
   compila. Como la RN-8 valida el sprint de detección desde el primer defecto, la tarea 1
   ya lo necesita. Se pidió a Agustín, dueño del `Sprint`, que adelante a `dev` la tarea 1 de
   la #8 (el sprint y su alta) con un PR temprano, igual que se hace acá con el `Defect`. La
   alternativa —que esta historia defina su propia forma de preguntar por un sprint— rompe
   la regla 2 del contrato, que hace dueño del `Sprint` a quien tiene la #8.

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios: [pendiente]

## Fase 2 — Diseño

Las decisiones transversales del proyecto —un solo paquete `internal/domain`, datos en
memoria, errores como valores del dominio comparables con `errors.Is`, mensajes en castellano
y validar todo antes de cambiar nada— están en [AGENTS.md](../../../AGENTS.md) y no se
repiten acá. Los nombres son los del contrato
([`docs/diseno-dominio.md`](../../diseno-dominio.md)): `Defect`, con `ID`, `Title`,
`Description`, `Severity`, `State`, `StoryID`, `DetectedIn` y `ResolvedIn`, registrado con
`AddDefect`, consultado con `Defects`, `DefectsByState` y `DefectsBySeverity`, cambiado de
estado con `ChangeDefectState` y resuelto con `ResolveDefect`.

### Enfoque

El defecto es una entidad nueva que **cuelga del proyecto**, igual que las historias (#6) y
los integrantes (#4): no existe sin su proyecto, y así el identificador único por proyecto
(RN-1, RN-16) sale solo, sin llevar ningún índice aparte. El identificador lo da un contador
del proyecto, como el de las historias.

Son tres operaciones —registrar, cambiar el estado y resolver— y las tres siguen el mismo
orden: **encontrar lo que se nombra** (el proyecto, el defecto, la historia, el sprint) →
**mirar el estado del defecto** → **validar los datos** → **aplicar**. Validar todo antes de
tocar nada es lo que hace que una operación rechazada deje el defecto como estaba (RN-18),
sin tener que deshacer.

Las consultas filtradas (RN-15) recorren la lista de defectos y arman el resultado en el
momento. La lista se mantiene en orden de registro, así el orden de las tres consultas sale
solo.

### Archivos afectados

| Archivo | Se crea o se modifica | Para qué |
|---|---|---|
| `internal/domain/defect.go` | se crea | el defecto, la severidad, el estado, las tres operaciones y las tres consultas |
| `internal/domain/defect_test.go` | se crea | los tests de los escenarios BDD de la Fase 1, nombrados como cada escenario |
| `internal/domain/errors.go` | se modifica | los errores propios de esta historia |
| `internal/domain/project.go` | se modifica | la lista de defectos y su contador dentro del proyecto |

### Decisiones

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| El identificador lo da un contador del proyecto que nunca retrocede | contar los defectos cargados | contar repite identificadores en cuanto exista borrado; es la misma decisión que tomó la #6 |
| La severidad deja libre su valor "ninguna" y el estado no | que las dos arranquen igual | la severidad se indica al registrar, y una no indicada tiene que rechazarse (RN-6), como el rol de la #4; el estado no se indica nunca, el defecto siempre nace Abierto (RN-3), como la historia en la #6 |
| Resolver es una operación propia, separada del cambio de estado | una sola operación de cambio de estado que, cuando el destino es Resuelto, además recibe el sprint | un dato que a veces se usa y a veces no termina en una operación que nadie entiende; además la RN-13 solo aplica al resolver |
| El texto de la severidad se convierte a severidad en el dominio | que la interfaz lo traduzca y rechace lo inválido | si la consola rechaza un texto inválido por su cuenta, la RN-6 queda duplicada fuera del dominio y sin tests, como ya se decidió en la #4 con los roles |
| Las consultas filtradas recorren la lista en el momento | guardar los defectos agrupados por estado y por severidad | con el cambio de estado habría que mantener los grupos sincronizados, y el volumen no lo justifica |
| Un error propio por cada rechazo | un único error de "defecto inválido" | la Fase 1 da siete mensajes distintos y los tests necesitan distinguir cuál salió |

### Dependencias y restricciones

- **Depende de la #3 y la #6**, que ya están en `dev`: el proyecto y la búsqueda de una
  historia dentro de él.
- **Depende del `Sprint` de la #8, que todavía no está en `dev`.** El contrato ya fija qué
  es un sprint y cómo se lo pide por su número, así que esta historia se escribe contra ese
  contrato. Lo único que necesita de la #8 es poder preguntarle al proyecto si existe el
  sprint con tal número, pero eso exige que el tipo `Sprint` y su alta estén en `dev`: sin
  ellos los tests no compilan (punto 7 de la puerta).
- **Con el `Sprint` en `dev`, los tests cargan los sprints que necesitan** con el alta de la
  #8, sin esperar el resto de esa historia.
- **No se toca `story.go`:** a la historia relacionada solo se la busca, y su estado no se
  mira.
- **`errors.go` es compartido:** los errores de esta historia se agregan ahí, y al traer
  `dev` hay que resolver ese conflicto a mano, como ya pasó en la #12.
- **El contrato no se toca:** ya nombra todas las operaciones (punto 1 de la puerta). Si al
  implementar hiciera falta cambiarlo, se avisa al grupo antes, como pide la regla 3 de
  `docs/diseno-dominio.md`.

> **Puerta 2** — ¿el diseño es viable y coherente con el resto del proyecto?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios: [pendiente]

## Fase 3 — Tareas

Cada tarea es un ciclo de TDD completo: el test que falla en un commit, la implementación que
lo hace pasar en otro. Están en el orden en que hay que hacerlas y cada una deja los tests en
verde. Siguen la convención de [AGENTS.md](../../../AGENTS.md): un test por escenario, con el
nombre del escenario y los errores verificados con `errors.Is`. Cada tarea de rechazo
comprueba además que el defecto o el proyecto quedó como estaba, que es la otra mitad de la
RN-18.

| # | Tarea | Escenario BDD | Reglas | Estado |
|---|---|---|---|---|
| 1 | Registrar un defecto válido: queda con identificador 1, en Abierto y sin sprint de resolución | Registrar un defecto | RN-1, RN-2, RN-3 | pendiente |
| 2 | El defecto siguiente recibe el identificador siguiente, y dos defectos con los mismos datos conviven | El defecto nuevo recibe el identificador siguiente | RN-1, RN-17 | pendiente |
| 3 | Rechazar un defecto sin título o sin descripción, y guardar los dos sin los espacios de los extremos | Rechazar un defecto sin titulo o sin descripcion / Guardar el titulo y la descripcion sin los espacios de los extremos | RN-4, RN-5 | pendiente |
| 4 | Aceptar las tres severidades y rechazar cualquier otra | Registrar un defecto con cada severidad / Rechazar una severidad que no existe | RN-6 | pendiente |
| 5 | La historia relacionada tiene que ser del proyecto, y una Terminada se acepta | Rechazar un defecto sobre una historia que no es del proyecto / Registrar un defecto sobre una historia terminada | RN-7 | pendiente |
| 6 | El sprint de detección tiene que existir en el proyecto, y uno cerrado se acepta | Rechazar un defecto con un sprint de deteccion que no existe / Registrar un defecto detectado en un sprint ya cerrado | RN-8 | pendiente |
| 7 | Pasar el defecto entre Abierto y En curso, en los dos sentidos y al estado que ya tenía | Cambiar el estado de un defecto no resuelto | RN-9 | pendiente |
| 8 | Rechazar el pase a Resuelto por cambio de estado | Rechazar el pase a Resuelto por cambio de estado | RN-10 | pendiente |
| 9 | Resolver un defecto Abierto y uno En curso: quedan Resueltos con su sprint | Resolver un defecto | RN-11 | pendiente |
| 10 | Rechazar la resolución con un sprint que no existe, dejando el defecto como estaba | Rechazar la resolucion con un sprint que no existe | RN-12, RN-18 | pendiente |
| 11 | Rechazar la resolución en un sprint anterior al de detección, y aceptar el mismo sprint | Rechazar la resolucion en un sprint anterior al de deteccion / Resolver un defecto en el mismo sprint en que se detecto | RN-13 | pendiente |
| 12 | Un defecto Resuelto no cambia de estado ni se vuelve a resolver | Un defecto resuelto no se vuelve a tocar | RN-14 | pendiente |
| 13 | Consultar todos los defectos, los de un estado y los de una severidad, en orden de registro, y una consulta sin resultados devuelve lista vacía | Listar los defectos del proyecto | RN-15 | pendiente |
| 14 | Los defectos de un proyecto no aparecen en otro, y cada proyecto numera desde 1 | Los defectos de un proyecto no aparecen en otro | RN-16 | pendiente |
| 15 | Operar sobre un defecto o un proyecto que no existe informa que no se encontró y no cambia nada | Operar sobre algo que no existe | RN-19 | pendiente |

### Por qué este orden

- **La 1 va primero** porque es el camino feliz del alta y deja la consulta de los defectos,
  contra la que verifican todas las demás.
- **De la 2 a la 6** se completa el alta, agregando una validación por vez y de la más propia
  del defecto (título, severidad) a la que sale hacia el proyecto (historia, sprint).
- **La 7 y la 8 son el cambio de estado**, que no se puede probar sin defectos cargados.
- **De la 9 a la 12 es resolver**, primero el camino feliz y después sus tres rechazos: el
  sprint que no existe, el sprint anterior y el defecto ya resuelto. La 12 va al final de ese
  grupo porque necesita un defecto resuelto, que recién existe después de la 9.
- **La 13 y la 14 son las consultas**, que necesitan varios defectos con estados y
  severidades distintas: por eso van cuando las tres operaciones ya funcionan.
- **La 15 va al final** porque el defecto o el proyecto inexistente es el borde más externo.

### Cobertura

Las diecinueve reglas quedan cubiertas: RN-1 (1, 2), RN-2 (1), RN-3 (1), RN-4 (3), RN-5 (3),
RN-6 (4), RN-7 (5), RN-8 (6), RN-9 (7), RN-10 (8), RN-11 (9), RN-12 (10), RN-13 (11),
RN-14 (12), RN-15 (13), RN-16 (14), RN-17 (2), RN-18 (3, 4, 5, 6, 8, 10, 11, 12), RN-19 (15).

| Criterio de aceptación | Tareas |
|---|---|
| Se registra un defecto con sus datos, en Abierto y sin sprint de resolución, y el identificador empieza en 1 y sigue de uno en uno | 1, 2 |
| No se puede registrar sin título o sin descripción, y se guardan sin espacios de más | 3 |
| La severidad solo puede ser Crítica, Mayor o Menor | 4 |
| La historia tiene que ser del proyecto, y una Terminada se acepta | 5 |
| El sprint de detección tiene que existir, y uno cerrado se acepta | 6 |
| Un defecto no resuelto pasa de Abierto a En curso y de vuelta | 7 |
| No se pasa a Resuelto sin indicar el sprint de resolución | 8 |
| Al resolver queda en Resuelto con su sprint, viniera de Abierto o de En curso | 9 |
| No se resuelve en un sprint anterior al de detección ni con uno que no existe, y sí en el mismo | 10, 11 |
| Un defecto Resuelto no cambia de estado ni se vuelve a resolver | 12 |
| Se listan todos, por estado y por severidad, en orden, y sin resultados da lista vacía | 13 |
| Los defectos de un proyecto no aparecen en otro, y cada proyecto numera desde 1 | 14 |
| Una operación rechazada deja todo como estaba, y sobre algo inexistente informa que no se encontró | 3, 4, 5, 6, 8, 10, 11, 12, 15 |

Los veinte escenarios BDD de la Fase 1 tienen su tarea, en la columna "Escenario BDD".

### Dependencias

- **Todas las tareas necesitan la #3 y la #6 en `dev`**, que ya lo están.
- **Las tareas 1, 6, 9, 10, 11 y 12 necesitan sprints**, que son de la #8. No hace falta la
  #8 completa: alcanza con su tarea 1 (el `Sprint` y su alta) en `dev`, que se le pidió
  adelantar a Agustín (punto 7 de la puerta). Si la #8 llega con otra forma de preguntar si
  un sprint existe, se ajusta acá y se anota en los desvíos.

### Entrega en dos pull requests

La #16 (métricas) necesita defectos para contar los detectados y los resueltos de cada
sprint, y en Go su test no puede armar un `Defect` si el tipo no existe en `dev`. Por eso
esta historia se parte en dos, que es el caso 1 de "Cuándo se puede partir una historia" en
`CLAUDE.md` (alguien está bloqueado esperándote):

1. **Primer PR (`Refs #15`):** esta spec con sus tres puertas y la **tarea 1**, que deja el
   `Defect`, su severidad, su estado, el registro y la lista. Necesita el `Sprint` de la #8
   en `dev`.
2. **Segundo PR (`Closes #15`):** las tareas 2 a 15.
- **La tarea 5 necesita una historia Terminada**, y los cambios de estado son de la #9 y la
  #10. El test la prepara desde dentro del paquete, igual que hizo la #12 en su tarea 9.

> **Puerta 3** — ¿las tareas cubren todos los criterios de aceptación?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios: [pendiente]

## Fase 4 — Implementación

_Se completa durante la implementación._

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #15 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | `internal/domain/defect_test.go` |
| Código Go | `internal/domain/defect.go` y sus errores en `internal/domain/errors.go` |
