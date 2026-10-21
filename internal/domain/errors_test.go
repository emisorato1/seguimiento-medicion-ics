package domain

import "testing"

// Sin escenario: defecto #60
// Cubre: los mensajes de la RN-4 y la RN-6 de la #3, y de la RN-5 y la RN-11 de la #4
//
// Defecto #60: los mensajes de error los lee el usuario tal cual (la consola los muestra
// sin cambios), asi que van en español bien escrito, con sus tildes. El resto de los tests
// compara los errores con errors.Is; este es el unico que mira el texto, y solo de los
// mensajes que tenian el defecto.
func TestLosMensajesDeErrorLlevanTildes(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want string
	}{
		{name: "fechas incoherentes", err: ErrProjectDatesInconsistent, want: "las fechas son incoherentes: la fecha de finalización es anterior a la de inicio"},
		{name: "proyecto no encontrado", err: ErrProjectNotFound, want: "no se encontró el proyecto"},
		{name: "integrante duplicado", err: ErrMemberDuplicate, want: "el integrante ya está cargado en el proyecto"},
		{name: "integrante no encontrado", err: ErrMemberNotFound, want: "no se encontró el integrante en el proyecto"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.err.Error() != c.want {
				t.Errorf("se esperaba el mensaje %q, salio: %q", c.want, c.err.Error())
			}
		})
	}
}
