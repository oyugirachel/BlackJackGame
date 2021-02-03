package main

import (
	"fmt"

	"github.com/oyugirachel/BlackJackGame/BlackJackAi"
)

func main() {
	game := BlackJackAi.New()
	winnings := game.Play(BlackJackAi.HumanAI())
	fmt.Println(winnings)
}
