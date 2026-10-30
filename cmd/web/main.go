// Command web abre el sistema en el navegador: carga los datos guardados, arma la web y
// muestra la direccion donde se usa (RN-2).
//
//	go run ./cmd/web                          en http://127.0.0.1:8080, con la base datos-ics.db
//	go run ./cmd/web -puerto 9090             en otro puerto
//	go run ./cmd/web -datos otra.db           con otra base
//	go run ./cmd/web -red                     acepta conexiones de la red local, para un celular
//	go run ./cmd/web -importar datos.json     trae el archivo a una base vacía y abre
//	go run ./cmd/web -exportar copia.json     escribe todos los datos en el archivo y termina
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/store"
	"github.com/emisorato1/seguimiento-medicion-ics/internal/web"
)

// errPortInUse es el aviso de que otro programa ya usa el puerto.
var errPortInUse = errors.New("ya lo está usando otro programa")

// errPortNotAllowed es el aviso de que el sistema no deja usar el puerto.
var errPortNotAllowed = errors.New("no se puede usar: el sistema no da permiso")

// Windows avisa el puerto ocupado y la falta de permiso con sus codigos de Winsock, 10048 y
// 10013, que no son syscall.EADDRINUSE ni syscall.EACCES. Esos numeros no los usa ningun
// otro sistema, asi que se reconocen en todos, y el test los prueba tambien en Linux y en
// macOS (revision de Maxi).
const (
	wsaeaddrinuse = syscall.Errno(10048)
	wsaeacces     = syscall.Errno(10013)
)

// options son las opciones del comando (D20).
type options struct {
	port       int
	dataPath   string
	lan        bool
	importPath string
	exportPath string
}

