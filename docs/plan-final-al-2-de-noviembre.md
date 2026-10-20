# Plan final — entregar el proyecto completo el 2 de noviembre

Decidido por el Agile Enabler el **20/10/2026**. Reemplaza las dos opciones de recorte de
[replan-al-2-de-noviembre.md](replan-al-2-de-noviembre.md), que quedan como registro del
análisis previo.

## La decisión

**Se entrega el alcance completo.** Las 20 historias, los 83 story points, los nueve
requerimientos mínimos terminados. Nada se presenta como "backlog pendiente".

El replan del 20/10 proponía recortar a unos 30 puntos y defender el resto como backlog. Se
descarta: el equipo decide subir el ritmo en vez de bajar el alcance.

## El número, sin maquillar

| | |
|---|---|
| Story points que quedan | **72** de 83 |
| Días hasta la presentación | **13** (20/10 al 1/11) |
| Costo medido en el Sprint 1 | 2,6 h por story point |
| Horas necesarias a ese costo | **187** |
| Por persona, por día, sábados y domingos incluidos | **~3,6 h** |

En el Sprint 1 cada uno puso unas **9,5 horas por semana** (≈1,4 h por día). Este plan pide
**dos veces y media eso, trece días seguidos**.

**Por qué igual cierra.** Las 2,6 h por punto midieron un proceso con cuatro o cinco pull
requests por historia y revisiones que tardaban un día. Con una historia por pull request,
las specs escritas de antemano y revisión en el día, el costo por punto tiene que bajar a
unas **2,0 h**. A ese costo hacen falta 144 horas: **2,8 h por persona por día**, y aparecen
unos 19 puntos de colchón.

**Lo que no se puede negociar.** No hay margen para que uno de los cuatro no sostenga el
ritmo. Si alguien no puede, se dice ahora y se recorta con tiempo, no el 31 a la noche.

## Calendario

| | Fechas | Días | SP | Objetivo |
|---|---|---|---|---|
| **Sprint 2** | mar 20 → dom 25/10 | 6 | 26 | Sprints, esfuerzo, defectos y consola |
| **Sprint 3** | lun 26 → vie 30/10 | 5 | 36 | Métricas, dashboard, Planning Poker y reportes |
| **Sprint 4** | sáb 31/10 → dom 1/11 | 2 | 10 | PDF, informe, manual y presentación |
| **Presentación** | lun 2/11 | — | — | Defensa |

Cada sprint cierra con review y retrospectiva **el último día**, no al día siguiente.

## Quién hace qué

Reparto total, pensado para que nadie espere a nadie:

| Quién | Historias | SP |
|---|---|---|
| **Agustín Salinas** | #7, #8, #9, #11, #18 | 17 |
| **Juliana Bustos** | #5, #14, #15, #17 | 19 |
| **Maximiliano Eula** | #12, #24, #13, #20 | 18 |
| **Emiliano Sorato** | #10, #16, #19 | 18 |
| | **Total** | **72** |

### Sprint 2 — del 20 al 25 de octubre (26 SP)

| Quién | Historias | SP |
|---|---|---|
| Agustín | #7 Priorizar el backlog (2), #8 Crear un sprint (3), #9 Asignar historias (5) | 10 |
| Juliana | #14 Registrar el esfuerzo real (3), #15 Gestionar defectos (5) | 8 |
| Maximiliano | #12 Story Points (2), #24 Interfaz de consola (3), #5 Consultar el estado (3) | 8 |
| Emiliano | **las specs de las 15 historias** y las revisiones | — |

**Sprint Goal:** que el sistema maneje sprints, esfuerzo y defectos, y que todo se pueda
operar desde la consola.

Emiliano no toma historias esta semana a propósito: su trabajo es dejar **todas las
especificaciones escritas y aprobadas antes del jueves 22**, para que en el Sprint 3 nadie
espere una puerta. Es el cuello de botella que mató al Sprint 1.

### Sprint 3 — del 26 al 30 de octubre (36 SP)

| Quién | Historias | SP |
|---|---|---|
| Emiliano | #10 Cerrar un sprint (5), #16 Calcular las métricas (8) | 13 |
| Juliana | #17 Mostrar el dashboard (8) | 8 |
| Maximiliano | #13 Votar en secreto en una ronda (8) | 8 |
| Agustín | #11 Consultar sprints anteriores (2), #18 Generar el reporte (5) | 7 |

**Sprint Goal:** que el sistema mida el proyecto y lo muestre: métricas, dashboard y reporte.

Es el sprint más cargado y el que decide la presentación.

### Sprint 4 — 31 de octubre y 1 de noviembre (10 SP)

| Quién | Historias | SP |
|---|---|---|
| Emiliano | #19 Exportar el reporte a PDF (5) | 5 |
| Maximiliano | #20 Resolver diferencias y acordar (5) | 5 |
| Agustín | documentación técnica | — |
| Juliana | presentación final | — |

**Sprint Goal:** cerrar el producto y tener todo lo que se entrega listo para presentar.

## Entregables que no son historias

