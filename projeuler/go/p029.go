package main

import (
	"fmt"
	"math/big"
)

const (
	amin = int64(2)
	amax = int64(100)
	bmin = int64(2)
	bmax = int64(100)
)

func main() {
	var terms []*big.Int
	I := new(big.Int)
	J := new(big.Int)
	for i := amin; i <= amax; i++ {
		I.SetInt64(i)
		for j := bmin; j <= bmax; j++ {
			J.SetInt64(j)
			terms = AppendBigIntIfMissing(terms, J.Exp(I, J, nil))
		}
	}
	fmt.Printf("Problem 29: %d\n", len(terms))
}
