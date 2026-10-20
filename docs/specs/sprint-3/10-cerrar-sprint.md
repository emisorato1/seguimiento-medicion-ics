# Spec: Cerrar un sprint

- **Historia:** #10
- **Sprint:** 3
- **Responsable:** Emiliano Sorato
- **Estado:** Borrador completo — pendiente de las tres puertas
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de la consigna y
  de los criterios de aceptación de la issue. El responsable lo revisa, lo corrige si hace
  falta y lo presenta a las puertas.
- **Depende de:** #8 (el sprint abierto), #9 (las historias que se le asignaron) y #6 (los
  estados de una historia); por ellas, también de la #3
- **La necesitan:** #11 (lista los sprints cerrados y su detalle), #16 (calcula velocidad,
  puntos completados y porcentaje de historias completas), y detrás de la #16 la #17 y la #18

## Fase 1 — Requisitos

### Objetivo

Que al terminar un sprint el equipo pueda cerrarlo dejando registrado qué historias se
completaron y cuáles no, para que las que quedaron vuelvan al Product Backlog y el sprint
quede como registro de lo que efectivamente pasó.

### Qué queda afuera

- **Calcular los story points completados, la velocidad, la desviación de horas y el
  porcentaje de historias completas.** Son la #16. El cierre **deja los datos**: qué se
  comprometió y qué se completó. Los números los calcula la #16 sobre esos datos.
- **Crear el sprint siguiente**, que es la #8. Cerrar deja al proyecto sin sprint abierto, y
  recién entonces se puede crear otro.
- **Consultar los sprints cerrados** y su detalle, que es la #11.
- **Asignar o quitar historias**, que es la #9. Al cerrar ya no se cambia el Sprint Backlog.
- **Reabrir un sprint cerrado.** No existe esa operación.
- **La review y la retrospectiva**: son reuniones del equipo y quedan en las actas, no en el
  sistema.
