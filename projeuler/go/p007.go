package main

import (
	"fmt"
)

func main() {
	c := 0
	sol := int64(0)
	primes := GetPrimesBelow(10000000)
	for prime, ok := <-primes; ok; prime, ok = <-primes {
		c++
		if c == 10001 {
			sol = prime
			break
		}
	}
	fmt.Printf("Problem 7: %d\n", sol)
}
