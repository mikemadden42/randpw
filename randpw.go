// Command randpw prints a random password.
//
// Based on example found at:
// https://www.socketloop.com/tutorials/golang-how-to-generate-random-string
package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
)

var dictionaries = map[string]string{
	"alphanum": "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
	"alpha":    "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
	"number":   "0123456789",
}

// symbolChars avoids quotes, backslashes, backticks and spaces so passwords
// can be pasted into a shell or config file without escaping.
const symbolChars = "!#$%&*+-=?@^_"

func randStr(strSize int, randType string, symbols bool) (string, error) {
	dictionary, ok := dictionaries[randType]
	if !ok {
		return "", fmt.Errorf("unknown type %q", randType)
	}
	if strSize < 1 {
		return "", fmt.Errorf("length must be at least 1, got %d", strSize)
	}
	if symbols {
		dictionary += symbolChars
	}

	// rand.Int picks uniformly from [0, n), avoiding the modulo bias of
	// reducing a random byte by the dictionary size.
	n := big.NewInt(int64(len(dictionary)))
	bytes := make([]byte, strSize)
	for k := range bytes {
		i, err := rand.Int(rand.Reader, n)
		if err != nil {
			return "", err
		}
		bytes[k] = dictionary[i.Int64()]
	}
	return string(bytes), nil
}

func main() {
	length := flag.Int("length", 16, "password length")
	randType := flag.String("type", "alphanum", "character set: alphanum, alpha or number")
	symbols := flag.Bool("symbols", false, "also use the symbols "+symbolChars)
	all := flag.Bool("all", false, "print one labeled password of each type")
	flag.Parse()

	if *length < 1 {
		fmt.Fprintf(os.Stderr, "randpw: -length must be at least 1, got %d\n", *length)
		os.Exit(2)
	}
	if _, ok := dictionaries[*randType]; !ok {
		fmt.Fprintf(os.Stderr, "randpw: -type must be alphanum, alpha or number, got %q\n", *randType)
		os.Exit(2)
	}

	if !*all {
		s, err := randStr(*length, *randType, *symbols)
		checkErr(err)
		fmt.Println(s)
		return
	}

	for _, t := range []struct{ label, randType string }{
		{"Alphanum", "alphanum"},
		{"Alpha", "alpha"},
		{"Numbers", "number"},
	} {
		s, err := randStr(*length, t.randType, *symbols)
		checkErr(err)
		fmt.Printf("%s: %s\n", t.label, s)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "randpw:", err)
		os.Exit(1)
	}
}
