# Acta — Review del Sprint 4

> Cerrada el 30/10 con `scripts/metricas.sh "Sprint 4"` y los datos del tablero. Acá no se
> escribe nada que alguien no haya dicho. La retrospectiva va en
> [un archivo aparte](sprint-4-retrospectiva.md). La review final del proyecto es la
> presentación del lunes 2/11.

- **Fecha:** 30/10, por escrito.
- **Participantes:** Emiliano Sorato, por escrito. Los demás integrantes no estaban disponibles
  el 30/10.
- **Facilita:** Emiliano Sorato (Agile Enabler)
- **Sprint:** del jueves 29 al viernes 30/10. Iba a durar un día, y siguió el 30/10 para
  terminar la #78 y la #19 (ver la [planning](sprint-4-planning.md)).
- **Sprint Goal:** cerrar el producto y tener todo lo que se entrega listo para presentar.
- **¿Se cumplió el Sprint Goal?** Sí. Las 21 historias del Product Backlog están terminadas, con
  sus 109 story points, el reporte se descarga en PDF y cada incremento se despliega solo en el
  servidor. Queda la presentación, que arma Juliana Bustos para el 2/11.

## Qué se completó

Una historia cuenta como completada solo si cumple la Definición de Terminado: spec con sus
tres puertas, escenarios BDD, tests en Go que pasan, revisada y mergeada a `dev`.

| # | Historia | SP | Estado | Responsable | Evidencia |
|---|---|---|---|---|---|
| 19 | Exportar el reporte a PDF | 5 | **Hecho** | Emiliano Sorato (con asistencia de IA) | PR #184 (tareas 1 a 12) y PR #222 (`r.WritePDF`, merge `9459bb6`, 30/10) |
| 78 | Operar todo desde la web (vuelve del Sprint 3) | 13 | **Hecho** | Emiliano Sorato (con asistencia de IA) | PRs #149, #163, #165 y #170, y el PR #224 con el reporte y su PDF (merge `bfd7742`, 30/10) |

- **Story points comprometidos:** 18.
- **Story points completados:** 18 (100%).
- **Velocidad del Sprint 4:** 18. No hubo adelantadas.
- **Historias completadas:** 2 de 2 (100%).
- El PR #224 lo revisó Emiliano Sorato, porque los demás no estaban disponibles el 30/10. La
  revisión fue la página del reporte en Chromium, a 1280 y a 390 de ancho, el error de un sprint
  mal escrito junto a su campo, el recorrido con el teclado y la descarga del PDF.

### Agregadas durante el sprint

| # | Qué | SP | Estado | Responsable | Evidencia |
|---|---|---|---|---|---|
| 219 | Desplegar la web en un servidor con el CD | sin estimar | **Hecho** | Emiliano Sorato (con asistencia de IA) | PRs #218 y #221. El servidor lo prestó Maximiliano Eula |
| 223 | Mejorar la presentación visual de la web para la demo | sin estimar | En review | Emiliano Sorato (con asistencia de IA) | PR #226, que revisa Juliana Bustos |

### Lo que se entrega además de las historias

| Qué | Responsable | Estado |
|---|---|---|
| Informe de métricas y cobertura | Emiliano Sorato | `docs/informe-de-metricas.md`, con los cuatro sprints |
| Informe final en PDF | Emiliano Sorato | `docs/informe-final.md`, y el PDF exportado desde Word |
| Manual breve de usuario | Maximiliano Eula | al día con el dashboard (PR #216) y el reporte (PR #224) |
| Documentación técnica | Agustín Salinas | al día con la #19 (PR #224) y el CI/CD (PR #227) |
| Presentación final | Juliana Bustos arma, los cuatro revisan | [pendiente] |
| Incremento final en `main`, con el tag `v1.0.0` | Emiliano Sorato | [pendiente] |

## Esfuerzo

| Historia | Estimadas | Reales | Desviación |
|---|---|---|---|
| #19 | 13,5 | 1 | −12,5 h |
| #78 | 35,1 | 2,5 | −32,6 h |
| **Total de las terminadas** | **48,6** | **3,5** | **−45,1 h (−93%)** |

- **Las dos se hicieron con asistencia de IA.** Sus horas reales son el tiempo de Emiliano Sorato
  para dirigir y revisar, y por eso la desviación da tan negativa. Las cargó él en el tablero el
  30/10.
- **La #78 se mide entera:** sus 2,5 h son 1,5 del Sprint 3, que ya se contaron ahí, y 1 del 4.
- **Horas trabajadas en el sprint:** 3,5. Son 1 de la #19, 1 de la #78 y 1,5 de la #219.
- **Horas por integrante:** Emiliano Sorato 3,5. Los demás no cargaron horas en el sprint.
- **Horas por story point completado:** 0,2 (3,5 h / 18 SP). Con todo hecho con asistencia de IA,
  no sirve para comparar el costo de un punto con los otros sprints.

## Defectos

| # | Defecto | Severidad | Detectado | Resuelto |
|---|---|---|---|---|
| 220 | Las diferencias de una ronda unen tres o más nombres con «y» | Menor | Sprint 4 (29/10, Maximiliano Eula) | Sprint 4 (30/10, PR #225) |

- **Detectados en el sprint:** 1. **Resueltos:** 1. **Abiertos al cierre:** ninguno.

## Calidad

Medido con `go test -cover ./...` sobre `dev` en `89e4afd` (30/10), con las dos historias
adentro. Los 426 tests pasan, en Linux y en Windows.

| Paquete | Cierre del Sprint 3 (`5d1d556`) | Cierre del Sprint 4 (`89e4afd`) |
|---|---|---|
| `internal/domain` | 100% | 100% |
| `internal/store` | 87,1% | 87,1% |
| `internal/web` | 94,1% | 94,4% |
| `cmd/app` | 96,2% | 96,2% |
| `cmd/web` | 91,1% | 91,1% |
| **Total** | **95,7%** | **95,8%** |
| Tests | 420 | 426 |

Con la #19 y la #78, la [trazabilidad](../trazabilidad.md) da todos los escenarios de las
historias con su test: la #19, 13 de 13, y la #78, 22 de 22.

## Demostración

La web quedó andando en el servidor desde el 30/10, con los datos de demo, y el
[guion de la demo](../demo/guion-de-demo.md) suma el reporte y su PDF. La demo de la review
final es la presentación del 2/11.

## Pendientes

- [x] Cerrar los números con `scripts/metricas.sh "Sprint 4"` (30/10).
- [x] Las horas reales de la #19 y la #78, cargadas el 30/10.
- [ ] El pull request de `dev` a `main` con la entrega final y su tag `v1.0.0`.
- [x] La [retrospectiva](sprint-4-retrospectiva.md) (falta la mirada del proyecto entero de
  Agustín Salinas).
