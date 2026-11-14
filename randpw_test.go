package main

import (
	"strings"
	"testing"
)

func TestRandStrLengthAndCharset(t *testing.T) {
	for randType, dictionary := range dictionaries {
		for _, symbols := range []bool{false, true} {
			allowed := dictionary
			if symbols {
				allowed += symbolChars
			}
			for _, size := range []int{1, 16, 128} {
				s, err := randStr(size, randType, symbols)
				if err != nil {
					t.Fatalf("randStr(%d, %q, %v) error: %v", size, randType, symbols, err)
				}
				if len(s) != size {
					t.Errorf("randStr(%d, %q, %v) length = %d", size, randType, symbols, len(s))
				}
				for _, c := range s {
					if !strings.ContainsRune(allowed, c) {
						t.Errorf("randStr(%d, %q, %v) contains %q, not in dictionary", size, randType, symbols, c)
					}
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
			if _, err := randStr(tt.size, tt.randType, false); err == nil {
				t.Errorf("randStr(%d, %q, false) expected error", tt.size, tt.randType)
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
		for _, symbols := range []bool{false, true} {
			allowed := dictionary
			if symbols {
				allowed += symbolChars
			}
			name := randType
			if symbols {
				name += "+symbols"
			}
			t.Run(name, func(t *testing.T) {
				s, err := randStr(len(allowed)*perChar, randType, symbols)
				if err != nil {
					t.Fatal(err)
				}
				counts := make(map[rune]int)
				for _, c := range s {
					counts[c]++
				}
				for _, c := range allowed {
					if got := counts[c]; got < perChar*9/10 || got > perChar*11/10 {
						t.Errorf("%q appeared %d times, want %d ±10%%", c, got, perChar)
					}
				}
			})
		}
	}
}
