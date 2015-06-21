package zlist

import (
	"math/rand"
	"testing"
)

func TestUvarint(t *testing.T) {
	in := make([]uint64, 20)
	vu := NewUvarint()
	for i := 0; i < len(in); i++ {
		in[i] = uint64(rand.Int63()) ^ (uint64(rand.Int63()) << 1)
		t.Logf("Appending %d\n", in[i])
		oldlen := len(vu.data)
		vu.Append(in[i])
		if len(vu.data) <= oldlen {
			t.Fatalf("Data not getting longer")
		}
	}
	next := vu.Iterator()
	count := 0
	for x, ok := next(); ok; x, ok = next() {
		if x != in[count] {
			t.Errorf("Mismatch on index %d, put %d, got %d", count, in[count], x)
		}
		count++
	}
	if count != len(in) {
		t.Errorf("Put %d elements in, got %d out", len(in), count)
	}
}

func TestVarint(t *testing.T) {
	in := make([]int64, 20)
	vi := NewVarint()
	for i := 0; i < len(in); i++ {
		in[i] = rand.Int63() ^ (rand.Int63() << 1)
		t.Logf("Appending %d\n", in[i])
		oldlen := len(vi.Uvarint.data)
		vi.Append(in[i])
		if len(vi.Uvarint.data) <= oldlen {
			t.Fatalf("Data not getting longer")
		}
	}
	next := vi.Iterator()
	count := 0
	for x, ok := next(); ok; x, ok = next() {
		if x != in[count] {
			t.Errorf("Mismatch on index %d, put %d, got %d", count, in[count], x)
		}
		count++
	}
	if count != len(in) {
		t.Errorf("Put %d elements in, got %d out", len(in), count)
	}
}
