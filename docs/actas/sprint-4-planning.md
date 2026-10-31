# Acta — Sprint Planning del Sprint 4

> Armada el 27/10, puesta al día el 29/10 y cerrada el 30/10, después de la review y la retro
> del Sprint 3. Lo que está entre corchetes, **[pendiente]**, se completa con lo que diga cada
> uno.

- **Fecha:** 29 y 30/10, por escrito.
- **Participantes:** Emiliano Sorato, Maximiliano Eula, Juliana Bustos y Agustín Salinas, por
  escrito.
- **Facilita:** Emiliano Sorato (Agile Enabler)
- **Duración del sprint:** jueves 29/10, un solo día. La review y la retrospectiva se hacen el
  29 a la noche. La review final es la presentación del lunes 2/11. Se decidió el 26/10 a
  pedido del equipo (ver el [plan](../plan-final-al-2-de-noviembre.md#calendario)). La #78 y la
  #19 no entraron el 29/10: el 30/10 Emiliano Sorato decidió seguir hasta terminarlas, antes de
  la presentación.

## Lo que traemos del Sprint 3

- **Velocidad del Sprint 3:** 33 (28 completados + 5 de la #73 adelantada), según la
  [review](sprint-3-review.md).
- **Lo que no se terminó en el Sprint 3** y vuelve al Product Backlog: la #78 (13 SP), con 22 de
  sus 24 tareas en `dev`. Entra a este sprint para terminarla (decisión de Emiliano Sorato, 30/10).

## Sprint Goal

Propuesta, que sale del plan: **cerrar el producto y tener todo lo que se entrega listo para
presentar.** Lo confirmaron los cuatro: Maximiliano Eula y Juliana Bustos el 29/10, y Emiliano
Sorato y Agustín Salinas el 30/10.

Es el objetivo que la consigna le da al Sprint 4: pulir, documentar y preparar la entrega
final, con la exportación del informe del proyecto a PDF.

## Sprint Backlog

| # | Qué | SP | Responsable | Estado al 29/10 |
|---|---|---|---|---|
| 19 | Exportar el reporte a PDF | 5 | Emiliano Sorato (con asistencia de IA) | las tareas 1 a 12 entraron a `dev` el 27/10 (PR #184). Con la #18 en `dev` desde el 29/10, falta el segundo PR: conectar el PDF con el reporte |
| 78 | Operar todo desde la web (vuelve del Sprint 3) | 13 | Emiliano Sorato (con asistencia de IA) | 22 de 24 tareas en `dev`. Faltan la 20 y la 21, el reporte y su PDF en la web, que necesitan la #19 |
| — | Informe de métricas y cobertura, y el informe final en PDF | — | Emiliano Sorato | `docs/informe-de-metricas.md` en `dev`; el informe final en una rama |
| — | Manual breve de usuario | — | Maximiliano Eula | `docs/manual-de-usuario.md` en `dev`; la puesta al día con la base de datos y las horas del defecto #192, en el PR #208 |
| — | Documentación técnica | — | Agustín Salinas | `docs/documentacion-tecnica.md` en `dev`, ya con la base SQLite de la #177 |
| — | Presentación final | — | Juliana Bustos arma, los cuatro revisan | [pendiente] |
| — | Pull request de `dev` a `main` con el incremento final, la entrega, y su tag `v1.0.0` | — | Emiliano Sorato | el 29/10, después del incremento del Sprint 3 con su tag `v0.3.0-sprint3` |
| — | Actas de [review](sprint-4-review.md) y [retrospectiva](sprint-4-retrospectiva.md) del Sprint 4 | — | Emiliano Sorato | borradores armados el 29/10 |

La #78 vuelve del Sprint 3 sin terminar y se termina en este sprint. Con ella va el defecto
#220 (Menor), que cargó Maximiliano Eula el 30/10: en el Planning Poker, tres o más nombres se
unen con «y».

- **Story points comprometidos:** 18, los 5 de la #19 y los 13 de la #78. La #19 no se adelantó
  al Sprint 3. La #73 no cuenta acá: ya está terminada, se adelantó al Sprint 3 y suma a la
  velocidad de ese sprint.
- **Horas estimadas: 48,6.** Son las 13,5 de la #19 (5 SP a 2,7 h por punto), que cargó en el
  tablero Emiliano Sorato, su responsable, el 27/10, y las 35,1 de la #78. De la #78 ya se
  contaron 1,5 h reales en el Sprint 3: en este sprint se cuentan solo las nuevas.

## Confirmación de cada responsable

- **Emiliano Sorato:** "De acuerdo con el objetivo. Y sumamos la #78, que vuelve del Sprint 3:
  le faltan el reporte y el PDF en la web, y los termino junto con la #19" (30/10).
- **Maximiliano Eula:** "de acuerdo con el objetivo" (29/10).
- **Juliana Bustos:** "De acuerdo: Cerrar el producto y dejar todo lo que se entrega listo para la
  defensa del lunes" (29/10).
- **Agustín Salinas:** "Sí, estoy de acuerdo con el objetivo: cerrar el producto y tener todo lo
  que se entrega listo para presentar" (30/10).

## Pendientes

- [x] El Sprint Goal confirmado por los cuatro (30/10).
- [x] Lo que vuelve del Sprint 3: la #78 (30/10).
- [x] Las horas estimadas de la #19: 13,5 h, el 27/10.
