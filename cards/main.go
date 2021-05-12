package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamond", newCard()}

	// Append to slice
	cards = append(cards, "Six of Clubs") // Note that append returns a new slice; it doesn't modifies the original one

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace of Spades"
}
