# Spec: Consultar sprints anteriores

- **Historia:** #11
- **Sprint:** 3
- **Responsable:** Agustín Salinas
- **Estado:** Borrador completo — pendiente de las tres puertas
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de la consigna y
  de los criterios de aceptación de la issue. El responsable lo revisa, lo corrige si hace
  falta y lo presenta a las puertas.
- **Depende de:** #8 (los sprints del proyecto y la consulta por número, en `dev`), #9 (las
  historias que se le asignaron a cada uno, en `dev`) y #10 (el estado Cerrado, la fecha de
  cierre y las historias completadas, todavía no en `dev`)
- **La necesitan:** #16 (recorre los sprints cerrados para medir), #17 (el dashboard los
  muestra), #18 (el reporte de sprint) y la interfaz

## Fase 1 — Requisitos

### Objetivo

Que el equipo pueda mirar los sprints que ya se cerraron —su meta, sus fechas, las historias
que se le comprometieron y cuáles se completaron— para comparar cómo viene el proyecto
sprint a sprint, sin poder cambiar nada de lo que ya pasó.

### Qué queda afuera

- **Los story points planificados y completados, la velocidad y el resto de los números.**
  Son la #16, que los calcula a partir de lo que esta historia entrega. Ver el punto 1 de
  "Para decidir en la puerta".
