package main

import (
	"fmt"
	"strings"

	"github.com/oyugirachel/deck"
)

// Hand type
type Hand []deck.Card

func (h Hand) String() string{
	strs := make([]string,len(h))
	for i := range h{
		strs[i] = h[i].String()
	}
	return strings.Join(strs,",")
}
// DealerString 
func(h Hand) DealerString() string{
	return h[0].String() + ",**HIDDEN**"

}
func main(){
	cards := deck.New(deck.Deck(3),deck.Shuffle)
	var card deck.Card
	var player, dealer Hand
	for i :=0; i<2; i++{
	  for _, hand :=range[]*Hand{&player,&dealer}{
		 card,cards = draw(cards)
		// Points to a slice 
		*hand = append(*hand,card)

	}
}
	var input string
	// s representing standing
	for input != "s"{
		fmt.Println("Player:",player)
		fmt.Println("Dealer:",dealer.DealerString())
		fmt.Println("What will you do? (h)it,(s)tand")
		// Reading the users input
		fmt.Scanf("%s\n",&input)
	

	}
   
	
}

func draw(cards []deck.Card) (deck.Card,[]deck.Card){
	return cards[0], cards[1:]

}