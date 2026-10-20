#!/usr/bin/env bash
#
# Calcula las metricas de un sprint a partir del tablero y de las issues.
# Son las que pide la consigna: story points planificados y completados, velocidad,
# horas estimadas y reales, desviacion, porcentaje de historias completadas y defectos.
#
# Uso: scripts/metricas.sh "Sprint 1"
#
# Limitacion conocida: el script lee el campo Sprint que cada historia tiene HOY en el
# tablero. Una historia que no se termino y paso al sprint siguiente deja de contar en el
# sprint donde se comprometio, asi que para un sprint cerrado los numeros pueden no
# coincidir con lo que paso (el Sprint 1 da 11 de 11 puntos, pero se comprometieron 16).
# Para un sprint cerrado, el registro que vale es su acta de review
# (docs/actas/sprint-N-review.md), que se escribe al cierre y no cambia. Este script sirve
# para el sprint abierto y para armar esa acta el dia del cierre.

set -u

SPRINT="${1:-Sprint 1}"
OWNER="emisorato1"
PROYECTO="3"
REPO="emisorato1/seguimiento-medicion-ics"

command -v gh >/dev/null || { echo "Falta gh. Instalar: brew install gh"; exit 1; }
command -v jq >/dev/null || { echo "Falta jq. Instalar: brew install jq"; exit 1; }

ITEMS=$(gh project item-list "$PROYECTO" --owner "$OWNER" --limit 200 --format json)

echo "# Metricas del $SPRINT"
echo
echo "_Calculadas el $(date '+%Y-%m-%d %H:%M') con scripts/metricas.sh._"
echo
echo "## Historias"
echo
echo "$ITEMS" | jq -r --arg s "$SPRINT" '
  [.items[] | select(.sprint == $s)]
  | (["| # | Historia | SP | Estado | Horas est. | Horas reales |", "|---|---|---|---|---|---|"]
     + (sort_by(.content.number) | map(
         "| #\(.content.number) | \(.title) | \(.["story Points"] // 0) | \(.status // "-") "
         + "| \(.["horas estimadas"] // "-") | \(.["horas reales"] // "-") |")))
  | join("\n")'

echo
echo "## Resumen"
echo
echo "$ITEMS" | jq -r --arg s "$SPRINT" '
  # Redondea al entero mas cercano (o a "decimales" decimales). floor no sirve: con
  # numeros negativos redondea para abajo y una desviacion de -12,1% salia -13%.
  def redondear: round;
  def redondear(decimales): (. * pow(10; decimales) | round) / pow(10; decimales);
  [.items[] | select(.sprint == $s)] as $todas
  | [$todas[] | select(.status == "Hecho")] as $hechas
  | ([$todas[] | .["story Points"] // 0] | add // 0) as $planificados
  | ([$hechas[] | .["story Points"] // 0] | add // 0) as $completados
  | ([$todas[] | .["horas estimadas"] // 0] | add // 0 | redondear(2)) as $estimadas
  | ([$todas[] | .["horas reales"] // 0] | add // 0 | redondear(2)) as $reales
  | ([$hechas[] | .["horas estimadas"] // 0] | add // 0 | redondear(2)) as $estimadas_hechas
  | ([$hechas[] | .["horas reales"] // 0] | add // 0 | redondear(2)) as $reales_hechas
  | "| Metrica | Valor |\n|---|---|"
  + "\n| Story points planificados | \($planificados) |"
  + "\n| Story points completados | \($completados) |"
  + "\n| **Velocidad del sprint** | **\($completados)** |"
  + "\n| Historias planificadas | \($todas | length) |"
  + "\n| Historias completadas | \($hechas | length) |"
  + "\n| Porcentaje de historias completadas | \(if ($todas|length) > 0 then (($hechas|length) * 100 / ($todas|length) | redondear) else 0 end)% |"
  + "\n| Horas estimadas, todas las historias | \($estimadas) |"
  + "\n| Horas reales, todas las historias | \($reales) |"
  + "\n| Horas estimadas, historias terminadas | \($estimadas_hechas) |"
  + "\n| Horas reales, historias terminadas | \($reales_hechas) |"
  + "\n| Desviacion de esfuerzo, historias terminadas | \($reales_hechas - $estimadas_hechas | redondear(2)) h (\(if $estimadas_hechas > 0 then (($reales_hechas - $estimadas_hechas) * 100 / $estimadas_hechas | redondear) else 0 end)%) |"
  + "\n| Horas por story point completado | \(if $completados > 0 then ($reales_hechas / $completados | redondear(1)) else "sin datos" end) |"'

echo
echo "## Defectos"
echo
DETECTADOS=$(gh issue list --repo "$REPO" --label defecto --state all --json number --jq 'length')
RESUELTOS=$(gh issue list --repo "$REPO" --label defecto --state closed --json number --jq 'length')
echo "| Metrica | Valor |"
echo "|---|---|"
echo "| Defectos detectados | $DETECTADOS |"
echo "| Defectos resueltos | $RESUELTOS |"
echo "| Defectos abiertos | $((DETECTADOS - RESUELTOS)) |"
echo
gh issue list --repo "$REPO" --label defecto --state all \
  --json number,title,state --jq '.[] | "- #\(.number) [\(if .state == "CLOSED" then "resuelto" else "abierto" end)] \(.title)"'

echo
echo "_La velocidad, la desviacion y las horas por story point solo cuentan historias en estado"
echo "Hecho: una historia a medias no suma story points, aunque tenga horas cargadas. Para un"
echo "sprint cerrado vale su acta de review: este script usa el Sprint que cada historia tiene hoy._"
