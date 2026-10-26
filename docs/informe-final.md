# Informe final del proyecto — esqueleto

> **Borrador del esqueleto, 26/10.** Es la estructura del informe en PDF que se entrega el
> 1/11. Lo arma Emiliano Sorato (Agile Enabler) con las métricas del repo, y Emiliano lo
> revisa antes de exportarlo. Cada sección dice de dónde salen sus datos. Lo que está entre
> corchetes, **[pendiente]**, se completa al cierre del producto (30/10) y del Sprint 4.
>
> El PDF se arma como `.docx` con los estilos de Word (Título 1, 2 y 3) y se exporta desde
> Word. La presentación final la arma Juliana Bustos y no va acá.

## Portada

Universidad Tecnológica Nacional, Facultad Regional San Rafael. Ingeniería y Calidad de
Software, 2026. Título: *Seguimiento y medición de proyectos de software*. Integrantes:
Agustín Salinas, Emiliano Sorato, Juliana Bustos y Maximiliano Eula [pendiente: legajos].

## 1. El producto

- Qué hace la aplicación y para quién, en un párrafo.
- Qué requerimientos mínimos de la consigna cubre, en una tabla: requerimiento, historia,
  estado. Fuente: el Product Backlog del tablero y `docs/specs/README.md`.
- Cómo se usa: consola y web. Fuente: `docs/manual-de-usuario.md`.
- Arquitectura en un diagrama: dominio en Go, guardado en archivo, consola y web. Fuente:
  `docs/documentacion-tecnica.md` y `docs/diseno-dominio.md`.

## 2. Cómo trabajamos

- **Scrum:** roles, calendario de los cinco sprints (0 a 4) y ceremonias. Fuente: el plan y
  las actas de `docs/actas/`.
- **Tablero:** columnas, campos (story points, sprint, horas estimadas y reales) y cómo se
  movía una historia. Fuente: el GitHub Project.
- **SDD, BDD y TDD:** spec con tres puertas antes del código, escenarios Given-When-Then y
  commits separados de test y de código. Fuente: `CLAUDE.md`, `AGENTS.md` y `docs/specs/`.
- **Uso de la IA:** qué historias se hicieron con asistencia de IA y cómo se revisaron. Las
  horas de esas historias son el tiempo de dirección y revisión.

## 3. Métricas por sprint

La tabla resumen de los cuatro sprints: story points planificados y completados, velocidad,
historias completadas, horas estimadas y reales, desviación, horas por story point y
defectos detectados y resueltos. Fuente: el [informe de métricas](informe-de-metricas.md),
que toma los sprints cerrados de sus actas de review y el abierto de
`scripts/metricas.sh "Sprint N"`.

- Sprint 1 y Sprint 2: ya cargados.
- Sprint 3 y Sprint 4: [pendiente].
- Las historias adelantadas y las que no tienen estimación, explicadas aparte, como en el
  informe de métricas.

## 4. Gráficos

[pendiente] Velocidad por sprint, horas estimadas contra reales y defectos detectados y
resueltos. Fuente: los mismos números de la sección 3, o capturas del dashboard de la
aplicación (#17) cargado con los datos de nuestro propio proyecto.

## 5. Pruebas y cobertura

- Cobertura por paquete y total al cierre de cada sprint. Fuente: `go test ./...
  -coverprofile` sobre `dev`, el mismo comando del CI.
- Cantidad de tests y de escenarios BDD con test. Fuente: `docs/trazabilidad.md`.
- El CI: qué corre y en qué sistemas.

## 6. Evidencia de TDD

Cuántos ciclos RED → GREEN quedaron en el historial y cómo se leen. Fuente:
`scripts/trazabilidad.sh`, que marca cada escenario con su commit de test y su commit de
código. Al 26/10: 250 ciclos completos, 44 tests que pasaron de entrada ("sin RED"), 13 sin
commit de código propio, 3 de andamiaje y 1 con test y código en el mismo commit.
[pendiente: actualizar al cierre]

## 7. Trazabilidad de punta a punta

El recorrido completo de una historia: historia → spec SDD → criterio de aceptación →
escenario BDD → test → código Go, con enlaces. [pendiente: elegir la historia junto con la
presentación, para que sea la misma que se muestre en la defensa.] Fuente:
`docs/trazabilidad.md`.

## 8. Retrospectivas

Lo que se acordó en cada retro y si se cumplió en el sprint siguiente, en una tabla. Fuente:
las actas `sprint-N-retrospectiva.md`. Sprint 1 y Sprint 2: cargados en sus actas. Sprint 3 y
4: [pendiente].

## 9. Conclusiones

[pendiente] Qué aprendimos midiéndonos: cómo cambió la estimación sprint a sprint, qué costó
más de lo pensado y qué haríamos distinto. Sale de la sección de conclusiones del informe de
métricas.

## Anexos

Enlaces al repositorio, al tablero, a las actas, a las specs, al manual de usuario y a la
documentación técnica.
