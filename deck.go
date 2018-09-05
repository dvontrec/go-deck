package main

import "fmt"

// Create a new type of deck that is a slica of strings
type deck []string

// Function that loops through the deck and prints out the values
func (d deck) print() {
	// loops through the deck and prints out the values
	for i, card := range d {
		fmt.Println(i, card)
	}
}
