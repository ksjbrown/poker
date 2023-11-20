package cards

// Rank defines the possible values of a Card
// internally represented by a uint8, so numeric comparison is possible.
type Rank uint8

const (
	ACE Rank = iota + 1
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)

var rankNames = map[Rank]string{
	ACE:   "Ace",
	TWO:   "Two",
	THREE: "Three",
	FOUR:  "Four",
	FIVE:  "Five",
	SIX:   "Six",
	SEVEN: "Seven",
	EIGHT: "Eight",
	NINE:  "Nine",
	TEN:   "Ten",
	JACK:  "Jack",
	QUEEN: "Queen",
	KING:  "King",
}

func (r Rank) String() string {
	return rankNames[r]
}

func AllRanks() [13]Rank {
	return [13]Rank{
		ACE,
		TWO,
		THREE,
		FOUR,
		FIVE,
		SIX,
		SEVEN,
		EIGHT,
		NINE,
		TEN,
		JACK,
		QUEEN,
		KING,
	}
}
