package main

// variables can be init outside of func, but not assigned values

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades" // := only for new variables, or above (var variableName type)
	// cannot do card = "Ace of Spades" (assigned value before init)
	// card = "Five of Diamonds" // reassigning

	// card := newCard()

	// fmt.Println(card)

	cards := deck{newCard(), newCard()}
	cards = append(cards, newCard())

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