- **El esfuerzo (#14) y los defectos (#15)**: se registran por su cuenta, el cierre no los
  toca.

### Entradas

| Dato | Obligatorio | Descripción |
|---|---|---|
| Proyecto | sí | en qué proyecto se cierra el sprint |
| Historias completadas | no | cuáles de las historias del sprint se completaron; puede no haber ninguna |
| Fecha de cierre | sí | el día en que el sprint se cierra de verdad |

**Qué sprint se cierra no es una entrada**: siempre es el sprint abierto del proyecto, que
es el único que se puede cerrar (#8, RN-8).

### Salidas

El sprint queda Cerrado, con su fecha de cierre y con el registro de cuáles de sus historias
se completaron. Las completadas quedan Terminadas y las demás vuelven al Product Backlog
como Pendientes. El proyecto deja de tener sprint abierto.

### Reglas de negocio

1. **[RN-1]** Cuando se cierra un sprint, el sistema debe dejarlo en estado **Cerrado** y
   registrar la fecha de cierre indicada.
2. **[RN-2]** Cuando se cierra un sprint, el sistema debe registrar **cuáles de sus historias
   se completaron**, aparte de la lista de historias que se le habían asignado.
3. **[RN-3]** Cuando se cierra un sprint, el sistema debe dejar en estado **Terminada** cada
   historia indicada como completada.
4. **[RN-4]** Cuando se cierra un sprint, el sistema debe devolver al Product Backlog, en
   estado **Pendiente**, cada historia asignada que no se indicó como completada.
5. **[RN-5]** El sistema debe conservar la lista de historias asignadas al sprint tal como
   estaba al cerrarlo: una historia que vuelve al Product Backlog sigue figurando entre las
   que se comprometieron en ese sprint.
6. **[RN-6]** Si el proyecto no tiene ningún sprint abierto, el sistema debe rechazar el
   cierre e informar que el proyecto no tiene un sprint abierto.
7. **[RN-7]** Si alguna de las historias indicadas como completadas no está asignada al
   sprint, el sistema debe rechazar el cierre e informar que esa historia no está en el
   sprint.
8. **[RN-8]** Si la misma historia se indica más de una vez como completada, el sistema debe
   contarla una sola vez y no rechazar el cierre.
9. **[RN-9]** Si no se indica ninguna historia completada, el sistema debe cerrar el sprint
   igual, con todas sus historias de vuelta en el Product Backlog.
10. **[RN-10]** Si falta la fecha de cierre, el sistema debe rechazar el cierre e informar
    que la fecha de cierre es obligatoria.
11. **[RN-11]** El sistema debe aceptar la fecha de cierre tal como se indique, sin
    compararla con las fechas del sprint ni con la fecha de hoy: un sprint se puede cerrar
    antes o después de su fecha de fin.
12. **[RN-12]** Mientras un sprint esté cerrado, el sistema debe rechazar cualquier intento
    de volver a cerrarlo o de cambiar lo que quedó registrado en él.
13. **[RN-13]** Si un cierre se rechaza, el sistema debe dejar el sprint y todas sus
    historias exactamente como estaban.
14. **[RN-14]** Cuando se cierra el sprint, el sistema debe dejar al proyecto sin sprint
    abierto, de modo que se pueda crear el siguiente.
15. **[RN-15]** Si se intenta cerrar el sprint de un proyecto que no existe, el sistema debe
    informar que no se encontró el proyecto y no cerrar nada.
16. **[RN-16]** Cuando se cierra el sprint, el sistema debe registrar los story points que
    tenía en ese momento cada historia asignada, y esos valores no deben cambiar después,
    aunque una de las historias vuelva al Product Backlog y se reestime.

Los estados de una historia —Pendiente, En curso y Terminada— son los que define la
historia #6. Esta historia usa las dos transiciones que faltaban: de En curso a Terminada y
de En curso de vuelta a Pendiente. Una vez Terminada, la historia ya no se modifica ni se
reestima (#6 RN-10, #12 RN-8).

### Casos límite

- Sprint abierto sin ninguna historia asignada: se cierra igual, sin nada completado.
- Todas las historias completadas: ninguna vuelve al Product Backlog.
- Ninguna completada: todas vuelven al Product Backlog.
- Una sola historia en el sprint, completada: el sprint queda con esa historia comprometida
  y esa misma completada.
- La misma historia indicada dos veces como completada: cuenta una sola vez.
- Una historia del proyecto que nunca se asignó al sprint, indicada como completada: se
  rechaza el cierre entero.
- Una historia asignada que ya estaba Terminada: no puede pasar, porque una historia
  Terminada no se asigna (#9, RN-6).
- Fecha de cierre anterior a la fecha de inicio del sprint: se acepta, el sistema no compara
  las fechas (RN-11).
- Cerrar dos veces seguidas: el segundo cierre se rechaza, porque ya no hay sprint abierto.
- Después de cerrar, crear un sprint nuevo: funciona, y recibe el número siguiente.

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| El proyecto no tiene sprint abierto, o ya se cerró | rechaza e informa que no hay un sprint abierto |
| Una historia indicada como completada no está en el sprint | rechaza e informa que esa historia no está en el sprint, y no cierra nada |
| Falta la fecha de cierre | rechaza e informa que la fecha de cierre es obligatoria |
| El proyecto no existe | informa que no se encontró el proyecto y no cierra nada |

### Criterios de aceptación

- [ ] Se puede cerrar el sprint abierto indicando qué historias se completaron, y queda en
      estado Cerrado con su fecha de cierre.
- [ ] Las historias indicadas como completadas quedan Terminadas.
- [ ] Las historias del sprint que no se indicaron quedan Pendientes, de vuelta en el
      Product Backlog.
- [ ] El sprint cerrado conserva la lista de las historias que se le habían asignado,
      incluidas las que volvieron al Product Backlog.
- [ ] El sprint cerrado deja registrado cuáles de sus historias se completaron.
- [ ] Se puede cerrar un sprint sin ninguna historia completada.
- [ ] Se puede cerrar un sprint que no tiene ninguna historia asignada.
- [ ] Indicar dos veces la misma historia no es un error y no la cuenta dos veces.
- [ ] Indicar como completada una historia que no está en el sprint rechaza el cierre y no
      cambia ninguna historia.
- [ ] No se puede cerrar sin fecha de cierre.
- [ ] No se puede cerrar dos veces el mismo sprint.
- [ ] Después de cerrar, el proyecto no tiene sprint abierto y se puede crear el siguiente.
- [ ] Cerrar el sprint de un proyecto que no existe informa que no se encontró el proyecto.
- [ ] Los story points que el sprint registró al cerrar no cambian si después se reestima una
      de sus historias.

### Escenarios BDD

```gherkin
Escenario: Cerrar el sprint registrando que se completo y que no
  Dado un proyecto con el sprint 1 abierto y las historias 1, 2 y 3 asignadas
  Cuando cierro el sprint el 25/10/2026 indicando que se completaron las historias 1 y 2
  Entonces el sprint 1 queda en estado "Cerrado"
  Y el sprint 1 queda con fecha de cierre 25/10/2026
  Y el sprint 1 registra como completadas las historias 1 y 2
  Y las historias 1 y 2 quedan en estado "Terminada"
  Y la historia 3 queda en estado "Pendiente"

Escenario: El sprint cerrado conserva lo que se le habia comprometido
  Dado un proyecto con el sprint 1 abierto y las historias 1, 2 y 3 asignadas
  Cuando cierro el sprint indicando que se completo solo la historia 1
  Entonces el sprint 1 sigue teniendo asignadas las historias 1, 2 y 3

Esquema del escenario: Cerrar con distinta cantidad de historias completadas
  Dado un proyecto con el sprint 1 abierto y 3 historias asignadas
  Cuando cierro el sprint indicando <completadas> historias completadas
  Entonces quedan <terminadas> historias en estado "Terminada"
  Y quedan <pendientes> historias en estado "Pendiente"
  Y el sprint 1 queda en estado "Cerrado"

  Ejemplos:
    | completadas | terminadas | pendientes |
    | 0           | 0          | 3          |
    | 1           | 1          | 2          |
    | 3           | 3          | 0          |

Escenario: Cerrar un sprint sin ninguna historia asignada
  Dado un proyecto con el sprint 1 abierto y sin historias asignadas
  Cuando cierro el sprint sin indicar historias completadas
  Entonces el sprint 1 queda en estado "Cerrado"
  Y el sprint 1 no registra ninguna historia completada

Escenario: La misma historia indicada dos veces cuenta una sola vez
  Dado un proyecto con el sprint 1 abierto y la historia 1 asignada
  Cuando cierro el sprint indicando dos veces la historia 1 como completada
  Entonces el sprint 1 registra una sola historia completada
  Y la historia 1 queda en estado "Terminada"

Escenario: Rechazar el cierre si una historia completada no esta en el sprint
  Dado un proyecto con el sprint 1 abierto, la historia 1 asignada y la historia 2 sin asignar
  Cuando intento cerrar el sprint indicando que se completaron las historias 1 y 2
  Entonces la operacion se rechaza
  Y se informa que la historia 2 no esta en el sprint
  Y el sprint 1 sigue en estado "Abierto"
  Y la historia 1 sigue en estado "En curso"

Escenario: Rechazar el cierre sin fecha de cierre
  Dado un proyecto con el sprint 1 abierto
  Cuando intento cerrar el sprint sin indicar la fecha de cierre
  Entonces la operacion se rechaza
  Y se informa que la fecha de cierre es obligatoria
  Y el sprint 1 sigue en estado "Abierto"

Escenario: Rechazar el cierre si no hay sprint abierto
  Dado un proyecto sin ningun sprint abierto
  Cuando intento cerrar el sprint
  Entonces la operacion se rechaza
  Y se informa que el proyecto no tiene un sprint abierto

Escenario: No se puede cerrar dos veces el mismo sprint
  Dado un proyecto cuyo sprint 1 ya se cerro
  Cuando intento cerrar el sprint otra vez
  Entonces la operacion se rechaza
  Y se informa que el proyecto no tiene un sprint abierto
  Y el sprint 1 conserva su fecha de cierre y sus historias completadas

Escenario: Despues de cerrar se puede crear el sprint siguiente
  Dado un proyecto con el sprint 1 abierto
  Cuando cierro el sprint 1
  Entonces el proyecto no tiene sprint abierto
  Y puedo crear el sprint 2

Escenario: Una fecha de cierre distinta de la fecha de fin se acepta
  Dado un proyecto con el sprint 1 abierto, con fin el 25/10/2026
  Cuando cierro el sprint el 23/10/2026
  Entonces el sprint 1 queda con fecha de cierre 23/10/2026

Escenario: Los puntos registrados al cerrar no cambian si se reestima la historia
  Dado un sprint abierto con la historia 3 estimada en 5 story points
  Y cierro el sprint sin indicar la historia 3 como completada
  Cuando reestimo la historia 3 en 8 story points
  Entonces el sprint cerrado sigue registrando 5 story points para la historia 3

Escenario: Cerrar el sprint de un proyecto que no existe
  Dado que no existe ningun proyecto con identificador 99
  Cuando intento cerrar el sprint del proyecto 99
  Entonces se informa que no se encontro el proyecto
```

**Para decidir en la puerta:**

1. **Falta un dato en el contrato del dominio: dónde queda registrado qué historias se
   completaron (RN-2).** Hoy el sprint guarda las historias que se le asignaron y la fecha
   de cierre, nada más. No alcanza con mirar el estado de cada historia: una historia que no
   se completó vuelve al Product Backlog y se puede terminar en el sprint siguiente, y
   entonces figuraría como completada **en los dos sprints**, inflando la velocidad del
   primero. Hace falta agregarle al sprint la lista de historias completadas, avisando al
   grupo antes de implementar, como pide la regla 3 del contrato. **Sin ese dato, la #16 no
   puede calcular ni los puntos completados ni la velocidad por sprint.**
2. **"Al cerrarlo quedan calculados los story points completados" se interpretó como "quedan
   registrados los datos para calcularlos".** El cálculo es de la #16, que es donde viven
   todas las métricas. Si el equipo prefiere guardar el número en el sprint, hay que
   agregarlo también al contrato y decidir qué pasa si alguien reestima una historia después
   del cierre.
3. **El cierre recibe una sola lista: las historias completadas.** Las demás son, por
   definición, las que no se completaron. Recibir las dos listas permitiría que se
   contradigan o que quede una historia en ninguna de las dos.
4. **La fecha de cierre es una entrada y no la fecha de hoy.** Es lo mismo que hizo la #8 con
   las fechas del sprint, y es lo que permite escribir tests que no dependan del día en que
   se corren. Si el equipo prefiere que la ponga el sistema, cambia la RN-10 y hay que
   decidir cómo se prueba.
5. **La fecha de cierre no se compara con nada (RN-11)**, ni siquiera con el inicio del
   sprint. Es el mismo criterio de la RN-14 de la #8. Se acepta incluso una fecha anterior al
   inicio, que es raro pero no rompe nada de lo que se mide; si el equipo quiere, se agrega
   la validación contra la fecha de inicio.
6. **Cerrar no exige que el sprint esté "terminado".** No hay ningún control de que no queden
   historias en curso: justamente lo normal es que queden, y ese es el dato que importa.
7. **Una historia que vuelve al Product Backlog sigue figurando entre las del sprint
   (RN-5).** Es lo que permite que la #16 sepa cuánto se había comprometido. Si se la sacara,
   todos los sprints cerrados parecerían haber completado el 100%.
8. **No existe reabrir un sprint.** Si el equipo cierra con los datos equivocados, hoy no hay
   forma de corregirlo. Agregar una corrección implica decidir qué pasa con las historias que
   ya volvieron al backlog y se reasignaron, así que se deja afuera a propósito.

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios: [pendiente]

## Fase 2 — Diseño

Las decisiones transversales del proyecto están en [AGENTS.md](../../../AGENTS.md) y no se
repiten acá. Los datos del sprint están en [el contrato del dominio](../../diseno-dominio.md).
Esta fase se apoya en la Fase 2 de la #8 (el sprint dentro del proyecto, su estado y la
consulta del sprint abierto) y en la de la #9 (la lista de historias asignadas).

### Enfoque

Cerrar es **la única transición de estado que tiene un sprint**, y es una sola operación: o
pasa todo, o no pasa nada. Se le pide al proyecto, porque el sprint que se cierra es siempre
el abierto y eso lo sabe el proyecto, y porque hay que cambiarles el estado a historias que
viven en el Product Backlog.

El orden es el de siempre, de lo más general a lo más particular:

1. **Hay un sprint abierto** (RN-6). Esto cubre también el intento de cerrar dos veces
   (RN-12): después del primer cierre ya no hay ninguno abierto.
2. **La fecha de cierre está** (RN-10), normalizada a día como el resto de las fechas.
3. **Todas las historias indicadas pertenecen al sprint** (RN-7). Se controla **la lista
   entera antes de tocar una sola historia**: si no, un cierre rechazado dejaría la mitad de
   las historias Terminadas y la otra mitad sin cambiar.
4. **Aplicar**: se arma la lista de completadas sin repetidos (RN-8), se les cambia el estado
   a las historias —Terminada a las completadas, Pendiente a las demás—, y se marca el sprint
   como Cerrado con su fecha.

La lista de historias asignadas **no se toca**: "volver al Product Backlog" (RN-4) es un
cambio de estado de la historia, no una baja del registro del sprint. El sprint queda como
la foto de lo que se comprometió y de lo que se logró, que es exactamente lo que la #16 y la
#11 necesitan leer después.

### Archivos afectados

| Archivo | Se crea o se modifica | Para qué |
|---|---|---|
| `internal/domain/sprint_close.go` | se crea | la operación de cierre y la lectura de las historias completadas de un sprint |
| `internal/domain/sprint_close_test.go` | se crea | los tests de los escenarios BDD de la Fase 1, nombrados como cada escenario |
| `internal/domain/sprint.go` | se modifica | el sprint pasa a tener el dato de las historias completadas |
| `internal/domain/errors.go` | se modifica | el error de la fecha de cierre obligatoria |
| `docs/diseno-dominio.md` | se modifica | agregar al contrato el dato de las historias completadas, avisando al grupo antes |

### Decisiones

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| El sprint guarda su propia lista de historias completadas | deducir lo completado mirando qué historias del sprint quedaron Terminadas | una historia que no se completó vuelve al backlog y puede terminarse en el sprint siguiente: deducirlo la contaría como completada en los dos, y la velocidad de cada sprint quedaría mal. **Es un dato que el contrato del dominio no tiene: hay que agregarlo y avisar al grupo antes de implementar** |
| La lista de historias asignadas no se modifica al cerrar | sacarle al sprint las historias que volvieron al Product Backlog | sin ellas no se puede saber cuánto se había comprometido, que es la mitad de lo que mide la #16 |
| El cierre es una sola operación que recibe las completadas | marcar historia por historia y después cerrar | con dos pasos, un cierre a medias deja historias Terminadas en un sprint que sigue abierto, y no hay forma de volver atrás |
| Se valida toda la lista de completadas antes de cambiar cualquier estado | ir aplicando a medida que se recorre la lista | es la única forma de cumplir la RN-13 sobre varias historias a la vez, sin tener que deshacer cambios ya hechos |
| La fecha de cierre se guarda por día, como el resto de las fechas | guardarla con hora | es la convención que ya fijaron la #3 y la #8 (defecto #59); si el cierre guardara la hora, dos fechas del mismo día dejarían de ser iguales |
| No se agrega ninguna operación para reabrir o corregir un sprint cerrado | dejarla por si acaso | ninguna regla la pide, y abriría la puerta a cambiar un sprint que la #16 ya midió. Que un sprint cerrado no se pueda modificar sale de que no exista la operación, no de un control |
| Lo de esta historia va en `sprint_close.go` | agregarlo a `sprint.go` | mismo criterio que la #9: cuatro historias tocan el sprint y conviene que no se pisen |
| Se reutiliza el error de "no hay sprint abierto" que declara la #9 | un error propio para el cierre | es la misma situación, y con dos errores quien consume tiene que mirar los dos |

### Dependencias y restricciones

- **Depende de la #8 y de la #9.** No se espera a que estén en `dev`: los tests arman el
  sprint abierto y sus historias asignadas dentro del paquete, contra el contrato.
- **Cambia el contrato del dominio.** Agregar la lista de historias completadas al sprint es
  un cambio del contrato: se avisa al grupo y se actualiza `docs/diseno-dominio.md` **en el
  mismo pull request**. La #11 y la #16 lo consumen.
- **No se toca la asignación.** Agregar y quitar historias es de la #9, y después del cierre
  ya no se puede hacer porque no hay sprint abierto.
- **No se calcula ninguna métrica.** Ni puntos completados, ni velocidad, ni porcentajes: son
  de la #16.
- **No se modifica el estado de ninguna historia fuera del sprint que se cierra.**

> **Puerta 2** — ¿el diseño es viable y coherente con el resto del proyecto?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios: [pendiente]

## Fase 3 — Tareas

Cada tarea es un ciclo de TDD completo: el test que falla en un commit, la implementación que
lo hace pasar en otro. Van en el orden en que hay que hacerlas, y cada una deja el código
compilando y los tests en verde.

Los tests siguen la convención de [AGENTS.md](../../../AGENTS.md): un test por escenario, con
el nombre del escenario. Cada tarea de rechazo verifica, además del error, que el sprint
siguió abierto y que ninguna historia cambió de estado (RN-13).

| # | Tarea | Escenario BDD | Reglas | Estado |
|---|---|---|---|---|
| 1 | Cerrar el sprint lo deja Cerrado, con su fecha de cierre, y el proyecto queda sin sprint abierto | Despues de cerrar se puede crear el sprint siguiente | RN-1, RN-14 | pendiente |
| 2 | Cerrar registra las historias completadas y deja Terminadas a esas, y Pendientes a las demás | Cerrar el sprint registrando que se completo y que no | RN-2, RN-3, RN-4 | pendiente |
| 3 | El sprint cerrado conserva la lista de las historias que se le habían asignado | El sprint cerrado conserva lo que se le habia comprometido | RN-5 | pendiente |
| 4 | Cerrar con ninguna, con algunas y con todas las historias completadas | Cerrar con distinta cantidad de historias completadas | RN-3, RN-4, RN-9 | pendiente |
| 5 | Cerrar un sprint sin ninguna historia asignada | Cerrar un sprint sin ninguna historia asignada | RN-9 | pendiente |
| 6 | La misma historia indicada dos veces se registra una sola vez | La misma historia indicada dos veces cuenta una sola vez | RN-8 | pendiente |
| 7 | Rechazar el cierre si una historia indicada no está en el sprint, sin cambiar ninguna historia | Rechazar el cierre si una historia completada no esta en el sprint | RN-7, RN-13 | pendiente |
| 8 | Rechazar el cierre sin fecha de cierre | Rechazar el cierre sin fecha de cierre | RN-10, RN-13 | pendiente |
| 9 | Rechazar el cierre cuando el proyecto no tiene sprint abierto | Rechazar el cierre si no hay sprint abierto | RN-6 | pendiente |
| 10 | Cerrar dos veces se rechaza y el sprint conserva lo que había quedado registrado | No se puede cerrar dos veces el mismo sprint | RN-12, RN-13 | pendiente |
| 11 | Una fecha de cierre distinta de la fecha de fin del sprint se acepta | Una fecha de cierre distinta de la fecha de fin se acepta | RN-11 | pendiente |
| 12 | Cerrar el sprint de un proyecto que no existe informa que no se encontró | Cerrar el sprint de un proyecto que no existe | RN-15 | pendiente |
| 13 | Al cerrar se registran los puntos de cada historia, y reestimarla después no los cambia | Los puntos registrados al cerrar no cambian si se reestima la historia | RN-16 | pendiente |

### Por qué este orden

- **La 1 va primero** porque es el cierre más chico que se puede escribir: el sprint cambia
  de estado y el proyecto se queda sin sprint abierto. Las demás verifican sobre eso.
- **La 2 agrega lo que define esta historia**: la separación entre lo completado y lo que
  vuelve al Product Backlog. La 3 y la 4 la completan desde los dos bordes.
- **La 5 y la 6 son los casos raros del camino feliz** y no necesitan ningún rechazo escrito.
- **De la 7 a la 10 van los rechazos**, de a uno, y recién ahí la 10 puede probar que el
  cierre no se repite.
- **La 11 y la 12 van al final** porque miran los bordes de afuera: la fecha y el proyecto.

### Cobertura

Las dieciséis reglas quedan cubiertas: RN-1 (1), RN-2 (2, 6), RN-3 (2, 4), RN-4 (2, 4),
RN-5 (3), RN-6 (9), RN-7 (7), RN-8 (6), RN-9 (4, 5), RN-10 (8), RN-11 (11), RN-12 (10),
RN-13 (7, 8, 10), RN-14 (1), RN-15 (12), RN-16 (13).

| Criterio de aceptación | Tareas |
|---|---|
| Se puede cerrar el sprint abierto y queda Cerrado con su fecha | 1 |
| Los story points registrados al cerrar no cambian si se reestima una historia | 13 |
| Las historias completadas quedan Terminadas | 2, 4 |
| Las que no se indicaron quedan Pendientes | 2, 4 |
| El sprint conserva las historias que se le habían asignado | 3 |
| El sprint deja registrado cuáles se completaron | 2, 6 |
| Se puede cerrar sin ninguna historia completada | 4 |
| Se puede cerrar un sprint sin historias asignadas | 5 |
| Indicar dos veces la misma historia no es un error | 6 |
| Una historia que no está en el sprint rechaza el cierre sin cambiar nada | 7 |
| No se puede cerrar sin fecha de cierre | 8 |
| No se puede cerrar dos veces | 9, 10 |
| Después de cerrar no hay sprint abierto y se puede crear el siguiente | 1 |
| Cerrar en un proyecto que no existe informa que no se encontró | 12 |

Los doce escenarios de la Fase 1 tienen su tarea en la columna "Escenario BDD".

### Dependencias

- **Todas las tareas necesitan la #8 y la #9**: el sprint abierto y sus historias asignadas.
  Mientras no estén en `dev`, los tests los arman dentro del paquete.
- **La tarea 1 necesita además la creación de sprints de la #8** para verificar que, cerrado
  el primero, se puede crear el siguiente.
- **El dato de las historias completadas tiene que estar agregado al contrato del dominio
  antes de la tarea 2**, con el aviso al grupo.

> **Puerta 3** — ¿las tareas cubren todos los criterios de aceptación?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios: [pendiente]

## Fase 4 — Implementación

_Se completa durante la implementación._

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #10 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | `internal/domain/sprint_close_test.go` |
| Código Go | `internal/domain/sprint_close.go`, con el dato de las historias completadas en `internal/domain/sprint.go` y su error en `internal/domain/errors.go` |
