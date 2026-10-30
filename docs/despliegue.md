# Integración y despliegue continuos (CI/CD)

Cómo se prueba cada cambio y cómo llega la web a un servidor donde se puede abrir desde
cualquier lado. Todo corre en GitHub Actions: los workflows están en `.github/workflows/` y
los scripts que usan, en `scripts/`.

## El recorrido de un cambio

1. **PR a `dev` o a `main` → CI** (`ci.yml`). Corre tres jobs:
   - *Formato, análisis y tests*: la coherencia de las specs, `gofmt`, `go vet`, los tests
     de Go y la cobertura. Es el único obligatorio para mergear.
   - *Tests en Windows*: los mismos tests en Windows (defecto #108).
   - *Prueba de humo de la web*: compila la web, la abre con los datos de demo en una base
     nueva y pide sus páginas principales (`scripts/prueba-humo.sh`). Los tests prueban cada
     parte por separado; la prueba de humo prueba que el programa armado arranca y atiende.
2. **Merge a `main` → CD** (`cd.yml`). `main` se mueve solo al cerrar un sprint, con el
   incremento. El CD:
   - corre los tests;
   - compila la consola y la web para Linux y para macOS, sin cgo;
   - le hace la prueba de humo al mismo binario que va a desplegar;
   - guarda los binarios como artefacto descargable de la corrida;
   - y **despliega la web en el servidor**. Si la versión nueva no responde, vuelve a poner
     la anterior y la corrida queda en rojo.

El CD también se puede correr a mano desde la pestaña *Actions* → *CD* → *Run workflow*,
siempre sobre `main`.

## El servidor

Es un servidor Linux (Ubuntu 24.04) de Maximiliano Eula, que lo compartimos con otra
aplicación suya. Por eso la web corre aparte y con límites:

- corre como **servicio de systemd** (`ics-web`), con el usuario `ics`, sin permisos de root;
- escucha en el **puerto 8080**; la otra aplicación usa el 80 y el 443, y no se toca;
- tiene un tope de **256 MB de memoria y medio procesador**, y solo puede escribir en su
  carpeta de datos;
- todo lo suyo está en `/opt/ics`: `app/` (la web), `datos/` (la base) y `nueva/` (donde el CD
  deja la versión que va a desplegar).

La dirección del servidor no está en el repo: está en los secrets de GitHub. Para tenerla,
pedírsela a Emiliano Sorato o a Maximiliano Eula.

**La base vuelve a empezar con los datos de demo en cada deploy.** La web no tiene usuarios
ni contraseñas (decisión D20 de la #77), así que cualquiera con la dirección puede cambiar los
datos: así cada versión arranca limpia. Entre un deploy y otro los datos se guardan, también
si el servidor se reinicia. La base anterior queda en `/opt/ics/datos/ics.db.anterior`.

## Cómo entra el CD al servidor

Con una **llave SSH que solo sirve para esto**: entra con el usuario `ics`, no puede abrir una
terminal ni reenviar puertos, y con `sudo` solo puede parar y arrancar el servicio `ics-web`.
No se usa la contraseña del servidor.

Los datos para entrar están en los secrets del **entorno `produccion`** del repo, que solo se
puede usar desde `main`:

| Secret | Qué es |
|---|---|
| `DEPLOY_HOST` | la dirección del servidor |
| `DEPLOY_KEY` | la llave privada del CD |
| `DEPLOY_KNOWN_HOSTS` | la huella del servidor, para que el CD no se conecte a otro que se haga pasar por él |

## Preparar el servidor (una sola vez)

Ya está hecho. Si hubiera que repetirlo en otro servidor, como root:

```bash
scp scripts/despliegue/instalar.sh scripts/despliegue/ics-web.service root@SERVIDOR:/tmp/
ssh root@SERVIDOR 'bash /tmp/instalar.sh "<llave pública del CD>"'
```

`instalar.sh` crea el usuario `ics` y las carpetas, instala el servicio, le deja a `ics`
pararlo y arrancarlo, y autoriza la llave del CD. Después hay que cargar los tres secrets y
correr el CD: el primer deploy pone la web en marcha.

## Si algo sale mal

| Qué pasa | Qué hacer |
|---|---|
| El CD queda en rojo en *Desplegar en el servidor* | el servidor sigue con la versión anterior; mirar el log de *Ponerla en marcha* |
| La web del servidor quedó con datos raros | correr el CD a mano: la base vuelve a los datos de demo |
| El CD no puede entrar al servidor | revisar que los tres secrets del entorno `produccion` estén cargados |
