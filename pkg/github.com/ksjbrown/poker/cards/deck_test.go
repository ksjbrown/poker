package cards

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	want := 52

	if len(deck.cards) != want {
		t.Errorf("expected %v cards got %v", want, len(deck.cards))
	}
}
