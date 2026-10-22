# Spec: Votar en secreto en una ronda de Planning Poker

- **Historia:** #13
- **Sprint:** 3
- **Responsable:** Maximiliano Eula
- **Estado:** Puerta 3 aprobada — **lista para implementar**
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de la consigna y
  de los criterios de aceptación de la issue. El responsable lo revisa, lo corrige si hace
  falta y lo presenta a las puertas. Revisado por el responsable el 21/10: decidió los
  puntos de la puerta, dio por resuelto el 7 con el contrato del PR #94, cerró todas las
  formas de imprimir la ronda y escribió qué se valida al cargar las rondas guardadas.
- **Depende de:** #4 (los integrantes son los que votan), #6 (la historia y su estado) y #12
  (la escala de story points); las tres ya están en `dev`
- **La necesitan:** #73 (diferencias, rondas nuevas y estimación acordada)

## Fase 1 — Requisitos

### Objetivo

Que el equipo pueda estimar una historia con Planning Poker dentro del sistema: abrir una
ronda, que cada integrante vote sin ver lo que votaron los demás y dar vuelta todos los votos
juntos cuando votaron todos. Es lo que evita que alguien se ancle en el número de otro, la
regla que más importa de la [guía del equipo](../../guia-planning-poker.md).

### Qué queda afuera

- **Lo que pasa después de revelar**: informar si hay diferencias, el voto más bajo y el más
  alto, abrir una ronda nueva y registrar la estimación acordada. Todo eso es la #73. En esta
  historia la historia que se vota **no cambia**: abrir, votar y revelar no la estiman.
- **Una segunda ronda sobre la misma historia**: hasta que esté la #73, cada historia tiene
  una sola ronda (punto 6 de "Para decidir en la puerta").
- **Las cartas "?" y "café"** de la guía (punto 5), **cambiar el voto** antes de revelar
  (punto 3), **cancelar una ronda** o sacar a alguien de ella.
- **Quién puede abrir o revelar**: cualquiera; la ronda no distingue roles.

### Entradas

| Dato | Obligatorio | Descripción |
|---|---|---|
| Historia | sí, al abrir | qué historia se vota, por su identificador |
| Integrante | sí, al votar | quién vota, por su nombre y apellido |
| Voto | sí, al votar | el valor que elige, de la escala 1, 2, 3, 5, 8, 13 y 21 |

### Salidas

Una ronda de votación sobre la historia, con su número y sus participantes. Mientras no se
revela, solo se puede saber quién votó y quién falta. Una vez revelada, se ven todos los votos
con el nombre de quien los dio.

### Reglas de negocio

1. **[RN-1]** Cuando se abre una ronda sobre una historia, el sistema debe registrarla con esa
   historia, el número 1, sin revelar y sin votos.
2. **[RN-2]** Cuando se abre una ronda, el sistema debe fijar como participantes a los
   integrantes que el proyecto tiene en ese momento. Los que se agreguen o se quiten después
   no cambian quiénes participan de esa ronda.
3. **[RN-3]** El sistema debe permitir abrir una ronda sobre una historia en cualquier estado
   salvo Terminada, esté estimada o no. Abrir la ronda no cambia la historia.
4. **[RN-4]** Si la historia está Terminada, el sistema debe rechazar la apertura e informar
   que la estimación de una historia terminada no se modifica.
5. **[RN-5]** Si la historia no existe en el proyecto, el sistema debe informar que no se
   encontró la historia.
6. **[RN-6]** Si el proyecto no tiene integrantes, el sistema debe rechazar la apertura e
   informar que no hay integrantes que puedan votar.
7. **[RN-7]** Mientras una historia tenga una ronda sin revelar, el sistema debe rechazar que
   se abra otra sobre la misma historia e informar que ya hay una ronda en curso.
8. **[RN-8]** Cuando un participante vota, el sistema debe registrar su voto y pasarlo de los
   que faltan a los que votaron.
9. **[RN-9]** Si el voto no es un valor de la escala de la #12, el sistema debe rechazarlo e
   informar cuáles son los valores válidos.
10. **[RN-10]** Si el participante ya votó en la ronda, el sistema debe rechazar el nuevo voto
    e informar que ya votó, sin decir qué votó. Queda el primer voto.
11. **[RN-11]** Si quien vota no es participante de la ronda, el sistema debe rechazar el voto
    e informar que no participa. El nombre se busca igual que en la #4 —ignorando mayúsculas
    y los espacios de los extremos— y el voto queda con el nombre tal como está cargado.
12. **[RN-12]** Mientras la ronda no esté revelada, el sistema debe rechazar cualquier pedido
    de ver los votos —también el de quien quiere ver el suyo— e informar que los votos están
    ocultos hasta que se revele la ronda. Ninguna respuesta ni aviso del sistema debe incluir
    el valor de un voto.
