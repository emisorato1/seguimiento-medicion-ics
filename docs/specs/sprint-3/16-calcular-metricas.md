# Spec: Calcular las métricas del proyecto y de cada sprint

- **Historia:** #16
- **Sprint:** 3
- **Responsable:** Emiliano Sorato
- **Estado:** Puerta 3 aprobada — **en implementación** (tareas 1 a 3 hechas)
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de la consigna y de los criterios de aceptación de la issue. El responsable lo revisa, lo corrige si hace falta y lo presenta a las puertas.
- **Depende de:** #8, #9 y #10 (los sprints, lo comprometido y lo completado), #11 (buscar un
  sprint por su número), #12 (los story points), #14 (las horas reales) y #15 (los defectos).
  Salvo la #12, ninguna está en `dev` todavía: se trabaja contra el contrato del dominio
- **La necesitan:** #17 (dashboard) y #18 (reporte), y detrás de la #18 la #19

## Fase 1 — Requisitos

### Objetivo

Que el equipo pueda pedir los números de un sprint o del proyecto entero —puntos, velocidad,
horas, historias completas y defectos—, calculados siempre de la misma manera, para saber con
datos cómo viene el proyecto y para planificar el sprint siguiente.

### Qué queda afuera

- **Mostrar los números**: el dashboard es la #17 y el reporte la #18. Esta historia devuelve
  números, no textos, gráficos ni valores redondeados para mostrar.
