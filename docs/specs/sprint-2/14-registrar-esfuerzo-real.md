# Spec: Registrar el esfuerzo real

- **Historia:** #14
- **Sprint:** 2
- **Responsable:** Juliana Bustos
- **Estado:** Borrador completo — pendiente de las tres puertas
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de la consigna y de
  los criterios de aceptación de la issue. El responsable lo revisa, lo corrige si hace falta
  y lo presenta a las puertas. Revisado por la responsable el 20/10 y presentado a las tres
  puertas por pull request.
- **Depende de:** #3 (el proyecto), #4 (el integrante que carga las horas) y #6 (la historia
  sobre la que se trabajó)
- **La necesitan:** #16 (horas reales y desviación contra las estimadas), #17 (dashboard),
  #18 (reporte)

## Fase 1 — Requisitos

### Objetivo

Que cada integrante pueda anotar las horas que trabajó de verdad en una historia, día por
día y diciendo qué hizo, para que después se pueda comparar el esfuerzo estimado con el real.
Sin este registro no hay con qué medir la desviación.

### Qué queda afuera

- **Calcular la desviación entre horas estimadas y reales**, las horas totales del proyecto y
  las de un sprint: todo eso es la #16. Acá solo se guardan las horas y se suman por historia
  y por integrante. Cargar las horas **estimadas** también es de la #16.
- **Atar el registro a un sprint.** Las horas se anotan sobre una historia; a qué sprint
  pertenecen se sabe por el sprint de esa historia, y eso lo resuelve la #16 (ver el punto 2
  de "Para decidir en la puerta").
- **Modificar o borrar un registro** ya cargado, y **aprobar o revisar** lo que carga cada
  uno: lo que se registra se guarda.
- **Controlar la jornada**: no hay máximo de horas por día, por semana ni por persona.
- Cómo se muestran las horas en pantalla: eso es de la interfaz.

### Entradas

| Dato | Obligatorio | Descripción |
|---|---|---|
| Proyecto | sí | en qué proyecto se registra |
| Historia | sí | sobre qué historia se trabajó |
| Integrante | sí | quién trabajó, por su nombre y apellido |
| Día | sí | qué día se trabajó |
| Horas | sí | cuántas horas, mayor que cero, admite fracciones |
| Actividad realizada | sí | qué se hizo en esas horas |

### Salidas

Un registro de esfuerzo guardado dentro del proyecto. Se pueden consultar todos los registros
del proyecto, el total de horas de una historia y el total de horas de un integrante.

### Reglas de negocio

1. **[RN-1]** Cuando se registra esfuerzo, el sistema debe guardar la historia, el
   integrante, el día, las horas y la actividad realizada.
2. **[RN-2]** El sistema debe aceptar como horas únicamente un valor mayor que cero, y debe
   admitir fracciones de hora, por ejemplo media hora.
3. **[RN-3]** Si las horas son cero o un valor negativo, el sistema debe rechazar la
   operación e informar que las horas tienen que ser mayores que cero.
4. **[RN-4]** Si el día del registro es posterior al día de hoy, el sistema debe rechazar la
   operación e informar que no se pueden cargar horas de un día que todavía no pasó.
5. **[RN-5]** Si no se indica el día, el sistema debe rechazar la operación e informar que la
   fecha es obligatoria.
6. **[RN-6]** Si la actividad realizada está vacía, el sistema debe rechazar la operación e
   informar que es obligatoria. Un texto compuesto solo por espacios cuenta como vacío, y lo
   que se guarda va sin los espacios del principio y del final.
7. **[RN-7]** Si quien registra no es integrante del proyecto, el sistema debe rechazar la
   operación e informar que no se encontró el integrante. El nombre se busca igual que en la
   #4: ignorando las mayúsculas y los espacios de los extremos.
8. **[RN-8]** Cuando se guarda el registro, el sistema debe guardar el nombre del integrante
   tal como está cargado en el proyecto, y no tal como se escribió al registrar.
9. **[RN-9]** Si la historia no pertenece al proyecto, el sistema debe rechazar la operación e
   informar que no se encontró la historia.
10. **[RN-10]** El sistema debe permitir registrar horas sobre una historia en cualquier
    estado, incluida una Terminada: las horas se cargan después de trabajar, y el trabajo
    puede haber terminado.
