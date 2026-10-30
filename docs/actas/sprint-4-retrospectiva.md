# Acta — Retrospectiva del Sprint 4

> Se hizo por escrito, después de la [review](sprint-4-review.md). Es la última del proyecto,
> así que también mira el proyecto entero. Lo que está entre corchetes, **[pendiente]**, lo
> completa lo que diga cada uno: acá no se escribe nada que alguien no haya dicho.

- **Fecha:** 29 y 30/10, por escrito.
- **Respondieron:** Maximiliano Eula y Juliana Bustos (29/10, mirando el proyecto entero) y
  Emiliano Sorato (30/10), por escrito. Agustín Salinas: [pendiente].
- **Facilita:** Emiliano Sorato (Agile Enabler)

Los números con los que se abre la retro salen del [acta de review](sprint-4-review.md): 18 de
18 story points y 2 de 2 historias, velocidad 18, −93% de desviación (las dos historias, con
asistencia de IA), un defecto detectado y resuelto, y 95,8% de cobertura con 426 tests. Con eso,
las 21 historias del Product Backlog están terminadas: 109 story points.

## Qué salió bien

- **Emiliano Sorato:** "Terminamos la #19 y la #78 el mismo día, con su TDD, y la trazabilidad
  quedó completa: todos los escenarios tienen su test. Y armamos el CI/CD: cada incremento se
  despliega solo en el servidor".
- **Maximiliano Eula:** respondió mirando el proyecto entero (más abajo).
- **Juliana Bustos:** respondió mirando el proyecto entero (más abajo).
- **Agustín Salinas:** [pendiente]

## Qué no salió bien

- **Emiliano Sorato:** "El sprint iba a ser de un día y se estiró, porque la #78 y la #19
  esperaban a la #18. Y varios PRs del 30/10 entraron sin la firma de un compañero, porque no
  estaban disponibles".
- **Maximiliano Eula:** respondió mirando el proyecto entero (más abajo).
- **Juliana Bustos:** respondió mirando el proyecto entero (más abajo).
- **Agustín Salinas:** [pendiente]

## Seguimiento de los acuerdos del Sprint 3

Los cuatro acuerdos de la [retrospectiva del Sprint 3](sprint-3-retrospectiva.md):

| Acuerdo | ¿Se cumplió? |
|---|---|
| ESTADO.md lo regenera una sola persona, una vez por día, y no cada PR | En parte: lo regeneró solo Emiliano Sorato, pero el 30/10 fueron dos veces (PRs #217 y #227) |
| No se mergea un PR con cambios pedidos sin responder | Sí: ningún PR del sprint tenía cambios pedidos sin responder |
| Se mergea apenas el PR tiene su revisión y el CI en verde | Sí |
| Un test que pasa sin RED se revisa contra los otros lugares donde se aplica la misma regla | Sí: la tarea 24 de la #78 pasó sin RED, y antes de darla por cubierta se comprobó que el test fallaba rompiendo la plantilla |

## Mirando el proyecto entero

Qué aprendimos midiéndonos, cómo cambió la estimación sprint a sprint y qué haríamos
distinto. Lo que salga acá va a las conclusiones del informe final y a la presentación.

- **Maximiliano Eula:** 
  - "Con las specs escritas de antemano y adelantando historias, la velocidad subió mucho (de
    6 a 52)."
  - "Las horas de las historias hechas con IA no sirven para estimar: habría que medirlas
    aparte desde el principio."
  - "Casi todos los defectos los encontramos revisando, así que la revisión con los tests
    corridos fue lo que más nos ayudó con la calidad."
- **Emiliano Sorato:**
  - "Lo que más nos sirvió fue escribir la spec antes y pasarla por las puertas: con el
    contrato del dominio acordado, cada uno avanzó sin esperar al otro, y la velocidad pasó de
    6 en el Sprint 1 a 52 en el 2 y 33 en el 3."
  - "Medirnos nos mostró que en horas estimábamos mal: con 2,7 h por punto quedamos siempre muy
    por encima de lo real (−52% en el Sprint 2 y −51% en el 3). Los story points fueron más
    estables. Las horas de lo hecho con asistencia de IA habría que haberlas medido aparte
    desde el Sprint 1, como dicen Maxi y Juli."
  - "La calidad salió de los tests y de revisar con los tests corridos: el dominio quedó al 100%
    y casi todos los defectos salieron de revisiones o de la prueba exploratoria, no del uso."
  - "Qué haría distinto: arrancar con la base de datos y la web desde el principio, y no
    sumarlas con el proyecto avanzado, porque la #177 y la #78 fueron las que más trabajo dieron
    al final. Y armar antes el CD con el deploy, para mostrar cada incremento andando."
- **Juliana Bustos:** "Las horas estimadas quedaron por encima de las reales (-72 % en historias
  terminadas del Sprint 3). Las horas no miden bien nuestro ritmo con IA, mientras que los
  story points resultaron más estables. También aprendí que hay que cargar las horas reales
  inmediatamente al terminar cada historia y que la capa web (accesibilidad y responsividad)
  requiere más trabajo que la regla de negocio que muestra".
- **Agustín Salinas:** [pendiente]

## Pendientes

- [ ] La mirada del proyecto entero de Agustín Salinas.
- [ ] Llevar las conclusiones al informe final y a la presentación.
