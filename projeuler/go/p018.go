package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var old []int
	var new []int
	big := -1

	f, _ := os.Open("../common/p018.txt")
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	old = lineToInts(scanner.Text())

	for scanner.Scan() {
		new = lineToInts(scanner.Text())
		longestPaths(old, new)
		old = new
	}

	for _, v := range old {
		if v > big {
			big = v
		}
	}

	fmt.Printf("Problem 18: %d\n", big)
}
