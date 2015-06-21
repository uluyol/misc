package main

import (
	"fmt"
	"strconv"
	"strings"
)

type pandigit struct {
	String string
	Int    int64
}

func main() {
	var (
		bigNum     pandigit
		curNum     pandigit
		curNum_a   []string
		curNum_len int
		tmp        string
		i          int64
		j          int64
		err        error
	)
	for i = 2; i < 100000; i++ {
		curNum_len = 0
		curNum_a = nil
		for j = 1; ; j++ {
			tmp = strconv.FormatInt(i*j, 10)
			curNum_a = append(curNum_a, tmp)
			curNum_len += len(tmp)
			if curNum_len > 9 || (curNum_len == 9 && j == 1) {
				break
			} else if curNum_len == 9 {
				curNum.String = strings.Join(curNum_a, "")
				if curNum.Int, err = strconv.ParseInt(curNum.String,
					10,
					64); err != nil {
					return
				}
				if curNum.Int > bigNum.Int && IsPandigital(curNum.Int) {
					bigNum = curNum
				}
				break
			}
		}
	}
	fmt.Printf("Problem 38: %s\n", bigNum.String)
}
