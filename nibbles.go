package main

// Nibbles is a slice of bytes with the least significant 4 bits
// of each representing a nibble.
type Nibbles []byte

// Nibble breaks down the input bytes into nibbles inside bytes.
// The returned byte slice is twice the length of the input slice
// and the least significant 4 bits encode the nibble.
func Nibble(source []byte) Nibbles {
	ns := make(Nibbles, 2*len(source))
	for i, b := range source {
		ns[i*2] = b >> 4
		ns[(i*2)+1] = (b & 15)
	}
	return ns
}

// Bytes is the inverse of Nibble; it assembles the nibbles into
// contiguous bytes.
func (ns Nibbles) Bytes() []byte {
	output := make([]byte, len(ns)/2)
	for i := range output {
		output[i] = (ns[i*2] << 4) | (ns[(i*2)+1])
	}
	return output
}
