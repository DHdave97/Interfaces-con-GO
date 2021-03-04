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

/* func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
} */

/* func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
} */

func (eb englishBot) getGreeting() string {
	// very custom login for generating an english greeting
	return "Hi men"
}
func (sb spanishBot) getGreeting() string {
	// very custom login for generating an english greeting
	return "Hola mijin"
}