- **El sprint abierto**: consultarlo es la #8.
- **Modificar, reabrir o borrar un sprint cerrado.** No existe ninguna de esas operaciones.
- **Comparar dos sprints entre sí, graficar o mostrar tendencias**: eso es el dashboard
  (#17) y los reportes (#18).
- **Filtrar o buscar sprints** por fecha, por goal o por cantidad de historias.
- **Cómo se muestra en pantalla**: eso es de la interfaz.

### Entradas

| Dato | Obligatorio | Descripción |
|---|---|---|
| Proyecto | sí | de qué proyecto se consultan los sprints |
| Número de sprint | solo para el detalle | cuál de los sprints se quiere ver en detalle |

### Salidas

El listado de los sprints cerrados del proyecto, y el detalle de un sprint pedido por su
número. De cada sprint se ve su número, su Sprint Goal, sus fechas, su estado, su fecha de
cierre, las historias que se le asignaron y cuáles de ellas se completaron.

### Reglas de negocio

1. **[RN-1]** El sistema debe permitir listar los sprints cerrados de un proyecto, ordenados
   por número, del más viejo al más nuevo.
2. **[RN-2]** El sistema debe dejar fuera del listado al sprint abierto: solo se listan los
   cerrados.
3. **[RN-3]** Si el proyecto no tiene ningún sprint cerrado, el sistema debe devolver un
   listado vacío, y eso no es un error.
4. **[RN-4]** El sistema debe mostrar de cada sprint cerrado: su número, su Sprint Goal, su
   fecha de inicio, su fecha de fin, su fecha de cierre, las historias que se le asignaron y
   cuáles de ellas se completaron.
5. **[RN-5]** El sistema debe permitir consultar el detalle de un sprint indicando su número
   dentro del proyecto.
6. **[RN-6]** Si se pide un sprint con un número que no existe en el proyecto, el sistema
   debe informar que no se encontró el sprint.
7. **[RN-7]** Mientras un sprint esté abierto, el sistema debe entregar su detalle igual que
   el de uno cerrado, con la fecha de cierre vacía.
8. **[RN-8]** El sistema debe entregar los sprints solo para leerlos: consultarlos no cambia
   nada, y lo que se recibe no permite modificar el sprint ni sus historias.
9. **[RN-9]** El sistema debe tratar cada proyecto por separado: el listado y el detalle solo
   traen sprints de ese proyecto.
10. **[RN-10]** Si se pide el listado o el detalle de un proyecto que no existe, el sistema
    debe informar que no se encontró el proyecto.
11. **[RN-11]** El sistema debe conservar los sprints cerrados mientras el programa esté
    corriendo: cerrar un sprint nuevo no cambia ni borra los anteriores.

### Casos límite

- Proyecto sin ningún sprint: el listado está vacío y no es un error.
- Proyecto con un solo sprint y abierto: el listado está vacío.
- Proyecto con un sprint cerrado y uno abierto: el listado trae solo el cerrado.
- Sprint cerrado sin ninguna historia asignada: aparece igual, con sus listas vacías.
- Sprint cerrado con todas las historias completadas y otro con ninguna: los dos aparecen,
  cada uno con lo suyo.
- Pedir el detalle del sprint número 0, o de un número mayor al último: no se encontró.
- Pedir el detalle del sprint que está abierto: se devuelve, con la fecha de cierre vacía.
- Dos proyectos con sprints cerrados: cada listado trae solo los suyos, aunque los números
  se repitan entre proyectos.

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| El sprint pedido no existe en el proyecto | informa que no se encontró el sprint |
| El proyecto no existe | informa que no se encontró el proyecto |
| El proyecto no tiene sprints cerrados | devuelve un listado vacío: **no es un error** |

### Criterios de aceptación

- [ ] Se listan los sprints cerrados de un proyecto con su número, su goal, sus fechas, su
      fecha de cierre y sus historias.
- [ ] De cada sprint cerrado se ve qué historias se le asignaron y cuáles se completaron.
- [ ] El sprint abierto no aparece en el listado.
- [ ] El listado viene ordenado por número, del más viejo al más nuevo.
- [ ] Un proyecto sin sprints cerrados devuelve un listado vacío y eso no es un error.
- [ ] Se puede pedir el detalle de un sprint por su número.
- [ ] Pedir un sprint con un número que no existe informa que no se encontró el sprint.
- [ ] Consultar un sprint no lo modifica, y lo que se recibe no permite modificarlo.
- [ ] El listado de un proyecto no trae sprints de otro.
- [ ] Consultar en un proyecto que no existe informa que no se encontró el proyecto.

### Escenarios BDD

```gherkin
Escenario: Listar los sprints cerrados de un proyecto
  Dado un proyecto con el sprint 1 cerrado y el sprint 2 cerrado
  Cuando pido los sprints cerrados del proyecto
  Entonces recibo los sprints 1 y 2, en ese orden
  Y de cada uno veo su goal, sus fechas, su fecha de cierre y sus historias

Escenario: El sprint abierto no aparece en el listado
  Dado un proyecto con el sprint 1 cerrado y el sprint 2 abierto
  Cuando pido los sprints cerrados del proyecto
  Entonces recibo solamente el sprint 1

Esquema del escenario: Un proyecto sin sprints cerrados devuelve un listado vacio
  Dado un proyecto <situacion>
  Cuando pido los sprints cerrados del proyecto
  Entonces el listado esta vacio
  Y eso no es un error

  Ejemplos:
    | situacion               |
    | "sin ningun sprint"     |
    | "con un sprint abierto" |

Escenario: Ver que historias se completaron en un sprint cerrado
  Dado un proyecto cuyo sprint 1 se cerro con las historias 1, 2 y 3 asignadas y la 1 completada
  Cuando pido el detalle del sprint 1
  Entonces veo que se le asignaron las historias 1, 2 y 3
  Y veo que se completo la historia 1

Escenario: Un sprint cerrado sin historias se consulta igual
  Dado un proyecto cuyo sprint 1 se cerro sin ninguna historia asignada
  Cuando pido el detalle del sprint 1
  Entonces recibo el sprint 1 con su goal y sus fechas
  Y sus listas de historias estan vacias

Escenario: Pedir el detalle de un sprint por su numero
  Dado un proyecto con el sprint 1 cerrado y el sprint 2 cerrado
  Cuando pido el detalle del sprint 2
  Entonces recibo el sprint 2 con su goal, sus fechas y su fecha de cierre

Escenario: Pedir el detalle del sprint abierto
  Dado un proyecto con el sprint 1 abierto
  Cuando pido el detalle del sprint 1
  Entonces recibo el sprint 1 en estado "Abierto"
  Y su fecha de cierre esta vacia

Esquema del escenario: Pedir un sprint que no existe
  Dado un proyecto con el sprint 1 cerrado
  Cuando pido el detalle del sprint <numero>
  Entonces se informa que no se encontro el sprint

  Ejemplos:
    | numero |
    | 0      |
    | 2      |
    | -1     |

Escenario: Consultar un sprint no lo modifica
  Dado un proyecto cuyo sprint 1 se cerro con 2 historias asignadas
  Cuando pido el detalle del sprint 1 y modifico la lista de historias que recibo
  Entonces el sprint 1 sigue teniendo sus 2 historias asignadas

Escenario: Cada proyecto lista solamente sus sprints
  Dado un proyecto A con su sprint 1 cerrado y un proyecto B con su sprint 1 cerrado
  Cuando pido los sprints cerrados del proyecto A
  Entonces recibo un solo sprint
  Y es el sprint 1 del proyecto A

Escenario: Cerrar un sprint nuevo no cambia los anteriores
  Dado un proyecto cuyo sprint 1 se cerro con la historia 1 completada
  Y el sprint 2 abierto con la historia 2 asignada
  Cuando cierro el sprint 2 sin historias completadas
  Entonces el sprint 1 sigue registrando a la historia 1 como completada

Escenario: Consultar los sprints de un proyecto que no existe
  Dado que no existe ningun proyecto con identificador 99
  Cuando pido los sprints cerrados del proyecto 99
  Entonces se informa que no se encontro el proyecto
```

**Para decidir en la puerta:**

1. **Los story points planificados y completados no se calculan acá.** La issue los pide en
   su segundo criterio, pero son dos de las once métricas de la #16, y el equipo decidió que
   las métricas vivan en una sola historia. Esta historia entrega lo que hace falta para
   calcularlos: las historias que se le asignaron a cada sprint y cuáles se completaron. Si
   el equipo prefiere ver los números sin esperar a la #16, hay que mover ese cálculo acá y
   la #16 queda sin su primera mitad.
2. **El detalle también se entrega del sprint abierto (RN-7).** Es el mismo dato y negarlo
   obligaría a la interfaz a usar dos caminos distintos para mostrar lo mismo. Lo que sí es
   solo de los cerrados es el **listado** (RN-2).
3. **"Sin poder editarlo" no se resuelve con una regla nueva.** Ninguna operación del sistema
   modifica un sprint cerrado: la #9 solo toca el abierto y la #10 rechaza cerrar dos veces.
   Esta historia lo verifica, no lo implementa, y por eso la RN-8 habla de que lo que se
   recibe tampoco permita cambiarlo de costado.
4. **El listado va del más viejo al más nuevo.** Para comparar sprint a sprint es el orden
   natural; una pantalla que quiera el último primero puede darlo vuelta. Si el equipo
   prefiere el orden inverso, cambia la RN-1 y la #17 lo hereda.
5. **No hay paginado ni filtros.** Son cuatro o cinco sprints por proyecto.
6. **La consulta de un sprint por su número ya la entregó la #8.** La RN-5 y la RN-6 quedan
   acá porque son criterios de esta issue, pero no se vuelven a construir. Lo que esta
   historia suma al detalle es lo del cierre, que trae la #10.

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios: [pendiente]

## Fase 2 — Diseño

Las decisiones transversales del proyecto están en [AGENTS.md](../../../AGENTS.md) y no se
repiten acá. Los datos del sprint están en [el contrato del dominio](../../diseno-dominio.md),
incluido el dato de las historias completadas que agrega la #10. Esta fase se apoya en la
Fase 2 de la #8: los sprints viven dentro del proyecto, en el orden en que se crearon.

### Enfoque

Esta historia **no cambia nada**: son dos consultas sobre los sprints que el proyecto ya
tiene. No hay que normalizar, ni validar, ni aplicar: solo encontrar y entregar. **La segunda
ya existe**: la #8 la adelantó (su RN-15), así que la única operación nueva es el listado de
cerrados.

- El **listado de cerrados** recorre los sprints del proyecto y deja los que están en estado
  Cerrado. Como la #8 los agrega en orden y el número nunca retrocede, el orden de la RN-1
  sale solo: no hay que ordenar nada.
- El **detalle** es la consulta por número de la #8: busca el sprint dentro del proyecto y
  lo devuelve, esté abierto o cerrado (RN-7). Si no está, informa que no se encontró (RN-6).
  Esta historia no la modifica; sus tests la recorren con los datos del cierre.

Lo que se entrega es el sprint guardado, no una copia, como ya decidió el proyecto para
todas sus entidades: los datos están ocultos fuera del paquete y solo se cambian con
operaciones que validan, y ninguna operación modifica un sprint cerrado. Las **listas de
historias sí se entregan copiadas**, igual que hace el Product Backlog con las suyas, para
que quien las recibe no pueda cambiar el registro del sprint sin querer (RN-8).

### Archivos afectados

| Archivo | Se crea o se modifica | Para qué |
|---|---|---|
| `internal/domain/sprint_query.go` | se crea | el listado de sprints cerrados |
| `internal/domain/sprint_query_test.go` | se crea | los tests de los escenarios BDD de la Fase 1, nombrados como cada escenario |

`errors.go` no se toca: el error de sprint no encontrado ya lo declaró la #8.

### Decisiones

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| El listado de cerrados es una operación del dominio | que el listado completo lo filtre quien lo muestra | las reglas viven en el dominio (AGENTS.md); con el filtro afuera, la RN-2 no tendría test y cada interfaz podría filtrar distinto |
| El orden sale del orden en que están guardados los sprints | ordenar explícitamente por número al consultar | la #8 agrega siempre al final con el número siguiente, así que la lista ya está ordenada; ordenar otra vez es trabajo que esconde esa garantía en lugar de apoyarse en ella |
| Los sprints se buscan recorriendo la lista | llevar además un índice por número para encontrarlo directo | son cuatro o cinco sprints por proyecto; un índice es un segundo lugar donde guardar lo mismo, que se puede desincronizar |
| Las listas de historias se entregan copiadas, el sprint no | copiar el sprint entero, o no copiar nada | el sprint tiene sus datos ocultos y ninguna operación lo modifica una vez cerrado, así que copiarlo no compra nada; las listas sí se podrían modificar desde afuera del paquete y por eso se copian, igual que ya hace el Product Backlog |
| Lo de esta historia va en `sprint_query.go` | agregarlo a `sprint.go` | mismo criterio que la #9 y la #10: cuatro historias tocan el sprint y no conviene que se pisen en el mismo archivo |
| El detalle reutiliza la consulta por número de la #8, sin cambiarla | escribir una consulta propia del detalle | la #8 la adelantó con su error porque la #15 la necesitaba antes; dos consultas para lo mismo podrían responder distinto |

### Dependencias y restricciones

- **Depende de la #8**, que ya está en `dev` (los sprints dentro del proyecto, el listado
  completo `p.Sprints()` y la consulta por número `p.Sprint(number)` con su error), y de la
  **#10** (el estado Cerrado, la fecha de cierre y el dato de las historias completadas),
  que todavía no. Mientras tanto los tests fijan el estado Cerrado desde dentro del paquete,
  como ya hace el test de la consulta por número de la #8.
