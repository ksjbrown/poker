package hands

import (
	"sort"

	"github.com/ksjbrown/poker/pkg/cards"
)

// getRankOccurences returns a map of Ranks and the amount of times they appear in the hand
// This is primarily used to check for four of a kind, full house, pairs, etc.
func getRankOccurences(h *Hand) map[cards.Rank]int {
	occurrences := make(map[cards.Rank]int)
	for _, card := range *h {
		occurrences[card.Rank]++
	}
	return occurrences
}

func filterRankByOccurence(h *Hand, occurrences ...int) []cards.Rank {
	expectedOccurrences := []int(occurrences)
	sort.Ints(expectedOccurrences)

	rankOccurrences := getRankOccurences(h)
	ranks := make([]cards.Rank, 0, len(expectedOccurrences))

	for i := len(expectedOccurrences) - 1; i >= 0; i-- {
		expectedOccurrence := expectedOccurrences[i]
		for rank, occurence := range rankOccurrences {
			if occurence >= expectedOccurrence {
				ranks = append(ranks, rank)
				delete(rankOccurrences, rank)
				break
			}
		}
	}
	return ranks
}

// expectOccurrences will check that a hand has cards with a number of ranks that occur as given by the argument occurrences.
func expectOccurrences(h *Hand, occurrences ...int) bool {
	ranks := filterRankByOccurence(h, occurrences...)
	return len(ranks) == len(occurrences)
}

func isStraight(h *Hand) bool {
	hand := *h
	if len(hand) != 5 {
		return false
	}
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
	if len(*h) != 5 {
		return false
	}
	return h.isStraight() && h.isFlush()
}

func (h *Hand) isFourOfAKind() bool {
	if len(*h) < 4 {
		return false
	}
	return expectOccurrences(h, 4)

}

func (h *Hand) isFullHouse() bool {
	if len(*h) != 5 {
		return false
	}
	return expectOccurrences(h, 3, 2)
}

func (h *Hand) isFlush() bool {
	hand := *h
	if len(hand) != 5 {
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

func (h *Hand) isStraight() bool {
	// solution must sort hand, make a copy so we don't affect the argument hand
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

func (h *Hand) isThreeOfAKind() bool {
	if len(*h) < 3 {
		return false
	}
	return expectOccurrences(h, 3)
}

func (h *Hand) isTwoPair() bool {
	if len(*h) < 4 {
		return false
	}
	return expectOccurrences(h, 2, 2)
}

func (h *Hand) isOnePair() bool {
	return expectOccurrences(h, 2)
}
