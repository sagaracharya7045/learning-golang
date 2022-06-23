package main

import "fmt"

func main() {
	cards, _ := deckFromFile("hello.txt")
	fmt.Println(cards)
}
