package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// paul := person{"Paul", "Atreides"}
	paul := person{firstName: "Paul", lastName: "Atreides"} // Better way

	fmt.Println(paul)
}
