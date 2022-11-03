package main

import "fmt"

func main() {

	colors := make(map[string]string)
	// var colors map[string]string

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	colors["white"] = "#ffffff"

	delete(colors, "white")

	fmt.Println(colors)
}
