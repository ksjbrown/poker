package cards

import (
	"strconv"
)

// Rank defines the possible values of a Card
// internally represented by a uint8, so numeric comparison is possible.
type Rank int

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

const rankCount = 13

func (r Rank) String() string {
	return rankNames[r]
}

func (r *Rank) Char() string {
	if 2 <= *r && *r <= 9 {
		return strconv.Itoa(int(*r))
	}
	return rankNames[*r][:1]
}

func (r *Rank) isValid() bool {
	return 0 < *r && *r <= rankCount
}

// Score is a function used to compare the typical strength of a card.
// In particular, it assumes that Ace is a high card (one higher than King).
func (r *Rank) Score() int {
	if *r == ACE {
		return int(KING + 1)
	}
	return int(*r)
}