13. **[RN-13]** El sistema debe permitir consultar en cualquier momento quiénes votaron y
    quiénes faltan, en el orden de los integrantes del proyecto. Una lista vacía no es un
    error.
14. **[RN-14]** Cuando se pide revelar y votaron todos los participantes, el sistema debe
    dejar la ronda revelada y, desde entonces, mostrar todos los votos con el nombre de quien
    los dio, en el orden de los integrantes.
15. **[RN-15]** Si se pide revelar y falta algún voto, el sistema debe rechazarlo e informar
    que todavía faltan votos. La ronda sigue sin revelar.
16. **[RN-16]** Si la ronda ya está revelada, el sistema debe rechazar un voto nuevo y un
    nuevo pedido de revelar, e informar que la ronda ya fue revelada.
17. **[RN-17]** Cuando se muestra una ronda sin revelar, el sistema debe mostrar su número, la
    historia y cuántos participantes votaron, sin ningún valor votado.
18. **[RN-18]** Si un voto o un pedido de revelar se rechaza, el sistema debe dejar la ronda
    como estaba.
19. **[RN-19]** El sistema debe permitir consultar la ronda actual de una historia, que es la
    última que se abrió sobre ella, esté revelada o no. Si la historia todavía no tiene
    ninguna ronda, debe informarlo, y eso no es un error. Si la historia no existe, debe
    informar que no se encontró la historia. *(Agregado del contrato del 20/10: punto 8 de
    "Para decidir en la puerta".)*

### Casos límite

- Un solo integrante: vota él y la ronda se puede revelar. Nadie votó: revelar se rechaza.
- Votar 1 o 21, los extremos de la escala: válido.
- Integrante que se agrega después de abrir: no vota en esa ronda. Integrante que se quita
  después de abrir: sigue participando y la ronda espera su voto.
- Dos historias del mismo proyecto con su ronda en curso a la vez: válido, son independientes.
- Historia sobre la que nunca se abrió una ronda: consultar su ronda actual informa que no
  tiene ninguna, y no es un error.

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| La historia no existe | informa que no se encontró la historia |
| La historia está Terminada | rechaza e informa que su estimación no se modifica |
| El proyecto no tiene integrantes | rechaza e informa que no hay quién vote |
| La historia ya tiene una ronda en curso | rechaza e informa que ya hay una ronda en curso |
| Voto fuera de la escala, cero o negativo | rechaza e informa los valores válidos |
| Segundo voto del mismo integrante | rechaza e informa que ya votó, sin decir qué |
| Quien vota no participa de la ronda | rechaza e informa que no participa |
| Ver los votos de una ronda sin revelar | rechaza e informa que están ocultos hasta revelar |
| Revelar con votos faltantes | rechaza e informa que todavía faltan votos |
| Votar o revelar en una ronda revelada | rechaza e informa que la ronda ya fue revelada |

### Criterios de aceptación

- [ ] Se puede abrir una ronda sobre una historia no Terminada, esté estimada o no, y queda
      como la ronda 1, sin revelar y sin votos.
- [ ] No se puede abrir una ronda sobre una historia Terminada ni sobre una que no existe.
- [ ] No se puede abrir una ronda en un proyecto sin integrantes.
- [ ] No se puede abrir otra ronda sobre una historia que tiene una sin revelar.
- [ ] Cada participante vota una sola vez por ronda, y el segundo voto no muestra el primero.
- [ ] El voto tiene que ser un valor de la escala 1, 2, 3, 5, 8, 13 o 21.
- [ ] Solo votan los integrantes que tenía el proyecto al abrir la ronda, y el nombre se
      reconoce aunque venga con otras mayúsculas o espacios de más.
- [ ] Sin revelar, nadie puede ver ningún voto, ni siquiera el propio; solo se sabe quién
      votó y quién falta.
- [ ] Mostrar una ronda sin revelar no deja ver ningún valor votado.
- [ ] La ronda se revela recién cuando votaron todos; antes, el pedido se rechaza.
- [ ] Al revelarse se ven todos los votos con el nombre de quien los dio.
- [ ] No se puede votar ni volver a revelar una ronda ya revelada.
- [ ] Un voto o un pedido de revelar rechazado deja la ronda como estaba.
- [ ] Se puede consultar la ronda actual de una historia, revelada o no; si no tiene ninguna,
      se informa y no es un error.

### Escenarios BDD

