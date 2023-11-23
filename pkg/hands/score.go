package hands

import (
	"github.com/ksjbrown/poker/pkg/cards"
)

// Score is a comparison metric for Hand structs.
//
// It is stored as a hexadecimal integer containing the values:
//   - 0xABCCC
//
// Where:
//   - A: the Major rank. 	This is the type of hand, e.g. high card, one pair, straight flush
//   - B: the Minor rank. 	Differentiates between types of hands at different strengths.
//   - C: the Micro rank. 	Differentiates between equal hands via the remaining "kicker".
//     There can be up to 4 kicker cards (for high card hand), so we need 4 hex digits here.
type Score int

const (
	digitWidth  int = 4
	microDigits int = 4
	minorDigits int = 1
	majorDigits int = 1
)

func NewScore(major int, minor int, micro ...int) *Score {
	var score Score
	// major
	score.set(major, majorDigits, minorDigits+microDigits)
	// minor
	score.set(minor, minorDigits, microDigits)
	// micro
	score.set(0, microDigits, 0)
	microIndexMax := min(len(micro), microDigits) - 1
	for i := 0; i <= microIndexMax; i++ {
		shift := microDigits - 1 - i
		score.set(micro[i], 1, shift)
	}
	return &score
}

func (s *Score) Major() int {
	return s.get(majorDigits, minorDigits+microDigits)
}

func (s *Score) Minor() int {
	return s.get(minorDigits, microDigits)
}

func (s *Score) Micro() int {
	return s.get(microDigits, 0)
}

func (s *Score) get(width int, shift int) int {
	width *= digitWidth
	shift *= digitWidth
	mask := ((1 << width) - 1) << shift
	return (int(*s) & mask) >> shift
}

// set unsets the bits where the score should be written (defined by the bit width and shift), and then writes the score to that bit position
func (s *Score) set(score int, width int, shift int) {
	width *= digitWidth
	shift *= digitWidth
	mask := ((1 << width) - 1) << shift
	oldValue := int(*s) & mask
	newValue := score << shift
	*s = *s + Score(newValue-oldValue)
}

func (s *Score) Compare(other Score) int {
	return int(*s - other)
}

func (h *Hand) Score() int {
	rank := h.Rank()
	score := NewScore(
		int(rank),
		calculateMinor(*h, rank),
		calculateMicro(*h, rank)...,
	)
	return int(*score)
}

// calculateMinor calculates the correct minor score based on the rank of the hand
func calculateMinor(h Hand, r Rank) int {
	// sort cards, so we can predict where e.g. highest card is located
	cs := cards.Cards(h)
	cs = cs.Copy()
	cs.Sort(cards.AceHighSort)

	switch r {

	case STRAIGHT_FLUSH, FLUSH, STRAIGHT, HIGH_CARD:
		// highest card
		return cs[len(cs)-1].Score()

	case FOUR_OF_A_KIND, FULL_HOUSE, THREE_OF_A_KIND, ONE_PAIR:
		// rank of the four of a kind
		// group by rank, sort by largest group, and take the rank of any card in the first group
		grouped := cs.GroupByRank()
		groups := cards.OrderGroupsByLength(grouped)
		return groups[0][0].Score()

	case TWO_PAIR:
		// get max score of any card from the two pairs (the first two groups)
		grouped := cs.GroupByRank()
		groups := cards.OrderGroupsByLength(grouped)
		return max(groups[0][0].Score(), groups[1][0].Score())

	default:
		return 0
	}
}

// calculateMicro calculates the correct micro score based on the rank of the hand
func calculateMicro(h Hand, r Rank) []int {
	// sort cards, so we can predict where e.g. highest card is located
	cs := cards.Cards(h)
	cs = cs.Copy()
	cs.Sort(cards.AceHighSort)

	switch r {
	case STRAIGHT_FLUSH, STRAIGHT:
		// no kicker cards
		return []int{0}

	case FOUR_OF_A_KIND, FULL_HOUSE:
		// rank of any card in the second largest group
		grouped := cs.GroupByRank()
		groups := cards.OrderGroupsByLength(grouped)
		return []int{groups[1][0].Score()}

	case THREE_OF_A_KIND, ONE_PAIR:
		// kickers are any cards not in the three of a kind
		grouped := cs.GroupByRank()
		groups := cards.OrderGroupsByLength(grouped)
		kickers := cs.GroupBySelection(groups[0])[1]
		kickerScores := make([]int, 0, len(kickers))
		for i := len(kickers) - 1; i >= 0; i-- {
			kickerScores = append(kickerScores, kickers[i].Score())
		}
		// kickers are still in face value order, since we sorted at start of the method
		// we can return them in reverse order to get most significant score at index 0
		return kickerScores

	case TWO_PAIR:
		grouped := cs.GroupByRank()
		groups := cards.OrderGroupsByLength(grouped)
		// first kicker is the pair with the lowest face value, second is the final group
		return []int{
			min(groups[0][0].Score(), groups[1][0].Score()),
			groups[2][0].Score(),
		}

	case HIGH_CARD, FLUSH:
		// return all cards except the highest card, in reverse order
		maxIndex := len(cs) - 1
		return []int{
			cs[maxIndex-1].Score(),
			cs[maxIndex-2].Score(),
			cs[maxIndex-3].Score(),
			cs[maxIndex-4].Score(),
		}

	default:
		return []int{0}
	}
}
