# Spec: Crear un sprint con su Sprint Goal

- **Historia:** #8
- **Sprint:** 2
- **Responsable:** Agustín Salinas
- **Estado:** Terminada en el Sprint 2 — **mergeada a `dev`** el 22/10 (PR #138)
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de
  la consigna y de los criterios de aceptación de la issue. El responsable lo
  revisa, lo corrige si hace falta y lo presenta a las puertas. Revisado por el
  responsable el 20/10: decidió los seis puntos de la puerta y sumó la persistencia
  y el número del sprint abierto en el error de la RN-8.
- **Depende de:** #3 (el sprint pertenece a un proyecto)
- **La necesitan:** #9 (asignar historias al sprint), #10 (cerrar el sprint),
  #11 (consultar sprints anteriores), #16 (métricas), #5 (el sprint en curso del proyecto) y
  #15 (los sprints de detección y de resolución de un defecto)

## Fase 1 — Requisitos

### Objetivo

Que el equipo pueda abrir un sprint en un proyecto, con una meta clara y un principio y un
fin, para organizar el trabajo en períodos cortos. Es la base de todo lo que viene después:
sin un sprint no hay dónde asignar historias, qué cerrar ni qué medir.

### Qué queda afuera

- Asignar historias al sprint y armar el Sprint Backlog, que es la historia #9.
- Cerrar el sprint y registrar qué se completó, que es la #10.
- Listar los sprints ya cerrados y mostrar su detalle, que es la #11. Esta historia solo
  permite encontrar un sprint por su número (RN-15).
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
15. **[RN-15]** El sistema debe permitir consultar un sprint de un proyecto por su número,
    esté abierto o cerrado. Si el proyecto no tiene ningún sprint con ese número, el sistema
    debe informar que no se encontró el sprint.

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
| Número de sprint que no existe en el proyecto | informa que no se encontró el sprint |

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
- [ ] Se puede consultar un sprint por su número, esté abierto o cerrado, y un número que no
      existe informa que no se encontró el sprint.

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

Escenario: Consultar un sprint por su numero
  Dado un proyecto cuyo sprint 1 esta cerrado y cuyo sprint 2 esta abierto
  Cuando pido el sprint 1 del proyecto
  Entonces recibo el sprint 1 con su goal, sus fechas y su estado "Cerrado"

Escenario: Consultar un sprint que no existe
  Dado un proyecto con un unico sprint, el numero 1
  Cuando pido el sprint 5 del proyecto
  Entonces se informa que no se encontro el sprint
```

**Decidido por el responsable antes de la puerta (20/10).** Los seis puntos que el
borrador dejaba abiertos quedan como estaban propuestos. Si quien aprueba la puerta no está
de acuerdo con alguno, es el momento de decirlo:

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
> Aprobó: Maximiliano Eula — Fecha: 2026-10-20 — Comentarios: aprobada en el PR #104 (https://github.com/emisorato1/seguimiento-medicion-ics/pull/104#issuecomment-6303567805). Las condiciones de la revisión de las puertas (de Maxi, Juliana y Emiliano) se aplicaron el 21/10 en 3c5acf2 y 5fbf1df: consulta por número (RN-15, tarea 16), qué se controla al cargar (tarea 15) y sacar "datos en memoria".

## Fase 2 — Diseño

Las decisiones transversales del proyecto —un solo paquete `internal/domain`, errores como
valores del dominio comparables con `errors.Is`, mensajes en castellano, la persistencia en
un archivo desde la #77 y validar todo antes de cambiar nada— están en
[AGENTS.md](../../../AGENTS.md) y no se repiten acá. Los datos del sprint y los nombres de
sus operaciones ya están fijados en [el contrato del dominio](../../diseno-dominio.md).
Esta fase registra solo lo propio de esta historia.

### Enfoque

El sprint es una entidad nueva, pero **no es una entidad suelta: vive dentro del proyecto**,
igual que el Product Backlog (#6) y los integrantes (#4). Un sprint no existe sin proyecto,
y guardarlo adentro hace que el número único por proyecto (RN-1, RN-9) salga solo, sin tener
que llevar un índice aparte.

Crear un sprint sigue el mismo camino que crear un proyecto o una historia:

1. **Normalizar** lo que llegó: sacarle los espacios de los extremos al Sprint Goal (RN-7) y
   dejar las fechas en día, sin hora.
2. **Validar** sobre lo ya normalizado: el goal, las dos fechas, su coherencia y que el
   proyecto no tenga otro sprint abierto (RN-4 a RN-6, RN-8).
3. **Aplicar**: recién ahí se le asigna el número y se agrega a los sprints del proyecto.

Ese orden es lo que garantiza la RN-12 sin tener que deshacer nada.

La regla de "un solo sprint abierto" (RN-8) y la consulta del sprint abierto (RN-10) son
**la misma pregunta**: antes de crear, el sistema se pregunta si ya hay uno abierto y, si lo
hay, lo informa en el rechazo. Por eso la consulta se implementa primero y la validación la
usa, en lugar de recorrer los sprints dos veces con dos criterios distintos.

Esta historia crea la entidad **con todos los datos que fija el contrato del dominio**,
aunque solo use algunos: el estado arranca en Abierto, la lista de historias arranca vacía y
la fecha de cierre arranca vacía. Las llenan la #9 y la #10.

### Archivos afectados

| Archivo | Se crea o se modifica | Para qué |
|---|---|---|
| `internal/domain/sprint.go` | se crea | la entidad sprint con los datos del contrato, su estado como tipo propio, la creación dentro del proyecto, el listado de sprints y la consulta del sprint abierto |
| `internal/domain/sprint_test.go` | se crea | los tests de los escenarios BDD de la Fase 1, nombrados como cada escenario |
| `internal/domain/project.go` | se modifica | el proyecto pasa a contener sus sprints y el contador de números |
| `internal/domain/errors.go` | se modifica | los seis errores de esta historia |
| `internal/domain/registry_json.go` | se modifica | los sprints se guardan y se vuelven a cargar con el proyecto |
| `internal/domain/registry_json_test.go` | se modifica | el caso de los sprints en el test de ida y vuelta |

### Decisiones

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| Los sprints viven dentro del proyecto | una colección de sprints aparte, indexada por proyecto | un sprint no existe sin proyecto, y adentro el número único por proyecto (RN-9) sale gratis, igual que el identificador de historia en la #6 |
| El contador de números es del proyecto y **nunca retrocede** | calcular el próximo número como "cantidad de sprints + 1" | contar se rompe apenas exista borrado y repetiría un número ya usado. Es la misma decisión que tomó la #6 para las historias |
| El estado del sprint es un tipo propio con dos valores, y **Abierto es el valor inicial** | un texto libre, o dejar un valor inicial "sin estado" como hizo la #4 con el rol | un texto libre deja pasar estados que no existen; y un sprint sin estado no existe: nace abierto (RN-3), así que reservar un valor para "sin estado" sería reservarlo para algo imposible |
| La consulta del sprint abierto informa aparte si hay alguno, en lugar de devolver un error cuando no hay | devolver el error "el proyecto no tiene sprint abierto" | la RN-10 dice que no tener sprint abierto **no es un error**. Con un error, quien consulta tendría que distinguir "no hay" de "falló", que es justo lo que la regla quiere evitar. La #9 y la #10 sí lo tratan como error en sus propias operaciones, y ese error lo declara la #9 |
| La RN-8 se verifica con la misma consulta de la RN-10 | guardar aparte cuál es el sprint abierto del proyecto | un dato aparte puede contradecir a la lista de sprints; con una sola fuente eso no puede pasar |
| Las fechas del sprint se guardan por día, reusando la normalización que ya hace el proyecto | guardarlas tal como llegan, con hora y zona | los casos límite de la Fase 1 piden que dos fechas del mismo día sean la misma fecha; reusar lo de la #3 evita que el sprint compare distinto que el proyecto (defecto #59) |
| La entidad se crea completa, con los datos del contrato que esta historia no usa | crear solo el número, el goal, las fechas y el estado | el contrato ya los fija, y si no están, la #9 y la #10 tienen que volver a tocar la entidad y chocan entre ellas en el mismo archivo |
| El rechazo por sprint abierto **envuelve el error del dominio con el número** del sprint que está abierto | un error fijo sin el número, o un error distinto por cada número | la RN-8 pide informar cuál está abierto, y un error del paquete es un valor fijo. Envuelto, `errors.Is` lo sigue reconociendo y el mensaje que ve el usuario dice el número |
| Los sprints se guardan en el archivo de datos junto con su proyecto, y al cargarlos se controla que el archivo sea coherente: el Sprint Goal no vacío y las fechas presentes y en orden (como en la RN-4 a la RN-6), **como mucho un sprint abierto por proyecto**, y números de sprint **únicos, desde 1 y menores que el contador** del proyecto. No se pueden usar literalmente "las mismas validaciones que al crear", porque crear siempre deja el sprint Abierto y un archivo puede traer sprints cerrados | dejarlos solo en memoria | desde la #77 la web guarda todo en un archivo, y `AGENTS.md` pide que toda entidad nueva entre en esa conversión con su caso de ida y vuelta: si no, la web perdería los sprints al cerrarse sin que ningún test lo note. Esta historia guarda los datos que crea; los del cierre los suma la #10 |
| **La consulta de un sprint por su número y su error los adelanta esta historia**, aunque el contrato se los asignaba a la #11 | esperar a la #11 | la #15 los necesita desde su primera tarea para validar los sprints de detección y de resolución de un defecto, y la #11 está en el Sprint 3. Son pocas líneas, van en el PR temprano y la #11 las usa tal cual. Lo pidieron Maxi y Juliana en la revisión de las puertas (PR #104), y el contrato quedó actualizado |
| Lo de esta historia va en `sprint.go`; la #9, la #10 y la #11 agregan archivos propios | que las cuatro historias del sprint escriban en `sprint.go` | es el criterio que ya se aprobó en la Puerta 2 de la #12 para `story_estimation.go`: separar por funcionalidad para que dos personas no editen el mismo archivo |

### Dependencias y restricciones

- **Depende de la #3.** Modifica `project.go`, así que el código no se empieza hasta que la
  #3 esté en `dev`.
- **`errors.go` es compartido.** Los errores de esta historia se agregan ahí, al final, para
  achicar el conflicto con las otras historias del sprint.
- **`registry_json.go` también es compartido.** La conversión del sprint va en funciones
  propias, y del código que ya existe solo se toca lo mínimo para que el proyecto guarde y
  cargue sus sprints.
- **No se toca nada de las historias:** ni el backlog, ni el estado, ni la estimación. El
  sprint nace sin historias asignadas y quien las asigna es la #9.
- **No se escribe el cierre.** El pase a Cerrado y la fecha de cierre son de la #10: esta
  historia solo deja el dato vacío.
- **Las operaciones son las que ya nombra el contrato del dominio** (`p.AddSprint(...)`,
  `p.Sprints()`, `p.OpenSprint()` y `p.Sprint(number)`). La consulta por número la adelanta
  esta historia; la #11 solo suma el listado de los cerrados.
- **Las usan:** la #9 (necesita el sprint abierto para asignar), la #10 (lo cierra), la #11
  (lista los cerrados), la #16 (mide sobre ellos), la #5 (el sprint en curso) y la #15 (busca
  los sprints de detección y de resolución por su número).

> **Puerta 2** — ¿el diseño es viable y coherente con el resto del proyecto?
> Aprobó: Maximiliano Eula — Fecha: 2026-10-20 — Comentarios: aprobada en el PR #104 (https://github.com/emisorato1/seguimiento-medicion-ics/pull/104#issuecomment-6303567805). Las condiciones de la revisión de las puertas (de Maxi, Juliana y Emiliano) se aplicaron el 21/10 en 3c5acf2 y 5fbf1df: consulta por número (RN-15, tarea 16), qué se controla al cargar (tarea 15) y sacar "datos en memoria".

## Fase 3 — Tareas

Cada tarea es un ciclo de TDD completo: el test que falla en un commit, la implementación que
lo hace pasar en otro. Están en el orden en que hay que hacerlas, y cada una deja el código
compilando y los tests en verde.

Los tests siguen la convención de [AGENTS.md](../../../AGENTS.md): un test por escenario, con
el nombre del escenario. Cada tarea de rechazo verifica, además del error, que el proyecto
quedó con los mismos sprints que antes: es la RN-12.

| # | Tarea | Escenario BDD | Reglas | Estado |
|---|---|---|---|---|
| 1 | Un proyecto recién creado no tiene sprints y no tiene sprint abierto, y eso no es un error | Un proyecto sin sprints no tiene sprint abierto | RN-10 | hecha |
| 2 | Crear el primer sprint: queda con número 1, en Abierto y sin historias asignadas | Crear el primer sprint de un proyecto | RN-1, RN-2, RN-3 | hecha |
| 3 | La consulta del sprint abierto devuelve ese sprint con su goal, sus fechas y su estado | Consultar el sprint abierto de un proyecto | RN-10 | hecha |
| 4 | Rechazar la creación sin Sprint Goal, con el goal de solo espacios incluido; el proyecto sigue sin sprints | Rechazar un sprint sin Sprint Goal | RN-4, RN-12 | hecha |
| 5 | Rechazar la creación sin fecha de inicio y sin fecha de fin, informando cuál de las dos falta | Rechazar un sprint sin las fechas obligatorias | RN-5, RN-12 | hecha |
| 6 | Rechazar la creación con la fecha de fin anterior a la de inicio | Rechazar un sprint con fechas incoherentes | RN-6, RN-12 | hecha |
| 7 | Un sprint de un solo día y un sprint con las dos fechas en el pasado se aceptan | Un sprint de un solo dia es valido / Un sprint con fechas pasadas es valido | RN-6, RN-14 | hecha |
| 8 | El Sprint Goal se guarda sin los espacios del principio y del final | Guardar el Sprint Goal sin espacios en los extremos | RN-7 | hecha |
| 9 | Con un sprint abierto, crear otro se rechaza informando cuál está abierto, y el proyecto sigue con un solo sprint | No se puede abrir un segundo sprint habiendo uno abierto | RN-8, RN-12 | hecha |
| 10 | Cerrado el sprint 1, el sprint nuevo recibe el número 2 y el proyecto queda con dos sprints | El sprint nuevo recibe el numero siguiente | RN-1, RN-2 | hecha |
| 11 | Un proyecto con todos sus sprints cerrados no tiene sprint abierto | Un proyecto con todos sus sprints cerrados no tiene sprint abierto | RN-10 | hecha |
| 12 | Dos sprints del mismo proyecto pueden tener el mismo goal y las mismas fechas | Dos sprints del mismo proyecto pueden tener el mismo goal y las mismas fechas | RN-11 | hecha |
| 13 | Dos proyectos tienen cada uno su sprint número 1, y el sprint abierto de uno no impide crear el del otro | Cada proyecto puede tener su propio sprint abierto | RN-9 | hecha |
| 14 | Crear un sprint o pedir el sprint abierto de un proyecto que no existe informa que no se encontró el proyecto | Operar sobre un proyecto que no existe | RN-13 | hecha |
| 15 | Los sprints de un proyecto se conservan al guardar y volver a cargar el registro, con su número, goal, fechas y estado, y el contador sigue desde donde estaba. Un archivo con el goal vacío, fechas faltantes o al revés, dos sprints abiertos, o un número repetido, menor que 1 o no menor que el contador se rechaza al cargar | — (casos en el test de ida y vuelta y en el de datos dañados de `registry_json_test.go`) | RN-1, RN-2, RN-4 a RN-6, RN-8 | hecha |
| 16 | Consultar un sprint por su número, abierto o cerrado, y un número que no existe informa que no se encontró el sprint | Consultar un sprint por su numero; Consultar un sprint que no existe | RN-15 | hecha |

**Las tareas 1, 2, 3 y 16 van en un primer pull request**, que se mergea apenas esté: crean
el sprint, la consulta del sprint abierto y la consulta por número, que es lo que necesitan la
#5, la #10, la #15 y la #16 para empezar. Es el caso 1 de "Cuándo se puede partir una historia"
en `CLAUDE.md`. El resto va en el segundo, que cierra la historia.

### Por qué este orden

- **La 1 va primero** porque la consulta del sprint abierto es lo que usan casi todas las
  demás para verificar el resultado, y en un proyecto vacío se puede escribir sin que exista
  todavía la creación.
- **La 2 es el camino feliz** y deja la entidad armada; de la 4 a la 6 se le agregan las
  validaciones de a una, que recién ahí tienen contra qué fallar.
- **La 9 va después de la 2** porque necesita un sprint ya creado para que haya uno abierto.
- **La 10, la 11 y la 12 necesitan un sprint cerrado**, así que van después de que la regla
  del sprint abierto esté cerrada: son las que prueban que el número sigue avanzando.
- **La 13 y la 14 van al final** porque miran el borde de afuera —otro proyecto, o ninguno— y
  solo tienen sentido cuando lo de adentro ya funciona.
- **La 15 cierra la historia**: guarda en el archivo todo lo que las anteriores dejaron
  armado. No sale de un escenario de la Fase 1 sino de la regla de persistencia de
  `AGENTS.md`, y por eso es un caso del test de ida y vuelta que ya existe.
- **La 16 se agregó en la revisión de las puertas** y va al final de la tabla para no
  renumerar las demás, pero **se hace inmediatamente después de la 3**: va en el primer pull
  request. Para probar el sprint cerrado, el test fija el estado desde dentro del paquete,
  como las tareas 10 a 12.

### Cobertura

Las quince reglas quedan cubiertas: RN-1 (2, 10, 15), RN-2 (2, 10, 15), RN-3 (2), RN-4 (4),
RN-5 (5), RN-6 (6, 7), RN-7 (8), RN-8 (9), RN-9 (13), RN-10 (1, 3, 11), RN-11 (12),
RN-12 (4, 5, 6, 9), RN-13 (14), RN-14 (7), RN-15 (16).

| Criterio de aceptación | Tareas |
|---|---|
| Se puede crear un sprint y queda registrado con su número y en Abierto | 2 |
| El primer sprint recibe el número 1 y cada sprint nuevo el siguiente | 2, 10 |
| El número no se repite dentro del mismo proyecto | 10 |
| Un sprint recién creado no tiene historias asignadas | 2 |
| No se puede crear un sprint sin Sprint Goal, y el goal de solo espacios cuenta como vacío | 4 |
| No se puede crear un sprint sin fecha de inicio ni sin fecha de fin | 5 |
| No se puede crear un sprint con fecha de fin anterior a la de inicio | 6 |
| Un sprint cuya fecha de fin es igual a la de inicio se acepta | 7 |
| No se puede crear un segundo sprint habiendo uno abierto, y el mensaje dice cuál es | 9 |
| Tener un sprint abierto en un proyecto no impide crear uno en otro proyecto | 13 |
| Se puede consultar el sprint abierto, y no tener ninguno no es un error | 1, 3, 11 |
| El Sprint Goal se guarda sin los espacios de los extremos | 8 |
| Una creación rechazada no agrega ningún sprint | 4, 5, 6, 9 |
| Operar sobre un proyecto inexistente informa que no se encontró | 14 |
| Se puede consultar un sprint por su número, y uno que no existe informa que no se encontró | 16 |

Los diecisiete escenarios de la Fase 1 tienen su tarea en la columna "Escenario BDD".

### Dependencias

- **Todas las tareas necesitan la #3 en `dev`**: el proyecto y su registro.
- **Las tareas 10, 11 y 12 necesitan un sprint cerrado**, y cerrar es la #10, que todavía no
  existe. Los tests lo preparan desde dentro del paquete, que puede fijar el estado
  directamente. Es lo mismo que hizo la #12 con las historias Terminadas.

> **Puerta 3** — ¿las tareas cubren todos los criterios de aceptación?
> Aprobó: Maximiliano Eula — Fecha: 2026-10-20 — Comentarios: aprobada en el PR #104 (https://github.com/emisorato1/seguimiento-medicion-ics/pull/104#issuecomment-6303567805). Las condiciones de la revisión de las puertas (de Maxi, Juliana y Emiliano) se aplicaron el 21/10 en 3c5acf2 y 5fbf1df: consulta por número (RN-15, tarea 16), qué se controla al cargar (tarea 15) y sacar "datos en memoria".

## Fase 4 — Implementación

Ciclo RED → GREEN → REFACTOR. **El test en un commit y la implementación en otro.**

### Primer pull request (PR #104): tareas 1, 2, 3 y 16

| Tarea | Commit del test | Commit de la implementación |
|---|---|---|
| 1 | `5e2b377` | `ff3264e` |
| 2 | `5d602fa` | `e3a9416` |
| 3 | `6cf7475` | `de2ac50` |
| 16 | `5de4cda` | `ee69eed` |

Las cuatro fallaron antes de su implementación, y cada commit de test dice por qué.

- **El contador de sprints guarda el último número asignado**, que arranca en cero, y no "el
  próximo". Así un proyecto que se carga de un archivo que todavía no guarda sprints (la
  tarea 15 va en el segundo pull request) numera igual desde 1, sin inicializar nada.
- **Este pull request todavía no valida ni guarda en el archivo**: las validaciones son las
  tareas 4 a 9 y la persistencia la 15. Nadie crea sprints desde la interfaz todavía (la
  web es la #78), así que no hay forma de que entre un sprint inválido mientras tanto.

### Segundo pull request: tareas 4 a 15

Va junto con la #9, en el mismo pull request, como permite el modo sprint final para dos
eslabones seguidos de la misma cadena.

| Tarea | Commit del test | Commit de la implementación |
|---|---|---|
| 4 | `07deed8` | `a2cadc9` |
| 5 | `6991115` | `20eb08c` |
| 6 | `8511e9d` | `734a903` |
| 7 | `fcbd577` | — (ver abajo) |
| 8 | `41396e9` | `eb536ef` |
| 9 | `ffaada8` | `c59f3b2` |
| 10 | `19394e3` | — (ver abajo) |
| 11 | `7309747` | — (ver abajo) |
| 12, 13 y 14 | `cfde49b` | — (ver abajo) |
| 15 | `b2a1db5` y `89b90fb` | `8e05211` |

### Desvíos respecto de la spec

- **Tareas sin commit de implementación: la 7 y de la 10 a la 14.** Sus tests pasaron en
  verde al escribirlos, como anticipaba el aviso de la Fase 3: la tarea 6 ya rechazaba solo
  un fin *anterior* al inicio y nada compara contra la fecha de hoy (7), el contador de la
  tarea 2 ya numeraba desde el último (10), la consulta de la tarea 3 ya buscaba un sprint
  abierto (11), nada compara goals ni fechas entre sprints (12), cada proyecto tiene sus
  sprints y su contador (13) y el registro de la #3 ya informa el proyecto inexistente (14).
  Cada commit lo explica. Forzar un RED habría sido simular TDD.
- **El test de la tarea 14 se llama `TestOperarSprintsSobreUnProyectoQueNoExiste`** y no con
  el nombre exacto del escenario, porque la #6 ya tiene un escenario "Operar sobre un
  proyecto que no existe" con su test. El comentario `// Esquema del escenario:` conserva el
  nombre, y por ahí lo encuentra la trazabilidad.
- **El test de la tarea 9 mira también el texto del error**, además de `errors.Is`: la RN-8
  pide que el aviso diga cuál es el sprint abierto, y eso es lo único que agrega el error
  envuelto.
- **La tarea 15 tiene un segundo commit de test (`89b90fb`)**: dos casos de fechas de sprint
  mal escritas en el archivo, que pasan con la implementación y cubren las dos lecturas de
  fecha que habían quedado sin test. Devuelven el dominio al 100% de cobertura.
- **La consola y la web todavía no usan los sprints**: eso es de la #24 y la #78.

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #8 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | `internal/domain/sprint_test.go` (un test por escenario, 17) y los casos de sprints en `internal/domain/registry_json_test.go` (tarea 15) |
| Código Go | `internal/domain/sprint.go`, con los sprints dentro de `internal/domain/project.go`, su conversión en `internal/domain/registry_json.go` y sus errores en `internal/domain/errors.go` |
