# Acta — Review del Sprint 1

> Los datos de métricas y evidencia están calculados con `scripts/metricas.sh "Sprint 1"` y
> verificados contra el tablero y el repositorio. Lo que está entre corchetes lo completa el
> equipo en la reunión. La retrospectiva va en un archivo aparte.

- **Fecha:** 19/10/2026. El sprint cerró el 16/10; la review se armó el 19/10, el mismo día
  que la retrospectiva y la planning del Sprint 2.
- **Participantes:** [quiénes estuvieron]
- **Facilita:** Emiliano Sorato (Agile Enabler)
- **Sprint Goal:** poder crear un proyecto con sus integrantes, cargarle el Product Backlog y
  estimar las historias en story points.
- **¿Se cumplió el Sprint Goal?** [sí / parcialmente / no, y por qué]

## Qué se completó

Una historia cuenta como completada solo si cumple la Definición de Terminado: spec con sus
tres puertas, escenarios BDD, tests en Go que pasan, revisada por otro integrante y mergeada
a `dev`.

| # | Historia | SP | ¿Terminada? | Evidencia |
|---|---|---|---|---|
| 3 | Crear y modificar proyectos | 3 | **sí** | PR #54 |
| 4 | Registrar integrantes | 3 | **sí** | PR #62 |
| 6 | Administrar el Product Backlog | 5 | no al cierre: se mergeó el 19/10 | PR #56 y #58 |
| 12 | Estimar con Story Points | 2 | no | 10 de 13 tareas, rama `historia-12-estimar-story-points` |
| 24 | Interfaz de consola | 3 | no | 7 de 15 tareas |

- **Story points comprometidos:** 16
- **Story points completados:** **6** (38%)
- **Velocidad del Sprint 1: 6**
- **Historias completadas:** 2 de 5 (40%)

Las dos terminadas, la #3 y la #4, cumplen la Definición de Terminado: spec con sus tres
puertas, escenarios BDD, tests en Go que pasan, revisadas por otro integrante y mergeadas a
`dev` antes del cierre, con el dominio al **100% de cobertura**.

Las historias no terminadas **vuelven al Product Backlog**, no se arrastran en silencio.

### Por qué la #6 no cuenta en el Sprint 1

_Corregido el 20/10._ La primera versión de esta acta contaba la #6 y daba velocidad 11,
3 de 5 historias y 2,6 h por story point. Al cierre, el 16/10, la #6 tenía todo hecho menos
una cosa, y la Definición de Terminado la exige: **no estaba mergeada a `dev`**.

| Fecha | Qué pasó |
|---|---|
| 02/10, 03/10 y 08/10 | Puertas 1, 2 y 3 de la spec, aprobadas por Maximiliano Eula |
| 09/10 | Se mergea el PR #56 (tareas 1 y 2) y se abre el PR #58 (tareas 3 a 14) |
| 15/10 | Último commit del PR #58, CI en verde, y aprobación de Maximiliano Eula |
| **16/10** | **Cierre del sprint: el PR #58 sigue abierto y la tarjeta en *En review*** |
| 19/10 | Se mergea el PR #58, sin cambios después de la aprobación, y la tarjeta pasa a *Hecho* |

La #6 se cuenta **en el Sprint 2**, igual que la #12: una historia cuenta en el sprint en que
se mergea a `dev`. Es el mismo criterio que se usa para los defectos, que cuentan en el sprint
en que se cierran (el #60 está más abajo).

**Otra cosa que hay que anotar de la #6:** el PR #56, que trajo sus tareas 1 y 2, se mergeó
sin que ninguna persona lo revisara. El resto de la historia (PR #58) sí tuvo revisión
cruzada. Es el mismo incumplimiento del acuerdo (b) del Sprint 0 que figura en la
retrospectiva.

## Esfuerzo

| Integrante | Horas estimadas | Horas reales |
|---|---|---|
| Emiliano Sorato | 15 | 13 |
| Agustín Salinas | 9 | 8 |
| Juliana Bustos | 9 | 8 |
| Maximiliano Eula | 15 | 9,25 |
| **Total** | **48** | **38,25** |

- **Desviación total:** −9,75 h (−20%).
- **Conversión recalibrada: 2,7 horas por story point completado** (16 h ÷ 6 SP), contando
  **solo las horas de las historias terminadas**. Se había estimado con 3 h por punto, así
  que **se estimó de más**: las historias costaron menos de lo previsto.

  Dividir las 38,25 horas totales por los 6 puntos terminados da 6,4 h, pero ese número
  mezcla las horas invertidas en historias que no entregaron ningún punto. Es el mismo error
  que el acta advierte más abajo para la desviación. **El número que sirve para estimar el
  Sprint 2 es 2,7.**

**Cómo leer la desviación.** El −20% mezcla las 48 horas estimadas de las cinco historias
contra las 38,25 reales de las dos terminadas más lo invertido en las tres que no. La
desviación limpia es la de cada historia terminada, y en las dos se estimó de más:

| Historia | Estimadas | Reales | Desviación |
|---|---|---|---|
| #3 | 9 | 8 | −1 h |
| #4 | 9 | 8 | −1 h |
| **Total** | **18** | **16** | **−2 h (−11%)** |

Las 15 h estimadas y las 13 reales de la #6 están en la tabla por integrante, porque se
trabajaron en este sprint, pero no entran en la conversión: la historia no se terminó acá.

## Defectos

| Detectados en el sprint | Resueltos | Abiertos |
|---|---|---|
| 2 | 1 | 1 |

| # | Defecto | Estado | Detectado en | Resuelto en |
|---|---|---|---|---|
| 59 | Las fechas de un proyecto se comparan con hora, no por día | resuelto | revisión del PR #54 | PR #61 |
| 60 | Los mensajes de error del dominio no llevan tildes | abierto | revisión del PR #54 | parcial: los de la #6 en el PR #58 |

Los dos los encontró Maximiliano Eula revisando código ajeno.

## Demostración

[Qué se mostró y qué dijeron los Product Architects. Si hubo pedidos de cambio, anotarlos: van
al Product Backlog como historias nuevas, no se arreglan sobre la marcha.]

## Pendientes

- [x] Cargar en el tablero las historias no terminadas de vuelta en Product Backlog: la #12
      y la #24 vuelven sin sprint asignado. La #24 conserva sus 9 tareas hechas.
- [x] Pasar la #6 al Sprint 2 en el tablero (20/10).
- [ ] Abrir el pull request de `dev` a `main` con el incremento y etiquetarlo `v0.1.0-sprint1`.
      Abierto el 20/10: PR #88, con `dev` al cierre del sprint (`4481d1c`). Falta mergearlo y
      poner el tag.
- [ ] Hacer la retrospectiva.
