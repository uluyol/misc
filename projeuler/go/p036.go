package main

import (
	"fmt"
	"strconv"
)

func main() {
	var f int64
	for i := int64(1); i < 1000000; i++ {
		i_bin := strconv.FormatInt(i, 2)
		if i == ReverseInt(i) && i_bin == ReverseString(i_bin) {
			f += i
		}
	}
	fmt.Printf("Problem 36: %d\n", f)
}
