#!/usr/bin/env bash
#
# Prepara el servidor para el CD. Se corre una sola vez, como root, en el servidor:
#
#   - crea el usuario ics, que es el unico con el que entra el CD;
#   - crea /opt/ics con sus carpetas: app (la web), datos (la base) y nueva (donde el CD
#     deja la version que va a desplegar);
#   - instala el servicio ics-web y le deja al usuario ics pararlo y arrancarlo, nada mas;
#   - autoriza la llave publica del CD para entrar como ics.
#
# No toca nada de lo que ya hay en el servidor: ni Docker ni los puertos 80 y 443.
#
# Uso: bash instalar.sh "<llave publica del CD>"
#      (con ics-web.service en la misma carpeta que este script)

set -eu

LLAVE="${1:?Uso: bash instalar.sh \"<llave publica del CD>\"}"
AQUI=$(cd "$(dirname "$0")" && pwd)

if ! id ics >/dev/null 2>&1; then
  useradd --system --home-dir /opt/ics --create-home --shell /bin/bash ics
fi
install -d -o ics -g ics -m 755 /opt/ics/app /opt/ics/datos /opt/ics/nueva
install -d -o ics -g ics -m 700 /opt/ics/.ssh

# La llave solo sirve para correr comandos y copiar archivos: sin reenvio de puertos ni
# terminal interactiva.
echo "restrict $LLAVE" >/opt/ics/.ssh/authorized_keys
chown ics:ics /opt/ics/.ssh/authorized_keys
chmod 600 /opt/ics/.ssh/authorized_keys

cat >/etc/sudoers.d/ics-web <<'EOF'
ics ALL=(root) NOPASSWD: /usr/bin/systemctl stop ics-web.service, /usr/bin/systemctl start ics-web.service
EOF
chmod 440 /etc/sudoers.d/ics-web
visudo -cf /etc/sudoers.d/ics-web

install -m 644 "$AQUI/ics-web.service" /etc/systemd/system/ics-web.service
systemctl daemon-reload
systemctl enable ics-web.service

echo "Listo. El primer deploy pone la web en marcha."
