# Candidatas para el Sprint 2

Preparado por el Agile Enabler antes de la planning del 16/10, para que la reunión se use en
decidir y no en reconstruir las dependencias. **No fija el alcance**: eso se decide en la
planning, y recién cuando se sepa la velocidad real del Sprint 1.

## Lo que queda del Product Backlog

Doce historias, **64 story points** de los 80 totales.

## Qué se puede hacer en el Sprint 2

Suponiendo que las cinco historias del Sprint 1 terminen, estas quedan sin bloqueos:

| # | Historia | SP | Depende de | Nota |
|---|---|---|---|---|
| 5 | Consultar el estado de un proyecto | 3 | #3, #4, #6 | lectura, sin reglas nuevas |
| 7 | Priorizar el Product Backlog | 2 | #6 | ordenar, es de las más chicas |
| 8 | Crear un sprint con su Sprint Goal | 3 | #3 | abre toda la rama de sprints |
| 13 | Estimar con Planning Poker | 13 | #12, #4 | la más grande del backlog |
| 14 | Registrar el esfuerzo real | 3 | #4, #6 | necesario para las métricas |

Y detrás de la #8, en cadena: **#9** (asignar historias, 5) → **#10** (cerrar sprint, 5) →
**#11** (consultar sprints anteriores, 2) y **#16** (métricas, 8).

La **#15** (defectos, 5) necesita #6 y #8. La **#17** (dashboard, 8), la **#18** (reportes, 5)
y la **#19** (PDF, 5) necesitan que haya métricas, así que son de los sprints 3 y 4.

## El plan de la cátedra pone "interfaz usable" en el Sprint 2

Eso tira hacia dos cosas a la vez: seguir con funcionalidad (la rama de sprints) y mejorar la
interfaz más allá de la consola mínima de la #24. Conviene decidirlo en la planning, porque
son dos caminos distintos:

- **Camino funcionalidad:** #8, #9, #10 y #7. Deja la gestión de sprints completa y habilita
  las métricas para el Sprint 3. Son 15 story points.
- **Camino interfaz:** mejorar la consola o empezar la interfaz definitiva, más las historias
  chicas de lectura (#5, #7, #11). La contra es que posterga lo que habilita las métricas, y
  las métricas son lo que más pesa en los entregables finales.

Lo que recomiendo discutir: **la rama de sprints (#8, #9, #10) es el cuello de botella de
todo lo que viene**. Sin sprints cerrados no hay velocidad, sin velocidad no hay métricas, y
sin métricas no hay dashboard ni reportes ni informe final. Postergarla aprieta los sprints 3
y 4.

## Dos cosas para decidir en la planning

1. **La #13 (Planning Poker, 13 SP)** está marcada desde el Sprint 0 como candidata a
   partirse en dos: "votación con cartas ocultas y revelado" y "rondas sucesivas hasta
   acordar". Si entra al Sprint 2, conviene partirla antes de estimarla.
2. **Cuánto comprometer.** El Sprint 1 comprometió 16 story points sin velocidad histórica.
   Para el Sprint 2 ya va a haber un número real: **se compromete esa velocidad, no más**,
   aunque sobre tiempo en el papel.

## Antes de la planning

- Estimar con **Planning Poker de verdad**, con votos ocultos. En el Sprint 1 se estimó por
  afinidad y quedó el compromiso de hacerlo bien esta vez.
- Recalibrar la conversión de horas por story point con los datos reales del Sprint 1.
