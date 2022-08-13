package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	deal := d[0:handSize]
	cards := d[handSize:]
	return deal, cards
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	err := ioutil.WriteFile(filename, []byte(d.toString()), 0777)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func deckFromFile(filename string) (deck, error) {
	card, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return deck{}, err
	}
	return deck(strings.Split(string(card), ",")), nil //converting byte to string with string()
}

func (d deck) randomizeDeck() deck {
	rand.Seed(time.Now().UnixNano())
	for i := range d {
		randomPosition := rand.Intn(len(d))
		d[i], d[randomPosition] = d[randomPosition], d[i] //oneliner

	}
	return d
}
