package bitio

import (
	"io"
)

type BitWriter interface {
	WriteBit(Bit) error
	WriteBits([]Bit) (int, error)
	Flush() (int, error)

	Write([]byte) (int, error)
	Close() error
}

type writerAdapter struct {
	wr io.Writer
	b  []Bit
}

func NewBitWriter(wr io.Writer) BitWriter {
	return &writerAdapter{wr: wr}
}

func (w *writerAdapter) WriteBit(b Bit) error {
	_, err := w.WriteBits([]Bit{b})
	return err
}

func (w *writerAdapter) WriteBits(b []Bit) (int, error) {
	w.b = append(w.b, b...)
	return w.flush(false)
}

func (w *writerAdapter) Flush() (int, error) {
	return w.flush(true)
}

func (w *writerAdapter) flush(empty bool) (int, error) {
	var payload []byte
	var n int

	if len(w.b) >= 8 {
		end := len(w.b) / 8 * 8
		bytes := ByterFromBits[[]byte](w.b[:end])

		payload = append(payload, bytes...)
		n += end
		w.b = w.b[end:]
	}

	if empty && len(w.b) > 0 {
		b := ByterFromBits[byte](w.b)
		
		payload = append(payload, b)
		n += len(w.b)
		w.b = w.b[:0]
	}

	nBytes, err := w.wr.Write(payload)
	return min(n, nBytes*8), err
}

func (w *writerAdapter) Write(b []byte) (int, error) {
	return w.WriteBits(Bits(b))
}

func (w *writerAdapter) Close() error {
	_, err := w.flush(true)
	return err
}
