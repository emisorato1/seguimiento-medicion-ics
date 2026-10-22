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
# coincidir con lo que paso (el Sprint 1 da 6 de 6 puntos, pero se comprometieron 16).
# Para un sprint cerrado, el registro que vale es su acta de review
# (docs/actas/sprint-N-review.md), que se escribe al cierre y no cambia. Este script sirve
# para el sprint abierto y para armar esa acta el dia del cierre.
#
# Otra limitacion: el tablero guarda un solo total de horas por historia. Si una historia
# viene de un sprint anterior, sus horas estimadas y reales se suman enteras en el sprint
# actual, aunque una parte ya se conto en el anterior (la #12, la #24 y la #6 traen 22,25 h
# reales del Sprint 1). Para no contarlas dos veces, se resta lo que dice el acta de review
# del sprint anterior; la regla esta en docs/informe-de-metricas.md.

set -u

SPRINT="${1:-Sprint 1}"
OWNER="emisorato1"
PROYECTO="3"
REPO="emisorato1/seguimiento-medicion-ics"

command -v gh >/dev/null || { echo "Falta gh. Instalar: brew install gh"; exit 1; }
command -v jq >/dev/null || { echo "Falta jq. Instalar: brew install jq"; exit 1; }

# Fechas de cada sprint, las de CLAUDE.md y las actas de planning. Sirven para saber en que
# sprint se detecto y se resolvio cada defecto. Un sprint que no esta aca cuenta todos.
# Lo que pasa entre un sprint y el siguiente cuenta en el siguiente: por eso el Sprint 2
# arranca el 17/10, el dia despues del cierre del Sprint 1, aunque empezo el 20/10.
case "$SPRINT" in
  "Sprint 1") DESDE="2026-09-28"; HASTA="2026-10-16" ;;
  "Sprint 2") DESDE="2026-10-17"; HASTA="2026-10-25" ;;
  "Sprint 3") DESDE="2026-10-26"; HASTA="2026-10-30" ;;
  "Sprint 4") DESDE="2026-10-31"; HASTA="2026-11-01" ;;
  *)          DESDE="0000-00-00"; HASTA="9999-99-99" ;;
esac

ITEMS=$(gh project item-list "$PROYECTO" --owner "$OWNER" --limit 200 --format json)

# Las historias con la etiqueta "adelantada" entraron al sprint en curso por la decision 11
# del plan (quien termina lo suyo adelanta su proxima historia). No suman a lo comprometido
# y se muestran aparte; si se terminan, suman a la velocidad del sprint. Si una no se termina
# y vuelve a su sprint del plan, hay que sacarle la etiqueta.
ADELANTADAS=$(gh issue list --repo "$REPO" --label adelantada --state all --limit 100 --json number --jq '[.[].number]')

echo "# Metricas del $SPRINT"
echo
echo "_Calculadas el $(date '+%Y-%m-%d %H:%M') con scripts/metricas.sh._"
echo
echo "## Historias"
echo
echo "$ITEMS" | jq -r --arg s "$SPRINT" --argjson adel "$ADELANTADAS" '
  [.items[] | select(.sprint == $s)]
  | (["| # | Historia | SP | Estado | Horas est. | Horas reales |", "|---|---|---|---|---|---|"]
     + (sort_by(.content.number) | map(
         "| #\(.content.number) | \(.title) | \(.["story Points"] // 0) | \(.status // "-")"
         + "\(if (.content.number | IN($adel[])) then " (adelantada)" else "" end) "
         + "| \(.["horas estimadas"] // "-") | \(.["horas reales"] // "-") |")))
  | join("\n")'

