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

Nada de esto se toca sin avisar: ya está implementado y con tests.

| Entidad | Datos | Dónde |
|---|---|---|
| `Project` | `ID`, `Name`, `Description`, `Start`, `End` | `project.go` |
| `Member` | `Name`, `Role` (`RoleAgileEnabler`, `RoleProductBuilder`) | `member.go` |
| `Story` | `ID`, `Title`, `Description`, `Priority`, `State`, `StoryPoints`, `AcceptanceCriteria` | `story.go` |

Las operaciones cuelgan del proyecto: `p.AddStory(...)`, `p.Members()`, `p.Story(id)`.
Lo nuevo sigue el mismo camino.

## Entidades nuevas

### `Sprint` — historias #8, #9, #10, #11

**Dueño:** Agustín (#8, #9, #11) y Emiliano (#10).

| Dato | Tipo | Qué es |
|---|---|---|
| `ID` | entero | identificador dentro del proyecto |
| `Goal` | texto | el Sprint Goal, obligatorio |
| `Start` | fecha | día de inicio |
| `End` | fecha | día de cierre planificado |
| `State` | `SprintPlanned`, `SprintActive`, `SprintClosed` | en qué etapa está |
| `Stories` | lista de identificadores de historia | las que se le asignaron |
| `ClosedOn` | fecha | cuándo se cerró de verdad; vacía si no cerró |

Cuelga del proyecto: `p.AddSprint(...)`, `p.Sprints()`, `p.Sprint(id)`.

**Lo que necesita saber quien lo consume:** un sprint cerrado tiene `State` en
`SprintClosed`, `ClosedOn` cargada, y sus historias quedan con el estado que tenían al
cierre. Las que no se terminaron **vuelven al Product Backlog**, no se arrastran.

### `Effort` — historia #14

**Dueño:** Juliana.

| Dato | Tipo | Qué es |
|---|---|---|
| `StoryID` | entero | sobre qué historia se trabajó |
| `MemberName` | texto | quién lo registró |
| `Hours` | decimal | horas reales, mayor que cero |
| `Day` | fecha | qué día |

Se registra sobre el proyecto: `p.LogEffort(...)`, `p.Efforts()`.

**Lo que necesita saber quien lo consume:** hay un registro por persona, historia y día. Las
horas de una historia son la suma de todos sus registros.

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
| `DetectedIn` | identificador de sprint | sprint donde se detectó |
| `ResolvedIn` | identificador de sprint | sprint donde se resolvió; vacío si sigue abierto |

Cuelga del proyecto: `p.AddDefect(...)`, `p.Defects()`, `p.ResolveDefect(id, sprintID)`.

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

Se pide así: `p.SprintMetrics(sprintID)` y `p.ProjectMetrics()`.

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
| `Sections` | lista de secciones | cada una con `Heading` y `Lines` |

Se pide: `p.ProjectReport()` y `p.SprintReport(sprintID)`.

**Lo que necesita saber la historia #19:** el reporte se arma una sola vez y la exportación a
PDF **solo lo formatea**. No vuelve a calcular nada.

### `VotingRound` — historias #13 y #73

**Dueño:** Maximiliano.

| Dato | Tipo | Qué es |
|---|---|---|
| `StoryID` | entero | qué historia se está estimando |
| `Number` | entero | número de ronda, empieza en 1 |
| `Votes` | lista de votos | cada uno con `MemberName` y `Points` |
| `Revealed` | booleano | si ya se mostraron los votos |
| `Agreed` | entero | estimación acordada; vacío hasta que se acuerda |

Se pide: `p.OpenVotingRound(storyID)`, `r.Vote(member, points)`, `r.Reveal()`,
`r.Agree(points)`.

**La regla que define la historia #13:** mientras `Revealed` sea falso, **ningún voto se
puede leer** desde afuera, ni siquiera el propio. Solo se sabe quién votó y quién no.

## Qué puede empezar hoy sin esperar nada

| Historia | Depende de | Cómo arranca igual |
|---|---|---|
| #16 Métricas | sprints cerrados, esfuerzo, defectos | arma un sprint cerrado y sus registros en el propio test |
| #17 Dashboard | #16 | arma un `Metrics` a mano en el test |
| #18 Reporte | #16 | ídem |
| #19 PDF | #18 | arma un `Report` a mano en el test |
| #73 Acuerdo | #13 | arma una ronda revelada en el test |
