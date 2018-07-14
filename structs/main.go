package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jonContact := contactInfo{
		email:   "jonathan@panoply.io",
		zipCode: 94117,
	}
	jon := person{
		firstName:   "Jon",
		lastName:    "Hong",
		contactInfo: jonContact,
	}
	jonPointer := &jon // the &: "give us the memory address of this variable"
	jonPointer.updateName("Jonathan")
	jon.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	// * operator: gives us access to that variable's memory address
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// (pass by) value types
// int, float, string, bool, struct

// (pass by) reference types
// slices, maps, channels, pointers, functions
