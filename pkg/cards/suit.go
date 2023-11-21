package cards

// Suit defines an enum of all possible card suits.
//
// A card can have a suit of Spades, Clubs, Hearts or Diamonds.
// Additionally, a CardColor (one of Red or Black) can be determined via the CardColor() method.
type Suit int

const (
	CLUBS Suit = iota + 1
	DIAMONDS
	HEARTS
	SPADES
)

var suitNames = map[Suit]string{
	CLUBS:    "Clubs",
	DIAMONDS: "Diamonds",
	HEARTS:   "Hearts",
	SPADES:   "Spades",
}

const suitCount = 4

func (s Suit) String() string {
	return suitNames[s]
}

// AllSuits() returns all possible Suit values, which can be used for iteration, checking if a string is a valid Suit value, etc.
func AllSuits() [4]Suit {
	return [4]Suit{CLUBS, DIAMONDS, HEARTS, SPADES}
}
