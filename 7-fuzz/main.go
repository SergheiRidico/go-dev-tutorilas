package main

import (
	"fmt"
	// "unicode/utf8"
	// "errors"
)

// Flawed version
// Function version that doesn't pass fuzzing due to:
// 1. Byte-by-bute reversing
// 2. Invalid UTF-8 strings inputs
func Reverse(s string) (string, error) {
	r := []byte(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}

// Fixed version
// func Reverse(s string) (string, error) {
// 	if !utf8.ValidString(s) {
// 		return s, errors.New("input is not valid UTF-8")
// 	}
// 	r := []rune(s)
// 	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
// 		r[i], r[j] = r[j], r[i]
// 	}
// 	return string(r), nil
// }

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, _ := Reverse(input)
	doubleRev, _ := Reverse(rev)

	fmt.Printf("Original: %q\n", input)
	fmt.Printf("Reversed: %q\n", rev)
	fmt.Printf("Reversed again: %q\n", doubleRev)
}
