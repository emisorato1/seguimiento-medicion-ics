# Spec: Mostrar el dashboard

- **Historia:** #17
- **Sprint:** 3
- **Responsable:** Juliana Bustos
- **Estado:** Terminada en el Sprint 3 — **mergeada a `dev`** el 29/10 (PR #209)
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de la consigna y
  de los criterios de aceptación de la issue. El responsable lo revisa, lo corrige si hace
  falta y lo presenta a las puertas. **Adaptado a la web el 20/10** (punto 1 de "Para decidir
  en la puerta").
- **Depende de:** #16 (las métricas ya calculadas), #8 y #10 (el sprint abierto y los
  cerrados), #3 (el proyecto) y #77 (la web: el layout, la barra y la página de "no se
  encontró"). La #16, la #3 y la #77 ya están en `dev`; los sprints se arman en los tests,
  contra el contrato del dominio
- **La necesitan:** nadie todavía

## Fase 1 — Requisitos

### Objetivo

Que el equipo entienda de un vistazo cómo viene un proyecto, en una página de la web: sus
números principales, el sprint en curso y unos gráficos que comparan lo planificado con lo
real, sin leer tablas.

### Qué queda afuera

- **Calcular las métricas**: es la #16. El dashboard las recibe calculadas y las presenta.
- **Los integrantes y el resumen del backlog**: son el estado del proyecto (#5).
- **El reporte y su PDF** (#18 y #19), y **el detalle de un sprint cerrado** (#11).
- **Elegir qué se grafica**, cambiar el tipo de gráfico, acercar o pasar el mouse para ver
  más: los gráficos son siempre los mismos y ya muestran sus valores.
- **Mostrarlo en la consola**: la consola (#24) no lo ofrece (punto 8 de "Para decidir").

### Entradas

| Dato | Obligatorio | Descripción |
|---|---|---|
| Proyecto | sí | de qué proyecto se quiere ver el dashboard |

Las métricas no las escribe nadie: el sistema las toma del cálculo de métricas (#16).

### Salidas

La página del dashboard del proyecto, en cuatro partes y en este orden: el encabezado, el
estado general, el sprint en curso y los gráficos. Abrirla no cambia nada.

### Reglas de negocio

1. **[RN-1]** Cuando se abre el dashboard de un proyecto, el sistema debe mostrarlo en cuatro
   partes, en este orden: el encabezado, el estado general, el sprint en curso y los gráficos.
2. **[RN-2]** El sistema debe mostrar las métricas tal como las entrega el cálculo de
   métricas, sin volver a calcular ninguna: lo único que hace con ellas es redondearlas para
   mostrarlas (RN-15).
3. **[RN-3]** El sistema debe mostrar en el encabezado el nombre del proyecto, su fecha de
   inicio y su fecha de finalización, o «sin fecha de fin» si no la tiene.
4. **[RN-4]** El sistema debe mostrar en el estado general las métricas del proyecto entero:
   story points planificados y completados, historias planificadas y completadas con el
   porcentaje de completas, horas estimadas y reales con su desviación, defectos detectados
   y resueltos, y la velocidad del equipo, que es el promedio de los sprints cerrados.
   Mientras no haya ningún sprint cerrado, en lugar de la velocidad debe mostrar «sin sprints
   cerrados», como el gráfico de velocidad (RN-9).
5. **[RN-5]** Mientras el proyecto tenga un sprint abierto, el sistema debe mostrar su número,
   su Sprint Goal, sus fechas, sus story points planificados, sus horas reales registradas y
   sus defectos detectados, y no debe mostrar sus puntos completados ni su velocidad, que se
   registran recién al cerrarlo.
6. **[RN-6]** Mientras el proyecto tenga sprints pero ninguno abierto, el sistema debe
   informar que no hay un sprint en curso y mostrar igual el resto del dashboard.
7. **[RN-7]** Mientras el proyecto no tenga ningún sprint, el sistema debe mostrar solo el
   encabezado y el aviso de que todavía no hay datos para mostrar, sin estado general, sin
   sprint en curso y sin gráficos, y eso no es un error.
8. **[RN-8]** El sistema debe mostrar cuatro gráficos, en este orden: story points
   planificados contra completados, horas estimadas contra reales, defectos detectados contra
   resueltos, y la velocidad de cada sprint cerrado.
9. **[RN-9]** El sistema debe mostrar en el gráfico de velocidad una barra por sprint cerrado,
   ordenadas por número, del más viejo al más nuevo. Mientras no haya ningún sprint cerrado,
   el gráfico debe informar que todavía no hay sprints cerrados.
10. **[RN-10]** El sistema debe dibujar cada gráfico con barras horizontales, una por valor,
    cada una con su rótulo y su valor escrito.
11. **[RN-11]** Cuando se dibuja un gráfico, el sistema debe darle todo el ancho del gráfico a
    la barra del valor más grande, y a cada una de las demás la parte del ancho proporcional a
    su valor.
12. **[RN-12]** Si un valor es cero, el sistema no debe dibujarle barra, y debe mostrar igual
    su rótulo y su valor.
13. **[RN-13]** Mientras un valor sea mayor que cero, el sistema debe darle a su barra al
    menos el 1 % del ancho, aunque su proporción sea menor, para que se vea.
14. **[RN-14]** Si todos los valores de un gráfico son cero, el sistema debe mostrarlo con sus
    rótulos y sus valores y sin ninguna barra, y eso no es un error.
15. **[RN-15]** El sistema debe mostrar los porcentajes y las horas con un decimal, coma
    decimal y su unidad (% o h); los story points, las historias, los defectos y la
    velocidad, sin decimales; la desviación, con su signo cuando es negativa y sin signo
    cuando es positiva o cero; y las fechas del encabezado y
    del sprint en curso como día/mes/año, por ejemplo 01/10/2026. Es la misma regla del
    reporte (#18, RN-14) y de las fechas del resto de la web (#77, RN-5): un número o una fecha
    se lee igual en el dashboard y en el reporte.
16. **[RN-16]** El sistema debe mostrar el dashboard sin modificar ningún dato del proyecto.
17. **[RN-17]** Si se pide el dashboard de un proyecto que no existe, el sistema debe informar
    que no se encontró el proyecto y ofrecer volver a la lista de proyectos, como el resto de
    la web (#77, RN-16).
18. **[RN-18]** Mientras se está dentro de un proyecto, el sistema debe ofrecer en la barra de
    navegación el acceso a su dashboard, junto a sus datos, sus integrantes y su backlog.
19. **[RN-19]** El sistema debe acompañar cada gráfico con un texto alternativo, para quien no
    lo ve, que diga qué compara y cada rótulo con su valor, y debe permitir ver esos mismos
    valores como una tabla.
20. **[RN-20]** El sistema no debe distinguir las barras solo por su color: cada una lleva su
    rótulo. Las barras deben tener un contraste de al menos 3 a 1 con el fondo.
21. **[RN-21]** Como el resto de la web, el dashboard debe verse completo sin conexión a
    internet, y en un celular de 360 puntos de ancho los gráficos deben ocupar el ancho de la
    pantalla sin que la página se desplace de costado (#77, RN-3 y RN-18).

### Casos límite

- **Un único sprint, abierto:** se ve todo; el gráfico de velocidad avisa que todavía no
  hay sprints cerrados, y el estado general muestra la velocidad como «sin sprints cerrados»
  en lugar de un cero.
- **Sprints cerrados con una velocidad promedio de cero:** el estado general muestra «0», no
  el aviso. El aviso depende de que no haya sprints cerrados, no de que el valor sea cero
  (RN-4).
- **Métricas en cero**, por ejemplo con un sprint recién abierto: gráficos sin barras, con
  sus rótulos y sus valores.
- **Más completado que planificado:** la barra que ocupa todo el ancho es la de completados.
- **Horas con decimales:** la barra se mide con el valor que llega; el redondeo a un decimal
  es solo para escribirlo.
- **Un valor muy chico al lado de uno muy grande**, como 1 contra 1000: su barra ocupa el
  1 % del ancho y se ve.
- **Un Sprint Goal con signos** como `<` o `&`: se ve tal cual (#77, RN-17).

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| El proyecto no existe | informa que no se encontró el proyecto y ofrece volver a la lista |
| El proyecto no tiene sprints | **no es un error**: muestra el aviso de que no hay datos |
| No hay sprint abierto | **no es un error**: lo informa en esa parte |
| Un gráfico tiene todos sus valores en cero | **no es un error**: se muestra sin barras |

### Criterios de aceptación

- [ ] El dashboard tiene cuatro partes en orden: encabezado, estado general, sprint en curso
      y gráficos.
- [ ] El encabezado muestra el nombre y las fechas, o «sin fecha de fin».
- [ ] El estado general muestra los mismos números que entregó el cálculo de métricas.
- [ ] El estado general muestra la velocidad promedio del equipo, o «sin sprints cerrados»
      si todavía no se cerró ninguno.
- [ ] Con un sprint abierto se ven su número, su goal, sus fechas, sus story points
      planificados, sus horas reales y sus defectos, y no se ven completados ni velocidad.
- [ ] Sin sprint abierto, se informa que no hay sprint en curso y el resto se ve igual.
- [ ] Un proyecto sin sprints muestra el encabezado y el aviso, sin gráficos y sin error.
- [ ] Se muestran cuatro gráficos: story points, horas, defectos y velocidad.
- [ ] El de velocidad tiene una barra por sprint cerrado, ordenadas por número, sin el abierto.
- [ ] Cada barra tiene su rótulo y su valor.
- [ ] La barra del valor más grande ocupa todo el ancho del gráfico, y las demás son
      proporcionales.
- [ ] Un valor cero no tiene barra; uno mayor que cero ocupa al menos el 1 % del ancho; un
      gráfico con todos sus valores en cero se muestra sin barras y sin error.
- [ ] Porcentajes y horas van con un decimal y coma; la desviación negativa, con su signo;
      las fechas, como día/mes/año.
- [ ] Cada gráfico tiene su texto alternativo con sus valores, y se pueden ver como tabla.
- [ ] Las barras no se distinguen solo por el color y tienen contraste con el fondo.
- [ ] El dashboard se abre desde la barra de navegación del proyecto.
- [ ] Se ve sin internet y en un celular, sin desplazar la página de costado.
- [ ] Abrir el dashboard no cambia nada del proyecto.
- [ ] Pedir el dashboard de un proyecto que no existe informa que no se encontró.

### Escenarios BDD

```gherkin
Escenario: Dashboard de un proyecto con un sprint en curso
  Dado un proyecto "Sistema de metricas" con el sprint 1 cerrado y el sprint 2 abierto
  Y sus metricas ya calculadas
  Cuando abro el dashboard del proyecto
  Entonces veo en este orden el encabezado, el estado general, el sprint en curso y los graficos
  Y los graficos son, en este orden, story points, horas, defectos y velocidad

Escenario: Abrir el dashboard desde la barra de navegacion
  Dado un proyecto "Sistema de metricas"
  Cuando entro a cualquier pagina del proyecto
  Entonces la barra de navegacion tiene el acceso a su dashboard

Escenario: Encabezado de un proyecto sin fecha de finalizacion
  Dado un proyecto "Sistema de metricas" con inicio 01/10/2026, sin fecha de finalizacion y con un sprint abierto
  Cuando abro el dashboard del proyecto
  Entonces el encabezado muestra el nombre "Sistema de metricas" y el inicio 01/10/2026
  Y muestra "sin fecha de fin"

Escenario: El estado general muestra las metricas tal como se calcularon
  Dado un proyecto con sprints cerrados cuyas metricas son 40 story points planificados y 30 completados
  Y 10 historias planificadas y 7 completadas
  Y 40 horas estimadas y 45 reales
  Y 5 defectos detectados y 3 resueltos
  Y una velocidad de 12
  Cuando abro el dashboard del proyecto
  Entonces el estado general muestra exactamente esos valores

Esquema del escenario: Porcentajes, horas y enteros se muestran con su formato
  Dado un proyecto con sprints cuyas metricas tienen <metrica> igual a <calculado>
  Cuando abro el dashboard del proyecto
  Entonces el estado general muestra <metrica> como <mostrado>

  Ejemplos:
    | metrica                             | calculado | mostrado  |
    | "porcentaje de historias completas" | 66,6666   | "66,7 %"  |
    | "porcentaje de historias completas" | 100       | "100,0 %" |
    | "horas reales"                      | 12,5      | "12,5 h"  |
    | "horas estimadas"                   | 40        | "40,0 h"  |
    | "desviacion de horas"               | -12,5     | "-12,5 %" |
    | "desviacion de horas"               | 0         | "0,0 %"   |
    | "desviacion de horas"               | 25        | "25,0 %"  |
    | "velocidad"                         | 12        | "12"      |
    | "story points completados"          | 30        | "30"      |

Escenario: El sprint en curso muestra su meta y lo que lleva registrado
  Dado un proyecto con el sprint 2 abierto del 26/10/2026 al 30/10/2026, con el goal "Tener el dashboard" y 13 story points planificados
  Cuando abro el dashboard del proyecto
  Entonces el sprint en curso muestra el sprint 2, su goal y las fechas 26/10/2026 y 30/10/2026
  Y muestra 13 story points planificados, sus horas reales y sus defectos detectados
  Y no muestra puntos completados ni velocidad

Escenario: Un proyecto con todos sus sprints cerrados no tiene sprint en curso
  Dado un proyecto cuyos sprints 1 y 2 estan cerrados
  Cuando abro el dashboard del proyecto
  Entonces se informa que no hay un sprint en curso
  Y se muestra igual el estado general

Escenario: Un proyecto sin sprints avisa que no hay datos
  Dado un proyecto registrado sin ningun sprint
  Cuando abro el dashboard del proyecto
  Entonces veo el encabezado del proyecto
  Y se informa que todavia no hay datos para mostrar
  Y no se muestra ningun grafico
  Y no se informa ningun error

Esquema del escenario: El largo de cada barra es proporcional al valor mas grande del grafico
  Dado un grafico cuyo valor mas grande es <maximo>
  Cuando se dibuja la barra de un valor <valor>
  Entonces la barra ocupa <ancho> del ancho del grafico

  Ejemplos:
    | maximo | valor | ancho     |
    | 40     | 40    | 100 %     |
    | 40     | 10    | 25 %      |
    | 8      | 5     | 62,5 %    |
    | 50     | 12,5  | 25 %      |
    | 1000   | 1     | 1 %       |
    | 40     | 0     | sin barra |

Esquema del escenario: Las barras de un grafico se miden contra su valor mas grande
  Dado el grafico de story points con <planificados> planificados y <completados> completados
  Cuando abro el dashboard del proyecto
  Entonces la barra de "Planificados" ocupa <barra planificados>
  Y la barra de "Completados" ocupa <barra completados>
  Y cada barra muestra su rotulo y su valor, aunque el valor sea cero
  Y no se informa ningun error

  Ejemplos:
    | planificados | completados | barra planificados | barra completados | caso                    |
    | 40           | 30          | 100 %              | 75 %              | normal                  |
    | 20           | 25          | 80 %               | 100 %             | mas de lo planificado   |
    | 40           | 40          | 100 %              | 100 %             | valores iguales         |
    | 3            | 0           | 100 %              | sin barra         | un valor en cero        |
    | 0            | 0           | sin barra          | sin barra         | todo el grafico en cero |

Escenario: El grafico de velocidad tiene una barra por sprint cerrado
  Dado un proyecto con los sprints 1, 2 y 3 cerrados, con velocidades 8, 13 y 10
  Cuando abro el dashboard del proyecto
  Entonces el grafico de velocidad tiene las barras "Sprint 1", "Sprint 2" y "Sprint 3", en ese orden
  Y esas barras ocupan 61,5 %, 100 % y 76,9 % del ancho

Escenario: Todavia no hay sprints cerrados para el grafico de velocidad
  Dado un proyecto cuyo unico sprint es el 1 y esta abierto
  Cuando abro el dashboard del proyecto
  Entonces el grafico de velocidad informa que todavia no hay sprints cerrados
  Y el sprint 1 no aparece como barra en ese grafico
  Y el estado general muestra la velocidad como "sin sprints cerrados"

Escenario: Cada grafico tiene su texto alternativo y su tabla
  Dado el grafico de story points con 40 planificados y 30 completados
  Cuando abro el dashboard del proyecto
  Entonces el grafico tiene un texto alternativo que dice que compara los story points, con "Planificados: 40" y "Completados: 30"
  Y se pueden ver como una tabla con las filas "Planificados" 40 y "Completados" 30

Escenario: Abrir el dashboard no cambia el proyecto
  Dado un proyecto con el sprint 1 abierto y 3 historias asignadas
  Cuando abro el dashboard del proyecto
  Entonces el sprint 1 sigue abierto con sus 3 historias
  Y el backlog sigue igual

Escenario: Pedir el dashboard de un proyecto que no existe
  Dado que no existe ningun proyecto con identificador 99
  Cuando abro el dashboard del proyecto 99
  Entonces se informa que no se encontro el proyecto
  Y veo el acceso para volver a la lista de proyectos

Escenario: Usar el dashboard en un celular y sin internet
  Dado una pantalla de 360 puntos de ancho y una computadora sin conexion a internet
  Cuando abro el dashboard de un proyecto con sprints
  Entonces los graficos ocupan el ancho de la pantalla y la pagina no se desplaza de costado
  Y la pagina se ve completa, con sus estilos
```

**Para decidir en la puerta:**

1. **El dashboard se muestra en la web, con gráficos de barras de verdad. Se adaptó el
   20/10**, cuando el equipo confirmó la interfaz web (plan final, "Interfaz web"). El
   borrador anterior los dibujaba con caracteres para la consola, porque no había otra
   interfaz. **Juliana, como responsable, lo revisa**: cambian la RN-10 a la RN-14 (las barras
   se miden en parte del ancho y no en caracteres), se suman la RN-17 a la RN-21 (lo propio de
   una página) y la Fase 2 y la Fase 3 se rehicieron. Las demás reglas siguen igual.
2. **Todo el ancho para el valor más grande y al menos el 1 % para cualquier valor mayor que
   cero (RN-11, RN-13)**, en vez de los 20 caracteres de la consola. Así el gráfico se adapta
   a la pantalla, y el mínimo hace que un valor chico no parezca un cero. El formato de los
   números (RN-15) no cambia.
3. **Cuatro gráficos, no tres.** Los tres que comparan lo planificado con lo real, más el de
   velocidad por sprint, que es el clásico de Scrum y la forma en que se ven los sprints
   cerrados (lo que la #11 espera del dashboard). La desviación y el porcentaje de completas
   van como número: la desviación puede ser negativa, y una barra no la representa.
4. **Del sprint en curso no se muestran completados ni velocidad (RN-5).** Se registran recién
   al cerrar el sprint (#10), así que mientras está abierto darían cero y parecería que no se
   hizo nada. La #16 sí calcula las métricas de un sprint abierto, con esos dos en cero.
5. **Sin sprints no se muestra el estado general (RN-7)**, porque todo daría cero: las horas
   cargadas antes del primer sprint no se ven hasta que haya uno.
6. **No muestra integrantes ni el resumen del backlog**: eso es el estado del proyecto (#5).
   Si el equipo los quiere acá, se toman de la #5 y esta historia pasa a depender de ella.
7. **Cambia el contrato del dominio.** Antes el dashboard salía del dominio como una lista de
   secciones con las barras ya dibujadas en las líneas, compartida con el reporte (#18).
   Ahora el dominio entrega **los datos de cada gráfico** —rótulo, valor y valor ya escrito—
   y la web dibuja las barras (Fase 2, D2). La sección queda solo para el reporte. **Ya está
   hecho**: el contrato y AGENTS.md se cambiaron el 20/10 (574a1a3). Se avisa al grupo, en
   especial a quien tiene la #18.
8. **La consola no muestra el dashboard.** La #24 no lo pide y la consigna pide
   "representaciones gráficas", que en una terminal no se defienden bien. Si el equipo lo
   quiere también en la consola, es otra historia.
9. **El texto alternativo y la tabla (RN-19)** hacen que alguien que no ve el gráfico, o que
   usa un lector de pantalla, tenga los mismos números. La tabla se abre desde el gráfico, para
   no repetir en pantalla lo que cada barra ya dice.
10. **La velocidad promedio va en el estado general, y sin sprints cerrados dice «sin sprints
    cerrados» (RN-4).** Lo agregó Juliana el 22/10: la consigna pide la velocidad del equipo y
    el contrato ya la entrega. La #16 la devuelve en cero cuando no hay sprints cerrados, pero
    un cero haría parecer que el equipo no hizo nada, así que el dashboard escribe el mismo
    aviso que el gráfico de velocidad (RN-9). Es lo único que el dashboard escribe distinto de
    como llega, y no es un cálculo: es cómo se muestra un dato que todavía no existe.

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: Agustín Salinas — Fecha: 2026-10-22 — Comentarios: aprobada en el PR #136 (https://github.com/emisorato1/seguimiento-medicion-ics/pull/136#issuecomment-6345137761), con tres menores que no bloquean. Se aplicaron el 23/10 en 9ce7080: el ejemplo de la desviación positiva y la velocidad sin decimales en la RN-15; el punto 7 queda como está, como dijo él.

## Fase 2 — Diseño

Las decisiones transversales del proyecto están en [AGENTS.md](../../../AGENTS.md) y no se
repiten acá. Las métricas y los sprints están en [el contrato del dominio](../../diseno-dominio.md),
y la forma de la web —el layout, la barra, las páginas de "no se encontró" y los tests con
`httptest`— en la [spec de la #77](../sprint-2/77-interfaz-web-base.md).

### Enfoque

El dashboard es **presentación**: no calcula, no guarda y no cambia nada. Tiene dos mitades,
cada una con lo suyo:

1. **El dominio arma los datos del dashboard** (`internal/domain`): del proyecto y de las
   métricas ya calculadas —del proyecto entero, del sprint abierto y de cada sprint cerrado—,
   el estado general, el sprint en curso y los cuatro gráficos. Cada número sale **ya escrito**
   con la regla de formato (RN-15), y cada barra trae además su valor como número, para que la
   web la mida. No dibuja nada.
2. **La web dibuja la página** (`internal/web`): una pantalla nueva con la convención de la
   #77 —un archivo de handlers, su plantilla y su test, y del layout solo la barra—. Muestra
   el encabezado con los datos del proyecto, las partes que entrega el dominio y cada gráfico
   como barras en SVG armado con `html/template`. La única regla propia de la web es **cuánto
   mide cada barra** (RN-11 a RN-14), que se prueba sola.

La conexión con la #16 es una operación chica del proyecto que busca los sprints, le pide
las métricas a la #16 y arma el dashboard. Es lo único que necesita la #16 implementada.

### Archivos afectados

| Archivo | Se crea o se modifica | Para qué |
|---|---|---|
| `internal/domain/dashboard.go` | se crea | armar los datos del dashboard a partir de las métricas, y la operación del proyecto que lo conecta con la #16 |
| `internal/domain/dashboard_test.go` | se crea | los tests de lo que arma el dominio |
| el archivo de la regla de formato | se crea o se modifica, solo si la #18 no lo hizo antes | escribir porcentajes, horas y enteros (RN-15), compartida con el reporte |
| `internal/web/dashboard.go` | se crea | la página del dashboard y el largo de cada barra |
| `internal/web/dashboard_test.go` | se crea | los tests de la página, con `httptest`, nombrados como cada escenario |
| `internal/web/templates/dashboard.html` | se crea | la página, con un SVG por gráfico, su texto alternativo y su tabla |
| `internal/web/templates/layout.html` | se modifica | el acceso al dashboard en la barra del proyecto (RN-18) |
| `internal/web/server.go` | se modifica | la ruta de la página, una línea |
| `internal/web/static/app.css` | se modifica | los estilos de los gráficos |
| `internal/web/server_test.go` | se modifica | sumar el dashboard a las páginas que revisan el layout y las tablas de la #77 |

### Decisiones

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| **D1** Los gráficos son barras horizontales en **SVG armado con `html/template`**, dentro de `internal/web` | una librería de gráficos de JavaScript (Chart.js) por CDN; una imagen PNG armada en Go; barras de caracteres | la web no usa JavaScript ni nada de afuera (#77, D3): la defensa puede ser sin internet. El SVG se adapta al ancho del celular, el texto de los rótulos se lee y se selecciona, y `html/template` escapa los rótulos solo. Una imagen no se adapta y no es accesible; los caracteres eran para la consola |
| **D2** **El dominio entrega los datos de cada gráfico**: para cada barra, su rótulo, su valor como número y su valor ya escrito. La web mide y dibuja | que el dominio entregue líneas con las barras dibujadas (el borrador anterior), o el SVG armado | cuánto mide una barra depende de la pantalla, y eso es de la interfaz. Lo que sí es regla de negocio, **cómo se escribe cada número** (RN-15), queda en el dominio con la misma regla que el reporte: un número no puede leerse distinto en el dashboard y en el reporte. Cambia el contrato y la decisión de AGENTS.md sobre "el núcleo arma texto": el dashboard ya no trae barras |
| **D3** El largo de cada barra es **una parte del ancho del gráfico, en porcentaje con un decimal**: el valor más grande, 100 %; uno en cero, sin barra; uno mayor que cero, al menos 1 %. **El ancho se escribe con punto decimal** (`width="61.5%"`), y no con la regla de formato: es sintaxis del SVG, y con coma el navegador no dibuja la barra. La coma de la RN-15 es solo para el texto que se lee | medir las barras en píxeles fijos | en porcentaje el gráfico ocupa el ancho de cualquier pantalla sin calcular nada más, y el test compara el ancho de la barra en la página, sin navegador |
| **D4** Cada gráfico es una figura con su título, el SVG con `role="img"` y su texto alternativo, que dice qué compara y cada rótulo con su valor, y debajo **la tabla de valores dentro de un desplegable** («Ver los valores») | solo el SVG; o la tabla siempre a la vista | el lector de pantalla lee el texto alternativo y tiene los números (RN-19); la tabla queda para quien la quiere sin repetir en pantalla lo que cada barra ya dice. El desplegable de HTML funciona sin JavaScript |
| **D5** Rótulo y valor escritos arriba de cada barra, el rótulo a la izquierda y el valor a la derecha | el valor al final de la barra | al final de la barra, el valor de una barra de 100 % no tiene lugar y en un celular se superpone. Arriba entran siempre, con cualquier largo |
| **D6** Dos colores para lo planificado y lo real, y uno para la velocidad, definidos en el CSS con contraste de al menos 3 a 1 con el fondo | un color por gráfico, o colores en la plantilla | cada barra ya tiene su rótulo, así que el color ayuda pero no es lo único que las distingue (RN-20); en el CSS se cambian en un solo lugar, igual que los de la #77 |
| **D7** La página es `GET /proyectos/{id}/dashboard` y la barra del proyecto suma «Dashboard» | ponerlo en la página del proyecto | es la convención de la #77: una pantalla, un handler, una plantilla y un acceso en la barra. Solo lee, así que va con el candado de lectura de la #77 (D12) |
| **D8** Las métricas se reciben calculadas | que el armado se las pida a la #16 | se escribe y se prueba sin esperar a la #16, con métricas armadas a mano; y como nunca calcula, no puede mostrar un número distinto del de la #16 o del reporte (RN-2). La conexión queda en una operación aparte de pocas líneas |
| **D9** Cada gráfico tiene su propia escala | una escala común para los cuatro | horas y defectos no se comparan entre sí: con una escala común, los defectos serían barras casi vacías al lado de las horas. Por eso cada barra lleva su valor escrito (RN-10) |
| **D10** Los tests de la web arman los datos del dashboard a mano | esperar a que el dominio los arme | la web y el dominio se pueden hacer a la vez, como la #77 con el guardado falso. Un test de punta a punta, al final, prueba que se conectan |

### Lo que entrega el dominio, según el contrato

Es la entrada del dashboard en el contrato desde el 20/10 (574a1a3), que reemplazó a "una
lista de `Section` con las barras ya dibujadas":

| Dato | Tipo | Qué es |
|---|---|---|
| `Dashboard.HasSprints` | booleano | si el proyecto tiene algún sprint; si no, la página muestra solo el aviso (RN-7) |
| `Dashboard.Overview` | lista de `Figure` | el estado general: cada métrica con su rótulo y su valor escrito (RN-4, RN-15) |
| `Dashboard.CurrentSprint` | `CurrentSprint`, o nada | el sprint abierto, o nada si no hay (RN-5, RN-6) |
| `CurrentSprint` | `Number`, `Goal`, `Start`, `End` y `Figures` (lista de `Figure`) | sus datos y lo que lleva registrado, sin completados ni velocidad. `Start` y `End` son fechas, y las escribe la web como día/mes/año (RN-15), igual que las del encabezado |
| `Dashboard.Charts` | lista de `Chart` | los cuatro gráficos, en el orden de la RN-8 |
| `Chart` | `Title` (texto), `Bars` (lista de `Bar`) y `Empty` (texto) | `Empty` es el aviso cuando el gráfico no tiene barras, como «Todavía no hay sprints cerrados» (RN-9) |
| `Bar` | `Label` (texto), `Value` (decimal) y `Text` (texto) | `Value` es para medir la barra; `Text`, el valor escrito con la RN-15 |
| `Figure` | `Label` (texto) y `Text` (texto) | un número del estado general o del sprint en curso, ya escrito |

Se pide con `p.Dashboard()`, que no falla: el proyecto que no existe lo informa el registro.
El encabezado no está en el contrato porque la web ya tiene el proyecto. **`Section` deja de
ser compartida**: queda solo para el reporte de la #18. La regla de formato sí se comparte:
la escribe la primera de la #17 y la #18 que llegue a `dev`, como pasó con `errors.go`.

### Dependencias y restricciones

- **La #16 ya está en `dev`** desde el 22/10 (PRs #81 y #140), con `p.ProjectMetrics()` y
  `p.SprintMetrics(number)`. Igual, las tareas del dominio arman las métricas a mano (D8), y
  la 8 y la 17, que usan las de verdad, ya no esperan nada.
- **Depende de los sprints de la #8 y la #10**: se arman dentro del paquete, como en la #11.
- **Las tareas de la web necesitan la #77 en `dev`**: el layout, la barra, la página de "no se
  encontró" y los tests. La #77 ya está en `dev` (PR #95, mergeado el 21/10).
- **El contrato ya tiene la entrada nueva del dashboard** (punto 7 de "Para decidir"): se
  cambió el 20/10 (574a1a3).
- **No se modifica la consola.**

> **Puerta 2** — ¿el diseño es viable y coherente con el resto del proyecto?
> Aprobó: Maximiliano Eula — Fecha: 2026-10-22 — Comentarios: aprobada en el PR #136 (https://github.com/emisorato1/seguimiento-medicion-ics/pull/136#issuecomment-6350299170), junto con la Puerta 3, con tres correcciones antes del código. Se aplicaron el 23/10 en 9ce7080: el contrato ya no figura como pendiente, la #16 figura en `dev` y el ancho de las barras va con punto (D3, tarea 13).

## Fase 3 — Tareas

Cada tarea es un ciclo de TDD completo: el test que falla en un commit, la implementación que
lo hace pasar en otro. Van en el orden en que hay que hacerlas, y cada una deja el código
compilando y los tests en verde.

Los tests siguen la convención de [AGENTS.md](../../../AGENTS.md): un test por escenario, con
el nombre del escenario. Los de la web son como los de la #77: `httptest`, un registro nuevo
por test y el texto del error tomado del propio error. Cuando un escenario tiene una parte en
el dominio y otra en la web, cada una tiene su test, en su paquete, con el mismo nombre.

**Dominio** (`internal/domain/dashboard_test.go`), con las métricas armadas a mano:

| # | Tarea | Escenario BDD | Reglas | Estado |
|---|---|---|---|---|
| 1 | Un proyecto sin sprints arma un dashboard sin estado general, sin sprint en curso y sin gráficos, y sin error | Un proyecto sin sprints avisa que no hay datos | RN-7 | hecha |
| 2 | El estado general trae cada métrica del proyecto con su rótulo y su valor, tal como llegó, incluida la velocidad promedio | El estado general muestra las metricas tal como se calcularon | RN-2, RN-4 | hecha |
| 3 | Porcentajes y horas con un decimal, coma y unidad; enteros sin decimales; la desviación negativa, con su signo | Porcentajes, horas y enteros se muestran con su formato | RN-15 | hecha |
| 4 | El sprint en curso trae el sprint abierto y lo que lleva registrado, sin completados ni velocidad; sin sprint abierto, no trae nada | El sprint en curso muestra su meta y lo que lleva registrado; Un proyecto con todos sus sprints cerrados no tiene sprint en curso | RN-5, RN-6 | hecha |
| 5 | Los gráficos de story points, horas y defectos, en ese orden, cada barra con su rótulo, su valor y su valor escrito | Dashboard de un proyecto con un sprint en curso | RN-8, RN-10 | hecha |
| 6 | El gráfico de velocidad va cuarto, con una barra por sprint cerrado ordenadas por número; sin sprints cerrados, sin barras y con el aviso, y la velocidad del estado general escrita como «sin sprints cerrados» | El grafico de velocidad tiene una barra por sprint cerrado; Todavia no hay sprints cerrados para el grafico de velocidad | RN-4, RN-8, RN-9 | hecha |
| 7 | Armar el dashboard no cambia nada del proyecto | Abrir el dashboard no cambia el proyecto | RN-16 | hecha |
| 8 | El dashboard se pide al proyecto y usa las métricas que calcula la #16 | — (conexión con la #16) | RN-2 | hecha |

**Web** (`internal/web/dashboard_test.go`), con los datos del dashboard armados a mano:

| # | Tarea | Escenario BDD | Reglas | Estado |
|---|---|---|---|---|
| 9 | El largo de una barra: 100 % para el valor más grande, la parte proporcional con un decimal, sin barra para un cero y al menos 1 % para un valor mayor que cero | El largo de cada barra es proporcional al valor mas grande del grafico | RN-11, RN-12, RN-13 | hecha |
| 10 | La página del dashboard, con el encabezado del proyecto, su fecha de inicio como día/mes/año y «sin fecha de fin», y el acceso en la barra; un proyecto que no existe da 404 con el mensaje y el acceso a la lista | Encabezado de un proyecto sin fecha de finalizacion; Abrir el dashboard desde la barra de navegacion; Pedir el dashboard de un proyecto que no existe | RN-3, RN-15, RN-17, RN-18 | hecha |
| 11 | Sin sprints, la página muestra el encabezado y el aviso, sin gráficos | Un proyecto sin sprints avisa que no hay datos | RN-7 | hecha |
| 12 | El estado general y el sprint en curso muestran lo que entrega el dominio, con las fechas del sprint como día/mes/año; sin sprint abierto, el aviso | El estado general muestra las metricas tal como se calcularon; El sprint en curso muestra su meta y lo que lleva registrado; Un proyecto con todos sus sprints cerrados no tiene sprint en curso | RN-4, RN-5, RN-6, RN-15 | hecha |
| 13 | Cada gráfico dibuja una barra por valor con el largo de la tarea 9, escrito con punto en el atributo del SVG (`61.5%`, D3), y su rótulo y su valor escritos con coma; un cero sin barra y un gráfico todo en cero sin barras; el de velocidad sin sprints cerrados, con el aviso | Las barras de un grafico se miden contra su valor mas grande; El grafico de velocidad tiene una barra por sprint cerrado; Todavia no hay sprints cerrados para el grafico de velocidad | RN-8 a RN-14 | hecha |
| 14 | La página tiene las cuatro partes en orden, y los gráficos en el orden en que llegan | Dashboard de un proyecto con un sprint en curso | RN-1, RN-8 | hecha |
| 15 | Cada gráfico tiene su texto alternativo con cada rótulo y su valor, y su tabla de valores | Cada grafico tiene su texto alternativo y su tabla | RN-19, RN-20 | hecha |
| 16 | El dashboard entra en los tests de la #77: en castellano, adaptado al ancho, con la barra, sin pedir nada de afuera, y su tabla dentro del recuadro | Usar el dashboard en un celular y sin internet | RN-21 | hecha |
| 17 | De punta a punta: con el dominio real, la página muestra los números que calcula la #16 | — (conexión de las dos mitades) | RN-2 | hecha |

**Verificación manual**, antes de abrir el PR y anotada en la Fase 4, como en la #77: el
contraste de las barras con un verificador (RN-20), el modo celular del navegador a 360 de
ancho (RN-21), un lector de pantalla leyendo un gráfico (RN-19) y la página con la red
desconectada.

### Por qué este orden

- **Del 1 al 7, el dominio**: no depende de la web. El 1 es el caso que corta antes de todo;
  el 2 y el 3 van separados para que un rojo diga si falla qué número va o cómo se escribe.
- **La 8 conecta con la #16**, y es la única del dominio que la necesita en `dev`.
- **La 9 primero en la web**: el largo de una barra no depende de la página y se prueba solo.
- **De la 10 a la 15, la página**, de afuera hacia adentro: la página y su encabezado, el caso
  sin sprints, los números, los gráficos, el orden de todo y la accesibilidad.
- **La 16 y la 17 al final**: la 16 suma la página a lo que ya revisa la #77, y la 17 prueba
  que las dos mitades se conectan.

### Cobertura

Las veintiuna reglas quedan cubiertas: RN-1 (14), RN-2 (2, 8, 17), RN-3 (10), RN-4 (2, 6, 12),
RN-5 (4, 12), RN-6 (4, 12), RN-7 (1, 11), RN-8 (5, 6, 13, 14), RN-9 (6, 13), RN-10 (5, 13),
RN-11 (9, 13), RN-12 (9, 13), RN-13 (9, 13), RN-14 (13), RN-15 (3, 10, 12), RN-16 (7), RN-17 (10),
RN-18 (10), RN-19 (15), RN-20 (15 y la verificación manual), RN-21 (16 y la verificación
manual).

| Criterio de aceptación | Tareas |
|---|---|
| Cuatro partes en orden: encabezado, estado general, sprint en curso y gráficos | 14 |
| El encabezado con el nombre y las fechas, o «sin fecha de fin» | 10 |
| El estado general con los mismos números del cálculo de métricas | 2, 8, 12, 17 |
| La velocidad promedio, o «sin sprints cerrados» | 2, 6, 12 |
| Con un sprint abierto, sus datos y lo registrado, sin completados ni velocidad | 4, 12 |
| Sin sprint abierto, se informa y el resto se muestra igual | 4, 12 |
| Un proyecto sin sprints muestra el aviso, sin gráficos y sin error | 1, 11 |
| Cuatro gráficos: story points, horas, defectos y velocidad | 5, 6, 14 |
| El gráfico de velocidad, una barra por sprint cerrado, sin el abierto | 6, 13 |
| Cada barra con su rótulo y su valor | 5, 13 |
| La barra más grande ocupa todo el ancho y las demás son proporcionales | 9, 13 |
| Un cero sin barra, un valor chico al menos 1 %, todo en cero sin error | 9, 13 |
| Porcentajes y horas con un decimal y coma; desviación negativa con signo; fechas como día/mes/año | 3, 10, 12 |
| Texto alternativo con sus valores y tabla | 15 |
| Las barras no se distinguen solo por el color y tienen contraste | 15 y la verificación manual |
| El dashboard se abre desde la barra | 10 |
| Sin internet y en un celular | 16 y la verificación manual |
| Abrir el dashboard no cambia nada | 7 |
| Un proyecto inexistente informa que no se encontró | 10 |

Los dieciséis escenarios de la Fase 1 tienen su tarea en la columna "Escenario BDD".

**Un test pasa solo** y por eso va con otro que falla o se anota en Desvíos: el de que armar
el dashboard no cambia nada (tarea 7), porque el armado nunca escribe. Las tareas 8 y 17 no
tienen escenario: sus tests van encabezados con `// Tarea 8:` y `// Tarea 17:`.

### Dependencias

- **Del 1 al 7 arman las métricas a mano**, contra el contrato, aunque la #16 ya esté en
  `dev`: así cada test elige sus números. Los sprints abiertos y cerrados (#8, #10) se arman
  dentro del paquete.
- **La 3 necesita la regla de formato**: si la #18 no la escribió todavía, la escribe esta
  tarea, y la #18 la usa. Al 23/10 todavía no está en `dev`.
- **La 8 y la 17 usan las métricas de la #16, que ya está en `dev`** (PRs #81 y #140): no
  esperan nada.
- **De la 9 a la 17 usan la #77, que ya está en `dev`**, y la 10 a la 16 se pueden escribir
  antes de que el dominio tenga el dashboard: arman sus datos a mano (D10).

> **Puerta 3** — ¿las tareas cubren todos los criterios de aceptación?
> Aprobó: Maximiliano Eula — Fecha: 2026-10-22 — Comentarios: aprobada en el PR #136 (https://github.com/emisorato1/seguimiento-medicion-ics/pull/136#issuecomment-6350299170), junto con la Puerta 2, con tres correcciones antes del código. Se aplicaron el 23/10 en 9ce7080: el contrato ya no figura como pendiente, la #16 figura en `dev` y el ancho de las barras va con punto (D3, tarea 13).

## Fase 4 — Implementación

Ciclo RED → GREEN → REFACTOR. El test en un commit, con la salida del test fallando después
de `RED:` en el mensaje, y la implementación en otro.

| Tarea | Commit del test | Commit de la implementación |
|---|---|---|
| 1 | `6367a0e` | `1b11f92` |
| 2 | `5941e21` | `a3487f5` |
| 3 | `92c6a8b` | `b66feed` |
| 4 | `6330fb4` | `e84abc3` |
| 5 | `ef5ab19` | `fc0c221` |
| 6 | `1be2d30` | `2e28197` |
| 7 | `4f8e6d1` (sin RED) | — |
| 8 | `792ee9a` | `acc6ff5` |
| 9 | `a176c80` | `1833a8d` |
| 10 | `9bf4313` | `132571d` |
| 11 | `9c4e01e` | `4a9caae` |
| 12 | `71ccebe` | `6f8bd5a` |
| 13 | `210f45d` | `cc37e50` (y el estilo de toda la página, `0ee0fc3`) |
| 14 | `3721850` (sin RED) | — |
| 15 | `fae9c92` | `d7803df` |
| 16 | `ad3b799` (sin RED) | — |
| 17 | `13048ad` (sin RED) | — |

### Decisiones tomadas al implementar

Detalles que la Fase 2 no fijaba y que no cambian ningún escenario:

- **El empate se redondea hacia afuera del cero** (tarea 3): 12,25 h se escribe "12,3 h". La
  biblioteca estándar redondea el empate al par y daría "12,2", que no es como se redondea a
  mano. Pasa con los cuartos de hora del esfuerzo.
- **Un negativo que se redondea a cero se escribe sin signo** (tarea 3): una desviación de
  -0,04 % es "0,0 %" y no "-0,0 %", porque la RN-15 escribe el cero sin signo.
- **La regla de formato quedó en `internal/domain/format.go`**, aparte del dashboard, para
  que la #18 la use tal como está.
- **El armado recibe las métricas de cada sprint por su número** (tarea 4): las del sprint
  abierto para el sprint en curso, y las mismas le van a servir a la tarea 6 para la
  velocidad de cada sprint cerrado.
- **Los gráficos usan las métricas del proyecto entero** (tarea 5), las mismas del estado
  general: un número no se lee distinto en dos partes de la página (RN-2).
- **Cada gráfico lleva un título corto y cada barra un rótulo corto** (tarea 5): «Story
  points» con «Planificados» y «Completados», «Horas» con «Estimadas» y «Reales», y
  «Defectos» con «Detectados» y «Resueltos». El título dice qué se compara y el rótulo, cuál
  es cada barra.
- **El test de la tarea 5 mira solo los tres primeros gráficos**: el de velocidad, cuarto, lo
  agrega la tarea 6 con su propio test, sin reescribir este.
- **El gráfico de velocidad se titula «Velocidad»** (tarea 6), corto como los otros tres y
  con el mismo rótulo que la velocidad del estado general.
- **El aviso «sin sprints cerrados» sale de los sprints, no de la velocidad** (tarea 6): una
  velocidad de cero no dice si hay sprints cerrados, porque la #16 da cero en los dos casos.
  El test de la tarea suma como subcaso el caso límite de los sprints cerrados con velocidad
  promedio de cero, que se escribe «0».
- **`p.Dashboard()` descarta el error de las métricas de cada sprint** (tarea 8), con un
  comentario: `SprintMetrics` solo falla con un sprint que no existe, y los números salen de
  los sprints del propio proyecto. Es lo mismo que hace el cierre de la #10 con `p.Story`, y
  así no se toca el código de la #16.
- **El test de la tarea 8 usa las operaciones de verdad** (las historias, el sprint, la
  asignación y el cierre de la #10) y fija los números que da la #16 con esos datos, en lugar
  de compararlos contra una llamada a la #16 dentro del test.
- **El largo de la barra es un número, no el texto del SVG** (tarea 9): la web lo mide en
  porcentaje con un decimal (61,5), y escribirlo con punto en el atributo (`61.5%`, D3) es de
  la tarea 13, que dibuja la barra. El test suma dos casos que el escenario no muestra: el
  redondeo a un decimal (8 de 13 da 61,5) y un gráfico todo en cero, que no divide por cero.
- **«Dashboard» va en la barra después de Datos** (tarea 10), antes de Integrantes: es la
  vista general del proyecto, al lado de sus datos (RN-18). La pestaña se titula «Dashboard
  de» y el nombre del proyecto, como «Esfuerzo de»; el h1 es el nombre, con el inicio y el
  fin debajo, escritos igual que en la página de datos.
- **Los tests de la barra y del proyecto que no existe son propios** (tarea 10), en
  `internal/web/dashboard_test.go` y con el nombre de su escenario, en lugar de filas en las
  tablas de la #77 y la #78. Así toda la trazabilidad de la #17 apunta a un solo archivo.
- **El aviso sin sprints lleva el acceso a crear el primero** (tarea 11): «Todavía no hay
  datos para mostrar.» en el recuadro de vacío, con el botón «Crear el primer sprint», como
  la pantalla de Sprints. El dashboard no cambia nada: el botón solo lleva a otra pantalla.
- **El estado general y el sprint en curso se escriben como en la página del proyecto**
  (tarea 12): cada número en una lista de rótulo y valor, el sprint en curso con su enlace,
  su goal y sus fechas («Sprint 2: … (26/10/2026 al 30/10/2026)»), y sin sprint abierto el
  mismo texto, «No hay ningún sprint en curso».
- **Cada gráfico es un solo SVG** (tarea 13), con los rótulos, los valores y las barras
  adentro, y no un SVG chico por barra: así la tarea 15 le pone el `role="img"` y el texto
  alternativo a un solo dibujo, como pide la D4. Lo eligió Juliana. Los gráficos van en una
  parte «Gráficos», debajo del sprint en curso, cada uno en su tarjeta con el título arriba.
- **El color de cada barra sale de su posición** (tarea 13), sin cambiar el contrato del
  dominio: en los tres gráficos que comparan, la primera barra es lo planificado y la segunda
  lo real; en el de velocidad, todas del mismo color. Lo eligió Juliana.
- **La paleta es la que eligió Juliana el 27/10** (tarea 13, D6), y es el foco de toda la
  página del dashboard, no solo de los gráficos, como pidió después: el nombre del proyecto
  en Siena, los títulos subrayados en Rosa Polvoriento, las tarjetas en alabastro con un
  borde de color, y los enlaces y el botón en Siena. Va en `app.css`, solo para esta página,
  sin tocar los colores de la #77 ni el resto de la web: cada gráfico en una tarjeta Alabastro Cálido
  (`#F5E6D3`), lo planificado en Siena Tostado (`#A0522D`, 4,58 a 1), lo real en Verde
  Salvia oscurecido (`#5E7A47`, 3,94 a 1) y la velocidad en Vara de Oro oscurecida
  (`#8A6A1E`, 4,12 a 1). Los originales de la paleta, `#87A96B` y `#D4A437`, no llegaban al
  3 a 1 de la RN-20 sobre ese fondo. Las fuentes de la web no cambian.
  **Cambiado el 29/10 por el rediseño de toda la web en Verde Salvia** (rama
  `estilo-paleta-salvia`, con el ok del grupo): el dashboard deja su paleta propia y toma la
  de toda la web. Los gráficos van en tarjetas blancas como el resto, y las barras en Siena
  (`#A0522D`, 5,62 a 1 sobre blanco), Salvia vivo (`#4A7C59`, 4,86) y Oro oscurecido
  (`#8A6A1E`, 5,05). Los números de contraste de la verificación manual de abajo son los de la
  paleta anterior.
- **El texto alternativo empieza con una frase por gráfico** (tarea 15) que dice qué compara,
  y sigue con cada rótulo y su valor: «Story points planificados contra completados.
  Planificados: 40, Completados: 30.». La frase la da la web por la posición del gráfico,
  como el color, sin cambiar el contrato del dominio. Va en el `aria-label` del SVG, que es
  una imagen para el lector de pantalla. Lo eligió Juliana.
- **La tabla de «Ver los valores» tiene dos columnas, «Barra» y «Valor»** (tarea 15), con
  cada rótulo como encabezado de su fila, dentro del recuadro de tablas de la #77. El
  gráfico de velocidad sin sprints cerrados tiene solo el aviso, sin dibujo ni tabla: el
  aviso ya es texto. Lo eligió Juliana.

### Desvíos respecto de la spec

- **Tarea 4:** el escenario "Un proyecto con todos sus sprints cerrados no tiene sprint en
  curso" se llama igual que uno de la #5, y los dos tests están en el mismo paquete. El del
  dashboard se llama `TestUnProyectoConTodosSusSprintsCerradosNoTieneSprintEnCursoEnElDashboard`,
  con el escenario exacto en su comentario. Además pasaría sin código nuevo, porque hasta esta
  tarea el sprint en curso venía siempre vacío; el RED lo dio el otro test de la tarea.
- **Tarea 7, sin RED:** el test de "Abrir el dashboard no cambia el proyecto" pasó al
  escribirlo, como anticipaba la Fase 3, porque el armado nunca escribe. No tiene commit de
  implementación. Para saber que sirve se rompió el armado a propósito dos veces —que
  cerrara el sprint abierto y que le vaciara las historias— y falló las dos; después se
  volvió atrás (el detalle está en el mensaje del commit). El backlog el armado no lo puede
  tocar, porque solo recibe los sprints: esa mitad del test cuida la operación de la
  tarea 8.
- **Tareas 12 a 15, los tests de la web usan datos reales (D10):** la D10 decía que los
  tests de la web arman el dashboard a mano. Pero la página se lo pide al proyecto con
  `p.Dashboard()`, así que armarlo a mano exigía sumarle al servidor de la #77 una forma de
  recibirlo. Juliana eligió el 27/10 que los tests armen el proyecto con las operaciones de
  verdad (historias estimadas, sprints, asignación y cierre), sin tocar el servidor. Cada
  test fija igual los números que da la #16 con sus datos.
- **Tarea 13, el caso «más de lo planificado» va con el gráfico de horas:** el escenario lo
  pide con 20 story points planificados y 25 completados, pero con datos reales lo
  completado sale de lo comprometido y nunca lo supera. Juliana eligió probarlo con 20 h
  estimadas y 25 h reales, que da lo mismo (80 % y 100 %). Los otros cuatro casos van con
  story points como dice el escenario.
- **Tarea 14, sin RED:** el test de "Dashboard de un proyecto con un sprint en curso" pasó al
  escribirlo, porque las tareas 12 y 13 ya dejaron las cuatro partes y los gráficos en orden.
  No tiene commit de implementación. Para saber que sirve se rompió el orden dos veces —el
  sprint en curso antes del estado general, en la plantilla, y Horas antes que Story points,
  al armar los gráficos— y falló las dos; después se volvió atrás. Lo eligió Juliana, como
  en la tarea 7.
- **Tarea 16, un test propio y sin RED:** la tarea decía sumar el dashboard a los tests de la
  #77 en `server_test.go`. Juliana eligió un test propio del escenario en
  `internal/web/dashboard_test.go`, con un proyecto con sprints, por la misma razón que en la
  tarea 10: toda la trazabilidad de la #17 en un archivo. Revisa lo mismo que la #77 y,
  además, que cada gráfico mida todo el ancho. Pasó al escribirlo, porque el dashboard ya
  usa el layout y el recuadro de tablas de la #77; se rompió dos veces —la tabla sin su
  recuadro y los gráficos con 600 px fijos— y falló las dos. Que la página no se desplace de
  costado a 360 de ancho queda para la verificación manual.
- **Tarea 17, sin RED:** el test de punta a punta iba a ser el que conectara la web con el
  dominio real, porque la D10 armaba los datos a mano. Como desde la tarea 12 los tests ya
  usan el dominio real, Juliana eligió igual un test completo, `TestElDashboardMuestraLosNumerosQueCalculaLaHistoria16`:
  un proyecto con historias, un sprint cerrado y otro abierto, esfuerzo y defectos, y cada
  número de la página (estado general, sprint en curso y cada barra) fijado con lo que da la
  #16 para esos datos. Es el recorrido de la RN-2 para la defensa. Pasó al escribirlo; se
  rompió la conexión dos veces en `p.Dashboard()` —el estado general con las métricas del
  sprint abierto, y el sprint en curso con las del proyecto— y falló las dos.

### Verificación manual

El 28/10, con los datos de demo (`docs/demo/datos-demo.json` importados a una base nueva) y
el navegador del panel de Claude Code (Chromium), sin script:

- **Contraste (RN-20 y RN-19 de la #77)**, con la fórmula de la WCAG sobre los colores que
  pinta la página: las barras sobre la tarjeta alabastro dan 4,58 a 1 (Siena), 3,94 (Salvia
  oscurecido) y 4,12 (Oro oscurecido), todas sobre el 3 a 1. El texto da 12,90 sobre la
  tarjeta; los rótulos y el aviso en gris, 5,22; el título y los enlaces en Siena, 5,24
  sobre el fondo y 4,58 sobre la tarjeta; el botón, 5,62 (7,55 al pasar el mouse); la
  tabla, 6,39 y 15,80. Todo sobre el 4,5 a 1 del texto.
- **A 360 de ancho (RN-21)**: la página no se desplaza de costado (ancho del documento 360).
  Cada gráfico mide 292 px, cada tarjeta termina en el borde con su margen, y las cuatro
  tablas abiertas entran en su recuadro sin desplazarse.
- **Sin pedir nada de afuera (RN-21)**: el único recurso que pide la página es
  `/static/app.css`, que sale del programa.
- **Texto alternativo (RN-19)**: en el árbol de accesibilidad, cada gráfico es una imagen
  con su nombre completo, por ejemplo «Story points planificados contra completados.
  Planificados: 109, Completados: 58.».

**Lo que queda para hacer a mano antes del PR:** escuchar un gráfico con un lector de
pantalla de verdad (NVDA o Narrador) y abrir la página con la red desconectada. Lo de arriba
lo anticipa, pero no lo reemplaza.

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #17 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | `internal/domain/dashboard_test.go` y `internal/web/dashboard_test.go` |
| Código Go | `internal/domain/dashboard.go`, `internal/web/dashboard.go` y `internal/web/templates/dashboard.html` |
