# Spec: Operar el sistema desde una interfaz de consola

- **Historia:** #24
- **Sprint:** 1
- **Responsable:** Maximiliano Eula — asignada por el Agile Enabler el 02/10 por carga
  de trabajo (era quien tenía menos story points)
- **Estado:** Puerta 3 aprobada — **lista para implementar**
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de la consigna y
  de los criterios de aceptación de la issue. El responsable lo revisa, lo corrige si hace
  falta y lo presenta a la Puerta 1
- **Depende de:** #3, #4, #6 y #12 — las cuatro tienen que estar en `dev`
- **La necesitan:** nadie todavía. En el Sprint 2 se reemplaza o se amplía con la interfaz
  definitiva

## Fase 1 — Requisitos

### Objetivo

Que el equipo pueda ejecutar a mano las funcionalidades del Sprint 1 sin escribir tests,
para poder probarlas y mostrarlas en la review del sprint.

### Qué queda afuera

- **Cualquier regla de negocio.** La consola solo pide datos, llama al núcleo y muestra el
  resultado. Si una validación aparece en la consola, está en el lugar equivocado.
- **Modificar un proyecto, quitar un integrante, modificar una historia y quitar la
  estimación.** Están en las historias #3, #4, #6 y #12, pero la issue no las pide para la
  consola y no entran en su estimación. Se pueden probar con los tests, y se suman a la
  interfaz en el Sprint 2.
- Sprints, defectos, métricas, dashboard y reportes: no existen todavía.
- Guardar los datos entre ejecuciones: al cerrar el programa se pierde todo.
- Colores, menús gráficos o cualquier adorno.

### Entradas

Comandos tipeados por el usuario, con sus datos. Las operaciones disponibles son las de la
RN-1.

### Salidas

El resultado de cada operación en pantalla: lo que se creó con su identificador, la lista
pedida, o el mensaje de error si la operación se rechazó.

### Reglas de negocio

1. **[RN-1]** El sistema debe ofrecer estas operaciones: crear proyecto, agregar integrante,
   listar integrantes, cargar historia, listar el backlog, estimar una historia y salir.
2. **[RN-2]** Cuando el usuario ejecuta una operación que el núcleo rechaza, el sistema debe
   mostrar el motivo del rechazo tal como lo informa el núcleo, y seguir funcionando.
3. **[RN-3]** El sistema debe mostrar el identificador de todo lo que se crea, porque es lo
   que hace falta para las operaciones siguientes.
4. **[RN-4]** Si el usuario escribe una operación que no existe, el sistema debe informarlo
   y mostrar las operaciones disponibles.
5. **[RN-5]** Si a una operación le falta un dato sin el cual no se puede ejecutar —por
   ejemplo, a qué proyecto o a qué historia se aplica—, o un dato no tiene la forma esperada
   —por ejemplo, letras donde va un número—, el sistema debe indicar qué dato falta o está
   mal escrito y no ejecutar la operación.
6. **[RN-6]** El sistema debe permitir salir en cualquier momento.
7. **[RN-7]** La consola no debe validar reglas de negocio por su cuenta. Si un dato está y
   tiene la forma esperada pero no cumple una regla —un nombre vacío, un valor de story
   points fuera de la escala—, lo rechaza el núcleo y la consola solo muestra el motivo
   (RN-2).

La diferencia entre la RN-5 y la RN-7 es si la operación **se puede ejecutar**: si falta a
qué historia se aplica, no hay operación posible y avisa la consola; si el valor está pero
no es válido, la operación se ejecuta y la rechaza el núcleo.

### Casos límite

- Listar los integrantes o el backlog cuando no hay ninguno: muestra que está vacío, no es
  un error.
- Operar sobre un identificador que no existe: muestra el error del núcleo y sigue.
- Entrada vacía: no hace nada y vuelve a pedir.
- Letras donde va un número, como un identificador o los story points: lo informa la
  consola, sin ejecutar la operación (RN-5).
- Un número válido que no cumple una regla, como 4 story points: lo rechaza el núcleo
  (RN-7).

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| Operación inexistente | informa y lista las operaciones disponibles |
| Falta un dato sin el cual la operación no se puede ejecutar | indica cuál falta y no la ejecuta |
| Un dato no tiene la forma esperada | indica cuál está mal escrito y no ejecuta la operación |
| El núcleo rechaza la operación | muestra el motivo y sigue funcionando |

### Criterios de aceptación

- [ ] Se puede crear un proyecto desde la consola y se ve su identificador.
- [ ] Se pueden agregar integrantes a un proyecto y listarlos.
- [ ] Se pueden cargar historias en el backlog y listarlas.
- [ ] Se le pueden asignar story points a una historia.
- [ ] Cuando una operación se rechaza, se ve el motivo y el programa sigue andando.
- [ ] Listar algo vacío no es un error.
- [ ] Una operación incompleta o con un dato mal escrito se informa sin ejecutarla.
- [ ] Se puede salir del programa.
- [ ] La consola no contiene ninguna validación de negocio propia.