```gherkin
Esquema del escenario: Abrir una ronda sobre una historia que no esta terminada
  Dado un proyecto con "Agustin Salinas", "Juliana Bustos" y "Maximiliano Eula"
  Y la historia 1 en estado <estado> y <estimacion>
  Cuando abro una ronda de votacion sobre la historia 1
  Entonces la ronda queda registrada con numero 1 y sin revelar
  Y no voto nadie y faltan votar los tres
  Y la historia 1 sigue <estimacion>

  Ejemplos:
    | estado      | estimacion      |
    | "Pendiente" | "sin estimar"   |
    | "Pendiente" | "estimada en 5" |
    | "En curso"  | "estimada en 3" |

Esquema del escenario: Rechazar una ronda sobre una historia que no se puede votar
  Dado un proyecto con integrantes y <caso>
  Cuando intento abrir una ronda de votacion sobre la historia <id>
  Entonces la operacion se rechaza
  Y se informa que <aviso>
  Y el proyecto sigue sin rondas

  Ejemplos:
    | caso                                    | id | aviso                                                    |
    | "la historia 1 en estado Terminada"     | 1  | "la estimacion de una historia terminada no se modifica" |
    | "ninguna historia con identificador 99" | 99 | "no se encontro la historia"                             |

Escenario: Rechazar una ronda en un proyecto sin integrantes
  Dado un proyecto sin integrantes y con la historia 1 en estado "Pendiente"
  Cuando intento abrir una ronda de votacion sobre la historia 1
  Entonces la operacion se rechaza
  Y se informa que no hay integrantes que puedan votar

Escenario: Una historia no puede tener dos rondas en curso
  Dado un proyecto con integrantes, las historias 1 y 2, y una ronda sin revelar sobre la historia 1
  Cuando intento abrir otra ronda sobre la historia 1
  Entonces la operacion se rechaza
  Y se informa que ya hay una ronda en curso
  Y abrir una ronda sobre la historia 2 se acepta

Escenario: Votar registra quien voto sin dejar ver los votos, ni siquiera el propio
  Dado una ronda sin revelar con "Agustin Salinas", "Juliana Bustos" y "Maximiliano Eula"
  Y "Agustin Salinas" voto 3 y "Juliana Bustos" voto 8
  Cuando "Juliana Bustos" pide ver los votos de la ronda
  Entonces se informa que los votos estan ocultos hasta que se revele la ronda
  Y solo se sabe que votaron "Agustin Salinas" y "Juliana Bustos" y que falta "Maximiliano Eula"

Esquema del escenario: Revelar cuando votaron todos
  Dado una ronda sin revelar donde <votos>
  Cuando pido revelar la ronda
  Entonces la ronda queda revelada
  Y se ve cada voto con el nombre de quien lo dio, en el orden de los integrantes

  Ejemplos:
    | votos                                                                     |
    | "Agustin Salinas voto 3, Juliana Bustos voto 8 y Maximiliano Eula voto 5" |
    | "Maximiliano Eula, el unico integrante, voto 13"                          |

Esquema del escenario: Rechazar revelar mientras falten votos
  Dado una ronda sin revelar con tres participantes, donde votaron <votaron>
  Cuando pido revelar la ronda
  Entonces la operacion se rechaza
  Y se informa que todavia faltan votos
  Y la ronda sigue sin revelar y con los votos ocultos

  Ejemplos:
    | votaron                            |
    | "nadie"                            |
    | "Agustin Salinas y Juliana Bustos" |

Esquema del escenario: Rechazar un voto fuera de la escala
  Dado una ronda sin revelar donde falta votar "Juliana Bustos"
  Cuando "Juliana Bustos" intenta votar <puntos>
  Entonces la operacion se rechaza
  Y se informa que los valores validos son 1, 2, 3, 5, 8, 13 y 21
  Y "Juliana Bustos" sigue entre los que faltan

  Ejemplos:
    | puntos |
    | 4      |
    | 7      |
    | 0      |
    | -3     |

Escenario: Rechazar un segundo voto del mismo integrante
  Dado una ronda sin revelar donde "Juliana Bustos" ya voto 13
  Cuando "Juliana Bustos" intenta votar 21
  Entonces la operacion se rechaza
  Y el aviso dice que ya voto en esta ronda, sin mostrar ningun valor
  Y al revelar la ronda el voto de "Juliana Bustos" es 13

Esquema del escenario: Rechazar el voto de quien no participa de la ronda
  Dado una ronda abierta con "Agustin Salinas", "Juliana Bustos" y "Maximiliano Eula"
  Y <situacion>
  Cuando <quien> intenta votar 5
  Entonces la operacion se rechaza
  Y se informa que no participa de esta ronda

  Ejemplos:
    | situacion                                                  | quien             |
    | "Pedro Gomez no es integrante del proyecto"                | "Pedro Gomez"     |
    | "Emiliano Sorato se agrego al proyecto despues de abrirla" | "Emiliano Sorato" |

Escenario: Quien sale del proyecto despues de abrir la ronda sigue participando
  Dado una ronda abierta con "Agustin Salinas", "Juliana Bustos" y "Maximiliano Eula"
  Cuando se quita a "Juliana Bustos" del proyecto
  Entonces "Juliana Bustos" sigue entre los que faltan votar
  Y la ronda no se puede revelar sin su voto

Escenario: Reconocer al integrante aunque el nombre venga escrito distinto
  Dado una ronda sin revelar donde falta votar "Juliana Bustos"
  Cuando "  juliana bustos  " vota 5
  Entonces el voto se acepta
  Y figura entre los que votaron "Juliana Bustos", escrito como esta cargado

Esquema del escenario: Una ronda revelada ya no acepta cambios
  Dado una ronda revelada donde todos votaron y "Juliana Bustos" voto 8
  Cuando intento <accion>
  Entonces la operacion se rechaza
  Y se informa que la ronda ya fue revelada
  Y los votos revelados siguen siendo los mismos

  Ejemplos:
    | accion                                 |
    | "votar 3 otra vez como Juliana Bustos" |
    | "revelarla de nuevo"                   |

Escenario: Mostrar una ronda sin revelar no deja ver los votos
  Dado una ronda sin revelar sobre la historia 1 donde "Agustin Salinas" voto 13 y "Juliana Bustos" voto 21
  Cuando se muestra la ronda
  Entonces se ve que es la ronda 1 de la historia 1 y que votaron 2 de 3
  Y no aparece ninguno de los valores votados

Esquema del escenario: Consultar la ronda actual de una historia
  Dado la historia 1 con <rondas>
  Cuando consulto su ronda actual
  Entonces <resultado>

  Ejemplos:
    | rondas                 | resultado                                                |
    | ninguna ronda          | se informa que no tiene ninguna ronda, y no es un error  |
    | la ronda 1 sin revelar | recibo la ronda 1, sin revelar                           |
    | la ronda 1 revelada    | recibo la ronda 1, revelada                              |

Escenario: Consultar la ronda actual de una historia que no existe
  Dado un proyecto sin historia 99
  Cuando consulto la ronda actual de la historia 99
  Entonces se informa que no se encontro la historia
```

