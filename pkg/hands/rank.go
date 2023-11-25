package hands

import (
	"github.com/ksjbrown/poker/pkg/cards"
)

// HandRank defines the well-known hand types of poker.
// These are used as the major ranking score when comparing hands
// Take card not to confuse this struct with the cards.HandRank struct, related to the face value of a card.
type HandRank int

const (
	STRAIGHT_FLUSH  HandRank = 9
	FOUR_OF_A_KIND  HandRank = 8
	FULL_HOUSE      HandRank = 7
	FLUSH           HandRank = 6
	STRAIGHT        HandRank = 5
	THREE_OF_A_KIND HandRank = 4
	TWO_PAIR        HandRank = 3
	ONE_PAIR        HandRank = 2
	HIGH_CARD       HandRank = 1
)

var rankNames = map[HandRank]string{
	STRAIGHT_FLUSH:  "Straight Flush",
	FOUR_OF_A_KIND:  "Four of a Kind",
	FULL_HOUSE:      "Full House",
	FLUSH:           "Flush",
	STRAIGHT:        "Straight",
	THREE_OF_A_KIND: "Three of a Kind",
	TWO_PAIR:        "Two Pair",
	ONE_PAIR:        "One Pair",
	HIGH_CARD:       "High Card",
}

func (r HandRank) String() string {
	return rankNames[r]
}

func isStraightFlush(h Hand) bool {
	return isStraight(h) && isFlush(h)
}

func isFourOfAKind(h Hand) bool {
	if len(h) < 4 {
		return false
	}
	return expectRankGroupSizes(h, 4, 1)
}

func isFullHouse(h Hand) bool {
	if len(h) != 5 {
		return false
	}
	return expectRankGroupSizes(h, 3, 2)
}

func isFlush(h Hand) bool {
	if len(h) != 5 {
		return false
	}
	return expectSuitGroupSizes(h, 5)
}

func isStraight(h Hand) bool {
	if len(h) != 5 {
		return false
	}
	// solution must sort hand, make a copy so we don't affect the argument hand
	cs := cards.Cards(h)
	cs = cs.Copy()
	cs.Sort(cards.StandardSort)

	// check special ace high case
	if cs[0].Rank == cards.ACE &&
		cs[1].Rank == cards.TEN &&
		cs[2].Rank == cards.JACK &&
		cs[3].Rank == cards.QUEEN &&
		cs[4].Rank == cards.KING {
		return true
	}
	// check each elements Rank value is one greater than the previous element i.e. ranks are sequential
	for i := 1; i < 5; i++ {
		if (cs[i].Rank - cs[i-1].Rank) != 1 {
			return false
		}
	}
	return true
}

func isThreeOfAKind(h Hand) bool {
	if len(h) < 3 {
		return false
	}
	return expectRankGroupSizes(h, 3)
}

func isTwoPair(h Hand) bool {
	if len(h) < 4 {
		return false
	}
	return expectRankGroupSizes(h, 2, 2)
}

func isOnePair(h Hand) bool {
	return expectRankGroupSizes(h, 2)
}

func expectRankGroupSizes(h Hand, sizes ...int) bool {
	cs := cards.Cards(h)
	groups := cs.GroupByRank()
	return cards.ExpectGroupSizes(groups, sizes...)
}

func expectSuitGroupSizes(h Hand, sizes ...int) bool {
	cs := cards.Cards(h)
	groups := cs.GroupBySuit()
	return cards.ExpectGroupSizes(groups, sizes...)
}
