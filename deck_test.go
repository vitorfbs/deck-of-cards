package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	expectedLength := 52
	cards := newDeck()
	if len(cards) != expectedLength {
		t.Errorf("Expected deck length of %d, but got %v", expectedLength, len(cards))
	}
}

func TestNewDeckHasCorrectCards(t *testing.T) {
	foundCardsCount := 0
	missingCards := deck{}

	cards := newDeck()

	cardSuites := [4]string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Queen", "King"}

	for _, value := range cardValues {
		for _, suite := range cardSuites {
			testCard := fmt.Sprintf("%s of %s", value, suite)
			foundCardsCountBeforeCheck := foundCardsCount
			for _, card := range cards {
				if card == testCard {
					foundCardsCount++
				}
			}
			if foundCardsCountBeforeCheck == foundCardsCount {
				missingCards = append(missingCards, testCard)
			}
		}
	}

	if len(missingCards) != 0 {
		t.Errorf("Expected no missing cards, but found these missing: %s", missingCards.toString())
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	expectedLength := 52
	os.Remove("_decktesting")

	cards := newDeck()
	cards.saveToFile("_decktesting")

	loadedCards := newDeckFromFile("_decktesting")

	if len(loadedCards) != expectedLength {
		t.Errorf("Expected deck length of %d, but got %v", expectedLength, len(cards))
	}

	os.Remove("_decktesting")
}
