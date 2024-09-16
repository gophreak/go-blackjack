package hand_test

import (
	"testing"

	"blackjack/deck"
	"blackjack/hand"

	"github.com/stretchr/testify/assert"
)

func TestHand_AddCard(t *testing.T) {
	h := hand.New()

	assert.Equal(t, 0, h.Count())

	h.AddCard(deck.MakeCard(deck.Diamonds, deck.Ace))
	assert.Equal(t, 1, h.Count())

	h.AddCard(deck.MakeCard(deck.Spades, deck.Ten))
	assert.Equal(t, 2, h.Count())
}

func TestHand_GeTotal(t *testing.T) {
	testCases := map[string]struct {
		cards    []*deck.Card
		minTotal int
		maxTotal int
	}{
		"Single Ace": {
			[]*deck.Card{
				deck.MakeCard(deck.Clubs, deck.Ace),
			},
			1,
			11,
		},
		"Single Three": {
			[]*deck.Card{
				deck.MakeCard(deck.Clubs, deck.Three),
			},
			3,
			3,
		},
		"Single Seven": {
			[]*deck.Card{
				deck.MakeCard(deck.Clubs, deck.Seven),
			},
			7,
			7,
		},
		"Single Jack": {
			[]*deck.Card{
				deck.MakeCard(deck.Clubs, deck.Jack),
			},
			10,
			10,
		},
		"Single Queen": {
			[]*deck.Card{
				deck.MakeCard(deck.Clubs, deck.Queen),
			},
			10,
			10,
		},
		"Single King": {
			[]*deck.Card{
				deck.MakeCard(deck.Clubs, deck.King),
			},
			10,
			10,
		},
		"Ace and King": {
			[]*deck.Card{
				deck.MakeCard(deck.Clubs, deck.King),
				deck.MakeCard(deck.Spades, deck.Ace),
			},
			11,
			21,
		},
		"Four and Seven": {
			[]*deck.Card{
				deck.MakeCard(deck.Clubs, deck.Four),
				deck.MakeCard(deck.Spades, deck.Seven),
			},
			11,
			11,
		},
		"Two and Two": {
			[]*deck.Card{
				deck.MakeCard(deck.Clubs, deck.Two),
				deck.MakeCard(deck.Spades, deck.Two),
			},
			4,
			4,
		},
		"Two, Six and Nine": {
			[]*deck.Card{
				deck.MakeCard(deck.Clubs, deck.Two),
				deck.MakeCard(deck.Hearts, deck.Six),
				deck.MakeCard(deck.Diamonds, deck.Nine),
			},
			17,
			17,
		},
		"Two Aces": {
			[]*deck.Card{
				deck.MakeCard(deck.Clubs, deck.Ace),
				deck.MakeCard(deck.Hearts, deck.Ace),
			},
			2,
			12,
		},
		"Three Aces": {
			[]*deck.Card{
				deck.MakeCard(deck.Clubs, deck.Ace),
				deck.MakeCard(deck.Hearts, deck.Ace),
				deck.MakeCard(deck.Hearts, deck.Ace),
			},
			3,
			13,
		},
		"Over normal": {
			[]*deck.Card{
				deck.MakeCard(deck.Clubs, deck.Nine),
				deck.MakeCard(deck.Hearts, deck.King),
				deck.MakeCard(deck.Hearts, deck.Seven),
			},
			26,
			26,
		},
		"Over and Under with Ace": {
			[]*deck.Card{
				deck.MakeCard(deck.Clubs, deck.Ace),
				deck.MakeCard(deck.Hearts, deck.King),
				deck.MakeCard(deck.Hearts, deck.Seven),
			},
			18,
			18,
		},
		"Over and Under with Ace last": {
			[]*deck.Card{
				deck.MakeCard(deck.Hearts, deck.King),
				deck.MakeCard(deck.Hearts, deck.Seven),
				deck.MakeCard(deck.Clubs, deck.Ace),
			},
			18,
			18,
		},
		"Over and Under with Ace middle": {
			[]*deck.Card{
				deck.MakeCard(deck.Hearts, deck.Seven),
				deck.MakeCard(deck.Clubs, deck.Ace),
				deck.MakeCard(deck.Hearts, deck.King),
			},
			18,
			18,
		},
		"Over and Under with 3 Aces": {
			[]*deck.Card{
				deck.MakeCard(deck.Hearts, deck.Five),
				deck.MakeCard(deck.Clubs, deck.Ace),
				deck.MakeCard(deck.Clubs, deck.Ace),
				deck.MakeCard(deck.Hearts, deck.Four),
			},
			11,
			21,
		},
	}

	for name, data := range testCases {
		t.Run(name, func(t *testing.T) {
			h := hand.New()
			assert.Equal(t, 0, h.Count())
			min, max := h.GetMinValue(), h.GetMaxValue()
			assert.Equal(t, 0, min)
			assert.Equal(t, 0, max)

			for _, c := range data.cards {
				h.AddCard(c)
			}

			min, max = h.GetMinValue(), h.GetMaxValue()
			assert.Equal(t, data.minTotal, min)
			assert.Equal(t, data.maxTotal, max)
		})
	}
}

