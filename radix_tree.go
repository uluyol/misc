package main

import (
	"bytes"
	"container/list"
	"fmt"
)

type node struct {
	children []*node
	data     interface{}
	label    string
}

func newNode(s string) *node {
	return &node{
		children: make([]*node, 0, 2),
		data:     nil,
		label:    s,
	}
}

func (n *node) isLeaf() bool {
	return n.data != nil
}

func (n *node) get(s string) (interface{}, bool) {
	if s == "" && n.isLeaf() {
		return n.data, true
	}
	for _, c := range n.children {
		end := len(c.label)
		if end > len(s) || s[:end] != c.label {
			continue
		}
		if end < len(s) && c.isLeaf() {
			continue
		}
		return c.get(s[end:])
	}
	return nil, false
}

func (n *node) search(s string) (*node, string, bool) {
	if s == "" && n.isLeaf() {
		return n, s, true
	}
	for _, c := range n.children {
		end := len(c.label)
		if end > len(s) || s[:end] != c.label {
			continue
		}
		if c.isLeaf() {
			return n, s, false
		}
		return c.search(s[end:])
	}
	return n, s, false
}

func (n *node) getParent(s string) (*node, int) {
	for i, c := range n.children {
		if len(c.label) > len(s) || s[:len(c.label)] != c.label {
			continue
		}
		if len(c.label) == len(s) && c.isLeaf() {
			return n, i
		}
		return c.getParent(s[len(c.label):])
	}
	return n, -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sharedPrefixLength(a, b string) int {
	i := 0
	for ; i < min(len(a), len(b)); i++ {
		if a[i] != b[i] {
			break
		}
	}
	return i
}

func (n *node) setParent(k string, v interface{}) {
	for i, c := range n.children {
		shared := sharedPrefixLength(k, c.label)
		if shared > 0 {
			splitParent := newNode(k[:shared])
			c.label = c.label[shared:]
			newChild := newNode(k[shared:])
			newChild.data = v
			splitParent.children = append(splitParent.children, c)
			splitParent.children = append(splitParent.children, newChild)
			n.children[i] = splitParent
			return
		}
	}
	newChild := newNode(k)
	newChild.data = v
	n.children = append(n.children, newChild)
}

func (n *node) set(k string, v interface{}) {
	top, newkey, found := n.search(k)
	if found {
		top.data = v
	} else {
		top.setParent(newkey, v)
	}
}

type Set struct {
	root *node
}

func NewSet() *Set {
	return &Set{newNode("")}
}

func (s Set) Has(k string) bool {
	_, ok := s.Get(k)
	return ok
}

func (s Set) Get(k string) (interface{}, bool) {
	return s.root.get(k)
}

func (s Set) Set(k string, v interface{}) {
	s.root.set(k, v)
}

func extend(a, b []*node) []*node {
	c := a
	for _, v := range b {
		c = append(c, v)
	}
	return c
}

func (s Set) Delete(k string) {
	p, i := s.root.getParent(k)
	if i < 0 {
		return // No such key
	}
	p.children[i] = p.children[len(p.children)-1]
	p.children = p.children[:len(p.children)-1]
	if len(p.children) == 1 {
		p.label += p.children[0].label
		p.data = p.children[0].data
		p.children = p.children[:0]
	}
}

type queue struct {
	l *list.List
}

func newQueue() *queue {
	return &queue{list.New()}
}

func (q queue) push(v interface{}) {
	q.l.PushBack(v)
}

func (q queue) pop() interface{} {
	e := q.l.Front()
	v := e.Value
	q.l.Remove(e)
	return v
}

func (q queue) more() bool {
	return q.l.Len() > 0
}

func (s Set) String() string {
	var buf bytes.Buffer
	q := newQueue()
	q.push(s.root)
	q.push(nil)
	for q.more() {
		c := q.pop()
		if c == nil {
			if q.more() {
				q.push(nil)
				buf.WriteString("\n")
				continue
			}
			break
		}
		n := c.(*node)
		buf.WriteString(fmt.Sprintf("%s(%t, %v)", n.label, n.isLeaf(), n.data))
		buf.WriteString("-")
		for _, child := range n.children {
			q.push(child)
		}
	}
	return buf.String()
}

func main() {
	pp := func(v interface{}, ok bool) string { return fmt.Sprintf("%v %t", v, ok) }
	s := NewSet()
	p := func(k string) { fmt.Printf("%s: %s\n", k, pp(s.Get(k))) }

	s.Set("Hello", "World")
	s.Set("Hell", "Fire")
	s.Set("Yellow", "Purple")
	s.Set("Yellower", "Green")
	p("Hello")
	p("Hell")
	p("Yellow")
	s.Delete("Yellow")
	p("Yellower")
	fmt.Println(s)
}
