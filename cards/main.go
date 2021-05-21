package main

import "fmt"

func main() {
	/* 	cards := newDeck()

	   	hand, remainingCards := deal(cards, 5)

	   	hand.print()
	   	remainingCards.print()
	*/

	str := "Hello World!"
	// Type conversion in Go
	fmt.Println([]byte(str))

	cards := newDeck()
	fmt.Println("Deck:\n" + cards.toString())
}
