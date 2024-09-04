package main

func main() {
	cards := newDeckFromFile("cards")
	cards.shuffle()

	hand := newDeckFromFile("hand")

	cards.print()
	hand.print()

	hand.saveToFile("hand")
	cards.saveToFile("cards")
}