- **Registrar los datos de los que salen**: sprints y cierre (#8 a #10), story points (#12),
  horas reales (#14) y defectos (#15). Acá solo se leen. La excepción son las **horas
  estimadas** de cada historia, que la #14 dejó para esta historia.
- **Guardar las métricas**: se calculan cada vez que se piden.
- Métricas por integrante, tendencias entre sprints, burndown y proyecciones.

### Entradas

| Dato | Obligatorio | Descripción |
|---|---|---|
| Proyecto | sí | de qué proyecto se piden las métricas |
| Sprint | solo para las de un sprint | el número del sprint dentro del proyecto |
| Historia | solo para cargar horas estimadas | a qué historia se le carga la estimación en horas |
| Horas estimadas | solo para cargar horas estimadas | mayor que cero, admite fracciones |

### Salidas

Once números, del sprint pedido o del proyecto entero:

| Métrica | Qué es |
|---|---|
| Story points planificados | los puntos de las historias comprometidas |
| Story points completados | los puntos de las historias completadas |
| Velocidad | los puntos completados del sprint; en el proyecto, el promedio de sus sprints cerrados |
| Horas estimadas | las horas que se estimaron para las historias |
| Horas reales | las horas registradas como esfuerzo sobre esas historias |
| Desviación de horas | cuánto se apartaron las reales de las estimadas, en porcentaje |
| Historias planificadas | cuántas historias se comprometieron |
| Historias completadas | cuántas se completaron |
| Porcentaje de historias completas | completadas sobre planificadas, en porcentaje |
| Defectos detectados | cuántos defectos se encontraron |
| Defectos resueltos | cuántos se resolvieron |

Además, una historia queda con sus horas estimadas cuando se le cargan.

### Reglas de negocio

1. **[RN-1]** El sistema debe calcular, a pedido, las once métricas de un sprint de un
   proyecto y las del proyecto entero.
2. **[RN-2]** El sistema debe calcular las métricas cada vez que se piden, con los datos de
   ese momento, sin guardarlas y sin modificar ningún dato.
3. **[RN-3]** Cuando se piden las métricas de un sprint, el sistema debe tomar como
   planificadas todas las historias comprometidas en ese sprint, incluidas las que al
   cerrarlo volvieron al Product Backlog. Si el sprint está cerrado, debe usar —para las
   planificadas y para las completadas— los story points que cada historia tenía al cerrar,
   no los actuales.
4. **[RN-4]** Cuando se piden las métricas de un sprint, el sistema debe tomar como
   completadas las historias que quedaron registradas como completadas al cerrarlo, y no
   debe deducirlas del estado actual de las historias.
5. **[RN-5]** El sistema debe tomar como velocidad de un sprint sus story points completados.
6. **[RN-6]** Mientras un sprint esté abierto, el sistema debe calcular sus métricas igual,
   con los puntos completados, las historias completadas y la velocidad en cero, porque lo
   completado se registra recién al cerrar.
7. **[RN-7]** Cuando se piden las métricas de un sprint, el sistema debe sumar las horas
   estimadas y las reales de cada historia comprometida en él solo si ese es el **último**
   sprint —el de número más alto— en que la historia estuvo comprometida. Las horas reales de
   una historia son todas las registradas sobre ella, de cualquier integrante y día.
8. **[RN-8]** Cuando se piden las métricas de un sprint, el sistema debe contar como
   detectados los defectos que lo tienen como sprint de detección, y como resueltos los que
   lo tienen como sprint de resolución.
9. **[RN-9]** Cuando se piden las métricas del proyecto, el sistema debe tomar como
   planificadas todas las historias del Product Backlog, cada una una vez, hayan entrado a un
   sprint o no, y como completadas las registradas como completadas en algún sprint cerrado.
10. **[RN-10]** Cuando se piden las métricas del proyecto, el sistema debe tomar como
    velocidad el promedio de las velocidades de sus sprints cerrados, redondeado al entero más
    cercano y con las mitades hacia arriba. Si no tiene ningún sprint cerrado, la velocidad es
    cero.
11. **[RN-11]** Cuando se piden las métricas del proyecto, el sistema debe sumar las horas
    estimadas de todas sus historias y todas las horas registradas en el proyecto, y contar
    como detectados todos sus defectos y como resueltos los que están en estado Resuelto.
12. **[RN-12]** El sistema debe calcular la desviación de horas como la diferencia entre las
    horas reales y las estimadas, dividida por las estimadas y multiplicada por 100: es
    positiva si se trabajó más de lo estimado y negativa si se trabajó menos.
13. **[RN-13]** El sistema debe calcular el porcentaje de historias completas como las
    completadas divididas por las planificadas, multiplicado por 100.
14. **[RN-14]** Si las horas estimadas son cero, el sistema debe devolver la desviación en
    cero; si no hay historias planificadas, el porcentaje en cero. Ninguno es un error.
15. **[RN-15]** Si no hay datos cargados —ni historias, ni horas, ni defectos—, el sistema
    debe devolver las once métricas en cero, y eso no es un error.
16. **[RN-16]** El sistema debe devolver las horas, la desviación y el porcentaje con sus
    decimales, sin redondear. Una historia sin estimar aporta cero story points.
17. **[RN-17]** El sistema debe permitir cargar las horas estimadas de una historia con un
    valor mayor que cero, que admite fracciones, y volver a cargarlas reemplazando el valor
    anterior. Una historia a la que nunca se le cargaron tiene cero, y eso no es un error.
18. **[RN-18]** Si las horas estimadas son cero o negativas, el sistema debe rechazar la
    carga, informar que tienen que ser mayores que cero y dejar la historia como estaba.
19. **[RN-19]** Si la historia está Terminada, el sistema debe rechazar la carga de horas
    estimadas, informar que la estimación de una historia terminada no se modifica y dejar la
    historia como estaba.
20. **[RN-20]** Si el sprint pedido no existe en el proyecto, el sistema debe informar que no
    se encontró el sprint; si la historia no existe, que no se encontró la historia; y si el
    proyecto no existe, que no se encontró el proyecto.

### Casos límite

- Horas estimadas en cero con horas reales cargadas: la desviación es cero, no infinita.
  Reales en cero con estimadas cargadas: −100.
- Una historia completada de ocho: 12,5%, sin redondear.
- Historia que no se completó en el sprint 1 y se completó en el 2: planificada en los dos,
  completada solo en el 2, y sus horas cuentan solo en el 2.
- Defecto detectado y resuelto en el mismo sprint: suma en los dos números de ese sprint.
- Historia comprometida a la que después le quitaron la estimación: aporta cero puntos, pero
  cuenta como historia planificada.
- Sprints cerrados de velocidad 10 y 13: la del proyecto es 12, porque el 11,5 se redondea
  hacia arriba.
- Historia que volvió al backlog y se reestimó: el sprint cerrado sigue con los puntos que
  tenía al cerrar.

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| Sprint, historia o proyecto inexistente | informa que no se encontró |
| Horas estimadas cero o negativas | rechaza e informa que tienen que ser mayores que cero |
| Horas estimadas sobre una historia Terminada | rechaza e informa que la estimación de una historia terminada no se modifica |

Que falten datos **no** es un error: las métricas dan cero (RN-14, RN-15).

### Criterios de aceptación

- [ ] Los planificados de un sprint son todas sus historias comprometidas, incluidas las que
      volvieron al Product Backlog.
- [ ] Los completados de un sprint salen de lo registrado al cerrarlo: una historia que se
      completó en un sprint posterior no cuenta en el anterior.
- [ ] Los puntos de un sprint cerrado son los que tenía cada historia al cerrarlo: reestimar
      una historia después no los cambia.
- [ ] La velocidad de un sprint es igual a sus puntos completados, y la del proyecto es el
      promedio de sus sprints cerrados redondeado a entero.
- [ ] Las métricas de un sprint abierto se calculan, con completados y velocidad en cero.
- [ ] Las horas de un sprint suman las de sus historias, y las de una historia que pasó por
      dos sprints cuentan solo en el último.
- [ ] La desviación es la diferencia porcentual entre horas reales y estimadas, con signo, y
      con cero horas estimadas es cero sin error.
- [ ] El porcentaje de historias completas tiene decimales, y sin historias es cero sin error.
- [ ] Los defectos de un sprint salen de su sprint de detección y de resolución; los del
      proyecto son todos los registrados y los que están resueltos.
- [ ] Las métricas del proyecto incluyen todas sus historias, hayan entrado a un sprint o no.
- [ ] Un sprint o un proyecto sin datos devuelve las once métricas en cero, sin error.
- [ ] Pedir las métricas no cambia ningún dato, y si se cargan datos nuevos dan distinto.
- [ ] Se pueden cargar y reemplazar las horas estimadas de una historia no terminada, con
      fracciones; con cero, con un negativo o sobre una Terminada se rechaza sin cambiar nada.
- [ ] Un sprint, una historia o un proyecto inexistentes informan que no se encontraron.

### Escenarios BDD

```gherkin
Escenario: Metricas de un sprint cerrado
  Dado un proyecto con el sprint 1 cerrado
  Y el sprint 1 tenia comprometidas las historias 1, 2, 3 y 4, de 3, 5, 2 y 3 puntos
  Y al cerrarlo se registraron como completadas las historias 1, 2 y 3
  Y las historias tienen 4, 6, 2 y 4 horas estimadas y 5, 8, 3 y 4 horas registradas
  Y hay 2 defectos detectados en el sprint 1, de los cuales 1 se resolvio en el sprint 1
  Cuando pido las metricas del sprint 1
  Entonces recibo 13 story points planificados, 10 completados y velocidad 10
  Y recibo 4 historias planificadas, 3 completadas y un porcentaje de 75
  Y recibo 16 horas estimadas, 20 reales y una desviacion de 25
  Y recibo 2 defectos detectados y 1 resuelto

Escenario: Una historia que volvio al backlog cuenta como completada solo donde se completo
  Dado un proyecto con las historias 1 y 2, de 3 y 5 puntos
  Y el sprint 1 cerrado con las historias 1 y 2 comprometidas y solo la 1 completada
  Y el sprint 2 cerrado con la historia 2 comprometida y completada
  Cuando pido las metricas de los sprints 1 y 2
  Entonces el sprint 1 tiene 8 puntos planificados y 3 completados
  Y el sprint 2 tiene 5 puntos planificados y 5 completados

Escenario: Los puntos de un sprint cerrado no cambian si se reestima una historia
  Dado un proyecto con las historias 1 y 2, de 3 y 5 puntos
  Y el sprint 1 cerrado con las historias 1 y 2 comprometidas y solo la 1 completada
  Cuando reestimo la historia 2 en 8 puntos
  Y pido las metricas del sprint 1
  Entonces recibo 8 story points planificados y 3 completados

Escenario: Las horas de una historia que paso por dos sprints cuentan en el ultimo
  Dado un proyecto con el sprint 1 cerrado, con la historia 2 comprometida y sin completar
  Y el sprint 2 abierto con la historia 2 comprometida
  Y la historia 2 tiene 6 horas estimadas y 7 registradas
  Cuando pido las metricas de los sprints 1 y 2
  Entonces el sprint 1 tiene 0 horas estimadas y 0 reales
  Y el sprint 2 tiene 6 horas estimadas y 7 reales

Escenario: Las metricas de un sprint abierto son parciales
  Dado un proyecto con el sprint 1 abierto y la historia 1, de 3 puntos, comprometida
  Y la historia 1 tiene 2 horas registradas
  Cuando pido las metricas del sprint 1
  Entonces recibo 3 story points planificados, 0 completados y velocidad 0
  Y recibo 2 horas reales
  Y no es un error

Esquema del escenario: Sin datos cargados las metricas valen cero
  Dado <ambito> sin historias, sin horas y sin defectos
  Cuando pido sus metricas
  Entonces las once metricas valen cero
  Y no es un error

  Ejemplos:
    | ambito                                     |
    | un proyecto con el sprint 1 recien abierto |
    | un proyecto sin sprints                    |

Esquema del escenario: Calcular la desviacion de horas
  Dado un sprint cuyas historias suman <estimadas> horas estimadas y <reales> horas reales
  Cuando pido las metricas del sprint
  Entonces la desviacion de horas es <desviacion>

  Ejemplos:
    | estimadas | reales | desviacion |
    | 16        | 20     | 25         |
    | 10        | 8      | -20        |
    | 8         | 0      | -100       |
    | 0         | 5      | 0          |
    | 0         | 0      | 0          |

Esquema del escenario: Calcular el porcentaje de historias completas
  Dado un sprint cerrado con <planificadas> historias comprometidas y <completadas> completadas
  Cuando pido las metricas del sprint
  Entonces el porcentaje de historias completas es <porcentaje>

  Ejemplos:
    | planificadas | completadas | porcentaje |
    | 4            | 3           | 75         |
    | 8            | 1           | 12,5       |
    | 3            | 3           | 100        |
    | 0            | 0           | 0          |

Escenario: Defectos detectados y resueltos de cada sprint
  Dado un proyecto con los sprints 1 y 2
  Y el defecto 1 detectado en el sprint 1 y resuelto en el sprint 2
  Y el defecto 2 detectado y resuelto en el sprint 2
  Y el defecto 3 detectado en el sprint 2 y sin resolver
  Cuando pido las metricas de los sprints 1 y 2
  Entonces el sprint 1 tiene 1 defecto detectado y 0 resueltos
  Y el sprint 2 tiene 2 defectos detectados y 2 resueltos

Escenario: Metricas del proyecto entero
  Dado un proyecto con las historias 1 a 5, de 3, 5, 2, 3 y 8 puntos
  Y el sprint 1 cerrado con las historias 1, 2 y 3 comprometidas y la 1 y la 2 completadas
  Y el sprint 2 abierto con la historia 3 comprometida
  Y las historias 4 y 5 nunca entraron a un sprint
  Y en total 30 horas estimadas, 33 horas registradas y 3 defectos, 1 de ellos resuelto
  Cuando pido las metricas del proyecto
  Entonces recibo 21 story points planificados, 8 completados y velocidad 8
  Y recibo 5 historias planificadas, 2 completadas y un porcentaje de 40
  Y recibo 30 horas estimadas, 33 reales y una desviacion de 10
  Y recibo 3 defectos detectados y 1 resuelto

Esquema del escenario: Velocidad del proyecto
  Dado un proyecto cuyos sprints cerrados tienen velocidades <velocidades>
  Cuando pido las metricas del proyecto
  Entonces la velocidad es <velocidad>

  Ejemplos:
    | velocidades | velocidad |
    | 8, 5 y 6    | 6         |
    | 10 y 13     | 12        |
    | ninguna     | 0         |

Escenario: Las metricas reflejan los datos cargados despues
  Dado un proyecto con el sprint 1 cerrado y la historia 1 comprometida con 5 horas registradas
  Cuando pido las metricas del sprint 1
  Y registro 2 horas mas en la historia 1
  Y vuelvo a pedir las metricas del sprint 1
  Entonces la primera vez recibo 5 horas reales y la segunda 7

Escenario: Cargar y reemplazar las horas estimadas de una historia
  Dado una historia en estado "Pendiente" a la que nunca se le cargaron horas estimadas
  Entonces tiene cero horas estimadas
  Cuando le cargo 6,5 horas estimadas y despues 8
  Entonces la historia queda con 8 horas estimadas

Esquema del escenario: Rechazar horas estimadas que no son mayores que cero
  Dado una historia en estado "Pendiente" con 4 horas estimadas
  Cuando intento cargarle <horas> horas estimadas
  Entonces la operacion se rechaza
  Y se informa que las horas estimadas tienen que ser mayores que cero
  Y la historia sigue con 4 horas estimadas

  Ejemplos:
    | horas |
    | 0     |
    | -2    |

Escenario: No cambiar las horas estimadas de una historia terminada
  Dado una historia en estado "Terminada" con 4 horas estimadas
  Cuando intento cargarle 6 horas estimadas
  Entonces la operacion se rechaza
  Y se informa que la estimacion de una historia terminada no se modifica
  Y la historia sigue con 4 horas estimadas

Esquema del escenario: Pedir algo que no existe
  Dado un proyecto sin sprint 99 ni historia 99, y que no existe el proyecto 99
  Cuando intento <operacion>
  Entonces se informa que no se encontro <que>

  Ejemplos:
    | operacion                                | que         |
    | pedir las metricas del sprint 99         | el sprint   |
    | cargar horas estimadas en la historia 99 | la historia |
    | pedir las metricas del proyecto 99       | el proyecto |
```

**Para decidir en la puerta:**

1. **Las horas estimadas son un dato de cada historia.** La métrica las necesitaba, ninguna
   entidad las guardaba y la #14 dejó su carga para esta historia. Se interpretó que es un
   dato **de cada historia**, como el campo *Horas estimadas* del tablero, con las reglas de
   los story points de la #12 (RN-17 a RN-19). **Ya está en el contrato** (`Story` suma
   `EstimatedHours`, con `p.EstimateHours`). Como `story.go` es compartido, se avisa al grupo
   antes de tocarlo. Sin este dato, las horas estimadas y la desviación valdrían siempre cero
   y el tercer criterio de la issue quedaría a medias.
2. **Las horas de un sprint caen en el último sprint de cada historia (RN-7).** Es la
   simplificación del contrato, porque el esfuerzo no tiene sprint. Se aplica también a las
   estimadas, para que la desviación compare las mismas historias. La consecuencia: las horas
   de un sprint **ya cerrado** bajan si una historia que no se completó se vuelve a
   comprometer. Evitarlo exige ponerle sprint al registro de esfuerzo (#14 y contrato).
3. **Los planificados de un sprint cerrado no cambian.** Una historia que volvió al backlog
   se puede reestimar (la #12 lo permite), y eso movía los números de un sprint viejo.
   **Decidido el 20/10:** la #10 registra los puntos de cada historia al cerrar (su RN-16), y
   esta historia los usa para los sprints cerrados (RN-3).
4. **Dividir por cero da cero (RN-14)**, como pide la issue. Un "sin datos" sería más exacto,
   pero el contrato define las métricas como números planos: quien muestra puede decir "sin
   datos" cuando las horas estimadas son cero.
5. **La velocidad del proyecto es el promedio de los sprints cerrados (RN-10).** El contrato
   la define solo para un sprint; en el proyecto, repetir los puntos completados no dice nada
   nuevo y el promedio es lo que se usa para planificar. Va redondeada a entero, con las
   mitades hacia arriba, porque así la define el contrato y así redondea el dashboard (#17).
   El sprint abierto no cuenta.
6. **El proyecto cuenta todo su Product Backlog como planificado (RN-9)**, para que el
   porcentaje diga cuánto del proyecto está hecho. Contar solo lo que entró a un sprint
   repetiría lo que ya miden las métricas por sprint.
7. **Las horas cuentan todas las historias del sprint; los puntos, solo las completadas.** Es
   el mismo criterio con el que el equipo mide su tablero en cada review: el esfuerzo en lo
   que no se terminó es esfuerzo,
   pero una historia a medias no suma velocidad. Y como ahí, el redondeo es de quien muestra.

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: Emiliano Sorato — Fecha: 2026-10-20 — Comentarios: aprobada en un comentario de la issue #16. La aprueba el responsable de la historia porque el resto del equipo no estaba disponible: es una excepción a la regla de que aprueba otra persona, y queda para que otro integrante la ratifique en el PR de la historia. Antes de aprobar se corrigió: la velocidad del proyecto se redondea, como dice el contrato (10 y 13 dan 12); se agregaron el criterio y el escenario de los puntos de un sprint cerrado (RN-3), que ningún escenario probaba; y se actualizó el punto 1, porque las horas estimadas ya estaban en el contrato.

## Fase 2 — Diseño

Las decisiones transversales —un solo paquete `internal/domain`, datos en memoria, errores
comparables con `errors.Is`, mensajes en castellano y validar todo antes de cambiar nada—
están en [AGENTS.md](../../../AGENTS.md). Los nombres son los del
[contrato del dominio](../../diseno-dominio.md): `Metrics` con sus once datos, pedido con
`SprintMetrics` y `ProjectMetrics`; `EstimatedHours` y `p.EstimateHours` en `Story`; y lo
que se lee de `Sprint` (`Stories`, `CompletedStories`, `PointsAtClose`, `State`), `Effort`
y `Defect`.

### Enfoque

Son dos piezas. **La carga de horas estimadas** sigue el camino de la estimación en story
points de la #12: buscar la historia, controlar que no esté Terminada, validar el valor y
recién ahí aplicarlo. La historia nace con cero horas estimadas, igual que nace sin estimar.

**El cálculo** es una consulta pura (RN-2). Las dos consultas trabajan en dos pasos:

1. **Elegir el ámbito**: qué historias cuentan como planificadas, cuáles como completadas, de
   cuáles se suman las horas y qué defectos cuentan. Acá está todo lo que distingue a un
   sprint del proyecto (RN-3 a RN-11).
2. **Calcular**: un único cálculo, compartido, convierte ese ámbito en los once números. La
   desviación y el porcentaje, con su control de división por cero, están en un solo lugar.

Para la RN-7, el cálculo de un sprint averigua en el momento cuál es el último sprint de cada
historia, recorriendo los sprints del proyecto. No se guarda en ningún lado.

### Archivos afectados

| Archivo | Se crea o se modifica | Para qué |
|---|---|---|
| `internal/domain/metrics.go` | se crea | las métricas y las dos consultas, por sprint y por proyecto |
| `internal/domain/metrics_test.go` | se crea | los tests de los escenarios de métricas |
| `internal/domain/story_hours.go` | se crea | la carga de horas estimadas de una historia |
| `internal/domain/story_hours_test.go` | se crea | los tests de los escenarios de horas estimadas |
| `internal/domain/story.go` | se modifica | la historia pasa a tener sus horas estimadas, con cero al nacer |
| `internal/domain/errors.go` | se modifica | el error de horas estimadas inválidas |

### Decisiones

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| Las horas estimadas son un dato de la historia, con el mismo nombre que la métrica | un registro aparte de estimaciones, o deducirlas de los story points con un factor | es como lo lleva el tablero, y un factor inventaría una regla que nadie pidió. **Ya está en el contrato; se avisa al grupo antes de tocar `story.go`** |
| Un error propio para las horas estimadas inválidas; para la Terminada, el de la #12 | reutilizar el de horas inválidas de la #14 | la #14 no está en `dev` y no se la espera; el caso de la Terminada es el mismo de la #12, con el mismo mensaje |
| Un solo cálculo para los dos ámbitos, con el ámbito elegido antes | una consulta completa para el sprint y otra para el proyecto | las fórmulas son las mismas; escritas dos veces, tarde o temprano dan distinto |
| Se calculan en cada pedido | guardarlas en el sprint al cerrarlo | un esfuerzo cargado después del cierre quedaría afuera sin aviso; con datos en memoria, recalcular no cuesta |
| Los completados salen del registro de completadas del sprint | del estado actual de la historia | lo fija el contrato: deducirlo contaría dos veces una historia terminada en el sprint siguiente |
| El último sprint de cada historia se averigua al calcular | que la #14 guarde el sprint en cada registro | es la simplificación que aceptó el contrato y no toca la #14 |
| La división por cero devuelve cero | un error, o un valor especial de "sin datos" | lo pide la issue, y el dashboard y el reporte no tienen que manejar un caso aparte |
| Sin redondear, salvo la velocidad del proyecto | redondear como `scripts/metricas.sh` | redondear es de quien muestra; la velocidad va redondeada a entero porque el contrato la define así |

### Dependencias y restricciones

- **Depende de la #8 a la #11, la #14 y la #15, y ninguna está en `dev`.** No se espera: los
  tests arman dentro del paquete los sprints —abiertos y cerrados, con lo comprometido y lo
  completado—, los registros de esfuerzo y los defectos, contra el contrato.
- **Se reutilizan los errores** de sprint no encontrado (#11), de historia y proyecto no
  encontrados y de historia terminada (#12).
- **Solo se leen** sprints, registros de esfuerzo y defectos: no se modifica nada de la #10,
  la #14 ni la #15.
- **`story.go` y `errors.go` son compartidos**: en `story.go` solo se agrega un dato y su
  lectura, sin tocar lo de la #6 y la #12.
- **La #17 y la #18 arman métricas a mano en sus tests**: si cambia el significado de un
  número, se les avisa.

> **Puerta 2** — ¿el diseño es viable y coherente con el resto del proyecto?
> Aprobó: Emiliano Sorato — Fecha: 2026-10-20 — Comentarios: aprobada en el mismo comentario de la issue #16, con la misma excepción que la Puerta 1. Se agregó `PointsAtClose` a lo que se lee de `Sprint`, y se sacó `docs/diseno-dominio.md` de los archivos afectados porque el contrato ya tiene las horas estimadas.

## Fase 3 — Tareas

Cada tarea es un ciclo de TDD completo: el test que falla en un commit, la implementación que
lo hace pasar en otro, y cada una deja los tests en verde. Siguen la convención de
[AGENTS.md](../../../AGENTS.md): un test por escenario, con su nombre, y cada esquema como un
solo test con tabla de casos. Los números con decimales se comparan con una tolerancia chica.

| # | Tarea | Escenario BDD | Reglas | Estado |
|---|---|---|---|---|
| 1 | Una historia nace con cero horas estimadas; cargarlas las registra con fracciones, y volver a cargarlas reemplaza | Cargar y reemplazar las horas estimadas de una historia | RN-17 | hecha |
| 2 | Rechazar horas estimadas cero o negativas, sin cambiar la historia | Rechazar horas estimadas que no son mayores que cero | RN-18 | hecha |
| 3 | Rechazar la carga en una historia Terminada, sin cambiarla | No cambiar las horas estimadas de una historia terminada | RN-19 | hecha |
| 4 | Las métricas de un sprint y de un proyecto sin datos valen cero | Sin datos cargados las metricas valen cero | RN-1, RN-14, RN-15 | pendiente |
| 5 | Planificados de un sprint: puntos y cantidad de todas sus historias comprometidas | Metricas de un sprint cerrado | RN-3, RN-16 | pendiente |
| 6 | Completados y velocidad de un sprint salen del registro del cierre | Una historia que volvio al backlog cuenta como completada solo donde se completo | RN-4, RN-5 | pendiente |
| 7 | Los puntos de un sprint cerrado salen de los registrados al cerrarlo, no de los actuales | Los puntos de un sprint cerrado no cambian si se reestima una historia | RN-3 | pendiente |
| 8 | Un sprint abierto devuelve sus métricas con completados y velocidad en cero | Las metricas de un sprint abierto son parciales | RN-6 | pendiente |
| 9 | Horas de un sprint, y las de una historia que pasó por dos sprints cuentan en el último | Las horas de una historia que paso por dos sprints cuentan en el ultimo | RN-7 | pendiente |
| 10 | Desviación con signo, y cero sin horas estimadas | Calcular la desviacion de horas | RN-12, RN-14 | pendiente |
| 11 | Porcentaje de historias completas con decimales, y cero sin historias | Calcular el porcentaje de historias completas | RN-13, RN-14, RN-16 | pendiente |
| 12 | Defectos detectados y resueltos de cada sprint | Defectos detectados y resueltos de cada sprint | RN-8 | pendiente |
| 13 | Métricas del proyecto: todo el backlog, todo el esfuerzo y todos los defectos | Metricas del proyecto entero | RN-9, RN-11 | pendiente |
| 14 | Velocidad del proyecto: promedio redondeado de los sprints cerrados, y cero sin ninguno | Velocidad del proyecto | RN-10 | pendiente |
| 15 | Pedir las métricas no cambia ningún dato y refleja lo que se cargó después | Las metricas reflejan los datos cargados despues | RN-2 | pendiente |
| 16 | Sprint, historia o proyecto inexistentes informan que no se encontraron | Pedir algo que no existe | RN-20 | en curso: fila de la historia |

**El escenario "Metricas de un sprint cerrado" se completa en varias tareas.** La 5 escribe su
test verificando los planificados, y la 6, la 9, la 10, la 11 y la 12 le agregan la
verificación de su número. Al terminar la 12, el escenario está entero.

### Por qué este orden

- **De la 1 a la 3, las horas estimadas**: la tarea 9 las necesita, y es lo único que escribe.
- **La 4 fija las dos consultas** y su respuesta sin datos: lo que sigue es llenar los ceros.
- **De la 5 a la 9 se arma el ámbito de un sprint**, de lo más simple (lo comprometido) a lo
  más delicado (el último sprint de cada historia). La 7 va después de la 6 porque congela
  los puntos que la 5 y la 6 ya calculan. La 10 y la 11 son las fórmulas con división, que
  necesitan sus dos términos ya calculados.
- **La 13 y la 14 agregan el proyecto**, que reutiliza el cálculo ya probado.
- **La 15 y la 16 van al final**: protegen la consulta pura y los bordes de afuera.

### Cobertura

Las veinte reglas quedan cubiertas: RN-1 (4), RN-2 (15), RN-3 (5, 7), RN-4 (6), RN-5 (6),
RN-6 (8), RN-7 (9), RN-8 (12), RN-9 (13), RN-10 (14), RN-11 (13), RN-12 (10), RN-13 (11),
RN-14 (4, 10, 11), RN-15 (4), RN-16 (5, 11), RN-17 (1), RN-18 (2), RN-19 (3), RN-20 (16).

| Criterio de aceptación | Tareas |
|---|---|
| Planificados de un sprint, incluidas las que volvieron al backlog | 5 |
| Completados del registro del cierre, sin contar dos veces | 6 |
| Puntos de un sprint cerrado fijos al cerrarlo | 7 |
| Velocidad del sprint y del proyecto | 6, 14 |
| Métricas de un sprint abierto | 8 |
| Horas de un sprint, y de una historia que pasó por dos | 9 |
| Desviación con signo, y cero sin horas estimadas | 10 |
| Porcentaje con decimales, y cero sin historias | 11 |
| Defectos de un sprint y del proyecto | 12, 13 |
| El proyecto incluye todo su backlog | 13 |
| Sin datos, once ceros sin error | 4 |
| Pedir no cambia nada y refleja los datos nuevos | 15 |
| Cargar y reemplazar horas estimadas, y sus rechazos | 1, 2, 3 |
| Sprint, historia o proyecto inexistentes | 16 |

Los dieciséis escenarios BDD de la Fase 1 tienen su tarea en la columna "Escenario BDD".

### Dependencias

- **Las tareas 1 a 3 tocan `story.go`, que es compartido**: se avisa al grupo antes (punto 1
  de la puerta). La 3 necesita una historia Terminada: el test la prepara desde dentro del
  paquete, como la #12 en su tarea 9.
- **De la 4 a la 16 se necesitan sprints, esfuerzo y defectos** (#8 a #11, #14 y #15), que no
  están en `dev`: los tests los arman a mano contra el contrato. No se espera a nadie. La 7
  reestima con la operación de la #12, que ya está en `dev`.
- **La 16 necesita el error de sprint no encontrado de la #11.** Si todavía no está en `dev`,
  se declara acá con el nombre que fija la #11 y se resuelve el conflicto al traer `dev`.

> **Puerta 3** — ¿las tareas cubren todos los criterios de aceptación?
> Aprobó: Emiliano Sorato — Fecha: 2026-10-20 — Comentarios: aprobada en el mismo comentario de la issue #16, con la misma excepción que la Puerta 1. Se agregó la tarea 7, de los puntos de un sprint cerrado, y se renumeraron las siguientes.

## Fase 4 — Implementación

Ciclo RED → GREEN → REFACTOR. El test en un commit y la implementación en otro.

| Tarea | Commit del test | Commit de la implementación |
|---|---|---|
| 1 | `56ff0bc` | `7e4e3b8` |
| 2 | `eb4ad3d` | `e636a5c` |
| 3 | `4d24c31` | `9d00e6f` |
| 16 (fila de la historia) | `15a8336` | — (sin RED, ver Desvíos) |
| 2 (NaN, hallazgo de la revisión) | `4708708` | `9dd7309` |

Los RED de las tareas 1 y 2 fueron por compilación (la operación y el error todavía no
existían); el de la tarea 3, por comportamiento: la carga aceptaba una historia Terminada.

### Hallazgo de la revisión

**La carga aceptaba NaN ("no es un número") como horas estimadas.** La RN-18 pide rechazar
lo que no sea mayor que cero, y NaN no lo es, pero la condición estaba escrita como "menor o
igual que cero", que con NaN da falso. Puede llegar desde la consola, porque la conversión de
texto a número de Go acepta "NaN". Se encontró en la revisión del PR #81, antes de entrar a
`dev`, así que no se abrió defecto: se agregó un test sin escenario que lo rechaza (RED por
comportamiento) y la condición pasó a "no es mayor que cero". El infinito sí es mayor que
cero y cumple la regla tal como está escrita, así que no se rechaza.

### Desvíos respecto de la spec

Ninguno de comportamiento. Uno de orden:

- **La fila "cargar horas estimadas en la historia 99" de la tarea 16 se adelantó.** Con las
  tareas 1 a 3 la carga ya busca la historia, así que el caso de la historia inexistente
  pasó sin código nuevo; sin su test, la cobertura del dominio bajaba a 99,6%. Las filas del
  sprint y del proyecto se agregan en la tarea 16.
- **El error de horas inválidas se llama `ErrEstimatedHoursInvalid`** y no
  `ErrStoryHoursInvalid`, para no confundirlo con `p.StoryHours(id)`, que en el contrato son
  las horas **reales** de una historia (#14).

Nota de proceso: las tareas 4 a 16 necesitan `Sprint` (#8), `Effort` (#14) y `Defect` (#15).
Como ninguna está en `dev` y esas entidades son de sus dueños, las tareas 1 a 3 entran en un
primer PR con `Refs #16`, y el resto sigue cuando esas entidades lleguen a `dev`.

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #16 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | `internal/domain/metrics_test.go` y `internal/domain/story_hours_test.go` |
| Código Go | `internal/domain/metrics.go`, `internal/domain/story_hours.go`, el dato de horas estimadas en `internal/domain/story.go` y su error en `internal/domain/errors.go` |
