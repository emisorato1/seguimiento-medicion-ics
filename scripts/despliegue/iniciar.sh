#!/usr/bin/env bash
#
# Abre la web en el servidor. Lo ejecuta el servicio ics-web: no se usa a mano.
#
# Si la base no existe (la primera vez, o despues de un deploy, que la vuelve a empezar),
# la crea con los datos de demo. Si existe, la abre como esta.

set -eu

APP=/opt/ics/app
BASE=/opt/ics/datos/ics.db

if [ -f "$BASE" ]; then
  exec "$APP/web" -red -puerto "$PUERTO" -datos "$BASE"
fi
exec "$APP/web" -red -puerto "$PUERTO" -datos "$BASE" -importar "$APP/datos-demo.json"
