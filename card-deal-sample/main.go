package main

import "fmt"

func main() {

	// calling the newDeck func
	cards := newDeck()
	fmt.Println("Dealing the deck")
	hand, remainingDeck := deal(cards, 5)
	hand.print()
	remainingDeck.print()
	// print as string
	fmt.Println(hand.toString())

	// save to file
	hand.saveToFile("my_hand")

	newDeck := newDeckFromFile("my_hand")

	fmt.Println("Deck form file")
	newDeck.print()

	fmt.Println("Shuffling")
	newDeck.shuffle()

	newDeck.print()

	// sampleSyntax()
}

func sampleSyntax() {

	//variable declaration
	var card string = "Ace of Spades"
	fmt.Println("Printing variables")
	fmt.Println(card)
	// another way to declare
	card1 := "Ace of Hearts"
	fmt.Println(card1)

	// var cards []string = []string{"Ace of Diamonds"}
	// above can be written as
	cardSlice := []string{"Ace of Diamonds", "Ace of Hearts"}
	cardSlice = append(cardSlice, "Six of Spades")

	fmt.Println("Printing the slice using Println")
	fmt.Println(cardSlice)

	// using a type
	cardDeck := deck{"Seven of Spades", newCard()}
	cardDeck = append(cardDeck, "Ace of Spades")

	fmt.Println("Printing the loop")
	// looping
	for i, card := range cardDeck {
		fmt.Println(i, card)
	}

	// convert to byte slice
	helloWorld := "Hello World!"
	fmt.Println([]byte(helloWorld))
}

func newCard() string {
	return "Five of Diamonds"
}
