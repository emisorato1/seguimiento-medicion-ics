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

### Sprint 1 — en curso

| Spec | Historia | SP | Responsable | Estado |
|---|---|---|---|---|
| [Crear y modificar proyectos](sprint-1/03-crear-proyectos.md) | #3 | 3 | Agustín Salinas | ✅ Terminada — mergeada a `dev` |
| [Registrar integrantes](sprint-1/04-registrar-integrantes.md) | #4 | 3 | Juliana Bustos | ✅ Terminada — mergeada a `dev` |
| [Administrar el Product Backlog](sprint-1/06-administrar-product-backlog.md) | #6 | 5 | Emiliano Sorato | 🔍 Implementada — en review |
| [Estimar con Story Points](sprint-1/12-estimar-story-points.md) | #12 | 2 | Maximiliano Eula | 🟢 Puerta 3 aprobada — lista para implementar |
| [Interfaz de consola](sprint-1/24-interfaz-de-consola.md) | #24 | 3 | Maximiliano Eula | ⚙️ En implementación |

Las Fases 1 se redactaron como borrador con asistencia de IA a partir de la consigna y de los
criterios de aceptación de cada issue, y cada responsable las revisó y corrigió antes de
presentarlas a la Puerta 1.

**Solo se escribe código de una historia que tiene las tres puertas aprobadas.** El orden es
Puerta 1 (requisitos) → Fase 2 y Puerta 2 (diseño) → Fase 3 y Puerta 3 (tareas) → recién ahí
la implementación.

### Sprint 2 — del 20 al 24 de octubre

| Spec | Historia | SP | Responsable | Estado |
|---|---|---|---|---|
| [Priorizar el Product Backlog](sprint-2/07-priorizar-backlog.md) | #7 | 2 | a definir | 📝 Requisitos — pendiente de Puerta 1 |
| [Crear un sprint con su Sprint Goal](sprint-2/08-crear-sprint.md) | #8 | 3 | a definir | 📝 Requisitos — pendiente de Puerta 1 |
| [Estimar con Story Points](sprint-1/12-estimar-story-points.md) | #12 | 2 | Maximiliano Eula | ⚙️ En implementación — viene del Sprint 1 |
| [Interfaz de consola](sprint-1/24-interfaz-de-consola.md) | #24 | 3 | Maximiliano Eula | ⚙️ En implementación — viene del Sprint 1 |

La #12 y la #24 conservan su spec en `sprint-1/`, que es el sprint en el que se
comprometieron. Las dos tienen sus tres puertas aprobadas.

Las Fases 1 de la #7 y la #8 son borradores preparados con asistencia de IA, igual que en el
Sprint 1: cada una tiene su sección "Para decidir en la puerta" con las decisiones ambiguas
que hay que cerrar en la Puerta 1.

**El plan para las dos semanas que quedan está en
[docs/replan-al-2-de-noviembre.md](../replan-al-2-de-noviembre.md).**

## Orden de trabajo del Sprint 1

Las historias tienen dependencias, así que el orden importa:

```
#3 Crear proyectos  ──┬── #4 Registrar integrantes ───────────┐
                      └── #6 Administrar Product Backlog ──┬──┴── #24 Interfaz de consola
                                                           └── #12 Estimar con Story Points
```

- **#3 va primero.** Sin proyecto no hay dónde cargar integrantes ni historias.
- #4 y #6 pueden hacerse en paralelo una vez que #3 esté en `dev`.
- #12 necesita que #6 esté en `dev`.
- **#24 va última**: necesita las cuatro anteriores en `dev`. La toma quien termine antes.

En la práctica: Agustín arranca primero, y los demás pueden ir escribiendo y aprobando sus
specs mientras esperan.

## Reglas

- El archivo se llama `NN-nombre-corto.md`, donde `NN` es el número de la issue.
- La spec va en la carpeta del sprint donde la historia **se comprometió**.
- Si una historia no se termina y vuelve al Product Backlog, cuando se vuelve a comprometer
  se **mueve** la spec con `git mv` a la carpeta del sprint nuevo: una historia tiene una
  sola spec, con su historial completo.
- En Requisitos no va nada técnico. En Diseño van decisiones, no código.
- Las specs se actualizan en el mismo pull request que las vuelve falsas.
