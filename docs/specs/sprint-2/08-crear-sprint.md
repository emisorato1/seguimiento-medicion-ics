# Spec: Crear un sprint con su Sprint Goal

- **Historia:** #8
- **Sprint:** 2
- **Responsable:** a definir en la planning
- **Estado:** Requisitos — **borrador, pendiente de Puerta 1**
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de
  la consigna y de los criterios de aceptación de la issue. El responsable lo
  revisa, lo corrige si hace falta y lo presenta a la Puerta 1.
- **Depende de:** #3 (el sprint pertenece a un proyecto)
- **La necesitan:** #9 (asignar historias al sprint), #10 (cerrar el sprint),
  #11 (consultar sprints anteriores), #16 (métricas)

## Fase 1 — Requisitos

### Objetivo

Que el equipo pueda abrir un sprint en un proyecto, con una meta clara y un principio y un
fin, para organizar el trabajo en períodos cortos. Es la base de todo lo que viene después:
sin un sprint no hay dónde asignar historias, qué cerrar ni qué medir.

### Qué queda afuera

- Asignar historias al sprint y armar el Sprint Backlog, que es la historia #9.
- Cerrar el sprint y registrar qué se completó, que es la #10.
- Consultar los sprints ya cerrados y su detalle, que es la #11.
- Las métricas del sprint (story points planificados, completados, velocidad), que son la #16.
- Modificar el Sprint Goal o las fechas de un sprint ya creado: en esta historia el sprint
  se crea y no se edita.
- Borrar un sprint.
- Comparar las fechas del sprint con las del proyecto o con la fecha de hoy (ver RN-14 y
  el punto 2 de "Para decidir en la puerta").

### Entradas

| Dato | Obligatorio | Descripción |
|---|---|---|
| Proyecto | sí | a qué proyecto pertenece el sprint |
| Sprint Goal | sí | la meta del sprint, en una frase |
| Fecha de inicio | sí | cuándo arranca |
| Fecha de fin | sí | cuándo termina |

El **número del sprint** no es una entrada: lo asigna el sistema al crearlo (RN-1). La
issue pide que el sprint tenga número, no que lo elija quien lo crea.

### Salidas

Un sprint registrado dentro del proyecto, con su número, su Sprint Goal, sus dos fechas,
en estado Abierto y sin historias asignadas. Se puede consultar cuál es el sprint abierto
del proyecto.

### Reglas de negocio

1. **[RN-1]** Cuando se crea un sprint, el sistema debe asignarle el número que sigue al del
   último sprint creado en ese proyecto, empezando por 1. El número no se repite dentro del
   proyecto y no cambia.
2. **[RN-2]** El sistema debe registrar para cada sprint: número, Sprint Goal, fecha de
   inicio, fecha de fin y estado.
3. **[RN-3]** Cuando se crea un sprint, el sistema debe dejarlo en estado **Abierto** y sin
   historias asignadas.
4. **[RN-4]** Si se intenta crear un sprint sin Sprint Goal, el sistema debe rechazar la
   operación e informar que el Sprint Goal es obligatorio. Un texto compuesto solo por
   espacios cuenta como vacío.
5. **[RN-5]** Si se intenta crear un sprint sin fecha de inicio o sin fecha de fin, el
   sistema debe rechazar la operación e informar cuál de las dos falta.
6. **[RN-6]** Si la fecha de fin es anterior a la fecha de inicio, el sistema debe rechazar
   la operación e informar que las fechas son incoherentes.
7. **[RN-7]** Cuando se registra el Sprint Goal, el sistema debe guardarlo sin los espacios
   del principio y del final.
8. **[RN-8]** Mientras un proyecto tenga un sprint abierto, el sistema debe rechazar la
   creación de otro sprint en ese proyecto e informar cuál es el sprint que ya está abierto.
9. **[RN-9]** El sistema debe tratar cada proyecto por separado: los números de sprint son
   únicos dentro de cada proyecto, y un sprint abierto en un proyecto no impide crear un
   sprint en otro.
10. **[RN-10]** El sistema debe permitir consultar cuál es el sprint abierto de un proyecto.
    Si el proyecto no tiene ninguno abierto, debe informarlo, y eso no es un error.
11. **[RN-11]** El sistema debe permitir que dos sprints del mismo proyecto tengan el mismo
    Sprint Goal y las mismas fechas: lo que los distingue es el número.
12. **[RN-12]** Si una creación se rechaza, el sistema debe dejar el proyecto y sus sprints
    como estaban antes.
13. **[RN-13]** Si se intenta crear un sprint o consultar el sprint abierto de un proyecto
    que no existe, el sistema debe informar que no se encontró el proyecto y no crear nada.
14. **[RN-14]** El sistema debe aceptar las fechas del sprint tal como se indiquen, sin
    compararlas con la fecha de hoy ni con las fechas del proyecto. La única comparación es
    la de la RN-6, entre las dos fechas del propio sprint.

### Estados de un sprint

| Estado | Qué significa |
|---|---|
| Abierto | es el sprint en curso del proyecto: se le asignan historias y se trabaja en él |
| Cerrado | ya terminó: no se le asignan historias ni se modifica |

