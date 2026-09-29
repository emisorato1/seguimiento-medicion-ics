# Sprint 0 — Definir el Product Backlog inicial

- **Ítem:** issue #2
- **Tipo:** tarea de planificación (no lleva el flujo de cuatro fases, ver el
  [README](README.md) de esta carpeta)
- **Estado:** **TERMINADO** (2026-09-28)
- **Hecho por:** Emiliano Sorato (Agile Enabler), con asistencia de IA,
  revisado por el equipo en el cierre del Sprint 0

## Objetivo

Traducir los nueve requerimientos mínimos de la consigna a historias de usuario cargadas en
el tablero, para que el equipo tenga qué estimar y qué comprometer en cada sprint.

## Decisiones tomadas

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| Una historia por cada capacidad del usuario, no una por requerimiento de la consigna | nueve historias, una por requerimiento | "Gestión de Sprints" como una sola historia daría 20 story points y sería imposible de estimar y de repartir |
| Cada historia se escribe como *Como / quiero / para* con criterios de aceptación | una lista de tareas técnicas | los criterios de aceptación son la entrada de los escenarios BDD y de los tests; sin ellos se corta la trazabilidad |
| Las historias no se crean sin criterios de aceptación | cargarlas primero y completarlas después | una historia sin criterios no se puede estimar ni dar por terminada |
| Ninguna funcionalidad entra al backlog sin estar en la consigna | agregar lo que pareciera útil | el producto se evalúa contra la consigna; lo demás es alcance que nadie pidió |
| El backlog queda sin estimar al cerrar el Sprint 0 | estimar sobre la marcha al crearlo | la estimación es del equipo y se hace en la planning, no de quien carga las tarjetas |

## Cobertura de la consigna

Los nueve requerimientos mínimos, con las historias que los cubren:

| # | Requerimiento | Historias |
|---|---|---|
| 1 | Gestión de proyectos | #3, #4, #5 |
| 2 | Product Backlog | #6, #7 |
| 3 | Gestión de Sprints | #8, #9, #10, #11 |
| 4 | Estimación y Planning Poker | #12, #13 |
| 5 | Registro de esfuerzo | #14 |
| 6 | Gestión de defectos | #15 |
| 7 | Métricas | #16 |
| 8 | Dashboard | #17 |
| 9 | Reportes (PDF en el Sprint 4) | #18, #19 |

17 historias en total. No hay ningún requerimiento sin historia, ni ninguna historia que no
salga de un requerimiento.

## Dependencias entre historias

Esto condiciona el orden de los sprints:

```
#3 Crear proyectos
 ├── #4 Registrar integrantes ──────────── #14 Registrar esfuerzo
 └── #6 Administrar Product Backlog
      ├── #7 Priorizar backlog
      ├── #12 Estimar con Story Points ─── #13 Planning Poker
      └── #8 Crear sprint
           └── #9 Asignar historias al sprint
                └── #10 Cerrar sprint
                     ├── #11 Consultar sprints anteriores
                     └── #16 Calcular metricas
                          ├── #17 Dashboard
                          └── #18 Reportes ── #19 Exportar a PDF
```

Y también:

- **#15 Gestionar defectos** necesita #6 (el defecto se relaciona con una historia) y #8
  (sprint de detección y de resolución).
- **#16 Métricas** necesita además #12 (story points) y #14 (horas reales) para poder
  calcular velocidad y desviación.

Por eso el Sprint 1 toma #3, #4, #6 y #12: son la raíz del árbol y las únicas que no
dependen de nada que todavía no exista.

## Evidencia

| Qué | Dónde |
|---|---|
| Las 17 historias | issues #3 a #19 del repositorio |
| Prioridades cargadas | [Project #3](https://github.com/users/emisorato1/projects/3) |
| Estimaciones | acta de planning del Sprint 1 |

## Qué quedó pendiente

- La estimación, que se hizo en la planning del Sprint 1.
- La historia **#13 (13 story points)** quedó marcada como candidata a partirse en dos
  cuando se acerque su implementación.
