package main

import (
	"fmt"
	"math/rand"
)

const rounds = 1000

func main() {
	successSwitch := 0
	successStay := 0
	actual := 0
	choice := 0

	for i := 0; i < rounds; i++ {
		actual = rand.Intn(3)
		choice = rand.Intn(3)
		if choice == actual {
			successStay++
			continue
		}
		successSwitch++
	}

	fmt.Printf("Staying:   % 2d%%\n", 100*successStay/rounds)
	fmt.Printf("Switching: % 2d%%\n", 100*successSwitch/rounds)
}