En esta historia el sprint **solo se crea en Abierto**. El pase a Cerrado lo hace la
historia #10.

### Casos límite

- Primer sprint de un proyecto: recibe el número 1.
- Proyecto sin ningún sprint: no tiene sprint abierto, y eso no es un error.
- Proyecto con todos sus sprints cerrados: tampoco tiene sprint abierto, y se puede crear
  uno nuevo.
- Fecha de fin **igual** a la de inicio: es válida, un sprint puede durar un día. Las fechas
  se comparan por día: si llegan con hora, la hora no cuenta.
- Fechas en el pasado: son válidas, el sistema no las compara con la fecha de hoy.
- Sprint Goal con espacios al principio y al final: se guarda sin esos espacios.
- Sprint Goal de un solo carácter: es válido.
- Dos proyectos, cada uno con su primer sprint: los dos llevan el número 1, y el sprint
  abierto de uno no impide crear el del otro.
- Fechas que se superponen con las de un sprint ya cerrado del mismo proyecto: se aceptan.

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| Sprint Goal vacío o solo espacios | rechaza e informa que el Sprint Goal es obligatorio |
| Falta la fecha de inicio | rechaza e informa que la fecha de inicio es obligatoria |
| Falta la fecha de fin | rechaza e informa que la fecha de fin es obligatoria |
| Fecha de fin anterior a la de inicio | rechaza e informa que las fechas son incoherentes |
| El proyecto ya tiene un sprint abierto | rechaza e informa cuál es el sprint que está abierto |
| Proyecto inexistente al crear o al consultar | informa que no se encontró el proyecto y no crea nada |

### Criterios de aceptación

- [ ] Se puede crear un sprint con Sprint Goal, fecha de inicio y fecha de fin, y queda
      registrado con su número y en estado Abierto.
- [ ] El primer sprint de un proyecto recibe el número 1, y cada sprint nuevo el número
      siguiente.
- [ ] El número no se repite dentro del mismo proyecto.
- [ ] Un sprint recién creado no tiene historias asignadas.
- [ ] No se puede crear un sprint sin Sprint Goal, y un Sprint Goal de solo espacios cuenta
      como vacío.
- [ ] No se puede crear un sprint sin fecha de inicio ni sin fecha de fin.
- [ ] No se puede crear un sprint con fecha de fin anterior a la de inicio.
- [ ] Un sprint cuya fecha de fin es igual a la de inicio se acepta.
- [ ] No se puede crear un segundo sprint mientras haya uno abierto en el mismo proyecto, y
      el mensaje dice cuál es el que está abierto.
- [ ] Tener un sprint abierto en un proyecto no impide crear un sprint en otro proyecto.
- [ ] Se puede consultar cuál es el sprint abierto de un proyecto; si no hay ninguno se
      informa y no es un error.
- [ ] El Sprint Goal se guarda sin los espacios del principio y del final.
- [ ] Una creación rechazada no agrega ningún sprint al proyecto.
- [ ] Crear un sprint o pedir el sprint abierto de un proyecto inexistente informa que no se
      encontró el proyecto.

### Escenarios BDD

