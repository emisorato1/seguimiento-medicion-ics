package domain

// VotingRound es una ronda de Planning Poker sobre una historia del proyecto (historia #13).
// Sus datos estan ocultos fuera del paquete: los votos solo se leen por la consulta que lo
// permite, y recien cuando la ronda esta revelada (RN-12).
type VotingRound struct {
	storyID      int
	number       int
	participants []string
	votes        []Vote
	revealed     bool
}

// Vote es un voto de una ronda revelada: quién lo dio y cuántos story points votó. Solo se
// entrega por Votes, y recién cuando la ronda está revelada (RN-12, RN-14).
type Vote struct {
	MemberName string
	Points     int
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

// Voted devuelve los nombres de los participantes que ya votaron, en el orden de los
// integrantes del proyecto y sin ningún valor (RN-13).
func (r *VotingRound) Voted() []string {
	return r.participantsWhere(true)
}

// Pending devuelve los nombres de los participantes que todavía no votaron, en el orden de
// los integrantes del proyecto (RN-13).
func (r *VotingRound) Pending() []string {
	return r.participantsWhere(false)
}

// participantsWhere devuelve, en el orden de los participantes, los que votaron o los que
// faltan. Es la única forma de saber algo de los votos antes de revelar: solo nombres.
func (r *VotingRound) participantsWhere(voted bool) []string {
	names := []string{}
	for _, name := range r.participants {
		if r.hasVoted(name) == voted {
			names = append(names, name)
		}
	}
	return names
}

// hasVoted indica si el participante con ese nombre ya votó en la ronda.
func (r *VotingRound) hasVoted(name string) bool {
	for _, vote := range r.votes {
		if vote.MemberName == name {
			return true
		}
	}
	return false
}

// Vote registra el voto de un participante y lo pasa de los que faltan a los que votaron
// (RN-8). La respuesta solo dice si se aceptó: nunca lleva el valor (RN-12). El voto tiene
// que ser un valor de la escala de la #12, con su mismo error (RN-9). Si se rechaza, la
// ronda queda como estaba (RN-18).
func (r *VotingRound) Vote(member string, points int) error {
	if !inScale(points) {
		return ErrStoryPointsInvalid
	}
	r.votes = append(r.votes, Vote{MemberName: member, Points: points})
	return nil
}

// Votes devuelve los votos de la ronda con el nombre de quien los dio. Es la única consulta
// que entrega valores, y mientras la ronda no esté revelada responde con un error, no con
// una lista vacía, que se confundiría con "nadie votó" (RN-12, Fase 2).
func (r *VotingRound) Votes() ([]Vote, error) {
	if !r.revealed {
		return nil, ErrVotesHidden
	}
	votes := make([]Vote, 0, len(r.votes))
	for _, name := range r.participants {
		for _, vote := range r.votes {
			if vote.MemberName == name {
				votes = append(votes, vote)
			}
		}
	}
	return votes, nil
}

// Reveal da vuelta todos los votos juntos: desde entonces Votes los entrega, cada uno con
// el nombre de quien lo dio y en el orden de los integrantes (RN-14). Si falta algún voto,
// se rechaza y la ronda sigue como estaba (RN-15, RN-18).
func (r *VotingRound) Reveal() error {
	if len(r.Pending()) > 0 {
		return ErrVotesMissing
	}
	r.revealed = true
	return nil
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
