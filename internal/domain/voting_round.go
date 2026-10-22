package domain

// VotingRound es una ronda de Planning Poker sobre una historia del proyecto (historia #13).
// Sus datos estan ocultos fuera del paquete: los votos solo se leen por la consulta que lo
// permite, y recien cuando la ronda esta revelada (RN-12).
type VotingRound struct {
	storyID      int
	number       int
	participants []string
	revealed     bool
}

// StoryID devuelve el identificador de la historia que se vota en la ronda.
func (r *VotingRound) StoryID() int {
	return r.storyID
}

// Number devuelve el numero de la ronda sobre su historia, que empieza en 1 (RN-1).
func (r *VotingRound) Number() int {
	return r.number
}

// Revealed indica si la ronda ya fue revelada.
func (r *VotingRound) Revealed() bool {
	return r.revealed
}

// Voted devuelve los nombres de los participantes que ya votaron, sin ningun valor (RN-13).
func (r *VotingRound) Voted() []string {
	return []string{}
}

// Pending devuelve los nombres de los participantes que todavia no votaron, en el orden de
// los integrantes del proyecto (RN-13).
func (r *VotingRound) Pending() []string {
	pending := make([]string, len(r.participants))
	copy(pending, r.participants)
	return pending
}

// OpenVotingRound abre una ronda de votacion sobre la historia con ese identificador. Los
// participantes son los integrantes que el proyecto tiene en este momento, y se fijan ahi
// (RN-2). Abrir la ronda no cambia la historia (RN-3). Primero busca la historia (RN-5) y
// después mira su estado: una Terminada no se vota, porque su estimación ya no se modifica
// (RN-4, con el error de la #12).
func (p *Project) OpenVotingRound(storyID int) (*VotingRound, error) {
	story, err := p.Story(storyID)
	if err != nil {
		return nil, err
	}
	if story.state == StoryDone {
		return nil, ErrStoryEstimateDone
	}
	if len(p.members) == 0 {
		return nil, ErrVotingRoundNoMembers
	}
	// Hasta que esté la #73, cada historia tiene una sola ronda: se rechaza toda segunda
	// ronda, esté revelada o no la primera (punto 6 de la puerta). La #73 lo afloja para
	// las reveladas.
	for _, round := range p.votingRounds {
		if round.storyID == storyID {
			return nil, ErrVotingRoundInProgress
		}
	}
	participants := make([]string, 0, len(p.members))
	for _, member := range p.members {
		participants = append(participants, member.name)
	}
	round := &VotingRound{storyID: storyID, number: 1, participants: participants}
	p.votingRounds = append(p.votingRounds, round)
	return round, nil
}
