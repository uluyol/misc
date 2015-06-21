package main

import (
	"fmt"
	"math/big"
)

const size = 20

func main() {

	final := new(big.Int)
	num := new(big.Int).SetInt64(int64(1))
	den := new(big.Int).SetInt64(int64(1))

	for i := int64(0); i < 2*size; i++ {
		final.SetInt64(2*size - i)
		num.Mul(num, final)
	}
	for i := int64(0); i < size; i++ {
		final.SetInt64((size - i) * (size - i))
		den.Mul(den, final)
	}
	final.Div(num, den)

	fmt.Printf("Problem 15: %s\n", final.String())
}
