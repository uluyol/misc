package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := new(big.Int).SetInt64(int64(1))
	t := new(big.Int)
	for i := int64(1); i <= 100; i++ {
		t.SetInt64(i)
		n.Mul(n, t)
	}
	fmt.Printf("Problem 20: %d\n", SumOfBigIntDigits(n))
}