```gherkin
Escenario: Crear el primer sprint de un proyecto
  Dado un proyecto registrado sin sprints
  Cuando creo un sprint con el goal "Tener el backlog estimado", inicio 01/11/2026 y fin 14/11/2026
  Entonces el sprint queda registrado con numero 1
  Y queda en estado "Abierto"
  Y no tiene historias asignadas

Escenario: El sprint nuevo recibe el numero siguiente
  Dado un proyecto cuyo unico sprint es el numero 1 y esta cerrado
  Cuando creo un sprint nuevo
  Entonces el sprint queda registrado con numero 2
  Y el proyecto queda con 2 sprints

Escenario: No se puede abrir un segundo sprint habiendo uno abierto
  Dado un proyecto con el sprint 1 en estado "Abierto"
  Cuando intento crear otro sprint
  Entonces la operacion se rechaza
  Y se informa que el sprint 1 ya esta abierto
  Y el proyecto sigue teniendo 1 sprint

Esquema del escenario: Rechazar un sprint sin Sprint Goal
  Dado un proyecto registrado sin sprints
  Cuando intento crear un sprint con el goal <goal>
  Entonces la operacion se rechaza
  Y se informa que el Sprint Goal es obligatorio
  Y el proyecto sigue sin sprints

  Ejemplos:
    | goal  |
    | ""    |
    | "   " |

Esquema del escenario: Rechazar un sprint sin las fechas obligatorias
  Dado un proyecto registrado sin sprints
  Cuando intento crear un sprint sin <fecha>
  Entonces la operacion se rechaza
  Y se informa que falta <fecha>
  Y el proyecto sigue sin sprints

  Ejemplos:
    | fecha             |
    | "fecha de inicio" |
    | "fecha de fin"    |

Escenario: Rechazar un sprint con fechas incoherentes
  Dado un proyecto registrado sin sprints
  Cuando intento crear un sprint con inicio 14/11/2026 y fin 01/11/2026
  Entonces la operacion se rechaza
  Y se informa que las fechas son incoherentes

Escenario: Un sprint de un solo dia es valido
  Dado un proyecto registrado sin sprints
  Cuando creo un sprint con inicio 01/11/2026 y fin 01/11/2026
  Entonces el sprint queda registrado

Escenario: Un sprint con fechas pasadas es valido
  Dado un proyecto registrado sin sprints
  Cuando creo un sprint con inicio 01/03/2026 y fin 14/03/2026
  Entonces el sprint queda registrado

Escenario: Guardar el Sprint Goal sin espacios en los extremos
  Dado un proyecto registrado sin sprints
  Cuando creo un sprint con el goal "  Tener el backlog estimado  "
  Entonces el sprint queda registrado con el goal "Tener el backlog estimado"

Escenario: Cada proyecto puede tener su propio sprint abierto
  Dado un proyecto A con un sprint abierto y un proyecto B sin sprints
  Cuando creo un sprint en el proyecto B
  Entonces el sprint del proyecto B queda registrado con numero 1
  Y el proyecto A sigue con su sprint abierto

Escenario: Consultar el sprint abierto de un proyecto
  Dado un proyecto con el sprint 1 en estado "Abierto"
  Cuando pido el sprint abierto del proyecto
  Entonces recibo el sprint 1 con su goal, sus fechas y su estado

Escenario: Un proyecto sin sprints no tiene sprint abierto
  Dado un proyecto registrado sin sprints
  Cuando pido el sprint abierto del proyecto
  Entonces se informa que el proyecto no tiene ningun sprint abierto
  Y eso no es un error

Escenario: Un proyecto con todos sus sprints cerrados no tiene sprint abierto
  Dado un proyecto cuyo unico sprint es el numero 1 y esta cerrado
  Cuando pido el sprint abierto del proyecto
  Entonces se informa que el proyecto no tiene ningun sprint abierto

Escenario: Dos sprints del mismo proyecto pueden tener el mismo goal y las mismas fechas
  Dado un proyecto cuyo sprint 1 esta cerrado con el goal "Tener el backlog estimado"
  Cuando creo un sprint nuevo con el goal "Tener el backlog estimado" y las mismas fechas
  Entonces el sprint queda registrado con numero 2
  Y los dos sprints tienen el mismo goal

Esquema del escenario: Operar sobre un proyecto que no existe
  Dado que no existe ningun proyecto con identificador 99
  Cuando intento <operacion> en el proyecto 99
  Entonces se informa que no se encontro el proyecto

  Ejemplos:
    | operacion                 |
    | "crear un sprint"         |
    | "pedir el sprint abierto" |
```

**Para decidir en la puerta:**

1. **RN-1 — el número lo pone el sistema, no quien crea el sprint.** La issue dice "crear un
   sprint con numero", y acá se interpretó que el sprint *tiene* número, no que se escribe a
   mano. Es lo mismo que ya se hizo con el identificador del proyecto (#3) y de la historia
   (#6), y evita tener que decidir qué pasa con un número repetido o salteado. Si el equipo
   quiere elegirlo (por ejemplo para numerar igual que la cátedra, del 0 al 4), hay que
   agregar la validación de número repetido y cambiar la RN-1.
2. **RN-14 — las fechas del sprint no se validan contra las del proyecto.** Parece razonable
   pedir que el sprint caiga adentro del proyecto, pero el proyecto puede no tener fecha de
   fin y sus fechas se pueden modificar después (#3, RN-8): esa validación dejaría sprints
   que se vuelven inválidos solos sin que nadie los toque.
3. **No se valida la superposición de fechas entre sprints.** Como solo puede haber uno
   abierto a la vez (RN-8), el solapamiento únicamente puede aparecer contra sprints ya
   cerrados, y no afecta a nada de lo que se mide.
4. **No existe un estado "Planificado": el sprint se crea ya abierto.** La consecuencia es
   que no se puede dejar preparado el sprint siguiente mientras el actual sigue abierto; hay
   que cerrar uno para crear el otro. Si el equipo quiere planificar con anticipación, hace
   falta un tercer estado y la RN-8 pasa a contar solo los abiertos.
5. **Un sprint de un solo día se acepta** (RN-6 compara por "anterior", no por "distinta").
   No se fija ni un mínimo ni un máximo de duración.
6. **La consulta del sprint abierto (RN-10) entra en esta historia** aunque consultar sprints
   sea la #11, porque sin ella no hay forma de verificar la RN-8 ni de que la #9 sepa a qué
   sprint asignar historias. La #11 consulta los **cerrados**.

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios: [pendiente]

## Fase 2 — Diseño

_Se escribe después de que la Puerta 1 esté aprobada._

## Fase 3 — Tareas

_Se escribe después de que la Puerta 2 esté aprobada._

## Fase 4 — Implementación

_Se escribe después de que la Puerta 3 esté aprobada._

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #8 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | pendiente (Fase 3) |
| Código Go | pendiente (Fase 3) |
