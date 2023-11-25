package poker

import "time"

type Phase struct {

	// enum for the PhaseType
	PhaseType PhaseType

	// Timestamp when the phase began
	Timestamp time.Time

	// slice of the player actions which were performed in this phase
	PlayerActions []PlayerAction
}

type PhaseType int

const (
	PREFLOP PhaseType = iota + 1
	FLOP
	TURN
	RIVER
)
