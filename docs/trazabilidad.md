# Trazabilidad

_Generado por `scripts/trazabilidad.sh`. No editar a mano: se regenera con el script cada vez
que entra a `dev` un pull request que toca tests o specs._

En la defensa eligen una funcionalidad y piden el recorrido completo: **historia → spec →
criterios de aceptación → escenario BDD → test → código**. Acá está ese recorrido para cada
historia que ya tiene tests en el repositorio.

- **Escenarios y criterios**: salen de la Fase 1 de cada spec.
- **Test de cada escenario**: el comentario `// Escenario:` arriba del test, como pide
  `AGENTS.md`. Si el título del escenario está en más de una spec, el test va a la historia
  cuya tabla de trazabilidad lista su archivo.
- **Código de cada test**: sale del historial de TDD. *RED* es el commit que agregó el test y
  *GREEN* el commit siguiente de su rama que toca código; los archivos son los que cambió el
  GREEN.
- Las specs no dicen qué escenario prueba cada criterio, así que los criterios se listan
  junto a la historia, sin cruzarlos.

Los huecos se marcan con ⚠️ y no se tapan.

## Resumen

| Historia | Escenarios | Con test | ⚠️ Sin test | Tests sin escenario |
|---|---|---|---|---|
| [#3 Crear y modificar proyectos](#3-crear-y-modificar-proyectos) | 20 | 20 | 0 | 1 |
| [#4 Registrar integrantes de un proyecto](#4-registrar-integrantes-de-un-proyecto) | 13 | 13 | 0 | 11 |
| [#6 Administrar el Product Backlog](#6-administrar-el-product-backlog) | 18 | 18 | 0 | 13 |
| [#12 Estimar historias con Story Points](#12-estimar-historias-con-story-points) | 10 | 10 | 0 | 3 |
| [#24 Operar el sistema desde una interfaz de consola](#24-operar-el-sistema-desde-una-interfaz-de-consola) | 13 | 13 | 0 | 5 |
| [#5 Consultar el estado de un proyecto](#5-consultar-el-estado-de-un-proyecto) | 10 | 10 | 0 | 0 |
| [#7 Priorizar el Product Backlog](#7-priorizar-el-product-backlog) | 15 | 15 | 0 | 1 |
| [#8 Crear un sprint con su Sprint Goal](#8-crear-un-sprint-con-su-sprint-goal) | 17 | 17 | 0 | 0 |
| [#9 Asignar historias a un sprint](#9-asignar-historias-a-un-sprint) | 14 | 14 | 0 | 0 |
| [#14 Registrar el esfuerzo real](#14-registrar-el-esfuerzo-real) | 22 | 22 | 0 | 2 |
| [#15 Gestionar defectos](#15-gestionar-defectos) | 20 | 20 | 0 | 3 |
| [#77 Interfaz web: proyectos, integrantes y backlog](#77-interfaz-web-proyectos-integrantes-y-backlog) | 15 | 15 | 0 | 13 |
| [#10 Cerrar un sprint](#10-cerrar-un-sprint) | 14 | 14 | 0 | 0 |
| [#11 Consultar sprints anteriores](#11-consultar-sprints-anteriores) | 12 | 12 | 0 | 0 |
| [#13 Votar en secreto en una ronda de Planning Poker](#13-votar-en-secreto-en-una-ronda-de-planning-poker) | 16 | 16 | 0 | 0 |
| [#16 Calcular las métricas del proyecto y de cada sprint](#16-calcular-las-métricas-del-proyecto-y-de-cada-sprint) | 16 | 16 | 0 | 2 |
| [#17 Mostrar el dashboard](#17-mostrar-el-dashboard) | 16 | 16 | 0 | 2 |
| [#177 Guardar los datos en una base de datos](#177-guardar-los-datos-en-una-base-de-datos) | 14 | 14 | 0 | 2 |
| [#18 Generar el reporte de un proyecto o de un sprint](#18-generar-el-reporte-de-un-proyecto-o-de-un-sprint) | 16 | 16 | 0 | 1 |
| [#78 Interfaz web: sprints, esfuerzo, defectos, Planning Poker y reportes](#78-interfaz-web-sprints-esfuerzo-defectos-planning-poker-y-reportes) | 22 | 22 | 0 | 6 |
| [#19 Exportar el reporte a PDF](#19-exportar-el-reporte-a-pdf) | 13 | 13 | 0 | 4 |
| [#73 Resolver diferencias y acordar la estimación](#73-resolver-diferencias-y-acordar-la-estimación) | 13 | 13 | 0 | 0 |

## Qué dice el historial de TDD

Cada test tiene en la columna *RED → GREEN* lo que muestra el historial de commits. Hay
cinco casos posibles:

| Marca | Qué quiere decir | Tests |
|---|---|---|
| `abc1234` → `def5678` | El ciclo completo: el test entró en un commit y el código que lo hizo pasar, en el siguiente. | 327 |
| `abc1234` → `def5678` (el test trae lo justo para compilar) | También es el ciclo completo. El commit del test agrega lo mínimo para que compile (una constante, una firma) y falla; la implementación viene en el commit siguiente. | 3 |
| `abc1234` (sin RED: ya pasaba) | El test pasó con el código que ya estaba, y quien lo escribió lo dejó dicho en el mensaje del commit. | 69 |
| `abc1234` → ⚠️ sin GREEN | Después del test no hay un commit de código en su rama: pasó con lo que ya estaba. Puede ser que una tarea anterior ya cubriera el caso o que el código se haya escrito antes que el test; el historial no alcanza para saber cuál de las dos. | 13 |
| `abc1234` ⚠️ test y código en el mismo commit | El test y el código entraron juntos, así que el historial no muestra la fase en rojo. Es un desvío del TDD. | 1 |

## #3 Crear y modificar proyectos

- **Historia:** issue #3
- **Spec:** [docs/specs/sprint-1/03-crear-proyectos.md](specs/sprint-1/03-crear-proyectos.md)
- **Código según la spec:** `internal/domain/project.go`, `internal/domain/project_registry.go`, `internal/domain/errors.go`

<details><summary>Criterios de aceptación (12)</summary>

- Se puede crear un proyecto con nombre, descripción, fecha de inicio y fecha de fin, y queda registrado con un identificador propio.
- Se puede crear un proyecto con solo nombre y fecha de inicio.
- No se puede crear ni modificar un proyecto sin nombre.
- No se puede crear ni modificar un proyecto sin fecha de inicio.
- No se puede crear ni modificar un proyecto con fecha de fin anterior a la de inicio.
- Se pueden modificar los datos de un proyecto ya creado y los cambios quedan guardados.
- Una modificación rechazada deja el proyecto como estaba.
- Modificar un proyecto sin indicar un dato no borra ese dato.
- Se le puede quitar la descripción o la fecha de finalización a un proyecto pidiéndolo explícitamente.
- Dos proyectos creados con el mismo nombre reciben identificadores distintos.
- Modificar un proyecto inexistente informa el error y no crea nada.
- Consultar un proyecto inexistente informa que no se encontró.

</details>

| Escenario BDD | Test | Código | RED → GREEN |
|---|---|---|---|
| [Crear un proyecto con todos los datos](specs/sprint-1/03-crear-proyectos.md?plain=1#L111) | `TestCrearUnProyectoConTodosLosDatos` [project_test.go:54](../internal/domain/project_test.go#L54) | [project.go](../internal/domain/project.go), [project_registry.go](../internal/domain/project_registry.go) | `37e4f2f` → `0fecae8` |
| [Crear un proyecto solo con los datos obligatorios](specs/sprint-1/03-crear-proyectos.md?plain=1#L117) | `TestCrearUnProyectoSoloConLosDatosObligatorios` [project_test.go:91](../internal/domain/project_test.go#L91) | [project.go](../internal/domain/project.go) | `11c7ed0` → `110c654` |
| [Rechazar un proyecto sin nombre](specs/sprint-1/03-crear-proyectos.md?plain=1#L122) | `TestRechazarUnProyectoSinNombre` [project_test.go:164](../internal/domain/project_test.go#L164) | [errors.go](../internal/domain/errors.go), [project_registry.go](../internal/domain/project_registry.go) | `1e45818` → `13e9e68` |
| [Rechazar un proyecto sin fecha de inicio](specs/sprint-1/03-crear-proyectos.md?plain=1#L132) | `TestRechazarUnProyectoSinFechaDeInicio` [project_test.go:197](../internal/domain/project_test.go#L197) | [errors.go](../internal/domain/errors.go), [project_registry.go](../internal/domain/project_registry.go) | `b03aa4a` → `2c40255` |
| [Guardar el nombre sin espacios en los extremos](specs/sprint-1/03-crear-proyectos.md?plain=1#L137) | `TestGuardarElNombreSinEspaciosEnLosExtremos` [project_test.go:134](../internal/domain/project_test.go#L134) | [project_registry.go](../internal/domain/project_registry.go) | `e9c0d69` → `2d6f386` |
| [Aceptar un nombre de un solo caracter](specs/sprint-1/03-crear-proyectos.md?plain=1#L141) | `TestAceptarUnNombreDeUnSoloCaracter` [project_test.go:149](../internal/domain/project_test.go#L149) | [project_registry.go](../internal/domain/project_registry.go) | `e9c0d69` → `2d6f386` |
| [Dos proyectos con el mismo nombre](specs/sprint-1/03-crear-proyectos.md?plain=1#L145) | `TestDosProyectosConElMismoNombre` [project_test.go:112](../internal/domain/project_test.go#L112) | el que ya estaba | `239e570` (sin RED: ya pasaba) |
| [Rechazar fechas incoherentes](specs/sprint-1/03-crear-proyectos.md?plain=1#L151) | `TestRechazarFechasIncoherentes` [project_test.go:209](../internal/domain/project_test.go#L209) | [errors.go](../internal/domain/errors.go), [project_registry.go](../internal/domain/project_registry.go) | `da19160` → `f6a0ffb` |
| [Aceptar un proyecto que empieza y termina el mismo dia](specs/sprint-1/03-crear-proyectos.md?plain=1#L156) | `TestAceptarUnProyectoQueEmpiezaYTerminaElMismoDia` [project_test.go:225](../internal/domain/project_test.go#L225) | [errors.go](../internal/domain/errors.go), [project_registry.go](../internal/domain/project_registry.go) | `da19160` → `f6a0ffb` |
| [Comparar las fechas por dia, sin importar la hora](specs/sprint-1/03-crear-proyectos.md?plain=1#L160) | `TestCompararLasFechasPorDiaSinImportarLaHora` [project_test.go:244](../internal/domain/project_test.go#L244) | [project.go](../internal/domain/project.go) | `21325fb` → `9ba417c` |
| [Modificar solo el nombre conserva los demas datos](specs/sprint-1/03-crear-proyectos.md?plain=1#L165) | `TestModificarSoloElNombreConservaLosDemasDatos` [project_test.go:319](../internal/domain/project_test.go#L319) | [project.go](../internal/domain/project.go), [project_registry.go](../internal/domain/project_registry.go) | `3bdd1c3` → `3c85a8c` |
| [Guardar sin espacios en los extremos el nombre modificado](specs/sprint-1/03-crear-proyectos.md?plain=1#L172) | `TestGuardarSinEspaciosEnLosExtremosElNombreModificado` [project_test.go:347](../internal/domain/project_test.go#L347) | [project.go](../internal/domain/project.go), [project_registry.go](../internal/domain/project_registry.go) | `732ebf0` → `66d9474` |
| [Modificar un proyecto sin cambiar nada](specs/sprint-1/03-crear-proyectos.md?plain=1#L177) | `TestModificarUnProyectoSinCambiarNada` [project_test.go:333](../internal/domain/project_test.go#L333) | [project.go](../internal/domain/project.go), [project_registry.go](../internal/domain/project_registry.go) | `3bdd1c3` → `3c85a8c` |
| [Rechazar una modificacion que deja el proyecto sin nombre](specs/sprint-1/03-crear-proyectos.md?plain=1#L184) | `TestRechazarUnaModificacionQueDejaElProyectoSinNombre` [project_test.go:363](../internal/domain/project_test.go#L363) | [project.go](../internal/domain/project.go), [project_registry.go](../internal/domain/project_registry.go) | `732ebf0` → `66d9474` |
| [Rechazar una fecha de inicio posterior a la de finalizacion al modificar](specs/sprint-1/03-crear-proyectos.md?plain=1#L191) | `TestRechazarUnaFechaDeInicioPosteriorALaDeFinalizacionAlModificar` [project_test.go:377](../internal/domain/project_test.go#L377) | [project.go](../internal/domain/project.go), [project_registry.go](../internal/domain/project_registry.go) | `732ebf0` → `66d9474` |
| [Quitar la fecha de finalizacion](specs/sprint-1/03-crear-proyectos.md?plain=1#L198) | `TestQuitarLaFechaDeFinalizacion` [project_test.go:391](../internal/domain/project_test.go#L391) | [project.go](../internal/domain/project.go), [project_registry.go](../internal/domain/project_registry.go) | `9e21841` → `587028c` |
| [Quitar la descripcion](specs/sprint-1/03-crear-proyectos.md?plain=1#L204) | `TestQuitarLaDescripcion` [project_test.go:408](../internal/domain/project_test.go#L408) | [project.go](../internal/domain/project.go), [project_registry.go](../internal/domain/project_registry.go) | `9e21841` → `587028c` |
| [Modificar la descripcion y la fecha de finalizacion](specs/sprint-1/03-crear-proyectos.md?plain=1#L211) | `TestModificarLaDescripcionYLaFechaDeFinalizacion` [project_test.go:422](../internal/domain/project_test.go#L422) | [project.go](../internal/domain/project.go), [project_registry.go](../internal/domain/project_registry.go) | `d35d276` → `9e5f78b` |
| [Modificar un proyecto que no existe](specs/sprint-1/03-crear-proyectos.md?plain=1#L218) | `TestModificarUnProyectoQueNoExiste` [project_test.go:296](../internal/domain/project_test.go#L296) | [errors.go](../internal/domain/errors.go), [project.go](../internal/domain/project.go), [project_registry.go](../internal/domain/project_registry.go) | `bcf6cb1` → `45d17a8` |
| [Consultar un proyecto que no existe](specs/sprint-1/03-crear-proyectos.md?plain=1#L224) | `TestConsultarUnProyectoQueNoExiste` [project_test.go:284](../internal/domain/project_test.go#L284) | [errors.go](../internal/domain/errors.go), [project.go](../internal/domain/project.go), [project_registry.go](../internal/domain/project_registry.go) | `bcf6cb1` → `45d17a8` |

**Tests que no citan un escenario de la Fase 1 (1).**

| Test | Cubre | Código | RED → GREEN |
|---|---|---|---|
| `TestLosMensajesDeErrorLlevanTildes` [errors_test.go:12](../internal/domain/errors_test.go#L12) | los mensajes de la RN-4 y la RN-6 de la #3, y de la RN-5 y la RN-11 de la #4 | [errors.go](../internal/domain/errors.go) | `73d0160` → `d2ff81b` |

## #4 Registrar integrantes de un proyecto

- **Historia:** issue #4
- **Spec:** [docs/specs/sprint-1/04-registrar-integrantes.md](specs/sprint-1/04-registrar-integrantes.md)
- **Código según la spec:** `internal/domain/member.go`, `internal/domain/project.go`, `internal/domain/errors.go`

<details><summary>Criterios de aceptación (12)</summary>

- Se puede agregar un integrante indicando su nombre y apellido y su rol.
- No se puede agregar un integrante sin nombre y apellido.
- No se puede agregar un integrante con un rol distinto de Agile Enabler o Product Builder.
- Se pueden listar los integrantes de un proyecto, con su rol.
- Listar los integrantes de un proyecto sin integrantes devuelve una lista vacía.
- No se puede cargar dos veces el mismo integrante en un proyecto, aunque cambien las mayúsculas o los espacios de los extremos.
- El mismo integrante se puede cargar en dos proyectos distintos.
- No se puede cargar un segundo Agile Enabler en el mismo proyecto.
- Se puede quitar un integrante del proyecto.
- Quitar un integrante que no está en el proyecto informa el error.
- Operar sobre un proyecto inexistente informa que no se encontró.
- Una operación rechazada deja los integrantes del proyecto como estaban.

</details>

| Escenario BDD | Test | Código | RED → GREEN |
|---|---|---|---|
| [Agregar el primer integrante](specs/sprint-1/04-registrar-integrantes.md?plain=1#L120) | `TestAgregarElPrimerIntegrante` [member_test.go:45](../internal/domain/member_test.go#L45) | [member.go](../internal/domain/member.go) | `8398ab5` → `c3ed475` |
| [Rechazar un integrante sin nombre y apellido](specs/sprint-1/04-registrar-integrantes.md?plain=1#L126) | `TestRechazarUnIntegranteSinNombreYApellido` [member_test.go:95](../internal/domain/member_test.go#L95) | [errors.go](../internal/domain/errors.go), [member.go](../internal/domain/member.go) | `27fa739` → `2f225f6` |
| [Guardar el nombre sin espacios en los extremos](specs/sprint-1/04-registrar-integrantes.md?plain=1#L138) | `TestGuardarElNombreDelIntegranteSinEspaciosEnLosExtremos` [member_test.go:73](../internal/domain/member_test.go#L73) | [member.go](../internal/domain/member.go) | `1821eb3` → `ad6d7df` |
| [Rechazar un rol invalido](specs/sprint-1/04-registrar-integrantes.md?plain=1#L143) | `TestRechazarUnRolInvalido` [member_test.go:123](../internal/domain/member_test.go#L123) | [member.go](../internal/domain/member.go) | `0216054` → `2e8beef` |
| [Rechazar un integrante duplicado](specs/sprint-1/04-registrar-integrantes.md?plain=1#L150) | `TestRechazarUnIntegranteDuplicado` [member_test.go:172](../internal/domain/member_test.go#L172) | [errors.go](../internal/domain/errors.go), [member.go](../internal/domain/member.go) | `598a735` → `db034c1` |
| [La misma persona en dos proyectos](specs/sprint-1/04-registrar-integrantes.md?plain=1#L163) | `TestLaMismaPersonaEnDosProyectos` [member_test.go:247](../internal/domain/member_test.go#L247) | ⚠️ ningún commit de código después del test | `ac05cb0` → ⚠️ sin GREEN |
| [Rechazar un segundo Agile Enabler](specs/sprint-1/04-registrar-integrantes.md?plain=1#L170) | `TestRechazarUnSegundoAgileEnabler` [member_test.go:276](../internal/domain/member_test.go#L276) | [errors.go](../internal/domain/errors.go), [member.go](../internal/domain/member.go) | `5432431` → `1da3059` |
| [Agregar un Agile Enabler despues de quitar al anterior](specs/sprint-1/04-registrar-integrantes.md?plain=1#L177) | `TestAgregarUnAgileEnablerDespuesDeQuitarAlAnterior` [member_test.go:424](../internal/domain/member_test.go#L424) | [member.go](../internal/domain/member.go) | `17736ed` → `c3b82b7` |
| [Listar los integrantes de un proyecto](specs/sprint-1/04-registrar-integrantes.md?plain=1#L183) | `TestListarLosIntegrantesDeUnProyecto` [member_test.go:485](../internal/domain/member_test.go#L485) | ⚠️ ningún commit de código después del test | `5b56329` → ⚠️ sin GREEN |
| [Listar un proyecto sin integrantes](specs/sprint-1/04-registrar-integrantes.md?plain=1#L189) | `TestListarUnProyectoSinIntegrantes` [member_test.go:29](../internal/domain/member_test.go#L29) | [member.go](../internal/domain/member.go), [project.go](../internal/domain/project.go) | `008b6ef` → `cc89760` |
| [Quitar el ultimo integrante](specs/sprint-1/04-registrar-integrantes.md?plain=1#L194) | `TestQuitarElUltimoIntegrante` [member_test.go:371](../internal/domain/member_test.go#L371) | [member.go](../internal/domain/member.go) | `17736ed` → `c3b82b7` |
| [Quitar un integrante que no pertenece al proyecto](specs/sprint-1/04-registrar-integrantes.md?plain=1#L199) | `TestQuitarUnIntegranteQueNoPerteneceAlProyecto` [member_test.go:453](../internal/domain/member_test.go#L453) | [errors.go](../internal/domain/errors.go), [member.go](../internal/domain/member.go) | `a29c335` → `102ee66` |
| [Operar sobre un proyecto que no existe](specs/sprint-1/04-registrar-integrantes.md?plain=1#L205) | `TestOperarSobreIntegrantesDeUnProyectoQueNoExiste` [member_test.go:557](../internal/domain/member_test.go#L557) | ⚠️ ningún commit de código después del test | `3e3d580` → ⚠️ sin GREEN |

**Tests que no citan un escenario de la Fase 1 (11).**

| Test | Cubre | Código | RED → GREEN |
|---|---|---|---|
| `TestUnIntegranteExponeSuNombreYSuRol` [member_test.go:13](../internal/domain/member_test.go#L13) | RN-1 | [member.go](../internal/domain/member.go) | `61fc94d` → `adf4e3a` |
| `TestUnRolSinInicializarNoEsNingunoDeLosValidos` [member_test.go:156](../internal/domain/member_test.go#L156) | RN-4 | [member.go](../internal/domain/member.go) | `0216054` → `2e8beef` |
| `TestNoSonDuplicadosLosNombresQueDifierenEnTildesOEspaciosDelMedio` [member_test.go:213](../internal/domain/member_test.go#L213) | RN-5 | [errors.go](../internal/domain/errors.go), [member.go](../internal/domain/member.go) | `598a735` → `db034c1` |
| `TestConUnAgileEnablerSePuedenAgregarProductBuilders` [member_test.go:301](../internal/domain/member_test.go#L301) | RN-7 | [errors.go](../internal/domain/errors.go), [member.go](../internal/domain/member.go) | `5432431` → `1da3059` |
| `TestConVariosErroresSeInformaElPrimero` [member_test.go:327](../internal/domain/member_test.go#L327) | RN-2, RN-4, RN-5, RN-7, RN-12 | ⚠️ ningún commit de código después del test | `c95cd0c` → ⚠️ sin GREEN |
| `TestQuitarUnIntegranteDejaALosDemas` [member_test.go:400](../internal/domain/member_test.go#L400) | RN-9 | [member.go](../internal/domain/member.go) | `17736ed` → `c3b82b7` |
| `TestCambiarLaListaRecibidaNoCambiaLosIntegrantesDelProyecto` [member_test.go:523](../internal/domain/member_test.go#L523) | RN-8 | ⚠️ ningún commit de código después del test | `5b56329` → ⚠️ sin GREEN |
| `TestConvertirElTextoEnUnRol` [member_test.go:605](../internal/domain/member_test.go#L605) | RN-4 | [errors.go](../internal/domain/errors.go), [member.go](../internal/domain/member.go) | `1d5d160` → `0244f63` |
| `TestRechazarUnTextoQueNoEsUnRol` [member_test.go:635](../internal/domain/member_test.go#L635) | RN-4 | [errors.go](../internal/domain/errors.go), [member.go](../internal/domain/member.go) | `1d5d160` → `0244f63` |
| `TestElErrorDeRolNombraLosValidos` [member_test.go:653](../internal/domain/member_test.go#L653) | RN-4 | [errors.go](../internal/domain/errors.go), [member.go](../internal/domain/member.go) | `1d5d160` → `0244f63` |
| `TestLosDosRolesSeMuestranYSeVuelvenAConvertir` [member_test.go:670](../internal/domain/member_test.go#L670) | RN-4 | [errors.go](../internal/domain/errors.go), [member.go](../internal/domain/member.go) | `1d5d160` → `0244f63` |

## #6 Administrar el Product Backlog

- **Historia:** issue #6
- **Spec:** [docs/specs/sprint-1/06-administrar-product-backlog.md](specs/sprint-1/06-administrar-product-backlog.md)
- **Código según la spec:** `internal/domain/story.go`, `internal/domain/errors.go`

<details><summary>Criterios de aceptación (13)</summary>

- Se puede crear una historia con título, descripción, prioridad y criterios de aceptación, y queda con identificador propio y estado Pendiente.
- El identificador no se repite dentro del mismo proyecto.
- No se puede crear una historia sin título, sin descripción, sin prioridad o sin criterios de aceptación.
- No se puede crear una historia con una prioridad que no sea Alta, Media o Baja.
- Una historia recién creada queda sin estimar.
- Se puede modificar una historia mientras no esté Terminada.
- No se puede modificar una historia Terminada.
- Se puede pedir el backlog completo de un proyecto.
- El backlog de un proyecto sin historias se devuelve vacío.
- Un criterio de aceptación vacío se descarta, y si no queda ninguno la operación se rechaza.
- Una operación rechazada deja la historia y el backlog como estaban.
- Modificar una historia inexistente informa el error y no crea ninguna.
- Operar sobre un proyecto inexistente informa que no se encontró.

</details>

| Escenario BDD | Test | Código | RED → GREEN |
|---|---|---|---|
| [Crear una historia en el backlog](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L144) | `TestCrearUnaHistoriaEnElBacklog` [story_test.go:251](../internal/domain/story_test.go#L251) | [project.go](../internal/domain/project.go), [project_registry.go](../internal/domain/project_registry.go), [story.go](../internal/domain/story.go) | `81382b9` → `3e5a29c` |
| [Los identificadores no se repiten](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L151) | `TestLosIdentificadoresDeLasHistoriasNoSeRepiten` [story_test.go:297](../internal/domain/story_test.go#L297) | [project.go](../internal/domain/project.go), [project_registry.go](../internal/domain/project_registry.go), [story.go](../internal/domain/story.go) | `81382b9` → `3e5a29c` |
| [Rechazar una historia sin criterios de aceptacion](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L156) | `TestRechazarUnaHistoriaSinLosDatosObligatorios` [story_test.go:342](../internal/domain/story_test.go#L342) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `d2e8d9f` → `de00fc7` |
| [Rechazar una prioridad invalida](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L162) | `TestRechazarUnaPrioridadInvalida` [story_test.go:138](../internal/domain/story_test.go#L138) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `4dc7d45` → `400a285` |
| ″ | `TestRechazarUnaHistoriaSinLosDatosObligatorios` [story_test.go:342](../internal/domain/story_test.go#L342) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `d2e8d9f` → `de00fc7` |
| [Modificar una historia pendiente](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L168) | `TestModificarUnaHistoria` [story_test.go:397](../internal/domain/story_test.go#L397) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `2bae7e0` → `13f6ad5` |
| [No modificar una historia terminada](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L173) | `TestNoModificarUnaHistoriaTerminada` [story_test.go:487](../internal/domain/story_test.go#L487) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `2bae7e0` → `13f6ad5` |
| [Backlog de un proyecto sin historias](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L179) | `TestBacklogDeUnProyectoSinHistorias` [story_test.go:235](../internal/domain/story_test.go#L235) | [project.go](../internal/domain/project.go), [project_registry.go](../internal/domain/project_registry.go), [story.go](../internal/domain/story.go) | `81382b9` → `3e5a29c` |
| [Rechazar una historia sin los datos obligatorios](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L184) | `TestRechazarUnaHistoriaSinLosDatosObligatorios` [story_test.go:342](../internal/domain/story_test.go#L342) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `d2e8d9f` → `de00fc7` |
| [Un texto con solo espacios cuenta como vacio](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L198) | `TestRechazarUnaHistoriaSinLosDatosObligatorios` [story_test.go:342](../internal/domain/story_test.go#L342) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `d2e8d9f` → `de00fc7` |
| [Descartar los criterios de aceptacion vacios](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L204) | `TestNormalizarLosDatosDeUnaHistoria` [story_test.go:311](../internal/domain/story_test.go#L311) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `d2e8d9f` → `de00fc7` |
| [Rechazar si todos los criterios quedan vacios](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L209) | `TestRechazarUnaHistoriaSinLosDatosObligatorios` [story_test.go:342](../internal/domain/story_test.go#L342) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `d2e8d9f` → `de00fc7` |
| [Guardar el titulo sin espacios en los extremos](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L215) | `TestNormalizarLosDatosDeUnaHistoria` [story_test.go:311](../internal/domain/story_test.go#L311) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `d2e8d9f` → `de00fc7` |
| [Una modificacion rechazada no cambia nada](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L220) | `TestUnaModificacionRechazadaNoCambiaNada` [story_test.go:470](../internal/domain/story_test.go#L470) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `2bae7e0` → `13f6ad5` |
| [Modificar una historia que no existe](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L226) | `TestModificarUnaHistoriaQueNoExiste` [story_test.go:507](../internal/domain/story_test.go#L507) | ⚠️ ningún commit de código después del test | `499eb03` → ⚠️ sin GREEN |
| [Reemplazar los criterios de aceptacion](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L232) | `TestReemplazarLosCriteriosDeAceptacion` [story_test.go:453](../internal/domain/story_test.go#L453) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `2bae7e0` → `13f6ad5` |
| [Los identificadores son unicos dentro de cada proyecto](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L237) | `TestLosIdentificadoresSonUnicosDentroDeCadaProyecto` [story_test.go:280](../internal/domain/story_test.go#L280) | [project.go](../internal/domain/project.go), [project_registry.go](../internal/domain/project_registry.go), [story.go](../internal/domain/story.go) | `81382b9` → `3e5a29c` |
| [Pedir el backlog de un proyecto con historias](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L243) | `TestPedirElBacklogDeUnProyectoConHistorias` [story_test.go:540](../internal/domain/story_test.go#L540) | ⚠️ ningún commit de código después del test | `499eb03` → ⚠️ sin GREEN |
| ″ | `TestCambiarElBacklogRecibidoNoCambiaElDelProyecto` [story_test.go:572](../internal/domain/story_test.go#L572) | ⚠️ ningún commit de código después del test | `499eb03` → ⚠️ sin GREEN |
| [Operar sobre un proyecto que no existe](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L250) | `TestOperarSobreUnProyectoQueNoExiste` [story_test.go:528](../internal/domain/story_test.go#L528) | ⚠️ ningún commit de código después del test | `499eb03` → ⚠️ sin GREEN |

**Tests que no citan un escenario de la Fase 1 (13).**

| Test | Cubre | Código | RED → GREEN |
|---|---|---|---|
| `TestUnaHistoriaExponeTodosSusDatos` [story_test.go:14](../internal/domain/story_test.go#L14) | RN-2 | [story.go](../internal/domain/story.go) | `bde9c51` → `aded90a` |
| `TestCambiarLosCriteriosRecibidosNoCambiaLosDeLaHistoria` [story_test.go:53](../internal/domain/story_test.go#L53) | RN-2 | [story.go](../internal/domain/story.go) | `bde9c51` → `aded90a` |
| `TestLosTresEstadosDeUnaHistoriaSeDistinguen` [story_test.go:69](../internal/domain/story_test.go#L69) | RN-2 | [story.go](../internal/domain/story.go) | `bde9c51` → `aded90a` |
| `TestConvertirElTextoEnUnaPrioridad` [story_test.go:93](../internal/domain/story_test.go#L93) | RN-4 | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `4dc7d45` → `400a285` |
| `TestUnaPrioridadQueNoExisteNoSeMuestraComoAlta` [story_test.go:124](../internal/domain/story_test.go#L124) | RN-4 | [story.go](../internal/domain/story.go) | `0d22bc7` ⚠️ test y código en el mismo commit |
| `TestElErrorDePrioridadNombraLasValidas` [story_test.go:156](../internal/domain/story_test.go#L156) | RN-4 | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `4dc7d45` → `400a285` |
| `TestLasTresPrioridadesSeMuestranYSeVuelvenAConvertir` [story_test.go:170](../internal/domain/story_test.go#L170) | RN-4 | ⚠️ ningún commit de código después del test | `1d7326d` → ⚠️ sin GREEN |
| `TestUnaHistoriaDevuelveSusStoryPoints` [story_test.go:193](../internal/domain/story_test.go#L193) | RN-2 | ⚠️ ningún commit de código después del test | `1d7326d` → ⚠️ sin GREEN |
| `TestDosHistoriasDelMismoProyectoPuedenTenerElMismoTitulo` [story_test.go:378](../internal/domain/story_test.go#L378) | RN-14 | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `d2e8d9f` → `de00fc7` |
| `TestModificarUnaHistoriaEnCurso` [story_test.go:419](../internal/domain/story_test.go#L419) | RN-8 | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `2bae7e0` → `13f6ad5` |
| `TestModificarUnaHistoriaSinCambiarNada` [story_test.go:437](../internal/domain/story_test.go#L437) | RN-9 | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `2bae7e0` → `13f6ad5` |
| `TestBuscarUnaHistoriaDelBacklog` [story_test.go:588](../internal/domain/story_test.go#L588) | RN-13 | ⚠️ ningún commit de código después del test | `499eb03` → ⚠️ sin GREEN |
| `TestModificarLaDescripcionDeUnaHistoria` [story_test.go:606](../internal/domain/story_test.go#L606) | RN-9 | ⚠️ ningún commit de código después del test | `e946fda` → ⚠️ sin GREEN |

## #12 Estimar historias con Story Points

- **Historia:** issue #12
- **Spec:** [docs/specs/sprint-1/12-estimar-story-points.md](specs/sprint-1/12-estimar-story-points.md)
- **Código según la spec:** `internal/domain/story_estimation.go`, `internal/domain/errors.go`

<details><summary>Criterios de aceptación (8)</summary>

- Se puede asignar story points a una historia con cualquiera de los valores 1, 2, 3, 5, 8, 13 o 21.
- No se puede asignar un valor fuera de esa escala, por ejemplo 4 o 7.
- No se puede asignar cero ni un valor negativo.
- Una historia sin estimar figura como no estimada, y eso no es un error.
- Se puede reestimar una historia mientras no esté Terminada, y queda el último valor.
- No se puede cambiar la estimación de una historia Terminada.
- Se le puede quitar la estimación a una historia no terminada.
- Estimar una historia que no existe informa que no se encontró.

</details>

| Escenario BDD | Test | Código | RED → GREEN |
|---|---|---|---|
| [Estimar con cada valor de la escala](specs/sprint-1/12-estimar-story-points.md?plain=1#L99) | `TestEstimarConCadaValorDeLaEscala` [story_estimation_test.go:25](../internal/domain/story_estimation_test.go#L25) | [story_estimation.go](../internal/domain/story_estimation.go) | `bf8961a` → `e343aed` |
| [Rechazar valores fuera de la escala](specs/sprint-1/12-estimar-story-points.md?plain=1#L114) | `TestRechazarValoresFueraDeLaEscala` [story_estimation_test.go:44](../internal/domain/story_estimation_test.go#L44) | [errors.go](../internal/domain/errors.go), [story_estimation.go](../internal/domain/story_estimation.go) | `5fc534d` → `83ee6a3` |
| [Consultar una historia sin estimar](specs/sprint-1/12-estimar-story-points.md?plain=1#L128) | `TestConsultarUnaHistoriaSinEstimar` [story_estimation_test.go:15](../internal/domain/story_estimation_test.go#L15) | el que ya estaba | `fba8844` (sin RED: ya pasaba) |
| [Reestimar una historia](specs/sprint-1/12-estimar-story-points.md?plain=1#L133) | `TestReestimarUnaHistoria` [story_estimation_test.go:80](../internal/domain/story_estimation_test.go#L80) | el que ya estaba | `46299d2` (sin RED: ya pasaba) |
| [Reestimar con el mismo valor](specs/sprint-1/12-estimar-story-points.md?plain=1#L138) | `TestReestimarConElMismoValor` [story_estimation_test.go:95](../internal/domain/story_estimation_test.go#L95) | el que ya estaba | `695b149` (sin RED: ya pasaba) |
| [Reestimar una historia en curso](specs/sprint-1/12-estimar-story-points.md?plain=1#L143) | `TestReestimarUnaHistoriaEnCurso` [story_estimation_test.go:110](../internal/domain/story_estimation_test.go#L110) | el que ya estaba | `47fd4b0` (sin RED: ya pasaba) |
| [No reestimar una historia terminada](specs/sprint-1/12-estimar-story-points.md?plain=1#L148) | `TestNoReestimarUnaHistoriaTerminada` [story_estimation_test.go:155](../internal/domain/story_estimation_test.go#L155) | [errors.go](../internal/domain/errors.go), [story_estimation.go](../internal/domain/story_estimation.go) | `1b01841` → `ed351db` |
| [Quitar la estimacion](specs/sprint-1/12-estimar-story-points.md?plain=1#L155) | `TestQuitarLaEstimacion` [story_estimation_test.go:125](../internal/domain/story_estimation_test.go#L125) | [story_estimation.go](../internal/domain/story_estimation.go) | `4375c8b` → `8d706ea` |
| [No quitar la estimacion de una historia terminada](specs/sprint-1/12-estimar-story-points.md?plain=1#L160) | `TestNoQuitarLaEstimacionDeUnaHistoriaTerminada` [story_estimation_test.go:181](../internal/domain/story_estimation_test.go#L181) | [story_estimation.go](../internal/domain/story_estimation.go) | `d305ca8` → `2dedbc5` |
| [Estimar una historia que no existe](specs/sprint-1/12-estimar-story-points.md?plain=1#L166) | `TestEstimarUnaHistoriaQueNoExiste` [story_estimation_test.go:196](../internal/domain/story_estimation_test.go#L196) | [story_estimation.go](../internal/domain/story_estimation.go) | `20e302e` → `e33c7bd` |

**Tests que no citan un escenario de la Fase 1 (3).**

| Test | Cubre | Código | RED → GREEN |
|---|---|---|---|
| `TestQuitarLaEstimacionDeUnaHistoriaYaSinEstimar` [story_estimation_test.go:140](../internal/domain/story_estimation_test.go#L140) | RN-7 | el que ya estaba | `ef2fd72` (sin RED: ya pasaba) |
| `TestQuitarLaEstimacionDeUnaHistoriaInexistente` [story_estimation_test.go:208](../internal/domain/story_estimation_test.go#L208) | RN-9 | [story_estimation.go](../internal/domain/story_estimation.go) | `f4e9086` → `4e93db8` |
| `TestConMasDeUnProblemaSeInformaElPrimero` [story_estimation_test.go:220](../internal/domain/story_estimation_test.go#L220) | RN-8, RN-9 | el que ya estaba | `ae2a6cc` (sin RED: ya pasaba) |

## #24 Operar el sistema desde una interfaz de consola

- **Historia:** issue #24
- **Spec:** [docs/specs/sprint-1/24-interfaz-de-consola.md](specs/sprint-1/24-interfaz-de-consola.md)
- **Código según la spec:** `cmd/app/console.go`, `cmd/app/main.go`

<details><summary>Criterios de aceptación (9)</summary>

- Se puede crear un proyecto desde la consola y se ve su identificador.
- Se pueden agregar integrantes a un proyecto y listarlos.
- Se pueden cargar historias en el backlog y listarlas.
- Se le pueden asignar story points a una historia.
- Cuando una operación se rechaza, se ve el motivo y el programa sigue andando.
- Listar algo vacío no es un error.
- Una operación incompleta o con un dato mal escrito se informa sin ejecutarla.
- Se puede salir del programa.
- La consola no contiene ninguna validación de negocio propia.

</details>

| Escenario BDD | Test | Código | RED → GREEN |
|---|---|---|---|
| [Crear un proyecto desde la consola](specs/sprint-1/24-interfaz-de-consola.md?plain=1#L103) | `TestCrearUnProyectoDesdeLaConsola` [console_test.go:74](../cmd/app/console_test.go#L74) | [console.go](../cmd/app/console.go) | `799f9e8` → `07c23ff` (el test trae lo justo para compilar) |
| [Mostrar el motivo de un rechazo](specs/sprint-1/24-interfaz-de-consola.md?plain=1#L109) | `TestMostrarElMotivoDeUnRechazo` [console_test.go:106](../cmd/app/console_test.go#L106) | [console.go](../cmd/app/console.go) | `b1bf255` → `4cd8d97` |
| [Agregar y listar integrantes](specs/sprint-1/24-interfaz-de-consola.md?plain=1#L115) | `TestAgregarYListarIntegrantes` [console_test.go:160](../cmd/app/console_test.go#L160) | [console.go](../cmd/app/console.go) | `385645f` → `ae3ffa5` |
| [Listar un backlog vacio](specs/sprint-1/24-interfaz-de-consola.md?plain=1#L121) | `TestListarUnBacklogVacio` [console_test.go:185](../cmd/app/console_test.go#L185) | [console.go](../cmd/app/console.go) | `11a3b48` → `01b3b07` |
| [Cargar una historia y listar el backlog](specs/sprint-1/24-interfaz-de-consola.md?plain=1#L126) | `TestCargarUnaHistoriaYListarElBacklog` [console_test.go:204](../cmd/app/console_test.go#L204) | [console.go](../cmd/app/console.go) | `3f578a2` → `fc3fa30` |
| [Estimar una historia desde la consola](specs/sprint-1/24-interfaz-de-consola.md?plain=1#L132) | `TestEstimarUnaHistoriaDesdeLaConsola` [console_test.go:243](../cmd/app/console_test.go#L243) | [console.go](../cmd/app/console.go) | `0d8941c` → `4d2fe0e` |
| [El nucleo rechaza un valor fuera de la escala](specs/sprint-1/24-interfaz-de-consola.md?plain=1#L137) | `TestElNucleoRechazaUnValorFueraDeLaEscala` [console_test.go:268](../cmd/app/console_test.go#L268) | el que ya estaba | `383dcf7` (sin RED: ya pasaba) |
| [Operacion incompleta](specs/sprint-1/24-interfaz-de-consola.md?plain=1#L143) | `TestOperacionIncompleta` [console_test.go:292](../cmd/app/console_test.go#L292) | [console.go](../cmd/app/console.go) | `bedd456` → `28b8aef` |
| [Dato con forma invalida](specs/sprint-1/24-interfaz-de-consola.md?plain=1#L149) | `TestDatoConFormaInvalida` [console_test.go:313](../cmd/app/console_test.go#L313) | [console.go](../cmd/app/console.go) | `7e64570` → `de946a6` |
| [Operar sobre un identificador que no existe](specs/sprint-1/24-interfaz-de-consola.md?plain=1#L155) | `TestOperarSobreUnIdentificadorQueNoExiste` [console_test.go:353](../cmd/app/console_test.go#L353) | el que ya estaba | `fb5bad5` (sin RED: ya pasaba) |
| [Operacion inexistente](specs/sprint-1/24-interfaz-de-consola.md?plain=1#L161) | `TestOperacionInexistente` [console_test.go:54](../cmd/app/console_test.go#L54) | [console.go](../cmd/app/console.go) | `35173b0` → `69fbd58` (el test trae lo justo para compilar) |
| [Entrada vacia](specs/sprint-1/24-interfaz-de-consola.md?plain=1#L166) | `TestEntradaVacia` [console_test.go:38](../cmd/app/console_test.go#L38) | [console.go](../cmd/app/console.go) | `44eaaf7` → `d9e27a8` (el test trae lo justo para compilar) |
| [Salir del programa](specs/sprint-1/24-interfaz-de-consola.md?plain=1#L172) | `TestSalirDelPrograma` [console_test.go:16](../cmd/app/console_test.go#L16) | [console.go](../cmd/app/console.go) | `7d31b89` → `cd3c73b` |

**Tests que no citan un escenario de la Fase 1 (5).**

| Test | Cubre | Código | RED → GREEN |
|---|---|---|---|
| `TestUnaFechaMalEscritaSeInformaYLaOperacionNoSeEjecuta` [console_test.go:130](../cmd/app/console_test.go#L130) | RN-5 | [console.go](../cmd/app/console.go) | `e48de53` → `42cfe49` |
| `TestCadaOperacionMuestraElRechazoDelNucleo` [console_test.go:376](../cmd/app/console_test.go#L376) | RN-2, RN-7 | el que ya estaba | `43fface` (sin RED: ya pasaba) |
| `TestElBacklogMuestraLosStoryPointsDeUnaHistoriaEstimada` [console_test.go:417](../cmd/app/console_test.go#L417) | RN-1 | el que ya estaba | `43fface` (sin RED: ya pasaba) |
| `TestLaEntradaTerminaSinSalir` [console_test.go:435](../cmd/app/console_test.go#L435) | RN-6 | el que ya estaba | `43fface` (sin RED: ya pasaba) |
| `TestListarLosIntegrantesDeUnProyectoSinIntegrantes` [console_test.go:448](../cmd/app/console_test.go#L448) | RN-1 | [console.go](../cmd/app/console.go) | `f0ca0e9` → `6120035` |

## #5 Consultar el estado de un proyecto

- **Historia:** issue #5
- **Spec:** [docs/specs/sprint-2/05-consultar-estado-proyecto.md](specs/sprint-2/05-consultar-estado-proyecto.md)
- **Código según la spec:** `internal/domain/project_status.go`

<details><summary>Criterios de aceptación (12)</summary>

- El estado muestra el identificador, el nombre, la descripción y las dos fechas del proyecto.
- Un proyecto sin descripción o sin fecha de finalización las muestra como no cargadas, sin error.
- El estado muestra cada integrante con su nombre y su rol, en el orden en que se cargó.
- Un proyecto sin integrantes muestra la lista vacía, sin error.
- El estado muestra la cantidad de historias totales, completadas y pendientes, y el total es igual a completadas más pendientes.
- Solo las historias Terminadas cuentan como completadas; las En curso cuentan como pendientes.
- Si hay un sprint abierto, el estado muestra su número, su goal, sus fechas y cuántas historias tiene.
- Si no hay sprint abierto, el estado lo informa y eso no es un error.
- Una consulta hecha después de un cambio en el proyecto lo refleja.
- Consultar el estado no modifica nada del proyecto.
- El estado de un proyecto no incluye nada de otro.
- Pedir el estado de un proyecto que no existe informa que no se encontró.

</details>

| Escenario BDD | Test | Código | RED → GREEN |
|---|---|---|---|
| [Consultar el estado de un proyecto recien creado](specs/sprint-2/05-consultar-estado-proyecto.md?plain=1#L123) | `TestConsultarElEstadoDeUnProyectoRecienCreado` [project_status_test.go:47](../internal/domain/project_status_test.go#L47) | [project_status.go](../internal/domain/project_status.go) | `7de777d` → `718ae62` |
| [Un proyecto sin descripcion ni fecha de finalizacion](specs/sprint-2/05-consultar-estado-proyecto.md?plain=1#L132) | `TestUnProyectoSinDescripcionNiFechaDeFinalizacion` [project_status_test.go:83](../internal/domain/project_status_test.go#L83) | [project_status.go](../internal/domain/project_status.go) | `7de777d` → `718ae62` |
| [El estado muestra los integrantes con su rol](specs/sprint-2/05-consultar-estado-proyecto.md?plain=1#L137) | `TestElEstadoMuestraLosIntegrantesConSuRol` [project_status_test.go:118](../internal/domain/project_status_test.go#L118) | [project_status.go](../internal/domain/project_status.go) | `d7aa701` → `674ff06` |
| [Resumen del backlog segun el estado de las historias](specs/sprint-2/05-consultar-estado-proyecto.md?plain=1#L143) | `TestResumenDelBacklogSegunElEstadoDeLasHistorias` [project_status_test.go:140](../internal/domain/project_status_test.go#L140) | [project_status.go](../internal/domain/project_status.go) | `a385451` → `431b247` |
| [El estado muestra el sprint en curso](specs/sprint-2/05-consultar-estado-proyecto.md?plain=1#L155) | `TestElEstadoMuestraElSprintEnCurso` [project_status_test.go:260](../internal/domain/project_status_test.go#L260) | [project_status.go](../internal/domain/project_status.go) | `27cd558` → `96b9633` |
| [Un proyecto con todos sus sprints cerrados no tiene sprint en curso](specs/sprint-2/05-consultar-estado-proyecto.md?plain=1#L161) | `TestUnProyectoConTodosSusSprintsCerradosNoTieneSprintEnCurso` [project_status_test.go:275](../internal/domain/project_status_test.go#L275) | [project_status.go](../internal/domain/project_status.go) | `27cd558` → `96b9633` |
| [El estado refleja los cambios del proyecto](specs/sprint-2/05-consultar-estado-proyecto.md?plain=1#L167) | `TestElEstadoReflejaLosCambiosDelProyecto` [project_status_test.go:170](../internal/domain/project_status_test.go#L170) | [project_status.go](../internal/domain/project_status.go) | `a385451` → `431b247` |
| [Consultar el estado no cambia nada](specs/sprint-2/05-consultar-estado-proyecto.md?plain=1#L174) | `TestConsultarElEstadoNoCambiaNada` [project_status_test.go:191](../internal/domain/project_status_test.go#L191) | [project_status.go](../internal/domain/project_status.go) | `a385451` → `431b247` |
| [Cada proyecto muestra solo lo suyo](specs/sprint-2/05-consultar-estado-proyecto.md?plain=1#L180) | `TestCadaProyectoMuestraSoloLoSuyo` [project_status_test.go:223](../internal/domain/project_status_test.go#L223) | [project_status.go](../internal/domain/project_status.go) | `a385451` → `431b247` |
| [Consultar el estado de un proyecto que no existe](specs/sprint-2/05-consultar-estado-proyecto.md?plain=1#L186) | `TestConsultarElEstadoDeUnProyectoQueNoExiste` [project_status_test.go:103](../internal/domain/project_status_test.go#L103) | [project_status.go](../internal/domain/project_status.go) | `7de777d` → `718ae62` |

## #7 Priorizar el Product Backlog

- **Historia:** issue #7
- **Spec:** [docs/specs/sprint-2/07-priorizar-backlog.md](specs/sprint-2/07-priorizar-backlog.md)
- **Código según la spec:** `PrioritizedBacklog()`, `internal/domain/story_priority.go`

<details><summary>Criterios de aceptación (14)</summary>

- Pedir el backlog ordenado de un proyecto devuelve primero las historias de prioridad Alta, después las Media y por último las Baja.
- Dentro de una misma prioridad, las historias quedan ordenadas por identificador de menor a mayor.
- El listado ordenado devuelve la misma cantidad de historias que tiene el backlog, sin omitir ni repetir ninguna.
- Las historias aparecen en el listado cualquiera sea su estado y estén estimadas o no.
- Cada historia del listado se devuelve con su identificador, título, descripción, prioridad, estado, story points y criterios de aceptación.
- Al cambiar la prioridad de una historia, el listado siguiente la ubica en el lugar que corresponde a su nueva prioridad.
- Cambiar una prioridad por la misma que ya tenía deja el listado igual.
- Pedir el listado ordenado no cambia ningún dato de las historias ni el orden en que el backlog las devuelve sin ordenar.
- Pedir el listado ordenado dos veces seguidas, sin cambiar nada en el medio, devuelve exactamente el mismo resultado.
- El backlog ordenado de un proyecto sin historias se devuelve vacío, y eso no es un error.
- Un proyecto con una sola historia devuelve un listado con esa única historia.
- Un backlog donde todas las historias tienen la misma prioridad se devuelve ordenado por identificador.
- Pedir el backlog ordenado de un proyecto que no existe informa que no se encontró.
- Las historias Terminadas aparecen al final del listado, después de todas las demás, y ordenadas entre sí por prioridad.

</details>

| Escenario BDD | Test | Código | RED → GREEN |
|---|---|---|---|
| [Listar el backlog ordenado por prioridad](specs/sprint-2/07-priorizar-backlog.md?plain=1#L150) | `TestListarElBacklogOrdenadoPorPrioridad` [story_priority_test.go:152](../internal/domain/story_priority_test.go#L152) | [story_priority.go](../internal/domain/story_priority.go) | `eb1c628` → `f11be53` |
| [Desempatar por identificador dentro de la misma prioridad](specs/sprint-2/07-priorizar-backlog.md?plain=1#L155) | `TestDesempatarPorIdentificadorDentroDeLaMismaPrioridad` [story_priority_test.go:252](../internal/domain/story_priority_test.go#L252) | [story_priority.go](../internal/domain/story_priority.go) | `50a9d3f` → `deef1b0` |
| [El orden se actualiza al cambiar una prioridad](specs/sprint-2/07-priorizar-backlog.md?plain=1#L160) | `TestElOrdenSeActualizaAlCambiarUnaPrioridad` [story_priority_test.go:176](../internal/domain/story_priority_test.go#L176) | [story_priority.go](../internal/domain/story_priority.go) | `eb1c628` → `f11be53` |
| [Bajar la prioridad de la historia que estaba primera](specs/sprint-2/07-priorizar-backlog.md?plain=1#L172) | `TestBajarLaPrioridadDeLaHistoriaQueEstabaPrimera` [story_priority_test.go:204](../internal/domain/story_priority_test.go#L204) | [story_priority.go](../internal/domain/story_priority.go) | `eb1c628` → `f11be53` |
| [Backlog ordenado de un proyecto sin historias](specs/sprint-2/07-priorizar-backlog.md?plain=1#L178) | `TestBacklogOrdenadoDeUnProyectoSinHistorias` [story_priority_test.go:50](../internal/domain/story_priority_test.go#L50) | [story_priority.go](../internal/domain/story_priority.go) | `41f2471` → `d194fa8` |
| [Backlog ordenado de un proyecto con una sola historia](specs/sprint-2/07-priorizar-backlog.md?plain=1#L184) | `TestBacklogOrdenadoDeUnProyectoConUnaSolaHistoria` [story_priority_test.go:140](../internal/domain/story_priority_test.go#L140) | [story_priority.go](../internal/domain/story_priority.go) | `ab5ea38` → `5c75d88` |
| [Todas las historias tienen la misma prioridad](specs/sprint-2/07-priorizar-backlog.md?plain=1#L189) | `TestTodasLasHistoriasTienenLaMismaPrioridad` [story_priority_test.go:264](../internal/domain/story_priority_test.go#L264) | [story_priority.go](../internal/domain/story_priority.go) | `50a9d3f` → `deef1b0` |
| [No hay historias de prioridad alta](specs/sprint-2/07-priorizar-backlog.md?plain=1#L194) | `TestNoHayHistoriasDePrioridadAlta` [story_priority_test.go:164](../internal/domain/story_priority_test.go#L164) | [story_priority.go](../internal/domain/story_priority.go) | `eb1c628` → `f11be53` |
| [El listado incluye las historias en cualquier estado](specs/sprint-2/07-priorizar-backlog.md?plain=1#L200) | `TestElListadoIncluyeLasHistoriasEnCualquierEstado` [story_priority_test.go:83](../internal/domain/story_priority_test.go#L83) | [story_priority.go](../internal/domain/story_priority.go) | `ab5ea38` → `5c75d88` |
| [Las historias del listado traen todos sus datos](specs/sprint-2/07-priorizar-backlog.md?plain=1#L205) | `TestLasHistoriasDelListadoTraenTodosSusDatos` [story_priority_test.go:109](../internal/domain/story_priority_test.go#L109) | [story_priority.go](../internal/domain/story_priority.go) | `ab5ea38` → `5c75d88` |
| [Pedir el listado ordenado no cambia el backlog](specs/sprint-2/07-priorizar-backlog.md?plain=1#L211) | `TestPedirElListadoOrdenadoNoCambiaElBacklog` [story_priority_test.go:234](../internal/domain/story_priority_test.go#L234) | [story_priority.go](../internal/domain/story_priority.go) | `eb1c628` → `f11be53` |
| [Dos listados seguidos devuelven lo mismo](specs/sprint-2/07-priorizar-backlog.md?plain=1#L218) | `TestDosListadosSeguidosDevuelvenLoMismo` [story_priority_test.go:276](../internal/domain/story_priority_test.go#L276) | [story_priority.go](../internal/domain/story_priority.go) | `50a9d3f` → `deef1b0` |
| [Cambiar una prioridad por la que ya tenia](specs/sprint-2/07-priorizar-backlog.md?plain=1#L223) | `TestCambiarUnaPrioridadPorLaQueYaTenia` [story_priority_test.go:219](../internal/domain/story_priority_test.go#L219) | [story_priority.go](../internal/domain/story_priority.go) | `eb1c628` → `f11be53` |
| [Pedir el backlog ordenado de un proyecto que no existe](specs/sprint-2/07-priorizar-backlog.md?plain=1#L229) | `TestPedirElBacklogOrdenadoDeUnProyectoQueNoExiste` [story_priority_test.go:68](../internal/domain/story_priority_test.go#L68) | [story_priority.go](../internal/domain/story_priority.go) | `41f2471` → `d194fa8` |
| [Las historias terminadas van al final](specs/sprint-2/07-priorizar-backlog.md?plain=1#L235) | `TestLasHistoriasTerminadasVanAlFinal` [story_priority_test.go:309](../internal/domain/story_priority_test.go#L309) | [story_priority.go](../internal/domain/story_priority.go) | `cc4e1f9` → `ba0c908` |

**Tests que no citan un escenario de la Fase 1 (1).**

| Test | Cubre | Código | RED → GREEN |
|---|---|---|---|
| `TestDesempatarPorIdentificadorAunqueElBacklogEsteGuardadoEnOtroOrden` [story_priority_test.go:295](../internal/domain/story_priority_test.go#L295) | RN-3 y la decision de la Fase 2 "El desempate por identificador se hace explicito" | [story_priority.go](../internal/domain/story_priority.go) | `50a9d3f` → `deef1b0` |

## #8 Crear un sprint con su Sprint Goal

- **Historia:** issue #8
- **Spec:** [docs/specs/sprint-2/08-crear-sprint.md](specs/sprint-2/08-crear-sprint.md)
- **Código según la spec:** `internal/domain/sprint.go`, `internal/domain/project.go`, `internal/domain/registry_json.go`, `internal/domain/errors.go`

<details><summary>Criterios de aceptación (15)</summary>

- Se puede crear un sprint con Sprint Goal, fecha de inicio y fecha de fin, y queda registrado con su número y en estado Abierto.
- El primer sprint de un proyecto recibe el número 1, y cada sprint nuevo el número siguiente.
- El número no se repite dentro del mismo proyecto.
- Un sprint recién creado no tiene historias asignadas.
- No se puede crear un sprint sin Sprint Goal, y un Sprint Goal de solo espacios cuenta como vacío.
- No se puede crear un sprint sin fecha de inicio ni sin fecha de fin.
- No se puede crear un sprint con fecha de fin anterior a la de inicio.
- Un sprint cuya fecha de fin es igual a la de inicio se acepta.
- No se puede crear un segundo sprint mientras haya uno abierto en el mismo proyecto, y el mensaje dice cuál es el que está abierto.
- Tener un sprint abierto en un proyecto no impide crear un sprint en otro proyecto.
- Se puede consultar cuál es el sprint abierto de un proyecto; si no hay ninguno se informa y no es un error.
- El Sprint Goal se guarda sin los espacios del principio y del final.
- Una creación rechazada no agrega ningún sprint al proyecto.
- Crear un sprint o pedir el sprint abierto de un proyecto inexistente informa que no se encontró el proyecto.
- Se puede consultar un sprint por su número, esté abierto o cerrado, y un número que no existe informa que no se encontró el sprint.

</details>

| Escenario BDD | Test | Código | RED → GREEN |
|---|---|---|---|
| [Crear el primer sprint de un proyecto](specs/sprint-2/08-crear-sprint.md?plain=1#L159) | `TestCrearElPrimerSprintDeUnProyecto` [sprint_test.go:29](../internal/domain/sprint_test.go#L29) | [sprint.go](../internal/domain/sprint.go) | `5d602fa` → `e3a9416` |
| [El sprint nuevo recibe el numero siguiente](specs/sprint-2/08-crear-sprint.md?plain=1#L166) | `TestElSprintNuevoRecibeElNumeroSiguiente` [sprint_test.go:281](../internal/domain/sprint_test.go#L281) | el que ya estaba | `19394e3` (sin RED: ya pasaba) |
| [No se puede abrir un segundo sprint habiendo uno abierto](specs/sprint-2/08-crear-sprint.md?plain=1#L172) | `TestNoSePuedeAbrirUnSegundoSprintHabiendoUnoAbierto` [sprint_test.go:254](../internal/domain/sprint_test.go#L254) | [errors.go](../internal/domain/errors.go), [sprint.go](../internal/domain/sprint.go) | `ffaada8` → `c59f3b2` |
| [Rechazar un sprint sin Sprint Goal](specs/sprint-2/08-crear-sprint.md?plain=1#L179) | `TestRechazarUnSprintSinSprintGoal` [sprint_test.go:138](../internal/domain/sprint_test.go#L138) | [errors.go](../internal/domain/errors.go), [sprint.go](../internal/domain/sprint.go) | `07deed8` → `a2cadc9` |
| [Rechazar un sprint sin las fechas obligatorias](specs/sprint-2/08-crear-sprint.md?plain=1#L191) | `TestRechazarUnSprintSinLasFechasObligatorias` [sprint_test.go:164](../internal/domain/sprint_test.go#L164) | [errors.go](../internal/domain/errors.go), [sprint.go](../internal/domain/sprint.go) | `6991115` → `20eb08c` |
| [Rechazar un sprint con fechas incoherentes](specs/sprint-2/08-crear-sprint.md?plain=1#L203) | `TestRechazarUnSprintConFechasIncoherentes` [sprint_test.go:191](../internal/domain/sprint_test.go#L191) | [errors.go](../internal/domain/errors.go), [sprint.go](../internal/domain/sprint.go) | `8511e9d` → `734a903` |
| [Un sprint de un solo dia es valido](specs/sprint-2/08-crear-sprint.md?plain=1#L209) | `TestUnSprintDeUnSoloDiaEsValido` [sprint_test.go:206](../internal/domain/sprint_test.go#L206) | el que ya estaba | `fcbd577` (sin RED: ya pasaba) |
| [Un sprint con fechas pasadas es valido](specs/sprint-2/08-crear-sprint.md?plain=1#L214) | `TestUnSprintConFechasPasadasEsValido` [sprint_test.go:221](../internal/domain/sprint_test.go#L221) | el que ya estaba | `fcbd577` (sin RED: ya pasaba) |
| [Guardar el Sprint Goal sin espacios en los extremos](specs/sprint-2/08-crear-sprint.md?plain=1#L219) | `TestGuardarElSprintGoalSinEspaciosEnLosExtremos` [sprint_test.go:236](../internal/domain/sprint_test.go#L236) | [sprint.go](../internal/domain/sprint.go) | `41396e9` → `eb536ef` |
| [Cada proyecto puede tener su propio sprint abierto](specs/sprint-2/08-crear-sprint.md?plain=1#L224) | `TestCadaProyectoPuedeTenerSuPropioSprintAbierto` [sprint_test.go:345](../internal/domain/sprint_test.go#L345) | el que ya estaba | `cfde49b` (sin RED: ya pasaba) |
| [Consultar el sprint abierto de un proyecto](specs/sprint-2/08-crear-sprint.md?plain=1#L230) | `TestConsultarElSprintAbiertoDeUnProyecto` [sprint_test.go:57](../internal/domain/sprint_test.go#L57) | [sprint.go](../internal/domain/sprint.go) | `6cf7475` → `de2ac50` |
| [Un proyecto sin sprints no tiene sprint abierto](specs/sprint-2/08-crear-sprint.md?plain=1#L235) | `TestUnProyectoSinSprintsNoTieneSprintAbierto` [sprint_test.go:11](../internal/domain/sprint_test.go#L11) | [project.go](../internal/domain/project.go), [sprint.go](../internal/domain/sprint.go) | `5e2b377` → `ff3264e` |
| [Un proyecto con todos sus sprints cerrados no tiene sprint abierto](specs/sprint-2/08-crear-sprint.md?plain=1#L241) | `TestUnProyectoConTodosSusSprintsCerradosNoTieneSprintAbierto` [sprint_test.go:304](../internal/domain/sprint_test.go#L304) | el que ya estaba | `7309747` (sin RED: ya pasaba) |
| [Dos sprints del mismo proyecto pueden tener el mismo goal y las mismas fechas](specs/sprint-2/08-crear-sprint.md?plain=1#L246) | `TestDosSprintsDelMismoProyectoPuedenTenerElMismoGoalYLasMismasFechas` [sprint_test.go:321](../internal/domain/sprint_test.go#L321) | el que ya estaba | `cfde49b` (sin RED: ya pasaba) |
| [Operar sobre un proyecto que no existe](specs/sprint-2/08-crear-sprint.md?plain=1#L252) | `TestOperarSprintsSobreUnProyectoQueNoExiste` [sprint_test.go:380](../internal/domain/sprint_test.go#L380) | el que ya estaba | `cfde49b` (sin RED: ya pasaba) |
| [Consultar un sprint por su numero](specs/sprint-2/08-crear-sprint.md?plain=1#L262) | `TestConsultarUnSprintPorSuNumero` [sprint_test.go:89](../internal/domain/sprint_test.go#L89) | [errors.go](../internal/domain/errors.go), [sprint.go](../internal/domain/sprint.go) | `5de4cda` → `ee69eed` |
| [Consultar un sprint que no existe](specs/sprint-2/08-crear-sprint.md?plain=1#L267) | `TestConsultarUnSprintQueNoExiste` [sprint_test.go:120](../internal/domain/sprint_test.go#L120) | [errors.go](../internal/domain/errors.go), [sprint.go](../internal/domain/sprint.go) | `5de4cda` → `ee69eed` |

## #9 Asignar historias a un sprint

- **Historia:** issue #9
- **Spec:** [docs/specs/sprint-2/09-asignar-historias-a-sprint.md](specs/sprint-2/09-asignar-historias-a-sprint.md)
- **Código según la spec:** `AssignStory`, `UnassignStory`, `internal/domain/sprint_backlog.go`, `internal/domain/registry_json.go`, `internal/domain/errors.go`

<details><summary>Criterios de aceptación (12)</summary>

- Se puede asignar una historia estimada al sprint abierto, y queda en el Sprint Backlog.
- Al asignarla, la historia queda En curso.
- No se puede asignar una historia sin estimar.
- No se puede asignar dos veces la misma historia al mismo sprint.
- No se puede asignar ni quitar una historia Terminada.
- No se puede asignar ni quitar nada si el proyecto no tiene un sprint abierto.
- Se puede quitar una historia del sprint mientras el sprint esté abierto, y vuelve al Product Backlog en estado Pendiente.
- No se puede quitar una historia que no está en el sprint.
- Se puede consultar el Sprint Backlog del sprint abierto, en el orden en que se asignaron las historias, y si está vacío eso no es un error.
- Una historia que quedó sin completar en un sprint cerrado se puede asignar al sprint abierto siguiente.
- Una operación rechazada no cambia el Sprint Backlog ni el estado de la historia.
- Asignar o quitar en un proyecto o sobre una historia que no existe informa que no se encontró.

</details>

| Escenario BDD | Test | Código | RED → GREEN |
|---|---|---|---|
| [Asignar una historia estimada al sprint abierto](specs/sprint-2/09-asignar-historias-a-sprint.md?plain=1#L148) | `TestAsignarUnaHistoriaEstimadaAlSprintAbierto` [sprint_backlog_test.go:54](../internal/domain/sprint_backlog_test.go#L54) | [sprint_backlog.go](../internal/domain/sprint_backlog.go) | `4eda9be` → `7405b81` |
| [El Sprint Backlog respeta el orden en que se asignaron las historias](specs/sprint-2/09-asignar-historias-a-sprint.md?plain=1#L154) | `TestElSprintBacklogRespetaElOrdenEnQueSeAsignaronLasHistorias` [sprint_backlog_test.go:73](../internal/domain/sprint_backlog_test.go#L73) | el que ya estaba | `3356937` (sin RED: ya pasaba) |
| [Un sprint recien abierto tiene el Sprint Backlog vacio](specs/sprint-2/09-asignar-historias-a-sprint.md?plain=1#L159) | `TestUnSprintRecienAbiertoTieneElSprintBacklogVacio` [sprint_backlog_test.go:22](../internal/domain/sprint_backlog_test.go#L22) | el que ya estaba | `7f9af97` (sin RED: ya pasaba) |
| [Rechazar una historia sin estimar](specs/sprint-2/09-asignar-historias-a-sprint.md?plain=1#L165) | `TestRechazarUnaHistoriaSinEstimar` [sprint_backlog_test.go:94](../internal/domain/sprint_backlog_test.go#L94) | [errors.go](../internal/domain/errors.go), [sprint_backlog.go](../internal/domain/sprint_backlog.go) | `9890acc` → `64ba2dc` |
| [Rechazar una historia que ya esta en el sprint](specs/sprint-2/09-asignar-historias-a-sprint.md?plain=1#L173) | `TestRechazarUnaHistoriaQueYaEstaEnElSprint` [sprint_backlog_test.go:116](../internal/domain/sprint_backlog_test.go#L116) | [errors.go](../internal/domain/errors.go), [sprint_backlog.go](../internal/domain/sprint_backlog.go) | `7b688bd` → `7c4dd2a` |
| [Rechazar una historia terminada](specs/sprint-2/09-asignar-historias-a-sprint.md?plain=1#L180) | `TestRechazarUnaHistoriaTerminada` [sprint_backlog_test.go:141](../internal/domain/sprint_backlog_test.go#L141) | [sprint_backlog.go](../internal/domain/sprint_backlog.go) | `cd46b06` → `936fb60` |
| [Rechazar quitar una historia terminada](specs/sprint-2/09-asignar-historias-a-sprint.md?plain=1#L187) | `TestRechazarQuitarUnaHistoriaTerminada` [sprint_backlog_test.go:167](../internal/domain/sprint_backlog_test.go#L167) | [sprint_backlog.go](../internal/domain/sprint_backlog.go) | `cd46b06` → `936fb60` |
| [Operar sin un sprint abierto](specs/sprint-2/09-asignar-historias-a-sprint.md?plain=1#L195) | `TestOperarSinUnSprintAbierto` [sprint_backlog_test.go:192](../internal/domain/sprint_backlog_test.go#L192) | [errors.go](../internal/domain/errors.go), [sprint_backlog.go](../internal/domain/sprint_backlog.go) | `3524b7f` → `670390e` |
| [Quitar una historia del sprint abierto](specs/sprint-2/09-asignar-historias-a-sprint.md?plain=1#L206) | `TestQuitarUnaHistoriaDelSprintAbierto` [sprint_backlog_test.go:216](../internal/domain/sprint_backlog_test.go#L216) | [sprint_backlog.go](../internal/domain/sprint_backlog.go) | `0d27e9e` → `edfbc46` |
| [Quitar una historia y volver a asignarla](specs/sprint-2/09-asignar-historias-a-sprint.md?plain=1#L212) | `TestQuitarUnaHistoriaYVolverAAsignarla` [sprint_backlog_test.go:238](../internal/domain/sprint_backlog_test.go#L238) | el que ya estaba | `f8681fa` (sin RED: ya pasaba) |
| [Rechazar quitar una historia que no esta en el sprint](specs/sprint-2/09-asignar-historias-a-sprint.md?plain=1#L217) | `TestRechazarQuitarUnaHistoriaQueNoEstaEnElSprint` [sprint_backlog_test.go:263](../internal/domain/sprint_backlog_test.go#L263) | [errors.go](../internal/domain/errors.go), [sprint_backlog.go](../internal/domain/sprint_backlog.go) | `c739a3a` → `7443a68` |
| [Una historia que quedo sin completar se puede asignar al sprint siguiente](specs/sprint-2/09-asignar-historias-a-sprint.md?plain=1#L224) | `TestUnaHistoriaQueQuedoSinCompletarSePuedeAsignarAlSprintSiguiente` [sprint_backlog_test.go:282](../internal/domain/sprint_backlog_test.go#L282) | el que ya estaba | `f22036f` (sin RED: ya pasaba) |
| [Cada proyecto arma su propio Sprint Backlog](specs/sprint-2/09-asignar-historias-a-sprint.md?plain=1#L232) | `TestCadaProyectoArmaSuPropioSprintBacklog` [sprint_backlog_test.go:315](../internal/domain/sprint_backlog_test.go#L315) | el que ya estaba | `f22036f` (sin RED: ya pasaba) |
| [Operar sobre algo que no existe](specs/sprint-2/09-asignar-historias-a-sprint.md?plain=1#L237) | `TestOperarSobreAlgoQueNoExiste` [sprint_backlog_test.go:347](../internal/domain/sprint_backlog_test.go#L347) | [defect.go](../internal/domain/defect.go), [errors.go](../internal/domain/errors.go) | `b45a9fc` → `69ddd84` |

## #14 Registrar el esfuerzo real

- **Historia:** issue #14
- **Spec:** [docs/specs/sprint-2/14-registrar-esfuerzo-real.md](specs/sprint-2/14-registrar-esfuerzo-real.md)
- **Código según la spec:** `internal/domain/effort.go`, `internal/domain/project.go`, `internal/domain/registry_json.go`, `internal/domain/errors.go`

<details><summary>Criterios de aceptación (15)</summary>

- Se puede registrar esfuerzo indicando integrante, historia, día, horas y actividad realizada, y el registro queda guardado con esos cinco datos.
- Se aceptan fracciones de hora, por ejemplo media hora.
- No se pueden registrar cero horas, un valor negativo, un valor infinito ni algo que no es un número.
- No se pueden registrar horas con fecha posterior al día de hoy, y el día de hoy sí se acepta.
- No se puede registrar esfuerzo sin día ni sin actividad realizada; una actividad de solo espacios cuenta como vacía y lo que se guarda va sin los espacios de los extremos.
- No se puede registrar esfuerzo de alguien que no es integrante del proyecto; el nombre se encuentra aunque venga con otras mayúsculas o espacios de más, y se guarda tal como está cargado en el proyecto.
- No se puede registrar sobre una historia que no pertenece al proyecto, y sí sobre una historia Terminada.
- El mismo integrante puede cargar dos veces la misma historia el mismo día, y las dos cargas quedan y suman.
- Se puede consultar el total de horas de una historia, y una historia sin registros da cero sin que sea un error.
- Se puede consultar el total de horas de un integrante, y un integrante sin registros da cero sin que sea un error.
- Consultar las horas de una historia que no es del proyecto, o de alguien que no es integrante, informa que no se encontró.
- Se pueden consultar todos los registros del proyecto, en el orden en que se cargaron.
- Las horas cargadas en un proyecto no suman en otro.
- Un registro rechazado no deja nada guardado.
- Registrar o consultar en un proyecto inexistente informa que no se encontró el proyecto.

</details>

| Escenario BDD | Test | Código | RED → GREEN |
|---|---|---|---|
| [Registrar horas trabajadas en una historia](specs/sprint-2/14-registrar-esfuerzo-real.md?plain=1#L179) | `TestRegistrarHorasTrabajadasEnUnaHistoria` [effort_test.go:42](../internal/domain/effort_test.go#L42) | [effort.go](../internal/domain/effort.go), [project.go](../internal/domain/project.go) | `d1b98f4` → `a82f98d` |
| [Registrar media hora](specs/sprint-2/14-registrar-esfuerzo-real.md?plain=1#L185) | `TestRegistrarMediaHora` [effort_test.go:66](../internal/domain/effort_test.go#L66) | el que ya estaba | `254b808` (sin RED: ya pasaba) |
| [Rechazar horas que no son mayores que cero](specs/sprint-2/14-registrar-esfuerzo-real.md?plain=1#L190) | `TestRechazarHorasQueNoSonMayoresQueCero` [effort_test.go:83](../internal/domain/effort_test.go#L83) | [effort.go](../internal/domain/effort.go), [errors.go](../internal/domain/errors.go) | `d7cebf8` → `10e8616` |
| [Rechazar horas de un dia que todavia no paso](specs/sprint-2/14-registrar-esfuerzo-real.md?plain=1#L204) | `TestRechazarHorasDeUnDiaQueTodaviaNoPaso` [effort_test.go:113](../internal/domain/effort_test.go#L113) | [effort.go](../internal/domain/effort.go), [errors.go](../internal/domain/errors.go) | `930478a` → `8796e5d` |
| [Registrar horas de un dia anterior al inicio del proyecto](specs/sprint-2/14-registrar-esfuerzo-real.md?plain=1#L211) | `TestRegistrarHorasDeUnDiaAnteriorAlInicioDelProyecto` [effort_test.go:130](../internal/domain/effort_test.go#L130) | [effort.go](../internal/domain/effort.go), [errors.go](../internal/domain/errors.go) | `930478a` → `8796e5d` |
| [Rechazar un registro sin fecha](specs/sprint-2/14-registrar-esfuerzo-real.md?plain=1#L216) | `TestRechazarUnRegistroSinFecha` [effort_test.go:147](../internal/domain/effort_test.go#L147) | [effort.go](../internal/domain/effort.go), [errors.go](../internal/domain/errors.go) | `a2a393b` → `d18d77a` |
| [Rechazar un registro sin actividad realizada](specs/sprint-2/14-registrar-esfuerzo-real.md?plain=1#L222) | `TestRechazarUnRegistroSinActividadRealizada` [effort_test.go:164](../internal/domain/effort_test.go#L164) | [effort.go](../internal/domain/effort.go), [errors.go](../internal/domain/errors.go) | `0ca1644` → `8b7645a` |
| [Guardar la actividad sin los espacios de los extremos](specs/sprint-2/14-registrar-esfuerzo-real.md?plain=1#L234) | `TestGuardarLaActividadSinLosEspaciosDeLosExtremos` [effort_test.go:185](../internal/domain/effort_test.go#L185) | [effort.go](../internal/domain/effort.go), [errors.go](../internal/domain/errors.go) | `0ca1644` → `8b7645a` |
| [Rechazar el registro de alguien que no es integrante del proyecto](specs/sprint-2/14-registrar-esfuerzo-real.md?plain=1#L239) | `TestRechazarElRegistroDeAlguienQueNoEsIntegranteDelProyecto` [effort_test.go:202](../internal/domain/effort_test.go#L202) | [effort.go](../internal/domain/effort.go) | `cece8ee` → `7ce0599` |
| [Encontrar al integrante aunque el nombre venga escrito distinto](specs/sprint-2/14-registrar-esfuerzo-real.md?plain=1#L246) | `TestEncontrarAlIntegranteAunqueElNombreVengaEscritoDistinto` [effort_test.go:219](../internal/domain/effort_test.go#L219) | [effort.go](../internal/domain/effort.go) | `cece8ee` → `7ce0599` |
| [Rechazar el registro sobre una historia que no es del proyecto](specs/sprint-2/14-registrar-esfuerzo-real.md?plain=1#L251) | `TestRechazarElRegistroSobreUnaHistoriaQueNoEsDelProyecto` [effort_test.go:236](../internal/domain/effort_test.go#L236) | [effort.go](../internal/domain/effort.go) | `2fb75e7` → `c696b48` |
| [Registrar horas sobre una historia terminada](specs/sprint-2/14-registrar-esfuerzo-real.md?plain=1#L257) | `TestRegistrarHorasSobreUnaHistoriaTerminada` [effort_test.go:253](../internal/domain/effort_test.go#L253) | [effort.go](../internal/domain/effort.go) | `2fb75e7` → `c696b48` |
| [Dos cargas del mismo integrante, la misma historia y el mismo dia](specs/sprint-2/14-registrar-esfuerzo-real.md?plain=1#L262) | `TestDosCargasDelMismoIntegranteLaMismaHistoriaYElMismoDia` [effort_test.go:275](../internal/domain/effort_test.go#L275) | [effort.go](../internal/domain/effort.go) | `c645a93` → `4ea1780` |
| [Consultar el total de horas de una historia](specs/sprint-2/14-registrar-esfuerzo-real.md?plain=1#L270) | `TestConsultarElTotalDeHorasDeUnaHistoria` [effort_test.go:326](../internal/domain/effort_test.go#L326) | [effort.go](../internal/domain/effort.go) | `2de2ee2` → `6d5d3cb` |
| [Consultar el total de horas de un integrante](specs/sprint-2/14-registrar-esfuerzo-real.md?plain=1#L275) | `TestConsultarElTotalDeHorasDeUnIntegrante` [effort_test.go:393](../internal/domain/effort_test.go#L393) | [effort.go](../internal/domain/effort.go) | `4504c33` → `79e2d56` |
| [Sin registros, el total es cero](specs/sprint-2/14-registrar-esfuerzo-real.md?plain=1#L280) | `TestSinRegistrosElTotalEsCero` [effort_test.go:340](../internal/domain/effort_test.go#L340) | [effort.go](../internal/domain/effort.go) | `2de2ee2` → `6d5d3cb` |
| [Consultar las horas de algo que no existe](specs/sprint-2/14-registrar-esfuerzo-real.md?plain=1#L291) | `TestConsultarLasHorasDeAlgoQueNoExiste` [effort_test.go:366](../internal/domain/effort_test.go#L366) | [effort.go](../internal/domain/effort.go) | `2de2ee2` → `6d5d3cb` |
| [Las horas cargadas en un proyecto no suman en otro](specs/sprint-2/14-registrar-esfuerzo-real.md?plain=1#L301) | `TestLasHorasCargadasEnUnProyectoNoSumanEnOtro` [effort_test.go:410](../internal/domain/effort_test.go#L410) | el que ya estaba | `92c6091` (sin RED: ya pasaba) |
| [Registrar esfuerzo en un proyecto que no existe](specs/sprint-2/14-registrar-esfuerzo-real.md?plain=1#L306) | `TestRegistrarEsfuerzoEnUnProyectoQueNoExiste` [effort_test.go:428](../internal/domain/effort_test.go#L428) | el que ya estaba | `a74dec4` (sin RED: ya pasaba) |
| [Un registro no puede tener mas de 24 horas](specs/sprint-2/14-registrar-esfuerzo-real.md?plain=1#L311) | `TestUnRegistroNoPuedeTenerMasDe24Horas` [effort_test.go:443](../internal/domain/effort_test.go#L443) | [effort.go](../internal/domain/effort.go), [errors.go](../internal/domain/errors.go), [efforts.go](../internal/web/efforts.go), [forms.go](../internal/web/forms.go) | `d17d90b` → `1d6e9dc` |
| [Dos registros del mismo dia que juntos pasan de 24 horas se aceptan](specs/sprint-2/14-registrar-esfuerzo-real.md?plain=1#L321) | `TestUnRegistroNoPuedeTenerMasDe24Horas` [effort_test.go:443](../internal/domain/effort_test.go#L443) | [effort.go](../internal/domain/effort.go), [errors.go](../internal/domain/errors.go), [efforts.go](../internal/web/efforts.go), [forms.go](../internal/web/forms.go) | `d17d90b` → `1d6e9dc` |
| [Un registro de mas de 24 horas que ya estaba guardado se sigue cargando](specs/sprint-2/14-registrar-esfuerzo-real.md?plain=1#L327) | `TestUnRegistroDeMasDe24HorasYaGuardadoSeCarga` [effort_test.go:474](../internal/domain/effort_test.go#L474) | [effort.go](../internal/domain/effort.go), [errors.go](../internal/domain/errors.go), [efforts.go](../internal/web/efforts.go), [forms.go](../internal/web/forms.go) | `d17d90b` → `1d6e9dc` |

**Tests que no citan un escenario de la Fase 1 (2).**

| Test | Cubre | Código | RED → GREEN |
|---|---|---|---|
| `TestUnEsfuerzoDeUnDiaPosteriorAHoySeCargaIgual` [registry_json_test.go:577](../internal/domain/registry_json_test.go#L577) | Fase 2 de la #14, la decisión sobre los datos guardados | [effort.go](../internal/domain/effort.go), [registry_json.go](../internal/domain/registry_json.go) | `c819ca6` → `4836ba4` |
| `TestElEsfuerzoDeUnIntegranteQuitadoSeCargaIgual` [registry_json_test.go:603](../internal/domain/registry_json_test.go#L603) | Fase 2 de la #14, la decisión sobre los datos guardados | [effort.go](../internal/domain/effort.go), [registry_json.go](../internal/domain/registry_json.go) | `dffc162` → `abe3750` |

## #15 Gestionar defectos

- **Historia:** issue #15
- **Spec:** [docs/specs/sprint-2/15-gestionar-defectos.md](specs/sprint-2/15-gestionar-defectos.md)
- **Código según la spec:** `internal/domain/defect.go`, `internal/domain/project.go`, `internal/domain/registry_json.go`, `internal/domain/errors.go`

<details><summary>Criterios de aceptación (13)</summary>

- Se registra un defecto con título, descripción, severidad, historia relacionada y sprint de detección, y queda con su identificador, en Abierto y sin sprint de resolución. El primero del proyecto recibe el 1 y cada uno nuevo el siguiente.
- No se puede registrar un defecto sin título o sin descripción; un texto de solo espacios cuenta como vacío y lo que se guarda va sin los espacios de los extremos.
- La severidad solo puede ser Crítica, Mayor o Menor, sin distinguir mayúsculas ni espacios de los extremos, y "Critica" sin tilde se acepta como "Crítica".
- No se puede registrar un defecto sobre una historia que no es del proyecto, y sí sobre una historia Terminada.
- No se puede registrar un defecto con un sprint de detección que no existe, y sí con uno ya cerrado.
- Un defecto no resuelto se puede pasar de Abierto a En curso y de En curso a Abierto.
- No se puede pasar un defecto a Resuelto cambiándole el estado: hay que indicar el sprint de resolución.
- Al resolver un defecto queda en Resuelto con su sprint de resolución, viniera de Abierto o de En curso.
- No se puede resolver un defecto en un sprint anterior al de detección ni con un sprint que no existe en el proyecto, y sí en el mismo sprint en que se detectó.
- Un defecto Resuelto no cambia de estado ni se vuelve a resolver.
- Se pueden listar todos los defectos, los de un estado y los de una severidad, en el orden en que se registraron, y una consulta sin resultados devuelve una lista vacía sin ser un error.
- Los defectos de un proyecto no aparecen en las consultas de otro, y cada proyecto numera desde 1.
- Una operación rechazada deja el proyecto y sus defectos como estaban, y operar sobre un defecto o un proyecto que no existe informa que no se encontró.

</details>

| Escenario BDD | Test | Código | RED → GREEN |
|---|---|---|---|
| [Registrar un defecto](specs/sprint-2/15-gestionar-defectos.md?plain=1#L185) | `TestRegistrarUnDefecto` [defect_test.go:50](../internal/domain/defect_test.go#L50) | [defect.go](../internal/domain/defect.go), [project.go](../internal/domain/project.go) | `77f495b` → `f54c976` |
| [El defecto nuevo recibe el identificador siguiente](specs/sprint-2/15-gestionar-defectos.md?plain=1#L192) | `TestElDefectoNuevoRecibeElIdentificadorSiguiente` [defect_test.go:73](../internal/domain/defect_test.go#L73) | el que ya estaba | `2b54c74` (sin RED: ya pasaba) |
| [Rechazar un defecto sin titulo o sin descripcion](specs/sprint-2/15-gestionar-defectos.md?plain=1#L198) | `TestRechazarUnDefectoSinTituloOSinDescripcion` [defect_test.go:91](../internal/domain/defect_test.go#L91) | [defect.go](../internal/domain/defect.go), [errors.go](../internal/domain/errors.go) | `efd321e` → `9a533f3` |
| [Guardar el titulo y la descripcion sin los espacios de los extremos](specs/sprint-2/15-gestionar-defectos.md?plain=1#L210) | `TestGuardarElTituloYLaDescripcionSinLosEspaciosDeLosExtremos` [defect_test.go:120](../internal/domain/defect_test.go#L120) | [defect.go](../internal/domain/defect.go), [errors.go](../internal/domain/errors.go) | `efd321e` → `9a533f3` |
| [Registrar un defecto con cada severidad](specs/sprint-2/15-gestionar-defectos.md?plain=1#L215) | `TestRegistrarUnDefectoConCadaSeveridad` [defect_test.go:138](../internal/domain/defect_test.go#L138) | [defect.go](../internal/domain/defect.go), [errors.go](../internal/domain/errors.go) | `91fb4c9` → `9e643a5` |
| [Rechazar una severidad que no existe](specs/sprint-2/15-gestionar-defectos.md?plain=1#L227) | `TestRechazarUnaSeveridadQueNoExiste` [defect_test.go:175](../internal/domain/defect_test.go#L175) | [defect.go](../internal/domain/defect.go), [errors.go](../internal/domain/errors.go) | `91fb4c9` → `9e643a5` |
| [Rechazar un defecto sobre una historia que no es del proyecto](specs/sprint-2/15-gestionar-defectos.md?plain=1#L234) | `TestRechazarUnDefectoSobreUnaHistoriaQueNoEsDelProyecto` [defect_test.go:207](../internal/domain/defect_test.go#L207) | [defect.go](../internal/domain/defect.go) | `231cc56` → `186dd0a` |
| [Registrar un defecto sobre una historia terminada](specs/sprint-2/15-gestionar-defectos.md?plain=1#L240) | `TestRegistrarUnDefectoSobreUnaHistoriaTerminada` [defect_test.go:224](../internal/domain/defect_test.go#L224) | [defect.go](../internal/domain/defect.go) | `231cc56` → `186dd0a` |
| [Rechazar un defecto con un sprint de deteccion que no existe](specs/sprint-2/15-gestionar-defectos.md?plain=1#L245) | `TestRechazarUnDefectoConUnSprintDeDeteccionQueNoExiste` [defect_test.go:238](../internal/domain/defect_test.go#L238) | [defect.go](../internal/domain/defect.go) | `7f4e8b6` → `2d6c903` |
| [Registrar un defecto detectado en un sprint ya cerrado](specs/sprint-2/15-gestionar-defectos.md?plain=1#L252) | `TestRegistrarUnDefectoDetectadoEnUnSprintYaCerrado` [defect_test.go:255](../internal/domain/defect_test.go#L255) | [defect.go](../internal/domain/defect.go) | `7f4e8b6` → `2d6c903` |
| [Cambiar el estado de un defecto no resuelto](specs/sprint-2/15-gestionar-defectos.md?plain=1#L257) | `TestCambiarElEstadoDeUnDefectoNoResuelto` [defect_test.go:281](../internal/domain/defect_test.go#L281) | [defect.go](../internal/domain/defect.go) | `a58c971` → `1c03838` |
| [Rechazar el pase a Resuelto por cambio de estado](specs/sprint-2/15-gestionar-defectos.md?plain=1#L268) | `TestRechazarElPaseAResueltoPorCambioDeEstado` [defect_test.go:309](../internal/domain/defect_test.go#L309) | [defect.go](../internal/domain/defect.go), [errors.go](../internal/domain/errors.go) | `90b1696` → `e06e055` |
| [Resolver un defecto](specs/sprint-2/15-gestionar-defectos.md?plain=1#L275) | `TestResolverUnDefecto` [defect_test.go:343](../internal/domain/defect_test.go#L343) | [defect.go](../internal/domain/defect.go) | `5aad442` → `39b8879` |
| [Resolver un defecto en el mismo sprint en que se detecto](specs/sprint-2/15-gestionar-defectos.md?plain=1#L286) | `TestResolverUnDefectoEnElMismoSprintEnQueSeDetecto` [defect_test.go:397](../internal/domain/defect_test.go#L397) | [defect.go](../internal/domain/defect.go), [errors.go](../internal/domain/errors.go) | `801f055` → `940825d` |
| [Rechazar la resolucion en un sprint anterior al de deteccion](specs/sprint-2/15-gestionar-defectos.md?plain=1#L292) | `TestRechazarLaResolucionEnUnSprintAnteriorAlDeDeteccion` [defect_test.go:376](../internal/domain/defect_test.go#L376) | [defect.go](../internal/domain/defect.go), [errors.go](../internal/domain/errors.go) | `801f055` → `940825d` |
| [Rechazar la resolucion con un sprint que no existe](specs/sprint-2/15-gestionar-defectos.md?plain=1#L299) | `TestRechazarLaResolucionConUnSprintQueNoExiste` [defect_test.go:361](../internal/domain/defect_test.go#L361) | [defect.go](../internal/domain/defect.go) | `28905e7` → `4f5f746` |
| [Un defecto resuelto no se vuelve a tocar](specs/sprint-2/15-gestionar-defectos.md?plain=1#L306) | `TestUnDefectoResueltoNoSeVuelveATocar` [defect_test.go:415](../internal/domain/defect_test.go#L415) | [defect.go](../internal/domain/defect.go), [errors.go](../internal/domain/errors.go) | `16c3069` → `7eacf55` |
| [Listar los defectos del proyecto](specs/sprint-2/15-gestionar-defectos.md?plain=1#L318) | `TestListarLosDefectosDelProyecto` [defect_test.go:453](../internal/domain/defect_test.go#L453) | [defect.go](../internal/domain/defect.go) | `65af2a1` → `3a51103` |
| [Los defectos de un proyecto no aparecen en otro](specs/sprint-2/15-gestionar-defectos.md?plain=1#L330) | `TestLosDefectosDeUnProyectoNoAparecenEnOtro` [defect_test.go:488](../internal/domain/defect_test.go#L488) | el que ya estaba | `f686721` (sin RED: ya pasaba) |
| [Operar sobre algo que no existe](specs/sprint-2/15-gestionar-defectos.md?plain=1#L336) | `TestOperarSobreAlgoQueNoExisteConDefectos` [defect_test.go:528](../internal/domain/defect_test.go#L528) | ⚠️ no se encontró el commit del test | — |

**Tests que no citan un escenario de la Fase 1 (3).**

| Test | Cubre | Código | RED → GREEN |
|---|---|---|---|
| `TestUnaSeveridadQueNoExisteSeMuestraSinSeveridad` [defect_test.go:197](../internal/domain/defect_test.go#L197) | RN-6 | el que ya estaba | `9e53b08` (sin RED: ya pasaba) |
| `TestRechazarUnEstadoQueNoExiste` [defect_test.go:328](../internal/domain/defect_test.go#L328) | RN-9, RN-18 | [defect.go](../internal/domain/defect.go), [errors.go](../internal/domain/errors.go) | `90b1696` → `e06e055` |
| `TestConvertirElTextoEnUnEstadoDeDefecto` [defect_test.go:589](../internal/domain/defect_test.go#L589) | RN-9, RN-10 | [defect.go](../internal/domain/defect.go), [errors.go](../internal/domain/errors.go), [registry_json.go](../internal/domain/registry_json.go) | `d92cfbd` → `809a617` |

## #77 Interfaz web: proyectos, integrantes y backlog

- **Historia:** issue #77
- **Spec:** [docs/specs/sprint-2/77-interfaz-web-base.md](specs/sprint-2/77-interfaz-web-base.md)
- **Código según la spec:** `internal/web/`, `internal/store/file.go`, `internal/domain/registry_json.go`, `internal/domain/story_estimation.go`, `cmd/web/main.go`

<details><summary>Criterios de aceptación (14)</summary>

- El sistema se abre en el navegador con un solo comando y funciona sin internet.
- Desde la web se crean, modifican y consultan proyectos.
- Desde la web se agregan y quitan integrantes, con su rol.
- Desde la web se administra el Product Backlog: alta, modificación, prioridad y estimación en story points.
- Los errores de validación se muestran junto al campo que falló, con el mensaje del sistema, sin perder lo que se había escrito.
- Las listas vacías muestran que no hay datos y qué se puede hacer.
- La página se puede usar en una pantalla de notebook y en un celular.
- Los datos cargados no se pierden al cerrar y volver a abrir el sistema.
- Cada operación se confirma con un mensaje, y recargar la página no la repite.
- Antes de quitar algo se pide confirmación, y cancelar no cambia nada.
- La barra de navegación muestra siempre en qué proyecto se está.
- Cada campo tiene su etiqueta, se ve el foco y todo se usa con el teclado.
- Si los datos guardados no se pueden leer, el sistema avisa y no los pisa.
- La web no contiene ninguna validación de negocio propia.

</details>

| Escenario BDD | Test | Código | RED → GREEN |
|---|---|---|---|
| [Crear un proyecto desde la web](specs/sprint-2/77-interfaz-web-base.md?plain=1#L156) | `TestCrearUnProyectoDesdeLaWeb` [projects_test.go:15](../internal/web/projects_test.go#L15) | [projects.go](../internal/web/projects.go), [server.go](../internal/web/server.go), [proyecto-nuevo.html](../internal/web/templates/proyecto-nuevo.html), [proyecto.html](../internal/web/templates/proyecto.html), [proyectos.html](../internal/web/templates/proyectos.html) | `8de188c` → `9723e86` |
| [Cada operacion se confirma con un mensaje](specs/sprint-2/77-interfaz-web-base.md?plain=1#L163) | `TestCadaOperacionSeConfirmaConUnMensaje` [server_test.go:935](../internal/web/server_test.go#L935) | [projects.go](../internal/web/projects.go), [server.go](../internal/web/server.go), [proyecto-form.html](../internal/web/templates/proyecto-form.html), [proyecto.html](../internal/web/templates/proyecto.html) | `b13d561` → `6160b7d` |
| [Consultar un proyecto](specs/sprint-2/77-interfaz-web-base.md?plain=1#L178) | `TestConsultarUnProyecto` [projects_test.go:69](../internal/web/projects_test.go#L69) | [projects.go](../internal/web/projects.go), [server.go](../internal/web/server.go), [app.css](../internal/web/static/app.css), [layout.html](../internal/web/templates/layout.html), [proyecto.html](../internal/web/templates/proyecto.html), [proyectos.html](../internal/web/templates/proyectos.html) | `c7624af` → `1e8dbd2` |
| [Las opciones para elegir son las del sistema](specs/sprint-2/77-interfaz-web-base.md?plain=1#L184) | `TestLasOpcionesParaElegirSonLasDelSistema` [server_test.go:1160](../internal/web/server_test.go#L1160) | [forms.go](../internal/web/forms.go), [members.go](../internal/web/members.go), [server.go](../internal/web/server.go), [integrante-form.html](../internal/web/templates/integrante-form.html), [integrantes.html](../internal/web/templates/integrantes.html), [layout.html](../internal/web/templates/layout.html) | `e9ded37` → `2a95134` |
| [Confirmar o cancelar antes de quitar](specs/sprint-2/77-interfaz-web-base.md?plain=1#L194) | `TestConfirmarOCancelarAntesDeQuitar` [server_test.go:1208](../internal/web/server_test.go#L1208) | [members.go](../internal/web/members.go), [server.go](../internal/web/server.go), [app.css](../internal/web/static/app.css), [confirmar.html](../internal/web/templates/confirmar.html), [integrantes.html](../internal/web/templates/integrantes.html) | `7907b18` → `e990442` |
| [Mostrar el rechazo sin perder lo escrito](specs/sprint-2/77-interfaz-web-base.md?plain=1#L207) | `TestMostrarElRechazoSinPerderLoEscrito` [server_test.go:417](../internal/web/server_test.go#L417) | [forms.go](../internal/web/forms.go), [projects.go](../internal/web/projects.go), [app.css](../internal/web/static/app.css), [layout.html](../internal/web/templates/layout.html), [proyecto-nuevo.html](../internal/web/templates/proyecto-nuevo.html) | `99d8188` → `20180ce` |
| [Las listas vacias dicen que hacer](specs/sprint-2/77-interfaz-web-base.md?plain=1#L225) | `TestSinDatosGuardadosSeArrancaSinProyectos` [file_test.go:54](../internal/store/file_test.go#L54) | [file.go](../internal/store/file.go) | `9e9ec91` → `fc414fb` |
| ″ | `TestLasListasVaciasDicenQueHacer` [server_test.go:311](../internal/web/server_test.go#L311) | [projects.go](../internal/web/projects.go), [proyectos.html](../internal/web/templates/proyectos.html) | `e9570b7` → `650752f` |
| [Pedir algo que no existe](specs/sprint-2/77-interfaz-web-base.md?plain=1#L237) | `TestPedirAlgoQueNoExiste` [server_test.go:1082](../internal/web/server_test.go#L1082) | el que ya estaba | `15a8336` (sin RED: ya pasaba) |
| [Un texto con signos se muestra tal cual](specs/sprint-2/77-interfaz-web-base.md?plain=1#L249) | `TestUnTextoConSignosSeMuestraTalCual` [projects_test.go:105](../internal/web/projects_test.go#L105) | el que ya estaba | `413660d` (sin RED: ya pasaba) |
| [Abrir el sistema con un solo comando y sin internet](specs/sprint-2/77-interfaz-web-base.md?plain=1#L253) | `TestAbrirElSistemaConUnSoloComandoYSinInternet` [main_test.go:26](../cmd/web/main_test.go#L26) | [projects.go](../internal/web/projects.go), [server.go](../internal/web/server.go), [app.css](../internal/web/static/app.css), [layout.html](../internal/web/templates/layout.html), [proyectos.html](../internal/web/templates/proyectos.html) | `11939e0` → `6669ec3` |
| ″ | `TestAbrirElSistemaConUnSoloComandoYSinInternet` [server_test.go:228](../internal/web/server_test.go#L228) | [projects.go](../internal/web/projects.go), [server.go](../internal/web/server.go), [app.css](../internal/web/static/app.css), [layout.html](../internal/web/templates/layout.html), [proyectos.html](../internal/web/templates/proyectos.html) | `11939e0` → `6669ec3` |
| [Los datos siguen al cerrar y volver a abrir](specs/sprint-2/77-interfaz-web-base.md?plain=1#L259) | `TestLosDatosSiguenAlCerrarYVolverAAbrir` [registry_json_test.go:291](../internal/domain/registry_json_test.go#L291) | [errors.go](../internal/domain/errors.go), [registry_json.go](../internal/domain/registry_json.go) | `1a885f0` → `61077ba` |
| ″ | `TestLoGuardadoSeVuelveALeerIgual` [file_test.go:69](../internal/store/file_test.go#L69) | [file.go](../internal/store/file.go) | `9e9ec91` → `fc414fb` |
| ″ | `TestLosDatosSiguenAlCerrarYVolverAAbrir` [server_test.go:1623](../internal/web/server_test.go#L1623) | [errors.go](../internal/domain/errors.go), [registry_json.go](../internal/domain/registry_json.go) | `1a885f0` → `61077ba` |
| [Los datos guardados estan danados](specs/sprint-2/77-interfaz-web-base.md?plain=1#L265) | `TestLosDatosGuardadosEstanDanados` [main_test.go:257](../cmd/web/main_test.go#L257) | [registry_json.go](../internal/domain/registry_json.go) | `ff7d24f` → `2ecdfad` |
| ″ | `TestLosDatosGuardadosEstanDanados` [registry_json_test.go:386](../internal/domain/registry_json_test.go#L386) | [registry_json.go](../internal/domain/registry_json.go) | `ff7d24f` → `2ecdfad` |
| ″ | `TestUnArchivoDanadoQuedaComoEstaba` [file_test.go:160](../internal/store/file_test.go#L160) | [file.go](../internal/store/file.go) | `6d73f7a` → `a83bfc5` |
| [Un cambio que no se pudo guardar se avisa](specs/sprint-2/77-interfaz-web-base.md?plain=1#L271) | `TestUnCambioQueNoSePudoGuardarSeAvisa` [file_test.go:202](../internal/store/file_test.go#L202) | [file.go](../internal/store/file.go) | `6d73f7a` → `a83bfc5` |
| ″ | `TestUnCambioQueNoSePudoGuardarSeAvisa` [server_test.go:1323](../internal/web/server_test.go#L1323) | [file.go](../internal/store/file.go) | `6d73f7a` → `a83bfc5` |
| [Dos pestanas a la vez](specs/sprint-2/77-interfaz-web-base.md?plain=1#L276) | `TestDosPestanasALaVez` [server_test.go:1348](../internal/web/server_test.go#L1348) | [server.go](../internal/web/server.go) | `4c925e3` → `1e7a4b1` |
| [Usar la pagina en un celular y con el teclado](specs/sprint-2/77-interfaz-web-base.md?plain=1#L281) | `TestUsarLaPaginaEnUnCelularYConElTeclado` [server_test.go:1409](../internal/web/server_test.go#L1409) | [historia-form.html](../internal/web/templates/historia-form.html) | `4ab0d49` → `add6a34` |

**Tests que no citan un escenario de la Fase 1 (13).**

| Test | Cubre | Código | RED → GREEN |
|---|---|---|---|
| `TestListarLosProyectosDeUnRegistroVacio` [project_registry_test.go:7](../internal/domain/project_registry_test.go#L7) | dependencia de la #77, RN-7 y RN-23 (la lista vacía no es un error) | [project_registry.go](../internal/domain/project_registry.go) | `1f94f7a` → `3e81aca` |
| `TestListarTodosLosProyectosEnOrdenDeNumero` [project_registry_test.go:22](../internal/domain/project_registry_test.go#L22) | dependencia de la #77 | [project_registry.go](../internal/domain/project_registry.go) | `1f94f7a` → `3e81aca` |
| `TestElMensajeDeDatosDanadosEstaEnCastellano` [registry_json_test.go:551](../internal/domain/registry_json_test.go#L551) | #77 RN-24 | [registry_json.go](../internal/domain/registry_json.go), [file.go](../internal/store/file.go) | `ff2b683` → `d0f0a49` |
| `TestUnaRondaReveladaSinParticipantesNoSeCarga` [registry_json_test.go:635](../internal/domain/registry_json_test.go#L635) | #77 RN-24, RN-6 de la #13 | [registry_json.go](../internal/domain/registry_json.go) | `684fcb2` → `87264b3` |
| `TestLeerLaEscalaDeStoryPoints` [story_estimation_test.go:253](../internal/domain/story_estimation_test.go#L253) | #77 RN-11 | [story_estimation.go](../internal/domain/story_estimation.go) | `4227227` → `5a6d632` |
| `TestElMensajeDeUnArchivoQueNoSePuedeLeerEstaEnCastellano` [file_test.go:234](../internal/store/file_test.go#L234) | #77 RN-24 | [registry_json.go](../internal/domain/registry_json.go), [file.go](../internal/store/file.go) | `ff2b683` → `d0f0a49` |
| `TestElMensajeDeUnCambioQueNoSePudoGuardarEstaEnCastellano` [file_test.go:251](../internal/store/file_test.go#L251) | #77 RN-25 | [file.go](../internal/store/file.go) | `6f761d9` → `3854359` |
| `TestResolverSinElegirElSprint` [defects_test.go:255](../internal/web/defects_test.go#L255) | RN-2, RN-9 | el que ya estaba | `8847aa0` (sin RED: ya pasaba) |
| `TestUnSprintQueNoEsUnNumeroSeMarcaJuntoAlCampo` [reports_test.go:124](../internal/web/reports_test.go#L124) | RN-2 | [reports.go](../internal/web/reports.go), [server.go](../internal/web/server.go), [app.css](../internal/web/static/app.css), [layout.html](../internal/web/templates/layout.html), [reporte.html](../internal/web/templates/reporte.html) | `13e67ec` → `1a58c18` |
| `TestUnPDFQueNoSePuedeGenerarSeAvisaSinEntregarArchivo` [reports_test.go:216](../internal/web/reports_test.go#L216) | RN-19 | [reports.go](../internal/web/reports.go), [server.go](../internal/web/server.go), [reporte.html](../internal/web/templates/reporte.html) | `95e5dfa` → `d1455ef` |
| `TestUnPedidoDeOtroSitioSeRechaza` [server_test.go:1744](../internal/web/server_test.go#L1744) | #77, D20 (sin usuarios ni contraseñas) | [server.go](../internal/web/server.go), [rechazado.html](../internal/web/templates/rechazado.html) | `7e4fdb9` → `d2fa1e9` |
| `TestQuitarDelSprintUnaHistoriaQueNoEstaAsignada` [sprints_test.go:190](../internal/web/sprints_test.go#L190) | RN-2 | el que ya estaba | `b7523ca` (sin RED: ya pasaba) |
| `TestCerrarConUnaFechaQueNoEsFecha` [sprints_test.go:368](../internal/web/sprints_test.go#L368) | RN-2, RN-9 | el que ya estaba | `09d23fa` (sin RED: ya pasaba) |

## #10 Cerrar un sprint

- **Historia:** issue #10
- **Spec:** [docs/specs/sprint-3/10-cerrar-sprint.md](specs/sprint-3/10-cerrar-sprint.md)
- **Código según la spec:** `internal/domain/sprint_close.go`, `internal/domain/registry_json.go`, `internal/domain/errors.go`

<details><summary>Criterios de aceptación (14)</summary>

- Se puede cerrar el sprint abierto indicando qué historias se completaron, y queda en estado Cerrado con su fecha de cierre.
- Las historias indicadas como completadas quedan Terminadas.
- Las historias del sprint que no se indicaron quedan Pendientes, de vuelta en el Product Backlog.
- El sprint cerrado conserva la lista de las historias que se le habían asignado, incluidas las que volvieron al Product Backlog.
- El sprint cerrado deja registrado cuáles de sus historias se completaron.
- Se puede cerrar un sprint sin ninguna historia completada.
- Se puede cerrar un sprint que no tiene ninguna historia asignada.
- Indicar dos veces la misma historia no es un error y no la cuenta dos veces.
- Indicar como completada una historia que no está en el sprint rechaza el cierre y no cambia ninguna historia.
- No se puede cerrar sin fecha de cierre.
- No se puede cerrar dos veces el mismo sprint.
- Después de cerrar, el proyecto no tiene sprint abierto y se puede crear el siguiente.
- Cerrar el sprint de un proyecto que no existe informa que no se encontró el proyecto.
- Los story points que el sprint registró al cerrar no cambian si después se reestima una de sus historias.

</details>

| Escenario BDD | Test | Código | RED → GREEN |
|---|---|---|---|
| [Cerrar el sprint registrando que se completo y que no](specs/sprint-3/10-cerrar-sprint.md?plain=1#L156) | `TestCerrarElSprintRegistrandoQueSeCompletoYQueNo` [sprint_close_test.go:59](../internal/domain/sprint_close_test.go#L59) | [sprint_close.go](../internal/domain/sprint_close.go) | `e2e1137` → `bf2ac19` |
| [El sprint cerrado conserva lo que se le habia comprometido](specs/sprint-3/10-cerrar-sprint.md?plain=1#L165) | `TestElSprintCerradoConservaLoQueSeLeHabiaComprometido` [sprint_close_test.go:82](../internal/domain/sprint_close_test.go#L82) | el que ya estaba | `aea7f11` (sin RED: ya pasaba) |
| [Cerrar con distinta cantidad de historias completadas](specs/sprint-3/10-cerrar-sprint.md?plain=1#L170) | `TestCerrarConDistintaCantidadDeHistoriasCompletadas` [sprint_close_test.go:96](../internal/domain/sprint_close_test.go#L96) | el que ya estaba | `e1a1b77` (sin RED: ya pasaba) |
| [Cerrar un sprint sin ninguna historia asignada](specs/sprint-3/10-cerrar-sprint.md?plain=1#L183) | `TestCerrarUnSprintSinNingunaHistoriaAsignada` [sprint_close_test.go:132](../internal/domain/sprint_close_test.go#L132) | el que ya estaba | `e1a1b77` (sin RED: ya pasaba) |
| [La misma historia indicada dos veces cuenta una sola vez](specs/sprint-3/10-cerrar-sprint.md?plain=1#L189) | `TestLaMismaHistoriaIndicadaDosVecesCuentaUnaSolaVez` [sprint_close_test.go:146](../internal/domain/sprint_close_test.go#L146) | [sprint_close.go](../internal/domain/sprint_close.go) | `e44c2af` → `e32d3e5` |
| [Rechazar el cierre si una historia completada no esta en el sprint](specs/sprint-3/10-cerrar-sprint.md?plain=1#L195) | `TestRechazarElCierreSiUnaHistoriaCompletadaNoEstaEnElSprint` [sprint_close_test.go:163](../internal/domain/sprint_close_test.go#L163) | [sprint_close.go](../internal/domain/sprint_close.go) | `ee65d1d` → `c5463df` |
| [Rechazar el cierre sin fecha de cierre](specs/sprint-3/10-cerrar-sprint.md?plain=1#L203) | `TestRechazarElCierreSinFechaDeCierre` [sprint_close_test.go:182](../internal/domain/sprint_close_test.go#L182) | [errors.go](../internal/domain/errors.go), [sprint_close.go](../internal/domain/sprint_close.go) | `ec1e59c` → `ceee9ed` |
| [Rechazar el cierre si no hay sprint abierto](specs/sprint-3/10-cerrar-sprint.md?plain=1#L210) | `TestRechazarElCierreSiNoHaySprintAbierto` [sprint_close_test.go:200](../internal/domain/sprint_close_test.go#L200) | [sprint_close.go](../internal/domain/sprint_close.go) | `0670e00` → `fe38044` |
| [No se puede cerrar dos veces el mismo sprint](specs/sprint-3/10-cerrar-sprint.md?plain=1#L216) | `TestNoSePuedeCerrarDosVecesElMismoSprint` [sprint_close_test.go:212](../internal/domain/sprint_close_test.go#L212) | el que ya estaba | `33fbd2c` (sin RED: ya pasaba) |
| [Despues de cerrar se puede crear el sprint siguiente](specs/sprint-3/10-cerrar-sprint.md?plain=1#L223) | `TestDespuesDeCerrarSePuedeCrearElSprintSiguiente` [sprint_close_test.go:27](../internal/domain/sprint_close_test.go#L27) | [sprint_close.go](../internal/domain/sprint_close.go) | `e3f37ce` → `2ebe154` |
| [Una fecha de cierre distinta de la fecha de fin se acepta](specs/sprint-3/10-cerrar-sprint.md?plain=1#L229) | `TestUnaFechaDeCierreDistintaDeLaFechaDeFinSeAcepta` [sprint_close_test.go:231](../internal/domain/sprint_close_test.go#L231) | el que ya estaba | `8b85cc2` (sin RED: ya pasaba) |
| [Los puntos registrados al cerrar no cambian si se reestima la historia](specs/sprint-3/10-cerrar-sprint.md?plain=1#L234) | `TestLosPuntosRegistradosAlCerrarNoCambianSiSeReestimaLaHistoria` [sprint_close_test.go:264](../internal/domain/sprint_close_test.go#L264) | [sprint_close.go](../internal/domain/sprint_close.go) | `f0ec86f` → `627b2b9` |
| [Cerrar el sprint de un proyecto que no existe](specs/sprint-3/10-cerrar-sprint.md?plain=1#L240) | `TestCerrarElSprintDeUnProyectoQueNoExiste` [sprint_close_test.go:249](../internal/domain/sprint_close_test.go#L249) | el que ya estaba | `2aa23a1` (sin RED: ya pasaba) |
| [Una historia sin estimar al cerrar registra cero puntos](specs/sprint-3/10-cerrar-sprint.md?plain=1#L245) | `TestUnaHistoriaSinEstimarAlCerrarRegistraCeroPuntos` [sprint_close_test.go:284](../internal/domain/sprint_close_test.go#L284) | [sprint_close.go](../internal/domain/sprint_close.go) | `f0ec86f` → `627b2b9` |

## #11 Consultar sprints anteriores

- **Historia:** issue #11
- **Spec:** [docs/specs/sprint-3/11-consultar-sprints-anteriores.md](specs/sprint-3/11-consultar-sprints-anteriores.md)
- **Código según la spec:** `internal/domain/sprint_query.go`, `p.Sprint(number)`, `internal/domain/sprint.go`

<details><summary>Criterios de aceptación (10)</summary>

- Se listan los sprints cerrados de un proyecto con su número, su goal, sus fechas, su fecha de cierre y sus historias.
- De cada sprint cerrado se ve qué historias se le asignaron y cuáles se completaron.
- El sprint abierto no aparece en el listado.
- El listado viene ordenado por número, del más viejo al más nuevo.
- Un proyecto sin sprints cerrados devuelve un listado vacío y eso no es un error.
- Se puede pedir el detalle de un sprint por su número.
- Pedir un sprint con un número que no existe informa que no se encontró el sprint.
- Consultar un sprint no lo modifica, y lo que se recibe no permite modificarlo.
- El listado de un proyecto no trae sprints de otro.
- Consultar en un proyecto que no existe informa que no se encontró el proyecto.

</details>

| Escenario BDD | Test | Código | RED → GREEN |
|---|---|---|---|
| [Listar los sprints cerrados de un proyecto](specs/sprint-3/11-consultar-sprints-anteriores.md?plain=1#L113) | `TestListarLosSprintsCerradosDeUnProyecto` [sprint_query_test.go:61](../internal/domain/sprint_query_test.go#L61) | [sprint_query.go](../internal/domain/sprint_query.go) | `252e761` → `ae68d7f` |
| [El sprint abierto no aparece en el listado](specs/sprint-3/11-consultar-sprints-anteriores.md?plain=1#L119) | `TestElSprintAbiertoNoApareceEnElListado` [sprint_query_test.go:96](../internal/domain/sprint_query_test.go#L96) | el que ya estaba | `9bdc11e` (sin RED: ya pasaba) |
| [Un proyecto sin sprints cerrados devuelve un listado vacio](specs/sprint-3/11-consultar-sprints-anteriores.md?plain=1#L124) | `TestUnProyectoSinSprintsCerradosDevuelveUnListadoVacio` [sprint_query_test.go:34](../internal/domain/sprint_query_test.go#L34) | [sprint_query.go](../internal/domain/sprint_query.go) | `b913e5c` → `0ca7b01` |
| [Ver que historias se completaron en un sprint cerrado](specs/sprint-3/11-consultar-sprints-anteriores.md?plain=1#L135) | `TestVerQueHistoriasSeCompletaronEnUnSprintCerrado` [sprint_query_test.go:134](../internal/domain/sprint_query_test.go#L134) | el que ya estaba | `957e198` (sin RED: ya pasaba) |
| [Un sprint cerrado sin historias se consulta igual](specs/sprint-3/11-consultar-sprints-anteriores.md?plain=1#L141) | `TestUnSprintCerradoSinHistoriasSeConsultaIgual` [sprint_query_test.go:153](../internal/domain/sprint_query_test.go#L153) | el que ya estaba | `a2de059` (sin RED: ya pasaba) |
| [Pedir el detalle de un sprint por su numero](specs/sprint-3/11-consultar-sprints-anteriores.md?plain=1#L147) | `TestPedirElDetalleDeUnSprintPorSuNumero` [sprint_query_test.go:114](../internal/domain/sprint_query_test.go#L114) | el que ya estaba | `af315e0` (sin RED: ya pasaba) |
| [Pedir el detalle del sprint abierto](specs/sprint-3/11-consultar-sprints-anteriores.md?plain=1#L152) | `TestPedirElDetalleDelSprintAbierto` [sprint_query_test.go:195](../internal/domain/sprint_query_test.go#L195) | el que ya estaba | `0c1ea94` (sin RED: ya pasaba) |
| [Pedir un sprint que no existe](specs/sprint-3/11-consultar-sprints-anteriores.md?plain=1#L158) | `TestPedirUnSprintQueNoExiste` [sprint_query_test.go:175](../internal/domain/sprint_query_test.go#L175) | el que ya estaba | `9d28eba` (sin RED: ya pasaba) |
| [Consultar un sprint no lo modifica](specs/sprint-3/11-consultar-sprints-anteriores.md?plain=1#L169) | `TestConsultarUnSprintNoLoModifica` [sprint_query_test.go:214](../internal/domain/sprint_query_test.go#L214) | el que ya estaba | `932af91` (sin RED: ya pasaba) |
| [Cada proyecto lista solamente sus sprints](specs/sprint-3/11-consultar-sprints-anteriores.md?plain=1#L175) | `TestCadaProyectoListaSolamenteSusSprints` [sprint_query_test.go:238](../internal/domain/sprint_query_test.go#L238) | el que ya estaba | `37400f3` (sin RED: ya pasaba) |
| [Cerrar un sprint nuevo no cambia los anteriores](specs/sprint-3/11-consultar-sprints-anteriores.md?plain=1#L181) | `TestCerrarUnSprintNuevoNoCambiaLosAnteriores` [sprint_query_test.go:263](../internal/domain/sprint_query_test.go#L263) | el que ya estaba | `1cdc45c` (sin RED: ya pasaba) |
| [Consultar los sprints de un proyecto que no existe](specs/sprint-3/11-consultar-sprints-anteriores.md?plain=1#L187) | `TestConsultarLosSprintsDeUnProyectoQueNoExiste` [sprint_query_test.go:292](../internal/domain/sprint_query_test.go#L292) | el que ya estaba | `f20e276` (sin RED: ya pasaba) |

## #13 Votar en secreto en una ronda de Planning Poker

- **Historia:** issue #13
- **Spec:** [docs/specs/sprint-3/13-votar-en-secreto.md](specs/sprint-3/13-votar-en-secreto.md)
- **Código según la spec:** `internal/domain/voting_round.go`, `internal/domain/project.go`, `internal/domain/errors.go`

<details><summary>Criterios de aceptación (14)</summary>

- Se puede abrir una ronda sobre una historia no Terminada, esté estimada o no, y queda como la ronda 1, sin revelar y sin votos.
- No se puede abrir una ronda sobre una historia Terminada ni sobre una que no existe.
- No se puede abrir una ronda en un proyecto sin integrantes.
- No se puede abrir otra ronda sobre una historia que tiene una sin revelar.
- Cada participante vota una sola vez por ronda, y el segundo voto no muestra el primero.
- El voto tiene que ser un valor de la escala 1, 2, 3, 5, 8, 13 o 21.
- Solo votan los integrantes que tenía el proyecto al abrir la ronda, y el nombre se reconoce aunque venga con otras mayúsculas o espacios de más.
- Sin revelar, nadie puede ver ningún voto, ni siquiera el propio; solo se sabe quién votó y quién falta.
- Mostrar una ronda sin revelar no deja ver ningún valor votado.
- La ronda se revela recién cuando votaron todos; antes, el pedido se rechaza.
- Al revelarse se ven todos los votos con el nombre de quien los dio.
- No se puede votar ni volver a revelar una ronda ya revelada.
- Un voto o un pedido de revelar rechazado deja la ronda como estaba.
- Se puede consultar la ronda actual de una historia, revelada o no; si no tiene ninguna, se informa y no es un error.

</details>

| Escenario BDD | Test | Código | RED → GREEN |
|---|---|---|---|
| [Abrir una ronda sobre una historia que no esta terminada](specs/sprint-3/13-votar-en-secreto.md?plain=1#L149) | `TestAbrirUnaRondaSobreUnaHistoriaQueNoEstaTerminada` [voting_round_test.go:32](../internal/domain/voting_round_test.go#L32) | [voting_round.go](../internal/domain/voting_round.go) | `03740ab` → `a3cc842` |
| [Rechazar una ronda sobre una historia que no se puede votar](specs/sprint-3/13-votar-en-secreto.md?plain=1#L163) | `TestRechazarUnaRondaSobreUnaHistoriaQueNoSePuedeVotar` [voting_round_test.go:74](../internal/domain/voting_round_test.go#L74) | [voting_round.go](../internal/domain/voting_round.go) | `689cd7b` → `8dcf72a` |
| [Rechazar una ronda en un proyecto sin integrantes](specs/sprint-3/13-votar-en-secreto.md?plain=1#L175) | `TestRechazarUnaRondaEnUnProyectoSinIntegrantes` [voting_round_test.go:105](../internal/domain/voting_round_test.go#L105) | [errors.go](../internal/domain/errors.go), [voting_round.go](../internal/domain/voting_round.go) | `50098ab` → `3ac3da3` |
| [Una historia no puede tener dos rondas en curso](specs/sprint-3/13-votar-en-secreto.md?plain=1#L181) | `TestUnaHistoriaNoPuedeTenerDosRondasEnCurso` [voting_round_test.go:124](../internal/domain/voting_round_test.go#L124) | [errors.go](../internal/domain/errors.go), [project.go](../internal/domain/project.go), [voting_round.go](../internal/domain/voting_round.go) | `29be0a2` → `72c965c` |
| [Votar registra quien voto sin dejar ver los votos, ni siquiera el propio](specs/sprint-3/13-votar-en-secreto.md?plain=1#L188) | `TestVotarRegistraQuienVotoSinDejarVerLosVotosNiSiquieraElPropio` [voting_round_test.go:161](../internal/domain/voting_round_test.go#L161) | [errors.go](../internal/domain/errors.go), [voting_round.go](../internal/domain/voting_round.go) | `ba90caa` → `6a78fd1` |
| [Revelar cuando votaron todos](specs/sprint-3/13-votar-en-secreto.md?plain=1#L195) | `TestRevelarCuandoVotaronTodos` [voting_round_test.go:188](../internal/domain/voting_round_test.go#L188) | [voting_round.go](../internal/domain/voting_round.go) | `ffc855c` → `f2bcb13` |
| [Rechazar revelar mientras falten votos](specs/sprint-3/13-votar-en-secreto.md?plain=1#L206) | `TestRechazarRevelarMientrasFaltenVotos` [voting_round_test.go:244](../internal/domain/voting_round_test.go#L244) | [errors.go](../internal/domain/errors.go), [voting_round.go](../internal/domain/voting_round.go) | `f08badb` → `a911077` |
| [Rechazar un voto fuera de la escala](specs/sprint-3/13-votar-en-secreto.md?plain=1#L218) | `TestRechazarUnVotoFueraDeLaEscala` [voting_round_test.go:278](../internal/domain/voting_round_test.go#L278) | [voting_round.go](../internal/domain/voting_round.go) | `0b37c88` → `3bfffc9` |
| [Rechazar un segundo voto del mismo integrante](specs/sprint-3/13-votar-en-secreto.md?plain=1#L232) | `TestRechazarUnSegundoVotoDelMismoIntegrante` [voting_round_test.go:300](../internal/domain/voting_round_test.go#L300) | [errors.go](../internal/domain/errors.go), [voting_round.go](../internal/domain/voting_round.go) | `16afd3a` → `0420074` |
| [Rechazar el voto de quien no participa de la ronda](specs/sprint-3/13-votar-en-secreto.md?plain=1#L239) | `TestRechazarElVotoDeQuienNoParticipaDeLaRonda` [voting_round_test.go:343](../internal/domain/voting_round_test.go#L343) | [errors.go](../internal/domain/errors.go), [voting_round.go](../internal/domain/voting_round.go) | `ef719c2` → `f7c5454` |
| [Quien sale del proyecto despues de abrir la ronda sigue participando](specs/sprint-3/13-votar-en-secreto.md?plain=1#L251) | `TestQuienSaleDelProyectoDespuesDeAbrirLaRondaSigueParticipando` [voting_round_test.go:390](../internal/domain/voting_round_test.go#L390) | [errors.go](../internal/domain/errors.go), [voting_round.go](../internal/domain/voting_round.go) | `ef719c2` → `f7c5454` |
| [Reconocer al integrante aunque el nombre venga escrito distinto](specs/sprint-3/13-votar-en-secreto.md?plain=1#L257) | `TestReconocerAlIntegranteAunqueElNombreVengaEscritoDistinto` [voting_round_test.go:375](../internal/domain/voting_round_test.go#L375) | [errors.go](../internal/domain/errors.go), [voting_round.go](../internal/domain/voting_round.go) | `ef719c2` → `f7c5454` |
| [Una ronda revelada ya no acepta cambios](specs/sprint-3/13-votar-en-secreto.md?plain=1#L263) | `TestUnaRondaReveladaYaNoAceptaCambios` [voting_round_test.go:413](../internal/domain/voting_round_test.go#L413) | [errors.go](../internal/domain/errors.go), [voting_round.go](../internal/domain/voting_round.go) | `b6e1bca` → `cc70d85` |
| [Mostrar una ronda sin revelar no deja ver los votos](specs/sprint-3/13-votar-en-secreto.md?plain=1#L275) | `TestMostrarUnaRondaSinRevelarNoDejaVerLosVotos` [voting_round_test.go:454](../internal/domain/voting_round_test.go#L454) | [voting_round.go](../internal/domain/voting_round.go) | `591adbc` → `ad93071` |
| [Consultar la ronda actual de una historia](specs/sprint-3/13-votar-en-secreto.md?plain=1#L281) | `TestConsultarLaRondaActualDeUnaHistoria` [voting_round_test.go:479](../internal/domain/voting_round_test.go#L479) | [voting_round.go](../internal/domain/voting_round.go) | `baa7833` → `ed97e3e` |
| [Consultar la ronda actual de una historia que no existe](specs/sprint-3/13-votar-en-secreto.md?plain=1#L292) | `TestConsultarLaRondaActualDeUnaHistoriaQueNoExiste` [voting_round_test.go:529](../internal/domain/voting_round_test.go#L529) | [voting_round.go](../internal/domain/voting_round.go) | `baa7833` → `ed97e3e` |

## #16 Calcular las métricas del proyecto y de cada sprint

- **Historia:** issue #16
- **Spec:** [docs/specs/sprint-3/16-calcular-metricas.md](specs/sprint-3/16-calcular-metricas.md)
- **Código según la spec:** `internal/domain/metrics.go`, `internal/domain/story_hours.go`, `internal/domain/story.go`, `internal/domain/registry_json.go`, `internal/domain/errors.go`

<details><summary>Criterios de aceptación (14)</summary>

- Los planificados de un sprint son todas sus historias comprometidas, incluidas las que volvieron al Product Backlog.
- Los completados de un sprint salen de lo registrado al cerrarlo: una historia que se completó en un sprint posterior no cuenta en el anterior.
- Los puntos de un sprint cerrado son los que tenía cada historia al cerrarlo: reestimar una historia después no los cambia.
- La velocidad de un sprint es igual a sus puntos completados, y la del proyecto es el promedio de sus sprints cerrados redondeado a entero.
- Las métricas de un sprint abierto se calculan, con completados y velocidad en cero.
- Las horas de un sprint suman las de sus historias, y las de una historia que pasó por dos sprints cuentan solo en el último.
- La desviación es la diferencia porcentual entre horas reales y estimadas, con signo, y con cero horas estimadas es cero sin error.
- El porcentaje de historias completas tiene decimales, y sin historias es cero sin error.
- Los defectos de un sprint salen de su sprint de detección y de resolución; los del proyecto son todos los registrados y los que están resueltos.
- Las métricas del proyecto incluyen todas sus historias, hayan entrado a un sprint o no.
- Un sprint o un proyecto sin datos devuelve las once métricas en cero, sin error.
- Pedir las métricas no cambia ningún dato, y si se cargan datos nuevos dan distinto.
- Se pueden cargar y reemplazar las horas estimadas de una historia no terminada, con fracciones; con cero, con un negativo o sobre una Terminada se rechaza sin cambiar nada.
- Un sprint, una historia o un proyecto inexistentes informan que no se encontraron.

</details>

| Escenario BDD | Test | Código | RED → GREEN |
|---|---|---|---|
| [Metricas de un sprint cerrado](specs/sprint-3/16-calcular-metricas.md?plain=1#L171) | `TestMetricasDeUnSprintCerrado` [metrics_test.go:122](../internal/domain/metrics_test.go#L122) | [metrics.go](../internal/domain/metrics.go) | `11f3554` → `fff2a46` |
| [Una historia que volvio al backlog cuenta como completada solo donde se completo](specs/sprint-3/16-calcular-metricas.md?plain=1#L183) | `TestUnaHistoriaQueVolvioAlBacklogCuentaComoCompletadaSoloDondeSeCompleto` [metrics_test.go:428](../internal/domain/metrics_test.go#L428) | [metrics.go](../internal/domain/metrics.go) | `9bacc10` → `8597d88` |
| [Los puntos de un sprint cerrado no cambian si se reestima una historia](specs/sprint-3/16-calcular-metricas.md?plain=1#L191) | `TestLosPuntosDeUnSprintCerradoNoCambianSiSeReestimaUnaHistoria` [metrics_test.go:464](../internal/domain/metrics_test.go#L464) | [metrics.go](../internal/domain/metrics.go) | `b6808c6` → `033613d` |
| [Las horas de una historia que paso por dos sprints cuentan en el ultimo](specs/sprint-3/16-calcular-metricas.md?plain=1#L198) | `TestLasHorasDeUnaHistoriaQuePasoPorDosSprintsCuentanEnElUltimo` [metrics_test.go:174](../internal/domain/metrics_test.go#L174) | [metrics.go](../internal/domain/metrics.go) | `cc515ff` → `b48e2be` |
| [Las metricas de un sprint abierto son parciales](specs/sprint-3/16-calcular-metricas.md?plain=1#L206) | `TestLasMetricasDeUnSprintAbiertoSonParciales` [metrics_test.go:254](../internal/domain/metrics_test.go#L254) | el que ya estaba | `1909bb4` (sin RED: ya pasaba) |
| [Sin datos cargados las metricas valen cero](specs/sprint-3/16-calcular-metricas.md?plain=1#L214) | `TestSinDatosCargadosLasMetricasValenCero` [metrics_test.go:53](../internal/domain/metrics_test.go#L53) | [metrics.go](../internal/domain/metrics.go) | `7f4a880` → `0faa0aa` |
| [Calcular la desviacion de horas](specs/sprint-3/16-calcular-metricas.md?plain=1#L225) | `TestCalcularLaDesviacionDeHoras` [metrics_test.go:214](../internal/domain/metrics_test.go#L214) | [metrics.go](../internal/domain/metrics.go) | `7bab460` → `96ed50e` |
| [Calcular el porcentaje de historias completas](specs/sprint-3/16-calcular-metricas.md?plain=1#L238) | `TestCalcularElPorcentajeDeHistoriasCompletas` [metrics_test.go:485](../internal/domain/metrics_test.go#L485) | [metrics.go](../internal/domain/metrics.go) | `b27a92c` → `0afff9a` |
| [Defectos detectados y resueltos de cada sprint](specs/sprint-3/16-calcular-metricas.md?plain=1#L250) | `TestDefectosDetectadosYResueltosDeCadaSprint` [metrics_test.go:280](../internal/domain/metrics_test.go#L280) | [metrics.go](../internal/domain/metrics.go) | `b6806d0` → `81328ca` |
| [Metricas del proyecto entero](specs/sprint-3/16-calcular-metricas.md?plain=1#L259) | `TestMetricasDelProyectoEntero` [metrics_test.go:320](../internal/domain/metrics_test.go#L320) | [metrics.go](../internal/domain/metrics.go) | `a03372b` → `d8b461e` |
| [Velocidad del proyecto](specs/sprint-3/16-calcular-metricas.md?plain=1#L271) | `TestVelocidadDelProyecto` [metrics_test.go:527](../internal/domain/metrics_test.go#L527) | [metrics.go](../internal/domain/metrics.go) | `9816638` → `ff6e1be` |
| [Las metricas reflejan los datos cargados despues](specs/sprint-3/16-calcular-metricas.md?plain=1#L282) | `TestLasMetricasReflejanLosDatosCargadosDespues` [metrics_test.go:383](../internal/domain/metrics_test.go#L383) | el que ya estaba | `771ffff` (sin RED: ya pasaba) |
| [Cargar y reemplazar las horas estimadas de una historia](specs/sprint-3/16-calcular-metricas.md?plain=1#L289) | `TestCargarYReemplazarLasHorasEstimadasDeUnaHistoria` [story_hours_test.go:18](../internal/domain/story_hours_test.go#L18) | [story.go](../internal/domain/story.go), [story_hours.go](../internal/domain/story_hours.go) | `56ff0bc` → `7e4e3b8` |
| [Rechazar horas estimadas que no son mayores que cero](specs/sprint-3/16-calcular-metricas.md?plain=1#L295) | `TestRechazarHorasEstimadasQueNoSonMayoresQueCero` [story_hours_test.go:46](../internal/domain/story_hours_test.go#L46) | [errors.go](../internal/domain/errors.go), [story_hours.go](../internal/domain/story_hours.go) | `eb4ad3d` → `e636a5c` |
| [No cambiar las horas estimadas de una historia terminada](specs/sprint-3/16-calcular-metricas.md?plain=1#L307) | `TestNoCambiarLasHorasEstimadasDeUnaHistoriaTerminada` [story_hours_test.go:79](../internal/domain/story_hours_test.go#L79) | [story_hours.go](../internal/domain/story_hours.go) | `4d24c31` → `9d00e6f` |
| [Pedir algo que no existe](specs/sprint-3/16-calcular-metricas.md?plain=1#L314) | `TestPedirAlgoQueNoExiste` [metrics_test.go:15](../internal/domain/metrics_test.go#L15) | el que ya estaba | `15a8336` (sin RED: ya pasaba) |

**Tests que no citan un escenario de la Fase 1 (2).**

| Test | Cubre | Código | RED → GREEN |
|---|---|---|---|
| `TestRechazarHorasEstimadasQueNoSonUnNumero` [story_hours_test.go:105](../internal/domain/story_hours_test.go#L105) | RN-18 | [story_hours.go](../internal/domain/story_hours.go) | `4708708` → `9dd7309` |
| `TestRechazarHorasEstimadasInfinitas` [story_hours_test.go:132](../internal/domain/story_hours_test.go#L132) | RN-18 | [story_hours.go](../internal/domain/story_hours.go) | `f0ee9fb` → `1a60943` |

## #17 Mostrar el dashboard

- **Historia:** issue #17
- **Spec:** [docs/specs/sprint-3/17-mostrar-dashboard.md](specs/sprint-3/17-mostrar-dashboard.md)
- **Código según la spec:** `internal/domain/dashboard.go`, `internal/web/dashboard.go`, `internal/web/templates/dashboard.html`

<details><summary>Criterios de aceptación (19)</summary>

- El dashboard tiene cuatro partes en orden: encabezado, estado general, sprint en curso y gráficos.
- El encabezado muestra el nombre y las fechas, o «sin fecha de fin».
- El estado general muestra los mismos números que entregó el cálculo de métricas.
- El estado general muestra la velocidad promedio del equipo, o «sin sprints cerrados» si todavía no se cerró ninguno.
- Con un sprint abierto se ven su número, su goal, sus fechas, sus story points planificados, sus horas reales y sus defectos, y no se ven completados ni velocidad.
- Sin sprint abierto, se informa que no hay sprint en curso y el resto se ve igual.
- Un proyecto sin sprints muestra el encabezado y el aviso, sin gráficos y sin error.
- Se muestran cuatro gráficos: story points, horas, defectos y velocidad.
- El de velocidad tiene una barra por sprint cerrado, ordenadas por número, sin el abierto.
- Cada barra tiene su rótulo y su valor.
- La barra del valor más grande ocupa todo el ancho del gráfico, y las demás son proporcionales.
- Un valor cero no tiene barra; uno mayor que cero ocupa al menos el 1 % del ancho; un gráfico con todos sus valores en cero se muestra sin barras y sin error.
- Porcentajes y horas van con un decimal y coma; la desviación negativa, con su signo; las fechas, como día/mes/año.
- Cada gráfico tiene su texto alternativo con sus valores, y se pueden ver como tabla.
- Las barras no se distinguen solo por el color y tienen contraste con el fondo.
- El dashboard se abre desde la barra de navegación del proyecto.
- Se ve sin internet y en un celular, sin desplazar la página de costado.
- Abrir el dashboard no cambia nada del proyecto.
- Pedir el dashboard de un proyecto que no existe informa que no se encontró.

</details>

| Escenario BDD | Test | Código | RED → GREEN |
|---|---|---|---|
| [Dashboard de un proyecto con un sprint en curso](specs/sprint-3/17-mostrar-dashboard.md?plain=1#L167) | `TestDashboardDeUnProyectoConUnSprintEnCurso` [dashboard_test.go:177](../internal/domain/dashboard_test.go#L177) | [dashboard.go](../internal/domain/dashboard.go) | `ef5ab19` → `fc0c221` |
| ″ | `TestDashboardDeUnProyectoConUnSprintEnCurso` [dashboard_test.go:444](../internal/web/dashboard_test.go#L444) | [dashboard.go](../internal/domain/dashboard.go) | `ef5ab19` → `fc0c221` |
| [Abrir el dashboard desde la barra de navegacion](specs/sprint-3/17-mostrar-dashboard.md?plain=1#L174) | `TestAbrirElDashboardDesdeLaBarraDeNavegacion` [dashboard_test.go:17](../internal/web/dashboard_test.go#L17) | [dashboard.go](../internal/web/dashboard.go), [server.go](../internal/web/server.go), [dashboard.html](../internal/web/templates/dashboard.html), [layout.html](../internal/web/templates/layout.html) | `9bf4313` → `132571d` |
| [Encabezado de un proyecto sin fecha de finalizacion](specs/sprint-3/17-mostrar-dashboard.md?plain=1#L179) | `TestEncabezadoDeUnProyectoSinFechaDeFinalizacion` [dashboard_test.go:52](../internal/web/dashboard_test.go#L52) | [dashboard.go](../internal/web/dashboard.go), [server.go](../internal/web/server.go), [dashboard.html](../internal/web/templates/dashboard.html), [layout.html](../internal/web/templates/layout.html) | `9bf4313` → `132571d` |
| [El estado general muestra las metricas tal como se calcularon](specs/sprint-3/17-mostrar-dashboard.md?plain=1#L185) | `TestElEstadoGeneralMuestraLasMetricasTalComoSeCalcularon` [dashboard_test.go:49](../internal/domain/dashboard_test.go#L49) | [dashboard.go](../internal/domain/dashboard.go), [format.go](../internal/domain/format.go) | `5941e21` → `a3487f5` |
| ″ | `TestElEstadoGeneralMuestraLasMetricasTalComoSeCalcularon` [dashboard_test.go:131](../internal/web/dashboard_test.go#L131) | [dashboard.go](../internal/domain/dashboard.go), [format.go](../internal/domain/format.go) | `5941e21` → `a3487f5` |
| [Porcentajes, horas y enteros se muestran con su formato](specs/sprint-3/17-mostrar-dashboard.md?plain=1#L194) | `TestPorcentajesHorasYEnterosSeMuestranConSuFormato` [dashboard_test.go:83](../internal/domain/dashboard_test.go#L83) | [format.go](../internal/domain/format.go) | `92c6a8b` → `b66feed` |
| [El sprint en curso muestra su meta y lo que lleva registrado](specs/sprint-3/17-mostrar-dashboard.md?plain=1#L211) | `TestElSprintEnCursoMuestraSuMetaYLoQueLlevaRegistrado` [dashboard_test.go:127](../internal/domain/dashboard_test.go#L127) | [dashboard.go](../internal/domain/dashboard.go) | `6330fb4` → `e84abc3` |
| ″ | `TestElSprintEnCursoMuestraSuMetaYLoQueLlevaRegistrado` [dashboard_test.go:162](../internal/web/dashboard_test.go#L162) | [dashboard.go](../internal/domain/dashboard.go) | `6330fb4` → `e84abc3` |
| [Un proyecto con todos sus sprints cerrados no tiene sprint en curso](specs/sprint-3/17-mostrar-dashboard.md?plain=1#L218) | `TestUnProyectoConTodosSusSprintsCerradosNoTieneSprintEnCursoEnElDashboard` [dashboard_test.go:159](../internal/domain/dashboard_test.go#L159) | [dashboard.go](../internal/domain/dashboard.go) | `6330fb4` → `e84abc3` |
| ″ | `TestUnProyectoConTodosSusSprintsCerradosNoTieneSprintEnCurso` [dashboard_test.go:195](../internal/web/dashboard_test.go#L195) | [project_status.go](../internal/domain/project_status.go) | `27cd558` → `96b9633` |
| [Un proyecto sin sprints avisa que no hay datos](specs/sprint-3/17-mostrar-dashboard.md?plain=1#L224) | `TestUnProyectoSinSprintsAvisaQueNoHayDatos` [dashboard_test.go:28](../internal/domain/dashboard_test.go#L28) | [dashboard.go](../internal/domain/dashboard.go) | `6367a0e` → `1b11f92` |
| ″ | `TestUnProyectoSinSprintsAvisaQueNoHayDatos` [dashboard_test.go:80](../internal/web/dashboard_test.go#L80) | [dashboard.go](../internal/domain/dashboard.go) | `6367a0e` → `1b11f92` |
| [El largo de cada barra es proporcional al valor mas grande del grafico](specs/sprint-3/17-mostrar-dashboard.md?plain=1#L232) | `TestElLargoDeCadaBarraEsProporcionalAlValorMasGrandeDelGrafico` [dashboard_test.go:242](../internal/web/dashboard_test.go#L242) | [dashboard.go](../internal/web/dashboard.go) | `a176c80` → `1833a8d` |
| [Las barras de un grafico se miden contra su valor mas grande](specs/sprint-3/17-mostrar-dashboard.md?plain=1#L246) | `TestLasBarrasDeUnGraficoSeMidenContraSuValorMasGrande` [dashboard_test.go:309](../internal/web/dashboard_test.go#L309) | [dashboard.go](../internal/web/dashboard.go), [app.css](../internal/web/static/app.css), [dashboard.html](../internal/web/templates/dashboard.html) | `210f45d` → `cc37e50` |
| [El grafico de velocidad tiene una barra por sprint cerrado](specs/sprint-3/17-mostrar-dashboard.md?plain=1#L262) | `TestElGraficoDeVelocidadTieneUnaBarraPorSprintCerrado` [dashboard_test.go:219](../internal/domain/dashboard_test.go#L219) | [dashboard.go](../internal/domain/dashboard.go) | `1be2d30` → `2e28197` |
| ″ | `TestElGraficoDeVelocidadTieneUnaBarraPorSprintCerrado` [dashboard_test.go:383](../internal/web/dashboard_test.go#L383) | [dashboard.go](../internal/domain/dashboard.go) | `1be2d30` → `2e28197` |
| [Todavia no hay sprints cerrados para el grafico de velocidad](specs/sprint-3/17-mostrar-dashboard.md?plain=1#L268) | `TestTodaviaNoHaySprintsCerradosParaElGraficoDeVelocidad` [dashboard_test.go:244](../internal/domain/dashboard_test.go#L244) | [dashboard.go](../internal/domain/dashboard.go) | `1be2d30` → `2e28197` |
| ″ | `TestTodaviaNoHaySprintsCerradosParaElGraficoDeVelocidad` [dashboard_test.go:426](../internal/web/dashboard_test.go#L426) | [dashboard.go](../internal/domain/dashboard.go) | `1be2d30` → `2e28197` |
| [Cada grafico tiene su texto alternativo y su tabla](specs/sprint-3/17-mostrar-dashboard.md?plain=1#L275) | `TestCadaGraficoTieneSuTextoAlternativoYSuTabla` [dashboard_test.go:485](../internal/web/dashboard_test.go#L485) | [dashboard.go](../internal/web/dashboard.go), [dashboard.html](../internal/web/templates/dashboard.html) | `fae9c92` → `d7803df` |
| [Abrir el dashboard no cambia el proyecto](specs/sprint-3/17-mostrar-dashboard.md?plain=1#L281) | `TestAbrirElDashboardNoCambiaElProyecto` [dashboard_test.go:284](../internal/domain/dashboard_test.go#L284) | el que ya estaba | `4f8e6d1` (sin RED: ya pasaba) |
| [Pedir el dashboard de un proyecto que no existe](specs/sprint-3/17-mostrar-dashboard.md?plain=1#L287) | `TestPedirElDashboardDeUnProyectoQueNoExiste` [dashboard_test.go:220](../internal/web/dashboard_test.go#L220) | [dashboard.go](../internal/web/dashboard.go), [server.go](../internal/web/server.go), [dashboard.html](../internal/web/templates/dashboard.html), [layout.html](../internal/web/templates/layout.html) | `9bf4313` → `132571d` |
| [Usar el dashboard en un celular y sin internet](specs/sprint-3/17-mostrar-dashboard.md?plain=1#L293) | `TestUsarElDashboardEnUnCelularYSinInternet` [dashboard_test.go:544](../internal/web/dashboard_test.go#L544) | el que ya estaba | `ad3b799` (sin RED: ya pasaba) |

**Tests que no citan un escenario de la Fase 1 (2).**

| Test | Cubre | Código | RED → GREEN |
|---|---|---|---|
| `TestElDashboardSePideAlProyectoYUsaLasMetricasDeLa16` [dashboard_test.go:312](../internal/domain/dashboard_test.go#L312) | RN-2 | [dashboard.go](../internal/domain/dashboard.go) | `792ee9a` → `acc6ff5` |
| `TestElDashboardMuestraLosNumerosQueCalculaLaHistoria16` [dashboard_test.go:602](../internal/web/dashboard_test.go#L602) | RN-2 | el que ya estaba | `13048ad` (sin RED: ya pasaba) |

## #177 Guardar los datos en una base de datos

- **Historia:** issue #177
- **Spec:** [docs/specs/sprint-3/177-guardar-en-base-de-datos.md](specs/sprint-3/177-guardar-en-base-de-datos.md)
- **Código según la spec:** `internal/store/database.go`, `rows.go`, `exchange.go`, `migrations.go`, `migrations/0001_inicial.sql`, `cmd/web/main.go`

<details><summary>Criterios de aceptación (7)</summary>

- Todo lo que se carga desde la web sigue igual al cerrar y volver a abrir el sistema.
- Un cambio que no se pudo guardar no deja ninguna parte guardada: la base queda como estaba, y el cambio entra con el próximo guardado que funcione.
- Al abrir, cada dato pasa por las mismas reglas que al cargarlo desde la web; uno que no las cumple se avisa y el sistema no abre.
- Exportar a un archivo e importarlo en una base nueva da exactamente los mismos datos.
- Los datos de demo y lo guardado con la versión anterior se pueden importar; lo de la versión anterior se importa solo al abrir con la base vacía, y se avisa.
- Importar un archivo con datos que no cumplen las reglas, o sobre una base con datos, no cambia nada y se avisa.
- Funciona en Linux, en macOS y en Windows sin instalar nada aparte, y los tests pasan en los dos sistemas del CI.

</details>

| Escenario BDD | Test | Código | RED → GREEN |
|---|---|---|---|
| [Los datos siguen al cerrar y volver a abrir](specs/sprint-3/177-guardar-en-base-de-datos.md?plain=1#L192) | `TestUnRegistroCompletoVuelveIgual` [database_test.go:209](../internal/store/database_test.go#L209) | [database.go](../internal/store/database.go) | `b305866` → `487bac5` |
| ″ | `TestLosDatosSiguenAlCerrarYVolverAAbrir` [server_test.go:1623](../internal/web/server_test.go#L1623) | [errors.go](../internal/domain/errors.go), [registry_json.go](../internal/domain/registry_json.go) | `1a885f0` → `61077ba` |
| [Un cambio que no se pudo guardar no deja nada a medias](specs/sprint-3/177-guardar-en-base-de-datos.md?plain=1#L199) | `TestUnGuardadoQueFallaNoDejaNadaAMedias` [database_test.go:323](../internal/store/database_test.go#L323) | el que ya estaba | `b8e6140` (sin RED: ya pasaba) |
| ″ | `TestUnGuardadoQueFallaEnCualquierTablaNoDejaNada` [database_test.go:439](../internal/store/database_test.go#L439) | el que ya estaba | `e3863eb` (sin RED: ya pasaba) |
| [Un dato guardado que no cumple las reglas no se carga](specs/sprint-3/177-guardar-en-base-de-datos.md?plain=1#L208) | `TestLoQueNoSePuedeAbrirSeAvisaYNoSeToca` [database_test.go:364](../internal/store/database_test.go#L364) | [database.go](../internal/store/database.go) | `d23737e` → `f6ffce5` |
| ″ | `TestUnaFilaQueApuntaANadaNoSeCarga` [database_test.go:467](../internal/store/database_test.go#L467) | el que ya estaba | `e3863eb` (sin RED: ya pasaba) |
| [Los datos difíciles vuelven iguales](specs/sprint-3/177-guardar-en-base-de-datos.md?plain=1#L215) | `TestLosDatosDificilesVuelvenIguales` [database_test.go:247](../internal/store/database_test.go#L247) | el que ya estaba | `582d427` (sin RED: ya pasaba) |
| [Exportar e importar da los mismos datos](specs/sprint-3/177-guardar-en-base-de-datos.md?plain=1#L228) | `TestExportarEImportarDesdeElComando` [main_test.go:345](../cmd/web/main_test.go#L345) | [main.go](../cmd/web/main.go), [exchange.go](../internal/store/exchange.go) | `61d92a4` → `50cc4c3` |
| ″ | `TestExportarEImportarDaLoMismo` [exchange_test.go:16](../internal/store/exchange_test.go#L16) | [main.go](../cmd/web/main.go), [exchange.go](../internal/store/exchange.go) | `61d92a4` → `50cc4c3` |
| [Traer los datos de la versión anterior](specs/sprint-3/177-guardar-en-base-de-datos.md?plain=1#L234) | `TestExportarEImportarDesdeElComando` [main_test.go:345](../cmd/web/main_test.go#L345) | [main.go](../cmd/web/main.go), [exchange.go](../internal/store/exchange.go) | `61d92a4` → `50cc4c3` |
| ″ | `TestTraerLosDatosDeDemo` [main_test.go:609](../cmd/web/main_test.go#L609) | el que ya estaba | `d5c5689` (sin RED: ya pasaba) |
| [Importar algo que no se puede no cambia nada](specs/sprint-3/177-guardar-en-base-de-datos.md?plain=1#L240) | `TestImportarSobreUnaBaseConDatosSeRechaza` [main_test.go:400](../cmd/web/main_test.go#L400) | [database.go](../internal/store/database.go), [exchange.go](../internal/store/exchange.go) | `e2f87d9` → `beca251` |
| ″ | `TestImportarAlgoQueNoSePuedeNoCambiaNada` [exchange_test.go:42](../internal/store/exchange_test.go#L42) | [database.go](../internal/store/database.go), [exchange.go](../internal/store/exchange.go) | `e2f87d9` → `beca251` |
| [Lo guardado con la versión anterior se importa solo](specs/sprint-3/177-guardar-en-base-de-datos.md?plain=1#L252) | `TestElArchivoDeLaVersionAnteriorSeImportaAlAbrir` [main_test.go:430](../cmd/web/main_test.go#L430) | [main.go](../cmd/web/main.go), [exchange.go](../internal/store/exchange.go) | `df065cc` → `641e723` |
| ″ | `TestElArchivoViejoSeImportaUnaSolaVez` [exchange_test.go:97](../internal/store/exchange_test.go#L97) | [main.go](../cmd/web/main.go), [exchange.go](../internal/store/exchange.go) | `df065cc` → `641e723` |
| [Lo de la versión anterior no vuelve aunque se borre todo](specs/sprint-3/177-guardar-en-base-de-datos.md?plain=1#L262) | `TestElArchivoViejoSeImportaUnaSolaVez` [exchange_test.go:97](../internal/store/exchange_test.go#L97) | [main.go](../cmd/web/main.go), [exchange.go](../internal/store/exchange.go) | `df065cc` → `641e723` |
| [Si lo de la versión anterior no se puede importar, no se toca](specs/sprint-3/177-guardar-en-base-de-datos.md?plain=1#L267) | `TestElArchivoViejoSeImportaUnaSolaVez` [exchange_test.go:97](../internal/store/exchange_test.go#L97) | [main.go](../cmd/web/main.go), [exchange.go](../internal/store/exchange.go) | `df065cc` → `641e723` |
| [Abrir por primera vez](specs/sprint-3/177-guardar-en-base-de-datos.md?plain=1#L276) | `TestLaWebGuardaEnLaBase` [main_test.go:300](../cmd/web/main_test.go#L300) | [main.go](../cmd/web/main.go) | `c218f64` → `f210ae8` |
| ″ | `TestAbrirUnaBaseNuevaCreaElEsquema` [database_test.go:41](../internal/store/database_test.go#L41) | [database.go](../internal/store/database.go), [migrations.go](../internal/store/migrations.go) | `5c1b0be` → `46f6516` |
| ″ | `TestGuardarYCargarUnRegistro` [database_test.go:99](../internal/store/database_test.go#L99) | [database.go](../internal/store/database.go), [rows.go](../internal/store/rows.go) | `5ea2432` → `5ad0f3b` |
| ″ | `TestLaBaseSeAbreEnLaRutaQueSeIndica` [database_test.go:534](../internal/store/database_test.go#L534) | [database.go](../internal/store/database.go) | `ece4fe9` → `20aec82` |
| [Una base de una versión más nueva no se abre](specs/sprint-3/177-guardar-en-base-de-datos.md?plain=1#L282) | `TestLoQueNoSePuedeAbrirSeAvisaYNoSeToca` [database_test.go:364](../internal/store/database_test.go#L364) | [database.go](../internal/store/database.go) | `d23737e` → `f6ffce5` |
| [Lo que no se puede pedir al exportar o importar](specs/sprint-3/177-guardar-en-base-de-datos.md?plain=1#L288) | `TestLoQueNoSePuedePedirAlExportarOImportar` [main_test.go:469](../cmd/web/main_test.go#L469) | [exchange.go](../internal/store/exchange.go) | `088e9d2` → `acd91b6` |
| ″ | `TestImportarYExportarJuntosOSinBase` [main_test.go:499](../cmd/web/main_test.go#L499) | [main.go](../cmd/web/main.go) | `f77f4e0` → `498b921` |
| ″ | `TestUnArchivoJSONComoBaseSeRechaza` [main_test.go:534](../cmd/web/main_test.go#L534) | [main.go](../cmd/web/main.go) | `dd4dc73` → `82fcd5f` |
| ″ | `TestExportarAlMismoArchivoDeLaBaseSeRechaza` [exchange_test.go:207](../internal/store/exchange_test.go#L207) | [exchange.go](../internal/store/exchange.go) | `088e9d2` → `acd91b6` |
| [Una base abierta no la abre otro sistema](specs/sprint-3/177-guardar-en-base-de-datos.md?plain=1#L301) | `TestExportarConLaWebAbiertaAvisaQueSeCierreLaWeb` [main_test.go:576](../cmd/web/main_test.go#L576) | [main.go](../cmd/web/main.go) | `0c662ab` → `cd2ee78` |
| ″ | `TestUnaBaseAbiertaNoLaAbreOtroSistema` [database_test.go:497](../internal/store/database_test.go#L497) | [database.go](../internal/store/database.go) | `02e7135` → `f6ec463` |

**Tests que no citan un escenario de la Fase 1 (2).**

| Test | Cubre | Código | RED → GREEN |
|---|---|---|---|
| `TestUnaBaseQueNoSePuedeCrearSeAvisa` [database_test.go:421](../internal/store/database_test.go#L421) | #177 RN-2, RN-4 | el que ya estaba | `9a1c02c` (sin RED: ya pasaba) |
| `TestUnDatoQueLaBaseNoConoceNoSeGuarda` [rows_test.go:11](../internal/store/rows_test.go#L11) | #177 D5, RN-2 | el que ya estaba | `b36baae` (sin RED: ya pasaba) |

## #18 Generar el reporte de un proyecto o de un sprint

- **Historia:** issue #18
- **Spec:** [docs/specs/sprint-3/18-generar-reporte.md](specs/sprint-3/18-generar-reporte.md)
- **Código según la spec:** `internal/domain/report.go`

<details><summary>Criterios de aceptación (12)</summary>

- Se puede generar el reporte de un sprint y el del proyecto entero, cada uno con su título, su día de generación y su alcance.
- El reporte tiene siempre las cuatro secciones en orden, y una sección sin datos dice que no hay.
- Todos los números del reporte coinciden con las métricas del mismo alcance.
- El Resumen de un sprint tiene proyecto, Sprint Goal, fechas y estado, y si el sprint está abierto avisa que los números son parciales.
- El Resumen del proyecto tiene nombre, fechas y cantidad de sprints.
- Incluye las historias planificadas y completadas, una por línea, con sus story points o "sin estimar".
- En el reporte de un sprint, cada historia figura como completada o no según el cierre de ese sprint.
- Incluye las estimaciones y el esfuerzo registrado: puntos, velocidad, horas estimadas y reales, y desviación.
- Incluye los defectos del período, uno por línea, con su severidad y su estado.
- Horas y porcentajes van con un decimal, coma y su unidad, la desviación negativa con su signo y la positiva sin signo, la velocidad sin decimales, y las fechas como día/mes/año.
- Generar el reporte no cambia ningún dato.
- Pedir el reporte de un sprint o de un proyecto inexistente informa que no se encontró.

</details>

| Escenario BDD | Test | Código | RED → GREEN |
|---|---|---|---|
| [Generar el reporte de un sprint cerrado](specs/sprint-3/18-generar-reporte.md?plain=1#L154) | `TestGenerarElReporteDeUnSprintCerrado` [report_test.go:31](../internal/domain/report_test.go#L31) | [report.go](../internal/domain/report.go) | `c6723df` → `74ee04d` |
| [Los numeros del reporte son los de las metricas](specs/sprint-3/18-generar-reporte.md?plain=1#L162) | `TestLosNumerosDelReporteSonLosDeLasMetricas` [report_test.go:68](../internal/domain/report_test.go#L68) | [report.go](../internal/domain/report.go) | `7ac4da6` → `ec5f0d8` |
| [El resumen dice en que estado esta el sprint](specs/sprint-3/18-generar-reporte.md?plain=1#L174) | `TestElResumenDiceEnQueEstadoEstaElSprint` [report_test.go:130](../internal/domain/report_test.go#L130) | [format.go](../internal/domain/format.go), [report.go](../internal/domain/report.go) | `bdf773f` → `26b23ff` |
| [Las historias del sprint figuran segun el cierre de ese sprint](specs/sprint-3/18-generar-reporte.md?plain=1#L185) | `TestLasHistoriasDelSprintFiguranSegunElCierreDeEseSprint` [report_test.go:170](../internal/domain/report_test.go#L170) | [report.go](../internal/domain/report.go) | `1771942` → `f92d2de` |
| [Una historia sin estimar figura como sin estimar](specs/sprint-3/18-generar-reporte.md?plain=1#L193) | `TestUnaHistoriaSinEstimarFiguraComoSinEstimar` [report_test.go:209](../internal/domain/report_test.go#L209) | [report.go](../internal/domain/report.go) | `56f50dd` → `e9e8980` |
| [Las historias de un sprint abierto figuran en curso](specs/sprint-3/18-generar-reporte.md?plain=1#L198) | `TestLasHistoriasDeUnSprintAbiertoFiguranEnCurso` [report_test.go:240](../internal/domain/report_test.go#L240) | [report.go](../internal/domain/report.go) | `6bf634c` → `e8739ff` |
| [Una seccion sin datos lo dice](specs/sprint-3/18-generar-reporte.md?plain=1#L203) | `TestUnaSeccionSinDatosLoDice` [report_test.go:294](../internal/domain/report_test.go#L294) | [report.go](../internal/domain/report.go) | `a955296` → `8c31622` |
| [Los defectos del sprint son los detectados o resueltos en el](specs/sprint-3/18-generar-reporte.md?plain=1#L214) | `TestLosDefectosDelSprintSonLosDetectadosOResueltosEnEl` [report_test.go:258](../internal/domain/report_test.go#L258) | [report.go](../internal/domain/report.go) | `0a9ac30` → `abf261b` |
| [Generar el reporte del proyecto entero](specs/sprint-3/18-generar-reporte.md?plain=1#L222) | `TestGenerarElReporteDelProyectoEntero` [report_test.go:323](../internal/domain/report_test.go#L323) | [report.go](../internal/domain/report.go) | `6387269` → `41f9ffd` |
| [Un proyecto sin fecha de fin](specs/sprint-3/18-generar-reporte.md?plain=1#L233) | `TestUnProyectoSinFechaDeFin` [report_test.go:389](../internal/domain/report_test.go#L389) | [report.go](../internal/domain/report.go) | `d747ae2` → `4b5d2eb` |
| [Un proyecto sin sprints cerrados no muestra una velocidad de cero](specs/sprint-3/18-generar-reporte.md?plain=1#L238) | `TestUnProyectoSinSprintsCerradosNoMuestraUnaVelocidadDeCero` [report_test.go:401](../internal/domain/report_test.go#L401) | [report.go](../internal/domain/report.go) | `d747ae2` → `4b5d2eb` |
| [Formato de los numeros](specs/sprint-3/18-generar-reporte.md?plain=1#L244) | `TestFormatoDeLosNumeros` [report_test.go:102](../internal/domain/report_test.go#L102) | [report.go](../internal/domain/report.go) | `d0ed15b` → `f56067d` |
| [Las historias y los defectos van ordenados por identificador](specs/sprint-3/18-generar-reporte.md?plain=1#L257) | `TestLasHistoriasYLosDefectosVanOrdenadosPorIdentificador` [report_test.go:418](../internal/domain/report_test.go#L418) | el que ya estaba | `745dff2` (sin RED: ya pasaba) |
| [Las historias de un sprint van ordenadas por identificador](specs/sprint-3/18-generar-reporte.md?plain=1#L264) | `TestLasHistoriasDeUnSprintVanOrdenadasPorIdentificador` [report_test.go:467](../internal/domain/report_test.go#L467) | [report.go](../internal/domain/report.go) | `0a3098d` → `6d7b256` |
| [El resumen del proyecto cuenta los sprints en singular o en plural](specs/sprint-3/18-generar-reporte.md?plain=1#L274) | `TestElResumenDelProyectoCuentaLosSprintsEnSingularOEnPlural` [report_test.go:593](../internal/domain/report_test.go#L593) | [report.go](../internal/domain/report.go) | `688cb5c` → `9fc560c` |
| [Pedir un reporte de algo que no existe](specs/sprint-3/18-generar-reporte.md?plain=1#L285) | `TestPedirUnReporteDeAlgoQueNoExiste` [report_test.go:555](../internal/domain/report_test.go#L555) | el que ya estaba | `61299e0` (sin RED: ya pasaba) |

**Tests que no citan un escenario de la Fase 1 (1).**

| Test | Cubre | Código | RED → GREEN |
|---|---|---|---|
| `TestLasOperacionesDelReportePidenLasMetricasYNoCambianNada` [report_test.go:507](../internal/domain/report_test.go#L507) ⚠️ no dice que no tiene escenario | RN-2, RN-15 | [report.go](../internal/domain/report.go) | `9b8b67e` → `d989b1b` |

## #78 Interfaz web: sprints, esfuerzo, defectos, Planning Poker y reportes

- **Historia:** issue #78
- **Spec:** [docs/specs/sprint-3/78-interfaz-web-operaciones.md](specs/sprint-3/78-interfaz-web-operaciones.md)
- **Código según la spec:** `internal/web/`, `cmd/web/main.go`

<details><summary>Criterios de aceptación (13)</summary>

- Desde la web se crea un sprint, se le asignan historias, se cierra y se consultan los anteriores.
- Desde la web se registra el esfuerzo real y se ven las horas por historia y por integrante.
- Desde la web se cargan defectos, se cambia su estado, se resuelven y se filtran por estado y por severidad.
- Desde la web se hace una ronda de Planning Poker: cada integrante vota sin ver los votos de los demás hasta que se revelan, y se acuerda la estimación.
- Desde la web se ve el reporte del proyecto o de un sprint, y se descarga en PDF.
- Los errores se muestran igual que en el resto de la interfaz web.
- Antes de revelar, ningún voto aparece en ninguna página ni mensaje, ni siquiera el propio.
- Antes de cerrar un sprint o resolver un defecto se pide confirmación, y cancelar no cambia nada.
- La fecha de cierre y el día del esfuerzo vienen propuestos con la fecha de hoy.
- Desde la web se cargan las horas estimadas de una historia.
- La barra de un proyecto lleva a sus sprints, esfuerzo, defectos, Planning Poker y reporte.
- Los sprints, el esfuerzo, los defectos y las rondas siguen al cerrar y volver a abrir.
- La web no contiene ninguna validación de negocio propia.

</details>

| Escenario BDD | Test | Código | RED → GREEN |
|---|---|---|---|
| [La barra de un proyecto lleva a todas sus secciones](specs/sprint-3/78-interfaz-web-operaciones.md?plain=1#L220) | `TestLaBarraDeUnProyectoLlevaATodasSusSecciones` [sprints_test.go:20](../internal/web/sprints_test.go#L20) | [server.go](../internal/web/server.go), [sprints.go](../internal/web/sprints.go), [layout.html](../internal/web/templates/layout.html), [sprints.html](../internal/web/templates/sprints.html) | `fb7c03e` → `56d1f95` |
| [Crear un sprint desde la web](specs/sprint-3/78-interfaz-web-operaciones.md?plain=1#L225) | `TestCrearUnSprintDesdeLaWeb` [sprints_test.go:62](../internal/web/sprints_test.go#L62) | [server.go](../internal/web/server.go), [sprints.go](../internal/web/sprints.go), [sprint-form.html](../internal/web/templates/sprint-form.html), [sprints.html](../internal/web/templates/sprints.html) | `cc2b293` → `211237f` |
| [Asignar y quitar historias del sprint abierto](specs/sprint-3/78-interfaz-web-operaciones.md?plain=1#L232) | `TestAsignarYQuitarHistoriasDelSprintAbierto` [sprints_test.go:122](../internal/web/sprints_test.go#L122) | [server.go](../internal/web/server.go), [sprints.go](../internal/web/sprints.go), [layout.html](../internal/web/templates/layout.html), [sprint.html](../internal/web/templates/sprint.html), [sprints.html](../internal/web/templates/sprints.html) | `f3a96de` → `5077131` |
| [Cerrar el sprint marcando las completadas](specs/sprint-3/78-interfaz-web-operaciones.md?plain=1#L239) | `TestCerrarElSprintMarcandoLasCompletadas` [sprints_test.go:243](../internal/web/sprints_test.go#L243) | [forms.go](../internal/web/forms.go), [server.go](../internal/web/server.go), [sprints.go](../internal/web/sprints.go), [app.css](../internal/web/static/app.css), [sprint-cerrar.html](../internal/web/templates/sprint-cerrar.html), [sprint.html](../internal/web/templates/sprint.html) | `8b60975` → `3550e3f` |
| [Consultar un sprint anterior](specs/sprint-3/78-interfaz-web-operaciones.md?plain=1#L249) | `TestConsultarUnSprintAnterior` [sprints_test.go:331](../internal/web/sprints_test.go#L331) | [sprints.go](../internal/web/sprints.go), [sprint.html](../internal/web/templates/sprint.html), [sprints.html](../internal/web/templates/sprints.html) | `e4476d7` → `d7b0785` |
| [Registrar esfuerzo y ver los totales](specs/sprint-3/78-interfaz-web-operaciones.md?plain=1#L255) | `TestRegistrarEsfuerzoYVerLosTotales` [efforts_test.go:15](../internal/web/efforts_test.go#L15) | [main.go](../cmd/web/main.go), [efforts.go](../internal/web/efforts.go), [server.go](../internal/web/server.go), [esfuerzo-form.html](../internal/web/templates/esfuerzo-form.html), [esfuerzo.html](../internal/web/templates/esfuerzo.html), [layout.html](../internal/web/templates/layout.html) | `587fd47` → `97bb41b` |
| [Cargar las horas estimadas de una historia](specs/sprint-3/78-interfaz-web-operaciones.md?plain=1#L264) | `TestCargarLasHorasEstimadasDeUnaHistoria` [efforts_test.go:69](../internal/web/efforts_test.go#L69) | [efforts.go](../internal/web/efforts.go), [forms.go](../internal/web/forms.go), [server.go](../internal/web/server.go), [historia.html](../internal/web/templates/historia.html), [horas-estimadas.html](../internal/web/templates/horas-estimadas.html) | `e8124f5` → `cf1d745` |
| [Cargar un defecto, pasarlo a en curso y resolverlo](specs/sprint-3/78-interfaz-web-operaciones.md?plain=1#L270) | `TestCargarUnDefectoPasarloAEnCursoYResolverlo` [defects_test.go:50](../internal/web/defects_test.go#L50) | [defects.go](../internal/web/defects.go), [server.go](../internal/web/server.go), [defecto-form.html](../internal/web/templates/defecto-form.html), [defecto.html](../internal/web/templates/defecto.html), [defectos.html](../internal/web/templates/defectos.html), [layout.html](../internal/web/templates/layout.html) | `3b14453` → `9611795` |
| [Filtrar los defectos](specs/sprint-3/78-interfaz-web-operaciones.md?plain=1#L280) | `TestFiltrarLosDefectos` [defects_test.go:188](../internal/web/defects_test.go#L188) | [defects.go](../internal/web/defects.go), [defectos.html](../internal/web/templates/defectos.html) | `9090625` → `7066111` |
| [Una ronda de Planning Poker sin ver los votos hasta revelar](specs/sprint-3/78-interfaz-web-operaciones.md?plain=1#L291) | `TestUnaRondaDePlanningPokerSinVerLosVotosHastaRevelar` [poker_test.go:28](../internal/web/poker_test.go#L28) | [poker.go](../internal/web/poker.go), [server.go](../internal/web/server.go), [layout.html](../internal/web/templates/layout.html), [planning-poker.html](../internal/web/templates/planning-poker.html), [ronda.html](../internal/web/templates/ronda.html) | `8f37017` → `4297f72` |
| [Nadie ve un voto antes de revelar](specs/sprint-3/78-interfaz-web-operaciones.md?plain=1#L303) | `TestNadieVeUnVotoAntesDeRevelar` [poker_test.go:142](../internal/web/poker_test.go#L142) | [poker.go](../internal/web/poker.go) | `fff8bab` → `cba14ef` |
| [Un voto rechazado no muestra el valor elegido](specs/sprint-3/78-interfaz-web-operaciones.md?plain=1#L310) | `TestUnVotoRechazadoNoMuestraElValorElegido` [poker_test.go:181](../internal/web/poker_test.go#L181) | [poker.go](../internal/web/poker.go) | `fff8bab` → `cba14ef` |
| [No se puede revelar con votos faltantes](specs/sprint-3/78-interfaz-web-operaciones.md?plain=1#L316) | `TestNoSePuedeRevelarConVotosFaltantes` [poker_test.go:87](../internal/web/poker_test.go#L87) | [poker.go](../internal/web/poker.go), [server.go](../internal/web/server.go), [ronda.html](../internal/web/templates/ronda.html) | `5acab35` → `06a195a` |
| [Diferencias, nueva ronda y acuerdo](specs/sprint-3/78-interfaz-web-operaciones.md?plain=1#L322) | `TestDiferenciasNuevaRondaYAcuerdo` [poker_test.go:215](../internal/web/poker_test.go#L215) | [poker.go](../internal/web/poker.go), [server.go](../internal/web/server.go), [ronda.html](../internal/web/templates/ronda.html), [rondas.html](../internal/web/templates/rondas.html) | `43e3e5a` → `4ad7a5b` |
| [Ver el reporte](specs/sprint-3/78-interfaz-web-operaciones.md?plain=1#L333) | `TestVerElReporte` [reports_test.go:48](../internal/web/reports_test.go#L48) | [reports.go](../internal/web/reports.go), [server.go](../internal/web/server.go), [app.css](../internal/web/static/app.css), [layout.html](../internal/web/templates/layout.html), [reporte.html](../internal/web/templates/reporte.html) | `13e67ec` → `1a58c18` |
| [Descargar el reporte en PDF](specs/sprint-3/78-interfaz-web-operaciones.md?plain=1#L344) | `TestDescargarElReporteEnPDF` [reports_test.go:156](../internal/web/reports_test.go#L156) | [reports.go](../internal/web/reports.go), [server.go](../internal/web/server.go), [reporte.html](../internal/web/templates/reporte.html) | `95e5dfa` → `d1455ef` |
| [Confirmar o cancelar antes de lo que no se deshace](specs/sprint-3/78-interfaz-web-operaciones.md?plain=1#L350) | `TestConfirmarOCancelarAntesDeResolverUnDefecto` [defects_test.go:137](../internal/web/defects_test.go#L137) | [defects.go](../internal/web/defects.go), [server.go](../internal/web/server.go), [defecto.html](../internal/web/templates/defecto.html) | `92d4364` → `bf4b303` |
| ″ | `TestConfirmarOCancelarAntesDeCerrarElSprint` [sprints_test.go:297](../internal/web/sprints_test.go#L297) | [forms.go](../internal/web/forms.go), [server.go](../internal/web/server.go), [sprints.go](../internal/web/sprints.go), [app.css](../internal/web/static/app.css), [sprint-cerrar.html](../internal/web/templates/sprint-cerrar.html), [sprint.html](../internal/web/templates/sprint.html) | `8b60975` → `3550e3f` |
| [Mostrar el rechazo sin perder lo escrito](specs/sprint-3/78-interfaz-web-operaciones.md?plain=1#L363) | `TestMostrarElRechazoSinPerderLoEscrito` [server_test.go:417](../internal/web/server_test.go#L417) | [forms.go](../internal/web/forms.go), [projects.go](../internal/web/projects.go), [app.css](../internal/web/static/app.css), [layout.html](../internal/web/templates/layout.html), [proyecto-nuevo.html](../internal/web/templates/proyecto-nuevo.html) | `99d8188` → `20180ce` |
| [Las listas vacias dicen que hacer](specs/sprint-3/78-interfaz-web-operaciones.md?plain=1#L380) | `TestLasListasVaciasDicenQueHacer` [server_test.go:311](../internal/web/server_test.go#L311) | [projects.go](../internal/web/projects.go), [proyectos.html](../internal/web/templates/proyectos.html) | `e9570b7` → `650752f` |
| [Pedir algo que no existe](specs/sprint-3/78-interfaz-web-operaciones.md?plain=1#L393) | `TestPedirAlgoQueNoExiste` [server_test.go:1082](../internal/web/server_test.go#L1082) | el que ya estaba | `15a8336` (sin RED: ya pasaba) |
| [El proyecto muestra su resumen](specs/sprint-3/78-interfaz-web-operaciones.md?plain=1#L405) | `TestElProyectoMuestraSuResumen` [projects_test.go:127](../internal/web/projects_test.go#L127) | [projects.go](../internal/web/projects.go), [proyecto.html](../internal/web/templates/proyecto.html) | `e10078c` → `b951ba3` |
| [Los datos nuevos siguen al cerrar y volver a abrir](specs/sprint-3/78-interfaz-web-operaciones.md?plain=1#L410) | `TestLosDatosNuevosSiguenAlCerrarYVolverAAbrir` [server_test.go:1677](../internal/web/server_test.go#L1677) | el que ya estaba | `97ef3a2` (sin RED: ya pasaba) |

**Tests que no citan un escenario de la Fase 1 (6).**

| Test | Cubre | Código | RED → GREEN |
|---|---|---|---|
| `TestUnDefectoYaResueltoDesdeUnaPaginaViejaMuestraElMensaje` [defects_test.go:275](../internal/web/defects_test.go#L275) ⚠️ no dice que no tiene escenario | #78 RN-2; #15 RN-14 | [defects.go](../internal/web/defects.go), [sprints.go](../internal/web/sprints.go), [defecto.html](../internal/web/templates/defecto.html), [sprint-cerrar.html](../internal/web/templates/sprint-cerrar.html) | `e3343b2` → `32033f9` |
| `TestResolverEnUnSprintQueNoExisteAvisaAntesDePreguntar` [defects_test.go:296](../internal/web/defects_test.go#L296) ⚠️ no dice que no tiene escenario | #78 RN-2, RN-9; #15 RN-12 | [defects.go](../internal/web/defects.go), [sprints.go](../internal/web/sprints.go), [defecto.html](../internal/web/templates/defecto.html), [sprint-cerrar.html](../internal/web/templates/sprint-cerrar.html) | `e3343b2` → `32033f9` |
| `TestUnFiltroEnMinusculasQuedaElegido` [defects_test.go:315](../internal/web/defects_test.go#L315) ⚠️ no dice que no tiene escenario | #78 RN-13, D7 | [defects.go](../internal/web/defects.go) | `ae34a7c` → `fd6228b` |
| `TestLasDiferenciasSeparanConComasTresOMasNombres` [poker_test.go:289](../internal/web/poker_test.go#L289) ⚠️ no dice que no tiene escenario | #78 RN-8 | [poker.go](../internal/web/poker.go) | `60ba49d` → `8304d32` |
| `TestUnaRondaVotadaDesdeDosPestanas` [poker_test.go:367](../internal/web/poker_test.go#L367) | #78 RN-4, RN-5, RN-7, RN-8; RN-20 de la #77 | el que ya estaba | `70d8529` (sin RED: ya pasaba) |
| `TestCerrarDesdeUnaPaginaViejaAvisaEnElFormulario` [sprints_test.go:390](../internal/web/sprints_test.go#L390) ⚠️ no dice que no tiene escenario | #78 RN-2 (RN-13 de la #77); #10 RN-6 | [defects.go](../internal/web/defects.go), [sprints.go](../internal/web/sprints.go), [defecto.html](../internal/web/templates/defecto.html), [sprint-cerrar.html](../internal/web/templates/sprint-cerrar.html) | `e3343b2` → `32033f9` |

## #19 Exportar el reporte a PDF

- **Historia:** issue #19
- **Spec:** [docs/specs/sprint-4/19-exportar-reporte-pdf.md](specs/sprint-4/19-exportar-reporte-pdf.md)
- **Código según la spec:** `internal/domain/report_pdf.go`, `r.WritePDF`, `Report`, `internal/domain/report.go`, `internal/domain/errors.go`

<details><summary>Criterios de aceptación (10)</summary>

- Exportar un reporte da un archivo PDF, que la web ofrece para descargar.
- El archivo se abre en Vista Previa y en un navegador y muestra el texto del reporte, que se puede seleccionar y buscar.
- El archivo tiene el título, la fecha de generación y todas las secciones con todas sus líneas, en el mismo orden que el reporte en pantalla, sin nada agregado ni quitado.
- El título va más grande y los encabezados en negrita.
- Una línea que no entra sigue en el renglón siguiente y nada queda fuera del margen.
- Un reporte que no entra en una página sigue en otra, cada página dice "Página N de M" y ningún encabezado queda solo al pie.
- Las tildes y la ñ se ven bien, y un carácter que no se puede representar sale como "?" sin impedir la exportación.
- Los paréntesis y las barras del texto no impiden abrir el archivo.
- Si el PDF no se puede entregar entero, se informa.
- Un reporte de proyecto y uno de sprint se exportan de la misma forma.

</details>

| Escenario BDD | Test | Código | RED → GREEN |
|---|---|---|---|
| [Exportar el reporte de un sprint a PDF](specs/sprint-4/19-exportar-reporte-pdf.md?plain=1#L128) | `TestExportarElReporteDeUnSprintAPDF` [report_pdf_test.go:46](../internal/domain/report_pdf_test.go#L46) | [report_pdf.go](../internal/domain/report_pdf.go) | `15d3634` → `b77dfb3` |
| ″ | `TestExportarElReporteDeUnSprintAPDFTextoYLetras` [report_pdf_test.go:425](../internal/domain/report_pdf_test.go#L425) | [report_pdf.go](../internal/domain/report_pdf.go) | `15f3415` → `87f84a8` |
| [El archivo se abre en un lector de PDF](specs/sprint-4/19-exportar-reporte-pdf.md?plain=1#L135) | `TestElArchivoSeAbreEnUnLectorDePDF` [report_pdf_test.go:349](../internal/domain/report_pdf_test.go#L349) | [report_pdf.go](../internal/domain/report_pdf.go) | `625c17b` → `6b90bf1` |
| [El PDF no agrega ni quita nada del reporte](specs/sprint-4/19-exportar-reporte-pdf.md?plain=1#L141) | `TestElPDFNoAgregaNiQuitaNadaDelReporte` [report_pdf_test.go:576](../internal/domain/report_pdf_test.go#L576) | el que ya estaba | `9711eb6` (sin RED: ya pasaba) |
| [Una linea larga sigue en el renglon siguiente](specs/sprint-4/19-exportar-reporte-pdf.md?plain=1#L152) | `TestUnaLineaLargaSigueEnElRenglonSiguiente` [report_pdf_test.go:135](../internal/domain/report_pdf_test.go#L135) | [report_pdf.go](../internal/domain/report_pdf.go) | `48a0757` → `a781f46` |
| [Una palabra mas larga que el renglon se corta en el margen](specs/sprint-4/19-exportar-reporte-pdf.md?plain=1#L158) | `TestUnaPalabraMasLargaQueElRenglonSeCortaEnElMargen` [report_pdf_test.go:189](../internal/domain/report_pdf_test.go#L189) | [report_pdf.go](../internal/domain/report_pdf.go) | `e083a2d` → `fa863dc` |
| [Un reporte largo ocupa varias paginas numeradas](specs/sprint-4/19-exportar-reporte-pdf.md?plain=1#L163) | `TestUnReporteLargoOcupaVariasPaginasNumeradas` [report_pdf_test.go:236](../internal/domain/report_pdf_test.go#L236) | [report_pdf.go](../internal/domain/report_pdf.go) | `b87ac8a` → `cae1730` |
| [Un titulo largo sigue en varios renglones con su letra](specs/sprint-4/19-exportar-reporte-pdf.md?plain=1#L175) | `TestUnTituloLargoSigueEnVariosRenglonesConSuLetra` [report_pdf_test.go:165](../internal/domain/report_pdf_test.go#L165) | [report_pdf.go](../internal/domain/report_pdf.go) | `48a0757` → `a781f46` |
| [Una linea vacia queda como renglon en blanco](specs/sprint-4/19-exportar-reporte-pdf.md?plain=1#L181) | `TestUnaLineaVaciaQuedaComoRenglonEnBlanco` [report_pdf_test.go:72](../internal/domain/report_pdf_test.go#L72) | [report_pdf.go](../internal/domain/report_pdf.go) | `15d3634` → `b77dfb3` |
| [Exportar dos veces el mismo reporte da el mismo archivo](specs/sprint-4/19-exportar-reporte-pdf.md?plain=1#L186) | `TestExportarDosVecesElMismoReporteDaElMismoArchivo` [report_pdf_test.go:605](../internal/domain/report_pdf_test.go#L605) | el que ya estaba | `916d3cb` (sin RED: ya pasaba) |
| [Un encabezado nunca queda solo al pie de una pagina](specs/sprint-4/19-exportar-reporte-pdf.md?plain=1#L191) | `TestUnEncabezadoNuncaQuedaSoloAlPieDeUnaPagina` [report_pdf_test.go:275](../internal/domain/report_pdf_test.go#L275) | [report_pdf.go](../internal/domain/report_pdf.go) | `cce7625` → `b919c4e` |
| [Letras del castellano y caracteres que no se pueden representar](specs/sprint-4/19-exportar-reporte-pdf.md?plain=1#L196) | `TestLetrasDelCastellanoYCaracteresQueNoSePuedenRepresentar` [report_pdf_test.go:479](../internal/domain/report_pdf_test.go#L479) | [report_pdf.go](../internal/domain/report_pdf.go) | `7a46cb9` → `109e222` |
| [Los parentesis y las barras no rompen el archivo](specs/sprint-4/19-exportar-reporte-pdf.md?plain=1#L208) | `TestLosParentesisYLasBarrasNoRompenElArchivo` [report_pdf_test.go:458](../internal/domain/report_pdf_test.go#L458) | [report_pdf.go](../internal/domain/report_pdf.go) | `7cab343` → `745621b` |
| [Un PDF que no se puede entregar se informa](specs/sprint-4/19-exportar-reporte-pdf.md?plain=1#L213) | `TestUnPDFQueNoSePuedeEntregarSeInforma` [report_pdf_test.go:519](../internal/domain/report_pdf_test.go#L519) | [errors.go](../internal/domain/errors.go), [report_pdf.go](../internal/domain/report_pdf.go) | `05d3a80` → `81215d9` |

**Tests que no citan un escenario de la Fase 1 (4).**

| Test | Cubre | Código | RED → GREEN |
|---|---|---|---|
| `TestUnReporteSinSeccionesTieneUnaPagina` [report_pdf_test.go:89](../internal/domain/report_pdf_test.go#L89) | RN-3 | [report_pdf.go](../internal/domain/report_pdf.go) | `15d3634` → `b77dfb3` |
| `TestAnchosDeHelvetica` [report_pdf_test.go:103](../internal/domain/report_pdf_test.go#L103) | RN-4 | [report_pdf.go](../internal/domain/report_pdf.go) | `48a0757` → `a781f46` |
| `TestUnEncabezadoSinLineasAlFinalNoAgregaPaginaVacia` [report_pdf_test.go:298](../internal/domain/report_pdf_test.go#L298) | RN-7 | [report_pdf.go](../internal/domain/report_pdf.go) | `cce7625` → `b919c4e` |
| `TestLosReportesQueGeneraLaHistoria18SeExportanTalCual` [report_pdf_test.go:625](../internal/domain/report_pdf_test.go#L625) | RN-2, RN-3, RN-13 | [report_pdf.go](../internal/domain/report_pdf.go) | `20ac523` → `709b953` |

## #73 Resolver diferencias y acordar la estimación

- **Historia:** issue #73
- **Spec:** [docs/specs/sprint-4/73-resolver-diferencias-y-acordar.md](specs/sprint-4/73-resolver-diferencias-y-acordar.md)
- **Código según la spec:** `internal/domain/voting_round_agreement.go`, `internal/domain/voting_round.go`, `internal/domain/errors.go`

<details><summary>Criterios de aceptación (11)</summary>

- Con la ronda revelada, el sistema informa si hay diferencias entre los votos.
- Informa el voto más bajo y el más alto, con el nombre de todos los que dieron cada uno.
- Si todos votaron lo mismo, informa que no hay diferencias.
- Con la ronda sin revelar no informa nada de los votos, ni siquiera si hay diferencias.
- Se puede abrir una ronda nueva sobre la misma historia cuando la última está revelada, y las anteriores quedan guardadas como estaban.
- Las rondas de cada historia se numeran en orden, empezando en 1, y se consultan en ese orden.
- Se puede registrar la estimación acordada, y la historia queda estimada con ese valor.
- La estimación acordada tiene que ser un valor de la escala, aunque nadie lo haya votado.
- No se puede acordar sobre una ronda sin revelar, ya acordada o que no es la última de su historia.
- No se puede acordar sobre una historia Terminada.
- Un acuerdo rechazado no cambia ni la ronda ni la historia.

</details>

| Escenario BDD | Test | Código | RED → GREEN |
|---|---|---|---|
| [Informar las diferencias de una ronda revelada](specs/sprint-4/73-resolver-diferencias-y-acordar.md?plain=1#L138) | `TestInformarLasDiferenciasDeUnaRondaRevelada` [voting_round_agreement_test.go:47](../internal/domain/voting_round_agreement_test.go#L47) | [voting_round_agreement.go](../internal/domain/voting_round_agreement.go) | `2ab9574` → `84e068b` |
| [Sin diferencias cuando todos votaron lo mismo](specs/sprint-4/73-resolver-diferencias-y-acordar.md?plain=1#L150) | `TestSinDiferenciasCuandoTodosVotaronLoMismo` [voting_round_agreement_test.go:83](../internal/domain/voting_round_agreement_test.go#L83) | [voting_round_agreement.go](../internal/domain/voting_round_agreement.go) | `2ab9574` → `84e068b` |
| [No se informan diferencias de una ronda sin revelar](specs/sprint-4/73-resolver-diferencias-y-acordar.md?plain=1#L161) | `TestNoSeInformanDiferenciasDeUnaRondaSinRevelar` [voting_round_agreement_test.go:118](../internal/domain/voting_round_agreement_test.go#L118) | [voting_round_agreement.go](../internal/domain/voting_round_agreement.go) | `2ab9574` → `84e068b` |
| [Acordar la estimacion de una ronda revelada](specs/sprint-4/73-resolver-diferencias-y-acordar.md?plain=1#L167) | `TestAcordarLaEstimacionDeUnaRondaRevelada` [voting_round_agreement_test.go:142](../internal/domain/voting_round_agreement_test.go#L142) | [voting_round.go](../internal/domain/voting_round.go), [voting_round_agreement.go](../internal/domain/voting_round_agreement.go) | `34f60c4` → `de58f0d` |
| [Una ronda revelada sin acuerdo figura como sin estimacion acordada](specs/sprint-4/73-resolver-diferencias-y-acordar.md?plain=1#L179) | `TestUnaRondaReveladaSinAcuerdoFiguraComoSinEstimacionAcordada` [voting_round_agreement_test.go:175](../internal/domain/voting_round_agreement_test.go#L175) | [voting_round.go](../internal/domain/voting_round.go), [voting_round_agreement.go](../internal/domain/voting_round_agreement.go) | `34f60c4` → `de58f0d` |
| [Rechazar un acuerdo que la estimacion no permite](specs/sprint-4/73-resolver-diferencias-y-acordar.md?plain=1#L185) | `TestRechazarUnAcuerdoQueLaEstimacionNoPermite` [voting_round_agreement_test.go:185](../internal/domain/voting_round_agreement_test.go#L185) | [voting_round.go](../internal/domain/voting_round.go), [voting_round_agreement.go](../internal/domain/voting_round_agreement.go) | `34f60c4` → `de58f0d` |
| [Rechazar un acuerdo sobre una ronda sin revelar](specs/sprint-4/73-resolver-diferencias-y-acordar.md?plain=1#L199) | `TestRechazarUnAcuerdoSobreUnaRondaSinRevelar` [voting_round_agreement_test.go:225](../internal/domain/voting_round_agreement_test.go#L225) | [errors.go](../internal/domain/errors.go), [voting_round_agreement.go](../internal/domain/voting_round_agreement.go) | `5808f95` → `b13c00e` |
| [Rechazar un segundo acuerdo sobre la misma ronda](specs/sprint-4/73-resolver-diferencias-y-acordar.md?plain=1#L206) | `TestRechazarUnSegundoAcuerdoSobreLaMismaRonda` [voting_round_agreement_test.go:248](../internal/domain/voting_round_agreement_test.go#L248) | [errors.go](../internal/domain/errors.go), [voting_round_agreement.go](../internal/domain/voting_round_agreement.go) | `5120a6d` → `01e8f5d` |
| [Estimar una historia en dos rondas](specs/sprint-4/73-resolver-diferencias-y-acordar.md?plain=1#L213) | `TestEstimarUnaHistoriaEnDosRondas` [voting_round_agreement_test.go:291](../internal/domain/voting_round_agreement_test.go#L291) | [voting_round.go](../internal/domain/voting_round.go), [voting_round_agreement.go](../internal/domain/voting_round_agreement.go) | `a6a5a5a` → `185bce2` |
| [Cada historia numera sus rondas desde 1](specs/sprint-4/73-resolver-diferencias-y-acordar.md?plain=1#L222) | `TestCadaHistoriaNumeraSusRondasDesde1` [voting_round_agreement_test.go:321](../internal/domain/voting_round_agreement_test.go#L321) | [voting_round.go](../internal/domain/voting_round.go), [voting_round_agreement.go](../internal/domain/voting_round_agreement.go) | `a6a5a5a` → `185bce2` |
| [Abrir una ronda nueva despues de acordar](specs/sprint-4/73-resolver-diferencias-y-acordar.md?plain=1#L227) | `TestAbrirUnaRondaNuevaDespuesDeAcordar` [voting_round_agreement_test.go:343](../internal/domain/voting_round_agreement_test.go#L343) | [voting_round.go](../internal/domain/voting_round.go), [voting_round_agreement.go](../internal/domain/voting_round_agreement.go) | `a6a5a5a` → `185bce2` |
| [Consultar las rondas de una historia](specs/sprint-4/73-resolver-diferencias-y-acordar.md?plain=1#L234) | `TestConsultarLasRondasDeUnaHistoria` [voting_round_agreement_test.go:367](../internal/domain/voting_round_agreement_test.go#L367) | [voting_round.go](../internal/domain/voting_round.go), [voting_round_agreement.go](../internal/domain/voting_round_agreement.go) | `a6a5a5a` → `185bce2` |
| [Solo se acuerda sobre la ultima ronda de la historia](specs/sprint-4/73-resolver-diferencias-y-acordar.md?plain=1#L244) | `TestSoloSeAcuerdaSobreLaUltimaRondaDeLaHistoria` [voting_round_agreement_test.go:391](../internal/domain/voting_round_agreement_test.go#L391) | [errors.go](../internal/domain/errors.go), [voting_round_agreement.go](../internal/domain/voting_round_agreement.go) | `b1c418c` → `c1ce166` |

