package bitio

import (
	"io"
)

type BitReader interface {
	BitRead([]Bit) (int, error)
	io.Reader
}

type readerAdapter struct {
	rd     io.Reader
	buf    []Bit
	err    error
}

func NewBitReader(rd io.Reader) BitReader {
	return &readerAdapter{rd: rd}
}

func (r *readerAdapter) BitRead(b []Bit) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}

	if len(r.buf) == 0 && r.err != nil {
		return 0, r.err
	}

	var n int
	for n < len(b) {
		if len(r.buf) == 0 {
			if r.err != nil {
				return n, r.err
			}

			r.fetch()
		}

		read := copy(b[n:], r.buf)
		r.buf = r.buf[read:]
		n += read
	}

	return n, r.err
}

func (r *readerAdapter) fetch() {
	var buf [1]byte
	_, r.err = r.rd.Read(buf[:])
	r.buf = Bits(buf[0])
}

func (r *readerAdapter) Read(b []byte) (int, error) {
	bytes := make([]byte, len(b))
	n, err := r.rd.Read(bytes)

	r.buf = append(r.buf, Bits(bytes)...)
	for i := range b {
		b[i] = Convert[byte](r.buf[:8])
		r.buf = r.buf[8:]
	}

	return n, err
}

type ByterReader[T Byter] struct {
	rd BitReader
}

func NewByterReader[T Byter](r io.Reader) *ByterReader[T] {
	if r, ok := r.(BitReader); ok {
		return &ByterReader[T]{ rd: r }
	}

	return &ByterReader[T]{ rd: NewBitReader(r) }
}

func (r *ByterReader[T]) ByterRead() (x T, n int, err error) {
	buf := make([]Bit, Size(x))
	n, err = r.BitRead(buf)
	if err != nil {
		return
	}

	x = Convert[T](buf)
	return
}

func (r *ByterReader[T]) BitRead(b []Bit) (int, error) {
	return r.rd.BitRead(b)
}

func (r *ByterReader[T]) Read(b []byte) (int, error) {
	return r.rd.Read(b)
}

func ByterRead[T Byter](r BitReader) (x T, n int, err error) {
	buf := make([]Bit, Size(x))
	n, err = r.BitRead(buf)
	if err != nil {
		return
	}

	x = Convert[T](buf)
	return
}
