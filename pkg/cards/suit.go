package cards

// Suit defines an enum of all possible card suits. 
// 
// A card can have a suit of Spades, Clubs, Hearts or Diamonds.
// Additionally, a CardColor (one of Red or Black) can be determined via the CardColor() method.
type Suit string

const (
	CLUBS    Suit = "Clubs"
	DIAMONDS Suit = "Diamonds"
	HEARTS   Suit = "Hearts"
	SPADES   Suit = "Spades"
)

// CardColor defines the color of a card, based on it's Suit value.
// It can be either Red or Black, and can be accessed from a Suit via the CardColor method().
type CardColor string

const (
	RED   CardColor = "Red"
	BLACK CardColor = "Black"
)

// suitColors is an internal mapping of a Suit to it's appropriate CardColor
var suitColors = map[Suit]CardColor{
	CLUBS:    BLACK,
	DIAMONDS: RED,
	HEARTS:   RED,
	SPADES:   BLACK,
}

// CardColor() returns the CardColor associated with the attached Suit.
func (s *Suit) CardColor() CardColor {
	return suitColors[*s]
}

// AllSuits() returns all possible Suit values, which can be used for iteration, checking if a string is a valid Suit value, etc.
func AllSuits() [4]Suit {
	return [4]Suit{CLUBS, DIAMONDS, HEARTS, SPADES}
}
