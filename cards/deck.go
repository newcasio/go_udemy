package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//create a new type of 'deck' which is a slice of strings
type deck []string

//Create function to create a new deck of cards.
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Clubs", "Hearts", "Spades", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suit)
		}
	}
	return cards
}

//Create a receiver.
//This is a method of type deck called 'print'.
//Only can be called on variables with the type 'deck'.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//Function to deal a specific number of cards from a deck
//takes arguments
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//take a deck (slice of strings), turn it into a single string
func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		//if error value is anything other than nil, log an error and quit program
		fmt.Println("Error: ", err)
		//use os package to exit if error
		os.Exit(1)
	}
}
