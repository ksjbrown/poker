package game

import "github.com/ksjbrown/poker/pkg/cards"

// PlayerState represents the state of the player
// This is separated from Player, so that it can be serialized, send to other PlayerIO instances
type PlayerState struct {

	// The unique ID Of the Player
	Id int

	// The name of the player
	Name string

	// The available balance of the player
	Credits int

	// The current cards a player has
	PocketCards cards.Cards
}

// Player represents a Poker player playing in this game.
type Player struct {

	// provides methods for reporting decisions to the game
	// and receiving updates from the game about other players decisions
	playerIO PlayerIO

	// The current state of the player
	PlayerState PlayerState
}
