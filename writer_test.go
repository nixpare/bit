package bit

import (
	"bytes"
	"slices"
	"testing"
)

func TestBitWriterWriteBits(t *testing.T) {
	type testData struct {
		to     []byte
		writes [][]Bit
	}

	data := []testData{
		{to: []byte{1, 5}, writes: [][]Bit{
			{ONE, ZERO, ZERO, ZERO, ZERO, ZERO, ZERO, ZERO},
			{ONE, ZERO, ONE, ZERO, ZERO, ZERO, ZERO, ZERO},
		}},
		{to: []byte{5, 1}, writes: [][]Bit{
			{ONE, ZERO, ONE, ZERO, ZERO},
			{ZERO, ZERO, ZERO, ONE, ZERO, ZERO, ZERO},
		}},
		{to: []byte{6, 9}, writes: [][]Bit{
			{ZERO, ONE, ONE, ZERO, ZERO, ZERO, ZERO, ZERO},
			{ONE, ZERO, ZERO, ONE},
		}},
	}

	for _, d := range data {
		buf := bytes.NewBuffer(nil)
		wr := NewBitWriter(buf)

		var n, nWrites int

		for _, write := range d.writes {
			nWrites += len(write)

			n2, err := wr.WriteBits(write)
			n += n2
			if err != nil {
				t.Errorf("error writing bits %v -> %v", write, err)
			}
		}

		n2, err := wr.Flush()
		n += n2
		if err != nil {
			t.Errorf("error flushing bits %v -> %v", d.writes, err)
		}

		if n != nWrites {
			t.Errorf("expected %d bits wrote, found %d (for %v)", nWrites, n, d.writes)
		}

		if !slices.Equal(buf.Bytes(), d.to) {
			t.Errorf("error writing bits %v -> expected %v, found %v", d.writes, d.to, buf.Bytes())
		}
	}
}
