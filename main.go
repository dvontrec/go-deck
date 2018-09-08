// Tells that this is an exacutable file
package main

// Function that is run when main is run.
func main() {
	// Creates a dynamic sized array(SLICE) of strings
	cards := newDeckFromFile("mycards")
	cards.print()
}
