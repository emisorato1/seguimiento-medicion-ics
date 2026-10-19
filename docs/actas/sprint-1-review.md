# Acta — Review del Sprint 1

> Los datos de métricas y evidencia están calculados con `scripts/metricas.sh "Sprint 1"` y
> verificados contra el tablero y el repositorio. Lo que está entre corchetes lo completa el
> equipo en la reunión. La retrospectiva va en un archivo aparte.

- **Fecha:** [completar — el sprint cerró el 16/10, la review se hizo el dd/mm]
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
| 6 | Administrar el Product Backlog | 5 | **sí** | PR #56 y #58 |
| 12 | Estimar con Story Points | 2 | no | 10 de 13 tareas, rama `historia-12-estimar-story-points` |
| 24 | Interfaz de consola | 3 | no | 7 de 15 tareas |

- **Story points comprometidos:** 16
- **Story points completados:** **11**
- **Velocidad del Sprint 1: 11**
- **Historias completadas:** 3 de 5 (60%)

Las tres terminadas cumplen la Definición de Terminado: spec con sus tres puertas, escenarios
BDD, tests en Go que pasan y mergeadas a `dev`, con el dominio al **100% de cobertura**.

**Con una excepción que hay que anotar:** el PR #56, que trajo las tareas 1 y 2 de la #6, se
mergeó sin que ninguna persona lo revisara. El resto de la historia (PR #58) sí tuvo revisión
cruzada, pero la Definición de Terminado no se cumple al pie de la letra para esa parte. Es
el mismo incumplimiento del acuerdo (b) del Sprint 0 que figura en la retrospectiva.

Las historias no terminadas **vuelven al Product Backlog**, no se arrastran en silencio.

## Esfuerzo

| Integrante | Horas estimadas | Horas reales |
|---|---|---|
| Emiliano Sorato | 15 | 13 |
| Agustín Salinas | 9 | 8 |
| Juliana Bustos | 9 | 8 |
| Maximiliano Eula | 15 | 9,25 |
| **Total** | **48** | **38,25** |

- **Desviación total:** −9,75 h (−20%).
- **Conversión recalibrada: 2,6 horas por story point completado** (29 h ÷ 11 SP), contando
  **solo las horas de las historias terminadas**. Se había estimado con 3 h por punto, así
  que **se estimó de más**: las historias costaron menos de lo previsto.

  Dividir las 38,25 horas totales por los 11 puntos terminados da 3,5 h, pero ese número
  mezcla las horas invertidas en historias que no entregaron ningún punto. Es el mismo error
  que el acta advierte más abajo para la desviación. **El número que sirve para estimar el
  Sprint 2 es 2,6.**

**Cómo leer la desviación.** El −21% mezcla las 48 horas estimadas de las cinco historias
contra las 38,25 reales de las tres terminadas más lo invertido en las dos que no. La
desviación limpia es la de cada historia terminada, y en las tres se estimó de más:

| Historia | Estimadas | Reales | Desviación |
|---|---|---|---|
| #3 | 9 | 8 | −1 h |
| #4 | 9 | 8 | −1 h |
| #6 | 15 | 13 | −2 h |

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
- [ ] Abrir el pull request de `dev` a `main` con el incremento y etiquetarlo `v0.1.0-sprint1`.
- [ ] Hacer la retrospectiva.
