package main

import (
	"fmt"
)

func main() {
	n := 0
	for i := 20; ; i += 20 {
		if IsDivByAll(i, 19) {
			n = i
			break
		}
	}
	fmt.Printf("Problem 5: %d\n", n)
}
