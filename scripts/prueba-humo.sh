#!/usr/bin/env bash
#
# Prueba de humo de la web: abre el binario ya compilado con los datos de demo en una base
# nueva y pide las paginas principales, como lo haria alguien desde el navegador. Si alguna
# no responde o no muestra lo que tiene que mostrar, falla. Los tests de Go prueban cada
# parte por separado; esto prueba que el programa armado arranca y atiende.
#
# Corre en CI con cada PR y en el CD antes de desplegar. No toca la red: solo 127.0.0.1.
#
# Uso: scripts/prueba-humo.sh <binario de la web> [puerto]
#      scripts/prueba-humo.sh dist/web-linux-amd64 8099

set -u

BINARIO="${1:?Uso: scripts/prueba-humo.sh <binario de la web> [puerto]}"
PUERTO="${2:-8099}"
BASE_URL="http://127.0.0.1:$PUERTO"
DATOS_DEMO="docs/demo/datos-demo.json"
FALLAS=0

carpeta=$(mktemp -d)
"$BINARIO" -datos "$carpeta/humo.db" -importar "$DATOS_DEMO" -puerto "$PUERTO" >"$carpeta/salida.txt" 2>&1 &
pid=$!
trap 'kill "$pid" 2>/dev/null; wait "$pid" 2>/dev/null; rm -rf "$carpeta"' EXIT

# La web tarda un momento en importar los datos y tomar el puerto.
for _ in $(seq 1 20); do
  curl -fsS -o /dev/null "$BASE_URL/proyectos" 2>/dev/null && break
  if ! kill -0 "$pid" 2>/dev/null; then
    echo "La web se cerro al arrancar:"
    cat "$carpeta/salida.txt"
    exit 1
  fi
  sleep 0.5
done

# pagina <ruta> <texto que tiene que aparecer>
pagina() {
  local cuerpo
  if ! cuerpo=$(curl -fsS "$BASE_URL$1" 2>&1); then
    echo "  FALLA: $1 no respondio ($cuerpo)"
    FALLAS=$((FALLAS + 1))
  elif ! grep -qF -e "$2" <<<"$cuerpo"; then
    echo "  FALLA: $1 no muestra «$2»"
    FALLAS=$((FALLAS + 1))
  else
    echo "  ok: $1"
  fi
}

echo "Prueba de humo de $BINARIO en $BASE_URL"
pagina /proyectos "Seguimiento y Medición ICS"
pagina /proyectos/1 "Seguimiento y Medición ICS"
pagina /proyectos/1/dashboard "Estado general"
pagina /proyectos/1/integrantes "Agile Enabler"
pagina /proyectos/1/backlog "Votar en secreto en una ronda de Planning Poker"
pagina /proyectos/1/historias/13 "Votar en secreto en una ronda de Planning Poker"
pagina /proyectos/1/sprints/1 "Sprint 1"
pagina /proyectos/1/planning-poker "Planning Poker"
pagina /proyectos/1/historias/77/rondas/1 "Hay diferencias"
pagina /proyectos/1/esfuerzo "Totales por historia"
pagina /proyectos/1/defectos "Defectos"
pagina /proyectos/1/reporte "Reporte del proyecto"
pagina "/proyectos/1/reporte?sprint=1" "Reporte del sprint 1"
pagina "/proyectos/1/reporte/pdf?sprint=1" "%PDF-1.4"
pagina /static/app.css ".sidebar"

if [ "$FALLAS" -gt 0 ]; then
  echo "$FALLAS paginas fallaron."
  exit 1
fi
echo "Todas las paginas responden."
