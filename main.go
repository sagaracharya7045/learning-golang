package main

// func doesFileExist(filename string) {
// 	fileName := ""
// 	_, error := os.Stat(fileName)
// 	// check if error is "file not exists"
// 	if os.IsNotExist(error) {
// 		fmt.Printf("%v file does not exist\n", fileName)
// 	} else {
// 		fmt.Printf("%v file exist\n", fileName)
// 	}
// }
func main() {
	// cards := newDeck()
	// cards.saveToFile("hello.txt")
	// doesFileExist("hello.txt")
	readcards, _ := deckFromFile("hello.txt")
	// fmt.Println(readcards)
	readcards.print()
}

// func main() {
// 	fmt.Println("Hello world")
// }
