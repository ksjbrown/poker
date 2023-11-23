package game

import "time"

// PlayerIO defines the methods expected by the server logic,
// so that it can query info from the player, and report if info is invalid or acceptable, updates of state, etc.
type PlayerIO interface {
	onActionRequested(timeout time.Duration) PlayerTurnAction
	onActionAccepted(pta PlayerTurnAction)
	onActionRejected(timeout time.Duration, pta PlayerTurnAction) PlayerTurnAction
	onActionTimedOut()
	onGameStateUpdated(gs *GameState)
}

type PlayerTurnAction struct {
	Type   PlayerTurnActionType
	Amount int // only set for some actions, blind, raise, call
}

type PlayerTurnActionType int

const (
	FOLD PlayerTurnActionType = iota + 1
	CHECK
	BLIND
	CALL
	RAISE
)
