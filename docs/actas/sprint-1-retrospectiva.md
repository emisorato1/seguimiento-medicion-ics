# Acta — Retrospectiva del Sprint 1

> Se hizo **por escrito, en el chat del grupo**, el 19/10, junto con la planning del Sprint 2,
> para cerrar las dos cosas en un rato en lugar de una reunión de dos horas. Lo que figura acá
> es lo que escribió cada uno, atribuido a quien lo dijo. Lo que está entre corchetes falta y
> no se da por hecho.

- **Fecha:** 19/10/2026, asincrónica por chat
- **Respondieron:** Emiliano Sorato, Agustín Salinas y un tercer integrante cuya respuesta
  llegó desde un número sin agendar: [confirmar si fue Juliana Bustos o Maximiliano Eula]
- **Sin respuesta registrada:** [el cuarto integrante]
- **Facilita:** Emiliano Sorato (Agile Enabler)

Los números con los que se abrió la retro, del [acta de review](sprint-1-review.md):
velocidad 11 de 16 story points, 3 de 5 historias terminadas, 38 horas reales contra 48
estimadas, 2 defectos encontrados y 1 resuelto, dominio al 100% de cobertura.

## Qué salió bien

- **Las revisiones encontraron fallas reales antes de que llegaran a `dev`.** Lo marcaron
  los tres. Emiliano: *"las puertas sirvieron de verdad — en las revisiones aparecieron tres
  bugs reales que ningún test habría agarrado"*. Agustín detalla los de la #3: el choque entre
  la RN-8 y la RN-9, la consulta de un proyecto inexistente, las fechas con hora y las zonas
  horarias. El tercer integrante también nombra lo de las zonas horarias.
- **La calidad técnica se sostuvo**: el dominio al 100% de cobertura (Agustín y el tercer
  integrante).
- **El TDD quedó demostrable** (Agustín): 24 commits separados entre test y código, y los
  desvíos anotados en la spec en vez de disimulados.
- **La #3 sirvió de base** (Agustín): dejó `errors.go` y el registro de proyectos, y eso
  destrabó a la #4, la #6 y la #24.
- **Los defectos quedaron registrados y documentados** (Agustín): el #59 resuelto y el #60
  listo para PR. El #60 se cerró después, el 19/10.

## Qué no salió bien

- **El reparto.** Maximiliano quedó con dos historias y las dos dependían de un PR de
  Emiliano (Emiliano). El tercer integrante lo nombra como sobrecarga.
- **No se respetó el flujo de aprobaciones.** Hubo merges con bypass de administrador —
  Emiliano lo reconoce: *"el acuerdo (b) no lo cumplí, mergeé varios PRs con bypass"*— y PRs
  mergeados por quienes no eran sus autores (tercer integrante).
- **Se mergeó un PR con un pedido de cambios sin resolver** (Agustín): el #31, donde había
  marcado un hueco en la decisión 5 de `AGENTS.md`. [El resto de su respuesta llegó cortada
  por el chat: completar.]

## Seguimiento de los acuerdos del Sprint 0

| Acuerdo | ¿Se cumplió? |
|---|---|
| Cada integrante toma al menos una historia y la lleva de punta a punta | **Parcial.** Agustín (#3), Juliana (#4) y Emiliano (#6) sí. Maximiliano terminó el sprint con la #12 en review y la #24 en progreso: las dos pasaron al Sprint 2. |
| Ningún PR se mergea sin la aprobación de otro integrante, sin bypass de admin | **No.** Lo reconocen Emiliano y el tercer integrante. |
| Revisar las fechas del campus cada semana y avisar tres días antes | [no se habló en la retro] |
| Las historias se estiman en la planning, antes de abrir el sprint | [no se habló en la retro]. El [acta de planning del Sprint 2](sprint-2-planning.md) registra que en el Sprint 1 se estimó por afinidad y no con Planning Poker. |

Un acuerdo que no se cumplió dos sprints seguidos no es un acuerdo: o se cambia o se saca.

## Qué vamos a cambiar en el Sprint 2

**Propuesto en la retro** por el tercer integrante: *"cumplir a rajatabla los acuerdos del
Sprint 0: 1 historia por persona, 0 bypass en PRs y que cada autor mergee su propio PR tras
la aprobación"*.

**Lo que cambió al día siguiente.** El 20/10 el equipo decidió entregar el proyecto completo
para el 2 de noviembre ([plan](../plan-final-al-2-de-noviembre.md)), y eso modificó dos de
las tres propuestas:

| Propuesta de la retro | Qué quedó | Por qué |
|---|---|---|
| Una historia por persona | **Reparto parejo por story points**: entre 17 y 19 por persona, nadie colgado de un PR ajeno | con 72 puntos en 13 días, una historia por persona no alcanza. Lo que la retro marcó como problema era la sobrecarga de uno solo, y eso es lo que se corrige |
| Cero bypass de admin | **Se mantiene.** Hacia `dev` deja de hacer falta: alcanza con el CI en verde y la revisión escrita. Hacia `main` sigue la aprobación obligatoria | el bypass apareció porque la aprobación obligatoria dejaba los PRs parados un día. Sacándola de `dev` desaparece la razón para saltearla |
| Cada autor mergea su propio PR | **Se mantiene**, ya estaba en `CLAUDE.md` | — |

| Acuerdo | Responsable |
|---|---|
| Nadie con más carga que el resto ni dependiendo de un PR ajeno para arrancar | Emiliano, en el plan y en cada planning |
| Cero bypass de administrador, ni hacia `dev` ni hacia `main` | todos |
| Cada autor mergea su propio PR | todos |
| No se mergea un PR con un pedido de cambios sin resolver | quien mergea |

[Confirmar estos cuatro acuerdos en el grupo.]

## Pendientes

- [ ] Confirmar quién es el tercer integrante que respondió.
- [ ] Completar la respuesta de Agustín, que llegó cortada.
- [ ] Sumar la respuesta del cuarto integrante, si la quiere dar.
- [ ] Confirmar los acuerdos en el grupo.
