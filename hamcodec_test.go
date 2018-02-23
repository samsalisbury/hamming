package main

import (
	"fmt"
	"io"
	"testing"
)

func TestHamCodec(t *testing.T) {

}

// Encode Hamming codes the input stream to the output stream.
func Encode(source io.Reader, dest io.Writer) error {
	in := [4]byte{}
	//nibbles := [8]Nibble{}
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
		if o != n {
			return fmt.Errorf("%d out of %d bytes written", o, n)
		}
	}
	return nil
}
