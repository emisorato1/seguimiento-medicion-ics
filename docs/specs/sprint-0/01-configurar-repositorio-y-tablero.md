# Sprint 0 — Configurar el repositorio y el tablero

- **Ítem:** issue #1
- **Tipo:** tarea de configuración (no lleva el flujo de cuatro fases, ver el
  [README](README.md) de esta carpeta)
- **Estado:** **TERMINADO** (2026-09-28)
- **Hecho por:** Emiliano Sorato (Agile Enabler), con asistencia de IA,
  revisado por el equipo en el cierre del Sprint 0

## Objetivo

Dejar el entorno de trabajo listo para que los cuatro integrantes puedan trabajar en
paralelo sin pisarse, y para que el proceso quede registrado como evidencia evaluable.

## Decisiones tomadas

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| El GitHub Project es la fuente de verdad del estado | un archivo de estado en el repo, mantenido a mano | un archivo a mano se desactualiza el primer día que alguien se olvida; el tablero se actualiza donde se trabaja |
| `ESTADO.md` se genera desde el tablero con `scripts/estado.sh` y se commitea | consultar el tablero siempre en vivo | así el estado se puede leer sin tener `gh` instalado ni autenticado |
| Estados del tablero: Product Backlog → Sprint Backlog → Especificación SDD → En progreso → En review → Hecho | los tres estados por defecto (Todo / In Progress / Done) | la columna de Especificación SDD hace visible en el tablero que la spec va antes del código, que es lo que se evalúa |
| Ramas `main` (producción) ← `dev` (integración) ← rama por historia | trabajar todos sobre `main` | `main` queda siempre estable y cada historia se integra revisada; es el flujo que se usa en un proyecto real |
| Una historia = una rama = un pull request | una rama por tarea de la spec | una rama por tarea daría PRs de tres líneas y el PR dejaría de cerrar una historia completa |
| Solo se permite merge commit: squash y rebase desactivados | squash para dejar el historial prolijo | el squash aplasta los commits y borra la evidencia del ciclo TDD, que es justamente lo que evalúa la cátedra |
| `main` y `dev` protegidas: PR obligatorio, una aprobación y CI en verde | confiar en que el equipo respete las reglas | una regla escrita que nadie hace cumplir no es una regla |
| CI corre `gofmt`, `go vet` y `go test -cover` en cada PR | correr los tests a mano | si el CI no lo verifica, tarde o temprano entra código roto a `dev` |
| Identificadores del código en inglés, documentación en español | todo en español | inglés es la convención de Go; el glosario de traducciones está en `AGENTS.md` para que no haya tres nombres para lo mismo |

## Qué se hizo

- Repositorio público `emisorato1/seguimiento-medicion-ics` con README, `CLAUDE.md`
  (reglas de trabajo), `AGENTS.md` (convenciones técnicas) y `.gitignore`.
- GitHub Project público #3, vinculado al repositorio, con los seis estados y los campos
  Story Points, Prioridad, Sprint, Horas estimadas y Horas reales.
- Colaboradores: los cuatro integrantes con permiso de administración, y los profesores
  `ccarrasco-git` y `frsrpprats` con permiso de lectura.
- Ramas `main` y `dev` protegidas, con workflows de CI y CD.
- Scripts `scripts/tablero.sh` y `scripts/estado.sh`.
- Esqueleto del proyecto en Go: `go.mod`, `cmd/app` e `internal/`.

## Evidencia

| Qué | Dónde |
|---|---|
| Creación del repositorio | commit `92b54f6` (2026-09-07) |
| README con el proyecto y los integrantes | commits `3417b65` y `1d8ec0c` |
| Reglas de trabajo y scripts | commit `5e23152` |
| Esqueleto en Go, CI y CD | commit `1af3d5b` |
| Reglas de ramas y merge | PR #20 |
| Tablero | [Project #3](https://github.com/users/emisorato1/projects/3) |

## Qué quedó pendiente

- El profesor `frsrpprats` todavía no aceptó la invitación de colaborador.
