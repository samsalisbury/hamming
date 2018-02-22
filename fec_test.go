package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"testing"
)

func randomlyCorrupt(b *byte) {
	*b = *b + (1 << (rand.Uint64() % 8))
}

func TestRandomlyCorrupt(t *testing.T) {

	source, err := ioutil.ReadFile("hc.png")
	if err != nil {
		t.Fatal(err)
	}

	for n := 1; n <= 10; n++ {

		corrupted := make([]byte, len(source))
		copy(corrupted, source)

		for i := range corrupted {
			if i < 2000 {
				continue
			}
			if rand.Intn(10) < 8 {
				randomlyCorrupt(&corrupted[i])
			}
		}

		name := fmt.Sprintf("corrupt-hc-%d.png", n)
		if err := ioutil.WriteFile(name, corrupted, 0666); err != nil {
			t.Fatal(err)
		}
	}
}

func TestDisasembleNibbles(t *testing.T) {

	source, err := ioutil.ReadFile("hc.png")
	if err != nil {
		t.Fatal(err)
	}

	// Nibbles will hold bytes where the 4 least significant bits
	// contain a nibble of the source data.
	nibbles := make([]byte, 2*len(source))

	fmt.Printf("Source length: %d bytes\n", len(source))

	for i, b := range source {
		first := b >> 4
		second := (b & 15)
		nibbles[i*2] = first
		nibbles[(i*2)+1] = second

		if i < 10 || len(source)-i < 10 {
			fmt.Printf("byte %05d %08b = nibbles %08b + %08b\n", i, b, first, second)
		}
	}

	output := make([]byte, len(source))
	for i := range output {
		output[i] = (nibbles[i*2] << 4) | (nibbles[(i*2)+1])
	}

	for i := range source {
		if source[i] != output[i] {
			t.Fatalf("reassembled output does not match input")
		}
	}

}
