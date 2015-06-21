package main

import (
	"fmt"
)

const k = 600851475143

func main() {
	var n int64
	t := GetFactors(k)
	for _, v := range t {
		if v > n {
			n = v
		}
	}
	fmt.Printf("Problem 3: %d\n", n)
}
