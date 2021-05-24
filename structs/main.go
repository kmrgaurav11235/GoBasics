package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	leto := person{
		firstName: "Leto",
		lastName:  "Atreides",
		contact: contactInfo{
			email:   "leto.da.boss@caladan.gov",
			zipCode: 560076,
		},
	}
	/*
		Note that the commas are present at the end of every line above, even when it is the last line.
		This is required when we define structs using multiple lines.
	*/

	fmt.Printf("%+v \n", leto)
}
