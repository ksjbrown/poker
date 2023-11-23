package cards

import (
	"fmt"
)

// Card represends a playing card in a standard Deck
// A Card has a Suit and Rank.
type Card struct {
	Rank Rank
	Suit Suit
}

// NewCard will return a new *Card with the argument Suit and Rank.
func NewCard(r Rank, s Suit) (*Card, error) {
	card := Card{r, s}
	if !card.isValid() {
		return &card, fmt.Errorf("invalid Card args: Rank=%d, Suit=%d", r, s)
	}
	return &card, nil
}

func (c *Card) String() string {
	return c.Rank.Char() + c.Suit.Char()
}

// IsValid returns true if the Suit and Rank for this card are valid values set to zero values.
//
// It
func (c *Card) isValid() bool {
	return (c.Rank.isValid() && c.Suit.isValid())
}

func (c *Card) Less(other Card) bool {
	return c.Rank < other.Rank
}

// Score is a function used to compare the typical strength of a card.
// In particular, it assumes that Ace is a high card (one higher than King).
func (c *Card) Score() int {
	return c.Rank.Score()
}
