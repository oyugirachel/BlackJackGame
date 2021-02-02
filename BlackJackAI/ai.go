package blackjack


// AI interface
type AI interface {
	Bet() int
	Play(hand []deck.Card, dealer deck.Card) Move
	Results(hand [][]deck.Card, dealer []deck.Card)



}
// HumanAI struct
type HumanAI struct{


}
// Bet function
func (ai *HumanAI) Bet() int{
	return 1
}
// Play function
func (ai *HumanAI) Play(hand []deck.Card, dealer deck.Card) Move{
	
	for{
		fmt.Println("Player:", hand)
	    fmt.Println("Dealer:", dealer)
	    fmt.Println("What will you do? (h)it,(s)tand")
		 
		 var input string
		// Reading the users input
		
	    fmt.Scanf("%s\n", &input)
	    switch input {
	    case "h":
		  return Hit
	    case "s":
		  return Stand
	    default:
		fmt.Println("Invalid Option", input)
	}
	}
	
}
// Results function
func (ai *HumanAI) Results(hand [][]deck.Card, dealer []deck.Card){
	fmt.Println("**FINAL HANDS**")
	fmt.Println("Player", hand)
	fmt.Println("Dealer", dealer)
}

// filler to be implemented later

// Move type
type Move func(GameState) GameState
// GameState struct
type GameState struct{}
// Hit function
func Hit(gs GameState) GameState{
	return gs
}
// Stand function
func Stand(gs GameState) GameState{
	return gs
}


