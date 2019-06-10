// Based on example found at:
// https://www.socketloop.com/tutorials/golang-how-to-generate-random-string

package main

import (
	"crypto/rand"
	"flag"
	"fmt"
)

func randStr(strSize int, randType string) string {

	var dictionary string

	if randType == "alphanum" {
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "alpha" {
		dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "number" {
		dictionary = "0123456789"
	}

	var bytes = make([]byte, strSize)
	_, err := rand.Read(bytes)
	checkErr(err)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

func main() {
	length := flag.Int("length", 16, "password length")
	flag.Parse()

	fmt.Println("Alphanum : ", randStr(*length, "alphanum"))

	fmt.Println("Alpha : ", randStr(*length, "alpha"))

	fmt.Println("Numbers : ", randStr(*length, "number"))

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
