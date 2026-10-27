# Spec: Generar el reporte de un proyecto o de un sprint

- **Historia:** #18
- **Sprint:** 3
- **Responsable:** Agustín Salinas
- **Estado:** Puerta 3 aprobada — **lista para implementar**
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de la consigna y de los criterios de aceptación de la issue. El responsable lo revisa, lo corrige si hace falta y lo presenta a las puertas.
- **Depende de:** #16 (las métricas), #8 (buscar un sprint por su número), #10 (qué se
  completó en cada sprint) y #15 (los defectos que se listan), todas en `dev`. La regla de
  formato de los números es la de la #17, en `internal/domain/format.go`, que ya está en
  `dev`; esta historia le suma las fechas
- **La necesitan:** #19 (exportar el reporte a PDF)

## Fase 1 — Requisitos

### Objetivo

Que el equipo pueda generar, para un sprint o para el proyecto entero, un reporte que junte
en un solo lugar qué se planificó y qué se completó, las estimaciones, el esfuerzo, las
métricas y los defectos, listo para mostrar en la review.

### Qué queda afuera

- **Calcular las métricas**: es la #16. El reporte toma sus números y los escribe.
- **Exportar a PDF**: es la #19. Acá el reporte se arma; la #19 lo escribe como PDF y la web lo
  ofrece como descarga (#78). No se guarda en disco.
- **Gráficos**: son del dashboard (#17).
- Guardar los reportes generados, llevar un historial de reportes o compararlos.
- Elegir qué secciones van: el reporte tiene siempre las mismas.
- El esfuerzo por historia o por integrante (ver el punto 3 de "Para decidir en la puerta").

### Entradas

| Dato | Obligatorio | Descripción |
|---|---|---|
| Proyecto | sí | de qué proyecto es el reporte |
| Sprint | solo para el reporte de un sprint | el número del sprint dentro del proyecto |

### Salidas

Un reporte con su título, el día en que se generó, si es de un sprint o del proyecto, los
números de las métricas de ese mismo alcance y cuatro secciones de texto, cada una con su
encabezado y sus líneas:

| Sección | En el reporte de un sprint | En el reporte del proyecto |
|---|---|---|
| Resumen | proyecto, Sprint Goal, fechas y estado del sprint | proyecto, fechas y cantidad de sprints |
| Historias | totales y una línea por historia comprometida, con su situación en el sprint | totales y una línea por historia del backlog, con su estado actual |
| Estimación y esfuerzo | puntos, velocidad, horas y desviación | lo mismo, con la velocidad promedio |
| Defectos | totales y los defectos detectados o resueltos en el sprint | totales y todos los defectos |

### Reglas de negocio

1. **[RN-1]** El sistema debe permitir generar el reporte de un sprint de un proyecto y el del
   proyecto entero.
2. **[RN-2]** Cuando se genera un reporte, el sistema debe tomar los números de las métricas
   del mismo alcance y no calcular ninguno por su cuenta: cada número del reporte tiene que
   coincidir con el que dan las métricas.
3. **[RN-3]** El sistema debe titular el reporte "Reporte del sprint N — Proyecto" o
   "Reporte del proyecto Proyecto", con el número y el nombre que correspondan, y registrar
   el día en que se generó y si es de un sprint o del proyecto.
4. **[RN-4]** El sistema debe armar siempre las cuatro secciones, en este orden: Resumen,
   Historias, Estimación y esfuerzo, Defectos.
5. **[RN-5]** Si una sección no tiene nada que listar, el sistema debe incluirla igual, con
   sus totales en cero y una línea que diga que no hay historias o que no hay defectos.
6. **[RN-6]** En el reporte de un sprint, el Resumen debe tener el nombre del proyecto, el
   Sprint Goal, las fechas de inicio y de fin, y el estado: cerrado, con su fecha de cierre,
   o abierto.
7. **[RN-7]** Mientras el sprint esté abierto, el sistema debe avisar en el Resumen que los
   números son parciales porque el sprint todavía no cerró.
8. **[RN-8]** En el reporte del proyecto, el Resumen debe tener el nombre, la fecha de
   inicio, la de fin —o "sin fecha de fin"— y cuántos sprints tiene y cuántos están cerrados.
9. **[RN-9]** La sección Historias debe tener las historias planificadas, las completadas con
   su porcentaje, y una línea por historia con su identificador, su título, sus story points
   —o "sin estimar"— y su situación.
10. **[RN-10]** En el reporte de un sprint, el sistema debe listar las historias comprometidas
    en él, con la situación "completada" o "no completada" según lo registrado al cerrarlo, y
    no según su estado actual. Mientras el sprint esté abierto, como todavía no tiene cierre,
    cada historia figura con sus story points actuales y la situación "en curso". En el del
    proyecto, debe listar todas las historias del backlog con su estado actual.
11. **[RN-11]** La sección Estimación y esfuerzo debe tener los story points planificados y
    completados, la velocidad —"velocidad promedio" en el reporte del proyecto—, las horas
    estimadas, las horas reales y la desviación. Mientras el proyecto no tenga ningún sprint
    cerrado, en lugar de la velocidad promedio debe decir "sin sprints cerrados", como el
    dashboard (#17, RN-4).
12. **[RN-12]** La sección Defectos debe tener los detectados y los resueltos, y una línea por
    defecto con su identificador, su título, su severidad, su estado y, si está resuelto, el
    sprint en que se resolvió. En el reporte de un sprint van los defectos detectados o
    resueltos en ese sprint; en el del proyecto, todos.
13. **[RN-13]** El sistema debe listar las historias y los defectos por su identificador, de
    menor a mayor.
14. **[RN-14]** El sistema debe escribir los porcentajes y las horas con un decimal, coma
    decimal y su unidad (% o h); los story points, las historias, los defectos y la
    velocidad, sin decimales; la desviación, con su signo cuando es negativa y sin signo
    cuando es positiva o cero; y las fechas como día/mes/año, por ejemplo 01/10/2026. Es la
    misma regla del dashboard (#17, RN-15): un número o una fecha se lee igual en el
    dashboard, en el reporte y en el PDF.
15. **[RN-15]** Cuando se genera un reporte, el sistema debe armarlo con los datos de ese
    momento, sin guardarlo y sin modificar ningún dato.
16. **[RN-16]** Si se pide el reporte de un sprint que no existe en el proyecto, el sistema
    debe informar que no se encontró el sprint; si el proyecto no existe, que no se encontró
    el proyecto.

### Casos límite

- Sprint sin historias comprometidas: la sección Historias tiene sus totales en cero y la
  línea "No hay historias en este sprint."
- Sprint sin defectos: la sección Defectos dice "No hay defectos en este sprint."
- Proyecto sin sprints, o sin ninguno cerrado: el Resumen dice "0 sprints, 0 cerrados" (o
  los que tenga) y la velocidad promedio dice "sin sprints cerrados", no 0.
- Historia que no se completó en el sprint 1 y sí en el 2: en el reporte del sprint 1 figura
  "no completada", en el del 2 "completada" y en el del proyecto "Terminada".
- Defecto detectado en el sprint 1 y resuelto en el 2: aparece en los reportes de los dos.
- Desviación positiva o cero: sin signo ("25,0 %", "0,0 %"); negativa: "-20,0 %".
  Porcentaje de 66,666…: "66,7 %".
- Proyecto sin fecha de fin: el Resumen lo dice en lugar de dejar el dato vacío.

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| El sprint pedido no existe en el proyecto | informa que no se encontró el sprint |
| El proyecto no existe | informa que no se encontró el proyecto |

Que falten datos **no** es un error: las secciones lo dicen (RN-5).

### Criterios de aceptación

- [ ] Se puede generar el reporte de un sprint y el del proyecto entero, cada uno con su
      título, su día de generación y su alcance.
- [ ] El reporte tiene siempre las cuatro secciones en orden, y una sección sin datos dice
      que no hay.
- [ ] Todos los números del reporte coinciden con las métricas del mismo alcance.
- [ ] El Resumen de un sprint tiene proyecto, Sprint Goal, fechas y estado, y si el sprint
      está abierto avisa que los números son parciales.
- [ ] El Resumen del proyecto tiene nombre, fechas y cantidad de sprints.
- [ ] Incluye las historias planificadas y completadas, una por línea, con sus story points
      o "sin estimar".
- [ ] En el reporte de un sprint, cada historia figura como completada o no según el cierre
      de ese sprint.
- [ ] Incluye las estimaciones y el esfuerzo registrado: puntos, velocidad, horas estimadas y
      reales, y desviación.
- [ ] Incluye los defectos del período, uno por línea, con su severidad y su estado.
- [ ] Horas y porcentajes van con un decimal, coma y su unidad, la desviación negativa con
      su signo y la positiva sin signo, la velocidad sin decimales, y las fechas como
      día/mes/año.
- [ ] Generar el reporte no cambia ningún dato.
- [ ] Pedir el reporte de un sprint o de un proyecto inexistente informa que no se encontró.

### Escenarios BDD

```gherkin
Escenario: Generar el reporte de un sprint cerrado
  Dado el proyecto "Seguimiento ICS" con el sprint 1 cerrado
  Cuando genero el reporte del sprint 1
  Entonces el reporte se titula "Reporte del sprint 1 — Seguimiento ICS"
  Y es un reporte de sprint, con fecha de generacion el dia de hoy
  Y tiene las secciones "Resumen", "Historias", "Estimación y esfuerzo" y "Defectos", en ese orden
  Y el proyecto queda como estaba

Escenario: Los numeros del reporte son los de las metricas
  Dado un sprint con 13 puntos planificados, 10 completados, velocidad 10, 16 horas estimadas, 20 reales y desviacion 25
  Cuando genero el reporte del sprint
  Entonces el reporte trae esas mismas metricas
  Y la seccion "Estimación y esfuerzo" tiene las lineas:
    | Story points planificados: 13 |
    | Story points completados: 10  |
    | Velocidad: 10                 |
    | Horas estimadas: 16,0 h       |
    | Horas reales: 20,0 h          |
    | Desviación: 25,0 %            |

Esquema del escenario: El resumen dice en que estado esta el sprint
  Dado el sprint 1 con el goal "Tener el MVP", del 01/10/2026 al 25/10/2026, <estado>
  Cuando genero el reporte del sprint 1
  Entonces el resumen tiene el goal "Tener el MVP" y las fechas 01/10/2026 y 25/10/2026
  Y el resumen dice <linea>

  Ejemplos:
    | estado                   | linea                                                   |
    | cerrado el 25/10/2026    | "Estado: cerrado el 25/10/2026"                         |
    | abierto                  | "Estado: abierto. Los números son parciales."           |

Escenario: Las historias del sprint figuran segun el cierre de ese sprint
  Dado el sprint 1 cerrado con las historias 1 y 2 comprometidas y solo la 1 completada
  Y la historia 2 se completo despues en el sprint 2 y hoy esta "Terminada"
  Cuando genero el reporte del sprint 1
  Entonces la seccion "Historias" dice 2 planificadas y 1 completada
  Y la historia 1 figura como "completada"
  Y la historia 2 figura como "no completada"

Escenario: Una historia sin estimar figura como sin estimar
  Dado el sprint 1 con la historia 3 comprometida y sin story points
  Cuando genero el reporte del sprint 1
  Entonces la linea de la historia 3 dice "sin estimar" en lugar de los puntos

Escenario: Las historias de un sprint abierto figuran en curso
  Dado el sprint 1 abierto con la historia 1 estimada en 3 story points
  Cuando genero el reporte del sprint 1
  Entonces la linea de la historia 1 dice "3 SP" y "en curso"

Esquema del escenario: Una seccion sin datos lo dice
  Dado el sprint 1 sin <que>
  Cuando genero el reporte del sprint 1
  Entonces la seccion <seccion> tiene sus totales en cero
  Y tiene la linea <linea>

  Ejemplos:
    | que       | seccion     | linea                              |
    | historias | "Historias" | "No hay historias en este sprint." |
    | defectos  | "Defectos"  | "No hay defectos en este sprint."  |

Escenario: Los defectos del sprint son los detectados o resueltos en el
  Dado el defecto 1 detectado en el sprint 1 y resuelto en el 2
  Y el defecto 2 detectado en el sprint 2 y abierto
  Y el defecto 3 detectado en el sprint 3
  Cuando genero el reporte del sprint 2
  Entonces la seccion "Defectos" lista los defectos 1 y 2
  Y la linea del defecto 1 dice su severidad, "Resuelto" y "sprint 2"

Escenario: Generar el reporte del proyecto entero
  Dado el proyecto "Seguimiento ICS", con inicio 01/09/2026 y fin 02/11/2026
  Y con los sprints 1 y 2 cerrados y el 3 abierto
  Y con las historias 1 a 4, la 4 sin haber entrado nunca a un sprint
  Cuando genero el reporte del proyecto
  Entonces el reporte se titula "Reporte del proyecto Seguimiento ICS"
  Y el resumen dice las dos fechas y "3 sprints, 2 cerrados"
  Y la seccion "Historias" lista las historias 1 a 4 con su estado actual
  Y la seccion "Estimación y esfuerzo" dice "Velocidad promedio"
  Y la seccion "Defectos" lista todos los defectos del proyecto

Escenario: Un proyecto sin fecha de fin
  Dado un proyecto sin fecha de fin
  Cuando genero el reporte del proyecto
  Entonces el resumen dice "sin fecha de fin"

Escenario: Un proyecto sin sprints cerrados no muestra una velocidad de cero
  Dado un proyecto con el sprint 1 abierto y ninguno cerrado
  Cuando genero el reporte del proyecto
  Entonces el resumen dice "1 sprint, 0 cerrados"
  Y la seccion "Estimación y esfuerzo" tiene la linea "Velocidad promedio: sin sprints cerrados"

Esquema del escenario: Formato de los numeros
  Dado un sprint cuyas metricas tienen <metrica> igual a <valor>
  Cuando genero el reporte del sprint
  Entonces el reporte tiene la linea <linea>

  Ejemplos:
    | metrica                       | valor  | linea                                |
    | horas reales                  | 7,5    | "Horas reales: 7,5 h"                |
    | desviacion                    | 25     | "Desviación: 25,0 %"                 |
    | desviacion                    | -20    | "Desviación: -20,0 %"                |
    | desviacion                    | 0      | "Desviación: 0,0 %"                  |
    | porcentaje con 2 completadas  | 66,666 | "Historias completadas: 2 (66,7 %)"  |

Escenario: Las historias y los defectos van ordenados por identificador
  Dado un proyecto con las historias 1, 2 y 3, priorizadas de modo que el backlog las muestra como 3, 1 y 2
  Y cuyos defectos 2 y 1 se resolvieron en ese orden
  Cuando genero el reporte del proyecto
  Entonces las historias aparecen como 1, 2 y 3
  Y los defectos aparecen como 1 y 2

Esquema del escenario: Pedir un reporte de algo que no existe
  Dado un proyecto sin sprint 99, y que no existe el proyecto 99
  Cuando intento generar <reporte>
  Entonces se informa que no se encontro <que>

  Ejemplos:
    | reporte                     | que         |
    | el reporte del sprint 99    | el sprint   |
    | el reporte del proyecto 99  | el proyecto |
```

**Para decidir en la puerta:**

1. **Las historias y los defectos se listan uno por uno**, además de los totales. La issue
   dice "incluye historias planificadas y completadas": en una review importa saber cuáles.
   La consecuencia es que el reporte de un proyecto grande es largo; no se recorta.
2. **En el reporte de un sprint, la situación de cada historia sale del cierre (RN-10)**, por
   la misma razón que el contrato guarda las completadas: una historia que no se terminó y se
   terminó después figuraría como completada en un sprint que no la completó.
3. **El esfuerzo va en totales**, sin desglose por historia ni por integrante. Por historia,
   las horas de un sprint siguen la regla del "último sprint" de la #16 y una línea por
   historia podría no sumar el total; por integrante no es una métrica de la consigna. Si el
   equipo lo quiere, se agrega una sección y se decide qué hacer con esa diferencia.
4. **El reporte escribe los números como texto (RN-14)**, con la misma regla que el
   dashboard (#17, RN-15): un número se lee igual en la pantalla, en el reporte y en el PDF,
   que copia las líneas sin tocarlas. Si el equipo cambia el formato, cambia en los dos. Por
   la misma razón, sin sprints cerrados la velocidad promedio dice "sin sprints cerrados"
   (RN-11), como el dashboard, y no 0: un cero haría parecer que el equipo no hizo nada.
5. **El día de generación lo indica quien pide el reporte**, no el sistema: el contrato pide
   `p.SprintReport(number, today)` y `p.ProjectReport(today)`, porque el dominio no lee el
   reloj (AGENTS.md). La web le pasa el día de hoy, y los tests una fecha fija.
6. **Se puede generar el reporte de un sprint abierto (RN-7)**, avisando que los números son
   parciales: sirve para mostrar cómo viene el sprint. La alternativa era rechazarlo.

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: Emiliano Sorato — Fecha: 2026-10-27 — Comentarios: aprobada en la issue #18 (https://github.com/emisorato1/seguimiento-medicion-ics/issues/18#issuecomment-6445539317), las tres puertas juntas sobre 505bc4d, con cuatro correcciones de texto antes del código que salen de la revisión de Juliana Bustos en el PR #181. Se aplicaron el 27/10 en d3fd46b: la fecha con el formato 02/01/2006 y format.go en Archivos afectados, los textos exactos con tilde, "sin sprints cerrados" según haya algún sprint cerrado, y sin "o la #17".

## Fase 2 — Diseño

Las decisiones transversales están en [AGENTS.md](../../../AGENTS.md) y no se repiten. Los
nombres son los del [contrato del dominio](../../diseno-dominio.md): `Report`, con `Title`,
`GeneratedOn`, `Scope` (`ScopeProject`, `ScopeSprint`), `Metrics` y `Sections`, pedido con
`ProjectReport` y `SprintReport`. Cada sección es una `Section`, con `Heading` y `Lines`
(textos ya armados). Con el dashboard (#17) se comparte **la regla de formato de los números**,
no la sección (cambio de contrato del 20/10, cuando el dashboard pasó a la web).

### Enfoque

El reporte se arma de una sola vez, en tres pasos: **pedir las métricas** del mismo alcance a
la #16, **elegir qué se lista** —las historias y los defectos del sprint o del proyecto— y
**escribir cada sección** línea por línea con las reglas de formato de la RN-14.

El armado **recibe las métricas ya calculadas**: no las pide él. Las dos operaciones del
contrato solo piden las métricas a la #16 y se las pasan. Así el reporte no tiene con qué
recalcular un número (RN-2), y los tests del formato pueden pasarle métricas armadas a mano
con los valores exactos de cada escenario.

Los dos reportes comparten las piezas que escriben cada sección; lo que cambia es el alcance:
qué historias, qué defectos y qué va en el Resumen.

### Archivos afectados

| Archivo | Se crea o se modifica | Para qué |
|---|---|---|
| `internal/domain/report.go` | se crea | el reporte, el armado de sus secciones y las dos operaciones |
| `internal/domain/report_test.go` | se crea | los tests de los escenarios BDD de la Fase 1 |
| `internal/domain/format.go` | se modifica | se le suma la fecha, en la tarea 4, que es la primera que la usa. Su test va en `report_test.go`: `format.go` no tiene archivo de test propio, y sus números los prueba la #17 en `dashboard_test.go` |

**`internal/domain/format.go` es de la #17** y ya está en `dev` (PR #204): la #18 usa sus
números tal como están. **Las fechas no las tiene**, porque en la #17 las escribe la web.
Esta historia las suma, con el mismo formato que la web (`02/01/2006`, día y mes con dos
dígitos), para que una fecha se lea igual en la pantalla, en el reporte y en el PDF. La #17
no cambia.

No se agregan errores: se reutiliza el de sprint no encontrado (#8). **El proyecto que no
existe lo informa el registro (#3)**, con su error de proyecto no encontrado: los dos
reportes se le piden a un proyecto, y sin proyecto no hay a quién pedírselos. Es lo mismo
que hacen la #8, la #11 y el contrato del Dashboard.

### Decisiones

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| El armado recibe las métricas calculadas; las operaciones públicas se las piden a la #16 | que el armado las pida él mismo | garantiza la RN-2: el reporte no tiene con qué recalcular un número, y los tests del formato le pasan valores exactos |
| El reporte guarda las métricas y además las líneas de texto | guardar solo las líneas | lo pide el contrato, y así un test lee números sin interpretar texto |
| Los números se escriben como texto acá, con reglas fijas | que cada consumidor los formatee | la #19 solo arma páginas; si cada uno formateara, el PDF y la pantalla podrían mostrar números distintos |
| Siempre las cuatro secciones, aunque estén vacías | omitir las vacías | el reporte tiene siempre la misma forma, y "no hay defectos" también es información |
| La fecha de generación la recibe la operación como dato de quien llama | sacarla del reloj del sistema | decidido por el Agile Enabler el 20/10 y escrito en AGENTS.md: el dominio no lee el reloj, así los tests fijan la fecha que necesitan |
| La situación de cada historia en un sprint sale del registro del cierre | del estado actual de la historia | ver el punto 2 de la puerta |
| El formato de los números y las fechas va en `format.go`, compartido con la #17, y lo declara la primera de las dos que llegue a `dev` | que cada uno formatee a su manera, o que la #18 espere a la #17 | un número tiene que leerse igual en el dashboard, en el reporte y en el PDF; y esperar a la #17 frenaría también a la #19. Es la misma regla que se usó con `errors.go` |
| **El reporte y sus secciones son datos simples**, con sus datos a la vista (título, día de generación, alcance, métricas y secciones; cada sección con su encabezado y sus líneas) | entregarlos detrás de operaciones de lectura | es lo que ya son `Metrics` y el Dashboard: el reporte no tiene nada que proteger, porque se arma de una vez y no se guarda (RN-15). Así la web (#78) y el PDF (#19) lo leen sin pasos extra |
| **"Sin sprints cerrados" sale de si el proyecto tiene algún sprint cerrado** (la consulta de la #11), no de que la velocidad promedio sea cero | mirar si la velocidad es 0 | con sprints cerrados, una velocidad promedio de 0 se escribe "0": el número solo no alcanza para decidir el aviso. Se resuelve igual que la tarea 6 de la #17, coordinado con Juliana Bustos |
| **El reporte se entrega como un valor**, y la exportación a PDF de la #19 lo recibe igual | entregar una referencia al reporte | cada pedido arma uno nuevo y nadie lo comparte, así que no hay nada que se pueda cambiar de costado. Acordado con Emiliano Sorato en el #181, para que el PDF lo lea de la misma forma |
| **El reporte del proyecto no informa errores; el de un sprint, solo el de sprint no encontrado (#8)** | que los dos informen también el proyecto que no existe | los dos se le piden a un proyecto que ya se encontró: el proyecto que no existe lo informa el registro (#3), como en el Dashboard. Un sprint sí puede no existir dentro del proyecto (RN-16) |

### Dependencias y restricciones

- **La #16, la #8, la #10 y la #15 ya están en `dev`.** Los tests arman los sprints, sus
  cierres y los defectos con sus operaciones reales, y las tareas 2 y 3 le pasan al armado
  métricas armadas a mano para fijar valores exactos.
- **No se calcula ninguna métrica** ni se cambia ningún dato.
- **La regla de formato es la de la #17** en `format.go` (ver Archivos afectados), y esta
  historia solo le suma la fecha. Hereda sus dos bordes del redondeo, anotados en la Fase 4
  de la #17: el empate se redondea hacia afuera del cero (12,25 h da "12,3 h") y un negativo
  que se redondea a cero va sin signo (-0,04 % da "0,0 %"). Ningún test espera otra cosa.
  Cualquier cambio se avisa a las dos historias.
- **Lo necesita la #19**, que arma un reporte a mano en sus tests. Si cambia el formato de
  las secciones, se avisa.

> **Puerta 2** — ¿el diseño es viable y coherente con el resto del proyecto?
> Aprobó: Emiliano Sorato — Fecha: 2026-10-27 — Comentarios: aprobada en la issue #18 (https://github.com/emisorato1/seguimiento-medicion-ics/issues/18#issuecomment-6445539317), las tres puertas juntas sobre 505bc4d, con cuatro correcciones de texto antes del código que salen de la revisión de Juliana Bustos en el PR #181. Se aplicaron el 27/10 en d3fd46b: la fecha con el formato 02/01/2006 y format.go en Archivos afectados, los textos exactos con tilde, "sin sprints cerrados" según haya algún sprint cerrado, y sin "o la #17".

## Fase 3 — Tareas

Cada tarea es un ciclo de TDD completo: el test que falla en un commit, la implementación que
lo hace pasar en otro. Están en el orden en que hay que hacerlas y cada una deja los tests en
verde. Siguen la convención de [AGENTS.md](../../../AGENTS.md): un test por escenario, con su
nombre, y cada esquema como un solo test con tabla de casos. Los textos de las líneas se
comparan completos: acá el texto **es** el resultado, no un mensaje de error.

| # | Tarea | Escenario BDD | Reglas | Estado |
|---|---|---|---|---|
| 1 | Armar el reporte de un sprint: título, alcance, día de generación y las cuatro secciones en orden | Generar el reporte de un sprint cerrado | RN-1, RN-3, RN-4 | pendiente |
| 2 | La sección Estimación y esfuerzo copia los números de las métricas recibidas | Los numeros del reporte son los de las metricas | RN-2, RN-11 | pendiente |
| 3 | Formato de horas, desviación y porcentajes | Formato de los numeros | RN-14 | pendiente |
| 4 | Resumen de un sprint, cerrado o abierto con su aviso | El resumen dice en que estado esta el sprint | RN-6, RN-7 | pendiente |
| 5 | Historias del sprint: totales y una línea por historia, con la situación del cierre | Las historias del sprint figuran segun el cierre de ese sprint | RN-9, RN-10 | pendiente |
| 6 | Una historia sin estimar figura como "sin estimar" | Una historia sin estimar figura como sin estimar | RN-9 | pendiente |
| 7 | Defectos del sprint: los detectados o resueltos en él, con su severidad y su estado | Los defectos del sprint son los detectados o resueltos en el | RN-12 | pendiente |
| 8 | Una sección sin nada que listar lo dice, con sus totales en cero | Una seccion sin datos lo dice | RN-5 | pendiente |
| 9 | Reporte del proyecto: título, resumen, todas las historias con su estado, todos los defectos y velocidad promedio | Generar el reporte del proyecto entero | RN-1, RN-3, RN-8, RN-10, RN-11, RN-12 | pendiente |
| 10 | Un proyecto sin fecha de fin lo dice en el resumen, y sin sprints cerrados la velocidad promedio lo dice | Un proyecto sin fecha de fin / Un proyecto sin sprints cerrados no muestra una velocidad de cero | RN-8, RN-11 | pendiente |
| 11 | Historias y defectos ordenados por identificador | Las historias y los defectos van ordenados por identificador | RN-13 | pendiente |
| 12 | Las operaciones del contrato piden las métricas a la #16: los números coinciden y ningún dato cambia | Generar el reporte de un sprint cerrado / Los numeros del reporte son los de las metricas | RN-2, RN-15 | pendiente |
| 13 | Un sprint inexistente lo informa el reporte; un proyecto inexistente, el registro | Pedir un reporte de algo que no existe | RN-16 | pendiente |
| 14 | Las historias de un sprint abierto figuran con sus puntos actuales y "en curso" | Las historias de un sprint abierto figuran en curso | RN-10 | pendiente |

### Por qué este orden

- **La 1 va primero** porque fija la forma del reporte, que verifican todas las demás.
- **La 2 y la 3 son los números**: copiarlos y escribirlos bien, antes de las secciones que
  los usan.
- **De la 4 a la 8 se completan las secciones del sprint**, una por vez, y al final el caso de
  la sección vacía.
- **La 9 a la 11 agregan el proyecto**, reutilizando las piezas del sprint.
- **La 12 va casi al final** porque conecta las operaciones del contrato con la #16 sobre todo
  lo anterior; **la 13**, los bordes de afuera.

### Cobertura

Las dieciséis reglas quedan cubiertas: RN-1 (1, 9), RN-2 (2, 12), RN-3 (1, 9), RN-4 (1),
RN-5 (8), RN-6 (4), RN-7 (4), RN-8 (9, 10), RN-9 (5, 6), RN-10 (5, 9, 14), RN-11 (2, 9, 10),
RN-12 (7, 9), RN-13 (11), RN-14 (3), RN-15 (12), RN-16 (13).

| Criterio de aceptación | Tareas |
|---|---|
| Reporte de un sprint y del proyecto, con título, día y alcance | 1, 9 |
| Cuatro secciones en orden, y una sección vacía dice que no hay | 1, 8 |
| Los números coinciden con las métricas | 2, 12 |
| Resumen de un sprint, con aviso si está abierto | 4 |
| Resumen del proyecto | 9, 10 |
| Historias planificadas y completadas, con puntos o "sin estimar" | 5, 6 |
| Situación de cada historia según el cierre del sprint, o en curso si sigue abierto | 5, 14 |
| Estimaciones y esfuerzo registrado | 2 |
| Defectos del período, con severidad y estado | 7, 9 |
| Formato de horas, porcentajes, desviación, velocidad y fechas | 3, 4, 10 |
| Generar no cambia ningún dato | 12 |
| Sprint o proyecto inexistentes | 13 |

Los catorce escenarios BDD de la Fase 1 tienen su tarea en la columna "Escenario BDD".

### Dependencias

- **La #16, la #8, la #10 y la #15 ya están en `dev`**: ninguna tarea espera a otra historia
  de dominio.
- **La 3, la 4 y la 10 usan la regla de formato de `format.go`** de la #17, que ya está en
  `dev`. La 4 le suma la fecha, con su test en `report_test.go`.
- **La 11 prioriza el backlog con la #7**, para que el orden en que se muestra no coincida
  con el de los identificadores.
- **La 13 usa el error de sprint no encontrado de la #8** y el de proyecto no encontrado de
  la #3.

> **Puerta 3** — ¿las tareas cubren todos los criterios de aceptación?
> Aprobó: Emiliano Sorato — Fecha: 2026-10-27 — Comentarios: aprobada en la issue #18 (https://github.com/emisorato1/seguimiento-medicion-ics/issues/18#issuecomment-6445539317), las tres puertas juntas sobre 505bc4d, con cuatro correcciones de texto antes del código que salen de la revisión de Juliana Bustos en el PR #181. Se aplicaron el 27/10 en d3fd46b: la fecha con el formato 02/01/2006 y format.go en Archivos afectados, los textos exactos con tilde, "sin sprints cerrados" según haya algún sprint cerrado, y sin "o la #17".

## Fase 4 — Implementación

_Se completa durante la implementación._

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #18 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | `internal/domain/report_test.go` |
| Código Go | `internal/domain/report.go` |