func TestHand_GetCards(t *testing.T) {
	h := hand.New()

	h.AddCard(deck.MakeCard(deck.Diamonds, deck.Ace))
	h.AddCard(deck.MakeCard(deck.Spades, deck.Ten))
	h.AddCard(deck.MakeCard(deck.Hearts, deck.Four))

	cards := h.GetCards()
	assert.Len(t, cards, 3)

	assert.Equal(t, deck.MakeCard(deck.Diamonds, deck.Ace), cards[0])
	assert.Equal(t, deck.MakeCard(deck.Spades, deck.Ten), cards[1])
	assert.Equal(t, deck.MakeCard(deck.Hearts, deck.Four), cards[2])
}

func TestHand_CanPromptReturnsTrueOnlyWhenPlayerNotBustOrBlackJack(t *testing.T) {
	h := hand.New()

	h.AddCard(deck.MakeCard(deck.Diamonds, deck.Five))
	h.AddCard(deck.MakeCard(deck.Spades, deck.Ten))

	assert.True(t, h.CanPrompt())

	h.AddCard(deck.MakeCard(deck.Spades, deck.Ten))

	assert.False(t, h.CanPrompt(), "hand should be bust")
}

func TestHand_CanPromptReturnsTrueOnlyWhenPlayerNotBust(t *testing.T) {
	h := hand.New()

	h.AddCard(deck.MakeCard(deck.Diamonds, deck.Five))
	h.AddCard(deck.MakeCard(deck.Spades, deck.Ten))

	assert.True(t, h.CanPrompt())

	h.AddCard(deck.MakeCard(deck.Spades, deck.Ten))

	assert.False(t, h.CanPrompt(), "hand should be bust")
}

func TestHand_IsBustReturnsTrueOnlyWhenPlayerIsBust(t *testing.T) {
	h := hand.New()

	h.AddCard(deck.MakeCard(deck.Diamonds, deck.Five))
	assert.False(t, h.IsBust(), "hand should not have bust")

	h.AddCard(deck.MakeCard(deck.Spades, deck.Ten)) // 15
	assert.False(t, h.IsBust(), "hand should not have bust")

	h.AddCard(deck.MakeCard(deck.Spades, deck.Ace)) // 16 / 26 - 16
	assert.False(t, h.IsBust(), "hand should not have bust")

	h.AddCard(deck.MakeCard(deck.Diamonds, deck.Four)) // 20 / 30 - 20
	assert.False(t, h.IsBust(), "hand should not have bust")

	h.AddCard(deck.MakeCard(deck.Hearts, deck.Two)) // 22 / 32 - 22
	assert.True(t, h.IsBust(), "hand should be bust")
}

func TestHand_HasBlackjackReturnsTrueOnlyIfBlackjackConditionIsMet(t *testing.T) {
	h := hand.New()

	h.AddCard(deck.MakeCard(deck.Diamonds, deck.Ten))
	h.AddCard(deck.MakeCard(deck.Diamonds, deck.King))
	h.AddCard(deck.MakeCard(deck.Diamonds, deck.Ace)) // 21 / 31 - 21 -> Not blackjack (3 cards)
	assert.False(t, h.HasBlackjack(), "hand should not have Blackjack with 3 cards")

	h2 := hand.New()
	h2.AddCard(deck.MakeCard(deck.Diamonds, deck.King))
	h2.AddCard(deck.MakeCard(deck.Diamonds, deck.Ace)) // 11 / 21 - 21 -> Blackjack (2 cards)
	assert.True(t, h2.HasBlackjack(), "hand should have Blackjack")

	h3 := hand.New()
	h3.AddCard(deck.MakeCard(deck.Diamonds, deck.Ace))
	h3.AddCard(deck.MakeCard(deck.Diamonds, deck.Ten)) // 11 / 21 - 21 -> Blackjack (2 cards)
	assert.True(t, h3.HasBlackjack(), "hand should have Blackjack")
}

