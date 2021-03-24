package main

import "fmt"

/**
 * Description: Complete the solution so that it reverses the string passed into it.
 * Example: "world"  =>  "dlrow"
 */

func reverse(word string) string {
	var reverseWord string

	for _, value := range word {
		reverseWord = string(value) + reverseWord
	}

	return reverseWord
}

func main() {
	fmt.Println(reverse("world"))
}
