package main

import "strings"

// Replacing the string.

func ReplaceWord(a string, b string, c string) string {
	return strings.ReplaceAll(a, string(b), string(c))
}
