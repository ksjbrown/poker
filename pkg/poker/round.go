package poker

import "github.com/ksjbrown/poker/pkg/cards"

type Round struct {

	// Id is the round number. The first round has an id of 0, second of 1, and so on
	Id RoundID

	// Players are the players who participate in this round.
	// The Dealer is always the first player in the slice.
	Players Players

	// CommunityCards are the cards available to all players for formation of hands
	CommunityCards cards.Cards

	// Phases represent the different stages in a round
	Phases []Phase

	// The winner of the round
	Winner PlayerID
}

type RoundID int

type Rounds []Round
