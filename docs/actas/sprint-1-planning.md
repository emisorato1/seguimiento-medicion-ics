# Acta — Sprint Planning del Sprint 1

- **Fecha:** 2026-09-28
- **Participantes:** Emiliano Sorato, Agustín Salinas, Juliana Bustos, Maximiliano Eula
- **Facilita:** Emiliano Sorato (Agile Enabler)

## Cómo estimamos

El equipo estimó **por afinidad**: se partió de una propuesta de valores preparada con
asistencia de IA, y el equipo la recorrió historia por historia, la discutió y la aprobó en
esta reunión. No se usó Planning Poker con votación oculta en esta planning.

Se acordó que en la planning del Sprint 2 se va a estimar con Planning Poker propiamente
dicho, al menos para las historias en las que no haya acuerdo inmediato. Sirve además para
entender de primera mano la historia #13, que es implementar ese mecanismo.

**Historia de referencia:** #4 — Registrar integrantes de un proyecto = **3 story points**.
Todo lo demás se estimó comparando contra esa.

## Estimaciones

| # | Historia | Prioridad | Story Points | Criterio |
|---|---|---|---|---|
| 4 | Registrar integrantes de un proyecto | Alta | 3 | referencia |
| 3 | Crear y modificar proyectos | Alta | 3 | igual que la referencia, más la validación de fechas |
| 5 | Consultar el estado de un proyecto | Media | 3 | lectura, pero toca varias entidades |
| 6 | Administrar el Product Backlog | Alta | 5 | la entidad más rica: siete campos, ID único, validaciones |
| 7 | Priorizar el Product Backlog | Media | 2 | ordenar una lista |
| 8 | Crear un sprint con su Sprint Goal | Alta | 3 | como #3, con la regla de un solo sprint abierto |
| 9 | Asignar historias a un sprint | Alta | 5 | reglas cruzadas entre historia y sprint |
| 10 | Cerrar un sprint | Alta | 5 | efectos en cascada: devolver historias, calcular SP, congelar |
| 11 | Consultar sprints anteriores | Media | 2 | lectura de datos ya calculados |
| 12 | Estimar historias con Story Points | Alta | 2 | validar que el valor esté en Fibonacci |
| 13 | Estimar con Planning Poker | Alta | 13 | rondas, votos ocultos, detección de diferencias, historial |
| 14 | Registrar el esfuerzo real | Alta | 3 | alta simple con validaciones de fecha y horas |
| 15 | Gestionar defectos | Media | 5 | entidad nueva más la regla de sprint de resolución |
| 16 | Calcular las métricas del proyecto | Alta | 8 | nueve cálculos, cada uno con su caso sin datos |
| 17 | Mostrar el dashboard | Media | 8 | depende de la interfaz y de que las métricas existan |
| 18 | Generar el reporte de proyecto o sprint | Media | 5 | integra todo lo anterior |
| 19 | Exportar el reporte a PDF | Baja | 5 | librería externa, riesgo de complicarse |

**Total del Product Backlog: 80 story points.**

Observación registrada: la historia **#13 (13 SP)** es candidata a partirse en dos —
"votación con cartas ocultas y revelado" y "rondas sucesivas hasta acordar"— cuando se
acerque el momento de implementarla.

## Sprint Goal

> Poder crear un proyecto con sus integrantes, cargarle el Product Backlog y estimar las
> historias en story points.

## Sprint Backlog

Es el primer sprint y no hay velocidad histórica, así que el compromiso es deliberadamente
conservador: se eligieron las historias que no dependen de ninguna otra, una por integrante.

| # | Historia | Story Points | Quién la toma |
|---|---|---|---|
| 6 | Administrar el Product Backlog | 5 | Emiliano Sorato |
| 3 | Crear y modificar proyectos | 3 | Agustín Salinas |
| 4 | Registrar integrantes de un proyecto | 3 | Juliana Bustos |
| 12 | Estimar historias con Story Points | 2 | Maximiliano Eula |

