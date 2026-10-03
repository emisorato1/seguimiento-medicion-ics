# Spec: Operar el sistema desde una interfaz de consola

- **Historia:** #24
- **Sprint:** 1
- **Responsable:** Maximiliano Eula — asignada por el Agile Enabler el 02/10 por carga
  de trabajo (era quien tenía menos story points)
- **Estado:** Requisitos — **borrador, pendiente de Puerta 1**
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
- Sprints, defectos, métricas, dashboard y reportes: no existen todavía.
- Guardar los datos entre ejecuciones: al cerrar el programa se pierde todo.
- Colores, menús gráficos o cualquier adorno.

### Entradas

Comandos tipeados por el usuario, con sus datos. Las operaciones disponibles son las cuatro
historias del sprint: crear proyecto, agregar integrante, cargar historia y estimar historia,
más las consultas para listar lo cargado.

### Salidas

El resultado de cada operación en pantalla: lo que se creó con su identificador, la lista
pedida, o el mensaje de error si la operación se rechazó.

### Reglas de negocio

1. **[RN-1]** El sistema debe ofrecer una operación por cada funcionalidad del Sprint 1:
   crear proyecto, agregar integrante, listar integrantes, cargar historia, listar el
   backlog y estimar una historia.
2. **[RN-2]** Cuando el usuario ejecuta una operación que el núcleo rechaza, el sistema debe
   mostrar el motivo del rechazo tal como lo informa el núcleo, y seguir funcionando.
3. **[RN-3]** El sistema debe mostrar el identificador de todo lo que se crea, porque es lo
   que hace falta para las operaciones siguientes.
4. **[RN-4]** Si el usuario escribe una operación que no existe, el sistema debe informarlo
   y mostrar las operaciones disponibles.
5. **[RN-5]** Si el usuario ejecuta una operación sin los datos obligatorios, el sistema debe
   indicar qué dato falta y no llamar al núcleo.
6. **[RN-6]** El sistema debe permitir salir en cualquier momento.
7. **[RN-7]** La consola no debe validar reglas de negocio por su cuenta: las validaciones
   las hace el núcleo y la consola solo muestra el resultado.

### Casos límite

- Listar los integrantes o el backlog cuando no hay ninguno: muestra que está vacío, no es
  un error.
- Operar sobre un identificador que no existe: muestra el error del núcleo y sigue.
- Entrada vacía: no hace nada y vuelve a pedir.

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| Operación inexistente | informa y lista las operaciones disponibles |
| Falta un dato obligatorio | indica cuál falta y no llama al núcleo |
| El núcleo rechaza la operación | muestra el motivo y sigue funcionando |

### Criterios de aceptación

- [ ] Se puede crear un proyecto desde la consola y se ve su identificador.
- [ ] Se pueden agregar integrantes a un proyecto y listarlos.
- [ ] Se pueden cargar historias en el backlog y listarlas.
- [ ] Se le pueden asignar story points a una historia.
- [ ] Cuando una operación se rechaza, se ve el motivo y el programa sigue andando.
- [ ] Listar algo vacío no es un error.
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

Escenario: Listar un backlog vacio
  Dado un proyecto recien creado
  Cuando pido el backlog
  Entonces veo que no hay historias cargadas

Escenario: Operacion inexistente
  Dado que el programa esta abierto
  Cuando escribo una operacion que no existe
  Entonces veo la lista de operaciones disponibles

Escenario: Estimar una historia desde la consola
  Dado un proyecto con una historia cargada
  Cuando le asigno 5 story points
  Entonces veo que la historia quedo estimada en 5
```

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios:

## Fase 2 — Diseño

_Se escribe después de que la Puerta 1 esté aprobada._

## Fase 3 — Tareas

_Se escribe después de que la Puerta 2 esté aprobada._

## Fase 4 — Implementación

_Se completa a medida que se avanza._

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #24 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | pendiente |
| Código Go | pendiente |
