package main

import (
	"fmt"
	"strings"
)

/*
	given a string of words, return the length of the shortest word(s).
	String will never be empty and you do not need to account for different data types.
*/

func findShort(s string) int {
	strArray := strings.Split(s, " ")
	shortestWord := strArray[0]
	shortestWordLen := len(shortestWord)

	for _, value := range strArray {
		if len(value) < shortestWordLen {
			shortestWord = value
			shortestWordLen = len(value)
		}
	}

	return shortestWordLen
}

func main() {
	fmt.Println(findShort("the shortest word"))
}
