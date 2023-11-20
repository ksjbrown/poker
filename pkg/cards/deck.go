package cards

import (
	"fmt"
	"math/rand"
)

const (
	deckSize int = 52
)

// Deck defines a container for playing Cards, and provides some useful methods for working with a Deck.
//
// A standard deck consists of 52 cards.
// Support for a 54 card Deck with 2 Jokers may be considered in the future.
type Deck struct {

	// array of pointers to Cards, every element should not be nil if created by NewDeck
	Cards [deckSize]*Card

	// index of the next card to be dealt.
	nextCard int
}

// NewDeck creates a new deck of cards with 52 playing cards.
func NewDeck() *Deck {
	deck := Deck{[deckSize]*Card{}, 0}
	for i, suit := range AllSuits() {
		for j, rank := range AllRanks() {
			deck.Cards[i*13+j] = NewCard(suit, rank)
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
	count := len(d.Cards)
	for i := range d.Cards {
		j := i + rand.Intn(count-i)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
	d.Reset()
}

// DealCards will return a slice of *Card, whose length is equal to the argument int value.
//
// If the index is negative, an empty slice is returned with an error value.
// If the index exceeds the number of remaining cards in the deck, an empty slice is returned with an error value.
// Otherwise, the dealt cards internal counter is incremented, and a slice is returned, containing the same number of cards as the argument int value.
func (d *Deck) DealCards(i int) ([]*Card, error) {
	if i < 0 {
		return []*Card{}, fmt.Errorf("cannot deal negative number of cards: %v", i)
	}
	if i > d.CardsRemaining() {
		return []*Card{}, fmt.Errorf("attempted to deal %v cards, but only %v remaining", i, d.CardsRemaining())
	}
	cards := d.Cards[d.nextCard : d.nextCard+i]
	d.nextCard = min(len(d.Cards), d.nextCard+i)

	return cards, nil
}

// DealCard is the same as DealCards, but deals a single card.
//
// In case of the same error conditions, a nil pointer is returned.
func (d *Deck) DealCard() (*Card, error) {
	cards, err := d.DealCards(1)
	if len(cards) != 1 {
		return nil, err
	}
	return cards[0], err
}
