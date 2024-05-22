package bitio

import "math"

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

func Byte(bits []Bit) byte {
    return byte(Uint8(bits))
}

func Bytes(bits []Bit) []byte {
    var bytes []byte
    for i := 0; i < len(bits); i += 8 {
        bytes = append(bytes, Byte(bits[i:i+8]))
    }

    return bytes
}

func Uint8(b []Bit) uint8 {
    if len(b) > 8 {
        panic("expected max 8 bits")
    }

    var x uint8
    for i, bit := range b {
        if bit == ONE {
            x |= 1 << i
        }
    }

    return x
}

func Uint16(b []Bit) uint16 {
    if len(b) > 16 {
        panic("expected max 16 bits")
    }

    var x uint16
    for i, bit := range b {
        if bit == ONE {
            x |= 1 << i
        }
    }

    return x
}

func Uint32(b []Bit) uint32 {
    if len(b) > 32 {
        panic("expected max 32 bits")
    }

    var x uint32
    for i, bit := range b {
        if bit == ONE {
            x |= 1 << i
        }
    }

    return x
}

func Uint64(b []Bit) uint64 {
    if len(b) > 64 {
        panic("expected max 64 bits")
    }

    var x uint64
    for i, bit := range b {
        if bit == ONE {
            x |= 1 << i
        }
    }

    return x
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
        for i := 0; i < len(bits); i++ {
            var x uint8
            if bits[i] == ONE {
                x |= 1 << i % 8
            }

            if i % 8 == 8-1 {
                *r = append(*r, x)
            }
        }
	case *[]int8:
		for i := 0; i < len(bits); i++ {
            var x int8
            if bits[i] == ONE {
                x |= 1 << i % 8
            }

            if i % 8 == 8-1 {
                *r = append(*r, x)
            }
        }

	case *[]uint16:
        for i := 0; i < len(bits); i++ {
            var x uint16
            if bits[i] == ONE {
                x |= 1 << i % 16
            }

            if i % 16 == 16-1 {
                *r = append(*r, x)
            }
        }
	case *[]int16:
		for i := 0; i < len(bits); i++ {
            var x int16
            if bits[i] == ONE {
                x |= 1 << i % 16
            }

            if i % 16 == 16-1 {
                *r = append(*r, x)
            }
        }

	case *[]uint32:
        for i := 0; i < len(bits); i++ {
            var x uint32
            if bits[i] == ONE {
                x |= 1 << i % 32
            }

            if i % 32 == 32-1 {
                *r = append(*r, x)
            }
        }
	case *[]int32:
		for i := 0; i < len(bits); i++ {
            var x int32
            if bits[i] == ONE {
                x |= 1 << i % 32
            }

            if i % 32 == 32-1 {
                *r = append(*r, x)
            }
        }
	case *[]float32:
		for i := 0; i < len(bits); i += 32 {
            end := min(i+32, len(bits))

            u := ByterFromBits[uint32](bits[i:end])
            *r = append(*r, math.Float32frombits(u))
		}

	case *[]uint64:
        for i := 0; i < len(bits); i++ {
            var x uint64
            if bits[i] == ONE {
                x |= 1 << i % 64
            }

            if i % 64 == 64-1 {
                *r = append(*r, x)
            }
        }
	case *[]int64:
		for i := 0; i < len(bits); i++ {
            var x int64
            if bits[i] == ONE {
                x |= 1 << i % 64
            }

            if i % 64 == 64-1 {
                *r = append(*r, x)
            }
        }
	case *[]float64:
		for i := 0; i < len(bits); i += 32 {
            end := min(i+32, len(bits))

            u := ByterFromBits[uint64](bits[i:end])
            *r = append(*r, math.Float64frombits(u))
		}

	default:
		panic("not implemented")

	}

    return *result
}

func ByterFromBytes[T Byter](bytes []byte) T {
    bits := Bits(bytes)
    return ByterFromBits[T](bits)
}