- **El contrato del dominio nombra el listado de cerrados** (`p.ClosedSprints()`); es lo
  único que esta historia crea.
- **No se modifica nada.** Ninguna operación de esta historia cambia un sprint, una historia
  ni el proyecto.
- **No se calcula ninguna métrica**: son de la #16, que consume estas consultas.

> **Puerta 2** — ¿el diseño es viable y coherente con el resto del proyecto?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios: [pendiente]

## Fase 3 — Tareas

Cada tarea es un ciclo de TDD completo: el test que falla en un commit, la implementación que
lo hace pasar en otro. Van en el orden en que hay que hacerlas, y cada una deja el código
compilando y los tests en verde.

Los tests siguen la convención de [AGENTS.md](../../../AGENTS.md): un test por escenario, con
el nombre del escenario.

| # | Tarea | Escenario BDD | Reglas | Estado |
|---|---|---|---|---|
| 1 | Un proyecto sin sprints cerrados devuelve un listado vacío, con y sin sprint abierto | Un proyecto sin sprints cerrados devuelve un listado vacio | RN-3 | pendiente |
| 2 | Listar los sprints cerrados con su goal, sus fechas, su fecha de cierre y sus historias, ordenados por número | Listar los sprints cerrados de un proyecto | RN-1, RN-4 | pendiente |
| 3 | El sprint abierto no aparece en el listado | El sprint abierto no aparece en el listado | RN-2 | pendiente |
| 4 | Pedir el detalle de un sprint por su número | Pedir el detalle de un sprint por su numero | RN-5 | pendiente |
| 5 | El detalle muestra qué historias se asignaron y cuáles se completaron | Ver que historias se completaron en un sprint cerrado | RN-4 | pendiente |
| 6 | Un sprint cerrado sin historias se consulta igual, con las listas vacías | Un sprint cerrado sin historias se consulta igual | RN-4 | pendiente |
| 7 | Pedir un número que no existe informa que no se encontró el sprint | Pedir un sprint que no existe | RN-6 | pendiente |
| 8 | El detalle del sprint abierto se entrega con la fecha de cierre vacía | Pedir el detalle del sprint abierto | RN-7 | pendiente |
| 9 | Modificar las listas de historias que devuelve la consulta no cambia el sprint | Consultar un sprint no lo modifica | RN-8 | pendiente |
| 10 | El listado de un proyecto no trae los sprints de otro | Cada proyecto lista solamente sus sprints | RN-9 | pendiente |
| 11 | Cerrar un sprint nuevo no cambia lo registrado en los anteriores | Cerrar un sprint nuevo no cambia los anteriores | RN-11 | pendiente |
| 12 | Consultar en un proyecto que no existe informa que no se encontró el proyecto | Consultar los sprints de un proyecto que no existe | RN-10 | pendiente |