func TestHand_CompareHand_PlayerLosesWhenBust(t *testing.T) {
	// Player hand is 23 - bust
	playerHand := hand.New()
	playerHand.AddCard(deck.MakeCard(deck.Clubs, deck.King))
	playerHand.AddCard(deck.MakeCard(deck.Clubs, deck.Queen))
	playerHand.AddCard(deck.MakeCard(deck.Clubs, deck.Three))

	testCases := map[string]struct {
		dealerHand *hand.Hand
	}{
		"Dealer is not bust": {
			dealerHand: func() *hand.Hand {
				h := hand.New()
				h.AddCard(deck.MakeCard(deck.Clubs, deck.King))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Queen))

				return h
			}(),
		},
		"Dealer is blackjack": {
			dealerHand: func() *hand.Hand {
				h := hand.New()
				h.AddCard(deck.MakeCard(deck.Clubs, deck.King))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Ace))

				return h
			}(),
		},
		"Dealer is 21 - not blackjack": {
			dealerHand: func() *hand.Hand {
				h := hand.New()
				h.AddCard(deck.MakeCard(deck.Clubs, deck.King))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Ten))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Ace))

				return h
			}(),
		},
		"Dealer is bust (lower)": {
			dealerHand: func() *hand.Hand {
				h := hand.New()
				h.AddCard(deck.MakeCard(deck.Clubs, deck.King))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Queen))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Two))

				return h
			}(),
		},
		"Dealer is bust (higher)": {
			dealerHand: func() *hand.Hand {
				h := hand.New()
				h.AddCard(deck.MakeCard(deck.Clubs, deck.King))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Queen))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Four))

				return h
			}(),
		},
		"Dealer is bust (equal)": {
			dealerHand: func() *hand.Hand {
				h := hand.New()
				h.AddCard(deck.MakeCard(deck.Clubs, deck.King))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Queen))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Three))

				return h
			}(),
		},
	}

	for name, data := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, 0, int(playerHand.CompareHand((*data.dealerHand))))
		})
	}
}

func TestHand_CompareHand_AgainstDealerBlackjack(t *testing.T) {
	// dealer hand is blackjack
	dealerHand := hand.New()
	dealerHand.AddCard(deck.MakeCard(deck.Clubs, deck.King))
	dealerHand.AddCard(deck.MakeCard(deck.Clubs, deck.Ace))

	testCases := map[string]struct {
		playerHand     *hand.Hand
		expectedStatus int
	}{
		"Player is bust - 22": {
			playerHand: func() *hand.Hand {
				h := hand.New()
				h.AddCard(deck.MakeCard(deck.Clubs, deck.King))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Queen))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Two))

				return h
			}(),
			expectedStatus: 0, // lost
		},
		"Player is 21 - not blackjack": {
			playerHand: func() *hand.Hand {
				h := hand.New()
				h.AddCard(deck.MakeCard(deck.Clubs, deck.King))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Ten))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Ace))

				return h
			}(),
			expectedStatus: 0, // lost
		},
		"Player is under - 20": {
			playerHand: func() *hand.Hand {
				h := hand.New()
				h.AddCard(deck.MakeCard(deck.Clubs, deck.King))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Ten))

				return h
			}(),
			expectedStatus: 0, // lost
		},
		"Player is blackjack": {
			playerHand: func() *hand.Hand {
				h := hand.New()
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Jack))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Ace))

				return h
			}(),
			expectedStatus: 1, // draw
		},
	}

	for name, data := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, data.expectedStatus, int(data.playerHand.CompareHand(*dealerHand)))
		})
	}
}

func TestHand_CompareHand_AgainstDealerBust(t *testing.T) {
	// dealer hand is bust - 23
	dealerHand := hand.New()
	dealerHand.AddCard(deck.MakeCard(deck.Clubs, deck.King))
	dealerHand.AddCard(deck.MakeCard(deck.Clubs, deck.Queen))
	dealerHand.AddCard(deck.MakeCard(deck.Clubs, deck.Three))

	testCases := map[string]struct {
		playerHand     *hand.Hand
		expectedStatus int
	}{
		"Player is bust - 22 (under)": {
			playerHand: func() *hand.Hand {
				h := hand.New()
				h.AddCard(deck.MakeCard(deck.Clubs, deck.King))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Queen))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Two))

				return h
			}(),
			expectedStatus: 0, // lost
		},
		"Player is bust - 24 (over)": {
			playerHand: func() *hand.Hand {
				h := hand.New()
				h.AddCard(deck.MakeCard(deck.Clubs, deck.King))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Queen))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Four))

				return h
			}(),
			expectedStatus: 0, // lost
		},
		"Player is 21 - not blackjack": {
			playerHand: func() *hand.Hand {
				h := hand.New()
				h.AddCard(deck.MakeCard(deck.Clubs, deck.King))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Ten))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Ace))

				return h
			}(),
			expectedStatus: 2, // win
		},
		"Player is under - 20": {
			playerHand: func() *hand.Hand {
				h := hand.New()
				h.AddCard(deck.MakeCard(deck.Clubs, deck.King))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Ten))

				return h
			}(),
			expectedStatus: 2, // win
		},
		"Player is under - 4": {
			playerHand: func() *hand.Hand {
				h := hand.New()
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Two))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Two))

				return h
			}(),
			expectedStatus: 2, // win
		},
		"Player is blackjack": {
			playerHand: func() *hand.Hand {
				h := hand.New()
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Jack))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Ace))

				return h
			}(),
			expectedStatus: 3, // blackjack
		},
	}

	for name, data := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, data.expectedStatus, int(data.playerHand.CompareHand(*dealerHand)))
		})
	}
}

