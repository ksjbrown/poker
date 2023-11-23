package game

import (
	"time"

	"github.com/ksjbrown/poker/pkg/cards"
)


// GameSettings defines all the configurable options about a game
type GameSettings struct {
	// TODO
}

// GameState contains all of the relevant information to the game
//
// A Poker game consists of any number of betting rounds (RoundState)
// Each Round consists of different betting stages ()
type GameState struct {
	Settings       GameSettings
	StartTime      time.Time
	Rounds []RoundState
	
	CommunityCards cards.Cards
}

type RoundState struct {

	// The state of the player in this round
	Players []PlayerState

	// calculate BigBlind and SmallBlind from Dealer
	Dealer Player

	// The time the round was started
	StartTime time.Time

	// The different phases that make up the round
	Phases []RoundPhase

}

// RoundPhaseState records ll of the relevant information regarding the 
type RoundPhaseState struct {

}

// RoundPhase is an enum indicating what phase of the round we are in
type RoundPhase int

const (
	NOT_STARTED RoundPhase = iota + 1
	PREFLOP
	FLOP
	TURN
	RIVER
)
