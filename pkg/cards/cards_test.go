package cards

import "testing"

func TestCard_NewCard(t *testing.T) {

	testCases := []struct {
		rank CardRank
		suit CardSuit
		err  bool
	}{
		{ACE, CLUBS, false},
		{0, CLUBS, true},
		{14, CLUBS, true},
		{ACE, 0, true},
		{ACE, 5, true},
	}

	for _, tc := range testCases {
		card, err := NewCard(tc.rank, tc.suit)
		if (err != nil) != tc.err {
			t.Errorf("expected error '%v', but was '%v'", tc.err, err)
		}
		if tc.err {
			continue
		}
		if card.Rank != tc.rank {
			t.Errorf("expected rank '%v', but was %v", card.Rank, tc.rank)
		}
		if card.Suit != tc.suit {
			t.Errorf("expected suit '%v', but was %v", card.Suit, tc.suit)
		}
	}
}

func TestCard_String(t *testing.T) {
	testCases := []struct {
		card     Card
		expected string
	}{
		{Card{ACE, CLUBS}, "AC"},
		{Card{TWO, DIAMONDS}, "2D"},
		{Card{TEN, SPADES}, "TS"},
		{Card{JACK, HEARTS}, "JH"},
		{Card{QUEEN, CLUBS}, "QC"},
		{Card{KING, CLUBS}, "KC"},
	}

	for _, tc := range testCases {
		actual := tc.card.String()
		if actual != tc.expected {
			t.Errorf("expected '%v', but got '%v'", tc.expected, actual)
		}
	}
}
