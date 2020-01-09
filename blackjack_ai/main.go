package main

import (
	"fmt"

	"github.com/kazijawad/Gophercises/blackjack_ai/blackjack"
)

func main() {
	game := blackjack.New()
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}