Van en paralelo con el código, no después:

| Entregable | Quién | Para cuándo |
|---|---|---|
| Informe de métricas y cobertura | Emiliano | 1/11 |
| Manual breve de usuario | Maximiliano | crece con cada tarea de la consola, cierra 1/11 |
| Documentación técnica | Agustín | 1/11 |
| Presentación final | Juliana arma, los cuatro revisan | 1/11 |
| Actas de planning, review y retrospectiva | Emiliano | al cierre de cada sprint |
| Retrospectiva del Sprint 1 (pendiente) | Emiliano | 21/10 |

## Las dependencias, y cómo no nos frenan

La cadena real del proyecto es:

```
#8 → #9 → #10 → #16 → #17
                  └──→ #18 → #19
```

Son 39 de los 72 puntos encadenados. Si se respeta tal cual, el dashboard y el reporte no
empiezan hasta el 29 y no llegan.

**Se rompe con un contrato escrito antes de implementar.** El documento
[diseno-dominio.md](diseno-dominio.md) fija, antes de que nadie escriba código, qué entidades
existen y qué datos tiene cada una: `Sprint`, `Effort`, `Defect`, `Metrics`, `Report` y
`VotingRound`. Con eso:

- **#17 y #18 arrancan sin esperar a #16**: se escriben contra el contrato y se prueban con
  datos de prueba armados a mano en el propio test.
- **#16 arranca sin esperar a #10**: la métrica se calcula sobre sprints cerrados, y un
  sprint cerrado de prueba se construye en el test.
- Cuando la pieza real aterriza en `dev`, se conectan. Si algo no encaja, **el que rompe el
  contrato avisa al grupo antes de cambiarlo**, no después.

Esto no es optativo: es lo único que hace que 36 puntos entren en cinco días.

## Decisiones de proceso

Diez cambios, vigentes hasta el 2/11.

1. **Las specs se escriben todas ahora, no sprint por sprint.** Emiliano deja las 15
   especificaciones completas antes del jueves 22. Las tres puertas de cada historia se
   aprueban en una sola pasada.
2. **El contrato de dominio va primero.** Nadie implementa una entidad compartida antes de
   que esté en `diseno-dominio.md`.
3. **Una historia, un pull request**, con la spec, los tests y el código. El orden SDD se
   demuestra en los commits: spec, después test, después código.
4. **Un pull request puede cerrar varias historias** si son eslabones seguidos de la misma
   cadena (por ejemplo #8 y #9).
5. **`dev` deja de pedir aprobación obligatoria.** El CI en verde sigue siendo obligatorio y
   toda PR sigue recibiendo una revisión escrita. `main` mantiene la aprobación obligatoria:
   es la rama que se presenta.
6. **Cada uno corre su historia con Claude Code**, con `CLAUDE.md`, `AGENTS.md` y su spec
   como entrada. Lo que vuelve se lee y se entiende antes de subirlo: en la defensa responde
   la persona, no la herramienta.
7. **Revisión en el día.** Un pull request abierto antes de las 18:00 se revisa ese mismo día.
8. **Daily asíncrona por el grupo, a las 21:00**, una línea: qué cerré, qué sigo, si estoy
   trabado.
9. **Regla de las dos horas.** Trabado más de dos horas, se avisa al grupo. No se pierde un
   día entero peleando solo.
10. **La #13 se parte en dos.** Queda #13 *Votar en secreto en una ronda* (8) y #20
    *Resolver diferencias y acordar la estimación* (5). Eran 13 puntos en una sola historia,
    demasiado para un pull request y demasiado para un dueño solo.

**Lo que no cambia: la Definición de Terminado.** Spec antes del código, escenarios BDD,
tests que pasan, revisión escrita y cobertura. Se recorta ceremonia, nunca calidad.

## Punto de control: domingo 25 a la noche

El Agile Enabler mira el tablero y decide:

- **Si el Sprint 2 cerró sus 26 puntos**, el plan sigue tal cual.
- **Si cerró 20 o más**, se sostiene el alcance y se recorta #19 (PDF) a exportar un texto
  plano con extensión `.pdf` mínima, o se deja como backlog.
- **Si cerró menos de 20**, se recorta de verdad: salen #19 y #20, y #17 queda en una
  versión de consola sin gráficos.

La decisión se toma **ese día y se avisa al grupo**, no se patea.

## Riesgos

| Riesgo | Qué hacemos |
|---|---|
| Alguien no sostiene el ritmo | La daily de las 21:00 lo muestra en un día, no en una semana. Se redistribuye al otro día. |
| El Sprint 3 tiene 36 puntos en 5 días | Es el riesgo principal. Lo contiene el contrato de dominio y el punto de control del 25. |
| #16 métricas depende de todo | Se construye contra datos de prueba armados en el test, no contra las otras historias. |
| El PDF (#19) sin librerías externas | Se resuelve con la salida mínima que exige la consigna. Es el primer candidato a recortarse. |
| La calidad se cae por correr | El CI no se toca: formato, análisis, tests y cobertura siguen bloqueando el merge. |