### Escenarios BDD

```gherkin
Escenario: Crear un proyecto desde la consola
  Dado que el programa esta abierto
  Cuando ejecuto la operacion de crear proyecto con nombre "Sistema de metricas" e inicio 01/10/2026
  Entonces veo que el proyecto se creo
  Y veo su identificador

Escenario: Mostrar el motivo de un rechazo
  Dado que el programa esta abierto
  Cuando intento crear un proyecto sin nombre
  Entonces veo el mensaje que informa que el nombre es obligatorio
  Y el programa sigue funcionando

Escenario: Agregar y listar integrantes
  Dado un proyecto creado desde la consola
  Cuando le agrego un integrante
  Y pido la lista de integrantes del proyecto
  Entonces veo al integrante en la lista

Escenario: Listar un backlog vacio
  Dado un proyecto recien creado
  Cuando pido el backlog
  Entonces veo que no hay historias cargadas

Escenario: Cargar una historia y listar el backlog
  Dado un proyecto creado desde la consola
  Cuando cargo una historia con titulo "Registrar esfuerzo"
  Y pido el backlog
  Entonces veo la historia "Registrar esfuerzo" con su identificador

Escenario: Estimar una historia desde la consola
  Dado un proyecto con una historia cargada
  Cuando le asigno 5 story points
  Entonces veo que la historia quedo estimada en 5

Escenario: El nucleo rechaza un valor fuera de la escala
  Dado un proyecto con una historia cargada
  Cuando intento asignarle 4 story points
  Entonces veo el mensaje del nucleo con los valores validos
  Y el programa sigue funcionando

Escenario: Operacion incompleta
  Dado que el programa esta abierto
  Cuando ejecuto la operacion de estimar sin indicar la historia
  Entonces veo que falta indicar la historia
  Y la operacion no se ejecuta

Escenario: Dato con forma invalida
  Dado un proyecto con una historia cargada
  Cuando intento asignarle "cinco" story points
  Entonces veo que los story points tienen que ser un numero
  Y la operacion no se ejecuta

Escenario: Operar sobre un identificador que no existe
  Dado que el programa esta abierto
  Cuando intento pedir el backlog del proyecto 99, que no existe
  Entonces veo el mensaje del nucleo que informa que no se encontro el proyecto
  Y el programa sigue funcionando

Escenario: Operacion inexistente
  Dado que el programa esta abierto
  Cuando escribo una operacion que no existe
  Entonces veo la lista de operaciones disponibles

Escenario: Entrada vacia
  Dado que el programa esta abierto
  Cuando envio una linea vacia
  Entonces no pasa nada
  Y el programa vuelve a pedir una operacion

Escenario: Salir del programa
  Dado que el programa esta abierto
  Cuando ejecuto la operacion de salir
  Entonces el programa termina
```

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: Emiliano Sorato — Fecha: 2026-10-03 — Comentarios: aprobada en el PR #38. Destacó
> la separación entre la RN-5 (errores de forma, los informa la consola) y la RN-7 (reglas de
> negocio, las rechaza el núcleo), y que quede escrito que modificar y quitar la estimación se
> prueban con los tests y se suman a la interfaz en el Sprint 2.

## Fase 2 — Diseño

Las decisiones transversales del proyecto están en [AGENTS.md](../../../AGENTS.md) y no se
repiten acá. Las dos que más pesan en esta historia: el núcleo **no sabe nada de la
interfaz**, y las validaciones viven **en el dominio**, nunca en la consola.

### Enfoque

La consola es una capa delgada: **pide datos, llama al núcleo y muestra lo que responde**.
No tiene reglas de negocio ni guarda datos propios.

Cada operación sigue el mismo recorrido:

1. El usuario escribe el nombre de la operación.
2. La consola pide cada dato en su propia línea.
3. Si falta un dato sin el cual la operación no se puede ejecutar, o un dato no tiene la
   forma esperada, la consola lo informa y no llama al núcleo (RN-5).
4. Si no, llama al núcleo y muestra el resultado o el motivo del rechazo, tal cual (RN-2).

La pregunta que separa el paso 3 del 4 es **si la operación se puede ejecutar**: si falta a
qué historia se aplica, no hay operación posible y avisa la consola; si el valor está pero no
es válido —un nombre vacío, 4 story points—, la operación se ejecuta y la rechaza el núcleo
(RN-7).

La consola trabaja sobre una entrada y una salida genéricas. El arranque del programa le
conecta el teclado y la pantalla; los tests le pasan líneas de texto y leen lo que escribió.

