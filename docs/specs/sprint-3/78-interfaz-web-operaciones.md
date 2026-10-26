# Spec: Interfaz web: sprints, esfuerzo, defectos, Planning Poker y reportes

- **Historia:** #78
- **Sprint:** 3
- **Responsable:** Emiliano Sorato
- **Estado:** En implementación. Parte 1 en dev (PR #149): sprints sin el cierre, esfuerzo,
  horas estimadas, Planning Poker hasta revelar y el resumen del proyecto. Parte 2 en dev
  (PR #163): los defectos. Parte 3 en su PR, rama `historia-78-cierre-acuerdo`: cerrar el
  sprint, consultar uno anterior y el acuerdo del Planning Poker. Maximiliano Eula **aprobó las tres puertas el 2026-10-22** en la issue #78
  (https://github.com/emisorato1/seguimiento-medicion-ics/issues/78#issuecomment-6348317557), con dos cambios que esta spec ya incorpora (la
  tarea 17 y el punto 8 de "Para decidir en la puerta") y una nota sobre los casos límite
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de la consigna, de
  los criterios de aceptación de la issue y de las specs de las historias de dominio que
  muestra. El responsable lo revisa, lo corrige si hace falta y lo presenta a las puertas.
- **Depende de:** #77 (la base de la web: estructura, navegación, errores, confirmaciones y
  guardado) y, para cada pantalla, de su historia de dominio: #8 y #9 (crear un sprint y
  asignarle historias), #10 (cerrarlo), #11 (sprints anteriores), #14 (esfuerzo), #15
  (defectos), #13 y #73 (Planning Poker), #18 y #19 (reporte y PDF). Además, #5 (resumen del
  proyecto) y #16 (horas estimadas). Al 25/10 están en `dev` la #5, la #8, la #9, la #13, la
  #14, la #15 y las horas estimadas de la #16; faltan la #10, la #11, la #73, la #18 y la #19
- **La necesitan:** nadie todavía

## Fase 1 — Requisitos

### Objetivo

Que el equipo pueda llevar el proyecto entero desde el navegador: armar y cerrar sus sprints,
cargar las horas trabajadas y los defectos, estimar con Planning Poker sin ver los votos de
los demás y sacar el reporte del proyecto o de un sprint, también en PDF. Completa la web que
empezó la #77, con la misma forma de trabajar.

### Qué queda afuera

