package state

import (
	"encoding/json"
	"io"
)

// State defines the interface for state storing instances
type State interface {
}

func SaveState(dest io.Writer, state State) error {
	jsonEncoder := json.NewEncoder(dest)
	return jsonEncoder.Encode(state)
}

func LoadState(source io.Reader, state State) error {
	jsonDecoder := json.NewDecoder(source)
	return jsonDecoder.Decode(state)
}
