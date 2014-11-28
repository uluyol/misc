package main

import (
	"container/list"
	"fmt"
	"time"
)

type lruObject struct {
	Key    interface{}
	Value  interface{}
	Expiry time.Time
}

type LRUCache struct {
	cap    int
	items  *list.List
	lookup map[interface{}]*list.Element
	expiry time.Duration
}

type cacheError struct{}

func (ce cacheError) Error() string {
	return "Cache Miss"
}

var miss cacheError

func NewLRUCache(cap int, exp string) (*LRUCache, error) {
	expiry, err := time.ParseDuration(exp)
	if err != nil {
		return nil, err
	}
	c := new(LRUCache)
	c.cap = cap
	c.items = list.New()
	c.lookup = make(map[interface{}]*list.Element, cap)
	c.expiry = expiry
	return c, nil
}

func Must(c *LRUCache, e error) *LRUCache {
	if e != nil {
		panic(e)
	}
	return c
}

func (c *LRUCache) Get(k interface{}) (interface{}, error) {
	e := c.lookup[k]
	if e == nil {
		fmt.Println("MISS")
		return nil, miss
	}
	o := e.Value.(lruObject)
	if o.Expiry.Before(time.Now()) {
		delete(c.lookup, k)
		c.items.Remove(e)
		fmt.Println("MISS")
		return nil, miss
	}
	fmt.Println("HIT")
	c.items.MoveToFront(e)
	return o.Value, nil
}

func (c *LRUCache) Set(k interface{}, v interface{}) {
	if c.cap <= c.items.Len() {
		delete(c.lookup, c.items.Back().Value.(lruObject).Key)
		c.items.Remove(c.items.Back())
	}
	e := c.items.PushFront(lruObject{Key: k, Value: v, Expiry: time.Now().Add(c.expiry)})
	c.lookup[k] = e
}

var fibCache = Must(NewLRUCache(20, "100ms"))

func fib(n int) int {
	c, err := fibCache.Get(n)
	if err == nil {
		return c.(int)
	}
	if n <= 2 {
		return 1
	}
	v := fib(n-1) + fib(n-2)
	fibCache.Set(n, v)
	return v
}

func main() {
	cache, _ := NewLRUCache(10, "5s")
	cache.Set("a", 4)
	v, _ := cache.Get("a")
	fmt.Printf("Got: %d\n", v.(int))
	fmt.Printf("fib(5): %d\n", fib(5))
	fmt.Printf("fib(6): %d\n", fib(6))
	fmt.Printf("fib(100): %d\n", fib(100))
	for i := 0; i < 11; i++ {
		c := 'a' + i
		cache.Set(i, c)
		fmt.Printf("i: %c\n", c)
	}
	for i := 0; i < 11; i++ {
		v, err := cache.Get(i)
		var c rune
		if err == nil {
			c = v.(rune)
		} else {
			c = rune('a' + i)
			cache.Set(i, c)
		}
		fmt.Printf("i: %c\n", c)
	}
	for i := 10; i >= 0; i-- {
		v, err := cache.Get(i)
		var c rune
		if err == nil {
			c = v.(rune)
		} else {
			c = rune('a' + i)
			cache.Set(i, c)
		}
		fmt.Printf("i: %c\n", c)
	}
}