func TestHand_CompareHand_AgainstDealerHighValue(t *testing.T) {
	// dealer hand is - 19
	dealerHand := hand.New()
	dealerHand.AddCard(deck.MakeCard(deck.Clubs, deck.Ten))
	dealerHand.AddCard(deck.MakeCard(deck.Clubs, deck.Nine))

	testCases := map[string]struct {
		playerHand     *hand.Hand
		expectedStatus int
	}{
		"Player is bust - 22": {
			playerHand: func() *hand.Hand {
				h := hand.New()
				h.AddCard(deck.MakeCard(deck.Clubs, deck.King))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Queen))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Two))

				return h
			}(),
			expectedStatus: 0, // lost
		},
		"Player is 21 - not blackjack": {
			playerHand: func() *hand.Hand {
				h := hand.New()
				h.AddCard(deck.MakeCard(deck.Clubs, deck.King))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Ten))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Ace))

				return h
			}(),
			expectedStatus: 2, // win
		},
		"Player is higher - 20": {
			playerHand: func() *hand.Hand {
				h := hand.New()
				h.AddCard(deck.MakeCard(deck.Clubs, deck.King))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Ten))

				return h
			}(),
			expectedStatus: 2, // win
		},
		"Player is equal - 19": {
			playerHand: func() *hand.Hand {
				h := hand.New()
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Nine))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Queen))

				return h
			}(),
			expectedStatus: 1, // draw
		},
		"Player is under - 18": {
			playerHand: func() *hand.Hand {
				h := hand.New()
				h.AddCard(deck.MakeCard(deck.Clubs, deck.King))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Eight))

				return h
			}(),
			expectedStatus: 0, // lost
		},
		"Player is blackjack": {
			playerHand: func() *hand.Hand {
				h := hand.New()
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Jack))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Ace))

				return h
			}(),
			expectedStatus: 3, // blackjack
		},
	}

	for name, data := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, data.expectedStatus, int(data.playerHand.CompareHand(*dealerHand)))
		})
	}
}

func TestStatus_String_ReturnsAppropriateMessage(t *testing.T) {
	// dealer hand is - 19
	dealerHand := hand.New()
	dealerHand.AddCard(deck.MakeCard(deck.Clubs, deck.Ten))
	dealerHand.AddCard(deck.MakeCard(deck.Clubs, deck.Nine))

	testCases := map[string]struct {
		playerHand     *hand.Hand
		expectedStatus string
	}{
		"Player Lost status": {
			playerHand: func() *hand.Hand {
				h := hand.New()
				// 18 < 19
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Jack))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Eight))

				return h
			}(),
			expectedStatus: "LOST",
		},
		"Player Draw status": {
			playerHand: func() *hand.Hand {
				h := hand.New()
				// 18 = 19
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Ten))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Nine))

				return h
			}(),
			expectedStatus: "DRAW",
		},
		"Player Win status - 20": {
			playerHand: func() *hand.Hand {
				h := hand.New()
				// 20 > 19
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Ten))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Jack))

				return h
			}(),
			expectedStatus: "WIN",
		},
		"Player Win status - 21 - not blackjack": {
			playerHand: func() *hand.Hand {
				h := hand.New()
				// 21 > 19
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Ten))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Jack))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Ace))

				return h
			}(),
			expectedStatus: "WIN",
		},
		"Player Blackjack status": {
			playerHand: func() *hand.Hand {
				h := hand.New()
				// 21 > 19
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Ten))
				h.AddCard(deck.MakeCard(deck.Clubs, deck.Ace))

				return h
			}(),
			expectedStatus: "BLACKJACK",
		},
	}

	for name, data := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, data.expectedStatus, data.playerHand.CompareHand(*dealerHand).String())
		})
	}
}
