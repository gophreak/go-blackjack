package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"blackjack/blackjack"
	"blackjack/dealer"
	"blackjack/player"
)

func main() {
	var game *blackjack.Game
	var gameError error

	var dealer = dealer.New()

	// find number of players
	for {
		fmt.Print("How many players? ")

		numPlayers := validateNum(readLine())

		players := make([]blackjack.Player, numPlayers)

		for i := 0; i < numPlayers; i++ {
			players[i] = blackjack.Player(player.New(fmt.Sprintf("Player %d", i+1)))
		}

		game, gameError = blackjack.NewGame(players, dealer)
		if game != nil && gameError == nil {
			break
		}

		fmt.Printf("%s\n", gameError.Error())
	}

	// Print dealer hand
	fmt.Println("")
	fmt.Println(game.Dealer().Name())
	fmt.Println("-------")
	for _, c := range game.Dealer().Hand().GetCards() {
		fmt.Printf("%s of %s\n", c.GetRank(), c.GetSuit())
	}
	fmt.Println("")
	fmt.Println(renderTotal(game.Dealer()))

	fmt.Println("-----------------")
	// Print player hand
	for {
		current := game.Player()
		if current == nil {
			break
		}

		fmt.Println("")
		fmt.Println(current.Name())
		fmt.Println("---------")
		for _, c := range current.Hand().GetCards() {
			fmt.Printf("%s of %s\n", c.GetRank(), c.GetSuit())
		}
		fmt.Println("")
		fmt.Println(renderTotal(current))
		fmt.Println("-----------------")
	}

	fmt.Println("=================")

	// Promp players to stand or draw
	for {
		current := game.Player()
		if current == nil {
			break
		}

		var response string

		for current.Hand().CanPrompt() && response != "S" && response != "s" {
			response = ""
			var errorMessage string

			for response != "D" && response != "S" && response != "d" && response != "s" {
				fmt.Printf("%s%s (%s), Draw(D) or Stick(S)? ", errorMessage, current.Name(), renderTotal(current))
				response = readLine()

				errorMessage = "Invalid option: "
			}
			if response == "D" || response == "d" {
				game.DrawCard(current)

				fmt.Println("")
				fmt.Println(current.Name())
				fmt.Println("---------")
				for _, c := range current.Hand().GetCards() {
					fmt.Printf("%s of %s\n", c.GetRank(), c.GetSuit())
				}
				fmt.Println("")
				fmt.Println(renderTotal(current))
				fmt.Println("-----------------")
			}
		}
	}

	fmt.Println("==================================")
	fmt.Println("")

	game.Finish()

	// Print all final hands
	fmt.Println(game.Dealer().Name())
	fmt.Println("-------")
	for _, c := range game.Dealer().Hand().GetCards() {
		fmt.Printf("%s of %s\n", c.GetRank(), c.GetSuit())
	}
	fmt.Println("")
	fmt.Println(renderTotal(game.Dealer()))

	fmt.Println("-----------------")
	for {
		current := game.Player()
		if current == nil {
			break
		}

		fmt.Println("")
		fmt.Println(current.Name())
		fmt.Println("---------")
		for _, c := range current.Hand().GetCards() {
			fmt.Printf("%s of %s\n", c.GetRank(), c.GetSuit())
		}
		fmt.Println("")
		fmt.Println(renderTotal(current))
		fmt.Println("-----------------")
	}

	fmt.Println("=================")
	fmt.Println("=================")

	fmt.Println("")

	fmt.Println(game.Dealer().Name(), renderTotal(game.Dealer()))
	for {
		current := game.Player()
		if current == nil {
			break
		}
		fmt.Printf("%s (%d): %s\n", current.Name(), current.Hand().GetMaxValue(), current.Hand().GetStatus().String())
	}
	fmt.Println("")
}

func validateNum(n string) int {
	num, err := strconv.Atoi(strings.Replace(n, "\n", "", 1))
	if err != nil {
		return 0
	}

	return num
}

func readLine() string {
	line, _ := bufio.NewReader(bufio.NewReader(os.Stdin)).ReadString('\n')

	return strings.Replace(line, "\n", "", 1)
}

func renderTotal(p blackjack.Player) string {
	var total string
	total += fmt.Sprintf("Total: %d", p.Hand().GetMinValue())
	if p.Hand().GetMinValue() != p.Hand().GetMaxValue() {
		total += fmt.Sprintf(" / %d", p.Hand().GetMaxValue())
	}
	if p.Hand().HasBlackjack() {
		total += " -- BLACKJACK"
	}
	if p.Hand().IsBust() {
		total += " -- BUST"
	}

	return total
}
