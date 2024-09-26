package main

import (
	"fmt"
)

func main() {
	checkstr := "Welcome to Group A's Week 4 Project!"
	fmt.Println("Welcome to Group A's Week 4 Project!")
	fmt.Println("Reversed String: ", Reversedstr(checkstr))
	fmt.Println("Palindrome: ", Palindrome(checkstr))

	fmt.Println("Replacing the Word from String:", ReplaceWord(checkstr, "t", "AA"))
}
