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
