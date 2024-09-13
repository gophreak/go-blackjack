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
