package game

// The main Poker game instance
//
// Upon creation, an ID for the game is generated.
// Players can Join a game.
// If a Player leaves, they can Join a game again.
//
// The Player who 

// Unique IDs will be generated for any new player, who also can provide a Display Name
type Poker struct {

	Players map[int]PlayerState

}
