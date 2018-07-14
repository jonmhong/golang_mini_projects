package main

import "fmt"

// another example
// type bot interface {
//	// to qualify for type bot, the type has to have all these functions and params
// 	getGreeting(string, int) (string, error)
// 	getBotVersion() float64
// 	respondToUser(user) string
// }

// we have a new custom type called "bot"
type bot interface {
	// if any type has a function called "getGreeting" then they're
	// are automatic members of the "bot" interface
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
	// very custom logic for generating english greeting
	return "Hi there!"
}

// can remove sb since we're not using sb in the function
func (spanishBot) getGreeting() string {
	return "Hola!"
}
