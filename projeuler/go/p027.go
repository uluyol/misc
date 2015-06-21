package main

import (
	"fmt"
)

type Pair struct {
	a int
	b int
}

func main() {
	cpair := new(Pair)
	pair_c := make(map[Pair]int)
	ab_max := 0
	max_c := 0
	for a := -999; a < 1000; a++ {
		for b := -999; b < 1000; b += 2 {
			cpair.a = a
			cpair.b = b
			for n := 0; ; n++ {
				if !IsPrime(int64(n*n + a*n + b)) {
					break
				}
				pair_c[*cpair]++
			}
		}
	}
	for pair, count := range pair_c {
		if count > max_c {
			max_c = count
			ab_max = pair.a * pair.b
		}
	}
	fmt.Printf("Problem 27: %d\n", ab_max)
}
