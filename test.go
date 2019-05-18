package main

func main() {
	// Sav file
	// cards := newdeck()
	// cards.savetoFile("Firstfile")

	// read from file
	// cards := newDeckFromFile("Firstfile")
	// cards.print()

	// shuffle the cards
	cards := newDeckFromFile("Firstfile")
	cards.shuffle()
	cards.print()

	// hand, remainingcards := deal(cards, 5)
	// //cards.print()
	// hand.print()
	// // remaining cards
	// remainingcards.print()
}
