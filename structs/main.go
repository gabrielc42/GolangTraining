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

	//johan.updateName("jimmy")
	johan.print()

	// johan := person{firstName: "Johan", lastName: "Bolitkny"}
	// fmt.Println(johan)

	//var johan person
	// johan.firstName = "Johan"
	// fmt.PrintF("%+v", johan)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
