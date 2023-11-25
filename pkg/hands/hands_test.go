package hands

import (
	"testing"

	"github.com/ksjbrown/poker/pkg/cards"
)

func TestHand_Score(t *testing.T) {
	// ignore errors for NewHand
	newHand := func(cs ...cards.Card) Hand {
		h, _ := NewHand(cs...)
		return *h
	}

	// ignore errors for NewCard
	newCard := func(r cards.CardRank, s cards.CardSuit) cards.Card {
		c, _ := cards.NewCard(r, s)
		return *c
	}

	tests := []struct {
		rank  HandRank
		score int
		hand  Hand
	}{
		// Test Cases
		{
			HIGH_CARD,
			0x1E6543,
			newHand(
				newCard(cards.ACE, cards.HEARTS),
				newCard(cards.SIX, cards.HEARTS),
				newCard(cards.FIVE, cards.HEARTS),
				newCard(cards.FOUR, cards.HEARTS),
				newCard(cards.THREE, cards.DIAMONDS),
			),
		},
		{
			ONE_PAIR,
			0x22DCA0,
			newHand(
				newCard(cards.TWO, cards.HEARTS),
				newCard(cards.TWO, cards.DIAMONDS),
				newCard(cards.KING, cards.HEARTS),
				newCard(cards.QUEEN, cards.HEARTS),
				newCard(cards.TEN, cards.HEARTS),
			),
		},
		{
			TWO_PAIR,
			0x332500,
			newHand(
				newCard(cards.THREE, cards.HEARTS),
				newCard(cards.THREE, cards.DIAMONDS),
				newCard(cards.TWO, cards.HEARTS),
				newCard(cards.TWO, cards.DIAMONDS),
				newCard(cards.FIVE, cards.HEARTS),
			),
		},
		{
			THREE_OF_A_KIND,
			0x43ED00,
			newHand(
				newCard(cards.THREE, cards.CLUBS),
				newCard(cards.THREE, cards.HEARTS),
				newCard(cards.THREE, cards.DIAMONDS),
				newCard(cards.ACE, cards.DIAMONDS),
				newCard(cards.KING, cards.HEARTS),
			),
		},
		{
			STRAIGHT,
			0x560000,
			newHand(
				newCard(cards.TWO, cards.DIAMONDS),
				newCard(cards.THREE, cards.HEARTS),
				newCard(cards.FOUR, cards.CLUBS),
				newCard(cards.FIVE, cards.HEARTS),
				newCard(cards.SIX, cards.DIAMONDS),
			),
		},
		{
			STRAIGHT, // ace high
			0x5E0000,
			newHand(
				newCard(cards.ACE, cards.DIAMONDS),
				newCard(cards.TEN, cards.HEARTS),
				newCard(cards.JACK, cards.CLUBS),
				newCard(cards.QUEEN, cards.HEARTS),
				newCard(cards.KING, cards.DIAMONDS),
			),
		},
		{
			FLUSH,
			0x6E6543,
			newHand(
				newCard(cards.ACE, cards.HEARTS),
				newCard(cards.SIX, cards.HEARTS),
				newCard(cards.FIVE, cards.HEARTS),
				newCard(cards.FOUR, cards.HEARTS),
				newCard(cards.THREE, cards.HEARTS),
			),
		},
		{
			FULL_HOUSE,
			0x7ED000,
			newHand(
				newCard(cards.ACE, cards.HEARTS),
				newCard(cards.ACE, cards.CLUBS),
				newCard(cards.ACE, cards.DIAMONDS),
				newCard(cards.KING, cards.HEARTS),
				newCard(cards.KING, cards.CLUBS),
			),
		},
		{
			FOUR_OF_A_KIND,
			0x8ED000,
			newHand(
				newCard(cards.ACE, cards.HEARTS),
				newCard(cards.ACE, cards.CLUBS),
				newCard(cards.ACE, cards.DIAMONDS),
				newCard(cards.ACE, cards.SPADES),
				newCard(cards.KING, cards.CLUBS),
			),
		},
		{
			STRAIGHT_FLUSH,
			0x960000,
			newHand(
				newCard(cards.TWO, cards.HEARTS),
				newCard(cards.THREE, cards.HEARTS),
				newCard(cards.FOUR, cards.HEARTS),
				newCard(cards.FIVE, cards.HEARTS),
				newCard(cards.SIX, cards.HEARTS),
			),
		},
		{
			STRAIGHT_FLUSH,
			0x9E0000,
			newHand(
				newCard(cards.ACE, cards.HEARTS),
				newCard(cards.TEN, cards.HEARTS),
				newCard(cards.JACK, cards.HEARTS),
				newCard(cards.QUEEN, cards.HEARTS),
				newCard(cards.KING, cards.HEARTS),
			),
		},
	}

	// Execute tests
	for _, tc := range tests {
		expected := tc.score
		actual := tc.hand.Score()
		if expected != actual {
			t.Errorf("error in test '%v': expected score '%x', got '%x'", tc.rank, expected, actual)
		}
	}
}
