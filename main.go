package main

func main() {

	// cards := newDeck()

	// cards.saveToFile("my_cards.docx")

	cards := newDeckFromFile("my_cards.txt")
	cards.print()
}
