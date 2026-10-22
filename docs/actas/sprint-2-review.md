# Acta — Review del Sprint 2

> **Esqueleto armado el 21/10.** Los números están calculados con `scripts/metricas.sh "Sprint 2"`
> **al 21/10, y se actualizan al cierre** del domingo 25. Lo que está entre corchetes lo completa
> el equipo en la review. La retrospectiva va en [un archivo aparte](sprint-2-retrospectiva.md).

- **Fecha:** [completar — el sprint cierra el domingo 25/10]
- **Participantes:** [quiénes estuvieron]
- **Facilita:** Emiliano Sorato (Agile Enabler)
- **Sprint Goal:** que el sistema maneje sprints, esfuerzo y defectos, y que todo se pueda operar
  desde la consola.
- **¿Se cumplió el Sprint Goal?** [sí / parcialmente / no, y por qué]

## Qué se completó

_Al 21/10, se actualiza al cierre._

Una historia cuenta como completada solo si cumple la Definición de Terminado: spec con sus
tres puertas, escenarios BDD, tests en Go que pasan, revisada por otro integrante y mergeada
a `dev`. Cuenta en el sprint en que se mergea, y lo que se mergea entre un sprint y el
siguiente cuenta en el siguiente.

| # | Historia | SP | Estado al 21/10 | Responsable | Evidencia |
|---|---|---|---|---|---|
| 6 | Administrar el Product Backlog | 5 | **Hecho** | Emiliano Sorato | PR #58, mergeado el 19/10, entre el Sprint 1 y el 2 |
| 12 | Estimar con Story Points | 2 | **Hecho** | Maximiliano Eula | PR #70 |
| 24 | Interfaz de consola | 3 | **Hecho** | Maximiliano Eula | PR #106 |
| 7 | Priorizar el Product Backlog | 2 | **Hecho** | Agustín Salinas | PR #105 |
| 8 | Crear un sprint con su Sprint Goal | 3 | En progreso | Agustín Salinas | el sprint ya está en `dev` desde el PR #104, con las tareas 1, 2, 3 y 16 (`Refs #8`); [completar al cierre] |
| 9 | Asignar historias a un sprint | 5 | Lista para implementar | Agustín Salinas | [completar al cierre] |
| 5 | Consultar el estado de un proyecto | 3 | **Hecho** | Juliana Bustos (la implementó Emiliano con asistencia de IA) | PR #102 |
| 14 | Registrar el esfuerzo real | 3 | **Hecho** | Juliana Bustos (la implementó Emiliano con asistencia de IA) | PR #100 |
| 15 | Gestionar defectos | 5 | Lista para implementar | Juliana Bustos (la implementa Emiliano con asistencia de IA) | spec en `dev` desde el PR #101; [completar al cierre] |
| 77 | Interfaz web: proyectos, integrantes y backlog | sin estimar | **Hecho** | Emiliano Sorato | PR #95 |

