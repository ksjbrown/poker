package hands

import (
	"fmt"
	"slices"

	"github.com/ksjbrown/poker/pkg/cards"
)

const (
	handMinSize = 2
	handMaxSize = 5
)

// Hand represents a combination of a minimum of 2 Cards and up to 5 Cards, which together determine a certain Score.
type Hand []cards.Card

// NewHand creates a new hand from the argument cards slice
func NewHand(c []cards.Card) (*Hand, error) {
	hand := Hand(c)
	handSize := len(hand)
	if handSize < handMinSize {
		return &hand, fmt.Errorf("hand must contain at least 2 cards")
	}
	if handSize > handMaxSize {
		return &hand, fmt.Errorf("hand must not contain more than 5 cards")
	}
	return &hand, nil
}

// calculateScore analyses the cards available to this hand and returns a HandScore indicating the strength of this Hand relative to other Hands.
func (h *Hand) calculateScore() Score {
	panic("not implemented")
}

// CompareTo is a shortcut for comparing a Hand's HandScore to another Hand's HandScore
func (h *Hand) Compare(other *Hand) int {
	return h.calculateScore().Compare(other.calculateScore())
}

func (h *Hand) Cards() []cards.Card {
	return []cards.Card(*h)
}

func (h *Hand) Copy() (*Hand, error) {
	copiedCards := make([]cards.Card, len(*h))
	copy(copiedCards, h.Cards())
	hand, err := NewHand(copiedCards)
	return hand, err
}

type HandSortAlgorithm func(left cards.Card, right cards.Card) int

func (h *Hand) Sort(algorithm HandSortAlgorithm) {
	slices.SortFunc(h.Cards(), algorithm)
}

func StandardSort(left cards.Card, right cards.Card) int {
	return int(left.Rank - right.Rank)
}

func AceHighSort(left cards.Card, right cards.Card) int {
	aceHighValue := func(card cards.Card) cards.Rank {
		if card.Rank == cards.ACE {
			return cards.KING + 1
		}
		return card.Rank
	}
	return int(aceHighValue(left) - aceHighValue(right))
}
