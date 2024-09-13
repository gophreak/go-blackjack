package player

import (
	"blackjack/hand"
)

type Player struct {
	name string
	hand *hand.Hand
}

func New(name string) *Player {
	return &Player{
		hand: hand.New(),
		name: name,
	}
}

func (p Player) Hand() *hand.Hand {
	return p.hand
}

func (p Player) Name() string {
	return p.name
}
