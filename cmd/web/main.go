// Command web abre el sistema en el navegador: carga los datos guardados, arma la web y
// muestra la direccion donde se usa (RN-2).
//
//	go run ./cmd/web                  en http://127.0.0.1:8080, con datos-ics.json
//	go run ./cmd/web -puerto 9090     en otro puerto
//	go run ./cmd/web -datos otro.json con otro archivo de datos
//	go run ./cmd/web -red             acepta conexiones de la red local, para un celular
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"syscall"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/store"
	"github.com/emisorato1/seguimiento-medicion-ics/internal/web"
)

// errPortInUse es el aviso de que otro programa ya usa el puerto.
var errPortInUse = errors.New("ya lo está usando otro programa")

// options son las opciones del comando (D20).
type options struct {
	port     int
	dataPath string
	lan      bool
}

// parseOptions lee las opciones de la linea de comandos. Una opcion que no existe es un
// error, que el comando informa al terminar.
func parseOptions(args []string) (options, error) {
	var opts options
	flags := flag.NewFlagSet("web", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	flags.IntVar(&opts.port, "puerto", 8080, "puerto donde escucha la web")
	flags.StringVar(&opts.dataPath, "datos", "datos-ics.json", "archivo donde se guardan los datos")
	flags.BoolVar(&opts.lan, "red", false, "aceptar conexiones de otras computadoras de la red local")
	if err := flags.Parse(args); err != nil {
		return options{}, fmt.Errorf("opciones: %v (las que existen son -puerto, -datos y -red)", err)
	}
	return opts, nil
}

// addr es donde escucha la web: solo en esta computadora, salvo con -red. Sin usuarios ni
// contraseñas, cualquiera en la misma red podria cambiar los datos (D20).
func (o options) addr() string {
	if o.lan {
		return fmt.Sprintf(":%d", o.port)
	}
	return fmt.Sprintf("127.0.0.1:%d", o.port)
}

// listen toma el puerto donde va a escuchar la web. Antes prueba el puerto en todas las
// direcciones: en macOS, otro programa que escucha en todas —como Docker en el 8080— no
// impide escuchar solo en 127.0.0.1, y el navegador podia terminar en el otro programa sin
// ningun aviso (hallazgo de la verificacion manual).
func listen(o options) (net.Listener, error) {
	probe, err := net.Listen("tcp", fmt.Sprintf(":%d", o.port))
	if err != nil {
		return nil, portError(o.port, err)
	}
	probe.Close()

	listener, err := net.Listen("tcp", o.addr())
	if err != nil {
		return nil, portError(o.port, err)
	}
	return listener, nil
}

// portError explica en castellano por que no se pudo tomar el puerto, y que hacer.
func portError(port int, err error) error {
	switch {
	case errors.Is(err, syscall.EADDRINUSE):
		return fmt.Errorf("el puerto %d %w. Probá con otro, por ejemplo: go run ./cmd/web -puerto %d", port, errPortInUse, port+1)
	case errors.Is(err, syscall.EACCES):
		return fmt.Errorf("no hay permiso para usar el puerto %d. Probá con uno mayor que 1024, por ejemplo: go run ./cmd/web -puerto 8081", port)
	default:
		return fmt.Errorf("no se pudo usar el puerto %d (%v). Probá con otro con -puerto", port, err)
	}
}

// open carga los datos guardados y arma la web, que guarda en el mismo archivo. Si los datos
// no se pueden leer devuelve el error y no toca el archivo (RN-24, D18).
func open(o options) (http.Handler, error) {
	file := store.NewFile(o.dataPath)
	registry, err := file.Load()
	if err != nil {
		return nil, err
	}
	return web.NewServer(registry, file), nil
}

// showAddress muestra donde abrir el sistema en el navegador y donde se guardan los datos.
func showAddress(out io.Writer, o options) {
	fmt.Fprintf(out, "Seguimiento y Medición ICS\n")
	fmt.Fprintf(out, "Abrilo en el navegador: http://127.0.0.1:%d\n", o.port)
	if o.lan {
		for _, ip := range lanAddresses() {
			fmt.Fprintf(out, "Desde otra computadora o un celular de la red: http://%s:%d\n", ip, o.port)
		}
	}
	fmt.Fprintf(out, "Los datos se guardan en %s\n", o.dataPath)
	fmt.Fprintf(out, "Para cerrarlo, Ctrl+C.\n")
}

// lanAddresses son las direcciones IPv4 de esta computadora en la red local.
func lanAddresses() []string {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips
	}
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ips = append(ips, ipNet.IP.String())
		}
	}
	return ips
}

func main() {
	opts, err := parseOptions(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "No se pudo abrir el sistema:", err)
		os.Exit(2)
	}
	handler, err := open(opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "No se pudo abrir el sistema: %v\nEl archivo %s quedó como estaba.\n", err, opts.dataPath)
		os.Exit(1)
	}

	// La direccion se muestra recien cuando el puerto es nuestro: si no, se mostraria una
	// direccion donde atiende otro programa.
	listener, err := listen(opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "No se pudo abrir el sistema:", err)
		os.Exit(1)
	}
	showAddress(os.Stdout, opts)
	if err := http.Serve(listener, handler); err != nil {
		fmt.Fprintln(os.Stderr, "La web se detuvo:", err)
		os.Exit(1)
	}
}
