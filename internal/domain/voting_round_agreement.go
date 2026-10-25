package domain

// Lo que pasa después de revelar una ronda (historia #73): las diferencias entre los votos,
// la ronda nueva y la estimación acordada. La ronda es la de la #13, en voting_round.go.

// Differences es lo que se informa de los votos de una ronda revelada (RN-1): si hay
// diferencias, y el voto más bajo y el más alto con quiénes los dieron.
type Differences struct {
	HasDifferences bool
	Lowest         VoteExtreme
	Highest        VoteExtreme
}

// VoteExtreme es el voto más bajo o el más alto de una ronda: su valor y todos los que lo
// dieron, en el orden de los participantes (RN-2).
type VoteExtreme struct {
	Points  int
	Members []string
}

// Differences compara los votos de la ronda. Los lee con Votes, la consulta que protege la
// #13: con la ronda sin revelar ese error sube tal cual, y no se informa nada, ni siquiera
// si hay diferencias (RN-4). Hay diferencias cuando al menos dos votos son distintos; si
// todos coinciden, el más bajo y el más alto son ese mismo voto (RN-3).
func (r *VotingRound) Differences() (Differences, error) {
	votes, err := r.Votes()
	if err != nil {
		return Differences{}, err
	}
	lowest, highest := votes[0].Points, votes[0].Points
	for _, vote := range votes {
		lowest = min(lowest, vote.Points)
		highest = max(highest, vote.Points)
	}
	return Differences{
		HasDifferences: lowest != highest,
		Lowest:         extreme(votes, lowest),
		Highest:        extreme(votes, highest),
	}, nil
}

// extreme junta a todos los que votaron ese valor, en el orden en que llegan los votos, que
// es el de los participantes.
func extreme(votes []Vote, points int) VoteExtreme {
	members := []string{}
	for _, vote := range votes {
		if vote.Points == points {
			members = append(members, vote.MemberName)
		}
	}
	return VoteExtreme{Points: points, Members: members}
}

// Agreed devuelve la estimación acordada de la ronda, o 0 si todavía no se acordó, que no
// es un error (RN-15). Es un registro de lo que se acordó ese día: si después se reestima la
// historia con la #12, la ronda no cambia (punto 6 de la puerta).
func (r *VotingRound) Agreed() int {
	return r.agreed
}

// Agree registra la estimación acordada y deja la historia estimada con ese valor (RN-9).
// Estima con la operación de la #12, que controla la escala y que la historia no esté
// Terminada, con sus mismos errores (RN-10, RN-11), y cambia la historia solo si pasa. Recién
// entonces la ronda registra el valor: si algo se rechaza, la ronda y la historia quedan como
// estaban (RN-16). Antes de estimar van los controles de la ronda: tiene que estar revelada
// (RN-12) y sin acuerdo (RN-13).
func (r *VotingRound) Agree(points int) error {
	if !r.revealed {
		return ErrRoundNotRevealed
	}
	if r.agreed != 0 {
		return ErrRoundAlreadyAgreed
	}
	if err := r.project.EstimateStory(r.storyID, points); err != nil {
		return err
	}
	r.agreed = points
	return nil
}

// VotingRounds devuelve las rondas de la historia, en orden de número, con sus votos y su
// acuerdo si lo tienen (RN-7, RN-8). Una historia sin rondas devuelve una lista vacía, que no
// es un error; una que no existe informa que no se encontró.
func (p *Project) VotingRounds(storyID int) ([]*VotingRound, error) {
	if _, err := p.Story(storyID); err != nil {
		return nil, err
	}
	rounds := []*VotingRound{}
	for _, round := range p.votingRounds {
		if round.storyID == storyID {
			rounds = append(rounds, round)
		}
	}
	return rounds, nil
}

// lastVotingRound devuelve la última ronda que se abrió sobre la historia, o nil si no tiene
// ninguna. El proyecto guarda las rondas en orden de apertura, así que es la última de la
// lista con esa historia.
func (p *Project) lastVotingRound(storyID int) *VotingRound {
	for i := len(p.votingRounds) - 1; i >= 0; i-- {
		if p.votingRounds[i].storyID == storyID {
			return p.votingRounds[i]
		}
	}
	return nil
}
