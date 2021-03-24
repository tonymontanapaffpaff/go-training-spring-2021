package main

import (
	"fmt"
	"regexp"
)

/*
	Write a simple regex to validate a username. Allowed characters are:
	lowercase letters,
	numbers,
	underscore
	Length should be between 4 and 16 characters (both included).
*/

func isUsername(username string) bool {
	pattern := "^[a-z0-9_]{4,16}$"
	matched, err := regexp.MatchString(pattern, username)

	if err != nil {
		return false
	}

	return matched
}

func main() {
	fmt.Println(isUsername("go"))
	fmt.Println(isUsername("Golang"))
	fmt.Println(isUsername("_g_o_"))
	fmt.Println(isUsername("go_1"))
	fmt.Println(isUsername("gooooooooooooooooo"))
}
