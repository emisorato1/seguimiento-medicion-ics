package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
	"github.com/emisorato1/seguimiento-medicion-ics/internal/store"
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
		// Desde la #177 los datos van a una base SQLite (D8).
		if opts.dataPath != "datos-ics.db" {
			t.Errorf("se esperaban los datos en datos-ics.db, salio %q", opts.dataPath)
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

		code := run([]string{"-puerto", puerto, "-datos", filepath.Join(t.TempDir(), "datos-ics.db")}, &out, &errOut, noServe(t))

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

		code := run([]string{"-puerto", "0", "-datos", filepath.Join(t.TempDir(), "datos-ics.db")}, &out, &errOut, serve)

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
		opts, _ := parseOptions([]string{"-datos", filepath.Join(t.TempDir(), "datos-ics.db")})

		handler, db, err := open(opts, io.Discard)
		if err != nil {
			t.Fatalf("no se esperaba error al abrir sin datos guardados, salio: %v", err)
		}
		defer db.Close()

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
// Cubre: RN-24; #177 RN-4 (la base dañada)
func TestLosDatosGuardadosEstanDanados(t *testing.T) {
	path := filepath.Join(t.TempDir(), "datos-ics.db")
	danado := []byte("no soy una base")
	if err := os.WriteFile(path, danado, 0o644); err != nil {
		t.Fatalf("no se pudo preparar el archivo danado: %v", err)
	}
	opts, _ := parseOptions([]string{"-datos", path})

	_, _, err := open(opts, io.Discard)

	if !errors.Is(err, domain.ErrSavedDataInvalid) {
		t.Fatalf("se esperaba ErrSavedDataInvalid, salio: %v", err)
	}
	quedo, _ := os.ReadFile(path)
	if !bytes.Equal(quedo, danado) {
		t.Errorf("los datos guardados tenian que quedar como estaban, y cambiaron a:\n%s", quedo)
	}

	t.Run("el comando avisa que la base quedo como estaba y cierra con el codigo 1", func(t *testing.T) {
		var out, errOut bytes.Buffer

		code := run([]string{"-datos", path}, &out, &errOut, noServe(t))

		if code != 1 {
			t.Errorf("se esperaba el codigo 1, salio %d", code)
		}
		if !strings.Contains(errOut.String(), "La base "+path+" quedó como estaba.") || out.Len() != 0 {
			t.Errorf("se esperaba solo el aviso de la base, salio:\n%s\n%s", out.String(), errOut.String())
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

// Escenario: Abrir por primera vez
// Cubre: #177 RN-1, RN-5, D8
func TestLaWebGuardaEnLaBase(t *testing.T) {
	path := filepath.Join(t.TempDir(), "datos-ics.db")
	var out, errOut bytes.Buffer
	serve := func(listener net.Listener, handler http.Handler) error {
		listener.Close()
		form := url.Values{"nombre": {"Sistema de métricas"}, "inicio": {"2026-10-01"}}
		req := httptest.NewRequest(http.MethodPost, "/proyectos/nuevo", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		res := httptest.NewRecorder()
		handler.ServeHTTP(res, req)
		if res.Code != http.StatusSeeOther {
			t.Errorf("se esperaba crear el proyecto con el estado %d, salio %d", http.StatusSeeOther, res.Code)
		}
		return http.ErrServerClosed
	}

	run([]string{"-puerto", "0", "-datos", path}, &out, &errOut, serve)

	if !strings.Contains(out.String(), "Los datos se guardan en "+path) {
		t.Errorf("se esperaba el aviso de donde se guardan los datos, salio:\n%s", out.String())
	}
	// Al terminar, la base quedo cerrada y tiene el proyecto: se puede volver a abrir y, en
	// Windows, borrar.
	db, err := store.OpenDatabase(path)
	if err != nil {
		t.Fatalf("no se pudo volver a abrir la base: %v", err)
	}
	registry, err := db.Load()
	db.Close()
	if err != nil {
		t.Fatalf("no se pudo cargar la base: %v", err)
	}
	if projects := registry.Projects(); len(projects) != 1 || projects[0].Name() != "Sistema de métricas" {
		t.Errorf("se esperaba en la base el proyecto creado desde la web, salio %v", projects)
	}
	if err := os.Remove(path); err != nil {
		t.Errorf("la base tenia que quedar cerrada al terminar, y no se pudo borrar: %v", err)
	}
}

// Escenarios: Exportar e importar da los mismos datos / Traer los datos de la versión
// anterior (con un archivo de prueba propio, no los datos de demo: Dependencias de la spec)
// Cubre: #177 RN-7, RN-8
func TestExportarEImportarDesdeElComando(t *testing.T) {
	dir := t.TempDir()
	base := filepath.Join(dir, "datos-ics.db")
	archivo := filepath.Join(dir, "exportado.json")
	registry := domain.NewProjectRegistry()
	if _, err := registry.Create(domain.ProjectData{Name: "Sistema de métricas", Start: time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC)}); err != nil {
		t.Fatalf("no se pudo preparar el proyecto: %v", err)
	}
	db, err := store.OpenDatabase(base)
	if err != nil {
		t.Fatalf("no se pudo preparar la base: %v", err)
	}
	if err := db.Save(registry); err != nil {
		t.Fatalf("no se pudo preparar la base: %v", err)
	}
	db.Close()

	t.Run("exportar escribe el archivo y termina sin abrir la web", func(t *testing.T) {
		var out, errOut bytes.Buffer

		code := run([]string{"-datos", base, "-exportar", archivo}, &out, &errOut, noServe(t))

		if code != 0 || !strings.Contains(out.String(), "Se exportaron los datos a "+archivo) {
			t.Errorf("se esperaba el codigo 0 y el aviso, salio %d:\n%s\n%s", code, out.String(), errOut.String())
		}
		leido, err := store.NewFile(archivo).Load()
		if err != nil || len(leido.Projects()) != 1 {
			t.Errorf("se esperaba el archivo con el proyecto, salio %v", err)
		}
	})

	t.Run("importar trae el archivo a una base nueva y abre con sus datos", func(t *testing.T) {
		var out, errOut bytes.Buffer
		var proyectos string
		serve := func(listener net.Listener, handler http.Handler) error {
			listener.Close()
			res := httptest.NewRecorder()
			handler.ServeHTTP(res, httptest.NewRequest(http.MethodGet, "/proyectos", nil))
			proyectos = res.Body.String()
			return http.ErrServerClosed
		}

		run([]string{"-puerto", "0", "-datos", filepath.Join(dir, "nueva.db"), "-importar", archivo}, &out, &errOut, serve)

		if !strings.Contains(out.String(), "Se importaron los datos de "+archivo) {
			t.Errorf("se esperaba el aviso de la importacion, salio:\n%s\n%s", out.String(), errOut.String())
		}
		if !strings.Contains(proyectos, "Sistema de métricas") {
			t.Errorf("se esperaba abrir con el proyecto importado, salio:\n%s", proyectos)
		}
	})
}

// Esquema del escenario: Importar algo que no se puede no cambia nada (desde el comando)
// Cubre: #177 RN-9, RN-10
func TestImportarSobreUnaBaseConDatosSeRechaza(t *testing.T) {
	dir := t.TempDir()
	base := filepath.Join(dir, "datos-ics.db")
	archivo := filepath.Join(dir, "otro.json")
	registry := domain.NewProjectRegistry()
	if _, err := registry.Create(domain.ProjectData{Name: "Sistema de métricas", Start: time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC)}); err != nil {
		t.Fatalf("no se pudo preparar el proyecto: %v", err)
	}
	db, err := store.OpenDatabase(base)
	if err != nil {
		t.Fatalf("no se pudo preparar la base: %v", err)
	}
	if err := db.Save(registry); err != nil {
		t.Fatalf("no se pudo preparar la base: %v", err)
	}
	db.Close()
	if err := store.NewFile(archivo).Save(domain.NewProjectRegistry()); err != nil {
		t.Fatalf("no se pudo preparar el archivo: %v", err)
	}
	var out, errOut bytes.Buffer

	code := run([]string{"-datos", base, "-importar", archivo}, &out, &errOut, noServe(t))

	if code != 1 || !strings.Contains(errOut.String(), store.ErrDatabaseNotEmpty.Error()) || !strings.Contains(errOut.String(), "La base "+base+" quedó como estaba.") {
		t.Errorf("se esperaba el codigo 1 y el aviso de que no importa sobre datos, salio %d:\n%s", code, errOut.String())
	}
}

// Escenario: Lo guardado con la versión anterior se importa solo
// Cubre: #177 RN-14, RN-15
func TestElArchivoDeLaVersionAnteriorSeImportaAlAbrir(t *testing.T) {
	dir := t.TempDir()
	viejo := filepath.Join(dir, "datos-ics.json")
	registry := domain.NewProjectRegistry()
	if _, err := registry.Create(domain.ProjectData{Name: "Sistema de metricas", Start: time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC)}); err != nil {
		t.Fatalf("no se pudo preparar el proyecto: %v", err)
	}
	if err := store.NewFile(viejo).Save(registry); err != nil {
		t.Fatalf("no se pudo preparar el archivo viejo: %v", err)
	}
	abrir := func() (string, string) {
		var out, errOut bytes.Buffer
		var proyectos string
		serve := func(listener net.Listener, handler http.Handler) error {
			listener.Close()
			res := httptest.NewRecorder()
			handler.ServeHTTP(res, httptest.NewRequest(http.MethodGet, "/proyectos", nil))
			proyectos = res.Body.String()
			return http.ErrServerClosed
		}
		run([]string{"-puerto", "0", "-datos", filepath.Join(dir, "datos-ics.db")}, &out, &errOut, serve)
		return out.String(), proyectos
	}

	out, proyectos := abrir()

	aviso := "Se importaron los datos de " + viejo + " y el archivo se renombró a " + viejo + ".importado"
	if !strings.Contains(out, aviso) || strings.Count(proyectos, "Sistema de metricas") != 1 {
		t.Errorf("se esperaba el aviso y el proyecto importado, salio:\n%s\n%s", out, proyectos)
	}
	// Al volver a abrir no se importa otra vez.
	out, proyectos = abrir()
	if strings.Contains(out, "Se importaron") || strings.Count(proyectos, "Sistema de metricas") != 1 {
		t.Errorf("la segunda vez no tenia que importar nada, salio:\n%s\n%s", out, proyectos)
	}
}

// Esquema del escenario: Lo que no se puede pedir al exportar o importar
// Cubre: #177 RN-16, RN-17, RN-18
func TestLoQueNoSePuedePedirAlExportarOImportar(t *testing.T) {
	dir := t.TempDir()
	base := filepath.Join(dir, "datos-ics.db")
	registry := domain.NewProjectRegistry()
	if _, err := registry.Create(domain.ProjectData{Name: "Sistema de métricas", Start: time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC)}); err != nil {
		t.Fatalf("no se pudo preparar el proyecto: %v", err)
	}
	db, err := store.OpenDatabase(base)
	if err != nil {
		t.Fatalf("no se pudo preparar la base: %v", err)
	}
	if err := db.Save(registry); err != nil {
		t.Fatalf("no se pudo preparar la base: %v", err)
	}
	db.Close()
	antes, _ := os.ReadFile(base)

	var out, errOut bytes.Buffer
	code := run([]string{"-datos", base, "-exportar", base}, &out, &errOut, noServe(t))

	if code != 1 || !strings.Contains(errOut.String(), store.ErrExportOverDatabase.Error()) {
		t.Errorf("exportar sobre la base: se esperaba el codigo 1 y el aviso, salio %d:\n%s", code, errOut.String())
	}
	if despues, _ := os.ReadFile(base); !bytes.Equal(despues, antes) {
		t.Error("exportar sobre la base: la base tenia que quedar como estaba")
	}
}

// Esquema del escenario: Lo que no se puede pedir al exportar o importar (las dos juntas, y
// exportar sin base)
// Cubre: #177 RN-17, RN-18
func TestImportarYExportarJuntosOSinBase(t *testing.T) {
	t.Run("importar y exportar a la vez se rechaza", func(t *testing.T) {
		dir := t.TempDir()
		args := []string{"-datos", filepath.Join(dir, "a.db"), "-importar", filepath.Join(dir, "x.json"), "-exportar", filepath.Join(dir, "y.json")}
		_, err := parseOptions(args)
		if err == nil || !strings.Contains(err.Error(), "-importar y -exportar") {
			t.Errorf("se esperaba el aviso de que no se pueden usar juntas, salio: %v", err)
		}
		var out, errOut bytes.Buffer
		code := run(args, &out, &errOut, noServe(t))
		if code != 2 {
			t.Errorf("se esperaba el codigo 2, salio %d:\n%s", code, errOut.String())
		}
		if entries, _ := os.ReadDir(dir); len(entries) != 0 {
			t.Errorf("no tenia que crear nada, y en la carpeta hay %d archivos", len(entries))
		}
	})

	t.Run("exportar sin base avisa y no crea nada", func(t *testing.T) {
		dir := t.TempDir()
		var out, errOut bytes.Buffer
		code := run([]string{"-datos", filepath.Join(dir, "no-existe.db"), "-exportar", filepath.Join(dir, "x.json")}, &out, &errOut, noServe(t))
		if code != 1 || !strings.Contains(errOut.String(), "no hay datos para exportar") {
			t.Errorf("se esperaba el codigo 1 y el aviso, salio %d:\n%s", code, errOut.String())
		}
		if entries, _ := os.ReadDir(dir); len(entries) != 0 {
			t.Errorf("no tenia que crear nada, y en la carpeta hay %d archivos", len(entries))
		}
	})
}