echo
echo "## Resumen"
echo
echo "$ITEMS" | jq -r --arg s "$SPRINT" --argjson adel "$ADELANTADAS" '
  # Redondea al entero mas cercano (o a "decimales" decimales). floor no sirve: con
  # numeros negativos redondea para abajo y una desviacion de -12,1% salia -13%.
  def redondear: round;
  def redondear(decimales): (. * pow(10; decimales) | round) / pow(10; decimales);
  [.items[] | select(.sprint == $s)] as $del_sprint
  | [$del_sprint[] | select((.content.number | IN($adel[])) | not)] as $todas
  | [$del_sprint[] | select(.content.number | IN($adel[]))] as $adelantadas
  | [$todas[] | select(.status == "Hecho")] as $hechas
  | [$adelantadas[] | select(.status == "Hecho")] as $adelantadas_hechas
  # Una historia sin estimar (sin story points, como la #77) no tiene contra que comparar
  # sus horas reales: no entra en la desviacion ni en las horas por story point. Una
  # adelantada terminada si entra, igual que en la velocidad.
  | [($hechas + $adelantadas_hechas)[] | select((.["story Points"] // 0) > 0)] as $hechas_estimadas
  | ([$adelantadas[] | .["story Points"] // 0] | add // 0) as $adelantados
  | ([$adelantadas_hechas[] | .["story Points"] // 0] | add // 0) as $adelantados_hechos
  | ([$todas[] | .["story Points"] // 0] | add // 0) as $planificados
  | ([$hechas[] | .["story Points"] // 0] | add // 0) as $completados
  | ([$todas[] | .["horas estimadas"] // 0] | add // 0 | redondear(2)) as $estimadas
  | ([$todas[] | .["horas reales"] // 0] | add // 0 | redondear(2)) as $reales
  | ([$hechas_estimadas[] | .["horas estimadas"] // 0] | add // 0 | redondear(2)) as $estimadas_hechas
  | ([$hechas_estimadas[] | .["horas reales"] // 0] | add // 0 | redondear(2)) as $reales_hechas
  | "| Metrica | Valor |\n|---|---|"
  + "\n| Story points planificados | \($planificados) |"
  + "\n| Story points completados | \($completados) |"
  + "\n| **Velocidad del sprint** | **\($completados + $adelantados_hechos)** (\($completados) + \($adelantados_hechos) adelantados) |"
  + "\n| Historias planificadas | \($todas | length) |"
  + "\n| Historias completadas | \($hechas | length) |"
  + "\n| Porcentaje de historias completadas | \(if ($todas|length) > 0 then (($hechas|length) * 100 / ($todas|length) | redondear) else 0 end)% |"
  + "\n| Horas estimadas, todas las historias | \($estimadas) |"
  + "\n| Horas reales, todas las historias | \($reales) |"
  + "\n| Horas reales, contando las adelantadas | \($reales + ([$adelantadas[] | .["horas reales"] // 0] | add // 0) | redondear(2)) |"
  + "\n| Horas estimadas, historias terminadas | \($estimadas_hechas) |"
  + "\n| Horas reales, historias terminadas | \($reales_hechas) |"
  + "\n| Desviacion de esfuerzo, historias terminadas | \($reales_hechas - $estimadas_hechas | redondear(2)) h (\(if $estimadas_hechas > 0 then (($reales_hechas - $estimadas_hechas) * 100 / $estimadas_hechas | redondear) else 0 end)%) |"
  + "\n| Horas por story point completado | \(if ($completados + $adelantados_hechos) > 0 then ($reales_hechas / ($completados + $adelantados_hechos) | redondear(1)) else "sin datos" end) |"
  + "\n\n**Adelantadas** (decision 11 del plan: no suman a lo comprometido)\n"
  + "\n| Metrica | Valor |\n|---|---|"
  + "\n| Historias adelantadas | \($adelantadas | length)\(if ($adelantadas | length) > 0 then " (" + ($adelantadas | map("#\(.content.number)") | join(", ")) + ")" else "" end) |"
  + "\n| Story points adelantados | \($adelantados) |"
  + "\n| Historias adelantadas completadas | \($adelantadas_hechas | length) |"
  + "\n| Story points adelantados completados | \($adelantados_hechos) |"
  + "\n| Horas estimadas de las adelantadas | \([$adelantadas[] | .["horas estimadas"] // 0] | add // 0 | redondear(2)) |"
  + "\n| Horas reales de las adelantadas | \([$adelantadas[] | .["horas reales"] // 0] | add // 0 | redondear(2)) |"'

echo
echo "## Defectos"
echo
DEFECTOS=$(gh issue list --repo "$REPO" --label defecto --state all --limit 200 \
  --json number,title,createdAt,closedAt)
# Un defecto se detecta el dia que se abre su issue y se resuelve el dia que se cierra. El
# dia se toma en la hora de Argentina (UTC-3): uno cerrado a las 22 h del ultimo dia del
# sprint es de ese sprint, aunque en UTC ya sea el dia siguiente.
echo "$DEFECTOS" | jq -r --arg desde "$DESDE" --arg hasta "$HASTA" '
  def dia: if . == null then null else (fromdate - 10800 | strftime("%Y-%m-%d")) end;
  def en_el_sprint: . != null and . >= $desde and . <= $hasta;
  [.[] | . + {detectado: (.createdAt | dia), resuelto: (.closedAt | dia)}] as $todos
  | [$todos[] | select(.detectado | en_el_sprint)] as $detectados
  | [$todos[] | select(.resuelto | en_el_sprint)] as $resueltos
  | [$todos[] | select(.detectado <= $hasta and (.resuelto == null or .resuelto > $hasta))] as $abiertos
  | "| Metrica | Valor |\n|---|---|"
  + "\n| Defectos detectados en el sprint | \($detectados | length) |"
  + "\n| Defectos resueltos en el sprint | \($resueltos | length) |"
  + "\n| Defectos abiertos al cierre | \($abiertos | length) |\n"
  + ($todos | sort_by(.number) | map("\n- #\(.number) detectado \(.detectado), "
      + "\(if .resuelto == null then "abierto" else "resuelto \(.resuelto)" end): \(.title)") | join(""))'

echo
echo "_La velocidad, la desviacion y las horas por story point solo cuentan historias en estado"
echo "Hecho: una historia a medias no suma story points, aunque tenga horas cargadas. Para un"
echo "sprint cerrado vale su acta de review: este script usa el Sprint que cada historia tiene hoy."
echo "Las horas de una historia que viene de un sprint anterior se suman enteras: lo que ya se"
echo "conto en ese sprint hay que restarlo a mano. Una historia terminada sin story points"
echo "suma sus horas reales al total, pero no a la desviacion ni a las horas por story point."
echo "Las adelantadas no suman a lo comprometido ni a sus horas; si se terminan, suman a la"
echo "velocidad, a la desviacion y a las horas por story point._"
