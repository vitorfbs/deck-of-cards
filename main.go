package main

func main() {
	// declare variables for cards and hand
	// cards: the cards still in the deck available
	// hand: your hand of cards
	var cards deck
	var hand deck

	cardsFilePath := "cards"
	handFilePath := "hand"

	// check if the cards file exists locally
	if checkIfFileOrDirectoryExists(cardsFilePath) == nil {
		cards = newDeckFromFile(cardsFilePath)
	} else {
		cards = newDeck()
	}

	if checkIfFileOrDirectoryExists(handFilePath) == nil {
		hand = newDeckFromFile(handFilePath)
	} else {
		hand, cards = deal(cards, 5)
	}

	cards.shuffle()

	cards.print()
	hand.print()

	cards.saveToFile(cardsFilePath)
	hand.saveToFile(handFilePath)
}
