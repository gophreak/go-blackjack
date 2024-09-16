package hand

import (
	"blackjack/deck"
)

const (
	blackjackValue = 21
)

type Status int

const (
	Lost Status = iota
	Draw
	Win
	Blackjack
)

type options int

const (
	ready options = iota
	pass
	bust
)

type Hand struct {
	cards    []*deck.Card
	min, max int
	option   options
}

func New() *Hand {
	return &Hand{}
}

func (h Hand) GetCards() []*deck.Card {
	return h.cards
}

func (h Hand) CanPrompt() bool {
	return ready == h.option
}

func (h Hand) IsBust() bool {
	return bust == h.option
}

func (h Hand) HasBlackjack() bool {
	return h.Count() == 2 && blackjackValue == h.max
}

func (h *Hand) AddCard(card *deck.Card) {
	h.cards = append(h.cards, card)

	h.calculate()
}

func (h Hand) GetMaxValue() int {
	return h.max
}

func (h Hand) GetMinValue() int {
	return h.min
}

func (h Hand) Count() int {
	return len(h.cards)
}

func (h *Hand) CompareHand(dealerHand Hand) Status {
	// hand is assumed lost, unless otherwise

	// hand bust, don't do anything
	if h.IsBust() {
		return Lost
	}

	// dealer has blackjack - if hand has blackjack then draw, otherwise lost
	if dealerHand.HasBlackjack() {
		if h.HasBlackjack() {
			return Draw
		}

		return Lost
	}

	// dealer is not blackjack, hand wins if blackjack
	if h.HasBlackjack() {
		return Blackjack
	}

	// dealer is bust, hand wins if not bust
	if dealerHand.IsBust() {
		return Win
	}

	// dealer is not bust or blackjack - compare maximum hand values
	valueToBeat := dealerHand.GetMaxValue()

	// hand has higher value - win
	if h.max > valueToBeat {
		return Win
	}

	// hand has same value - draw
	if h.max == valueToBeat {
		return Draw
	}

	// hand is not bust, does not have blackjack, and has less maximum hand value than dealer - lost (default)
	return Lost
}

func (h *Hand) calculate() {
	h.min, h.max = 0, 0

	for i := 0; i < len(h.cards); i++ {
		cardVal := h.cards[i].GetValue()
		h.min += cardVal
		h.max += cardVal

		// only apply higher value if doesn't exceed 21
		if cardVal == 1 && h.max+10 <= blackjackValue {
			h.max += 10
		}
	}
	// if max is over 21, just return min value
	if h.max > blackjackValue {
		h.max = h.min
	}

	if h.max == blackjackValue {
		h.option = pass

		return
	}

	if h.GetMinValue() > blackjackValue {
		h.option = bust

		return
	}
}

func (s Status) String() string {
	return [...]string{"LOST", "DRAW", "WIN", "BLACKJACK"}[s]
}
