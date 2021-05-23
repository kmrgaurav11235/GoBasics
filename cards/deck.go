package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck' which is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

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

// Go supports multiple return values
func deal(d deck, handSize int) (deck, deck) {
	// Slice range syntax variable[startIndexInclusive:endIndexNonInclusive]. Default values are 0 and 'n' respectively.
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	byteSlice, err := ioutil.ReadFile(fileName)

	// This pattern of checking the error immediately after function call is a very common pattern in Go.
	if err != nil {
		fmt.Println("[ERROR]:", err)
		// Exit causes the current program to exit with the given status code. Parameter 0 indicates success, non-zero an error.
		// The program terminates immediately; deferred functions are not run
		os.Exit(1)
	}

	s := strings.Split(string(byteSlice), ",") // string slice
	return deck(s)
}

func (d deck) shuffle() {
	// Seed value for random number generator
	seed := time.Now().UnixNano()
	rand.Seed(seed)

	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
