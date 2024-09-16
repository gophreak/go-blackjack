package deck

type deck [52]*Card

var d deck

func init() {
	z := 0

	for x := 0; x < 4; x++ {
		for y := 0; y < 13; y++ {
			d[z] = MakeCard(Suit(x), Rank(y))
			z++
		}
	}
}

func Deck() deck {
	return d
}


type Rank int

const (
	Ace Rank = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	RankCount int = iota
)

func (r Rank) String() string {
	return [...]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}[r]
}

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
