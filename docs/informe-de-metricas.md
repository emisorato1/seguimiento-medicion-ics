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
| Cobertura | `go test ./... -coverprofile`, el mismo comando que corre el CI, sobre `dev` en el commit de cierre |
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
| Tests | 52 | _pendiente_ | _pendiente_ | _pendiente_ |

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

Medido sobre `dev` en el último commit antes del cierre (`4481d1c`, 15/10):

| Paquete | Cobertura |
|---|---|
| `internal/domain` (núcleo y reglas de negocio) | **100%** |
| `cmd/app` (punto de entrada, sin lógica todavía) | 0% |
| **Total** | **99,2%** |

- **52 tests** en el dominio, de las tres historias terminadas.
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

- **El tablero ya no muestra el Sprint 1 como se comprometió.** La #12 y la #24 volvieron al
  backlog y hoy tienen cargado *Sprint 2*, así que `scripts/metricas.sh "Sprint 1"` da 11 de
  11 puntos (100%). Lo comprometido fueron 16: este informe toma los números del acta de
  planning y de la review, que se escribieron en su momento.
- **Las horas de la #12 cuentan en el sprint donde se terminó.** Las 4 horas reales que
  tiene cargadas se trabajaron casi todas en el Sprint 1, pero la historia cerró en el
  Sprint 2. Es la misma simplificación que usa la aplicación: el esfuerzo no tiene sprint, y
  las horas de una historia caen en el último sprint donde estuvo.
- **Redondeo.** Los porcentajes se redondean al entero más cercano.
