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

Las Fases 1 se redactaron como borrador con asistencia de IA a partir de la consigna y de los
criterios de aceptación de cada issue, y cada responsable las revisó y corrigió antes de
presentarlas a la Puerta 1.

**Solo se escribe código de una historia que tiene las tres puertas aprobadas.** El orden es
Puerta 1 (requisitos) → Fase 2 y Puerta 2 (diseño) → Fase 3 y Puerta 3 (tareas) → recién ahí
la implementación.

**El Sprint 1 cerró con 6 de los 16 story points comprometidos** (velocidad 6). Las dos
historias terminadas, la #3 y la #4, están en `dev` con el dominio al 100% de cobertura. Los
números completos están en el [acta de review](../actas/sprint-1-review.md).

La **#6**, la **#12** y la **#24** se comprometieron en el Sprint 1 y cuentan en el Sprint 2,
así que figuran en su tabla. La #6 tenía todo aprobado al cierre, pero se mergeó a `dev` el
19/10. La #12 y la #24 volvieron a medio hacer. Las tres conservan su carpeta `sprint-1/`, que
es donde se comprometieron, y no hay que rehacer sus specs.

### Sprint 2 — del 20 al 25 de octubre (31 SP)

| Spec | Historia | SP | Responsable | Estado |
|---|---|---|---|---|
| [Priorizar el Product Backlog](sprint-2/07-priorizar-backlog.md) | #7 | 2 | Agustín Salinas | ✅ Terminada — mergeada a `dev` (PR #105) |
| [Crear un sprint con su Sprint Goal](sprint-2/08-crear-sprint.md) | #8 | 3 | Agustín Salinas | ✅ Terminada — mergeada a `dev` (PR #138) |
| [Asignar historias a un sprint](sprint-2/09-asignar-historias-a-sprint.md) | #9 | 5 | Agustín Salinas | ✅ Terminada — mergeada a `dev` (PR #138) |
| [Registrar el esfuerzo real](sprint-2/14-registrar-esfuerzo-real.md) | #14 | 3 | Juliana Bustos | ✅ Terminada — mergeada a `dev` (PR #100) |
| [Gestionar defectos](sprint-2/15-gestionar-defectos.md) | #15 | 5 | Juliana Bustos | ✅ Terminada — mergeada a `dev` (PR #129) |
| [Administrar el Product Backlog](sprint-1/06-administrar-product-backlog.md) | #6 | 5 | Emiliano Sorato | ✅ Terminada — mergeada a `dev` el 19/10, entre el Sprint 1 y el 2 |
| [Estimar con Story Points](sprint-1/12-estimar-story-points.md) | #12 | 2 | Maximiliano Eula | ✅ Terminada — mergeada a `dev` (PR #70) |
| [Interfaz de consola](sprint-1/24-interfaz-de-consola.md) | #24 | 3 | Maximiliano Eula | ✅ Terminada — mergeada a `dev` (PR #106) |
| [Consultar el estado de un proyecto](sprint-2/05-consultar-estado-proyecto.md) | #5 | 3 | Juliana Bustos | ✅ Terminada — mergeada a `dev` (PR #102) |
| [Interfaz web: proyectos, integrantes y backlog](sprint-2/77-interfaz-web-base.md) | #77 | sin estimar | Emiliano Sorato | ✅ Terminada — mergeada a `dev` (PR #95) |

### Sprint 3 — del 26 al 30 de octubre (36 SP)

| Spec | Historia | SP | Responsable | Estado |
|---|---|---|---|---|
| [Cerrar un sprint](sprint-3/10-cerrar-sprint.md) | #10 | 5 | Emiliano Sorato | ✅ Terminada — mergeada a `dev` (PR #150) |
| [Calcular las métricas del proyecto](sprint-3/16-calcular-metricas.md) | #16 | 8 | Emiliano Sorato | ✅ Terminada — mergeada a `dev` (PRs #81, #140 y #151) |
| [Mostrar el dashboard](sprint-3/17-mostrar-dashboard.md) | #17 | 8 | Juliana Bustos | ✅ Lista para implementar — puertas aprobadas en el PR #136 |
| [Votar en secreto en una ronda](sprint-3/13-votar-en-secreto.md) | #13 | 8 | Maximiliano Eula | ✅ Terminada — mergeada a `dev` (PR #141), adelantada al Sprint 2 |
| [Consultar sprints anteriores](sprint-3/11-consultar-sprints-anteriores.md) | #11 | 2 | Agustín Salinas | 🔍 Implementada — en review (PR #143) |
| [Generar el reporte de proyecto o sprint](sprint-3/18-generar-reporte.md) | #18 | 5 | Agustín Salinas | 📝 Borrador completo — pendiente de las tres puertas |
| [Interfaz web: sprints, esfuerzo, defectos, Planning Poker y reportes](sprint-3/78-interfaz-web-operaciones.md) | #78 | sin estimar | Emiliano Sorato | 📝 Borrador completo — pendiente de las tres puertas |
| [Guardar los datos en una base de datos](sprint-3/177-guardar-en-base-de-datos.md) | #177 | sin estimar | Emiliano Sorato | 🟡 Puertas aprobadas por el responsable (26/10) — pendiente de ratificación de un compañero |

### Sprint 4 — 31 de octubre y 1 de noviembre (10 SP)

| Spec | Historia | SP | Responsable | Estado |
|---|---|---|---|---|
| [Exportar el reporte a PDF](sprint-4/19-exportar-reporte-pdf.md) | #19 | 5 | Emiliano Sorato | 🟢 Puerta 3 aprobada — lista para implementar (espera la #18) |
| [Resolver diferencias y acordar la estimación](sprint-4/73-resolver-diferencias-y-acordar.md) | #73 | 5 | Maximiliano Eula | ✅ Terminada — mergeada a `dev` (PR #148), adelantada al Sprint 3 |

### Qué es un "borrador completo"

Las 14 specs que quedan se escribieron **enteras y por adelantado** el 20/10 —las tres fases,
con asistencia de IA— en vez de fase por fase esperando cada puerta. Es el modo acordado
hasta el 2/11 (ver *Modo sprint final* en [CLAUDE.md](../../CLAUDE.md)): esperar las puertas
una por una fue el cuello de botella medido del Sprint 1.

Que sea un borrador completo **no quiere decir que esté aprobado**. Las tres puertas siguen
pendientes y las aprueba una persona, en una sola revisión, antes de que se escriba código.
Cada spec lo dice en su encabezado, y `scripts/coherencia.sh` verifica que un documento que
se declara borrador no tenga ninguna puerta firmada.

Cada una tiene su sección **"Para decidir en la puerta"** con las decisiones ambiguas que
tomó el borrador y qué cambiaría si el equipo decide distinto: es lo que hace que la
revisión sirva para algo. El responsable la revisa, la corrige y la presenta.

**El plan completo —calendario, reparto y riesgos— está en
[docs/plan-final-al-2-de-noviembre.md](../plan-final-al-2-de-noviembre.md).**

## Orden de trabajo hasta el 2 de noviembre

Las historias que quedan están encadenadas:

```
#8 → #9 → #10 → #16 ─┬─→ #17 Dashboard
                     └─→ #18 → #19 PDF

#13 → #73            (Planning Poker, cadena aparte)
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
