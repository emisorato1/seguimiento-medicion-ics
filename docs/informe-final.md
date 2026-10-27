# Informe final del proyecto — esqueleto

> **Borrador, 26/10.** Ya tiene cargados el Sprint 0, el 1 y el 2, la cobertura al 26/10, la
> evidencia de TDD y la trazabilidad de la #13. Es la estructura del informe en PDF que queda listo el
> 29/10 y se presenta el 2/11. Lo arma y lo revisa Emiliano Sorato (Agile Enabler), con las métricas del repo, antes de
> exportarlo. Cada sección dice de dónde salen sus datos. Lo que está entre corchetes,
> **[pendiente]**, se completa al cierre del Sprint 3 (28/10) y del Sprint 4 (29/10).
>
> El PDF se arma como `.docx` con los estilos de Word (Título 1, 2 y 3) y se exporta desde
> Word. **Antes de exportar:** la portada va sin símbolos ni logos (solo universidad, facultad,
> materia, título, integrantes con legajo y año), y se limpian los metadatos del Word y del PDF
> (autor y programa que lo generó), como pide el `CLAUDE.md` de la materia. La presentación
> final la arma Juliana Bustos y no va acá.

## Portada

Universidad Tecnológica Nacional, Facultad Regional San Rafael. Ingeniería y Calidad de
Software, 2026. Título: *Seguimiento y medición de proyectos de software*. Integrantes:
Agustín Salinas (legajo 9105), Emiliano Sorato (legajo 8378), Juliana Bustos (legajo 9975) y
Maximiliano Eula (legajo 8665).

## 1. El producto

- Qué hace la aplicación y para quién, en un párrafo.
- Qué requerimientos mínimos de la consigna cubre, en una tabla: requerimiento, historia,
  estado. Fuente: el Product Backlog del tablero y `docs/specs/README.md`.
- Cómo se usa: consola y web. Fuente: `docs/manual-de-usuario.md`.
- Arquitectura en un diagrama: dominio en Go, consola y web, y el guardado de los datos en una
  base de datos SQLite (#177), con el archivo JSON para importar y exportar. [pendiente: la #177
  está en implementación; se confirma cuando se mergee a `dev`.] Fuente:
  `docs/documentacion-tecnica.md` y `docs/diseno-dominio.md`.

![Arquitectura: los paquetes de Go y los archivos de datos](graficos/arquitectura.png)

El diagrama sale de `scripts/diagrama-arquitectura.py`. Las flechas son los import entre
paquetes: la web y la consola solo usan el dominio, y el guardado está aparte, en
`internal/store`.

## 2. Cómo trabajamos

- **Scrum:** roles, calendario de los cinco sprints (0 a 4) y ceremonias. Fuente: el plan y
  las actas de `docs/actas/`.
- **Tablero:** columnas, campos (story points, sprint, horas estimadas y reales) y cómo se
  movía una historia. Fuente: el GitHub Project.
- **SDD, BDD y TDD:** spec con tres puertas antes del código, escenarios Given-When-Then y
  commits separados de test y de código. Fuente: `CLAUDE.md`, `AGENTS.md` y `docs/specs/`.
- **Uso de la IA:** qué historias se hicieron con asistencia de IA y cómo se revisaron. Las
  horas de esas historias son el tiempo de dirección y revisión.

## 3. Métricas por sprint

Los números de un sprint cerrado salen de su acta de review, que se escribe el día del cierre
y no cambia. Si después se estima una historia que se terminó en ese sprint, el recálculo va en
el acta aparte y con fecha, como pasó con la #77 en el Sprint 2. El [informe de métricas](informe-de-metricas.md) explica cómo se calcula cada
uno y por qué el tablero solo no alcanza para un sprint cerrado.

