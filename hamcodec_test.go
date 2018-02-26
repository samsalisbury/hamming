package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"testing"
)

func TestHamCodec(t *testing.T) {

	source, err := ioutil.ReadFile("hc.png")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Source: %d bytes\n", len(source))

	encoded := &bytes.Buffer{}
	if err := Encode(bytes.NewReader(source), encoded); err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Encoded: %d bytes\n", encoded.Len())

	corrupted := make([]byte, encoded.Len())
	copy(corrupted, encoded.Bytes())

	// Randomly corrupt 80% of the encoded bytes...
	for i := range corrupted {
		if rand.Intn(10) < 8 {
			RandomlyCorrupt(&corrupted[i])
		}
	}

	numCorrupted := 0
	encodedBytes := encoded.Bytes()

	// Check it really is corrupted..
	for i, c := range corrupted {
		b := encodedBytes[i]
		if c != b {
			numCorrupted++
		}
	}

	if numCorrupted == 0 {
		t.Fatalf("setup failed to corrupt bytes")
	}

	fmt.Printf("Corrupted: %d bytes\n", numCorrupted)

	decoded := &bytes.Buffer{}
	if err := Decode(bytes.NewReader(corrupted), decoded); err != nil {
		t.Fatal(err)
	}

	decodedBytes := decoded.Bytes()

	fmt.Printf("Decoded: %d bytes\n", decoded.Len())

	numStillCorrupted := 0
	for i, d := range decodedBytes {
		if d != source[i] {
			numStillCorrupted++
		}
	}
	if numStillCorrupted != 0 {
		t.Fatalf("Bytes still corrupted: %d", numStillCorrupted)
	}

}

// Decode Hamming coded input stream to the output stream.
func Decode(source io.Reader, dest io.Writer) error {
	in := [8]byte{}
	nibbles := [8]Nibble{}
	bs := []byte{}
	out := [4]byte{}
	for {
		n, err := source.Read(in[:])
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		for i := range in {
			syn := HamParityCheck(in[i])
			if syn != 0 {
				in[i] = HamErrCorrect(in[i], syn)
			}
			nibbles[i] = HamDecode(in[i])
		}
		bs = Nibbles(nibbles[:]).Bytes()
		for i := range bs {
			out[i] = bs[i]
		}

		o, err := dest.Write(out[:n/2])
		if err != nil {
			return err
		}
		if o != n/2 {
			return fmt.Errorf("%d out of %d bytes written", o, n/2)
		}
	}
	return nil
}

// Encode Hamming codes the input stream to the output stream.
func Encode(source io.Reader, dest io.Writer) error {
	in := [4]byte{}
	out := [8]byte{}
	for {
		n, err := source.Read(in[:])
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		nibbles := BytesToNibbles(in[:n])
		for i, b := range nibbles {
			out[i] = HamCode(b)
		}

		o, err := dest.Write(out[:n*2])
		if err != nil {
			return err
		}
		if o != 2*n {
			return fmt.Errorf("%d out of %d bytes written", o, 2*n)
		}
	}
	return nil
}
