package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := longest("mandar", "anand", "joshi")
	fmt.Println(s)
	magicOfScopes()
}

func longest(s ...string) (result string) {
	for _, str := range s {
		if len(str) > len(result) {
			result = str
		}
	}
	return result
}

func magicOfScopes() {
	s := "abcde"
	for _, s := range s {
		s := unicode.ToUpper(s)
		fmt.Print(string(s))
	}
	fmt.Println("\n" + s)
}
