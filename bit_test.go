package bitio

import (
	"slices"
	"testing"
)

func TestBitFromByte(t *testing.T) {
	type testData struct {
		from byte
		pos  int
		to   Bit
	}

	data := []testData{
		{from: 1, pos: 0, to: ONE},
		{from: 1, pos: 1, to: ZERO},
		{from: 1, pos: 1, to: ZERO},

		{from: 5, pos: 0, to: ONE},
		{from: 5, pos: 1, to: ZERO},
		{from: 5, pos: 2, to: ONE},
		{from: 5, pos: 3, to: ZERO},
	}

	for _, d := range data {
		if result := IndexByte(d.from, d.pos); result != d.to {
			t.Errorf("error conversion from byte to bit: from %b (pos %d) -> expected %v, found %v", d.from, d.pos, d.to, result)
		}
	}
}

func TestBitsFromByte(t *testing.T) {
	type testData struct {
		from byte
		to   []Bit
	}

	data := []testData{
		{from: 0, to: []Bit{ZERO, ZERO, ZERO, ZERO, ZERO, ZERO, ZERO, ZERO}},
		{from: 1, to: []Bit{ONE, ZERO, ZERO, ZERO, ZERO, ZERO, ZERO, ZERO}},
		{from: 64, to: []Bit{ZERO, ZERO, ZERO, ZERO, ZERO, ZERO, ONE, ZERO}},
		{from: 128, to: []Bit{ZERO, ZERO, ZERO, ZERO, ZERO, ZERO, ZERO, ONE}},
		{from: 255, to: []Bit{ONE, ONE, ONE, ONE, ONE, ONE, ONE, ONE}},
	}

	for _, d := range data {
		if result := Bits(d.from); !slices.Equal(result, d.to) {
			t.Errorf("error conversion from byte to bit: from %b -> expected %v, found %v", d.from, d.to, result)
		}
	}
}

func TestBitsAndReverse(t *testing.T) {
	type testData32 struct {
		from float32
		to   []Bit
	}

	data32 := []testData32{
		{from: 3.14, to: []Bit{
			ONE, ONE, ZERO, ZERO, ZERO, ZERO, ONE, ONE,
			ONE, ZERO, ONE, ZERO, ONE, ONE, ONE, ONE,
			ZERO, ZERO, ZERO, ONE, ZERO, ZERO, ONE, ZERO,
			ZERO, ZERO, ZERO, ZERO, ZERO, ZERO, ONE, ZERO,
		}},
	}

	for _, d := range data32 {
		if result := Bits(d.from); !slices.Equal(result, d.to) {
			t.Errorf("error conversion from float32 to bit: from %v -> expected %v, found %v", d.from, d.to, result)
		}

		if reverse := ByterFromBits[float32](d.to); reverse != d.from {
			t.Errorf("error conversion from bits to float32: from %v -> expected %v, found %v", d.to, d.from, reverse)
		}
	}

	type testData64 struct {
		from float64
		to   []Bit
	}

	data64 := []testData64{
		{from: 3.34567896523, to: []Bit{
			ONE, ONE, ONE, ONE, ZERO, ONE, ZERO, ONE,
			ONE, ONE, ONE, ONE, ONE, ZERO, ZERO, ONE,
			ZERO, ZERO, ONE, ZERO, ONE, ZERO, ONE, ZERO,
			ONE, ZERO, ONE, ZERO, ONE, ZERO, ONE, ZERO,

			ONE, ONE, ZERO, ZERO, ONE, ONE, ONE, ONE,
			ONE, ONE, ZERO, ZERO, ZERO, ZERO, ONE, ONE,
			ZERO, ONE, ZERO, ONE, ZERO, ZERO, ZERO, ZERO,
			ZERO, ZERO, ZERO, ZERO, ZERO, ZERO, ONE, ZERO,
		}},
	}

	for _, d := range data64 {
		if result := Bits(d.from); !slices.Equal(result, d.to) {
			t.Errorf("error conversion from float64 to bit: from %b -> expected %v, found %v", d.from, d.to, result)
		}

		if reverse := ByterFromBits[float64](d.to); reverse != d.from {
			t.Errorf("error conversion from bits to float64: from %v -> expected %v, found %v", d.to, d.from, reverse)
		}
	}
}
