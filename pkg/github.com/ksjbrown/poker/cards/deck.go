package cards

// Deck is a standard
type Deck struct {
	cards []Card
}

func NewDeck() Deck {
	deck := Deck{make([]Card, 0, 52)}
	for _, suit := range AllSuits() {
		for _, rank := range AllRanks() {
			deck.cards = append(deck.cards, NewCard(suit, rank))
		}
	}
	return deck
}
