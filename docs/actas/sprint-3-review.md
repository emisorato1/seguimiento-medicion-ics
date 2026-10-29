# Acta — Review del Sprint 3

> **Borrador armado el 27/10 y puesto al día el 29/10 con los datos del tablero.** Se cierra
> con `scripts/metricas.sh "Sprint 3"` y con lo que diga cada uno. Lo que está entre
> corchetes, **[pendiente]**, todavía no se sabe o nadie lo dijo: acá no se escribe nada que
> alguien no haya dicho. La retrospectiva va en [un archivo aparte](sprint-3-retrospectiva.md).

- **Fecha:** [pendiente: el día del cierre, y si fue por escrito o con reunión]
- **Participantes:** [pendiente]
- **Facilita:** Emiliano Sorato (Agile Enabler)
- **Sprint:** del lunes 26 al miércoles 28/10. Al principio iba hasta el viernes 30/10, y el
  26/10 se acortó a pedido del equipo, para terminar todo el 29/10.
- **Sprint Goal:** que el sistema mida el proyecto y lo muestre: métricas, dashboard y reporte.
- **¿Se cumplió el Sprint Goal?** [pendiente]

## Qué se completó

_Al 29/10, según el tablero._

Una historia cuenta como completada solo si cumple la Definición de Terminado: spec con sus
tres puertas, escenarios BDD, tests en Go que pasan, revisada por otro integrante y mergeada
a `dev`.

| # | Historia | SP | Estado al 29/10 | Responsable | Evidencia |
|---|---|---|---|---|---|
| 10 | Cerrar un sprint | 5 | **Hecho** | Emiliano Sorato (con asistencia de IA) | PR #150, 26/10 |
| 16 | Calcular las métricas | 8 | **Hecho** | Emiliano Sorato (con asistencia de IA) | PRs #81, #140 y #151 |
| 11 | Consultar sprints anteriores | 2 | **Hecho** | Agustín Salinas | PR #143, 26/10 |
| 17 | Mostrar el dashboard | 8 | En review | Juliana Bustos | PR #209, abierto al 29/10 |
| 18 | Generar el reporte | 5 | En review | Agustín Salinas | PR #181, abierto al 29/10, con las tres puertas aprobadas el 27/10 por Emiliano Sorato. [pendiente: el cierre, cuando entre a `dev`] |
| 78 | Operar todo desde la web | 13 | En progreso | Emiliano Sorato (con asistencia de IA) | PRs #149, #163, #165 y #170 en `dev`; 17 de 22 escenarios con test al 29/10 |

- **Story points comprometidos:** 41. Son los 28 de la planning más los 13 de la #78, que se
  estimó en la misma planning ([planning del Sprint 3](sprint-3-planning.md)).
- **Story points completados:** [pendiente: al cierre]. Al 29/10 son 15 (37%).
- **Velocidad del Sprint 3:** [pendiente: al cierre]. Al 29/10 son 20 (15 + 5 de la #73
  adelantada). La #177 no suma: se terminó sin story points.
- **Historias completadas:** [pendiente: al cierre]. Al 29/10 son 3 de 6 (50%).

Las historias no terminadas al cierre **vuelven al Product Backlog**, no se arrastran en
silencio.

### Historias adelantadas

