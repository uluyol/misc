package main

import (
	"fmt"
)

const k = 1000

func main() {
	var n, a, b, c uint32

	for a = 1; a < k-1; a++ {
		for b = 1; b < k-a; b++ {
			c = 1000 - b - a
			if (a*a + b*b) == (c * c) {
				goto endl
			}
		}
	}
endl:
	n = a * b * c

	fmt.Printf("Problem 9: %d\n", n)
}
