package main

import (
	"fmt"
)

func main() {
	var (
		n int64
		f int64
		i int64
		j int64
	)
	for i = 999; i > 99; i-- {
		for j = 999; j > 99; j-- {
			n = i * j
			if n == ReverseInt(n) && n > f {
				f = n
			}
		}
	}
	fmt.Printf("Problem 4: %d\n", f)
}
