# Especificaciones SDD

Una carpeta por sprint, y dentro una spec por historia:

```
docs/specs/
├── PLANTILLA.md              se copia para cada spec nueva
├── sprint-1/
│   ├── 03-crear-proyectos.md
│   └── 04-registrar-integrantes.md
├── sprint-2/
└── ...
```

- El nombre del archivo es `NN-nombre-corto.md`, donde `NN` es el número de la issue.
- La spec va en la carpeta del sprint en el que la historia **se comprometió**.
- Si una historia no se termina y vuelve al Product Backlog, cuando se vuelve a comprometer
  se **mueve** la spec a la carpeta del sprint nuevo (`git mv`), no se copia: la historia
  tiene una sola spec y hay que poder seguir su historial.
- Las specs se versionan junto al código y se actualizan si la implementación se desvía.
