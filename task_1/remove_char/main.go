package main

import "fmt"

/*
 Description: It's pretty straightforward. Your goal is to create a function
 that removes the first and last characters of a string.
 You're given one parameter, the original string.
 You don't have to worry with strings with less than two characters.
*/

func removeChar(word string) string {
	wordByteArr := []byte(word)
	wordByteArr = append(wordByteArr[1:(len(wordByteArr) - 1)])

	return string(wordByteArr)
}

func main() {
	fmt.Println(removeChar("Call me Golang or just"))
	fmt.Println(removeChar("Go"))
}
