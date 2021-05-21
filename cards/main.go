package main

func main() {
	/* 	cards := newDeck()

	   	hand, remainingCards := deal(cards, 5)

	   	hand.print()
	   	remainingCards.print()
	*/

	/* str := "Hello World!"
	// Type conversion in Go
	fmt.Println([]byte(str)) */

	cards := newDeck()
	// fmt.Println("Deck:\n" + cards.toString())
	cards.saveToFile("my_cards")

	newCards := newDeckFromFile("my_cards")
	newCards.print()

	// Error check
	// anotherCards := newDeckFromFile("invalid_file")
	// anotherCards.print()
}
