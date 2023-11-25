package state

type PlayerStates []PlayerState

type PlayerState struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	IsAdmin bool   `json:"isAdmin"`
}
