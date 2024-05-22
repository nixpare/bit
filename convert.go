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