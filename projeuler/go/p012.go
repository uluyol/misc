package main

import (
	"fmt"
)

func main() {
	var n int64
	var t []int64
	var prime_count map[int64]int16
	var div_count int16

	for i := int64(1); ; i++ {
		n += i
		t = GetFactors(n)
		prime_count = make(map[int64]int16)
		for _, v := range t {
			prime_count[v]++
		}
		div_count = 1
		for _, v := range prime_count {
			div_count *= (v + 1)
		}
		if div_count >= 500 {
			break
		}
	}
	fmt.Printf("Problem 12: %d\n", n)
}
