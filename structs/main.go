package main

import "fmt"

func main() {
	mySlice := []string{"Fear", "is", "the", "mind-killer"}
	updateSlice(mySlice)
	fmt.Println(mySlice) // This will print "[Fear is the little-death]" even though we passed by value
	// This is because slice is a reference type not a value type
	/*
		Value Types:
			* int, float, string, bool, struct
			* Use pointers to change the value types in a function
		Reference Types:
			* slice, map, channel, pointer, function
			* Don't worry about pointers with reference types
	*/
}

func updateSlice(s []string) {
	s[3] = "little-death"
}

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // is equivalent to "contactInfo contactInfo"
}

func mainOld() {
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

	// letoPointer := &leto
	// // &variable = Give me the memory address of the value that this variable is pointing at.
	// letoPointer.updateFitstName("Paul")

	// Instead of the above code, we can use a shortcut
	leto.updateFitstName("Paul")

	leto.print()
}

// Below *person is a type description. It means we are working with a "Pointer to a person".
func (pointerToPerson *person) updateFitstName(newFirstName string) {
	// (*pointerToPerson).firstName = newFirstName
	// // Above *pointer is an operator. It means, "Give me the value at this memory address".

	// Instead of the above code, we can use a shortcut
	(*pointerToPerson).firstName = newFirstName
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
