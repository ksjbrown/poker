package cards

// CardSuit defines an enum of all possible card suits.
//
// A card can have a suit of Spades, Clubs, Hearts or Diamonds.
// Additionally, a CardColor (one of Red or Black) can be determined via the CardColor() method.
type CardSuit int

const (
	CLUBS CardSuit = iota + 1
	DIAMONDS
	HEARTS
	SPADES
)

var suitNames = map[CardSuit]string{
	CLUBS:    "Clubs",
	DIAMONDS: "Diamonds",
	HEARTS:   "Hearts",
	SPADES:   "Spades",
}

const suitCount = 4

func (s *CardSuit) isValid() bool {
	return 0 < *s && *s <= suitCount
}

func (s CardSuit) String() string {
	return suitNames[s]
}

func (s *CardSuit) Char() string {
	return s.String()[:1]
}
