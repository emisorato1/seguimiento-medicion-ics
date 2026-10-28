# Documentación técnica

Cómo está construido **Software Metrics & Estimation**: su arquitectura, el modelo de dominio,
las decisiones de diseño, el manejo de errores, la persistencia y cómo se prueba. Es el
entregable "documentación técnica" de la consigna; el uso de la aplicación está en el manual
de usuario.

- **Responsable:** Agustín Salinas
- **Estado:** en construcción. Las secciones marcadas *(en construcción)* se completan a medida
  que las historias entran a `dev`. Última actualización: 27/10/2026, con lo que está en `dev`:
  proyectos, integrantes, backlog, sprints (#8 a #11), estimación y Planning Poker (#12, #13 y
  #73), esfuerzo (#14), defectos (#15), métricas (#16) y la web (#77 y #78, salvo el reporte).

Las reglas de trabajo del equipo están en [CLAUDE.md](../CLAUDE.md) y las convenciones
técnicas detalladas en [AGENTS.md](../AGENTS.md). Este documento no las repite: las explica y
muestra cómo encajan.

---

## 1. Visión general

Una aplicación para estimar, seguir y medir proyectos de software: proyectos e integrantes,
Product Backlog, sprints, estimación con story points y Planning Poker, registro de esfuerzo,
defectos, métricas, dashboard y reportes.

| | |
|---|---|
| Lenguaje | Go 1.27, con la biblioteca estándar y **una sola dependencia externa**, acordada con el equipo: el driver de SQLite `modernc.org/sqlite`, en Go puro y sin CGO (#177) |
| Interfaces | consola (historia #24) y web (historias #77 y #78) |
| Datos | en memoria en la consola; en una base SQLite en la web (#177). El archivo JSON de la #77 queda para importar y exportar |
| Pruebas | `testing` de la biblioteca estándar, escritas con TDD |
| Integración continua | GitHub Actions: formato, análisis estático, tests, cobertura y coherencia de las specs; los tests también en Windows |

---

## 2. Arquitectura

El código se divide en tres capas, con una regla que no se rompe: **las dependencias apuntan
siempre hacia el dominio**, y el dominio no depende de nadie.

```
cmd/app                        cmd/web + internal/web
consola (#24), en memoria      web (#77), con los datos en una base SQLite
        │                              │
        │                              ▼
        │                      internal/store     lee y escribe la base y el archivo de datos
        │                              │
        ▼                              ▼
internal/domain    el núcleo: entidades, reglas de negocio, validaciones, métricas, errores
```

| Paquete | Qué hace | Qué **no** hace |
|---|---|---|
| `internal/domain` | todas las reglas de negocio y validaciones; arma el texto del reporte y del PDF, y los valores del dashboard ya escritos | no imprime, no lee del teclado, no toca archivos, no lee el reloj del sistema, no dibuja |
| `internal/store` | lee y escribe la base SQLite y el archivo de datos, e importa y exporta | no valida datos: eso lo hace el dominio, con su conversión |
| `cmd/app` | la consola: lee lo que escribe el usuario, llama al dominio y muestra el resultado | no valida nada por su cuenta: le pregunta al dominio |
| `cmd/web` | lee las opciones (`-puerto`, `-datos`, `-red`), carga los datos guardados, toma el puerto avisando si está ocupado o no hay permiso, y arranca la web con la fecha de la computadora | no guarda ni valida nada: lo hacen `store` y el dominio |
| `internal/web` (#77 y #78) | lee los formularios, llama al dominio, guarda con `store` y muestra las páginas. En el dashboard (#17) mide y dibuja las barras de los gráficos en SVG | no valida reglas: las valida el dominio |

Las dos interfaces usan **el mismo dominio**: una regla se escribe y se testea una sola vez,
y vale igual en la consola y en la web.

**Por qué así.** Con el dominio aislado:

- **Se testea solo**, sin pantalla, teclado ni archivos. Por eso el dominio puede tener el
  100% de cobertura.
- **La interfaz se puede cambiar** sin tocar una regla. Pasó de verdad: el Sprint 2 sumó la
  web al lado de la consola, y ninguna regla de negocio cambió; el dominio solo sumó dos
  consultas que la web necesitaba (listar los proyectos y la escala de story points).
- **Una validación no se puede saltear**: como la interfaz no puede modificar los datos del
  dominio directamente (ver la sección 4), cualquier cambio pasa por las reglas.

**Un solo paquete de dominio.** Todo el núcleo vive en `internal/domain`, con un archivo por
entidad (`project.go`, `story.go`, `member.go`…) y archivos extra por funcionalidad cuando
varias historias tocan la misma entidad (`story_estimation.go`, `project_registry.go`).
Separarlo en un paquete por entidad crearía ciclos de importación apenas una historia
referencia a un proyecto; separar por archivo evita además que dos personas editen el mismo.

---

## 3. Modelo de dominio

**Todo cuelga del proyecto.** Un proyecto contiene sus integrantes, su Product Backlog, sus
sprints, el esfuerzo registrado, sus defectos y las rondas de Planning Poker de sus historias;
nada de eso existe sin él. Los proyectos los guarda el **registro de proyectos**, que
es el único que les asigna número.

```
ProjectRegistry ─── Project ─┬── Member       (integrantes, #4)
                             ├── Story        (Product Backlog, #6 y #7; estimación, #12)
                             ├── VotingRound  (Planning Poker de una historia, #13 y #73)
                             ├── Sprint       (sprints, #8; Sprint Backlog, #9; cierre, #10)
                             ├── Effort       (esfuerzo real, #14)
                             └── Defect       (defectos, #15)
```

| Entidad | Historia | Datos principales | Archivo |
|---|---|---|---|
| `ProjectRegistry` | #3 | los proyectos y el contador de números | `project_registry.go` |
| `Project` | #3, #5 | nombre, descripción, fecha de inicio y de fin; su estado (#5) | `project.go`, `project_status.go` |
| `Member` | #4 | nombre y rol (Agile Enabler o Product Builder) | `member.go` |
| `Story` | #6, #7, #12, #16 | título, descripción, prioridad, estado, story points, horas estimadas, criterios de aceptación; el backlog ordenado por prioridad (#7) | `story.go`, `story_priority.go`, `story_estimation.go`, `story_hours.go` |
| `Sprint` | #8, #9, #10, #11 | número, Sprint Goal, fechas, estado y Sprint Backlog (las historias asignadas, en orden); al cerrar, la fecha de cierre, las historias completadas y los story points de cada una (#10); la consulta de los cerrados (#11) | `sprint.go`, `sprint_backlog.go`, `sprint_close.go`, `sprint_query.go` |
| `Effort` | #14 | historia, integrante, actividad, horas y día | `effort.go` |
| `Metrics` | #16 | las once métricas de un sprint o del proyecto, calculadas cada vez que se piden | `metrics.go` |
| `Defect` | #15 | título, descripción, severidad, estado, historia, sprint de detección y de resolución | `defect.go` |
| `VotingRound` | #13, #73 | historia, número de ronda, participantes fijados al abrirla, votos ocultos hasta revelar, diferencias y estimación acordada | `voting_round.go`, `voting_round_agreement.go` |
| `Report` | #18, #19 | — | *(en construcción)* |

**Un solo sprint abierto por proyecto.** El sprint nace abierto y se cierra con la #10: no
hay un estado "planificado". El Sprint Backlog guarda solo los números de las historias, y
el detalle de cada una se pide al Product Backlog, así hay una sola fuente de verdad. La
lista **no se vacía al cerrar**: es el registro de lo que se comprometió, y sobre eso miden
las métricas.

**Identificadores.** Los asigna el dominio, empezando en 1. Los proyectos son únicos en todo
el sistema; las historias y los sprints, únicos dentro de su proyecto: la primera historia de
cada proyecto es la 1. Los contadores **nunca retroceden**, así un número no se reutiliza
aunque algún día exista el borrado.

**Fechas.** Son días, sin hora ni zona horaria: dos fechas del mismo día son iguales aunque
lleguen con distinta hora o zona (defecto #59).

El contrato completo de las entidades que todavía no están implementadas, con los datos y las
operaciones de cada una, está en [docs/diseno-dominio.md](diseno-dominio.md).

---

## 4. Decisiones de diseño

Las decisiones que afectan a todo el proyecto están en la tabla *Decisiones de arquitectura*
de [AGENTS.md](../AGENTS.md), cada una con la alternativa descartada y su porqué. Las más
importantes, en una línea:

| Decisión | Por qué |
|---|---|
| **Los datos de las entidades están ocultos fuera del paquete**; solo se cambian con operaciones que validan | la interfaz no puede saltear una validación aunque quiera: el compilador no se lo permite |
| **El registro entrega la entidad guardada, no una copia** | con copias, cada historia tendría que acordarse de guardar el cambio, y un olvido lo pierde sin ningún error |
| **Toda operación valida todo antes de cambiar cualquier dato** | una operación rechazada deja todo como estaba, sin tener que deshacer nada |
| **Crear y modificar comparten la misma validación** (por ejemplo `prepare` en `project.go`) | las reglas no se pueden desincronizar entre las dos operaciones |
| **El pedido de modificación distingue "no tocar", "cambiar" y "quitar"** (`ProjectChanges`, `OptionalChange`) | "no mandé el dato" y "quiero borrarlo" no se confunden, y un pedido contradictorio no se puede escribir |
| **El dominio no lee el reloj del sistema**: la fecha de hoy la recibe de quien llama | los tests no dependen del día en que se corren |
| **Una sola dependencia externa, acordada con el equipo**: el driver de SQLite en Go puro (`modernc.org/sqlite`), sin CGO | así compila igual en Windows, Linux y macOS; fuera de eso, solo la biblioteca estándar, porque el proyecto es chico y se evalúa el código propio |

Las decisiones propias de cada funcionalidad están en la Fase 2 (Diseño) de su spec, en
[docs/specs/](specs/README.md).

---

## 5. Manejo de errores

**Los errores son valores del dominio**, declarados todos en
[`internal/domain/errors.go`](../internal/domain/errors.go):

```go
ErrProjectNameRequired = errors.New("el nombre del proyecto es obligatorio")
```

| Regla | Por qué |
|---|---|
| **Se devuelven, nunca `panic`** | un error de validación es parte normal del uso, no una falla del programa |
| **Se comparan con `errors.Is`**, nunca por el texto | el mensaje se puede corregir (pasó con las tildes, defecto #60) sin romper ningún test ni ninguna interfaz |
| **El nombre va en inglés y el mensaje en español** | el mensaje lo lee el usuario: la consola y la web lo muestran tal cual |
| **Un error puede llevar un dato extra envolviéndolo** con `fmt.Errorf("%w: …")` | sigue cumpliendo `errors.Is` y el mensaje dice más. Por ejemplo, cuando la base ya está abierta por otro sistema, el error dice cuál es la base (`ErrDatabaseInUse` en `internal/store/database.go`); y crear un sprint habiendo otro abierto responde *"ya hay un sprint abierto en el proyecto: el sprint 1"* (`AddSprint` en `sprint.go`) |
| **Cada error tiene un comentario con la historia y la regla que lo originan** | se puede rastrear de dónde sale sin buscar |
| **Lo que no es un error no se informa como error** | un backlog vacío o un proyecto sin sprint abierto son situaciones normales y devuelven una lista vacía o un "no hay", no un error |

Como toda operación valida antes de cambiar nada (sección 4), **un error nunca deja datos a
medio cambiar**.

---

## 6. Persistencia

La consola trabaja en memoria. La web guarda todo en **una base SQLite**, `datos-ics.db`
(historia #177). Hasta la #177 lo guardaba en un archivo JSON, que ahora sirve para importar
y exportar los datos.

| Parte | Dónde | Qué hace |
|---|---|---|
| Conversión a JSON y desde JSON | `internal/domain/registry_json.go` | vive en el dominio porque es lo único que puede leer los datos ocultos sin exportarlos; al cargar, valida cada entidad |
| La base | `internal/store/database.go` | guarda y carga el registro, en una transacción |
| El esquema | `internal/store/migrations/` | las migraciones, numeradas y embebidas en el binario |
| El puente entre las dos | `internal/store/rows.go` | la forma del archivo, como tipos propios del store |
| Importar y exportar | `internal/store/exchange.go` y `internal/store/file.go` | el archivo JSON y el de la versión anterior; `file.go` lee y escribe el archivo sin dejarlo a medias |

- **El dominio no sabe de la base.** Al guardar, el registro se convierte con la conversión
  del dominio, se desarma en filas y se escribe **en una transacción** que vacía las tablas y
  las vuelve a llenar: si algo falla, la base queda como estaba. Al cargar, las filas se
  rearman en el mismo formato y se le pasan a la conversión del dominio. El texto nunca se
  guarda en una columna: **hay una tabla por entidad**, y tablas aparte para lo que una
  entidad tiene varias veces (criterios, historias de un sprint, participantes y votos).
- **Al cargar, cada entidad pasa por las mismas validaciones que al crearla**: una base
  editada a mano no puede meter datos inválidos.
- **Y además se controla lo que solo los datos guardados pueden traer mal**, porque usando el
  sistema nunca pasaría. Por ejemplo, en los sprints: números repetidos o que el contador no
  alcanza, más de un sprint abierto, y en el Sprint Backlog una historia que no existe o que
  aparece dos veces.
- **Los contadores de la numeración se guardan como vienen** (en `registry` y en
  `projects`), nunca se deducen del número más alto: así no se reutiliza el número de algo
  que se quitó.
- **Toda historia que agrega una entidad o un dato lo suma a la conversión**, con su caso en
  el test de ida y vuelta (`registry_json_test.go`), **y a la base**: a los tipos de
  `rows.go`, a una tabla con **una migración nueva** (`0002_…sql`; nunca se edita una que ya
  se usó), y al test `TestUnRegistroCompletoVuelveIgual`. Si se suma a la conversión y no a
  la base, el guardado falla con «No quedó guardado» en vez de perder el dato en silencio.
- Las fechas se guardan como `AAAA-MM-DD`, y las prioridades, estados y roles con su nombre:
  la base y el archivo se pueden leer y no se rompen si se reordenan las constantes.
- La base guarda en `PRAGMA user_version` el número de la última migración aplicada. Una base
  con un número mayor que el que conoce el sistema no se abre.
- **Una base la usa un solo sistema por vez**: se abre con el bloqueo exclusivo de SQLite
  (`locking_mode=EXCLUSIVE`), y otro sistema que la quiera abrir recibe «la base ya está
  abierta por otro sistema». El sistema operativo suelta el bloqueo al cerrar o si el programa
  se corta.

### Cómo se guarda, paso a paso

1. **Al abrir**, `cmd/web` abre la base con `store.OpenDatabase(ruta)`, que la crea con su
   esquema si no existe y toma el bloqueo para que otro sistema no la abra. Si la base está
   vacía y junto a ella hay un `datos-ics.json` de la versión anterior, lo importa una sola vez
   (sección 10). Después lee los datos con `Load()`: si la base no se puede leer, o el dominio
   no acepta lo que tiene, el sistema no abre: muestra el error (`ErrSavedDataInvalid`) y deja
   la base como estaba.
2. **Cada operación aceptada** termina en `done` (`internal/web/server.go`), que guarda el
   registro entero, deja el mensaje que la confirma y redirige a la página del resultado.
3. **El guardado** (`(*Database).Save` en `internal/store/database.go`) pasa el registro por la
   conversión del dominio y, en una sola transacción, vacía las tablas y escribe todo de nuevo:
   si algo falla, la transacción se deshace y la base queda como estaba, sin nada a medias.
4. **Si el guardado falla**, la página avisa «No quedó guardado» y el cambio sigue en memoria:
   el próximo guardado que funcione lo incluye (`ErrNotSaved`).
5. **Los pedidos usan el registro de a uno**: las páginas que solo leen pueden atenderse a la
   vez, y las que cambian algo van solas, así una operación y su guardado quedan juntos.
6. **En los tests de la web** casi no se escribe nada: el servidor recibe un guardado falso que
   cuenta las veces que se guardó y puede simular una falla. Solo los dos escenarios de cerrar
   y volver a abrir guardan de verdad, con el archivo y con la base, en una carpeta temporal.

---

## 7. Pruebas y cobertura

### TDD, con la evidencia en el historial

Cada regla de negocio se escribió con el ciclo **RED → GREEN → REFACTOR**, y **el test va en
un commit y la implementación en otro**. Así el historial de Git muestra el test fallando
antes de que exista el código:

```
37e4f2f test: crear un proyecto con todos los datos (tarea 1, historia 3)
0fecae8 feat: registro de proyectos y alta de un proyecto (tarea 1, historia 3)
```

Cuando un test pasó en verde sin código nuevo, no se forzó un RED: quedó solo el commit del
test, explicado en su mensaje y en la Fase 4 de la spec (por ejemplo la tarea 3 de la #3).

### Cómo se escriben los tests

La convención completa está en la sección *Cómo se escriben los tests* de
[AGENTS.md](../AGENTS.md). Lo esencial:

- **Un test por escenario BDD**, con el nombre del escenario en PascalCase y un comentario
  que dice qué escenario y qué reglas cubre:

  ```go
  // Escenario: Rechazar un proyecto sin nombre
  // Cubre: RN-2
  func TestRechazarUnProyectoSinNombre(t *testing.T) {
  ```

- Un **"Esquema del escenario"** es un solo test con una tabla de casos, uno por fila.
- **Cada test arranca de cero**: arma lo que necesita y no comparte estado con otros.
- Los errores se comparan con `errors.Is`.
- Los tests viven al lado del código que prueban, en `archivo_test.go`.

### Cobertura

El CI mide la cobertura en cada pull request y la publica en el resumen de la corrida. El
objetivo es **100% en `internal/domain`**, que es donde están las reglas de negocio. Los
números de cada sprint están en el [informe de métricas](informe-de-metricas.md).

---

## 8. Calidad automatizada

Cada pull request y cada push a `dev` o `main` pasa por el CI
([`.github/workflows/ci.yml`](../.github/workflows/ci.yml)). **No se mergea con el CI en
rojo.**

| Paso | Qué verifica |
|---|---|
| Coherencia de las specs (`scripts/coherencia.sh`) | que cada spec, el índice y el registro de las puertas digan lo mismo, y que no queden marcas de conflicto |
| Formato (`gofmt`) | que todo el código esté formateado igual |
| Análisis estático (`go vet`) | errores comunes que compilan pero están mal |
| Tests y cobertura | que todos los tests pasen, con el reporte de cobertura en el resumen de la corrida y como archivo descargable (`cobertura`) |
| Tests en Windows | los mismos tests en una máquina Windows. Es un chequeo aparte y no es obligatorio para mergear: existe porque parte del equipo trabaja en Windows (defecto #108) |

Cuando algo entra a `main`, el CD
([`.github/workflows/cd.yml`](../.github/workflows/cd.yml)) vuelve a correr los tests y
compila los binarios para Linux (amd64) y macOS (arm64).

`main` y `dev` están protegidas: no se puede pushear directo, y todo entra por pull request.

---

## 9. Trazabilidad

Cada funcionalidad se puede recorrer de punta a punta:

```
Historia (issue) → Spec SDD → Criterio de aceptación → Escenario BDD → Test → Código Go
```

**Ejemplo real, de la historia #3:**

| Eslabón | Dónde |
|---|---|
| Historia | issue #3, "Crear y modificar proyectos" |
| Spec | [docs/specs/sprint-1/03-crear-proyectos.md](specs/sprint-1/03-crear-proyectos.md) |
| Regla de negocio | RN-2: no se puede crear ni modificar un proyecto sin nombre |
| Criterio de aceptación | "No se puede crear ni modificar un proyecto sin nombre" |
| Escenario BDD | "Rechazar un proyecto sin nombre" |
| Test | `TestRechazarUnProyectoSinNombre` en `internal/domain/project_test.go` |
| Evidencia TDD | commit `1e45818` (test en rojo) y `13e9e68` (implementación) |
| Código | `prepare` en `internal/domain/project.go` y `ErrProjectNameRequired` en `errors.go` |

El recorrido de **todas** las historias está en [docs/trazabilidad.md](trazabilidad.md).
Cada spec cierra con su propia tabla de trazabilidad, y su Fase 4 tiene el commit del test y
el de la implementación de cada tarea.

---

## 10. Para desarrolladores

```bash
go test ./...                  # todos los tests
go test ./internal/domain/ -v  # los del dominio, con el nombre de cada escenario
go test ./... -cover           # con cobertura
go test -race ./...            # busca accesos al mismo dato desde dos pedidos a la vez
go vet ./...                   # análisis estático
gofmt -l .                     # archivos sin formatear (no tiene que devolver nada)
scripts/coherencia.sh          # coherencia de las specs
go run ./cmd/app               # la consola
go run ./cmd/web               # la web, en http://127.0.0.1:8080, con la base datos-ics.db
go run ./cmd/web -puerto 9090  # en otro puerto
go run ./cmd/web -red          # acepta conexiones de la red local, por ejemplo un celular
go run ./cmd/web -exportar copia.json            # todos los datos a un archivo
go run ./cmd/web -datos otra.db -importar copia.json  # un archivo a una base nueva
```

Si en la carpeta de la base está el `datos-ics.json` de la versión anterior y la base está
vacía, la web lo importa sola al abrir, una sola vez: después lo renombra a
`datos-ics.json.importado` y lo avisa.

Antes de abrir un pull request: los tests en verde, `gofmt -l .` vacío y `dev` traído a la
rama. El flujo completo de ramas y pull requests está en [CLAUDE.md](../CLAUDE.md).
