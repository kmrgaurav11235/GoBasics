package main

func main() {
	cards := deck{"Ace of Diamond", newCard()}

	// Append to slice
	cards = append(cards, "Six of Clubs") // Note that append returns a new slice; it doesn't modifies the original one

	cards.print()
}

func newCard() string {
	return "Ace of Spades"
}
