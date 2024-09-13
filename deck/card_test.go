package deck_test

import (
	"testing"

	"blackjack/deck"

	"github.com/stretchr/testify/assert"
)

var testCases = map[string]struct {
	suit               deck.Suit
	rank               deck.Rank
	suitName, rankName string
	faceValue          int
}{
	"Ace of Spades": {
		deck.Spades, deck.Ace, "Spades", "Ace", 1,
	},
	"Three of Diamonds": {
		deck.Diamonds, deck.Three, "Diamonds", "Three", 3,
	},
	"Nine of Hearts": {
		deck.Hearts, deck.Nine, "Hearts", "Nine", 9,
	},
	"King of Clubs": {
		deck.Clubs, deck.King, "Clubs", "King", 10,
	},
	"Four of Spades": {
		deck.Spades, deck.Four, "Spades", "Four", 4,
	},
}

func TestMake_GetRank(t *testing.T) {
	for name, data := range testCases {
		t.Run(name, func(t *testing.T) {
			c := deck.MakeCard(data.suit, data.rank)

			assert.Equal(t, data.rankName, c.GetRank())
		})
	}
}

func TestMake_GetSuit(t *testing.T) {
	for name, data := range testCases {
		t.Run(name, func(t *testing.T) {
			c := deck.MakeCard(data.suit, data.rank)

			assert.Equal(t, data.suitName, c.GetSuit())
		})
	}
}

func TestMake_Display(t *testing.T) {
	for name, data := range testCases {
		t.Run(name, func(t *testing.T) {
			c := deck.MakeCard(data.suit, data.rank)

			assert.Equal(t, data.suit.String(), c.GetSuit())
			assert.Equal(t, data.rank.String(), c.GetRank())
		})
	}
}

func TestMake_GetFaceValue(t *testing.T) {
	for name, data := range testCases {
		t.Run(name, func(t *testing.T) {
			c := deck.MakeCard(data.suit, data.rank)

			assert.Equal(t, data.faceValue, c.GetValue())
		})
	}
}
