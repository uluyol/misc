package zlist

import (
	"math/rand"
	"testing"
)

func randuint64() uint64 {
	return uint64(randint64())
}

func randint64() int64 {
	return rand.Int63() ^ (rand.Int63() << 1)
}

func testUintSliceCount(t *testing.T, count int, bs int) {
	data := make([]uint64, count)
	s := NewUintSlice(Gzip)
	if bs > 0 {
		s.SetBlocksize(bs)
	}
	for i := 0; i < count; i++ {
		data[i] = randuint64()
		s.Append(data[i])
	}

	for i := 0; i < count; i++ {
		x := s.Get(i)
		if data[i] != x {
			t.Errorf("Put in %d, got %d", data[i], x)
		}
	}
}

func testIntSliceCount(t *testing.T, count int, bs int) {
	data := make([]int64, count)
	s := NewIntSlice(Gzip)
	if bs > 0 {
		s.SetBlocksize(bs)
	}
	for i := 0; i < count; i++ {
		data[i] = randint64()
		s.Append(data[i])
	}

	for i := 0; i < count; i++ {
		x := s.Get(i)
		if data[i] != x {
			t.Errorf("Put in %d, got %d", data[i], x)
		}
	}
}

func TestUintSliceSmall(t *testing.T) {
	testUintSliceCount(t, 200, -1)
}

func TestIntSliceSmall(t *testing.T) {
	testIntSliceCount(t, 200, -1)
}

func TestUintSlice(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping large test")
	}

	testUintSliceCount(t, 1<<19, 1<<13)
}

func TestIntSlice(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping large test")
	}

	testIntSliceCount(t, 1<<19, 1<<13)
}
