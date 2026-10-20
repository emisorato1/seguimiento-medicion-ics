# Manual de usuario

Guía breve para usar el sistema de seguimiento y medición de proyectos. Crece a medida que
se terminan las historias: hoy cubre la **interfaz de consola** (historia #24). La interfaz
web (#77 y #78) se suma cuando esté en `dev`.

## Cómo se abre

Hace falta Go instalado. Desde la carpeta del proyecto:

```bash
go run ./cmd/app
```

El programa muestra su versión y las operaciones disponibles, y queda esperando:

```
Seguimiento y Medicion ICS 0.1.0
Escribí el nombre de una operación. Operaciones disponibles: crear-proyecto, agregar-integrante, listar-integrantes, cargar-historia, listar-backlog, estimar, salir
>
```

**Los datos viven mientras el programa está abierto.** Al salir se pierden: la consola no
guarda nada en disco.

## Cómo se usa

Se escribe el **nombre de la operación** y se aprieta Enter. Después el programa pide cada
dato en su propia línea, con una etiqueta que dice qué espera. Cuando termina, vuelve a
mostrar `>` para la operación siguiente.

- Una línea vacía en lugar de una operación no hace nada.
- Si la operación no existe, el programa lo dice y lista las disponibles.
- Los proyectos y las historias se indican por su **identificador**, el número que el
  programa muestra al crearlos. El primer proyecto es el 1, y la primera historia de cada
  proyecto también es la 1.
- Las fechas se escriben **dd/mm/aaaa**, por ejemplo `01/10/2026`.

## Operaciones

| Operación | Qué hace | Qué pide |
|---|---|---|
| `crear-proyecto` | crea un proyecto y muestra su identificador | nombre, descripción (opcional), fecha de inicio y fecha de finalización (opcional) |
| `agregar-integrante` | suma una persona al equipo del proyecto | proyecto, nombre y apellido, rol (`Agile Enabler` o `Product Builder`) |
| `listar-integrantes` | muestra el equipo del proyecto con el rol de cada uno | proyecto |
| `cargar-historia` | agrega una historia al Product Backlog y muestra su identificador | proyecto, título, descripción, prioridad (`Alta`, `Media` o `Baja`) y los criterios de aceptación |
| `listar-backlog` | muestra las historias del backlog | proyecto |
| `estimar` | asigna story points a una historia | proyecto, historia y story points (1, 2, 3, 5, 8, 13 o 21) |
| `salir` | cierra el programa | nada |

Para los datos opcionales, se deja la línea vacía. Los **criterios de aceptación** se
escriben uno por línea, y una línea vacía termina la lista. Los roles y las prioridades se
pueden escribir con mayúsculas o minúsculas.

## Un ejemplo completo

Lo que escribe el usuario va después de cada etiqueta:

```
> crear-proyecto
Nombre: Sistema de métricas
Descripción (opcional):
Fecha de inicio (dd/mm/aaaa): 01/10/2026
Fecha de finalización (dd/mm/aaaa, opcional):
Proyecto creado con identificador 1
> agregar-integrante
Proyecto (identificador): 1
Nombre y apellido: Juliana Bustos
Rol (Agile Enabler o Product Builder): Product Builder
Integrante agregado al proyecto 1
> cargar-historia
Proyecto (identificador): 1
Título: Registrar esfuerzo
Descripción: Cargar las horas trabajadas en una historia
Prioridad (Alta, Media o Baja): Alta
Criterios de aceptación (uno por línea, línea vacía para terminar):
Se cargan las horas de una historia
No se aceptan horas negativas

Historia cargada con identificador 1
> estimar
Proyecto (identificador): 1
Historia (identificador): 1
Story points: 5
Historia 1 estimada en 5 story points
> listar-backlog
Proyecto (identificador): 1
- #1 Registrar esfuerzo (Alta, Pendiente, 5 story points)
> salir
```

Una historia que todavía no se estimó aparece como **sin estimar**.

## Cuando algo no se puede hacer

El programa nunca se cierra por un error: avisa y vuelve a pedir una operación.

- **Un dato mal escrito** (una fecha que no es dd/mm/aaaa, letras donde va un número, o
  estimar sin indicar la historia): el programa dice qué dato está mal y **la operación no se
  ejecuta**. Hay que volver a escribir la operación desde el principio.
- **Una regla del sistema que no se cumple** (un proyecto sin nombre, un rol o una prioridad
  que no existen, un integrante repetido, un valor que no está en la escala de story points,
  un proyecto o una historia que no existen, estimar una historia terminada): el programa
  muestra el motivo, por ejemplo
  `los story points tienen que ser uno de estos valores: 1, 2, 3, 5, 8, 13 y 21`, y no cambia
  nada.