11. **[RN-11]** El sistema debe permitir más de un registro del mismo integrante, sobre la
    misma historia y el mismo día: cada registro es una carga de horas con su propia
    actividad, y todas suman.
12. **[RN-12]** El sistema debe tratar el día del registro como un día, sin hora: dos
    registros del mismo día son del mismo día aunque lleguen con horas distintas.
13. **[RN-13]** Cuando se piden las horas de una historia, el sistema debe devolver la suma
    de las horas de todos sus registros. Una historia sin registros tiene cero horas, y eso
    no es un error.
14. **[RN-14]** Cuando se piden las horas de un integrante, el sistema debe devolver la suma
    de las horas de todos los registros que cargó en ese proyecto. Un integrante sin
    registros tiene cero horas, y eso no es un error.
15. **[RN-15]** El sistema debe permitir consultar todos los registros de esfuerzo del
    proyecto, en el orden en que se cargaron. Un proyecto sin registros devuelve una lista
    vacía, que no es un error.
16. **[RN-16]** El sistema debe tratar cada proyecto por separado: los registros de un
    proyecto no suman ni aparecen en otro.
17. **[RN-17]** El sistema debe comparar el día del registro únicamente con el día de hoy
    (RN-4), sin compararlo con las fechas del proyecto ni con las de ningún sprint.
18. **[RN-18]** Si un registro se rechaza, el sistema debe dejar el proyecto y sus registros
    como estaban antes.
19. **[RN-19]** Si se intenta registrar esfuerzo o consultar horas en un proyecto que no
    existe, el sistema debe informar que no se encontró el proyecto y no registrar nada.

### Casos límite

- El día de hoy: válido, es el límite exacto de la RN-4. El día de mañana: rechazado.
- Media hora: válida. Cero horas y horas negativas: rechazadas.
- Historia sin registros, integrante sin registros y proyecto sin registros: cero horas y
  lista vacía, y nada de eso es un error.
- Dos registros del mismo integrante, la misma historia y el mismo día: quedan los dos y
  suman.
- Día cargado con hora: cuenta el día, la hora no. Día muy anterior al inicio del proyecto:
  válido.
- Horas sobre una historia Terminada: válido.
- Nombre del integrante con otras mayúsculas o con espacios de más: se encuentra igual, y se
  guarda el nombre tal como está cargado.
- El mismo nombre y apellido en dos proyectos distintos: cada proyecto suma solo lo suyo.

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| Horas cero o negativas | rechaza e informa que las horas tienen que ser mayores que cero |
| Día posterior al de hoy | rechaza e informa que no se pueden cargar horas de un día que todavía no pasó |
| Falta el día | rechaza e informa que la fecha es obligatoria |
| Actividad vacía o de solo espacios | rechaza e informa que la actividad realizada es obligatoria |
| Quien registra no es integrante del proyecto | rechaza e informa que no se encontró el integrante |
| La historia no pertenece al proyecto | rechaza e informa que no se encontró la historia |
| Proyecto inexistente al registrar o al consultar | informa que no se encontró el proyecto y no registra nada |

### Criterios de aceptación

- [ ] Se puede registrar esfuerzo indicando integrante, historia, día, horas y actividad
      realizada, y el registro queda guardado con esos cinco datos.
- [ ] Se aceptan fracciones de hora, por ejemplo media hora.
- [ ] No se pueden registrar cero horas ni un valor negativo.
- [ ] No se pueden registrar horas con fecha posterior al día de hoy, y el día de hoy sí se
      acepta.
- [ ] No se puede registrar esfuerzo sin día ni sin actividad realizada; una actividad de
      solo espacios cuenta como vacía y lo que se guarda va sin los espacios de los extremos.
- [ ] No se puede registrar esfuerzo de alguien que no es integrante del proyecto; el nombre
      se encuentra aunque venga con otras mayúsculas o espacios de más, y se guarda tal como
      está cargado en el proyecto.
- [ ] No se puede registrar sobre una historia que no pertenece al proyecto, y sí sobre una
      historia Terminada.
- [ ] El mismo integrante puede cargar dos veces la misma historia el mismo día, y las dos
      cargas quedan y suman.
- [ ] Se puede consultar el total de horas de una historia, y una historia sin registros da
      cero sin que sea un error.
