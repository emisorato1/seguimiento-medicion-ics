# Imagen de la web para Azure Container Apps. No compila: usa el mismo binario de Linux que
# arma el job "publicar" del CD (sin cgo, no depende de librerias del sistema). La base de
# datos vive adentro del contenedor y arranca cada vez con los datos de demo.
FROM alpine:3.20

RUN adduser -D -H ics && mkdir /datos && chown ics /datos

COPY web datos-demo.json /app/

USER ics
EXPOSE 8080
CMD ["/app/web", "-red", "-puerto", "8080", "-datos", "/datos/ics.db", "-importar", "/app/datos-demo.json"]
