package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreetings() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	print(eb)
	print(sb)

}

func (eb englishBot) getGreetings() string {
	return "English bot"
}

func (sb spanishBot) getGreetings() string {
	return "Spanish bot"
}

func print(b bot) {
	fmt.Println(b.getGreetings())
}
