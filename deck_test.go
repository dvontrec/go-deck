package main

import "testing"

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
