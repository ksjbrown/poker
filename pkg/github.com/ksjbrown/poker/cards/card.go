package cards

import "fmt"

// Card represends a playing card in a standard Deck
// A Card has a Suit and Rank
type Card struct {
	Suit      Suit
	Rank      Rank
	CardColor CardColor
}

func NewCard(s Suit, r Rank) Card {
	return Card{s, r, s.CardColor()}
}

func (c Card) String() string {
	return fmt.Sprintf("%v of %v", c.Rank, c.Suit)
}