### Archivos afectados

| Archivo | Se crea o se modifica | Para qué |
|---|---|---|
| `cmd/app/main.go` | se modifica | arranca la consola conectada al teclado y la pantalla |
| `cmd/app/console.go` | se crea | las operaciones: pedir los datos, llamar al núcleo y mostrar el resultado |
| `cmd/app/console_test.go` | se crea | los tests de los escenarios BDD de la Fase 1 |

### Decisiones

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| El código vive en `cmd/app`, repartido en arranque y operaciones | un paquete nuevo `internal/console` | respeta la estructura actual y la consola sigue siendo reemplazable desde `cmd/` |
| La consola trabaja sobre una entrada y una salida genéricas | leer y escribir directo del teclado y la pantalla | es lo que permite que cada escenario BDD sea un test automatizado |
| Se escribe el nombre de la operación y la consola pide cada dato en su propia línea | toda la operación en una sola línea | los nombres con espacios se leen sin comillas ni reglas de separación |
| La consola no guarda datos: usa el registro de proyectos que define la #3 | que la consola arme su propia lista | guardar datos es responsabilidad del núcleo y no se pierde al cambiar de interfaz |
| Los errores se muestran con el mensaje del núcleo, tal cual | textos propios de la consola para cada error | lo pide la RN-2, y evita duplicar los mensajes que viven en el dominio |
| Una historia con 0 story points se muestra como "sin estimar" | mostrar el número | la #12 representa "sin estimar" con 0 y deja en quien muestra el dato cumplir su RN-4 |

### Dependencias y restricciones

- **Depende de la #3, la #4, la #6 y la #12**: el código no se empieza hasta que las cuatro
  estén en `dev`.
- **Depende de que la Fase 2 de la #3 defina dónde se guardan los proyectos** mientras corre
  el programa. Hoy no está definido en ninguna spec ni en AGENTS.md.
- **No se modifica nada de `internal/domain`.** Si a la consola le falta algo del núcleo, se
  pide en la historia que corresponde, no se agrega desde acá.
- **Ninguna validación de negocio en la consola** (RN-7): solo se controla la forma de los
  datos sin los cuales la operación no se puede ejecutar (RN-5).

> **Puerta 2** — ¿el diseño es viable y coherente con el resto del proyecto?
> Aprobó: Emiliano Sorato — Fecha: 2026-10-05 — Comentarios: aprobada en el PR #40. La entrada y salida genericas son lo que convierte cada escenario BDD de la consola en un test automatizado.

## Fase 3 — Tareas

Cada tarea es un ciclo de TDD completo: el test que falla en un commit, la implementación que
lo hace pasar en otro. Están en el orden en que hay que hacerlas, y cada una deja el código
compilando y los tests en verde. La única excepción es la 15, que se explica abajo.

Los tests siguen la convención de [AGENTS.md](../../../AGENTS.md): un test por escenario, con
el nombre del escenario. Cada test le pasa a la consola líneas de texto como entrada y
verifica lo que escribió como salida.

| # | Tarea | Escenario BDD | Reglas | Necesita en `dev` | Estado |
|---|---|---|---|---|---|
| 1 | La consola pide una operación sobre una entrada y una salida genéricas, y termina con la operación de salir | Salir del programa | RN-6 | nada | hecha |
| 2 | Una línea vacía no hace nada y vuelve a pedir una operación | Entrada vacia | caso límite | nada | hecha |
| 3 | Una operación que no existe se informa y se listan las disponibles | Operacion inexistente | RN-4 | nada | hecha |
| 4 | Crear un proyecto: pide los datos, lo crea en el registro y muestra su identificador | Crear un proyecto desde la consola | RN-1, RN-3 | #3 | pendiente |
| 5 | Si el núcleo rechaza la operación, se muestra su mensaje tal cual y la consola sigue | Mostrar el motivo de un rechazo | RN-2, RN-7 | #3 | pendiente |
| 6 | Agregar un integrante y listar los del proyecto; el rol se pasa al dominio como texto | Agregar y listar integrantes | RN-1 | #4 | pendiente |
| 7 | Listar el backlog de un proyecto sin historias muestra que está vacío | Listar un backlog vacio | RN-1 | #6 | pendiente |
| 8 | Cargar una historia y listar el backlog con su identificador; una historia con 0 story points se muestra "sin estimar" | Cargar una historia y listar el backlog | RN-1, RN-3 | #6, #12 | pendiente |
| 9 | Estimar una historia y ver su nuevo valor | Estimar una historia desde la consola | RN-1 | #12 | pendiente |
| 10 | Si el núcleo rechaza un valor fuera de la escala, se muestra su mensaje y la consola sigue | El nucleo rechaza un valor fuera de la escala | RN-2, RN-7 | #12 | pendiente |
| 11 | Si falta indicar a qué historia se aplica, se informa y la operación no se ejecuta | Operacion incompleta | RN-5 | #12 | pendiente |
| 12 | Letras donde va un número se informan y la operación no se ejecuta | Dato con forma invalida | RN-5 | #12 | pendiente |
| 13 | Operar sobre un identificador que no existe muestra el mensaje del núcleo y la consola sigue | Operar sobre un identificador que no existe | RN-2 | #6 | pendiente |
| 14 | Una fecha mal escrita se informa y la operación no se ejecuta | — (RN-5, sin escenario) | RN-5 | #3 | pendiente |
| 15 | El programa arranca la consola conectada al teclado y la pantalla, con un registro de proyectos nuevo | — (arranque) | — | todas | pendiente |

