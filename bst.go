package main

import "fmt"

type Comparable interface {
	Compare(b Comparable) int // <0 for less, 0 for equals, >1 for b greater
}

type Num int

func (a Num) Compare(c Comparable) int {
	b := c.(Num)
	return int(b - a)
}

type Node struct {
	V     Comparable
	Left  *Node
	Right *Node
}

type Tree struct {
	Root   *Node
	Height int
}

func (n *Node) insert(c *Node, depth int) int {
	if n.V.Compare(c.V) <= 0 {
		if n.Left == nil {
			n.Left = c
			return depth
		}
		return n.Left.insert(c, depth+1)
	}
	if n.Right == nil {
		n.Right = c
		return depth
	}
	return n.Right.insert(c, depth+1)
}

func (n *Node) search(v Comparable) bool {
	if n == nil {
		return false
	}
	c := n.V.Compare(v)
	if c == 0 {
		return true
	}
	if c < 0 {
		return n.Left.search(v)
	}
	return n.Right.search(v)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (t *Tree) Insert(v Comparable) {
	n := new(Node)
	n.V = v

	if t.Root == nil {
		t.Root = n
		t.Height = 1
		return
	}

	t.Height = max(t.Height, t.Root.insert(n, 1))
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
	t.Insert(Num(5))
	t.Insert(Num(2))
	t.Insert(Num(100))
	fmt.Println("content:", t)
	fmt.Println("height :", t.Height)
	fmt.Println("has 100:", t.Has(Num(100)))
	fmt.Println("has 2  :", t.Has(Num(2)))
	fmt.Println("has 4  :", t.Has(Num(4)))
	fmt.Println("has 5  :", t.Has(Num(5)))
}
