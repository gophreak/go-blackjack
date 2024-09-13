package chute

import (
	"errors"
	"time"

	"blackjack/deck"

	"golang.org/x/exp/rand"
)

type Chute struct {
	cards chan *deck.Card
}

func New(numDecks int) (Chute, error) {
	if numDecks < 1 {
		return Chute{}, errors.New("must have at least 1 deck")
	}
	var cards = make([]*deck.Card, numDecks*len(deck.Deck()))

	var i int
	for x := 0; x < numDecks; x++ {
		for _, card := range deck.Deck() {
			cards[i] = card
			i++
		}
	}

	return Chute{
		cards: shuffle(cards),
	}, nil
}

func shuffle(a []*deck.Card) chan *deck.Card {
	rand.Seed(uint64(time.Now().UnixNano()))
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	ch := make(chan *deck.Card, len(a))
	for i := 0; i < len(a); i++ {
		ch <- a[i]
	}

	close(ch)

	return ch
}

func (c Chute) Draw() *deck.Card {

	// cloe channel and range over
	for card := range c.cards {
		return card
	}

	return nil
}
