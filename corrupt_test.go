package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"testing"
)

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
				RandomlyCorrupt(&corrupted[i])
			}
		}

		name := fmt.Sprintf("corrupt-hc-%d.png", n)
		if err := ioutil.WriteFile(name, corrupted, 0666); err != nil {
			t.Fatal(err)
		}
	}
}
