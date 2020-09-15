package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sp := spanishBot{}

	printGreeting(eb)
	printGreeting(sp)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "hi There"
}

// alternative way
// func (englishBot) getGreeting() string {
// 	return "hi There"
// }

func (sp spanishBot) getGreeting() string {
	return "hola"
}
