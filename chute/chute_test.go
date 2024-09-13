package chute_test

import (
	"testing"

	"blackjack/chute"
	"blackjack/deck"

	"github.com/stretchr/testify/assert"
)

func TestNew_FailsOnZeroDeck(t *testing.T) {
	_, err := chute.New(0)

	assert.Error(t, err)
	assert.Equal(t, "must have at least 1 deck", err.Error())
}

func TestDraw_ReturnsCard(t *testing.T) {
	gameChute, err := chute.New(1)

	assert.NoError(t, err)
	assert.NotEmpty(t, gameChute.Draw())
}

func TestDraw_ReturnsAllCardsInDeck(t *testing.T) {
	testCases := map[string]struct {
		numDecks int
		numCards int
	}{
		"single deck": {
			1, 52,
		},
		"two decks": {
			2, 104,
		},
		"four decks": {
			4, 208,
		},
		"six decks": {
			6, 312,
		},
	}

	for name, data := range testCases {
		t.Run(name, func(t *testing.T) {
			gameChute, err := chute.New(data.numDecks)
			assert.NoError(t, err)

			var numCards, numSameCard int

			for {
				drawn := gameChute.Draw()
				if drawn == nil {
					break
				}

				if drawn.GetRank() == deck.Ace.String() && drawn.GetSuit() == deck.Spades.String() {
					numSameCard++
				}
				numCards++
			}

			assert.Equal(t, data.numDecks, numSameCard)
			assert.Equal(t, data.numCards, numCards)
		})
	}
}
