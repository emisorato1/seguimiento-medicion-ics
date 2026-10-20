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

**El caso de la #12 y la #24.** Al cierre del Sprint 1 tenían cargadas 4 h y 5,25 h reales:
son las 9,25 h de Maximiliano que ya están en el Sprint 1. En el Sprint 2 se restan del total
que da el script. Sus horas estimadas (6 y 9) tampoco se vuelven a sumar, porque ya están en
las 48 del Sprint 1: el acta de planning del Sprint 2 lista 69,6 h porque las incluye, y en
este informe el Sprint 2 cuenta 54,6 h estimadas de historias nuevas.

Cuando la #12 se mide como historia terminada en el Sprint 2, entra entera: 6 h estimadas
contra 4 h reales. Esas 4 h aparecen en dos métricas distintas —las horas que puso el equipo
en el Sprint 1 y lo que costó la #12—, pero ninguna de las dos las suma dos veces.

## Resumen por sprint

| Métrica | Sprint 1 | Sprint 2 | Sprint 3 | Sprint 4 |
|---|---|---|---|---|
| Fechas | 28/09 → 16/10 | 20/10 → 25/10 | 26/10 → 30/10 | 31/10 → 01/11 |
| SP planificados | 16 | _pendiente_ | _pendiente_ | _pendiente_ |
| SP completados (velocidad) | **11** | _pendiente_ | _pendiente_ | _pendiente_ |
| Historias completadas | 3 de 5 (60%) | _pendiente_ | _pendiente_ | _pendiente_ |
| Horas estimadas | 48 | _pendiente_ | _pendiente_ | _pendiente_ |
| Horas reales | 38,25 | _pendiente_ | _pendiente_ | _pendiente_ |
| Desviación total | −20% | _pendiente_ | _pendiente_ | _pendiente_ |
| Desviación de las historias terminadas | −12% | _pendiente_ | _pendiente_ | _pendiente_ |
| Horas por story point completado | 2,6 | _pendiente_ | _pendiente_ | _pendiente_ |
| Defectos detectados / resueltos | 2 / 1 | _pendiente_ | _pendiente_ | _pendiente_ |
| Cobertura del dominio | 100% | _pendiente_ | _pendiente_ | _pendiente_ |
| Tests | 72 | _pendiente_ | _pendiente_ | _pendiente_ |

**Velocidad del proyecto:** 11 (el promedio de los sprints cerrados; por ahora, uno solo).

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
| 6 | Administrar el Product Backlog | 5 | sí | Emiliano Sorato |
| 12 | Estimar historias con Story Points | 2 | no (10 de 13 tareas) | Maximiliano Eula |
| 24 | Operar el sistema desde una interfaz de consola | 3 | no (7 de 15 tareas) | Maximiliano Eula |

- Comprometidos: **16 SP**. Completados: **11 SP**, el 69% de los puntos.
- Historias completadas: **3 de 5**, el 60%.
- **Velocidad del Sprint 1: 11.**

La #12 y la #24 volvieron al Product Backlog y se retomaron en el Sprint 2.

### Esfuerzo

Las horas estimadas se cargaron el 02/10, antes de empezar a programar, con una conversión
plana de 3 horas por story point porque no había ningún dato anterior.

**Por historia terminada**

| # | Estimadas | Reales | Desviación |
|---|---|---|---|
| 3 | 9 | 8 | −1 h (−11%) |
| 4 | 9 | 8 | −1 h (−11%) |
| 6 | 15 | 13 | −2 h (−13%) |
| **Total** | **33** | **29** | **−4 h (−12%)** |

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
- La **de las historias terminadas** (−12%) compara cada historia con lo que realmente costó
  terminarla. Es la que sirve para estimar.

En las tres historias terminadas **se estimó de más**. El costo real fue de **2,6 horas por
story point** (29 h ÷ 11 SP), contra las 3 que se usaron para estimar. Ese 2,6 es el número
con el que se planificó el Sprint 2.

### Defectos

| # | Defecto | Detectado | Resuelto |
|---|---|---|---|
| 59 | Las fechas de un proyecto se comparan con hora, no por día | 12/10, en la revisión del PR #54 | 15/10, PR #61 (Sprint 1) |
| 60 | Los mensajes de error del dominio no llevan tildes | 12/10, en la revisión del PR #54 | 20/10 (después del cierre del Sprint 1) |

- Detectados en el sprint: **2**. Resueltos en el sprint: **1**.
- Los dos aparecieron en una revisión de código, no en un test: los encontró Maximiliano Eula
  revisando un PR ajeno.

### Cobertura y tests

Medido sobre `dev` en `87635b9` (19/10), el merge del PR #58, que completó la #6. La #6
entró a `dev` después del 16/10, pero la review la cuenta en el Sprint 1, así que la
cobertura se mide ahí: es el último cambio de código del Sprint 1.

| Paquete | Cobertura |
|---|---|
| `internal/domain` (núcleo y reglas de negocio) | **100%** |
| `cmd/app` (punto de entrada, sin lógica todavía) | 0% |
| **Total** | **99,5%** |

- **72 tests** en el dominio (130 casos, contando cada fila de los tests con tabla).
- El CI corre formato, análisis estático y los tests con cobertura en cada pull request.

### Evidencia de TDD

El test va en un commit y la implementación en otro. Commits de cada historia en `dev`:

| # | Commits de test | Commits de implementación |
|---|---|---|
| 3 | 12 | 11 |
| 4 | 15 | 11 |
| 6 | 8 | 5 |

_[Completar: revisar que en cada par el test esté antes que la implementación, y explicar
los casos donde no hay un commit de implementación por cada test.]_

## Sprint 2 — interfaz usable

_[Se completa el 25/10 con la review. Mismas secciones que el Sprint 1.]_

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
  `scripts/metricas.sh` lee el sprint que cada historia tiene hoy. La #12 y la #24 volvieron
  al backlog y ahora tienen cargado *Sprint 2*, así que `scripts/metricas.sh "Sprint 1"` da 11
  de 11 puntos (100%), cuando lo comprometido fueron 16. Por eso, **para un sprint cerrado el
  registro que vale es su acta de review**, que se escribe el día del cierre y no cambia. Este
  informe toma de ahí los números de cada sprint cerrado, y el script queda para el sprint
  abierto y para armar esa acta.
- **Horas de las historias que pasan de sprint.** Cada hora cuenta una sola vez. La regla
  está en *Historias que pasan de un sprint a otro*, al principio del informe.
- **Redondeo.** Los porcentajes se redondean al entero más cercano.
