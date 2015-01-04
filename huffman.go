package main

import (
	"container/heap"
	"fmt"
)

type pair struct {
	sym  interface{}
	freq float32
}

type keyPair struct {
	sym  interface{}
	code string
}

type pairHeap []pair

func (h pairHeap) Len() int            { return len(h) }
func (h pairHeap) Less(i, j int) bool  { return h[i].freq < h[j].freq }
func (h pairHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *pairHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *pairHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}

func combine(a, b pair) pair {
	return pair{[2]pair{a, b}, a.freq + b.freq}
}

type CodeBuilder struct {
	pairs pairHeap
}

type Code struct {
	encode map[interface{}]string
	decode map[string]interface{}
}

func newCode() *Code {
	c := new(Code)
	c.encode = make(map[interface{}]string)
	c.decode = make(map[string]interface{})
	return c
}

func (c *Code) Encode(i interface{}) string { return c.encode[i] }
func (c *Code) Decode(k string) interface{} { return c.decode[k] }
func (c *Code) String() string {
	s := ""
	for k, v := range c.encode {
		s = fmt.Sprintf("%s%v: %s\n", s, k, v)
	}
	return s[:len(s)-1]
}

func (c *CodeBuilder) AddSymbol(sym interface{}, prob float32) {
	c.pairs = append(c.pairs, pair{sym, prob})
}

func (c *CodeBuilder) Build() *Code {
	heap.Init(&c.pairs)
	num := len(c.pairs)
	for len(c.pairs) > 1 {
		a := heap.Pop(&c.pairs)
		b := heap.Pop(&c.pairs)
		heap.Push(&c.pairs, combine(a.(pair), b.(pair)))
	}

	ch := make(chan keyPair)
	pre := ""
	if num == 1 {
		pre = "0"
	}
	go recursiveBuild(c.pairs[0], pre, ch)
	code := newCode()
	for i := 0; i < num; i++ {
		kp := <-ch
		code.encode[kp.sym] = kp.code
		code.decode[kp.code] = kp.sym
	}
	return code
}

func recursiveBuild(item pair, pre string, c chan keyPair) {
	if p2, ok := item.sym.([2]pair); ok {
		recursiveBuild(p2[0], pre+"0", c)
		recursiveBuild(p2[1], pre+"1", c)
		return
	}
	c <- keyPair{item.sym, pre}
}

func main() {
	syms := []string{"a", "b", "c", "d", "e"}
	probs := []float32{0.25, 0.25, 0.3, 0.1, 0.1}

	cb := new(CodeBuilder)
	for i := 0; i < len(syms); i++ {
		cb.AddSymbol(syms[i], probs[i])
	}

	code := cb.Build()
	fmt.Println(code)

	fmt.Println()

	cb2 := new(CodeBuilder)
	cb2.AddSymbol("x", 1)
	fmt.Println(cb2.Build())
}
