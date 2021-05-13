package main

import "fmt"

// Create a new type of 'deck' which is a slice of string
type deck []string

func (d deck) print() {
	// (d deck) is a Receiver. Any variable of type "deck" now gets access to this "print" method
	for i, card := range d {
		fmt.Println(i, card)
	}
}
