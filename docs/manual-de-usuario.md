# Manual de usuario

Seguimiento y Medición ICS sirve para cargar los proyectos del equipo, sus integrantes y su
Product Backlog, y estimar las historias en story points. Se puede usar de dos maneras:
desde el navegador o desde la consola.

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

Esta parte la escribe Maximiliano Eula, con la interfaz de consola (historia #24).
