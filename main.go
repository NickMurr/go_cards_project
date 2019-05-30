package main

func main() {

	cards := newDeck()

	hand, remaindingCards := deal(cards, 5)

	hand.print()
	remaindingCards.print()

}
