## Qué hace

[Una o dos frases. Si cierra la historia, poner `Closes #N`.]

- Historia: #
- Spec: `docs/specs/NN-nombre.md`

## Checklist

- [ ] La spec de la historia está aprobada y las puertas quedaron registradas en el archivo.
- [ ] Los escenarios BDD de la spec tienen su test.
- [ ] Los tests pasan (`go test ./...`) y el CI está en verde.
- [ ] El código está formateado (`gofmt -l .` no devuelve nada).
- [ ] Los commits muestran el ciclo TDD: el test en un commit y la implementación en otro.
- [ ] Si la implementación se desvió de la spec, la spec quedó actualizada en este mismo PR.
- [ ] La historia está movida en el tablero y `ESTADO.md` regenerado.

## Cómo probarlo

[Qué correr para verificar que funciona.]

## Revisión

Revisa un integrante que **no** escribió este código.
