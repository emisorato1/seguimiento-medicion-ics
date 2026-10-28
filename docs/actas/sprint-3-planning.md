# Acta — Sprint Planning del Sprint 3

- **Fecha:** del 24 al 26/10/2026
- **Cómo se hizo:** por escrito, en la [issue #147](https://github.com/emisorato1/seguimiento-medicion-ics/issues/147), con un comentario de cada uno
- **Participantes:** Emiliano Sorato, Maximiliano Eula, Juliana Bustos y Agustín Salinas
- **Facilita:** Emiliano Sorato (Agile Enabler)
- **Duración del sprint:** del lunes 26 al miércoles 28/10. Al principio iba hasta el viernes 30/10;
  el 26/10 se acortó a pedido del equipo, para terminar todo el 29/10 (ver el
  [plan](../plan-final-al-2-de-noviembre.md#calendario))

## Lo que traemos del Sprint 2

- **31 de 31 story points** comprometidos terminados, más los 8 de la #13 adelantada:
  **velocidad 39**. Los números completos están en el [acta de review](sprint-2-review.md).
- **Punto de control del 25/10:** el Sprint 2 cerró más de 26 puntos, así que, por decisión de
  Emiliano Sorato del 24/10, **el plan sigue tal cual**.
- **Conversión para estimar: 2,7 horas por story point**, la misma del Sprint 2.

## Sprint Goal

> **Que el sistema mida el proyecto y lo muestre: métricas, dashboard y reporte.**

Lo propuso Emiliano en la #147, y los cuatro contestaron **"de acuerdo"**.

## Sprint Backlog

Cada responsable confirmó sus historias en la #147. El compromiso son **41 story points**: los
36 del plan menos los 8 de la #13, que se terminó en el Sprint 2 (28), más los 13 de la #78, que
entró sin estimar y se estimó el 26/10 en el Planning Poker de esta planning. Como dice el
[anexo de la planning del Sprint 2](sprint-2-planning.md#historias-nuevas-la-interfaz-web), una
historia sin estimar no suma al compromiso hasta que el equipo la estima.

| # | Historia | SP | Responsable | Horas estimadas | Estado al 26/10 | Lo que dijo el responsable |
|---|---|---|---|---|---|---|
| 10 | Cerrar un sprint | 5 | Emiliano Sorato | 13,5 | **Hecho** (PR #150, 26/10) | "confirmo la #10" |
| 16 | Calcular las métricas | 8 | Emiliano Sorato | 21,6 | **Hecho** (PRs #81, #140 y #151, 26/10) | "confirmo [...] el resto de la #16" |
| 17 | Mostrar el dashboard | 8 | Juliana Bustos | 21,6 | En progreso | "confirmo la #17 (8 SP). Ya está en progreso: las tareas 1 a 3 de 17 están hechas en su rama" |
| 11 | Consultar sprints anteriores | 2 | Agustín Salinas | 5,4 | Lista para implementar | "confirmo la #11 (2 SP; puertas aprobadas por Juli en el #143; 6 de sus 12 tareas esperan a la #10)" |
| 18 | Generar el reporte | 5 | Agustín Salinas | 13,5 | Especificación SDD | "[confirmo] la #18 (5 SP; espera a la #16 y reutiliza el `format.go` de la #17 para no duplicar la regla de formato)" |
| 78 | Operar todo desde la web | 13 (acordado el 26/10) | Emiliano Sorato | 35,1 | En progreso (3 partes en `dev`, 17 de 22 escenarios con test) | "confirmo [...] la #78" |

- **Horas estimadas de lo comprometido: 110,7** (41 SP a 2,7 h por punto). Cada uno las cargó
  en el tablero. Las 35,1 h de la #78 (13 × 2,7, la conversión del equipo) las cargó Emiliano
  Sorato, su responsable, el 26/10, cuando la historia ya tenía story points.
- **La #10 y la #16 ya están terminadas el 26/10**, el primer día del sprint. Sus horas reales
  son de Emiliano Sorato, que las hizo con asistencia de IA: 0,5 h en la #10 y 1 h en la #16.

### Historia adelantada

Por la decisión 11 del [plan](../plan-final-al-2-de-noviembre.md), Maximiliano Eula adelanta su
próxima historia. No suma a los 28 comprometidos: se cuenta aparte.

| # | Historia | SP | Responsable | Horas estimadas | Estado al 26/10 |
|---|---|---|---|---|---|
| 73 | Resolver diferencias y acordar la estimación | 5 | Maximiliano Eula | 13,5 | **Hecho** (PR #148, con el arreglo del defecto #156 en el #157) |

Maximiliano escribió: "adelanto la #73 (5 SP), que tiene sus tres puertas aprobadas (PR #142) y
la #13 ya en `dev`. Horas estimadas cargadas: 13,5 h". La #73 era del Sprint 4. Sus horas
reales son **3 h**, contra 13,5 estimadas (−10,5 h, −78%): las cargó Maximiliano Eula en el
tablero y lo dejó dicho el 26/10 [en el PR #169](https://github.com/emisorato1/seguimiento-medicion-ics/pull/169#issuecomment-6428626374).

### Historia agregada durante el sprint

| # | Historia | SP | Responsable | Horas estimadas | Estado al 26/10 |
|---|---|---|---|---|---|
| 177 | Guardar los datos en una base de datos (SQLite) | [pendiente] | Emiliano Sorato | sin horas | En progreso |

La agregó Emiliano Sorato el 26/10, ya empezado el sprint, y no estaba en el compromiso de los
28 story points. Sus story points salen del Planning Poker: Emiliano votó 8, y los demás todavía
no votaron [pendiente]. Hasta que tenga story points no tiene horas estimadas. Trae una
dependencia nueva para el proyecto: Maximiliano Eula dio el ok en persona a Emiliano el 26/10.

## Planning Poker de la #77 y la #78

Ninguna de las dos tenía story points. La #77 ya está terminada: se estima solo para la
velocidad, como propuso Maximiliano.

**Ronda 1** (publicada en la #147):

| | Agustín | Juliana | Maximiliano | Emiliano |
|---|---|---|---|---|
| #77 Web base | 5 | 5 | 13 | 13 |
| #78 Operaciones web | 13 | 8 | 13 | 13 |

Hubo diferencias, así que se explicaron los extremos:

- **Emiliano Sorato, 13 en la #77:** "son ~4.950 líneas en 3 PRs, con servidor, pantallas,
  guardado en archivo y seguridad de formularios; la #8 y la #9 juntas, 8 SP, fueron ~1.100".
- **Juliana Bustos, 5 en la #77:** "la estimé por la lógica nueva, no por el volumen. No agrega
  reglas de dominio: son formularios y tablas sobre operaciones que ya existían. Me quedaron
  afuera el servidor, el guardado en archivo y la seguridad de los formularios, que la #77 armó
  desde cero".
- **Juliana Bustos, 8 en la #78:** "la complejidad está en el Planning Poker completo (votos
  ocultos, revelado, rondas y el acuerdo de la #73), más conectar las pantallas de sprints,
  esfuerzo, defectos y reportes con el PDF. Es más que la #77, pero sigue apoyándose en la
  estructura que ella dejó armada".
- **Agustín Salinas, 5 en la #77:** "la estimé por la lógica nueva y no por el volumen de
  líneas. No agrega reglas de negocio: muestra y opera lo que el dominio ya resolvía. Leyendo
  la explicación del 13 (servidor, guardado en archivo y seguridad de formularios armados desde
  cero), lo reconsidero en la ronda 2".
- **Maximiliano Eula:** "mis 13 y 13 coinciden con los tuyos; la explicación del lado alto ya la
  diste".

**Ronda 2** (cada uno la mandó por privado a Emiliano, para que siguiera siendo secreta):

| | Agustín | Juliana | Maximiliano | Emiliano |
|---|---|---|---|---|
| #77 Web base | 13 | 13 | 13 | 13 |
| #78 Operaciones web | 13 | 13 | 13 | 13 |

Maximiliano Eula mandó su ronda 2 el 26/10, por WhatsApp a Emiliano. Mantiene los votos de
la ronda 1: "La #77 armó desde cero el servidor, el guardado en archivo y la seguridad de los
formularios, y la #78 suma cinco pantallas, con el Planning Poker completo, sobre esa base."

Juliana Bustos mandó su ronda 2 el 26/10, por WhatsApp a Emiliano. Cambió los dos votos:

> "#77: 13 SP. Escuchando los argumentos de la ronda 1 y habiendo revisado el alcance real
> (servidor HTTP, persistencia, seguridad de formularios y plantillas, que superaron las 4.900
> líneas), reconozco que 5 SP se quedaba corto al haber mirado inicialmente solo el dominio.
> #78: 13 SP. Viendo que solo la primera parte (PR #149) ya sumó más de 2.000 líneas en 25
> archivos, y que aún resta conectar defectos, reportes, gráficos SVG y generación de PDF, la
> carga de integración y operativa justifica los 13 SP."

**Estimación acordada: #77 = 13 y #78 = 13.** En la ronda 2 los cuatro votaron 13 y 13: hubo
consenso. Se cargó en el tablero el 26/10.

**Si no votan a tiempo** (decisión de Emiliano Sorato, del 26/10): si Juliana y Maximiliano no
mandan su ronda 2 antes del 28/10, queda el valor más alto. La
[guía de Planning Poker](../guia-planning-poker.md) dice que eso se hace "después de 3 rondas".
Aplicarlo después de la ronda 2 es una decisión de Emiliano Sorato. Con los votos que hay, eso daría 13 para la #77 y 13 para la #78.
Todavía no es la estimación acordada.

La regla no hizo falta: los cuatro votaron antes del 28/10 y coincidieron.

## Puertas y revisiones

- **Agustín Salinas:** "aprobé las de la #10, con dos ajustes (en la issue #10)".
- **Maximiliano Eula:** "las de la #78 y la #19 las aprobé el 22/10 en sus issues, con
  observaciones; falta registrarlas en las specs, que siguen 'pendientes de las tres puertas'.
  El #88 ya tiene la aprobación de Juli y se puede mergear".

Al 26/10 las puertas de la #78 y la #19 ya están registradas en sus specs, como aprobadas por
Maximiliano Eula el 22/10. La de la #19 entró con el PR #152.
El #88 ya se mergeó a `main`.

## Daily

Sigue **por el grupo a las 21 h**, una línea cada uno: qué cerré, qué sigo, si estoy trabado.
Los cuatro contestaron "ok".

## Pendientes

- [x] La ronda 2 del Planning Poker de Juliana, y la estimación acordada de la #77 y la #78:
  13 y 13, el 26/10.
- [x] Las horas estimadas de la #78: 35,1 h, cargadas por Emiliano Sorato el 26/10.
- [ ] El Planning Poker de la #177 (falta el voto de Juliana, Maximiliano y Agustín) y sus horas
  estimadas.
- [x] Las horas reales de la #73: 3 h, cargadas por Maximiliano Eula el 26/10.
