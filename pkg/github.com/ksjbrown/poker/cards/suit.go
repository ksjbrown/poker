package cards

// Suit defines an enum of all possible card suits
// A card can have a suit of Spades, Clubs, Hearts or Diamonds
type Suit string

const (
	CLUBS    Suit = "Clubs"
	DIAMONDS Suit = "Diamonds"
	HEARTS   Suit = "Hearts"
	SPADES   Suit = "Spades"
)

type CardColor string

const (
	RED   CardColor = "Red"
	BLACK CardColor = "Black"
)

var suitColors = map[Suit]CardColor{
	CLUBS:    BLACK,
	DIAMONDS: RED,
	HEARTS:   RED,
	SPADES:   BLACK,
}

func (s Suit) CardColor() CardColor {
	return suitColors[s]
}

func AllSuits() [4]Suit {
	return [4]Suit{CLUBS, DIAMONDS, HEARTS, SPADES}
}
