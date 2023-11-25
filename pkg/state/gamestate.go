package state

import (
	"encoding/json"
	"io"
)

type GameStater interface {
	GetGameState() *GameState
	SetGameState(*GameState) error
}

// GameState represents the JSON information required to be stored on disk to recreate the game state at program startup
// GameState also is responsible for saving the other states.
type GameState struct {
	Id     int        `json:"gameId"`
	Status GameStatus `json:"gameStatus"`

	// Additional states
	PlayerStates   PlayerStates   `json:"players"`
	RoundStates    RoundStates    `json:"rounds"`
	SettingsStates SettingsStates `json:"settings"`
}

type GameStatus int

const (
	NOT_STARTED GameStatus = iota
	IN_PROGRESS
	COMPLETED
)

func LoadGameState(gsr GameStater, src *io.Reader) error {
	var gs GameState
	jsonDecoder := json.NewDecoder(*src)
	if err := jsonDecoder.Decode(&gs); err != nil {
		// TODO: log
		return err
	}
	if err := gsr.SetGameState(&gs); err != nil {
		// TODO: log
		return err
	}
	return nil
}

func SaveGameState(gsr GameStater, dest *io.Writer) error {
	gs := gsr.GetGameState()
	jsonEncoder := json.NewEncoder(*dest)
	if err := jsonEncoder.Encode(gs); err != nil {
		// TODO: log
		return err
	}

	return nil
}
