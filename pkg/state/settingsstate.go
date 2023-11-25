package state

// SettingsStates contains a slice of SettingsState updates.
// The current settings are the result of all SettingsState updates applied one after the other.
// The resultant GameSettings can be calcualted via the SettingsStates.Current() method
type SettingsStates []SettingsState

type SettingsState struct {
	Round  int                                `json:"round"`
	Values map[GameSettingsOption]interface{} `json:"values"`
}

type GameSettingsOption string

const (
	CREDITS    GameSettingsOption = "credits"
	BIGBLIND   GameSettingsOption = "bigBlind"
	SMALLBLIND GameSettingsOption = "smallBlind"
)
