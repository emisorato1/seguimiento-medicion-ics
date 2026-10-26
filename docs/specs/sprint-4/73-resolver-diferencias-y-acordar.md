# Spec: Resolver diferencias y acordar la estimación

- **Historia:** #73
- **Sprint:** 4
- **Responsable:** Maximiliano Eula
- **Estado:** Implementada — **en review**. Adelantada al Sprint 3 con la regla del 21/10 de CLAUDE.md
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de la consigna y
  de los criterios de aceptación de la issue. El responsable lo revisa, lo corrige si hace
  falta y lo presenta a las puertas. Revisado por el responsable el 22/10: decidió los
  puntos de la puerta, dio por resuelto el 7 con el contrato, sumó la tarea de guardar el
  acuerdo en el archivo y actualizó las dependencias con la #13 ya escrita (PR #141).
- **Depende de:** #13 (la ronda, el voto y el revelado, en el PR #141) y #12 (la operación de
  estimar con la que se graba el acuerdo, ya en `dev`). El código empieza cuando la #13 esté
  en `dev` (ver las dependencias de la Fase 3)
- **La necesitan:** nadie todavía: cierra la cadena del Planning Poker

## Fase 1 — Requisitos

### Objetivo

Que, después de dar vuelta los votos, el equipo vea si hay diferencias y quiénes tienen que
explicar su número —los que votaron más bajo y más alto—, pueda volver a votar en una ronda
nueva y, cuando se pone de acuerdo, deje la historia estimada con ese valor. Son los pasos 5 a
8 de la [guía del equipo](../../guia-planning-poker.md).

### Qué queda afuera

- **Abrir la primera ronda, votar, ocultar y revelar**: es la #13.
- **Calcular el promedio o proponer un valor**: el sistema muestra las diferencias y el valor
  lo decide el equipo, que es lo que pide la issue ("en vez de quedarme con el promedio").
- **Limitar la cantidad de rondas** (punto 4 de "Para decidir en la puerta").
- **Cambiar o quitar un acuerdo** ya registrado, y borrar rondas. Para cambiar la estimación
  se abre otra ronda o se reestima con la #12.
- **Las cartas "?" y "café"**, que quedaron afuera en la #13.

### Entradas

| Dato | Obligatorio | Descripción |
|---|---|---|
| Ronda | sí | sobre qué ronda se piden las diferencias o se acuerda |
| Historia | sí, al abrir una ronda nueva o consultar sus rondas | por su identificador |
| Estimación acordada | sí, al acordar | el valor que decidió el equipo |

### Salidas

De una ronda revelada: si hay diferencias, y el voto más bajo y el más alto con quiénes los
dieron. Una ronda nueva sobre la misma historia, con el número siguiente, y la lista de rondas
de la historia. Al acordar, la ronda con su estimación acordada y la historia estimada con
ese valor.

### Reglas de negocio

1. **[RN-1]** Cuando se piden las diferencias de una ronda revelada, el sistema debe informar
   si hay diferencias, el voto más bajo y el voto más alto. Hay diferencias cuando al menos
   dos votos tienen valores distintos.
2. **[RN-2]** El sistema debe informar, junto al voto más bajo y al más alto, el nombre de
   todos los integrantes que lo dieron, en el orden de los participantes de la ronda.
3. **[RN-3]** Si todos los votos tienen el mismo valor, el sistema debe informar que no hay
   diferencias, con el más bajo y el más alto iguales a ese valor. Vale también para una ronda
   de un solo integrante.
4. **[RN-4]** Mientras la ronda no esté revelada, el sistema debe rechazar el pedido de
   diferencias con el aviso de votos ocultos de la #13 (su RN-12): no informa ni siquiera si
   hay diferencias.
5. **[RN-5]** Cuando se abre una ronda sobre una historia cuya última ronda ya está revelada,
   el sistema debe aceptarla y darle el número que sigue al de esa última ronda. Mientras la
   última siga sin revelar, vale la RN-7 de la #13.
6. **[RN-6]** El sistema debe numerar las rondas por historia: la primera ronda de cada
   historia es la 1, aunque otras historias del proyecto ya tengan rondas.
7. **[RN-7]** Cuando se abre una ronda nueva, el sistema debe conservar las anteriores tal
   como estaban, con sus votos y su estimación acordada si la tenían. La nueva arranca como
   cualquier ronda de la #13: sin votos, sin revelar y con los integrantes de ese momento.
