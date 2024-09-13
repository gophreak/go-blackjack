package blackjack

import (
	"fmt"

	"blackjack/chute"
	"blackjack/hand"
)

const (
	MinPlayers = 1
	MaxPlayers = 6
)

const (
	// 4 decks (208 cards)
	// max 6 players + dealer - impossible to run out of cards in chute per game
	// gives enough randomness to make it difficult to calculate cards
	numDecksForGame = 4

	// dealer must draw to minimum 16
	minDealerValue = 16
)

type Players []Player

type Player interface {
	Name() string
	Hand() *hand.Hand
}

type Dealer interface {
	Player
	Status() string
}

type Game struct {
	chute     chute.Chute
	dealer    Dealer
	players   []Player
	playerIdx int
}

func NewGame(players []Player, dealer Dealer) (*Game, error) {
	if MinPlayers > len(players) || MaxPlayers < len(players) {
		return nil, fmt.Errorf("invalid number of players. (%d - %d)", MinPlayers, MaxPlayers)
	}

	chute, err := chute.New(numDecksForGame)
	if err != nil {
		panic("developer error - invalid number of decks")
	}

	game := &Game{
		chute:   chute,
		dealer:  dealer,
		players: players,
	}

	game.setup()

	return game, nil
}

func (g *Game) Player() Player {
	idx := g.playerIdx
	if idx == len(g.players) {
		g.playerIdx = 0

		return nil
	}

	g.playerIdx++

	return g.players[idx]
}

func (g *Game) DrawCard(p Player) {
	p.Hand().AddCard(g.chute.Draw())
}

func (g *Game) Finish() {
	g.drawDealerToFinish()

	for {
		current := g.Player()
		if current == nil {
			break
		}

		current.Hand().CompareHand(*g.Dealer().Hand())
	}
}

func (g Game) Dealer() Dealer {
	return g.dealer
}

func (g *Game) setup() {
	for x := 0; x < 2; x++ {
		for {
			current := g.Player()
			if current == nil {
				break
			}
			g.DrawCard(current)

		}
		// although random, mimic dealer draws on first round
		if x == 0 {
			g.DrawCard(g.Dealer())
		}
	}
}

func (g *Game) drawDealerToFinish() {
	for g.dealer.Hand().GetMaxValue() < minDealerValue {
		g.DrawCard(g.dealer)
	}
}
