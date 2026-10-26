# Manual de usuario

Guía breve para usar el sistema de seguimiento y medición de proyectos. Crece a medida que
se terminan las historias: hoy cubre la **interfaz web** (historia #77) y la **interfaz de
consola** (historia #24). Las pantallas de la #78 se suman cuando estén en `dev`.

Las dos interfaces hacen lo mismo con los proyectos, los integrantes y el backlog, pero **no
comparten los datos**: la web los guarda en un archivo y los encuentra al volver a abrirla;
la consola los pierde al cerrarse.

## Desde el navegador

### Abrir el sistema

En una terminal, dentro de la carpeta del proyecto, escribí:

```bash
go run ./cmd/web
```

El sistema muestra la dirección donde se abre, **http://127.0.0.1:8080**. Copiala en el
navegador. No hace falta internet. Para cerrarlo, volvé a la terminal y apretá `Ctrl+C`.

- **Si dice que el puerto ya lo está usando otro programa**, abrilo en otro puerto con
  `go run ./cmd/web -puerto 8081` y usá la dirección que muestra.
- **Los datos se guardan solos** después de cada cambio, en el archivo `datos-ics.json` de
  esa carpeta. Al volver a abrir el sistema está todo como lo dejaste. Para usar otro archivo,
  por ejemplo uno para una demo: `go run ./cmd/web -datos demo.json`.
- **Si el archivo de datos está dañado**, el sistema avisa y no arranca, y no lo toca. Los
  datos se pueden recuperar arreglando el archivo, o se arranca con otro usando `-datos`.
- **Para abrirlo desde un celular** en la misma red, agregá `-red`: el sistema muestra
  también la dirección para el celular. No tiene usuarios ni contraseñas, así que usalo solo
  en una red de confianza.

### Cómo se usa

- **La barra de arriba** lleva siempre a la lista de proyectos. Dentro de un proyecto,
  muestra su nombre y sus tres partes: **Datos**, **Integrantes** y **Backlog**.
- **Cada cambio se confirma** con un mensaje en verde arriba de la página. Recargar la página
  no repite el cambio.
- **Si algo está mal**, el formulario vuelve con lo que escribiste y el motivo en rojo,
  debajo del campo que hay que corregir. Corregilo y volvé a enviar.
- **Antes de quitar algo**, el sistema pregunta. Con «Cancelar» no cambia nada.
- Todo se puede hacer **con el teclado**: Tab para pasar de un campo a otro y Enter para
  enviar. En Safari, para pasar también por los enlaces, se usa Option+Tab.

### Proyectos

La primera pantalla es la lista de proyectos. Con **Nuevo proyecto** (o, la primera vez,
**Crear el primer proyecto**) se carga uno: nombre y fecha de inicio son obligatorios; la
descripción y la fecha de fin, no.

![Lista de proyectos](manual/proyectos.png)

Si falta un dato o las fechas no cierran (la de fin antes que la de inicio), el formulario
lo marca en el campo:

![Un proyecto sin nombre](manual/proyecto-con-error.png)

Al hacer clic en el nombre de un proyecto se ven sus datos. Con **Modificar proyecto** se
cambian. Si borrás la descripción o la fecha de fin, el proyecto queda sin ellas.

![Datos de un proyecto](manual/proyecto.png)

### Integrantes

En **Integrantes** se ve quién trabaja en el proyecto y con qué rol. Con **Agregar
integrante** se carga uno, con su nombre y apellido y su rol, elegido de la lista. Un
proyecto tiene un solo Agile Enabler, y no se puede cargar dos veces a la misma persona: si
pasa, el formulario lo avisa en el campo.

![Integrantes de un proyecto](manual/integrantes.png)

Con **Quitar** se saca a un integrante. Antes, el sistema pide confirmarlo:

![Confirmar antes de quitar](manual/confirmar.png)

### Backlog e historias

En **Backlog** están las historias del proyecto, con su prioridad, su estado y sus story
points. Con **Cargar historia** se agrega una: título, descripción, prioridad y al menos un
criterio de aceptación, uno por renglón. La historia empieza Pendiente y sin estimar.

![Product Backlog](manual/backlog.png)

Al hacer clic en una historia se ven todos sus datos. Desde ahí:

- **Modificar historia** cambia el título, la descripción, la prioridad o los criterios.
- **Estimar** (o **Reestimar**) le asigna story points, elegidos de la escala 1, 2, 3, 5, 8,
  13 y 21.
- **Quitar estimación** la deja sin estimar, después de confirmarlo.

![Una historia](manual/historia.png)

### Si algo sale mal

| Lo que ves | Qué hacer |
|---|---|
| Un mensaje en rojo debajo de un campo | corregir ese campo y volver a enviar |
| «No se encontró el proyecto» o «la historia» | volver a la lista con el botón y elegirlo de ahí |
| «El cambio no quedó guardado» | el cambio sigue cargado y se guarda con el próximo que funcione; revisá que se pueda escribir en la carpeta de los datos |
| Al abrir, «el puerto ya lo está usando otro programa» | abrirlo con `-puerto 8081` |
| Al abrir, «los datos guardados no se pueden leer» | arreglar el archivo de datos o arrancar con otro usando `-datos` |

Las pantallas de sprints, esfuerzo, defectos, Planning Poker y reportes se suman a este
manual cuando estén terminadas (historia #78).

## Desde la consola

### Cómo se abre

Hace falta Go instalado. Desde la carpeta del proyecto:

```bash
go run ./cmd/app
```

El programa muestra su versión y las operaciones disponibles, y queda esperando:

```
Seguimiento y Medición ICS 0.1.0
Escribí el nombre de una operación. Operaciones disponibles: crear-proyecto, agregar-integrante, listar-integrantes, cargar-historia, listar-backlog, estimar, salir
>
```

**Los datos viven mientras el programa está abierto.** Al salir se pierden: la consola no
guarda nada en disco.

### Cómo se usa

Se escribe el **nombre de la operación** y se aprieta Enter. Después el programa pide cada
dato en su propia línea, con una etiqueta que dice qué espera. Cuando termina, vuelve a
mostrar `>` para la operación siguiente.

- Una línea vacía en lugar de una operación no hace nada.
- Si la operación no existe, el programa lo dice y lista las disponibles.
- Los proyectos y las historias se indican por su **identificador**, el número que el
  programa muestra al crearlos. El primer proyecto es el 1, y la primera historia de cada
  proyecto también es la 1.
- Las fechas se escriben **dd/mm/aaaa**, por ejemplo `01/10/2026`.

### Operaciones

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

### Un ejemplo completo

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

### Cuando algo no se puede hacer

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
