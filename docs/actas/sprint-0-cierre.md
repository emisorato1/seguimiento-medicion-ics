# Acta — Cierre del Sprint 0

- **Objetivo del Sprint 0:** dejar el entorno de trabajo listo — repositorio, tablero y
  Product Backlog inicial — para poder empezar a construir en el Sprint 1.
- **Inicio:** 2026-09-07 (primer commit del repositorio)
- **Cierre:** 2026-09-28
- **Participantes:** Emiliano Sorato, Agustín Salinas, Juliana Bustos, Maximiliano Eula
- **Facilita:** Emiliano Sorato (Agile Enabler)

## Qué se hizo

| Ítem | Estado | Evidencia |
|---|---|---|
| Repositorio público en GitHub | Hecho | `92b54f6` (2026-09-07) |
| README explicando el proyecto | Hecho | `3417b65` (2026-09-14) |
| Roles de Scrum asignados | Hecho | `1d8ec0c` (2026-09-14) |
| Guía de Planning Poker y acta de planning | Hecho | `43caf9c` (2026-09-14) |
| GitHub Project público y vinculado al repo | Hecho | Project #3 |
| Tablero con estados y campos propios | Hecho | Project #3 |
| Profesores agregados como colaboradores (Read) | Hecho | `ccarrasco-git` aceptó; `frsrpprats` pendiente |
| Integrantes agregados al repo y al tablero | Hecho | los tres aceptaron |
| Product Backlog inicial (17 historias) | Hecho | issues #3 a #19 |
| Ítems de Sprint 0 en el tablero | Hecho | issues #1 y #2, cerradas |
| Reglas de trabajo del proyecto (CLAUDE.md, AGENTS.md, plantilla de spec) | Hecho | `5e23152`, `d0f90df` |
| Flujo de ramas, CI y CD | Hecho | `1af3d5b`, PR #20 |

## Qué quedó pendiente

- Las historias del backlog están **sin estimar**. La estimación con Planning Poker quedó
  para la planning del Sprint 1.
- La entrega "Repositorio proyecto" del campus se subió fuera de término: el repositorio
  estaba listo desde el 2026-09-14 pero la entrega no se cargó antes del cierre
  (2026-09-28 00:00).

## Retrospectiva

### Qué salió bien

- Se cumplió el objetivo del sprint: repositorio y Project públicos, Product Backlog de 17
  historias, README y colaboradores, todo listo.
- Quedó definido un flujo de trabajo concreto y escrito —SDD en cuatro fases con puertas,
  ramas `main`/`dev`, CI/CD y protección de ramas— y no solamente la intención de trabajar
  ordenados.
- El backlog salió directo de los nueve requerimientos mínimos de la consigna, así que no
  hay funcionalidad inventada ni falta ninguna.
- El estado del proyecto quedó automatizado: el tablero es la fuente de verdad y un script
  lo vuelca al repositorio, así los dos no se desincronizan.

### Qué no salió bien

- Todo el Sprint 0 tiene un solo autor en el historial de commits. El resto del equipo no
  participó de la construcción.
- La entrega del campus venció sin subirse, aunque el repositorio estaba listo catorce días
  antes. Nadie estaba mirando la fecha de cierre.
- El sprint terminó sin planning, así que el backlog quedó sin estimar y el Sprint 1 no
  puede arrancar de una.
- El primer pull request se mergeó salteando la protección de rama, sin que ninguna persona
  lo revisara.

### Qué vamos a cambiar en el Sprint 1

| Acuerdo | Responsable |
|---|---|
| Cada integrante toma al menos una historia y la lleva de punta a punta: spec, tests, código y PR | cada uno |
| Ningún PR se mergea sin la aprobación de otro integrante. No se usa el bypass de admin | Emiliano Sorato (Agile Enabler) |
| Revisar las fechas del campus al empezar cada semana y avisar al equipo tres días antes de cada cierre | Emiliano Sorato (Agile Enabler) |
| Las historias se estiman en la planning, antes de abrir el sprint | todo el equipo |
