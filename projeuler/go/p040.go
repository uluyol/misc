package main

import (
	"fmt"
	"strconv"
)

func main() {
	frac_part := []rune{' '}
	for i := int64(1); ; i++ {
		frac_part = append(frac_part, []rune(strconv.FormatInt(i, 10))...)
		if len(frac_part) >= 1000001 {
			break
		}
	}
	fmt.Printf("Problem 40: %d\n", RuneToInt(frac_part[1])*
		RuneToInt(frac_part[10])*
		RuneToInt(frac_part[100])*
		RuneToInt(frac_part[1000])*
		RuneToInt(frac_part[10000])*
		RuneToInt(frac_part[100000])*
		RuneToInt(frac_part[1000000]))
}
