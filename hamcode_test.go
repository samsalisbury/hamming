package main

import (
	"fmt"
	"testing"
)

func TestHamCode(t *testing.T) {

	cases := map[Nibble]byte{
		0:  0,
		1:  105,
		2:  42,
		3:  67,
		4:  76,
		11: 51,
		15: 127,
	}

	for in := Nibble(0); in < 15; in++ {
		want, ok := cases[in]
		if !ok {
			continue
		}
		got := HamCode(in)
		if got != want {
			t.Errorf("%02d:%03d  HamCode(%04b) == %07b; want %07b", in, want, in, got, want)
		}
	}

	// Generate full set of codes...
	for i := Nibble(0); i < 16; i++ {
		o := HamCode(i)
		fmt.Printf("HamCode(%04b) == %07b\n", i, o)
	}
}

// HamCode returns 7/4 ham encoded byte of input data nibble.
func HamCode(b Nibble) byte {
	matrix := [7]Nibble{13, 11, 8, 7, 4, 2, 1}
	var code byte
	for i := byte(0); i < 7; i++ {
		code |= byte((PopCount(matrix[i]&b) % 2) << (6 - i))
	}
	return code
}
