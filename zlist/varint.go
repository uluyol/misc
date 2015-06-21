package zlist

const (
	defaultSize = 128
)

// A Uvarint is a list of uint64 which are stored in memory using a variable-length
// encoding.
type Uvarint struct {
	data []byte
}

func NewUvarint() *Uvarint {
	return &Uvarint{make([]byte, 0, defaultSize)}
}

func NewUvarintCap(cap int) *Uvarint {
	return &Uvarint{make([]byte, 0, cap)}
}

func (v *Uvarint) Append(x uint64) {
	for x >= 0x80 {
		b := byte(x | 0x80)
		v.data = append(v.data, b)
		x >>= 7
	}
	v.data = append(v.data, byte(x))
}

// Iterator returns an iterator for the list. The first return value from the iterator is
// the next number and the second indicates whether or not the end of the list has been
// reached.
func (v Uvarint) Iterator() func() (n uint64, ok bool) {
	pos := 0
	return func() (uint64, bool) {
		if pos >= len(v.data) {
			return 0, false
		}
		x := uint64(0)
		for i, b := range v.data[pos:] {
			x |= uint64(b&0x7f) << uint(i*7)
			if b&0x80 == 0 {
				pos += i + 1
				return x, true
			}
		}
		panic("incomplete array")
	}
}

// A Varint is a list of int64 which are stored in memory using a variable-length
// encoding.
type Varint struct {
	*Uvarint
}

func NewVarint() *Varint {
	return &Varint{NewUvarint()}
}

func NewVarintCap(cap int) *Varint {
	return &Varint{NewUvarintCap(cap)}
}

func (v *Varint) Append(x int64) {
	v.Uvarint.Append(uint64(x))
}

// Iterator returns an iterator for the list. The first return value from the iterator is
// the next number and the second indicates whether or not the end of the list has been
// reached.
func (v Varint) Iterator() func() (n int64, ok bool) {
	next := v.Uvarint.Iterator()
	return func() (int64, bool) {
		u, ok := next()
		return int64(u), ok
	}
}