- **Story points comprometidos:** 31 (los 26 del plan más los 5 de la #6)
- **Story points completados al 21/10:** 18
- **Velocidad del Sprint 2:** [se fija al cierre]
- **Historias completadas al 21/10:** 6 de 9 (67%), sin contar la #77, que también está terminada
  (PR #95) pero no está estimada

La #77 no suma story points ni horas mientras el equipo no la estime (ver el
[anexo de la planning](sprint-2-planning.md#historias-nuevas-la-interfaz-web)).

Las historias no terminadas al cierre **vuelven al Product Backlog**, no se arrastran en
silencio.

## Esfuerzo

_Al 21/10, se actualiza al cierre._

Las horas se cuentan como dice el [informe de métricas](../informe-de-metricas.md#historias-que-pasan-de-un-sprint-a-otro):
cada hora una sola vez, en el sprint donde se trabajó.

- **Horas estimadas de las historias nuevas:** 56,7 (a 2,7 h por story point). La #6, la #12
  y la #24 traen sus 30 h estimadas del Sprint 1, donde ya se contaron.
- **Horas trabajadas en el Sprint 2:** 15,5. Es lo que da el script (37,75) menos las 22,25 h que
  la #6, la #12 y la #24 ya tenían cargadas al cierre del Sprint 1.

| Integrante | Horas reales en el Sprint 2 |
|---|---|
| Agustín Salinas | [completar al cierre] |
| Juliana Bustos | [completar al cierre] |
| Maximiliano Eula | [completar al cierre] |
| Emiliano Sorato | [completar al cierre] |
| **Total** | **[completar al cierre]** |

**Historias terminadas**, medidas enteras —todas sus horas, de los dos sprints— en el sprint
donde se terminaron:

| Historia | Estimadas | Reales | Desviación |
|---|---|---|---|
| #6 | 15 | 13 | −2 h |
| #12 | 6 | 4 | −2 h |
| #24 | 9 | 7,25 | −1,75 h |
| #7 | 5,4 | 4 | −1,4 h |
| #5 | 8,1 | 0,5 | −7,6 h |
| #14 | 8,1 | 0,5 | −7,6 h |
| **Total** | **51,6** | **29,25** | **−22,35 h (−43%)** |

**La #5 y la #14 se implementaron con asistencia de IA**, que la consigna permite. Sus horas
reales son el tiempo humano de dirección y revisión (0,5 h cada una, de Emiliano Sorato), y por
eso la desviación da muy negativa. Sin esas dos, las otras cuatro dan −20% (35,4 h estimadas
contra 28,25 reales).

La #77 también se hizo con asistencia de IA (1,5 h de Emiliano Sorato), pero no entra en esta
tabla porque no tiene horas estimadas con qué compararla. Sus horas sí cuentan en las
trabajadas del sprint.

- **Horas por story point completado al 21/10:** 1,6 (29,25 h ÷ 18 SP). Sin la #5 y la #14
  da 2,4 (28,25 h ÷ 12 SP).

## Defectos

_Al 21/10, se actualiza al cierre._

| Detectados en el sprint | Resueltos en el sprint | Abiertos al cierre |
|---|---|---|
| 1 | 2 | 0 |

| # | Defecto | Historia | Detectado | Resuelto |
|---|---|---|---|---|
| 60 | Los mensajes de error del dominio no llevan tildes | #3 y #6 | Sprint 1 (12/10) | Sprint 2 (20/10) |
| 108 | Tres tests de `internal/store` fallan en Windows | #77 | Sprint 2 (20/10, Agustín Salinas) | Sprint 2 (21/10, PR #116) |

## Calidad

_Al 21/10, se actualiza al cierre._

Medido sobre `dev` en `6ad7ac7` (21/10):

| Paquete | Cobertura |
|---|---|
| `internal/domain` | 100% |
| `internal/store` | 96,0% |
| `cmd/app` (consola) | 96,2% |
| **Total** | **98,8%** |

- **121 tests** (234 casos, contando cada fila de los tests con tabla).
- Desde el arreglo del defecto #108, el CI corre los tests también en Windows. Ese chequeo no
  bloquea el merge: el obligatorio sigue siendo el de Linux.
- La [trazabilidad](../trazabilidad.md) de cada historia terminada se genera con
  `scripts/trazabilidad.sh`.

## Demostración

[Qué se mostró y qué dijeron los Product Architects. Si hubo pedidos de cambio, anotarlos: van
al Product Backlog como historias nuevas, no se arreglan sobre la marcha.]

## Pendientes

- [ ] Actualizar los números de esta acta con `scripts/metricas.sh "Sprint 2"` el 25/10.
- [ ] Devolver al Product Backlog las historias no terminadas.
- [ ] Abrir el pull request de `dev` a `main` con el incremento y etiquetarlo `v0.2.0-sprint2`.
- [ ] Hacer la [retrospectiva](sprint-2-retrospectiva.md).
