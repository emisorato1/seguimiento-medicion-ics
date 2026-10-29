# Guía para la defensa

En la defensa los profes eligen una funcionalidad cualquiera y piden el recorrido completo:
**historia → spec → criterio de aceptación → escenario BDD → test → código Go**. Esta guía
tiene, para cada historia, dónde está cada paso, y abajo las preguntas que nos pueden hacer,
con una respuesta corta.

El recorrido completo de todos los escenarios, con los commits RED y GREEN de cada test, está
en [trazabilidad.md](trazabilidad.md). Si nos dejan elegir, mostramos la #13 (votar en secreto
en el Planning Poker): sus 16 escenarios tienen el ciclo RED → GREEN completo, y sus commits de
test traen la salida del RED.

## El recorrido de cada historia

| # | Historia | Spec y escenario principal | Test | Función Go |
|---|---|---|---|---|
| #3 | Crear y modificar proyectos | [03-crear-proyectos.md](specs/sprint-1/03-crear-proyectos.md?plain=1#L111): Crear un proyecto con todos los datos | `TestCrearUnProyectoConTodosLosDatos` [project_test.go:54](../internal/domain/project_test.go#L54) | `(*ProjectRegistry).Create` [project_registry.go:20](../internal/domain/project_registry.go#L20) |
| #4 | Registrar integrantes de un proyecto | [04-registrar-integrantes.md](specs/sprint-1/04-registrar-integrantes.md?plain=1#L120): Agregar el primer integrante | `TestAgregarElPrimerIntegrante` [member_test.go:45](../internal/domain/member_test.go#L45) | `(*Project).AddMember` [member.go:71](../internal/domain/member.go#L71) |
| #5 | Consultar el estado de un proyecto | [05-consultar-estado-proyecto.md](specs/sprint-2/05-consultar-estado-proyecto.md?plain=1#L123): Consultar el estado de un proyecto recien creado | `TestConsultarElEstadoDeUnProyectoRecienCreado` [project_status_test.go:47](../internal/domain/project_status_test.go#L47) | `(*Project).Status` [project_status.go:54](../internal/domain/project_status.go#L54) |
| #6 | Administrar el Product Backlog | [06-administrar-product-backlog.md](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L144): Crear una historia en el backlog | `TestCrearUnaHistoriaEnElBacklog` [story_test.go:251](../internal/domain/story_test.go#L251) | `(*Project).AddStory` [story.go:199](../internal/domain/story.go#L199) |
| #7 | Priorizar el Product Backlog | [07-priorizar-backlog.md](specs/sprint-2/07-priorizar-backlog.md?plain=1#L150): Listar el backlog ordenado por prioridad | `TestListarElBacklogOrdenadoPorPrioridad` [story_priority_test.go:152](../internal/domain/story_priority_test.go#L152) | `(*Project).PrioritizedBacklog` [story_priority.go:11](../internal/domain/story_priority.go#L11) |
| #8 | Crear un sprint con su Sprint Goal | [08-crear-sprint.md](specs/sprint-2/08-crear-sprint.md?plain=1#L159): Crear el primer sprint de un proyecto | `TestCrearElPrimerSprintDeUnProyecto` [sprint_test.go:29](../internal/domain/sprint_test.go#L29) | `(*Project).AddSprint` [sprint.go:126](../internal/domain/sprint.go#L126) |
| #9 | Asignar historias a un sprint | [09-asignar-historias-a-sprint.md](specs/sprint-2/09-asignar-historias-a-sprint.md?plain=1#L148): Asignar una historia estimada al sprint abierto | `TestAsignarUnaHistoriaEstimadaAlSprintAbierto` [sprint_backlog_test.go:54](../internal/domain/sprint_backlog_test.go#L54) | `(*Project).AssignStory` [sprint_backlog.go:6](../internal/domain/sprint_backlog.go#L6) |
| #10 | Cerrar un sprint | [10-cerrar-sprint.md](specs/sprint-3/10-cerrar-sprint.md?plain=1#L156): Cerrar el sprint registrando que se completo y que no | `TestCerrarElSprintRegistrandoQueSeCompletoYQueNo` [sprint_close_test.go:59](../internal/domain/sprint_close_test.go#L59) | `(*Project).CloseSprint` [sprint_close.go:11](../internal/domain/sprint_close.go#L11) |
| #11 | Consultar sprints anteriores | [11-consultar-sprints-anteriores.md](specs/sprint-3/11-consultar-sprints-anteriores.md?plain=1#L113): Listar los sprints cerrados de un proyecto | `TestListarLosSprintsCerradosDeUnProyecto` [sprint_query_test.go:61](../internal/domain/sprint_query_test.go#L61) | `(*Project).ClosedSprints` [sprint_query.go:8](../internal/domain/sprint_query.go#L8) |
| #12 | Estimar historias con Story Points | [12-estimar-story-points.md](specs/sprint-1/12-estimar-story-points.md?plain=1#L99): Estimar con cada valor de la escala (esquema del escenario) | `TestEstimarConCadaValorDeLaEscala` [story_estimation_test.go:25](../internal/domain/story_estimation_test.go#L25) | `(*Story).Estimate` [story_estimation.go:25](../internal/domain/story_estimation.go#L25) |
| #13 | Votar en secreto en una ronda de Planning Poker | [13-votar-en-secreto.md](specs/sprint-3/13-votar-en-secreto.md?plain=1#L188): Votar registra quien voto sin dejar ver los votos, ni siquiera el propio | `TestVotarRegistraQuienVotoSinDejarVerLosVotosNiSiquieraElPropio` [voting_round_test.go:161](../internal/domain/voting_round_test.go#L161) | `(*VotingRound).Votes` [voting_round.go:135](../internal/domain/voting_round.go#L135), que oculta los votos; el voto lo registra `(*VotingRound).Vote` [voting_round.go:101](../internal/domain/voting_round.go#L101) |
| #14 | Registrar el esfuerzo real | [14-registrar-esfuerzo-real.md](specs/sprint-2/14-registrar-esfuerzo-real.md?plain=1#L169): Registrar horas trabajadas en una historia | `TestRegistrarHorasTrabajadasEnUnaHistoria` [effort_test.go:42](../internal/domain/effort_test.go#L42) | `(*Project).LogEffort` [effort.go:59](../internal/domain/effort.go#L59) |
| #15 | Gestionar defectos | [15-gestionar-defectos.md](specs/sprint-2/15-gestionar-defectos.md?plain=1#L185): Registrar un defecto | `TestRegistrarUnDefecto` [defect_test.go:50](../internal/domain/defect_test.go#L50) | `(*Project).AddDefect` [defect.go:167](../internal/domain/defect.go#L167) |
| #16 | Calcular las metricas del proyecto | [16-calcular-metricas.md](specs/sprint-3/16-calcular-metricas.md?plain=1#L171): Metricas de un sprint cerrado | `TestMetricasDeUnSprintCerrado` [metrics_test.go:122](../internal/domain/metrics_test.go#L122) | `(*Project).SprintMetrics` [metrics.go:91](../internal/domain/metrics.go#L91) |
| #17 | Mostrar el dashboard | [17-mostrar-dashboard.md](specs/sprint-3/17-mostrar-dashboard.md?plain=1#L167): Dashboard de un proyecto con un sprint en curso | [pendiente] | [pendiente] |
| #18 | Generar el reporte de proyecto o sprint | [18-generar-reporte.md](specs/sprint-3/18-generar-reporte.md?plain=1#L144): Generar el reporte de un sprint cerrado | [pendiente] | [pendiente] |
| #19 | Exportar el reporte a PDF | [19-exportar-reporte-pdf.md](specs/sprint-4/19-exportar-reporte-pdf.md?plain=1#L128): Exportar el reporte de un sprint a PDF | [pendiente] | [pendiente] |
| #24 | Operar el sistema desde una interfaz de consola | [24-interfaz-de-consola.md](specs/sprint-1/24-interfaz-de-consola.md?plain=1#L103): Crear un proyecto desde la consola | `TestCrearUnProyectoDesdeLaConsola` [console_test.go:74](../cmd/app/console_test.go#L74) | `createProject` [console.go:97](../cmd/app/console.go#L97) (la despacha `runConsole` en `:64`; llama a `(*ProjectRegistry).Create`) |
| #73 | Resolver diferencias y acordar la estimacion | [73-resolver-diferencias-y-acordar.md](specs/sprint-4/73-resolver-diferencias-y-acordar.md?plain=1#L167): Acordar la estimacion de una ronda revelada (esquema del escenario) | `TestAcordarLaEstimacionDeUnaRondaRevelada` [voting_round_agreement_test.go:142](../internal/domain/voting_round_agreement_test.go#L142) | `(*VotingRound).Agree` [voting_round_agreement.go:67](../internal/domain/voting_round_agreement.go#L67) |
| #77 | Usar el sistema desde el navegador: proyectos, integrantes y backlog | [77-interfaz-web-base.md](specs/sprint-2/77-interfaz-web-base.md?plain=1#L156): Crear un proyecto desde la web | `TestCrearUnProyectoDesdeLaWeb` [projects_test.go:15](../internal/web/projects_test.go#L15) | `(*Server).createProject` [projects.go:77](../internal/web/projects.go#L77) (ruta `POST /proyectos/nuevo` en `internal/web/server.go:74`) |
| #78 | Operar sprints, esfuerzo, defectos, Planning Poker y reportes desde el navegador | [78-interfaz-web-operaciones.md](specs/sprint-3/78-interfaz-web-operaciones.md?plain=1#L224): Crear un sprint desde la web | `TestCrearUnSprintDesdeLaWeb` [sprints_test.go:60](../internal/web/sprints_test.go#L60) | `(*Server).createSprint` [sprints.go:60](../internal/web/sprints.go#L60) (ruta `POST /proyectos/{id}/sprints/nuevo` en `internal/web/server.go:95`) |
| #177 | Guardar los datos en una base de datos | [177-guardar-en-base-de-datos.md](specs/sprint-3/177-guardar-en-base-de-datos.md?plain=1#L192): Los datos siguen al cerrar y volver a abrir | `TestUnRegistroCompletoVuelveIgual` [database_test.go:209](../internal/store/database_test.go#L209) | `(*Database).Save` [database.go:93](../internal/store/database.go#L93), y lo vuelve a leer `(*Database).Load` [database.go:233](../internal/store/database.go#L233) |

