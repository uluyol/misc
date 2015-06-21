package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
)

func main() {
	var (
		file      *os.File
		err       error
		vals      []string
		total     int
		charscore int
	)
	if file, err = os.Open("../common/names.txt"); err != nil {
		return
	}
	csv_reader := csv.NewReader(file)
	if vals, err = csv_reader.Read(); err != nil {
		return
	}
	if err = file.Close(); err != nil {
		return
	}
	sort.Strings(vals)
	for i := 0; i < len(vals); i++ {
		charscore = 0
		for _, v := range vals[i] {
			charscore += int(v) - 64
		}
		total += (charscore * (i + 1))
	}
	fmt.Printf("Problem 22: %d\n", total)
}
