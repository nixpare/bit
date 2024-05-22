package bitio

import (
	"math"
	"slices"
)

type Bit bool

const (
	ZERO Bit = false
	ONE  Bit = true
)

func (b Bit) String() string {
	if b == ONE {
		return "1"
	}
	return "0"
}

func Int(b Bit) int {
	if b == ZERO {
		return 0
	}
	return 1
}

type OneByter interface {
	~uint8 | ~int8
}

type TwoByter interface {
	~uint16 | ~int16
}

type FourByter interface {
	~uint32 | ~int32 | ~float32
}

type EightByter interface {
	~uint64 | ~int64 | ~float64
}

type Byter interface {
	OneByter | TwoByter | FourByter | EightByter
}

type Endianess int

const (
	LITTLE_ENDIAN Endianess = iota
	BIG_ENDIAN
)

func Bits[T Byter](b T, e Endianess) []Bit {
	var from []byte

	switch b := any(b).(type) {
	case uint8:
		from = []byte{byte(b)}
	case int8:
		from = []byte{byte(b)}

	case uint16:
		from = []byte{byte(b), byte(b >> 8)}
	case int16:
		from = []byte{byte(b), byte(b >> 8)}

	case uint32:
		from = []byte{byte(b), byte(b >> 8), byte(b >> 16), byte(b >> 24)}
	case int32:
		from = []byte{byte(b), byte(b >> 8), byte(b >> 16), byte(b >> 24)}
	case float32:
		u := math.Float32bits(b)
		from = []byte{byte(u), byte(u >> 8), byte(u >> 16), byte(u >> 24)}

	case uint64:
		from = []byte{byte(b), byte(b >> 8), byte(b >> 16), byte(b >> 24), byte(b >> 32), byte(b >> 40), byte(b >> 48), byte(b >> 56)}
	case int64:
		from = []byte{byte(b), byte(b >> 8), byte(b >> 16), byte(b >> 24), byte(b >> 32), byte(b >> 40), byte(b >> 48), byte(b >> 56)}
	case float64:
		u := math.Float64bits(b)
		from = []byte{byte(u), byte(u >> 8), byte(u >> 16), byte(u >> 24), byte(u >> 32), byte(u >> 40), byte(u >> 48), byte(u >> 56)}
	}

	if e == BIG_ENDIAN {
		slices.Reverse(from)
	}

	bits := make([]Bit, 0, len(from)*8)
	for _, b := range from {
		bits = append(bits, BitsFromByte(b)...)
	}
	return bits
}

func BitFromByte(b byte, pos int) Bit {
	if pos >= 8 {
		panic("a byte has only 8 bits")
	}

	if (b>>pos)&1 == 1 {
		return ONE
	}
	return ZERO
}

func BitsFromByte(b byte) []Bit {
	bits := make([]Bit, 8)

	for i := range 8 {
		if (b>>i)&1 == 1 {
			bits[i] = ONE
		} else {
			bits[i] = ZERO
		}
	}

	return bits
}

func BitsFromBytes(b []byte) []Bit {
	bits := make([]Bit, 0, len(b)*8)
	for _, x := range b {
		bits = append(bits, BitsFromByte(x)...)
	}

	return bits
}

func ReverseEndianess(b []Bit) {
	for i := 0; i < len(b)/2; i += 8 {
		for j := 0; j < 8; j++ {
			left, right := i+j, len(b)-(i+8)+j
			println(left, right)

			b[left], b[right] = b[right], b[left]
		}
	}
}