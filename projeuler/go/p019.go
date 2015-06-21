package main

import (
	"fmt"
	"os"
)

const endyear = 2000

const (
	Jan = iota
	Feb
	Mar
	Apr
	May
	Jun
	Jul
	Aug
	Sep
	Oct
	Nov
	Dec
)

const (
	Sun = iota
	Mon
	Tue
	Wed
	Thu
	Fri
	Sat
)

func main() {
	curday := Sun
	count := 0
	is_leapyear := false
	month := 0

	for year := 1901; year <= endyear; year++ {
		is_leapyear = (year%4) == 0 && (year%100) != 0 ||
			(year%400) == 0
		for month = Jan; month <= Dec; month++ {
			count += btoi(curday == Sun)
			switch month {
			case Jan, Mar, May, Jul, Aug, Oct, Dec:
				curday = (curday + 31) % 7
			case Apr, Jun, Sep, Nov:
				curday = (curday + 30) % 7
			case Feb:
				curday = (curday + 28 + btoi(is_leapyear)) % 7
			default:
				fmt.Fprintf(os.Stderr, "Should be impossible\n")
			}
		}
	}

	fmt.Printf("Problem 19: %d\n", count)
}