8. **[RN-8]** El sistema debe permitir consultar las rondas de una historia, en orden de
   número. Una historia sin rondas devuelve una lista vacía, que no es un error; una que no
   existe informa que no se encontró la historia.
9. **[RN-9]** Cuando se acuerda una estimación sobre una ronda revelada, el sistema debe
   registrarla en la ronda y dejar la historia estimada con ese valor, reemplazando el que
   tuviera, igual que al estimar con la #12.
10. **[RN-10]** El sistema debe aceptar como estimación acordada cualquier valor de la escala
    de la #12, lo haya votado alguien o no.
11. **[RN-11]** Si el valor acordado no está en la escala, o si la historia está Terminada, el
    sistema debe rechazar el acuerdo con los mismos avisos de la #12 (sus RN-2, RN-3 y RN-8).
12. **[RN-12]** Si la ronda todavía no se reveló, el sistema debe rechazar el acuerdo e
    informar que primero hay que revelarla.
13. **[RN-13]** Si la ronda ya tiene una estimación acordada, el sistema debe rechazar otro
    acuerdo e informar que esa ronda ya se acordó.
14. **[RN-14]** Si la historia tiene una ronda posterior a la que se quiere acordar, el sistema
    debe rechazar el acuerdo e informar que solo se acuerda sobre la última ronda.
15. **[RN-15]** Mientras una ronda no tenga acuerdo, el sistema debe informarla como sin
    estimación acordada, y eso no es un error.
16. **[RN-16]** Si un acuerdo se rechaza, el sistema debe dejar la ronda y la historia como
    estaban.

### Casos límite

- Ronda de un solo integrante, o todos votaron lo mismo: no hay diferencias, y el más bajo y
  el más alto son ese voto.
- Empate en el más bajo o en el más alto: se nombra a todos los que lo dieron.
- Acordar un valor que nadie votó, por ejemplo 5 con votos de 3 y 8: válido.
- Acordar sobre una historia que ya estaba estimada: queda el valor acordado.
- Historia que pasó a Terminada entre el revelado y el acuerdo: el acuerdo se rechaza.
- Abrir una ronda nueva después de acordar: válido, es reestimar; la historia conserva su
  estimación hasta el próximo acuerdo.
