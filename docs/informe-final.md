# Informe final del proyecto — esqueleto

> **Borrador, 26/10.** Ya tiene cargados el Sprint 0, el 1 y el 2, la cobertura al 26/10, la
> evidencia de TDD y la trazabilidad de la #10. Es la estructura del informe en PDF que se entrega el
> 1/11. Lo arma Emiliano Sorato (Agile Enabler) con las métricas del repo, y Emiliano lo
> revisa antes de exportarlo. Cada sección dice de dónde salen sus datos. Lo que está entre
> corchetes, **[pendiente]**, se completa al cierre del producto (30/10) y del Sprint 4.
>
> El PDF se arma como `.docx` con los estilos de Word (Título 1, 2 y 3) y se exporta desde
> Word. La presentación final la arma Juliana Bustos y no va acá.

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
- Arquitectura en un diagrama: dominio en Go, guardado en archivo, consola y web. Fuente:
  `docs/documentacion-tecnica.md` y `docs/diseno-dominio.md`.

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
y no cambia. El [informe de métricas](informe-de-metricas.md) explica cómo se calcula cada
uno y por qué el tablero solo no alcanza para un sprint cerrado.

| Métrica | Sprint 0 | Sprint 1 | Sprint 2 | Sprint 3 | Sprint 4 |
|---|---|---|---|---|---|
| Fechas | 07/09 → 28/09 | 28/09 → 16/10 | 20/10 → 25/10 | 26/10 → 30/10 | 31/10 → 01/11 |
| SP planificados | — | 16 | 31 | 28 | [pendiente] |
| SP completados (velocidad) | — | **6** | **39** (31 + 8 adelantados) | [pendiente] | [pendiente] |
| Historias completadas | 2 ítems de preparación | 2 de 5 (40%) | 9 de 9 (100%) | [pendiente] | [pendiente] |
| Horas estimadas | — | 48 | 56,7 | 75,6 | [pendiente] |
| Horas reales trabajadas | no se registraron | 38,25 | 31 | [pendiente] | [pendiente] |
| Desviación de las historias terminadas | — | −11% | −52% (−36% sin las hechas con asistencia de IA) | [pendiente] | [pendiente] |
| Horas por story point completado | — | 2,7 | 1,3 (1,8 sin las hechas con asistencia de IA) | [pendiente] | [pendiente] |
| Defectos detectados / resueltos | 0 / 0 | 2 / 1 | 1 / 2 | [pendiente] | [pendiente] |

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
(8 SP), así que la velocidad fue de 39. Las horas trabajadas en el sprint fueron 31: Agustín
Salinas 20, Maximiliano Eula 8, Emiliano Sorato 3 y Juliana Bustos 0, según el tablero.
Juliana Bustos anotó además 14 horas propias que no están en el tablero (la spec de la #17,
revisiones y gestión). La #5, la #14 y la #15 se implementaron con asistencia de IA. Sus horas
reales son el tiempo de dirección y revisión de Emiliano Sorato, y por eso la desviación da
tan negativa; por eso se muestra también la cuenta sin esas tres.

**Sprint 3 y Sprint 4:** [pendiente: se completan con sus reviews, el 30/10 y el 1/11].

**Historias adelantadas y sin estimar.** Desde el 21/10, quien termina lo suyo adelanta su
próxima historia (decisión 11 del [plan](plan-final-al-2-de-noviembre.md)). Una adelantada no
suma a lo comprometido; si se termina, suma a la velocidad, aparte. La #77 y la #78 entraron
sin story points: sus horas cuentan como trabajadas, pero no entran en la desviación.

## 4. Gráficos

[pendiente] Velocidad por sprint, horas estimadas contra reales y defectos detectados y
resueltos. Fuente: los mismos números de la sección 3, o capturas del dashboard de la
aplicación (#17) cargado con los datos de nuestro propio proyecto.

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

**Un par de ejemplo**, de la #10 (PR #150, 26/10):

- RED: `e2e1137` — *test: cerrar el sprint registrando que se completo y que no (historia 10,
  tarea 2)*. Agrega el test a `internal/domain/sprint_close_test.go` y trae en el mensaje la
  salida del compilador: `sprint.CompletedStories undefined`.
- GREEN: `bf2ac19` — *feat: cerrar registra las completadas y devuelve las demas al backlog
  (historia 10, tarea 2)*. Toca solo `internal/domain/sprint_close.go`.

[pendiente: actualizar los números al cierre del Sprint 3 y del Sprint 4.]

## 7. Trazabilidad de punta a punta: la #10 *Cerrar un sprint*

[pendiente: confirmar con la presentación que la historia que se muestra en la defensa sea
esta.] El recorrido de todas las historias está en [docs/trazabilidad.md](trazabilidad.md).

1. **Historia de usuario** — [issue #10](https://github.com/emisorato1/seguimiento-medicion-ics/issues/10):
   *Como integrante del equipo quiero cerrar un sprint registrando qué historias se
   completaron, para poder medir la velocidad del equipo y planificar el siguiente sprint.*
   Tiene 5 SP, es de Emiliano Sorato y se implementó con asistencia de IA.
2. **Especificación SDD** — [docs/specs/sprint-3/10-cerrar-sprint.md](specs/sprint-3/10-cerrar-sprint.md):
   objetivo, entradas, salidas, reglas de negocio (RN-1 a RN-7 y siguientes), casos límite y
   condiciones de error. Agustín Salinas aprobó sus tres puertas el 25/10, en un
   [comentario de la issue](https://github.com/emisorato1/seguimiento-medicion-ics/issues/10#issuecomment-6406625442),
   con dos ajustes que se aplicaron antes de empezar el código (commit `98a4400`).
3. **Criterio de aceptación** — de los 14 de la spec: *"Se puede cerrar el sprint abierto
   indicando qué historias se completaron, y queda en estado Cerrado con su fecha de
   cierre"*, junto con *"Las historias indicadas como completadas quedan Terminadas"* y *"Las
   historias del sprint que no se indicaron quedan Pendientes, de vuelta en el Product
   Backlog"*.
4. **Escenario BDD** — en la [misma spec](specs/sprint-3/10-cerrar-sprint.md?plain=1#L156):

   ```gherkin
   Escenario: Cerrar el sprint registrando que se completo y que no
     Dado un proyecto con el sprint 1 abierto y las historias 1, 2 y 3 asignadas
     Cuando cierro el sprint el 25/10/2026 indicando que se completaron las historias 1 y 2
     Entonces el sprint 1 queda en estado "Cerrado"
     Y el sprint 1 queda con fecha de cierre 25/10/2026
     Y el sprint 1 registra como completadas las historias 1 y 2
     Y las historias 1 y 2 quedan en estado "Terminada"
     Y la historia 3 queda en estado "Pendiente"
   ```

5. **Test en Go** — `TestCerrarElSprintRegistrandoQueSeCompletoYQueNo`, en
   [internal/domain/sprint_close_test.go:59](../internal/domain/sprint_close_test.go#L59). Arriba
   lleva los comentarios `// Escenario:` con el título del escenario y `// Cubre: RN-1, RN-2,
   RN-3, RN-4`. Entró en el commit RED `e2e1137`.
6. **Código Go** — `(*Project).CloseSprint`, en
   [internal/domain/sprint_close.go:11](../internal/domain/sprint_close.go#L11). Entró en el
   commit GREEN `bf2ac19` y llegó a `dev` con el PR #150.

La #10 tiene 14 escenarios y los 14 tienen su test. En 8 está el ciclo RED → GREEN completo, y
en 6 el test pasó con el código que ya estaba y el commit lo dice.

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