- **Total comprometido: 13 story points** de los 80 del backlog.
- **Historias que quedaron afuera:** las de gestión de sprints (#8 a #11) dependen de que
  exista el Product Backlog, y las de métricas, dashboard y reportes (#16 a #19) dependen de
  que haya datos cargados. Entran en los sprints siguientes.

## Ajustes posteriores a la planning

**2026-09-29.** El equipo sumó al Sprint 1 la historia **#24 — Operar el sistema desde una
interfaz de consola (3 SP, prioridad Media)**.

Motivo: el plan de la cátedra pone "interfaz usable" recién en el Sprint 2, así que sin una
consola mínima lo único que se podría mostrar en la review del Sprint 1 serían tests
pasando. La consola no tiene reglas de negocio propias: solo llama al núcleo en Go.

Depende de #3, #4, #6 y #12, así que se hace al final del sprint.

**Total comprometido actualizado: 16 story points.**

## Acuerdos del equipo

- **Definición de Terminado:** una historia está terminada cuando tiene su especificación SDD
  escrita antes del código, sus escenarios BDD, tests en Go que pasan, el código revisado por
  otro integrante en un pull request y está mergeada a `dev`.
- **Revisión cruzada:** ningún pull request se mergea sin la aprobación de otro integrante.
  No se usa el bypass de administrador.
- **Daily:** asincrónica por el grupo de chat, de lunes a viernes. Cada uno responde antes de
  las 22:00 qué hizo, qué va a hacer y qué lo está trabando. Además, una sincrónica corta los
  **miércoles 21:00**, de 15 minutos.
- **Cierre del Sprint 1, review y retrospectiva:** **viernes 16/10/2026**.

## Calendario del sprint

Definido por el Agile Enabler el **02/10** porque había quedado pendiente desde la planning.
Sin fecha de cierre no hay velocidad, y sin velocidad no se puede planificar el Sprint 2.

| Fecha | Qué |
|---|---|
| lunes 05/10 | todas las Fases 2 escritas y presentadas a la Puerta 2 |
| miércoles 07/10 | todas las Puertas 2 aprobadas y las Fases 3 presentadas a la Puerta 3 |
| jueves 08/10 | todas las Puertas 3 aprobadas — **a partir de acá se programa** |
| viernes 16/10 | cierre del sprint, review y retrospectiva |

El sprint arrancó el 28/09, así que son tres semanas en vez de dos. Es más de lo que debería
durar, pero la primera semana se fue íntegra en specs y es el primer sprint del equipo. El
Sprint 2 vuelve a dos semanas.

**Se evaluó acortarlo a dos semanas (cierre el 09/10) y se descartó.** De los diez días
hábiles, siete ya se habían ido en specs y todavía faltaban las Puertas 2 y 3 de cuatro
historias: quedaban cuatro días de código, unas 32 horas entre los cuatro, que a 3 horas por
story point dan unos 10 de los 16 comprometidos. Como en Scrum el sprint es una caja de
tiempo fija y lo que se ajusta es el alcance, la alternativa real era sacar dos historias.
Se eligió mantener las cinco y estirar el cierre una semana para que **los cuatro integrantes
tengan su historia en este sprint**, que era uno de los acuerdos de la retrospectiva del
Sprint 0. El Sprint 2 vuelve a dos semanas, y ahí el alcance se ajusta a la velocidad medida.

**Si alguien no está de acuerdo con estas fechas, se habla en la daily del lunes 05/10.**
Pasada esa fecha quedan firmes.

## Asignación de la historia #24

La **#24 (Interfaz de consola, 3 SP)** había quedado "a definir" para quien terminara antes.
Como nadie terminó, la asigna el Agile Enabler el 02/10 por carga de trabajo:

| Integrante | Story points antes | Con la #24 |
|---|---|---|
| Emiliano Sorato | 5 | 5 |
| Agustín Salinas | 3 | 3 |
| Juliana Bustos | 3 | 3 |
| **Maximiliano Eula** | **2** | **5** |

Queda para **Maximiliano Eula**, que es quien tiene la carga más baja. La #24 depende de las
otras cuatro, así que la arranca cuando las demás estén en `dev`.

## Pendientes

- [x] Definir día y horario de la daily, y la fecha de la review.
- [x] Asignar responsable a la historia #24.
- [x] Cada integrante escribe la Fase 1 (Requisitos) de la spec de su historia.
- [ ] Cada integrante escribe la Fase 2 (Diseño) y la presenta a la Puerta 2.
