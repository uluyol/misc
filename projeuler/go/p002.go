package main

import "fmt"

func main() {
	i := 1
	j := 2
	t := 0
	n := 0
	for {
		if (j % 2) == 0 {
			n += j
		}
		t = j
		j += i
		i = t
		if j > 4000000 {
			break
		}
	}
	fmt.Printf("Problem 2: %d\n", n)
}