- Historia sin rondas: la consulta devuelve una lista vacía.

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| Pedir las diferencias de una ronda sin revelar | rechaza e informa que los votos están ocultos hasta revelar |
| Acordar sobre una ronda sin revelar | rechaza e informa que primero hay que revelarla |
| Acordar dos veces la misma ronda | rechaza e informa que esa ronda ya se acordó |
| Acordar sobre una ronda que no es la última de su historia | rechaza e informa que solo se acuerda sobre la última |
| Valor acordado fuera de la escala | rechaza e informa los valores válidos |
| La historia está Terminada al acordar | rechaza e informa que su estimación no se modifica |
| Abrir una ronda nueva con la última sin revelar | rechaza e informa que ya hay una ronda en curso (#13) |
| Consultar las rondas de una historia que no existe | informa que no se encontró la historia |

### Criterios de aceptación

- [ ] Con la ronda revelada, el sistema informa si hay diferencias entre los votos.
- [ ] Informa el voto más bajo y el más alto, con el nombre de todos los que dieron cada uno.
- [ ] Si todos votaron lo mismo, informa que no hay diferencias.
- [ ] Con la ronda sin revelar no informa nada de los votos, ni siquiera si hay diferencias.
- [ ] Se puede abrir una ronda nueva sobre la misma historia cuando la última está revelada,
      y las anteriores quedan guardadas como estaban.
- [ ] Las rondas de cada historia se numeran en orden, empezando en 1, y se consultan en ese
      orden.
- [ ] Se puede registrar la estimación acordada, y la historia queda estimada con ese valor.
- [ ] La estimación acordada tiene que ser un valor de la escala, aunque nadie lo haya votado.
- [ ] No se puede acordar sobre una ronda sin revelar, ya acordada o que no es la última de
      su historia.
- [ ] No se puede acordar sobre una historia Terminada.
- [ ] Un acuerdo rechazado no cambia ni la ronda ni la historia.

### Escenarios BDD

```gherkin
Esquema del escenario: Informar las diferencias de una ronda revelada
  Dado una ronda revelada donde <votos>
  Cuando pido las diferencias de la ronda
  Entonces se informa que hay diferencias
  Y el voto mas bajo es <bajo>
  Y el voto mas alto es <alto>

  Ejemplos:
    | votos                                                                     | bajo                      | alto                                         |
    | "Agustin Salinas voto 3, Juliana Bustos voto 8 y Maximiliano Eula voto 5" | "3, de Agustin Salinas"   | "8, de Juliana Bustos"                       |
    | "Agustin Salinas voto 3, Juliana Bustos voto 8 y Maximiliano Eula voto 8" | "3, de Agustin Salinas"   | "8, de Juliana Bustos y Maximiliano Eula"    |

Esquema del escenario: Sin diferencias cuando todos votaron lo mismo
  Dado una ronda revelada donde <votos>
  Cuando pido las diferencias de la ronda
  Entonces se informa que no hay diferencias
  Y el voto mas bajo y el mas alto son <valor>, de <quienes>

  Ejemplos:
    | votos                                            | valor | quienes                |
    | "los tres integrantes votaron 5"                 | 5     | "los tres integrantes" |
    | "Maximiliano Eula, el unico integrante, voto 13" | 13    | "Maximiliano Eula"     |

Escenario: No se informan diferencias de una ronda sin revelar
  Dado una ronda sin revelar donde votaron todos
  Cuando pido las diferencias de la ronda
  Entonces se informa que los votos estan ocultos hasta que se revele la ronda
  Y no se informa si hay diferencias

Esquema del escenario: Acordar la estimacion de una ronda revelada
  Dado la historia 1 <antes>, con una ronda revelada donde se voto 3 y 8
  Cuando acuerdo la estimacion en <acordado>
  Entonces la ronda queda con la estimacion acordada <acordado>
  Y la historia 1 queda estimada en <acordado>

  Ejemplos:
    | antes           | acordado | por que                         |
    | "sin estimar"   | 8        | el valor mas alto               |
    | "estimada en 3" | 8        | reemplaza la estimacion anterior |
    | "sin estimar"   | 5        | un valor que nadie voto         |

Escenario: Una ronda revelada sin acuerdo figura como sin estimacion acordada
  Dado una ronda revelada sobre la que todavia no se acordo nada
  Cuando consulto su estimacion acordada
  Entonces figura como "sin acordar"
  Y eso no es un error

Esquema del escenario: Rechazar un acuerdo que la estimacion no permite
  Dado una ronda revelada sobre la historia 1, que esta <historia>
  Cuando intento acordar la estimacion en <puntos>
  Entonces la operacion se rechaza
  Y se informa que <aviso>
  Y la ronda sigue sin acordar
  Y la historia 1 sigue <historia>

  Ejemplos:
    | historia                    | puntos | aviso                                                    |
    | "Pendiente y sin estimar"   | 4      | "los valores validos son 1, 2, 3, 5, 8, 13 y 21"         |
    | "Pendiente y sin estimar"   | 0      | "los valores validos son 1, 2, 3, 5, 8, 13 y 21"         |
    | "Terminada y estimada en 3" | 5      | "la estimacion de una historia terminada no se modifica" |

Escenario: Rechazar un acuerdo sobre una ronda sin revelar
  Dado la historia 1 sin estimar, con una ronda sin revelar donde votaron todos
  Cuando intento acordar la estimacion en 5
  Entonces la operacion se rechaza
  Y se informa que primero hay que revelar la ronda
  Y la historia 1 sigue sin estimar

Escenario: Rechazar un segundo acuerdo sobre la misma ronda
  Dado una ronda revelada sobre la historia 1 con la estimacion acordada en 5
  Cuando intento acordar la estimacion en 8
  Entonces la operacion se rechaza
  Y se informa que esa ronda ya tiene una estimacion acordada
  Y la ronda sigue acordada en 5 y la historia 1 sigue estimada en 5

Escenario: Estimar una historia en dos rondas
  Dado un proyecto con "Agustin Salinas" y "Juliana Bustos", y la historia 1 sin estimar
  Y la ronda 1 de la historia 1, donde votaron 3 y 8, revelada con diferencias
  Y la ronda 2 de la historia 1, donde los dos votaron 5, revelada
  Cuando acuerdo la estimacion de la ronda 2 en 5
  Entonces la historia 1 queda estimada en 5
  Y las rondas de la historia 1 son la 1 y la 2, en ese orden
  Y la ronda 1 sigue con sus votos 3 y 8 y sin acordar

Escenario: Cada historia numera sus rondas desde 1
  Dado un proyecto donde la historia 1 ya tuvo las rondas 1 y 2
  Cuando abro una ronda sobre la historia 2
  Entonces la ronda queda registrada con numero 1

Escenario: Abrir una ronda nueva despues de acordar
  Dado la historia 1 estimada en 5 por el acuerdo de su ronda 1
  Cuando abro una ronda nueva sobre la historia 1
  Entonces la ronda nueva tiene el numero 2
  Y la historia 1 sigue estimada en 5
  Y la ronda 1 sigue acordada en 5

Esquema del escenario: Consultar las rondas de una historia
  Dado un proyecto con la historia 2 sin rondas
  Cuando consulto las rondas de la historia <id>
  Entonces <resultado>

  Ejemplos:
    | id | resultado                                      |
    | 2  | "recibo una lista vacia, y eso no es un error" |
    | 99 | "se informa que no se encontro la historia"    |

Escenario: Solo se acuerda sobre la ultima ronda de la historia
  Dado la historia 1 sin estimar, con la ronda 1 revelada sin acuerdo y la ronda 2 abierta
  Cuando intento acordar la estimacion de la ronda 1 en 5
  Entonces la operacion se rechaza
  Y se informa que solo se puede acordar sobre la ultima ronda de la historia
  Y la historia 1 sigue sin estimar
```

**Para decidir en la puerta:**

1. **Con empate se nombra a todos (RN-2).** La guía dice que hablan "el que puso el número
   más alto y el que puso el más bajo", en singular; con un empate no hay forma de elegir uno
   sin un criterio arbitrario. Si el equipo quiere uno solo, hay que decidir cuál.
   **Decidido el 22/10: se nombra a todos.**
2. **El acuerdo puede ser cualquier valor de la escala (RN-10)**, aunque nadie lo haya
   votado: con votos de 3 y 8 el equipo puede quedar en 5. Si se quiere que sea uno de los
   votados, es una regla más.
   **Decidido el 22/10: se acepta como está.**
3. **Se puede acordar aunque haya diferencias.** La guía misma cierra así: si después de tres
   rondas no hay acuerdo, "se toma el valor más alto". Exigir que no haya diferencias
   bloquearía esa salida.
   **Decidido el 22/10: se acepta como está.**
4. **No hay límite de rondas.** Las tres de la guía son una regla del equipo, no del sistema.
   Si se quiere que el sistema rechace la cuarta, es una regla más.
   **Decidido el 22/10: sin límite.**
5. **Solo se acuerda la última ronda, y una sola vez (RN-13 y RN-14).** Acordar una ronda que
   ya fue superada por otra dejaría la historia estimada mientras se sigue votando. Para
   cambiar un acuerdo se abre una ronda nueva.
   **Decidido el 22/10: se acepta como está.**
6. **El acuerdo de la ronda es un registro y no se sincroniza con la historia.** Si después
   se reestima o se le quita la estimación con la #12, la ronda conserva lo que se acordó ese
   día; lo que vale para planificar es la estimación de la historia.
   **Decidido el 22/10: se acepta como está.**
7. **Lo que le faltaba al contrato.** El borrador marcaba que
   [`docs/diseno-dominio.md`](../../diseno-dominio.md) no nombraba la consulta de las rondas
   de una historia ni la de las diferencias.
   **Resuelto en el contrato (PR #94, 20/10):** ya nombra la consulta de las diferencias, la
   de las rondas de una historia y el acuerdo. No hace falta tocarlo antes de implementar; los
   nombres están en la Fase 2.

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: Juliana Bustos — Fecha: 2026-10-23 — Comentarios: aprobada en el PR #142 (https://github.com/emisorato1/seguimiento-medicion-ics/pull/142#issuecomment-6357143382), junto con las otras dos puertas. Sus tres observaciones menores (la cita de la RN-3 en la RN-11, los nombres de funciones fuera de la Fase 1 y una sola acción en el Cuando del escenario de dos rondas) se aplicaron antes del código.

## Fase 2 — Diseño

Las decisiones transversales del proyecto están en [AGENTS.md](../../../AGENTS.md) y no se
repiten acá. Los nombres son los del [contrato del dominio](../../diseno-dominio.md): `Agreed`,
`Agree` y `Differences` sobre la ronda, y `VotingRounds` sobre el proyecto. Esta fase se apoya en la de la #13 —la ronda dentro del proyecto y
la consulta protegida de los votos— y en la de la #12, la operación de estimar.

### Enfoque

Esta historia **no crea ninguna entidad**: le agrega a la ronda de la #13 tres cosas
—comparar sus votos, aceptar una ronda nueva sobre la misma historia y registrar el acuerdo—
y le agrega al proyecto la consulta de las rondas de una historia.

**Las diferencias** se calculan cada vez que se piden, a partir de los votos leídos por la
consulta que protege la #13. Así la RN-4 no se programa de nuevo: si la ronda no está
revelada, esa consulta ya rechaza y el error sube tal cual. Las tres respuestas —si hay
diferencias, el más bajo y el más alto con quiénes— salen juntas, de los mismos votos.

**La ronda nueva** sale de aflojar la apertura de la #13: el proyecto ya no rechaza abrir si
la historia tiene rondas, sino solo si la última sigue sin revelar, y el número es el de la
última más uno.

**El acuerdo** tiene dos efectos —registrar el valor en la ronda y estimar la historia— y
tiene que ser todo o nada (RN-16). Primero van los controles de la ronda (revelada → sin
acuerdo → es la última); después la estimación con la operación de la #12, que hace sus
propios controles (Terminada → valor) y cambia la historia solo si pasan; y recién si salió
bien, la ronda registra el valor. Lo único que puede fallar va antes de lo único que no.

"Sin acuerdo" se representa con 0, por el mismo motivo que "sin estimar" en la #12: el 0
nunca es un acuerdo válido.

### Archivos afectados

| Archivo | Se crea o se modifica | Para qué |
|---|---|---|
| `internal/domain/voting_round_agreement.go` | se crea | las diferencias de una ronda, el acuerdo y la consulta de las rondas de una historia |
| `internal/domain/voting_round_agreement_test.go` | se crea | los tests de los escenarios BDD de la Fase 1, nombrados como cada escenario |
| `internal/domain/voting_round.go` | se modifica | que la ronda recuerde el proyecto que la abrió, y que la apertura acepte una ronda nueva cuando la última está revelada |
| `internal/domain/errors.go` | se modifica | los tres errores nuevos: ronda sin revelar al acordar, ronda ya acordada y ronda que no es la última |
| `internal/domain/registry_json.go` y su test | se modifican | el acuerdo de cada ronda entra en la conversión del registro, y al cargar cada ronda recupera su proyecto (regla de la #77) |

### Decisiones

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| El acuerdo estima la historia con la operación de la #12 | que la ronda valide la escala y escriba los story points | lo anticipó la Fase 2 de la #12: la escala y el control de historia terminada quedan en un solo lugar, y si la #12 cambia, el acuerdo la sigue sin tocarlo |
| Orden del acuerdo: controles de la ronda → estimar con la #12 → registrar en la ronda | registrar el acuerdo y después estimar | si la historia está Terminada la estimación falla, y con el otro orden quedaría una ronda acordada con la historia sin estimar |
| La ronda recuerda el proyecto que la abrió | pedirle el acuerdo al proyecto, o que la ronda guarde solo su historia | el contrato pone el acuerdo en la ronda (`r.Agree`), y solo el proyecto puede entregar la historia y decir si hay una ronda posterior (RN-14) |
| Las diferencias leen los votos por la consulta protegida de la #13 | leerlos directo, que dentro del paquete se puede | el secreto queda en un solo lugar; leyendo directo, la RN-4 habría que reprogramarla, y un olvido daría el más bajo y el más alto de una ronda sin revelar |
| El número de la ronda nueva es el de la última de la historia más uno | un contador por historia guardado aparte, como el de las historias de la #6 | las rondas no se borran, así que contar no repite números; si algún día se pueden borrar, se revisa |
| Lo nuevo va en `voting_round_agreement.go`; `voting_round.go` solo cambia en la apertura | agregarlo todo al archivo de la #13 | la trazabilidad queda directa por historia, como se aprobó en la #12 |
| El acuerdo se guarda en el archivo con su ronda, y al cargar se controla igual que al acordar: valor de la escala y solo sobre una ronda revelada; además, los números de las rondas de cada historia van seguidos desde 1, ninguna ronda sin revelar tiene otra después, y cada ronda recupera el proyecto que la abrió. Una ronda acordada **sí puede no ser la última**: después de acordar se puede abrir otra (corregido al implementar, ver los desvíos) | no guardarlo; guardarlo sin controlar | `AGENTS.md` pide que todo dato nuevo entre en la conversión, y sin el proyecto una ronda cargada no se podría acordar. Controlar al cargar es lo que hace la #77 con las demás entidades: un archivo editado a mano no mete datos que el dominio rechazaría |

### Dependencias y restricciones

- **Depende de la #13** (la ronda, su lista en el proyecto y la consulta protegida de los
  votos) y de la **#12** (estimar), que ya está en `dev`.
- **No se toca `story_estimation.go`**: el acuerdo usa la operación de estimar tal como está.
- **Ningún test de la #13 se modifica.** El que rechaza una segunda ronda con la primera sin
  revelar tiene que seguir pasando: es el que protege que la apertura no se afloje de más. El
  caso de ida y vuelta de `registry_json_test.go` se amplía con el acuerdo, sin cambiar lo
  que ya verifica.
- **`errors.go` es compartido**: los errores nuevos van al final.

> **Puerta 2** — ¿el diseño es viable y coherente con el resto del proyecto?
> Aprobó: Juliana Bustos — Fecha: 2026-10-23 — Comentarios: aprobada en el PR #142 (https://github.com/emisorato1/seguimiento-medicion-ics/pull/142#issuecomment-6357143382), junto con las otras dos puertas. Sus tres observaciones menores (la cita de la RN-3 en la RN-11, los nombres de funciones fuera de la Fase 1 y una sola acción en el Cuando del escenario de dos rondas) se aplicaron antes del código.

## Fase 3 — Tareas

Cada tarea es un ciclo de TDD completo: el test que falla en un commit, la implementación que
lo hace pasar en otro. Van en orden, cada una deja los tests en verde y **cada una tiene al
menos un test que falla antes de implementarla**: los escenarios que salen del mismo código
van en la misma tarea. Siguen la convención de [AGENTS.md](../../../AGENTS.md): un test por
escenario, con su nombre, y los errores verificados con `errors.Is`. Cada rechazo comprueba
además que la ronda y la historia quedaron como estaban (RN-16).

| # | Tarea | Escenario BDD | Reglas | Estado |
|---|---|---|---|---|
| 1 | Las diferencias de una ronda revelada: si hay, y el más bajo y el más alto con todos los que los dieron; sin diferencias cuando todos coinciden; sin revelar se rechaza con el aviso de votos ocultos | Informar las diferencias de una ronda revelada / Sin diferencias cuando todos votaron lo mismo / No se informan diferencias de una ronda sin revelar | RN-1, RN-2, RN-3, RN-4 | hecha |
| 2 | Acordar sobre una ronda revelada registra el valor en la ronda y estima la historia con la #12; un valor fuera de la escala o una historia Terminada se rechazan con los errores de la #12 y la ronda queda sin acuerdo | Acordar la estimacion de una ronda revelada / Una ronda revelada sin acuerdo figura como sin estimacion acordada / Rechazar un acuerdo que la estimacion no permite | RN-9, RN-10, RN-11, RN-15, RN-16 | hecha |
| 3 | Rechazar el acuerdo sobre una ronda sin revelar, antes de estimar la historia | Rechazar un acuerdo sobre una ronda sin revelar | RN-12, RN-16 | hecha |
| 4 | Rechazar un segundo acuerdo sobre la misma ronda; quedan el primero en la ronda y en la historia | Rechazar un segundo acuerdo sobre la misma ronda | RN-13, RN-16 | hecha |
| 5 | Abrir una ronda nueva cuando la última está revelada: lleva el número siguiente de su historia, las anteriores quedan como estaban y se consultan en orden, y la ronda que abre el proyecto ya se puede acordar | Estimar una historia en dos rondas / Cada historia numera sus rondas desde 1 / Abrir una ronda nueva despues de acordar / Consultar las rondas de una historia | RN-5, RN-6, RN-7, RN-8 | hecha |
| 6 | Rechazar el acuerdo sobre una ronda que no es la última de su historia, antes de estimarla | Solo se acuerda sobre la ultima ronda de la historia | RN-14, RN-16 | hecha |
| 7 | El acuerdo entra en la conversión del registro: una ronda acordada y una sin acordar se reconstruyen iguales, y una ronda cargada se puede acordar. Rechaza la carga un acuerdo fuera de la escala o sobre una ronda sin revelar, números de ronda que no van seguidos desde 1 en su historia y una ronda sin revelar seguida de otra. Una ronda acordada que ya no es la última se carga: es el estado del escenario "Abrir una ronda nueva despues de acordar" | — (regla de la #77: Los datos siguen al cerrar y volver a abrir) | — | hecha |

### Por qué este orden

- **La 1 va primero** porque es la más independiente: solo lee una ronda revelada y no cambia
  nada.
- **La 2 es el camino feliz del acuerdo y ya trae los rechazos de la #12**, porque salen de la
  misma llamada: separarlos daría tareas que pasan en verde, el error que dejó anotado la
  Fase 4 de la #12.
- **La 3 y la 4** agregan los controles propios de la ronda, delante de la llamada a la #12:
  sus tests verifican que la historia no se estimó, que es lo que prueba ese orden.
- **La 5 es la única que toca la apertura de la #13** y va después porque es la que necesita su
  código. La 6 necesita dos rondas de la misma historia, que recién existen con la 5.
- **La 7 cierra la historia**: guarda lo que las anteriores dejaron armado. No sale de un
  escenario de la Fase 1 sino de la regla de persistencia de `AGENTS.md`, como la tarea 14 de
  la #13.

### Cobertura

Las dieciséis reglas quedan cubiertas: RN-1 (1), RN-2 (1), RN-3 (1), RN-4 (1), RN-5 (5),
RN-6 (5), RN-7 (5), RN-8 (5), RN-9 (2), RN-10 (2), RN-11 (2), RN-12 (3), RN-13 (4),
RN-14 (6), RN-15 (2), RN-16 (2, 3, 4, 6).

| Criterio de aceptación | Tareas |
|---|---|
| Con la ronda revelada, se informa si hay diferencias | 1 |
| Se informan el más bajo y el más alto, con todos los que los dieron | 1 |
| Si todos votaron lo mismo, no hay diferencias | 1 |
| Sin revelar no se informa nada de los votos | 1 |
| Se abre una ronda nueva y las anteriores quedan guardadas | 5 |
| Las rondas de cada historia se numeran desde 1 y se consultan en orden | 5 |
| Se registra la estimación acordada y la historia queda estimada | 2 |
| El valor acordado tiene que ser de la escala, aunque nadie lo haya votado | 2 |
| No se acuerda sobre una ronda sin revelar, ya acordada o que no es la última | 3, 4, 6 |
| No se acuerda sobre una historia Terminada | 2 |
| Un acuerdo rechazado no cambia ni la ronda ni la historia | 2, 3, 4, 6 |

Los trece escenarios de la Fase 1 tienen su tarea, en la columna "Escenario BDD". La tarea 7
no cubre una regla de la Fase 1: es la regla de la #77 de que los datos sigan al cerrar y
volver a abrir.

### Dependencias

- **Todas las tareas necesitan el código de la #13 en `dev`**, que está escrito en el PR #141:
  la ronda, el voto, el revelado y la conversión de las rondas. Es la condición de la regla
  del 21/10 para adelantar una historia: lo que necesita tiene que estar en `dev`. Las tareas 1
  a 4 arman sus rondas con las operaciones de la #13; la 5 y la 6 cambian su apertura; la 7
  amplía su conversión.
- **La tarea 2 necesita una historia Terminada**: el test fija el estado desde dentro del
  paquete, como hizo la #12.

> **Puerta 3** — ¿las tareas cubren todos los criterios de aceptación?
> Aprobó: Juliana Bustos — Fecha: 2026-10-23 — Comentarios: aprobada en el PR #142 (https://github.com/emisorato1/seguimiento-medicion-ics/pull/142#issuecomment-6357143382), junto con las otras dos puertas. Sus tres observaciones menores (la cita de la RN-3 en la RN-11, los nombres de funciones fuera de la Fase 1 y una sola acción en el Cuando del escenario de dos rondas) se aplicaron antes del código.

## Fase 4 — Implementación

Ciclo RED → GREEN → REFACTOR. El test en un commit, con la salida del test fallando después
de `RED:` en el mensaje, y la implementación en otro.

| Tarea | Commit del test | Commit de la implementación |
|---|---|---|
| 1 | `2ab9574` | `84e068b` |
| 2 | `34f60c4` | `de58f0d` |
| 3 | `5808f95` | `b13c00e` |
| 4 | `5120a6d` | `01e8f5d` |
| 5 | `a6a5a5a` | `185bce2` |
| 6 | `b1c418c` | `c1ce166` |
| 7 | `03e2280` | `3fbd33a` |
| refactor | — | `ed24b67` |

Las siete tareas tuvieron RED, y cada commit se verificó por separado: los de test fallan y
los de implementación pasan. `internal/domain` queda en 100% de cobertura, y ningún test de la
#13 se modificó.

### Desvíos respecto de la spec

- **Tarea 7: un control de la carga estaba mal escrito y se sacó.** Decía que la carga
  rechazara "un acuerdo en una ronda que no es la última". Pero el escenario "Abrir una ronda
  nueva despues de acordar" deja justamente ese estado —la ronda 1 acordada y la 2 abierta—,
  que es válido: la RN-14 prohíbe *acordar* sobre una ronda que no es la última, no que una
  ronda acordada quede antes de otra. Con ese control, la web no habría podido volver a abrir
  su archivo después de reestimar una historia. En su lugar, la carga rechaza lo que de verdad
  no puede pasar: una ronda sin revelar seguida de otra de la misma historia, que además
  cubre el caso de "dos rondas sin revelar" de la #13. El test de ida y vuelta guarda una
  ronda acordada que ya no es la última. Lo decidió el responsable el 24/10, al encontrarlo, y
  se corrigió la tarea 7 y la decisión de la Fase 2.
- **La historia se adelantó al Sprint 3** (`32e4d17`), con la regla del 21/10 de CLAUDE.md:
  tenía las tres puertas aprobadas y la #13 ya estaba en `dev`.
- **Refactor `ed24b67`:** `CurrentVotingRound` (#13) pasa a usar la misma búsqueda de la
  última ronda que la apertura. No cambia comportamiento: los tests pasan sin tocarlos.

### Decisiones tomadas al implementar

Detalles que la Fase 2 no fijaba y que no cambian ningún escenario:

- **Las diferencias son un dato `Differences`**, con `HasDifferences`, `Lowest` y `Highest`;
  cada extremo es un `VoteExtreme` con el valor (`Points`) y todos los que lo dieron
  (`Members`), en el orden de los participantes. Es lo que la #78 muestra.
- **"Sin acordar" es 0** en `Agreed()`, como "sin estimar" en la #12.
- **El número de la ronda nueva sale de la última ronda de la historia**, con la misma
  búsqueda que usa `Agree` para saber si una ronda es la última.

### Defecto encontrado después del merge: #156

En la revisión del PR #148 (25/10), Juliana Bustos pidió un cambio antes de mergear. El #148
se mergeó por error con ese pedido sin responder, así que se registró como el defecto #156 y
se arregló en un PR aparte, a pedido de Emiliano Sorato.

- **Qué pasaba:** la carga aceptaba una ronda revelada **sin participantes**, porque a nadie le
  faltaba votar, y después `Differences()` cortaba con un pánico (`index out of range [0]`) al
  leer el primer voto. En la web (#78), una página rota. También aceptaba un participante con
  **nombre vacío** o **repetido**, y con el repetido `Differences()` nombraba dos veces a la
  misma persona.
- **El arreglo:** la carga rechaza los tres casos, con `ErrSavedDataInvalid`, como propuso
  Juliana. Al abrir una ronda ninguno puede pasar (RN-6 de la #13 y los integrantes de la #4),
  así que solo un archivo puede traerlos. Con eso el primer voto queda garantizado y
  `Differences()` no se tocó.
- **TDD:** `684fcb2` es el test, con su RED: el pánico, y las dos filas nuevas de datos
  dañados. `87264b3` es el arreglo.

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #73 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | `internal/domain/voting_round_agreement_test.go` |
| Código Go | `internal/domain/voting_round_agreement.go`, la apertura en `internal/domain/voting_round.go` y sus errores en `internal/domain/errors.go` |
