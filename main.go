// Tells that this is an exacutable file
package main

// Function that is run when main is run.
func main() {
	// Creates a dynamic sized array(SLICE) of strings
	cards := newDeck()

	// the hand and  the new deck are the return values of the deal function
	hand, remainingCards := deal(cards, 5)

	// calls the print method on the cards deck
	cards.print()
	hand.print()
	remainingCards.print()
}
