/* Binary Search Tree Implementation. Doesn't store duplicates. */

package main

import "fmt"

type Comparable interface {
	Cmp(b Comparable) int // <0 for less, 0 for equals, >1 for b greater
}

type Num int

func (a Num) Cmp(c Comparable) int {
	b := c.(Num)
	return int(b - a)
}

type Node struct {
	V     Comparable
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func (n *Node) insert(c *Node) {
	cmp := n.V.Cmp(c.V)
	if cmp == 0 {
		return
	}
	if cmp < 0 {
		if n.Left == nil {
			n.Left = c
			return
		}
		n.Left.insert(c)
		return
	}
	if n.Right == nil {
		n.Right = c
		return
	}
	n.Right.insert(c)
}

func (n *Node) search(v Comparable) bool {
	if n == nil {
		return false
	}
	c := n.V.Cmp(v)
	if c == 0 {
		return true
	}
	if c < 0 {
		return n.Left.search(v)
	}
	return n.Right.search(v)
}

func (t *Tree) Add(v Comparable) {
	n := new(Node)
	n.V = v

	if t.Root == nil {
		t.Root = n
		return
	}

	t.Root.insert(n)
}

func (t *Tree) Has(v Comparable) bool {
	return t.Root.search(v)
}

func (n *Node) String() string {
	s := fmt.Sprint(n.V)
	if n.Left != nil {
		s = fmt.Sprintf("%s %s", n.Left.String(), s)
	}
	if n.Right != nil {
		s = fmt.Sprintf("%s %s", s, n.Right.String())
	}
	return s
}

func (t *Tree) String() string {
	return t.Root.String()
}

func main() {
	t := new(Tree)
	t.Add(Num(5))
	t.Add(Num(5))
	t.Add(Num(2))
	t.Add(Num(1))
	t.Add(Num(100))
	fmt.Println("content:", t)
	fmt.Println("has 100:", t.Has(Num(100)))
	fmt.Println("has 2  :", t.Has(Num(2)))
	fmt.Println("has 4  :", t.Has(Num(4)))
	fmt.Println("has 5  :", t.Has(Num(5)))
}
