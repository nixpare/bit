package bitio

import (
	"io"
)

type BitWriter interface {
	BitWrite(...Bit) (int, error)
	Flush() (int, error)
	io.WriteCloser
}

type writerAdapter struct {
	wr io.Writer
	b  []Bit
}

func NewBitWriter(w io.Writer) BitWriter {
	return &writerAdapter{wr: w}
}

func (w *writerAdapter) BitWrite(b ...Bit) (int, error) {
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
		bytes := Convert[[]byte](w.b[:end])

		payload = append(payload, bytes...)
		n += end
		w.b = w.b[end:]
	}

	if empty && len(w.b) > 0 {
		b := Convert[byte](w.b)
		
		payload = append(payload, b)
		n += len(w.b)
		w.b = w.b[:0]
	}

	nBytes, err := w.wr.Write(payload)
	return min(n, nBytes*8), err
}

func (w *writerAdapter) Write(b []byte) (int, error) {
	n, err := w.BitWrite(Bits(b)...)
	n /= 8
	return n, err 
}

func (w *writerAdapter) Close() error {
	_, err := w.flush(true)
	return err
}

type ByterWriter[T Byter] struct {
	wr BitWriter
}

func NewByterWriter[T Byter](w io.Writer) *ByterWriter[T] {
	if w, ok := w.(BitWriter); ok {
		return &ByterWriter[T]{ wr: w }
	}

	return &ByterWriter[T]{ wr: NewBitWriter(w) }
}

func (w *ByterWriter[T]) ByterWrite(x T) (int, error) {
	b := Bits(x)
	return w.BitWrite(b...)
}

func (w *ByterWriter[T]) BitWrite(b ...Bit) (int, error) {
	return w.wr.BitWrite(b...)
}

func (w *ByterWriter[T]) Flush() (int, error) {
	return w.wr.Flush()
}

func (w *ByterWriter[T]) Write(b []byte) (int, error) {
	return w.wr.Write(b)
}

func (w *ByterWriter[T]) Close() error {
	return w.wr.Close()
}

func ByterWrite[T Byter](w BitWriter, x T) (int, error) {
	b := Bits(x)
	return w.BitWrite(b...)
}
