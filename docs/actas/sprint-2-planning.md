# Acta — Sprint Planning del Sprint 2

> Se completa **durante** la reunión. Lo que está entre corchetes hay que reemplazarlo.

- **Fecha:** [19/10/2026]
- **Participantes:** [quiénes estuvieron]
- **Facilita:** Emiliano Sorato (Agile Enabler)

## Lo que traemos del Sprint 1

- **Velocidad medida: 11 story points.** El Sprint 2 compromete esa velocidad, no más,
  aunque sobre tiempo en el papel.
- **Conversión recalibrada: 2,6 horas por story point completado** (29 h de las historias
  terminadas ÷ 11 SP), contra las 3 que se usaron para estimar el Sprint 1: se estimó de más.
- Los números completos están en el [acta de review](sprint-1-review.md).

## Historias confirmadas antes de la planning

El equipo confirmó por chat el **19/10**, antes de la reunión, que estas dos historias
entran al Sprint 2:

| # | Historia | SP | Estado al confirmarse |
|---|---|---|---|
| 12 | Estimar con Story Points | 2 | las tres puertas aprobadas, 10 de 13 tareas hechas |
| 24 | Interfaz de consola | 3 | las tres puertas aprobadas, 7 de 15 tareas hechas |

**Por qué antes y no en la reunión.** Las dos vuelven del Sprint 1 con sus tres puertas
aprobadas y a medio hacer —la #12 con 10 de 13 tareas y la #24 con 7 de 15—: no había forma
de que no entraran, y esperar a la reunión dejaba a su responsable sin poder avanzar durante
todo el día. La regla de no programar fuera del sprint existe para que nadie trabaje en algo
que el equipo no comprometió, y acá ese riesgo no existía. Quedó escrito como regla en
`CLAUDE.md`, con la aclaración de que la confirmación la da el equipo y no el Agile Enabler
por su cuenta.

## Alcance del Sprint 2

**Decidido: opción A del replan** (profundidad: la cadena de sprints y métricas). El
razonamiento y la opción descartada están en
[docs/replan-al-2-de-noviembre.md](../replan-al-2-de-noviembre.md).

Son **dos semanas, del 20/10 al 31/10**, porque las presentaciones arrancan el 2/11. El
presupuesto real son 30 story points: 15 por semana, calculados con el costo medido del
Sprint 1 (2,6 h por punto) y 10 horas por persona por semana.

### Semana 1 — del 20 al 24 de octubre

| # | Historia | SP | Quién |
|---|---|---|---|
| 12 | Estimar con Story Points | 2 | Maximiliano Eula |
| 24 | Interfaz de consola | 3 | Maximiliano Eula |
| 8 | **Crear un sprint** | 3 | Agustín Salinas |
| 9 | **Asignar historias a un sprint** | 5 | Agustín Salinas |

**La #8 y la #9 van en el mismo pull request**: son eslabones seguidos de la cadena y la #9
no tiene sentido sin la #8.

### Semana 2 — del 27 al 31 de octubre

| # | Historia | SP | Quién |
|---|---|---|---|
| 10 | **Cerrar un sprint** | 5 | Juliana Bustos |
| 14 | Registrar el esfuerzo real | 3 | Juliana Bustos |
| 16 | **Calcular las métricas** | 8 | Emiliano Sorato |

- **Total comprometido: 29 story points.**
- **Reparto:** Maximiliano 5, Agustín 8, Juliana 8, Emiliano 8. En el Sprint 1 uno quedó con
  dos historias y el resto con una; esta vez está parejo.

### El camino crítico

```
#8 → #9 → #10 → #16
```

Son **21 de los 29 puntos en serie**: cada una necesita que la anterior esté en `dev`. Si un
eslabón se atrasa un día, se atrasa todo lo que viene detrás y las métricas quedan afuera.

La cadena **tiene prioridad sobre cualquier otra cosa**: si alguien se libera, va ahí aunque
su propia historia no esté terminada.

### Lo que queda afuera

