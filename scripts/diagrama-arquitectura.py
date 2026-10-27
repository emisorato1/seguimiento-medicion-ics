#!/usr/bin/env python3
#
# Dibuja el diagrama de arquitectura del informe final y de la presentacion, y lo deja en
# docs/graficos/arquitectura.png. Esta pensado para leerse impreso en blanco y negro: cajas
# con borde negro, las de archivos con borde de rayas y fondo gris claro.
#
# Uso: ~/.venvs/facultad/bin/python scripts/diagrama-arquitectura.py
#
# Las flechas son los import entre paquetes (quien usa a quien), tal como estan en el
# codigo. Si cambia algun import, hay que cambiar FLECHAS. Necesita matplotlib.

import os

import matplotlib

matplotlib.use("Agg")
import matplotlib.pyplot as plt  # noqa: E402
from matplotlib.patches import FancyArrowPatch, FancyBboxPatch  # noqa: E402

RAIZ = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
SALIDA = os.path.join(RAIZ, "docs", "graficos", "arquitectura.png")

NEGRO = "#0b0b0b"
GRIS = "#52514e"
FONDO = "#ffffff"
FONDO_ARCHIVO = "#ececea"
FONDO_DOMINIO = "#d9d8d4"

ANCHO = 2.6
ALTO = 0.9

# nombre: (x del centro, y del centro, titulo, detalle, tipo)
CAJAS = {
    "navegador": (3.9, 6.2, "Navegador", "la persona usa la web", "persona"),
    "terminal": (8.8, 6.2, "Terminal", "la persona usa la consola", "persona"),
    "cmdweb": (3.9, 4.7, "cmd/web", "arranca el servidor web", "paquete"),
    "cmdapp": (8.8, 4.7, "cmd/app", "la consola (en memoria)", "paquete"),
    "store": (2.4, 3.0, "internal/store", "guardar y cargar", "paquete"),
    "web": (5.4, 3.0, "internal/web", "pantallas y formularios", "paquete"),
    "dominio": (5.9, 1.2, "internal/domain", "las reglas de negocio", "dominio"),
    "base": (0.2, 1.2, "datos-ics.db", "base SQLite", "archivo"),
    "json": (3.0, 1.2, "datos-ics.json", "importar y exportar", "archivo"),
}

# (desde, hasta, texto): "usa" es un import entre paquetes.
FLECHAS = [
    ("navegador", "cmdweb", ""),
    ("terminal", "cmdapp", ""),
    ("cmdweb", "web", "usa"),
    ("cmdweb", "store", "usa"),
    ("web", "dominio", "usa"),
    ("store", "dominio", "usa"),
    ("cmdapp", "dominio", "usa"),
    ("store", "base", "lee y escribe"),
    ("store", "json", "importa y exporta"),
]


def caja(eje, x, y, titulo, detalle, tipo):
    estilo = "-"
    fondo = FONDO
    grosor = 1.4
    if tipo == "archivo":
        estilo = "--"
        fondo = FONDO_ARCHIVO
    if tipo == "dominio":
        fondo = FONDO_DOMINIO
        grosor = 2.2
    if tipo == "persona":
        estilo = ":"
    rect = FancyBboxPatch((x - ANCHO / 2, y - ALTO / 2), ANCHO, ALTO,
                          boxstyle="round,pad=0.02,rounding_size=0.12",
                          linewidth=grosor, linestyle=estilo, edgecolor=NEGRO, facecolor=fondo)
    eje.add_patch(rect)
    eje.text(x, y + 0.13, titulo, ha="center", va="center", fontsize=11,
             fontweight="bold", color=NEGRO, family="monospace" if tipo != "persona" else None)
    eje.text(x, y - 0.2, detalle, ha="center", va="center", fontsize=8.5, color=GRIS)


def borde(x, y, hacia_x, hacia_y):
    """El punto del borde de la caja (x, y) en la direccion de (hacia_x, hacia_y)."""
    dx = hacia_x - x
    dy = hacia_y - y
    if dx == 0 and dy == 0:
        return x, y
    escala_x = 9999
    escala_y = 9999
    if dx != 0:
        escala_x = (ANCHO / 2 + 0.05) / abs(dx)
    if dy != 0:
        escala_y = (ALTO / 2 + 0.05) / abs(dy)
    escala = min(escala_x, escala_y)
    return x + dx * escala, y + dy * escala


def flecha(eje, desde, hasta, texto):
    x1, y1 = CAJAS[desde][0], CAJAS[desde][1]
    x2, y2 = CAJAS[hasta][0], CAJAS[hasta][1]
    a = borde(x1, y1, x2, y2)
    b = borde(x2, y2, x1, y1)
    eje.add_patch(FancyArrowPatch(a, b, arrowstyle="-|>", mutation_scale=14,
                                  linewidth=1.2, color=NEGRO))
    if texto:
        # El texto va en el medio de la flecha, con fondo blanco para que la linea no lo tape.
        eje.text((a[0] + b[0]) / 2, (a[1] + b[1]) / 2, texto, fontsize=8, color=GRIS,
                 ha="center", va="center", zorder=3,
                 bbox={"boxstyle": "round,pad=0.2", "facecolor": FONDO, "edgecolor": "none"})


def main():
    figura, eje = plt.subplots(figsize=(9.2, 5.4), dpi=200)
    figura.patch.set_facecolor(FONDO)
    eje.set_xlim(-1.3, 10.3)
    eje.set_ylim(0.1, 6.9)
    eje.axis("off")
    for desde, hasta, texto in FLECHAS:
        flecha(eje, desde, hasta, texto)
    for nombre in CAJAS:
        x, y, titulo, detalle, tipo = CAJAS[nombre]
        caja(eje, x, y, titulo, detalle, tipo)
    eje.text(-1.2, 0.15, "Las flechas \"usa\" son los import de Go. Borde de rayas: archivos de datos. Borde de puntos: quien usa el sistema.",
             fontsize=8, color=GRIS, ha="left", va="bottom")
    os.makedirs(os.path.dirname(SALIDA), exist_ok=True)
    figura.tight_layout()
    figura.savefig(SALIDA, facecolor=FONDO)
    plt.close(figura)
    print("guardado:", os.path.relpath(SALIDA, RAIZ))


main()
