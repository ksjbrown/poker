package hands

import "github.com/ksjbrown/poker/pkg/cards"

type Solver struct {
	cards []*cards.Card
}

func NewSolver(cards []*cards.Card) Solver {
	solver := Solver{cards}
	return solver
}

func (s *Solver) Solve() Score {
	panic("not implemented")
}

// Returns true if a straight flush can be created from the cards assigned to the solver.
func (s *Solver) isStraightFlush() bool {
	panic("not implemented")
}