- [ ] Se puede consultar el total de horas de un integrante, y un integrante sin registros da
      cero sin que sea un error.
- [ ] Se pueden consultar todos los registros del proyecto, en el orden en que se cargaron.
- [ ] Las horas cargadas en un proyecto no suman en otro.
- [ ] Un registro rechazado no deja nada guardado.
- [ ] Registrar o consultar en un proyecto inexistente informa que no se encontró el proyecto.

### Escenarios BDD

```gherkin
Escenario: Registrar horas trabajadas en una historia
  Dado un proyecto con la integrante "Juliana Bustos" y la historia 1
  Cuando registro 3 horas de "Juliana Bustos" en la historia 1 el dia de hoy con la actividad "Escribir los tests de la estimacion"
  Entonces el registro queda guardado con la historia 1, "Juliana Bustos", el dia de hoy, 3 horas y esa actividad
  Y el proyecto pasa a tener 1 registro de esfuerzo

Escenario: Registrar media hora
  Dado un proyecto con la integrante "Juliana Bustos" y la historia 1
  Cuando registro media hora en la historia 1
  Entonces el registro queda guardado con media hora

Esquema del escenario: Rechazar horas que no son mayores que cero
  Dado un proyecto con la integrante "Juliana Bustos" y la historia 1
  Cuando intento registrar <horas> horas en la historia 1
  Entonces la operacion se rechaza
  Y se informa que las horas tienen que ser mayores que cero
  Y el proyecto sigue sin registros de esfuerzo

  Ejemplos:
    | horas |
    | 0     |
    | -2    |

Escenario: Rechazar horas de un dia que todavia no paso
  Dado un proyecto con la integrante "Juliana Bustos" y la historia 1
  Cuando intento registrar 2 horas con fecha del dia de manana
  Entonces la operacion se rechaza
  Y se informa que no se pueden cargar horas de un dia que todavia no paso
  Y el proyecto sigue sin registros de esfuerzo

Escenario: Registrar horas de un dia anterior al inicio del proyecto
  Dado un proyecto que empezo el 01/10/2026, con la integrante "Juliana Bustos" y la historia 1
  Cuando registro 2 horas con fecha 15/09/2026
  Entonces el registro queda guardado

Escenario: Rechazar un registro sin fecha
  Dado un proyecto con la integrante "Juliana Bustos" y la historia 1
  Cuando intento registrar 2 horas sin indicar el dia
  Entonces la operacion se rechaza
  Y se informa que la fecha es obligatoria

Esquema del escenario: Rechazar un registro sin actividad realizada
  Dado un proyecto con la integrante "Juliana Bustos" y la historia 1
  Cuando intento registrar 2 horas con la actividad <actividad>
  Entonces la operacion se rechaza
  Y se informa que la actividad realizada es obligatoria
  Y el proyecto sigue sin registros de esfuerzo

  Ejemplos:
    | actividad |
    | ""        |
    | "   "     |

Escenario: Guardar la actividad sin los espacios de los extremos
  Dado un proyecto con la integrante "Juliana Bustos" y la historia 1
  Cuando registro 2 horas con la actividad "  Revisar el pull request  "
  Entonces el registro queda guardado con la actividad "Revisar el pull request"

Escenario: Rechazar el registro de alguien que no es integrante del proyecto
  Dado un proyecto con la integrante "Juliana Bustos" y la historia 1
  Cuando intento registrar 2 horas a nombre de "Pedro Gomez"
  Entonces la operacion se rechaza
  Y se informa que no se encontro el integrante
  Y el proyecto sigue sin registros de esfuerzo

Escenario: Encontrar al integrante aunque el nombre venga escrito distinto
  Dado un proyecto con la integrante "Juliana Bustos" y la historia 1
  Cuando registro 2 horas a nombre de "  juliana bustos  "
  Entonces el registro queda guardado a nombre de "Juliana Bustos"

Escenario: Rechazar el registro sobre una historia que no es del proyecto
  Dado un proyecto con la integrante "Juliana Bustos" y sin ninguna historia con identificador 99
  Cuando intento registrar 2 horas en la historia 99
  Entonces la operacion se rechaza
  Y se informa que no se encontro la historia

Escenario: Registrar horas sobre una historia terminada
  Dado un proyecto con la integrante "Juliana Bustos" y la historia 1 en estado "Terminada"
  Cuando registro 2 horas en la historia 1
  Entonces el registro queda guardado

Escenario: Dos cargas del mismo integrante, la misma historia y el mismo dia
  Dado un proyecto con la integrante "Juliana Bustos" y la historia 1
  Cuando registro 2 horas en la historia 1 hoy con la actividad "Escribir los tests"
  Y registro 1 hora en la historia 1 hoy con la actividad "Corregir la revision"
  Entonces el proyecto tiene 2 registros de esfuerzo
  Y la historia 1 tiene 3 horas

Escenario: Consultar el total de horas de una historia
  Dado un proyecto con la historia 1, con registros de 2 horas y de 1 hora y media
  Cuando consulto las horas de la historia 1
  Entonces recibo 3 horas y media

Escenario: Consultar el total de horas de un integrante
  Dado un proyecto donde "Juliana Bustos" cargo 2 horas en la historia 1 y 3 horas en la historia 2
  Cuando consulto las horas de "Juliana Bustos"
  Entonces recibo 5 horas

Esquema del escenario: Sin registros, el total es cero
  Dado un proyecto con la historia 1 y la integrante "Juliana Bustos", sin ningun registro
  Cuando consulto las horas de <quien>
  Entonces recibo cero horas
  Y eso no es un error

  Ejemplos:
    | quien            |
    | la historia 1    |
    | "Juliana Bustos" |

Escenario: Las horas cargadas en un proyecto no suman en otro
  Dado un proyecto A donde "Juliana Bustos" cargo 4 horas y un proyecto B donde cargo 1 hora
  Cuando consulto las horas de "Juliana Bustos" en el proyecto B
  Entonces recibo 1 hora

Escenario: Registrar esfuerzo en un proyecto que no existe
  Dado que no existe ningun proyecto con identificador 99
  Cuando intento registrar 2 horas en el proyecto 99
  Entonces se informa que no se encontro el proyecto
```

