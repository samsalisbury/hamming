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

func TestHamParityCheck(t *testing.T) {

	// Generate full set of valid codes, ensure they have zero syndrome.
	for i := Nibble(0); i < 16; i++ {
		in := HamCode(i)
		if got := HamParityCheck(in); got != 0 {
			t.Errorf("HamParityCheck(%07b) == %04b; want %04b", in, got, 0)
		}
	}

	cases := []struct {
		in   byte
		want Nibble
	}{
		{in: 52, want: 1},
		{in: 86, want: 4},
		{in: 43, want: 7},
	}

	for _, c := range cases {
		got := HamParityCheck(c.in)
		if got != c.want {
			t.Errorf("HamParityCheck(%07b) == %04b; want %04b", c.in, got, c.want)
		}
	}

}

func TestHamErrCorrect(t *testing.T) {

	cases := []struct {
		faulty   byte
		syndrome Nibble
		want     byte
	}{
		{faulty: 118, syndrome: 5, want: 102},
	}
	for _, c := range cases {
		got := HamErrCorrect(c.faulty, c.syndrome)
		if got != c.want {
			t.Errorf("HamErrCorrect(%07b, %04b) == %07b; want %07b",
				c.faulty, c.syndrome, got, c.want)
		}
	}

}

func TestHamDecode(t *testing.T) {

	cases := map[byte]Nibble{
		0:   0,
		105: 1,
		42:  2,
		67:  3,
		76:  4,
		51:  11,
		127: 15,
	}

	for i := byte(0); i < 127; i++ {
		want, ok := cases[i]
		if !ok {
			continue
		}
		got := HamDecode(i)
		if got != want {
			t.Errorf("HamDecode(%07b) == %04b; want %04b", i, got, want)
		}
	}
}