Gestión de defectos (#15), dashboard (#17), reportes (#18 y #19), Planning Poker (#13),
consultar el estado de un proyecto (#5), priorizar el backlog (#7) y consultar sprints
anteriores (#11): **43 story points**, más de la mitad del backlog.

No se ocultan. Se presentan como backlog pendiente, con su estimación y sus dependencias, que
es lo que un Product Backlog tiene que mostrar al cierre de un proyecto.

### Entregables que no son historias

Van **en paralelo**, no después del código:

| Entregable | Quién | Cuándo |
|---|---|---|
| Informe de métricas y cobertura | Emiliano | se arma con `scripts/metricas.sh` desde ya |
| Manual breve de usuario | Maximiliano | crece con cada tarea de la consola |
| Documentación técnica | sale de las specs, que ya están | revisar al final |
| Actas de retrospectivas | Emiliano | ya están las del Sprint 0 y 1 |
| Presentación final | los cuatro | semana del 27 |

## Sprint Goal

> Poder planificar y cerrar un sprint, registrar el esfuerzo, y que el sistema calcule las
> métricas del proyecto a partir de eso.

## Acuerdos del equipo

- **Definición de Terminado:** sin cambios respecto del Sprint 1.
- **Reparto:** en el Sprint 1 un integrante quedó con dos historias y el resto con una. Esta
  vez la #12 y la #24 son de Maximiliano porque ya las empezó, y no se le asigna nada más.
- **Daily:** asincrónica por chat de lunes a viernes. Con dos semanas de plazo, el que se traba avisa el mismo día y no en la daily siguiente.
- **Fecha de cierre, review y retrospectiva:** viernes 31/10. Las presentaciones arrancan el 2/11.

## Pendientes

- [ ] Hacer la retrospectiva del Sprint 1 antes de esta planning.
- [ ] Estimar con Planning Poker de verdad las historias nuevas, con votos ocultos: en el
      Sprint 1 se estimó por afinidad y quedó el compromiso de hacerlo bien esta vez.
- [ ] Decidir si la #13 (13 SP) se parte en dos antes de estimarla.

---

## Anexo — cambio de alcance decidido el 20/10

Después de esta planning, el Agile Enabler llevó al equipo la decisión de **entregar el
alcance completo** en lugar de los 29 story points que se habían comprometido arriba. Lo que
está más arriba en esta acta queda como lo que se acordó en la reunión; esto es lo que
cambió el mismo día.

### Qué cambia

| | En la planning | Después del cambio |
|---|---|---|
| Alcance | 29 SP, 43 quedaban como backlog | **72 SP, nada queda afuera** |
| Sprints hasta el 2/11 | uno (20 al 31/10) | **tres**: Sprint 2 (20–25/10), Sprint 3 (26–30/10), Sprint 4 (31/10–1/11) |
| Ritmo pedido | ~10 h por persona por semana | **~3,6 h por persona por día**, fines de semana incluidos |
| Historias por persona | una o dos | cuatro o cinco, repartidas 17/19/18/18 SP |

### Por qué

Se presenta el **2 de noviembre** y el equipo prefiere defender el producto entero antes que
defender un recorte. Los cuatro trabajan con Claude Code, así que el costo por story point
puede bajar bastante respecto de las 2,6 horas medidas en el Sprint 1 —ese número incluía
cuatro o cinco pull requests por historia y revisiones que tardaban un día— y el plan
reorganiza el trabajo justamente para atacar eso.

### Qué se resolvió de los pendientes de arriba

- **La #13 se parte en dos**, como quedó pendiente de decidir: queda la #13 *Votar en secreto
  en una ronda* (8 SP) y la nueva #73 *Resolver diferencias y acordar la estimación* (5 SP).
  Eran 13 puntos en una sola historia: demasiado para un pull request y para un solo
  responsable con este plazo.
- **Las estimaciones no se rehicieron con Planning Poker.** Se mantienen las que ya estaban,
  salvo el reparto de la #13 entre sus dos mitades. Con el plazo encima no había margen para
  una sesión de estimación, y queda anotado como desvío: sigue siendo el compromiso
  pendiente desde el Sprint 1.

### Lo que no cambia

La Definición de Terminado. Spec antes del código, escenarios BDD, tests que pasan, revisión
escrita y cobertura. **Se recorta ceremonia, no calidad.**

Si el ritmo no da, el **domingo 25 a la noche** se recorta con el criterio escrito en
[el plan](../plan-final-al-2-de-noviembre.md), y se avisa ese día.

El plan completo —calendario, reparto por persona, dependencias y riesgos— está en
[docs/plan-final-al-2-de-noviembre.md](../plan-final-al-2-de-noviembre.md).

### Historias nuevas: la interfaz web

El mismo 20/10, día de la planning, Emiliano sumó dos historias al tablero:

| # | Historia | Sprint | SP |
|---|---|---|---|
| 77 | Usar el sistema desde el navegador: proyectos, integrantes y backlog | Sprint 2 | sin estimar |
| 78 | Operar sprints, esfuerzo, defectos, Planning Poker y reportes desde el navegador | Sprint 3 | sin estimar |

**Por qué.** La consigna pide una interfaz usable y un dashboard con gráficos, y la consola
sola no alcanza para eso.

**Están sin estimar.** No tienen story points ni horas estimadas cargadas. Las estima el
equipo por chat y, hasta entonces, no suman puntos ni horas al compromiso del sprint.

### La #6 cuenta en este sprint

_Agregado el 20/10._ La #6 *Administrar el Product Backlog* (5 SP) estaba aprobada al cierre
del Sprint 1, pero se mergeó a `dev` recién el 19/10, y la Definición de Terminado exige el
merge. Por eso no cuenta en el Sprint 1 y entra en este sprint, igual que la #12: una
historia cuenta en el sprint en que se mergea, y lo que se mergea entre un sprint y el
siguiente, como la #6 el 19/10, cuenta en el siguiente. Las fechas están en el
[acta de review del Sprint 1](sprint-1-review.md).

- **El compromiso del Sprint 2 queda en 31 story points**: los 26 del plan más los 5 de la
  #6, que llega **ya terminada** (se mergeó el 19/10).
- Con la #6 afuera del Sprint 1, **la velocidad del Sprint 1 pasa de 11 a 6** y la conversión
  de 2,6 a **2,7 horas por story point** (16 h ÷ 6 SP). Lo que está más arriba en esta acta
  usa el 11 y el 2,6, que eran los números de ese momento.

### Horas estimadas del Sprint 2

Se cargaron en el tablero el 20/10, antes de empezar cada historia, con la conversión de
**2,7 horas por story point** que sale del Sprint 1. No se usó el 2,0 del plan porque es una
meta, no algo medido. Primero se cargaron a 2,6 y el mismo día se corrigieron a 2,7, cuando
la #6 salió del Sprint 1.

| # | Historia | SP | Horas estimadas |
|---|---|---|---|
| 5 | Consultar el estado de un proyecto | 3 | 8,1 |
| 7 | Priorizar el Product Backlog | 2 | 5,4 |
| 8 | Crear un sprint con su Sprint Goal | 3 | 8,1 |
| 9 | Asignar historias a un sprint | 5 | 13,5 |
| 14 | Registrar el esfuerzo real | 3 | 8,1 |
| 15 | Gestionar defectos | 5 | 13,5 |
| 12 | Estimar con Story Points | 2 | 6 |
| 24 | Interfaz de consola | 3 | 9 |
| 6 | Administrar el Product Backlog | 5 | 15 |
| | **Total** | **31** | **86,7** |

La #12, la #24 y la #6 conservan las horas que se les cargaron en el Sprint 1, a 3 horas por
punto: ya estaban estimadas antes de empezar y cambiarlas ahora no sería una estimación
honesta. Las seis historias nuevas suman 56,7 h.

**Cada responsable puede ajustar las horas de su historia al empezarla**, antes de cargar
horas reales. La #77 queda sin horas hasta que tenga story points.

### Confirmación del equipo

El 20/10 el equipo confirmó por el grupo las dos decisiones de este anexo: **entregar el
alcance completo** con el plan al 2 de noviembre, y **sumar la interfaz web** (#77 y #78).

- **Confirmaron:** Agustín Salinas, Juliana Bustos y Maximiliano Eula, los tres Product
  Builders. Lo registra Emiliano Sorato, que llevó la propuesta al grupo.
