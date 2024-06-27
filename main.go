package main

import (
	"fmt"

	"github.com/marktran77/go-pocker/deck"
)

func main() {

	for i := 0; i < 10; i++ {
		d := deck.New()
		fmt.Println(d)
		fmt.Println()
	}

}