**Para decidir en la puerta:**

1. **Los participantes se fijan al abrir (RN-2)**: votan los que se sentaron a la mesa al
   empezar. Si alguien sale del proyecto con la ronda abierta, la ronda espera su voto y no se
   puede revelar; como acá no hay forma de cancelarla, la salida es estimar directo con la
   #12. La alternativa es tomar siempre los integrantes del momento: el que llega tiene que
   votar antes de revelar y el voto del que sale desaparece.
   **Decidido el 21/10: se acepta como está, con su consecuencia escrita.** Como la ronda no
   se puede cancelar y la RN-7 no deja abrir otra mientras haya una sin revelar, **esa
   historia ya no se puede votar con Planning Poker**: se estima directo con la #12. Que
   alguien salga del proyecto en medio de una ronda es raro, y cancelar rondas queda afuera
   de esta historia.
2. **Un solo integrante puede votar solo.** "Nadie estima solo" ([CLAUDE.md](../../../CLAUDE.md))
   es una regla del equipo, no del sistema, y un proyecto de una persona no podría usar el
   Planning Poker. Si se quiere un mínimo de dos, cambia la RN-6.
   **Decidido el 21/10: se acepta como está.**
3. **El voto no se cambia (RN-10)**, porque la issue dice "uno solo por ronda". Si el equipo
   quiere poder cambiarlo antes de revelar, el segundo reemplaza al primero y el secreto no se
   afecta.
   **Decidido el 21/10: el voto no se cambia.**
4. **Revelar es un pedido explícito (RN-14)**, no pasa solo con el último voto: como dice la
   guía, las cartas se dan vuelta juntas cuando alguien lo pide. Y **revelar dos veces es un
   error (RN-16)**, con el criterio que la #4 y la #9 aplicaron a los repetidos.
   **Decidido el 21/10: se acepta como está.**
5. **Sin "?" ni "café"**: la issue pide un valor válido de Fibonacci. Agregarlas obliga a
   decidir cómo se guardan y a dejarlas fuera del más bajo y el más alto de la #73.
   **Decidido el 21/10: sin "?" ni "café".**
6. **Una sola ronda por historia hasta que esté la #73.** Mientras tanto, pedir una segunda
   ronda se rechaza con el aviso de la RN-7 aunque la primera ya esté revelada. La RN-7 sigue
   valiendo después de la #73 para las rondas sin revelar.
   **Decidido el 21/10: se acepta como está.** La #73 también es de Maximiliano y la afloja.
