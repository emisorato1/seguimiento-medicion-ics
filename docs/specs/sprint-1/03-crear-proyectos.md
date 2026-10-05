# Spec: Crear y modificar proyectos

- **Historia:** #3
- **Sprint:** 1
- **Responsable:** Agustín Salinas
- **Estado:** Tareas — **pendiente de Puerta 3**
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de
  la consigna y de los criterios de aceptación de la issue. El responsable lo
  revisa, lo corrige si hace falta y lo presenta a la Puerta 1. Corregido por el
  responsable y presentado a la Puerta 1 por pull request
- **Depende de:** nada. Es la base del Sprint 1.
- **La necesitan:** #4 (integrantes), #6 (Product Backlog), #8 (sprints)

## Fase 1 — Requisitos

### Objetivo

Que un integrante del equipo pueda registrar en el sistema el proyecto de software sobre el
que va a trabajar, y corregir sus datos cuando cambien. Es la entidad raíz: todo lo demás
—integrantes, historias, sprints, métricas— cuelga de un proyecto.

### Qué queda afuera

- Borrar proyectos.
- Archivar o cerrar un proyecto.
- Permisos: en esta historia cualquiera puede crear y modificar cualquier proyecto.
- Los integrantes del proyecto, que son la historia #4.

### Entradas

| Dato | Obligatorio | Descripción |
|---|---|---|
| Nombre | sí | cómo se llama el proyecto |
| Descripción | no | texto libre que explica de qué se trata |
| Fecha de inicio | sí | cuándo arranca |
| Fecha de finalización | no | cuándo se estima terminarlo |

### Salidas

Un proyecto registrado, con un identificador propio que el sistema asigna y que no cambia.
Al consultarlo se devuelven sus datos y ese identificador.

### Reglas de negocio

1. **[RN-1]** El sistema debe asignar a cada proyecto un identificador único que no se
   repite y que no cambia durante toda la vida del proyecto.
2. **[RN-2]** Si se intenta crear o modificar un proyecto sin nombre, o con un nombre
   compuesto solo por espacios, el sistema debe rechazar la operación e informar que el
   nombre es obligatorio.
3. **[RN-3]** Si se intenta crear o modificar un proyecto sin fecha de inicio, el sistema debe
   rechazar la operación e informar que la fecha de inicio es obligatoria.
4. **[RN-4]** Si se indica una fecha de finalización anterior a la fecha de inicio, el sistema
   debe rechazar la operación e informar que las fechas son incoherentes.
5. **[RN-5]** Cuando se modifica un proyecto existente, el sistema debe aplicar las mismas
   validaciones que al crearlo. Si alguna falla, el proyecto debe quedar como estaba antes.
6. **[RN-6]** Si se intenta modificar un proyecto que no existe, el sistema debe informar que
   no se encontró el proyecto y no modificar ni crear nada.
7. **[RN-7]** El sistema debe permitir que dos proyectos distintos tengan el mismo nombre:
   lo que los distingue es el identificador.
8. **[RN-8]** Cuando se modifica un proyecto, el sistema debe cambiar solo los datos que se
   indicaron en esa modificación y dejar el resto como estaba. Un dato que no se indica
   nunca se borra.
9. **[RN-9]** Cuando se pide explícitamente quitar la descripción o la fecha de finalización
   de un proyecto, el sistema debe dejarlo sin ese dato y conservar el resto. Quitar un dato
   es un pedido distinto de no indicarlo (RN-8). El nombre y la fecha de inicio no se pueden
   quitar (RN-2 y RN-3).
10. **[RN-10]** Cuando se crea o se modifica un proyecto, el sistema debe guardar su nombre
    sin los espacios del principio y del final.

### Casos límite

