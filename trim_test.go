package main

import (
	"testing"
	"bytes"
	"unicode"
	"strings"
)

// Strings with various amounts of spaces
var trimSpaceStrings = map[string]string{
	"none": "this string is not trimmed",
	"one": " this string has one space",
	"many": "       this string has lots of spaces",
	"reallybig": strings.Repeat(" ", 1000) + "has a really big number of spaces",
}

// Strings with various amounts of different whitespace
var trimWhitespaceStrings = map[string]string{
	"none": "this string is not trimmed",
	"onetab": "\tthis string has a tab prefix",
	"onespace": " this string has many spaces",
	"variety": "\t \n \f    \t  \r  this string has lots of whitespace",
	"reallybig": strings.Repeat("\t \n", 500) + "has a really big amount of whitespace",
}

// Extracted from unicode.IsSpace
const whitespace = "\t\n\v\f\r \x85\xA0"

// Tests performance of bytes.TrimLeft
func BenchmarkBytesTrimLeftSpace(b *testing.B) {
	for name, val := range trimSpaceStrings {
		raw := []byte(val)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n += 1 {
				_ = bytes.TrimLeft(raw, whitespace)
			}
		})
	}
}

// Tests performance of bytes.TrimFunc + unicode.IsSpace
func BenchmarkBytesTrimLeftFuncSpace(b *testing.B) {
	for name, val := range trimSpaceStrings {
		raw := []byte(val)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n += 1 {
				_ = bytes.TrimLeftFunc(raw, unicode.IsSpace)
			}
		})
	}
}

// Tests performance of strings.TrimLeft
func BenchmarkStringsTrimLeftSpace(b *testing.B) {
	for name, val := range trimSpaceStrings {
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n += 1 {
				_ = strings.TrimLeft(val, whitespace)
			}
		})
	}
}

// Tests performance of strings + unicode.IsSpace
func BenchmarkStringsTrimLeftFuncSpace(b *testing.B) {
	for name, val := range trimSpaceStrings {
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n += 1 {
				_ = strings.TrimLeftFunc(val, unicode.IsSpace)
			}
		})
	}
}