### Por qué este orden

- **La 1 va primero** porque el listado vacío se puede escribir sin que exista todavía
  ningún sprint cerrado, y deja la operación creada para las que siguen.
- **La 2 y la 3** llenan el listado y después le sacan el sprint abierto: al revés, el filtro
  no tendría nada contra qué probarse.
- **De la 4 a la 8** va el detalle, del camino feliz al rechazo y al caso del sprint abierto.
- **La 9, la 10 y la 11 van al final** porque son las garantías de que consultar no rompe
  nada: ni el sprint que se devuelve, ni otro proyecto, ni los sprints anteriores.
- **La 12 cierra** con el borde de afuera, el proyecto que no existe.
- **La 4, la 7, la 9, la 10 y la 12 se espera que pasen sin código nuevo.** La consulta por
  número, su error y la copia de las historias asignadas ya los trae la #8; un proyecto que
  no existe lo informa el registro antes de llegar a sus sprints, igual que en la #8; y una
  vez que existe el listado, cada proyecto solo tiene los suyos. Cada una va en un solo
  commit de test con "sin RED" en el asunto, explicando por qué: forzar un RED sería simular
  TDD.
- **Mientras la #10 no esté en `dev`, el orden real cambia.** La 2, la 5, la 6, la 8 y la 11
  necesitan la fecha de cierre o las historias completadas (ver Dependencias), así que se
  hacen primero la 1, 3, 4, 7, 9, 10 y 12, y después la 2, 5, 6, 8 y 11. La 3 puede ir antes
  que la 2 porque su escenario ya trae un sprint cerrado contra el cual probar el filtro.

