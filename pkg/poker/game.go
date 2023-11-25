package poker

import (
	"fmt"

	"github.com/ksjbrown/poker/pkg/state"
)

const (
	MIN_PLAYER_COUNT = 2
)

// The main Game game instance
//
// Upon creation, an ID for the game is generated.
// Players can Join and Leave games
//
// Unique IDs will be generated for any new player, who also can provide a Display Name
type Game struct {
	state.GameState
}

func NewGame(id int) *Game {
	g := Game{
		GameState: state.GameState{Id: id},
	}
	return &g
}

func (g *Game) GetGameState() *state.GameState {
	return &g.GameState
}

func (g *Game) SetGameState(gs *state.GameState) error {
	if gs == nil {
		return fmt.Errorf("no game state provided")
	}
	g.GameState = *gs
	return nil
}

func (g *Game) StartGame() error {
	// player count must be above minimum
	if !checkOkPlayerCount(g) {
		return fmt.Errorf("expected min %v players, only %v in game %v", MIN_PLAYER_COUNT, len(g.PlayerStates), g.Id)
	}
	g.Status = state.IN_PROGRESS
	return nil
}

func checkOkPlayerCount(g *Game) bool {
	return len(g.PlayerStates) >= MIN_PLAYER_COUNT
}
