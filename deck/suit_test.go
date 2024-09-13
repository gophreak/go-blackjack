package deck_test

import (
	"testing"

	"blackjack/deck"

	"github.com/stretchr/testify/assert"
)

func TestSuit_String(t *testing.T) {
	assert.Equal(t, "Diamonds", deck.Diamonds.String())
	assert.Equal(t, "Spades", deck.Spades.String())
	assert.Equal(t, "Hearts", deck.Hearts.String())
	assert.Equal(t, "Clubs", deck.Clubs.String())
}

func TestRankSizeIs4(t *testing.T) {
	assert.Equal(t, 4, deck.SuitCount)
}
