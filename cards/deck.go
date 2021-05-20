package main

import "fmt"

// Create a new type of 'deck' which is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "One", "Two", "Three"}

	// "_" tells Go that this value will not be used. Without it, we will get "declared but not used" error.
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	// (d deck) is a Receiver. Any variable of type "deck" now gets access to this "print" method
	for i, card := range d {
		fmt.Println(i, card)
	}
}