| Métrica | Sprint 1 | Sprint 2 | Sprint 3 | Sprint 4 |
|---|---|---|---|---|
| Fechas | 28/09 → 16/10 | 20/10 → 25/10 | 26/10 → 28/10 | 29/10 |
| SP planificados | 16 | 31 (44 con la #77) | 41 | [pendiente] |
| SP completados (velocidad) | **6** | **52** (44 + 8 adelantados) | [pendiente] | [pendiente] |
| Historias completadas | 2 de 5 (40%) | 10 de 10 (100%) | [pendiente] | [pendiente] |
| Horas estimadas | 48 | 56,7 | 110,7 | [pendiente] |
| Horas reales trabajadas | 38,25 | 31 | [pendiente] | [pendiente] |
| Desviación de las historias terminadas | −11% | −52% (−36% sin las hechas con asistencia de IA) | [pendiente] | [pendiente] |
| Horas por story point completado | 2,7 | 1,0 (1,8 sin las hechas con asistencia de IA) | [pendiente] | [pendiente] |
| Defectos detectados / resueltos | 2 / 1 | 1 / 2 | [pendiente] | [pendiente] |

_El Sprint 0 fue de preparación, sin story points: lo que se hizo está en el párrafo del Sprint 0, abajo._

**Sprint 0 — preparación** ([acta](actas/sprint-0-cierre.md)). No tuvo story points: era
dejar listos el repositorio, el tablero y el Product Backlog. Se cumplió el objetivo: el
repositorio y el Project públicos, el README, los colaboradores, las reglas de trabajo, el CI
y un Product Backlog de 17 historias (issues #3 a #19), sacado de los requerimientos mínimos
de la consigna. Quedaron dos cosas mal: el backlog terminó sin estimar, y la entrega del
campus se subió fuera de término.

**Sprint 1 — MVP** ([planning](actas/sprint-1-planning.md),
[review](actas/sprint-1-review.md), [retro](actas/sprint-1-retrospectiva.md)). Se
comprometieron 16 SP y se terminaron 6: la #3 y la #4. La #12 y la #24 volvieron al backlog.
La #6 se mergeó el 19/10, después del cierre, y por eso cuenta en el Sprint 2. La desviación
total fue de −20% (48 h estimadas contra 38,25 trabajadas) y la de las terminadas, de −11%
(18 contra 16). De ahí sale la conversión de **2,7 horas por story point** que se usó para
estimar los sprints siguientes.

**Sprint 2 — interfaz usable** ([planning](actas/sprint-2-planning.md),
[review](actas/sprint-2-review.md), [retro](actas/sprint-2-retrospectiva.md)). Se
terminaron las 9 historias comprometidas (31 SP) y además la #13, adelantada desde el Sprint 3
(8 SP): al cierre, la velocidad fue de 39. La #77, la base de la web, también se terminó en el
sprint, pero sin story points. El 26/10 el equipo la estimó en 13 y sumó al sprint, que quedó
con 44 SP terminados y velocidad 52. Las horas trabajadas en el sprint fueron 31: Agustín
Salinas 20, Maximiliano Eula 8, Emiliano Sorato 3 y Juliana Bustos 0, según el tablero.
Juliana Bustos anotó además 14 horas propias que no están en el tablero (la spec de la #17,
revisiones y gestión). La #5, la #14 y la #15 se implementaron con asistencia de IA. Sus horas
reales son el tiempo de dirección y revisión de Emiliano Sorato, y por eso la desviación da
tan negativa; por eso se muestra también la cuenta sin esas tres. La #77 tampoco entra en la
desviación, porque no tiene horas estimadas: se estimó cuando ya estaba terminada.

**Sprint 3 y Sprint 4:** [pendiente: se completan con sus reviews, el 28/10 y el 29/10]. Del
Sprint 3, la tabla ya tiene lo planificado: 41 SP y 110,7 horas estimadas.

**Historias adelantadas, agregadas y estimadas después.**

- **Adelantadas.** Desde el 21/10, quien termina lo suyo adelanta su próxima historia (decisión
  11 del [plan](plan-final-al-2-de-noviembre.md)). Una adelantada no suma a lo comprometido; si
  se termina, suma a la velocidad, aparte. Fueron la #13 en el Sprint 2 y la #73 en el 3.
- **Agregadas durante el sprint.** La #177 (guardar los datos en una base de datos) entró al
  Sprint 3 el 26/10, ya empezado. Se cuenta igual que una adelantada: no suma a lo comprometido
  y, si se termina, suma a la velocidad. Sus story points están [pendientes] del Planning Poker.
- **Estimadas después.** La #77 y la #78, las de la web, entraron sin story points. El equipo
  las estimó el 26/10, en la ronda 2 del Planning Poker de la planning del Sprint 3: 13 y 13,
  con los cuatro de acuerdo. La #77 ya estaba terminada y suma al Sprint 2, donde se terminó.
  La #78 entra en el compromiso del Sprint 3, que queda en 41 SP, porque se estimó en esa
  planning y estuvo en el sprint desde el primer día, como decidió Emiliano Sorato. Sus 35,1 h
  estimadas salen de los 13 SP a 2,7 h por punto. La #77 no tiene horas estimadas, así que no
  entra en la desviación.

## 4. Gráficos

Los tres gráficos salen de la tabla de la sección 3, con `scripts/graficos-informe.py`. Se
regeneran cada vez que cambia la tabla. Un sprint con [pendiente] todavía no aparece.

![Velocidad por sprint](graficos/velocidad.png)

![Horas estimadas contra horas reales por sprint](graficos/horas.png)

![Defectos detectados y resueltos por sprint](graficos/defectos.png)

[pendiente: si la #17 llega a `dev`, sumar una captura del dashboard con los datos del
proyecto.]

## 5. Pruebas y cobertura

Medido con `go test -cover ./...` sobre `dev` en `4f7644b`, el 26/10, el primer día del
Sprint 3. Los cierres anteriores salen de las actas de review.

| Paquete | Qué tiene | Sprint 1 (`4481d1c`, 15/10) | Sprint 2 (`b45cbe6`, 24/10) | 26/10 (`4f7644b`) |
|---|---|---|---|---|
| `internal/domain` | el núcleo y las reglas de negocio | **100%** | **100%** | **100%** |
| `internal/store` | el guardado en archivo | — | 96,0% | 96,0% |
| `internal/web` | las pantallas de la web | — | 93,8% | 94,0% |
| `cmd/app` | la consola | 0% (sin lógica todavía) | 96,2% | 96,2% |
| `cmd/web` | el arranque de la web | — | 56,2% | 94,0% |
| **Total** | | **99,2%** | **96,0%** | **96,9%** |
| Tests | | 52 | 263 | 317 |

- Al 26/10 son **317 tests**, y 469 subtests más, que salen de los tests con tabla.
- [pendiente: la cobertura al cierre del Sprint 3 y del Sprint 4.]
- El CI corre formato, análisis estático y los tests con cobertura en cada pull request. Desde
  el arreglo del defecto #108 (PR #116) los corre también en Windows.
- Escenarios BDD con test: ver la sección 6.

## 6. Evidencia de TDD

La regla del proyecto es que el test va en un commit (`test:`, con la salida del RED en el
mensaje) y el código que lo hace pasar, en el siguiente (`feat:`). No se aplasta la historia,
así que el ciclo queda en el historial de `dev`. `scripts/trazabilidad.sh` lee ese historial y
le pone a cada test una marca. Al 26/10, sobre `4f7644b`:

| Qué muestra el historial | Tests |
|---|---|
| Ciclo completo: el test en un commit y el código que lo hace pasar en el siguiente | 250 |
| Ciclo completo, con el test trayendo lo justo para compilar (una firma o una constante) | 3 |
| Sin RED: el test pasó con el código que ya estaba, y el commit lo dice | 44 |
| Sin GREEN: después del test no hay un commit de código | 13 |
| Test y código en el mismo commit (desvío del TDD) | 1 |

El único desvío es `0d22bc7`, de la #6, en el Sprint 1. Los "sin RED" son casos que una tarea
anterior ya resolvía. Desde el 20/10 se dejan dichos en el asunto del commit, para que no
parezca que el test se escribió después.

Los 13 "sin GREEN" no son tests escritos después del código. Son 13 tests del Sprint 1, del 9
y el 13/10: 5 de la #4 y 8 de la #6. Uno por uno, cada test pasa en su propio commit con código
que ya había escrito un par RED → GREEN anterior de la misma historia (por ejemplo, el "no se
encontró" o la copia de la lista). Son "sin RED" que no quedaron marcados, porque esa costumbre
empezó el 20/10.

**Un par de ejemplo**, de la #13 (PR #141), el mismo de la sección 7:

- RED: `ba90caa` — *test: votar registra quién votó sin dejar ver los votos (historia 13, tarea
  5)*. Agrega el test a `internal/domain/voting_round_test.go` y trae en el mensaje la salida
  del compilador: `round.Vote undefined` y `undefined: ErrVotesHidden`.
- GREEN: `6a78fd1` — *feat: votar sin que se pueda leer ningún voto antes de revelar (historia
  13, tarea 5)*.

[pendiente: actualizar los números al cierre del Sprint 3 y del Sprint 4.]

## 7. Trazabilidad de punta a punta: la #13 *Votar en secreto en una ronda de Planning Poker*

Es el Planning Poker con votos ocultos que pide la consigna, y la hizo Maximiliano Eula. De sus
16 escenarios, los 16 tienen el ciclo RED → GREEN completo, y sus 13 commits de test traen en el
mensaje la salida del RED. El recorrido de todas las historias está en
[docs/trazabilidad.md](trazabilidad.md).

1. **Historia de usuario** — [issue #13](https://github.com/emisorato1/seguimiento-medicion-ics/issues/13):
   *Como integrante del equipo quiero votar la estimación de una historia sin ver lo que votaron
   los demás, para que nadie se deje influir y la estimación salga del criterio de cada uno.*
   Tiene 8 SP. Se adelantó desde el Sprint 3 y se terminó en el Sprint 2 (PR #141, 23/10), con
   la revisión de Juliana Bustos.
2. **Especificación SDD** — [docs/specs/sprint-3/13-votar-en-secreto.md](specs/sprint-3/13-votar-en-secreto.md):
   objetivo, entradas, salidas, reglas de negocio (desde la RN-1), casos límite y condiciones de
   error. Agustín Salinas aprobó sus tres puertas el 21/10, en el
   [PR #115](https://github.com/emisorato1/seguimiento-medicion-ics/pull/115#issuecomment-6328013297).
3. **Criterio de aceptación** — de la issue y de la spec: *"Mientras la ronda no esté revelada,
   ningún voto se puede leer desde afuera, ni siquiera el propio. Solo se sabe quién votó y
   quién falta."*
4. **Escenario BDD** — en la [misma spec](specs/sprint-3/13-votar-en-secreto.md?plain=1#L188):

   ```gherkin
   Escenario: Votar registra quien voto sin dejar ver los votos, ni siquiera el propio
     Dado una ronda sin revelar con "Agustin Salinas", "Juliana Bustos" y "Maximiliano Eula"
     Y "Agustin Salinas" voto 3 y "Juliana Bustos" voto 8
     Cuando "Juliana Bustos" pide ver los votos de la ronda
     Entonces se informa que los votos estan ocultos hasta que se revele la ronda
     Y solo se sabe que votaron "Agustin Salinas" y "Juliana Bustos" y que falta "Maximiliano Eula"
   ```

5. **Test en Go** — `TestVotarRegistraQuienVotoSinDejarVerLosVotosNiSiquieraElPropio`, en
   [internal/domain/voting_round_test.go:161](../internal/domain/voting_round_test.go#L161). Arriba
   lleva los comentarios `// Escenario:` con el título del escenario y `// Cubre: RN-8, RN-12,
   RN-13`. Entró en el commit RED `ba90caa` (22/10), que trae en el mensaje la salida del
   compilador: `round.Vote undefined`, `round.Votes undefined` y `undefined: ErrVotesHidden`.
6. **Código Go** — `(*VotingRound).Vote`, en
   [internal/domain/voting_round.go:101](../internal/domain/voting_round.go#L101), y
   `(*VotingRound).Votes`, en [voting_round.go:135](../internal/domain/voting_round.go#L135), que
   devuelve `ErrVotesHidden` mientras la ronda no se revela. Entraron en el commit GREEN
   `6a78fd1` (22/10) y llegaron a `dev` con el PR #141.

**Un detalle para contar en la defensa.** Además de `Votes()`, la ronda tiene
`(VotingRound).Format`, en [voting_round.go:56](../internal/domain/voting_round.go#L56). Sin eso,
imprimir la ronda con `%#v`, la forma de depuración de Go, mostraba los votos ocultos.
Maximiliano Eula lo nombra en la [retro del Sprint 2](actas/sprint-2-retrospectiva.md) entre los
problemas que aparecieron en las revisiones antes de llegar a `dev`. Tiene su propio par RED →
GREEN (tarea 12, commit GREEN `ad93071`).

## 8. Retrospectivas

Lo que se acordó en cada retro y si se cumplió en el sprint siguiente, en una tabla. Fuente:
las actas `sprint-N-retrospectiva.md`. Sprint 1 y Sprint 2: cargados en sus actas. Sprint 3 y
4: [pendiente].

## 9. Conclusiones

[pendiente] Qué aprendimos midiéndonos: cómo cambió la estimación sprint a sprint, qué costó
más de lo pensado y qué haríamos distinto. Sale de la sección de conclusiones del informe de
métricas.

## Anexos

Enlaces al repositorio, al tablero, a las actas, a las specs, al manual de usuario y a la
documentación técnica.
