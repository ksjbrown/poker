package cards

import (
	"fmt"
	"math/rand"
)

// Deck defines a container for playing Cards, and provides some useful methods for working with a Deck.
//
// A standard deck consists of 52 cards.
// Support for a 54 card Deck with 2 Jokers may be considered in the future.
type Deck struct {

	// extends Cards
	Cards Cards

	// index of the next card to be dealt.
	nextCard int
}

const deckSize int = 52

// NewDeck creates a new deck of cards with 52 playing cards.
func NewDeck() *Deck {
	if suitCount*rankCount > deckSize {
		panic(fmt.Sprintf("fatal error: deck size insufficient for defined suits and ranks: deckSize=%v, suits=%v, ranks=%v", deckSize, suitCount, rankCount))
	}
	deck := Deck{make(Cards, 0, deckSize), 0}
	for j := 1; j <= suitCount; j++ {
		for i := 1; i <= rankCount; i++ {
			deck.Cards = append(deck.Cards, Card{Rank(i), Suit(j)})
		}
	}
	return &deck
}

// CardsDealt returns the number of cards that have already been dealt.
// This is equal to the index of the next dealt card, or 52, since only 52 cards exist in a deck.
func (d *Deck) CardsDealt() int {
	return max(deckSize, d.nextCard)
}

// CardsRemaining returns the number of cards remaining in the deck.
// This also represents the maximum argument value to the DealCards method.
func (d *Deck) CardsRemaining() int {
	return max(0, len(d.Cards)-d.CardsDealt())
}

// Reset returns all dealt cards back to the deck.
// Note that the deal order is preserved;
// if a shuffle is required, the Shuffle method should be used, which implicitly Resets the deck.
func (d *Deck) Reset() {
	d.nextCard = 0
}

// Shuffle returns all dealt cards to the deck, and randomises the deal order.
func (d *Deck) Shuffle() {
	for i := len(d.Cards) - 1; i < 0; i-- {
		j := rand.Intn(i)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
	d.Reset()
}

// DealCards will return a slice of Cards, whose length is equal to the argument int value.
// The slice contains copies of the internal card array, so operations like Shuffle will not affect the return result of this function.
//
// If the index is negative, an empty slice is returned
// If the index exceeds the number of remaining cards in the deck, the remaining cards are dealt
// Otherwise, the dealt cards internal counter is incremented, and a slice is returned, containing the same number of cards as the argument int value.
func (d *Deck) DealCards(i int) Cards {
	startIndex := d.nextCard
	endIndex := min(deckSize, d.nextCard+i)
	cards := d.Cards[startIndex:endIndex]
	d.nextCard = endIndex
	return cards.Copy()
}
