package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := getCardSuites()
	cardValues := getCardValues()

	for _, value := range cardValues {
		for _, suite := range cardSuites {
			card := fmt.Sprintf("%s of %s", value, suite)
			cards = append(cards, card)

		}
	}

	return cards
}

func newDeckFromFile(filename string) deck {
	byteSliceCards, err := readFromFile(filename)
	if err != nil {
		fmt.Println(err.Error() + ": " + filename)
		os.Exit(1)
	}
	return deck(strings.Split(string(byteSliceCards), ","))
}

func deal(d deck, numberOfCards int) (deck, deck) {
	start := len(d) - numberOfCards

	if start < 0 {
		start = 0
	}

	return d[start:], d[:start]
}

func (d deck) shuffle() deck {
	for i := range d {
		targetIndex := rand.Intn(len(d) - 1)

		d[i], d[targetIndex] = d[targetIndex], d[i]
	}

	return d
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	byteSliceDeck := []byte(d.toString())
	return writeToFile(filename, byteSliceDeck)
}
