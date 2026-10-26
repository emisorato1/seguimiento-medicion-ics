# Acta — Review del Sprint 2

> **Armada el 24/10 con los datos del tablero**, calculados con `scripts/metricas.sh "Sprint 2"`,
> y la calidad medida sobre `dev`. **Completada el 26/10** con lo que escribió cada uno en la
> [issue #147](https://github.com/emisorato1/seguimiento-medicion-ics/issues/147). Lo que sigue
> entre corchetes, **[pendiente]**, nadie lo contestó: acá no se escribe nada que alguien no haya
> dicho. La retrospectiva va en [un archivo aparte](sprint-2-retrospectiva.md).

- **Fecha:** del 24 al 26/10/2026, por escrito en la issue #147, sin reunión
- **Participantes:** Emiliano Sorato, Maximiliano Eula, Juliana Bustos y Agustín Salinas
- **Facilita:** Emiliano Sorato (Agile Enabler)
- **Sprint Goal:** que el sistema maneje sprints, esfuerzo y defectos, y que todo se pueda operar
  desde la consola.
- **¿Se cumplió el Sprint Goal?** Las 9 historias comprometidas están terminadas. La parte de
  "operar desde la consola" la cuentan en la demo: según Juliana y Agustín, sus historias
  todavía no tienen pantalla ni opción en la consola.
  - **Emiliano Sorato (26/10): cumplido en parte.** El sistema maneja sprints, esfuerzo y
    defectos, pero se opera desde la web (#78) y no desde la consola, por la decisión del 20/10.
  - **Maximiliano Eula ([26/10, en el PR #169](https://github.com/emisorato1/seguimiento-medicion-ics/pull/169#issuecomment-6428626374)): cumplido en parte.** "Terminamos las
    9 historias del sprint, y el sistema ya maneja sprints, esfuerzo y defectos. Lo de operarlo
    todo desde la consola cambió en el camino: el 20/10 decidimos usar la web para lo nuevo, y
    esas pantallas llegaron con la #78, ya en el Sprint 3. La consola quedó con proyectos,
    integrantes, backlog y estimación."
  - **Juliana Bustos y Agustín Salinas:** [pendiente: en la #147 no lo dijeron].

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
- **Horas trabajadas en el Sprint 2:** 31. Es lo que da el script contando
  las adelantadas (53,25) menos las 22,25 h que la #6, la #12 y la #24 ya tenían cargadas al
  cierre del Sprint 1.

| Integrante | Horas reales en el Sprint 2 |
|---|---|
| Integrante | Historias | Horas en el tablero |
|---|---|---|
| Agustín Salinas | #7 5 h, #8 6 h, #9 9 h | 20 |
| Maximiliano Eula | #12 4 h y #24 7,25 h (menos las 9,25 h que ya contaban en el Sprint 1), #13 6 h | 8 |
| Emiliano Sorato | #5, #14 y #15 0,5 h cada una, #77 1,5 h (las 13 h de la #6 son del Sprint 1) | 3 |
| Juliana Bustos | sus historias (#5, #14 y #15) las implementó Emiliano | 0 |
| **Total** | | **31** |

Cada uno confirmó en la #147 que sus horas están cargadas. El tablero guarda las horas por
historia; el reparto por persona sale de quién trabajó cada historia.

**Las horas propias de Juliana Bustos no están en el tablero.** [Escribió](https://github.com/emisorato1/seguimiento-medicion-ics/issues/147#issuecomment-6400529426): "Mis
horas propias del Sprint 2, para la tabla por integrante del acta: 14 h. Spec de la #17 (Fases
1, 2 y 3, y respuestas a las observaciones de Agustín, Maxi y Emiliano): 6 h. Revisiones de PRs
y de puertas (#88, #136, #141, #142 y #143): 6 h. Tablero, criterios y actas: 2 h". No suman a
las 31 h ni a la desviación, porque no son de una historia del Sprint 2: la spec de la #17 es de
una historia del Sprint 3, y el resto es revisión y gestión. Contando las 14 h, el equipo
trabajó **45 h** en el sprint.

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
| #13 (adelantada) | 21,6 | 6 | −15,6 h |
| **Total** | **108,3** | **51,75** | **−56,55 h (−52%)** |

La #7 tiene ahora 5 h reales en el tablero; el 21/10 tenía 4. Las horas de la #8 y la #9 son
de Agustín Salinas.

**La #5, la #14 y la #15 se implementaron con asistencia de IA**, que la consigna permite. Sus
horas reales son el tiempo humano de dirección y revisión (0,5 h cada una, de Emiliano Sorato),
y por eso la desviación da muy negativa. Sin esas tres, las otras siete dan −36% (78,6 h
estimadas contra 50,25 reales).

La #13 tiene 6 h reales contra 21,6 estimadas (−72%). Maximiliano Eula lo explicó el 24/10 en
la [issue #147](https://github.com/emisorato1/seguimiento-medicion-ics/issues/147#issuecomment-6391163190): "faltó cargar horas, ya las corregí (eran 3); aun así se
avanzó más rápido por el tiempo que se acortó para la entrega del proyecto". Las 21,6 salen
de la conversión de 2,7 h por punto.

La #77 también se hizo con asistencia de IA (1,5 h de Emiliano Sorato), pero no entra en esta
tabla porque no tiene horas estimadas con qué compararla. Sus horas sí cuentan en las
trabajadas del sprint.

- **Horas por story point completado:** 1,3 (51,75 h ÷ 39 SP). Sin la #5, la #14 y la #15 da
  1,8 (50,25 h ÷ 28 SP).

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

No hubo reunión: cada uno probó lo suyo sobre `dev` y lo contó en la #147.

- **[Emiliano Sorato](https://github.com/emisorato1/seguimiento-medicion-ics/issues/147#issuecomment-6390684355)** (sobre `1537556`): "compila todo, la web abre en la página
  de Proyectos y la consola muestra su menú. Lo mío del Sprint 2 (la web base #77, la #5 y la #14
  y la #15 que tomé de Juli) anda".
- **[Maximiliano Eula](https://github.com/emisorato1/seguimiento-medicion-ics/issues/147#issuecomment-6391163190)** (sobre `1537556`): "En la consola (#24) creé un proyecto,
  agregué un integrante, cargué una historia y la estimé en 8 (#12); un 4 se rechaza con el
  mensaje de la escala y el backlog la muestra con 8 story points. En la web la estimación
  también anda y queda guardada en el archivo. La #13 todavía no tiene pantalla (la conecta la
  #78): se prueba con sus tests, que pasan".
- **[Juliana Bustos](https://github.com/emisorato1/seguimiento-medicion-ics/issues/147#issuecomment-6400529426)** (sobre `1537556`): "Mis historias del Sprint 2 (#5, #14 y
  #15, que implementó Emiliano) todavía no tienen pantalla ni operación en la consola: la consola
  ofrece proyectos, integrantes, historias, backlog y estimación, y la web llega hasta el
  backlog. Las conecta la #78. Las probé con sus tests, que pasan todos: #5 10 de 10, #14 19 de
  19 y #15 22 de 22".
- **[Agustín Salinas](https://github.com/emisorato1/seguimiento-medicion-ics/issues/147#issuecomment-6406773200)** (sobre `cdda205`): "compila todo y los tests pasan. Mis
  historias del Sprint 2 (#7, #8 y #9) todavía no tienen pantalla ni opción en la consola: las
  conecta la #78. Las probé con sus tests, que pasan todos: #7 16 de 16, #8 17 de 17 y #9 14 de
  14".

Los Product Architects no participaron de esta review por escrito.

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
- [x] Avisar al grupo la decisión del punto de control: está en la [issue #147](https://github.com/emisorato1/seguimiento-medicion-ics/issues/147).
- [ ] Revisar los números con `scripts/metricas.sh "Sprint 2"` el 25/10, antes de cerrar.
- [x] Devolver al Product Backlog las historias no terminadas: no quedó ninguna.
- [ ] Abrir el pull request de `dev` a `main` con el incremento y etiquetarlo `v0.2.0-sprint2`.
- [ ] Hacer la [retrospectiva](sprint-2-retrospectiva.md).