7. **Lo que le faltaba al contrato.** El borrador marcaba que
   [`docs/diseno-dominio.md`](../../diseno-dominio.md) no tenía los participantes de la ronda
   ni nombraba las consultas de quién votó, quién falta y los votos revelados.
   **Resuelto en el contrato (PR #94, 20/10):** ya tiene `Participants`, `r.Voted()`,
   `r.Pending()`, `r.Votes()` y `p.CurrentVotingRound(storyID)`. No hace falta tocarlo antes
   de implementar.
8. **Agregados del contrato del 20/10 (PR #94), para que los revise Maximiliano en la
   puerta.** Los pidió la web (#78) y los escribió el Agile Enabler al ajustar el contrato;
   no son decisión de esta spec:
   - **La ronda actual de una historia (RN-19)**: entre un pedido y otro, la web tiene que
     volver a encontrar la ronda que está votando, y la lista de rondas de una historia es de
     la #73. Devuelve la última ronda, revelada o no. Sin ninguna ronda se informa y no es un
     error, igual que el sprint abierto de la #8.
     **Decidido el 21/10: se acepta.**
   - **Las rondas entran en lo que se guarda**, como toda entidad nueva desde la #77 (tarea
     14). La consecuencia: los votos de una ronda sin revelar quedan escritos en el archivo de
     datos, y quien lo abra con un editor los ve. La alternativa —no guardar las rondas
     abiertas— las pierde al cerrar el sistema. **Lo decide esta puerta**; la web (#78) no
     cambia en ningún caso.
     **Decidido el 21/10: las rondas sin revelar se guardan con sus votos.** El secreto de la
     RN-12 vale **dentro del sistema**: ninguna pantalla, consulta ni aviso muestra un voto
     antes de revelar. Quien abre el archivo de datos con un editor ve los votos, igual que
     puede ver o cambiar cualquier otro dato: el archivo no es una interfaz del sistema. No
     guardarlas haría perder una ronda a medio votar si se cierra la web, y cifrar los votos
     sería una protección de mentira, porque la clave estaría en el mismo programa.

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: Agustín Salinas — Fecha: 2026-10-21 — Comentarios: aprobada en el PR #115 (https://github.com/emisorato1/seguimiento-medicion-ics/pull/115#issuecomment-6328013297), junto con las otras dos puertas. Nota que no bloquea: en la demo no hay que quitar un integrante con una ronda a medio votar (punto 1).

## Fase 2 — Diseño

Las decisiones transversales del proyecto —un solo paquete `internal/domain`, errores como
valores comparables con `errors.Is`, mensajes en castellano, validar todo antes de cambiar
nada y sumar cada entidad nueva a lo que se guarda— están en [AGENTS.md](../../../AGENTS.md) y
no se repiten acá. Los datos y las operaciones son los del
[contrato del dominio](../../diseno-dominio.md): `VotingRound` con `StoryID`, `Number`,
`Participants`, `Votes` y `Revealed`; `OpenVotingRound` y `CurrentVotingRound` sobre el
proyecto, y `Vote`, `Reveal`, `Voted`, `Pending` y `Votes` sobre la ronda. `Agreed`,
`Agree`, `Differences` y `VotingRounds` son de la #73.

### Enfoque

La ronda es una entidad nueva que **cuelga del proyecto**, como los integrantes y las
historias: el proyecto guarda todas sus rondas en el orden en que se abrieron, y cada ronda
sabe de qué historia es por su identificador. Abrir se le pide al proyecto, porque hacen falta
la historia y los integrantes. Votar y revelar se le piden a la ronda, que ya tiene lo que
necesita: los participantes, que copió al abrirse (RN-2), y los votos.

Abrir controla, en este orden, que la historia exista, que no esté Terminada, que el proyecto
tenga integrantes y que la historia no tenga ya una ronda; recién entonces la crea.

#### Cómo se garantiza el secreto

Es la decisión que define la historia. Que los datos de la ronda estén ocultos fuera del
paquete, como los de toda entidad, es necesario pero **no alcanza**: después de revelar los
votos se tienen que poder leer, así que existe una consulta que los devuelve y es ahí donde
vive la regla; y las funciones de impresión de Go muestran también los datos ocultos de una
entidad, así que un simple "imprimir la ronda" desde la consola los mostraría. El secreto se
sostiene con cuatro cosas, y cada una tiene su test:

1. **Los valores salen por una sola consulta**, la de los votos, que mientras la ronda no está
   revelada responde con un error (RN-12). No hay consulta del voto de un integrante.
2. **Quién votó y quién falta son consultas aparte** que devuelven solo nombres (RN-13).
3. **Votar solo responde si se aceptó**, y ningún aviso lleva un valor (RN-10).
4. **La ronda define cómo se imprime, en todas las formas de imprimir**: número, historia y
   cuántos votaron, sin valores (RN-17). No alcanza con definir el texto de la ronda: Go tiene
   una forma de imprimir pensada para depurar que lo saltea y muestra los datos ocultos tal
   cual. La ronda controla también esa, y el test prueba cada una.

Lo que no se protege, y conviene tenerlo claro para la defensa: **dentro** de
`internal/domain` cualquier código puede leer los datos ocultos, como con toda entidad. El
"desde afuera" de la regla es desde fuera del paquete, donde viven la consola y cualquier
otra interfaz. Adentro lo cuidan la revisión y la #73, que lee los votos por la misma
consulta protegida.

### Archivos afectados

| Archivo | Se crea o se modifica | Para qué |
|---|---|---|
| `internal/domain/voting_round.go` | se crea | la ronda; abrirla, votar y revelar; las consultas de quién votó, quién falta y los votos; la ronda actual de una historia; cómo se imprime |
| `internal/domain/registry_json.go` y su test | se modifican | las rondas entran en la conversión del registro, con su caso en el test de ida y vuelta (regla de la #77; punto 8 de la puerta) |
| `internal/domain/voting_round_test.go` | se crea | los tests de los escenarios BDD de la Fase 1, nombrados como cada escenario |
| `internal/domain/project.go` | se modifica | la lista de rondas dentro del proyecto, al lado de los integrantes y el backlog |
| `internal/domain/errors.go` | se modifica | los siete errores nuevos: sin integrantes, ronda en curso, no participa, ya votó, votos ocultos, faltan votos y ronda revelada |

### Decisiones

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| Los valores solo se leen por la consulta de los votos, que sin revelar responde con un error | que sin revelar devuelva una lista vacía, o los votos con el valor tapado | la lista vacía se confunde con "nadie votó" y una interfaz distraída la mostraría así; el valor tapado obliga a un valor especial que se puede tomar por un voto |
| Quién votó y quién falta son consultas propias, que devuelven nombres | que la consulta de los votos, sin revelar, traiga los nombres sin el valor | ninguna consulta queda "a medias": la que da valores, o los da todos o rechaza |
| Una sola entidad de ronda, con su indicador de revelada | un tipo para la ronda sin revelar, sin forma de leer votos, y otro para la revelada | daría la garantía al compilar, pero el contrato define una sola `VotingRound` con `Revealed`, y la #73 guarda rondas en los dos estados en la misma lista |
| La ronda define cómo se imprime, sin valores, **en todas las formas de imprimir** | confiar en que los datos ocultos alcanzan, o definir solo el texto de la ronda | la impresión de Go recorre los datos ocultos: un print de depuración en la consola revelaría los votos sin que se rompa ninguna regla a la vista. Y definir solo el texto deja abierta la forma de imprimir para depurar, que no lo usa |
| Al cargar las rondas guardadas se validan con las mismas reglas que al votar | confiar en el archivo | un archivo editado a mano podría traer un voto fuera de la escala o de alguien que no participa, y el dominio terminaría con datos que nunca habría aceptado. Es lo que hace la #77 con las demás entidades |
| Se vota con el nombre y apellido, buscado como en la #4, y el voto guarda el nombre como está cargado | recibir el integrante ya buscado | es lo que hace la #14 con el esfuerzo: la consola tiene el texto, y fuera del dominio no hay una búsqueda de integrante por nombre |
| Orden de los controles al votar: ronda revelada → participa → ya votó → valor | validar primero el valor | si la ronda ya se reveló nada más importa, y no tiene sentido validar el valor de quien no puede votar; es el criterio de la #9 y la #12 |
| Se reutilizan la escala y los errores de la #12 (valor inválido, historia terminada) y el de historia inexistente de la #6 | escala y errores propios del Planning Poker | mismo caso y mismo mensaje; con dos escalas, cambiar una deja la otra mintiendo |
| En esta historia la apertura rechaza toda segunda ronda sobre una historia; la #73 la afloja | numerar desde ya contando las rondas anteriores | sería implementar sin test una regla de la #73 y dejarle una tarea que pasa en verde (punto 6 de la puerta) |
| El proyecto guarda las rondas en una sola lista, en orden de apertura | que cada historia guarde sus rondas | no se toca `story.go`, que es de la #6, y la ronda ya sabe de qué historia es |

### Dependencias y restricciones

- **Depende de la #4, la #6 y la #12**, que ya están en `dev`. No hay que esperar a nadie.
- **No se tocan `story.go`, `story_estimation.go` ni `member.go`**: la historia y los
  integrantes solo se leen, y la escala se usa tal como está. `errors.go` es compartido.
- **Lo que le deja a la #73**: la consulta de los votos es la única puerta a los valores, y la
  #73 la usa; la apertura rechaza toda segunda ronda, y la #73 la afloja sin tocar el test de
  la ronda en curso, que tiene que seguir pasando.

> **Puerta 2** — ¿el diseño es viable y coherente con el resto del proyecto?
> Aprobó: Agustín Salinas — Fecha: 2026-10-21 — Comentarios: aprobada en el PR #115 (https://github.com/emisorato1/seguimiento-medicion-ics/pull/115#issuecomment-6328013297), junto con las otras dos puertas. Nota que no bloquea: en la demo no hay que quitar un integrante con una ronda a medio votar (punto 1).

## Fase 3 — Tareas

Cada tarea es un ciclo de TDD completo: el test que falla en un commit, la implementación que
lo hace pasar en otro. Van en orden, cada una deja los tests en verde y **cada una tiene al
menos un test que falla antes de implementarla**: los escenarios que salen del mismo código
van en la misma tarea, que es la lección de la Fase 4 de la #12. Siguen la convención de
[AGENTS.md](../../../AGENTS.md): un test por escenario, con su nombre, y los errores
verificados con `errors.Is`. Cada rechazo comprueba además que la ronda quedó como estaba.

| # | Tarea | Escenario BDD | Reglas | Estado |
|---|---|---|---|---|
| 1 | Abrir una ronda sobre una historia no terminada: queda como la ronda 1, sin revelar, sin votos y con todos los integrantes entre los que faltan; la historia no cambia | Abrir una ronda sobre una historia que no esta terminada | RN-1, RN-2, RN-3, RN-13 | pendiente |
| 2 | Rechazar la apertura sobre una historia Terminada o inexistente, con los errores que ya existen | Rechazar una ronda sobre una historia que no se puede votar | RN-4, RN-5 | pendiente |
| 3 | Rechazar la apertura en un proyecto sin integrantes | Rechazar una ronda en un proyecto sin integrantes | RN-6 | pendiente |
| 4 | Rechazar una segunda ronda sobre una historia con una ronda sin revelar; otra historia abre la suya | Una historia no puede tener dos rondas en curso | RN-7 | pendiente |
| 5 | Votar pasa al integrante de los que faltan a los que votaron, y los votos no se pueden leer mientras no se revele | Votar registra quien voto sin dejar ver los votos, ni siquiera el propio | RN-8, RN-12, RN-13 | pendiente |
| 6 | Revelar con todos los votos deja ver cada voto con su nombre, en el orden de los integrantes, también con un solo integrante | Revelar cuando votaron todos | RN-14 | pendiente |
| 7 | Rechazar revelar mientras falte algún voto; la ronda sigue sin revelar | Rechazar revelar mientras falten votos | RN-15, RN-18 | pendiente |
| 8 | Rechazar un voto fuera de la escala con el error de la #12 | Rechazar un voto fuera de la escala | RN-9, RN-18 | pendiente |
| 9 | Rechazar el segundo voto de un integrante, antes que el control del valor y sin mostrar el primero, que es el que queda | Rechazar un segundo voto del mismo integrante | RN-10, RN-12, RN-18 | pendiente |
| 10 | Rechazar el voto de quien no participa, antes que el control del valor; reconocer el nombre escrito distinto y guardarlo como está cargado; quien sale del proyecto sigue participando | Rechazar el voto de quien no participa de la ronda / Reconocer al integrante aunque el nombre venga escrito distinto / Quien sale del proyecto despues de abrir la ronda sigue participando | RN-2, RN-11, RN-18 | pendiente |
| 11 | Rechazar votar o revelar de nuevo en una ronda revelada, antes que cualquier otro control | Una ronda revelada ya no acepta cambios | RN-16, RN-18 | pendiente |
| 12 | Imprimir una ronda sin revelar muestra número, historia y cuántos votaron, sin ningún valor, en todas las formas de imprimir de Go (`%v`, `%+v`, `%#v` y `%s`) | Mostrar una ronda sin revelar no deja ver los votos | RN-17 | pendiente |
| 13 | La ronda actual de una historia: la última, revelada o no; sin rondas se informa y no es un error; historia inexistente | Consultar la ronda actual de una historia / Consultar la ronda actual de una historia que no existe | RN-19 | pendiente |
| 14 | Las rondas entran en la conversión del registro: una sin revelar y una revelada, con sus participantes y sus votos, se reconstruyen iguales. Una ronda inválida en el archivo rechaza la carga: historia que no existe, voto fuera de la escala, voto de quien no participa o dos votos de la misma persona, dos rondas sin revelar sobre la misma historia, o una ronda revelada a la que le faltan votos | — (regla de la #77: Los datos siguen al cerrar y volver a abrir) | — | pendiente |

**Cómo se verifica que no se filtra ningún valor.** Los tests de las tareas 9 y 12 votan 13 y
21, valores que no aparecen en ningún otro dato de la ronda (historia 1, ronda 1, nombres sin
números), y comprueban que ni el aviso de voto repetido ni la ronda impresa los contengan. No
se compara el texto completo, que puede cambiar: se verifica lo que la RN-12 prohíbe.

### Por qué este orden

- **La 1 va primero** porque trae las consultas de quién votó y quién falta, con las que
  verifican todas las demás. De la 2 a la 4, los rechazos de la apertura de a uno.
- **La 5 trae el voto y, en el mismo test, el secreto**: la consulta de los votos nace
  rechazando, porque revelar todavía no existe, y la regla que define la historia queda
  protegida desde el primer voto. La 6 y la 7 son revelar y su rechazo.
- **De la 8 a la 11 los rechazos del voto se agregan del último al primero** del orden de la
  Fase 2: cada control nuevo va delante de los que ya existen, y su test incluye un caso donde
  también fallaría el anterior. Así el orden queda probado sin una tarea propia que pasaría
  en verde.
- **La 12 cierra la última vía por la que se podía ver un voto.**
- **La 13 y la 14 son los agregados del contrato del 20/10** (punto 8 de la puerta): no
  cambian nada de lo anterior, y la 14 va última porque guarda todo lo que la ronda ya tiene.

### Cobertura

Las diecinueve reglas quedan cubiertas: RN-1 (1), RN-2 (1, 10), RN-3 (1), RN-4 (2), RN-5 (2),
RN-6 (3), RN-7 (4), RN-8 (5), RN-9 (8), RN-10 (9), RN-11 (10), RN-12 (5, 9), RN-13 (1, 5),
RN-14 (6), RN-15 (7), RN-16 (11), RN-17 (12), RN-18 (7, 8, 9, 10, 11), RN-19 (13).

| Criterio de aceptación | Tareas |
|---|---|
| Se puede abrir una ronda sobre una historia no Terminada, y queda como la ronda 1 | 1 |
| No se puede abrir una ronda sobre una historia Terminada ni inexistente | 2 |
| No se puede abrir una ronda en un proyecto sin integrantes | 3 |
| No se puede abrir otra ronda sobre una historia con una sin revelar | 4 |
| Cada participante vota una sola vez, y el segundo voto no muestra el primero | 5, 9 |
| El voto tiene que ser un valor de la escala | 8 |
| Solo votan los integrantes del momento de abrir, y el nombre se reconoce escrito distinto | 10 |
| Sin revelar nadie ve ningún voto; solo se sabe quién votó y quién falta | 5 |
| Mostrar una ronda sin revelar no deja ver ningún valor | 12 |
| La ronda se revela recién cuando votaron todos | 7 |
| Al revelarse se ven todos los votos con su nombre | 6 |
| No se puede votar ni volver a revelar una ronda revelada | 11 |
| Un rechazo deja la ronda como estaba | 7, 8, 9, 10, 11 |
| Se consulta la ronda actual; sin ninguna se informa y no es error | 13 |

Los dieciséis escenarios de la Fase 1 tienen su tarea, en la columna "Escenario BDD". La 14
no tiene escenario propio: cumple la regla de la #77 de sumar cada entidad nueva a lo que se
guarda.

### Dependencias

- **Ninguna tarea espera a otra historia**: la #4, la #6 y la #12 ya están en `dev`. La
  tarea 2 necesita una historia Terminada, que es de la #10: el test fija el estado desde
  dentro del paquete, como hizo la #12.
- **El contrato ya tiene los participantes y los nombres de las consultas** (punto 7 de la
  puerta), así que la tarea 1 no espera nada.

> **Puerta 3** — ¿las tareas cubren todos los criterios de aceptación?
> Aprobó: Agustín Salinas — Fecha: 2026-10-21 — Comentarios: aprobada en el PR #115 (https://github.com/emisorato1/seguimiento-medicion-ics/pull/115#issuecomment-6328013297), junto con las otras dos puertas. Nota que no bloquea: en la demo no hay que quitar un integrante con una ronda a medio votar (punto 1).

## Fase 4 — Implementación

_Se completa durante la implementación._

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #13 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | `internal/domain/voting_round_test.go` |
| Código Go | `internal/domain/voting_round.go`, la lista de rondas en `internal/domain/project.go` y sus errores en `internal/domain/errors.go` |
