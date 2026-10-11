#!/usr/bin/env bash
#
# Mantiene el tablero del proyecto al dia sin tener que entrar a la web.
# El GitHub Project es la fuente de verdad del estado del proyecto.
#
# Uso:
#   scripts/tablero.sh ver
#   scripts/tablero.sh estado <issue> "<Estado>"
#   scripts/tablero.sh puntos <issue> <numero>
#   scripts/tablero.sh sprint <issue> "Sprint N"
#   scripts/tablero.sh prioridad <issue> <Alta|Media|Baja>
#
# Estados validos: Product Backlog, Sprint Backlog, Especificacion SDD,
#                  Lista para implementar, En progreso, En review, Hecho
#
# OJO: no agregues ni renombres estados editando las opciones del campo Status
# por la API. GitHub les cambia el identificador y TODAS las tarjetas que lo
# tenian quedan sin estado. Hacelo desde la web del Project, que conserva los
# identificadores.

set -e

OWNER="emisorato1"
NUMERO="3"

uso() {
  sed -n '3,16p' "$0" | sed 's/^# \{0,1\}//'
  exit 1
}

requisitos() {
  command -v gh >/dev/null || { echo "Falta gh. Instalar: brew install gh"; exit 1; }
  command -v jq >/dev/null || { echo "Falta jq. Instalar: brew install jq"; exit 1; }
}

proyecto_id() {
  gh project view "$NUMERO" --owner "$OWNER" --format json | jq -r '.id'
}

item_id() {
  gh project item-list "$NUMERO" --owner "$OWNER" --limit 200 --format json \
    | jq -r --argjson n "$1" '.items[] | select(.content.number == $n) | .id'
}

campo_id() {
  gh project field-list "$NUMERO" --owner "$OWNER" --limit 50 --format json \
    | jq -r --arg n "$1" '.fields[] | select(.name == $n) | .id'
}

opcion_id() {
  gh project field-list "$NUMERO" --owner "$OWNER" --limit 50 --format json \
    | jq -r --arg n "$1" --arg o "$2" \
      '.fields[] | select(.name == $n) | .options[] | select(.name == $o) | .id'
}

editar_opcion() {  # $1 campo  $2 opcion  $3 item
  local fid oid
  fid=$(campo_id "$1")
  oid=$(opcion_id "$1" "$2")
  [ -n "$oid" ] || { echo "El campo '$1' no tiene la opcion '$2'."; exit 1; }
  gh project item-edit --id "$3" --project-id "$(proyecto_id)" \
    --field-id "$fid" --single-select-option-id "$oid" >/dev/null
}

es_numero() {
  case "$1" in
    ''|*[!0-9]*) return 1 ;;
    *) return 0 ;;
  esac
}

resolver_item() {  # $1 numero de issue
  local item
  es_numero "$1" || { echo "'$1' no es un numero de issue valido."; exit 1; }
  item=$(item_id "$1")
  [ -n "$item" ] || { echo "La issue #$1 no esta en el tablero."; exit 1; }
  echo "$item"
}

requisitos
[ $# -ge 1 ] || uso

case "$1" in
  ver)
    gh project item-list "$NUMERO" --owner "$OWNER" --limit 200 --format json \
      | jq -r '.items[]
          | "#\(.content.number)\t\(.status // "sin estado")\t\((to_entries | map(select(.key | ascii_downcase == "story points")) | .[0].value) // "-") SP\t\(.title)"' \
      | sort -t'#' -k2 -n
    ;;
  estado)
    [ $# -eq 3 ] || uso
    editar_opcion "Status" "$3" "$(resolver_item "$2")"
    echo "Issue #$2 -> $3"
    ;;
  sprint)
    [ $# -eq 3 ] || uso
    editar_opcion "Sprint" "$3" "$(resolver_item "$2")"
    echo "Issue #$2 -> $3"
    ;;
  prioridad)
    [ $# -eq 3 ] || uso
    editar_opcion "Prioridad" "$3" "$(resolver_item "$2")"
    echo "Issue #$2 -> prioridad $3"
    ;;
  puntos)
    [ $# -eq 3 ] || uso
    es_numero "$3" || { echo "Los story points tienen que ser un numero. Recibi: '$3'"; exit 1; }
    gh project item-edit --id "$(resolver_item "$2")" --project-id "$(proyecto_id)" \
      --field-id "$(campo_id "Story Points")" --number "$3" >/dev/null
    echo "Issue #$2 -> $3 story points"
    ;;
  *)
    uso
    ;;
esac
