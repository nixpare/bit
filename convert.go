package bitio

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

func Uint32(b []Bit, e Endianess) uint32 {
    if len(b) > 32 {
        panic("expected 32 bits")
    }

    var x uint32
    for i, bit := range b {
        if bit == ONE {
            if e == LITTLE_ENDIAN {
                x |= 1 << i
            } else {
                x |= 1 << (31 - i)
            }
        }
    }

    return x
}