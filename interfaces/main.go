package main

import "fmt"

type bot interface {
	getGreeting() string
} /*
The above code basically says:
	1. Our program now has a new type called 'bot'.
	2. If you are a type in this program with a function named 'getGreeting()' that returns string, then
		you are now an honorary member of the type 'bot'.
*/

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
} // Any bot type can be passed here, i.e. any type with a function named 'getGreeting()' that returns string.

// We can skip the variable name for 'englishBot' in the receiver because we are not using it inside the function.
func (englishBot) getGreeting() string {
	// Custom logic for generating English greeting.
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	// Custom logic for generating Spanish greeting.
	return "Hola!"
}
