// Command web abre el sistema en el navegador: carga los datos guardados, arma la web y
// muestra la direccion donde se usa (RN-2).
//
//	go run ./cmd/web                  en http://127.0.0.1:8080, con datos-ics.json
//	go run ./cmd/web -puerto 9090     en otro puerto
//	go run ./cmd/web -datos otro.json con otro archivo de datos
//	go run ./cmd/web -red             acepta conexiones de la red local, para un celular
package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/store"
	"github.com/emisorato1/seguimiento-medicion-ics/internal/web"
)

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

	showAddress(os.Stdout, opts)
	if err := http.ListenAndServe(opts.addr(), handler); err != nil {
		fmt.Fprintln(os.Stderr, "La web se detuvo:", err)
		os.Exit(1)
	}
}
