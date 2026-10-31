# Acta — Retrospectiva del Sprint 2

- **Fecha:** del 24 al 26/10/2026
- **Cómo se hizo:** por escrito, en la [issue #147](https://github.com/emisorato1/seguimiento-medicion-ics/issues/147), con un comentario de cada uno
- **Respondieron:** Emiliano Sorato (24/10), Maximiliano Eula (24/10), Juliana Bustos (25/10) y Agustín Salinas (25/10)
- **Sin respuesta registrada:** nadie
- **Facilita:** Emiliano Sorato (Agile Enabler)

Lo que sigue es lo que escribió cada uno, con sus palabras. Los enlaces llevan a su comentario:
[Emiliano](https://github.com/emisorato1/seguimiento-medicion-ics/issues/147#issuecomment-6390684355),
[Maximiliano](https://github.com/emisorato1/seguimiento-medicion-ics/issues/147#issuecomment-6391163190),
[Juliana](https://github.com/emisorato1/seguimiento-medicion-ics/issues/147#issuecomment-6400529426) y
[Agustín](https://github.com/emisorato1/seguimiento-medicion-ics/issues/147#issuecomment-6406773200).

Los números con los que se abre la retro salen del [acta de review](sprint-2-review.md): **31 de
31 story points** comprometidos y las 9 historias terminadas, **velocidad 39** con la #13
adelantada, **31 horas** trabajadas, desviación de las terminadas de **−52%** (−36% sin las
historias hechas con asistencia de IA), **ningún defecto abierto** y **96,0%** de cobertura.

## Qué salió bien

- **Emiliano Sorato:** "cerramos el sprint completo (31 de 31) y con las specs escritas de
  antemano casi nadie esperó una puerta".
- **Maximiliano Eula:** "terminé la #12, la #24 y la #13, y en las revisiones aparecieron
  problemas reales antes de que llegaran a `dev` (el puerto ocupado en Windows, el `%#v` que
  dejaba ver los votos ocultos y el ancho de las barras del dashboard escrito con coma)".
- **Juliana Bustos:** "el flujo SDD, BDD y TDD se respetó. Las tres puertas de varias historias
  complejas (#17, #73 y #11) se aprobaron antes de tocar código, y la integración a dev no trajo
  sorpresas".
- **Agustín Salinas:** "con las specs escritas de antemano pude encadenar la #8 y la #9 sin
  esperar a nadie, y las revisiones propias antes de pedir revisión encontraron problemas
  reales (en la #9, una historia repetida en el Sprint Backlog que se aceptaba al cargar el
  archivo)".

## Qué no salió bien

- **Emiliano Sorato:** "hubo comentarios en los PRs que primero marcaban algo como mal y después
  como bien, y eso confundió; además, Maxi se quedó parado esperando la planning y la cadena #8
  → #9 → #10 frenó a varios".
- **Maximiliano Eula:** "el #141 estuvo casi un día sin revisión, los comentarios que mandé
  desde PowerShell salieron sin saltos de línea, y quedé parado esperando hasta que acordamos
  adelantar historias".
- **Juliana Bustos:** "tardé casi 24 h en revisar el #141 (reconozco lo que marcó Maxi). Al
  juntarse revisiones hacia el final del sprint se armaron cuellos de botella puntuales".
- **Agustín Salinas:** "`registry_json.go` lo tocamos varios a la vez y al juntar la #8 y la #9
  con la #14 hubo conflictos; además, mis commits de test no traían la línea `RED:` con la
  salida del test, como marcó Emiliano".

## Seguimiento de los acuerdos del Sprint 1

Los acuerdos salen de la [retrospectiva del Sprint 1](sprint-1-retrospectiva.md). En la #147 se
preguntó por tres: cero bypass, revisión en el día y no mergear con cambios pedidos sin
responder. Los otros dos no se preguntaron.

| Acuerdo | Responsable | ¿Se cumplió? |
|---|---|---|
| Nadie con más carga que el resto ni dependiendo de un PR ajeno para arrancar | Emiliano | no se preguntó en la #147 |
| Cero bypass de administrador, ni hacia `dev` ni hacia `main` | todos | **cumplido**, según los cuatro |
| Cada autor mergea su propio PR | todos | se cambió el 20/10 (ver abajo); no se preguntó en la #147 |
| Un PR con cambios pedidos no se mergea hasta que se respondan | quien mergea | **cumplido**, según los cuatro |
| Un PR abierto antes de las 18:00 se revisa ese mismo día | quien revisa; si a las 20:00 nadie lo hizo, Emiliano | **en parte** (ver abajo) |

Lo que dijo cada uno:

- **Emiliano Sorato:** "cero bypass, cumplido. Revisión en el día, cumplido casi siempre. Cambios
  pedidos sin responder: el #81 tenía un pedido de Agus y no se mergeó hasta corregirlo,
  cumplido aunque tardó".
- **Maximiliano Eula:** "se cumplieron".
- **Juliana Bustos:** "Cero bypass: cumplido. En el Sprint 2 no vi merges con bypass de
  administrador. Revisión en el día: no la cumplí con el #141. Por eso propongo la franja fija.
  No mergear con cambios pedidos sin responder: cumplido. El #141 lo mergeé después de dejar la
  revisión escrita y con el CI en verde, sin pedidos pendientes".
- **Agustín Salinas:** "cero bypass, cumplido. Revisión en el día: cumplida con los PRs; con las
  puertas de la #10 no, me las pidieron el 22/10 y las aprobé el 25/10. No mergear con cambios
  pedidos sin responder: cumplido (el #81 no se mergeó hasta corregir lo que pedí)".

El acuerdo "cada autor mergea su propio PR" se cambió el 20/10: hacia `dev` mergea la sesión de
revisión, después de dejar la revisión escrita y con el CI en verde, para que un PR ya revisado
no quede esperando a su autor (está en `CLAUDE.md`).

## El modo sprint final

Desde el 20/10 rigen los cambios del *Modo sprint final* de `CLAUDE.md`. En la #147 no se
preguntó por cada uno. Lo que sigue es lo que alguien escribió sobre ellos; donde nadie dijo
nada, queda así.

| Cambio | Lo que se dijo |
|---|---|
| Las specs se escriben todas por adelantado, con las tres puertas en una sola revisión | Emiliano y Agustín dicen que, con las specs escritas de antemano, casi nadie esperó una puerta y se pudo encadenar la #8 y la #9. Juliana, que las puertas de la #17, la #73 y la #11 se aprobaron antes de tocar código |
| El contrato del dominio (`docs/diseno-dominio.md`) va antes que el código | nadie dijo nada |
| No se espera a que la dependencia esté en `dev`: se arma a mano en el test | nadie dijo nada |
| Una historia, un pull request | nadie dijo nada |
| Revisión en el día | ver el seguimiento de los acuerdos |
| Daily asíncrona a las 21:00, una línea cada uno | los cuatro contestaron "ok" para seguirla en el Sprint 3 |
| Regla de las dos horas: trabado más de dos horas, se avisa | nadie dijo nada |

## Qué vamos a cambiar en el Sprint 3

**Lo que se propuso en la retro:**

- **Emiliano Sorato:** "aprobar las puertas en el día igual que las revisiones, y seguir
  adelantando historias cuando alguien termina lo suyo".
- **Maximiliano Eula:** "pedir revisor apenas se abre un PR, y seguir adelantando la próxima
  historia del plan apenas uno termina lo suyo".
- **Juliana Bustos:** "reservar una franja fija por día para las revisiones de PRs asignados,
  antes de seguir con el código propio, para cumplir la revisión en el día". Además propone
  "mantener los commits separados de test: (con la salida del RED en el mensaje) y feat:, que es
  la evidencia del TDD".
- **Agustín Salinas:** "poner la salida del RED en cada commit de test y "sin RED" en el asunto
  cuando el test pasa de entrada, y avisar por el grupo antes de tocar un archivo compartido
  como `registry_json.go`".

**Lo que quedó:**

Emiliano Sorato votó el 26/10 por tres propuestas, y Maximiliano Eula por las mismas tres el
29/10. Los votos de Juliana y Agustín están [pendientes].

| Propuesta | Emiliano | Maximiliano | Juliana | Agustín |
|---|---|---|---|---|
| Reservar una franja fija por día para revisar los PRs asignados | sí | sí | [pendiente] | [pendiente] |
| Pedir revisor apenas se abre el PR | sí | sí | [pendiente] | [pendiente] |
| Seguir adelantando la próxima historia apenas uno termina lo suyo | sí | sí | [pendiente] | [pendiente] |

| Acuerdo | Responsable |
|---|---|
| [pendiente: queda como acuerdo lo que elija el equipo] | [pendiente] |

Dos propuestas coinciden entre varios. Emiliano y Maximiliano proponen seguir adelantando
historias, y eso ya es la decisión 11 del plan desde el 21/10. Juliana y Agustín proponen el
RED en el commit del test, que es regla de `AGENTS.md` desde el 20/10.

## Pendientes

- [ ] Que el equipo elija los acuerdos del Sprint 3 entre las propuestas.
