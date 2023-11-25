package cards

import (
	"fmt"
	"slices"
)

// Card represends a playing card in a standard Deck
// A Card has a Suit and Rank.
type Card struct {
	Rank CardRank
	Suit CardSuit
}

// NewCard will return a new *Card with the argument Suit and Rank.
func NewCard(r CardRank, s CardSuit) (*Card, error) {
	card := Card{r, s}
	if !card.isValid() {
		return &card, fmt.Errorf("invalid Card args: Rank=%d, Suit=%d", r, s)
	}
	return &card, nil
}

func (c *Card) String() string {
	return c.Rank.Char() + c.Suit.Char()
}

// IsValid returns true if the Suit and Rank for this card are valid values set to zero values.
//
// It
func (c *Card) isValid() bool {
	return (c.Rank.isValid() && c.Suit.isValid())
}

func (c *Card) Less(other Card) bool {
	return c.Rank < other.Rank
}

// Score is a function used to compare the typical strength of a card.
// In particular, it assumes that Ace is a high card (one higher than King).
func (c *Card) Score() int {
	return c.Rank.Score()
}

type Cards []Card

func (cs *Cards) Copy() Cards {
	return slices.Clone(*cs)
}

// SortAlgorithm represents a method of sorting cards
// Implementations give the difference between left and right based on some sorting metric
type SortAlgorithm func(left Card, right Card) int

// StandardSort will sort the cards in this hand according to their face rank, i.e. Ace first, Two second, and so on, up til King
func StandardSort(left Card, right Card) int {
	return int(left.Rank - right.Rank)
}

// AceHighSort will sort cards, with Ace being considered the highest card, i.e. Two first, Three second, ... Queen, King, then Ace
func AceHighSort(left Card, right Card) int {
	return left.Score() - right.Score()
}

// Sort will sort the cards in this hand based on the argument HandSortAlgorithm
func (cs *Cards) Sort(algorithm SortAlgorithm) {
	slices.SortFunc(*cs, algorithm)
}

// GroupByRank takes a slice of cards, and returns a map of rank to cards, where each card in a slice has the same rank
func (cs *Cards) GroupByRank() map[CardRank]Cards {
	grouped := make(map[CardRank]Cards)
	for _, card := range *cs {
		grouped[card.Rank] = append(grouped[card.Rank], card)
	}
	return grouped
}

// GroupBySuit takes a slice of cards, and returns a map of suit to cards, where each card in a slice has the same suit
func (cs *Cards) GroupBySuit() map[CardSuit]Cards {
	grouped := make(map[CardSuit]Cards)
	for _, card := range *cs {
		grouped[card.Suit] = append(grouped[card.Suit], card)
	}
	return grouped
}

// GroupBySelection returns a slice of Cards, where the first element is each Card in this Cards which is also in the selection Cards.
// The second element contains the remaining cards
func (cs *Cards) GroupBySelection(selection Cards) []Cards {
	grouped := make([]Cards, 2)
	const found int = 0
	const notFound int = 1

	for _, card := range *cs {
		group := notFound
		for _, selected := range selection {
			if card == selected {
				group = found
				break
			}
		}
		grouped[group] = append(grouped[group], card)
	}
	return grouped
}

// OrderGroupsByLength returns a slice Cards, the Cards at index 0 is the largest, at index 1 the next largest, and so on.
func OrderGroupsByLength[T comparable](mcs map[T]Cards) []Cards {
	// unpack values from map into slice
	groups := make([]Cards, 0, len(mcs))
	for _, group := range mcs {
		groups = append(groups, group)
	}
	// sort by reverse order of length of each group
	slices.SortFunc(groups, func(cs1 Cards, cs2 Cards) int {
		return -1 * (len(cs1) - len(cs2))
	})
	return groups
}

func ExpectGroupSizes[T comparable](mcs map[T]Cards, sizes ...int) bool {
	// sort groups and sizes largest first
	groups := OrderGroupsByLength(mcs)
	slices.SortFunc(sizes, func(i1 int, i2 int) int {
		return -1 * (i1 - i2)
	})
	// we can check fewer sizes than groups, but not more
	if len(sizes) > len(groups) {
		return false
	}
	// check group sizes match expected sizes
	for i := 0; i < len(sizes); i++ {
		if len(groups[i]) != sizes[i] {
			return false
		}
	}
	// all groups match
	return true
}
