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
type spanishBot struct{} /* Note that interfaces are implicit, i.e. we don't have to manually say that our
custom types satisy the interface.
*/

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

/*
More points about Interfaces:
* Interfaces are not generic types. Go (famously) does not have generic types.
* Interfaces are a contract to help us manage types.
*/
