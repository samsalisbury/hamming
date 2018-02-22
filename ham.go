package main

// HamCode returns 7/4 ham encoded byte of input data nibble.
func HamCode(b Nibble) byte {
	matrix := [7]Nibble{13, 11, 8, 7, 4, 2, 1}
	var code byte
	for i := byte(0); i < 7; i++ {
		code |= byte((PopCount(matrix[i]&b) % 2) << (6 - i))
	}
	return code
}

// HamParityCheck returns a number indicating which bit is wrong, or zero if all
// bits are correct.
func HamParityCheck(b byte) Nibble {
	matrix := [3]byte{85, 51, 15}
	var syndrome Nibble
	for i := Nibble(0); i < 3; i++ {
		syndrome |= (PopCountByte(matrix[i]&b) % 2) << (2 - i)
	}
	return syndrome
}

// HamErrCorrect takes a code word e with a single incorrect bit and the result
// of HamParityCheck to return the corrected code word.
func HamErrCorrect(e byte, syn Nibble) byte {
	v := byte(1 << (syn - 1))
	return e ^ v
}

// HamDecode takes a correct code word and returns the original data decoded.
func HamDecode(b byte) Nibble {
	matrix := [4]byte{16, 4, 2, 1}
	var n Nibble
	for i := Nibble(0); i < 4; i++ {
		n |= (PopCountByte(matrix[i]&b) % 2) << (3 - i)
	}
	return n
}
