package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// creating a slice of deck
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newdeck() deck {
	cards := deck{}
	cardsuits := []string{"Spade", "Diamond", "Hearts", "Club"}
	cardvalue := []string{"Ace", "Two", "Three", "Four"}
	for _, i := range cardvalue {
		for _, j := range cardsuits {
			cards = append(cards, i+" of "+j)
		}
	}
	return cards
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) savetoFile(filename string) {
	ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

// create new deck from file

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range d {
		newposition := r.Intn(len(d) - 1)

		// swap the position of deck
		d[index], d[newposition] = d[newposition], d[index]
	}
}
