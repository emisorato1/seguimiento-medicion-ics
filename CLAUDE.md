# Reglas de trabajo del proyecto

Proyecto integrador de **Ingeniería y Calidad de Software** — UTN FRSR, 2026.

Este archivo es el punto de entrada. Lo lee cualquier integrante del equipo y cualquier
Claude Code que se abra sobre el repositorio. Si estás empezando una sesión, leelo entero
antes de tocar nada: acá está **qué hay que construir, cómo se trabaja y cómo seguir**.

Las convenciones técnicas (comandos, estructura, estilo de código, git) están en
[AGENTS.md](AGENTS.md).

---

## Cómo sigo desde acá

Siempre, al empezar una sesión:

0. **Sincronizar antes que nada.** El repositorio lo tocan cuatro personas:

   ```bash
   git checkout dev && git pull
   ```

   Nunca empezar a trabajar sobre una copia vieja. Si ya estabas en una rama de trabajo,
   traé `dev` primero: `git pull origin dev`.

1. **Mirar si hay pull requests esperando tu revisión:**

   ```bash
   scripts/revisiones.sh
   ```

   Si hay alguno, **eso va antes que ponerte a programar**: un PR parado frena al que lo
   abrió y a todos los que dependen de esa historia. Claude prepara la revisión (muestra el
   diff, señala lo que no cierra con la spec), pero **la aprobación la da una persona**.

2. Leer [ESTADO.md](ESTADO.md) — dice qué historias hay, en qué estado y cuáles están
   estimadas. Si tenés `gh`, regeneralo primero con `scripts/estado.sh`.
3. Elegir la historia a trabajar: la que te toca según
   [docs/specs/README.md](docs/specs/README.md), respetando el orden de dependencias que ahí
   se indica. Si la historia que te toca depende de otra que todavía no está en `dev`,
   podés avanzar con su spec pero no con su código.
4. Mirar en qué estado del tablero está y hacer lo que corresponde:

| Si la historia está en... | Lo que sigue es... |
|---|---|
| Product Backlog | No se toca: entra al sprint en la planning. Si no está estimada, se estima con Planning Poker. |
| Sprint Backlog | Escribir la **Fase 1 (Requisitos)** de su spec en `docs/specs/sprint-N/NN-nombre.md` y pedir la Puerta 1. Mover a *Especificación SDD*. |
| Especificación SDD | Completar Diseño y Tareas con sus puertas. Cuando pasa la Puerta 3, mover a *Lista para implementar*. |
| Lista para implementar | Las tres puertas están aprobadas. Si lo que necesita ya está en `dev`, empezar el ciclo TDD y mover a *En progreso*. Si todavía depende de otra historia, no se toca. |
| En progreso | Implementar con TDD, tarea por tarea. Cuando está lista, abrir el pull request y mover a *En review*. |
| En review | Esperar la revisión de otro integrante. Si hay correcciones, volver a *En progreso*. |
| Hecho | Nada. Elegir otra historia. |

**Regla de oro: no se escribe código de una historia que no tiene su spec aprobada.**

Si no hay ninguna historia en Sprint Backlog, el sprint no está planificado: eso es una
reunión del equipo, no algo que se resuelva en una sesión de código.

---

## Qué hay que construir

Una aplicación para **estimar, seguir y medir proyectos de software**. Puede ser web, de
escritorio o de consola, pero **el núcleo y todas las reglas de negocio van en Go**.

Los nueve requerimientos mínimos de la consigna, cada uno con sus historias en el tablero:

| # | Requerimiento | Historias |
|---|---|---|
| 1 | Gestión de proyectos | #3, #4, #5 |
| 2 | Product Backlog | #6, #7 |
| 3 | Gestión de Sprints | #8, #9, #10, #11 |
| 4 | Estimación y Planning Poker | #12, #13 |
| 5 | Registro de esfuerzo | #14 |
| 6 | Gestión de defectos | #15 |
| 7 | Métricas | #16 |
| 8 | Dashboard | #17 |
| 9 | Reportes (PDF en el Sprint 4) | #18, #19 |

La consigna completa está en el campus. **Ninguna funcionalidad nueva se agrega sin que
exista su historia en el tablero.**

## Cómo se evalúa

| Peso | Qué |
|---|---|
| 25% | Producto funcional |
| 25% | Aplicación de SDD, BDD y TDD |
| 20% | Calidad del software |
| 20% | Gestión del proyecto |
| 10% | Equipo y presentación |

