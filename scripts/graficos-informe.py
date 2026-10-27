#!/usr/bin/env python3
#
# Arma los tres graficos del informe final a partir de la tabla de la seccion 3 de
# docs/informe-final.md: velocidad por sprint, horas estimadas contra reales, y defectos
# detectados y resueltos. Los deja en docs/graficos/, que es donde los busca el informe.
#
# Uso: ~/.venvs/facultad/bin/python scripts/graficos-informe.py
#
# Se corre de nuevo cada vez que cambia la tabla. Un sprint que todavia dice [pendiente] en
# una fila no aparece en el grafico de esa fila. Necesita matplotlib.

import os
import re
import sys

import matplotlib

matplotlib.use("Agg")
import matplotlib.pyplot as plt  # noqa: E402
from matplotlib.ticker import MaxNLocator  # noqa: E402

RAIZ = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
INFORME = os.path.join(RAIZ, "docs", "informe-final.md")
SALIDA = os.path.join(RAIZ, "docs", "graficos")

# Colores: las dos primeras posiciones de una paleta probada para daltonismo. El texto va
# siempre en gris oscuro, nunca en el color de la serie.
AZUL = "#2a78d6"
NARANJA = "#eb6834"
FONDO = "#fcfcfb"
TEXTO = "#0b0b0b"
TEXTO_SUAVE = "#52514e"
GRILLA = "#e4e3df"


def leer_tabla():
    """Devuelve los sprints del encabezado y un diccionario fila -> celdas."""
    texto = open(INFORME, encoding="utf-8").read()
    inicio = texto.index("| Métrica |")
    fin = texto.index("\n\n", inicio)
    lineas = texto[inicio:fin].split("\n")
    encabezado = [c.strip() for c in lineas[0].strip().strip("|").split("|")]
    sprints = encabezado[1:]
    filas = {}
    for linea in lineas[2:]:
        celdas = [c.strip() for c in linea.strip().strip("|").split("|")]
        filas[celdas[0]] = celdas[1:]
    return sprints, filas


def numeros(celda):
    """Los numeros de una celda, con coma decimal. Una celda [pendiente] no tiene ninguno."""
    if "pendiente" in celda:
        return []
    encontrados = re.findall(r"\d+(?:,\d+)?", celda.replace("**", ""))
    valores = []
    for n in encontrados:
        valores.append(float(n.replace(",", ".")))
    return valores


def con_coma(valor):
    if valor == int(valor):
        return str(int(valor))
    return ("%.2f" % valor).rstrip("0").rstrip(".").replace(".", ",")


def preparar(titulo):
    figura, eje = plt.subplots(figsize=(6.4, 3.6), dpi=200)
    figura.patch.set_facecolor(FONDO)
    eje.set_facecolor(FONDO)
    eje.set_title(titulo, loc="left", color=TEXTO, fontsize=11, fontweight="bold")
    for lado in ("top", "right", "left"):
        eje.spines[lado].set_visible(False)
    eje.spines["bottom"].set_color(TEXTO_SUAVE)
    eje.tick_params(colors=TEXTO_SUAVE, labelsize=9, length=0)
    eje.yaxis.grid(True, color=GRILLA, linewidth=0.8)
    eje.set_axisbelow(True)
    return figura, eje


def etiquetas(eje, barras, valores):
    # Lugar arriba para el numero de la barra mas alta, que si no choca con el titulo.
    if valores:
        eje.set_ylim(0, max(valores) * 1.18)
    eje.yaxis.set_major_locator(MaxNLocator(integer=True))
    for barra, valor in zip(barras, valores):
        eje.annotate(con_coma(valor), (barra.get_x() + barra.get_width() / 2, barra.get_height()),
                     xytext=(0, 3), textcoords="offset points", ha="center", va="bottom",
                     color=TEXTO, fontsize=9)


def guardar(figura, nombre):
    figura.tight_layout()
    ruta = os.path.join(SALIDA, nombre)
    figura.savefig(ruta, facecolor=FONDO)
    plt.close(figura)
    print("guardado:", os.path.relpath(ruta, RAIZ))


def una_serie(sprints, celdas, titulo, nombre):
    nombres = []
    valores = []
    for sprint, celda in zip(sprints, celdas):
        n = numeros(celda)
        if n:
            nombres.append(sprint)
            valores.append(n[0])
    figura, eje = preparar(titulo)
    barras = eje.bar(nombres, valores, width=0.45, color=AZUL)
    etiquetas(eje, barras, valores)
    eje.set_ylabel("story points", color=TEXTO_SUAVE, fontsize=9)
    guardar(figura, nombre)


def dos_series(sprints, celdas_a, celdas_b, nombre_a, nombre_b, titulo, unidad, nombre):
    """Columnas de a pares. Si a un sprint le falta uno de los dos datos, no se muestra."""
    nombres = []
    valores_a = []
    valores_b = []
    for i in range(len(sprints)):
        a = numeros(celdas_a[i])
        b = numeros(celdas_b[i])
        if a and b:
            nombres.append(sprints[i])
            valores_a.append(a[0])
            valores_b.append(b[0])
    figura, eje = preparar(titulo)
    ancho = 0.32
    posiciones = range(len(nombres))
    izquierda = []
    derecha = []
    for p in posiciones:
        izquierda.append(p - ancho / 2 - 0.01)
        derecha.append(p + ancho / 2 + 0.01)
    barras_a = eje.bar(izquierda, valores_a, width=ancho, color=AZUL, label=nombre_a)
    # La segunda serie lleva rayado, para que se distinga tambien impresa en blanco y negro.
    barras_b = eje.bar(derecha, valores_b, width=ancho, color=NARANJA, label=nombre_b,
                       hatch="//", edgecolor=FONDO, linewidth=0)
    etiquetas(eje, barras_a, valores_a)
    etiquetas(eje, barras_b, valores_b)
    if valores_a and valores_b:
        eje.set_ylim(0, max(max(valores_a), max(valores_b)) * 1.18)
    eje.set_xticks(list(posiciones))
    eje.set_xticklabels(nombres)
    eje.set_ylabel(unidad, color=TEXTO_SUAVE, fontsize=9)
    leyenda = eje.legend(frameon=False, fontsize=9, loc="upper left", bbox_to_anchor=(1, 1))
    for t in leyenda.get_texts():
        t.set_color(TEXTO)
    guardar(figura, nombre)


def main():
    if not os.path.exists(INFORME):
        print("No encuentro", INFORME)
        sys.exit(1)
    os.makedirs(SALIDA, exist_ok=True)
    sprints, filas = leer_tabla()

    una_serie(sprints, filas["SP completados (velocidad)"], "Velocidad por sprint", "velocidad.png")

    dos_series(sprints, filas["Horas estimadas"], filas["Horas reales trabajadas"],
               "Estimadas", "Reales", "Horas estimadas contra horas reales", "horas", "horas.png")

    detectados = []
    resueltos = []
    for celda in filas["Defectos detectados / resueltos"]:
        n = numeros(celda)
        if len(n) == 2:
            detectados.append(con_coma(n[0]))
            resueltos.append(con_coma(n[1]))
        else:
            detectados.append("[pendiente]")
            resueltos.append("[pendiente]")
    dos_series(sprints, detectados, resueltos, "Detectados", "Resueltos",
               "Defectos detectados y resueltos", "defectos", "defectos.png")


main()
