#!/usr/bin/env bash
#
# Regenera ESTADO.md leyendo el GitHub Project, que es la fuente de verdad.
# Correr despues de cualquier cambio en el tablero y commitear el resultado,
# asi el estado queda disponible aunque no se tenga gh instalado.
#
# Uso: scripts/estado.sh

set -e

OWNER="emisorato1"
NUMERO="3"
SALIDA="ESTADO.md"

command -v gh >/dev/null || { echo "Falta gh. Instalar: brew install gh"; exit 1; }
command -v jq >/dev/null || { echo "Falta jq. Instalar: brew install jq"; exit 1; }

ITEMS=$(gh project item-list "$NUMERO" --owner "$OWNER" --limit 200 --format json)
URL=$(gh project view "$NUMERO" --owner "$OWNER" --format json | jq -r '.url')

tabla() {  # $1 = estado a filtrar
  echo "$ITEMS" | jq -r --arg e "$1" '
    [.items[] | select((.status // "") == $e)]
    | if length == 0 then "_No hay historias en este estado._"
      else (["| # | Historia | SP | Prioridad | Sprint |", "|---|---|---|---|---|"]
            + (sort_by(.content.number)
               | map("| #\(.content.number) | \(.title) | \((to_entries | map(select(.key | ascii_downcase == "story points")) | .[0].value) // "-") | \(.prioridad // "-") | \(.sprint // "-") |")))
           | join("\n")
      end'
}

{
  echo "# Estado del proyecto"
  echo
  echo "_Generado por \`scripts/estado.sh\` el $(date '+%Y-%m-%d %H:%M'). No editar a mano._"
  echo
  echo "La fuente de verdad es el tablero: $URL"
  echo "Este archivo es una copia para poder leer el estado sin \`gh\` instalado."
  echo
  echo "## Resumen"
  echo
  echo "$ITEMS" | jq -r '
    .items as $i
    | "- Historias en el tablero: \($i | length)"
    + "\n- Sin estimar: \([$i[] | select(((to_entries | map(select(.key | ascii_downcase == "story points")) | .[0].value) // 0) == 0)] | length)"
    + "\n- Story points cargados: \([$i[] | (to_entries | map(select(.key | ascii_downcase == "story points")) | .[0].value) // 0] | add)"
    + "\n- Terminadas: \([$i[] | select(.status == "Hecho")] | length)"'
  echo
  for estado in "En review" "En progreso" "Especificacion SDD" "Sprint Backlog" "Product Backlog" "Hecho"; do
    echo "## $estado"
    echo
    tabla "$estado"
    echo
  done
  echo "## Ultimos commits"
  echo
  echo '```'
  git log --format='%h %ad %s' --date=short -8
  echo '```'
} > "$SALIDA"

echo "Escrito $SALIDA"
