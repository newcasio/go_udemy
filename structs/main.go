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
	jim := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "what@here.com",
			zipCode: 12123,
		},
	}
	jimPointer := &jim
	jimPointer.updateName("Jim")
	// these are equal
	jim.updateName("Jim")

	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
