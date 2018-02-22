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

	fmt.Printf("Source length: %d bytes\n", len(source))

	// Nibbles will hold bytes where the 4 least significant bits
	// contain a nibble of the source data.
	nibbles := Nibble(source)
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
