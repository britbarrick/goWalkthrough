package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// custom logic for generating greeting in English
	return "Hello there"
}

func (spanishBot) getGreeting() string {
	// custom logic for generating greeting in Spanish
	return "Hola"
}
