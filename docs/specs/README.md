# Especificaciones SDD

Una carpeta por sprint, y dentro un documento por historia:

```
docs/specs/
├── PLANTILLA.md     se copia para cada spec nueva
├── sprint-0/        tareas de configuracion (sin flujo de cuatro fases)
└── sprint-1/        una spec por historia comprometida
```

## Estado de las specs

### Sprint 0 — terminado

Son tareas de configuración y planificación, no funcionalidad, así que documentan
**decisiones y evidencia** en lugar de seguir el flujo de cuatro fases. El motivo está en el
[README del Sprint 0](sprint-0/README.md).

| Documento | Ítem | Estado | Hecho por |
|---|---|---|---|
| [Configurar el repositorio y el tablero](sprint-0/01-configurar-repositorio-y-tablero.md) | #1 | ✅ terminado | Emiliano Sorato |
| [Definir el Product Backlog inicial](sprint-0/02-definir-product-backlog-inicial.md) | #2 | ✅ terminado | Emiliano Sorato |

### Sprint 1 — cerrado el 16/10

| Spec | Historia | SP | Responsable | Estado |
|---|---|---|---|---|
| [Crear y modificar proyectos](sprint-1/03-crear-proyectos.md) | #3 | 3 | Agustín Salinas | ✅ Terminada — mergeada a `dev` |
| [Registrar integrantes](sprint-1/04-registrar-integrantes.md) | #4 | 3 | Juliana Bustos | ✅ Terminada — mergeada a `dev` |
| [Administrar el Product Backlog](sprint-1/06-administrar-product-backlog.md) | #6 | 5 | Emiliano Sorato | ✅ Terminada — mergeada a `dev` |
| [Estimar con Story Points](sprint-1/12-estimar-story-points.md) | #12 | 2 | Maximiliano Eula | 🔍 Implementada — en review (pasa al Sprint 2) |
| [Interfaz de consola](sprint-1/24-interfaz-de-consola.md) | #24 | 3 | Maximiliano Eula | ⚙️ En implementación (pasa al Sprint 2) |

Las Fases 1 se redactaron como borrador con asistencia de IA a partir de la consigna y de los
criterios de aceptación de cada issue, y cada responsable las revisó y corrigió antes de
presentarlas a la Puerta 1.

**Solo se escribe código de una historia que tiene las tres puertas aprobadas.** El orden es
Puerta 1 (requisitos) → Fase 2 y Puerta 2 (diseño) → Fase 3 y Puerta 3 (tareas) → recién ahí
la implementación.

**El Sprint 1 cerró con 11 de los 16 story points comprometidos** (velocidad 11). Las tres
historias terminadas están en `dev` con el dominio al 100% de cobertura. Los números
completos están en el [acta de review](../actas/sprint-1-review.md).

La **#12** y la **#24** vuelven al Sprint 2 con sus tres puertas aprobadas y a medio hacer:
no hay que rehacer sus specs, y conservan su carpeta `sprint-1/`, que es donde se
comprometieron.

### Sprint 2 — del 20 al 25 de octubre (26 SP)

| Spec | Historia | SP | Responsable | Estado |
|---|---|---|---|---|
| [Priorizar el Product Backlog](sprint-2/07-priorizar-backlog.md) | #7 | 2 | Agustín Salinas | 📝 Requisitos — pendiente de Puerta 1 |
| [Crear un sprint con su Sprint Goal](sprint-2/08-crear-sprint.md) | #8 | 3 | Agustín Salinas | 📝 Requisitos — pendiente de Puerta 1 |
| Asignar historias a un sprint | #9 | 5 | Agustín Salinas | ⏳ spec pendiente |
| Registrar el esfuerzo real | #14 | 3 | Juliana Bustos | ⏳ spec pendiente |
| Gestionar defectos | #15 | 5 | Juliana Bustos | ⏳ spec pendiente |
| [Estimar con Story Points](sprint-1/12-estimar-story-points.md) | #12 | 2 | Maximiliano Eula | 🔍 En review |
| [Interfaz de consola](sprint-1/24-interfaz-de-consola.md) | #24 | 3 | Maximiliano Eula | ⚙️ En implementación |
| Consultar el estado de un proyecto | #5 | 3 | Maximiliano Eula | ⏳ spec pendiente |

### Sprint 3 — del 26 al 30 de octubre (36 SP)

| Spec | Historia | SP | Responsable | Estado |
|---|---|---|---|---|
| Cerrar un sprint | #10 | 5 | Emiliano Sorato | ⏳ spec pendiente |
| Calcular las métricas del proyecto | #16 | 8 | Emiliano Sorato | ⏳ spec pendiente |
| Mostrar el dashboard | #17 | 8 | Juliana Bustos | ⏳ spec pendiente |
| Votar en secreto en una ronda | #13 | 8 | Maximiliano Eula | ⏳ spec pendiente |
| Consultar sprints anteriores | #11 | 2 | Agustín Salinas | ⏳ spec pendiente |
| Generar el reporte de proyecto o sprint | #18 | 5 | Agustín Salinas | ⏳ spec pendiente |

### Sprint 4 — 31 de octubre y 1 de noviembre (10 SP)

| Spec | Historia | SP | Responsable | Estado |
|---|---|---|---|---|
| Exportar el reporte a PDF | #19 | 5 | Emiliano Sorato | ⏳ spec pendiente |
| Resolver diferencias y acordar la estimación | #20 | 5 | Maximiliano Eula | ⏳ spec pendiente |

Las Fases 1 son borradores preparados con asistencia de IA, igual que en el Sprint 1: cada
una tiene su sección "Para decidir en la puerta" con las decisiones ambiguas que hay que
cerrar en la Puerta 1, y cada responsable la revisa y corrige antes de presentarla.

**El plan completo —calendario, reparto y riesgos— está en
[docs/plan-final-al-2-de-noviembre.md](../plan-final-al-2-de-noviembre.md).**

## Orden de trabajo hasta el 2 de noviembre

Las historias que quedan están encadenadas:

```
#8 → #9 → #10 → #16 ─┬─→ #17 Dashboard
                     └─→ #18 → #19 PDF

#13 → #20            (Planning Poker, cadena aparte)
#5, #7, #14, #15, #24  (sin dependencias, se hacen cuando sea)
```

Son 39 de los 72 puntos en serie. **Si cada uno espera a que la anterior esté en `dev`, el
dashboard y los reportes no llegan.**

Por eso, desde el 20/10, **la cadena no se espera: se trabaja contra el contrato**. Las
entidades compartidas están definidas en [docs/diseno-dominio.md](../diseno-dominio.md)
antes de que nadie las implemente, y quien consume una que todavía no existe arma los datos
a mano en su propio test. Cuando la pieza real aterriza en `dev`, se conectan.

El contrato se cambia avisando al grupo **antes**, no después: un cambio silencioso rompe el
trabajo de otra persona que ya está escrito.


## Reglas

- El archivo se llama `NN-nombre-corto.md`, donde `NN` es el número de la issue.
- La spec va en la carpeta del sprint donde la historia **se comprometió**.
- Si una historia no se termina y vuelve al Product Backlog, cuando se vuelve a comprometer
  se **mueve** la spec con `git mv` a la carpeta del sprint nuevo: una historia tiene una
  sola spec, con su historial completo.
- En Requisitos no va nada técnico. En Diseño van decisiones, no código.
- Las specs se actualizan en el mismo pull request que las vuelve falsas.
