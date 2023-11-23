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

func (s *Suit) isValid() bool {
	return 0 < *s && *s <= suitCount
}

func (s Suit) String() string {
	return suitNames[s]
}

func (s *Suit) Char() string {
	return s.String()[:1]
}
