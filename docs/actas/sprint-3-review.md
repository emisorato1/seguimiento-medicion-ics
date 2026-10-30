# Acta — Review del Sprint 3

> Cerrada el 30/10 con `scripts/metricas.sh "Sprint 3"` y los datos del tablero. Acá no se
> escribe nada que alguien no haya dicho. La retrospectiva va en
> [un archivo aparte](sprint-3-retrospectiva.md).

- **Fecha:** 30/10, por escrito. El sprint terminó el 28/10. La #17 y la #18 entraron a `dev`
  el 29/10, un día después, y cuentan en este sprint porque estaban comprometidas en él y la
  review se hizo después (decisión de Emiliano Sorato, 30/10).
- **Participantes:** Emiliano Sorato, Maximiliano Eula y Juliana Bustos, por escrito.
- **Facilita:** Emiliano Sorato (Agile Enabler)
- **Sprint:** del lunes 26 al miércoles 28/10. Al principio iba hasta el viernes 30/10, y el
  26/10 se acortó a pedido del equipo, para terminar todo el 29/10.
- **Sprint Goal:** que el sistema mida el proyecto y lo muestre: métricas, dashboard y reporte.
- **¿Se cumplió el Sprint Goal?** En parte. El sistema mide el proyecto (#16) y lo muestra en el
  dashboard (#17), y el reporte (#18) ya se arma. Lo que faltó es verlo todo desde la web: la
  #78 vuelve al Product Backlog.

## Qué se completó

_Al cierre, según el tablero._

Una historia cuenta como completada solo si cumple la Definición de Terminado: spec con sus
tres puertas, escenarios BDD, tests en Go que pasan, revisada por otro integrante y mergeada
a `dev`.

| # | Historia | SP | Estado al 29/10 | Responsable | Evidencia |
|---|---|---|---|---|---|
| 10 | Cerrar un sprint | 5 | **Hecho** | Emiliano Sorato (con asistencia de IA) | PR #150, 26/10 |
| 16 | Calcular las métricas | 8 | **Hecho** | Emiliano Sorato (con asistencia de IA) | PRs #81, #140 y #151 |
| 11 | Consultar sprints anteriores | 2 | **Hecho** | Agustín Salinas | PR #143, 26/10 |
| 17 | Mostrar el dashboard | 8 | **Hecho** | Juliana Bustos | PR #204 (dominio) y PR #209 (merge `5d1d556`, 29/10) |
| 18 | Generar el reporte | 5 | **Hecho** | Agustín Salinas | PR #181 (merge `556edca`, 29/10), con las tres puertas aprobadas el 27/10 por Emiliano Sorato |
| 78 | Operar todo desde la web | 13 | En progreso: **vuelve al Product Backlog** | Emiliano Sorato (con asistencia de IA) | PRs #149, #163, #165 y #170 en `dev`: 22 de sus 24 tareas. Faltan la 20 y la 21, el reporte y su PDF en la web, que esperaban la #18 y la #19 |

- **Story points comprometidos:** 41. Son los 28 de la planning más los 13 de la #78, que se
  estimó en la misma planning ([planning del Sprint 3](sprint-3-planning.md)).
- **Story points completados:** 28 (68%).
- **Velocidad del Sprint 3:** 33 (28 + 5 de la #73 adelantada). La #177 no suma: se terminó sin
  story points.
- **Historias completadas:** 5 de 6 (83%). Falta la #78.

Las historias no terminadas al cierre **vuelven al Product Backlog**, no se arrastran en
silencio.

### Historias adelantadas

| # | Historia | SP | Estado | Responsable | Evidencia |
|---|---|---|---|---|---|
| 73 | Resolver diferencias y acordar la estimación | 5 | **Hecho** | Maximiliano Eula | PR #148, con los arreglos de los defectos #156 (PR #157) y #172 (PR #173) |
| 19 | Exportar el reporte a PDF | 5 | No se adelantó: la conexión con el reporte de la #18 queda para el Sprint 4 | Emiliano Sorato (con asistencia de IA) | PR #184, las tareas 1 a 12 |

### Historias agregadas durante el sprint

| # | Historia | SP | Estado al 28/10 | Responsable | Evidencia |
|---|---|---|---|---|---|
| 177 | Guardar los datos en una base de datos (SQLite) | sin estimar | **Hecho** | Emiliano Sorato (con asistencia de IA) | spec en el PR #180 e implementación en el PR #186 (merge `59bef23`, 28/10) |

## Esfuerzo

_Al cierre, según el tablero._

| Historia | Estimadas | Reales | Desviación |
|---|---|---|---|
| #10 | 13,5 | 0,5 | −13 h |
| #16 | 21,6 | 1 | −20,6 h |
| #11 | 5,4 | 4,5 | −0,9 h |
| #17 | 21,6 | 25 | +3,4 h |
| #18 | 13,5 | 10 | −3,5 h |
| #73 (adelantada) | 13,5 | 3 | −10,5 h |
| **Total de las terminadas** | **89,1** | **44** | **−45,1 h (−51%)** |

- **Horas estimadas de lo comprometido:** 110,7 (41 SP a 2,7 h por punto, con las 35,1 h de
  la #78).
- **La #10 y la #16 se hicieron con asistencia de IA.** Sus horas reales son el tiempo de
  Emiliano Sorato para dirigir y revisar, y por eso la desviación da muy negativa. Sin esas
  dos, la #11, la #17, la #18 y la #73 dan −21% (54 h estimadas contra 42,5 reales).
- **La #78 y la #177: 1,5 h reales cada una**, dadas por Emiliano Sorato el 28/10. También son
  su tiempo de dirección y revisión con asistencia de IA. La #78 sigue en progreso, con 35,1 h
  estimadas. La #177 se terminó sin story points: sus horas suman al total, pero no entran en la
  desviación ni en las horas por story point.
- **La #17: 25 h reales**, dadas por Juliana Bustos el 29/10. Incluyen también sus revisiones
  de PRs y los conflictos de integración del sprint.
- **Horas reales cargadas en el sprint:** 47. Son 42,5 de lo comprometido, 3 de la
  #73 y 1,5 de la #177.
- **Horas por integrante**, a nombre del responsable de cada historia, porque el tablero las
  guarda por historia: Juliana Bustos 25 (#17), Agustín Salinas 14,5 (#11 y #18), Emiliano
  Sorato 4,5 (#10, #16, #78 y #177) y Maximiliano Eula 3 (#73).
- **Horas por story point completado:** 1,3 (44 h / 33 SP).

## Defectos

| # | Defecto | Severidad | Detectado | Resuelto |
|---|---|---|---|---|
| 156 | Una ronda revelada sin participantes, cargada del archivo, hace caer `Differences()` | — | Sprint 3 (26/10) | Sprint 3 (26/10, PR #157) |
| 172 | Una ronda con un participante repetido con otras mayúsculas, o con espacios de más, queda trabada | — | Sprint 3 (26/10, Maximiliano Eula) | Sprint 3 (26/10, PR #173) |
| 190 | Con una página vieja, algunos errores no se ven o muestran otra página | Menor | Sprint 3 (26/10, prueba exploratoria de la web) | Sprint 3 (27/10, PR #194) |
| 191 | Una palabra muy larga ensancha la página en el celular | Menor | Sprint 3 (26/10, prueba exploratoria) | Sprint 3 (27/10, PR #195) |
| 192 | Las horas aceptan notación científica y no tienen máximo | Mayor | Sprint 3 (26/10, prueba exploratoria) | Sprint 3 (28/10, PR #196) |
| 193 | El filtro de estado en minúsculas no queda elegido en la lista | Menor | Sprint 3 (26/10, prueba exploratoria) | Sprint 3 (27/10, PR #197) |

- **Detectados en el sprint:** 6. **Resueltos:** 6. **Abiertos al cierre:** ninguno.
- La #190 a la #193 salieron de la prueba exploratoria de la web del 26/10. Lo que esa prueba
  dejó como decisión, y no como defecto, está en los *Límites conocidos* de la
  [guía de defensa](../guia-de-defensa.md).

## Calidad

Medida con las historias del sprint ya en `dev` (`5d1d556`, 29/10).

| Paquete | 26/10 (`4f7644b`) | 29/10 (`5d1d556`) |
|---|---|---|
| `internal/domain` | 100% | 100% |
| `internal/store` | 96,0% | 87,1% |
| `internal/web` | 94,0% | 94,1% |
| `cmd/app` | 96,2% | 96,2% |
| `cmd/web` | 94,0% | 91,1% |
| **Total** | **96,9%** | **95,7%** |
| Tests | 317 | 420 |

Medido con `go test -cover ./...`; los 420 tests pasan, con la #17 y la #18 ya en `dev`. Baja `internal/store`, porque la #177
cambió el guardado en archivo por la base de datos.

## Demostración

El 29/10 se ensayó el [guion de la demo](../demo/guion-de-demo.md) contra la web de `dev`, con
una base nueva y los datos de demo: todo funcionó y salieron tres ajustes al guion (PR #212).
El 30/10 se le sumó el dashboard de la #17 (PR #217).

## Pendientes

- [x] Cerrar los números con `scripts/metricas.sh "Sprint 3"` (30/10).
- [x] Las horas reales de la #78 y de la #177: 1,5 h cada una, dadas por Emiliano Sorato el 28/10
  y cargadas en el tablero.
- [x] Devolver al Product Backlog las historias no terminadas: la #78, que se termina en el
  Sprint 4 (30/10).
- [ ] El pull request de `dev` a `main` con el incremento del Sprint 3, el 28/10 a la noche, y su
  tag `v0.3.0-sprint3`. Cada sprint tiene su incremento: el del Sprint 4 es la entrega final,
  `v1.0.0`, el 29/10 (decisión del 27/10).
- [x] La [retrospectiva](sprint-3-retrospectiva.md) (falta la respuesta de Agustín Salinas).
