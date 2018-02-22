package main

import "testing"

func TestPopCount(t *testing.T) {

	testCases := []struct {
		in, want Nibble
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 1},
		{5, 2},
		{6, 2},
		{7, 3},
		{8, 1},
	}

	for _, tc := range testCases {
		got := PopCount(tc.in)
		if got != tc.want {
			t.Errorf("PopCount(%08b) == %d; want %d", tc.in, got, tc.want)
		}
	}
}
