package bitio

import "io"

type BitReader interface {
	ReadBit() (Bit, error)
	ReadBits([]Bit) (int, error)
	Read([]byte) (int, error)
}

type stage int

const (
	fetch stage = iota
	use
	end
)

type readerAdapter struct {
	rd     io.Reader
	stage  stage
	offset int
	b      byte
	err    error
}

func NewBitReader(rd io.Reader) BitReader {
	return &readerAdapter{rd: rd}
}

func (r *readerAdapter) ReadBit() (Bit, error) {
	if r.stage == end {
		return ZERO, r.err
	}

	if r.stage == fetch {
		r.fetch()
		if r.err != nil {
			r.stage = end
		} else {
			r.stage++
		}
	} else if r.offset == 8 {
		r.offset = 0
		r.stage = fetch

		return r.ReadBit()
	}

	bit := IndexByte(r.b, r.offset)
	r.offset++

	return bit, r.err
}

func (r *readerAdapter) ReadBits(b []Bit) (int, error) {
	if r.stage == end {
		return 0, r.err
	}

	if len(b) == 0 {
		return 0, nil
	}

	if r.stage == fetch {
		r.fetch()
		if r.err != nil {
			r.stage = end
		} else {
			r.stage++
		}
	} else if r.offset == 8 {
		r.offset = 0
		r.stage = fetch

		return r.ReadBits(b)
	}

	var n int

	for i := 0; i < len(b) && r.offset < 8; i++ {
		b[i] = IndexByte(r.b, r.offset)
		r.offset++
		n++
	}

	if n == len(b) {
		return n, r.err
	}

	for ; n < len(b); n++ {
		if r.offset == 8 {
			r.offset = 0

			r.fetch()
			if r.err != nil {
				r.stage = end
				return n, r.err
			}
		}

		b[n] = IndexByte(r.b, r.offset)
		r.offset++
	}

	return n, r.err
}

func (r *readerAdapter) fetch() {
	var buf [1]byte
	_, r.err = r.rd.Read(buf[:])
	r.b = buf[0]
}

func (r *readerAdapter) Read([]byte) (int, error) {
	panic("To be implemented")
}
