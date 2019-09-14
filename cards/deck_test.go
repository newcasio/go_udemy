package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 52 {
		t.Errorf("Expected deck length 52, but got %v", len(deck))
	}

}

func TestFirstCard(t *testing.T) {
	deck := newDeck()

	if deck[0] != "Ace of Clubs" {
		t.Errorf("Expected first card to be Ace of Clubs, but got %v", deck[0])
	}
}

func TestLastCard(t *testing.T) {
	deck := newDeck()

	if deck[len(deck)-1] != "King of Diamonds" {
		t.Errorf("Expect last card to be King of Diamonds, but got %v", deck[len(deck)-1])
	}
}

//save deck and load deck
func TestSaveDeckToFileAndNewDeckFromFile(t *testing.T) {
	//remove any saved deck with same deck name so no false positive when testing if file was created.
	os.Remove("_deckTesting")

	_deckTesting := newDeck()

	if len(_deckTesting) != 52 {
		t.Errorf("Expected deck length 52, but got %v", len(_deckTesting))
	}

	os.Remove("_deckTesting")
}
