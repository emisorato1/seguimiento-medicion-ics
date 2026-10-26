# Especificaciones del Sprint 0

El Sprint 0 no produjo funcionalidad del producto: preparó el entorno de trabajo. Sus dos
ítems —configurar el repositorio y el tablero, y definir el Product Backlog inicial— son
tareas de configuración y de planificación, no historias de usuario.

Por eso **no llevan el flujo de cuatro fases con puertas de aprobación**. El apunte de la
cátedra lo dice explícitamente en el capítulo de anti-patrones:

> **Sobreespecificación.** Síntoma: la spec es más larga que el código que genera. […]
> Causa: no calibrar la intensidad de SDD al tamaño del problema. Aplicar el flujo completo
> de cuatro fases a un cambio que podría resolverse con un prompt directo.
> Solución: el principio de proporcionalidad. No todo necesita SDD.

Lo que sí corresponde documentar, y es lo que hay en esta carpeta, son las **decisiones**
que se tomaron: qué se eligió, qué se descartó y por qué. Son decisiones que condicionan
todo el resto del proyecto, y en la defensa hay que poder justificarlas.

| Documento | Ítem |
|---|---|
| [01-configurar-repositorio-y-tablero.md](01-configurar-repositorio-y-tablero.md) | issue #1 |
| [02-definir-product-backlog-inicial.md](02-definir-product-backlog-inicial.md) | issue #2 |

A partir del Sprint 1, cada historia sí lleva su spec completa con las cuatro fases.
