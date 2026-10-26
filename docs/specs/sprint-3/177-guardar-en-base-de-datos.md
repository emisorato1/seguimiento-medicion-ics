# Spec: Guardar los datos en una base de datos

- **Historia:** #177
- **Sprint:** 3
- **Responsable:** Emiliano Sorato
- **Estado:** Puerta 1, Puerta 2 y Puerta 3 aprobadas por el responsable el 2026-10-26 en la issue #177 (https://github.com/emisorato1/seguimiento-medicion-ics/issues/177#issuecomment-6429388256), como excepción — **pendientes de que un compañero las ratifique**; ningún PR de la historia entra a `dev` antes de eso
- **Depende de:** #77 (guardar los datos de la web), y de que cada historia de dominio ya
  haya sumado sus datos a lo que se guarda (#8 a #15, #13, #73)
- **La necesitan:** nadie todavía
- **Origen:** decisión de Emiliano Sorato como integrante del equipo, el 26/10. Cambia dos
  reglas de `AGENTS.md` (la de cómo se guardan los datos y la de no usar dependencias
  externas); el cambio lo aplica Emiliano Sorato en un commit aparte, antes del merge de la
  implementación.
- **Acuerdo del equipo para la dependencia:** Maximiliano Eula dio el ok en persona a
  Emiliano el 26/10 (según lo que registró Emiliano).
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de las decisiones
  de Emiliano y de la spec de la #77. El responsable lo revisa, lo corrige si hace falta y lo
  presenta a las puertas.

## Fase 1 — Requisitos

### Objetivo

Que lo que el equipo carga desde la web quede guardado en una base de datos en lugar de un
archivo de texto que se reescribe entero, sin que quien usa el sistema note otra diferencia
que esta: un cambio queda guardado completo o no queda, nunca a medias. Los datos se pueden
seguir llevando a un archivo y traer desde uno, para mover el proyecto de una computadora a
otra, para cargar los datos de demo y para traer lo guardado con la versión anterior.

### Qué queda afuera

- **La consola (#24)**: sigue con los datos en memoria, como hoy.
- **Varias personas en computadoras distintas sobre los mismos datos**: el sistema sigue
  siendo uno solo, abierto en una computadora; las pestañas y el celular por la red (#77)
  funcionan como hoy.
- **Deshacer cambios o ver un historial** de lo que se modificó.
- **Copias de seguridad automáticas.** Exportar a un archivo sirve para hacer una a mano.
- **Cambiar el formato del archivo** que se exporta: es el mismo que usa la web hoy.
- **Usuarios y contraseñas**: siguen sin existir (#77).

### Entradas

| Dato | Obligatorio | Descripción |
|---|---|---|
| Dónde está la base de datos | no | por defecto, en la carpeta desde donde se abre el sistema, como hoy el archivo de datos |
| Archivo para importar | no | un archivo con los datos en el formato de hoy, para traerlos a la base al abrir el sistema |
| Archivo para exportar | no | dónde escribir todos los datos, en el formato de hoy |

### Salidas

- Los datos de cada operación aceptada, guardados en la base.
- Al abrir el sistema, todos los datos guardados, como estaban al cerrarlo.
- Al exportar, un archivo con todos los datos.
- Los avisos en castellano cuando algo no se puede guardar, leer, importar o exportar.

### Reglas de negocio

1. **[RN-1]** Cuando una operación de la web se acepta, el sistema debe guardarla en la base
   de datos antes de confirmarla con su mensaje, igual que hoy (RN-21 de la #77).
2. **[RN-2]** Si un cambio no se puede guardar, el sistema no debe dejar guardada ninguna
   parte de él: la base queda como estaba antes de la operación. Avisa que no quedó guardado,
   y el cambio se guarda junto con la próxima operación que funcione, como hoy (RN-25 de la
   #77).
3. **[RN-3]** Cuando se abre, el sistema debe recuperar todos los datos guardados:
   proyectos, integrantes, historias con sus criterios, sprints con lo asignado y su cierre,
   esfuerzo, defectos y rondas de Planning Poker con sus votos y su acuerdo, y seguir la
   numeración donde había quedado (RN-22 de la #77).
4. **[RN-4]** Cuando recupera los datos, el sistema debe controlar cada uno con las mismas
   reglas que al cargarlo desde la web. Si alguno no las cumple, debe avisar en castellano
   qué dato está mal y no abrir, sin cambiar nada de la base (RN-24 de la #77).
5. **[RN-5]** Mientras no haya una base de datos, como la primera vez, el sistema debe abrir
   sin datos y crearla.
6. **[RN-6]** Mientras una ronda de Planning Poker no esté revelada, sus votos deben seguir
   ocultos en la web después de cerrar y volver a abrir el sistema (RN-4 de la #78).
7. **[RN-7]** Cuando se pide exportar, el sistema debe escribir todos los datos en un archivo
   con el formato que usa hoy la web, y terminar sin abrir la página.
8. **[RN-8]** Cuando se pide importar un archivo, el sistema debe controlar todos sus datos
   con las mismas reglas que al cargarlos desde la web y, solo si los cumplen todos,
   guardarlos en la base y abrir con ellos.
9. **[RN-9]** Si el archivo para importar no existe, no se puede leer o tiene algún dato que
   no cumple las reglas, el sistema debe avisar en castellano qué pasó y no cambiar nada de
   la base.
10. **[RN-10]** Si se pide importar y la base ya tiene datos, el sistema debe rechazarlo y
    avisar, sin cambiar nada. **Decidido por Emiliano Sorato el 26/10.**
11. **[RN-11]** Exportar los datos e importarlos en una base nueva debe dar exactamente los
    mismos datos.
12. **[RN-12]** El sistema debe funcionar igual en Linux, en macOS y en Windows, sin instalar
    nada aparte, como hoy.
13. **[RN-13]** Si la base fue creada por una versión más nueva del sistema, que la guarda de
    otra forma, el sistema debe avisar que no la puede abrir y no tocarla.
14. **[RN-14]** Cuando el sistema abre con la base vacía y encuentra en la carpeta el archivo
    de datos de la versión anterior (`datos-ics.json`), debe importarlo solo, con las mismas
    reglas que al importar a pedido (RN-8 y RN-9), y avisar que lo importó y desde qué
    archivo. **Decidido por Emiliano Sorato el 26/10.**
15. **[RN-15]** Cuando el archivo de la versión anterior se importa bien, el sistema debe
    renombrarlo a `datos-ics.json.importado` y avisar que lo renombró, para que no se vuelva a
    importar: aunque después se borre todo, al volver a abrir no reaparecen los datos viejos.
    Si la importación falla, el archivo no se toca, la base queda vacía y el sistema avisa
    qué pasó y no abre (RN-9). **Decidido por Emiliano Sorato el 26/10.**

### Casos límite

- La base nueva, sin ningún proyecto: el sistema abre vacío, como la primera vez.
- Un proyecto sin integrantes, sin historias, sin sprints ni nada más.
- Textos con tildes, eñes, comillas, signos de HTML y saltos de línea, como los criterios de
  aceptación de varios renglones: vuelven exactamente iguales.
- Horas con decimales (2,5): vuelven exactas, sin redondeos.
- Fechas: vuelven como el mismo día, sin correrse por la hora ni por la zona horaria (como el
  defecto #59).
- Una historia sin estimar, un sprint abierto sin cierre, un defecto sin resolver, una ronda
  sin revelar con votos y una ronda acordada.
- Un integrante que se quitó del proyecto pero tiene horas cargadas: se recupera sin error,
  porque es algo que el uso normal produce (Fase 2 de la #14).
- Después de borrar el último elemento (por ejemplo, quitar el único integrante), al volver
  a abrir no reaparece.
- La numeración: después de reabrir, la historia, el sprint y el defecto siguientes reciben el
  número que sigue, aunque se hayan quitado elementos.
- El archivo de datos de la versión anterior (`datos-ics.json`) en la carpeta y la base
  vacía: se importa solo, se renombra a `datos-ics.json.importado` y se avisa (RN-14, RN-15).
  Con la base ya con datos, se lo ignora.
- Después de importarlo, se borra todo y se vuelve a abrir: no se vuelve a importar, porque
  el archivo ya está renombrado (RN-15).
- El archivo de la versión anterior con un dato que no cumple las reglas: se avisa qué dato
  está mal y el sistema no abre; la base queda vacía y el archivo, sin tocar (RN-15).
- Ya existe un `datos-ics.json.importado` de una importación anterior: el renombre no puede
  pisarlo, así que se avisa y no se abre, sin cambiar nada (ver "Condiciones de error").

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| La base no se puede abrir o está dañada | avisa en castellano y no abre, sin tocarla |
| La base tiene un dato que no cumple las reglas | avisa qué dato está mal y no abre, sin tocarla |
| La base es de una versión más nueva del sistema | avisa que no la puede abrir y no la toca |
| Un cambio no se puede guardar (por ejemplo, la carpeta no permite escribir) | avisa "No quedó guardado", la base queda como antes y el cambio se guarda con la próxima operación que funcione |
| El archivo para importar no existe, no se puede leer o no cumple las reglas | avisa qué pasó y no cambia nada |
| Se pide importar y la base ya tiene datos | avisa que no importa sobre datos existentes y no cambia nada |
| El archivo de la versión anterior, importado solo al abrir, no cumple las reglas | avisa qué dato está mal y no abre; la base queda vacía y el archivo, sin tocar |
| El archivo de la versión anterior se importó, pero no se puede renombrar (por ejemplo, ya existe un `datos-ics.json.importado`) | avisa en castellano y no abre, y la base queda vacía, para que al resolverlo se importe una sola vez |
| El archivo para exportar no se puede escribir | avisa en castellano y termina sin escribir nada a medias |

### Criterios de aceptación

- [ ] Todo lo que se carga desde la web sigue igual al cerrar y volver a abrir el sistema.
- [ ] Un cambio que no se pudo guardar no deja ninguna parte guardada: la base queda como
      estaba, y el cambio entra con el próximo guardado que funcione.
- [ ] Al abrir, cada dato pasa por las mismas reglas que al cargarlo desde la web; uno que no
      las cumple se avisa y el sistema no abre.
- [ ] Exportar a un archivo e importarlo en una base nueva da exactamente los mismos datos.
- [ ] Los datos de demo y lo guardado con la versión anterior se pueden importar; lo de la
      versión anterior se importa solo al abrir con la base vacía, y se avisa.
- [ ] Importar un archivo con datos que no cumplen las reglas, o sobre una base con datos, no
      cambia nada y se avisa.
- [ ] Funciona en Linux, en macOS y en Windows sin instalar nada aparte, y los tests pasan en
      los dos sistemas del CI.

### Escenarios BDD

```gherkin
Escenario: Los datos siguen al cerrar y volver a abrir
  Dado un proyecto con integrantes, historias, un sprint cerrado y otro abierto, horas,
    defectos y una ronda revelada y acordada
  Cuando cierro el sistema y lo vuelvo a abrir
  Entonces veo todos los datos exactamente como estaban
  Y la historia, el sprint y el defecto siguientes reciben el número que sigue

Escenario: Un cambio que no se pudo guardar no deja nada a medias
  Dado un proyecto guardado en la base
  Y que el guardado se corta a mitad de camino
  Cuando cargo una historia nueva
  Entonces veo el aviso "No quedó guardado"
  Y la base tiene exactamente lo que tenía antes
  Cuando hago otra operación y el guardado funciona
  Entonces la base tiene las dos

Escenario: Un dato guardado que no cumple las reglas no se carga
  Dado una base donde una historia tiene 4 story points
  Cuando abro el sistema
  Entonces veo un aviso que dice qué dato está mal
  Y el sistema no abre
  Y la base queda como estaba

Esquema del escenario: Los datos difíciles vuelven iguales
  Dado <dato>
  Cuando lo guardo, cierro el sistema y lo vuelvo a abrir
  Entonces <resultado>

  Ejemplos:
    | dato                                                         | resultado                                   |
    | el criterio "Se ve <b>el total</b> y \"las horas\"" en dos renglones | el criterio aparece igual, en dos renglones |
    | 2,5 horas de "Ana Pérez" el 27/10/2026                       | son 2,5 horas del 27/10/2026                |
    | una ronda sin revelar con el voto 13 de "Ana Pérez"          | sigue sin verse ningún voto                 |
    | "Ana Pérez" quitada del proyecto con 2 horas cargadas         | el sistema abre y las 2 horas siguen        |
    | quitar el único integrante                                   | el proyecto sigue sin integrantes           |

Escenario: Exportar e importar da los mismos datos
  Dado una base con un proyecto completo
  Cuando exporto los datos a un archivo
  E importo ese archivo en una base nueva
  Entonces la base nueva tiene exactamente los mismos datos

Escenario: Traer los datos de la versión anterior
  Dado el archivo de datos de demo (docs/demo/datos-demo.json) y ninguna base
  Cuando abro el sistema pidiendo importar ese archivo
  Entonces veo el proyecto "Seguimiento y Medición ICS" con sus 21 historias y sus dos sprints
  Y al cerrar y volver a abrir, sin pedir importar, siguen ahí

Esquema del escenario: Importar algo que no se puede no cambia nada
  Dado <situación>
  Cuando abro el sistema pidiendo importar <archivo>
  Entonces veo el aviso <aviso>
  Y la base queda como estaba

  Ejemplos:
    | situación             | archivo                              | aviso                                     |
    | una base vacía        | un archivo que no existe             | que no se pudo abrir el archivo           |
    | una base vacía        | un archivo con una historia de 4 SP  | qué dato está mal                         |
    | una base con datos    | los datos de demo                    | que no importa sobre datos existentes     |

Escenario: Lo guardado con la versión anterior se importa solo
  Dado el archivo "datos-ics.json" de la versión anterior con el proyecto "Sistema de metricas"
  Y la base vacía
  Cuando abro el sistema
  Entonces veo el aviso de que se importaron los datos de "datos-ics.json"
  Y que el archivo se renombró a "datos-ics.json.importado"
  Y veo el proyecto "Sistema de metricas"
  Cuando cierro el sistema y lo vuelvo a abrir
  Entonces no se vuelve a importar, y veo el proyecto "Sistema de metricas" una sola vez

Escenario: Lo de la versión anterior no vuelve aunque se borre todo
  Dado que el sistema ya importó "datos-ics.json" y lo renombró a "datos-ics.json.importado"
  Cuando quito todos los datos, cierro el sistema y lo vuelvo a abrir
  Entonces no se importa nada y veo la lista de proyectos vacía

Escenario: Si lo de la versión anterior no se puede importar, no se toca
  Dado el archivo "datos-ics.json" con una historia de 4 story points
  Y la base vacía
  Cuando abro el sistema
  Entonces veo un aviso que dice qué dato está mal
  Y el sistema no abre
  Y el archivo "datos-ics.json" queda como estaba, sin renombrar
  Y la base sigue vacía

Escenario: Abrir por primera vez
  Dado que no hay ninguna base de datos
  Cuando abro el sistema
  Entonces veo la lista de proyectos vacía
  Y al cargar un proyecto queda guardado en una base nueva

Escenario: Una base de una versión más nueva no se abre
  Dado una base creada por una versión más nueva del sistema
  Cuando abro el sistema
  Entonces veo un aviso que dice que no la puede abrir
  Y la base queda como estaba
```

### Decidido antes de la puerta

Emiliano Sorato decidió el 26/10 los puntos que el borrador dejaba abiertos:

1. **Importar sobre una base con datos se rechaza** (RN-10): importar reemplazaría todo lo
   que hay, y un error de tipeo en el comando borraría el proyecto.
2. **El archivo de la versión anterior se importa solo** si la base está vacía, y se avisa
   (RN-14). El borrador proponía solo avisar.
3. **Se importa una sola vez** (RN-15): después de importarlo bien se renombra a
   `datos-ics.json.importado`. Así, alguien que borra todo y vuelve a abrir no recibe otra
   vez los datos viejos. Si la importación falla, el archivo no se toca. Decidido por
   Emiliano Sorato el 26/10.

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: Emiliano Sorato — Fecha: 2026-10-26 — Comentarios: aprobada en un comentario de la issue #177 (https://github.com/emisorato1/seguimiento-medicion-ics/issues/177#issuecomment-6429388256), sobre el commit `4d74453`. La aprueba el responsable de la historia: es una excepción a la regla de que aprueba otra persona, y queda para que un compañero la ratifique con un comentario en la issue #177, desde su propia cuenta. Ningún PR de la historia entra a `dev` antes de esa ratificación.
> Ratificó: [pendiente]

## Fase 2 — Diseño

### Enfoque

Solo cambian `internal/store` y el arranque de `cmd/web`; **el dominio no se entera**
(decisión de Emiliano). La web ya pide guardar después de cada operación aceptada a algo que
"sabe guardar": hoy es el archivo; con esta historia es la base, con la misma forma de
guardar y de cargar, así los handlers no cambian.

El problema de fondo es que los datos de las entidades están ocultos fuera del dominio:
`internal/store` no los puede leer ni reconstruir. La solución es usar como puente **la
conversión que el dominio ya tiene** (la de la #77):

- **Al guardar**, el registro se convierte con la conversión del dominio, ese resultado se
  desarma en filas —una tabla por entidad— y las filas se escriben en la base en una sola
  transacción.
- **Al cargar**, se leen las filas, se rearma el mismo formato y se lo pasa a la conversión
  del dominio, que reconstruye todo **con las mismas validaciones de hoy** (RN-4).

Ese formato es solo un paso en memoria: **la base guarda tablas, nunca el texto**. Y es el
mismo formato del archivo que se importa y se exporta, así que importar y exportar son leer y
escribir ese archivo con lo que ya existe.

### Archivos afectados

| Archivo | Se crea o se modifica | Para qué |
|---|---|---|
| `internal/store/database.go` y su test | se crean | abrir la base, aplicar la migración, guardar y cargar con transacción |
| `internal/store/migrations/0001_inicial.sql` | se crea | el esquema inicial, versionado y embebido en el binario |
| `internal/store/file.go` | se modifica (poco) | queda para importar y exportar el archivo |
| `cmd/web/main.go` y su test | se modifica | la base por defecto, las opciones para importar y exportar, y el aviso del archivo viejo |
| `go.mod` y `go.sum` | se modifican | el driver de la base |
| `docs/manual-de-usuario.md` | se modifica | cómo se guardan los datos, importar y exportar |
| `docs/documentacion-tecnica.md` | se modifica | el esquema y cómo se suma un dato nuevo |
| `AGENTS.md` | se modifica: lo aplica Emiliano Sorato en un commit aparte, antes del merge de la implementación | la fila de cómo se guardan los datos y la regla de las dependencias, ya acordadas |

### Decisiones

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| **D1** SQLite con el driver `modernc.org/sqlite`, por `database/sql` | un driver con CGO (`mattn/go-sqlite3`); otro motor con servidor (PostgreSQL) | es Go puro: compila y corre en el CI de Windows y en las tres plataformas del CD sin un compilador de C. Un motor con servidor obliga a instalar algo aparte (RN-12). Decidido por Emiliano |
| **D2** Una tabla por entidad, y tablas aparte para lo que una entidad tiene varias veces: los criterios de cada historia, las historias asignadas y completadas de cada sprint con sus puntos al cerrar, los participantes y los votos de cada ronda. **Los contadores de la numeración se guardan tal como los trae la conversión**: el próximo número de proyecto, en una tabla del registro con una sola fila; el próximo número de historia, el último sprint y el último defecto de cada proyecto, como columnas de la tabla de proyectos. **Nunca se deducen del número más alto guardado** | una tabla con el texto del registro entero; deducir el próximo número del máximo guardado | se pidió explícitamente: guardar el texto en una columna no es una base de datos de verdad (decidido por Emiliano). Con D4, que borra y vuelve a insertar, deducir el contador del máximo reutilizaría el número del último elemento quitado; la conversión ya trae los contadores (RN-22 de la #77), así que se guardan como vienen. Las rondas se numeran por historia y el dominio controla al cargar que vayan seguidas: no llevan contador |
| **D3** El puente es la conversión del dominio (Enfoque) | que el dominio exporte constructores y lectores para cada dato; que `internal/store` reimplemente las validaciones | lo primero cambia el contrato del dominio, que esta historia no toca; lo segundo duplica las reglas y las deja divergir. Así la carga valida exactamente lo mismo que hoy, por el mismo código |
| **D4** Guardar reescribe todas las tablas dentro de **una transacción**: borra y vuelve a insertar | actualizar solo lo que cambió | la web pide guardar el registro entero (D13 de la #77), y saber qué cambió exigiría que el dominio lo informe. Los datos son pocos. La transacción es lo que garantiza todo o nada (RN-2) |
| **D5** Si la conversión trae un dato que la base no tiene dónde guardar, **el guardado falla** con "No quedó guardado" | ignorar el dato | es lo que hoy obliga a cada historia a sumar su dato a la conversión (AGENTS.md): ahora también a una tabla, con una migración nueva. Fallar en voz alta evita perder datos en silencio, y el test de ida y vuelta lo detecta antes del merge |
| **D6** El esquema se crea con **migraciones numeradas**, embebidas en el binario; la base guarda en qué número está, y al abrir se aplican las que faltan, en una transacción | crear las tablas "si no existen" | permite cambiar el esquema después sin perder datos. Una base con un número mayor que el que conoce el sistema no se abre (RN-13) |
| **D7** Una sola conexión abierta, y el guardado sigue usando el candado de la web | varias conexiones | todos los cambios ya pasan de a uno por el candado (#77); con una conexión no hay bloqueos entre conexiones del propio sistema |
| **D8** La base por defecto es `datos-ics.db` en la carpeta desde donde se abre, y la opción para elegir otra es la de hoy (`-datos`). Importar y exportar son dos opciones nuevas del arranque. El archivo de la versión anterior se busca con su nombre de siempre, `datos-ics.json`, en la misma carpeta que la base, y se renombra en esa carpeta (RN-14, RN-15) | elegir el formato según la extensión del archivo | el JSON queda solo para importar y exportar (decisión de Emiliano); una sola forma de guardar es más simple de probar y de explicar. Buscarlo junto a la base es donde lo dejaba la versión anterior |
| **D9** Para probar que la transacción no deja datos a medias, el test hace fallar la base a mitad del guardado con una regla de la propia base (un disparador que rechaza una inserción) | agregarle al código un punto para inyectar fallas | no agrega código que solo usan los tests, y prueba la transacción real |
| **D10** El orden al importar solo el archivo viejo es: importar y guardar en la base, después renombrar el archivo. Si el renombre falla, se vacía la base otra vez antes de avisar | renombrar primero y después importar | renombrando primero, una importación que falla dejaría el archivo renombrado y no se volvería a intentar (RN-15). Vaciar la base si el renombre falla mantiene la promesa de importar una sola vez |

### Dependencias y restricciones

- **La dependencia está acordada con el equipo** (AGENTS.md: "agregar una dependencia se
  acuerda con el equipo"): Maximiliano Eula dio el ok en persona a Emiliano el 26/10 (según lo
  que registró Emiliano). El cambio de `AGENTS.md` lo aplica Emiliano Sorato en un commit
  aparte, antes del merge de la implementación.
- **No se toca** `internal/domain` (decisión de Emiliano), `internal/web` (los handlers ya
  guardan a través de algo que sabe guardar) ni `cmd/app` (la consola sigue en memoria).
- **El CI baja el driver** al compilar (`go mod download`); hoy no baja nada.
- Los tests de errores de archivos se arman con el lugar de las cosas y no con permisos de
  Unix, que Windows ignora (defecto #108).
- **El escenario "Traer los datos de la versión anterior" nombra los datos de demo**
  (`docs/demo/datos-demo.json`), que están en la rama `datos-demo` y todavía no en `dev`. La
  tarea 10 no depende de esa rama: **su test usa un archivo de prueba propio** con los mismos
  tipos de datos. Si `datos-demo` está en `dev` al implementar, se importa además a mano y se
  anota en la Fase 4.

> **Puerta 2** — ¿el diseño es viable y coherente con el resto del proyecto?
> Aprobó: Emiliano Sorato — Fecha: 2026-10-26 — Comentarios: aprobada en un comentario de la issue #177 (https://github.com/emisorato1/seguimiento-medicion-ics/issues/177#issuecomment-6429388256), sobre el commit `4d74453`. La aprueba el responsable de la historia: es una excepción a la regla de que aprueba otra persona, y queda para que un compañero la ratifique con un comentario en la issue #177, desde su propia cuenta. Ningún PR de la historia entra a `dev` antes de esa ratificación.
> Ratificó: [pendiente]

## Fase 3 — Tareas

Cada tarea es un ciclo de TDD: el test en un commit, con su salida en rojo, y la
implementación en el siguiente.

| # | Tarea | Escenario BDD | Regla | Estado |
|---|---|---|---|---|
| 1 | Abrir una base nueva crea el esquema; volver a abrirla no lo vuelve a crear | Abrir por primera vez | RN-5 | pendiente |
| 2 | Guardar y cargar un registro vacío y uno con un proyecto | Abrir por primera vez | RN-1, RN-3 | pendiente |
| 3 | Ida y vuelta completa: un registro con todas las entidades vuelve idéntico, comparado con la conversión del dominio | Los datos siguen al cerrar y volver a abrir | RN-3 | pendiente |
| 4 | Los datos difíciles: signos y renglones, decimales, fechas, ronda sin revelar, integrante quitado con horas, borrar el último, la numeración | Los datos difíciles vuelven iguales | RN-3, RN-6 | pendiente |
| 5 | Un guardado que falla a mitad deja la base como estaba, y el siguiente que funciona guarda todo | Un cambio que no se pudo guardar no deja nada a medias | RN-2 | pendiente |
| 6 | Un dato de la conversión que la base no conoce hace fallar el guardado | — (D5) | RN-2 | pendiente |
| 7 | Un dato guardado que no cumple las reglas, una base dañada y una de versión más nueva: avisan y no se tocan | Un dato guardado que no cumple las reglas no se carga / Una base de una versión más nueva no se abre | RN-4, RN-13 | pendiente |
| 8 | Una carpeta donde no se puede escribir: "No quedó guardado" | — (Condiciones de error) | RN-2 | pendiente |
| 9 | El arranque usa la base por defecto | Abrir por primera vez | RN-5 | pendiente |
| 10 | Exportar escribe el archivo y termina; importar en una base nueva; exportar e importar dan lo mismo | Exportar e importar da los mismos datos / Traer los datos de la versión anterior | RN-7, RN-8, RN-11 | pendiente |
| 11 | Importar algo que no se puede: archivo que no existe, con un dato inválido, o sobre una base con datos | Importar algo que no se puede no cambia nada | RN-9, RN-10 | pendiente |
| 12 | El archivo de la versión anterior se importa solo con la base vacía, se renombra y se avisa; no vuelve aunque se borre todo; uno inválido no se toca y no abre; un renombre que no se puede hacer no abre | Lo guardado con la versión anterior se importa solo / Lo de la versión anterior no vuelve aunque se borre todo / Si lo de la versión anterior no se puede importar, no se toca | RN-14, RN-15, RN-9 | pendiente |
| 13 | La web guarda en la base: el escenario de la #77 y la #78 de cerrar y volver a abrir, con la base | Los datos siguen al cerrar y volver a abrir | RN-1, RN-3 | pendiente |
| 14 | El manual de usuario y la documentación técnica | — | — | pendiente |

Los tests corren en los dos sistemas del CI, Linux y Windows (RN-12): ninguno depende de
permisos de Unix ni de rutas con `/`.

> **Puerta 3** — ¿las tareas cubren todos los criterios de aceptación?
> Aprobó: Emiliano Sorato — Fecha: 2026-10-26 — Comentarios: aprobada en un comentario de la issue #177 (https://github.com/emisorato1/seguimiento-medicion-ics/issues/177#issuecomment-6429388256), sobre el commit `4d74453`. La aprueba el responsable de la historia: es una excepción a la regla de que aprueba otra persona, y queda para que un compañero la ratifique con un comentario en la issue #177, desde su propia cuenta. Ningún PR de la historia entra a `dev` antes de esa ratificación.
> Ratificó: [pendiente]

## Fase 4 — Implementación

Se completa al implementar.

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #177 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | `internal/store/database_test.go`, `cmd/web/main_test.go` y los de persistencia de `internal/web` |
| Código Go | `internal/store/database.go`, `internal/store/migrations/`, `cmd/web/main.go` |
