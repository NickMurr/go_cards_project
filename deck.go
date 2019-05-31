package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

/*
Create a new type of 'deck'
which is a slice if strings
*/

type deck []string

// Create a Deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamomds", "Hearts", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Print all Cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, "-", card)
	}
}

// Deal cards to Hand
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Converting deck to String
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Saving data to the File
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Read From The File
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {

		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	outStringSlice := strings.Split(string(bs), ",")
	return deck(outStringSlice)
}

// Shuffle cards
func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
