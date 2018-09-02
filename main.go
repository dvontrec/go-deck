// Tells that this is an exacutable file
package main

import "fmt"

// Function that is run when main is run.
func main() {
	// Creates a dynamic sized array(SLICE) of strings
	cards := []string{newCard(), newCard()}
	// creates a new array with the new card appended to the cards array
	cards = append(cards, newCard())

	// Loops through each element in the cards array as card
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// Defines a function that returns a string
func newCard() string {
	return "Five of Diamonds"
}
