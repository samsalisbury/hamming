package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestDisasembleNibbles(t *testing.T) {

	source, err := ioutil.ReadFile("hc.png")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Source length: %d bytes\n", len(source))

	// Nibbles will hold bytes where the 4 least significant bits
	// contain a nibble of the source data.
	nibbles := BytesToNibbles(source)
	if len(nibbles) != 2*len(source) {
		t.Fatalf("got %d nibbles; want %d (2x)", len(nibbles), 2*len(source))
	}
	for i := 0; i < len(source); i++ {
		if i == 10 {
			// Fast forward to last 10 bytes.
			i = len(source) - 10
		}
		first, second := nibbles[2*i], nibbles[(2*i)+1]
		b := source[i]
		fmt.Printf("byte %05d %08b = nibbles %08b + %08b\n", i, b, first, second)
	}

	output := nibbles.Bytes()

	for i := range source {
		if source[i] != output[i] {
			t.Fatalf("reassembled output does not match input")
		}
	}

}
