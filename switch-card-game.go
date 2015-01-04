package main

import (
	"fmt"
	"math/rand"
)

const ROUNDS = 1000

func main() {
	success_switch := 0
	success_stay := 0
	actual := 0
	choice := 0

	for i := 0; i < ROUNDS; i++ {
		actual = rand.Intn(3)
		choice = rand.Intn(3)
		if choice == actual {
			success_stay++
			continue
		}
		success_switch++
	}

	fmt.Printf("Staying:   % 2d%%\n", 100*success_stay/ROUNDS)
	fmt.Printf("Switching: % 2d%%\n", 100*success_switch/ROUNDS)
}
