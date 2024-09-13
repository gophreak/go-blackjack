package deck

type Suit int

const (
	Hearts Suit = iota
	Diamonds
	Spades
	Clubs
	SuitCount int = iota
)

func (s Suit) String() string {
	return [...]string{"Hearts", "Diamonds", "Spades", "Clubs"}[s]
}
