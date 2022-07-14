package main

import (
	"fmt"
)

func main() {
	// cards := newDeck()
	// cards.saveToFile("hello.txt")
	readcards, _ := deckFromFile("hello.txt")
	fmt.Println(readcards)
}
