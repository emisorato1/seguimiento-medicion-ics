#!/usr/bin/env bash
#
# Dice si tengo pull requests esperando mi revision, y como estan los mios.
# Correr al empezar cada sesion de trabajo: un PR parado frena a todo el equipo.
#
# Uso: scripts/revisiones.sh

set -e

REPO="emisorato1/seguimiento-medicion-ics"

command -v gh >/dev/null || { echo "Falta gh. Instalar: brew install gh"; exit 1; }
command -v jq >/dev/null || { echo "Falta jq. Instalar: brew install jq"; exit 1; }

YO=$(gh api user --jq '.login')
PRS=$(gh pr list --repo "$REPO" --state open --json \
  number,title,author,reviewDecision,isDraft,headRefName,statusCheckRollup)

echo "== Esperan tu revision =="
echo "$PRS" | jq -r --arg yo "$YO" '
  [ .[] | select(.author.login != $yo and .isDraft == false and .reviewDecision != "APPROVED") ]
  | if length == 0 then "  nada pendiente"
    else .[] | "  #\(.number)  \(.author.login)  \(.title)"
    end'

echo
echo "== Tus pull requests =="
echo "$PRS" | jq -r --arg yo "$YO" '
  [ .[] | select(.author.login == $yo) ]
  | if length == 0 then "  ninguno abierto"
    else .[] |
      "  #\(.number)  \(.title)\n" +
      "     revision: \(.reviewDecision // "sin revisar")   " +
      "CI: \([.statusCheckRollup[]? | .conclusion // .status] | join(",") // "sin correr")"
    end'

echo
echo "Recorda: la aprobacion la hace una persona, no Claude."
echo "Para revisar:  gh pr diff <numero>   |   gh pr view <numero> --web"