// parseOptions lee las opciones de la linea de comandos. Una opcion que no existe es un
// error, que el comando informa al terminar.
func parseOptions(args []string) (options, error) {
	var opts options
	flags := flag.NewFlagSet("web", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	flags.IntVar(&opts.port, "puerto", 8080, "puerto donde escucha la web")
	flags.StringVar(&opts.dataPath, "datos", "datos-ics.db", "base donde se guardan los datos")
	flags.BoolVar(&opts.lan, "red", false, "aceptar conexiones de otras computadoras de la red local")
	flags.StringVar(&opts.importPath, "importar", "", "archivo con datos para traer a la base al abrir")
	flags.StringVar(&opts.exportPath, "exportar", "", "archivo donde escribir todos los datos, sin abrir la web")
	if err := flags.Parse(args); err != nil {
		return options{}, fmt.Errorf("opciones: %v (las que existen son -puerto, -datos, -red, -importar y -exportar)", err)
	}
	// Juntas, no queda claro que se quiere hacer primero: se rechaza sin hacer ninguna (RN-17).
	if opts.importPath != "" && opts.exportPath != "" {
		return options{}, errors.New("opciones: -importar y -exportar no se pueden usar juntas: importá primero y exportá después")
	}
	// Antes de la #177, -datos era el archivo .json. Usado asi, SQLite fallaba con el archivo
	// o creaba en silencio una base con ese nombre: se avisa sin crear nada (RN-20).
	if strings.EqualFold(filepath.Ext(opts.dataPath), ".json") {
		return options{}, errors.New("opciones: -datos es la base de datos (.db); para traer un archivo .json usá -importar")
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
	case errors.Is(err, syscall.EADDRINUSE), errors.Is(err, wsaeaddrinuse):
		return fmt.Errorf("el puerto %d %w. Probá con otro, por ejemplo: go run ./cmd/web -puerto %d", port, errPortInUse, port+1)
	case errors.Is(err, syscall.EACCES), errors.Is(err, wsaeacces):
		return fmt.Errorf("el puerto %d %w (puede ser menor que 1024 o estar reservado). Probá con otro, por ejemplo: go run ./cmd/web -puerto %d", port, errPortNotAllowed, port+1)
	default:
		return fmt.Errorf("no se pudo usar el puerto %d (%v). Probá con otro con -puerto", port, err)
	}
}

// open abre la base, carga los datos guardados y arma la web, que guarda en la misma base
// (historia #177, D8). Con -importar, antes trae a la base los datos del archivo (RN-8).
// Devuelve tambien la base, que quien llama cierra al terminar. Si los datos no se pueden
// leer devuelve el error y no toca la base (RN-24 de la #77, RN-4).
func open(o options, out io.Writer) (http.Handler, *store.Database, error) {
	db, err := store.OpenDatabase(o.dataPath)
	if err != nil {
		return nil, nil, err
	}
	if o.importPath != "" {
		if err := store.Import(db, o.importPath); err != nil {
			db.Close()
			return nil, nil, err
		}
		fmt.Fprintf(out, "Se importaron los datos de %s\n", o.importPath)
	} else {
		// Sin -importar, lo que guardo la version anterior junto a la base se importa solo,
		// una sola vez (RN-14, RN-15).
		imported, err := store.ImportLegacy(db, filepath.Dir(o.dataPath))
		if err != nil {
			db.Close()
			return nil, nil, err
		}
		if imported {
			legacy := filepath.Join(filepath.Dir(o.dataPath), store.LegacyFile)
			fmt.Fprintf(out, "Se importaron los datos de %s y el archivo se renombró a %s.importado\n", legacy, legacy)
		}
	}
	registry, err := db.Load()
	if err != nil {
		db.Close()
		return nil, nil, err
	}
	return web.NewServer(registry, db, time.Now), db, nil
}

// export escribe todos los datos de la base en el archivo de -exportar y termina sin abrir la
// web (RN-7).
func export(o options, out, errOut io.Writer) int {
	// Sin base no hay nada que exportar: abrirla la crearia vacia (RN-18).
	if _, err := os.Stat(o.dataPath); errors.Is(err, fs.ErrNotExist) {
		fmt.Fprintf(errOut, "No se pudieron exportar los datos: no hay datos para exportar, porque la base %s no existe.\n", o.dataPath)
		return 1
	}
	db, err := store.OpenDatabase(o.dataPath)
	// Con la web abierta sobre la misma base, lo que hay que hacer es cerrarla y volver a
	// exportar, no "volver a abrir este" (RN-19).
	if errors.Is(err, store.ErrDatabaseInUse) {
		fmt.Fprintf(errOut, "No se pudieron exportar los datos: la base %s ya está abierta por la web. Cerrá la web y volvé a exportar.\n", o.dataPath)
		return 1
	}
	if err != nil {
		fmt.Fprintf(errOut, "No se pudieron exportar los datos: %v\nLa base %s quedó como estaba.\n", err, o.dataPath)
		return 1
	}
	defer db.Close()
	if err := store.Export(db, o.exportPath); err != nil {
		fmt.Fprintln(errOut, "No se pudieron exportar los datos:", err)
		return 1
	}
	fmt.Fprintf(out, "Se exportaron los datos a %s\n", o.exportPath)
	return 0
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
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}
	return localIPv4(addrs)
}

// localIPv4 deja, de las direcciones de la computadora, las IPv4 que no son de ella misma.
func localIPv4(addrs []net.Addr) []string {
	var ips []string
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ips = append(ips, ipNet.IP.String())
		}
	}
	return ips
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr, http.Serve))
}

// run abre el sistema y devuelve el codigo con el que termina el comando. La web la atiende
// serve, que en el comando es http.Serve y en los tests una funcion que no se queda esperando.
func run(args []string, out, errOut io.Writer, serve func(net.Listener, http.Handler) error) int {
	opts, err := parseOptions(args)
	if err != nil {
		fmt.Fprintln(errOut, "No se pudo abrir el sistema:", err)
		return 2
	}
	if opts.exportPath != "" {
		return export(opts, out, errOut)
	}
	handler, db, err := open(opts, out)
	if err != nil {
		fmt.Fprintf(errOut, "No se pudo abrir el sistema: %v\nLa base %s quedó como estaba.\n", err, opts.dataPath)
		return 1
	}
	// En Windows una base abierta no se puede borrar ni mover: se cierra en todos los caminos.
	defer db.Close()

	// La direccion se muestra recien cuando el puerto es nuestro: si no, se mostraria una
	// direccion donde atiende otro programa.
	listener, err := listen(opts)
	if err != nil {
		fmt.Fprintln(errOut, "No se pudo abrir el sistema:", err)
		return 1
	}
	showAddress(out, opts)
	if err := serve(listener, handler); err != nil {
		fmt.Fprintln(errOut, "La web se detuvo:", err)
		return 1
	}
	return 0
}
