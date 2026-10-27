# Contrato del dominio

Acordado el **20/10/2026** por el Agile Enabler. Vigente hasta el cierre del proyecto.

Este documento existe por una sola razón: **que nadie espere a nadie**. Las historias que
quedan están encadenadas (`#8 → #9 → #10 → #16 → #17/#18 → #19`) y, si cada uno espera a que
la anterior esté en `dev`, el dashboard y los reportes no llegan al 2 de noviembre.

Acá se fija, **antes de que nadie implemente**, qué entidades existen y qué datos tiene cada
una. Con eso, quien consume una entidad que todavía no está puede escribir su historia
completa —spec, tests y código— armando los datos de prueba a mano en su propio test.

**No es una especificación.** No dice cómo se comporta nada: eso va en la spec de cada
historia, en `docs/specs/`. Acá solo están los nombres y los datos.

## Reglas de uso

1. **Antes de implementar una entidad compartida, se mira que esté acá.** Si falta algo, se
   agrega acá primero.
2. **El dueño de la entidad es el dueño de la historia que la crea** (la columna *Dueño*).
   Es el único que define su comportamiento.
3. **Si hay que romper el contrato, se avisa al grupo antes de cambiarlo**, no después. Un
   cambio silencioso rompe el trabajo de otra persona que ya está escrito.
4. Se respetan las convenciones que ya están en `AGENTS.md`: identificadores en inglés,
   mensajes y comentarios en castellano, campos privados con métodos de lectura, errores como
   valores del paquete, todo en `internal/domain`.

## Lo que ya existe en `dev`

Nada de esto se toca sin avisar: ya está implementado y con tests. Al 27/10:

