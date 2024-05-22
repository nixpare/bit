package bitio

import (
	"math"
)

func IndexByte(b byte, pos int) Bit {
	if pos >= 8 {
		panic("a byte has only 8 bits")
	}

	if (b>>pos)&1 == 1 {
		return ONE
	}
	return ZERO
}

func ByterFromBits[T Byter](bits []Bit) T {
    result := new(T)

	switch r := any(result).(type) {

	case *uint8:
        for i := 0; i < len(bits) && i < 8; i++ {
            if bits[i] == ONE {
                *r |= 1 << i
            }
        }
	case *int8:
		for i := 0; i < len(bits) && i < 8; i++ {
            if bits[i] == ONE {
                *r |= 1 << i
            }
        }

	case *uint16:
		for i := 0; i < len(bits) && i < 16; i++ {
            if bits[i] == ONE {
                *r |= 1 << i
            }
        }
	case *int16:
		for i := 0; i < len(bits) && i < 16; i++ {
            if bits[i] == ONE {
                *r |= 1 << i
            }
        }

	case *uint32:
		for i := 0; i < len(bits) && i < 32; i++ {
            if bits[i] == ONE {
                *r |= 1 << i
            }
        }
	case *int32:
		for i := 0; i < len(bits) && i < 32; i++ {
            if bits[i] == ONE {
                *r |= 1 << i
            }
        }
	case *float32:
		u := ByterFromBits[uint32](bits)
        *r = math.Float32frombits(u)

	case *uint64:
		for i := 0; i < len(bits) && i < 64; i++ {
            if bits[i] == ONE {
                *r |= 1 << i
            }
        }
	case *int64:
		for i := 0; i < len(bits) && i < 64; i++ {
            if bits[i] == ONE {
                *r |= 1 << i
            }
        }
	case *float64:
        u := ByterFromBits[uint64](bits)
        *r = math.Float64frombits(u)

	case *[]uint8:
        var b *uint8
        for i := 0; i < len(bits); i++ {
            if i % 8 == 0 {
                *r = append(*r, 0)
                b = &(*r)[len(*r)-1]
            }

            if bits[i] == ONE {
                *b |= 1 << i % 8
            }
        }

	case *[]int8:
		var b *int8
        for i := 0; i < len(bits); i++ {
            if i % 8 == 0 {
                *r = append(*r, 0)
                b = &(*r)[len(*r)-1]
            }

            if bits[i] == ONE {
                *b |= 1 << i % 8
            }
        }

	case *[]uint16:
        var b *uint16
        for i := 0; i < len(bits); i++ {
            if i % 16 == 0 {
                *r = append(*r, 0)
                b = &(*r)[len(*r)-1]
            }

            if bits[i] == ONE {
                *b |= 1 << i % 16
            }
        }

	case *[]int16:
		var b *int16
        for i := 0; i < len(bits); i++ {
            if i % 16 == 0 {
                *r = append(*r, 0)
                b = &(*r)[len(*r)-1]
            }

            if bits[i] == ONE {
                *b |= 1 << i % 16
            }
        }

	case *[]uint32:
        var b *uint32
        for i := 0; i < len(bits); i++ {
            if i % 32 == 0 {
                *r = append(*r, 0)
                b = &(*r)[len(*r)-1]
            }

            if bits[i] == ONE {
                *b |= 1 << i % 32
            }
        }

	case *[]int32:
		var b *int32
        for i := 0; i < len(bits); i++ {
            if i % 32 == 0 {
                *r = append(*r, 0)
                b = &(*r)[len(*r)-1]
            }

            if bits[i] == ONE {
                *b |= 1 << i % 32
            }
        }
	case *[]float32:
		for i := 0; i < len(bits); i += 32 {
            end := min(i+32, len(bits))

            u := ByterFromBits[uint32](bits[i:end])
            *r = append(*r, math.Float32frombits(u))
		}

	case *[]uint64:
        var b *uint64
        for i := 0; i < len(bits); i++ {
            if i % 64 == 0 {
                *r = append(*r, 0)
                b = &(*r)[len(*r)-1]
            }

            if bits[i] == ONE {
                *b |= 1 << i % 64
            }
        }

	case *[]int64:
		var b *int64
        for i := 0; i < len(bits); i++ {
            if i % 64 == 0 {
                *r = append(*r, 0)
                b = &(*r)[len(*r)-1]
            }

            if bits[i] == ONE {
                *b |= 1 << i % 64
            }
        }
	case *[]float64:
		for i := 0; i < len(bits); i += 32 {
            end := min(i+32, len(bits))

            u := ByterFromBits[uint64](bits[i:end])
            *r = append(*r, math.Float64frombits(u))
		}

	default:
		panic("type not implemented")

	}

    return *result
}

func ByterFromBytes[T Byter](bytes []byte) T {
    bits := Bits(bytes)
    return ByterFromBits[T](bits)
}
