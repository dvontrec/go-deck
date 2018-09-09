package main

import (
	"os"
	"testing"
)

// Tests  the NewDeck function
func TestNewDeck(t *testing.T) {
	d := newDeck()
	// Tests the length of the deck to be 52 cards
	if len(d) != 52 {
		// Go template literal
		t.Errorf("Expected deck length of 52 but got %v.", len(d))
	}

	// Tests the deck base contents
	if d[0] != "Ace of Spades" {
		t.Errorf(("Expected Ace Of Spades, but got %v."), d[0])
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected Four of Clubs, but got %v.", d[len(d)-1])
	}
}

// Tests saving and loading deck on harddrive
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// removes old test files
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	// Checks that the deck was properly saved and can be loaded
	if len(loadedDeck) != 52 {
		t.Errorf("Expected a 52 card deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
