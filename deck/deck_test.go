package deck_test

import (
	"testing"

	"blackjack/deck"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	playable := deck.Deck()

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

func TestDeck_Draw(t *testing.T) {

}
