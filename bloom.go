package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"hash"
	"hash/fnv"
)

type Bitstring []byte

func (bs Bitstring) Len() uint {
	return uint(len(bs) * 8)
}

func (bs Bitstring) Set(n uint) {
	pos := n / 8
	bit := n % 8
	bs[pos] |= 1 << bit
}

func (bs Bitstring) Get(n uint) bool {
	pos := n / 8
	bit := n % 8
	return (bs[pos] >> bit) & 1 == 1
}

func (bs Bitstring) String() string {
	buf := bytes.Buffer{}
	for _, b := range bs {
		buf.WriteString(fmt.Sprintf("%08b", b))
	}
	return buf.String()
}

func NewBitstring(siz int) Bitstring {
	return make([]byte, siz)
}

type BloomFilter struct {
	hfuncs [](func() hash.Hash64)
	bits   Bitstring
}

// NewBloomFilter generates a bloom filter with length 8*n.
func NewBloomFilter(n int) *BloomFilter {
	bf := new(BloomFilter)
	bf.bits = NewBitstring(n)
	bf.hfuncs = [](func() hash.Hash64){fnv.New64, fnv.New64a}
	return bf
}

func makeHashes(d []byte, h hash.Hash64) ([]uint32, error) {
	_, err := h.Write(d)
	if err != nil {
		return nil, err
	}
	hashed := h.Sum64()
	lower := uint32(hashed)
	upper := uint32(hashed >> 32)
	return []uint32{lower, upper}, nil
}

func (bf *BloomFilter) makeAllHashes(v interface{}) ([]uint, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	data := buf.Bytes()
	if err != nil {
		return nil, err
	}

	hashes := []uint{}
	for _, hfunc := range bf.hfuncs {
		out, err := makeHashes(data, hfunc())
		if err != nil {
			return nil, err
		}
		for _, h := range out {
			hashes = append(hashes, uint(h))
		}
	}
	return hashes, nil
}

func (bf *BloomFilter) Add(v interface{}) error {
	hashes, err := bf.makeAllHashes(v)
	if err != nil {
		return err
	}

	for _, h := range hashes {
		bf.bits.Set(h % bf.bits.Len())
	}
	return nil
}

func (bf *BloomFilter) Has(v interface{}) (bool, error) {
	hashes, err := bf.makeAllHashes(v)
	if err != nil {
		return false, err
	}

	contains := true
	for _, h := range hashes {
		contains = contains && bf.bits.Get(h % bf.bits.Len())
	}
	return contains, nil
}

func (bf *BloomFilter) String() string {
	return bf.bits.String()
}

func main() {
	bf := NewBloomFilter(10)
	fmt.Println(bf)
	bf.Add("Happy")
	fmt.Println(bf)
	bf.Add("Sad")
	fmt.Println(bf)
	contained, _ := bf.Has("Happy")
	fmt.Printf("Contains \"Happy\": %t\n", contained)
	contained, _ = bf.Has("Sad")
	fmt.Printf("Contains \"Sad\": %t\n", contained)
	contained, _ = bf.Has("Not sad")
	fmt.Printf("Contains \"Not sad\": %t\n", contained)
}
