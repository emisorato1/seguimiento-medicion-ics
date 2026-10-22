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
// (RN-2). Abrir la ronda no cambia la historia (RN-3).
func (p *Project) OpenVotingRound(storyID int) (*VotingRound, error) {
	if _, err := p.Story(storyID); err != nil {
		return nil, err
	}
	participants := make([]string, 0, len(p.members))
	for _, member := range p.members {
		participants = append(participants, member.name)
	}
	return &VotingRound{storyID: storyID, number: 1, participants: participants}, nil
}
