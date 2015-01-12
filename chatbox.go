package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func nonword(c rune) bool {
	return !(unicode.IsLetter(c) || unicode.IsNumber(c))
}

func makeWords(q string) []string {
	words := strings.FieldsFunc(q, nonword)
	for i, w := range words {
		words[i] = strings.ToLower(w)
	}
	return words
}

type chatter struct {
	responses map[string]map[string]int
	words     map[string]map[string]int
}

func newChatter() *chatter {
	return &chatter{
		responses: make(map[string]map[string]int),
		words:     make(map[string]map[string]int),
	}
}

func (c *chatter) add(q, a string) {
	if c.responses[q] == nil {
		c.responses[q] = make(map[string]int)
	}
	c.responses[q][a]++

	words := makeWords(q)
	for _, w := range words {
		if c.words[w] == nil {
			c.words[w] = make(map[string]int)
		}
		c.words[w][a]++
	}
}

func getMostUsed(useMap map[string]int) (string, int) {
	mostUsed := ""
	highest := -1
	for s, u := range useMap {
		if u > highest {
			mostUsed = s
			highest = u
		}
	}
	return mostUsed, highest
}

func (c *chatter) getBestResponse(q string) (string, bool) {
	if c.responses[q] == nil {
		return "", false
	}
	resp, _ := getMostUsed(c.responses[q])
	c.responses[q][resp]++
	return resp, true
}

func (c *chatter) getBestFromWords(q string) (string, bool) {
	words := makeWords(q)

	var (
		source []string
		resps  []string
		usage  []int
	)

	for _, w := range words {
		if c.words[w] == nil {
			continue
		}
		resp, use := getMostUsed(c.words[w])
		source = append(source, w)
		resps = append(resps, resp)
		usage = append(usage, use)
	}

	if len(resps) == 0 {
		return "", false
	}

	src := ""
	resp := ""
	use := -1

	for i := 0; i < len(resps); i++ {
		if usage[i] > use {
			src = source[i]
			resp = resps[i]
			use = usage[i]
		}
	}

	c.words[src][resp]++

	return resp, true
}

func (c *chatter) getLeastUsed() string {
	var (
		src  string
		resp string
		use  = int(^uint(0) >> 1) // Max integer
	)

	for q, m := range c.responses {
		for a, u := range m {
			if u < use {
				src = q
				resp = a
				use = u
			}
		}
	}

	c.responses[src][resp]++
	return resp
}

func (c *chatter) getBest(q string) string {
	var (
		resp string
		ok   bool
	)

	if resp, ok = c.getBestResponse(q); ok {
		return resp
	}

	if resp, ok = c.getBestFromWords(q); ok {
		return resp
	}

	return c.getLeastUsed()
}

func main() {
	var (
		chat = newChatter()
		in   = bufio.NewReader(os.Stdin)
		cur  string
		prev = "Hello!\n"
		err  error
	)

	if len(os.Args) >= 2 {
		prev = os.Args[2]
	}

	for {
		fmt.Print(prev)
		fmt.Print("> ")
		cur, err = in.ReadString('\n')
		if err != nil {
			fmt.Println("bye bye")
			break
		}

		chat.add(prev, cur)
		prev = cur
		cur = chat.getBest(prev)
		prev = cur
	}
}