| # | Historia | SP | Estado | Responsable | Evidencia |
|---|---|---|---|---|---|
| 73 | Resolver diferencias y acordar la estimación | 5 | **Hecho** | Maximiliano Eula | PR #148, con los arreglos de los defectos #156 (PR #157) y #172 (PR #173) |
| 19 | Exportar el reporte a PDF | 5 | [pendiente: se adelanta desde el Sprint 4 si la #18 entra a `dev` antes del cierre] | Emiliano Sorato (con asistencia de IA) | PR #184 |

### Historias agregadas durante el sprint

| # | Historia | SP | Estado al 28/10 | Responsable | Evidencia |
|---|---|---|---|---|---|
| 177 | Guardar los datos en una base de datos (SQLite) | sin estimar | **Hecho** | Emiliano Sorato (con asistencia de IA) | spec en el PR #180 e implementación en el PR #186 (merge `59bef23`, 28/10) |

## Esfuerzo

_Al 29/10, según el tablero._

| Historia | Estimadas | Reales | Desviación |
|---|---|---|---|
| #10 | 13,5 | 0,5 | −13 h |
| #16 | 21,6 | 1 | −20,6 h |
| #11 | 5,4 | 4,5 | −0,9 h |
| #73 (adelantada) | 13,5 | 3 | −10,5 h |
| **Total de las terminadas** | **54** | **9** | **−45 h (−83%)** |

- **Horas estimadas de lo comprometido:** 110,7 (41 SP a 2,7 h por punto, con las 35,1 h de
  la #78).
- **La #10 y la #16 se hicieron con asistencia de IA.** Sus horas reales son el tiempo de
  Emiliano Sorato para dirigir y revisar, y por eso la desviación da muy negativa. Sin esas
  dos, la #11 y la #73 dan −60% (18,9 h estimadas contra 7,5 reales).
- **La #78 y la #177: 1,5 h reales cada una**, dadas por Emiliano Sorato el 28/10. También son
  su tiempo de dirección y revisión con asistencia de IA. La #78 sigue en progreso, con 35,1 h
  estimadas. La #177 se terminó sin story points: sus horas suman al total, pero no entran en la
  desviación ni en las horas por story point.
- **La #18 tiene 10 h reales** cargadas, y está en review: no suma hasta que entre a `dev`.
- **Horas reales cargadas en el sprint:** 22 al 29/10. Son 17,5 de lo comprometido, 3 de la
  #73 y 1,5 de la #177.
- **Horas por integrante:** [pendiente: cada uno las completa en la review].
- **Horas por story point completado:** [pendiente: al cierre]. Al 29/10, 0,5 (9 h / 20 SP).

## Defectos

| # | Defecto | Severidad | Detectado | Resuelto |
|---|---|---|---|---|
| 156 | Una ronda revelada sin participantes, cargada del archivo, hace caer `Differences()` | — | Sprint 3 (26/10) | Sprint 3 (26/10, PR #157) |
| 172 | Una ronda con un participante repetido con otras mayúsculas, o con espacios de más, queda trabada | — | Sprint 3 (26/10, Maximiliano Eula) | Sprint 3 (26/10, PR #173) |
| 190 | Con una página vieja, algunos errores no se ven o muestran otra página | Menor | Sprint 3 (26/10, prueba exploratoria de la web) | Sprint 3 (27/10, PR #194) |
| 191 | Una palabra muy larga ensancha la página en el celular | Menor | Sprint 3 (26/10, prueba exploratoria) | Sprint 3 (27/10, PR #195) |
| 192 | Las horas aceptan notación científica y no tienen máximo | Mayor | Sprint 3 (26/10, prueba exploratoria) | Sprint 3 (28/10, PR #196) |
| 193 | El filtro de estado en minúsculas no queda elegido en la lista | Menor | Sprint 3 (26/10, prueba exploratoria) | Sprint 3 (27/10, PR #197) |

- **Detectados en el sprint:** 6. **Resueltos:** 6. **Abiertos:** ninguno, al 29/10.
- La #190 a la #193 salieron de la prueba exploratoria de la web del 26/10. Lo que esa prueba
  dejó como decisión, y no como defecto, está en los *Límites conocidos* de la
  [guía de defensa](../guia-de-defensa.md).

## Calidad

[pendiente: medir de nuevo si entra más código antes del cierre.]

| Paquete | 26/10 (`4f7644b`) | 29/10 (`17ec982`) |
|---|---|---|
| `internal/domain` | 100% | 100% |
| `internal/store` | 96,0% | 87,1% |
| `internal/web` | 94,0% | 93,9% |
| `cmd/app` | 96,2% | 96,2% |
| `cmd/web` | 94,0% | 91,1% |
| **Total** | **96,9%** | **95,5%** |
| Tests | 317 | 384 |

Medido con `go test -cover ./...`; los 384 tests pasan. Baja `internal/store`, porque la #177
cambió el guardado en archivo por la base de datos.

## Demostración

[pendiente: lo que cuenta cada uno de su demo.]

## Pendientes

- [ ] Cerrar los números con `scripts/metricas.sh "Sprint 3"` (corrido al 29/10; falta el cierre).
- [x] Las horas reales de la #78 y de la #177: 1,5 h cada una, dadas por Emiliano Sorato el 28/10
  y cargadas en el tablero.
- [ ] Devolver al Product Backlog las historias no terminadas.
- [ ] El pull request de `dev` a `main` con el incremento del Sprint 3, el 28/10 a la noche, y su
  tag `v0.3.0-sprint3`. Cada sprint tiene su incremento: el del Sprint 4 es la entrega final,
  `v1.0.0`, el 29/10 (decisión del 27/10).
- [ ] La [retrospectiva](sprint-3-retrospectiva.md).
