package deck_test

import (
	"testing"

	"blackjack/deck"

	"github.com/stretchr/testify/assert"
)

func TestRank_String(t *testing.T) {
	assert.Equal(t, "Ace", deck.Ace.String())
	assert.Equal(t, "Two", deck.Two.String())
	assert.Equal(t, "Three", deck.Three.String())
	assert.Equal(t, "Four", deck.Four.String())
	assert.Equal(t, "Five", deck.Five.String())
	assert.Equal(t, "Six", deck.Six.String())
	assert.Equal(t, "Seven", deck.Seven.String())
	assert.Equal(t, "Eight", deck.Eight.String())
	assert.Equal(t, "Nine", deck.Nine.String())
	assert.Equal(t, "Ten", deck.Ten.String())
	assert.Equal(t, "Jack", deck.Jack.String())
	assert.Equal(t, "Queen", deck.Queen.String())
	assert.Equal(t, "King", deck.King.String())
}

func TestRankSizeIs13(t *testing.T) {
	assert.Equal(t, 13, deck.RankCount)
}
