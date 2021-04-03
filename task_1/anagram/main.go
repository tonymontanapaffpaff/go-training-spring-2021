package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
  Description: An anagram is the result of rearranging the letters of a Word to produce a new Word (see wikipedia https://en.wikipedia.org/wiki/Anagram).
  Note: anagrams are case insensitive

  Complete the function to return true if the two arguments given are anagrams of each other; return false otherwise.

  Examples:
  "foefet" is an anagram of "toffee"
  "Buckethead" is an anagram of "DeathCubeK"
*/

func isAnagram(test, original string) bool {
	sortStr := func(s string) string {
		strArr := strings.Split(strings.ToLower(s), "")
		sort.Strings(strArr)
		return strings.Join(strArr, "")
	}

	sortedTest := sortStr(test)
	sortedOriginal := sortStr(original)

	return sortedTest == sortedOriginal
}

func main() {
	fmt.Println(isAnagram("foefet", "toffee"))
	fmt.Println(isAnagram("Buckethead", "DeathCubeK"))
}
