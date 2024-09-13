package dealer_test

import (
	"testing"

	"blackjack/dealer"
	"blackjack/deck"

	"github.com/stretchr/testify/assert"
)

func TestDealer_CreatesHand(t *testing.T) {
	assert.NotNil(t, dealer.New().Hand())
}

func TestDealer_Name_ReturnsName(t *testing.T) {
	assert.Equal(t, "Dealer", dealer.New().Name())
}

func TestDealer_Status_ReturnsDefault(t *testing.T) {
	assert.Equal(t, "STANDS", dealer.New().Status())
}

func TestDealer_Status_ReutnsBustWhenTooMany(t *testing.T) {
	d := dealer.New()

	d.Hand().AddCard(deck.MakeCard(deck.Hearts, deck.King))
	d.Hand().AddCard(deck.MakeCard(deck.Hearts, deck.King))
	d.Hand().AddCard(deck.MakeCard(deck.Hearts, deck.King))

	assert.Equal(t, "BUST", d.Status())
}

func TestDealer_Status_ReutnsBlackJackWhenBlackjack(t *testing.T) {
	d := dealer.New()

	d.Hand().AddCard(deck.MakeCard(deck.Hearts, deck.King))
	d.Hand().AddCard(deck.MakeCard(deck.Hearts, deck.Ace))

	assert.Equal(t, "BLACKJACK", d.Status())
}
