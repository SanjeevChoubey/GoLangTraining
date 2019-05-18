package main

import (
	"os"
	"testing"
)

func TestNewdeck(t *testing.T) {

	d := newdeck()

	if len(d) != 16 {
		t.Errorf("Expected lenghth of deck is %v", len(d))
	}
}

func TestSavetoFileandnewDeckFromFile(t *testing.T) {
	os.Remove("_deckfile")

	deck := newdeck()
	deck.savetoFile("_deckfile")
	d := newDeckFromFile("_deckfile")
	if len(d) != 16 {
		t.Errorf("Expecting length of deck as 16 but got %v", len(d))
	}
}
