# Acta — Sprint Planning del Sprint 3

- **Fecha:** del 24 al 26/10/2026
- **Cómo se hizo:** por escrito, en la [issue #147](https://github.com/emisorato1/seguimiento-medicion-ics/issues/147), con un comentario de cada uno
- **Participantes:** Emiliano Sorato, Maximiliano Eula, Juliana Bustos y Agustín Salinas
- **Facilita:** Emiliano Sorato (Agile Enabler)
- **Duración del sprint:** del lunes 26 al viernes 30/10

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

Cada responsable confirmó sus historias en la #147. El compromiso son **28 story points**: los
36 del plan menos los 8 de la #13, que se terminó en el Sprint 2.

| # | Historia | SP | Responsable | Horas estimadas | Estado al 26/10 | Lo que dijo el responsable |
|---|---|---|---|---|---|---|
| 10 | Cerrar un sprint | 5 | Emiliano Sorato | 13,5 | **Hecho** (PR #150, 26/10) | "confirmo la #10" |
| 16 | Calcular las métricas | 8 | Emiliano Sorato | 21,6 | **Hecho** (PRs #81, #140 y #151, 26/10) | "confirmo [...] el resto de la #16" |
| 17 | Mostrar el dashboard | 8 | Juliana Bustos | 21,6 | En progreso | "confirmo la #17 (8 SP). Ya está en progreso: las tareas 1 a 3 de 17 están hechas en su rama" |
| 11 | Consultar sprints anteriores | 2 | Agustín Salinas | 5,4 | Lista para implementar | "confirmo la #11 (2 SP; puertas aprobadas por Juli en el #143; 6 de sus 12 tareas esperan a la #10)" |
| 18 | Generar el reporte | 5 | Agustín Salinas | 13,5 | Especificación SDD | "[confirmo] la #18 (5 SP; espera a la #16 y reutiliza el `format.go` de la #17 para no duplicar la regla de formato)" |
| 78 | Operar todo desde la web | sin estimar | Emiliano Sorato | sin horas | En progreso (3 partes en `dev`, 17 de 22 escenarios con test) | "confirmo [...] la #78" |

- **Horas estimadas de lo comprometido: 75,6** (28 SP a 2,7 h por punto). Cada uno las cargó en
  el tablero. La #78 queda sin horas hasta que se cierre su estimación.
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
reales están [pendientes: las carga Maximiliano Eula].

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
| #77 Web base | 13 | [pendiente] | [pendiente] | 13 |
| #78 Operaciones web | 13 | [pendiente] | [pendiente] | 13 |

**Estimación acordada:** [pendiente: se cierra cuando estén los votos de Juliana y
Maximiliano].

**Si no votan a tiempo** (decisión de Emiliano Sorato, del 26/10): si Juliana y Maximiliano no
mandan su ronda 2 antes del 28/10, se aplica la [guía de Planning Poker](../guia-planning-poker.md):
después de las rondas, queda el valor más alto. Con los votos que hay, eso daría 13 para la #77
y 13 para la #78. Todavía no es la estimación acordada.

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

- [ ] La ronda 2 del Planning Poker de Juliana y Maximiliano, y la estimación acordada de la #77
  y la #78.
- [ ] Las horas estimadas de la #78, cuando tenga story points.
- [ ] Las horas reales de la #73 (Maximiliano Eula).
