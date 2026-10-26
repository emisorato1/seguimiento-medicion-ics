# Datos de demo: nuestro propio proyecto

`datos-demo.json` es el proyecto integrador cargado en el sistema **tal como estaba al cierre
del Sprint 2 (25/10/2026)**, para mostrarlo en la presentación. Todo sale de lo registrado:
el [tablero](https://github.com/users/emisorato1/projects/3), los issues de cada historia y
las [actas](../actas). Lo que no está registrado quedó afuera (ver más abajo).

## Cómo se abre

El sistema guarda cada cambio en el archivo que se le indica, así que conviene abrir una
copia para no tocar este:

```bash
cp docs/demo/datos-demo.json /tmp/demo.json
go run ./cmd/web -datos /tmp/demo.json
```

## Qué tiene

| Qué | De dónde sale |
|---|---|
| **El proyecto**: nombre, inicio el 07/09 (primer commit) y fin el 02/11 (presentación) | `README.md`, acta de cierre del Sprint 0 y `CLAUDE.md` |
| **Los cuatro integrantes** con su rol: Emiliano Sorato (Agile Enabler), Agustín Salinas, Juliana Bustos y Maximiliano Eula (Product Builders) | `CLAUDE.md` y las actas |
| **Las 21 historias** con su número de issue, título, descripción («Como… quiero… para…»), criterios de aceptación, prioridad, story points y horas estimadas | los issues y los campos del tablero |
| **Sprint 1** (28/09 al 16/10, cerrado el 16/10): su Sprint Goal, las 5 historias asignadas y las 2 completadas (la #3 y la #4) | actas de planning y review del Sprint 1 |
| **Sprint 2** (20 al 25/10, cerrado el 25/10): su Sprint Goal y las 10 historias, todas completadas, con la #13 adelantada | anexo de la planning y review del Sprint 2 |
| **13 registros de esfuerzo**, uno por historia con horas reales, 69,25 h en total | campo *Horas reales* del tablero; las 38,25 h del Sprint 1 y las 31 h del Sprint 2 de las actas de review |
| **El Planning Poker de la #77 y la #78**, del 26/10: la ronda 1 de cada una, revelada con los cuatro votos y sus diferencias, y la ronda 2, revelada con los cuatro votos (13 de todos) y la estimación acordada en 13. Con eso la #77 y la #78 tienen 13 story points, como en el tablero | acta de planning del Sprint 3 (PR #169 y #171); el voto de Juliana Bustos en la ronda 2 lo registró Emiliano Sorato el 26/10, y el acta todavía lo tiene pendiente |
| **3 defectos**, todos resueltos, con su número de issue: el #59 (historia #3, detectado y resuelto en el Sprint 1), el #60 (historia #3, detectado en el Sprint 1 y resuelto en el 2) y el #108 (historia #77, detectado y resuelto en el Sprint 2) | los issues con la etiqueta `defecto` (título, «Qué pasa» y «Datos para el informe») y la review de cada sprint |

El estado de cada historia es el que tenía al cierre del Sprint 2: Terminada si se completó
en el Sprint 1 o en el 2, Pendiente si no. La #10, la #16 y la #73 se terminaron después, en
el Sprint 3, y acá figuran Pendientes.

## Lo que hubo que adaptar

- **Las rondas son del 26/10**, del Sprint 3, aunque el resto es la foto al cierre del Sprint
  2: se suman porque son el único Planning Poker registrado. Por eso la #77 y la #78 figuran
  con los 13 story points que se acordaron ese día, y la #78 con las 35,1 horas estimadas que
  se le cargaron después en el tablero. La #77 figura Terminada y con sus rondas, porque se
  estimó después de terminarla, para la velocidad (acta de planning del Sprint 3).
- **La severidad de los defectos la asignó Emiliano Sorato el 26/10**: ningún registro la
  tenía y el sistema la pide. El #59 y el #108 son Mayor; el #60, Menor.
- **El #60 tocaba la #3 y la #6**, y el sistema relaciona cada defecto con una sola historia:
  figura con la #3, que es la que su issue nombra primero.
- **El día de cada registro de esfuerzo es el día en que la historia se mergeó a `dev`.**
  El tablero guarda las horas reales por historia, sin día, y el sistema pide uno. Las horas
  se trabajaron en los días anteriores, no todas ese día.
- **Un registro por historia, a nombre de su responsable**, con las horas totales del
  tablero. Para la #5, la #14, la #15 y la #77, la review del Sprint 2 dice que son horas de
  Emiliano Sorato dirigiendo y revisando la implementación con asistencia de IA, y así figura
  la actividad.
- **La #77 está Terminada pero fuera del Sprint 2**: se terminó en ese sprint sin story
  points (se estimó recién el 26/10), y el sistema solo asigna historias estimadas.

## Lo que quedó afuera, y por qué

| Qué | Por qué |
|---|---|
| **El defecto #156** | se detectó en el Sprint 3, después de esta foto |
| **Rondas de Planning Poker de los Sprints 1 y 2** | no hay votos registrados: el Sprint 1 se estimó por afinidad y el Sprint 2 mantuvo esas estimaciones (actas de planning) |
| **El Sprint 3** | estaba en curso al momento de esta foto |
| **El reparto de horas por persona del Sprint 2** | la review lo tiene pendiente; las horas van a nombre del responsable de cada historia |

El archivo se armó con un script que no está en el repo; para regenerarlo con el tablero
actualizado, pedírselo a la sesión de la web.
