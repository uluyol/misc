package main

import (
	"fmt"
)

func main() {
	var (
		sol      int64
		count    int
		digits   []int64
		tr_prime bool
	)
	primes := GetPrimesBelow(10000001)
	for prime, ok := <-primes; ok; prime, ok = <-primes {
		if prime < 10 {
			continue
		}
		digits = IntToDig(prime)
		tr_prime = true
		if IsPrime(int64(digits[0])) &&
			IsPrime(int64(digits[len(digits)-1])) {
			for i := 1; i < len(digits); i++ {
				if !IsPrime(DigToInt(digits[i:])) ||
					!IsPrime(DigToInt(digits[:len(digits)-i])) {
					tr_prime = false
					break
				}
			}
			if tr_prime {
				sol += prime
				count++
			}
			if count == 11 {
				break
			}
		}
	}
	fmt.Printf("Problem 37: %d\n", sol)
}
