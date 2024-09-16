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

type Hand struct {
	cards []*deck.Card
}

func New() *Hand {
	return &Hand{}
}

func (h Hand) GetCards() []*deck.Card {
	return h.cards
}

func (h Hand) CanPrompt() bool {
	return !h.IsBust()
}

func (h Hand) IsBust() bool {
	return h.GetMinValue() > blackjackValue
}

func (h Hand) HasBlackjack() bool {
	return h.Count() == 2 && blackjackValue == h.GetMaxValue()
}

func (h *Hand) AddCard(card *deck.Card) {
	h.cards = append(h.cards, card)
}

func (h Hand) GetMaxValue() int {
	var val int
	for i := 0; i < len(h.cards); i++ {
		cardVal := h.cards[i].GetValue()
		val += cardVal

		// only apply higher value if doesn't exceed 21
		if cardVal == 1 && val+10 <= blackjackValue {
			val += 10
		}
	}
	// if max is over 21, just return min value
	if val > blackjackValue {
		return h.GetMinValue()
	}

	return val
}

func (h Hand) GetMinValue() int {
	var val int

	for i := 0; i < len(h.cards); i++ {
		val += h.cards[i].GetValue()
	}

	return val
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

	max := h.GetMaxValue()
	// hand has higher value - win
	if max > valueToBeat {
		return Win
	}

	// hand has same value - draw
	if max == valueToBeat {
		return Draw
	}

	// hand is not bust, does not have blackjack, and has less maximum hand value than dealer - lost (default)
	return Lost
}

func (s Status) String() string {
	return [...]string{"LOST", "DRAW", "WIN", "BLACKJACK"}[s]
}
