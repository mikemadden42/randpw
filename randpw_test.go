package main

import (
	"strings"
	"testing"
)

func TestRandStrLengthAndCharset(t *testing.T) {
	for randType, dictionary := range dictionaries {
		for _, size := range []int{1, 16, 128} {
			s, err := randStr(size, randType)
			if err != nil {
				t.Fatalf("randStr(%d, %q) error: %v", size, randType, err)
			}
			if len(s) != size {
				t.Errorf("randStr(%d, %q) length = %d", size, randType, len(s))
			}
			for _, c := range s {
				if !strings.ContainsRune(dictionary, c) {
					t.Errorf("randStr(%d, %q) contains %q, not in dictionary", size, randType, c)
				}
			}
		}
	}
}

func TestRandStrErrors(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		randType string
	}{
		{"zero length", 0, "alphanum"},
		{"negative length", -1, "alphanum"},
		{"unknown type", 16, "symbols"},
		{"empty type", 16, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := randStr(tt.size, tt.randType); err == nil {
				t.Errorf("randStr(%d, %q) expected error", tt.size, tt.randType)
			}
		})
	}
}

// TestRandStrDistribution checks that every character is roughly equally
// likely. The old modulo approach favored some alphanum and alpha characters
// by 25%; the 10% tolerance catches that while staying well clear of random
// noise. The old digit bias (26/256 vs 25/256) is too small to detect here.
func TestRandStrDistribution(t *testing.T) {
	const perChar = 2000
	for randType, dictionary := range dictionaries {
		t.Run(randType, func(t *testing.T) {
			s, err := randStr(len(dictionary)*perChar, randType)
			if err != nil {
				t.Fatal(err)
			}
			counts := make(map[rune]int)
			for _, c := range s {
				counts[c]++
			}
			for _, c := range dictionary {
				if got := counts[c]; got < perChar*9/10 || got > perChar*11/10 {
					t.Errorf("%q appeared %d times, want %d ±10%%", c, got, perChar)
				}
			}
		})
	}
}
