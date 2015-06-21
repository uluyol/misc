package main

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

func btoi(b bool) int {
	if b {
		return 1;
	}
	return 0;
}

// Only works for numbers <=9 digits
func IsPandigital(n int64) bool {
	occur_c := make(map[int]int)
	digits := IntToDig(n)
	for i := 0; i < len(digits); i++ {
		occur_c[int(digits[i])]++
	}
	for i := 1; i <= len(digits); i++ {
		if occur_c[i] != 1 {
			return false
		}
	}
	return true
}

func GetPrimesBelow(max int64) <-chan int64 {
	primes := make(chan int64)
	go func() {
		primes <- int64(2)
		sieve := make([]bool, max)
		for i := int64(3); i < max; i += 2 {
			if !sieve[i] {
				primes <- i
				for c := 3 * i; c < max; c += i << 1 {
					sieve[c] = true
				}
			}
		}
		close(primes)
	}()
	return primes
}

func IsPrime(n int64) bool {
	if n <= 1 {
		return false
	}
	for i := int64(2); i < int64(math.Sqrt(float64(n)))+1; i++ {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}

func RuneToInt(r rune) int {
	return int(r) - 48
}

func SumOfBigIntDigits(n *big.Int) int64 {
	sum := int64(0)
	for _, v := range n.String() {
		sum += int64(RuneToInt(v))
	}
	return sum
}

/* Gets digits of an integer */
func IntToDig(n int64) []int64 {
	var (
		f            []int64
		t            int64
		exp          int64
		last         int
		search_digit bool
	)
	search_digit = false
	for e := 19; e >= 0; e-- {
		exp = int64(math.Pow10(e))
		t = n / exp
		if search_digit {
			f[last-e] = t
			n = n % exp
		} else if t != 0 {
			f = make([]int64, e+1)
			last = e
			f[0] = t
			n = n % exp
			search_digit = true
		}
	}
	return f
}

func DigToInt(digits []int64) int64 {
	var out int64
	for i := len(digits); i > 0; i-- {
		out += digits[len(digits)-i] * int64(math.Pow10(i-1))
	}
	return out
}

func ReverseInt(n int64) int64 {
	digits := IntToDig(n)
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return DigToInt(digits)
}

func ReverseString(s string) string {
	reversed := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	return string(reversed)
}

func GetFactors(n int64) []int64 {
	var f []int64
	var i int64
	for i = 2; i < int64(math.Sqrt(float64(n))); i++ {
		if (n % i) == 0 {
			f = append(f, i)
			f = append(f, GetFactors(n/i)...)
			return f
		}
	}
	return []int64{n}
}

func IsDivByAll(k int, d int) bool {
	for i := d; i > 1; i-- {
		if (k % i) != 0 {
			return false
		}
	}
	return true
}

func HightlightMatrix(mx [20][20]uint8, i int, j int, pi bool, pj, ii bool, ij bool) {
	var mxc [20][20]string

	for i := 0; i < len(mx); i++ {
		for j := 0; j < len(mx); j++ {
			mxc[i][j] = "\x1b[0m"
		}
	}

	for c := 0; c < 4; c++ {
		if ii && ij {
			if pi && pj {
				mxc[i+c][j+c] = "\x1b[1;92m"
			} else if pi {
				mxc[i+c][j-c] = "\x1b[1;92m"
			} else if pj {
				mxc[i-c][j+c] = "\x1b[1;92m"
			} else {
				mxc[i-c][j-c] = "\x1b[1;92m"
			}
		} else if ii {
			if pi {
				mxc[i+c][j] = "\x1b[1;92m"
			} else {
				mxc[i-c][j] = "\x1b[1;92m"
			}
		} else if ij {
			if pj {
				mxc[i][j+c] = "\x1b[1;92m"
			} else {
				mxc[i][j-c] = "\x1b[1;92m"
			}
		}
	}

	for i := 0; i < len(mx); i++ {
		for j := 0; j < len(mx); j++ {
			if mx[i][j] < 10 {
				fmt.Printf("%s0%d ", mxc[i][j], mx[i][j])
			} else {
				fmt.Printf("%s%d ", mxc[i][j], mx[i][j])
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

/* Gets digits of an unsigned integer */
func Uitod64(n uint64) []uint8 {
	var f []uint8
	var t uint8
	var exp uint64
	search_digit := false
	for e := 19; e >= 0; e-- {
		exp = uint64(math.Pow10(e))
		t = uint8(n / exp)
		if search_digit || t != 0 {
			f = append(f, t)
			n = n % exp
			search_digit = true
		}
	}
	return f
}

/* Used by Itobstr */
func dtobstr(i uint8) string {
	eng := [10]string{"", "one", "two", "three", "four", "five", "six",
		"seven", "eight", "nine"}
	if i < 10 {
		return eng[i]
	}
	return ""
}

/* Converts an integer < 9999 to a string in written form.
 * Example: 121 one hundred and twenty-one
 * Argument: set true to keep spaces/hypens, false otherwise.
 */
func Itobstr(i int, ksep bool) (string, error) {
	teens := [10]string{"ten", "eleven", "twelve", "thirteen", "fourteen",
		"fifteen", "sixteen", "seventeen", "eighteen",
		"nineteen"}
	tens := [10]string{"", "ten", "twenty", "thirty", "forty", "fifty",
		"sixty", "seventy", "eighty", "ninety"}
	digit := Uitod64(uint64(i))
	fs, t := "", ""
	and := ""
	space := ""
	puberty := false

	if ksep {
		and = " and "
		space = " "
	} else {
		and = "and"
	}

	for {
		switch {
		case len(digit) > 4:
			return "", errors.New("Itobstr: Integers > 9999 not handled yet")
		case len(digit) == 4, len(digit) == 3:
			t = dtobstr(digit[0])
			if t != "" {
				if fs != "" {
					fs += and
				}
				fs += t
				fs += space
				if len(digit) == 4 {
					fs += "thousand"
				} else {
					fs += "hundred"
				}
			}
			digit = digit[1:len(digit)]
		/* English is an AWEFUL language */
		case len(digit) == 2:
			switch {
			case digit[0] == 1:
				/* English is a TERRIBLE language */
				t = teens[digit[1]]
				puberty = true
			case digit[0] < 10:
				t = tens[digit[0]]
			default:
				t = ""
			}
			if t != "" {
				if fs != "" {
					fs += and
				}
				fs += t
				if puberty {
					return fs, nil
				}
				if digit[1] != 0 {
					fs += "-"
				}
			}
			digit = digit[1:len(digit)]
		case len(digit) == 1:
			t = dtobstr(digit[0])
			if t != "" {
				if fs != "" {
					if fs[len(fs)-1] != '-' {
						fs += and
					} else if fs[len(fs)-1] == '-' && !ksep {
						fs = fs[0 : len(fs)-1]
					}
				}
				fs += t
			}
			return fs, nil
		}
	}
	return "", errors.New("Itobstr: something broke")
}

func BigFibonacci(term int) *big.Int {
	n := new(big.Int).SetInt64(int64(1))
	n_1 := new(big.Int).SetInt64(int64(1))
	n_2 := new(big.Int).SetInt64(int64(1))
	for i := 3; i <= term; i++ {
		n.Add(n_1, n_2)
		n_2.Set(n_1)
		n_1.Set(n)
	}
	return n_1
}

func AppendBigIntIfMissing(slice []*big.Int, i *big.Int) []*big.Int {
	for _, v := range slice {
		if i.Cmp(v) == 0 {
			return slice
		}
	}
	return append(slice, new(big.Int).Set(i))
}

func longestPaths(old []int, new []int) {
	new[0] += old[0]
	new[len(new)-1] += old[len(old)-1]

	for i := 1; i < len(new)-1; i++ {
		if old[i-1] > old[i] {
			new[i] += old[i-1]
		} else {
			new[i] += old[i]
		}
	}
}

func lineToInts(line string) []int {
	vals := strings.Fields(line)
	ints := make([]int, len(vals))
	for i, v := range vals {
		ints[i], _ = strconv.Atoi(v)
	}
	return ints
}
