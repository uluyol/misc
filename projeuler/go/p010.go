package main

import (
	"fmt"
)

const k = 2000000

func main() {

	f := int64(0)

	primes := GetPrimesBelow(k)
	for prime, ok := <-primes; ok; prime, ok = <-primes {
		f += prime
	}
	fmt.Printf("Problem 10: %d\n", f)
}
