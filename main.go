// Tells that this is an exacutable file
package main

import "fmt"

// Function that is run when main is run.
func main() {
	// Creates a new infered string variable called card
	card := newCard()
	fmt.Println(card)
}

// Defines a function that returns a string
func newCard() string {
	return "Five of Diamonds"
}
