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
	return fmt.Sprintf("%v of %v", c.Rank, c.Suit)
}

// IsValid returns true if the Suit and Rank for this card are valid values set to zero values.
//
// It
func (c *Card) isValid() bool {
	return (c.Rank > 0 && c.Rank <= rankCount) && (c.Suit > 0 && c.Suit <= suitCount)
}

func (c *Card) Less(other *Card) bool {
	return c.Rank < other.Rank
}
