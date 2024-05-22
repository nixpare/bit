package bitio

import (
	"bytes"
	"errors"
	"io"
	"slices"
	"testing"
)

func TestBitReaderReadBit(t *testing.T) {
	type testData struct {
		from []byte
		to   []Bit
	}

	data := []testData{
		{from: []byte{1}, to: []Bit{ONE, ZERO, ZERO}},
		{from: []byte{5}, to: []Bit{ONE, ZERO, ONE, ZERO}},
		{from: []byte{6, 9}, to: []Bit{
			ZERO, ONE, ONE, ZERO, ZERO, ZERO, ZERO, ZERO,
			ONE, ZERO, ZERO, ONE, ZERO, ZERO, ZERO, ZERO,
		}},
	}

	for _, d := range data {
		rd := NewBitReader(bytes.NewReader(d.from))

		result, err := rd.ReadBit()
		for i := 0; i < len(d.from)*8; i++ {
			if err != nil {
				t.Errorf("error reading bit %d from %v -> %v", i, d.from, err)
			}

			if i < len(d.to) && result != d.to[i] {
				t.Errorf("error reading bit %d from %v -> expected %v, found %v", i, d.from, d.to[i], result)
			}

			result, err = rd.ReadBit()
		}

		result, err = rd.ReadBit()
		if result != ZERO || !errors.Is(err, io.EOF) {
			t.Errorf("expected (ZERO, io.EOF) at the end of read, found (%v, %v)", result, err)
		}
	}
}

func TestBitReaderReadBits(t *testing.T) {
	type testData struct {
		from  []byte
		reads [][]Bit
	}

	data := []testData{
		{from: []byte{1, 5}, reads: [][]Bit{
			{ONE, ZERO, ZERO, ZERO, ZERO, ZERO, ZERO, ZERO},
			{ONE, ZERO, ONE, ZERO, ZERO, ZERO, ZERO, ZERO},
		}},
		{from: []byte{5, 1}, reads: [][]Bit{
			{ONE, ZERO, ONE, ZERO, ZERO, ZERO, ZERO, ZERO},
			{ONE, ZERO, ZERO, ZERO, ZERO, ZERO, ZERO, ZERO},
		}},
		{from: []byte{5, 1}, reads: [][]Bit{
			{ONE, ZERO, ONE, ZERO, ZERO},
			{ZERO, ZERO, ZERO, ONE, ZERO, ZERO, ZERO},
		}},
		{from: []byte{6, 9}, reads: [][]Bit{
			{ZERO, ONE, ONE, ZERO, ZERO, ZERO, ZERO, ZERO},
			{ONE, ZERO, ZERO, ONE, ZERO, ZERO, ZERO, ZERO},
		}},
	}

	for _, d := range data {
		rd := NewBitReader(bytes.NewReader(d.from))

		for i, read := range d.reads {
			buf := make([]Bit, len(read))
			n, err := rd.ReadBits(buf)
			if err != nil {
				t.Errorf("error read n°%d of bits from %v -> %v", i, d.from, err)
			}

			if n != len(read) || n != len(buf) {
				t.Errorf("error read n°%d of bits from %v -> n result diff from read: %d, buf: %d", i, d.from, len(read), len(buf))
			}

			if !slices.Equal(buf, read) {
				t.Errorf("error reading n°%d of bits from %v -> expected %v, found %v", i, d.from, read, buf)
			}
		}
	}
}
