# Trazabilidad

_Generado por `scripts/trazabilidad.sh`. No editar a mano: se regenera con el script._

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
| [#3 Crear y modificar proyectos](#3-crear-y-modificar-proyectos) | 20 | 20 | 0 | 0 |
| [#4 Registrar integrantes de un proyecto](#4-registrar-integrantes-de-un-proyecto) | 13 | 13 | 0 | 11 |
| [#6 Administrar el Product Backlog](#6-administrar-el-product-backlog) | 18 | 17 | 1 | 14 |
| [#12 Estimar historias con Story Points](#12-estimar-historias-con-story-points) | 10 | 10 | 0 | 3 |
| [#77 Interfaz web: proyectos, integrantes y backlog](#77-interfaz-web-proyectos-integrantes-y-backlog) | 15 | 4 | 11 | 6 |
| [Tests de ninguna historia](#tests-de-ninguna-historia) | — | — | — | 1 |

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
| [Dos proyectos con el mismo nombre](specs/sprint-1/03-crear-proyectos.md?plain=1#L145) | `TestDosProyectosConElMismoNombre` [project_test.go:112](../internal/domain/project_test.go#L112) | ⚠️ ningún commit de código después del test | `239e570` → ⚠️ sin GREEN |
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
| [Agregar el primer integrante](specs/sprint-1/04-registrar-integrantes.md?plain=1#L120) | `TestAgregarElPrimerIntegrante` [member_test.go:44](../internal/domain/member_test.go#L44) | [member.go](../internal/domain/member.go) | `8398ab5` → `c3ed475` |
| [Rechazar un integrante sin nombre y apellido](specs/sprint-1/04-registrar-integrantes.md?plain=1#L126) | `TestRechazarUnIntegranteSinNombreYApellido` [member_test.go:94](../internal/domain/member_test.go#L94) | [errors.go](../internal/domain/errors.go), [member.go](../internal/domain/member.go) | `27fa739` → `2f225f6` |
| [Guardar el nombre sin espacios en los extremos](specs/sprint-1/04-registrar-integrantes.md?plain=1#L138) | `TestGuardarElNombreDelIntegranteSinEspaciosEnLosExtremos` [member_test.go:72](../internal/domain/member_test.go#L72) | [member.go](../internal/domain/member.go) | `1821eb3` → `ad6d7df` |
| [Rechazar un rol invalido](specs/sprint-1/04-registrar-integrantes.md?plain=1#L143) | `TestRechazarUnRolInvalido` [member_test.go:122](../internal/domain/member_test.go#L122) | [member.go](../internal/domain/member.go) | `0216054` → `2e8beef` |
| [Rechazar un integrante duplicado](specs/sprint-1/04-registrar-integrantes.md?plain=1#L150) | `TestRechazarUnIntegranteDuplicado` [member_test.go:170](../internal/domain/member_test.go#L170) | [errors.go](../internal/domain/errors.go), [member.go](../internal/domain/member.go) | `598a735` → `db034c1` |
| [La misma persona en dos proyectos](specs/sprint-1/04-registrar-integrantes.md?plain=1#L163) | `TestLaMismaPersonaEnDosProyectos` [member_test.go:244](../internal/domain/member_test.go#L244) | ⚠️ ningún commit de código después del test | `ac05cb0` → ⚠️ sin GREEN |
| [Rechazar un segundo Agile Enabler](specs/sprint-1/04-registrar-integrantes.md?plain=1#L170) | `TestRechazarUnSegundoAgileEnabler` [member_test.go:273](../internal/domain/member_test.go#L273) | [errors.go](../internal/domain/errors.go), [member.go](../internal/domain/member.go) | `5432431` → `1da3059` |
| [Agregar un Agile Enabler despues de quitar al anterior](specs/sprint-1/04-registrar-integrantes.md?plain=1#L177) | `TestAgregarUnAgileEnablerDespuesDeQuitarAlAnterior` [member_test.go:418](../internal/domain/member_test.go#L418) | [member.go](../internal/domain/member.go) | `17736ed` → `c3b82b7` |
| [Listar los integrantes de un proyecto](specs/sprint-1/04-registrar-integrantes.md?plain=1#L183) | `TestListarLosIntegrantesDeUnProyecto` [member_test.go:479](../internal/domain/member_test.go#L479) | ⚠️ ningún commit de código después del test | `5b56329` → ⚠️ sin GREEN |
| [Listar un proyecto sin integrantes](specs/sprint-1/04-registrar-integrantes.md?plain=1#L189) | `TestListarUnProyectoSinIntegrantes` [member_test.go:28](../internal/domain/member_test.go#L28) | [member.go](../internal/domain/member.go), [project.go](../internal/domain/project.go) | `008b6ef` → `cc89760` |
| [Quitar el ultimo integrante](specs/sprint-1/04-registrar-integrantes.md?plain=1#L194) | `TestQuitarElUltimoIntegrante` [member_test.go:366](../internal/domain/member_test.go#L366) | [member.go](../internal/domain/member.go) | `17736ed` → `c3b82b7` |
| [Quitar un integrante que no pertenece al proyecto](specs/sprint-1/04-registrar-integrantes.md?plain=1#L199) | `TestQuitarUnIntegranteQueNoPerteneceAlProyecto` [member_test.go:447](../internal/domain/member_test.go#L447) | [errors.go](../internal/domain/errors.go), [member.go](../internal/domain/member.go) | `a29c335` → `102ee66` |
| [Operar sobre un proyecto que no existe](specs/sprint-1/04-registrar-integrantes.md?plain=1#L205) | `TestOperarSobreIntegrantesDeUnProyectoQueNoExiste` [member_test.go:550](../internal/domain/member_test.go#L550) | ⚠️ ningún commit de código después del test | `3e3d580` → ⚠️ sin GREEN |

**Tests que no citan un escenario de la Fase 1 (11).**

| Test | Cubre | Código | RED → GREEN |
|---|---|---|---|
| `TestUnIntegranteExponeSuNombreYSuRol` [member_test.go:12](../internal/domain/member_test.go#L12) | RN-1 | [member.go](../internal/domain/member.go) | `61fc94d` → `adf4e3a` |
| `TestUnRolSinInicializarNoEsNingunoDeLosValidos` [member_test.go:154](../internal/domain/member_test.go#L154) ⚠️ no dice que no tiene escenario | RN-4 | [member.go](../internal/domain/member.go) | `0216054` → `2e8beef` |
| `TestNoSonDuplicadosLosNombresQueDifierenEnTildesOEspaciosDelMedio` [member_test.go:210](../internal/domain/member_test.go#L210) ⚠️ no dice que no tiene escenario | RN-5 | [errors.go](../internal/domain/errors.go), [member.go](../internal/domain/member.go) | `598a735` → `db034c1` |
| `TestConUnAgileEnablerSePuedenAgregarProductBuilders` [member_test.go:297](../internal/domain/member_test.go#L297) ⚠️ no dice que no tiene escenario | RN-7 | [errors.go](../internal/domain/errors.go), [member.go](../internal/domain/member.go) | `5432431` → `1da3059` |
| `TestConVariosErroresSeInformaElPrimero` [member_test.go:322](../internal/domain/member_test.go#L322) | RN-2, RN-4, RN-5, RN-7, RN-12 | ⚠️ ningún commit de código después del test | `c95cd0c` → ⚠️ sin GREEN |
| `TestQuitarUnIntegranteDejaALosDemas` [member_test.go:394](../internal/domain/member_test.go#L394) ⚠️ no dice que no tiene escenario | RN-9 | [member.go](../internal/domain/member.go) | `17736ed` → `c3b82b7` |
| `TestCambiarLaListaRecibidaNoCambiaLosIntegrantesDelProyecto` [member_test.go:516](../internal/domain/member_test.go#L516) ⚠️ no dice que no tiene escenario | RN-8 | ⚠️ ningún commit de código después del test | `5b56329` → ⚠️ sin GREEN |
| `TestConvertirElTextoEnUnRol` [member_test.go:597](../internal/domain/member_test.go#L597) ⚠️ no dice que no tiene escenario | RN-4 | [errors.go](../internal/domain/errors.go), [member.go](../internal/domain/member.go) | `1d5d160` → `0244f63` |
| `TestRechazarUnTextoQueNoEsUnRol` [member_test.go:626](../internal/domain/member_test.go#L626) ⚠️ no dice que no tiene escenario | RN-4 | [errors.go](../internal/domain/errors.go), [member.go](../internal/domain/member.go) | `1d5d160` → `0244f63` |
| `TestElErrorDeRolNombraLosValidos` [member_test.go:643](../internal/domain/member_test.go#L643) ⚠️ no dice que no tiene escenario | RN-4 | [errors.go](../internal/domain/errors.go), [member.go](../internal/domain/member.go) | `1d5d160` → `0244f63` |
| `TestLosDosRolesSeMuestranYSeVuelvenAConvertir` [member_test.go:659](../internal/domain/member_test.go#L659) ⚠️ no dice que no tiene escenario | RN-4 | [errors.go](../internal/domain/errors.go), [member.go](../internal/domain/member.go) | `1d5d160` → `0244f63` |

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
| [Crear una historia en el backlog](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L144) | `TestCrearUnaHistoriaEnElBacklog` [story_test.go:243](../internal/domain/story_test.go#L243) | [project.go](../internal/domain/project.go), [project_registry.go](../internal/domain/project_registry.go), [story.go](../internal/domain/story.go) | `81382b9` → `3e5a29c` |
| ⚠️ [Los identificadores no se repiten](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L151) | **sin test** | | |
| [Rechazar una historia sin criterios de aceptacion](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L156) | `TestRechazarUnaHistoriaSinLosDatosObligatorios` [story_test.go:334](../internal/domain/story_test.go#L334) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `d2e8d9f` → `de00fc7` |
| [Rechazar una prioridad invalida](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L162) | `TestRechazarUnaPrioridadInvalida` [story_test.go:133](../internal/domain/story_test.go#L133) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `4dc7d45` → `400a285` |
| ″ | `TestRechazarUnaHistoriaSinLosDatosObligatorios` [story_test.go:334](../internal/domain/story_test.go#L334) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `d2e8d9f` → `de00fc7` |
| [Modificar una historia pendiente](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L168) | `TestModificarUnaHistoria` [story_test.go:390](../internal/domain/story_test.go#L390) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `2bae7e0` → `13f6ad5` |
| [No modificar una historia terminada](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L173) | `TestNoModificarUnaHistoriaTerminada` [story_test.go:479](../internal/domain/story_test.go#L479) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `2bae7e0` → `13f6ad5` |
| [Backlog de un proyecto sin historias](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L179) | `TestBacklogDeUnProyectoSinHistorias` [story_test.go:227](../internal/domain/story_test.go#L227) | [project.go](../internal/domain/project.go), [project_registry.go](../internal/domain/project_registry.go), [story.go](../internal/domain/story.go) | `81382b9` → `3e5a29c` |
| [Rechazar una historia sin los datos obligatorios](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L184) | `TestRechazarUnaHistoriaSinLosDatosObligatorios` [story_test.go:334](../internal/domain/story_test.go#L334) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `d2e8d9f` → `de00fc7` |
| [Un texto con solo espacios cuenta como vacio](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L198) | `TestRechazarUnaHistoriaSinLosDatosObligatorios` [story_test.go:334](../internal/domain/story_test.go#L334) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `d2e8d9f` → `de00fc7` |
| [Descartar los criterios de aceptacion vacios](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L204) | `TestNormalizarLosDatosDeUnaHistoria` [story_test.go:303](../internal/domain/story_test.go#L303) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `d2e8d9f` → `de00fc7` |
| [Rechazar si todos los criterios quedan vacios](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L209) | `TestRechazarUnaHistoriaSinLosDatosObligatorios` [story_test.go:334](../internal/domain/story_test.go#L334) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `d2e8d9f` → `de00fc7` |
| [Guardar el titulo sin espacios en los extremos](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L215) | `TestNormalizarLosDatosDeUnaHistoria` [story_test.go:303](../internal/domain/story_test.go#L303) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `d2e8d9f` → `de00fc7` |
| [Una modificacion rechazada no cambia nada](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L220) | `TestUnaModificacionRechazadaNoCambiaNada` [story_test.go:462](../internal/domain/story_test.go#L462) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `2bae7e0` → `13f6ad5` |
| [Modificar una historia que no existe](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L226) | `TestModificarUnaHistoriaQueNoExiste` [story_test.go:499](../internal/domain/story_test.go#L499) | ⚠️ ningún commit de código después del test | `499eb03` → ⚠️ sin GREEN |
| [Reemplazar los criterios de aceptacion](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L232) | `TestModificarUnaHistoria` [story_test.go:390](../internal/domain/story_test.go#L390) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `2bae7e0` → `13f6ad5` |
| ″ | `TestReemplazarLosCriteriosDeAceptacion` [story_test.go:445](../internal/domain/story_test.go#L445) | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `2bae7e0` → `13f6ad5` |
| [Los identificadores son unicos dentro de cada proyecto](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L237) | `TestLosIdentificadoresSonUnicosDentroDeCadaProyecto` [story_test.go:272](../internal/domain/story_test.go#L272) | [project.go](../internal/domain/project.go), [project_registry.go](../internal/domain/project_registry.go), [story.go](../internal/domain/story.go) | `81382b9` → `3e5a29c` |
| [Pedir el backlog de un proyecto con historias](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L243) | `TestPedirElBacklogDeUnProyectoConHistorias` [story_test.go:532](../internal/domain/story_test.go#L532) | ⚠️ ningún commit de código después del test | `499eb03` → ⚠️ sin GREEN |
| ″ | `TestCambiarElBacklogRecibidoNoCambiaElDelProyecto` [story_test.go:564](../internal/domain/story_test.go#L564) | ⚠️ ningún commit de código después del test | `499eb03` → ⚠️ sin GREEN |
| [Operar sobre un proyecto que no existe](specs/sprint-1/06-administrar-product-backlog.md?plain=1#L250) | `TestOperarSobreUnProyectoQueNoExiste` [story_test.go:520](../internal/domain/story_test.go#L520) | ⚠️ ningún commit de código después del test | `499eb03` → ⚠️ sin GREEN |

- ⚠️ `TestModificarUnaHistoria` también cita «Modificar una historia sin cambiar nada», que no es un escenario de la spec.

**Tests que no citan un escenario de la Fase 1 (14).**

| Test | Cubre | Código | RED → GREEN |
|---|---|---|---|
| `TestUnaHistoriaExponeTodosSusDatos` [story_test.go:13](../internal/domain/story_test.go#L13) | RN-2 | [story.go](../internal/domain/story.go) | `bde9c51` → `aded90a` |
| `TestCambiarLosCriteriosRecibidosNoCambiaLosDeLaHistoria` [story_test.go:51](../internal/domain/story_test.go#L51) ⚠️ no dice que no tiene escenario | RN-2 | [story.go](../internal/domain/story.go) | `bde9c51` → `aded90a` |
| `TestLosTresEstadosDeUnaHistoriaSeDistinguen` [story_test.go:66](../internal/domain/story_test.go#L66) ⚠️ no dice que no tiene escenario | RN-2 | [story.go](../internal/domain/story.go) | `bde9c51` → `aded90a` |
| `TestConvertirElTextoEnUnaPrioridad` [story_test.go:89](../internal/domain/story_test.go#L89) ⚠️ no dice que no tiene escenario | RN-4 | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `4dc7d45` → `400a285` |
| `TestUnaPrioridadQueNoExisteNoSeMuestraComoAlta` [story_test.go:119](../internal/domain/story_test.go#L119) ⚠️ no dice que no tiene escenario | RN-4 | [story.go](../internal/domain/story.go) | `0d22bc7` ⚠️ test y código en el mismo commit |
| `TestElErrorDePrioridadNombraLasValidas` [story_test.go:150](../internal/domain/story_test.go#L150) ⚠️ no dice que no tiene escenario | RN-4 | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `4dc7d45` → `400a285` |
| `TestLasTresPrioridadesSeMuestranYSeVuelvenAConvertir` [story_test.go:163](../internal/domain/story_test.go#L163) ⚠️ no dice que no tiene escenario | RN-4 | ⚠️ ningún commit de código después del test | `1d7326d` → ⚠️ sin GREEN |
| `TestUnaHistoriaDevuelveSusStoryPoints` [story_test.go:185](../internal/domain/story_test.go#L185) ⚠️ no dice que no tiene escenario | RN-2 | ⚠️ ningún commit de código después del test | `1d7326d` → ⚠️ sin GREEN |
| `TestLosIdentificadoresDeLasHistoriasNoSeRepiten` [story_test.go:289](../internal/domain/story_test.go#L289) | RN-1 | [project.go](../internal/domain/project.go), [project_registry.go](../internal/domain/project_registry.go), [story.go](../internal/domain/story.go) | `81382b9` → `3e5a29c` |
| `TestDosHistoriasDelMismoProyectoPuedenTenerElMismoTitulo` [story_test.go:369](../internal/domain/story_test.go#L369) | RN-14 | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `d2e8d9f` → `de00fc7` |
| `TestModificarUnaHistoriaEnCurso` [story_test.go:411](../internal/domain/story_test.go#L411) | RN-8 | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `2bae7e0` → `13f6ad5` |
| `TestModificarUnaHistoriaSinCambiarNada` [story_test.go:429](../internal/domain/story_test.go#L429) | RN-9 | [errors.go](../internal/domain/errors.go), [story.go](../internal/domain/story.go) | `2bae7e0` → `13f6ad5` |
| `TestBuscarUnaHistoriaDelBacklog` [story_test.go:579](../internal/domain/story_test.go#L579) | RN-13 | ⚠️ ningún commit de código después del test | `499eb03` → ⚠️ sin GREEN |
| `TestModificarLaDescripcionDeUnaHistoria` [story_test.go:596](../internal/domain/story_test.go#L596) | RN-9 | ⚠️ ningún commit de código después del test | `e946fda` → ⚠️ sin GREEN |

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
| ⚠️ [Crear un proyecto desde la web](specs/sprint-2/77-interfaz-web-base.md?plain=1#L156) | **sin test** | | |
| ⚠️ [Cada operacion se confirma con un mensaje](specs/sprint-2/77-interfaz-web-base.md?plain=1#L163) | **sin test** | | |
| ⚠️ [Consultar un proyecto](specs/sprint-2/77-interfaz-web-base.md?plain=1#L178) | **sin test** | | |
| ⚠️ [Las opciones para elegir son las del sistema](specs/sprint-2/77-interfaz-web-base.md?plain=1#L184) | **sin test** | | |
| ⚠️ [Confirmar o cancelar antes de quitar](specs/sprint-2/77-interfaz-web-base.md?plain=1#L194) | **sin test** | | |
| ⚠️ [Mostrar el rechazo sin perder lo escrito](specs/sprint-2/77-interfaz-web-base.md?plain=1#L207) | **sin test** | | |
| [Las listas vacias dicen que hacer](specs/sprint-2/77-interfaz-web-base.md?plain=1#L225) | `TestSinDatosGuardadosSeArrancaSinProyectos` [file_test.go:53](../internal/store/file_test.go#L53) | [file.go](../internal/store/file.go) | `9e9ec91` → `fc414fb` |
| ⚠️ [Pedir algo que no existe](specs/sprint-2/77-interfaz-web-base.md?plain=1#L237) | **sin test** | | |
| ⚠️ [Un texto con signos se muestra tal cual](specs/sprint-2/77-interfaz-web-base.md?plain=1#L249) | **sin test** | | |
| ⚠️ [Abrir el sistema con un solo comando y sin internet](specs/sprint-2/77-interfaz-web-base.md?plain=1#L253) | **sin test** | | |
| [Los datos siguen al cerrar y volver a abrir](specs/sprint-2/77-interfaz-web-base.md?plain=1#L259) | `TestLosDatosSiguenAlCerrarYVolverAAbrir` [registry_json_test.go:116](../internal/domain/registry_json_test.go#L116) | [errors.go](../internal/domain/errors.go), [registry_json.go](../internal/domain/registry_json.go) | `1a885f0` → `61077ba` |
| ″ | `TestLoGuardadoSeVuelveALeerIgual` [file_test.go:68](../internal/store/file_test.go#L68) | [file.go](../internal/store/file.go) | `9e9ec91` → `fc414fb` |
| [Los datos guardados estan danados](specs/sprint-2/77-interfaz-web-base.md?plain=1#L265) | `TestLosDatosGuardadosEstanDanados` [registry_json_test.go:172](../internal/domain/registry_json_test.go#L172) | [registry_json.go](../internal/domain/registry_json.go) | `ff7d24f` → `2ecdfad` |
| ″ | `TestUnArchivoDanadoQuedaComoEstaba` [file_test.go:124](../internal/store/file_test.go#L124) | [file.go](../internal/store/file.go) | `6d73f7a` → `a83bfc5` |
| [Un cambio que no se pudo guardar se avisa](specs/sprint-2/77-interfaz-web-base.md?plain=1#L271) | `TestUnCambioQueNoSePudoGuardarSeAvisa` [file_test.go:165](../internal/store/file_test.go#L165) | [file.go](../internal/store/file.go) | `6d73f7a` → `a83bfc5` |
| ⚠️ [Dos pestanas a la vez](specs/sprint-2/77-interfaz-web-base.md?plain=1#L276) | **sin test** | | |
| ⚠️ [Usar la pagina en un celular y con el teclado](specs/sprint-2/77-interfaz-web-base.md?plain=1#L281) | **sin test** | | |

**Tests que no citan un escenario de la Fase 1 (6).**

| Test | Cubre | Código | RED → GREEN |
|---|---|---|---|
| `TestListarLosProyectosDeUnRegistroVacio` [project_registry_test.go:7](../internal/domain/project_registry_test.go#L7) | dependencia de la #77, RN-7 y RN-23 (la lista vacía no es un error) | [project_registry.go](../internal/domain/project_registry.go) | `1f94f7a` → `3e81aca` |
| `TestListarTodosLosProyectosEnOrdenDeNumero` [project_registry_test.go:22](../internal/domain/project_registry_test.go#L22) | dependencia de la #77 | [project_registry.go](../internal/domain/project_registry.go) | `1f94f7a` → `3e81aca` |
| `TestElMensajeDeDatosDanadosEstaEnCastellano` [registry_json_test.go:236](../internal/domain/registry_json_test.go#L236) | #77 RN-24 | [registry_json.go](../internal/domain/registry_json.go), [file.go](../internal/store/file.go) | `ff2b683` → `d0f0a49` |
| `TestLeerLaEscalaDeStoryPoints` [story_estimation_test.go:253](../internal/domain/story_estimation_test.go#L253) | #77 RN-11 | [story_estimation.go](../internal/domain/story_estimation.go) | `4227227` → `5a6d632` |
| `TestElMensajeDeUnArchivoQueNoSePuedeLeerEstaEnCastellano` [file_test.go:217](../internal/store/file_test.go#L217) | #77 RN-24 | [registry_json.go](../internal/domain/registry_json.go), [file.go](../internal/store/file.go) | `ff2b683` → `d0f0a49` |
| `TestElMensajeDeUnCambioQueNoSePudoGuardarEstaEnCastellano` [file_test.go:237](../internal/store/file_test.go#L237) | #77 RN-25 | [file.go](../internal/store/file.go) | `6f761d9` → `3854359` |

## Tests de ninguna historia

Ninguna tabla de trazabilidad lista su archivo, y el test no cita un escenario ni nombra
una historia en `Cubre:`.

| Test | Cubre | Código | RED → GREEN |
|---|---|---|---|
| `TestLosMensajesDeErrorLlevanTildes` [errors_test.go:9](../internal/domain/errors_test.go#L9) ⚠️ no dice que no tiene escenario | ⚠️ sin `Cubre:` | [errors.go](../internal/domain/errors.go) | `73d0160` → `d2ff81b` |

