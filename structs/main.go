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

	johan := person{
		firstName: "Johan",
		lastName:  "Bolitkny",
		contactInfo: contactInfo{
			email:   "bolitknyj@yahoo.com",
			zipCode: 00001,
		},
	}
	// johanPointer := &johan //&variable is a reference, memory address pointer
	johan.updateName("Jimmy")
	johan.print()

	// turn address into value with *address
	// turn value into address with &value

	// johan := person{firstName: "Johan", lastName: "Bolitkny"}
	// fmt.Println(johan)

	//var johan person
	// johan.firstName = "Johan"
	// fmt.PrintF("%+v", johan)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) { //*type description pointer to the type
	(*pointerToPerson).firstName = newFirstName //*variable is value pointer, operator
}