- Proyecto sin descripción y sin fecha de finalización: es válido, son opcionales.
- Fecha de finalización **igual** a la de inicio: es válida, un proyecto puede durar un día.
- Nombre con espacios al principio y al final: se guarda sin esos espacios.
- Nombre de un solo carácter: es válido.
- Modificar un proyecto sin cambiar nada: es válido y no da error.
- Modificar solo la fecha de inicio de un proyecto que ya tiene fecha de finalización: si la
  nueva fecha de inicio queda después de la de finalización, se rechaza por fechas
  incoherentes, aunque la fecha de finalización no se haya tocado.

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| Nombre vacío o solo espacios | rechaza e informa que el nombre es obligatorio |
| Falta la fecha de inicio | rechaza e informa que la fecha de inicio es obligatoria |
| Fecha de fin anterior a la de inicio | rechaza e informa que las fechas son incoherentes |
| Identificador inexistente al modificar | informa que no se encontró el proyecto |

### Criterios de aceptación

- [ ] Se puede crear un proyecto con nombre, descripción, fecha de inicio y fecha de fin, y
      queda registrado con un identificador propio.
- [ ] Se puede crear un proyecto con solo nombre y fecha de inicio.
- [ ] No se puede crear ni modificar un proyecto sin nombre.
- [ ] No se puede crear ni modificar un proyecto sin fecha de inicio.
- [ ] No se puede crear ni modificar un proyecto con fecha de fin anterior a la de inicio.
- [ ] Se pueden modificar los datos de un proyecto ya creado y los cambios quedan guardados.
- [ ] Una modificación rechazada deja el proyecto como estaba.
- [ ] Modificar un proyecto sin indicar un dato no borra ese dato.
- [ ] Se le puede quitar la descripción o la fecha de finalización a un proyecto pidiéndolo
      explícitamente.
- [ ] Dos proyectos creados con el mismo nombre reciben identificadores distintos.
- [ ] Modificar un proyecto inexistente informa el error y no crea nada.

### Escenarios BDD

