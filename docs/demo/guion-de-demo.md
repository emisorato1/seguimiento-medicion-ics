# Guion de la demo

Un recorrido de 8 a 10 minutos por el sistema con nuestro propio proyecto, cargado con los
[datos de demo](README.md). Cada paso dice qué abrir, qué hacer y qué contar. Las capturas
de cada pantalla están en [capturas/](capturas).

Los lugares del **dashboard (#17)**, del **reporte (#18)** y del **PDF (#19)** están marcados:
se completan cuando esas historias entren a `dev`.

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
    presentar) y se confirma.
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

> **[Dashboard — #17]** Cuando entre la #17: abrir el dashboard y mostrar los gráficos de la
> velocidad, las horas estimadas contra las reales y los defectos. Tiempo: 1 minuto. Captura:
> `capturas/18-dashboard.png`.

## 6. El reporte y el PDF

> **[Reporte — #18]** Cuando entre la #18: abrir el reporte del proyecto y el del Sprint 2, y
> mostrar sus secciones. Captura: `capturas/19-reporte.png`.
>
> **[PDF — #19]** Cuando entre la #19: descargar el reporte en PDF y abrirlo. Captura:
> `capturas/20-pdf.png`.

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
