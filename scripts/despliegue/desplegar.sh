#!/usr/bin/env bash
#
# Pone en marcha en el servidor la version que el CD dejo en /opt/ics/nueva: el binario de
# la web, iniciar.sh y los datos de demo. Lo ejecuta el CD por SSH con el usuario ics.
#
# La base vuelve a empezar con los datos de demo en cada deploy: la web no tiene usuarios,
# asi que cualquiera con la direccion puede cambiar los datos, y asi cada version arranca
# limpia. La base anterior queda en ics.db.anterior.
#
# Si la version nueva no responde, vuelve a poner la anterior con su base y falla, para que
# el CD lo muestre en rojo.
#
# Uso (en el servidor): bash /opt/ics/nueva/desplegar.sh

set -eu

APP=/opt/ics/app
NUEVA=/opt/ics/nueva
DATOS=/opt/ics/datos
PUERTO="${PUERTO:-8080}"

responde() {
  for _ in $(seq 1 20); do
    curl -fsS -o /dev/null "http://127.0.0.1:$PUERTO/proyectos/1/dashboard" 2>/dev/null && return 0
    sleep 0.5
  done
  return 1
}

# La base se abre en modo exclusivo: hay que cerrar la web antes de moverla.
sudo systemctl stop ics-web.service

rm -rf "$APP.anterior"
cp -a "$APP" "$APP.anterior"
rm -f "$DATOS/ics.db.anterior"
if [ -f "$DATOS/ics.db" ]; then
  mv "$DATOS/ics.db" "$DATOS/ics.db.anterior"
fi

install -m 755 "$NUEVA/web" "$APP/web"
install -m 755 "$NUEVA/iniciar.sh" "$APP/iniciar.sh"
install -m 644 "$NUEVA/datos-demo.json" "$APP/datos-demo.json"

sudo systemctl start ics-web.service
if responde; then
  echo "La version nueva responde en el puerto $PUERTO."
  exit 0
fi

echo "La version nueva no responde: vuelve la anterior."
sudo systemctl stop ics-web.service
rm -rf "$APP"
mv "$APP.anterior" "$APP"
rm -f "$DATOS/ics.db"
if [ -f "$DATOS/ics.db.anterior" ]; then
  mv "$DATOS/ics.db.anterior" "$DATOS/ics.db"
fi
sudo systemctl start ics-web.service
exit 1
