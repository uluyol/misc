package zlist

import (
	"bytes"
	"compress/gzip"
	"compress/lzw"
	"errors"
	"io"
)

const (
	blocksize  = 1 << 27 // So that blocksize in bytes becomes 128 MB
	uint64size = 8
)

type ReaderGen func(r io.Reader) io.ReadCloser
type WriterGen func(w io.Writer) io.WriteCloser

type Compressor int

const (
	Gzip Compressor = iota
	Lzw
)

func newGzipReader(r io.Reader) io.ReadCloser {
	gr, err := gzip.NewReader(r)
	if err != nil {
		panic(err)
	}
	return gr
}

var (
	ErrUsed      = errors.New("the slice has already been used")
	ErrUnaligned = errors.New("the size must be a multiple of the data size")

	readerGen = map[Compressor]ReaderGen{
		Gzip: newGzipReader,
		Lzw:  func(r io.Reader) io.ReadCloser { return lzw.NewReader(r, lzw.LSB, 8) },
	}
	writerGen = map[Compressor]WriterGen{
		Gzip: func(w io.Writer) io.WriteCloser { return gzip.NewWriter(w) },
		Lzw:  func(w io.Writer) io.WriteCloser { return lzw.NewWriter(w, lzw.LSB, 8) },
	}
)

type zblock []byte

// A UintSlice stores large slices of uint64 that are compressed in large blocks.
type UintSlice struct {
	blocks []zblock
	cur    struct {
		slice    []byte
		pos      int
		modified bool
	}
	last      []byte
	blocksize int
	newReader ReaderGen
	newWriter WriterGen
}

func NewUintSlice(c Compressor) *UintSlice {
	return NewUintSliceCustom(readerGen[c], writerGen[c])
}

func NewUintSliceCustom(rg ReaderGen, wg WriterGen) *UintSlice {
	return &UintSlice{
		blocksize: blocksize,
		newReader: rg,
		newWriter: wg,
	}
}

func (s *UintSlice) SetBlocksize(bs int) error {
	if len(s.blocks) > 0 || s.last != nil {
		return ErrUsed
	}
	if bs%uint64size != 0 {
		return ErrUnaligned
	}
	s.blocksize = bs
	return nil
}

func bytesof(n uint64) [uint64size]byte {
	var raw [uint64size]byte
	for i := 0; i < uint64size; i++ {
		raw[i] = byte(n)
		n >>= 8
	}
	return raw
}

func uint64of(raw []byte) uint64 {
	var n uint64
	for i := 0; i < uint64size; i++ {
		n |= uint64(raw[i]) << uint(i*8)
	}
	return n
}

func (s *UintSlice) makeBlock() []byte {
	return make([]byte, 0, s.blocksize)
}

// Opt to panic instead of return errors since it's highly unlikely that
// an error will occur. bytes.Buffer will not error, the compression
// algorithm probably won't either.
func (s *UintSlice) compressBlock(block []byte) zblock {
	var buf bytes.Buffer
	w := s.newWriter(&buf)
	_, err := w.Write(block)
	w.Close()
	if err != nil {
		panic(err)
	}
	return (&buf).Bytes()
}

func (s *UintSlice) decompressBlock(dest []byte, block zblock) []byte {
	if dest == nil {
		dest = make([]byte, s.blocksize)
	}
	buf := bytes.NewReader(block)
	r := s.newReader(buf)
	_, err := r.Read(dest)
	if err != nil && err != io.EOF {
		panic(err)
	}
	return dest
}

func (s *UintSlice) Append(n uint64) {
	if s.last == nil {
		s.last = s.makeBlock()
	}

	for _, b := range bytesof(n) {
		s.last = append(s.last, b)
	}

	if len(s.last) == cap(s.last) {
		s.blocks = append(s.blocks, s.compressBlock(s.last))
		s.saveCur()
		s.cur.slice, s.last = s.last, s.cur.slice[:0]
		s.cur.pos = len(s.blocks)
		s.cur.modified = false
		s.last = nil
	}
}

func (s *UintSlice) Len() int {
	return (len(s.blocks)*s.blocksize + len(s.last)) / uint64size
}

func (s *UintSlice) decomposePos(i int) (int, int) {
	blocki := (i * uint64size) / s.blocksize
	bytei := i*uint64size - blocki*s.blocksize
	return blocki, bytei
}

func (s *UintSlice) saveCur() {
	if s.cur.modified {
		s.blocks[s.cur.pos] = s.compressBlock(s.cur.slice)
	}
}

func (s *UintSlice) loadCur(pos int) {
	if s.cur.pos == pos {
		return
	}
	s.saveCur()
	s.cur.slice = s.decompressBlock(s.cur.slice, s.blocks[pos])
	s.cur.pos = pos
	s.cur.modified = false
}

func (s *UintSlice) Get(i int) uint64 {
	blocki, bytei := s.decomposePos(i)
	if blocki >= len(s.blocks) {
		return uint64of(s.last[bytei:])
	}
	s.loadCur(blocki)
	return uint64of(s.cur.slice[bytei:])
}

func (s *UintSlice) Set(i int, n uint64) {
	blocki, bytei := s.decomposePos(i)
	if blocki >= len(s.blocks) {
		for j, b := range bytesof(n) {
			s.last[bytei+j] = b
		}
	}
	s.loadCur(blocki)
	for j, b := range bytesof(n) {
		s.cur.slice[bytei+j] = b
	}
	s.cur.modified = true
}

// A IntSlice stores large slices of int64 that are compressed in large blocks.
type IntSlice struct {
	*UintSlice
}

func NewIntSlice(c Compressor) *IntSlice {
	return &IntSlice{NewUintSlice(c)}
}

func NewIntSliceCustom(rg ReaderGen, wg WriterGen) *IntSlice {
	return &IntSlice{NewUintSliceCustom(rg, wg)}
}

func (s *IntSlice) Append(n int64) {
	s.UintSlice.Append(uint64(n))
}

func (s *IntSlice) Get(i int) int64 {
	return int64(s.UintSlice.Get(i))
}

func (s *IntSlice) Set(i int, n int64) {
	s.UintSlice.Set(i, uint64(n))
}
