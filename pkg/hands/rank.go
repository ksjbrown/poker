package hands

// Rank defines the well-known hand types of poker.
// These are used as the major ranking score when comparing hands
// Take card not to confuse this struct with the cards.Rank struct, related to the face value of a card.
type Rank int

const (
	HIGH_CARD Rank = iota + 1
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	STRAIGHT
	FLUSH
	FULL_HOUSE
	FOUR_OF_A_KIND
	STRAIGHT_FLUSH
)
