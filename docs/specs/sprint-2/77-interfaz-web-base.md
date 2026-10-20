# Spec: Interfaz web: proyectos, integrantes y backlog

- **Historia:** #77
- **Sprint:** 2
- **Responsable:** Emiliano Sorato
- **Estado:** Puerta 3 aprobada — **lista para implementar** (las aprobó el responsable como excepción; falta que un compañero las ratifique en la issue #77)
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de la consigna y
  de los criterios de aceptación de la issue. El responsable lo revisa, lo corrige si hace
  falta y lo presenta a las puertas.
- **Depende de:** #3, #4, #6 y #12, que ya tienen su código en `dev`
- **La necesitan:** #78 (usa su estructura, su navegación y su forma de mostrar errores) y
  #17 (el dashboard se muestra en la web)

## Fase 1 — Requisitos

### Objetivo

Que el equipo pueda cargar y consultar sus proyectos desde una página web clara, que se
entiende sin explicación y se usa también en un celular. Es la base de la interfaz web: la
navegación, la forma de mostrar errores y confirmaciones, y las pantallas de lo que el
sistema ya hace hoy.

### Qué queda afuera

- Las pantallas de sprints, esfuerzo, defectos, Planning Poker y reportes (#78) y el
  dashboard (#17).
- El resumen del estado del proyecto (#5) y el backlog ordenado por prioridad (#7): punto 6
  de "Para decidir en la puerta".
- **Reglas de negocio nuevas.** La web pide los datos, se los pasa al núcleo —las reglas del
  sistema, las mismas que usa la consola— y muestra lo que responde.
- Borrar proyectos o historias, que el núcleo no permite (#3, #6), y usuarios o permisos.
- Que la consola vea lo que se carga en la web: punto 2 de "Para decidir en la puerta".

### Entradas

| Formulario | Datos obligatorios | Datos opcionales |
|---|---|---|
| Proyecto (crear y modificar) | nombre, fecha de inicio | descripción, fecha de fin |
| Integrante | nombre y apellido, rol | — |
| Historia (cargar y modificar) | título, descripción, prioridad, al menos un criterio de aceptación | — |
| Estimación | story points | — |

Las fechas se eligen en un calendario; el rol, la prioridad y los story points, de una lista
(RN-11); los criterios se escriben uno por renglón. "Obligatorio" es lo que exige el núcleo:
la página no impide enviar un formulario incompleto (RN-15).

### Salidas

Las páginas con los proyectos, sus integrantes y su backlog, y los datos cambiados y
guardados según cada operación aceptada. Las tablas muestran: proyectos con número, nombre,
inicio y fin; integrantes con nombre y rol; backlog con número, título, prioridad, estado y
story points.

### Reglas de negocio

1. **[RN-1]** El sistema debe ofrecer desde el navegador: crear, modificar y consultar
   proyectos; agregar, listar y quitar integrantes con su rol; cargar, modificar y consultar
   historias del backlog con su prioridad; y estimar una historia o quitarle la estimación.
2. **[RN-2]** El sistema debe abrirse con un solo comando, que muestra la dirección donde se
   usa desde el navegador.
3. **[RN-3]** El sistema debe funcionar completo sin conexión a internet.
4. **[RN-4]** El sistema debe mostrar en todas las páginas una barra de navegación fija con
   el acceso a la lista de proyectos y, mientras se está dentro de un proyecto, su nombre y
   el acceso a sus datos, sus integrantes y su backlog.
5. **[RN-5]** El sistema debe mostrar sus textos en castellano y las fechas como dd/mm/aaaa.
6. **[RN-6]** El sistema debe mostrar los listados como tablas con encabezado, y un dato
   opcional sin cargar como «sin descripción», «sin fecha de fin» o «sin estimar», nunca
   vacío ni con un cero.
7. **[RN-7]** Mientras una lista esté vacía, el sistema debe decir que no hay datos y
   ofrecer ahí mismo el acceso para cargar el primero.
8. **[RN-8]** Cuando una operación termina bien, el sistema debe mostrar un mensaje que la
   confirma y nombra lo que cambió, por ejemplo «Se creó el proyecto Sistema de métricas».
9. **[RN-9]** Cuando se recarga la página que sigue a una operación, el sistema no debe
   repetir la operación ni volver a mostrar su mensaje.
10. **[RN-10]** Antes de quitar un integrante o una estimación, el sistema debe pedir
    confirmación; si se cancela, no cambia nada.
11. **[RN-11]** El sistema debe ofrecer el rol, la prioridad y los story points como una
    lista para elegir, con exactamente los valores que acepta el núcleo.
12. **[RN-12]** Si el núcleo rechaza una operación, el sistema debe volver a mostrar el
    formulario con lo que se había escrito y, junto al campo que lo causó, el mismo mensaje
    que da el núcleo.
13. **[RN-13]** Si el rechazo no corresponde a ningún campo, el sistema debe mostrarlo
    arriba del formulario.
14. **[RN-14]** Si un dato no tiene la forma esperada —una fecha que no es una fecha—, el
    sistema debe indicarlo junto al campo y no ejecutar la operación.
15. **[RN-15]** La web no debe validar reglas de negocio por su cuenta ni impedir que se
    envíe un formulario incompleto: lo rechaza el núcleo (RN-12).
16. **[RN-16]** Si se pide un proyecto o una historia que no existe, el sistema debe mostrar
    el mensaje del núcleo y ofrecer volver a la lista de proyectos.
17. **[RN-17]** El sistema debe mostrar cada texto tal como se escribió, incluidos signos
    como `<` o `&`, sin que cambie cómo se ve la página.
18. **[RN-18]** El sistema debe poder usarse en una notebook y en un celular de 360 puntos
    de ancho sin que la página se desplace de costado; una tabla que no entra se recorre
    dentro de su recuadro.
19. **[RN-19]** El sistema debe poder usarse solo con el teclado, con cada campo nombrado
    por una etiqueta visible, el foco siempre a la vista y el texto con un contraste de al
    menos 4,5 a 1 con su fondo.
20. **[RN-20]** Mientras se use desde varias pestañas a la vez, el sistema debe aplicar las
    operaciones de a una, sin perder ni mezclar ninguna.
21. **[RN-21]** Cuando una operación termina bien, el sistema debe guardar los datos, de
    modo que al cerrarlo y volver a abrirlo esté todo lo cargado, con los mismos números.
22. **[RN-22]** Cuando se vuelve a abrir, el sistema debe seguir la numeración donde había
    quedado: un proyecto o una historia nuevos nunca reciben un número ya usado.
23. **[RN-23]** Mientras no haya datos guardados, como la primera vez, el sistema debe
    arrancar sin proyectos, y eso no es un error.
24. **[RN-24]** Si los datos guardados no se pueden leer o contienen algo que el núcleo no
    acepta, el sistema debe avisarlo al arrancar y no arrancar, dejándolos como estaban.
25. **[RN-25]** Si un cambio no se puede guardar, el sistema debe avisar que no quedó guardado.

### Casos límite

- La primera vez, sin datos guardados: la lista de proyectos vacía, con su mensaje.
- Al modificar un proyecto, borrar la descripción o la fecha de fin las quita (#3, RN-9).
- Criterios con renglones en blanco: el núcleo los ignora; todos en blanco, rechaza (#6).
- Quitar el único integrante: la lista vuelve a mostrarse vacía, con su mensaje. «Quitar
  estimación» se ofrece solo en una historia estimada.
- Un enlace a un proyecto con letras en lugar del número: no se encontró el proyecto.
- Una historia Terminada (cuando existan los sprints) se ve, pero el núcleo rechaza
  modificarla o estimarla y el motivo va arriba del formulario.

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| El núcleo rechaza los datos de un formulario | vuelve a mostrarlo con lo escrito y el mensaje junto al campo |
| El rechazo no corresponde a un campo | lo muestra arriba del formulario |
| Una fecha que no es una fecha | lo indica junto al campo y no ejecuta la operación |
| Proyecto o historia que no existen | muestra el mensaje del núcleo y ofrece volver a la lista |
| Los datos guardados no se pueden leer o no se aceptan | avisa al arrancar, no arranca y no los toca |
| Un cambio no se puede guardar | avisa que el cambio no quedó guardado |

### Criterios de aceptación

Los ocho primeros son los de la issue.

- [ ] El sistema se abre en el navegador con un solo comando y funciona sin internet.
- [ ] Desde la web se crean, modifican y consultan proyectos.
- [ ] Desde la web se agregan y quitan integrantes, con su rol.
- [ ] Desde la web se administra el Product Backlog: alta, modificación, prioridad y
      estimación en story points.
- [ ] Los errores de validación se muestran junto al campo que falló, con el mensaje del
      sistema, sin perder lo que se había escrito.
- [ ] Las listas vacías muestran que no hay datos y qué se puede hacer.
- [ ] La página se puede usar en una pantalla de notebook y en un celular.
- [ ] Los datos cargados no se pierden al cerrar y volver a abrir el sistema.
- [ ] Cada operación se confirma con un mensaje, y recargar la página no la repite.
- [ ] Antes de quitar algo se pide confirmación, y cancelar no cambia nada.
- [ ] La barra de navegación muestra siempre en qué proyecto se está.
- [ ] Cada campo tiene su etiqueta, se ve el foco y todo se usa con el teclado.
- [ ] Si los datos guardados no se pueden leer, el sistema avisa y no los pisa.
- [ ] La web no contiene ninguna validación de negocio propia.

### Escenarios BDD

```gherkin
Escenario: Crear un proyecto desde la web
  Dado que no hay ningun proyecto cargado
  Cuando envio el formulario de proyecto con nombre "Sistema de metricas" e inicio 01/11/2026
  Entonces veo el mensaje "Se creo el proyecto Sistema de metricas"
  Y el proyecto figura en la lista con el numero 1
  Y si recargo la pagina, sigue habiendo un solo proyecto y el mensaje ya no aparece

Esquema del escenario: Cada operacion se confirma con un mensaje
  Dado <situacion>
  Cuando <accion>
  Entonces veo el mensaje <mensaje>
  Y <resultado>

  Ejemplos:
    | situacion                           | accion                                                  | mensaje                                       | resultado                                        |
    | un proyecto sin integrantes         | agrego a "Ana Perez" con el rol "Product Builder"       | "Se agrego a Ana Perez"                       | figura en los integrantes como "Product Builder" |
    | un proyecto sin historias           | cargo "Registrar esfuerzo" con prioridad "Alta"         | "Se cargo la historia 1"                      | figura como "Pendiente" y "sin estimar"          |
    | la historia 1 sin estimar           | elijo 5 story points                                    | "Se estimo la historia 1 en 5 story points"   | figura en el backlog con 5 story points          |
    | el proyecto "Sistema de metricas"   | cambio su nombre por "Metricas ICS"                     | "Se modifico el proyecto Metricas ICS"        | figura con el nombre "Metricas ICS"              |
    | el proyecto 1 con descripcion       | borro su descripcion                                    | "Se modifico el proyecto Sistema de metricas" | figura "sin descripcion"                         |
    | la historia 1 con prioridad "Media" | cambio su prioridad por "Alta"                          | "Se modifico la historia 1"                   | figura con prioridad "Alta"                      |

Escenario: Consultar un proyecto
  Dado el proyecto "Sistema de metricas" con inicio 01/11/2026, sin descripcion ni fecha de fin
  Cuando entro a su pagina
  Entonces veo su numero, su nombre, "sin descripcion", el inicio 01/11/2026 y "sin fecha de fin"
  Y la barra de navegacion muestra "Sistema de metricas" y el acceso a sus integrantes y su backlog

Esquema del escenario: Las opciones para elegir son las del sistema
  Cuando abro el formulario de <formulario>
  Entonces las opciones de <lista> son exactamente <opciones>

  Ejemplos:
    | formulario | lista        | opciones                       |
    | integrante | rol          | Agile Enabler, Product Builder |
    | historia   | prioridad    | Alta, Media, Baja              |
    | estimacion | story points | 1, 2, 3, 5, 8, 13, 21          |

Esquema del escenario: Confirmar o cancelar antes de quitar
  Dado el integrante "Ana Perez" y la historia 1 estimada en 5
  Cuando pido <accion>
  Entonces el sistema me pide confirmarlo
  Y si <respuesta>, <resultado>

  Ejemplos:
    | accion                       | respuesta | resultado                                   |
    | quitar a "Ana Perez"         | confirmo  | "Ana Perez" ya no figura en los integrantes |
    | quitar a "Ana Perez"         | cancelo   | "Ana Perez" sigue en los integrantes        |
    | quitar la estimacion de la 1 | confirmo  | la historia 1 figura "sin estimar"          |
    | quitar la estimacion de la 1 | cancelo   | la historia 1 sigue estimada en 5           |

Esquema del escenario: Mostrar el rechazo sin perder lo escrito
  Dado <situacion>
  Cuando envio el formulario de <formulario> con <dato>
  Entonces el formulario vuelve a mostrarse con todo lo que escribi
  Y <donde> veo el mensaje <mensaje>
  Y no cambio ningun dato

  Ejemplos:
    | situacion                     | formulario | dato                               | donde                       | mensaje                                                   |
    | que no hay proyectos          | proyecto   | el nombre vacio                    | junto al nombre             | "el nombre del proyecto es obligatorio"                   |
    | que no hay proyectos          | proyecto   | inicio 01/11/2026 y fin 01/10/2026 | junto a la fecha de fin     | "las fechas son incoherentes"                             |
    | que no hay proyectos          | proyecto   | una fecha de inicio que no es fecha | junto a la fecha de inicio | "la fecha no es valida"                                   |
    | un proyecto con "Ana Perez"   | integrante | el nombre "ana perez"              | junto al nombre             | "el integrante ya esta cargado en el proyecto"            |
    | un proyecto con Agile Enabler | integrante | el rol "Agile Enabler"             | junto al rol                | "el proyecto ya tiene un Agile Enabler"                   |
    | un proyecto sin historias     | historia   | el titulo vacio                    | junto al titulo             | "el titulo de la historia es obligatorio"                 |
    | un proyecto sin historias     | historia   | los criterios en blanco            | junto a los criterios       | "la historia necesita al menos un criterio de aceptacion" |
    | la historia 1 "Terminada"     | historia   | otro titulo                        | arriba del formulario       | "una historia terminada no se modifica"                   |

Esquema del escenario: Las listas vacias dicen que hacer
  Dado <situacion>
  Cuando abro <pagina>
  Entonces veo el mensaje <mensaje>
  Y el acceso para <accion>

  Ejemplos:
    | situacion                      | pagina                | mensaje                                      | accion                     |
    | el sistema sin datos guardados | la lista de proyectos | "Todavia no hay proyectos"                   | crear el primero           |
    | un proyecto sin integrantes    | sus integrantes       | "Este proyecto todavia no tiene integrantes" | agregar el primero         |
    | un proyecto sin historias      | su backlog            | "El backlog esta vacio"                      | cargar la primera historia |

Esquema del escenario: Pedir algo que no existe
  Dado que solo existe el proyecto 1, sin historias
  Cuando abro la pagina de <pedido>
  Entonces veo el mensaje <mensaje>
  Y el acceso para volver a la lista de proyectos

  Ejemplos:
    | pedido                        | mensaje                      |
    | el proyecto 99                | "no se encontro el proyecto" |
    | el proyecto "abc"             | "no se encontro el proyecto" |
    | la historia 99 del proyecto 1 | "no se encontro la historia" |

Escenario: Un texto con signos se muestra tal cual
  Cuando creo el proyecto "<b>Metricas</b> & Cia"
  Entonces en la lista veo el nombre "<b>Metricas</b> & Cia" escrito tal cual, sin negrita

Escenario: Abrir el sistema con un solo comando y sin internet
  Dado una computadora sin conexion a internet
  Cuando ejecuto el comando que abre el sistema
  Entonces veo la direccion para abrirlo en el navegador
  Y en esa direccion la lista de proyectos se ve completa, con sus estilos

Escenario: Los datos siguen al cerrar y volver a abrir
  Dado que cargue el proyecto 1 con el integrante "Ana Perez" y la historia 1 estimada en 3
  Cuando cierro el sistema y lo vuelvo a abrir
  Entonces veo el proyecto 1, a "Ana Perez" y la historia 1 estimada en 3
  Y el siguiente proyecto que creo recibe el numero 2

Escenario: Los datos guardados estan danados
  Dado que los datos guardados no se pueden leer
  Cuando ejecuto el comando que abre el sistema
  Entonces veo un aviso de que no se pudieron leer y el sistema no arranca
  Y los datos guardados quedan como estaban

Escenario: Un cambio que no se pudo guardar se avisa
  Dado que el lugar donde se guardan los datos no permite escribir
  Cuando creo un proyecto
  Entonces veo un aviso de que el cambio no quedo guardado

Escenario: Dos pestanas a la vez
  Dado el sistema abierto en dos pestanas
  Cuando en las dos creo un proyecto al mismo tiempo
  Entonces quedan dos proyectos, con numeros distintos

Escenario: Usar la pagina en un celular y con el teclado
  Dado una pantalla de 360 puntos de ancho
  Cuando abro el backlog de un proyecto con historias y recorro su formulario solo con el teclado
  Entonces la pagina no se desplaza de costado y la tabla se recorre dentro de su recuadro
  Y llego a cada campo en orden, con su etiqueta y viendo donde esta el foco
```

**Para decidir en la puerta:**

1. **Los datos se guardan solos después de cada operación** (RN-21), sin botón. El *cómo*
   es una **decisión de arquitectura nueva** que reemplaza el "datos en memoria" de
   AGENTS.md y obliga a cada historia que agrega datos a sumarlos a lo que se guarda (D13 a
   D19): la aprueba el equipo, no solo quien revisa.
2. **La consola no ve lo que se carga en la web.** La #24 está aprobada con "al cerrar el
   programa se pierde todo" y no se cambia desde acá. Puede usar la misma forma de guardar,
   pero las dos abiertas a la vez sobre los mismos datos se pisarían: es un cambio a la #24.
3. **"Con un solo comando"**: arranca el sistema y muestra la dirección. Que abra el
   navegador solo depende del sistema operativo y no se puede probar.
4. **"En un celular"**: la página se adapta a la pantalla de un celular (RN-18) y se
   verifica con el modo celular del navegador. Usarla desde un celular real exige aceptar
   conexiones de otras computadoras, que sin usuarios ni contraseñas queda apagado (D20).
5. **Elegir de una lista** el rol, la prioridad y los story points (RN-11): el rechazo de
   "4 story points" ya no se provoca desde la página, pero el núcleo lo sigue rechazando y un
   test lo comprueba enviándolo igual.
6. **La #5 y la #7 no entran**: no están en `dev` y ni la #77 ni la #78 las nombran.
   Propuesta: el backlog se muestra en el orden del núcleo, o en el de la #7 si llega antes
   de la tarea 19; el resumen de la #5 se suma a la página del proyecto en la #78.
7. **Un error por vez, sin el aviso propio del navegador** (RN-15): con dos campos mal se
   marca el primero que informa el núcleo y, al corregirlo, aparece el otro. Se pierde un
   aviso instantáneo, pero el mensaje es siempre el del núcleo, igual que en la consola.

**Resolución del responsable (20/10).** Emiliano Sorato, como responsable de la historia y
Agile Enabler, acepta los siete puntos tal como están planteados arriba, y también la
estructura nueva que `AGENTS.md` pide consultar antes de crear: `cmd/web`, `internal/web` e
`internal/store`. Las tres puertas las aprobó después el propio responsable, como excepción, porque el resto
del equipo no estaba disponible (ver las líneas de cada puerta).

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: Emiliano Sorato — Fecha: 2026-10-20 — Comentarios: aprobada en un comentario de la issue #77 (https://github.com/emisorato1/seguimiento-medicion-ics/issues/77#issuecomment-6299914339). La aprueba el responsable de la historia porque el resto del equipo no estaba disponible: es una excepción a la regla de que aprueba otra persona, y queda para que un compañero la ratifique con un comentario en la issue #77, desde su propia cuenta. Ningún PR de la historia entra a `dev` antes de esa ratificación.

## Fase 2 — Diseño

Las decisiones de [AGENTS.md](../../../AGENTS.md) valen igual y no se repiten. Esta historia
suma una segunda interfaz, dos paquetes y datos que sobreviven al cierre, así que tiene más
decisiones que una de dominio; las que cambian AGENTS.md están marcadas **(AGENTS.md)**.

### Enfoque

La web es una capa delgada, igual que la consola (#24): **lee el formulario, llama al núcleo
y muestra lo que responde.** No valida reglas ni guarda datos propios. Toda operación que
cambia algo sigue el mismo recorrido:

1. **Controlar la forma** —una fecha que no es fecha, un número de la dirección que no es
   número—; si falla, se vuelve a mostrar el formulario sin llamar al núcleo (RN-14).
2. **Tomar el candado** (D12) y llamar a la operación del dominio.
3. **Si rechaza**, se vuelve a mostrar el formulario con estado 422, lo escrito y el mensaje
   junto a su campo (D6).
4. **Si acepta**, se guarda el registro (D13), se deja el mensaje de una sola vez (D8) y se
   redirige con 303 a la página del resultado (D5); si el guardado falla, se avisa (D19).

Las páginas que solo muestran llenan un **layout común** —barra, mensaje y contenido—, y lo
que no existe responde 404 con el mensaje del dominio. Las direcciones van en castellano y
cuelgan del proyecto (`/proyectos/{id}/historias/{n}/editar`); cada formulario se pide por
GET y se envía por POST a la misma dirección. **Para la #78 queda la convención:** una
pantalla nueva es un archivo de handlers, sus plantillas y su test, y del layout solo toca
la barra. **Lo que se guarda** tiene tres responsables: el dominio convierte el registro a
JSON y lo reconstruye validando (D14, D15), `internal/store` lee y reescribe el archivo
(D17, D18), y la web solo pide guardar después de cada operación aceptada.

### Archivos afectados

| Archivo | Se crea o se modifica | Para qué |
|---|---|---|
| `cmd/web/main.go` y `main_test.go` | se crean | el comando: lee las opciones, carga los datos guardados, arma la web y muestra la dirección |
| `internal/web/server.go` | se crea | las rutas, las plantillas embebidas, el candado, el guardado y el mensaje de una sola vez |
| `internal/web/forms.go` | se crea | leer formularios, controlar la forma de fechas y números, y la tabla que asocia cada error del dominio con su campo |
| `internal/web/projects.go`, `members.go`, `backlog.go` | se crean | las páginas de proyectos, de integrantes y del backlog con su estimación |
| `internal/web/*_test.go` | se crean | `server_test.go`, `projects_test.go`, `members_test.go`, `backlog_test.go` y `forms_test.go`, con los esquemas que cruzan pantallas |
| `internal/web/templates/` y `static/app.css` | se crean | el layout, una plantilla por página y los estilos propios |
| `internal/store/file.go` y `file_test.go` | se crean | leer el archivo de datos al arrancar y reescribirlo entero sin dejarlo a medias |
| `internal/domain/registry_json.go` y su test | se crean | convertir el registro completo a JSON y reconstruirlo validando |
| `internal/domain/story_estimation.go` y su test | se modifican | leer la escala, como copia, para ofrecerla como opciones |
| `internal/domain/project_registry.go` | se modifica, **lo hace Emiliano Sorato** | la consulta que lista todos los proyectos del registro, que hoy no existe y que la página de inicio necesita. No es una tarea de esta historia (ver Dependencias) |
| `.github/workflows/cd.yml` | se modifica | compilar cada comando con su nombre: con dos `main`, `go build -o archivo ./cmd/...` falla y el CD se rompe |
| `.gitignore` y `README.md` | se modifican | no versionar el archivo de datos; cómo se abre la web |
| `AGENTS.md` | se modifica | la persistencia reemplaza a "datos en memoria"; la estructura suma `internal/web`, `internal/store` y `cmd/web`; la regla de sumar cada dato nuevo a la conversión; `go run ./cmd/web` en los comandos |
| `docs/diseno-dominio.md` | se modifica | la consulta de todos los proyectos, la lectura de la escala y la regla de la conversión |

### Decisiones

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| **D1** Solo biblioteca estándar: `net/http` con los patrones por método del `ServeMux`, `html/template` y `embed` | un framework web (Gin, Echo) o un motor de plantillas externo | AGENTS.md no permite dependencias sin acuerdo; desde Go 1.22 el `ServeMux` distingue el método y lee las variables de la dirección, que era lo que justificaba un router externo; `html/template` escapa los textos solo (RN-17) |
| **D2** Plantillas y CSS embebidos en el binario | leerlos del disco al arrancar | el binario que publica el CD funciona solo, desde cualquier carpeta; un archivo que falta no puede romper una página en la defensa |
| **D3** CSS propio, sin CDN ni framework de JavaScript, y **ningún JavaScript en esta historia** | Bootstrap o Tailwind por CDN; React o htmx | la defensa puede ser sin internet y un CDN deja la página sin estilos; un framework agrega otra tecnología que hay que defender y que `httptest` no prueba. Si la #78 necesita JavaScript, es mínimo y solo mejora |
| **D4** La web vive en `internal/web` y arranca con su propio comando, **`cmd/web`** | un modo de `cmd/app` (por ejemplo `go run ./cmd/app -web`) | `cmd/app/main.go` lo modifica la #24 para arrancar la consola: compartirlo hace que las dos historias choquen en el mismo archivo y que una interfaz no arranque sin la otra. Con dos comandos cada una se compila, se prueba y se publica sola, y la #78 suma pantallas sin tocar ningún arranque. AGENTS.md pide preguntar antes de crear un paquete: se pregunta en esta puerta **(AGENTS.md)** |
| **D5** Formularios por POST con redirección 303 después del éxito | responder la página desde el mismo POST | recargar repetiría la operación y crearía un proyecto duplicado (RN-9) |
| **D6** Al rechazar, se vuelve a mostrar el formulario con estado 422 y lo escrito; una tabla de la web asocia cada error del dominio con su campo comparando con `errors.Is`, y un error que no está en la tabla va arriba | redirigir también al fallar; mensajes propios de la web | redirigir pierde lo escrito; mensajes propios duplican los del dominio. Un error nuevo que nadie sumó a la tabla igual se muestra, arriba, y no se pierde (RN-13) |
| **D7** Ninguna validación en la web, ni siquiera marcar los campos como requeridos para el navegador; solo se controla la forma | dejar que el navegador frene los campos vacíos | el navegador mostraría su propio mensaje, en su idioma, y el del dominio no aparecería nunca: la regla quedaría duplicada. Controlar la forma es lo mismo que la RN-5 de la #24 |
| **D8** El mensaje de confirmación viaja en una cookie que se borra al mostrarlo | un parámetro en la dirección (`?ok=...`) | con el parámetro el mensaje vuelve a aparecer al recargar o al guardar el enlace (RN-9) |
| **D9** Confirmar antes de quitar con una página intermedia que tiene el botón que confirma. Solo hay dos cosas que se quitan: un integrante y una estimación | un diálogo de JavaScript | funciona en cualquier navegador y se prueba con `httptest`; el diálogo no. Borrar proyectos o historias no existe en el dominio ni en ninguna historia, y no se pide |
| **D10** Rol, prioridad y story points salen del dominio: las constantes de rol y prioridad con su nombre, y la escala con una lectura nueva. La web manda el texto elegido y el dominio lo convierte y lo valida con lo que ya tiene | escribir las opciones en la plantilla, o campos de texto libre | escritas a mano se desactualizan si cambia el dominio; con texto libre el error es lo normal. Un valor que no estaba en la lista igual lo rechaza el dominio |
| **D11** Las fechas se eligen con el calendario del navegador y se muestran dd/mm/aaaa | un campo de texto dd/mm/aaaa, como la consola | el calendario evita los errores de tipeo, funciona sin JavaScript y en el celular abre el selector propio del teléfono |
| **D12** Un único candado de lectura y escritura en la web, alrededor de cada operación **y de su guardado** | hacer seguro el dominio para varios pedidos a la vez | `net/http` atiende cada pedido en paralelo y los datos del dominio no lo soportan (RN-20). La concurrencia es un problema de la interfaz: el dominio sigue simple |
| **D13** Los datos se guardan en un archivo JSON con `encoding/json`, reescrito entero después de cada operación aceptada **(AGENTS.md)** | una base de datos (SQLite); guardar al cerrar; un registro de cambios | SQLite es una dependencia externa; guardar al cerrar pierde todo si el programa se corta; los datos son pocos y reescribir entero es lo más simple de probar |
| **D14** La conversión vive en `internal/domain`, con tipos intermedios no exportados; leer y escribir el archivo es de `internal/store` **(AGENTS.md)** | la conversión en la web o en `internal/store` | los datos de las entidades están ocultos fuera del paquete: afuera no se pueden leer ni reconstruir sin exportarlos, y exportarlos permitiría saltear las validaciones. El dominio sigue sin tocar archivos: convierte a texto y desde texto |
| **D15** Al reconstruir, cada entidad pasa por las mismas validaciones que al crearla, se conservan los números y los contadores, y un dato inválido rechaza la carga entera | confiar en el archivo | el archivo es texto y se puede editar o romper a mano; un proyecto sin nombre o 4 story points romperían reglas que el resto da por ciertas. Guardar los contadores es lo que hace que la numeración no retroceda (RN-22) |
| **D16** En el archivo, las fechas van como día (aaaa-mm-dd) y el rol, la prioridad y el estado con su nombre. No lleva número de versión: un dato que falta se carga vacío. **Un dato que el sistema no conoce rechaza la carga** | guardar la fecha con hora y zona, los números de las constantes y una versión del formato; **ignorar los datos que el sistema no conoce** | el día evita el defecto #59; con nombres el archivo se lee y no se rompe si se reordenan las constantes. Sin versión, un archivo de antes de los sprints carga con cero sprints, que es correcto. **Ignorar lo desconocido haría que una versión vieja del programa cargue un archivo con datos de una historia posterior (por ejemplo los sprints) sin ellos, y los borre en silencio al volver a guardar.** Lo de los datos desconocidos lo aceptó el responsable el 20/10, después de las puertas, a partir de la revisión del PR #92: queda para la ratificación |
| **D17** Se escribe en un archivo temporal de la misma carpeta y se renombra | escribir encima del archivo | un corte a mitad de la escritura dejaría un archivo roto y, por D15 y D18, el sistema no volvería a arrancar |
| **D18** Al arrancar: sin archivo, se empieza vacío; con un archivo que no se puede leer, el comando termina con el motivo y no lo toca | empezar vacío en los dos casos | el primer guardado pisaría los datos (RN-24) |
| **D19** Si el guardado falla, el cambio queda en memoria y se avisa con estado 500; el siguiente guardado exitoso lo incluye | deshacer el cambio | el dominio no tiene cómo deshacer, y como se guarda el registro entero nada se pierde si el siguiente guardado funciona |
| **D20** Escucha en `127.0.0.1:8080`, con las opciones `-puerto`, `-datos` (por defecto `datos-ics.json` en la carpeta donde se corre) y `-red` para aceptar conexiones de la red local | escuchar siempre en la red | sin usuarios ni contraseñas, cualquiera en la misma red podría cambiar los datos. `-red` existe para mostrarlo en un celular en la defensa |
| **D21** Los handlers se prueban con `net/http/httptest` contra el dominio real, con un registro nuevo y un guardado falso en memoria por test: estado, redirección, HTML y mensajes. **Un test por escenario**, test y código en commits separados | probar en un navegador real (Selenium, chromedp) | `httptest` es de la biblioteca estándar y corre en el CI sin navegador; el guardado falso permite simular que falla. Lo que así no se ve (contraste, ancho, foco) se verifica a mano (Fase 3) |
| **D22** Accesibilidad y diseño adaptable desde la base: idioma de la página en castellano, cada campo con su etiqueta asociada, el error enlazado al campo para los lectores de pantalla, foco visible, contraste de 4,5 a 1, una columna que se ensancha en pantallas grandes y tablas dentro de un recuadro que se desplaza | dejarlo para el final | si no se decide en la base, la #78 copia pantallas que no lo cumplen |

### Dependencias y restricciones

- **Depende de la #3, la #4, la #6 y la #12**, ya en `dev`, y de **una consulta nueva del
  dominio que liste todos los proyectos del registro**: hoy el registro solo crea, obtiene y
  modifica, y la página de inicio es esa lista. La agrega Emiliano Sorato en `project_registry.go`,
  con su test y su lugar en el contrato; esta spec no le pone nombre.
- **No se toca `cmd/app`**, que la #24 modifica para arrancar la consola (D4).
- **Dos cambios chicos más en el dominio, que sí son de esta historia**: la lectura de la
  escala (archivo de la #12) y la conversión, que es nueva. Se avisan al grupo antes de
  escribirlos, como pide la regla 3 del contrato.
- **Lo que se guarda crece con el dominio.** La historia que agrega un dato a una entidad
  (#8, #9, #10, #13, #14, #15, #73) lo suma también a la conversión, con su caso en el test
  de ida y vuelta: si no, se pierde al cerrar sin que ningún test lo note **(AGENTS.md y
  contrato)**.
- **Lo usa la #78**: el layout, la barra, los errores por campo, el mensaje de una sola vez,
  la página de confirmación, el candado y el guardado.

> **Puerta 2** — ¿el diseño es viable y coherente con el resto del proyecto?
> Aprobó: Emiliano Sorato — Fecha: 2026-10-20 — Comentarios: aprobada en un comentario de la issue #77 (https://github.com/emisorato1/seguimiento-medicion-ics/issues/77#issuecomment-6299914339). La aprueba el responsable de la historia porque el resto del equipo no estaba disponible: es una excepción a la regla de que aprueba otra persona, y queda para que un compañero la ratifique con un comentario en la issue #77, desde su propia cuenta. Ningún PR de la historia entra a `dev` antes de esa ratificación.

## Fase 3 — Tareas

Cada tarea es un ciclo de TDD completo: el test que falla en un commit y la implementación
que lo hace pasar en otro, en este orden y dejando todo en verde. Los tests de la web arman
un registro nuevo y el guardado falso (D21) y, como el HTML no se compara con `errors.Is`,
buscan en la página el texto del error del dominio tomado del propio error, nunca escrito a
mano. Cada esquema es **un solo test con tabla**: las tareas que lo comparten le suman sus
filas, que son lo que falla primero. Un test que pase sin código nuevo va a Desvíos.

| # | Tarea | Escenario BDD | Reglas | Estado |
|---|---|---|---|---|
| 1 | La escala de story points se puede leer, y cambiar lo recibido no la cambia | — (lectura nueva del dominio) | RN-11 | pendiente |
| 2 | Ida y vuelta: un registro con proyectos, integrantes e historias (estimada, sin estimar y Terminada) se convierte a JSON y se reconstruye igual, con sus números y contadores | Los datos siguen al cerrar y volver a abrir | RN-21, RN-22 | pendiente |
| 3 | Reconstruir rechaza un texto que no es JSON y los datos que el dominio no acepta: sin nombre, 4 story points, rol inexistente, dos Agile Enabler | Los datos guardados estan danados | RN-24 | pendiente |
| 4 | El archivo: si no existe, registro vacío; lo escrito se vuelve a leer igual; no queda ningún temporal | Los datos siguen al cerrar y volver a abrir / Las listas vacias dicen que hacer | RN-21, RN-23 | pendiente |
| 5 | Un archivo dañado da error y queda igual byte a byte; una carpeta donde no se puede escribir da error al guardar | Los datos guardados estan danados / Un cambio que no se pudo guardar se avisa | RN-24, RN-25 | pendiente |
| 6 | El layout: la raíz lleva a la lista de proyectos; toda página está en castellano, se adapta al ancho y tiene la barra; el CSS sale del binario y ninguna página pide nada de afuera | Abrir el sistema con un solo comando y sin internet | RN-3, RN-4 | pendiente |
| 7 | La lista de proyectos vacía dice que no hay y ofrece crear el primero | Las listas vacias dicen que hacer | RN-7, RN-23 | pendiente |
| 8 | Crear un proyecto redirige con 303 a su página y queda en la lista con su número; recargar no crea otro | Crear un proyecto desde la web | RN-1, RN-9 | pendiente |
| 9 | El mensaje de confirmación se ve después de crear y, al recargar, ya no | Crear un proyecto desde la web | RN-8, RN-9 | pendiente |
| 10 | Proyecto sin nombre o con fechas incoherentes: 422, lo escrito, el mensaje del dominio junto al campo y nada guardado | Mostrar el rechazo sin perder lo escrito | RN-12, RN-15 | pendiente |
| 11 | Una fecha que no es fecha se marca junto al campo y no se llama al dominio | Mostrar el rechazo sin perder lo escrito | RN-14 | pendiente |
| 12 | La página del proyecto: sus datos, «sin descripción», «sin fecha de fin», fechas dd/mm/aaaa y la barra con su nombre y sus secciones | Consultar un proyecto | RN-4, RN-5, RN-6 | pendiente |
| 13 | Modificar un proyecto con el formulario precargado: cambiar el nombre, y borrar la descripción la quita | Cada operacion se confirma con un mensaje | RN-1, RN-8 | pendiente |
| 14 | Proyecto 99, proyecto "abc" e historia 99: 404 con el mensaje del dominio y el acceso a la lista | Pedir algo que no existe | RN-16 | pendiente |
| 15 | Un nombre con `<`, `>` y `&` se ve tal cual | Un texto con signos se muestra tal cual | RN-17 | pendiente |
| 16 | Integrantes: lista vacía con su mensaje; agregar con el rol elegido entre las opciones del dominio | Cada operacion se confirma con un mensaje / Las opciones para elegir son las del sistema / Las listas vacias dicen que hacer | RN-1, RN-7, RN-11 | pendiente |
| 17 | Integrante duplicado y segundo Agile Enabler: el mensaje junto a su campo | Mostrar el rechazo sin perder lo escrito | RN-12 | pendiente |
| 18 | Quitar un integrante pide confirmación; confirmar lo quita con mensaje, cancelar no cambia nada | Confirmar o cancelar antes de quitar | RN-8, RN-10 | pendiente |
| 19 | Backlog: vacío con su mensaje; cargar una historia con criterios uno por renglón y la prioridad elegida; figura Pendiente y «sin estimar» | Cada operacion se confirma con un mensaje / Las opciones para elegir son las del sistema / Las listas vacias dicen que hacer | RN-1, RN-6, RN-7, RN-11 | pendiente |
| 20 | Historia sin título o con los criterios en blanco: el mensaje junto a su campo | Mostrar el rechazo sin perder lo escrito | RN-12 | pendiente |
| 21 | Modificar una historia cambia su prioridad; una Terminada, cargada con la conversión, muestra el rechazo arriba | Cada operacion se confirma con un mensaje / Mostrar el rechazo sin perder lo escrito | RN-1, RN-13 | pendiente |
| 22 | Estimar eligiendo de la escala del dominio; un valor fuera de la escala enviado igual muestra el mensaje del dominio junto al campo | Cada operacion se confirma con un mensaje / Las opciones para elegir son las del sistema | RN-1, RN-11, RN-15 | pendiente |
| 23 | Quitar la estimación pide confirmación; confirmar la quita, cancelar no | Confirmar o cancelar antes de quitar | RN-10 | pendiente |
| 24 | Con un guardado que falla, la operación responde 500 con el aviso de que no quedó guardada | Un cambio que no se pudo guardar se avisa | RN-25 | pendiente |
| 25 | Con el archivo real en una carpeta temporal: cargar, armar otra web desde el mismo archivo y ver todo; el proyecto nuevo recibe el número siguiente | Los datos siguen al cerrar y volver a abrir | RN-21, RN-22 | pendiente |
| 26 | Veinte creaciones en paralelo dejan veinte proyectos con números distintos, también con `go test -race` | Dos pestanas a la vez | RN-20 | pendiente |
| 27 | En cada formulario, todo campo tiene su etiqueta asociada y su error enlazado; toda tabla está dentro de su recuadro desplazable | Usar la pagina en un celular y con el teclado | RN-18, RN-19 | pendiente |
| 28 | El comando arma el servidor en `127.0.0.1:8080` con `datos-ics.json`, respeta `-puerto`, `-datos` y `-red`, muestra la dirección, y con datos dañados termina con el motivo | Abrir el sistema con un solo comando y sin internet / Los datos guardados estan danados | RN-2, RN-24 | pendiente |

**Verificación manual.** Lo que `httptest` no ve se comprueba a mano antes de abrir el PR y
se anota en la Fase 4: el contraste con un verificador, el modo celular del navegador a 360
de ancho, los formularios recorridos con el teclado mirando el foco, y el binario compilado
abierto con la red desconectada, en Chrome, Firefox y Safari.

### Por qué este orden

- **Del 1 al 5, el dominio y el archivo**: no dependen de la web ni de la lista de
  proyectos, y la 21 necesita la conversión para preparar una historia Terminada, que hoy
  nada más puede crear.
- **La 6 y la 7, el layout**; **del 8 al 15, proyectos**, donde se resuelve lo común
  —mensaje, rechazo, forma, no encontrado—; **del 16 al 23, integrantes y backlog**, que ya
  solo suman filas a los esquemas.
- **Del 24 al 27, lo transversal**, cuando hay pantallas sobre las que probarlo; **la 28 al
  final**, porque el comando solo conecta piezas ya probadas.

### Cobertura

Las veinticinco reglas quedan cubiertas: RN-1 (8, 13, 16, 19, 21, 22), RN-2 (28), RN-3 (6),
RN-4 (6, 12), RN-5 (12), RN-6 (12, 19), RN-7 (7, 16, 19), RN-8 (9, 13, 18), RN-9 (8, 9),
RN-10 (18, 23), RN-11 (1, 16, 19, 22), RN-12 (10, 17, 20), RN-13 (21), RN-14 (11),
RN-15 (10, 22), RN-16 (14), RN-17 (15), RN-18 (27), RN-19 (27), RN-20 (26),
RN-21 (2, 4, 25), RN-22 (2, 25), RN-23 (4, 7), RN-24 (3, 5, 28), RN-25 (5, 24). La RN-18 y
la RN-19 se completan con la verificación manual.

| Criterio de aceptación | Tareas |
|---|---|
| Se abre con un solo comando y funciona sin internet | 6, 28 |
| Se crean, modifican y consultan proyectos | 8, 12, 13 |
| Se agregan y quitan integrantes, con su rol | 16, 18 |
| Se administra el backlog: alta, modificación, prioridad y estimación | 19, 21, 22, 23 |
| Los errores van junto al campo, con el mensaje del sistema y sin perder lo escrito | 10, 11, 17, 20, 21 |
| Las listas vacías dicen que no hay datos y qué hacer | 7, 16, 19 |
| Se usa en notebook y en celular | 6, 27 y la verificación manual |
| Los datos no se pierden al cerrar y volver a abrir | 2, 4, 25 |
| Cada operación se confirma y recargar no la repite | 8, 9 |
| Antes de quitar se pide confirmación | 18, 23 |
| La barra muestra en qué proyecto se está | 12 |
| Etiquetas, foco y teclado | 27 y la verificación manual |
| Si los datos guardados no se pueden leer, avisa y no los pisa | 3, 5, 28 |
| La web no contiene validaciones de negocio propias | 10, 22 |

Los quince escenarios de la Fase 1 tienen su tarea en la columna "Escenario BDD".

### Dependencias

- **La lista de todos los proyectos** (la consulta que agrega Emiliano Sorato) tiene que estar en
  `dev` antes de la tarea 6: la página de inicio es esa lista. Las tareas 1 a 5 no la usan.
- **La 1 toca `story_estimation.go`, de la #12**: se avisa al grupo antes y, si la #12 sigue
  en review, se escribe después de su merge para no chocar.

> **Puerta 3** — ¿las tareas cubren todos los criterios de aceptación?
> Aprobó: Emiliano Sorato — Fecha: 2026-10-20 — Comentarios: aprobada en un comentario de la issue #77 (https://github.com/emisorato1/seguimiento-medicion-ics/issues/77#issuecomment-6299914339). La aprueba el responsable de la historia porque el resto del equipo no estaba disponible: es una excepción a la regla de que aprueba otra persona, y queda para que un compañero la ratifique con un comentario en la issue #77, desde su propia cuenta. Ningún PR de la historia entra a `dev` antes de esa ratificación.

## Fase 4 — Implementación

Ciclo RED → GREEN → REFACTOR. El test en un commit y la implementación en otro.

| Tarea | Commit del test | Commit de la implementación |
|---|---|---|
| Lista de todos los proyectos (dependencia, sin número) | `1f94f7a` | `3e81aca` |
| 1 | `4227227` | `5a6d632` |
| 2 | `1a885f0` | `61077ba` |
| 3 | `ff7d24f` | `2ecdfad` |
| 4 | `9e9ec91` | `fc414fb` |
| 5 | `6d73f7a` | `a83bfc5` |

Los RED de la lista de proyectos y de las tareas 1, 2, 4 y 5 fueron por compilación (la
operación o el error todavía no existían); el de la tarea 3, por comportamiento: 13 de sus
14 casos se cargaban sin error. La lista de proyectos la pide la Fase 2 sin ponerle nombre;
se llama `r.Projects()`, como en el contrato, y su test usa doce proyectos para que el orden
no salga bien por casualidad.

Nombres que la Fase 2 no fija: la conversión es `r.ToJSON()` y `RegistryFromJSON(texto)`;
el archivo, `store.NewFile(ruta)` con `Load()` y `Save(registro)`. Los errores son
`ErrSavedDataInvalid` en el dominio (datos guardados que no se pueden leer o no se aceptan)
y `ErrNotSaved` en `internal/store` («el cambio no quedó guardado»). Un rechazo por una
regla del dominio cumple `errors.Is` con los dos: con `ErrSavedDataInvalid` y con el error
de esa regla.

### Desvíos respecto de la spec

Ninguno de comportamiento. Tres agregados y una nota:

- **La tarea 3 tiene más casos que los cinco de la Fase 3**: fechas que no son fechas,
  número de proyecto repetido, contadores que retroceden o están en cero, historia sin
  título, prioridad y estado inexistentes. Cada uno prueba un control distinto de la
  reconstrucción; sin ellos el dominio no llegaba al 100%. El del estado inexistente pasó sin
  código nuevo: leer el estado por su nombre ya lo rechazaba desde la tarea 2.
- **Los contadores se controlan al reconstruir**: el próximo número tiene que ser al menos 1
  y mayor que todos los usados, y ninguno se repite. Sale de la D15 y de la RN-22: si no, un
  archivo editado a mano podía repartir un número ya usado.
- **Desvío respecto de la D16: un dato que el sistema no conoce rechaza la carga** en vez de
  ignorarse. La D16 solo dice que un dato que falta se carga vacío y no habla de los que
  sobran. Se decidió rechazarlos porque, si se ignoraran, un programa viejo que abre un
  archivo con datos de una historia posterior (por ejemplo los sprints) los cargaría sin
  ellos y los borraría en silencio al volver a guardar. Rechazarlos es lo que pide la RN-24
  para un archivo con algo que el núcleo no acepta. Lo aceptó el responsable el 20/10 y
  quedó escrito en la D16 de la Fase 2; queda para la ratificación.
- **`internal/store` queda en 95,5%**: lo único sin probar es el error de convertir el
  registro a texto, que con textos, números y listas `encoding/json` no puede dar. Se dejó
  igual, para que se avise si algún día un dato nuevo sí puede fallar.

### Hallazgos de la revisión del PR #92

Se encontraron antes de que entrara a `dev`, así que no se abrió defecto:

- **Se aceptaba texto sobrante después de los datos** («{…} basura» cargaba sin error),
  porque la lectura de JSON se detiene en el primer valor. Se agregó ese caso al test de la
  tarea 3 (`312cc45`, RED por comportamiento) y ahora se rechaza (`19ccd1b`).
- **El temporal no se forzaba al disco antes de renombrarlo**: un corte de luz justo después
  del renombre podía dejar el archivo vacío. Se agregó la sincronización (`Sync`, `6a39f92`) antes de
  cerrarlo. **No tiene test**: un corte de luz no se puede simular desde un test, y en el
  camino normal la línea ya queda cubierta por el test de la tarea 4.

Nota para la #16: si el PR #81 (horas estimadas) entra a `dev` antes que esta conversión, o
después, la que llegue segunda suma `estimatedHours` a `registry_json.go` con su caso en el
test de ida y vuelta, como pide la regla de Dependencias.

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #77 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | `internal/web/*_test.go`, `internal/store/file_test.go`, `internal/domain/registry_json_test.go`, `internal/domain/story_estimation_test.go` y `cmd/web/main_test.go` |
| Código Go | `internal/web/` (handlers, plantillas y CSS embebidos), `internal/store/file.go`, `internal/domain/registry_json.go`, la lectura de la escala en `internal/domain/story_estimation.go` y `cmd/web/main.go` |
