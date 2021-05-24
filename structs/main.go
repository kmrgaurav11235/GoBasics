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

	var leto person            // 'leto' is declared but not assigned a value
	fmt.Println(leto)          // Will show the default "" value for strings firstName and lastName
	fmt.Printf("%+v \n", leto) // Prints the field and value

	leto.lastName = "Atreides"
	leto.firstName = "Leto"
	fmt.Println(leto)
	fmt.Printf("%+v \n", leto)
}