La mitad de la nota es **proceso**, no producto. Por eso el tablero, las specs, las actas y
el historial de commits se cuidan igual que el código.

## Roles

| Rol | Quién | Qué hace |
|---|---|---|
| Product Architect | los profesores | definen qué se construye y revisan el producto |
| Agile Enabler | Emiliano Sorato | facilita los eventos y las **puertas de aprobación**, cuida que el proceso se cumpla |
| Product Builders | Agustín Salinas, Juliana Bustos, Maximiliano Eula | construyen el producto |

El Agile Enabler **no** es el jefe: no asigna tareas ni decide por el equipo. El equipo se
autogestiona. También programa como el resto.

---

## El tablero

El **GitHub Project es la fuente de verdad** del estado del proyecto:
https://github.com/users/emisorato1/projects/3

`ESTADO.md` es una copia generada del tablero, para poder leer el estado sin `gh`. Nunca se
edita a mano: se regenera con `scripts/estado.sh` y se commitea.

### Estados

```
Product Backlog → Sprint Backlog → Especificacion SDD → Lista para implementar
  → En progreso → En review → Hecho
```

**Para agregar o renombrar un estado, hacelo desde la web del Project, no por la API.**
Reescribir las opciones del campo por la API les cambia el identificador y deja sin estado a
todas las tarjetas que lo tenían. Pasó una vez: 16 tarjetas quedaron en blanco y hubo que
restaurarlas a mano.

### Campos

| Campo | Para qué |
|---|---|
| Story Points | la estimación acordada en Planning Poker (Fibonacci) |
| Prioridad | Alta / Media / Baja, la define el equipo con los profes |
| Sprint | Sprint 0 a 4 |
| Horas estimadas / Horas reales | para medir la desviación |

### Mantenerlo al día

**La historia se mueve en el momento en que cambia, no al final del día.** Un tablero que no
refleja la realidad es peor que no tener tablero.

```bash
scripts/revisiones.sh                           # PRs esperando tu revision
scripts/coherencia.sh                           # specs, indice y puertas coherentes
scripts/tablero.sh ver                          # estado de todas las historias
scripts/tablero.sh estado 3 "En progreso"       # mover una historia
scripts/tablero.sh puntos 3 5                   # cargar story points
scripts/tablero.sh sprint 3 "Sprint 1"          # asignarla a un sprint
scripts/tablero.sh prioridad 3 Alta
scripts/estado.sh                               # regenerar ESTADO.md y commitearlo
```

---

## Sprints

Plan de la cátedra:

| Sprint | Objetivo |
|---|---|
| 0 | preparación: repositorio, tablero, backlog inicial |
| 1 | MVP |
| 2 | interfaz usable |
| 3 | funcionalidad y calidad (visualización) |
| 4 | cierre, informe en PDF y review final |

### Eventos

| Evento | Cuándo | Qué queda documentado |
|---|---|---|
| Planning | al empezar el sprint | acta en `docs/actas/sprint-N-planning.md`: Sprint Goal, estimaciones y Sprint Backlog |
| Daily | acordado por el equipo, aunque sea por chat | no se documenta |
| Review | al cerrar el sprint | acta con qué se completó y qué no |
| Retrospectiva | después de la review | acta con qué salió bien, qué no y los acuerdos |

Las actas de retrospectiva son **entregable final**, así que ninguna se saltea.

### Reglas de sprint

- Un solo sprint abierto a la vez.
- Lo que no se termina vuelve al Product Backlog; no se arrastra en silencio.
- El alcance del sprint no se amplía a mitad de camino.
- Al cerrar, se registran los story points completados: esa es la **velocidad**, y es lo que
  se usa para planificar el sprint siguiente.

---

## Historias y story points

### Formato de una historia

Toda historia del tablero tiene:

```
**Como** [rol]
**quiero** [qué]
**para** [para qué]

## Criterios de aceptación
- [ ] ...
```

Los criterios de aceptación son la base de los escenarios BDD y de los tests. Una historia
sin criterios de aceptación no entra a un sprint.

### Estimación

- Se estima en **Story Points**, que miden tamaño y complejidad, no horas.
- La escala es Fibonacci: **1, 2, 3, 5, 8, 13, 21**. No existe el 4 ni el 7.
- Se estima con **Planning Poker**, con el equipo junto y los votos ocultos hasta que votan
  todos. El procedimiento está en [docs/guia-planning-poker.md](docs/guia-planning-poker.md).