| Entidad | Datos | Dónde |
|---|---|---|
| `ProjectRegistry` | los proyectos y el contador de números; `r.Create`, `r.Get`, `r.Update`, `r.Projects()` | `project_registry.go` |
| `Project` | `ID`, `Name`, `Description`, `Start`, `End`; su estado con `p.Status()` (#5) | `project.go`, `project_status.go` |
| `Member` | `Name`, `Role` (`RoleAgileEnabler`, `RoleProductBuilder`) | `member.go` |
| `Story` | `ID`, `Title`, `Description`, `Priority`, `State`, `StoryPoints`, `EstimatedHours`, `AcceptanceCriteria`; `p.Backlog()` y `p.PrioritizedBacklog()` (#7) | `story.go`, `story_priority.go`, `story_estimation.go`, `story_hours.go` |
| `Sprint` (#8 a #11), `Effort` (#14), `Defect` (#15), `Metrics` (#16) y `VotingRound` (#13, #73) | lo que dicen sus secciones de abajo, que ya están implementadas | ver cada sección |

**Lo que todavía no está en `dev`:** `Report` (#18 y #19) y `Dashboard` (#17).

Las operaciones cuelgan del proyecto: `p.AddStory(...)`, `p.Members()`, `p.Story(id)`.
Lo nuevo sigue el mismo camino.

**Agregados de la #77 (interfaz web):** el registro suma `r.Projects()`, que lista todos los
proyectos en orden de número, y el dominio suma `StoryPointScale()`, que entrega una copia de
la escala válida para ofrecerla como opciones. Los dos los implementa Emiliano, como parte de la #77. La conversión
a JSON vive en `internal/domain/registry_json.go`: **toda entidad nueva entra ahí** (ver
`AGENTS.md`).

**Cambio a lo que ya existe: `Story` suma `EstimatedHours`** (decimal, cero mientras no se
cargue), con su operación `p.EstimateHours(id, hours)`. Lo agrega la #16, que es la que lo
necesita para la desviación, y sigue las mismas reglas que los story points: mayor que cero,
se puede reemplazar, no se toca si la historia está Terminada. **Es el único cambio a
`story.go`**, y se avisa al grupo antes de hacerlo.

## Entidades nuevas

### `Sprint` — historias #8, #9, #10, #11

**Dueño:** Agustín (#8, #9, #11) y Emiliano (#10).

| Dato | Tipo | Qué es |
|---|---|---|
| `Number` | entero | número del sprint **dentro del proyecto**, lo asigna el sistema empezando en 1 |
| `Goal` | texto | el Sprint Goal, obligatorio |
| `Start` | fecha | día de inicio |
| `End` | fecha | día de cierre planificado |
| `State` | `SprintOpen`, `SprintClosed` | en qué etapa está |
| `Stories` | lista de identificadores de historia | las que se le asignaron; **no se vacía nunca**, ni al cerrar |
| `CompletedStories` | lista de identificadores de historia | las que estaban terminadas al cerrar; vacía mientras está abierto |
| `PointsAtClose` | story points por historia | los puntos que tenía cada historia de `Stories` al cerrar; vacío mientras está abierto |
| `ClosedOn` | fecha | cuándo se cerró de verdad; vacía si no cerró |

Cuelga del proyecto: `p.AddSprint(...)`, `p.OpenSprint()`, `p.Sprints()` y
`p.Sprint(number)` (#8), `p.AssignStory(id)` y `p.UnassignStory(id)` (#9),
`p.CloseSprint(...)` (#10) y `p.ClosedSprints()` (#11).

**`p.Sprints()` y `p.Sprint(number)` los adelanta la #8**, con el error de sprint no
encontrado, aunque al principio eran de la #11: la #15 los necesita desde su primera tarea para
validar los sprints de detección y de resolución de un defecto. Lo pidieron Maxi y Juliana en
la revisión de las puertas de la #8 (PR #104), el 20 y 21/10.

**Por qué `CompletedStories` es un dato y no se deduce.** Una historia que no se terminó
vuelve al Product Backlog y se puede terminar en el sprint siguiente: si "completada" se
dedujera del estado actual de la historia, contaría en los dos sprints y la velocidad del
primero quedaría inflada. Se fija al cerrar y no cambia más. Lo detectó la spec de la #10.

**Por qué `PointsAtClose`.** Por la misma razón: una historia que volvió al backlog se puede
reestimar, y si los planificados de un sprint cerrado se calcularan con los puntos actuales,
el número de un sprint viejo cambiaría sin que nadie lo toque. Decidido el 20/10.

**Son solo dos estados, no tres.** La Fase 1 de la #8 descartó un estado "Planificado": el
sprint se crea ya abierto, y solo puede haber uno abierto por proyecto a la vez. El
razonamiento está en su sección "Para decidir en la puerta", punto 4.

**El número es por proyecto, no global.** Dos proyectos pueden tener los dos su sprint 1.

**Lo que necesita saber quien lo consume:** un sprint cerrado tiene `State` en
`SprintClosed`, `ClosedOn` cargada y `CompletedStories` fija. Las historias que no se
terminaron **vuelven al Product Backlog** al cerrar (pasan a Pendiente), pero siguen en
`Stories`: así la #16 sabe cuánto se había comprometido. Los puntos de un sprint cerrado —planificados
y completados— salen de `PointsAtClose` y `CompletedStories`, nunca de los puntos ni del
estado actual de las historias.

### `Effort` — historia #14

**Dueño:** Juliana.

| Dato | Tipo | Qué es |
|---|---|---|
| `StoryID` | entero | sobre qué historia se trabajó |
| `MemberName` | texto | quién lo registró |
| `Activity` | texto | qué se hizo, obligatorio (lo pide el primer criterio de la #14) |
| `Hours` | decimal | horas reales, mayor que cero |
| `Day` | fecha | qué día |

Se registra sobre el proyecto: `p.LogEffort(...)`, `p.Efforts()`, `p.StoryHours(id)` y
`p.MemberHours(name)`.

**Lo que necesita saber quien lo consume:** las horas de una historia son la suma de todos
sus registros. **El esfuerzo no tiene sprint**: las horas reales de un sprint (#16) se
calculan sumando las de las historias asignadas a ese sprint (`Sprint.Stories`). Si una
historia pasó por dos sprints, todas sus horas caen en el último donde estuvo asignada; es
una simplificación aceptada y queda escrita en la spec de la #14.

### `Defect` — historia #15

**Dueño:** Juliana.

| Dato | Tipo | Qué es |
|---|---|---|
| `ID` | entero | identificador dentro del proyecto |
| `Title` | texto | obligatorio |
| `Description` | texto | obligatorio |
| `Severity` | `SeverityCritical`, `SeverityMajor`, `SeverityMinor` | gravedad |
| `State` | `DefectOpen`, `DefectInProgress`, `DefectResolved` | en qué anda |
| `StoryID` | entero | a qué historia pertenece |
| `DetectedIn` | número de sprint | sprint donde se detectó |
| `ResolvedIn` | número de sprint | sprint donde se resolvió; vacío si sigue abierto |

Cuelga del proyecto: `p.AddDefect(...)`, `p.Defects()`, `p.ChangeDefectState(id, state)`,
`p.ResolveDefect(id, sprintNumber)`, `p.DefectsByState(state)` y
`p.DefectsBySeverity(severity)`.

**`DefectResolved` solo se alcanza con `ResolveDefect`**, porque resolver exige el sprint de
resolución. Un defecto resuelto no se reabre.

### `Metrics` — historia #16

**Dueño:** Emiliano.

Es un resultado calculado, no algo que se guarda. Se pide por sprint o por proyecto entero.

| Dato | Tipo | Qué es |
|---|---|---|
| `PlannedPoints` | entero | story points comprometidos |
| `CompletedPoints` | entero | story points de las historias terminadas |
| `Velocity` | entero | puntos completados del sprint |
| `EstimatedHours` | decimal | horas estimadas |
| `ActualHours` | decimal | horas registradas en `Effort` |
| `HoursDeviation` | decimal | diferencia en porcentaje entre estimadas y reales |
| `StoriesPlanned` | entero | cuántas historias entraron |
| `StoriesCompleted` | entero | cuántas se terminaron |
| `CompletionRate` | decimal | porcentaje de historias completas |
| `DefectsDetected` | entero | defectos detectados |
| `DefectsResolved` | entero | defectos resueltos |

Se pide así: `p.SprintMetrics(sprintNumber)` y `p.ProjectMetrics()`.

**La velocidad del proyecto** es el promedio de la velocidad de los sprints cerrados,
redondeado a entero. Sin sprints cerrados, es cero.

**Un sprint abierto también tiene métricas**: las planificadas y las horas reales se
calculan igual, y los completados y la velocidad dan cero hasta que cierra, porque
`CompletedStories` recién se fija al cerrar.

**Lo que necesita saber quien lo consume (#17 y #18):** `Metrics` es un dato plano, sin
métodos de cálculo. Se puede armar a mano en un test para probar el dashboard o el reporte
sin que la historia #16 esté implementada.

### `Report` — historias #18 y #19

**Dueño:** Agustín (#18) y Emiliano (#19).

| Dato | Tipo | Qué es |
|---|---|---|
| `Title` | texto | encabezado del reporte |
| `GeneratedOn` | fecha | cuándo se generó |
| `Scope` | `ScopeProject`, `ScopeSprint` | de qué es |
| `Metrics` | `Metrics` | los números |
| `Sections` | lista de `Section` | ver abajo |

Se pide: `p.ProjectReport(today)` y `p.SprintReport(sprintNumber, today)`: el día lo pasa quien
llama, porque el dominio no lee el reloj.

**`Section` es del reporte.** Tiene `Heading` (texto) y `Lines` (lista de textos ya armados).
Lo que comparten el reporte (#18) y el dashboard (#17) es **la regla de formato de los
números** (horas con "h", el signo de la desviación, los decimales): la escribe la primera de
las dos historias que llegue a `dev`, como pasó con `errors.go`, y la otra la reutiliza. Así
un número se lee igual en el dashboard, en el reporte y en el PDF.

### `Dashboard` — historia #17

**Dueño:** Juliana. Cambió el 20/10, cuando el dashboard pasó a mostrarse en la web con
gráficos de verdad: el dominio entrega los valores ya escritos y **la web dibuja las barras**.

| Dato | Tipo | Qué es |
|---|---|---|
| `HasSprints` | booleano | si el proyecto tiene algún sprint; sin sprints, la página muestra solo el aviso |
| `Overview` | lista de `Figure` | el estado general del proyecto |
| `CurrentSprint` | el sprint abierto, o nada | `Number`, `Goal`, `Start`, `End` y `Figures` (sin completados ni velocidad) |
| `Charts` | lista de `Chart` | los cuatro gráficos, en orden |

- `Chart`: `Title`, `Bars` (lista de `Bar`) y `Empty` (el aviso cuando no tiene barras, por
  ejemplo "Todavía no hay sprints cerrados").
- `Bar`: `Label`, `Value` (decimal, para medir la barra) y `Text` (el valor ya escrito con la
  regla de formato).
- `Figure`: `Label` y `Text` (ya escrito).

Se pide con `p.Dashboard()`, que no falla: si el proyecto no existe, lo informa el registro.

**Lo que necesita saber la historia #19:** el reporte se arma una sola vez y la exportación a
PDF **solo lo formatea**. No vuelve a calcular nada. Se exporta con `r.WritePDF(w)`, que **escribe el PDF en un
`io.Writer`** y solo lee `Title`, `GeneratedOn` y `Sections`. El dominio no toca archivos: **la web
lo manda como descarga (#78)** y la consola no exporta, porque queda con lo de la #24
(decidido el 20/10, al escribir la spec de la #78; reemplaza a `r.ExportPDF(path)`).

**`Metrics` se declara antes que su cálculo.** Es un dato plano y sus campos son públicos, porque
la #17 y la #18 lo arman a mano en sus tests. Lo declara la #16 en un PR chico, sin
comportamiento y por lo tanto sin test, para que la #17, la #18 y la #19 no esperen a que
existan el sprint, el esfuerzo y los defectos (decidido el 21/10).

### `VotingRound` — historias #13 y #73

**Dueño:** Maximiliano.

| Dato | Tipo | Qué es |
|---|---|---|
| `StoryID` | entero | qué historia se está estimando |
| `Number` | entero | número de ronda, empieza en 1 |
| `Votes` | lista de votos | cada uno con `MemberName` y `Points` |
| `Revealed` | booleano | si ya se mostraron los votos |
| `Agreed` | entero | estimación acordada; vacío hasta que se acuerda |

Además guarda `Participants`: los nombres de los integrantes que tenía el proyecto al abrir
la ronda. Se fijan ahí y no cambian aunque después se sume o se vaya alguien.

Se pide: `p.OpenVotingRound(storyID)`, `r.Vote(member, points)` —`member` es el nombre y
apellido, como en la #14—, `r.Reveal()`, `r.Voted()` y `r.Pending()` (solo nombres),
`r.Votes()` (los valores; **con la ronda sin revelar devuelve error**, no una lista vacía),
y `p.CurrentVotingRound(storyID)` (la última ronda de la historia, para volver a encontrarla
entre un pedido y otro de la web); para la #73, `r.Differences()`, `r.Agree(points)` y
`p.VotingRounds(storyID)`. Las rondas entran en la conversión a JSON como toda entidad nueva.

Las firmas, tal como están en `dev`:

| Operación | Devuelve |
|---|---|
| `p.OpenVotingRound(storyID)` | `(*VotingRound, error)` |
| `r.Vote(member, points)`, `r.Reveal()`, `r.Agree(points)` | `error` |
| `r.Voted()`, `r.Pending()` | `[]string` |
| `r.Votes()` | `([]Vote, error)`, cada `Vote` con `MemberName` y `Points` |
| `r.Differences()` | `(Differences, error)`, con `HasDifferences`, `Lowest` y `Highest` |
| `r.Agreed()` | `int`, 0 mientras no se acordó |
| `p.CurrentVotingRound(storyID)` | `(*VotingRound, bool, error)`: el `bool` dice si la historia tiene alguna ronda |
| `p.VotingRounds(storyID)` | `([]*VotingRound, error)` |

`Agreed` vacío se guarda como 0, igual que "sin estimar" en la #12.

**La regla que define la historia #13:** mientras `Revealed` sea falso, **ningún voto se
puede leer** desde afuera, ni siquiera el propio. Solo se sabe quién votó y quién no.

## Dos decisiones transversales, ya tomadas

Aparecieron en varias specs a la vez y las decidió el Agile Enabler el 20/10. Están escritas
en las *Decisiones de arquitectura* de [AGENTS.md](../AGENTS.md):

1. **El dominio no lee el reloj del sistema.** Cuando una regla necesita la fecha de hoy
   —rechazar horas a futuro en la #14, fechar el reporte en la #18—, la recibe de quien
   llama. Por eso `p.LogEffort(...)`, `p.ProjectReport(today)` y `p.SprintReport(number,
   today)` reciben el día.
2. **El núcleo sí arma texto formateado.** El dashboard (#17), el reporte (#18) y el PDF
   (#19) entregan sus líneas ya armadas desde `internal/domain`. Lo que sigue prohibido es
   imprimir o leer del teclado.

## Qué puede empezar hoy sin esperar nada

*Histórico: es la tabla del 20/10, cuando se acordó el contrato. Al 27/10, la #16 y la #73
están en `dev`, y la #19 empezó su código sin esperar a la #18 (PR #184, en revisión).*

| Historia | Depende de | Cómo arranca igual |
|---|---|---|
| #16 Métricas | sprints cerrados, esfuerzo, defectos | arma un sprint cerrado y sus registros en el propio test |
| #17 Dashboard | #16 | arma un `Metrics` a mano en el test |
| #18 Reporte | #16 | ídem |
| #19 PDF | #18 | arma un `Report` a mano en el test |
| #73 Acuerdo | #13 | arma una ronda revelada en el test |
