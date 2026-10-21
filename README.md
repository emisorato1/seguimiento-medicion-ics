# Seguimiento y Medición ICS

Proyecto integrador de **Ingeniería y Calidad de Software** — UTN FRSR, 2026.

## De qué se trata

Es una aplicación para **estimar, seguir y medir proyectos de software** (Software Metrics
& Estimation). La idea es que un equipo pueda cargar su proyecto, armar el Product Backlog,
planificar sprints, estimar las historias con Planning Poker, registrar el esfuerzo real y
los defectos, y después ver en métricas cómo viene el proyecto comparado con lo que se había
planificado.

Lo que va a permitir hacer:

- Crear proyectos y cargar los integrantes del equipo.
- Administrar el Product Backlog (historias con prioridad, estado, story points y criterios
  de aceptación).
- Armar sprints con su objetivo, asignarles historias y cerrarlos.
- Estimar con Story Points usando Planning Poker: los votos quedan ocultos hasta que votan
  todos, se detectan las diferencias y se puede hacer otra ronda hasta acordar.
- Registrar las horas reales que cada integrante le dedicó a cada historia.
- Cargar defectos con su severidad y estado, indicando en qué sprint se detectaron y en cuál
  se resolvieron.
- Ver métricas: story points planificados vs. completados, velocidad del equipo, horas
  estimadas vs. reales, desviación, porcentaje de historias terminadas y defectos detectados
  y resueltos.
- Un dashboard con gráficos y reportes por proyecto o por sprint.

## Tecnología

El núcleo de la solución y todas las reglas de negocio están desarrollados en **Go**.

## Cómo correrlo

Para usar el sistema sin programar está el **[manual de usuario](docs/manual-de-usuario.md)**.

Hace falta **Go 1.23 o posterior**. No hay dependencias externas: todo es biblioteca estándar.

```bash
git clone https://github.com/emisorato1/seguimiento-medicion-ics.git
cd seguimiento-medicion-ics

go test ./...                 # correr todos los tests
go test ./... -cover          # con el porcentaje de cobertura
go run ./cmd/app              # ejecutar la aplicación
```

### Qué funciona hoy

El **núcleo del dominio** está implementado y **cubierto al 100% por tests**: crear y
modificar proyectos, registrar integrantes, administrar el Product Backlog y estimar con
Story Points. Se puede ver funcionando corriendo los tests, que están escritos con el nombre
de cada escenario de las especificaciones:

```bash
go test ./internal/domain/ -v
```

La **interfaz de consola** está en curso (historia #24): por ahora `go run ./cmd/app`
muestra el nombre y la versión. Cuando esté terminada, va a ser la forma de usar el sistema
a mano.

## Cómo trabajamos

El proyecto se desarrolla con **Scrum** (sprints, planning, daily, review y retrospectiva) y
siguiendo tres prácticas que se piden en la materia:

- **SDD**: antes de programar una funcionalidad se escribe su especificación (objetivo,
  entradas, salidas, reglas de negocio, casos límite y criterios de aceptación), y esa
  especificación se versiona junto con el código.
- **BDD**: los criterios de aceptación se escriben como escenarios Given-When-Then, cubriendo
  los casos normales, los alternativos, los límite y los de error.
- **TDD**: se programa con el ciclo RED → GREEN → REFACTOR, primero el test y después la
  implementación. Por eso el historial de commits se mantiene sin aplastar, porque es la
  evidencia de cómo se fue construyendo.

Todo queda trazado de punta a punta: historia de usuario → especificación → criterios de
aceptación → escenario BDD → test → código.

## Equipo

| Apellido y Nombre | Rol | Usuario de GitHub |
|---|---|---|
| Sorato, Emiliano | Agile Enabler | [@emisorato1](https://github.com/emisorato1) |
| Salinas, Agustín | Product Builder | [@Salinas5](https://github.com/Salinas5) |
| Bustos, Juliana | Product Builder | [@bustosjuliana](https://github.com/bustosjuliana) |
| Eula, Maximiliano | Product Builder | [@maxi627](https://github.com/maxi627) |

El rol de Product Architect lo ocupan los profesores de la cátedra.

## Documentación del proceso

- [CLAUDE.md](CLAUDE.md) — **las reglas de trabajo del proyecto**. Es el punto de entrada: qué construimos, cómo trabajamos y cómo seguir.
- [ESTADO.md](ESTADO.md) — estado actual del tablero (generado, no se edita a mano).
- [AGENTS.md](AGENTS.md) — comandos, estructura, estilo, flujo de git y límites del proyecto.
- [Índice de especificaciones](docs/specs/README.md) — **qué historia le toca a cada uno, en qué estado está y en qué orden**. Es por donde se empieza para trabajar.
- [Plantilla de especificación SDD](docs/specs/PLANTILLA.md) — el flujo de cuatro fases con sus puertas de aprobación.
- [Plan final al 2 de noviembre](docs/plan-final-al-2-de-noviembre.md) — **el plan vigente**: calendario de los sprints que quedan, reparto por persona y riesgos.
- [Contrato del dominio](docs/diseno-dominio.md) — las entidades compartidas, acordadas antes de implementarlas, para que nadie espere a nadie.
- [Guía de Planning Poker](docs/guia-planning-poker.md) — cómo estimamos las historias.
- [Actas](docs/actas/) — plannings, reviews y retrospectivas de cada sprint.
- [Informe de métricas y cobertura](docs/informe-de-metricas.md) — las métricas del propio proyecto, sprint por sprint, y la cobertura de los tests.
- [Candidatas del Sprint 2](docs/sprint-2-candidatos.md) — el análisis de dependencias previo a la planning del Sprint 2.
- [Replanificación al 2 de noviembre](docs/replan-al-2-de-noviembre.md) — el análisis de recorte que se descartó, conservado como registro de la decisión.

## Tablero

El seguimiento del trabajo se lleva en el GitHub Project del equipo:
https://github.com/users/emisorato1/projects/3
