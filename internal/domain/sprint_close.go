package domain

import "time"

// CloseSprint cierra el sprint abierto del proyecto con la fecha de cierre indicada (historia
// #10, RN-1). El sprint que se cierra es siempre el abierto: por eso se le pide al proyecto, y
// cerrado el sprint, el proyecto queda sin sprint abierto y puede crear el siguiente (RN-14).
func (p *Project) CloseSprint(completed []int, closedOn time.Time) error {
	sprint, _ := p.OpenSprint()
	sprint.state = SprintClosed
	sprint.closedOn = toDay(closedOn)
	return nil
}

// ClosedOn devuelve el día en que el sprint se cerró de verdad, o la fecha cero si sigue
// abierto.
func (s *Sprint) ClosedOn() time.Time {
	return s.closedOn
}
