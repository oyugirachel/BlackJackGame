package main

import (
	"fmt"

	"github.com/oyugirachel/BlackJackGame/BlackJackAI"
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

func (ai *basicAI) Play(hand []deck.Card, dealer deck.Card) BlackJackAI.Move {
	score := BlackJackAI.Score(hand...)
	if len(hand) == 2 {
		if hand[0] == hand[1] {
			cardScore := BlackJackAI.Score(hand[0])
			if cardScore >= 8 && cardScore != 10 {
				return BlackJackAI.MoveSplit
			}
		}
		if (score == 10 || score == 11) && !BlackJackAI.Soft(hand...) {
			return BlackJackAI.MoveDouble
		}
	}
	dScore := BlackJackAI.Score(dealer)
	if dScore >= 5 && dScore <= 6 {
		return BlackJackAI.MoveStand
	}
	if score < 13 {
		return BlackJackAI.MoveHit
	}
	return BlackJackAI.MoveStand
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
	score := BlackJackAI.Score(card)
	switch {
	case score >= 10:
		ai.score--
	case score <= 6:
		ai.score++
	}
	ai.seen++
}
func main() {
	opts := BlackJackAI.Options{
		Decks:           5,
		Hands:           1000,
		BlackjackPayout: 1.5,
	}

	game := BlackJackAI.New(opts)
	winnings := game.Play(&basicAI{
		decks: 5,
	})
	fmt.Println(winnings)
}
