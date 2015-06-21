package main

import "fmt"

func main() {
	i := 0
	n := 0
	for i = 3; i < 1000; i += 3 {
		n += i
	}
	for i = 5; i < 1000; i += 5 {
		if (i % 3) != 0 {
			n += i
		}
	}
	fmt.Printf("Problem 1: %d\n", n)
}
