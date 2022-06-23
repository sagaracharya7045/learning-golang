package main

import "fmt"

func main() {
	// cards := newDeck()
	// handSize := 3
	// hand, remainingDeck := deal(cards, handSize)
	// fmt.Println(cards.toString())
	// cards.saveToFile("hello.txt")
	pards, _ := deckFromFile("hello.txt")
	fmt.Println(pards)
	//check
	// hand.print()
	// remainingDeck.print()
}
