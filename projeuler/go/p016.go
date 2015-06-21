package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := new(big.Int).SetInt64(int64(1))
	two := new(big.Int).SetInt64(int64(2))
	for i := 0; i < 1000; i++ {
		n.Mul(n, two)
	}
	fmt.Printf("Problem 16: %d\n", SumOfBigIntDigits(n))
}
