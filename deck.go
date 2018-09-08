package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Create a new type of deck that is a slica of strings
type deck []string

// Constructor function that returns a new deck of cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// Loops through the suits and values to create the appropriate card
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Function that loops through the deck and prints out the values
func (d deck) print() {
	// loops through the deck and prints out the values
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Function for dealing a hand returns 2 decks
func deal(d deck, handSize int) (deck, deck) {
	// Returns everything from the start of the deck to the handsize and return the deck with the handsized deck taken away
	return d[:handSize], d[handSize:]

}

// Creates a deck method that returns the dec as a string
func (d deck) toString() string {
	// converts deck into a string slice then joins each elemant by a comma into one string.
	return strings.Join([]string(d), ",")
}

// Creates a deck method that saves the deck to the harddrive
func (d deck) saveToFile(filename string) error {
	// Writes the deck to the given file as a byteslice
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// function that saves a new deck from a file
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	// If there is an error
	if err != nil {
		// Print the error
		fmt.Println("Error: ", err)
		// Quits the program
		os.Exit(1)
	}

	// Creates a slice of strings from the bytes
	s := strings.Split(string(bs), ",")
	// Returns the slice of srings converted into a deck
	return deck(s)
}
