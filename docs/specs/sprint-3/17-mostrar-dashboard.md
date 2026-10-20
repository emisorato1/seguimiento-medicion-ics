# Spec: Mostrar el dashboard

- **Historia:** #17
- **Sprint:** 3
- **Responsable:** Juliana Bustos
- **Estado:** Borrador completo — pendiente de las tres puertas
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de la consigna y
  de los criterios de aceptación de la issue. El responsable lo revisa, lo corrige si hace
  falta y lo presenta a las puertas.
- **Depende de:** #16 (las métricas ya calculadas), #8 y #10 (el sprint abierto y los
  cerrados) y #3 (el proyecto). No espera a ninguna: se escribe contra el contrato del
  dominio
- **La necesitan:** nadie todavía; la interfaz es la que lo muestra

## Fase 1 — Requisitos

### Objetivo

Que el equipo entienda de un vistazo cómo viene un proyecto: sus números principales, el
sprint en curso y unos gráficos que comparan lo planificado con lo real, sin leer tablas.

### Qué queda afuera

- **Calcular las métricas**: es la #16. El dashboard las recibe calculadas y las presenta.
- **Los integrantes y el resumen del backlog**: son el estado del proyecto (#5).
- **El reporte y su PDF** (#18 y #19), y **el detalle de un sprint cerrado** (#11).
- **Gráficos con imágenes, colores o ventanas**: son barras de caracteres (punto 1 de "Para
  decidir en la puerta"). Tampoco se elige qué se grafica ni el ancho.
- **Imprimirlo en pantalla**: el comando de la consola es de la interfaz. Esta historia
  entrega el dashboard armado, línea por línea.

### Entradas

| Dato | Obligatorio | Descripción |
|---|---|---|
| Proyecto | sí | de qué proyecto se quiere ver el dashboard |

Las métricas no las escribe nadie: el sistema las toma del cálculo de métricas (#16).

### Salidas

El dashboard del proyecto, listo para mostrar, en cuatro partes y en este orden: el
encabezado, el estado general, el sprint en curso y los gráficos. Pedirlo no cambia nada.

### Reglas de negocio

1. **[RN-1]** Cuando se pide el dashboard de un proyecto, el sistema debe armarlo en cuatro
   partes, en este orden: el encabezado, el estado general, el sprint en curso y los gráficos.
2. **[RN-2]** El sistema debe mostrar las métricas tal como las entrega el cálculo de
   métricas, sin volver a calcular ninguna: lo único que hace con ellas es redondearlas para
   mostrarlas (RN-15).
3. **[RN-3]** El sistema debe mostrar en el encabezado el nombre del proyecto, su fecha de
   inicio y su fecha de finalización, o que no tiene fecha de finalización si no la tiene.
4. **[RN-4]** El sistema debe mostrar en el estado general las métricas del proyecto entero:
   story points planificados y completados, historias planificadas y completadas con el
   porcentaje de completas, horas estimadas y reales con su desviación, y defectos detectados
   y resueltos.
5. **[RN-5]** Mientras el proyecto tenga un sprint abierto, el sistema debe mostrar su número,
   su Sprint Goal, sus fechas, sus story points planificados, sus horas reales registradas y
   sus defectos detectados, y no debe mostrar sus puntos completados ni su velocidad, que se
   registran recién al cerrarlo.
6. **[RN-6]** Mientras el proyecto tenga sprints pero ninguno abierto, el sistema debe
   informar que no hay un sprint en curso y armar igual el resto del dashboard.
7. **[RN-7]** Mientras el proyecto no tenga ningún sprint, el sistema debe mostrar solo el
   encabezado y el aviso de que todavía no hay datos para mostrar, sin estado general, sin
   sprint en curso y sin gráficos, y eso no es un error.
8. **[RN-8]** El sistema debe mostrar cuatro gráficos, en este orden: story points
   planificados contra completados, horas estimadas contra reales, defectos detectados contra
   resueltos, y la velocidad de cada sprint cerrado.
9. **[RN-9]** El sistema debe mostrar en el gráfico de velocidad una barra por sprint cerrado,
   ordenadas por número, del más viejo al más nuevo. Mientras no haya ningún sprint cerrado,
   el gráfico debe informar que todavía no hay sprints cerrados.
10. **[RN-10]** El sistema debe dibujar cada gráfico con barras horizontales hechas de
    caracteres, una por valor, cada una con su rótulo y su valor al lado.
11. **[RN-11]** Cuando se dibuja un gráfico, el sistema debe darle 20 caracteres a la barra
    del valor más grande de ese gráfico, y a cada una de las demás un largo proporcional a su
    valor, redondeado al entero más cercano y con las mitades hacia arriba.
12. **[RN-12]** Si un valor es cero, el sistema no debe dibujarle ningún carácter de barra, y
    debe mostrar igual su rótulo y su valor.
13. **[RN-13]** Mientras un valor sea mayor que cero, el sistema debe dibujarle al menos un
    carácter de barra, aunque su proporción redondee a cero.
14. **[RN-14]** Si todos los valores de un gráfico son cero, el sistema debe mostrarlo con sus
    rótulos y sus valores y sin ninguna barra, y eso no es un error.
15. **[RN-15]** El sistema debe mostrar los porcentajes y las horas con un decimal, coma
    decimal y su unidad (% o h); los story points, las historias y los defectos, sin
    decimales; y la desviación, con su signo cuando es negativa.
16. **[RN-16]** El sistema debe armar el dashboard sin modificar ningún dato del proyecto.
17. **[RN-17]** Si se pide el dashboard de un proyecto que no existe, el sistema debe informar
    que no se encontró el proyecto y no armar ningún dashboard.

### Casos límite

- **Un único sprint, abierto:** se ve todo; el gráfico de velocidad avisa que todavía no
  hay sprints cerrados.
- **Métricas en cero**, por ejemplo con un sprint recién abierto: gráficos sin barras.
- **Más completado que planificado:** la barra de 20 es la de completados.
- **Horas con decimales:** la barra se calcula con el valor que llega; el redondeo a un
  decimal es solo para escribirlo.

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| El proyecto no existe | informa que no se encontró el proyecto y no arma ningún dashboard |
| El proyecto no tiene sprints | **no es un error**: muestra el aviso de que no hay datos |
| No hay sprint abierto | **no es un error**: lo informa en esa parte |
| Un gráfico tiene todos sus valores en cero | **no es un error**: se muestra sin barras |

### Criterios de aceptación

- [ ] El dashboard tiene cuatro partes en orden: encabezado, estado general, sprint en curso
      y gráficos.
- [ ] El encabezado muestra el nombre y las fechas, o que no hay fecha de finalización.
- [ ] El estado general muestra los mismos números que entregó el cálculo de métricas.
- [ ] Con un sprint abierto se ven su número, su goal, sus fechas, sus story points
      planificados, sus horas reales y sus defectos, y no se ven completados ni velocidad.
- [ ] Sin sprint abierto, se informa que no hay sprint en curso y el resto se ve igual.
- [ ] Un proyecto sin sprints muestra el encabezado y el aviso, sin gráficos y sin error.
- [ ] Se muestran cuatro gráficos: story points, horas, defectos y velocidad.
- [ ] El de velocidad tiene una barra por sprint cerrado, ordenadas por número, sin el abierto.
- [ ] Cada barra tiene su rótulo y su valor al lado.
- [ ] La barra del valor más grande ocupa exactamente 20 caracteres, y las demás son
      proporcionales, con las mitades redondeadas hacia arriba.
- [ ] La barra de un valor cero no ocupa ningún carácter.
- [ ] Un valor mayor que cero ocupa al menos un carácter, y un gráfico con todos sus valores
      en cero se muestra sin barras y sin error.
- [ ] Porcentajes y horas van con un decimal y coma; la desviación negativa, con su signo.
- [ ] Pedir el dashboard no cambia nada del proyecto.
- [ ] Pedir el dashboard de un proyecto que no existe informa que no se encontró.

### Escenarios BDD

```gherkin
Escenario: Dashboard de un proyecto con un sprint en curso
  Dado un proyecto "Sistema de metricas" con el sprint 1 cerrado y el sprint 2 abierto
  Y sus metricas ya calculadas
  Cuando pido el dashboard del proyecto
  Entonces veo en este orden el encabezado, el estado general, el sprint en curso y los graficos
  Y los graficos son, en este orden, story points, horas, defectos y velocidad

Escenario: Encabezado de un proyecto sin fecha de finalizacion
  Dado un proyecto "Sistema de metricas" con inicio 01/10/2026, sin fecha de finalizacion y con un sprint abierto
  Cuando pido el dashboard del proyecto
  Entonces el encabezado muestra el nombre "Sistema de metricas" y el inicio 01/10/2026
  Y muestra que no tiene fecha de finalizacion

Escenario: El estado general muestra las metricas tal como se calcularon
  Dado un proyecto con sprints cuyas metricas son 40 story points planificados y 30 completados
  Y 10 historias planificadas y 7 completadas
  Y 40 horas estimadas y 45 reales
  Y 5 defectos detectados y 3 resueltos
  Cuando pido el dashboard del proyecto
  Entonces el estado general muestra exactamente esos valores

Esquema del escenario: Porcentajes, horas y enteros se muestran con su formato
  Dado un proyecto con sprints cuyas metricas tienen <metrica> igual a <calculado>
  Cuando pido el dashboard del proyecto
  Entonces el estado general muestra <metrica> como <mostrado>

  Ejemplos:
    | metrica                             | calculado | mostrado  |
    | "porcentaje de historias completas" | 66,6666   | "66,7 %"  |
    | "porcentaje de historias completas" | 100       | "100,0 %" |
    | "horas reales"                      | 12,5      | "12,5 h"  |
    | "horas estimadas"                   | 40        | "40,0 h"  |
    | "desviacion de horas"               | -12,5     | "-12,5 %" |
    | "desviacion de horas"               | 0         | "0,0 %"   |
    | "story points completados"          | 30        | "30"      |

Escenario: El sprint en curso muestra su meta y lo que lleva registrado
  Dado un proyecto con el sprint 2 abierto, con el goal "Tener el dashboard" y 13 story points planificados
  Cuando pido el dashboard del proyecto
  Entonces el sprint en curso muestra el sprint 2, su goal y sus fechas
  Y muestra 13 story points planificados, sus horas reales y sus defectos detectados
  Y no muestra puntos completados ni velocidad

Escenario: Un proyecto con todos sus sprints cerrados no tiene sprint en curso
  Dado un proyecto cuyos sprints 1 y 2 estan cerrados
  Cuando pido el dashboard del proyecto
  Entonces se informa que no hay un sprint en curso
  Y se muestra igual el estado general

Escenario: Un proyecto sin sprints avisa que no hay datos
  Dado un proyecto registrado sin ningun sprint
  Cuando pido el dashboard del proyecto
  Entonces veo el encabezado del proyecto
  Y se informa que todavia no hay datos para mostrar
  Y no se muestra ningun grafico
  Y no se informa ningun error

Esquema del escenario: El largo de cada barra es proporcional al valor mas grande del grafico
  Dado un grafico cuyo valor mas grande es <maximo>
  Cuando se dibuja la barra de un valor <valor>
  Entonces la barra ocupa <caracteres> caracteres

  Ejemplos:
    | maximo | valor | caracteres |
    | 40     | 40    | 20         |
    | 40     | 10    | 5          |
    | 8      | 5     | 13         |
    | 50     | 12,5  | 5          |
    | 100    | 1     | 1          |
    | 40     | 0     | 0          |

Esquema del escenario: Las barras de un grafico se miden contra su valor mas grande
  Dado el grafico de story points con <planificados> planificados y <completados> completados
  Cuando se dibuja el grafico
  Entonces la barra de "Planificados" ocupa <barra planificados> caracteres
  Y la barra de "Completados" ocupa <barra completados> caracteres
  Y cada barra muestra su rotulo y su valor, aunque el valor sea cero
  Y no se informa ningun error

  Ejemplos:
    | planificados | completados | barra planificados | barra completados | caso                    |
    | 40           | 30          | 20                 | 15                | normal                  |
    | 20           | 25          | 16                 | 20                | mas de lo planificado   |
    | 40           | 40          | 20                 | 20                | valores iguales         |
    | 3            | 0           | 20                 | 0                 | un valor en cero        |
    | 0            | 0           | 0                  | 0                 | todo el grafico en cero |

Escenario: El grafico de velocidad tiene una barra por sprint cerrado
  Dado un proyecto con los sprints 1, 2 y 3 cerrados, con velocidades 8, 13 y 10
  Cuando pido el dashboard del proyecto
  Entonces el grafico de velocidad tiene las barras "Sprint 1", "Sprint 2" y "Sprint 3", en ese orden
  Y esas barras ocupan 12, 20 y 15 caracteres

Escenario: Todavia no hay sprints cerrados para el grafico de velocidad
  Dado un proyecto cuyo unico sprint es el 1 y esta abierto
  Cuando pido el dashboard del proyecto
  Entonces el grafico de velocidad informa que todavia no hay sprints cerrados
  Y el sprint 1 no aparece como barra en ese grafico

Escenario: Pedir el dashboard no cambia el proyecto
  Dado un proyecto con el sprint 1 abierto y 3 historias asignadas
  Cuando pido el dashboard del proyecto
  Entonces el sprint 1 sigue abierto con sus 3 historias
  Y el backlog sigue igual

Escenario: Pedir el dashboard de un proyecto que no existe
  Dado que no existe ningun proyecto con identificador 99
  Cuando pido el dashboard del proyecto 99
  Entonces se informa que no se encontro el proyecto
  Y no recibo ningun dashboard
```

**Para decidir en la puerta:**

1. **"Gráfico" quiere decir barras horizontales dibujadas con caracteres.** Es una decisión de
   alcance: la aplicación es de consola y el proyecto no usa librerías externas. Cada barra
   lleva su valor al lado para no perder precisión. Imágenes o una página web necesitan una
   librería o una interfaz nueva: sería otra historia. El porqué completo está en la Fase 2.
2. **Barras de 20 caracteres como máximo (RN-11), coma decimal y un decimal (RN-15).** Los 20
   entran en una consola de 80 columnas junto con el rótulo y el valor; la coma es como se
   escribe en castellano. Si se elige otra cosa, solo cambian los ejemplos.
3. **Cuatro gráficos, no tres.** Los tres que comparan lo planificado con lo real, más el de
   velocidad por sprint, que es el clásico de Scrum y la forma en que se ven los sprints
   cerrados (lo que la #11 espera del dashboard). La desviación y el porcentaje de completas
   van como número: la desviación puede ser negativa, y una barra no la representa.
4. **Del sprint en curso no se muestran completados ni velocidad (RN-5).** Se registran recién
   al cerrar el sprint (#10), así que mientras está abierto darían cero y parecería que no se
   hizo nada. Si la #16 decide no calcular métricas de un sprint abierto, esa parte queda con
   número, goal y fechas.
5. **Sin sprints no se muestra el estado general (RN-7)**, porque todo daría cero: las horas
   cargadas antes del primer sprint no se ven hasta que haya uno.
6. **No muestra integrantes ni el resumen del backlog**: eso es el estado del proyecto (#5).
   Si el equipo los quiere acá, se toman de la #5 y esta historia pasa a depender de ella.
7. **El contrato del dominio no tiene una entrada para el dashboard.** Se propone que salga
   armado igual que el reporte de la #18, en partes con un título y sus líneas. Es un cambio
   de contrato, y se avisa al grupo antes, como pide el propio contrato.

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios: [pendiente]

## Fase 2 — Diseño

Las decisiones transversales del proyecto están en [AGENTS.md](../../../AGENTS.md) y no se
repiten acá. Las métricas y los sprints están en [el contrato del dominio](../../diseno-dominio.md).

### Enfoque

El dashboard es **presentación**: no calcula, no guarda y no cambia nada. Recibe números y
los convierte en texto listo para mostrar, en dos piezas y una conexión:

1. **Dibujar un gráfico**: de rótulos con sus valores, una línea por valor con el rótulo, la
   barra y el valor. Tiene todas las reglas del dibujo (RN-10 a RN-14) y no sabe nada de
   proyectos ni de métricas, así que se prueba sola y primero.
2. **Armar el dashboard**: del proyecto y las métricas ya calculadas —del proyecto entero, del
   sprint abierto y de cada sprint cerrado—, las cuatro partes en orden, cada una con su
   título y sus líneas (RN-1). Los gráficos salen de la primera pieza.
3. **Conectar con la #16**: una operación chica del proyecto que busca los sprints, le pide
   las métricas a la #16 y arma el dashboard. Es lo único que necesita la #16 implementada.

Todos los números se escriben con **una sola regla de formato** (RN-15), para que un valor no
salga escrito de dos maneras. Los tests de los gráficos cuentan los caracteres de cada
barra; los de las partes verifican el rótulo y el valor de cada línea, sin comparar el texto
completo, como hizo la #12 con su mensaje.

### Archivos afectados

| Archivo | Se crea o se modifica | Para qué |
|---|---|---|
| `internal/domain/dashboard.go` | se crea | dibujar un gráfico de barras, armar el dashboard con sus cuatro partes y conectarlo con la #16 |
| `internal/domain/dashboard_test.go` | se crea | los tests de los escenarios BDD de la Fase 1, nombrados como cada escenario |
| el archivo del reporte de la #18 | se crea o se modifica, solo si la #18 no lo hizo antes | declarar la sección con título y líneas que comparten el dashboard y el reporte |
| `docs/diseno-dominio.md` | se modifica | agregar el dashboard al contrato, avisando antes al grupo |

No se modifica `errors.go`: el único error posible es el proyecto inexistente, que ya
informa el registro de la #3.

### Decisiones

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| **Los gráficos son barras horizontales de caracteres**, dibujadas con `#` | gráficos de verdad con una librería, o exportar una imagen o una página web | **decisión de alcance**: la aplicación es de consola y el proyecto no usa dependencias externas (AGENTS.md). Con caracteres se cumple la issue con la biblioteca estándar, y cada barra se verifica contando caracteres. Se usa `#` y no un bloque Unicode porque se ve igual en cualquier consola, incluida la de Windows |
| El dashboard sale del dominio con sus líneas ya armadas, barras incluidas | que el dominio entregue los números y la consola dibuje | las reglas del dibujo son verificables y tienen que tener test; en la consola se perderían al cambiar de interfaz, y la #24 dejó la consola sin reglas propias. **Se aparta de la letra de AGENTS.md ("el núcleo no formatea")**, igual que el reporte del contrato, que también trae líneas: para decidir en la puerta. Si se acepta, AGENTS.md se corrige en el mismo PR, como pasó con la #12 |
| El armado recibe las métricas ya calculadas | que el armado se las pida a la #16 | se escribe y se prueba sin esperar a la #16, con métricas armadas a mano; y como nunca calcula, no puede mostrar un número distinto del de la #16 o del reporte (RN-2). La conexión queda en una operación aparte de pocas líneas |
| Cada gráfico tiene su propia escala | una escala común para los cuatro | horas y defectos no se comparan entre sí: con una escala común, los defectos serían barras de un carácter al lado de las horas. Por eso cada barra lleva su valor (RN-10) |
| Ancho fijo de 20 caracteres | ajustarlo al ancho de la consola | no hay forma portable de saber el ancho sin dependencias, y con un ancho fijo los tests cuentan caracteres exactos |
| Las partes tienen la misma forma que las secciones del reporte | una forma propia del dashboard | la consola muestra dashboard y reporte de la misma manera, y la #19 podría exportar el dashboard sin cambios. El costo es compartir esa forma con la #18: la declara la primera de las dos que llegue, como se hizo con `errors.go` |

### Dependencias y restricciones

- **Depende de la #16, que no está implementada. No se la espera**: el contrato dice que las
  métricas son un dato plano que se puede armar a mano, y así lo hacen los tests de las
  tareas 3 a 9. Solo la tarea 10, que conecta, necesita la #16 en `dev`.
- **Depende de los sprints de la #8 y la #10**: se arman dentro del paquete, como en la #11.
- **El contrato no tiene una entrada para el dashboard.** Solo lo consume la interfaz, pero
  usa la sección del reporte, que es compartida: antes de implementar se agrega al contrato
  y se avisa al grupo, en especial a quien tiene la #18.
- **El contrato no dice qué entrega la #16 para un sprint abierto.** Se supone que sus
  métricas, con los completados en cero; si no las calcula, se ajustan la RN-5 y la tarea 7.
- **No se modifica ningún archivo del dominio que ya exista, ni la consola.**

> **Puerta 2** — ¿el diseño es viable y coherente con el resto del proyecto?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios: [pendiente]

## Fase 3 — Tareas

Cada tarea es un ciclo de TDD completo: el test que falla en un commit, la implementación que
lo hace pasar en otro. Van en el orden en que hay que hacerlas, y cada una deja el código
compilando y los tests en verde.

Los tests siguen la convención de [AGENTS.md](../../../AGENTS.md): un test por escenario, con
el nombre del escenario. Como en la #5 y la #7, cuando dos escenarios salen del mismo código
van en la misma tarea, para que ninguna pase en verde sin código nuevo.

| # | Tarea | Escenario BDD | Reglas | Estado |
|---|---|---|---|---|
| 1 | El largo de una barra es proporcional al valor más grande del gráfico, con las mitades hacia arriba, cero caracteres para un valor cero y al menos uno para cualquier valor mayor que cero | El largo de cada barra es proporcional al valor mas grande del grafico | RN-11, RN-12, RN-13 | pendiente |
| 2 | Un gráfico dibuja una línea por valor con su rótulo, su barra y su valor; un valor cero se muestra sin barra, y un gráfico todo en cero, sin barras y sin error | Las barras de un grafico se miden contra su valor mas grande | RN-10, RN-12, RN-14 | pendiente |
| 3 | El dashboard empieza por el encabezado, con el nombre y las fechas del proyecto o la falta de fecha de finalización; un proyecto inexistente informa que no se encontró | Encabezado de un proyecto sin fecha de finalizacion; Pedir el dashboard de un proyecto que no existe | RN-3, RN-17 | pendiente |
| 4 | Un proyecto sin sprints muestra solo el encabezado y el aviso de que no hay datos, sin error | Un proyecto sin sprints avisa que no hay datos | RN-7 | pendiente |
| 5 | El estado general muestra las métricas del proyecto tal como llegaron | El estado general muestra las metricas tal como se calcularon | RN-2, RN-4 | pendiente |
| 6 | Porcentajes y horas con un decimal, coma y unidad; enteros sin decimales; la desviación negativa, con su signo | Porcentajes, horas y enteros se muestran con su formato | RN-15 | pendiente |
| 7 | La parte del sprint en curso muestra el sprint abierto, sin completados ni velocidad; sin sprint abierto, informa que no hay sprint en curso | El sprint en curso muestra su meta y lo que lleva registrado; Un proyecto con todos sus sprints cerrados no tiene sprint en curso | RN-5, RN-6 | pendiente |
| 8 | Los gráficos empiezan por el de velocidad: una barra por sprint cerrado, ordenadas por número; sin sprints cerrados, el aviso | El grafico de velocidad tiene una barra por sprint cerrado; Todavia no hay sprints cerrados para el grafico de velocidad | RN-8, RN-9 | pendiente |
| 9 | Los gráficos de story points, horas y defectos van antes del de velocidad, y el dashboard queda con sus cuatro partes en orden; pedirlo no cambia nada | Dashboard de un proyecto con un sprint en curso; Pedir el dashboard no cambia el proyecto | RN-1, RN-8, RN-16 | pendiente |
| 10 | El dashboard se pide al proyecto y usa las métricas que calcula la #16 | — (conexión con la #16) | RN-2 | pendiente |

**Dos tests pasan solos** y por eso van con uno que falla: el del proyecto inexistente
(tarea 3), que resuelve el registro de la #3, y el de que pedir el dashboard no cambia nada
(tarea 9). La tarea 10 no tiene escenario: su test, encabezado con `// Tarea 10:`, verifica
que el dashboard de un proyecto muestre los mismos números que la #16 calcula para él.

### Por qué este orden

- **La 1 y la 2 primero**: el dibujo no depende ni del proyecto ni de las métricas. Van
  separadas para que un rojo diga si falla la cuenta de una barra o la línea del gráfico.
- **La 3 crea el dashboard** con su primera parte, y **la 4** es el caso que corta antes de
  todo lo demás.
- **La 5 y la 6 separadas**: qué números van y cómo se escriben. Un rojo en la 6 es formato.
- **La 8 y la 9 al final**, porque usan la 1 y la 2; la 9 cierra con el orden de la RN-1.
- **La 10 es la última** porque es la única que necesita la #16 en `dev`.

### Cobertura

Las diecisiete reglas quedan cubiertas: RN-1 (9), RN-2 (5, 10), RN-3 (3), RN-4 (5),
RN-5 (7), RN-6 (7), RN-7 (4), RN-8 (8, 9), RN-9 (8), RN-10 (2), RN-11 (1), RN-12 (1, 2),
RN-13 (1), RN-14 (2), RN-15 (6), RN-16 (9), RN-17 (3).

| Criterio de aceptación | Tareas |
|---|---|
| Cuatro partes en orden: encabezado, estado general, sprint en curso y gráficos | 9 |
| El encabezado con el nombre y las fechas, o sin fecha de finalización | 3 |
| El estado general con los mismos números del cálculo de métricas | 5, 10 |
| Con un sprint abierto, sus datos y lo registrado, sin completados ni velocidad | 7 |
| Sin sprint abierto, se informa y el resto se muestra igual | 7 |
| Un proyecto sin sprints muestra el aviso, sin gráficos y sin error | 4 |
| Cuatro gráficos: story points, horas, defectos y velocidad | 8, 9 |
| El gráfico de velocidad, una barra por sprint cerrado, sin el abierto | 8 |
| Cada barra con su rótulo y su valor | 2 |
| La barra más grande ocupa 20 caracteres y las demás son proporcionales | 1, 2 |
| Un valor cero no ocupa ningún carácter | 1, 2 |
| Un valor mayor que cero ocupa al menos un carácter; un gráfico todo en cero, sin error | 1, 2 |
| Porcentajes y horas con un decimal y coma; desviación negativa con signo | 6 |
| Pedir el dashboard no cambia nada | 9 |
| Un proyecto inexistente informa que no se encontró | 3 |

Los trece escenarios de la Fase 1 tienen su tarea en la columna "Escenario BDD".

### Dependencias

- **La 1 y la 2 no necesitan nada.**
- **De la 3 a la 9 necesitan métricas, y la #16 no está implementada: no se la espera.** Los
  tests las arman a mano, contra el contrato del dominio.
- **La 4, la 7, la 8 y la 9 necesitan sprints abiertos y cerrados** (#8 y #10): los tests los
  arman dentro del paquete.
- **La 3 necesita la sección del reporte**: si la #18 no la declaró todavía, la declara esta
  tarea, como quede en el contrato, y se avisa al grupo.
- **La 10 es la única que necesita la #16 en `dev`.** Si se demora, la historia no se cierra
  hasta la 10, pero no frena ninguna de las anteriores.

> **Puerta 3** — ¿las tareas cubren todos los criterios de aceptación?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios: [pendiente]

## Fase 4 — Implementación

_Se completa durante la implementación._

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #17 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | `internal/domain/dashboard_test.go` |
| Código Go | `internal/domain/dashboard.go` |