- Una historia de 21 casi seguro hay que partirla.
- **Nadie estima solo, y nadie cambia una estimación fuera de una planning.**

### Lista para entrar al sprint

Una historia puede entrar a un sprint solo si: tiene criterios de aceptación, está estimada,
y el equipo entiende qué hay que hacer.

---

## El flujo de una historia, de punta a punta

```
Historia (issue) → Spec SDD → Criterios de aceptación → Escenarios BDD → Tests → Código Go
```

### SDD — cuatro fases con puertas

```
Requisitos → [P1] → Diseño → [P2] → Tareas → [P3] → Implementación
```

Cada fase produce un documento que se revisa **antes** de avanzar. La plantilla está en
[docs/specs/PLANTILLA.md](docs/specs/PLANTILLA.md) y la spec de cada historia va en
`docs/specs/sprint-N/NN-nombre-de-la-funcionalidad.md`, versionada junto al código.

**El índice de specs, con el estado de cada una, quién la tiene y en qué orden hay que
hacerlas, está en [docs/specs/README.md](docs/specs/README.md).** Es lo primero que hay que
mirar para saber si una historia se puede empezar o está esperando otra.

No todo lleva el flujo completo: las tareas de configuración y los bugs chicos se resuelven
directo. SDD se reserva para funcionalidad de complejidad media o alta.

Reglas de las puertas:

- **La puerta la aprueba alguien distinto del que escribió la fase**, y queda registrado
  quién y cuándo en la propia spec.
- Puede ser un comentario en el pull request: no hace falta reunión.
- Se revisa en las puertas, no durante la implementación.
- El Agile Enabler es responsable de que las puertas efectivamente ocurran.

Calibrar el esfuerzo al tamaño del problema: una historia chica puede tener un diseño de
tres líneas. Escribir specs enormes para cosas simples es un anti-patrón, igual que escribir
specs que nadie revisa de verdad.

### Qué va y qué no va en una spec

Las dos reglas que más se incumplen, y que el apunte de la cátedra marca como
anti-patrones:

- **En Requisitos no va nada técnico.** Ni base de datos, ni estructuras, ni nombres de
  funciones: solo comportamiento. La prueba es si alguien que conoce el dominio pero no
  programa puede entender y validar el requisito. Si solo lo entiende un programador,
  tiene decisiones técnicas disfrazadas.
- **En Diseño van decisiones, no código.** Nada de structs, firmas ni pseudocódigo: si el
  diseño se vuelve código en prosa, el agente lo traduce mecánicamente y se pierde el valor
  de trabajar con IA. La heurística: *si dos personas con experiencia tomarían la misma
  decisión sin discutirla, no va en el diseño; si podrían elegir distinto y las
  consecuencias cambian, sí va.*

Si la spec termina más larga que el código que genera, está sobreespecificada.

Las reglas de negocio se escriben con la notación **EARS**, que es la que recomienda el
apunte: *Cuando [evento], el sistema debe…* / *Mientras [condición], el sistema debe…* /
*Si [situación de error], el sistema debe…*. La plantilla tiene los cinco patrones.

### BDD

Los criterios de aceptación se escriben como escenarios **Given-When-Then** dentro de la
Fase 1 de la spec, cubriendo casos normales, alternativos, límite y de error. Cada escenario
tiene que terminar en al menos un test.

### TDD

Ciclo **RED → GREEN → REFACTOR**, siempre el test primero.

**El test va en un commit y la implementación en otro.** El historial de commits es la
evidencia del TDD que evalúa la cátedra: no se aplasta la historia, no se hace `push --force`
sobre `main`, y no se sube un test junto con la implementación ya terminada.

### Trazabilidad

En la defensa eligen una funcionalidad al azar y piden el recorrido completo: historia →
spec → criterios → escenario BDD → test → código. Cada spec cierra con esa tabla, y
**cualquiera del equipo tiene que poder recorrerla**, no solo quien la escribió.

---

## Git y ramas

Tres niveles, como en un proyecto real:

```
main                 rama "de producción": siempre estable, siempre compila, siempre verde
 └── dev             integración: acá se juntan las historias terminadas del sprint
      └── historia-3-crear-proyecto     una rama por historia
      └── historia-4-registrar-integrantes
```

