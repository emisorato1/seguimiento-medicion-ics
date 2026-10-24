package main

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"testing"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// Escenario: Abrir el sistema con un solo comando y sin internet
// Cubre: RN-2. Esta es la parte del comando; que la pagina no pida nada de afuera lo prueba
// internal/web.
func TestAbrirElSistemaConUnSoloComandoYSinInternet(t *testing.T) {
	t.Run("sin opciones escucha solo en esta computadora, en el 8080", func(t *testing.T) {
		opts, err := parseOptions(nil)
		if err != nil {
			t.Fatalf("no se esperaba error, salio: %v", err)
		}
		if opts.addr() != "127.0.0.1:8080" {
			t.Errorf("se esperaba escuchar en 127.0.0.1:8080, salio %q", opts.addr())
		}
		if opts.dataPath != "datos-ics.json" {
			t.Errorf("se esperaban los datos en datos-ics.json, salio %q", opts.dataPath)
		}
	})

	t.Run("respeta -puerto, -datos y -red", func(t *testing.T) {
		opts, err := parseOptions([]string{"-puerto", "9090", "-datos", "otros.json", "-red"})
		if err != nil {
			t.Fatalf("no se esperaba error, salio: %v", err)
		}
		if opts.addr() != ":9090" {
			t.Errorf("con -red se esperaba escuchar en toda la red en el 9090, salio %q", opts.addr())
		}
		if opts.dataPath != "otros.json" {
			t.Errorf("se esperaban los datos en otros.json, salio %q", opts.dataPath)
		}
	})

	t.Run("una opcion que no existe es un error", func(t *testing.T) {
		if _, err := parseOptions([]string{"-color"}); err == nil {
			t.Fatal("se esperaba un error por la opcion -color, y no salio ninguno")
		}
	})

	t.Run("muestra la direccion para abrirlo en el navegador", func(t *testing.T) {
		opts, _ := parseOptions([]string{"-puerto", "9090"})
		var out bytes.Buffer

		showAddress(&out, opts)

		if !strings.Contains(out.String(), "http://127.0.0.1:9090") {
			t.Errorf("se esperaba la direccion http://127.0.0.1:9090, salio:\n%s", out.String())
		}
	})

	// Hallazgo de la verificacion manual: con Docker en el 8080, la web arrancaba igual en
	// 127.0.0.1 sin avisar y el navegador mostraba la otra aplicacion. Se prueba con el
	// puerto tomado en todas las direcciones, como lo toma Docker, y solo en 127.0.0.1.
	for _, ocupado := range []string{":0", "127.0.0.1:0"} {
		t.Run("con el puerto ocupado en "+ocupado+" avisa en castellano y sugiere -puerto", func(t *testing.T) {
			otro, err := net.Listen("tcp", ocupado)
			if err != nil {
				t.Fatalf("no se pudo ocupar el puerto para la prueba: %v", err)
			}
			defer otro.Close()
			puerto := otro.Addr().(*net.TCPAddr).Port
			opts, _ := parseOptions([]string{"-puerto", fmt.Sprint(puerto)})

			listener, err := listen(opts)

			if listener != nil {
				listener.Close()
			}
			if !errors.Is(err, errPortInUse) {
				t.Fatalf("se esperaba errPortInUse, salio: %v", err)
			}
			if !strings.Contains(err.Error(), "-puerto") || strings.Contains(err.Error(), "address already in use") {
				t.Errorf("se esperaba el aviso en castellano con la sugerencia de -puerto, salio: %v", err)
			}
		})
	}

	// Revision del PR #95 en Windows: ahi el puerto ocupado llega con el codigo de Winsock
	// WSAEADDRINUSE (10048), que no es syscall.EADDRINUSE, y el aviso salia con el texto del
	// sistema. Se arma el error como lo devuelve Windows, asi se prueba en cualquier sistema.
	t.Run("con el puerto ocupado en Windows tambien avisa en castellano", func(t *testing.T) {
		deWindows := &net.OpError{Op: "listen", Net: "tcp", Err: os.NewSyscallError("bind", syscall.Errno(10048))}

		err := portError(8080, deWindows)

		if !errors.Is(err, errPortInUse) {
			t.Fatalf("se esperaba errPortInUse, salio: %v", err)
		}
		if strings.Contains(err.Error(), syscall.Errno(10048).Error()) {
			t.Errorf("no se esperaba el texto del sistema en el aviso, salio: %v", err)
		}
	})

	// Lo mismo con el permiso: Windows lo avisa con WSAEACCES (10013), por ejemplo cuando el
	// puerto esta reservado por el sistema.
	for _, sinPermiso := range []syscall.Errno{syscall.EACCES, syscall.Errno(10013)} {
		t.Run("sin permiso para el puerto ("+sinPermiso.Error()+") avisa en castellano", func(t *testing.T) {
			err := portError(8080, &net.OpError{Op: "listen", Net: "tcp", Err: os.NewSyscallError("bind", sinPermiso)})

			if !errors.Is(err, errPortNotAllowed) {
				t.Fatalf("se esperaba errPortNotAllowed, salio: %v", err)
			}
			if strings.Contains(err.Error(), sinPermiso.Error()) || !strings.Contains(err.Error(), "-puerto") {
				t.Errorf("se esperaba el aviso en castellano con la sugerencia de -puerto, salio: %v", err)
			}
		})
	}

	// Con -red se muestran las direcciones de la red local para abrirlo desde un celular. De
	// las direcciones de la computadora quedan solo las IPv4 que no son de ella misma.
	t.Run("con -red muestra solo las direcciones IPv4 de la red local", func(t *testing.T) {
		addrs := []net.Addr{
			&net.IPNet{IP: net.ParseIP("127.0.0.1"), Mask: net.CIDRMask(8, 32)},
			&net.IPNet{IP: net.ParseIP("192.168.1.20"), Mask: net.CIDRMask(24, 32)},
			&net.IPNet{IP: net.ParseIP("fe80::1"), Mask: net.CIDRMask(64, 128)},
			&net.IPNet{IP: net.ParseIP("::1"), Mask: net.CIDRMask(128, 128)},
			&net.IPAddr{IP: net.ParseIP("10.0.0.5")},
			&net.IPNet{IP: net.ParseIP("10.0.0.7"), Mask: net.CIDRMask(8, 32)},
		}

		ips := localIPv4(addrs)

		if strings.Join(ips, " ") != "192.168.1.20 10.0.0.7" {
			t.Errorf("se esperaban 192.168.1.20 y 10.0.0.7, salio %v", ips)
		}
	})

	t.Run("con -red muestra una linea por cada direccion de la red local", func(t *testing.T) {
		opts, _ := parseOptions([]string{"-puerto", "9090", "-red"})
		var out bytes.Buffer

		showAddress(&out, opts)

		lineas := strings.Count(out.String(), "Desde otra computadora o un celular de la red: http://")
		if lineas != len(lanAddresses()) {
			t.Errorf("se esperaba una linea por cada una de %v, salio:\n%s", lanAddresses(), out.String())
		}
		if !strings.Contains(out.String(), "http://127.0.0.1:9090") {
			t.Errorf("se esperaba tambien la direccion local, salio:\n%s", out.String())
		}
	})

	t.Run("una opcion que no existe cierra el comando con el codigo 2 y sin abrir nada", func(t *testing.T) {
		var out, errOut bytes.Buffer

		code := run([]string{"-color"}, &out, &errOut, noServe(t))

		if code != 2 {
			t.Errorf("se esperaba el codigo 2, salio %d", code)
		}
		if !strings.Contains(errOut.String(), "No se pudo abrir el sistema:") || out.Len() != 0 {
			t.Errorf("se esperaba solo el aviso en la salida de errores, salio:\n%s\n%s", out.String(), errOut.String())
		}
	})

	t.Run("con el puerto ocupado no muestra ninguna direccion y cierra con el codigo 1", func(t *testing.T) {
		otro, err := net.Listen("tcp", ":0")
		if err != nil {
			t.Fatalf("no se pudo ocupar el puerto para la prueba: %v", err)
		}
		defer otro.Close()
		puerto := fmt.Sprint(otro.Addr().(*net.TCPAddr).Port)
		var out, errOut bytes.Buffer

		code := run([]string{"-puerto", puerto, "-datos", filepath.Join(t.TempDir(), "datos-ics.json")}, &out, &errOut, noServe(t))

		if code != 1 {
			t.Errorf("se esperaba el codigo 1, salio %d", code)
		}
		if out.Len() != 0 || !strings.Contains(errOut.String(), errPortInUse.Error()) {
			t.Errorf("se esperaba solo el aviso del puerto, salio:\n%s\n%s", out.String(), errOut.String())
		}
	})

	t.Run("con todo en orden muestra la direccion y atiende ahi la web", func(t *testing.T) {
		var out, errOut bytes.Buffer
		var atendio string
		serve := func(listener net.Listener, handler http.Handler) error {
			defer listener.Close()
			atendio = listener.Addr().String()
			res := httptest.NewRecorder()
			handler.ServeHTTP(res, httptest.NewRequest(http.MethodGet, "/proyectos", nil))
			if res.Code != http.StatusOK {
				t.Errorf("se esperaba la lista de proyectos con el estado %d, salio %d", http.StatusOK, res.Code)
			}
			return http.ErrServerClosed
		}

		code := run([]string{"-puerto", "0", "-datos", filepath.Join(t.TempDir(), "datos-ics.json")}, &out, &errOut, serve)

		if !strings.HasPrefix(atendio, "127.0.0.1:") {
			t.Errorf("se esperaba atender en 127.0.0.1, salio %q", atendio)
		}
		if !strings.Contains(out.String(), "Abrilo en el navegador: http://127.0.0.1:0") {
			t.Errorf("se esperaba la direccion para el navegador, salio:\n%s", out.String())
		}
		if code != 1 || !strings.Contains(errOut.String(), "La web se detuvo:") {
			t.Errorf("cuando la web se detiene se esperaba el codigo 1 y el aviso, salio %d:\n%s", code, errOut.String())
		}
	})

	t.Run("con el puerto libre escucha en esa direccion", func(t *testing.T) {
		opts, _ := parseOptions([]string{"-puerto", "0"})

		listener, err := listen(opts)

		if err != nil {
			t.Fatalf("no se esperaba error con un puerto libre, salio: %v", err)
		}
		defer listener.Close()
		if ip := listener.Addr().(*net.TCPAddr).IP.String(); ip != "127.0.0.1" {
			t.Errorf("se esperaba escuchar solo en 127.0.0.1, salio %s", ip)
		}
	})

	t.Run("en esa direccion la lista de proyectos se ve completa, con sus estilos", func(t *testing.T) {
		opts, _ := parseOptions([]string{"-datos", filepath.Join(t.TempDir(), "datos-ics.json")})

		handler, err := open(opts)
		if err != nil {
			t.Fatalf("no se esperaba error al abrir sin datos guardados, salio: %v", err)
		}

		for _, pagina := range []string{"/proyectos", "/static/app.css"} {
			res := httptest.NewRecorder()
			handler.ServeHTTP(res, httptest.NewRequest(http.MethodGet, pagina, nil))
			if res.Code != http.StatusOK {
				t.Errorf("se esperaba %s con el estado %d, salio %d", pagina, http.StatusOK, res.Code)
			}
		}
	})
}

