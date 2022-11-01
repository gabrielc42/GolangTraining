package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	johan := person{firstName: "Johan", lastName: "Bolitkny"}
	fmt.Println(johan)

	//var johan person
	// johan.firstName = "Johan"
	// fmt.PrintF("%+v", johan)
}