- **Nadie commitea directo a `main` ni a `dev`.** Todo entra por pull request.
- Una rama por historia, sacada siempre de `dev` actualizado:

  ```bash
  git checkout dev && git pull
  git checkout -b historia-3-crear-proyecto
  ```

- Nombre de la rama: `historia-<numero>-<descripcion-corta>`. Para un defecto:
  `defecto-<numero>-<descripcion-corta>`.
- **Una tarjeta del tablero = una historia = una rama = un pull request.** No mezclar dos
  historias en una rama: hace imposible revisar y rompe la trazabilidad.
- Al mergear el PR, **GitHub borra la rama solo**. No hace falta limpiarla a mano. Para
  borrar tu copia local de las ramas que ya no existen en el remoto:

  ```bash
  git fetch --prune
  ```
- Antes de abrir el PR, traer `dev` a tu rama y resolver los conflictos vos:

  ```bash
  git checkout dev && git pull
  git checkout historia-3-crear-proyecto && git merge dev
  ```

- No se hace `push --force` sobre ramas compartidas ni se aplasta el historial: los commits
  del ciclo TDD son la evidencia que se evalúa.

### Protección de ramas

`main` y `dev` están protegidas en GitHub, así que las reglas de arriba no dependen de la
buena memoria de nadie:

- No se puede pushear directo: hace falta un pull request.
- El PR necesita **1 aprobación** y el **CI en verde** para poder mergearse.
- Si se suben commits nuevos, las aprobaciones anteriores se descartan.
- No se permite `push --force` ni borrar la rama.
- Solo se puede mergear con **merge commit**: squash y rebase están desactivados.
- Los comentarios del PR tienen que estar resueltos antes de mergear.

El dueño del repositorio puede saltear la protección en una emergencia. Si alguna vez se
usa, se avisa al equipo y se deja dicho por qué.

### Historias y tareas: qué es una rama y qué no

| Nivel | Qué es | Dónde vive |
|---|---|---|
| Historia | una tarjeta del tablero, algo que le sirve al usuario | **una rama y un PR** |
| Tarea | un paso de la Fase 3 de la spec, un ciclo de TDD | **commits** dentro de la rama de la historia |

Las tareas **no** llevan rama propia. Una tarea es "el test de que no se puede crear un
proyecto sin nombre" más su implementación: dos commits. Partir eso en ramas daría PRs de
tres líneas, multiplicaría las revisiones y rompería la trazabilidad, porque el PR ya no
cerraría una historia completa.

Si una historia se siente demasiado grande para una sola rama (13 o 21 story points), el
problema no es la rama: **hay que partir la historia en dos historias**, cada una con su
tarjeta, su spec y su valor propio. Eso se decide en la planning o en el refinamiento, no
mientras se programa.

### Cómo se mergea

Solo está habilitado el **merge commit**. El squash y el rebase están desactivados a
propósito: aplastan los commits y borrarían la evidencia del ciclo TDD, que es lo que
la cátedra evalúa.

### Dos personas, la misma historia

El tablero es el candado: si una historia está en *En progreso*, ya la está haciendo alguien.
No la tomes. Si necesitás trabajar sobre algo que otro está tocando, se habla primero.

## Pull requests

El PR se abre cuando la historia cumple la Definición de Terminado. La plantilla
([.github/pull_request_template.md](.github/pull_request_template.md)) se completa entera:
es parte de la documentación que se evalúa, no un trámite.

- El PR enlaza la **historia** y la **spec**, y usa `Closes #N` para que la issue se cierre sola.
- **Revisa alguien que no escribió el código.** Es la puerta de calidad del flujo.
- Quien revisa mira: que el código haga lo que dice la spec, que los tests cubran los
  escenarios BDD, y que se entienda. Si no se entiende, se pide que se aclare.
- Plazo: si a las 48 h nadie revisó, el Agile Enabler busca revisor. Un PR parado frena a todos.
- No se mergea con el CI en rojo. Nunca.
- **Mergea el autor**, una vez que tiene la aprobación y el CI en verde. Es quien resuelve
  los conflictos, mueve la tarjeta y regenera `ESTADO.md`.
- Al mergear a `dev`: la historia pasa a *Hecho* en el tablero y se regenera `ESTADO.md`.

### Quién aprueba

| Pull request de… | Lo aprueba… |
|---|---|
| Un Product Builder | Emiliano Sorato (Agile Enabler) |
| Emiliano Sorato | cualquiera de los otros tres |
| `dev` → `main` (cierre de sprint) | el equipo en la review |

