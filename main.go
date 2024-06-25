package main

import (
	"fmt"
	"strconv"

	"github.com/marktran77/go-pocker/deck"
)

func main() {
	card := deck.NewCard(deck.Spades, 1)

	number := 3

	numberString := strconv.Itoa(number)

	fmt.Println(card)
}
