package main

import (
	"fmt"
)

func main() {
	checkstr := "Welcome to this Project!"
	fmt.Println("Welcome to this Project!")
	fmt.Println(Reversedstr(checkstr))

	// CONCATENATE STRINGS - By Ponni Sajeevan
	str1 := "The course is WINP2000!\n"                                  //String 1
	str2 := "The contributors are: "                                     //String 2
	str3 := "Ponni, Rose, Nikhitha, Dipanshu, Piyush, Swaroop, Abhishek" //String 3
	concatoutput := str1 + str2 + str3                                   //Concat Strings and save it to the variable concatoutput
	fmt.Println(concatoutput)                                            //Print the output

	fmt.Println("Convert to Uppercase:", Uppercase(checkstr)) //Uppercase

	fmt.Println("Reversed String: ", Reversedstr(checkstr)) //Reversedstr
	fmt.Println("Palindrome: ", Palindrome(checkstr))       //Palindrome
	fmt.Println("StringLength: ", StringLength(checkstr))   //Stringlength

	fmt.Println("Replacing the Characters from String:", ReplaceWord(checkstr, "t", "AA"))

	fmt.Println("ToCamelCase: ", ToCamelCase(checkstr)) //ToCamelCase
}