### Cobertura

Las once reglas quedan cubiertas: RN-1 (2), RN-2 (3), RN-3 (1), RN-4 (2, 5, 6), RN-5 (4),
RN-6 (7), RN-7 (8), RN-8 (9), RN-9 (10), RN-10 (12), RN-11 (11).

| Criterio de aceptación | Tareas |
|---|---|
| Se listan los sprints cerrados con su goal, sus fechas, su cierre y sus historias | 2 |
| De cada sprint se ve qué se asignó y qué se completó | 5, 6 |
| El sprint abierto no aparece en el listado | 3 |
| El listado viene ordenado por número | 2 |
| Un proyecto sin sprints cerrados devuelve un listado vacío | 1 |
| Se puede pedir el detalle de un sprint por su número | 4, 8 |
| Un número que no existe informa que no se encontró el sprint | 7 |
| Consultar no modifica el sprint | 9, 11 |
| El listado de un proyecto no trae sprints de otro | 10 |
| Consultar en un proyecto que no existe informa que no se encontró | 12 |

Los doce escenarios de la Fase 1 tienen su tarea en la columna "Escenario BDD".

### Dependencias

- **Todas las tareas necesitan los sprints de la #8**, que ya están en `dev`, y salvo la 1 y
  la 8, **sprints cerrados**. El estado Cerrado los tests lo fijan dentro del paquete.
- **Las tareas 2, 5, 6, 8 y 11 necesitan leer la fecha de cierre o las historias
  completadas**, que trae la #10. Esas cinco esperan a que la #10 esté en `dev`; las otras
  siete se pueden hacer antes.

> **Puerta 3** — ¿las tareas cubren todos los criterios de aceptación?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios: [pendiente]

## Fase 4 — Implementación

_Se completa durante la implementación._

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #11 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | `internal/domain/sprint_query_test.go` |
| Código Go | `internal/domain/sprint_query.go`; el detalle es `p.Sprint(number)` de `internal/domain/sprint.go` (#8) |
