package main

import "fmt"

func main() {

	cards := newDeck()

	// hand, remaindingCards := deal(cards, 5)

	// hand.print()
	// remaindingCards.print()

	var test = cards.toString()
	fmt.Println([]byte(test))
}