**Para decidir en la puerta:**

1. **La actividad realizada es obligatoria.** La issue la pide como dato del registro ("se
   registra integrante, fecha, actividad realizada y horas trabajadas"). **Resuelto en el
   contrato:** [`docs/diseno-dominio.md`](../../diseno-dominio.md) ya tiene la actividad
   como dato obligatorio del `Effort` y nombra las dos consultas de totales, por historia y
   por integrante. No hace falta tocar el contrato antes de implementar.
2. **El registro no se ata a un sprint.** El contrato no le puso sprint al `Effort`, así que
   las horas de un sprint salen de las historias asignadas a ese sprint, y esa cuenta es de
   la #16. La consecuencia: si una historia pasa por dos sprints, todas sus horas caen en el
   sprint donde esté asignada. Si el equipo quiere horas por sprint exactas, hay que
   agregarle el sprint al registro y cambiar el contrato.
3. **Varias cargas el mismo día sobre la misma historia (RN-11).** El contrato dice que "hay
   un registro por persona, historia y día"; acá se interpretó que eso describe qué
   identifica a un registro, no que no pueda haber dos. Con la actividad adentro, obligar a
   uno solo por día obligaría a rechazar la segunda carga o a acumular las horas perdiendo
   una de las dos actividades. Para la #16 da lo mismo, porque suma.
4. **El día de hoy se acepta y el de mañana no (RN-4).** Es la lectura literal de la issue
   ("fecha posterior al día de hoy"). Nadie carga horas que todavía no trabajó, pero sí a
   última hora del día en curso.
5. **No hay tope de horas y un registro no se corrige ni se borra.** Poner un máximo
   obligaría a elegir el número y a decidir si se cuenta por registro o por día, y corregir
   un registro es otra historia con su tarjeta. Ni la issue ni la consigna piden ninguna de
   las dos cosas.

6. **Nota del Agile Enabler — rechazar días futuros obliga al dominio a leer el reloj.**
   Sería la primera vez: hoy ninguna parte de `internal/domain` lo hace, y la #8 decidió lo
   contrario para las fechas del sprint (su RN-14 no las compara con hoy). Leer el reloj
   adentro hace que los tests dependan del día en que se corren. Las dos salidas razonables:
   que el "hoy" lo pase quien llama (la consola), o sacar la regla del día futuro y aceptar
   cualquier fecha, como la #8. **Conviene decidirlo antes de la Puerta 2.**
   **Decidido el 20/10: el "hoy" lo pasa quien llama.** La regla del día futuro se mantiene.

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios: [pendiente]

## Fase 2 — Diseño

Las decisiones transversales del proyecto —un solo paquete `internal/domain`, datos en
memoria, errores como valores del dominio comparables con `errors.Is`, mensajes en castellano
y validar todo antes de cambiar nada— están en [AGENTS.md](../../../AGENTS.md) y no se
repiten acá. Los nombres de los datos y de las operaciones son los del contrato
([`docs/diseno-dominio.md`](../../diseno-dominio.md)): `Effort`, con `StoryID`, `MemberName`,
`Activity`, `Hours` y `Day`, registrado con `LogEffort` y consultado con `Efforts`,
`StoryHours` y `MemberHours`.

### Enfoque

El registro de esfuerzo es una entidad nueva que **cuelga del proyecto**, como los
integrantes (#4) y las historias (#6): no existe sin su proyecto, y así las horas de un
proyecto nunca se mezclan con las de otro (RN-16) sin llevar ningún índice aparte.

Registrar es una sola operación que valida en orden de lo más general a lo más particular
—integrante → historia → día → horas → actividad— y recién cuando pasó todo guarda el
registro, que es lo que hace que un rechazo deje el proyecto como estaba (RN-18).

Los totales **se calculan sumando los registros cada vez que se piden**, no se guardan: un
total que no se guarda no puede quedar desactualizado, y es lo que dice el contrato.

### Archivos afectados

| Archivo | Se crea o se modifica | Para qué |
|---|---|---|
| `internal/domain/effort.go` | se crea | el registro de esfuerzo, la operación de registrar y las tres consultas (todos los registros, horas de una historia, horas de un integrante) |
| `internal/domain/effort_test.go` | se crea | los tests de los escenarios BDD de la Fase 1, nombrados como cada escenario |
| `internal/domain/errors.go` | se modifica | los errores propios de esta historia: horas inválidas, día futuro, fecha faltante y actividad faltante |
| `internal/domain/project.go` | se modifica | la lista de registros de esfuerzo dentro del proyecto, al lado de los integrantes y el backlog |

### Decisiones

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| El día de hoy lo recibe la operación como dato de quien llama | que el dominio lea el reloj del sistema | decidido por el Agile Enabler el 20/10 y escrito en AGENTS.md: los tests no dependen del día en que se corren y el dominio sigue sin efectos externos. La validación del día futuro sigue en el dominio; lo que viene de afuera es solo qué día es hoy |
| Los totales se calculan sumando los registros en cada consulta | guardar un acumulado por historia y por integrante y actualizarlo al registrar | un acumulado hay que mantenerlo sincronizado y puede mentir; el volumen de un trabajo práctico no justifica el riesgo |
| Los totales los responde el proyecto | devolver la lista de registros y que cada interfaz sume | si la suma vive afuera del dominio, las RN-13 y RN-14 no quedan cubiertas por ningún test y cada interfaz la reimplementa distinto |
| Se guarda el nombre del integrante ya resuelto contra la lista del proyecto (RN-8) | guardar el texto tal como se escribió | la #16 agrupa las horas por nombre: dos formas de escribir el mismo nombre partirían el total de la misma persona en dos |
| Un error propio para cada rechazo (horas, día futuro, falta la fecha, falta la actividad) | un único error de "registro inválido" | la Fase 1 da cuatro mensajes distintos, y los tests distinguen cuál salió con `errors.Is` |
| Las horas se validan como "mayores que cero", y no como "no menores o iguales que cero" | comprobar que no sean cero ni negativas | las dos parecen lo mismo, pero solo la primera rechaza un valor que no es un número (`NaN`), que puede llegar desde la consola porque convertir el texto "NaN" a número no da error. Es el mismo criterio que usa la #16 para las horas estimadas (hallazgo de la revisión del PR #81) |
| El registro admite varias cargas del mismo integrante, historia y día (RN-11) | un único registro por esa combinación, que acumula o rechaza | ver el punto 3 de "Para decidir en la puerta": con la actividad adentro, acumular pierde información |

### Dependencias y restricciones

- **Depende de la #3, la #4 y la #6**, que ya están en `dev`: el proyecto, la búsqueda de un
  integrante por su nombre y la de una historia dentro del proyecto. Acá solo se usan.
- **No depende de la #8:** el contrato no le puso sprint al registro, así que esta historia
  se implementa completa sin que el `Sprint` exista. Cruzar horas con sprints es de la #16.
- **No se toca `member.go` ni `story.go`:** al integrante y a la historia solo se los busca.
- **`errors.go` es compartido:** los errores de esta historia se agregan ahí, y al traer
  `dev` hay que resolver ese conflicto a mano, como ya pasó en la #12.
- **El contrato no se toca:** ya tiene la actividad y las dos consultas de totales (punto 1
  de la puerta). Si al implementar hiciera falta cambiarlo, se avisa al grupo antes, como
  pide la regla 3 de `docs/diseno-dominio.md`.

> **Puerta 2** — ¿el diseño es viable y coherente con el resto del proyecto?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios: [pendiente]

## Fase 3 — Tareas

Cada tarea es un ciclo de TDD completo: el test que falla en un commit, la implementación que
lo hace pasar en otro. Están en el orden en que hay que hacerlas y cada una deja los tests en
verde. Siguen la convención de [AGENTS.md](../../../AGENTS.md): un test por escenario, con el
nombre del escenario y los errores verificados con `errors.Is`. Cada tarea de rechazo
comprueba además que el proyecto quedó sin registros, que es la otra mitad de la RN-18.

| # | Tarea | Escenario BDD | Reglas | Estado |
|---|---|---|---|---|
| 1 | Registrar un esfuerzo válido: queda guardado con sus cinco datos y aparece en los registros del proyecto | Registrar horas trabajadas en una historia | RN-1, RN-15 | pendiente |
| 2 | Las horas admiten fracciones | Registrar media hora | RN-2 | pendiente |
| 3 | Rechazar cero horas, horas negativas y un valor que no es un número, sin dejar rastro | Rechazar horas que no son mayores que cero | RN-3, RN-18 | pendiente |
| 4 | El día de mañana se rechaza; el de hoy y uno anterior al inicio del proyecto se aceptan | Rechazar horas de un dia que todavia no paso / Registrar horas de un dia anterior al inicio del proyecto | RN-4, RN-17 | pendiente |
| 5 | Rechazar un registro sin fecha | Rechazar un registro sin fecha | RN-5 | pendiente |
| 6 | La actividad es obligatoria, un texto de solo espacios cuenta como vacío y se guarda sin los espacios de los extremos | Rechazar un registro sin actividad realizada / Guardar la actividad sin los espacios de los extremos | RN-6 | pendiente |
| 7 | Quien registra tiene que ser integrante del proyecto; el nombre se busca ignorando mayúsculas y espacios, y se guarda como está cargado | Rechazar el registro de alguien que no es integrante del proyecto / Encontrar al integrante aunque el nombre venga escrito distinto | RN-7, RN-8 | pendiente |
| 8 | La historia tiene que pertenecer al proyecto, y una Terminada acepta registros | Rechazar el registro sobre una historia que no es del proyecto / Registrar horas sobre una historia terminada | RN-9, RN-10 | pendiente |
| 9 | El mismo integrante puede cargar dos veces la misma historia el mismo día, y las dos cargas quedan y suman | Dos cargas del mismo integrante, la misma historia y el mismo dia | RN-11, RN-12 | pendiente |
| 10 | Las horas de una historia son la suma de sus registros, y sin registros son cero | Consultar el total de horas de una historia / Una historia sin registros tiene cero horas | RN-13 | pendiente |
| 11 | Las horas de un integrante son la suma de sus registros, y sin registros son cero | Consultar el total de horas de un integrante / Un integrante sin registros tiene cero horas | RN-14 | pendiente |
| 12 | Los registros de un proyecto no suman en otro | Las horas cargadas en un proyecto no suman en otro | RN-16 | pendiente |
| 13 | Registrar o consultar en un proyecto que no existe informa que no se encontró el proyecto | Registrar esfuerzo en un proyecto que no existe | RN-19 | pendiente |

**Tarea sin escenario propio.** La RN-12 (el día se guarda sin la hora) no tiene escenario
suyo en la Fase 1: se verifica dentro de la tarea 9, cargando los dos registros del mismo día
con horas distintas y comprobando que los dos quedan guardados con el mismo día.

### Por qué este orden

- **La 1 va primero** porque es el camino feliz y deja la consulta de los registros, contra
  la que verifican todas las demás.
- **De la 2 a la 6** son las validaciones de los datos del propio registro, agregando un
  error por vez.
- **La 7 y la 8** son las que salen del registro hacia el proyecto: el integrante y la
  historia.
- **La 9 cierra el alta** y deja dos registros cargados, que es la base de los totales.
- **La 10, la 11 y la 12 son las consultas**: sin registros no hay nada que sumar.
- **La 13 va al final** porque el proyecto inexistente es el borde más externo.

### Cobertura

Las diecinueve reglas quedan cubiertas: RN-1 (1), RN-2 (2), RN-3 (3), RN-4 (4), RN-5 (5),
RN-6 (6), RN-7 (7), RN-8 (7), RN-9 (8), RN-10 (8), RN-11 (9), RN-12 (9), RN-13 (10),
RN-14 (11), RN-15 (1), RN-16 (12), RN-17 (4), RN-18 (3, 5, 6, 7, 8), RN-19 (13).

| Criterio de aceptación | Tareas |
|---|---|
| Se registra integrante, historia, día, horas y actividad, y queda guardado | 1 |
| El día de hoy se acepta | 1, 4 |
| Se aceptan fracciones de hora | 2 |
| No se pueden registrar cero horas ni un valor negativo | 3 |
| No se pueden cargar horas de un día posterior a hoy, y hoy sí se acepta | 4 |
| No se puede registrar sin día ni sin actividad, y la actividad se guarda sin los espacios de los extremos | 5, 6 |
| El integrante tiene que ser del proyecto, se encuentra escrito distinto y se guarda como está cargado | 7 |
| No se puede registrar sobre una historia ajena al proyecto, y sí sobre una Terminada | 8 |
| Dos cargas del mismo integrante, historia y día quedan y suman | 9 |
| Se consulta el total de horas de una historia, y sin registros da cero | 10 |
| Se consulta el total de horas de un integrante, y sin registros da cero | 11 |
| Se consultan todos los registros del proyecto en orden de carga | 1 |
| Las horas de un proyecto no suman en otro | 12 |
| Un registro rechazado no deja nada guardado | 3, 5, 6, 7, 8 |
| Registrar o consultar en un proyecto inexistente informa que no se encontró | 13 |

Los dieciocho escenarios BDD de la Fase 1 tienen su tarea, en la columna "Escenario BDD".

### Dependencias

- **Todas las tareas necesitan la #3, la #4 y la #6 en `dev`**, que ya lo están.
- **La tarea 8 necesita una historia Terminada**, y los cambios de estado son de la #9 y la
  #10, que todavía no existen. El test la prepara desde dentro del paquete, que puede fijar
  el estado directamente, igual que hizo la #12 en su tarea 9.
- **Ninguna tarea necesita la #8.** Esta historia se implementa completa sin que el `Sprint`
  exista.
- **La actividad realizada ya está en el contrato**, así que la tarea 1 puede guardar los
  cinco datos sin esperar nada.

### Entrega en dos pull requests

La #16 (métricas) necesita registros de esfuerzo para sus tareas 4 a 16, y en Go su test no
puede armar un `Effort` si el tipo todavía no existe en `dev`. Por eso esta historia se parte
en dos, que es el caso 1 de "Cuándo se puede partir una historia" en `CLAUDE.md` (alguien
está bloqueado esperándote):

1. **Primer PR (`Refs #14`):** esta spec con sus tres puertas y la **tarea 1**, que deja el
   `Effort`, el registro y la consulta de los registros. Con eso la #16 ya puede cargar
   esfuerzo en sus tests.
2. **Segundo PR (`Closes #14`):** las tareas 2 a 13.

> **Puerta 3** — ¿las tareas cubren todos los criterios de aceptación?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios: [pendiente]

## Fase 4 — Implementación

_Se completa durante la implementación._

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #14 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | `internal/domain/effort_test.go` |
| Código Go | `internal/domain/effort.go` y sus errores en `internal/domain/errors.go` |