```gherkin
Escenario: Crear un proyecto con todos los datos
  Dado que no existe ningun proyecto registrado
  Cuando creo un proyecto con nombre "Sistema de metricas", inicio 01/10/2026 y fin 30/11/2026
  Entonces el proyecto queda registrado
  Y recibo su identificador

Escenario: Crear un proyecto solo con los datos obligatorios
  Dado que no existe ningun proyecto registrado
  Cuando creo un proyecto con nombre "Sistema de metricas" e inicio 01/10/2026
  Entonces el proyecto queda registrado sin descripcion y sin fecha de finalizacion

Esquema del escenario: Rechazar un proyecto sin nombre
  Cuando intento crear un proyecto con nombre <nombre> e inicio 01/10/2026
  Entonces la operacion se rechaza
  Y se informa que el nombre es obligatorio

  Ejemplos:
    | nombre      |
    | ""          |
    | "   "       |

Escenario: Rechazar un proyecto sin fecha de inicio
  Cuando intento crear un proyecto con nombre "Sistema de metricas" y sin fecha de inicio
  Entonces la operacion se rechaza
  Y se informa que la fecha de inicio es obligatoria

Escenario: Guardar el nombre sin espacios en los extremos
  Cuando creo un proyecto con nombre "  Sistema de metricas  " e inicio 01/10/2026
  Entonces el proyecto queda registrado con nombre "Sistema de metricas"

Escenario: Aceptar un nombre de un solo caracter
  Cuando creo un proyecto con nombre "X" e inicio 01/10/2026
  Entonces el proyecto queda registrado con nombre "X"

Escenario: Dos proyectos con el mismo nombre
  Dado un proyecto registrado con nombre "Sistema de metricas"
  Cuando creo otro proyecto con nombre "Sistema de metricas" e inicio 01/10/2026
  Entonces el proyecto queda registrado
  Y su identificador es distinto del identificador del primero

Escenario: Rechazar fechas incoherentes
  Cuando intento crear un proyecto con inicio 01/10/2026 y fin 30/09/2026
  Entonces la operacion se rechaza
  Y se informa que las fechas son incoherentes

Escenario: Aceptar un proyecto que empieza y termina el mismo dia
  Cuando creo un proyecto con inicio 01/10/2026 y fin 01/10/2026
  Entonces el proyecto queda registrado

Escenario: Modificar solo el nombre conserva los demas datos
  Dado un proyecto registrado con nombre "Sistema de metricas", descripcion "TP de ICS",
    inicio 01/10/2026 y fin 30/11/2026
  Cuando modifico solo su nombre a "Seguimiento y medicion"
  Entonces el proyecto pasa a llamarse "Seguimiento y medicion"
  Y conserva la descripcion "TP de ICS", el inicio 01/10/2026 y el fin 30/11/2026

Escenario: Guardar sin espacios en los extremos el nombre modificado
  Dado un proyecto registrado con nombre "Sistema de metricas"
  Cuando modifico su nombre a "  Seguimiento y medicion  "
  Entonces el proyecto pasa a llamarse "Seguimiento y medicion"

Escenario: Modificar un proyecto sin cambiar nada
  Dado un proyecto registrado con nombre "Sistema de metricas", descripcion "TP de ICS",
    inicio 01/10/2026 y fin 30/11/2026
  Cuando lo modifico indicando el mismo nombre que ya tiene y ningun otro dato
  Entonces la operacion no da error
  Y el proyecto conserva todos sus datos, incluidas la descripcion y la fecha de finalizacion

Escenario: Rechazar una modificacion que deja el proyecto sin nombre
  Dado un proyecto registrado con nombre "Sistema de metricas"
  Cuando intento modificar su nombre a "   "
  Entonces la operacion se rechaza
  Y se informa que el nombre es obligatorio
  Y el proyecto sigue llamandose "Sistema de metricas"

Escenario: Rechazar una fecha de inicio posterior a la de finalizacion al modificar
  Dado un proyecto registrado con inicio 01/10/2026 y fin 30/11/2026
  Cuando intento modificar su fecha de inicio a 01/12/2026
  Entonces la operacion se rechaza
  Y se informa que las fechas son incoherentes
  Y el proyecto conserva inicio 01/10/2026 y fin 30/11/2026

Escenario: Quitar la fecha de finalizacion
  Dado un proyecto registrado con nombre "Sistema de metricas", inicio 01/10/2026 y fin 30/11/2026
  Cuando pido explicitamente quitarle la fecha de finalizacion
  Entonces el proyecto queda sin fecha de finalizacion
  Y conserva su nombre y su fecha de inicio

Escenario: Quitar la descripcion
  Dado un proyecto registrado con nombre "Sistema de metricas", descripcion "TP de ICS",
    inicio 01/10/2026 y fin 30/11/2026
  Cuando pido explicitamente quitarle la descripcion
  Entonces el proyecto queda sin descripcion
  Y conserva su nombre, su fecha de inicio y su fecha de finalizacion

Escenario: Modificar un proyecto que no existe
  Dado que no existe ningun proyecto con identificador 99
  Cuando intento modificar el proyecto 99
  Entonces se informa que no se encontro el proyecto
  Y no se crea ningun proyecto
```

