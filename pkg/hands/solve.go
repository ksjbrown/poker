package hands

import "github.com/ksjbrown/poker/pkg/cards"

// rankOccurences returns a map of Ranks and the amount of times they appear in the hand
// This is primarily used to check for four of a kind, full house, pairs, etc.
func rankOccurences(h *Hand) map[cards.Rank]int {
	occurrences := make(map[cards.Rank]int)
	for _, card := range *h {
		occurrences[card.Rank]++
	}
	return occurrences
}

func isStraight(h *Hand) bool {
	hand := *h
	for i := 1; i < len(hand); i++ {
		// check special case for ace high
		if i == len(hand)-1 && hand[i].Rank == cards.ACE && hand[i-1].Rank == cards.KING {
			continue
		}
		if (hand[i-1].Rank - hand[i].Rank) != 1 {
			return false
		}
	}
	return true
}

func (h *Hand) isStraightFlush() bool {
	if len(*h) < handMaxSize {
		return false
	}
	return h.isStraight() && h.isFlush()
}

func (h *Hand) isFourOfAKind() bool {
	for _, occurence := range rankOccurences(h) {
		if occurence >= 4 {
			return true
		}
	}
	return false
}

func (h *Hand) isStraight() bool {
	hand, err := h.Copy()
	if err != nil {
		// maybe some logging here or something
		return false
	}
	for _, algorithm := range []HandSortAlgorithm{StandardSort, AceHighSort} {
		hand.Sort(algorithm)
		if isStraight(hand) {
			return true
		}
	}
	return false
}

func (h *Hand) isFlush() bool {
	hand := *h
	if len(hand) != handMaxSize {
		return false
	}
	targetSuit := hand[0].Suit
	for _, card := range hand {
		if card.Suit != targetSuit {
			return false
		}
	}
	return true
}
