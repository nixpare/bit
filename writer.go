package bitio

import "io"

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

	for len(w.b) >= 8 {
		var p byte
		for i := range 8 {
			if w.b[i] == ONE {
				p |= 1 << i
			}
		}

		payload = append(payload, p)
		n += 8
		w.b = w.b[8:]
	}

	if empty && len(w.b) > 0 {
		n += len(w.b)

		var p byte
		for i := range len(w.b) {
			if w.b[i] == ONE {
				p |= 1 << i
			}
		}

		payload = append(payload, p)
		w.b = w.b[:0]
	}

	nBytes, err := w.wr.Write(payload)
	return min(n, nBytes*8), err
}

func (w *writerAdapter) Write(b []byte) (int, error) {
	return w.WriteBits(BitsFromBytes(b))
}

func (w *writerAdapter) Close() error {
	_, err := w.flush(true)
	return err
}
