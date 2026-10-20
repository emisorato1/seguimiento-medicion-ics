# Spec: Exportar el reporte a PDF

- **Historia:** #19
- **Sprint:** 4
- **Responsable:** Emiliano Sorato
- **Estado:** Borrador completo — pendiente de las tres puertas
- **Redacción inicial:** borrador preparado con asistencia de IA a partir de la consigna y de los criterios de aceptación de la issue. El responsable lo revisa, lo corrige si hace falta y lo presenta a las puertas.
- **Depende de:** #18 (el reporte ya armado), y por ella de la #16. Ninguna está en `dev`
  todavía: se trabaja contra el contrato del dominio
- **La necesitan:** la #78, que ofrece el PDF para descargar desde la web. Es el último
  eslabón de la cadena, y el PDF es lo que se muestra en la review final

## Fase 1 — Requisitos

### Objetivo

Que el equipo pueda descargar desde la web un reporte de proyecto o de sprint como archivo
PDF, para entregarlo y presentarlo en la review final sin depender de tener el programa
abierto.

### Qué queda afuera

- **Armar el reporte y calcular sus números**: son la #18 y la #16. El PDF escribe lo que el
  reporte ya trae, sin calcular ni reorganizar nada.
- **Gráficos, tablas dibujadas, colores, logos e imágenes**: el PDF es un documento de texto.
  Los gráficos son del dashboard (#17) y no pasan al PDF (punto 1 de "Para decidir en la
  puerta").
- Exportar el dashboard (#17) a PDF: la issue pide el reporte.
- Elegir la tipografía, el tamaño de página o los márgenes: son fijos.
- Exportar a otros formatos, abrir el archivo, imprimirlo, mandarlo por correo, protegerlo con
  contraseña o firmarlo.
- **El botón de descarga**, que es de la #78 (su tarea 21).
- **Exportar desde la consola**: la consola queda con lo que especifica la #24, y el PDF se
  usa desde la web (punto 4 de "Para decidir en la puerta").

### Entradas

| Dato | Obligatorio | Descripción |
|---|---|---|
| Reporte | sí | un reporte ya generado, de un proyecto o de un sprint |

### Salidas

El reporte como archivo PDF, que la web ofrece para descargar (#78). Lo que tiene que tener
el archivo:

| Qué | Cómo |
|---|---|
| Formato | un PDF que se abre en cualquier lector común (Vista Previa, Acrobat Reader, un navegador) sin instalar nada |
| Páginas | A4 vertical, con márgenes; tantas como haga falta |
| Primera página | arriba, el título del reporte con letra más grande, y debajo "Generado el" con la fecha |
| Secciones | cada encabezado en negrita y debajo sus líneas, en el orden del reporte |
| Pie | en cada página, "Página N de M" |
| Texto | se puede seleccionar y buscar en el lector: no es una imagen |

### Reglas de negocio

1. **[RN-1]** Cuando se exporta un reporte, el sistema debe producir un archivo PDF que se
   pueda abrir con un lector común, sin instalar nada adicional.
2. **[RN-2]** El sistema debe escribir en el archivo el título del reporte, su día de
   generación y cada sección con su encabezado y todas sus líneas, con el mismo texto y en el
   mismo orden que el reporte.
3. **[RN-3]** El sistema no debe calcular, agregar, quitar ni reordenar nada del contenido
   del reporte. Lo único que agrega son las palabras "Generado el" delante de la fecha y la
   numeración de las páginas.
4. **[RN-4]** El sistema debe usar páginas A4 verticales con márgenes, y ningún texto debe
   quedar fuera de los márgenes.
5. **[RN-5]** El sistema debe escribir el título con letra más grande que el resto, y los
   encabezados de sección en negrita.
6. **[RN-6]** Cuando una línea no entra en el ancho de la página, el sistema debe seguirla en
   el renglón siguiente, cortando entre palabras; si una palabra sola no entra, debe cortarla
   donde llega al margen.
7. **[RN-7]** Cuando el contenido no entra en una página, el sistema debe seguir en una
   página nueva. Un encabezado de sección nunca debe quedar como último renglón de una página.
8. **[RN-8]** El sistema debe poner al pie de cada página "Página N de M".
9. **[RN-9]** El sistema debe escribir el texto de modo que se pueda seleccionar y buscar en
   el lector.
10. **[RN-10]** El sistema debe mostrar bien las letras del castellano —tildes, ñ, ü, ¿ y ¡—.
    Si aparece un carácter que el archivo no puede representar, debe reemplazarlo por un
    signo de pregunta y exportar igual.
11. **[RN-11]** El sistema debe escribir tal cual los paréntesis, las barras y cualquier otro
    signo del texto, sin que el archivo deje de abrirse.
12. **[RN-12]** Si el PDF no se puede entregar entero, el sistema debe informar que no se pudo
    generar el PDF.
13. **[RN-13]** El sistema debe exportar de la misma forma un reporte de proyecto y uno de
    sprint.

### Casos límite

- Reporte que entra en una página: una sola, con "Página 1 de 1".
- Reporte que necesita tres páginas: "Página 1 de 3", "Página 2 de 3" y "Página 3 de 3".
- Encabezado que caería en el último renglón de una página: pasa a la página siguiente.
- Una palabra de 200 caracteres sin espacios: se corta en el margen, en varios renglones.
- Título muy largo: sigue en varios renglones, también con la letra grande.
- Línea vacía dentro de una sección: queda como renglón en blanco.
- Texto con "(", ")" y "\": se ve tal cual.
- Texto con un emoji o con letras de otro alfabeto: en su lugar aparece "?", y el resto de la
  línea se ve bien.
- Exportar dos veces el mismo reporte: da dos archivos iguales.

### Condiciones de error

| Situación | Qué hace el sistema |
|---|---|
| El PDF no se puede entregar entero | informa que no se pudo generar el PDF |

### Criterios de aceptación

- [ ] Exportar un reporte da un archivo PDF, que la web ofrece para descargar.
- [ ] El archivo se abre en Vista Previa y en un navegador y muestra el texto del reporte,
      que se puede seleccionar y buscar.
- [ ] El archivo tiene el título, la fecha de generación y todas las secciones con todas sus
      líneas, en el mismo orden que el reporte en pantalla, sin nada agregado ni quitado.
- [ ] El título va más grande y los encabezados en negrita.
- [ ] Una línea que no entra sigue en el renglón siguiente y nada queda fuera del margen.
- [ ] Un reporte que no entra en una página sigue en otra, cada página dice "Página N de M" y
      ningún encabezado queda solo al pie.
- [ ] Las tildes y la ñ se ven bien, y un carácter que no se puede representar sale como "?"
      sin impedir la exportación.
- [ ] Los paréntesis y las barras del texto no impiden abrir el archivo.
- [ ] Si el PDF no se puede entregar entero, se informa.
- [ ] Un reporte de proyecto y uno de sprint se exportan de la misma forma.

### Escenarios BDD

```gherkin
Escenario: Exportar el reporte de un sprint a PDF
  Dado el reporte "Reporte del sprint 1 — Seguimiento ICS", generado el 25/10/2026, con sus cuatro secciones
  Cuando lo exporto a PDF
  Entonces recibo un archivo PDF
  Y la primera pagina empieza con el titulo en letra grande y "Generado el 25/10/2026"
  Y el archivo tiene cada encabezado en negrita y cada linea del reporte, en el mismo orden

Escenario: El archivo se abre en un lector de PDF
  Dado el reporte del sprint 1 exportado a PDF
  Cuando abro el archivo con Vista Previa y con un navegador
  Entonces veo el titulo y todas las lineas del reporte
  Y puedo seleccionar y buscar el texto

Esquema del escenario: El PDF no agrega ni quita nada del reporte
  Dado <reporte> con 4 secciones y 30 lineas en total
  Cuando lo exporto a PDF
  Entonces el archivo tiene exactamente esas 4 secciones y esas 30 lineas
  Y ademas solo "Generado el" y la numeracion de las paginas

  Ejemplos:
    | reporte                  |
    | el reporte del proyecto  |
    | el reporte de un sprint  |

Escenario: Una linea larga sigue en el renglon siguiente
  Dado un reporte con una linea mas larga que el ancho de la pagina
  Cuando lo exporto a PDF
  Entonces la linea sigue en el renglon siguiente, cortada entre dos palabras
  Y ningun renglon pasa el margen

Escenario: Una palabra mas larga que el renglon se corta en el margen
  Dado un reporte con una palabra de 200 caracteres sin espacios
  Cuando lo exporto a PDF
  Entonces la palabra ocupa varios renglones, cortada donde llega al margen

Esquema del escenario: Un reporte largo ocupa varias paginas numeradas
  Dado un reporte que necesita <paginas> paginas
  Cuando lo exporto a PDF
  Entonces el archivo tiene <paginas> paginas
  Y la ultima dice <pie>

  Ejemplos:
    | paginas | pie              |
    | 1       | "Pagina 1 de 1"  |
    | 3       | "Pagina 3 de 3"  |

Escenario: Un encabezado nunca queda solo al pie de una pagina
  Dado un reporte en el que un encabezado caeria en el ultimo renglon de la pagina 1
  Cuando lo exporto a PDF
  Entonces ese encabezado aparece al principio de la pagina 2, junto con sus lineas

Esquema del escenario: Letras del castellano y caracteres que no se pueden representar
  Dado un reporte con una linea que tiene <contenido>
  Cuando lo exporto a PDF
  Entonces el archivo muestra <resultado>
  Y el resto de la linea se ve igual que en el reporte

  Ejemplos:
    | contenido                                   | resultado                        |
    | tildes, enie, dieresis y signos de apertura | esas letras como en el reporte   |
    | un emoji                                    | un signo de pregunta en su lugar |

Escenario: Los parentesis y las barras no rompen el archivo
  Dado un reporte con la linea "Desviacion (sobre lo estimado): 25,0 % \ revisar"
  Cuando lo exporto a PDF
  Entonces el archivo se abre y muestra esa linea tal cual

Escenario: Un PDF que no se puede entregar se informa
  Dado un reporte del sprint 1
  Cuando lo exporto y la entrega del archivo falla
  Entonces la exportacion se rechaza
  Y se informa que no se pudo generar el PDF
```

**Para decidir en la puerta:**

1. **Alcance: el PDF es un documento de texto.** Sin gráficos, tablas dibujadas, colores ni
   logos: solo el texto del reporte, ordenado en páginas, con dos tipos de letra (normal y
   negrita). La razón es que el proyecto no usa librerías externas, así que el archivo hay
   que escribirlo "a mano", y cada cosa que no sea texto multiplica el trabajo de la historia
   más arriesgada del último sprint. Si el equipo quiere gráficos en el PDF, hace falta
   acordar una librería —y cambiar esa regla del proyecto— o crear otra historia.
2. **La tipografía es una de las básicas que trae cualquier lector**, así que el archivo se ve
   igual en cualquier computadora sin llevar la tipografía adentro. La consecuencia es que
   solo se ven las letras de los idiomas de Europa occidental: lo demás sale como "?"
   (RN-10). Para los reportes de este proyecto, que son en castellano, no se pierde nada.
3. **"El mismo contenido que el reporte en pantalla" se interpretó como "las mismas líneas"**:
   el PDF solo agrega "Generado el" y los números de página, y acomoda el texto en páginas.
   No copia el aspecto de la página web.
4. **"Se descarga": el núcleo escribe el PDF en un `io.Writer` y la web lo descarga (#78).**
   Decidido el 20/10, al escribir la spec de la #78, con un cambio del contrato (PR #94). El
   primer borrador había tomado «se descarga» como «se guarda donde se indica» porque la
   aplicación no era web; con la #77 lo es, y desde el navegador no se ve el disco de la
   computadora donde corre el sistema. **La consola no exporta**: queda con lo que especifica
   la #24, y todo lo nuevo —sprints, esfuerzo, defectos, Planning Poker, reportes y PDF— se
   usa desde la web (decidido el 20/10, en el plan). Por eso esta historia no guarda el PDF
   en el disco: no hay quién lo pida.
5. **El contrato ya nombra la operación de exportar**: el reporte escribe su PDF donde se le
   pida (PR #94, 20/10). Reemplaza la propuesta del primer borrador, que lo grababa en una
   ubicación.
6. **El tercer criterio de la issue ("previsto para el Sprint 4") es de planificación**, no de
   comportamiento: se cumple con que la historia y su spec estén en el Sprint 4, y no lleva
   escenario ni test.

> **Puerta 1** — ¿los requisitos son correctos, completos y sin decisiones técnicas metidas?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios: [pendiente]

## Fase 2 — Diseño

Las decisiones transversales están en [AGENTS.md](../../../AGENTS.md) y no se repiten. Lo que
se lee del reporte son los datos del [contrato del dominio](../../diseno-dominio.md):
`Title`, `GeneratedOn` y `Sections`, cada una una `Section` con su `Heading` y sus `Lines`
(la misma que usa el dashboard). **No se leen `Metrics` ni `Scope`**: los números ya están
escritos en las líneas por la #18.

### Enfoque

Un PDF es un archivo de texto con una estructura fija: una cabecera con la versión, una serie
de objetos numerados (el catálogo, el árbol de páginas, cada página con su contenido y las
tipografías), una tabla con la posición exacta de cada objeto dentro del archivo y una marca
de fin. Escribirlo a mano con la biblioteca estándar es viable **si el alcance se limita a lo
que necesita un reporte de texto**, que es la decisión central de esta historia.

La exportación tiene tres pasos separados, cada uno con sus tests:

1. **Maquetar**: convertir el reporte en páginas de renglones, cada uno con su tipo de letra
   —título, encabezado o línea—. Acá viven el corte de renglones (RN-6), el paso de página
   (RN-7) y la numeración (RN-8). Este paso no sabe nada de PDF.
2. **Escribir**: convertir esas páginas en el contenido del archivo, con su tabla de
   posiciones. Acá vive la codificación de los caracteres (RN-10, RN-11).
3. **Entregar**: el archivo se arma entero en memoria y se escribe de una vez donde lo pida
   quien llama (`r.WritePDF(w)`, contrato del PR #94). El dominio no toca archivos: la web lo
   escribe en la respuesta, como descarga (#78).

Separar el maquetado de la escritura permite probar el corte de renglones y de páginas sin
leer PDF, y probar que el archivo es válido sin depender del maquetado.

### Archivos afectados

| Archivo | Se crea o se modifica | Para qué |
|---|---|---|
| `internal/domain/report_pdf.go` | se crea | la exportación del reporte: maquetar y escribir el PDF donde se pida |
| `internal/domain/report_pdf_test.go` | se crea | los tests de los escenarios BDD de la Fase 1 |
| `internal/domain/errors.go` | se modifica | el error de un PDF que no se pudo entregar |

### Decisiones

| Decisión | Alternativa descartada | Por qué |
|---|---|---|
| **Alcance: un PDF mínimo, escrito a mano con la biblioteca estándar: texto en páginas A4, en dos tipografías, sin gráficos, imágenes, tablas dibujadas ni tipografías incrustadas** | una librería de PDF (gofpdf, pdfcpu) | AGENTS.md no permite dependencias externas sin acuerdo, y la cátedra evalúa el código propio; un reporte de texto no necesita más. **Es la decisión que hace entrar la historia en el Sprint 4** |
| Generar el PDF directamente | generar una página HTML y que el usuario la imprima como PDF desde el navegador | el criterio pide un archivo PDF; imprimir es un paso manual que depende del navegador y no se puede testear |
| Helvetica y Helvetica-Bold, de las 14 tipografías estándar que trae todo lector, sin incrustar | incrustar una tipografía TrueType | incrustarla obliga a leer el archivo de la tipografía y armar sus tablas de caracteres y de anchos: es el grueso de un generador de PDF. Las estándar se ven igual en cualquier lector (RN-1) |
| Codificación WinAnsi; lo que no está en esa tabla sale como "?" | Unicode, con una tipografía incrustada | con las tipografías estándar es la única que muestra tildes y ñ, y cubre todo el castellano. Consecuencia: emojis y otros alfabetos no se ven (punto 2 de la puerta) |
| El contenido de las páginas va sin comprimir | comprimirlo con zlib, que está en la biblioteca estándar | el archivo pesa pocos KB, y sin comprimir el texto se encuentra en los bytes del archivo: es como los tests verifican la RN-2 sin un lector de PDF |
| El ancho del renglón se mide en caracteres: 85 en las líneas (cuerpo de 10 puntos), 75 en los encabezados y 50 en el título | medir el ancho real de cada letra con la tabla de anchos de Helvetica | la tabla son más de 200 números copiados a mano y sin forma de testearlos; con un máximo conservador el texto entra igual. Consecuencia: algún renglón queda más corto de lo que podría |
| Página A4 con márgenes de 2 cm y un largo fijo de renglones por página | calcular el largo según el tamaño de cada renglón | con un interlineado fijo el paso de página se cuenta en renglones y se prueba con números enteros |
| Maquetar, escribir y entregar son tres pasos separados | armar el archivo mientras se recorre el reporte | cada paso se prueba por su lado (ver Enfoque) |
| El archivo se arma entero en memoria y se escribe de una vez | ir escribiendo a medida que se arma | un error del maquetado no deja medio PDF en el destino, y el archivo es chico |
| **La exportación escribe el PDF donde lo pida quien llama (`r.WritePDF(w)`)** | que el reporte reciba la ubicación y grabe el archivo (`r.ExportPDF(path)`, el primer borrador) | el dominio no toca archivos (AGENTS.md) y la web necesita mandarlo como descarga sin pasar por el disco (#78). Decidido el 20/10 con el cambio del contrato del PR #94 |
| Si el destino falla al escribir, se informa con un error propio, «no se pudo generar el PDF», y no con el del destino | devolver el error del destino tal cual | el mensaje es para el usuario y va en castellano (AGENTS.md); el del destino viene en inglés. La web lo muestra con su aviso (#78, RN-19) |
| Vive en `internal/domain/report_pdf.go` | un paquete aparte, `internal/pdf` | AGENTS.md pide un solo paquete y crear otro se acuerda con el equipo. Que el núcleo arme texto con formato ya lo permite AGENTS.md desde el 20/10 |

### Dependencias y restricciones

- **Depende de la #18, que no está en `dev`.** No se espera: los tests arman un `Report` a
  mano dentro del paquete, como dice el contrato, con secciones y líneas inventadas.
- **No se lee nada fuera del reporte**: ni el proyecto, ni las métricas, ni los sprints. Es lo
  que garantiza la RN-3 por construcción.
- **`errors.go` es compartido**: el error de esta historia se agrega ahí.
- **Esta historia no toca ninguna interfaz**: la descarga es de la #78, y la consola no
  exporta (punto 4 de la puerta).

> **Puerta 2** — ¿el diseño es viable y coherente con el resto del proyecto?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios: [pendiente]

## Fase 3 — Tareas

Cada tarea es un ciclo de TDD completo: el test que falla en un commit, la implementación que
lo hace pasar en otro, y cada una deja los tests en verde. Siguen la convención de
[AGENTS.md](../../../AGENTS.md): un test por escenario, con su nombre, cada esquema como un
solo test con tabla de casos, y los errores verificados con `errors.Is`. Los tests escriben
el PDF en un destino en memoria: no tocan el disco.

**Cómo se verifica que el archivo es un PDF válido sin un lector.** El test de la tarea 6
comprueba que el archivo empieza con la marca de la versión y termina con la de fin, que la
tabla de posiciones apunta exactamente al comienzo de cada objeto, que el árbol de páginas
cuenta todas las páginas y que cada una declara el tamaño A4. Es lo que un lector controla
antes de mostrar nada. Que **de verdad abra** se verifica a mano (ver abajo).

| # | Tarea | Escenario BDD | Reglas | Estado |
|---|---|---|---|---|
| 1 | Maquetar un reporte chico: una página con el título, "Generado el" y la fecha, los encabezados y las líneas en orden, y "Página 1 de 1" | Exportar el reporte de un sprint a PDF | RN-2, RN-3, RN-8 | pendiente |
| 2 | Una línea más larga que el renglón se corta entre palabras | Una linea larga sigue en el renglon siguiente | RN-4, RN-6 | pendiente |
| 3 | Una palabra más larga que el renglón se corta en el margen | Una palabra mas larga que el renglon se corta en el margen | RN-6 | pendiente |
| 4 | Un reporte largo sigue en páginas nuevas, numeradas "Página N de M" | Un reporte largo ocupa varias paginas numeradas | RN-7, RN-8 | pendiente |
| 5 | Un encabezado que caería en el último renglón pasa a la página siguiente | Un encabezado nunca queda solo al pie de una pagina | RN-7 | pendiente |
| 6 | Escribir: el archivo tiene cabecera, objetos, tabla de posiciones exacta, árbol de páginas y marca de fin, con páginas A4 | El archivo se abre en un lector de PDF | RN-1, RN-4 | pendiente |
| 7 | El texto de cada renglón está en el contenido de su página, con el título más grande y los encabezados en negrita | Exportar el reporte de un sprint a PDF | RN-2, RN-5, RN-9 | pendiente |
| 8 | Los paréntesis y las barras invertidas se escriben sin romper el archivo | Los parentesis y las barras no rompen el archivo | RN-11 | pendiente |
| 9 | Tildes, ñ, ü, ¿ y ¡ se escriben en la codificación WinAnsi, y lo que no entra en ella sale como "?" | Letras del castellano y caracteres que no se pueden representar | RN-10 | pendiente |
| 10 | Si el destino falla al escribir, la exportación informa que no se pudo generar el PDF | Un PDF que no se puede entregar se informa | RN-12 | pendiente |
| 11 | Un reporte de proyecto y uno de sprint se exportan igual, sin nada agregado ni quitado | El PDF no agrega ni quita nada del reporte | RN-3, RN-13 | pendiente |

**Verificación manual, que no es un ciclo de TDD.** Antes de abrir el pull request, se exporta
un reporte real (de la #18, o armado a mano si la #18 no llegó) —descargado desde la web si la
#78 ya lo ofrece— y se abre en Vista Previa y en un navegador, comprobando que se ve el
texto, que se puede buscar y que las tildes salen bien. El resultado, con la fecha y los
lectores usados, se anota en la Fase 4. Cubre la parte del escenario "El archivo se abre en
un lector de PDF" que ningún test automático puede probar.

### Por qué este orden

- **De la 1 a la 5 va el maquetado**, que no sabe nada de PDF: primero el caso de una página,
  después los cortes de renglón y por último los de página, que se apoyan en ellos.
- **De la 6 a la 9 va la escritura**: primero la estructura vacía pero válida, después el
  texto, y al final los caracteres que necesitan un trato especial.
- **La 10 es la entrega que falla**, que recién tiene sentido con un archivo válido.
- **La 11 va al final** porque recorre los tres pasos con los dos tipos de reporte.

### Cobertura

Las trece reglas quedan cubiertas: RN-1 (6), RN-2 (1, 7), RN-3 (1, 11), RN-4 (2, 6),
RN-5 (7), RN-6 (2, 3), RN-7 (4, 5), RN-8 (1, 4), RN-9 (7), RN-10 (9), RN-11 (8),
RN-12 (10), RN-13 (11).

| Criterio de aceptación | Tareas |
|---|---|
| Exportar da un PDF, que la web ofrece para descargar | 6 (la descarga es de la #78) |
| Se abre en Vista Previa y en un navegador, con texto seleccionable | 6, 7 y la verificación manual |
| Título, fecha y todas las secciones en orden, sin nada agregado ni quitado | 1, 7, 11 |
| Título más grande y encabezados en negrita | 7 |
| Las líneas largas siguen en el renglón siguiente, nada fuera del margen | 2, 3 |
| Varias páginas numeradas, sin encabezados solos al pie | 4, 5 |
| Tildes y ñ, y "?" para lo que no se puede representar | 9 |
| Paréntesis y barras no rompen el archivo | 8 |
| Si el PDF no se puede entregar, se informa | 10 |
| Proyecto y sprint se exportan igual | 11 |

Los diez escenarios BDD de la Fase 1 tienen su tarea en la columna "Escenario BDD". El tercer
criterio de la issue (Sprint 4) no lleva tarea: es de planificación (punto 6 de la puerta).

### Dependencias

- **Ninguna tarea espera a la #18**: todas arman el `Report` a mano dentro del paquete.
- **La operación ya está nombrada en el contrato** (PR #94): todas las tareas escriben en un
  destino en memoria, y la 10 usa uno que falla a propósito.
- **La verificación manual** usa un reporte real si la #18 ya está en `dev`; si no, uno armado
  a mano con el mismo formato de líneas.

> **Puerta 3** — ¿las tareas cubren todos los criterios de aceptación?
> Aprobó: [pendiente] — Fecha: [pendiente] — Comentarios: [pendiente]

## Fase 4 — Implementación

_Se completa durante la implementación._

## Trazabilidad

| Eslabón | Dónde está |
|---|---|
| Historia de usuario | issue #19 |
| Especificación SDD | este archivo |
| Criterios de aceptación | Fase 1 |
| Escenarios BDD | Fase 1 |
| Tests | `internal/domain/report_pdf_test.go` |
| Código Go | `internal/domain/report_pdf.go` y su error en `internal/domain/errors.go` |
