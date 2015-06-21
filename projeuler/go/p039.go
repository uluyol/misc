package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		a        float64
		b        float64
		c        float64
		bigcount int
		per      float64
		pers     map[float64]int
	)
	pers = make(map[float64]int)
	for a = 1; a <= 1000; a++ {
		for b = a; b <= 1000; b++ {
			c = math.Sqrt(a*a + b*b)
			if math.Floor(c) != c {
				continue
			}
			per = a + b + c
			if per > 1000 {
				break
			}
			pers[per]++
		}
	}
	per = 0
	for cper, count := range pers {
		if count > bigcount {
			bigcount = count
			per = cper
		}
	}
	fmt.Printf("Problem 39: %.0f\n", per)
}
