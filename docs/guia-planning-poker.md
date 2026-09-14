# Guía de Planning Poker

Esto es para leer **antes** de la planning, así la reunión no se va en explicar cómo se juega.

## Para qué sirve

Estimar cuánto esfuerzo tiene una historia comparándola con otras, no adivinando horas. La
unidad son **Story Points**: mide tamaño (cuánto trabajo, qué tan complicada es, cuánta
incertidumbre tiene), no tiempo de reloj.

## La baraja

Se usa la sucesión de Fibonacci: **1, 2, 3, 5, 8, 13, 21**, más dos cartas especiales:

- **?** — no entiendo la historia, necesito que me la expliquen.
- **∞ / café** — es demasiado grande, hay que partirla en historias más chicas.

Los números crecen separados a propósito: cuanto más grande la historia, menos precisión
tiene sentido. No existe el 4 ni el 7.

## La historia de referencia

Antes de estimar nada, el equipo elige **una historia chica y conocida** y le asigna un
valor de arranque (normalmente 3). Todo lo demás se compara contra esa: "¿esto es más
grande o más chico que la historia de referencia?".

Sin referencia, cada uno estima en una escala distinta y los números no significan nada.

## Cómo se juega cada historia

1. Alguien lee la historia y sus criterios de aceptación en voz alta.
2. Se hacen las preguntas que hagan falta (esta es la parte que más valor tiene).
3. Cada uno elige su carta **sin decirla** y sin mirar la de los demás.
4. Se dan vuelta todas **al mismo tiempo**.
5. Si todos coincidieron, se anota y se pasa a la siguiente.
6. Si hay diferencias, hablan **el que puso el número más alto y el que puso el más bajo**,
   y explican por qué. No para convencer, para contar qué vieron que el otro no vio.
7. Se vuelve a votar. Se hacen como mucho **3 rondas**.
8. Si después de 3 rondas no hay acuerdo, se toma el valor más alto o se parte la historia.

## Reglas que importan

- **Los votos van ocultos hasta el final.** Si alguien canta su número antes, el resto se
  ancla y la estimación no sirve. Esto también es un criterio de aceptación de la historia
  #13, así que conviene vivirlo para después programarlo.
- **No estima el que más sabe, estiman todos.** Si el que va a programar la historia dice
  "3" primero, se terminó la discusión y se perdió la información del resto.
- **No se negocia el número para que entre en el sprint.** Lo que se ajusta es cuántas
  historias entran, no cuánto mide cada una.
- **Estimar no es comprometerse.** El compromiso viene después, cuando se arma el sprint.
- Si una historia sale 21, casi seguro hay que partirla.

## Herramientas

Se puede jugar con cartas de verdad, con una app de Planning Poker online, o cada uno
escribiendo su número en un papel y dándolo vuelta a la cuenta de tres. Lo único que importa
es que los votos no se vean hasta que votaron todos.
