# Informe de métricas y cobertura

Proyecto integrador de **Ingeniería y Calidad de Software** — UTN FRSR, 2026.
Responsable: Emiliano Sorato (Agile Enabler).

> **Borrador en construcción.** Tiene cargado el Sprint 1. Los Sprints 2, 3 y 4 se completan
> al cierre de cada uno, con su review. La versión final se entrega el 1/11.

## De qué se trata

La aplicación que construimos mide proyectos de software. Este informe aplica esas mismas
métricas a **nuestro propio proyecto**: cuánto planificamos, cuánto terminamos, cuánto
tardamos contra lo que estimamos y cuántos defectos aparecieron. Al final, la cobertura de
los tests en cada cierre de sprint.

### De dónde salen los datos

| Dato | Fuente |
|---|---|
| Story points, estado, sprint y horas de cada historia | el tablero del proyecto (https://github.com/users/emisorato1/projects/3) |
| Defectos | las issues con la etiqueta `defecto` |
| Compromiso de cada sprint | el acta de planning (`docs/actas/sprint-N-planning.md`) |
| Lo que se terminó | el acta de review (`docs/actas/sprint-N-review.md`) |
| Cobertura | `go test ./... -coverprofile`, el mismo comando que corre el CI, sobre `dev` en el último cambio de código del sprint |
| Evidencia de TDD | el historial de commits de `dev` |

Los números del tablero se sacan con `scripts/metricas.sh "Sprint N"`.

### Cómo se calcula cada métrica

Son las nueve que pide la consigna.

| Métrica | Cómo se calcula |
|---|---|
| Story points planificados | la suma de los puntos de las historias comprometidas en la planning |
| Story points completados | la suma de los puntos de las historias que cumplen la Definición de Terminado al cierre |
| Velocidad | los story points completados del sprint |
| Horas estimadas | las que se cargaron en el tablero **antes** de empezar cada historia |
| Horas reales | las que cargó cada integrante mientras trabajaba |
| Desviación de esfuerzo | (reales − estimadas) ÷ estimadas, en porcentaje. Negativa quiere decir que se estimó de más |
| Porcentaje de historias completadas | historias terminadas ÷ historias comprometidas |
| Defectos detectados | issues `defecto` abiertas durante el sprint |
| Defectos resueltos | issues `defecto` cerradas durante el sprint |

Una historia a medias **no suma puntos**, aunque tenga horas cargadas. Si sumara, la
velocidad quedaría inflada y el sprint siguiente se planificaría mal.

**Una historia cuenta en el sprint en que se mergea a `dev`**, porque la Definición de
Terminado exige el merge, y un defecto cuenta en el sprint en que se cierra su issue. **Lo que
se mergea o se cierra entre un sprint y el siguiente cuenta en el siguiente.** Una historia
aprobada pero sin mergear al cierre no está terminada: es lo que pasó con la #6, que se mergeó
el 19/10, entre el cierre del Sprint 1 (16/10) y el arranque del Sprint 2 (20/10), y por eso
cuenta en el Sprint 2.

**Una historia adelantada** (decisión 11 del plan, desde el 21/10) entra al sprint en curso
sin sumar a lo comprometido. Si se termina, suma a los story points completados y a la
velocidad de ese sprint, mostrada aparte, y el sprint siguiente se planifica sin ella. El
porcentaje de historias completadas se calcula solo sobre las comprometidas. Si no se
termina, vuelve al sprint que tenía en el plan.

### Historias que pasan de un sprint a otro

Una historia que no se termina vuelve al backlog y sigue en el sprint siguiente. El tablero
guarda un solo total de horas por historia, así que hay que decidir dónde cuenta cada hora
para no sumarla dos veces:

- **Las horas reales cuentan en el sprint donde se trabajaron.** Lo que la historia tenía
  cargado al cierre queda en ese sprint, y lo fija el acta de review. En el sprint siguiente
  solo cuenta lo que se cargue después. Como `scripts/metricas.sh` suma el total de la
  historia, a su resultado se le resta lo que ya quedó en el sprint anterior.
- **Las horas estimadas cuentan una sola vez**, en el sprint donde la historia se comprometió
  por primera vez.
- **La desviación de una historia terminada se mide con la historia entera** —todas sus
  horas, de todos los sprints— en el sprint donde se termina. Es la que sirve para estimar, y
  así cada historia entra una sola vez y completa.

**El caso de la #12, la #24 y la #6.** Al cierre del Sprint 1 tenían cargadas 4 h, 5,25 h y
13 h reales: son 22,25 h que ya están en el Sprint 1. En el Sprint 2 se restan del total que
da el script. Sus horas estimadas (6, 9 y 15) tampoco se vuelven a sumar, porque ya están en
las 48 del Sprint 1: el acta de planning del Sprint 2 lista 86,7 h porque las incluye, y en
este informe el Sprint 2 cuenta 56,7 h estimadas de historias nuevas.

Cuando la #12 y la #6 se miden como historias terminadas en el Sprint 2, entran enteras: la
#12 con 6 h estimadas contra 4 reales, y la #6 con 15 contra 13. Esas horas aparecen en dos
métricas distintas —las horas que puso el equipo en el Sprint 1 y lo que costó cada
historia—, pero ninguna de las dos las suma dos veces.

**La aplicación no puede hacer este reparto, y el informe sí.** En la aplicación el registro
de esfuerzo no tiene sprint, así que la #16 (RN-7) pone todas las horas de una historia en el
último sprint donde estuvo. El informe puede repartirlas porque el acta de review dejó
escritas las horas que tenía cada historia al cierre.

## Resumen por sprint

| Métrica | Sprint 1 | Sprint 2 | Sprint 3 | Sprint 4 |
|---|---|---|---|---|
| Fechas | 28/09 → 16/10 | 20/10 → 25/10 | 26/10 → 30/10 | 31/10 → 01/11 |
| SP planificados | 16 | 31 (44 con la #77) ‡ | _pendiente_ | _pendiente_ |
| SP completados (velocidad) | **6** | 44 + 8 adelantados = **52** ‡ | _pendiente_ | _pendiente_ |
| Historias completadas | 2 de 5 (40%) | 10 de 10 (100%) ‡ | _pendiente_ | _pendiente_ |
| Horas estimadas | 48 | 56,7 | _pendiente_ | _pendiente_ |
| Horas reales | 38,25 | 31 * | _pendiente_ | _pendiente_ |
| Desviación total | −20% | al cierre | _pendiente_ | _pendiente_ |
| Desviación de las historias terminadas | −11% | −52% * † | _pendiente_ | _pendiente_ |
| Horas por story point completado | 2,7 | 1,0 † ‡ | _pendiente_ | _pendiente_ |
| Defectos detectados / resueltos | 2 / 1 | 1 / 2 * | _pendiente_ | _pendiente_ |
| Cobertura del dominio | 100% | 100% * | _pendiente_ | _pendiente_ |
| Tests | 52 | 263 * | _pendiente_ | _pendiente_ |

_* Sprint 2 al 24/10: se revisa al cierre, el domingo 25._

_† En el Sprint 2 la #5, la #14 y la #15 se implementaron con asistencia de IA, que la consigna
permite. Sus horas reales son el tiempo humano de dirección y revisión, y por eso la
desviación da muy negativa (ver el [Sprint 2](#sprint-2--interfaz-usable))._

_‡ La #77 se terminó en el Sprint 2 sin story points y el equipo la estimó en 13 el 26/10, en la
planning del Sprint 3. Ya estimada, suma al Sprint 2: al cierre eran 31 de 31, velocidad 39 y
1,3 h por story point._

**Velocidad del proyecto:** 29 (el promedio de los sprints cerrados: 6 en el Sprint 1 y 52 en el Sprint 2).

## Sprint 1 — MVP

- **Sprint Goal:** poder crear un proyecto con sus integrantes, cargarle el Product Backlog y
  estimar las historias en story points.
- **Duración:** tres semanas, del 28/09 al 16/10. La primera se fue entera en specs.
- Actas: [planning](actas/sprint-1-planning.md), [review](actas/sprint-1-review.md) y
  [retrospectiva](actas/sprint-1-retrospectiva.md).

### Historias y story points

| # | Historia | SP | ¿Terminada? | Responsable |
|---|---|---|---|---|
| 3 | Crear y modificar proyectos | 3 | sí | Agustín Salinas |
| 4 | Registrar integrantes de un proyecto | 3 | sí | Juliana Bustos |
| 6 | Administrar el Product Backlog | 5 | no al cierre (se mergeó el 19/10) | Emiliano Sorato |
| 12 | Estimar historias con Story Points | 2 | no (10 de 13 tareas) | Maximiliano Eula |
| 24 | Operar el sistema desde una interfaz de consola | 3 | no (7 de 15 tareas) | Maximiliano Eula |

- Comprometidos: **16 SP**. Completados: **6 SP**, el 38% de los puntos.
- Historias completadas: **2 de 5**, el 40%.
- **Velocidad del Sprint 1: 6.**

La #12 y la #24 volvieron al Product Backlog y se retomaron en el Sprint 2. La #6 cuenta en el
Sprint 2: al 16/10 tenía el código aprobado por Maximiliano Eula y el CI en verde, pero el
PR #58 se mergeó el 19/10. Las fechas están en el
[acta de review](actas/sprint-1-review.md#por-qué-la-6-no-cuenta-en-el-sprint-1). La primera
versión del acta la contaba y daba velocidad 11; se corrigió el 20/10.

### Esfuerzo

Las horas estimadas se cargaron el 02/10, antes de empezar a programar, con una conversión
plana de 3 horas por story point porque no había ningún dato anterior.

**Por historia terminada**

| # | Estimadas | Reales | Desviación |
|---|---|---|---|
| 3 | 9 | 8 | −1 h (−11%) |
| 4 | 9 | 8 | −1 h (−11%) |
| **Total** | **18** | **16** | **−2 h (−11%)** |

**Por integrante**, contando también las historias que no se terminaron

| Integrante | Estimadas | Reales |
|---|---|---|
| Emiliano Sorato | 15 | 13 |
| Agustín Salinas | 9 | 8 |
| Juliana Bustos | 9 | 8 |
| Maximiliano Eula | 15 | 9,25 |
| **Total** | **48** | **38,25** |

Hay dos desviaciones y conviene no mezclarlas:

- La **total** (−20%) compara las horas estimadas de las cinco historias contra lo que se
  trabajó en las cinco, terminadas o no. Sirve para ver cuánto tiempo puso el equipo.
- La **de las historias terminadas** (−11%) compara cada historia con lo que realmente costó
  terminarla. Es la que sirve para estimar.

En las dos historias terminadas **se estimó de más**. El costo real fue de **2,7 horas por
story point** (16 h ÷ 6 SP), contra las 3 que se usaron para estimar. El Sprint 2 se
planificó primero con 2,6, que contaba la #6, y el 20/10 sus horas se corrigieron a 2,7.

### Defectos

| # | Defecto | Detectado | Resuelto |
|---|---|---|---|
| 59 | Las fechas de un proyecto se comparan con hora, no por día | 12/10, en la revisión del PR #54 | 15/10, PR #61 (Sprint 1) |
| 60 | Los mensajes de error del dominio no llevan tildes | 12/10, en la revisión del PR #54 | 20/10 (después del cierre del Sprint 1) |

- Detectados en el sprint: **2**. Resueltos en el sprint: **1**.
- Los dos aparecieron en una revisión de código, no en un test: los encontró Maximiliano Eula
  revisando un PR ajeno.

### Cobertura y tests

Medido sobre `dev` en `4481d1c` (15/10), el último commit antes del cierre. La #6 entró
después (`87635b9`, 19/10), así que su código y sus tests cuentan en el Sprint 2.

| Paquete | Cobertura |
|---|---|
| `internal/domain` (núcleo y reglas de negocio) | **100%** |
| `cmd/app` (punto de entrada, sin lógica todavía) | 0% |
| **Total** | **99,2%** |

- **52 tests** en el dominio (102 casos, contando cada fila de los tests con tabla).
- El CI corre formato, análisis estático y los tests con cobertura en cada pull request.

### Evidencia de TDD

El test va en un commit y la implementación en otro. Commits de cada historia en `dev`:

| # | Commits de test | Commits de implementación |
|---|---|---|
| 3 | 12 | 11 |
| 4 | 15 | 11 |

La #6 (8 commits de test y 5 de implementación) va en el Sprint 2, donde se terminó.

_[Completar: revisar que en cada par el test esté antes que la implementación, y explicar
los casos donde no hay un commit de implementación por cada test.]_

## Sprint 2 — interfaz usable

> **Al 24/10, se revisa al cierre** del domingo 25 con el acta de review.

- **Sprint Goal:** que el sistema maneje sprints, esfuerzo y defectos, y que todo se pueda
  operar desde la consola.
- **Duración:** del 20/10 al 25/10, seis días.
- Actas: [planning](actas/sprint-2-planning.md), [review](actas/sprint-2-review.md) y
  [retrospectiva](actas/sprint-2-retrospectiva.md), las dos últimas en esqueleto hasta el cierre.

### Historias y story points

El compromiso es de **31 SP**: los 26 del plan más los 5 de la #6, que se mergeó el 19/10,
entre el Sprint 1 y el 2, y por eso cuenta acá. La #77 (interfaz web) entró al sprint sin
estimar: no suma puntos ni horas hasta que el equipo la estime.

| # | Historia | SP | Estado al 24/10 |
|---|---|---|---|
| 6 | Administrar el Product Backlog | 5 | Hecho |
| 12 | Estimar historias con Story Points | 2 | Hecho |
| 24 | Operar el sistema desde una interfaz de consola | 3 | Hecho |
| 7 | Priorizar el Product Backlog | 2 | Hecho |
| 8 | Crear un sprint con su Sprint Goal | 3 | Hecho (PR #138) |
| 9 | Asignar historias a un sprint | 5 | Hecho (PR #138) |
| 5 | Consultar el estado de un proyecto | 3 | Hecho |
| 14 | Registrar el esfuerzo real | 3 | Hecho |
| 15 | Gestionar defectos | 5 | Hecho (PR #129) |
| 77 | Interfaz web: proyectos, integrantes y backlog | sin estimar | Hecho (PR #95) |

- Adelantada desde el 21/10: la **#13** (8 SP y 21,6 h estimadas, Maximiliano Eula), que no
  suma a los 31 comprometidos. **Se terminó el 23/10** (PR #141), así que la velocidad del
  sprint es **31 + 8 = 39**, y el Sprint 3 se planifica sin esos 8 puntos. Está en el
  [acta de review](actas/sprint-2-review.md#historias-adelantadas).
- **La #77, estimada el 26/10 en 13 SP**, suma al Sprint 2, donde se terminó: el sprint queda
  con 44 SP terminados, 10 de 10 historias y **velocidad 52** (44 + 8 de la #13). Está en el
  [acta de review](actas/sprint-2-review.md).
- Completados al 24/10: **31 SP**, las 9 historias comprometidas. Tres llegaron del Sprint 1, y
  las otras seis son nuevas del sprint: la #7, la #5, la #14, la #8, la #9 y la #15.

### Esfuerzo

- Horas estimadas de las historias nuevas: **56,7**, a 2,7 h por story point. Las de la #6,
  la #12 y la #24 (30 h) ya están en el Sprint 1.
- Horas trabajadas en el sprint: **31**, contando la #13. Es el total del tablero
  (53,25) menos las 22,25 h que esas tres historias tenían cargadas al cierre del Sprint 1.
- Historias terminadas, medidas enteras: **108,3 h estimadas contra 51,75 reales (−52%)**, 1,3 h
  por story point. Entran la #13 (21,6 h contra 6; Maxi corrigió sus horas de 3 a 6) y la #15 (13,5 h contra 0,5).
- **La #5, la #14, la #15 y la #77 se implementaron con asistencia de IA**, que la consigna
  permite. Sus horas reales son el tiempo humano de dirección y revisión, todas de Emiliano
  Sorato: 0,5 h en la #5, en la #14 y en la #15, y 1,5 h en la #77. Por eso la desviación da muy negativa.
  Sin la #5, la #14 y la #15, las otras siete terminadas dan **−36%** (78,6 h contra 50,25) y
  1,8 h por story point.
- **Cuánto cuesta un punto en el Sprint 2.** La #6, la #12 y la #24 se trabajaron casi todas
  en el Sprint 1. Las nuevas del sprint hechas a mano, la #7, la #8 y la #9 (las tres de
  Agustín Salinas), dan **27 h estimadas contra 20 reales (−26%)**, 2 h por story point.
- La #77 no entra en la desviación porque no tiene horas estimadas con qué comparar su 1,5 h.
  Desde que tiene story points (13, desde el 26/10), sí entra en las horas por story point: el
  sprint da **1,0** (53,25 h ÷ 52 SP). Sin las cuatro hechas con asistencia de IA (la #5, la #14,
  la #15 y la #77), sigue en 1,8. También cuenta en las horas trabajadas.

### Defectos

| # | Defecto | Detectado | Resuelto |
|---|---|---|---|
| 60 | Los mensajes de error del dominio no llevan tildes | Sprint 1 (12/10) | Sprint 2 (20/10) |
| 108 | Tres tests de `internal/store` fallan en Windows | Sprint 2 (20/10) | Sprint 2 (21/10, PR #116) |

- Detectados en el sprint: **1**. Resueltos en el sprint: **2**. Abiertos: **0**.

### Cobertura y tests

Medido sobre `dev` en `b45cbe6` (24/10): dominio **100%**, `internal/store` 96,0%,
`internal/web` 93,8%, consola (`cmd/app`) 96,2%, arranque de la web (`cmd/web`) 56,2% y
**96,0% en total**, con **263 tests** (589 casos). Al 21/10 eran 121 tests y 98,8%: el total
baja porque entró la web, y en `cmd/web` no tienen test `main` ni `lanAddresses`. Desde el arreglo del
#108 el CI corre los tests también en Windows.

### Evidencia de TDD

El recorrido de cada historia —escenario, test, código y los commits RED y GREEN— está en la
[trazabilidad](trazabilidad.md), que se regenera con `scripts/trazabilidad.sh`.

## Sprint 3 — funcionalidad y calidad

_[Se completa el 30/10 con la review.]_

## Sprint 4 — cierre

_[Se completa el 01/11 con la review.]_

## Evolución del proyecto

_[Al final: velocidad sprint por sprint, horas por story point, cómo cambió la desviación y
la cobertura. Un gráfico de velocidad y uno de horas estimadas contra reales.]_

## Conclusiones

_[Al final. Qué aprendimos midiéndonos: si estimamos mejor sprint a sprint, qué nos costó
más de lo pensado y qué haríamos distinto.]_

## Notas sobre los datos

- **Limitación conocida: el tablero no conserva un sprint cerrado tal como se comprometió.**
  `scripts/metricas.sh` lee el sprint que cada historia tiene hoy. La #12, la #24 y la #6
  ahora tienen cargado *Sprint 2*, así que `scripts/metricas.sh "Sprint 1"` da 6 de 6 puntos
  (100%), cuando lo comprometido fueron 16. Por eso, **para un sprint cerrado el
  registro que vale es su acta de review**, que se escribe el día del cierre y no cambia. Este
  informe toma de ahí los números de cada sprint cerrado, y el script queda para el sprint
  abierto y para armar esa acta.
- **Limitación conocida: `scripts/metricas.sh` suma enteras las horas de una historia que
  viene de un sprint anterior.** El tablero guarda un solo total por historia, así que en el
  Sprint 2 la #12, la #24 y la #6 aparecen con sus 30 h estimadas y sus 22,25 h reales, que
  ya están en el Sprint 1. Para no contarlas dos veces, a lo que da el script se le resta lo que dice
  el acta de review del sprint anterior. La regla está en *Historias que pasan de un sprint a
  otro*, al principio del informe.
- **Defectos por sprint.** Hasta el 20/10, `scripts/metricas.sh` contaba todos los defectos
  del proyecto en cualquier sprint, y el Sprint 1 daba 2 resueltos cuando fue 1 (el #60 se
  cerró el 20/10). Ahora los filtra por las fechas de cada sprint: un defecto se detecta el día
  que se abre su issue y se resuelve el día que se cierra. Lo que cae entre dos sprints va al
  siguiente, así que para el script el Sprint 2 arranca el 17/10, el día después del cierre
  del Sprint 1.
- **Historias sin estimar.** Una historia terminada sin story points, como la #77 hasta el
  26/10, suma sus horas reales a las trabajadas en el sprint, pero `scripts/metricas.sh` no la
  cuenta en la desviación ni en las horas por story point, porque no tiene con qué compararlas.
  Cuando el equipo la estima, suma sus puntos al sprint donde se terminó. Si tiene story points
  pero no horas estimadas, como la #77, el script la cuenta en las horas por story point pero no
  en la desviación: para el Sprint 2 da −52%, igual que el acta.
- **Historias adelantadas.** `scripts/metricas.sh` las reconoce por la etiqueta `adelantada`
  de la issue: no las cuenta en lo comprometido ni en sus horas y las muestra en un bloque
  aparte. Si una no se termina y vuelve a su sprint del plan, hay que sacarle la etiqueta.
- **Historias agregadas durante el sprint.** Una historia que entra con el sprint ya empezado,
  sin haber estado en la planning (como la #177 en el Sprint 3), lleva la etiqueta `agregada`.
  `scripts/metricas.sh` la cuenta igual que una adelantada: no entra en lo comprometido ni en el
  porcentaje de historias completadas, la muestra en un bloque aparte y, si se termina, la suma
  a la velocidad.
- **Defectos del sprint.** El bloque de defectos del script lista solo los del sprint: los que se
  detectaron o se resolvieron en él, y los que quedaron abiertos al cierre.
- **Redondeo.** Los porcentajes se redondean al entero más cercano.