Además, **todos pueden comentar en cualquier PR aunque no sean el aprobador**, y conviene
que lo hagan: es la forma más barata de que los cuatro conozcan todo el proyecto, y en la
defensa pueden preguntarle a cualquiera por cualquier parte del código.

### De `dev` a `main`

Al cerrar el sprint se abre un PR de `dev` a `main` con el incremento completo. Ese PR lo
revisa el equipo junto en la review, y al mergear se etiqueta:

```bash
git tag -a v0.1.0-sprint1 -m "Incremento del Sprint 1"
git push origin v0.1.0-sprint1
```

Cada tag es la evidencia de un incremento entregado.

## CI y CD

| Workflow | Cuándo corre | Qué hace |
|---|---|---|
| [CI](.github/workflows/ci.yml) | en cada PR y push a `dev` o `main` | coherencia de las specs, `gofmt`, `go vet` y `go test ./... -cover` |
| [CD](.github/workflows/cd.yml) | al entrar algo a `main` | corre los tests, compila y publica los binarios |

Si el CI se pone en rojo en `dev` o `main`, **arreglarlo es la prioridad número uno** del
equipo, antes que seguir con cualquier historia. Una rama rota bloquea a los cuatro.

## Medir nuestro propio proyecto

El informe final pide métricas, y las más fáciles de defender son las de este mismo proyecto.
Para tenerlas hay que ir cargando los datos mientras se trabaja:

- **Horas**: al empezar una historia se carga *Horas estimadas* en el tablero; al terminarla,
  *Horas reales*. Con eso sale la desviación.
- **Defectos**: los errores que encontramos entre nosotros se cargan como issue con la
  etiqueta `defecto`, indicando la historia relacionada y en qué sprint se detectó y se
  resolvió. Sin esto no hay datos de calidad para el informe.
- **Velocidad**: sale sola de los story points completados al cerrar cada sprint, siempre que
  el tablero esté al día.

---

## Documentación

| Dónde | Qué |
|---|---|
| `README.md` | qué es el proyecto y quiénes lo hacen |
| `CLAUDE.md` | este archivo: las reglas de trabajo |
| `AGENTS.md` | convenciones técnicas: comandos, estructura, estilo, git, límites |
| `ESTADO.md` | estado del tablero (generado, no editar) |
| `docs/specs/sprint-N/` | una spec por funcionalidad, agrupadas por sprint |
| `docs/actas/` | plannings, reviews y retrospectivas |
| `docs/guia-planning-poker.md` | cómo estimamos |

Simple y al día vale más que extenso. Si un documento dejó de ser cierto, se corrige en el
mismo pull request que lo volvió falso.

---

## Uso de IA

La cátedra permite usar IA para análisis, especificaciones, código, tests, refactor, review
y documentación. La condición es que **todo lo generado quede entendido y validado por el
equipo**: respondemos por el código igual que si lo hubiéramos escrito a mano.

En la práctica: nadie mergea algo que no sabe explicar. Si en la defensa te preguntan por
una línea, la tenés que poder justificar.

### Reglas para Claude Code

**Siempre:**

- Hacer `git pull` sobre `dev` al empezar la sesión, antes de leer nada más.
- Correr `scripts/revisiones.sh` al empezar y avisar si hay PRs esperando revisión, antes de
  proponer cualquier otra cosa.
- Leer `ESTADO.md` y este archivo al empezar una sesión.
- Trabajar sobre una historia del tablero, y mover la historia cuando cambia de estado.
- Trabajar siempre en una rama de historia sacada de `dev`.
- Respetar el orden: spec aprobada antes que código, test antes que implementación.
- Regenerar `ESTADO.md` y commitearlo cuando se toca el tablero.

**Preguntar antes de:**

- Crear historias nuevas, cambiar prioridades o mover algo a *Hecho*.
- Cambiar una decisión que ya está escrita en una spec aprobada.
- Agregar dependencias o cambiar la estructura del proyecto.

**Nunca:**

- Commitear directo a `main` o a `dev`.
- Mergear un pull request con el CI en rojo, o mergear el PR propio sin que lo revise otro.
- Escribir código de una historia sin spec aprobada.
- Completar actas o retrospectivas con contenido inventado: lo que dijo el equipo lo escribe
  el equipo.
- Aprobar una puerta. Las puertas las aprueban personas.
- Cargar story points que no salieron de una planning.
