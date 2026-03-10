package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		input string
		want  string
	}{
		{" ", " "},
		{"Hello, world", "dlrow ,olleH"},
		{"12345!", "!54321"},
	}

	for _, tc := range testcases {
		rev, err := Reverse(tc.input)
		if err != nil {
			t.Errorf("Input is not a valid UTF-8 string: %q\n", tc.input)
		}
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q\n", rev, tc.want)
		}
	}
}

func FuzzReverse(f *testing.F) {
	testCases := []string{" ", "!12345", "Hello, world"}
	for _, tc := range testCases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
