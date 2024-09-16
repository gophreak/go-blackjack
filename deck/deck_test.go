package deck_test

import (
	"testing"

	"blackjack/deck"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	playable := deck.Init()

	assert.Len(t, playable, 52)
	assert.Equal(t, "Ace", playable[0].GetRank())
	assert.Equal(t, "Hearts", playable[0].GetSuit())

	assert.Equal(t, "Two", playable[1].GetRank())
	assert.Equal(t, "Hearts", playable[1].GetSuit())

	assert.Equal(t, "King", playable[12].GetRank())
	assert.Equal(t, "Hearts", playable[12].GetSuit())

	assert.Equal(t, "Ace", playable[13].GetRank())
	assert.Equal(t, "Diamonds", playable[13].GetSuit())

	assert.Equal(t, "King", playable[51].GetRank())
	assert.Equal(t, "Clubs", playable[51].GetSuit())
}

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

func TestSuit_String(t *testing.T) {
	assert.Equal(t, "Diamonds", deck.Diamonds.String())
	assert.Equal(t, "Spades", deck.Spades.String())
	assert.Equal(t, "Hearts", deck.Hearts.String())
	assert.Equal(t, "Clubs", deck.Clubs.String())
}

func TestRankSizeIs4(t *testing.T) {
	assert.Equal(t, 4, deck.SuitCount)
}