**Decidido en la revisión de la Puerta 1 (PR #26):** al modificar se indica solo lo que
cambia, y para quitar la descripción o la fecha de finalización hay que pedirlo
explícitamente (opción 2 de las que planteó Emiliano Sorato). Así "no tocar un dato" y
"borrar un dato" son pedidos distintos (RN-8 y RN-9). Se permite quitar esos datos porque son
opcionales y un proyecto puede no tener fecha de fin estimada todavía.

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: Emiliano Sorato — Fecha: 2026-10-02 — Comentarios: en la primera revisión (PR #26)
> se pidieron cambios porque la RN-8 y la RN-9 se pisaban: no se distinguía "no mandé el dato"
> de "mandé el dato vacío para borrarlo". El responsable eligió resolverlo con un pedido
> explícito de quitar el dato, y actualizó la RN-8, la RN-9, el criterio de aceptación y los
> escenarios de quitar la descripción y la fecha de finalización. Verificado y aprobado.

## Fase 2 — Diseño

Las decisiones transversales del proyecto —un solo paquete `internal/domain` con un archivo
por entidad, datos en memoria, errores como valores del dominio, mensajes en español,
validaciones en el dominio— están en [AGENTS.md](../../../AGENTS.md) y no se repiten acá.
Esta fase registra solo lo propio de esta historia.

### Enfoque

Es la primera historia con código de dominio, así que además de resolver el proyecto deja
dos cosas que usan las demás: **el registro de proyectos**, que es donde viven los proyectos
mientras corre el programa, y **`errors.go`**, donde se declaran los errores del dominio.

El registro es el único que guarda proyectos y el único que les asigna identificador. Crear,
modificar y buscar un proyecto se hace siempre a través de él. La consola (#24) y las
historias que cuelgan de un proyecto (#4, #6) lo usan para encontrar el proyecto con el que
trabajan.

Crear y modificar siguen el mismo orden que la #6:

1. **Armar los datos resultantes, aparte del proyecto.** Al crear, son los datos que
   llegaron. Al modificar, son los datos actuales del proyecto con los cambios pedidos
   aplicados. Si el proyecto a modificar no existe, se informa que no se encontró y no se
   arma ni se guarda nada (RN-6).
2. **Normalizar:** sacar los espacios de los extremos del nombre (RN-10).
3. **Validar el resultado completo** (RN-2, RN-3, RN-4).
4. **Aplicarlo** al proyecto recién si todo validó.

Como el resultado se arma aparte y se aplica al final, una modificación rechazada nunca llegó
a tocar el proyecto (RN-5). Y como se valida el resultado y no solo lo que llegó, una fecha de
inicio nueva se compara contra la fecha de finalización que ya estaba (caso límite de la
Fase 1).

### Archivos afectados

| Archivo | Se crea o se modifica | Para qué |
|---|---|---|
| `internal/domain/project.go` | se crea | el proyecto, sus validaciones y el pedido de modificación |
| `internal/domain/project_registry.go` | se crea | el registro de proyectos: alta, modificación, búsqueda e identificadores |
| `internal/domain/project_test.go` | se crea | los tests de los escenarios BDD de la Fase 1 |
| `internal/domain/errors.go` | se crea | los errores de esta historia; las demás agregan los suyos acá |

### Decisiones

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| Los proyectos se numeran desde 1, **únicos en todo el sistema**, con un contador del registro que **nunca retrocede** | calcular el próximo como "cantidad de proyectos + 1" | AGENTS.md fija la numeración de lo que está *dentro* de un proyecto, pero no la de los proyectos. Es el mismo criterio que la #6 usa para las historias: contar se rompe en cuanto exista borrado |
| Los proyectos viven en un **registro que se crea explícitamente** y se pasa a quien lo necesita | una lista global del paquete | con una lista global los tests comparten datos y dependen del orden en que corren; con un registro por test, cada escenario arranca de cero ("no existe ningún proyecto registrado") |
| El registro está **en un archivo propio**, separado del proyecto | todo en `project.go` | la #4 y la #6 van a modificar `project.go` para agregar integrantes y backlog; separar el registro reduce los choques en ese archivo |
| El pedido de modificación indica, para cada dato, si **no se toca**, si **cambia a un valor nuevo** o, solo para la descripción y la fecha de finalización, si **se quita** | usar un valor vacío para decir "quitar" | un valor vacío vuelve a mezclar "no lo mandé" con "lo quiero borrar", que es justo lo que se separó en la Puerta 1 (RN-8 y RN-9) |
| Para el nombre y la fecha de inicio **no existe la opción de quitar** | aceptar el pedido y rechazarlo con un error | si no se puede pedir, no hace falta un error ni un escenario para eso. Dejar el nombre vacío o en blanco sigue cubierto por la RN-2 |
| Las fechas son **días**, sin hora | fecha y hora | la Fase 1 compara días: un proyecto que empieza y termina el mismo día es válido. Con hora, dos fechas del mismo día podrían salir distintas |

### Dependencias y restricciones

- **No depende de nada**, pero la #4, la #6 y la #12 dependen de esta: modifican
  `project.go` o agregan errores a `errors.go`. Mientras la #3 no esté en `dev`, no pueden
  escribir código sin chocar.
- **`project.go` tiene solo lo de esta historia.** Los integrantes los agrega la #4 y el
  backlog la #6; esta historia no les deja lugares reservados.
- **No se toca `cmd/`.** Pedir los datos por teclado, el formato en que se escriben las
  fechas y cómo se muestran los errores es de la consola (#24).
- La #24 depende de esta fase para saber dónde se guardan los proyectos: es el registro.
  Lo crea quien arranca el programa (la consola), y hay uno solo mientras el programa corre.

**Decidido en la Puerta 2:** cuando alguien busca un proyecto en el registro, **recibe el
proyecto guardado**, no una copia, y los datos del proyecto están ocultos fuera del paquete:
solo se cambian con operaciones que validan.

La opción la propuso Maximiliano Eula y la respaldó Juliana Bustos en este PR. El argumento
que la define: la consola vive en `cmd/`, así que no puede tocar datos ocultos de
`internal/domain` y no tiene forma de saltear las validaciones; y dentro del paquete la copia
no protege nada, porque es el mismo paquete. O sea que la copia paga el costo —obligar a la
#4, la #6 y la #12 a guardar cada cambio por el registro— sin comprar la protección. Peor
aún, si alguna de esas historias se olvida de guardar, el cambio se pierde en silencio, sin
error y sin que ningún escenario BDD lo muestre.

Va con una condición, que forma parte de la decisión: como se trabaja sobre el proyecto
guardado, **cada operación valida todo antes de cambiar cualquier dato**. Eso es lo que
sostiene la RN-5 de esta historia, la RN-12 de la #4 y la RN-11 de la #6.

Las dos quedaron registradas en [AGENTS.md](../../../AGENTS.md) porque valen para todo el
proyecto.

> **Puerta 2** — ¿el diseño es viable y coherente con el resto del proyecto?
> Aprobó: Emiliano Sorato — Fecha: 2026-10-05 — Comentarios: aprobada con la decision de la pregunta abierta: el registro entrega el proyecto guardado, con sus datos ocultos fuera del paquete (opcion propuesta por Maximiliano Eula y respaldada por Juliana Bustos en el PR #41). Queda registrada en AGENTS.md porque vale para todo el proyecto.

## Fase 3 — Tareas

Cada tarea es un ciclo de TDD completo: el test que falla en un commit, la implementación que
lo hace pasar en otro. Están en el orden en que hay que hacerlas, y cada una deja el código
compilando y los tests en verde. Los tests siguen la convención de
[AGENTS.md](../../../AGENTS.md): uno por escenario, con el nombre del escenario.

| # | Tarea | Reglas que cubre | Escenarios | Estado |
|---|---|---|---|---|
| 1 | Crear un proyecto con todos los datos en un registro vacío: queda registrado, recibe el identificador 1 y se puede consultar por ese identificador | RN-1 | Crear un proyecto con todos los datos | pendiente |
| 2 | Crear un proyecto solo con nombre y fecha de inicio: queda sin descripción y sin fecha de finalización | — (datos opcionales de la Fase 1) | Crear un proyecto solo con los datos obligatorios | pendiente |
| 3 | Dos proyectos, aunque tengan el mismo nombre, reciben identificadores distintos; el contador nunca retrocede | RN-1, RN-7 | Dos proyectos con el mismo nombre | pendiente |
| 4 | El nombre se guarda sin los espacios de los extremos; un nombre de un solo carácter es válido | RN-10 | Guardar el nombre sin espacios en los extremos; Aceptar un nombre de un solo caracter | pendiente |
| 5 | Rechazar el nombre vacío o de solo espacios, y que la creación rechazada no registre nada. Crea `errors.go` con su primer error | RN-2 | Rechazar un proyecto sin nombre | pendiente |
| 6 | Rechazar la creación sin fecha de inicio | RN-3 | Rechazar un proyecto sin fecha de inicio | pendiente |
| 7 | Rechazar la fecha de finalización anterior a la de inicio; la misma fecha en las dos es válida | RN-4 | Rechazar fechas incoherentes; Aceptar un proyecto que empieza y termina el mismo dia | pendiente |
| 8 | Modificar un proyecto que no existe informa que no se encontró y no crea ninguno | RN-6 | Modificar un proyecto que no existe | pendiente |
| 9 | Modificar indicando solo lo que cambia: el resto se conserva, y modificar sin cambiar nada no da error | RN-8 | Modificar solo el nombre conserva los demas datos; Modificar un proyecto sin cambiar nada | pendiente |
| 10 | Modificar normaliza y valida el resultado completo con las mismas reglas que crear, todo antes de cambiar cualquier dato: una modificación rechazada deja el proyecto como estaba | RN-2, RN-3, RN-4, RN-5, RN-10 | Guardar sin espacios en los extremos el nombre modificado; Rechazar una modificacion que deja el proyecto sin nombre; Rechazar una fecha de inicio posterior a la de finalizacion al modificar | pendiente |
| 11 | Quitar la descripción o la fecha de finalización con un pedido explícito, conservando el resto | RN-9 | Quitar la fecha de finalizacion; Quitar la descripcion | pendiente |

### Por qué este orden

- **La 1 es el camino feliz** y arma lo mínimo para todo lo demás: el registro, el proyecto y
  la consulta por identificador. Sin eso no hay contra qué escribir ningún otro test.
- **La 4 va antes que la 5** porque la validación corre sobre el nombre ya normalizado: es lo
  que hace que `"   "` cuente como vacío. Al revés, el test de la 5 pasaría por el motivo
  equivocado. Es el mismo criterio que la #6.
- **Las validaciones (5 a 7) van de a una**, para que cada test nuevo falle por una sola razón.
- **La 8 abre la modificación por el camino más corto**: encontrar el proyecto o informar que
  no existe. Es lo primero que hace cualquier modificación.
- **La 10 no se separa en "valida" y "no cambia nada si se rechaza"**, a diferencia de la #6.
  Con la decisión de AGENTS.md de validar todo antes de cambiar cualquier dato, las dos cosas
  salen del mismo código: un segundo test pasaría en verde sin escribir nada, y eso no es un
  ciclo de TDD.

### Cobertura

**Reglas de negocio:** RN-1 (1, 3), RN-2 (5, 10), RN-3 (6, 10), RN-4 (7, 10), RN-5 (10),
RN-6 (8), RN-7 (3), RN-8 (9), RN-9 (11), RN-10 (4, 10).

**Criterios de aceptación:**

| Criterio | Tareas |
|---|---|
| Crear con todos los datos y un identificador propio | 1 |
| Crear solo con nombre y fecha de inicio | 2 |
| No crear ni modificar sin nombre | 5, 10 |
| No crear ni modificar sin fecha de inicio | 6, 10 |
| No crear ni modificar con fecha de fin anterior a la de inicio | 7, 10 |
| Modificar y que los cambios queden guardados | 9 |
| Una modificación rechazada deja el proyecto como estaba | 10 |
| Modificar sin indicar un dato no lo borra | 9 |
| Quitar la descripción o la fecha de fin pidiéndolo explícitamente | 11 |
| Dos proyectos con el mismo nombre reciben identificadores distintos | 3 |
| Modificar un proyecto inexistente informa el error y no crea nada | 8 |

**Escenarios BDD:** los 17 escenarios de la Fase 1 tienen tarea, como muestra la columna
*Escenarios* de la tabla.

Que los datos del proyecto estén ocultos fuera del paquete (decisión de la Puerta 2) no lleva
tarea propia: no es un comportamiento que se pruebe con un escenario, sino una restricción
que el compilador hace cumplir.

> **Puerta 3** — ¿las tareas cubren todos los criterios de aceptación?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios:

## Fase 4 — Implementación

_Se completa a medida que se avanza._

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #3 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | pendiente |
| Código Go | pendiente |
