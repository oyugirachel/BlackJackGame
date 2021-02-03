package main

import (
	"fmt"

	"github.com/oyugirachel/BlackJackGame/BlackJackAi"
	"github.com/oyugirachel/deck"
)

type basicAI struct {
	score int
	seen  int
	decks int
}

func (ai *basicAI) Bet(shuffled bool) int {
	if shuffled {
		ai.score = 0
		ai.seen = 0
	}
	trueScore := ai.score / ((ai.decks*52 - ai.seen) / 52)
	switch {
	case trueScore >= 14:
		return 100000
	case trueScore >= 8:
		return 5000
	default:
		return 100
	}
}

func (ai *basicAI) Play(hand []deck.Card, dealer deck.Card) BlackJackAi.Move {
	score := BlackJackAi.Score(hand...)
	if len(hand) == 2 {
		if hand[0] == hand[1] {
			cardScore := BlackJackAi.Score(hand[0])
			if cardScore >= 8 && cardScore != 10 {
				return BlackJackAi.MoveSplit
			}
		}
		if (score == 10 || score == 11) && !BlackJackAi.Soft(hand...) {
			return BlackJackAi.MoveDouble
		}
	}
	dScore := BlackJackAi.Score(dealer)
	if dScore >= 5 && dScore <= 6 {
		return BlackJackAi.MoveStand
	}
	if score < 13 {
		return BlackJackAi.MoveHit
	}
	return BlackJackAi.MoveStand
}

func (ai *basicAI) Results(hands [][]deck.Card, dealer []deck.Card) {
	for _, card := range dealer {
		ai.count(card)
	}
	for _, hand := range hands {
		for _, card := range hand {
			ai.count(card)
		}
	}
}

func (ai *basicAI) count(card deck.Card) {
	score := BlackJackAi.Score(card)
	switch {
	case score >= 10:
		ai.score--
	case score <= 6:
		ai.score++
	}
	ai.seen++
}
func main() {
	opts := BlackJackAi.Options{
		Decks:           5,
		Hands:           1000,
		BlackjackPayout: 1.5,
	}

	game := BlackJackAi.New(opts)
	winnings := game.Play(&basicAI{
		decks: 5,
	})
	fmt.Println(winnings)
}
