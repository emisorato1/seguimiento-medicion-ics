# Acta — Review del Sprint 2

> **Armada el 24/10 con los datos del tablero**, calculados con `scripts/metricas.sh "Sprint 2"`,
> y la calidad medida sobre `dev`. El domingo 25 se revisan otra vez antes de cerrar. Lo que
> está entre corchetes, **[pendiente]**, lo completa el equipo en la review: acá no se escribe
> nada que alguien no haya dicho. La retrospectiva va en [un archivo aparte](sprint-2-retrospectiva.md).

- **Fecha:** [pendiente — el sprint cierra el domingo 25/10]
- **Participantes:** [pendiente]
- **Facilita:** Emiliano Sorato (Agile Enabler)
- **Sprint Goal:** que el sistema maneje sprints, esfuerzo y defectos, y que todo se pueda operar
  desde la consola.
- **¿Se cumplió el Sprint Goal?** [pendiente: lo dice el equipo en la review]

## Qué se completó

_Al 24/10._

Una historia cuenta como completada solo si cumple la Definición de Terminado: spec con sus
tres puertas, escenarios BDD, tests en Go que pasan, revisada por otro integrante y mergeada
a `dev`. Cuenta en el sprint en que se mergea, y lo que se mergea entre un sprint y el
siguiente cuenta en el siguiente.

| # | Historia | SP | Estado al 24/10 | Responsable | Evidencia |
|---|---|---|---|---|---|
| 6 | Administrar el Product Backlog | 5 | **Hecho** | Emiliano Sorato | PR #58, mergeado el 19/10, entre el Sprint 1 y el 2 |
| 12 | Estimar con Story Points | 2 | **Hecho** | Maximiliano Eula | PR #70 |
| 24 | Interfaz de consola | 3 | **Hecho** | Maximiliano Eula | PR #106 |
| 7 | Priorizar el Product Backlog | 2 | **Hecho** | Agustín Salinas | PR #105 |
| 8 | Crear un sprint con su Sprint Goal | 3 | **Hecho** | Agustín Salinas | PR #138, después de las tareas 1, 2, 3 y 16 en el PR #104 |
| 9 | Asignar historias a un sprint | 5 | **Hecho** | Agustín Salinas | PR #138 |
| 5 | Consultar el estado de un proyecto | 3 | **Hecho** | Juliana Bustos (la implementó Emiliano con asistencia de IA) | PR #102 |
| 14 | Registrar el esfuerzo real | 3 | **Hecho** | Juliana Bustos (la implementó Emiliano con asistencia de IA) | PR #100 |
| 15 | Gestionar defectos | 5 | **Hecho** | Juliana Bustos (la implementó Emiliano con asistencia de IA) | PR #129 |
| 77 | Interfaz web: proyectos, integrantes y backlog | sin estimar | **Hecho** | Emiliano Sorato | PR #95 |

