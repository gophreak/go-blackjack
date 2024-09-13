package dealer

import "blackjack/hand"

const dealerName = "Dealer"

type Dealer struct {
	hand *hand.Hand
}

func New() Dealer {
	return Dealer{
		hand: hand.New(),
	}
}

func (d Dealer) Hand() *hand.Hand {
	return d.hand
}

func (d Dealer) Name() string {
	return dealerName
}

func (d Dealer) Status() string {
	if d.hand.IsBust() {
		return "BUST"
	}

	if d.hand.HasBlackjack() {
		return "BLACKJACK"
	}

	return "STANDS"
}
