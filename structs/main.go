package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // is equivalent to "contactInfo contactInfo"
}

func main() {
	leto := person{
		firstName: "Leto",
		lastName:  "Atreides",
		contactInfo: contactInfo{
			email:   "leto.da.boss@caladan.gov",
			zipCode: 560076,
		},
	}
	/*
		Note that the commas are present at the end of every line above, even when it is the last line.
		This is required when we define structs using multiple lines.
	*/

	letoPointer := &leto
	// &variable = Give me the memory address of the value that this variable is pointing at.
	letoPointer.updateFitstName("Paul")
	leto.print()
}

// Below *person is a type description. It means we are working with a "Pointer to a person".
func (pointerToPerson *person) updateFitstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
	// Above *pointer is an operator. It means, "Give me the value at this memory address".
}

/*
As we see above "*Type" is completely different from "*PointerVariable"
	1. *type is a type description.
	2. *pointer is an operator.

Also, the two operators:
	1. Turn address into value with "*address"
	2. Turn value into address with "&value"
*/

func (p person) print() {
	fmt.Printf("%+v \n", p)
}
