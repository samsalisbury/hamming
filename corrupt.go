package main

import "math/rand"

func RandomlyCorrupt(b *byte) {
	*b = *b + (1 << (rand.Uint64() % 8))
}
