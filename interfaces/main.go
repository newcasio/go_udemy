package main

import "fmt"

//define an interface
//create a new type of bot which is an interface, any function called getGreeting with a return value of string now is also of type bot.
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

//imagine each getGreeting function has very different code
func (eb englishBot) getGreeting() string {
	return "Hi there"
}

//can remove receiver instance (sb) from (sb spanishBot) if we are not using it in the code body
func (spanishBot) getGreeting() string {
	return "Hola"
}

//Print functions for each types
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }
// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

//use interface for printGreeting function
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
