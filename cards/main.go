package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"

	cards := deck{"Ace of Spades", "Queen of Clubs"}

	cards = append(cards, newCard())

	// for index, instance := range card {
	// 	fmt.Println(index, instance)
	// }
	cards.print()

	// newDeck().print()
	fmt.Println(newDeck())
}

func newCard() string {
	return "Two of Hearts"
}