- **El dashboard** (#17): tiene su propia historia.
- **Reglas de negocio nuevas.** Igual que en la #77, la web pide los datos, se los pasa al
  núcleo y muestra lo que responde. Lo que se puede o no se puede hacer lo deciden las
  historias de dominio de cada pantalla.
- **Lo que el núcleo no permite**: editar o borrar un sprint, un registro de esfuerzo o un
  defecto; reabrir un sprint o un defecto resuelto; cambiar un voto; borrar una ronda.
- **Usuarios y permisos**: como en la #77, no hay inicio de sesión. Ver el punto 1 de "Para
  decidir en la puerta".
- **Lo que ya resolvió la #77** —barra, mensajes, errores, confirmaciones, guardado, celular
  y teclado— no se vuelve a decidir: se aplica (RN-2).

### Entradas

| Formulario | Datos obligatorios | Datos opcionales |
|---|---|---|
| Sprint (crear) | Sprint Goal, fecha de inicio, fecha de fin | — |
| Asignar al sprint abierto | la historia, de una lista | — |
| Cerrar el sprint abierto | fecha de cierre | las historias completadas, marcadas en la lista del sprint |
| Esfuerzo | historia, integrante, día, horas, actividad | — |
| Horas estimadas de una historia | horas | — |
| Defecto (cargar) | título, descripción, severidad, historia, sprint de detección | — |
| Cambiar el estado de un defecto | el estado nuevo (Abierto o En curso) | — |
| Resolver un defecto | sprint de resolución | — |
| Filtro de defectos | — | un estado o una severidad |
| Abrir una ronda de Planning Poker | la historia | — |
| Voto | quién vota, story points | — |
| Acuerdo | story points | — |
| Reporte | — | el sprint; sin sprint, es el del proyecto |

Las fechas se eligen en un calendario y lo demás, cuando tiene valores fijos o sale de lo
cargado (historia, integrante, sprint, severidad, estado, story points), de una lista
(RN-11 de la #77). Las horas admiten decimales.

### Salidas

Las páginas con los datos de cada parte y los datos cambiados y guardados según cada
operación aceptada:

- **Sprints**: la lista de sprints del proyecto (número, Sprint Goal, inicio, fin, estado y
  fecha de cierre); el sprint abierto con su Sprint Backlog (número, título, story points y
  estado de cada historia); el detalle de un sprint anterior, con sus historias asignadas y
  cuáles se completaron.
- **Esfuerzo**: los registros (día, integrante, historia, actividad y horas) y el total de
  horas por historia y por integrante.
- **Defectos**: la lista (número, título, severidad, estado, historia, sprint de detección y
  sprint de resolución), completa o filtrada.
- **Planning Poker**: las rondas de una historia; de la ronda en curso, quién votó, quién
  falta y cuántos votaron; de una ronda revelada, cada voto con su integrante, las
  diferencias y la estimación acordada.
- **Reporte**: su título, la fecha en que se generó y sus secciones; y el mismo reporte como
  archivo PDF para guardar.
- **Proyecto**: la página del proyecto de la #77 suma su resumen (historias totales,
  completadas y pendientes, y el sprint en curso).

### Reglas de negocio

1. **[RN-1]** El sistema debe ofrecer desde el navegador: crear un sprint, asignarle
   historias y quitárselas, cerrarlo y consultar los anteriores; registrar el esfuerzo y ver
   las horas por historia y por integrante; cargar las horas estimadas de una historia;
   cargar defectos, cambiar su estado, resolverlos y filtrarlos por estado o por severidad;
   abrir una ronda de Planning Poker, votar, revelar, ver las diferencias, abrir otra ronda y
   acordar la estimación; y ver el reporte del proyecto o de un sprint y descargarlo en PDF.
2. **[RN-2]** El sistema debe aplicar en todas estas páginas las reglas de la #77 tal como
   están: la barra fija (su RN-4), los textos en castellano y las fechas dd/mm/aaaa (RN-5),
   las tablas y los datos sin cargar (RN-6), las listas vacías (RN-7), el mensaje que
   confirma y no se repite al recargar (RN-8, RN-9), la confirmación antes de quitar
   (RN-10), las opciones para elegir (RN-11), los rechazos junto al campo o arriba del
   formulario (RN-12 a RN-15), lo que no existe (RN-16), los signos (RN-17), el celular y el
   teclado (RN-18, RN-19), varias pestañas (RN-20) y el guardado (RN-21 a RN-25). Como en la
   #77, cada lista para elegir —historia, integrante, sprint, severidad, estado, story points
   y quién vota— empieza con una opción vacía que pide elegir, para que nada quede elegido
   sin que nadie lo elija. Si se envía sin elegir, lo rechaza el núcleo cuando el dato es un
   nombre (integrante, severidad, estado, quién vota), y la web como un dato sin la forma
   esperada (RN-14 de la #77) cuando es un número (historia, sprint, story points).
3. **[RN-3]** Mientras se está dentro de un proyecto, la barra de navegación debe ofrecer,
   además de lo que suma la #77, el acceso a sus sprints, su esfuerzo, sus defectos, su
   Planning Poker y su reporte.
4. **[RN-4]** Mientras una ronda no esté revelada, el sistema no debe mostrar ningún voto en
   ninguna página ni en ningún mensaje, ni siquiera a quien lo dio: solo quién votó, quién
   falta y cuántos votaron de cuántos.
5. **[RN-5]** Cuando se registra un voto, el sistema debe confirmarlo nombrando a quien votó
   y sin el valor, por ejemplo «Se registró el voto de Ana Perez».
6. **[RN-6]** Si el núcleo rechaza un voto, el sistema debe volver a mostrar el formulario
   con el mensaje, pero **sin** el valor elegido. Es la única excepción a la RN-12 de la #77,
   que pide volver a mostrar lo escrito.
7. **[RN-7]** Mientras una ronda no esté revelada, el sistema debe ofrecer para votar solo a
   los participantes que todavía no votaron.
8. **[RN-8]** Cuando una ronda está revelada, el sistema debe mostrar cada voto con su
   integrante y, si hay diferencias, el voto más bajo y el más alto con quiénes los dieron,
   tal como los informa el núcleo.
9. **[RN-9]** Antes de cerrar un sprint o de resolver un defecto, que no se deshacen, el
   sistema debe pedir confirmación como antes de quitar en la #77; si se cancela, no cambia
   nada.
10. **[RN-10]** Cuando se cierra un sprint, el sistema debe mostrar sus historias para marcar
    las completadas, ninguna marcada de entrada, y en la confirmación decir cuáles quedan
    completadas y cuáles vuelven al Product Backlog.
11. **[RN-11]** El sistema debe proponer la fecha de hoy como fecha de cierre de un sprint y
    como día de un registro de esfuerzo; quien carga la puede cambiar.
12. **[RN-12]** El sistema debe tomar como fecha de hoy la de la computadora donde corre,
    para todo lo que el núcleo necesite: rechazar horas de un día que todavía no pasó y
    fechar el reporte.
13. **[RN-13]** Cuando se filtran los defectos por un estado o por una severidad, el sistema
    debe mostrar solo los que cumplen; sin filtro, todos. Se filtra por uno de los dos a la
    vez.
14. **[RN-14]** El sistema debe mostrar el reporte con su título, su fecha de generación y
    sus secciones en el orden y con las líneas que arma el núcleo, sin cambiarles el formato
    a los números ni a las fechas.
15. **[RN-15]** Cuando se pide descargar el reporte, el sistema debe entregar el PDF de ese
    mismo reporte como un archivo para guardar, llamado `reporte-proyecto-N.pdf` o
    `reporte-proyecto-N-sprint-M.pdf`, con N el número del proyecto y M el del sprint.
16. **[RN-16]** Las listas para elegir deben mostrar todo lo que el núcleo tiene cargado —por
    ejemplo, todas las historias del backlog para asignar al sprint—, sin quitar las que el
    núcleo después va a rechazar: el rechazo lo da el núcleo, con su mensaje (RN-15 de la
    #77).
17. **[RN-17]** Si una lista para elegir no tiene ninguna opción —un proyecto sin
    integrantes al registrar esfuerzo, sin sprints al cargar un defecto, sin historias al
    abrir una ronda—, el sistema debe decirlo en lugar del formulario y ofrecer el acceso
    para cargar lo que falta.
18. **[RN-18]** Si se pide un sprint, un defecto o una ronda que no existe, el sistema debe
    mostrar el mensaje del núcleo y ofrecer volver al proyecto.
19. **[RN-19]** Si el PDF no se puede generar, el sistema debe avisarlo con el mensaje del
    núcleo y no entregar ningún archivo.
20. **[RN-20]** Donde el sistema ya permita consultar el estado de un proyecto (#5), la
    página del proyecto debe mostrar su resumen: historias totales, completadas y
    pendientes, y el sprint en curso o que no hay ninguno.

### Casos límite

- Proyecto sin sprints: la lista de sprints vacía con su mensaje, y el acceso para crear el
  primero. Sin sprint abierto, no se ofrece asignar ni cerrar, y se dice que no hay ninguno.
- Cerrar un sprint sin marcar ninguna historia: se acepta, y todas vuelven al Product
  Backlog (#10).
- Una historia que no se completó en un sprint y se vuelve a asignar en el siguiente: figura
  en los dos, completada solo donde se completó.
- Horas con decimales («2,5»): se aceptan; un día igual a hoy se acepta y uno de mañana se
  rechaza (#14).
- Un defecto resuelto ya no ofrece cambiar de estado ni resolverse otra vez.
- Una ronda con un solo participante: se vota, se revela y no hay diferencias.
- Recargar la página de una ronda sin revelar después de votar: sigue sin mostrar el valor.
- Un integrante que se suma al proyecto con la ronda abierta no aparece para votar (#13).
- Un integrante que se quitó del proyecto con la ronda abierta sigue apareciendo para votar,
  porque la lista sale de los que faltan según el núcleo (#13, su punto 1). No es un error.
- El reporte de un sprint abierto se muestra igual, con el aviso de números parciales que
  arma el núcleo (#18).

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| El núcleo rechaza los datos de un formulario | lo vuelve a mostrar con lo escrito y el mensaje junto al campo, como la #77 |
| El núcleo rechaza un voto | vuelve a mostrar el formulario con el mensaje y sin el valor elegido |
| El rechazo no corresponde a un campo (por ejemplo, revelar con votos faltantes) | lo muestra arriba del formulario |
| Una fecha o unas horas que no tienen la forma esperada | lo indica junto al campo y no ejecuta la operación |
| Sprint, defecto, ronda, historia o proyecto que no existen | muestra el mensaje del núcleo y ofrece volver al proyecto |
| El PDF no se pudo generar | lo avisa con el mensaje del núcleo y no entrega archivo |
| Un cambio no se pudo guardar | avisa que no quedó guardado, como la #77 |

### Criterios de aceptación

Los seis primeros son los de la issue.

- [ ] Desde la web se crea un sprint, se le asignan historias, se cierra y se consultan los
      anteriores.
- [ ] Desde la web se registra el esfuerzo real y se ven las horas por historia y por
      integrante.
- [ ] Desde la web se cargan defectos, se cambia su estado, se resuelven y se filtran por
      estado y por severidad.
- [ ] Desde la web se hace una ronda de Planning Poker: cada integrante vota sin ver los votos
      de los demás hasta que se revelan, y se acuerda la estimación.
- [ ] Desde la web se ve el reporte del proyecto o de un sprint, y se descarga en PDF.
- [ ] Los errores se muestran igual que en el resto de la interfaz web.
- [ ] Antes de revelar, ningún voto aparece en ninguna página ni mensaje, ni siquiera el
      propio.
- [ ] Antes de cerrar un sprint o resolver un defecto se pide confirmación, y cancelar no
      cambia nada.
- [ ] La fecha de cierre y el día del esfuerzo vienen propuestos con la fecha de hoy.
- [ ] Desde la web se cargan las horas estimadas de una historia.
- [ ] La barra de un proyecto lleva a sus sprints, esfuerzo, defectos, Planning Poker y
      reporte.
- [ ] Los sprints, el esfuerzo, los defectos y las rondas siguen al cerrar y volver a abrir.
- [ ] La web no contiene ninguna validación de negocio propia.

### Escenarios BDD

```gherkin
Escenario: La barra de un proyecto lleva a todas sus secciones
  Dado el proyecto "Sistema de metricas"
  Cuando entro a cualquiera de sus paginas
  Entonces la barra muestra "Sistema de metricas" y el acceso a sus datos, integrantes, backlog, sprints, esfuerzo, defectos, Planning Poker y reporte

Escenario: Crear un sprint desde la web
  Dado un proyecto sin sprints
  Cuando envio el formulario de sprint con el goal "Tener el MVP", inicio 26/10/2026 y fin 30/10/2026
  Entonces veo el mensaje "Se creo el sprint 1"
  Y el sprint 1 figura "Abierto" en la lista de sprints
  Y si recargo la pagina, sigue habiendo un solo sprint y el mensaje ya no aparece

Escenario: Asignar y quitar historias del sprint abierto
  Dado el sprint 1 abierto y las historias 1 y 2 estimadas en 3 y 5 puntos
  Cuando asigno la historia 1 y despues la 2, eligiendolas de la lista
  Entonces veo el mensaje "Se asigno la historia 2 al sprint 1"
  Y el Sprint Backlog muestra las historias 1 y 2 "En curso"
  Y si pido quitar la historia 1 y confirmo, solo queda la 2 y la 1 vuelve a "Pendiente"

Escenario: Cerrar el sprint marcando las completadas
  Dado el sprint 1 abierto con las historias 1 y 2 asignadas
  Y hoy es 30/10/2026
  Cuando abro el formulario para cerrar el sprint
  Entonces la fecha de cierre propuesta es 30/10/2026 y ninguna historia esta marcada
  Cuando marco la historia 1 y envio
  Entonces el sistema me pide confirmarlo y dice que la 1 queda completada y la 2 vuelve al Product Backlog
  Y si confirmo, veo el mensaje "Se cerro el sprint 1" y el sprint figura "Cerrado"
  Y no hay sprint abierto

Escenario: Consultar un sprint anterior
  Dado el sprint 1 cerrado el 30/10/2026 con las historias 1 y 2 asignadas y solo la 1 completada
  Cuando entro a la pagina del sprint 1
  Entonces veo su goal, sus fechas, "Cerrado" y la fecha de cierre 30/10/2026
  Y veo las historias 1 y 2 con su titulo, la 1 como completada y la 2 como no completada

Escenario: Registrar esfuerzo y ver los totales
  Dado el integrante "Ana Perez", la historia 1 y hoy 27/10/2026
  Cuando abro el formulario de esfuerzo
  Entonces el dia propuesto es 27/10/2026
  Cuando registro 2,5 horas de "Ana Perez" en la historia 1 con la actividad "Escribir tests"
  Entonces veo el mensaje "Se registraron 2,5 horas de Ana Perez en la historia 1"
  Y el registro figura en la tabla de esfuerzo
  Y el total de la historia 1 y el de "Ana Perez" son 2,5 horas

Escenario: Cargar las horas estimadas de una historia
  Dado la historia 1 sin horas estimadas
  Cuando le cargo 6,5 horas estimadas
  Entonces veo el mensaje "Se cargaron las horas estimadas de la historia 1"
  Y la historia 1 figura con 6,5 horas estimadas

Escenario: Cargar un defecto, pasarlo a en curso y resolverlo
  Dado la historia 1 y los sprints 1 y 2
  Cuando cargo el defecto "Suma mal las horas" con severidad "Mayor", la historia 1 y detectado en el sprint 1
  Entonces veo el mensaje "Se cargo el defecto 1" y figura "Abierto"
  Cuando lo paso a "En curso"
  Entonces figura "En curso"
  Cuando pido resolverlo en el sprint 2 y confirmo
  Entonces veo el mensaje "Se resolvio el defecto 1 en el sprint 2" y figura "Resuelto"
  Y ya no se ofrece cambiar su estado ni resolverlo

Esquema del escenario: Filtrar los defectos
  Dado el defecto 1 "Abierto" de severidad "Critica" y el defecto 2 "En curso" de severidad "Menor"
  Cuando filtro por <filtro>
  Entonces veo solo <defectos>

  Ejemplos:
    | filtro               | defectos          |
    | el estado "Abierto"  | el defecto 1      |
    | la severidad "Menor" | el defecto 2      |
    | nada                 | los defectos 1 y 2 |

Escenario: Una ronda de Planning Poker sin ver los votos hasta revelar
  Dado los integrantes "Ana Perez" y "Juan Gomez" y la historia 1
  Cuando abro una ronda para la historia 1
  Entonces veo el mensaje "Se abrio la ronda 1 de la historia 1"
  Y para votar se ofrecen "Ana Perez" y "Juan Gomez"
  Cuando "Ana Perez" vota 5
  Entonces veo el mensaje "Se registro el voto de Ana Perez"
  Y la ronda muestra que voto "Ana Perez", que falta "Juan Gomez" y que votaron 1 de 2
  Y para votar solo se ofrece "Juan Gomez"
  Cuando "Juan Gomez" vota 8 y revelo la ronda
  Entonces veo el voto de "Ana Perez" en 5 y el de "Juan Gomez" en 8

Escenario: Nadie ve un voto antes de revelar
  Dado una ronda de la historia 1 sin revelar en la que "Ana Perez" voto 13
  Cuando recorro todas las paginas del proyecto y recargo la de la ronda
  Entonces ninguna pagina muestra la tabla de votos
  Y en ninguna pagina aparece el 13 junto a "Ana Perez"
  Y el mensaje que confirma su voto no lleva ningun numero

Escenario: Un voto rechazado no muestra el valor elegido
  Dado una ronda de la historia 1 sin revelar en la que "Ana Perez" ya voto
  Cuando envio otro voto de "Ana Perez" con 21
  Entonces el formulario vuelve con el mensaje "ya voto en esta ronda"
  Y el 21 no aparece en la pagina

Escenario: No se puede revelar con votos faltantes
  Dado una ronda de la historia 1 en la que falta votar "Juan Gomez"
  Cuando pido revelarla
  Entonces arriba del formulario veo el mensaje "todavia faltan votos"
  Y la ronda sigue sin revelar

Escenario: Diferencias, nueva ronda y acuerdo
  Dado la ronda 1 de la historia 1 revelada con 3 de "Ana Perez" y 8 de "Juan Gomez"
  Cuando entro a la ronda 1
  Entonces veo que hay diferencias: el mas bajo es 3, de "Ana Perez", y el mas alto 8, de "Juan Gomez"
  Cuando abro una ronda nueva y los dos votan 5 y la revelo
  Entonces veo que no hay diferencias
  Cuando acuerdo 5 story points
  Entonces veo el mensaje "Se acordo la historia 1 en 5 story points"
  Y la historia 1 figura en el backlog con 5 story points
  Y la ronda 1 sigue en la lista de rondas de la historia 1, sin acordar

Esquema del escenario: Ver el reporte
  Dado el proyecto 1 con el sprint 1 cerrado y hoy 31/10/2026
  Cuando abro el reporte de <ambito>
  Entonces veo el titulo "<titulo>" y "Generado el 31/10/2026"
  Y veo las secciones Resumen, Historias, Estimacion y esfuerzo y Defectos, en ese orden, con las lineas que arma el sistema

  Ejemplos:
    | ambito       | titulo                                       |
    | el proyecto  | Reporte del proyecto Sistema de metricas     |
    | el sprint 1  | Reporte del sprint 1 — Sistema de metricas   |

Escenario: Descargar el reporte en PDF
  Dado el reporte del sprint 1 del proyecto 1
  Cuando pido descargarlo en PDF
  Entonces recibo un archivo para guardar llamado "reporte-proyecto-1-sprint-1.pdf"
  Y el archivo es un PDF con el mismo titulo y las mismas lineas que veo en la pagina

Esquema del escenario: Confirmar o cancelar antes de lo que no se deshace
  Dado el sprint 1 abierto y el defecto 1 "Abierto"
  Cuando pido <accion>
  Entonces el sistema me pide confirmarlo
  Y si <respuesta>, <resultado>

  Ejemplos:
    | accion                            | respuesta | resultado                          |
    | cerrar el sprint 1                | confirmo  | el sprint 1 figura "Cerrado"       |
    | cerrar el sprint 1                | cancelo   | el sprint 1 sigue "Abierto"        |
    | resolver el defecto 1 en el sprint 1 | confirmo  | el defecto 1 figura "Resuelto"  |
    | resolver el defecto 1 en el sprint 1 | cancelo   | el defecto 1 sigue "Abierto"    |

Esquema del escenario: Mostrar el rechazo sin perder lo escrito
  Dado <situacion>
  Cuando envio el formulario de <formulario> con <dato>
  Entonces el formulario vuelve a mostrarse con lo que escribi
  Y <donde> veo el mensaje del sistema <mensaje>
  Y no cambio ningun dato

  Ejemplos:
    | situacion                          | formulario | dato                                | donde                    | mensaje                                                        |
    | un proyecto sin sprints            | sprint     | el goal vacio                       | junto al goal            | "el Sprint Goal es obligatorio"                                |
    | el sprint 1 abierto                | sprint     | un goal valido                      | arriba del formulario    | "el sprint 1 ya esta abierto"                                  |
    | el sprint 1 abierto                | asignar    | la historia 3, sin estimar          | junto a la historia      | "solo se asignan historias estimadas"                          |
    | hoy 27/10/2026                     | esfuerzo   | el dia 28/10/2026                   | junto al dia             | "no se pueden cargar horas de un dia que todavia no paso"      |
    | hoy 27/10/2026                     | esfuerzo   | 0 horas                             | junto a las horas        | "las horas tienen que ser mayores que cero"                    |
    | los sprints 1 y 2                  | defecto    | el titulo vacio                     | junto al titulo          | "falta el titulo"                                              |
    | el defecto 1 detectado en el sprint 2 | resolver | el sprint 1                        | junto al sprint          | "un defecto no se resuelve en un sprint anterior al de deteccion" |

Esquema del escenario: Las listas vacias dicen que hacer
  Dado <situacion>
  Cuando abro <pagina>
  Entonces veo el mensaje <mensaje>
  Y el acceso para <accion>

  Ejemplos:
    | situacion                    | pagina                 | mensaje                                    | accion                  |
    | un proyecto sin sprints      | sus sprints            | "Todavia no hay sprints"                   | crear el primero        |
    | un proyecto sin esfuerzo     | su esfuerzo            | "Todavia no hay horas registradas"         | registrar las primeras  |
    | un proyecto sin defectos     | sus defectos           | "Todavia no hay defectos"                  | cargar el primero       |
    | un proyecto sin integrantes  | el formulario de esfuerzo | "Este proyecto todavia no tiene integrantes" | agregar el primero   |

Esquema del escenario: Pedir algo que no existe
  Dado que el proyecto 1 no tiene sprint 99, defecto 99 ni historia 99
  Cuando abro la pagina de <pedido>
  Entonces veo el mensaje del sistema <mensaje>
  Y el acceso para volver al proyecto

  Ejemplos:
    | pedido                            | mensaje                      |
    | el sprint 99                      | "no se encontro el sprint"   |
    | el defecto 99                     | "no se encontro el defecto"  |
    | las rondas de la historia 99      | "no se encontro la historia" |

Escenario: El proyecto muestra su resumen
  Dado el proyecto 1 con 3 historias, 1 terminada, y el sprint 2 abierto
  Cuando entro a su pagina
  Entonces veo 3 historias en total, 1 completada y 2 pendientes, y que el sprint en curso es el 2

Escenario: Los datos nuevos siguen al cerrar y volver a abrir
  Dado que cargue el sprint 1, un registro de esfuerzo, el defecto 1 y una ronda revelada
  Cuando cierro el sistema y lo vuelvo a abrir
  Entonces veo el sprint 1, el registro, el defecto 1 y la ronda con sus votos
  Y el siguiente sprint que creo recibe el numero 2
```

**Para decidir en la puerta:**

1. **Quién vota, sin usuarios.** Quien vota se elige de la lista de los que faltan. El voto es
   secreto frente a la pantalla —nadie lo ve hasta revelar (RN-4)—, pero nada impide elegir el
   nombre de otro. Para un equipo que vota junto alcanza; impedirlo exige usuarios, que la
   #77 dejó afuera.
   **Decidido por el responsable (Emiliano Sorato) el 21/10: se acepta como está.**
2. **Los votos sin revelar quedan en el archivo de datos**, porque las rondas entran en lo que
   se guarda (contrato, PR #94). Quien abra el archivo con un editor los ve. La alternativa
   —no guardar las rondas abiertas— las pierde al cerrar el sistema. Lo que se guarde de los
   votos sin revelar lo decide la #13 en su puerta (su spec propone guardarlas con sus votos,
   en el PR #115); la #78 lo respeta, y la web no cambia en ningún caso.
   **Decidido por el responsable (Emiliano Sorato) el 21/10: se acepta como está.**
3. **Las listas para elegir no filtran por reglas** (RN-16). Para asignar se ofrecen todas
   las historias, aunque la #9 rechace las que no están estimadas, las Terminadas y las ya
   asignadas: filtrarlas en la web copiaría reglas de la #9 fuera del núcleo. Ofrecer solo
   las que se pueden asignar exige una consulta nueva en el núcleo, que el contrato no tiene.
   La excepción es votar (RN-7): la lista de los que faltan sí la da el núcleo (#13).
   **Decidido por el responsable (Emiliano Sorato) el 21/10: se acepta como está.**
4. **Un filtro de defectos por vez** (RN-13). El núcleo filtra por estado o por severidad,
   cada uno por su lado; combinarlos sería una regla en la web.
   **Decidido por el responsable (Emiliano Sorato) el 21/10: se acepta como está.**
5. **Confirmación también antes de cerrar un sprint y de resolver un defecto** (RN-9), que no
   se deshacen (#10, #15). La #77 la pedía solo para quitar. No se pide antes de revelar
   —solo se puede con todos los votos— ni antes de acordar —se puede abrir otra ronda y
   acordar de nuevo—.
   **Decidido por el responsable (Emiliano Sorato) el 21/10: se acepta como está.**
6. **Las horas estimadas (#16) y el resumen del proyecto (#5) entran acá.** La #77 dejó el
   resumen para la #78 (su punto 6), y sin cargar horas estimadas desde la web la desviación
   de horas da siempre cero. El resumen depende de que la #5 llegue a `dev` (RN-20).
   **Decidido por el responsable (Emiliano Sorato) el 21/10: se acepta como está.**
7. **Cada pantalla se conecta cuando su historia de dominio está en `dev`.** La issue dice
   que hasta entonces «se prueba contra el contrato», pero una pantalla llama al núcleo y no
   se puede conectar a una operación que todavía no existe. Las tareas siguen el orden
   en que llegan las historias de dominio, y se pueden entregar por partes (D13).
   **Decidido por el responsable (Emiliano Sorato) el 21/10: se acepta como está.**
8. **El estado de un defecto llega como texto.** Lo señaló Maximiliano Eula al aprobar las
   puertas. La RN-2 dice que, si se envía sin elegir, el estado lo rechaza el núcleo, pero la
   #15 solo convierte desde texto la severidad (`ParseSeverity`): `ChangeDefectState` recibe
   el estado como valor del dominio. Si la web tradujera «Abierto» y «En curso» por su cuenta,
   copiaría una regla fuera del núcleo, que es lo que la #4 y la #15 evitaron con el rol y la
   severidad. Hay dos caminos: que la #15 sume la conversión del estado igual que la de la
   severidad (ya tiene una sin exportar, para la carga del archivo), o que la RN-2 diga cómo
   se trata el estado. **Resuelto el 26/10 por el primer camino: el dominio exporta
   `ParseDefectState`, que acepta «Abierto», «En curso» y «Resuelto» y rechaza el resto con
   `ErrDefectStateUnknown` (PR #153, Refs #15). La tarea 12 arranca cuando entre a dev.**

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: Maximiliano Eula — Fecha: 2026-10-22 — Comentarios: en un comentario de la issue #78 (https://github.com/emisorato1/seguimiento-medicion-ics/issues/78#issuecomment-6348317557), desde su cuenta. Aprueba las tres puertas y acepta los puntos 1 a 7 de "Para decidir en la puerta". Pide dos cambios antes de las tareas que los usan: que el test de la tarea 17 no busque el número suelto (incorporado en el escenario y en la tarea) y resolver el estado de un defecto como texto (punto 8, pendiente). Nota: sumar a los casos límite al integrante que se quitó con la ronda abierta (incorporado).

## Fase 2 — Diseño

Las decisiones de [AGENTS.md](../../../AGENTS.md) y **las de la #77, de la D1 a la D22**,
valen igual y no se repiten: biblioteca estándar, plantillas y CSS embebidos, nada de
JavaScript, formularios por POST con redirección, el rechazo con lo escrito y el mensaje
junto a su campo, la tabla que asocia cada error del dominio con su campo, el mensaje de una
sola vez, la página intermedia para confirmar, el candado, el guardado del registro entero y
los tests con `httptest`. Esta fase registra solo lo propio de la #78.

### Enfoque

Es la **convención que la #77 dejó para la #78**: una pantalla nueva es un archivo de
handlers, sus plantillas y su test, y del layout solo se toca la barra. Son cinco pantallas
—sprints, esfuerzo, defectos, Planning Poker y reporte— y cada una sigue el recorrido de la
#77: controlar la forma, tomar el candado, llamar al dominio, y si rechaza volver a mostrar el
formulario y si acepta guardar, dejar el mensaje y redirigir.

Lo único que cambia en ese recorrido es de dónde sale la **fecha de hoy**: el servidor la
recibe al armarse, como una forma de preguntar la fecha, y la pasa al dominio cuando una
operación la necesita (#14, #18). El comando le da la de la computadora; los tests, una fija.

Las direcciones siguen la forma de la #77 y cuelgan del proyecto: `/proyectos/{id}/sprints`
y `/proyectos/{id}/sprints/{n}`, `/proyectos/{id}/esfuerzo`, `/proyectos/{id}/defectos` y
`/proyectos/{id}/defectos/{n}`, `/proyectos/{id}/historias/{n}/rondas` y
`/proyectos/{id}/historias/{n}/rondas/{r}`, y `/proyectos/{id}/reporte` con el sprint como
parámetro opcional. Los sprints y los defectos se nombran por su número dentro del proyecto,
igual que las historias.

### Archivos afectados

| Archivo | Se crea o se modifica | Para qué |
|---|---|---|
| `internal/web/sprints.go` y `sprints_test.go` | se crean | crear un sprint, asignar y quitar historias, cerrarlo, la lista y el detalle |
| `internal/web/efforts.go` y `efforts_test.go` | se crean | registrar el esfuerzo, los totales y las horas estimadas de una historia |
| `internal/web/defects.go` y `defects_test.go` | se crean | cargar defectos, cambiar su estado, resolverlos y filtrarlos |
| `internal/web/poker.go` y `poker_test.go` | se crean | abrir rondas, votar, revelar, las diferencias y el acuerdo |
| `internal/web/reports.go` y `reports_test.go` | se crean | ver el reporte y descargar el PDF |
| `internal/web/templates/` | se modifica | una plantilla por página nueva; en el layout, la barra suma las secciones |
| `internal/web/server.go` | se modifica | las rutas nuevas y la fecha de hoy que recibe al armarse |
| `internal/web/forms.go` | se modifica | la tabla de errores suma, con su campo, los errores de cada historia de dominio que llega |
| `internal/web/projects.go` | se modifica | el resumen de la #5 en la página del proyecto |
| `cmd/web/main.go` | se modifica | le da al servidor la fecha de la computadora |
| `internal/web/static/app.css` | se modifica | lo mínimo para las páginas nuevas (la ronda y el reporte) |

**No se modifica nada de `internal/domain` ni de `internal/store`.** Si a una pantalla le
falta algo del núcleo, se pide en la historia de dominio que corresponde (ver Dependencias).

### Decisiones

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| **D1** La fecha de hoy la recibe el servidor al armarse y la pasa al dominio en cada operación que la necesita; los tests la fijan | que el dominio lea el reloj; que la escriba quien usa la página | AGENTS.md: el dominio no lee el reloj. Fijada en el test, el escenario «hoy es 27/10/2026» se prueba igual cualquier día |
| **D2** La página de una ronda sin revelar se arma solo con lo que el núcleo deja leer antes de revelar —quién votó, quién falta, cuántos votaron— y un test recorre todas las páginas y respuestas buscando el valor votado | ocultar los votos con CSS; mostrarle su voto a quien votó | el secreto lo garantiza el núcleo (#13), que no deja leer un voto sin revelar. Ocultarlo con estilos lo dejaría en el HTML, a un «ver código fuente» de distancia |
| **D3** Al rechazar un voto, el formulario vuelve sin el valor elegido (RN-6) | volver a mostrar lo escrito, como pide la #77 | mostrar el valor en pantalla rompería la RN-4 delante de quien esté mirando |
| **D4** Quien vota se elige de la lista de los que faltan, que da el núcleo | un campo de texto con el nombre; usuarios con contraseña | sin usuarios (#77), elegir de la lista evita errores de tipeo y solo ofrece a quien puede votar; escribir el nombre igual lo valida el núcleo |
| **D5** La descarga usa la exportación de la #19, que escribe el PDF donde se le pida (`r.WritePDF`, contrato del PR #94): la web lo arma en memoria y, si salió bien, lo envía como archivo para guardar con su nombre (RN-15) | escribirlo en un archivo temporal y enviarlo; guardarlo en una carpeta del servidor y mostrar la ruta | quien usa la página no ve el disco donde corre el sistema y, con la opción de red de la #77, puede ni siquiera estar en la misma computadora; además, el dominio no toca archivos. La #19 había tomado «se descarga» como «se guarda donde se indica» porque la aplicación no era web: se decidió el 20/10, al escribir esta spec, que la #19 escriba el PDF donde se le pida y que la web lo descargue (contrato, PR #94). La consola no exporta: queda con lo que especifica la #24, y todo lo nuevo se usa desde la web (plan, 20/10). Armarlo en memoria antes de enviar es lo que garantiza que un error no deje un PDF cortado a medio descargar (RN-19) |
| **D6** El reporte se muestra como una lista de líneas por sección, tal como las arma el núcleo, sin tabla y sin reformatear | armar una tabla con los números de las métricas | el formato es una regla del núcleo (#18, AGENTS.md): si la web lo rearmara, la página y el PDF podrían mostrar el mismo número distinto |
| **D7** El filtro de defectos va en la dirección, como parámetro, y es uno por vez | un formulario por POST; combinar los dos filtros | filtrar no cambia nada, así que se pide por GET y la página filtrada se puede recargar o guardar. Combinarlos sería una regla nueva fuera del núcleo (punto 4) |
| **D8** Cerrar el sprint es un solo formulario con la fecha y una casilla por historia, y la confirmación muestra cuáles quedan completadas y cuáles vuelven al backlog | marcar las completadas una por una y cerrar aparte | la #10 recibe todo junto, en una sola lista (su punto 3), y marcar de a una dejaría un estado intermedio que el núcleo no tiene |
| **D9** La confirmación de la #77 (D9) se usa también para cerrar un sprint y resolver un defecto | confirmar solo al quitar | no se deshacen (#10, #15). Es la misma página intermedia, sin JavaScript |
| **D10** Las listas para elegir salen de las consultas del núcleo sin filtrar por reglas: el backlog para asignar, abrir rondas o registrar esfuerzo; los integrantes; los sprints; las severidades y los estados con su nombre. Las del backlog van en el orden del listado ordenado de la #7, tal como quede en `dev`; mientras la #7 no esté, en el orden del núcleo | filtrar en la web las que el núcleo va a rechazar; ordenarlas en la web | es la D7 y la D10 de la #77 aplicadas a las listas nuevas (punto 3). La única lista filtrada, la de los que faltan votar, la filtra el núcleo. El orden de la #7 pone primero lo que queda por hacer, que es lo que se elige al planificar, y ordenar en la web copiaría su regla fuera del núcleo |
| **D11** Las horas se leen como número con coma o con punto, y lo que no es número se marca junto al campo sin llamar al dominio. **«NaN» e «Inf» (y sus variantes, como «infinity») también se marcan como «no es un número»**, aunque la conversión de texto a número de Go los acepta | dejar pasar el texto y que lo rechace el dominio | es controlar la forma, como las fechas en la #77 (D7): el dominio recibe horas, no texto. Que sean mayores que cero lo decide el dominio, pero el infinito es mayor que cero y pasaría, y NaN no es una cantidad de horas: ninguno de los dos es un número que alguien escriba como horas |
| **D12** El resumen del proyecto (#5) se suma a la página del proyecto de la #77 solo cuando la #5 está en `dev` | una página aparte; esperarla para empezar la #78 | la #77 lo dejó así (su punto 6), y el resto de la #78 no depende de ella |
| **D13** Cada pantalla se entrega cuando su historia de dominio está en `dev`, en PRs con `Refs #78`, y el último cierra la historia | esperar a que estén todas para un solo PR | son cinco pantallas sobre nueve historias de dominio de cuatro personas. Esperarlas a todas deja la web entera para el último día. Es el caso 1 de CLAUDE.md para partir una historia |

### Dependencias y restricciones

- **Depende de la #77 entera en `dev`**: layout, barra, errores por campo, mensaje de una
  sola vez, confirmación, candado y guardado.
- **Cada pantalla depende de su historia de dominio**, y no se empieza hasta que esté en
  `dev` (punto 7):

  | Pantalla | Historias de dominio |
  |---|---|
  | Sprints | #8, #9, #10, #11 |
  | Esfuerzo y horas estimadas | #14, #16 |
  | Defectos | #15 (y #8, porque un defecto se detecta en un sprint) |
  | Planning Poker | #13, #73 |
  | Reporte y PDF | #18, #19 (y por ellas, #16) |
  | Resumen del proyecto | #5 |

- **Lo que se guarda lo suma cada historia de dominio** a la conversión de la #77, con su caso
  en el test de ida y vuelta. La web no agrega nada: guarda el registro entero. Si una
  historia de dominio no lo suma, sus datos se pierden al cerrar y el último escenario falla.
- **Lo que el contrato no tenía y esta spec necesitaba.** No se inventó acá: el contrato se
  ajustó en el PR #94 (20/10).
  1. **El PDF sin pasar por un archivo**: la #19 lo escribe donde se le pida (`r.WritePDF`).
  2. **Volver a encontrar la ronda en curso** entre un pedido y otro:
     `p.CurrentVotingRound(storyID)`, de la #13. Las rondas entran en lo que se guarda.
  3. **La fecha de hoy en el reporte**: `p.ProjectReport(today)` y
     `p.SprintReport(number, today)`.
- **Lo que define cada historia de dominio**, y la web usa tal como llegue: la forma de lo que
  devuelven las diferencias y la lectura de la estimación acordada (#73), el nombre de la
  consulta del resumen (#5) y los nombres y textos de los errores nuevos. La web no depende
  del texto: los tests lo toman del propio error.
- **No se toca** `internal/domain`, `internal/store` ni `cmd/app` (la consola, #24), ni el
  dashboard (#17).

> **Puerta 2** — ¿el diseño es viable y coherente con el resto del proyecto?
> Aprobó: Maximiliano Eula — Fecha: 2026-10-22 — Comentarios: en un comentario de la issue #78 (https://github.com/emisorato1/seguimiento-medicion-ics/issues/78#issuecomment-6348317557), desde su cuenta. Aprueba las tres puertas y acepta los puntos 1 a 7 de "Para decidir en la puerta". Pide dos cambios antes de las tareas que los usan: que el test de la tarea 17 no busque el número suelto (incorporado en el escenario y en la tarea) y resolver el estado de un defecto como texto (punto 8, pendiente). Nota: sumar a los casos límite al integrante que se quitó con la ronda abierta (incorporado).

## Fase 3 — Tareas

Cada tarea es un ciclo de TDD completo, con el test que falla en un commit y la implementación
en otro, igual que en la #77: los tests arman un registro nuevo, el guardado falso y una
fecha de hoy fija; buscan en la página el texto del error tomado del propio error del
dominio, nunca escrito a mano; y cada esquema es un solo test con tabla, al que las tareas que
lo comparten le suman sus filas. Un test que pase sin código nuevo va a Desvíos.

| # | Tarea | Escenario BDD | Reglas | Estado |
|---|---|---|---|---|
| 1 | La barra de un proyecto suma sprints, esfuerzo, defectos, Planning Poker y reporte; el servidor recibe la fecha de hoy | La barra de un proyecto lleva a todas sus secciones | RN-3, RN-12 | hecha |
| 2 | Sprints: lista vacía con su mensaje; crear un sprint redirige, queda «Abierto» con su número y recargar no crea otro | Crear un sprint desde la web / Las listas vacias dicen que hacer | RN-1, RN-2 | hecha |
| 3 | Goal vacío junto al campo y segundo sprint abierto arriba del formulario | Mostrar el rechazo sin perder lo escrito | RN-2 | hecha |
| 4 | Asignar desde la lista de todo el backlog y quitar con confirmación; el Sprint Backlog con su estado | Asignar y quitar historias del sprint abierto | RN-1, RN-16 | hecha |
| 5 | Asignar una historia sin estimar: el mensaje del dominio junto a la historia | Mostrar el rechazo sin perder lo escrito | RN-2, RN-16 | hecha |
| 6 | Cerrar el sprint: fecha de hoy propuesta, casillas sin marcar, confirmación con lo que queda completado y lo que vuelve; confirmar cierra, cancelar no | Cerrar el sprint marcando las completadas / Confirmar o cancelar antes de lo que no se deshace | RN-9, RN-10, RN-11 | hecha |
| 7 | Lista de sprints y detalle de uno anterior con los títulos de sus historias y cuáles se completaron | Consultar un sprint anterior | RN-1 | hecha |
| 8 | Sprint 99: el mensaje del dominio y el acceso al proyecto | Pedir algo que no existe | RN-18 | hecha |
| 9 | Esfuerzo: lista vacía, día propuesto hoy, registrar con integrante e historia de la lista, la tabla y los totales por historia y por integrante | Registrar esfuerzo y ver los totales / Las listas vacias dicen que hacer | RN-1, RN-11 | hecha |
| 10 | Día futuro y cero horas junto a su campo; unas horas que no son número se marcan sin llamar al dominio; sin integrantes, el aviso en lugar del formulario | Mostrar el rechazo sin perder lo escrito / Las listas vacias dicen que hacer | RN-2, RN-12, RN-17 | hecha |
| 11 | Cargar las horas estimadas de una historia | Cargar las horas estimadas de una historia | RN-1 | hecha |
| 12 | Defectos: lista vacía; cargar con severidad, historia y sprint de la lista; cambiar a «En curso» (punto 8: `ParseDefectState`) | Cargar un defecto, pasarlo a en curso y resolverlo / Las listas vacias dicen que hacer | RN-1, RN-16 | hecha |
| 13 | Resolver con confirmación; un resuelto no ofrece cambiar ni resolver | Cargar un defecto, pasarlo a en curso y resolverlo / Confirmar o cancelar antes de lo que no se deshace | RN-9 | hecha |
| 14 | Filtrar por un estado o por una severidad, y sin filtro todos | Filtrar los defectos | RN-13 | hecha |
| 15 | Título vacío, resolver en un sprint anterior y defecto 99 | Mostrar el rechazo sin perder lo escrito / Pedir algo que no existe | RN-2, RN-18 | hecha |
| 16 | Abrir una ronda; quién votó, quién falta y cuántos; votar eligiendo de los que faltan, con el mensaje sin el valor | Una ronda de Planning Poker sin ver los votos hasta revelar | RN-1, RN-5, RN-7 | hecha |
| 17 | Ningún valor antes de revelar en ninguna página, mensaje ni recarga; un voto rechazado vuelve sin el valor. El test no busca el número suelto, porque el formulario de votar ofrece la escala entera y el número puede estar en una fecha o un identificador: comprueba que la tabla de votos no existe hasta revelar, que ninguna fila ni elemento que nombra a quien votó lleva su valor, y que el mensaje que confirma el voto no lleva ningún número (revisión de Maximiliano Eula) | Nadie ve un voto antes de revelar / Un voto rechazado no muestra el valor elegido | RN-4, RN-6 | hecha |
| 18 | Revelar: con votos faltantes, el mensaje arriba; con todos, cada voto con su integrante | No se puede revelar con votos faltantes / Una ronda de Planning Poker sin ver los votos hasta revelar | RN-2, RN-8 | hecha |
| 19 | Diferencias, ronda nueva, acuerdo y la lista de rondas de una historia; historia 99 | Diferencias, nueva ronda y acuerdo / Pedir algo que no existe | RN-1, RN-8, RN-18 | hecha |
| 20 | El reporte del proyecto y el de un sprint, con la fecha de hoy y las líneas del dominio sin reformatear | Ver el reporte | RN-1, RN-12, RN-14 | espera la #18 |
| 21 | Descargar el PDF con su nombre como archivo para guardar; si no se puede generar, el aviso y ningún archivo | Descargar el reporte en PDF | RN-15, RN-19 | espera la #19 |
| 22 | El resumen de la #5 en la página del proyecto | El proyecto muestra su resumen | RN-20 | hecha |
| 23 | Con el archivo real: sprints, esfuerzo, defectos y rondas siguen al reabrir, y el sprint nuevo recibe el número siguiente | Los datos nuevos siguen al cerrar y volver a abrir | RN-2 | hecha: también el cierre, el acuerdo y el sprint siguiente |
| 24 | En cada formulario nuevo, todo campo tiene su etiqueta y su error enlazado, y toda tabla está dentro de su recuadro | — (RN-18 y RN-19 de la #77, por la RN-2) | RN-2 | en parte: los formularios hechos, también los de defectos, cierre y acuerdo; falta el del reporte |

**Verificación manual**, como en la #77: el contraste, el modo celular a 360 de ancho, los
formularios nuevos recorridos con el teclado, una ronda votada desde dos pestañas y el PDF
descargado abierto en el visor del sistema, en Chrome, Firefox y Safari. Se anota en la
Fase 4.

### Por qué este orden

- **La 1 va primero**: la barra y la fecha de hoy las usan todas las demás.
- **De la 2 a la 8, sprints**, porque los defectos (sprint de detección), el reporte de un
  sprint y el último escenario los necesitan. Siguen la cadena de dominio: crear (#8),
  asignar (#9), cerrar (#10) y consultar (#11).
- **De la 9 a la 15, esfuerzo y defectos**, que no dependen entre sí: si la #15 llega antes
  que la #14, se invierten.
- **De la 16 a la 19, Planning Poker**, que es una cadena aparte (#13 → #73). La 17 va
  apenas después de votar porque protege la regla que define la historia.
- **La 20 y la 21, reporte y PDF**, al final de la cadena de dominio (#16 → #18 → #19).
- **De la 22 a la 24, lo transversal**, cuando hay pantallas sobre las que probarlo. La 22
  puede ir antes si la #5 llega antes.

### Cobertura

Las veinte reglas quedan cubiertas: RN-1 (2, 4, 7, 9, 11, 12, 16, 19, 20), RN-2 (2, 3, 5, 10,
15, 18, 23, 24), RN-3 (1), RN-4 (17), RN-5 (16), RN-6 (17), RN-7 (16), RN-8 (18, 19), RN-9 (6,
13), RN-10 (6), RN-11 (6, 9), RN-12 (1, 10, 20), RN-13 (14), RN-14 (20), RN-15 (21), RN-16 (4,
5, 12), RN-17 (10), RN-18 (8, 15, 19), RN-19 (21), RN-20 (22).

| Criterio de aceptación | Tareas |
|---|---|
| Sprint: crear, asignar, cerrar y consultar los anteriores | 2, 4, 6, 7 |
| Esfuerzo y horas por historia y por integrante | 9 |
| Defectos: cargar, cambiar de estado, resolver y filtrar | 12, 13, 14 |
| Planning Poker con los votos ocultos y el acuerdo | 16, 17, 18, 19 |
| Reporte del proyecto o de un sprint, y su PDF | 20, 21 |
| Los errores como en el resto de la web | 3, 5, 10, 15, 18 |
| Ningún voto visible antes de revelar | 17 |
| Confirmación antes de cerrar o resolver | 6, 13 |
| La fecha de hoy propuesta | 6, 9 |
| Horas estimadas desde la web | 11 |
| La barra lleva a todas las secciones | 1 |
| Los datos nuevos siguen al reabrir | 23 |
| La web no contiene validaciones de negocio propias | 4, 5, 10, 12 |

Los veintidós escenarios de la Fase 1 tienen su tarea en la columna "Escenario BDD".

### Dependencias

- **Ninguna tarea empieza sin la #77 en `dev`**, y cada grupo espera su historia de dominio
  (tabla de la Fase 2). La 1 solo necesita la #77.
- **De la 16 a la 18 vuelven a encontrar la ronda con `p.CurrentVotingRound`** de la #13, así
  que no esperan a la #73; la 19 sí, porque las diferencias, el acuerdo y la lista de rondas
  son de la #73.
- **La 21 usa `r.WritePDF` de la #19** (D5).
- **La 23 necesita que cada historia de dominio haya sumado su entidad a la conversión.**

> **Puerta 3** — ¿las tareas cubren todos los criterios de aceptación?
> Aprobó: Maximiliano Eula — Fecha: 2026-10-22 — Comentarios: en un comentario de la issue #78 (https://github.com/emisorato1/seguimiento-medicion-ics/issues/78#issuecomment-6348317557), desde su cuenta. Aprueba las tres puertas y acepta los puntos 1 a 7 de "Para decidir en la puerta". Pide dos cambios antes de las tareas que los usan: que el test de la tarea 17 no busque el número suelto (incorporado en el escenario y en la tarea) y resolver el estado de un defecto como texto (punto 8, pendiente). Nota: sumar a los casos límite al integrante que se quitó con la ronda abierta (incorporado).

## Fase 4 — Implementación

Se implementa en la rama `historia-78-web`, PR #149, con `Refs #78`. Las tareas van en el
orden en que su historia de dominio llegó a `dev` (punto 7, D13).

### Commits por tarea

Cada tarea tiene el test en un commit, con su salida en rojo después de `RED:`, y la
implementación en el siguiente.

| Tarea | Test (RED) | Implementación (GREEN) |
|---|---|---|
| 1. Barra: Sprints | `fb7c03e` | `56d1f95` |
| 2. Crear un sprint y la lista vacía | `cc2b293` | `211237f` |
| 3. Rechazos del sprint | `4aad7ae` | `485f145` |
| 4. Asignar y quitar con confirmación | `f3a96de` | `5077131` |
| 5. Asignar una historia sin estimar | `73f2f68` | `03ad1a0` |
| 8. Sprint 99 y volver al proyecto | `93be502` | `1813b0d` |
| 9. Esfuerzo, día de hoy y totales; la barra suma Esfuerzo | `587fd47` | `97bb41b` |
| 10. Rechazos del esfuerzo y aviso sin integrantes | `7e8b4d6` | `bd1088b` |
| 11. Horas estimadas | `e8124f5` | `cf1d745` |
| 16. Abrir una ronda y votar; la barra suma Planning Poker | `8f37017` | `4297f72` |
| 17. Ningún voto antes de revelar; voto rechazado sin su valor | `fff8bab` | `cba14ef` |
| 18. Revelar | `5acab35` | `06a195a` |
| 22. Resumen del proyecto (#5) | `e10078c` | `b951ba3` |
| 23. Los datos siguen al reabrir (en parte) | `97ef3a2` | — (pasa sin código nuevo) |
| 24. Teclado y celular en los formularios nuevos (en parte) | `499e7a2` | — (pasa sin código nuevo) |
| RN-18 en cada dirección nueva | `7c8ec8b` | — (pasa sin código nuevo) |
| 12. Cargar un defecto y pasarlo a en curso; la barra suma Defectos | `3b14453` | `9611795` |
| 13. Resolver con confirmación; el resuelto no ofrece cambios | `92d4364` | `bf4b303` |
| 14. Filtrar por un estado o por una severidad | `9090625` | `7066111` |
| 15. Rechazos de los defectos junto a su campo y defecto 99 | `8847aa0` | `153b2a6` |
| 23 y 24. Defectos al reabrir y sus formularios con el teclado | `e04ef0b` | — (sin RED: pasa sin código nuevo) |
| Rechazos que quedaban sin probar (cobertura) | `b7523ca` | — (sin RED: prueba código existente) |
| 6. Cerrar el sprint con las completadas y confirmación | `8b60975` | `3550e3f` |
| 7. Consultar un sprint anterior | `e4476d7` | `d7b0785` |
| 19. Diferencias, ronda nueva, acuerdo y lista de rondas | `43e3e5a` | `4ad7a5b` |
| 23 y 24. Cierre, acuerdo y sprint 2 al reabrir; sus formularios | `44abfe2` | — (sin RED: pasa sin código nuevo) |
| Fechas y casillas mal formadas al cerrar (cobertura) | `09d23fa` | — (sin RED: prueba código existente) |

### Desvíos

- **La barra suma cada sección junto con su pantalla**, y no todas en la tarea 1. Así `dev`
  nunca tiene un enlace a una página que todavía no existe. El test de la barra es una tabla
  y cada pantalla le suma su fila. Por la misma razón, **la fecha de hoy (D1) entró con la
  tarea 9**, la primera que la usa, y no con la 1: `NewServer` recibe de dónde tomarla, el
  comando le da `time.Now` y los tests una fija (27/10/2026).
- **La página de Planning Poker del proyecto** (`/proyectos/{id}/planning-poker`) no estaba
  entre las direcciones del Enfoque. La pide la barra (RN-3) y la entrada «abrir una ronda:
  la historia», que se elige de una lista: muestra cada historia con su ronda actual y el
  formulario para abrir una. La ronda queda en `/proyectos/{id}/historias/{n}/rondas/{r}`,
  como dice el Enfoque.
- **Asignar y quitar van a `/proyectos/{id}/sprints/asignar` y `/quitar`**, sin el número del
  sprint: el núcleo asigna siempre al sprint abierto (#9), y un número en la dirección
  podría no coincidir con él.
- **Una ronda que la historia no tiene** se informa con un error de la web, «no se encontró
  la ronda» (`errRoundNotFound`), como la página que no existe de la #77: el núcleo da la
  ronda actual de una historia, y la web la compara con el número de la dirección.
- **Mensajes con tilde**: la spec escribe los escenarios sin tildes («Se creo el sprint 1»)
  y la web los muestra con tilde («Se creó el sprint 1»), igual que en la #77.
- **El mensaje del rechazo se busca sin el cierre del elemento** en el esquema «Mostrar el
  rechazo sin perder lo escrito»: el núcleo le suma al sprint ya abierto su número («…: el
  sprint 1»), y el test toma el mensaje del propio error, sin escribirlo a mano.
- **Tests que pasaron sin código nuevo**, como pide la Fase 3. Sus commits no dicen «sin
  RED» en el asunto, como sí hacen los de la #10; para no reescribir la historia del PR, la
  marca queda acá con el hash de cada uno:
  - el segundo sprint abierto va arriba del formulario: lo hace la tabla de errores de la
    #77, que manda arriba lo que no es de ningún campo (tarea 3, `4aad7ae`);
  - asignar sin elegir la historia y unas horas que no son número (`dos`): los controla la
    forma que ya traían las tareas 4 y 9 (tareas 5 y 10);
  - **«Nadie ve un voto antes de revelar»** (tarea 17): el secreto lo garantiza el núcleo
    (D2), porque `Votes` no entrega nada antes de revelar. Para saber que el test sirve, se
    hizo que la plantilla mostrara el voto y el test falló en las dos lecturas de la ronda;
    después se volvió atrás;
  - los datos nuevos siguen al reabrir (tarea 23): lo que se guarda lo suma cada historia de
    dominio y la web guarda el registro entero (Fase 2, Dependencias; `97ef3a2`);
  - los formularios nuevos con etiqueta, error enlazado y foco, y las tablas en su recuadro
    (tarea 24, `499e7a2`): las plantillas usan las partes comunes de la #77;
  - las filas nuevas de «Pedir algo que no existe» (`7c8ec8b`).
- **El test de la tarea 9 falló por compilación**: cambió la forma de armar el servidor
  (`NewServer` con la fecha de hoy), y ese error de compilación es su RED.

- **Un arreglo de CSS sin test de Go**: la verificación manual encontró que a 360 de ancho
  el sprint con una historia asignada medía 526 px. El texto oculto «Quitar la historia 1»
  es `position: absolute` y, como el recuadro `.tabla` no estaba posicionado, se ubicaba
  respecto de la página y la ensanchaba. Se arregló con `position: relative` en `.tabla`,
  que cubre también la tabla de integrantes de la #77. Un test de Go no mide el ancho
  dibujado; el RED y el GREEN son las dos corridas del script de Playwright de abajo.

- **Dos arreglos de tests después de traer dev el 26/10**, que no cambian código Go:
  - con la #10, la carga pide la fecha de cierre de un sprint Cerrado, y el resumen del
    proyecto se armaba con uno sin ella: se le sumó `closed_on`;
  - en Windows las plantillas pueden bajar con CRLF, y cuatro tests buscaban filas con
    «\n» y la sangría (revisión de Agustín): ahora quitan los espacios entre etiquetas
    (`joinTags`). El RED se reprodujo con una copia del repo con las plantillas en CRLF.

- **Defectos (parte 2)**:
  - **Un filtro que no es de la lista**, escrito a mano en la dirección (`?estado=Cerrado`),
    lo rechaza el núcleo (`ParseDefectState`, `ParseSeverity`): la lista se ve completa,
    con el mensaje arriba. **Si llegan los dos filtros se usa el estado**: la página no
    los puede mandar juntos (son dos formularios), y combinarlos sería una regla en la
    web (punto 4). Ninguno de los dos casos está en los Ejemplos.
  - **El cambio de estado ofrece Abierto y En curso**, como dice Entradas: a Resuelto se
    llega resolviendo, con el sprint. El filtro sí ofrece los tres estados.
  - **Resolver**: elegir el sprint lleva a la confirmación (D9); si no se eligió ninguno,
    o si el núcleo rechaza el sprint, se vuelve a la página del defecto con el error junto
    al sprint, como pide el ejemplo «resolver en un sprint anterior».
  - **Filas que no están en los Ejemplos**: la severidad y el estado sin elegir (los
    rechaza el núcleo, RN-2), el sprint sin elegir (forma) y resolver sin sprint.
  - **Pasaron sin código nuevo**, con «sin RED» en el asunto: el defecto sin elegir el
    sprint y las filas de «Pedir algo que no existe» (tarea 15, anotadas en su commit), y
    las tareas 23 y 24 con los defectos (se le sacó la etiqueta a la severidad y el test
    falló, para saber que sirve).
  - **Cobertura**: review-merge observó al mergear el #149 que `internal/web` había bajado
    de 93,8% a 92,5%. Había tres caminos de rechazo sin ningún test (abrir una ronda de
    una historia que no existe, quitar del sprint una historia no asignada y el filtro que
    no existe); con sus tests queda en 93,7%.

### Verificación manual de la parte 1

El 26/10, con el script de Playwright de la #77 (fuera del repo) ampliado con las
pantallas de esta parte, en Chromium 151, Firefox 153 y WebKit 26.5 (el motor de Safari):

- con el teclado solo: crear un sprint, asignar una historia, quitarla con confirmación y
  cancelar; registrar 2,5 horas (el día viene con la fecha de hoy); cargar las horas
  estimadas; abrir una ronda, votar sin que se vea el voto y revelarla;
- el foco se ve en todo el recorrido y ningún pedido sale del programa;
- a 360 de ancho, ninguna de las 21 páginas se desplaza de costado (después del arreglo
  de arriba);
- después de matar el proceso y volver a abrirlo siguen el sprint, las horas y la ronda
  revelada.

Resultado: 90 controles bien en los tres motores. Falló solo el control de la #77 que
espera el 8080 ocupado por Docker: ese día Docker no estaba corriendo y el puerto
estaba libre, así que no es un error del programa. Queda para el final: la ronda votada
desde dos pestañas, el contraste de lo nuevo y el PDF.

- **Cierre y acuerdo (parte 3)**:
  - **La confirmación del cierre lleva las completadas en un solo campo oculto**
    (`completadas=1,2`), porque la página de confirmación de la #77 pasa un valor por
    campo. Algo que no es un número ahí no viene de las casillas: va arriba del formulario.
  - **Una fecha de cierre que no es fecha** vuelve al formulario sin preguntar, con las
    casillas como estaban; **una vacía** sí llega a la confirmación («¿Cerrar el sprint 1?»,
    sin fecha) y al confirmar la rechaza el núcleo junto a la fecha (#10, RN-10).
  - **Un sprint cerrado muestra los story points que tenía al cerrarse** (`PointsAtClose`
    de la #10), no los de hoy, y una columna «Completada» con Sí o No.
  - **La ronda se busca entre todas las de la historia** (`VotingRounds`), así se ven las
    anteriores; solo la última sin acuerdo ofrece acordar y abrir una ronda nueva, que es
    el mismo formulario de la tarea 16 con la historia ya elegida.
  - **CSS**: la Fase 2 nombraba solo la ronda y el reporte; se sumaron las reglas de las
    casillas del cierre (la casilla al lado de su texto).
  - **El test del RED de la tarea 19 buscaba mal una fila del backlog** (los story points
    pegados al título, y la fila tiene antes la prioridad y el estado); se corrigió en el
    commit de la implementación, que lo dice.
  - **Pasaron sin código nuevo**, con «sin RED» en el asunto: los rechazos del cierre y del
    acuerdo, que se escribieron con las tareas 6 y 19 porque las necesitaban para volver a
    mostrar el formulario; los datos al reabrir; y los formularios con el teclado. Se probó
    que sirven sacando una línea de la tabla de errores y la marca de las casillas.
  - **Cobertura**: `internal/web` queda en 94,0%.

### Verificación manual de la parte 2

El 26/10, con el mismo script ampliado con los defectos, en Chromium 151, Firefox 153 y
WebKit 26.5: con el teclado solo, cargar un defecto, filtrar por una severidad sin
defectos y volver a verlos todos, pasarlo a En curso, pedir resolverlo y cancelar, y
resolverlo confirmando; el resuelto ya no tiene formularios. A 360 de ancho ninguna de
las 26 páginas se desplaza de costado, y el defecto sigue después de reiniciar.
Resultado: 114 controles bien; falló solo el del 8080 ocupado, por lo mismo que en la
parte 1 (Docker no estaba corriendo).

### Verificación manual de la parte 3

El 26/10, con el mismo script, en Chromium 151, Firefox 153 y WebKit 26.5: con el teclado,
acordar la estimación de una ronda revelada y ver la lista de rondas, y cerrar el sprint
marcando una historia con la barra espaciadora, con su confirmación. El sprint cerrado y su
historia completada siguen después de reiniciar, y a 360 de ancho ninguna de las 27
páginas se desplaza de costado. Resultado: 135 controles bien; falló solo el del 8080, por
lo mismo que en las partes 1 y 2.

### Revisión de accesibilidad de las partes 1 a 3

El 26/10, con un script de Playwright aparte (fuera del repo), sobre 37 páginas con datos
de todas las pantallas, incluidos tres formularios rechazados y el enlace para saltar la
barra con el foco: el contraste de cada texto visible contra su fondo (WCAG AA: 4,5 a 1, o
3 a 1 si es grande), un solo `h1` por página, los títulos sin saltos de nivel y que todo
enlace y botón tenga nombre. Resultado: ningún problema. Para saber que el control del
contraste sirve, se le agregó a una página un texto gris claro y lo marcó (1,79 a 1).

### Lo que falta

- **Reporte y PDF (tareas 20 y 21)**: esperan la #18 y la #19.
- **La verificación manual** del reporte y su PDF, y la ronda votada desde dos pestañas.

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #78 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | `internal/web/sprints_test.go`, `efforts_test.go`, `defects_test.go`, `poker_test.go`, `reports_test.go` y los de la #77 que suman filas |
| Código Go | `internal/web/` (handlers, plantillas y CSS embebidos) y `cmd/web/main.go` |
