/*

This is an attempt to implement manual memory management in Go. Please DO NOT
use this. This was just for fun and would likely be worse than Go's garbage
collection.

*/

package main

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

var pool = newmemoryPool()

type bufferNode struct {
	start []byte
	next  *bufferNode
}

type memNode struct {
	start uintptr
	len   int
	next  *memNode
}

type memoryPool struct {
	sync.Mutex
	cur    *memNode
	newCap int

	// Needed so that data isn't garbage collected
	bufs *bufferNode
}

func newmemoryPool() *memoryPool {
	pool := new(memoryPool)
	pool.newCap = 4096
	return pool
}

func newbufferNode(siz int) *bufferNode {
	return &bufferNode{make([]byte, siz), nil}
}

func (bn *bufferNode) memNode() *memNode {
	start := (*reflect.SliceHeader)(unsafe.Pointer(&bn.start)).Data
	return &memNode{start, cap(bn.start), nil}
}

func (p *memoryPool) grow() {
	// Ensure the lock is held
	n := newbufferNode(p.newCap)
	n.next = p.bufs
	p.bufs = n
	p.newCap <<= 1

	m := n.memNode()
	m.next = p.cur
	p.cur = m
}

func (p *memoryPool) String() string {
	s := ""
	for cur := p.cur; cur != nil; cur = cur.next {
		s += fmt.Sprint(*cur, "")
	}
	return s
}

const sizeInt = 8

func Malloc(siz int) unsafe.Pointer {
	if siz <= 0 {
		return nil
	}

	var addr unsafe.Pointer
	pool.Lock()

	ssiz := siz + sizeInt
	fmt.Println("State:", pool)

	searching := true
	for searching {
		fmt.Println("At:", pool.cur)
		var prev *memNode = nil
		for cur := pool.cur; cur != nil; prev, cur = cur, cur.next {
			if cur.len < ssiz {
				continue
			}

			searching = false
			if prev == nil {
				pool.cur = cur.next
			} else {
				prev.next = cur.next
			}

			fmt.Println("found:", cur.start)

			h := (*int)(unsafe.Pointer(cur.start))
			*h = ssiz
			addr = unsafe.Pointer(cur.start + sizeInt)

			if cur.len-ssiz > 0 {
				n := &memNode{cur.start + uintptr(ssiz), cur.len - ssiz, pool.cur}
				pool.cur = n
			} else if cur.len-ssiz < 0 {
				panic("somehow are using small block")
			}
			break
		}
		if searching {
			pool.grow()
		}
	}

	pool.Unlock()
	fmt.Println("State:", pool)
	fmt.Println("Returned:", uintptr(addr))
	return addr
}

// Not very smart
func Free(addr unsafe.Pointer) {
	if addr == nil {
		return
	}

	hptr := uintptr(addr) - sizeInt
	h := (*int)(unsafe.Pointer(hptr))
	n := &memNode{hptr, *h, nil}

	pool.Lock()
	n.next = pool.cur
	pool.cur = n
	pool.Unlock()
}

func main() {
	ip := (*int)(Malloc(sizeInt))
	*ip = 3
	fmt.Println(*ip)
	Free(unsafe.Pointer(ip))
	jp := (*int)(Malloc(sizeInt))
	*jp = 4
	fmt.Println(*ip, "?=", *jp) // Should be equal to show that we are free-ing
}
