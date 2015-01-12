/* Binary Search Tree Implementation. Doesn't store duplicates. */

package main

import "fmt"

// Comparable objects can be compared.
type Comparable interface {
	Cmp(b Comparable) int // <0 for less, 0 for equals, >1 for b greater
}

type num int

func (a num) Cmp(c Comparable) int {
	b := c.(num)
	return int(b - a)
}

type node struct {
	V     Comparable
	Left  *node
	Right *node
}

// Tree represents a binary search tree that can be used as a set.
type Tree struct {
	root *node
}

func (n *node) insert(c *node) {
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

func (n *node) search(v Comparable) bool {
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

// Add adds v to the set.
func (t *Tree) Add(v Comparable) {
	n := new(node)
	n.V = v

	if t.root == nil {
		t.root = n
		return
	}

	t.root.insert(n)
}

// Has checks to see if v is in the set.
func (t *Tree) Has(v Comparable) bool {
	return t.root.search(v)
}

func recursiveRemove(n **node, v Comparable) bool {
	if *n == nil {
		return false
	}
	cmp := (*n).V.Cmp(v)
	if cmp < 0 {
		return recursiveRemove(&(*n).Left, v)
	} else if cmp > 0 {
		return recursiveRemove(&(*n).Right, v)
	}
	// Found
	if (*n).Left == nil && (*n).Right == nil {
		*n = nil
	} else if (*n).Left == nil {
		*n = (*n).Right
	} else if (*n).Right == nil {
		*n = (*n).Left
	} else {
		(*n).V = (*n).Right.V
		recursiveRemove(&(*n).Right, (*n).V)
	}
	return true
}

// Remove will remove v from the set.
func (t *Tree) Remove(v Comparable) bool {
	return recursiveRemove(&t.root, v)
}

func (n *node) String() string {
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
	return t.root.String()
}

func main() {
	t := new(Tree)
	t.Add(num(5))
	t.Add(num(5))
	t.Add(num(2))
	t.Add(num(1))
	t.Add(num(100))
	fmt.Println("content:", t)
	fmt.Println("has 100:", t.Has(num(100)))
	fmt.Println("has 2  :", t.Has(num(2)))
	fmt.Println("has 4  :", t.Has(num(4)))
	fmt.Println("has 5  :", t.Has(num(5)))
	t.Remove(num(5))
	fmt.Println("has 5  :", t.Has(num(5)))
	fmt.Println("content:", t)
}
