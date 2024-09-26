// ---------------------------------------------------------------------------------------
// Function Name: ToCamelCase
// ---------------------------------------------------------------------------------------                                                                                                       */
// Description: This function will create the input words in Camel Case format. For
// example, the input text is "How are you", the output will be "howAreYou"
// Author: Rose Sauri
// ---------------------------------------------------------------------------------------

package main

import "strings"

func ToCamelCase(sInput string) string {

	//Split the string by words
	oWords := strings.Split(sInput, " ")
	oOutput := strings.Split(sInput, " ")

	//Loop thru each word
	for i, sWord := range oWords {
		if i == 0 {
			//For the first word, make the word all lower case
			oOutput[i] = strings.ToLower(sWord)
		} else {
			//For the succeeding words, make the words in proper case
			//Split each word in individual chars and make the first char to Upper
			oChars := strings.Split(strings.ToLower(sWord), "")
			oChars[0] = strings.ToUpper(oChars[0])
			oOutput[i] = strings.Join(oChars, "")
		}

	}
	return strings.Join(oOutput, "")
}
