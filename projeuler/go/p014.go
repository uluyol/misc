package main

import (
	"fmt"
)

type IterSeq struct {
	start int64
	terms int
}

func main() {

	longest := IterSeq{}
	current := IterSeq{}

	var tmp int64

	for i := 2; i < 1000000; i++ {
		current.start = int64(i)
		tmp = current.start
		current.terms = 1
		for {
			if (tmp % 2) == 0 {
				tmp = tmp / 2
			} else {
				tmp = 3*tmp + 1
			}
			current.terms++
			if tmp == 1 {
				break
			}
		}
		if current.terms > longest.terms {
			longest = current
		}
	}

	fmt.Printf("Problem 14: %d\n", longest.start)
}
