# Replanificación: las presentaciones empiezan el 2 de noviembre

> **Decisión superada el 20/10/2026.** Este documento analizaba recortar el alcance y
> ofrecía dos opciones. El Agile Enabler decidió lo contrario: entregar el proyecto
> completo. El plan vigente es
> [plan-final-al-2-de-noviembre.md](plan-final-al-2-de-noviembre.md).
>
> Se conserva porque el análisis —el costo por punto, la cobertura de los nueve
> requerimientos y el camino crítico— sigue siendo la base del plan nuevo, y porque la
> opción descartada es parte de cómo se tomó la decisión.

Preparado por el Agile Enabler el **20/10/2026**, después de que la cátedra anunciara en
clase que las presentaciones arrancan el **2 de noviembre**.

## El problema, en números

| | |
|---|---|
| Backlog total | 83 story points |
| Terminados | 11 (13%) |
| **Quedan** | **72** |
| Semanas hasta el 2/11 | **2** |
| Necesario para terminar todo | **36 story points por semana** |
| Velocidad medida en el Sprint 1 | **5,5 por semana** |

**Terminar las 15 historias que quedan no es posible.** No con más horas, no con mejor
proceso, no con menos reuniones: son seis veces y media lo que el equipo produjo en tres
semanas.

La decisión ya no es *cómo llegamos* sino **qué entregamos**.

## Lo que sí entra

El Sprint 1 costó **2,6 horas por story point** contando solo las historias terminadas. Con
4 personas a 10 horas por semana, eso da unos **15 story points por semana**: **30 en las
dos semanas**.

Ese es el presupuesto real. Todo lo que se planifique por encima de 30 puntos es ilusión.

## Qué priorizar: cobertura de requerimientos, no puntos

La consigna se evalúa por los **nueve requerimientos mínimos**, no por story points. Conviene
tocar los nueve aunque sea de forma mínima, antes que terminar algunos y dejar otros en cero.

Estado actual de la cobertura:

| # | Requerimiento | Estado | Qué falta |
|---|---|---|---|
| 1 | Gestión de proyectos | ✅ casi | #5 consultar estado (3) |
| 2 | Product Backlog | ✅ casi | #7 priorizar (2) |
| 3 | Gestión de Sprints | ❌ | #8 (3), #9 (5), #10 (5), #11 (2) |
| 4 | Estimación y Planning Poker | ⚠️ parcial | #12 (2, casi lista), #13 Planning Poker (13) |
| 5 | Registro de esfuerzo | ❌ | #14 (3) |
| 6 | Gestión de defectos | ❌ | #15 (5) |
| 7 | Métricas | ❌ | #16 (8) |
| 8 | Dashboard | ❌ | #17 (8) |
| 9 | Reportes | ❌ | #18 (5), #19 (5) |

## Las dos opciones, con sus números

### Opción A — profundidad: la cadena de sprints y métricas

| Semana | Historias | SP |
|---|---|---|
| 1 (20–24/10) | #12 (2), #24 (3), #8 (3), #9 (5) | 13 |
| 2 (27–31/10) | #10 (5), #14 (3), #16 (8) | 16 |
| | **Total** | **29** |

Cubre los requerimientos 3, 5 y 7, y deja el sistema demostrable de punta a punta: se crea un
proyecto, se carga el backlog, se planifica un sprint, se cierra y salen métricas reales.

**Queda en cero:** defectos (6), dashboard (8), reportes (9) y Planning Poker.

### Opción B — amplitud: tocar los nueve requerimientos

| Semana | Historias | SP |
|---|---|---|
| 1 (20–24/10) | #12 (2), #24 (3), #8 (3), #14 (3), #5 (3) | 14 |
| 2 (27–31/10) | #9 (5), #15 (5), #7 (2), #11 (2) | 14 |
| | **Total** | **28** |

Cubre los requerimientos 1, 2, 3 (parcial), 4 (parcial), 5 y 6.

**Queda en cero:** métricas (7), dashboard (8) y reportes (9) — y sin métricas no hay informe
final, que es un entregable.

### Por qué recomiendo la A

Las métricas son el **corazón del producto**: la aplicación se llama "seguimiento y medición".
Entregar un sistema que gestiona proyectos pero no mide nada deja afuera lo que le da sentido.
Además el informe de métricas es entregable, y sin la #16 no hay de dónde sacarlo.

Y hay un argumento de dependencias: la cadena `#8 → #9 → #10 → #16` es lo único que **no se
puede hacer después en paralelo**. Las historias de la opción B son casi todas independientes
y se pueden sumar en cualquier momento.

## Lo que no está en ninguna de las dos opciones

- Informe de métricas y cobertura
- Documentación técnica y manual breve de usuario
- Actas de retrospectivas
- La presentación final

**Son entregables y hay que hacerlos igual.** Con dos semanas, hay que empezarlos ya y en
paralelo, no después del código.

## Lo que hay que hacer esta semana, y es lo más urgente

1. **Preguntarle al profesor qué día presenta el grupo.** Las presentaciones *arrancan* el
   2/11, pero si son varios días y al grupo le toca el 9 o el 16, el plan cambia por completo.
   Es la pregunta que más información aporta y cuesta un mensaje.
2. **Avisarle a la cátedra que el alcance completo no entra**, con estos números y las dos
   opciones. Es mejor que lo sepan ahora a que lo descubran en la presentación.
3. **Decidir A o B como equipo**, y comprometer 30 puntos, no más.

## Cómo trabajamos estas dos semanas

Los cambios de proceso siguen valiendo, y ahora son obligatorios:

- **Una historia, un pull request**: spec, tests y código juntos, las tres puertas en una
  revisión.
- **Fases 2 y 3 fusionadas** en las historias de hasta 3 puntos.
- **Revisión en el día.** Un PR parado un día se come el 10% del plazo.
- **La Definición de Terminado no se toca.** Tests, revisión cruzada y cobertura siguen igual:
  lo que se recorta es alcance y ceremonia, nunca calidad. Entregar menos funcionalidad bien
  hecha se defiende; entregar más a medias, no.
