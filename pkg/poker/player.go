package poker

import (
	"github.com/ksjbrown/poker/pkg/state"
)

// Player represents a Poker player playing in this game.
type Player struct {
	// promote PlayerState
	state.PlayerState

	
}

type PlayerID int

type Players struct {
	Players map[PlayerID]Player
	Admins  map[PlayerID]bool
}

type PlayerAction struct {
	Type   PlayerActionType
	Amount int
}

type PlayerActionType int

const (
	FOLD PlayerActionType = iota + 1
	CHECK
	BLIND
	CALL
	RAISE
)
