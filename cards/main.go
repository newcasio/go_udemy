package main

func main() {
	cards := newDeck()

	cards.saveToFile("my_deck_of_cards")
}
