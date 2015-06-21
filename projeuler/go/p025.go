package main

import (
	"fmt"
	"math/big"
)

func main() {
	i := 0
	n := new(big.Int).SetInt64(int64(1))
	n_1 := new(big.Int).SetInt64(int64(1))
	n_2 := new(big.Int).SetInt64(int64(1))
	for i = 3; ; i++ {
		n.Add(n_1, n_2)
		if len(n.String()) >= 1000 {
			break
		}
		n_2.Set(n_1)
		n_1.Set(n)
	}
	fmt.Printf("Problem 25: %d\n", i)
}
