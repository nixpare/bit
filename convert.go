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