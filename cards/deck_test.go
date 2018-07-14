package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedSize := 52

	if len(d) != expectedSize {
		t.Errorf("Expected deck length of %v but got %v", expectedSize, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Two of Hearts" {
		t.Errorf("Expected last card to be Two of Hearts but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// 1 defensive delete any files named "_decktesting"
	filename := "_decktesting"
	os.Remove(filename)
	// 2 create a deck
	deck := newDeck()
	// 3 save to file "_decktesting"
	deck.saveToFile(filename)
	// 4 load from file
	loadedDeck := newDeckFromFile(filename)
	// 5 assert deck length
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards but got %v", len(loadedDeck))
	}
	// 2 again delete any files named "_decktesting"
	os.Remove(filename)
}
