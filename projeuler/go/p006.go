package main

import (
	"fmt"
)

func main() {
	n := 0
	sq_sum, sum_sq := 1, 1
	for i := 2; i <= 100; i++ {
		sum_sq += i * i
		sq_sum += i
	}
	sq_sum *= sq_sum
	n = sq_sum - sum_sq
	fmt.Printf("Problem 6: %d\n", n)
}
