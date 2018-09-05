package main

import "fmt"

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
