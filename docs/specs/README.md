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
historias terminadas están en `dev` con el dominio al 100% de cobertura.

La **#12** y la **#24** volvieron al Product Backlog sin sprint asignado. Las dos conservan
sus tres puertas aprobadas y la #24 sus 9 tareas hechas, así que cuando se replanifiquen
arrancan desde donde quedaron: no hay que rehacer las specs.

Los números completos están en el [acta de review](../actas/sprint-1-review.md).

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
