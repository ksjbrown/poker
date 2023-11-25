package hands

import (
	"fmt"

	"github.com/ksjbrown/poker/pkg/cards"
)

// Hand represents a combination of a minimum of 2 Cards and up to 5 Cards, which together determine a certain Score.
type Hand cards.Cards

// NewHand creates a new hand from the argument cards slice, ensuring that number of cards is between 2 and 5
func NewHand(cs ...cards.Card) (*Hand, error) {
	handSize := len(cs)
	if handSize < 2 {
		return nil, fmt.Errorf("hand must contain at least 2 cards")
	}
	if handSize > 5 {
		return nil, fmt.Errorf("hand must not contain more than 5 cards")
	}
	hand := Hand(cs)
	copied := hand.Copy()
	return &copied, nil
}

// Copy creates a copy of the underlying array of cards this Hand is based on, and returns a new Hand from these copied cards.
func (h *Hand) Copy() Hand {
	cards := cards.Cards(*h)
	copied := cards.Copy()
	return Hand(copied)
}

func (h *Hand) String() string {
	return rankNames[h.Rank()]
}

func (h *Hand) Rank() HandRank {
	if isStraightFlush(*h) {
		return STRAIGHT_FLUSH
	}
	if isFourOfAKind(*h) {
		return FOUR_OF_A_KIND
	}
	if isFullHouse(*h) {
		return FULL_HOUSE
	}
	if isFlush(*h) {
		return FLUSH
	}
	if isStraight(*h) {
		return STRAIGHT
	}
	if isThreeOfAKind(*h) {
		return THREE_OF_A_KIND
	}
	if isTwoPair(*h) {
		return TWO_PAIR
	}
	if isOnePair(*h) {
		return ONE_PAIR
	}
	return HIGH_CARD
}
