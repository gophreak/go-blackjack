package deck

type Card struct {
	suit Suit
	rank Rank
}

func MakeCard(s Suit, r Rank) *Card {
	return &Card{
		suit: s,
		rank: r,
	}
}

func (c Card) GetSuit() string {
	return c.suit.String()
}

func (c Card) GetRank() string {
	return c.rank.String()
}

func (c Card) GetValue() int {
	return func(idx int) int {
		val := idx + 1 // zero index
		switch {
		case val > 10:
			return 10
		}

		return val
	}(int(c.rank))
}