// Escenario: Los datos guardados estan danados
// Cubre: RN-24
func TestLosDatosGuardadosEstanDanados(t *testing.T) {
	path := filepath.Join(t.TempDir(), "datos-ics.json")
	danado := []byte(`{"next_id": 2, "projects": [{"id": 1, "name": ""`)
	if err := os.WriteFile(path, danado, 0o644); err != nil {
		t.Fatalf("no se pudo preparar el archivo danado: %v", err)
	}
	opts, _ := parseOptions([]string{"-datos", path})

	_, err := open(opts)

	if !errors.Is(err, domain.ErrSavedDataInvalid) {
		t.Fatalf("se esperaba ErrSavedDataInvalid, salio: %v", err)
	}
	quedo, _ := os.ReadFile(path)
	if !bytes.Equal(quedo, danado) {
		t.Errorf("los datos guardados tenian que quedar como estaban, y cambiaron a:\n%s", quedo)
	}

	t.Run("el comando avisa que el archivo quedo como estaba y cierra con el codigo 1", func(t *testing.T) {
		var out, errOut bytes.Buffer

		code := run([]string{"-datos", path}, &out, &errOut, noServe(t))

		if code != 1 {
			t.Errorf("se esperaba el codigo 1, salio %d", code)
		}
		if !strings.Contains(errOut.String(), "El archivo "+path+" quedó como estaba.") || out.Len() != 0 {
			t.Errorf("se esperaba solo el aviso del archivo, salio:\n%s\n%s", out.String(), errOut.String())
		}
	})
}

// noServe es la web que no se tiene que llegar a atender: el comando se cierra antes.
func noServe(t *testing.T) func(net.Listener, http.Handler) error {
	return func(listener net.Listener, _ http.Handler) error {
		listener.Close()
		t.Error("no se esperaba que la web llegara a atender")
		return nil
	}
}
