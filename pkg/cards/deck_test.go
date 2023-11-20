package cards

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	var deck = NewDeck()
	var want = 52

	// fmt.Print(deck)

	if len(deck.Cards) != want {
		t.Errorf("expected %v cards got %v", want, len(deck.Cards))
	}
}
