package blackjack_test

import (
	"testing"

	"blackjack/blackjack"
	"blackjack/hand"

	"github.com/stretchr/testify/assert"
)

func Test_NewGame_ReturnsErrorNotEnoughPlayers(t *testing.T) {
	game, err := blackjack.NewGame([]blackjack.Player{}, newMockDealer())

	assert.Nil(t, game)
	assert.Error(t, err)
	assert.Equal(t, "invalid number of players. (1 - 6)", err.Error())
}

func Test_NewGame_ReturnsErrorTooManyPlayers(t *testing.T) {
	game, err := blackjack.NewGame([]blackjack.Player{
		newMockPlayer(),
		newMockPlayer(),
		newMockPlayer(),
		newMockPlayer(),
		newMockPlayer(),
		newMockPlayer(),
		newMockPlayer(),
	}, newMockDealer())

	assert.Nil(t, game)
	assert.Error(t, err)
	assert.Equal(t, "invalid number of players. (1 - 6)", err.Error())
}

func Test_NewGame_ReturnsGameMinPlayers(t *testing.T) {
	game, err := blackjack.NewGame([]blackjack.Player{newMockPlayer()}, newMockDealer())

	assert.NotNil(t, game)
	assert.NoError(t, err)
}

func Test_NewGame_ReturnsGameMaxPlayers(t *testing.T) {
	game, err := blackjack.NewGame([]blackjack.Player{
		newMockPlayer(),
		newMockPlayer(),
		newMockPlayer(),
		newMockPlayer(),
		newMockPlayer(),
		newMockPlayer(),
	}, newMockDealer())

	assert.NotNil(t, game)
	assert.NoError(t, err)
}

func Test_NewGame_SetsUpNumberOfPlayers(t *testing.T) {
	testCases := map[string]struct {
		players []blackjack.Player
	}{
		"six players": {
			[]blackjack.Player{
				newMockPlayer(),
				newMockPlayer(),
				newMockPlayer(),
				newMockPlayer(),
				newMockPlayer(),
				newMockPlayer(),
			},
		},
		"five players": {
			[]blackjack.Player{
				newMockPlayer(),
				newMockPlayer(),
				newMockPlayer(),
				newMockPlayer(),
				newMockPlayer(),
			},
		},
		"four players": {
			[]blackjack.Player{
				newMockPlayer(),
				newMockPlayer(),
				newMockPlayer(),
				newMockPlayer(),
			},
		},
		"three players": {
			[]blackjack.Player{
				newMockPlayer(),
				newMockPlayer(),
				newMockPlayer(),
			},
		},
		"two players": {
			[]blackjack.Player{
				newMockPlayer(),
				newMockPlayer(),
			},
		},
		"one player": {
			[]blackjack.Player{
				newMockPlayer(),
			},
		},
	}

	for name, data := range testCases {
		t.Run(name, func(t *testing.T) {
			game, err := blackjack.NewGame(data.players, newMockDealer())
			assert.NotNil(t, game)
			assert.NoError(t, err)

			var counter int
			for {
				current := game.Player()
				if current == nil {
					break
				}
				counter++
			}

			assert.Equal(t, counter, len(data.players))
		})
	}
}

func TestGame_DrawCard_AddsCardToHand(t *testing.T) {
	game, err := blackjack.NewGame([]blackjack.Player{newMockPlayer()}, newMockDealer())
	assert.NoError(t, err)
	assert.NotNil(t, game)

	current := game.Player()
	assert.Equal(t, 2, current.Hand().Count())

	game.DrawCard(current)
	assert.Equal(t, 3, current.Hand().Count())

	game.DrawCard(current)
	assert.Equal(t, 4, current.Hand().Count())

	game.DrawCard(current)
	assert.Equal(t, 5, current.Hand().Count())
}

func TestGame_Draws2CardsPerPlayer1ForDealer(t *testing.T) {
	sixPlayers := []blackjack.Player{newMockPlayer(), newMockPlayer(), newMockPlayer(), newMockPlayer(), newMockPlayer(), newMockPlayer()}
	game, err := blackjack.NewGame(sixPlayers, newMockDealer())
	assert.NoError(t, err)
	assert.NotNil(t, game)

	for {
		current := game.Player()
		if current == nil {
			break
		}
		assert.Equal(t, 2, current.Hand().Count())
	}

	assert.Equal(t, 1, game.Dealer().Hand().Count())
}

func TestGame_DrawCardToFinishALwaysHitsMinimum_16(t *testing.T) {
	// simulate 100 times to ensure never below 16
	for x := 0; x < 100; x++ {
		game, err := blackjack.NewGame([]blackjack.Player{newMockPlayer()}, newMockDealer())
		assert.NoError(t, err)
		assert.NotNil(t, game)

		game.Finish()

		assert.LessOrEqual(t, 16, game.Dealer().Hand().GetMaxValue())
	}
}

type MockPlayer struct {
	hand *hand.Hand
	name string
}

func (p MockPlayer) Hand() *hand.Hand {
	return p.hand
}
func (p MockPlayer) Name() string {
	return p.name
}

type MockDealer struct {
	hand   *hand.Hand
	name   string
	status string
}

func (d MockDealer) Hand() *hand.Hand {
	return d.hand
}
func (d MockDealer) Name() string {
	return d.name
}
func (d MockDealer) Status() string {
	return d.status
}

func newMockPlayer() *MockPlayer {
	return &MockPlayer{
		hand: hand.New(),
	}
}

func newMockDealer() *MockDealer {
	return &MockDealer{
		hand: hand.New(),
	}
}
