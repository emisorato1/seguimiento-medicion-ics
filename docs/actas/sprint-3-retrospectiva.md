# Acta — Retrospectiva del Sprint 3

> Se hizo por escrito, después de la [review](sprint-3-review.md). Lo que está entre corchetes,
> **[pendiente]**, lo completa lo que diga cada uno: acá no se escribe nada que alguien no haya
> dicho.

- **Fecha:** 29 y 30/10, por escrito.
- **Respondieron:** Maximiliano Eula y Juliana Bustos (29/10), y Emiliano Sorato y Agustín
  Salinas (30/10), por escrito.
- **Facilita:** Emiliano Sorato (Agile Enabler)

Los números con los que se abre la retro salen del [acta de review](sprint-3-review.md):
28 de los 41 story points (68%) y 5 de las 6 historias; la #78 vuelve al Product Backlog.
Velocidad 33, con los 5 de la #73 adelantada. Las historias terminadas dan −51% de desviación
en horas (−21% sin la #10 y la #16, hechas con asistencia de IA). Se detectaron 6 defectos y
se resolvieron los 6. La cobertura quedó en 95,7%, con 420 tests.

## Qué salió bien

- **Emiliano Sorato:** "Con las specs escritas antes, la #10 y la #16 entraron el 26/10, el
  primer día, y destrabaron la #17 y la #18. La prueba exploratoria de la web del 26/10 sirvió
  mucho: salieron cuatro defectos y los cuatro se resolvieron antes del cierre. Y con la #177
  los datos quedaron en una base de verdad sin romper lo que ya andaba".
- **Maximiliano Eula:** "terminé la #73 y me dediqué mucho a revisar PRs. Varias revisiones encontraron errores
  antes de que llegaran a dev, como el defecto #172, el tema del -datos con JSON y el orden del
  reporte".
- **Juliana Bustos:** "Cerré la #17 (Dashboard) con las 17 tareas y los 16 escenarios con test, TDD
  estricto y 100 % de cobertura en el dominio. Partir la historia ayudó: el PR #204 con el
  dominio y format.go destrabó a Agus para la #18. Las revisiones cruzadas sirvieron para
  detectar detalles reales a tiempo (como el orden de las historias en la RN-13)".
- **Agustín Salinas:** "Terminé la #11 y la #18 con TDD y dentro de lo estimado: la #11 en
  4,5 h de 5,4, y la #18 en 10 h de 13,5. Cada test quedó en un commit propio con su salida
  RED, antes del código. El historial muestra el ciclo completo".

## Qué no salió bien

- **Emiliano Sorato:** "La #78 no llegó: le quedaron el reporte y el PDF, que dependían de la
  #18 y de la #19, y como tenía muchas cosas a la vez los dejé para el final. También abrí
  muchos PRs de documentación que chocaban en ESTADO.md y en el índice de specs. Y las horas de
  lo que hice con asistencia de IA no se parecen en nada a lo estimado (la #10 y la #16 dan
  −33,6 h entre las dos), así que no sirven para estimar".
- **Maximiliano Eula:** "mi #73 se mergeó con un cambio pedido por Juli sin responder, y después hubo que
  arreglarlo (#156). La #177 entró a mitad de sprint sin estimar. Y perdimos bastante tiempo
  resolviendo conflictos en ESTADO.md y en los documentos".
- **Juliana Bustos:** "Se acumularon muchos PRs al cierre y los archivos compartidos (ESTADO.md,
  índice de specs y diseno-dominio.md) chocaban seguido. Además, cuatro tareas pasaron sin RED
  directo porque el código previo ya las cubría (se puede ordenar mejor la spec)".
- **Agustín Salinas:** "La spec de la #18 necesitó varias vueltas de revisión antes de quedar
  cerrada. La tarea 11 pasó sin RED y me dio una falsa seguridad. La regla de ordenar por
  identificador (RN-13) estaba rota en el reporte de sprint, y la encontraron Juliana y Maxi en
  la revisión. Al resolver un conflicto del índice de specs cambié por error la fila de la #78
  junto con la de la #18. Lo encontró Emiliano".

## Seguimiento de los acuerdos del Sprint 2

En la [retrospectiva del Sprint 2](sprint-2-retrospectiva.md) se propusieron tres acuerdos para
el Sprint 3. Votaron los cuatro en la [issue #147](https://github.com/emisorato1/seguimiento-medicion-ics/issues/147),
todos a favor de las tres: Emiliano Sorato (26/10), Juliana Bustos (26/10), Agustín Salinas
(29/10) y Maximiliano Eula (29/10). Juliana Bustos puso una condición a la tercera: "siempre que
tengan las tres puertas aprobadas y lo que necesitan esté en dev".

| Propuesta | ¿Quedó como acuerdo? | ¿Se cumplió? |
|---|---|---|
| Reservar una franja fija por día para revisar los PRs asignados | Sí, 4 de 4 | No se midió |
| Pedir revisor apenas se abre el PR | Sí, 4 de 4 | No se midió |
| Seguir adelantando la próxima historia apenas uno termina lo suyo | Sí, 4 de 4, con la condición de Juliana Bustos | Sí: la #73 se adelantó con sus tres puertas aprobadas y se terminó el 26/10 |

Lo que ya muestra el tablero: en el Sprint 3 se adelantó la #73 (Maximiliano Eula), que se
terminó el 26/10.

## Lo que cambió durante el sprint

Para que la retro lo evalúe, esto es lo que cambió en el proceso entre el 26 y el 28/10:

- **El sprint se acortó** del 30/10 al 28/10, a pedido del equipo, para terminar todo el 29/10
  y presentar el 2/11.
- **Entró una historia con el sprint empezado:** la #177 (base de datos SQLite), que se cuenta
  aparte con la etiqueta `agregada`.
- **La #77 y la #78 se estimaron en 13 y 13** en la ronda 2, con los cuatro de acuerdo.
- **Firma de un compañero:** desde el 26/10, los PRs que salen de la cuenta de Emiliano Sorato
  esperan la firma de otro integrante antes del merge.
- **Prueba exploratoria de la web** el 26/10: salieron cuatro defectos (#190 a #193).

## Qué vamos a cambiar en el Sprint 4

El Sprint 4 dura un día (29/10), así que lo que salga acá sirve sobre todo para la
presentación y para la retro final.

- **Emiliano Sorato:** "Que ESTADO.md lo regenere una sola persona una vez por día, y lo tomo
  yo. Terminar primero la #78 y la #19, y después los documentos. Y lo que dijo Maxi: no
  mergear con cambios pedidos sin responder".
- **Maximiliano Eula:** "no mergear nada con cambios pedidos sin responder, aunque estemos apurados. Y regenerar
  ESTADO.md una vez por día, no en cada PR".
- **Juliana Bustos:** "mergear apenas hay revisión y unificar quién regenera ESTADO.md".
- **Agustín Salinas:** "Cuando un test pasa sin RED, revisar los otros lugares donde se aplica
  la misma regla antes de darla por cubierta. Resolver los conflictos de documentación fila por
  fila, nunca con un reemplazo global".

Los tres primeros acuerdos salen de lo que propusieron Maximiliano Eula, Juliana Bustos y
Emiliano Sorato, que coinciden; el cuarto, de lo que propuso Agustín Salinas:

| Acuerdo | Responsable |
|---|---|
| ESTADO.md lo regenera una sola persona, una vez por día, y no cada PR | Emiliano Sorato |
| No se mergea un PR con cambios pedidos sin responder | todos |
| Se mergea apenas el PR tiene su revisión y el CI en verde | el autor de cada PR |
| Un test que pasa sin RED se revisa contra los otros lugares donde se aplica la misma regla | el autor de cada historia |

## Pendientes

- [x] Las respuestas de los cuatro (30/10).
