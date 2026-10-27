# Acta — Review del Sprint 3

> **Borrador armado el 27/10 con los datos del tablero.** Se cierra el miércoles 28/10 a la
> noche, con `scripts/metricas.sh "Sprint 3"` y con lo que diga cada uno. Lo que está entre
> corchetes, **[pendiente]**, todavía no se sabe o nadie lo dijo: acá no se escribe nada que
> alguien no haya dicho. La retrospectiva va en [un archivo aparte](sprint-3-retrospectiva.md).

- **Fecha:** 28/10/2026 [pendiente: cómo se hizo, por escrito o con reunión]
- **Participantes:** [pendiente]
- **Facilita:** Emiliano Sorato (Agile Enabler)
- **Sprint:** del lunes 26 al miércoles 28/10. Al principio iba hasta el viernes 30/10, y el
  26/10 se acortó a pedido del equipo, para terminar todo el 29/10.
- **Sprint Goal:** que el sistema mida el proyecto y lo muestre: métricas, dashboard y reporte.
- **¿Se cumplió el Sprint Goal?** [pendiente]

## Qué se completó

_Al 27/10, según el tablero._

Una historia cuenta como completada solo si cumple la Definición de Terminado: spec con sus
tres puertas, escenarios BDD, tests en Go que pasan, revisada por otro integrante y mergeada
a `dev`.

| # | Historia | SP | Estado al 27/10 | Responsable | Evidencia |
|---|---|---|---|---|---|
| 10 | Cerrar un sprint | 5 | **Hecho** | Emiliano Sorato (con asistencia de IA) | PR #150, 26/10 |
| 16 | Calcular las métricas | 8 | **Hecho** | Emiliano Sorato (con asistencia de IA) | PRs #81, #140 y #151 |
| 11 | Consultar sprints anteriores | 2 | **Hecho** | Agustín Salinas | PR #143, 26/10 |
| 17 | Mostrar el dashboard | 8 | En progreso | Juliana Bustos | rama `historia-17-dashboard`, sin PR al 27/10 |
| 18 | Generar el reporte | 5 | Especificación SDD | Agustín Salinas | spec en el PR #181, con las puertas pedidas |
| 78 | Operar todo desde la web | 13 | En progreso | Emiliano Sorato (con asistencia de IA) | PRs #149, #163, #165 y #170 en `dev`; 17 de 22 escenarios con test |

- **Story points comprometidos:** 41. Son los 28 de la planning más los 13 de la #78, que se
  estimó en la misma planning ([planning del Sprint 3](sprint-3-planning.md)).
- **Story points completados:** [pendiente: al cierre]. Al 27/10 son 15 (37%).
- **Velocidad del Sprint 3:** [pendiente: al cierre]. Al 27/10 son 20 (15 + 5 de la #73
  adelantada).
- **Historias completadas:** [pendiente: al cierre]. Al 27/10 son 3 de 6 (50%).

Las historias no terminadas al cierre **vuelven al Product Backlog**, no se arrastran en
silencio.

### Historias adelantadas

| # | Historia | SP | Estado | Responsable | Evidencia |
|---|---|---|---|---|---|
| 73 | Resolver diferencias y acordar la estimación | 5 | **Hecho** | Maximiliano Eula | PR #148, con los arreglos de los defectos #156 (PR #157) y #172 (PR #173) |
| 19 | Exportar el reporte a PDF | 5 | [pendiente: se adelanta desde el Sprint 4 si la #18 entra a `dev` antes del cierre] | Emiliano Sorato (con asistencia de IA) | PR #184 |

### Historias agregadas durante el sprint

| # | Historia | SP | Estado al 27/10 | Responsable | Evidencia |
|---|---|---|---|---|---|
| 177 | Guardar los datos en una base de datos (SQLite) | [pendiente: Planning Poker] | En progreso | Emiliano Sorato (con asistencia de IA) | spec en el PR #180 e implementación en el PR #186 |

## Esfuerzo

_Al 27/10, según el tablero._

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
- **Horas por integrante:** [pendiente: cada uno las completa en la review].
- **Horas por story point completado:** [pendiente: al cierre].

## Defectos

| # | Defecto | Severidad | Detectado | Resuelto |
|---|---|---|---|---|
| 156 | Una ronda revelada sin participantes, cargada del archivo, hace caer `Differences()` | — | Sprint 3 (26/10) | Sprint 3 (26/10, PR #157) |
| 172 | Una ronda con un participante repetido con otras mayúsculas, o con espacios de más, queda trabada | — | Sprint 3 (26/10, Maximiliano Eula) | Sprint 3 (26/10, PR #173) |
| 190 | Con una página vieja, algunos errores no se ven o muestran otra página | Menor | Sprint 3 (26/10, prueba exploratoria de la web) | [pendiente: PR #194] |
| 191 | Una palabra muy larga ensancha la página en el celular | Menor | Sprint 3 (26/10, prueba exploratoria) | [pendiente: PR #195] |
| 192 | Las horas aceptan notación científica y no tienen máximo | Mayor | Sprint 3 (26/10, prueba exploratoria) | [pendiente: PR #196] |
| 193 | El filtro de estado en minúsculas no queda elegido en la lista | Menor | Sprint 3 (26/10, prueba exploratoria) | [pendiente: PR #197] |

- **Detectados en el sprint:** 6. **Resueltos:** [pendiente: al cierre]; al 27/10, 2.
- La #190 a la #193 salieron de la prueba exploratoria de la web del 26/10. Lo que esa prueba
  dejó como decisión, y no como defecto, está en los *Límites conocidos* de la
  [guía de defensa](../guia-de-defensa.md).

## Calidad

[pendiente: medir sobre `dev` en el último cambio de código del sprint, con `go test ./...
-coverprofile`.] Al 26/10, sobre `4f7644b`: dominio 100%, `internal/store` 96,0%,
`internal/web` 94,0%, `cmd/app` 96,2%, `cmd/web` 94,0% y **96,9% en total**, con 317 tests.

## Demostración

[pendiente: lo que cuenta cada uno de su demo.]

## Pendientes

- [ ] Cerrar los números con `scripts/metricas.sh "Sprint 3"` el 28/10.
- [ ] Las horas reales de la #78 y de la #177 (Emiliano Sorato, por el manager).
- [ ] Devolver al Product Backlog las historias no terminadas.
- [ ] [pendiente: si el Sprint 3 tiene su propio pull request de `dev` a `main` o si va todo junto
  el 29/10, como dice el plan para el Sprint 4.]
- [ ] La [retrospectiva](sprint-3-retrospectiva.md).
