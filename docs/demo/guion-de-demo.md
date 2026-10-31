# Guion de la demo

Un recorrido de 8 a 10 minutos por el sistema con nuestro propio proyecto, cargado con los
[datos de demo](README.md). Cada paso dice qué abrir, qué hacer y qué contar. Las capturas
de cada pantalla están en [capturas/](capturas).

El recorrido muestra todas las historias, también el **dashboard (#17)**, el **reporte (#18)** y
su **PDF (#19)**, que se abren desde la web (#78).

## Antes de presentar

1. Traer `dev` y abrir el sistema con los datos de demo en una base nueva:

   ```bash
   git checkout dev && git pull
   go run ./cmd/web -datos demo.db -importar docs/demo/datos-demo.json
   ```

   La primera vez dice «Se importaron los datos de docs/demo/datos-demo.json». Para volver a
   empezar de cero, se borra `demo.db` y se repite el comando.

   *Si la #177 (la base de datos) todavía no está en `dev`*, se abre con una copia del archivo:
   `cp docs/demo/datos-demo.json /tmp/demo.json && go run ./cmd/web -datos /tmp/demo.json`. En
   ese caso el paso 7 se cuenta con el archivo y no con la base.
2. Abrir `http://127.0.0.1:8080` en el navegador, con la ventana grande.
3. Tener a mano una terminal visible para el paso 7.

## 1. El proyecto (1 minuto)

- **Abrir** la lista de proyectos y entrar a *Seguimiento y Medición ICS*.
  ([01](capturas/01-proyectos.png), [02](capturas/02-proyecto-y-resumen.png))
- **Contar:** "Cargamos nuestro propio proyecto con lo que registramos en el tablero y en
  las actas. El resumen dice que tenemos 21 historias, 13 completadas al cierre del Sprint
  2 y 8 pendientes."
- **Mostrar** *Integrantes*: los cuatro, con Emiliano como Agile Enabler.
  ([03](capturas/03-integrantes.png))

## 2. El backlog y una historia (1 minuto)

- **Abrir** *Backlog*: cada historia con su número de issue, prioridad, estado y story
  points. ([04](capturas/04-backlog.png))
- **Entrar** a la historia 13, *Votar en secreto en una ronda de Planning Poker*.
  ([05](capturas/05-historia-con-criterios.png))
- **Contar:** "Cada historia tiene sus criterios de aceptación, los mismos de la issue. De
  ahí salen los escenarios BDD y los tests."

## 3. El Planning Poker de la #77 (2 minutos)

- **Abrir** *Planning Poker* y, en la #77, *Todas las rondas de la historia*.
  ([06](capturas/06-planning-poker.png), [10](capturas/10-rondas-de-la-historia.png))
- **Ronda 1**, revelada: ([08](capturas/08-ronda-1-revelada-con-diferencias.png))
  - **Contar:** "Esta es la estimación que hicimos de verdad el 26/10. En la primera ronda
    Agustín y Juliana votaron 5, y Maxi y Emiliano 13. El sistema marca la diferencia: el más
    bajo y el más alto, y quiénes los dieron."
- **Ronda 2**, acordada: ([09](capturas/09-ronda-2-acordada.png))
  - **Contar:** "Después de hablarlo votamos de nuevo, los cuatro 13, y quedó acordada en 13.
    La historia tomó esos story points."
- **Votos ocultos:** mostrar la captura de la ronda 2 antes de que votara Juliana.
  ([07](capturas/07-ronda-con-votos-ocultos.png))
  - **Contar:** "Mientras la ronda está abierta nadie ve ningún voto, ni siquiera el propio:
    solo quién votó y quién falta."
  - En vivo se muestra con una **ronda de práctica**, al final (paso 8), sobre una copia de
    la base.

## 4. Un sprint y su cierre (2 minutos)

- **Abrir** *Sprints* y entrar al Sprint 1. ([11](capturas/11-sprints.png),
  [12](capturas/12-sprint-1-cerrado.png))
  - **Contar:** "El Sprint 1 cerró el 16/10. De las cinco historias se completaron dos, la #3
    y la #4: 6 story points, que fue nuestra velocidad. Las otras volvieron al backlog y se
    terminaron en el Sprint 2."
- **Crear el Sprint 3** como figura en su acta: *Nuevo sprint*, con el goal «Que el sistema
  mida el proyecto y lo muestre: métricas, dashboard y reporte.», del 26/10 al 28/10.
  Asignarle la #10, la #16, la #17, la #11, la #18 y la #78.
  ([13](capturas/13-sprint-3-abierto.png))
- **Cerrar el sprint:** *Cerrar el sprint* propone la fecha de hoy y una casilla por historia.
  ([14](capturas/14-cerrar-sprint.png))
  - **Cambiar la fecha** por la del cierre real del Sprint 3 (28/10, o la que diga el acta de
    la review). El formulario acepta una fecha posterior al fin del sprint: si se deja la de
    hoy, el sprint queda cerrado el día de la presentación.
  - Se marcan **solo las historias que de verdad se terminaron** (mirar el tablero antes de
    presentar) y se confirma. En la captura están las cinco terminadas al 30/10: todas menos
    la #78.
  - **Contar:** "Antes de cerrar, el sistema dice cuáles quedan completadas y cuáles vuelven
    al Product Backlog. Un sprint cerrado no se puede volver a abrir."

## 5. El esfuerzo y los defectos (1,5 minutos)

- **Abrir** *Esfuerzo*. ([15](capturas/15-esfuerzo.png))
  - **Contar:** "Son las horas reales del tablero de los dos primeros sprints, unas 69, con el
    total por historia y por integrante."
  - La página no muestra el total general, solo los de cada historia y cada integrante
    (16 + 28 + 8 + 17,25 = 69,25). El dashboard lo redondea a 69,3.
- **Abrir** *Defectos*: están los tres. ([16](capturas/16-defectos.png))
  - **Contar:** "Son los defectos que nos encontramos entre nosotros, con la historia y el
    sprint en que se detectaron y se resolvieron."
- **Entrar** al defecto 60, *Menor*, y volver. ([17](capturas/17-defecto.png))
- **Filtrar** por la severidad *Mayor*: quedan el 59 y el 108.

## 6. El dashboard, el reporte y el PDF

- **Abrir** *Dashboard* (1 minuto). ([18](capturas/18-dashboard.png))
  - **Contar:** "El dashboard junta todo lo que medimos. Con el Sprint 3 recién cerrado, el
    proyecto lleva 86 de 109 story points y 18 de 21 historias completadas. La velocidad es
    el promedio de los tres sprints cerrados: 6, 52 y 28, que da unos 29."
  - **Mostrar** los gráficos: los story points planificados contra los completados, las horas
    estimadas contra las reales, los defectos y la velocidad de cada sprint. Con *Ver los
    valores* se ven en una tabla.
  - *Sprint en curso* dice que no hay ninguno, porque el Sprint 3 se acaba de cerrar.
  - Si preguntan por el Sprint 3: en la demo da 28 y no 33 como en su review, porque la #73 se
    adelantó y no está asignada en estos datos. Y las horas reales son solo las de los dos
    primeros sprints (69,3 h): las del Sprint 3 no se cargan en la demo. La desviación de cada
    sprint está en su acta de review.


- **Abrir** *Reporte* (1 minuto), elegir el Sprint 2 en *Reporte de* y apretar *Ver el
  reporte*. ([19](capturas/19-reporte.png))
  - **Contar:** "El reporte dice lo mismo que el dashboard, pero en texto y por secciones: el
    resumen, las historias, la estimación y el esfuerzo, y los defectos. Sale de la misma cuenta
    de métricas que el dashboard, así que no puede mostrar otra cosa. El Sprint 2 da 52 story
    points de 52 y los dos defectos resueltos."
  - Si preguntan por las horas: el reporte mide cada historia entera, y por eso da 53,3 h. El
    acta de review cuenta 31, las trabajadas en el sprint, porque tres historias venían del
    Sprint 1.
- **Descargar en PDF** y abrir el archivo. ([20](capturas/20-pdf.png))
  - **Contar:** "El PDF es el mismo reporte, con el nombre del proyecto y del sprint:
    reporte-proyecto-1-sprint-2.pdf. Lo escribe nuestro código en Go, sin librerías."

## 7. La base de datos: cerrar y volver a abrir (1 minuto)

- **En la terminal**, apretar `Ctrl+C`: el sistema se cierra.
- **Volver a abrirlo sin importar nada**: `go run ./cmd/web -datos demo.db`, y recargar la
  página.
- **Mostrar** que está todo, también lo que se hizo en vivo: el Sprint 3 cerrado.
- **Contar:** "Los datos se guardan en una base SQLite, con una tabla por cada cosa del
  sistema. Cada cambio se guarda entero o no se guarda: nunca queda a medias. Y con
  `-exportar` se pueden llevar todos los datos a un archivo."

## 8. Una ronda de práctica, sobre una copia (1 minuto)

Para mostrar el Planning Poker en vivo sin que los datos reales del proyecto queden con votos
de práctica, la ronda se hace **sobre una copia de la base**.

- **En la terminal**, `Ctrl+C`, y después:

  ```bash
  cp demo.db practica.db
  go run ./cmd/web -datos practica.db
  ```

- **Abrir** *Planning Poker*, elegir la #19 y *Abrir una ronda*. Votar con dos integrantes.
- **Mostrar** que la página dice quién votó y quién falta, pero ningún voto.
- **Contar:** "Esta ronda es de práctica, sobre una copia: la base del proyecto no cambió."
- Al terminar, `Ctrl+C`. `practica.db` se puede borrar.

## Cierre (30 segundos)

- **Contar:** "Todo lo que mostramos tiene su historia en el tablero, su spec, sus escenarios
  y sus tests. Si quieren, elegimos una funcionalidad y la recorremos de punta a punta."

## Si algo sale mal en vivo

| Qué pasa | Qué hacer |
|---|---|
| «el puerto ya lo está usando otro programa» | abrir con `-puerto 8081` y usar esa dirección |
| «la base ya tiene datos y no se importa sobre datos existentes» | ya estaba importada: abrir sin `-importar`, o borrar `demo.db` y volver a importar |
| Algo de lo cargado en vivo salió mal | cerrar, borrar `demo.db` e importar de nuevo: tarda un segundo |
| La ronda de práctica quedó en la base del proyecto | se hizo sin la copia: cerrar, borrar `demo.db`, importar de nuevo y repetir el Sprint 3 |
