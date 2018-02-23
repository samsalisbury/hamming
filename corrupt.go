package main

import "math/rand"

// RandomlyCorrupt randomly flips a bit in the byte.
func RandomlyCorrupt(b *byte) {
	*b = *b + (1 << (rand.Uint64() % 8))
}
