package main

// Nibbles is a slice of bytes with the least significant 4 bits of each
// representing a nibble.
type Nibbles []Nibble

// A Nibble is a byte representing a nibble in its 4 least significant bits.
// The 4 most significant bits are always zero.
type Nibble byte

// BytesToNibbles breaks down the input bytes into nibbles inside bytes.
// The returned byte slice is twice the length of the input slice
// and the least significant 4 bits encode the nibble.
func BytesToNibbles(source []byte) Nibbles {
	ns := make(Nibbles, 2*len(source))
	for i, b := range source {
		ns[i*2] = Nibble(b >> 4)
		ns[(i*2)+1] = Nibble((b & 15))
	}
	return ns
}

// Bytes is the inverse of Nibble; it assembles the nibbles into
// contiguous bytes.
func (ns Nibbles) Bytes() []byte {
	output := make([]byte, len(ns)/2)
	for i := range output {
		output[i] = byte((ns[i*2] << 4) | (ns[(i*2)+1]))
	}
	return output
}

// PopCount returns the number of on bits in the Nibble.
func PopCount(b Nibble) Nibble {
	var c Nibble
	for b != 0 {
		b &= b - 1
		c++
	}
	return c
}

// PopCountByte returns the number of on bits in the Nibble.
func PopCountByte(b byte) Nibble {
	var c Nibble
	for b != 0 {
		b &= b - 1
		c++
	}
	return c
}
