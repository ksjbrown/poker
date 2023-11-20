package hands

import "github.com/ksjbrown/poker/pkg/cards"

// Hand represents a Players current hand during a round.
//
// Initially, two cards are dealt to the player, which are assigned as the player's pocket cards.
// As the round continues, the player has access to more community cards, which are set via the Add*Cards methods.
//
// Hand should be a long-lived object, which remains with the Player instance.
// When new hands are dealt, the Cards in the Hand should be updated.
type Hand struct {
	PocketCards    *[2]*cards.Card
	CommunityCards *[5]*cards.Card
}

// AllCards returns a slice of cards available to the player to form a hand.
func (h *Hand) AllCards() []*cards.Card {
	allCards := h.PocketCards[:2]
	for _, communityCard := range h.CommunityCards {
		if communityCard != nil {
			allCards = append(allCards, communityCard)
		}
	}
	return allCards
}

func NewHand(p1 *cards.Card, p2 *cards.Card, c *[5]*cards.Card) *Hand {
	return &Hand{&[2]*cards.Card{p1, p2}, c}
}

// calculateScore analyses the cards available to this hand and returns a HandScore indicating the strength of this Hand relative to other Hands.
func (h *Hand) calculateScore() Score {
	// deck := cards.NewDeck()
	// pocket, _ := deck.DealCards(2)

	// community := [5]*cards.Card{}

	// hand := NewHand(pocket[0], pocket[1], &community)

	// flop, _ := deck.DealCards(3)
	// for i, c := range flop {
	// 	community[i] = c
	// }

	// turn, _ := deck.DealCard()
	// community[3] = turn

	// river, _ := deck.DealCard()
	// community[4] = river

	solver := NewSolver(h.AllCards())
	return solver.Solve()
}

// CompareTo is a shortcut for comparing a Hand's HandScore to another Hand's HandScore
func (h *Hand) CompareTo(other *Hand) int {
	return h.calculateScore().CompareTo(other.calculateScore())
}
