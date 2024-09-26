package main

import (
	"fmt"
)

func main() {
	checkstr := "Welcome to Group A's Week 4 Project!"
	fmt.Println("Welcome to Group A's Week 4 Project!")
	fmt.Println(Reversedstr(checkstr))

	//CONCATENATE STRINGS - By Ponni Sajeevan
	fmt.Println("\nCONCATENATE STRINGS - By Ponni Sajeevan")
	str1 := "'The course is WINP2000!\n"                                  //String 1
	str2 := "The contributors are: "                                      //String 2
	str3 := "Ponni, Rose, Nikhitha, Dipanshu, Piyush, Swaroop, Abhishek'" //String 3
	concatoutput := str1 + str2 + str3                                    //Concat Strings and save it to the variable concatoutput
	fmt.Println(concatoutput)                                             //Print the output

	fmt.Println("Reversed String: ", Reversedstr(checkstr))
	fmt.Println("Palindrome: ", Palindrome(checkstr))

	fmt.Println("Replacing the Word from String:", ReplaceWord(checkstr, "t", "AA"))

	// Convert to Uppercase - Nikhitha
	input := "winp group a!"
	output := Uppercase(input)
	fmt.Println("Convert to Uppercase:", output)

	// Returns the length of the input string.
	fmt.Println(StringLength(checkstr))

}
