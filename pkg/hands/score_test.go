package hands

import (
	"testing"

	"github.com/ksjbrown/poker/pkg/cards"
)

func TestNewScore(t *testing.T) {
	type expected = struct {
		value int
		major int
		minor int
		micro int
	}
	testCases := []struct {
		score    Score
		expected expected
	}{
		// happy paths
		{*NewScore(0, 0, 0), expected{0, 0, 0, 0}},
		{*NewScore(1, 2, 3, 4), expected{0x123400, 1, 2, 0x3400}},
		{*NewScore(1, 2, 3, 4, 5, 6), expected{0x123456, 1, 2, 0x3456}},

		// sad paths
		{*NewScore(0, 0, 0, 0, 0, 0, 1), expected{0, 0, 0, 0}},
	}

	for _, tc := range testCases {
		value := int(tc.score)
		major := tc.score.Major()
		minor := tc.score.Minor()
		micro := tc.score.Micro()

		if value != tc.expected.value {
			t.Errorf("expected value '%x', got '%x'", tc.expected.value, value)
		}
		if major != tc.expected.major {
			t.Errorf("expected major '%x', got '%x'", tc.expected.major, major)
		}
		if minor != tc.expected.minor {
			t.Errorf("expected minor '%x', got '%x'", tc.expected.minor, minor)
		}
		if micro != tc.expected.micro {
			t.Errorf("expected major '%x', got '%x'", tc.expected.micro, micro)
		}
	}
}

func TestNewScoreExtraValuesIgnores(t *testing.T) {
	score := NewScore(0, 0, 0, 0, 0, 0, 1)
	if int(*score) != 0 {
		t.Errorf("expected 0, got %x", score)
	}
}

func TestIsOnePair(t *testing.T) {
	hand := Hand{
		cards.Card{Rank: cards.ACE, Suit: cards.CLUBS},
		cards.Card{Rank: cards.TWO, Suit: cards.CLUBS},
		cards.Card{Rank: cards.THREE, Suit: cards.CLUBS},
		cards.Card{Rank: cards.FOUR, Suit: cards.CLUBS},
		cards.Card{Rank: cards.ACE, Suit: cards.DIAMONDS},
	}
	if !hand.isOnePair() {
		t.Error("should be one pair")
	}
	if hand.isTwoPair() {
		t.Error("should not be two pair")
	}
}

func TestIsTwoPair(t *testing.T) {
	hand := Hand{
		cards.Card{Rank: cards.ACE, Suit: cards.CLUBS},
		cards.Card{Rank: cards.TWO, Suit: cards.CLUBS},
		cards.Card{Rank: cards.THREE, Suit: cards.CLUBS},
		cards.Card{Rank: cards.ACE, Suit: cards.DIAMONDS},
		cards.Card{Rank: cards.TWO, Suit: cards.DIAMONDS},
	}
	if !hand.isOnePair() {
		t.Error("should be one pair")
	}
	if !hand.isTwoPair() {
		t.Error("should be two pair")
	}
}