- **Story points comprometidos:** 31 (los 26 del plan más los 5 de la #6)
- **Story points completados:** **31 de 31**
- **Velocidad del Sprint 2:** **39 = 31 + 8 adelantados** (la #13, ver abajo).
- **Historias completadas:** 9 de 9 (100%), sin contar la #77, que también está terminada
  (PR #95) pero no está estimada

La #77 no suma story points ni horas mientras el equipo no la estime (ver el
[anexo de la planning](sprint-2-planning.md#historias-nuevas-la-interfaz-web)).

Las historias no terminadas al cierre **vuelven al Product Backlog**, no se arrastran en
silencio.

### Historias adelantadas

La decisión 11 del [plan](../plan-final-al-2-de-noviembre.md) dice que quien termina lo suyo
puede adelantar su próxima historia del plan sin esperar la planning. Se agregó el 21/10 y el
equipo la confirmó por chat. La historia tiene que tener las tres puertas aprobadas y lo que
necesita ya en `dev`. Entra al sprint en curso, queda registrada en esta acta y, si no se
termina, vuelve al sprint que tenía en el plan. **No suma a los 31 story points
comprometidos**: se cuenta aparte.

| # | Historia | SP | Entró | Estado al 24/10 | Responsable | Evidencia |
|---|---|---|---|---|---|---|
| 13 | Votar en secreto en una ronda de Planning Poker | 8 | 21/10 | **Hecho** | Maximiliano Eula | PR #141, mergeado el 23/10; la issue la cerró Juliana Bustos |

- La #13 cumple las dos condiciones. Agustín Salinas aprobó sus tres puertas en el PR #115,
  y la #4, la #6 y la #12, de las que depende, ya están en `dev`.
- **Horas estimadas de la #13: 21,6.** Se cargaron el 21/10, antes de empezarla, con la
  conversión de 2,7 horas por punto, como las demás del sprint (ver el
  [anexo de la planning](sprint-2-planning.md#horas-estimadas-del-sprint-2)). Maximiliano Eula
  las puede ajustar al empezarla. Sus horas reales suman a las trabajadas en el sprint:
  `scripts/metricas.sh` las muestra en la fila *Horas reales, contando las adelantadas*.
- **Se terminó en el Sprint 2**, así que suma a los story points completados y a la velocidad
  del sprint, mostrada aparte (31 + 8). El porcentaje de historias completadas se sigue
  calculando sobre las 9 comprometidas, y el Sprint 3 se planifica sin esos 8 puntos.
- **La #16 no es una adelantada: sigue en el Sprint 3.** Depende de la #8, la #9, la #10 y la
  #11, que no están en `dev`, así que no cumple la decisión 11. Su parte 1 (PR #81, las horas
  estimadas de una historia) se viene adelantando desde el 20/10, antes de la regla. Se
  programa contra el contrato del dominio, como permite el *Modo sprint final* de `CLAUDE.md`.

## Esfuerzo

_Al 24/10._

Las horas se cuentan como dice el [informe de métricas](../informe-de-metricas.md#historias-que-pasan-de-un-sprint-a-otro):
cada hora una sola vez, en el sprint donde se trabajó.

- **Horas estimadas de las historias nuevas:** 56,7 (a 2,7 h por story point). La #6, la #12
  y la #24 traen sus 30 h estimadas del Sprint 1, donde ya se contaron.
- **Horas trabajadas en el Sprint 2:** 28. Es lo que da el script contando
  las adelantadas (50,25) menos las 22,25 h que la #6, la #12 y la #24 ya tenían cargadas al
  cierre del Sprint 1.

| Integrante | Horas reales en el Sprint 2 |
|---|---|
| Agustín Salinas | [pendiente] |
| Juliana Bustos | [pendiente] |
| Maximiliano Eula | [pendiente] |
| Emiliano Sorato | [pendiente] |
| **Total** | **[pendiente]** |

El tablero guarda las horas por historia y no por persona, así que este reparto lo completa
cada uno en la review.

**Historias terminadas**, medidas enteras —todas sus horas, de los dos sprints— en el sprint
donde se terminaron:

| Historia | Estimadas | Reales | Desviación |
|---|---|---|---|
| #6 | 15 | 13 | −2 h |
| #12 | 6 | 4 | −2 h |
| #24 | 9 | 7,25 | −1,75 h |
| #7 | 5,4 | 5 | −0,4 h |
| #8 | 8,1 | 6 | −2,1 h |
| #9 | 13,5 | 9 | −4,5 h |
| #5 | 8,1 | 0,5 | −7,6 h |
| #14 | 8,1 | 0,5 | −7,6 h |
| #15 | 13,5 | 0,5 | −13 h |
| #13 (adelantada) | 21,6 | 3 | −18,6 h |
| **Total** | **108,3** | **48,75** | **−59,55 h (−55%)** |

La #7 tiene ahora 5 h reales en el tablero; el 21/10 tenía 4. Las horas de la #8 y la #9 son
de Agustín Salinas.

**La #5, la #14 y la #15 se implementaron con asistencia de IA**, que la consigna permite. Sus
horas reales son el tiempo humano de dirección y revisión (0,5 h cada una, de Emiliano Sorato),
y por eso la desviación da muy negativa. Sin esas tres, las otras siete dan −40% (78,6 h
estimadas contra 47,25 reales).

La #13 tiene 3 h reales contra 21,6 estimadas (−86%). Por qué: [pendiente, lo explica
Maximiliano Eula en la review].

La #77 también se hizo con asistencia de IA (1,5 h de Emiliano Sorato), pero no entra en esta
tabla porque no tiene horas estimadas con qué compararla. Sus horas sí cuentan en las
trabajadas del sprint.

- **Horas por story point completado:** 1,3 (48,75 h ÷ 39 SP). Sin la #5, la #14 y la #15 da
  1,7 (47,25 h ÷ 28 SP).

## Defectos

_Al 24/10: no cambió desde el 21/10._

| Detectados en el sprint | Resueltos en el sprint | Abiertos al cierre |
|---|---|---|
| 1 | 2 | 0 |

| # | Defecto | Historia | Detectado | Resuelto |
|---|---|---|---|---|
| 60 | Los mensajes de error del dominio no llevan tildes | #3 y #6 | Sprint 1 (12/10) | Sprint 2 (20/10) |
| 108 | Tres tests de `internal/store` fallan en Windows | #77 | Sprint 2 (20/10, Agustín Salinas) | Sprint 2 (21/10, PR #116) |

## Calidad

Medido sobre `dev` en `b45cbe6` (24/10), con `go test -cover ./...`. Todos los tests pasan.

| Paquete | Cobertura |
|---|---|
| `internal/domain` | 100% |
| `internal/store` | 96,0% |
| `internal/web` | 93,8% |
| `cmd/app` (consola) | 96,2% |
| `cmd/web` (arranque de la web) | 56,2% |
| **Total** | **96,0%** |

- **263 tests** (589 casos, contando cada fila de los tests con tabla). Al 21/10 eran 121.
- El total baja de 98,8% a 96,0% porque entró la web. En `cmd/web` lo que no tiene test es
  `main` y `lanAddresses` (las direcciones de la red local que muestra al arrancar); abrir el
  puerto y el navegador sí están probados.
- Desde el arreglo del defecto #108, el CI corre los tests también en Windows. Ese chequeo no
  bloquea el merge: el obligatorio sigue siendo el de Linux.
- La [trazabilidad](../trazabilidad.md) de cada historia terminada se genera con
  `scripts/trazabilidad.sh`.

## Demostración

[Qué se mostró y qué dijeron los Product Architects. Si hubo pedidos de cambio, anotarlos: van
al Product Backlog como historias nuevas, no se arreglan sobre la marcha.]

## Punto de control del 25/10

El [plan](../plan-final-al-2-de-noviembre.md#punto-de-control-domingo-25-a-la-noche) pide que el
domingo 25 el Agile Enabler mire el tablero y decida con estos umbrales:

| Si el Sprint 2 cerró… | Qué se hace |
|---|---|
| sus 26 puntos | el plan sigue tal cual |
| 20 o más | se sostiene el alcance y la #19 (PDF) se recorta o queda en el backlog |
| menos de 20 | salen la #19 y la #73, y la #17 queda en consola sin gráficos |

**Los números al 24/10:**

- **Completados de lo comprometido: 31 SP**, las 9 historias. Los 26 del plan más los 5 de la
  #6, que pasó del Sprint 1. Supera el primer umbral.
- **Más 8 SP adelantados** (la #13), que salen del Sprint 3. Velocidad del sprint: 39.
- **El Sprint 3 queda en 28 SP** (36 del plan menos los 8 de la #13): la #10, la #11, la #16,
  la #17 y la #18, más la #78 sin estimar. La #16 tiene dos partes en `dev` (12 de 16
  escenarios con test) y la #17 está en progreso.
- **El Sprint 4** tiene la #19 y la #73 (10 SP). La #73 ya tiene sus puertas aprobadas.
- **Defectos abiertos: 0.**

**Decisión de Emiliano Sorato, del 24/10: el plan sigue tal cual**, como dice el plan para un
Sprint 2 que cerró 26 puntos o más. No se recorta nada.

## Pendientes

- [x] Cargar las horas reales de la #15: 0,5 h, el 24/10.
- [ ] Avisar al grupo la decisión del punto de control.
- [ ] Revisar los números con `scripts/metricas.sh "Sprint 2"` el 25/10, antes de cerrar.
- [x] Devolver al Product Backlog las historias no terminadas: no quedó ninguna.
- [ ] Abrir el pull request de `dev` a `main` con el incremento y etiquetarlo `v0.2.0-sprint2`.
- [ ] Hacer la [retrospectiva](sprint-2-retrospectiva.md).
