// Tells that this is an exacutable file
package main

// Function that is run when main is run.
func main() {
	// Creates a dynamic sized array(SLICE) of strings
	cards := deck{newCard(), newCard()}
	// creates a new array with the new card appended to the cards array
	cards = append(cards, "Six of diamonds")

	// calls the print method on the cards deck
	cards.print()
}

// Defines a function that returns a string
func newCard() string {
	return "Five of Diamonds"
}
