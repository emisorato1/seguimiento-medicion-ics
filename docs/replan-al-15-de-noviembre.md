# Replanificación: terminar el proyecto al 15 de noviembre

Preparado por el Agile Enabler el **19/10/2026**, después de medir la velocidad del Sprint 1.

## El problema, en números

| | |
|---|---|
| Backlog total | 83 story points |
| Terminados | 11 (13%) |
| **Quedan** | **72** |
| Semanas hasta el 15/11 | **4** |
| Necesario | **18 story points por semana** |
| Velocidad medida en el Sprint 1 | **5,5 por semana** |

Hace falta **3,3 veces** la velocidad actual. Eso no se consigue apretando el acelerador:
hay que cambiar cómo trabajamos, porque el problema no es que se programe lento.

## Dónde se fue el tiempo

El Sprint 1 produjo 11 story points en 38 horas: **3,5 horas por punto**. Pero el código de
las tres historias terminadas se escribió en los últimos días. Lo que consumió el sprint fue
el **recorrido de cada historia por el proceso**:

Cada historia necesitó entre **cuatro y cinco pull requests** — Fase 1, Fase 2, Fase 3,
implementación, y los registros de las puertas — y cada uno esperó entre horas y días a que
alguien lo aprobara. En el Sprint 1 hubo **cuatro días** de un PR parado que bloqueaba a dos
personas, y varias idas y vueltas por aprobaciones descartadas al actualizar ramas.

La espera, no la escritura, es lo que hay que recortar.

## Los cambios

### 1. Una historia, un pull request

**La spec completa, los tests y el código van en el mismo PR**, y las tres puertas se
aprueban en esa única revisión.

El orden SDD se mantiene y se sigue demostrando: los commits muestran spec → test → código,
en ese orden y con fechas. Lo que desaparece es la **latencia** de cuatro revisiones
separadas por historia.

El apunte de la cátedra lo habilita: *"las cuatro fases representan el flujo completo, pero
no todos los problemas necesitan las cuatro con la misma profundidad… la habilidad está en
calibrar la intensidad del proceso al tamaño del problema"*.

### 2. Las Fases 2 y 3 se fusionan en las historias de hasta 3 puntos

Para una historia chica, el diseño son tres líneas y las tareas son cinco. Separarlas en dos
documentos con dos puertas cuesta más que lo que aporta. Se escriben juntas, bajo el título
**Diseño y tareas**, y pasan una sola puerta.

Las historias de 5 puntos o más mantienen las tres fases separadas.

### 3. Sprints de una semana

Cuatro sprints: del 20/10 al 15/11. Un sprint corto obliga a partir el trabajo en pedazos que
entran, y el error de planificación se descubre en días y no en semanas.

### 4. Revisión en el día

**Un PR abierto antes de las 18:00 se revisa ese mismo día.** Si a las 20:00 sigue sin
revisar, el Agile Enabler lo aprueba o consigue quién. Con sprints de una semana, un PR
parado dos días se come el 40% del sprint.

### 5. La prioridad es el camino crítico

Hay una cadena de historias que depende una de otra y no se puede paralelizar:

```
#8 crear sprint → #9 asignar historias → #10 cerrar sprint → #16 métricas
     → #18 reportes → #19 exportar a PDF
```

Son 31 de los 72 puntos **en serie**. Si una se atrasa, se atrasa todo lo que viene detrás, y
el dashboard, los reportes y el informe final quedan afuera. **Esa cadena tiene prioridad
sobre cualquier otra historia**, aunque las otras parezcan más urgentes.

## El plan, semana por semana

### Semana 1 — del 20 al 24 de octubre (18 SP)

| # | Historia | SP | Nota |
|---|---|---|---|
| 12 | Estimar con Story Points | 2 | le faltan 3 tareas |
| 24 | Interfaz de consola | 3 | le faltan 8 tareas |
| 8 | **Crear un sprint** | 3 | **arranca el camino crítico, va primero** |
| 9 | **Asignar historias a un sprint** | 5 | empieza apenas la #8 esté en `dev` |
| 7 | Priorizar el Product Backlog | 2 | borrador de spec ya escrito |
| 5 | Consultar el estado de un proyecto | 3 | |

### Semana 2 — del 27 al 31 de octubre (18 SP)

| # | Historia | SP | Nota |
|---|---|---|---|
| 10 | **Cerrar un sprint** | 5 | **camino crítico** |
| 14 | Registrar el esfuerzo real | 3 | la necesitan las métricas |
| 15 | Gestionar defectos | 5 | |
| 13a | Planning Poker: votación oculta y revelado | 5 | **la #13 se parte en dos** |

### Semana 3 — del 3 al 7 de noviembre (18 SP)

| # | Historia | SP | Nota |
|---|---|---|---|
| 16 | **Calcular las métricas** | 8 | **camino crítico**, necesita #10, #12 y #14 |
| 13b | Planning Poker: rondas hasta acordar | 8 | |
| 11 | Consultar sprints anteriores | 2 | |

### Semana 4 — del 10 al 14 de noviembre (18 SP)

| # | Historia | SP | Nota |
|---|---|---|---|
| 17 | Mostrar el dashboard | 8 | necesita #16 |
| 18 | **Generar el reporte** | 5 | **camino crítico** |
| 19 | Exportar el reporte a PDF | 5 | necesita #18, se hace al final de la semana |

## Lo que este plan NO incluye

Los entregables finales que no son historias del backlog y hay que escribir igual:

- Informe de métricas y cobertura
- Documentación técnica y manual breve de usuario
- Actas de las retrospectivas
- La presentación final

**Eso es trabajo real que no está en los 72 puntos.** Hay que hacerlo en paralelo, no en la
última semana: el informe de métricas se puede ir armando desde ya con
`scripts/metricas.sh`, y el manual de usuario crece con cada historia de la consola.

## Lo que hace falta de cada uno

72 puntos a 3,5 horas por punto son **252 horas**: 63 por persona, **16 horas por semana**.
Eso no es sostenible con otras materias encima.

Con los cambios de proceso el objetivo es bajar a **2 horas por punto**, que da 144 horas en
total: **9 horas por semana por persona**. Eso sí es sostenible, y es la apuesta de este
plan.

**Si en la semana 1 no bajamos de 3,5 horas por punto, el plan no cierra** y hay que hablar
con los profesores antes del 27/10, no en noviembre. La semana 1 es la prueba.

## Riesgos

| Riesgo | Qué hacemos |
|---|---|
| El camino crítico se atrasa | tiene prioridad absoluta: si alguien se libera, va ahí |
| Un PR queda parado | revisión en el día, y el Agile Enabler destraba a las 20:00 |
| La #13 (13 SP) no entra | se parte en dos antes de empezarla, 5 + 8 |
| Alguien se enferma o tiene parcial | el plan no tiene colchón: hay que avisar el mismo día |
| Bajar la intensidad del proceso empeora la calidad | la Definición de Terminado no cambia: tests, revisión cruzada y cobertura siguen igual |
