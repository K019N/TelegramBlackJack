package main

import "fmt"

func main() {
	makeCards()
	var card = randIntIn(2, 11)
	var exists, name = existsCard(card)
	if exists {
		fmt.Println(name, card)
	}
}
