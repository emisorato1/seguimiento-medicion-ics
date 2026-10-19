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
| 12 | Estimar con Story Points | 2 | no | no se empezó |
| 24 | Interfaz de consola | 3 | no | 9 de 15 tareas |

- **Story points comprometidos:** 16
- **Story points completados:** **11**
- **Velocidad del Sprint 1: 11**
- **Historias completadas:** 3 de 5 (60%)

Las tres terminadas cumplen la Definición de Terminado: spec con sus tres puertas, escenarios
BDD, tests en Go que pasan, revisadas por otro integrante y mergeadas a `dev`. La cobertura
del dominio es del **100%**.

Las historias no terminadas **vuelven al Product Backlog**, no se arrastran en silencio.

## Esfuerzo

| Integrante | Horas estimadas | Horas reales |
|---|---|---|
| Emiliano Sorato | 15 | 13 |
| Agustín Salinas | 9 | 8 |
| Juliana Bustos | 9 | 8 |
| Maximiliano Eula | 15 | 8,25 |
| **Total** | **48** | **38,25** |

- **Desviación total:** −9,75 h (−21%).
- **Conversión recalibrada: 3,5 horas por story point completado** (38,25 h ÷ 11 SP). Se
  había estimado con 3 h por punto, así que la calibración inicial estaba bien.

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