Sacado de `dev` en `203859d` (26/10), y la #177 en `59bef23` (28/10), con cada línea
verificada. Las historias web (#77 y #78) entran por una ruta HTTP: la función es el handler, y
el handler llama al dominio. La #24 entra por la consola: `createProject` llama a
`(*ProjectRegistry).Create`, la misma función de la #3. La #12 y la #73 tienen un *Esquema del
escenario*, con una tabla de ejemplos que el test recorre con subtests. [pendiente: completar
la #17, la #18 y la #19 cuando se mergeen.]

## Cómo mostrarlo en vivo

1. Abrir la issue en GitHub y leer la historia ("Como… quiero… para…").
2. Abrir la spec en `docs/specs/` y mostrar las tres puertas aprobadas al final de cada fase.
3. Bajar a *Escenarios BDD* y leer el escenario de la tabla.
4. Abrir el test y mostrar el comentario `// Escenario:` arriba, con el mismo título.
5. Ir a la función Go que llama el test.
6. En [trazabilidad.md](trazabilidad.md), buscar el test y abrir sus dos commits: primero el
   `test:`, con la salida del RED en el mensaje, y después el `feat:`.

## Límites conocidos

Salieron en la prueba exploratoria de la web del 26/10 y quedaron así a propósito. Si los
preguntan, son decisiones, no defectos:

- **La fecha 01/01/0001 se rechaza como "obligatoria":** es el valor cero de las fechas en Go,
  y nadie la va a escribir de verdad.
- **Se acepta esfuerzo de un día anterior al inicio del proyecto:** la #14 no compara la fecha
  del esfuerzo con la del proyecto.
- **Se puede cerrar un sprint con una fecha de cierre cualquiera, incluso anterior a su
  inicio:** lo decidió la #10 (RN-11: la fecha de cierre no se compara con ninguna otra).
- **No hay un tope de horas por día sumando registros:** un registro no puede pasar las 24 h
  (defecto #192), pero dos registros del mismo día que juntos las pasan se aceptan. Lo decidió
  el equipo el 21/10 en la #14 (decisión 5).
- **No hay largo máximo para los nombres y los títulos:** se acepta un nombre de proyecto de
  3000 caracteres. Lo que sí se controla es que no estén vacíos.

Los defectos que salieron de esa misma prueba están cargados como issues: #190, #191, #192 y #193.

## Preguntas que nos pueden hacer

**1. ¿Por qué Go solo con la biblioteca estándar?**
El proyecto es chico y la cátedra evalúa el código que escribimos nosotros, así que no hacía
falta un framework. Los tests usan el paquete `testing` de Go, sin librerías aparte. La regla
está en `AGENTS.md`: agregar una dependencia se acuerda con el equipo. La única que se acordó es
el driver de SQLite de la #177.

**2. ¿Por qué SQLite?**
La base es un solo archivo, así que no hay que instalar un servidor para usar el sistema. El
driver (`modernc.org/sqlite`) está escrito en Go puro, sin C, así que compila igual en el CI,
que prueba en Linux y en Windows, y en las dos plataformas para las que el CD arma los binarios
(Linux amd64 y macOS arm64). El archivo JSON de antes quedó para importar y
exportar, y lo guardado con la versión anterior se importa solo al abrir con la base vacía.

**3. ¿Cómo se ve el TDD en el historial?**
Cada test entra en un commit `test:` que trae en el mensaje la salida del RED (el `--- FAIL` o
el error de compilación), y el código que lo hace pasar entra en el commit siguiente, `feat:`.
No se aplasta la historia ni se hace squash: en GitHub solo está habilitado el merge commit.
`scripts/trazabilidad.sh` recorre el historial y marca cada test. Cuando un test pasa de
entrada, porque una tarea anterior ya lo resolvía, el commit lo dice ("sin RED").

**4. ¿Qué es una puerta?**
La spec tiene cuatro fases: Requisitos, Diseño, Tareas e Implementación. Antes de pasar de una
a la siguiente, alguien distinto del que la escribió la revisa y la aprueba. Eso es una puerta,
y queda registrado en la spec quién la aprobó y cuándo. Así nadie programa sobre requisitos
que no revisó otro.

**5. ¿Cómo usaron la IA y cómo la validaron?**
La consigna la permite, y el equipo usó asistencia de IA. Lo que tenemos anotado:
- **Emiliano** implementó con asistencia de IA la #6, la #5, la #14, la #15, la #77, la #10, la
  #16, la #78, la #177 y la #19. También el arreglo del defecto #108 (PR #116), el del defecto
  #156 (PR #157) y los tests del arranque de la web (PR #146). La IA trabajó bajo su dirección
  y con su revisión.
- **Maxi** la usó en la #12, la #24, la #13 y la #73: la IA proponía el test y el código, y él
  aprobaba cada paso y cada commit.

Con IA o sin ella, todo siguió el mismo camino:
- **Puertas:** la spec tiene sus tres puertas aprobadas por otra persona, antes del código.
- **TDD:** primero el test, que falla y deja la salida del RED en el commit (RED), y después el
  código que lo hace pasar, en el commit siguiente (GREEN).
- **Revisión:** revisión escrita de cada PR y, desde el 26/10, la firma de un compañero.
- **CI:** los tests corren en Linux y en Windows en cada PR.

Cada uno responde por lo suyo, igual que si lo hubiera escrito a mano. En las historias que
implementó Emiliano, las horas reales son el tiempo de dirigir y revisar, y por eso el informe
muestra la desviación con ellas y sin ellas.

**6. ¿Cómo garantizan que los votos del Planning Poker son secretos?**
Mientras la ronda no se revela, `Votes()` no devuelve ningún voto, ni siquiera el propio: da el
error `ErrVotesHidden`. Solo se sabe quién votó y quién falta. Además, la ronda tiene su propio
`Format`, para que imprimirla con `%#v` (la forma de depuración de Go) no muestre los votos.
Eso apareció en una revisión, y se arregló con su test (#13).

**7. ¿Cómo calculan la velocidad y la desviación?**
La velocidad es la suma de los story points de las historias terminadas en el sprint: una
historia a medias no suma. La desviación es (horas reales − estimadas) ÷ estimadas, solo sobre
las historias terminadas que tienen horas estimadas. Las horas estimadas se cargan en el
tablero antes de empezar cada historia. Desde el Sprint 2 se estima a 2,7 h por punto, que es
lo que nos costó un punto en el Sprint 1. Todo sale del tablero con `scripts/metricas.sh` y está explicado en el
[informe de métricas](informe-de-metricas.md).

**8. ¿Por qué la desviación da tan negativa?**
Porque en las historias que implementó Emiliano con asistencia de IA las horas reales son
pocas: son solo el tiempo de dirigir y revisar. Por eso el informe da también la cuenta sin
ellas: en el Sprint 2, −52% con todas y −36% sin la #5, la #14 y la #15.

**9. ¿Qué pasa con una historia que no se termina en el sprint?**
No suma puntos y vuelve al Product Backlog. Una historia cuenta en el sprint en que se mergea a
`dev`. Lo que se mergea entre dos sprints cuenta en el siguiente, como la #6, que se mergeó el
19/10 y contó en el Sprint 2.

**10. ¿Cómo registran los defectos?**
Cada defecto es una issue con la etiqueta `defecto`. Dice qué pasa, cómo reproducirlo, a qué
historias afecta y en qué sprint se detectó y se resolvió. Se arregla en su propia rama, con un
test que primero falla y después pasa. Por ejemplo, el #156 salió en la revisión del PR de la
#73, y el #172 en la revisión de su arreglo. Los dos se resolvieron en el Sprint 3.

**11. ¿Cómo saben que cada escenario tiene su test?**
Cada test lleva arriba el comentario `// Escenario:` con el título del escenario.
`scripts/trazabilidad.sh` arma [trazabilidad.md](trazabilidad.md) cruzando las specs con los
tests y marca con ⚠️ los escenarios que no tienen test. El CI corre `scripts/coherencia.sh`,
que controla que las specs, el índice y las puertas digan lo mismo.

**12. ¿Qué cobertura tienen?**
Al 28/10 (`dev` en `59bef23`, ya con la #177), 95,6% en total y 100% en el dominio
(`internal/domain`), que es donde están las reglas de negocio. La más baja es `internal/store`,
el guardado en la base, con 87,1%. El CI la mide en cada PR con `go test ./... -coverprofile`.
El detalle por sprint está en el [informe de métricas](informe-de-metricas.md).

**13. ¿Por qué hay consola y web?**
La consola (#24) fue la primera interfaz. El 20/10 decidimos sumar una web (#77 y #78), porque
la consigna pide una interfaz usable y un dashboard con gráficos, y eso se muestra mucho mejor
en el navegador. La consola quedó con proyectos, integrantes, backlog y estimación, y lo nuevo
se hace desde la web.

**14. ¿Cómo se revisa un PR y quién lo mergea?**
Nadie pushea directo: todo entra por PR. A `dev` hace falta el CI en verde, con una revisión
escrita de cada PR y, desde el 26/10, la firma de un compañero. A `main` hace falta además una
aprobación. Solo se mergea con merge commit, así queda el historial del TDD.

**15. ¿Por qué algunas historias tienen ciclos de TDD en pocos minutos?**
Las más rápidas son historias hechas con asistencia de IA, las de la lista de la pregunta 5.
Por ejemplo, las 12 tareas de la #19 se commitearon en siete minutos (26/10, de 20:09 a 20:16).
Cada ciclo es real: el test se escribe y se corre antes que el código, y el commit `test:` trae
la salida del RED. Si se vuelve a ese commit y se corren los tests, fallan igual, y en el
`feat:` siguiente pasan. Maxi lo comprobó commit por commit en la revisión del PR #184. Lo que
lleva tiempo en esas historias no es escribir el código sino dirigir y revisar el trabajo, y eso
es lo que se carga como horas reales.
