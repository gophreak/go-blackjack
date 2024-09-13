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
	val := int(c.rank) + 1
	if val > 10 {
		return 10
	}
	return val
}
