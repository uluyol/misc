package main

import (
	"fmt"
)

func main() {

	s := ""
	t := ""

	for i := 1; i <= 1000; i++ {
		t, _ = Itobstr(i, false)
		s += t
	}

	fmt.Printf("Problem 17: %d\n", len(s))
}