### Decisiones de esta fase

Bajar el diseño a tareas dejó cuatro puntos que la Fase 2 no resolvía:

- **Los criterios de aceptación de una historia se ingresan uno por línea**, y una línea vacía
  termina la lista. Es el mismo criterio de pedir cada dato en su propia línea, y no choca con
  ningún carácter que pueda aparecer en un criterio.
- **Las fechas se escriben `dd/mm/aaaa`**, que es el formato de todos los escenarios. La Fase 2
  de la #3 deja en la consola el formato de las fechas, así que una fecha mal escrita es un
  error de forma (RN-5) y la informa la consola: es la tarea 14.
- **La prioridad y el rol se pasan al dominio como texto.** La #4 ya adoptó que el dominio
  convierta el texto en un rol y rechace con su RN-4. Para la prioridad, la tarea 8 supone lo
  mismo, que está propuesto en la revisión de la Fase 3 de la #6 (PR #44). Si la #6 no lo
  adopta, la tarea 8 se ajusta antes de la Puerta 3.
- **La tarea 15 no tiene test unitario.** Solo conecta el teclado, la pantalla y el registro, y
  todo lo que hace la consola ya está probado por las tareas 1 a 14 sobre la entrada y la
  salida genéricas. Se verifica corriendo el programa con `go run ./cmd/app`.

### Por qué este orden

- **Las tareas 1 a 3 no dependen de ninguna otra historia**: son el ciclo de pedir una
  operación, la entrada vacía y la operación inexistente. Se pueden programar apenas pase la
  Puerta 3, mientras las demás historias llegan a `dev`.
- **De la 4 a la 13**, cada operación entra cuando su historia está en `dev`, en el orden de
  dependencias del sprint: #3, #4, #6 y #12. En cada una, primero el camino feliz y después
  los rechazos.
- **La 14 va después de la 4**, porque necesita la operación de crear proyecto para tener una
  fecha que escribir mal.
- **La 15 va al final**, cuando todas las operaciones existen.

### Cobertura

Las siete reglas de negocio quedan cubiertas: RN-1 (4, 6, 7, 8, 9), RN-2 (5, 10, 13),
RN-3 (4, 8), RN-4 (3), RN-5 (11, 12, 14), RN-6 (1), RN-7 (5, 10).

| Criterio de aceptación | Tareas |
|---|---|
| Se puede crear un proyecto desde la consola y se ve su identificador | 4 |
| Se pueden agregar integrantes a un proyecto y listarlos | 6 |
| Se pueden cargar historias en el backlog y listarlas | 8 |
| Se le pueden asignar story points a una historia | 9 |
| Cuando una operación se rechaza, se ve el motivo y el programa sigue andando | 5, 10, 13 |
| Listar algo vacío no es un error | 7 |
| Una operación incompleta o con un dato mal escrito se informa sin ejecutarla | 11, 12, 14 |
| Se puede salir del programa | 1 |
| La consola no contiene ninguna validación de negocio propia | 5, 10 |

Los trece escenarios BDD de la Fase 1 tienen su tarea, en la columna "Escenario BDD".

> **Puerta 3** — ¿las tareas cubren todos los criterios de aceptación?
> Aprobó: Emiliano Sorato — Fecha: 2026-10-08 — Comentarios: aprobada en el PR #49. La columna 'Necesita en dev' de la tabla de tareas indica que historia tiene que estar mergeada antes de cada tarea: las tareas 1 a 3 no dependen de ninguna.

## Fase 4 — Implementación

Ciclo RED → GREEN → REFACTOR. El test en un commit y la implementación en otro.

| Tarea | Commit del test | Commit de la implementación |
|---|---|---|
| 1 | `7d31b89` | `cd3c73b` |
| 2 | `44eaaf7` | `d9e27a8` |
| 3 | `35173b0` | `69fbd58` |

### Desvíos respecto de la spec

Ninguno por ahora.

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #24 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | `cmd/app/console_test.go` |
| Código Go | `cmd/app/console.go` |
