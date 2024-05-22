package bitio

import (
	"math"
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

func (b Bit) Int() int {
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

type SingleByter interface {
	OneByter | TwoByter | FourByter | EightByter
}

type MultiByter interface {
	~[]uint8 | ~[]int8 |
	~[]uint16 | ~[]int16 |
	~[]uint32 | ~[]int32 | ~[]float32 |
	~[]uint64 | ~[]int64 | ~[]float64
}

type Byter interface {
	SingleByter | MultiByter
}

type Endianess int

const (
	LITTLE_ENDIAN Endianess = iota
	BIG_ENDIAN
)

func Bits[T Byter](b T) []Bit {
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

	case []uint8:
		for _, x := range b {
			from = append(from, byte(x))
		}
	case []int8:
		for _, x := range b {
			from = append(from, byte(x))
		}

	case []uint16:
		for _, x := range b {
			from = append(from, byte(x), byte(x >> 8))
		}
	case []int16:
		for _, x := range b {
			from = append(from, byte(x), byte(x >> 8))
		}

	case []uint32:
		for _, x := range b {
			from = append(from, byte(x), byte(x >> 8), byte(x >> 16), byte(x >> 24))
		}
	case []int32:
		for _, x := range b {
			from = append(from, byte(x), byte(x >> 8), byte(x >> 16), byte(x >> 24))
		}
	case []float32:
		for _, x := range b {
			u := math.Float32bits(x)
			from = append(from, byte(u), byte(u >> 8), byte(u >> 16), byte(u >> 24))
		}

	case []uint64:
		for _, x := range b {
			from = append(from, byte(x), byte(x >> 8), byte(x >> 16), byte(x >> 24), byte(x >> 32), byte(x >> 40), byte(x >> 48), byte(x >> 56))
		}
	case []int64:
		for _, x := range b {
			from = append(from, byte(x), byte(x >> 8), byte(x >> 16), byte(x >> 24), byte(x >> 32), byte(x >> 40), byte(x >> 48), byte(x >> 56))
		}
	case []float64:
		for _, x := range b {
			u := math.Float64bits(x)
			from = append(from, byte(u), byte(u >> 8), byte(u >> 16), byte(u >> 24), byte(u >> 32), byte(u >> 40), byte(u >> 48), byte(u >> 56))
		}

	default:
		panic("type not implemented")

	}

	bits := make([]Bit, len(from) * 8)
	for i, b := range from {
		for j := range 8 {
			if (b >> j) & 1 == 1 {
				bits[i*8 + j] = ONE
			} else {
				bits[i*8 + j] = ZERO
			}
		}
	}
	return bits
}

func MultiBits[T Byter](b []T) []Bit {
	var tmp T
	bits := make([]Bit, 0, BitsNum(tmp) * len(b))
	for _, x := range b {
		bits = append(bits, Bits(x)...)
	}

	return bits
}

func BitsNum[T Byter](b T) int {
	switch b := any(b).(type) {
	case uint8, int8:
		return 8

	case uint16, int16:
		return 16

	case uint32, int32, float32:
		return 32

	case uint64, int64, float64:
		return 64

	case []uint8:
		return 8 * len(b)
	case []int8:
		return 8 * len(b)

	case []uint16:
		return 16 * len(b)
	case []int16:
		return 16 * len(b)

	case []uint32:
		return 32 * len(b)
	case []int32:
		return 32
	case []float32:
		return 32

	case []uint64:
		return 64 * len(b)
	case []int64:
		return 64 * len(b)
	case []float64:
		return 64 * len(b)
	
	default:
		panic("type not implemented")
	}
}

func MinBitsNum[T Byter](b T) int {
	bits := Bits(b)
	var num int

	for i, x := range bits {
		if x == ONE {
			num = i+1
		}
	}

	return num
}

func ReverseEndianess(b []Bit) {
	for i := 0; i < len(b)/2; i += 8 {
		for j := 0; j < 8; j++ {
			left, right := i+j, len(b)-(i+8)+j
			
			b[left], b[right] = b[right], b[left]
		}
	}
}
