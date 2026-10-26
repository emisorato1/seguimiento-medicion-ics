#!/usr/bin/env bash
#
# Calcula las metricas de un sprint a partir del tablero y de las issues.
# Son las que pide la consigna: story points planificados y completados, velocidad,
# horas estimadas y reales, desviacion, porcentaje de historias completadas y defectos.
#
# Uso: scripts/metricas.sh "Sprint 1"

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
  [.items[] | select(.sprint == $s)] as $todas
  | [$todas[] | select(.status == "Hecho")] as $hechas
  | ([$todas[] | .["story Points"] // 0] | add) as $planificados
  | ([$hechas[] | .["story Points"] // 0] | add) as $completados
  | ([$todas[] | .["horas estimadas"] // 0] | add) as $estimadas
  | ([$todas[] | .["horas reales"] // 0] | add) as $reales
  | "| Metrica | Valor |\n|---|---|"
  + "\n| Story points planificados | \($planificados) |"
  + "\n| Story points completados | \($completados) |"
  + "\n| **Velocidad del sprint** | **\($completados)** |"
  + "\n| Historias planificadas | \($todas | length) |"
  + "\n| Historias completadas | \($hechas | length) |"
  + "\n| Porcentaje de historias completadas | \(if ($todas|length) > 0 then (($hechas|length) * 100 / ($todas|length) | floor) else 0 end)% |"
  + "\n| Horas estimadas | \($estimadas) |"
  + "\n| Horas reales | \($reales) |"
  + "\n| Desviacion de esfuerzo | \($reales - $estimadas) h (\(if $estimadas > 0 then ((($reales - $estimadas) * 100 / $estimadas) | floor) else 0 end)%) |"
  + "\n| Horas por story point completado | \(if $completados > 0 then (($reales * 10 / $completados | floor) / 10) else "sin datos" end) |"'

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
echo "_La velocidad y la desviacion solo cuentan historias en estado Hecho: una historia a"
echo "medias no suma story points, aunque tenga horas cargadas._"
