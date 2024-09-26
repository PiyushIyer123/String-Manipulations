package main

func Palindrome(i string) bool {
	reversed := Reversedstr(i)
	return i == reversed
}
